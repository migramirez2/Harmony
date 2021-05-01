package main

import (
	"errors"
	"fmt"

	"github.com/harmony-one/harmony/api/service/legacysync"
	nodeconfig "github.com/harmony-one/harmony/internal/configs/node"
	goversion "github.com/hashicorp/go-version"
	"github.com/pelletier/go-toml"
)

const legacyConfigVersion = "1.0.4"

func doMigrations(confVersion string, confTree *toml.Tree) error {
	Ver, err := goversion.NewVersion(confVersion)
	if err != nil {
		return fmt.Errorf("invalid or missing config file version - '%s'", confVersion)
	}
	legacyVer, _ := goversion.NewVersion(legacyConfigVersion)
	migrationKey := confVersion
	if Ver.LessThan(legacyVer) {
		migrationKey = legacyConfigVersion
	}

	migration, found := migrations[migrationKey]

	// Version does not match any of the migration criteria
	if !found {
		return fmt.Errorf("unrecognized config version - %s", confVersion)
	}

	for confVersion != tomlConfigVersion {
		confTree = migration(confTree)
		confVersion = confTree.Get("Version").(string)
		migration = migrations[confVersion]
	}
	return nil
}

func migrateConf(confBytes []byte) (harmonyConfig, string, error) {
	var (
		migratedFrom string
	)
	confTree, err := toml.LoadBytes(confBytes)
	if err != nil {
		return harmonyConfig{}, "", fmt.Errorf("config file parse error - %s", err.Error())
	}
	confVersion, found := confTree.Get("Version").(string)
	if !found {
		return harmonyConfig{}, "", errors.New("config file invalid - no version entry found")
	}
	migratedFrom = confVersion
	if confVersion != tomlConfigVersion {
		err = doMigrations(confVersion, confTree)
		if err != nil {
			return harmonyConfig{}, "", err
		}
	}

	// At this point we must be at current config version so
	// we can safely unmarshal it
	var config harmonyConfig
	if err := confTree.Unmarshal(&config); err != nil {
		return harmonyConfig{}, "", err
	}
	return config, migratedFrom, nil
}

var (
	migrations = make(map[string]configMigrationFunc)
)

type configMigrationFunc func(*toml.Tree) *toml.Tree

func init() {
	migrations["1.0.4"] = func(confTree *toml.Tree) *toml.Tree {
		ntStr := confTree.Get("Network.NetworkType").(string)
		nt := parseNetworkType(ntStr)

		defDNSSyncConf := getDefaultDNSSyncConfig(nt)

		zoneField := confTree.Get("Network.DNSZone")
		if zone, ok := zoneField.(string); ok {
			confTree.Set("DNSSync.Zone", zone)
		}

		var port = int64(nodeconfig.DefaultDNSPort)
		portField := confTree.Get("Network.DNSPort")
		if p, ok := portField.(int64); ok {
			if p != nodeconfig.DefaultDNSPort {
				port = p - legacysync.SyncingPortDifference
			}
			confTree.Set("DNSSync.Port", port)
		}

		syncingField := confTree.Get("Network.LegacySyncing")
		if syncing, ok := syncingField.(bool); ok {
			confTree.Set("DNSSync.LegacySyncing", syncing)
		}

		clientField := confTree.Get("Sync.LegacyClient")
		if client, ok := clientField.(bool); ok {
			confTree.Set("DNSSync.Client", client)
		} else {
			confTree.Set("DNSSync.Client", defDNSSyncConf.Client)
		}

		serverField := confTree.Get("Sync.LegacyServer")
		if server, ok := serverField.(bool); ok {
			confTree.Set("DNSSync.Server", server)
		} else {
			confTree.Set("DNSSync.Server", defDNSSyncConf.Client)
		}

		serverPort := defDNSSyncConf.ServerPort
		serverPortField := confTree.Get("Sync.LegacyServerPort")
		if port, ok := serverPortField.(int64); ok {
			serverPort = int(port)
		}
		confTree.Set("DNSSync.ServerPort", serverPort)

		rosettaPort := confTree.Get("HTTP.RosettaPort")
		if rosettaPort == nil {
			confTree.Set("HTTP.RosettaPort", defaultConfig.HTTP.RosettaPort)
		}

		confTree.Set("Version", "2.0.0")
		return confTree
	}
}

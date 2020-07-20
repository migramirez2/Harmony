package main

import (
	"github.com/harmony-one/harmony/internal/cli"
	nodeconfig "github.com/harmony-one/harmony/internal/configs/node"
	"github.com/spf13/cobra"
)

var networkFlags = []cli.Flag{
	networkTypeFlag,
	bootNodeFlag,
	dnsZoneFlag,
	dnsPortFlag,
	legacyDNSZoneFlag,
	legacyDNSPortFlag,
	legacyDNSFlag,
	legacyNetworkTypeFlag,
}

var (
	networkTypeFlag = cli.StringFlag{
		Name:      "network",
		Shorthand: "n",
		DefValue:  "mainnet",
		Usage:     "network to join (mainnet, testnet, pangaea, localnet, partner, stressnet, devnet)",
	}
	bootNodeFlag = cli.StringSliceFlag{
		Name:  "bootnodes",
		Usage: "a list of bootnode multiaddress (delimited by ,)",
	}
	dnsZoneFlag = cli.StringFlag{
		Name:  "dns.zone",
		Usage: "use customized peers from the zone for state syncing",
	}
	// TODO: 9500 as default
	dnsPortFlag = cli.IntFlag{
		Name:     "dns.port",
		DefValue: nodeconfig.DefaultDNSPort,
		Usage:    "port of customized dns node",
	}
	legacyDNSZoneFlag = cli.StringFlag{
		Name:       "dns_zone",
		Usage:      "use peers from the zone for state syncing",
		Deprecated: "use --dns.zone",
	}
	legacyDNSPortFlag = cli.IntFlag{
		Name:       "dns_port",
		Usage:      "port of dns node",
		Deprecated: "use --dns.zone",
	}
	legacyDNSFlag = cli.BoolFlag{
		Name:       "dns",
		DefValue:   true,
		Usage:      "use dns for syncing",
		Deprecated: "set to false only to use self discovery peers for syncing",
	}
	legacyNetworkTypeFlag = cli.StringFlag{
		Name:       "network_type",
		Usage:      "network to join (mainnet, testnet, pangaea, localnet, partner, stressnet, devnet)",
		Deprecated: "use --network",
	}
)

func getNetworkType(cmd *cobra.Command) nodeconfig.NetworkType {
	var raw string
	if cmd.Flags().Changed(legacyNetworkTypeFlag.Name) {
		raw = cli.GetStringFlagValue(cmd, legacyNetworkTypeFlag)
	} else {
		raw = cli.GetStringFlagValue(cmd, networkTypeFlag)
	}
	nt := parseNetworkType(raw)
	return nt
}

func applyNetworkFlags(cmd *cobra.Command, cfg *hmyConfig) {
	fs := cmd.Flags()

	if fs.Changed(bootNodeFlag.Name) {
		cfg.Network.BootNodes = cli.GetStringSliceFlagValue(cmd, bootNodeFlag)
	}

	if fs.Changed(dnsZoneFlag.Name) {
		cfg.Network.DNSZone = cli.GetStringFlagValue(cmd, dnsZoneFlag)
	} else if fs.Changed(legacyDNSZoneFlag.Name) {
		cfg.Network.DNSZone = cli.GetStringFlagValue(cmd, legacyDNSZoneFlag)
	} else if fs.Changed(legacyDNSFlag.Name) {
		val := cli.GetBoolFlagValue(cmd, legacyDNSFlag)
		if val {
			cfg.Network.DNSZone = mainnetDnsZone
		} else {
			cfg.Network.LegacySyncing = true
		}
	}

	if fs.Changed(dnsPortFlag.Name) {
		cfg.Network.DNSPort = cli.GetIntFlagValue(cmd, dnsPortFlag)
	} else if fs.Changed(legacyDNSPortFlag.Name) {
		cfg.Network.DNSPort = cli.GetIntFlagValue(cmd, legacyDNSPortFlag)
	}
}

func parseNetworkType(nt string) nodeconfig.NetworkType {
	switch nt {
	case "mainnet":
		return nodeconfig.Mainnet
	case "testnet":
		return nodeconfig.Testnet
	case "pangaea", "staking", "stk":
		return nodeconfig.Pangaea
	case "partner":
		return nodeconfig.Partner
	case "stressnet", "stress", "stn":
		return nodeconfig.Stressnet
	case "localnet":
		return nodeconfig.Localnet
	case "devnet", "dev":
		return nodeconfig.Devnet
	default:
		return ""
	}
}

func getDefaultNetworkConfig(nt nodeconfig.NetworkType) networkConfig {
	bn := nodeconfig.GetDefaultBootNodes(nt)
	zone := nodeconfig.GetDefaultDNSZone(nt)
	port := nodeconfig.GetDefaultDNSPort(nt)
	return networkConfig{
		NetworkType: string(nt),
		BootNodes:   bn,
		DNSZone:     zone,
		DNSPort:     port,
	}
}

var p2pFlags = []cli.Flag{
	p2pPortFlag,
	p2pKeyFileFlag,
	legacyKeyFileFlag,
}

var (
	p2pPortFlag = cli.IntFlag{
		Name:     "p2p.port",
		Usage:    "port to listen for p2p communication",
		DefValue: defaultConfig.P2P.Port,
	}
	p2pKeyFileFlag = cli.StringFlag{
		Name:     "p2p.keyfile",
		Usage:    "the p2p key file of the harmony node",
		DefValue: defaultConfig.P2P.KeyFile,
	}
	legacyKeyFileFlag = cli.StringFlag{
		Name:       "key",
		Usage:      "the p2p key file of the harmony node",
		DefValue:   defaultConfig.P2P.KeyFile,
		Deprecated: "use --p2p.keyfile",
	}
)

func applyP2PFlags(cmd *cobra.Command, config *hmyConfig) {
	fs := cmd.Flags()

	if fs.Changed(p2pPortFlag.Name) {
		config.P2P.Port = cli.GetIntFlagValue(cmd, p2pPortFlag)
	}

	if fs.Changed(p2pKeyFileFlag.Name) {
		config.P2P.KeyFile = cli.GetStringFlagValue(cmd, p2pKeyFileFlag)
	} else if fs.Changed(legacyKeyFileFlag.Name) {
		config.P2P.KeyFile = cli.GetStringFlagValue(cmd, legacyKeyFileFlag)
	}
}

var rpcFlags = []cli.Flag{
	rpcEnabledFlag,
	rpcIPFlag,
	rpcPortFlag,
	legacyRPCIPFlag,
	legacyPublicRPCFlag,
}

var (
	rpcEnabledFlag = cli.BoolFlag{
		Name:     "http",
		Usage:    "enable HTTP / RPC requests",
		DefValue: defaultConfig.RPC.Enabled,
	}
	rpcIPFlag = cli.StringFlag{
		Name:     "http.ip",
		Usage:    "ip address to listen for RPC calls",
		DefValue: defaultConfig.RPC.IP,
	}
	rpcPortFlag = cli.IntFlag{
		Name:     "http.port",
		Usage:    "rpc port to listen for RPC calls",
		DefValue: defaultConfig.RPC.Port,
	}
	legacyRPCIPFlag = cli.StringFlag{
		Name:       "ip",
		Usage:      "ip of the node",
		DefValue:   defaultConfig.RPC.IP,
		Deprecated: "use --http.ip",
	}
	legacyPublicRPCFlag = cli.BoolFlag{
		Name:       "public_rpc",
		Usage:      "Enable Public RPC Access (default: false)",
		DefValue:   defaultConfig.RPC.Enabled,
		Deprecated: "please use --http.ip to specify the ip address to listen",
	}
)

func applyRPCFlags(cmd *cobra.Command, config *hmyConfig) {
	fs := cmd.Flags()

	var isRPCSpecified bool

	if fs.Changed(rpcIPFlag.Name) {
		config.RPC.IP = cli.GetStringFlagValue(cmd, rpcIPFlag)
		isRPCSpecified = true
	} else if fs.Changed(legacyRPCIPFlag.Name) {
		config.RPC.IP = cli.GetStringFlagValue(cmd, legacyRPCIPFlag)
		isRPCSpecified = true
	}

	if fs.Changed(rpcPortFlag.Name) {
		config.RPC.Port = cli.GetIntFlagValue(cmd, rpcPortFlag)
		isRPCSpecified = true
	}

	if fs.Changed(rpcEnabledFlag.Name) {
		config.RPC.Enabled = cli.GetBoolFlagValue(cmd, rpcEnabledFlag)
	} else if isRPCSpecified {
		config.RPC.Enabled = true
	}
}

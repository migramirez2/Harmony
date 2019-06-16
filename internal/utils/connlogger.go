package utils

import (
	"github.com/ethereum/go-ethereum/log"
	net "github.com/libp2p/go-libp2p-net"
	ma "github.com/multiformats/go-multiaddr"
)

// ConnLogger is a LibP2P connection logger that logs to an Ethereum logger.
// It logs all listener/connection/stream open/close activities at debug level.
// To use one, add it on a LibP2P host swarm as a notifier, ex:
//
//   connLogger := utils.NewConnLogger(
//   host.Network().Notify(connLogger)
type ConnLogger struct {
	l log.Logger
}

func (cl ConnLogger) Listen(net net.Network, ma ma.Multiaddr) {
	cl.l.Debug("[CONNECTIONS] Listener starting", "net", net, "addr", ma)
}

func (cl ConnLogger) ListenClose(net net.Network, ma ma.Multiaddr) {
	cl.l.Debug("[CONNECTIONS] Listener closing", "net", net, "addr", ma)
}

func (cl ConnLogger) Connected(net net.Network, conn net.Conn) {
	cl.l.Debug("[CONNECTIONS] Connected", "net", net,
		"localPeer", conn.LocalPeer(), "localAddr", conn.LocalMultiaddr(),
		"remotePeer", conn.RemotePeer(), "remoteAddr", conn.RemoteMultiaddr(),
	)
}

func (cl ConnLogger) Disconnected(net net.Network, conn net.Conn) {
	cl.l.Debug("[CONNECTIONS] Disconnected", "net", net,
		"localPeer", conn.LocalPeer(), "localAddr", conn.LocalMultiaddr(),
		"remotePeer", conn.RemotePeer(), "remoteAddr", conn.RemoteMultiaddr(),
	)
}

func (cl ConnLogger) OpenedStream(net net.Network, stream net.Stream) {
	conn := stream.Conn()
	cl.l.Debug("[CONNECTIONS] Stream opened", "net", net,
		"localPeer", conn.LocalPeer(), "localAddr", conn.LocalMultiaddr(),
		"remotePeer", conn.RemotePeer(), "remoteAddr", conn.RemoteMultiaddr(),
		"protocol", stream.Protocol(),
	)
}

func (cl ConnLogger) ClosedStream(net net.Network, stream net.Stream) {
	conn := stream.Conn()
	cl.l.Debug("[CONNECTIONS] Stream closed", "net", net,
		"localPeer", conn.LocalPeer(), "localAddr", conn.LocalMultiaddr(),
		"remotePeer", conn.RemotePeer(), "remoteAddr", conn.RemoteMultiaddr(),
		"protocol", stream.Protocol(),
	)
}

// NewConnLogger returns a new connection logger that uses the given
// Ethereum logger.  See ConnLogger for usage.
func NewConnLogger(l log.Logger) *ConnLogger {
	return &ConnLogger{l: l}
}

// RootConnLogger is a LibP2P connection logger that logs to Ethereum root
// logger.  See ConnLogger for usage.
var RootConnLogger = NewConnLogger(log.Root())

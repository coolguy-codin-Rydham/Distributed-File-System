package p2p

import "net"

// RPC Holds any arbitrary Message or Data that is
// being sent over each transport between two nodes
type RPC struct {
	From    net.Addr
	Payload []byte
}

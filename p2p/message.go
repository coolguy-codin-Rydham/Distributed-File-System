package p2p

import "net"

// Message Holds any arbitrary Message or Data that is
// being sent over each transport between two nodes
type Message struct {
	From    net.Addr
	Payload []byte
}

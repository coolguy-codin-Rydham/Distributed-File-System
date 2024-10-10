package p2p

// Peer is an interface that represents the remote node
type Peer interface {
}

// Its anything that handles the communication
// between nodes in the network
// form (TCP, UDP, websockets......)
type Transport interface {
	ListenAndAccept() error
}

package p2p

// Handshakefunc....?
type HandShakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}

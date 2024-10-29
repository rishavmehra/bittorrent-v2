package handshake

// <pstrlen><pstr><reserved><info_hash><peer_id>
type Handshake struct {
	Length   byte
	Protocol [19]byte
	Reserved [8]byte
	InfoHash [20]byte
	PeerID   [20]byte
}

func NewHandshake(peerID [20]byte, infoHash [20]byte) *Handshake {
	msg := &Handshake{
		Length:   19,
		InfoHash: infoHash,
		PeerID:   peerID,
	}
	copy(msg.Protocol[:], "BitTorrent protocol")
	return msg
}

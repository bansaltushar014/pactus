package peerset

import (
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/sync/bundle/message"
	"github.com/pactus-project/pactus/sync/service"
)

type Peer struct {
	Status          StatusCode
	Moniker         string
	Agent           string
	Address         string
	PeerID          peer.ID
	ConsensusKeys   []*bls.PublicKey
	Services        service.Services
	LastSent        time.Time
	LastReceived    time.Time
	LastBlockHash   hash.Hash
	Height          uint32
	ReceivedBundles int
	InvalidBundles  int
	ReceivedBytes   map[message.Type]int64
	SentBytes       map[message.Type]int64
}

func NewPeer(peerID peer.ID) *Peer {
	return &Peer{
		ConsensusKeys: make([]*bls.PublicKey, 0),
		Status:        StatusCodeUnknown,
		PeerID:        peerID,
		ReceivedBytes: make(map[message.Type]int64),
		SentBytes:     make(map[message.Type]int64),
	}
}

func (p *Peer) IsKnownOrTrusty() bool {
	return p.Status == StatusCodeKnown || p.Status == StatusCodeTrusty
}

func (p *Peer) IsBanned() bool {
	return p.Status == StatusCodeBanned
}

func (p *Peer) HasNetworkService() bool {
	return p.Services.IsNetwork()
}

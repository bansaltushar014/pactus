package peerset

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pactus-project/pactus/sync/bundle/message"
	"github.com/pactus-project/pactus/util"
)

type Session struct {
	lk   sync.RWMutex
	data sessionData
}

type sessionData struct {
	SessionID        int
	PeerID           peer.ID
	LastResponseCode message.ResponseCode
	LastActivityAt   time.Time
}

func newSession(id int, peerID peer.ID) *Session {
	return &Session{
		data: sessionData{
			SessionID:      id,
			PeerID:         peerID,
			LastActivityAt: util.Now(),
		},
	}
}

func (s *Session) SetLastResponseCode(code message.ResponseCode) {
	s.lk.Lock()
	defer s.lk.Unlock()

	s.data.LastResponseCode = code
	s.data.LastActivityAt = util.Now()
}

func (s *Session) PeerID() peer.ID {
	s.lk.RLock()
	defer s.lk.RUnlock()

	return s.data.PeerID
}

func (s *Session) SessionID() int {
	s.lk.RLock()
	defer s.lk.RUnlock()

	return s.data.SessionID
}

func (s *Session) LastActivityAt() time.Time {
	s.lk.RLock()
	defer s.lk.RUnlock()

	return s.data.LastActivityAt
}

package mapping

import (
	"github.com/gorilla/websocket"
	"sync"
)

func NewSidConnMapping() *SidConnMapping {
	return &SidConnMapping{relations: make(map[uint64]*websocket.Conn)}
}

type SidConnMapping struct {
	relations map[uint64]*websocket.Conn

	mu sync.RWMutex
}

func (m *SidConnMapping) Add(sid uint64, conn *websocket.Conn) {
	m.mu.Lock()
	m.relations[sid] = conn
	m.mu.Unlock()
}

func (m *SidConnMapping) Delete(sid uint64) {
	m.mu.Lock()
	delete(m.relations, sid)
	m.mu.Unlock()
}

func (m *SidConnMapping) Get(sid uint64) (*websocket.Conn, bool) {
	m.mu.RLock()
	conn, has := m.relations[sid]
	m.mu.RUnlock()
	return conn, has
}

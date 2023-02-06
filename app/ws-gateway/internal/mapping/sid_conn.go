package mapping

import (
	"github.com/gorilla/websocket"
	"sync"
)

func NewSidConnMapping() *SidConnMapping {
	return &SidConnMapping{relations: make(map[string]*websocket.Conn)}
}

type SidConnMapping struct {
	relations map[string]*websocket.Conn

	mu sync.RWMutex
}

func (m *SidConnMapping) Add(sid string, conn *websocket.Conn) {
	m.mu.Lock()
	m.relations[sid] = conn
	m.mu.Unlock()
}

func (m *SidConnMapping) Delete(sid string) {
	m.mu.Lock()
	delete(m.relations, sid)
	m.mu.Unlock()
}

func (m *SidConnMapping) Get(sid string) (*websocket.Conn, bool) {
	m.mu.RLock()
	conn, has := m.relations[sid]
	m.mu.RUnlock()
	return conn, has
}

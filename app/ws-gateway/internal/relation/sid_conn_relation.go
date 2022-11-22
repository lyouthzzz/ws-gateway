package relation

import (
	"github.com/gorilla/websocket"
	"sync"
)

func NewSidConnRelation() *SidConnRelation {
	return &SidConnRelation{relations: make(map[uint64]*websocket.Conn)}
}

type SidConnRelation struct {
	relations map[uint64]*websocket.Conn

	mu sync.RWMutex
}

func (relation *SidConnRelation) Add(sid uint64, conn *websocket.Conn) {
	relation.mu.Lock()
	relation.relations[sid] = conn
	relation.mu.Unlock()
}

func (relation *SidConnRelation) Delete(sid uint64) {
	relation.mu.Lock()
	delete(relation.relations, sid)
	relation.mu.Unlock()
}

func (relation *SidConnRelation) Get(sid uint64) (*websocket.Conn, bool) {
	relation.mu.RLock()
	conn, has := relation.relations[sid]
	relation.mu.RUnlock()
	return conn, has
}

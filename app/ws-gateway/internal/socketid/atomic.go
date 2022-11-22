package socketid

import "sync/atomic"

var _ Generator = (*atomicGenerator)(nil)

func NewAtomicGenerator(initV uint64) Generator {
	return &atomicGenerator{v: initV}
}

type atomicGenerator struct {
	v uint64
}

func (g *atomicGenerator) NextSid() (uint64, error) {
	return atomic.AddUint64(&g.v, 1), nil
}

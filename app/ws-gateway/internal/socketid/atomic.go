package socketid

import (
	"strconv"
	"sync/atomic"
)

var _ Generator = (*atomicGenerator)(nil)

func NewAtomicGenerator(initV uint64) Generator {
	return &atomicGenerator{v: initV}
}

type atomicGenerator struct {
	v uint64
}

func (g *atomicGenerator) NextSid() (string, error) {
	v := atomic.AddUint64(&g.v, 1)
	return strconv.FormatUint(v, 10), nil
}

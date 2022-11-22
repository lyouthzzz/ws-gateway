package socketid

type Generator interface {
	NextSid() (uint64, error)
}

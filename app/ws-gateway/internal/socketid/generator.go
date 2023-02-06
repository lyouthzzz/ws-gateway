package socketid

type Generator interface {
	NextSid() (string, error)
}

package request

type State int64

const (
	PENDING State = iota
	COMPLETE
)

type Request[T any] struct {
	id      string
	state   State
	payload T
}

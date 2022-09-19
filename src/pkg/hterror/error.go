package hterror

type Errors uint64

const (
	ErrNotFound Errors = 1 << iota
	ErrorConnection
)

func (e Errors) Error() string {
	switch e {
	case ErrNotFound:
		return "record not found"
	case ErrorConnection:
		return "internal server error"
	default:
		return "unknown error"
	}
}
func (e Errors) String() string {
	return e.Error()
}

package hterror

type Errors uint64

const (
	ErrNotFound Errors = iota
	ErrConnection
	ErrInternal
)

func (e Errors) Error() string {
	switch e {
	case ErrNotFound:
		return "record not found"
	case ErrConnection:
		return "internal server error"
	case ErrInternal:
		return "internal server error"
	default:
		return "unknown error"
	}
}
func (e Errors) String() string {
	return e.Error()
}

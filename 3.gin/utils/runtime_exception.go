package utils

type RuntimeException struct {
	err error
	Result
}

func NewRuntimeException() *RuntimeException {
	return &RuntimeException{
		Result: Result{
			Code: -1,
			Msg:  "Runtime exception",
		},
	}
}
func (b *RuntimeException) PanicError(err error) {
	b.err = err
	panic(*b)
}

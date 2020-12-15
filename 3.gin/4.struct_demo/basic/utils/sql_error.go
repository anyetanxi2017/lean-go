package utils

type SQLError struct {
	err error
	Result
}

func NewSQLError() *SQLError {
	return &SQLError{
		Result: Result{
			Code: -1,
			Msg:  "SQL Exception",
		},
	}
}
func (b *SQLError) PanicError(err error) {
	b.err = err
	panic(*b)
}

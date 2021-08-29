package errorex

import "fmt"

type MvcError struct {
	code    int
	message string
}

func (e MvcError) Code() int {
	return e.code
}
func (e MvcError) Error() string {
	return e.message
}

func New(format string, args ...interface{}) error {
	return &MvcError{
		message: fmt.Sprintf(format, args...),
	}
}

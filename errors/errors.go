package errors

import "fmt"

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}

// Newf formats according to a format specifier and returns the string as a
// value that satisfies error.
func Newf(format string, a ...any) error {
	return &errorString{fmt.Sprintf(format, a...)}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

// Error
func (e *errorString) Error() string {
	return e.s
}

package errors

import "fmt"

// AssertPanic
// if condition is true then panic text
func AssertPanic(condition bool, text string) {
	if condition {
		panic(text)
	}
}

// AssertPanicf
// if condition is true then panic text
func AssertPanicf(condition bool, format string, vals ...any) {
	AssertPanic(condition, fmt.Sprintf(format, vals...))
}

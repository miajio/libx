package libx_test

import (
	"testing"

	"github.com/miajio/libx/errors"
)

func TestAssertPanic(t *testing.T) {
	errors.AssertPanic(1 != 1, "number fail")
}

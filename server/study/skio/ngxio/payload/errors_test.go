package payload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpError(t *testing.T) {
	at := assert.New(t)

	tests := []struct {
		op        string
		err       error
		temporary bool
		errStr    string
	}{
		{"read", errPaused, true, "read: paused"},
		{"read", errTimeout, false, "read: timeout"},
	}

	for _, test := range tests {
		var err error
		err = newOpError(test.op, test.err)

		at.Equal(test.errStr, err.Error())

		re, ok := err.(Error)
		at.True(ok)

		at.Equal(test.temporary, re.Temporary())
	}
}

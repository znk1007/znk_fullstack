package sio

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNamespaceHandler(t *testing.T) {
	should := assert.New(t)
	must := require.New(t)

	h := newHandler()

	onConnectCalled := false
	h.OnConnect(func(c Conn) error{
		onConnectCalled = true
		return
	})
	disconnectMsg := ""
	
}

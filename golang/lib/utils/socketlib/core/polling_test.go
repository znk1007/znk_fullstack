package core

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/znk_fullstack/golang/lib/utils/socketlib/protos/pbs"
)

func TestDialOpen(t *testing.T) {
	cp := pbs.ConnParameters{
		PingInterval: int64(time.Second),
		PingTimeout:  int64(time.Minute),
		SID:          "abcdefg",
		Upgrades:     []string{"polling"},
	}
	should := assert.New(t)
	must := require.New(t)

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		query := r.URL.Query()
		should.NotEmpty(r.URL.Query().Get("t"))
		sid := query.Get("sid")
		if sid == "" {
			buf := bytes.NewBuffer(nil)
			writeTo(cp, buf)
		}
	}
}

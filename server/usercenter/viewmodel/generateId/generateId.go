package generate

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/segmentio/ksuid"
)

//GenerateID 生成唯一ID
func GenerateID() string {
	var uID string
	ks, err := ksuid.NewRandom()
	uID = ks.String()
	if err != nil {
		t := time.Unix(1000000, 0)
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		ud := ulid.MustNew(ulid.Timestamp(t), entropy)
		uID = ud.String()
	}
	return uID
}

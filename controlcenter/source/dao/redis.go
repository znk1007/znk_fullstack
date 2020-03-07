package ccdb

import (
	"github.com/go-redis/redis/v7"
)

func init() {
	_ = redis.NewClient(&redis.Options{})
}

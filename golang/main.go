package main

import (
	"znk/golang/lib/index"
	_ "znk/golang/lib/utils/database/cockroachz"
	_ "znk/golang/lib/utils/database/mongodbz"
	_ "znk/golang/lib/utils/database/redisz"
)

func main() {
	index.StartServer()
}

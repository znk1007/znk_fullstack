package main

import (
	"znk/lib/index"
	_ "znk/lib/utils/database/cockroachz"
	_ "znk/lib/utils/database/mongodbz"
	_ "znk/lib/utils/database/redisz"
)

func main() {
	index.StartServer()
}

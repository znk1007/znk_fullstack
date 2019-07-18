package main

import (
	"github.com/znk1007/fullstack/lib/index"
	_ "github.com/znk1007/fullstack/lib/utils/database/cockroachz"
	_ "github.com/znk1007/fullstack/lib/utils/database/mongodbz"
	_ "github.com/znk1007/fullstack/lib/utils/database/redisz"
)

func main() {
	index.StartServer()
}

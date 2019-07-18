package main

import (
	"github.com/znk_fullstack/golang/lib/index"
	_ "github.com/znk_fullstack/golang/lib/utils/database/cockroachz"
	_ "github.com/znk_fullstack/golang/lib/utils/database/mongodbz"
	_ "github.com/znk_fullstack/golang/lib/utils/database/redisz"
)

func main() {
	index.StartServer()
}

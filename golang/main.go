/*
 * @Author: your name
 * @Date: 2019-12-19 22:23:54
 * @LastEditTime : 2019-12-21 18:16:47
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /golang/main.go
 */
package main

import (
	"lib/index"
	_ "github.com/znk_fullstack/golang/lib/utils/database/cockroachz"
	_ "github.com/znk_fullstack/golang/lib/utils/database/mongodbz"
	_ "github.com/znk_fullstack/golang/lib/utils/database/redisz"
)

func main() {
	index.StartServer()
}

package webdemo

import (
	"fmt"
	"net"
)

func RequestJSON(jd *JSONDemo) {
	conn, err := net.Dial("tcp", "120.0.0.1")
	if err != nil {
		fmt.Println("dial err: ", err.Error())
		return
	}
	go processConn(conn)
}

func processConn(conn net.Conn) {
	
}

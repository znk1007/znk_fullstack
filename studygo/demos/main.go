package main

import (
	"github.com/znk_fullstack/studygo/demos/fourteenth"
)

func main() {

	// webdemo.StartServer(false, func(mux *http.ServeMux) {
	// 	webdemo.Hello("world", mux)
	// 	webdemo.JSONRequest(mux)
	// })

	fourteenth.StringLiteralTemplate()
	filePath := "/Users/huangsam/Desktop/golang/src/github.com/znk_fullstack/studygo/demos/testfile.templ"
	fourteenth.FileTemplate(filePath)
	fourteenth.DotActionTemplate()
	filePath = "/Users/huangsam/Desktop/golang/src/github.com/znk_fullstack/studygo/demos/test1.templ"
	fourteenth.AgeInfoTemplate(filePath)
}

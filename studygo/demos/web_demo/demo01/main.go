package main

import (
	"github.com/znk_fullstack/studygo/demos/web_demo/demo01/httprouterdemo"
)

func main() {
	httprouterdemo.GetMethodParam("greet", ":user", "hello world")
	httprouterdemo.PostMethod("test/post")
	httprouterdemo.StartEngine()
}

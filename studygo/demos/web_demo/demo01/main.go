package main

import (
	"github.com/znk_fullstack/studygo/demos/web_demo/demo01/gindemo"
)

func main() {
	// httprouterdemo.GetMethodParam("greet", ":user", "hello world")
	// httprouterdemo.PostMethod("test/post")
	// httprouterdemo.HeadMethod("test/head")
	// httprouterdemo.StartEngine()
	gindemo.SlashGet()
	gindemo.WelcomeGet()
	gindemo.UploadFile()
	gindemo.UploadFiles()
	gindemo.Run(":9527")
}

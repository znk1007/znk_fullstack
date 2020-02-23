package httprouterdemo

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type httprouterDemo struct {
	router httprouter
}

var routerDemo httprouter

func init() {
	routerDemo = httprouterDemo{
		router: httprouter.New(),
	}
}

func StartEngine() {
	http.ListenAndServe(":8080", routerDemo.router)
}

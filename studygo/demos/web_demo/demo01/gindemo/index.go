package gindemo

import "github.com/gin-gonic/gin"

type gindemo struct {
	router *gin
}

var demo gindemo

func init() {
	demo = gindemo{
		router: &gin.Default(),
	}
}

package tools

import "github.com/gin-gonic/gin"


type GinTool struct {
	serve *http.Server
	router *gin.Engie
}

fu
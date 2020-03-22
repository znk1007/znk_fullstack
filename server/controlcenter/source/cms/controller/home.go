package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Home 主页
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "home page",
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "succ",
	// })
}

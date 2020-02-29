package gindemo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

//Run 运行
func Run(port string) {
	router.Run(port)
}

//Router 路由
func Router() *gin.Engine {
	return router
}

//SlashGet 请求
func SlashGet() {
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
}

//WelcomeGet 。。
func WelcomeGet() {
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		//nickname := c.Query("nickname")// c.Request.URL.Query().Get("nickname")的简写
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})
}

//LoginRequest 登录
func LoginRequest() {
	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert") //默认值
		username := c.PostForm("username")
		password := c.PostForm("password")
		//hobbys := c.PostFormMap("hobby")
		//hobbys := c.QueryArray("hobby")
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s, hobby is %v", type1, username, password, hobbys))
	})
}

//UploadFile 上传单个文件
func UploadFile() {
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//upload the file to specific dst
		/*
			c.SaveUploadedFile(file, file.Filename)
			out, err := os.Create(file.Filename)
			defer out.Close()
			_, err = io.Copy(out, file)
		*/
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}

//UploadFiles 上传多个文件
func UploadFiles() {
	router.MaxMultipartMemory = 8 << 20
	router.POST("/uploadfiles", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("uploaded successfully %d files", len(files)))
	})
}

//GroupRequest 路由组
func GroupRequest() {
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	}
}

func loginEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

func submitEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

func readEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

//Login 登录对象
type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

//LoginJSON json登录信息
func LoginJSON() {
	router.POST("/loginJSON", func(c *gin.Context) {
		var lj Login
		router.POST("/loginJSON", func(c *gin.Context) {
			if err := c.ShouldBindJSON(&lj); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
				return
			}
			if lj.User != "ls" || lj.Password != "123456" {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		})
	})
}

//LoginForm 表单json
func LoginForm() {
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		if form.User != "ls" || form.Password != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}

//LoginURI uri请求
func LoginURI() {
	router.GET("/:user/:password", func(c *gin.Context) {
		var uri Login
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"username": uri.User, "password": uri.Password})
	})
}

//IndexTempl 加载模板
func IndexTempl() {
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.templ", gin.H{
			"title": "Main website",
		})
	})
}

//IndexTempl1 加载模板
func IndexTempl1() {
	router.LoadHTMLGlob("templates/**/*", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.templ", gin.H{
			"title": "Posts",
		})
		c.HTML(http.StatusOK, "users/index.templ", gin.H{
			"title": "Users",
		})
	})
}

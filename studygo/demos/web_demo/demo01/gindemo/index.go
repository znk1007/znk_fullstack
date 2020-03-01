package gindemo

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/multitemplate"
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
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.templ", gin.H{
			"title": "Posts",
		})
		c.HTML(http.StatusOK, "users/index.templ", gin.H{
			"title": "Users",
		})
	})
}

//IndexTempl2 加载模板
func IndexTempl2() {
	html := template.Must(template.ParseFiles("file1", "file2"))
	router.SetHTMLTemplate(html)
}

//IndexTempl3 加载模板
func IndexTempl3() {
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("./index.templ")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./index.templ", "<a href='https://lianshiclass.com'>练识课堂</a>")
	})
}

//LoadTemplates 加载自定义模板
func LoadTemplates(templatesDir string) {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.templ")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.templ")
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	router.HTMLRender = r
	router.GET("/index", indexFunc)
	router.GET("/home", homeFunc)
}

func indexFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "index.templ", nil)
}

func homeFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "home.templ", nil)
}

//LoadStaticFiles 加载静态文件
func LoadStaticFiles() {
	router.StaticFS("/showDir", http.Dir("."))
	router.StaticFS("/files", http.Dir("/bin"))
	router.StaticFile("/image", "./assets/image.jpg")
}

//Redirect 重定向
func Redirect() {
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
}

//RouterRedirect 路由重定向
func RouterRedirect() {
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
}

func LongAsync() {
	router.GET("/long_async", func(c *gin.Context) {
		cCP := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + cCP.Request.URL.Path)
		}()
	})
	router.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path " + c.Request.URL.Path)
	})
}

//MiddleWare 中间件
func MiddleWare() {
	hf := func(c *gin.Context) {
		t := time.Now()
		fmt.Println("before middleware")
		c.Set("request", "client_request")
		//发送request之前
		c.Next()
		//发送request之后
		status := c.Writer.Status()
		fmt.Println("after middleware", status)
		t2 := time.Since(t)
		fmt.Println("time: ", t2)
	}
	router.Use(hf)
	router.GET("/middleware", func(c *gin.Context) {
		request := c.MustGet("request").(string)
		req, _ := c.Get("request")
		fmt.Println("request: ", request)
		c.JSON(http.StatusOK, gin.H{
			"middle_request": request,
			"request":        req,
		})
	})
}

func AuthMiddleware() {
	secrets := gin.H{
		"ls": gin.H{
			"email": "ls@lianshiclass.com",
			"phone": "123456",
		},
		"yang": gin.H{
			"email": "yang@lianshiclass.com",
			"phone": "111111",
		},
		"edu": gin.H{
			"email": "edu@lianshiclass.com",
			"phone": "6666666",
		},
	}
	authed := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"ls":   "123",
		"yang": "111",
		"edu":  "666",
		"lucy": "4321",
	}))
	authed.GET("/secrets", func(c *gin.Context) {
		//获取提交的用户名(AuthUserKey).(string)
		u := c.MustGet(gin.AuthUserKey).(string)
		if sec, ok := secrets[u]; ok {
			c.JSON(http.StatusOK, gin.H{
				"user":   u,
				"secret": sec,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"user":   u,
				"secret": "NO SECRET",
			})
		}
	})
}

//CookieRequet cookie相关请求
func CookieRequet() {
	router.POST("/login", cookieMiddleware, func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err != nil {
			c.String(http.StatusOK, "登录失败")
			return
		}
		if login.User != "ls" && login.Password != "123456" {
			return
		}
		c.String(http.StatusOK, "登录成功")
	})
}

func cookieMiddleware(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		return
	}
	if login.User != "ls" && login.Password != "123456" {
		return
	}
	c.SetCookie(
		"username",
		login.User,
		0,
		"/login",
		"localhost",
		true,
		true,
	)
	c.SetCookie(
		"session",
		login.Password,
		0,
		"/login",
		"localhost",
		true,
		true,
	)
}

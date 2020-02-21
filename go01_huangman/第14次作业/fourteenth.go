package fourteenth

import (
	// "crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//User 用户模板
type User struct {
	Name string
	Age  int
}

//StringLiteralTemplate 字符串模板
func StringLiteralTemplate() {
	s := "My name is {{.Name}}. I am {{.Age}} year old.\n"
	t, err := template.New("testuser").Parse(s)
	if err != nil {
		log.Fatal("Parse string literal template error: ", err)
	}
	u := User{
		Name: "lianshi",
		Age:  18,
	}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute string literal template error: ", err)
	}
}

//FileTemplate 文件模板 /Users/huangsam/Desktop/golang/src/
func FileTemplate(fp string) {
	t, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal("Parse file template error: ", err)
	}
	u := User{
		Name: "ls",
		Age:  18,
	}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute file template error: ", err)
	}
	fmt.Println("")
}

func (u User) String() string {
	return fmt.Sprintf("(name:%s age: %d)", u.Name, u.Age)
}

//DotActionTemplate 点动作
func DotActionTemplate() {
	s := "The user is {{.}}"
	t, err := template.New("test1").Parse(s)
	if err != nil {
		log.Fatal("Parse error: ", err)
	}
	u := User{
		Name: "lianshi",
		Age:  18,
	}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute error: ", err)
	}
}

//AgeInfo 年龄信息
type AgeInfo struct {
	Age           int
	GreaterThan60 bool
	GreaterThan40 bool
}

//AgeInfoTemplate 年龄信息模板
func AgeInfoTemplate(fp string) {
	t, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal("Parse file template error: ", err)
	}
	rand.Seed(time.Now().Unix())
	age := rand.Intn(100)
	info := AgeInfo{
		Age:           age,
		GreaterThan60: age > 60,
		GreaterThan40: age > 40,
	}
	err = t.Execute(os.Stdout, info)
	if err != nil {
		log.Fatal("Execute error: ", err)
	}
}

//Item ...
type Item struct {
	Name  string
	Price int
}

//ItemTemplate Item模板
func ItemTemplate(file string) {
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal("Parse error: ", err)
		return
	}
	items := []Item{
		{"iPhone", 699},
		{"iPad", 799},
		{"iWatch", 199},
		{"MacBook", 999},
	}
	err = t.Execute(os.Stdout, items)
	if err != nil {
		log.Fatal("Execute error: ", err)
	}
}

//Pet 宠物
type Pet struct {
	Name  string
	Age   int
	Owner User
}

//PetTemplate 宠物模板
func PetTemplate(file string) {
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal("Parse error: ", err)
	}
	p := Pet{
		Name: "Orange",
		Age:  2,
		Owner: User{
			Name: "ls",
			Age:  18,
		},
	}
	err = t.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal("Execute error: ", err)
	}
	fmt.Println("")
}

//NestTemplate 包含动作模板
func NestTemplate(file1 string, file2 string) {
	t, err := template.ParseFiles(file1, file2)
	if err != nil {
		log.Fatal("nest parse err: ", err)
	}
	err = t.Execute(os.Stdout, "test data")
	if err != nil {
		log.Fatal("nest execute error: ", err)
	}
}

//User1 用户模型1
type User1 struct {
	FirstName string
	LastName  string
}

func (u User1) FullName() string {
	return u.FirstName + " " + u.LastName
}

//User1Template 用户模板1
func User1Template(file string) {
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal("user1 parse err: ", err)
	}
	err = t.Execute(os.Stdout, User1{
		FirstName: "lianshi",
		LastName:  "ls",
	})
	if err != nil {
		log.Fatal("user1 execute err: ", err)
	}
}

//Login 登录
type Login struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	templfile string
}

func (l Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(l.templfile) == 0 {
		log.Fatal("template file cannot be empty")
		return
	}
	t, err := template.ParseFiles(l.templfile)
	if err != nil {
		w.Write([]byte("parse template error: " + err.Error()))
		return
	}

	err = t.Execute(w, l)
	if err != nil {
		w.Write([]byte("execute login error: " + err.Error()))
		return
	}
}

//LoginServe 登录服务
func LoginServe(mux *http.ServeMux, templfile string) {
	if len(templfile) == 0 {
		return
	}
	mux.Handle("/loginpage", Login{templfile: templfile})
}

//LoginResponse 登录响应
func LoginResponse(mux *http.ServeMux) {
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		data := make([]byte, r.ContentLength)
		cnttype := r.Header.Get("Content-Type")
		fmt.Println("content type: ", cnttype)
		_, err := r.Body.Read(data)
		defer r.Body.Close()
		if err != nil && err != io.EOF {
			w.Write([]byte("get data err: " + err.Error()))
			return
		}
		fmt.Println("data: ", string(data))
		lg := &Login{}
		err = json.Unmarshal(data, lg)
		if err != nil {
			fmt.Println("unmarshal err: ", err.Error())
			w.Write([]byte("unmarshal err: " + err.Error()))
			return
		}
		fmt.Println(lg)
		if !(lg.Name == "lianshi" && lg.Password == "123456") {
			w.Write([]byte("登录失败"))
			fmt.Println("登录失败")
			return
		}
		fmt.Println("登录成功")
		w.Write([]byte("欢迎lianshi 同学"))
	})
}

//StartServer 开启服务
func StartServer(muxHandler func(mux *http.ServeMux)) {
	mux := http.NewServeMux()
	if muxHandler != nil {
		muxHandler(mux)
	}
	serve := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := serve.ListenAndServe(); err != nil {
		if muxHandler != nil {
			muxHandler(nil)
		}
	}
}

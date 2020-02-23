package httprouterdemo

//lsof -i tcp:8080
//kill -9 8080
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type httprouterdemo struct {
	router *httprouter.Router
}

var demo httprouterdemo

func init() {
	demo = httprouterdemo{
		router: httprouter.New(),
	}
}

//StartEngine 启动服务
func StartEngine() {
	err := http.ListenAndServe(":8080", demo.router)
	if err != nil {
		log.Fatal("listen and serve err: ", err)
	}
}

func getMethodParamHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("user: ", ps.ByName("user"))
	w.Write([]byte(ps.ByName("user")))
}

//GetMethodParam 带参数get请求
func GetMethodParam(route string, id string, data interface{}) {
	path := route
	if !strings.HasPrefix(route, "/") {
		path = "/" + path
	}
	if len(id) > 0 && strings.HasPrefix(id, ":") {
		path = path + "/" + id
	}
	if len(path) == 0 {
		return
	}
	demo.router.GET(path, getMethodParamHandler)
}

func postMethoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	defer r.Body.Close()
	if r.Header.Get("Content-Type") != "application/json" {
		w.Write([]byte("content-type err"))
		return
	}
	type test struct {
		Name string `json: "name"`
		Sex  int    `json:"sex"`
	}
	var t = &test{}
	err := json.Unmarshal(bs, t)
	if err != nil {
		fmt.Println("unmarshal err: ", err)
		return
	}
	fmt.Println(t)
}

//PostMethod POST请求
func PostMethod(route string) {
	path := route
	if !strings.HasPrefix(route, "/") {
		path = "/" + path
	}
	demo.router.POST(path, postMethoHandler)
}

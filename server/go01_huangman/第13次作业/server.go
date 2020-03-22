package webdemo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//JSONDemo json测试样例
type JSONDemo struct {
	Name string `json:"name"`
	Sex  int    `json:"sex"`
}

//JSONRequest json请求
func JSONRequest(mux *http.ServeMux) {
	if mux == nil {
		http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
			// data, err := ioutil.ReadAll(r.Body)
			data := make([]byte, r.ContentLength)
			_, err := r.Body.Read(data)
			defer r.Body.Close()

			if err != nil {
				fmt.Fprintf(w, "报错了"+err.Error())
				return
			}
			cnttype := r.Header.Get("Content-Type")
			if cnttype != "application/json" {
				fmt.Fprintf(w, "报错了: content-type不是application/json")
				return
			}
			jd := &JSONDemo{}
			err = json.Unmarshal(data, jd)
			if err != nil {
				fmt.Fprintf(w, "报错了: "+err.Error())
				return
			}
			fmt.Println("json demo: ", jd)
			fmt.Fprintf(w, "收到了")
		})
		return
	}
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Fprintf(w, "报错了: "+err.Error())
			return
		}
		cnttype := r.Header.Get("Content-Type")
		if cnttype != "application/json" {
			fmt.Fprintf(w, "报错了: content-type不是application/json")
			return
		}
		fmt.Println("content type: ", cnttype)
		jd := &JSONDemo{}
		err = json.Unmarshal(data, jd)
		if err != nil {
			fmt.Fprintf(w, "报错了: "+err.Error())
			return
		}
		fmt.Println("json demo: ", jd)
		fmt.Fprintf(w, "收到了")
	})
}

//Hello 函数
func Hello(body string, mux *http.ServeMux) {
	if mux == nil {
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello "+body)
		})
		return
	}
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello "+body)
	})
}

//StartServer 启动服务
func StartServer(byDefault bool, muxHandler func(*http.ServeMux)) {
	if byDefault {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
		if muxHandler != nil {
			muxHandler(nil)
		}
		return
	}
	mux := http.NewServeMux()
	if muxHandler != nil {
		muxHandler(mux)
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe" + err.Error())
		if muxHandler != nil {
			muxHandler(nil)
		}
	}
}

package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println(r.URL.Path)
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Method == "GET" {
			t, _ := template.ParseFiles("login.gtpl")
			fmt.Println(t.Execute(w, nil))
		} else {
			fmt.Println("username", r.Form["username"])
			fmt.Println("password", r.Form["password"])
			fmt.Fprintf(w, "ok")
		}
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			curTime := time.Now().Unix()
			h := md5.New()
			io.WriteString(h, strconv.FormatInt(curTime, 10))
			token := fmt.Sprintf("%x", h.Sum(nil))
			t, _ := template.ParseFiles("upload.gtpl")
			t.Execute(w, token)
		} else {
			//r.ParseForm()
			//if token := r.Form.Get("token"); token == "" {
			//	fmt.Fprintf(w, "token 不能为空")
			//	return
			//}else {
			//	fmt.Println(token)
			//}
			r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("uploadFile")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Fprintf(w, "%v", handler.Header)
			f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		}

	})
	fmt.Println("start server http://localhost:9000 ...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("error", err)
	}

}

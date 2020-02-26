package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
	// "io"
	"net/http"
)

var myTemplate *template.Template

type Person struct {
	Name string
	age  string
}

func userInfo(w http.ResponseWriter, r *http.Request) {

	p := Person{Name: "safly", age: "30"}
	myTemplate.Execute(w, p)

}

func mp4Func(w http.ResponseWriter, r *http.Request) {

	video, err := os.Open("./1.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer video.Close()

	http.ServeContent(w, r, "test.mp4", time.Now(), video)
}

func initTemplate(fileName string) (err error) {
	myTemplate, err = template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

/*
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
*/

func main() {
	initTemplate("./test.html")
	http.HandleFunc("/user/info", userInfo)
	http.HandleFunc("/user/mp4", mp4Func)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

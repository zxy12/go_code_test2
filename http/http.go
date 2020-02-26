package main

import (
	"fmt"
	"net/http"
	"time"
)

func userInfo(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format("2006-01-02 15:04:05")
	w.Write([]byte(tm))
}

func main() {
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

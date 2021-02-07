package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/hello", helloHandle)

	log.Fatal(http.ListenAndServe(":9999", nil)) //第二个参数nil，使用标准库中的函数处理
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
}

func helloHandle(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "req.head[%q] = %q\n", k, v)
	}
}

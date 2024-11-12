package main

import (
	"fmt"
	"log"
	"net/http"
)

// 首页处理函数
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "!Welcome to the Go Server!")
}

// 处理 /hello 的请求并返回一个JSON响应
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

// 注册路由
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", helloHandler)
	log.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

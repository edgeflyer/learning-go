// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	// 使用 http.HandleFunc 来注册路由和处理函数
// 	http.HandleFunc("/", handler)
// 	// 启动 HTTP 服务器，监听 localhost:8000
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// 输出请求路径
// 	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
// }

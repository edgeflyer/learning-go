// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		defer resp.Body.Close() // 确保在函数结束时关闭响应体
// 		input := bufio.NewScanner(os.Stdin)
// 		fmt.Println("Enter address:")
// 		input.Scan()
// 		addr := input.Text()
// 		dst, err := os.Create(addr)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
// 			os.Exit(1)
// 		}
// 		defer dst.Close() // 确保在函数结束时关闭文件

// 		n, err := io.Copy(dst, resp.Body)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "%v", err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("Downloaded %d bytes to file %s\n", n, addr)
// 	}
// }

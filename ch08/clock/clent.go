// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// )

// func main() {
// 	// 连接到 TCP 服务器，假设服务器地址是 localhost:8000
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	// 读取并打印从服务器接收到的每一行时间
// 	for {
// 		// 创建一个缓冲区来读取数据
// 		buffer := make([]byte, 1024)
// 		n, err := conn.Read(buffer)
// 		if err != nil && err != io.EOF {
// 			log.Fatal(err)
// 		}
// 		if n > 0 {
// 			// 打印接收到的数据
// 			fmt.Print(string(buffer[:n]))
// 		}
// 	}
// }

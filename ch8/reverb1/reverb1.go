// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"strings"
// 	"time"
// )

// func echo(c net.Conn, shout string, delay time.Duration) {
// 	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", shout)
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", strings.ToLower(shout))
// }

// func handleConn(c net.Conn) {
// 	input := bufio.NewScanner(c)
// 	for input.Scan() {
// 		echo(c, input.Text(), 1*time.Second)
// 	}
// 	c.Close()
// }

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for {
// 		conn, err := listener.Accept() // 阻塞
// 		if err != nil {
// 			log.Print(err)
// 			continue
// 		}
// 		go handleConn(conn) // 一次处理一个连接,在这里加个go就能同时处理多个请求，每次运行到这里开一个goroutine
// 	}
// }

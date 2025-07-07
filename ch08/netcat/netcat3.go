package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 捕获中断信号（Ctrl+C）
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // 指示主goroutine
	}() // 注意这里一定要有括号
	mustCopy(conn, os.Stdin)

	<-sigChan // 等待中断信号（Ctrl+C）
	log.Println("Received SIGINT, waiting for goroutines to finish...")
	log.Println("goroutine finished...")
	// conn.Close()
	// 等待后台 goroutine 完成
	conn.Close()
	<-done
	// 输出 wanle 后退出
	fmt.Println("wanle")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

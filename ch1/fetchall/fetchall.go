// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"time"
// )

// // fetch并发获取url并报告他们的时间和大小
// func main() {
// 	start := time.Now()     //记录程序开始时间
// 	ch := make(chan string) //使用make创建一个字符串通道（用于接收fetch函数返回的数据）。对于每个命令行参数，go语句在第一轮循环中启动一个新的routine，异步调用fetch来获取url内容。
// 	for _, url := range os.Args[1:] {
// 		go fetch(url, ch) //// 启动一个新的 goroutine，执行 fetch(url, ch)
// 	}
// 	// 从通道中接收数据，直到接收完所有 URL 对应的结果
// 	for range os.Args[1:] {
// 		fmt.Println(<-ch) //// 输出从 ch 通道接收到的每个 fetch 函数的返回结果
// 	}
// 	// 输出总的耗时
// 	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds)
// }

// func fetch(url string, ch chan<- string) {
// 	start := time.Now() // 记录 fetch 开始的时间
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprint(err) // 如果出错，将错误信息发送到通道
// 		return
// 	}
// 	// 用 io.Copy 把响应体内容丢弃，通过 io.Discard 流进行丢弃，只关注字节数
// 	nbytes, err := io.Copy(io.Discard, resp.Body) //读取响应内容，通过写入io.Discard输出流进行丢弃。copy返回字节数以及出现的任何错误。每一个结果返回时，fetch发送一行汇总信息到通道ch，main中的第二轮循环接收并输出那些汇总行
// 	resp.Body.Close()                             //不要泄露资源
// 	if err != nil {
// 		ch <- fmt.Sprintf("while reading %s: %v", url, err) // 读取内容出错时发送错误信息
// 		return
// 	}
// 	// 计算请求花费的时间并通过通道发送结果
// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) //Sprintf作用是格式化字符串，让所有里面的变量格式化为一个字符串
// }

// /*
// 通道是一种用于在goroutine之间传递数据的核心特性，提供了一种安全高效的方式来实现不同goroutine之间的通信和同步，通道就像一个线程安全队列，可以在并发程序中传递信息或数据，避免了显式的锁和其他同步机制的使用。
// 通过通道，数据可以在一个goroutine中发送，然后在另一个goroutine中接收。
// 默认情况下通道时双向的，可以用于发送和接收数据，可以将通道限制为单向的。
// 发送/接收数据用<-
// 会有阻塞行为，和缓冲区类似
// 无缓冲区的通道每次发送和接受都会直接阻塞，直到双方准备好
// 接收关闭的通道，如果通道中有值返回的是值和布尔值，没有数据的话就返回false
// */

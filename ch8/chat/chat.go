package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // 监听8000号端口
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster() // 广播goroutine,会一直运行
	for {
		conn, err := listener.Accept() // 阻塞到监听到tcp请求
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // 处理每个连接,每当有新的客户端连接时，创建一个性的goroutine来处理该连接
	}
}

type client chan<- string //双向通道，用来发送和接收消息
var (                     // 全局通道
	entering = make(chan client) // 接收新的clinet连接，用来通知广播器有新的客户端连接
	leaving  = make(chan client) // 接收客户端退出的信号。当客户端离开时，它会从clients列表中删除该客户端，并关闭其消息通道
	messages = make(chan string) // 所有接收的客户消息，当广播者接收到其中一个事件时，他把消息广播给所有在线客户
)

func broadcaster() {
	clients := make(map[client]bool) // 记录所有已连接的客户端，键是client类型的通道值用来标记客户端是否在线
	for {
		select { // 使用select监听三个通道
		case msg := <-messages: // 如果接收到消息（客户端发的消息），则广播给所有已连接的客户端
			// 把所有接受的消息广播给所有的客户
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering: //当有新的客户端连接时，将该客户端添加到clients中
			clients[cli] = true
		case cli := <-leaving: // 当客户端断开连接时，从clients中删除该客户端，并关闭消息通道
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)   // 为每个客户端创建一个消息通道ch，用于发送消息给客户端
	go clientWriter(conn, ch) // 启动一个goroutine来负责将消息从ch通道发送到客户端

	who := conn.RemoteAddr().String() // 获取客户端的地址
	ch <- "You are " + who            // 向客户端发送欢迎消息，告诉它它的地址
	messages <- who + " has arrived"  // 将客户端到来的消息广播给所有客户端
	entering <- ch                    // 将该客户端的通道添加到entering通道，广播器会接收到这个信号并将该客户端加入clients中

	input := bufio.NewScanner(conn) // 使用bufio.NewScanner来扫描客户端发送的数据
	for input.Scan() {
		messages <- who + ": " + input.Text() // 每当客户端发送一条消息时，将其发送到messages通道，广播给所有其他客户端
	}
	// 注意，忽略input.Err()中可能的错误

	leaving <- ch // 客户端断开连接时，将其通道发送到leaving通道，广播器会删除该客户端并关闭通道
	messages <- who + " has left"
	conn.Close() // 关闭与客户端的连接
}

// 负责将接受到的消息写入到客户端的连接
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch { // 从ch通道中接受消息，并通过fmt.Fprintln将消息发送到客户端
		fmt.Fprintln(conn, msg) // 忽略网络层面的错误
	}
}

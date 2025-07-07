package bank

// 通过用通道模仿信号量机制来保证进程的同步
var (
	sema    = make(chan struct{}, 1) // 用来保护balance的二进制信号量
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // 获取令牌
	balance = balance + amount
	<-sema // 释放令牌
}

func Balance() int {
	sema <- struct{}{} // 获取令牌
	b := balance
	<-sema // 释放令牌
	return b
}

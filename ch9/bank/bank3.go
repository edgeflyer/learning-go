package bank

// 使用sync.Mutex互斥锁
import "sync"

var (
	mu      sync.Mutex // 保护balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

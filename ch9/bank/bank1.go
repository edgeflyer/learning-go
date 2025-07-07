package bank

var deposits = make(chan int) // 发放存款额
var balances = make(chan int) // 接受余额

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance被限制在teller goroutine中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

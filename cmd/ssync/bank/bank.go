package bank

var deposits = make(chan int) //发送存款额
var balances = make(chan int) // 接受余额

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int //balance 被限制再teller goroutine中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // 启动监控goroutine
}

// 返回结果应该要表明交易是成功了还是因为没有足够资金失败了。这条消息会被发送给monitor的
// goroutine，且消息需要包含取款的额度和一个新的channel，这个新channel会被monitor
// goroutine来把boolean结果发回给Withdraw。
/* var isDepositOK = make(chan bool)

func Withdraw(amount int) bool {

}
*/

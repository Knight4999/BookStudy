package bank3

import "sync"

var (
	mu      sync.Mutex //保护balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance

}

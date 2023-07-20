package bank4

import "sync"

var (
	mu      sync.RWMutex //读写互斥锁
	balance int
)

func Balance() int {
	mu.RLock() //读锁
	defer mu.RUnlock()
	return balance
}

package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//产生自然数
	go func() {
		for x := 0; x < 10; x++ {
			ch1 <- x
		}
		close(ch1)
	}()
	//计算平方
	go func() {
		/* for {
			x, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- x * x
		} */
		for x := range ch1 {
			ch2 <- x * x
		}
		close(ch2)
	}()
	for x := range ch2 {
		fmt.Println(x)
	}
}

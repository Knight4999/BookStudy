// 测试定时器相关内容
package test5

import (
	"fmt"
	"time"
)

func A() {
	//Timer 时间到了，执行一次
	//1.基本使用
	timer := time.NewTimer(1 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := timer.C
	fmt.Printf("t2: %v\n", t2)

	/* //2.验证Timer只执行一次
	timer2 := time.NewTimer(2 * time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到")
	} */

	// 3.timer实现延时的功能
	//(1)
	time.Sleep(time.Second)
	//(2)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	//(3)
	<-time.After(2 * time.Second)
	fmt.Println("2秒到")

	//4.停止定时器
	timer4 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4 已关闭")
	}

	//5.重置定时器
	timer5 := time.NewTimer(4 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

}

func B() {
	//Ticker，时间到了，多次执行
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			<-ticker.C
			i++
			fmt.Println(i)
			if i == 10 {
				ticker.Stop()
				fmt.Println("ticker已停止")
			}
		}
	}()

}

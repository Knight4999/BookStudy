package main

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct{}

func (d dryer) dry() {
	fmt.Println("甩干器")
}

type haier struct {
	dryer //嵌入甩干器
}

func (h haier) wash() {
	fmt.Println("海尔洗衣机")
}

func main() {
	var w WashingMachine
	w = haier{}
	w.wash()
	w.dry()
}

package main

import (
	"fmt"
	"time"
)

var channel1 = make(chan interface{}, 1)
var channel2 = make(chan interface{}, 1)

func num1() {
	for i := 1; i < 100; {
		<-channel1
		fmt.Println(i)
		channel2 <- 0

		i += 2
	}
}
func num2() {
	for i := 2; i < 101; {
		<-channel2
		fmt.Println(i)
		channel1 <- 0
		i += 2
	}
}
func main() {
	channel1 <- 0
	go num1()
	go num2()
	time.Sleep(3 * time.Second)
}

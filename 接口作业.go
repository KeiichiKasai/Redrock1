package main

import "fmt"

func (c cat) say() {
	fmt.Println("喵")
}
func (d dog) say() {
	fmt.Println("汪")
}

type sayer interface {
	say()
}

type cat struct{}
type dog struct{}

func main() {
	var a sayer
	A := cat{}
	a = A
	a.say()
	B := dog{}
	a = B
	a.say()
}

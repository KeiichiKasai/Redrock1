package main

import "fmt"

func main() {
	fmt.Println("Calculator Start Link")
	fmt.Println("input:")
	var a int
	var ch string
	var b int
	fmt.Scan(&a, &ch, &b)
	switch ch {
	case "+":
		fmt.Printf("%d", a+b)
	case "-":
		fmt.Printf("%d", a*b)
	case "*":
		fmt.Printf("%d", a*b)
	case "/":
		fmt.Printf("%d", a/b)
	}
}

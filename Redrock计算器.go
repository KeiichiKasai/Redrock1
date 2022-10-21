package main

import "fmt"

func add(a int, ch string, b int) {
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
func main() {
	var a int
	var b int
	var ch string
	fmt.Scan(&a, &ch, &b)
	add(a, ch, b)

}

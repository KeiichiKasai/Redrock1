package main

import "fmt"

func main() {
	var a string
	fmt.Scanln(&a)
	b := []rune(a)
	for i := 0; i < len(b)/2; i++ {
		if b[i] == b[len(b)-1-i] {
			fmt.Printf("%s", string(b[i]))
			continue
		} else {
			fmt.Println("false")
			break
		}

	}

}

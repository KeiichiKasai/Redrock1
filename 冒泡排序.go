package main

import "fmt"

func BubbleSort(array [14]int) {
	for i := 0; i < (len(array) - 1); i++ {
		for j := 0; j < (len(array) - 1); j++ {

			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	fmt.Println(array)
}
func main() {
	var array = [...]int{9, 6, -8, 5, 2, 1, 11, 0, 54, 2, 1, 34, 1, 9}
	BubbleSort(array)
}

package main

import (
	"fmt"
	"sort"
)

type arr [14]int

func (a *arr) Len() int {
	return len(a)
}
func (a *arr) Less(i, j int) bool {
	return a[i] < a[j]
}
func (a *arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func main() {
	array := arr{9, 6, -8, 5, 2, 1, 11, 0, 54, 2, 1, 34, 1, 9}
	sort.Sort(&array)
	fmt.Println(array)
}

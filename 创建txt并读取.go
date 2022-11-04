package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFile(filename string) {
	_, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("create file successfully")
	}
}
func ReadFile4(filename string) {
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0400)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	fmt.Printf("%s", data)
}
func WriteString(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0200)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := file.WriteString(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	n--
	file.Close()
}

func main() {
	CreateFile("test.txt")
	WriteString("test.txt", "梁明海好帅！！！！！！！")
	ReadFile4("test.txt")
}

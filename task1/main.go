package main

import "fmt"

func IsValid(id int, text string) bool {
	return id > 0 && text != ""
}

func main() {
	fmt.Println(IsValid(0, "hello world"))
	fmt.Println(IsValid(-22, "hello world"))
	fmt.Println(IsValid(22, ""))
	fmt.Println(IsValid(225, "hello world"))

}

package main

import "fmt"

// BEGIN (write your solution here)
const (
	okMsg        = "OK"
	cancelledMsg = "CANCELLED"
)
const (
	okCode = iota
	cancelledCode
	UnknownCode
)

func ErrorMessageToCode(msg string) int {
	switch msg {
	case okMsg:
		return okCode
	case cancelledMsg:
		return cancelledCode
	default:
		return UnknownCode
	}
}
func main() {
	fmt.Println(ErrorMessageToCode("CANCELLED"))
	fmt.Println(ErrorMessageToCode(""))
	fmt.Println(ErrorMessageToCode("OK"))
}

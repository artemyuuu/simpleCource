package main

import (
	"fmt"
	"strings"
)

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string
	Age       int
}

// BEGIN (write your solution here)

var invalidRequest = "invalid request"

func Validate(req UserCreateRequest) string {
	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return invalidRequest
	}
	if req.Age <= 0 || req.Age > 150 {
		return invalidRequest
	}
	return ""
}
func main() {
	userCreate1 := UserCreateRequest{
		FirstName: "Artyom",
		Age:       156,
	}
	fmt.Println(Validate(userCreate1))
}

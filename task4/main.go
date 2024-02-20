package main

import (
	"fmt"
	"strings"
)

func ModifySpaces(s, mode string) string {
	switch mode {
	case "dash":
		s = strings.ReplaceAll(s, " ", "-")
	case "underscore":
		s = strings.ReplaceAll(s, " ", "_")
	default:
		s = strings.ReplaceAll(s," ", "*")
	}
	return s
}
func main(){
	fmt.Println(ModifySpaces(" hello world ", "dash"))
}

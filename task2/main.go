package main

import (
	"fmt"
	"strings"
)

func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)
	return fmt.Sprintf("Привет, %s!", name)
}
func main() {
	result := Greetings(" IVan ")
	fmt.Println(result)
}

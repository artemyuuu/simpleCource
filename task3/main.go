package main

import (
	"fmt" 
)

func DomainForLocale(domain, locale string) string {
	if locale == ""{
		locale = "en"
	}
	return fmt.Sprintf("%s.%s", locale, domain)
}
func main() {
	website := DomainForLocale("english.com", "")
	fmt.Println(website)
}

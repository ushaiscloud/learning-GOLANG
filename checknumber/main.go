package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, char := range arg {
		if char >= '0' && char <= '9' {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(CheckNumber("hello"))
	fmt.Println(CheckNumber("me12"))
}

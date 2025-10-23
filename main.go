package main

import "fmt"

func main() {
	for char := 'z'; char >= 'a'; char-- {
		fmt.Printf("%c", char)
	}
	fmt.Println()
}

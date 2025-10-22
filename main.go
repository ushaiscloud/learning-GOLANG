package main

import "fmt"

func RetainFirstHalf(str string) string {
	length := len(str)

	// If string is empty, return empty string
	if length == 0 {
		return ""
	}

	// If string length is 1, return the string
	if length == 1 {
		return str
	}

	// Calculate half length (round down for odd numbers)
	half := length / 2

	// Return first half of the string
	return str[:half]
}

func main() {
	// Test cases
	fmt.Println(RetainFirstHalf("Hello"))       // "He"
	fmt.Println(RetainFirstHalf("World!"))      // "Wor"
	fmt.Println(RetainFirstHalf("Go"))          // "G"
	fmt.Println(RetainFirstHalf("A"))           // "A"
	fmt.Println(RetainFirstHalf(""))            // ""
	fmt.Println(RetainFirstHalf("Programming")) // "Progr"
}

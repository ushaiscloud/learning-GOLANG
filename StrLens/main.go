package main

import (
	"fmt"
)

func StrLen(s string) int {
	count := 0

	for range s {
		count++
	}
	return count
}

func main() {
	i := StrLen("odumod")
	fmt.Println(i)
}

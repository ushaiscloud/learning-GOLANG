package main

import "fmt"

func PrintIfNot(str string) string {

	if len(str) >= 3 || str == "" {
		return "G\n"

	}

	return "invalid Input\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ANSI color codes for colorful output
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
)

func main() {
	fmt.Println(Cyan + "✨ Welcome to Talk It Tuesday! ✨\n" + Reset)

	reader := bufio.NewReader(os.Stdin) // allows reading input including spaces

	var num int
	fmt.Print("How many participants are sharing today? ")
	fmt.Scan(&num)

	participants := make([]string, 0, num)
	responses := make([]string, 0, num)

	// Collect participant names and responses
	for i := 0; i < num; i++ {
		fmt.Printf("\nEnter participant %d name: ", i+1)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Printf("Hi %s! What’s your Talk Tuesday response? ", name)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		participants = append(participants, name)
		responses = append(responses, response)
	}

	// Colors to rotate through
	colors := []string{Red, Green, Yellow, Blue, Purple}

	fmt.Println("\n Participants' Responses:\n")

	for i, name := range participants {
		response := responses[i]
		color := colors[i%len(colors)]
		fmt.Printf("%s%s:%s ", Cyan, name, Reset)
		displayAnimated(color, response)
		fmt.Println()
	}

	fmt.Println("\n" + Green + "🙏 Thank you all for sharing during Talk Tuesday!" + Reset)
}

// displayAnimated prints text word by word with color and delay
func displayAnimated(color, text string) {
	words := strings.Split(text, " ")
	for _, word := range words {
		fmt.Print(color + word + " " + Reset)
		time.Sleep(300 * time.Millisecond)
	}
}

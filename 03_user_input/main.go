package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter your name: ")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Print("thank you ", input)
}

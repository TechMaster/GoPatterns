package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println("Enter text: ")
	text2 := ""
	fmt.Scanln(&text2)
	fmt.Println(text2)

	var X int
	var Y int

	fmt.Printf("\nIntital X: %d, Y: %d", X, Y)

	fmt.Sscan("100\n200", &X, &Y)
	fmt.Printf("\nSscan X: %d, Y: %d", X, Y)

	fmt.Sscanf("(10, 20)", "(%d, %d)", &X, &Y)
	fmt.Printf("\nSscanf X: %d, Y: %d", X, Y)

	fmt.Sscanln("50\n50", &X, &Y)
	fmt.Printf("\nSscanln X: %d, Y: %d", X, Y)
}
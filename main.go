package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	scanner *bufio.Scanner
	choice  string
)

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	for {
		mainMenu()
		fmt.Print("Select an option: ")
		choice := warraInput()
		switch choice {
		case "1":
			fmt.Println("You selected Option One")
		case "2":
			fmt.Println("You selected Option Two")
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func warraInput() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func mainMenu() {
	fmt.Println("Main Menu")
	fmt.Println("1. Option One")
	fmt.Println("2. Option Two")
	fmt.Println("3. Exit")
}

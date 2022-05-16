package main

import (
	"fmt"
	"os"
)

func main() {

	showIntroduction()
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Quitting the program")
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		os.Exit(-1)
	}
}

func showIntroduction() {
	var name string = "John"
	var version float32 = 1.1
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is running on version", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Quit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The command chosen was", command)

	return command
}

func quitProgram() {

}

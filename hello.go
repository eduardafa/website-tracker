package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		showIntroduction()
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
		case 0:
			fmt.Println("Quit program")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
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

func startMonitoring() {
	fmt.Println("Monitoring...")
	website := "https://random-status-code.herokuapp.com/"
	response, _ := http.Get(website)
	if response.StatusCode == 200 {
		fmt.Println("Website:", website, "was loaded successfully!")
	} else {
		fmt.Println("Something wrong with the website:", website, "- Status Code:", response.StatusCode)
	}
}

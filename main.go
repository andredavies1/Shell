package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		//Shows the $ prompt
		fmt.Print("$ ")

		//Capture the user's command in the command variable
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		//If reading failed, stops the program
		if err != nil {
			return
		}
		//If command is equal to exit, end program
		if command == "exit\n" {
			os.Exit(0)
		}

		//Cleans the invisible \n off the end of the command
		command = strings.TrimSpace(command)
		//Prints the message, with command filling the %s blank
		fmt.Printf("%s: command not found\n", command)
	}
}

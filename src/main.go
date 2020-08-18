package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// todo
	// 1) go get API key

	fmt.Println("Hello, and welcome to this awesome thing.")
	setup()
}

func setup() {
	var apiKey string

	fmt.Print("thisis a test:")
	fmt.Scan(&apiKey)

	if _, err := os.Stat("settings.mjw"); err == nil {
		// file exists and read in.
		file, err := os.Open("settings.mjw")
		if err != nil {
			panic(err)
		}

		var choice string
		fmt.Println("It appears that you have already set up an API key.")
		fmt.Print("Would you like to use the saved key? (Y/n)")
		fmt.Scan(&choice)

		if choice == "Y" || choice == "y" {
			// use key
		} else {
			getCreds(&apiKey)
		}

		scanner := bufio.NewScanner(file)
		count := 0
		for scanner.Scan() {
			// todo
			count++
		}
	} else {
		getCreds(&apiKey)
		var saveCreds string

		fmt.Println()
		fmt.Print("Would you like to save your API key for easier access next time?  (Y/n)")
		fmt.Scan(&saveCreds)

		if saveCreds == "Y" || saveCreds == "y" {
			// write file
			file, err := os.Create("settings.mjw")
			defer file.Close()

			if err != nil {
				panic(err)
			}

			fmt.Fprintln(file, apiKey)
		}
	}
}

func getCreds(apiKey *string) {
	// io to get creds
	fmt.Print("Please enter your API key: ")
	fmt.Scan(apiKey)
}

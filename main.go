package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Welcome to helixjoe's file organizer. We provide..\n")
	fmt.Printf("--------------------------- \n")
	fmt.Printf(" 1. Directory Scanning \n")
	fmt.Printf("--------------------------- \n")
	fmt.Printf("Enter an option : ")
	var option int
	fmt.Scan(&option)
	switch option {
	case 1:
		files, err := os.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name(), file.Type(), file.IsDir())
		}
	default:
		fmt.Println("Choose a better option")
	}

}

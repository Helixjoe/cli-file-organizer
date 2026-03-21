package main

import (
	"fmt"
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
		dir_scan()
	default:
		fmt.Println("Choose a better option")
	}

}

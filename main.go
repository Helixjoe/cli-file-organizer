package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Welcome to helixjoe's file organizer. We provide..\n")
	var option int
	for option != 3 {
		fmt.Printf("--------------------------- \n")
		fmt.Printf(" 1. Directory Scanning \n")
		fmt.Printf("--------------------------- \n")
		fmt.Printf("Enter an option : ")
		fmt.Scan(&option)
		switch option {
		case 1:
			dir_scan(true)
			break
		case 2:
			organize_files_by_type()
			break
		case 3:
			revert_organization()
			break
		case 4:
			break
		default:
			fmt.Println("Choose a better option")
			break
		}
	}
}

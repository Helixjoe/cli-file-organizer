package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Directory Scanning \n")
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.Type(), file.IsDir())
	}
}

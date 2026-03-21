package main

import (
	"fmt"
	"log"
	"os"
)

func dir_scan() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n 📄 - FILES, 📂 - FOLDERS")
	fmt.Println("----------------")
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println("📄", file.Name())
		} else {
			fmt.Println("📂", file.Name())
		}
	}
	fmt.Println("----------------")
}

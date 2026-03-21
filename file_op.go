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

	for _, file := range files {
		fmt.Println(file.Name(), file.Type(), file.IsDir())
	}
}

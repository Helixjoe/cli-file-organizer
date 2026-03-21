package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func dir_scan(legend bool) {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	if legend {
		fmt.Println("\n 📄 - FILES, 📂 - FOLDERS")
	}
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

func organize_files_by_type() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	// create map of file types
	fileTypes := make(map[string]int)
	for _, file := range files {
		if !file.IsDir() {
			fileType := strings.Split(file.Name(), ".")[1]
			if fileTypes[fileType] == 0 {
				os.Mkdir(fileType, 0750)
			}
			fileTypes[fileType] += 1
			os.Rename(file.Name(), fileType+"/"+file.Name())
		}
	}
	fmt.Println("FILES ARE NOW ORGANIZED INTO FOLDERS")
	dir_scan(false)
}

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
	fileTypes := make(map[string]int) // created map of file types
	for _, file := range files {
		if !file.IsDir() {
			fileType := strings.Split(file.Name(), ".")[1]
			if fileTypes[fileType] == 0 {
				os.Mkdir(fileType, 0750) // directory created
			}
			fileTypes[fileType] += 1
			os.Rename(file.Name(), fileType+"/"+file.Name())
		}
	}
	fmt.Println("FILES ARE NOW ORGANIZED INTO FOLDERS")
	dir_scan(false)
}

func revert_organization() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			dirFiles, err := os.ReadDir(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, dirFile := range dirFiles {
				err := os.Rename(file.Name()+"/"+dirFile.Name(), dirFile.Name())
				if err != nil {
					log.Fatal(err)
				}
			}
			os.Remove(file.Name())
		}
	}
	// dir_scan(false)
}

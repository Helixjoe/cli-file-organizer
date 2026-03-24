package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
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

func delete_duplicate_files() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	freqTable := make(map[string]int)
	noDuplicatesFound := true
	duplicatesDeleted := 0
	for _, file := range files {
		if !file.IsDir() {
			hash, err := hashFile(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			freqTable[hash] += 1
			if freqTable[hash] > 1 {
				fmt.Println(file.Name() + "DELETED")
				os.Remove(file.Name())
				duplicatesDeleted += 1
				noDuplicatesFound = false
			}
		}
	}
	if noDuplicatesFound {
		fmt.Println("\n NO DUPLICATES FOUND")
	}
	if duplicatesDeleted > 0 {
		fmt.Println("\n", duplicatesDeleted, " DUPLICATE(S) DELETED")
	}
	// prevent repeated deletion
	for k := range freqTable {
		delete(freqTable, k)
	}

}

// hashFile generates the SHA256 hash of a file.
func hashFile(filePath string) (string, error) {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	// Ensure the file is closed when the function returns
	defer f.Close()

	// Create a new SHA256 hash instance
	h := sha256.New()

	// Copy the file content to the hash instance
	// io.Copy reads the file in chunks (32KB buffer by default)
	// and writes them to the hasher, which calculates the hash incrementally.
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	// Get the finalized hash result as a byte slice and encode it to a hexadecimal string
	hashInBytes := h.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}

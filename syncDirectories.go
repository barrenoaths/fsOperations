package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("--- Sync stuff ---")

	dirPath1 := "/opt/learnGo"
	fileToHashPath := "/opt/learnGo/syncDirectories/orchid.txt"

	if dirExists(dirPath1) {
		fmt.Printf("Directory %s exists\n", dirPath1)
	}

	if fileExist(fileToHashPath) {
		fmt.Printf("File %s exists\n", fileToHashPath)
	}

	wantedHash, _ := getFileHash(fileToHashPath)
	fmt.Printf("Hash of the %s is %s\n", fileToHashPath, wantedHash)

	fmt.Println("---------------------------")

	printTheContentsOfDir(dirPath1)

}

func dirExists(absPath string) bool {
	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func fileExist(absPath string) bool {
	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func printTheContentsOfDir(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory.", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func printTheContentsOfFile(filePath string) {
	fmt.Println("placeholder")
}

func copyFile() {
	fmt.Println("placeholder")
}

func copyDir() {
	fmt.Println("placeholder")
}

func copyDirWithFiles() {
	fmt.Println("placeholder")
}

func findLineInFile() {
	fmt.Println("placeholder")
}

func findPhraseInFile() {
	fmt.Println("placeholder")
}

func getFileMetadata() {
	fmt.Println("placeholder")
}

func modifyFileMetadata() {
	fmt.Println("placeholder")
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <file-path>")
        os.Exit(1)
    }

    filePath := os.Args[1]
	lines := make(chan string)

	go readLines(filePath, lines)

	processLines(lines)
}

func processLines(lines <-chan string) {
	for line := range lines {
		fmt.Println(line)
	}
}

func readLines(filePath string, lines chan<- string) {
	defer close(lines)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

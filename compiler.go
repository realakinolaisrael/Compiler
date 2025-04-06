package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    filePath := "example.txt"

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close() // Ensures the file is closed when the function exits

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // Print each line
        fmt.Println(scanner.Text())
    }

    // Checking for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}
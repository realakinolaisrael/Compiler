package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Specify the file to read
    filePath := "example.txt" // Replace with your file name

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close() // Ensure the file is closed when the function exits

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // Print each line
        fmt.Println(scanner.Text())
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}
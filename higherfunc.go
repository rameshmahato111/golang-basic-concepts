package main

import "fmt"

// Simple function to add two numbers
func addNumber(a int, b int) int {
    return a + b
}

// Higher-order function that takes another function as an argument
func higerFunc(a int, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    // Call the higher-order function with the add function as an argument
    result := higerFunc(30, 10, addNumber)

    // Print the result
    fmt.Println("This is the higher-order function result and the value is", result)
}

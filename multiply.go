package main

import "fmt"

func main() {
    const a, b int = 20, 30
    const c = a * b
    fmt.Println("area is", c)

    if c%2 == 0 {
        fmt.Println("this is an even number")
    } else {
        fmt.Println("this is an odd number")
    }
}

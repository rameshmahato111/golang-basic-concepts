package main

import "fmt"

// Function with parameters
// func greet(name string) {
//     fmt.Println("Hello,", name)
// }

// func main() {
//     greet("Alice")
// }


func birthday(name string, age int) {
	fmt.Println("hello ", name, "happy", age, "birth day")
}

func main(){
	birthday("Riyashna", 2)
}
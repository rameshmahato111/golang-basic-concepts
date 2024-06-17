
//  Pointers in Go are an essential concept that allows you to work with memory addresses directly. 
//  Understanding pointers is crucial for writing efficient and effective Go programs. 
//  Here’s a detailed explanation of pointers, along with examples to illustrate their use.



package main
import "fmt"

func main(){
	// var ptr *int
	// fmt.Println("this is the pointer", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("this is the pointer", ptr)
	fmt.Println("this is the pointer", *ptr)



}
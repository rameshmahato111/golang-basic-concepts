package main

import "fmt"

type User struct {
	Name string
	Email string
	Status bool
	Age int

	
}
func main(){

	fmt.Println("structs in go lang")

	hitesh := User{"ramesh", "ramesh@gmail.com", true, 28}
	fmt.Println("name is", hitesh.Name)
	fmt.Println("email is", hitesh.Email)
	fmt.Println("status is", hitesh.Status)
	fmt.Println("age is", hitesh.Age)





	
}



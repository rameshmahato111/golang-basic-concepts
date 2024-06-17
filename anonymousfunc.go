// this is the type of anonymous function


package main

import "fmt"


func main(){
	greet := func(name string){
    
		fmt.Println("hello", name)
  
	}
	greet("ramesh")


	func(){
		fmt.Println("this is the anonymous function")
	} ()
}
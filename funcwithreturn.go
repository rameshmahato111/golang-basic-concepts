// this is the function with return values



package main

import "fmt"


func giveBack (a, b int) int{

   return a+b;
}

func main(){
	result :=  giveBack (30, 50)

	fmt.Println("the result value is ", result)
}
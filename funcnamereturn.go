// in this function we can return name value as shown in the below code 


package main

import "fmt"


func nameReturn (a, b int)(add int, multiply int){

      add = a +b
      multiply = a*b

	  return add, multiply
}


func main(){

	sum, multi :=  nameReturn(30, 40)

	

   fmt. Println("add is ", sum, "multiple value is", multi)
}
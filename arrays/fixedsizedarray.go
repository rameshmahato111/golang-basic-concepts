package main

import "fmt"



func main(){
	// an array of five integer
	var arr [5] int   
	fmt.Println("arrays are", arr)
    
	// initializing value of arays
	
	 initialarr := [5] int {1, 2, 3, 4, 5} 

	fmt.Println("initialized array values", initialarr)

	// type of implicit arrays

	implicit :=[] int {1, 2, 3}

	fmt.Println("type of implicit arays", implicit)


}
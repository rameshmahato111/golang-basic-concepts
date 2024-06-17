// this is the method of accessing value of arrays


package main

import "fmt"


func main(){

	arr := [5] string {"apple", "banana", "mango", "pineapple", "orage"}

	fmt.Println("accessing the first element of array", arr[0])
	fmt.Println("accessing the second element of array", arr[1])
	fmt.Println("accessing the third element of array", arr[2])
	fmt.Println("accessing the fourth element of array", arr[3])
	fmt.Println("accessing the fourth element of array", arr[4])

	// insert elements in array 
    // it replace the previous value 
	arr[1] = "blueberry"


	fmt.Println("accessing the whole array", arr)
}
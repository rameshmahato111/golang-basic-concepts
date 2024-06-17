// declaring and initializing multidimensional arrays

package main

import "fmt"

func main(){
    //  here is the declaration of 2*3 dimensional array
	var matrix [2][3] int  

	// asignin value 
	// asign 1 value at the first row and first column of the matrix
	matrix [0][0]= 1

	matrix [1][2] = 5

	fmt.Println("array of matrix is", matrix)

}
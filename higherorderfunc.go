// in the higher order function, we can take function as arguement or return functions
// igher-order functions are functions that take other functions as arguments or return functions as their result. They are powerful tools for creating flexible and reusable code.


package main
import "fmt"


func higherorder (a, b int, op func(int, int) int) int{
	
    return op (a, b)
}



func main(){
	add := func (a, b int) int {
		return a + b
	}

	result := higherorder(40, 30, add)

	fmt.Println("this is the higher order function", result)
}
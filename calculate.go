package main
import "fmt"


func add(a, b int) int {

	return  a +b
}


func subtract(a, b int) int {
	return a -b
}

func multiply(a, b int) int {
	return a * b
}


func allOperation(a, b int, op func(int, int) int) int {

	return op(a, b)
}


func main(){
	resultAdd := allOperation(20, 30, add)
	fmt.Println(resultAdd)
	resultSub := allOperation(30, 20, subtract)
	fmt.Println(resultSub)
	resultMultiply := allOperation(40, 20, multiply)
	fmt.Println(resultMultiply)



}
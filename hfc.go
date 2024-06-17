package main
import "fmt"



func add(a, b int) int{
	return a+b
}

func hfc(a, b int, operation func(int, int)int) int {
	return operation(a, b)

}


func main(){
	sum:= hfc(3, 4, add)

	fmt.Println("the total sum is", sum)
}
package main
import "fmt"


func multipleReturn(a, b int ) (int, int){
   quotient := a/b

   remainder := a%b

   return quotient, remainder
}



func main(){
	// quto  and remain are name given. we can put any name as variable
   quto, remain := multipleReturn(20, 4)
   fmt.Println("the answer of the quotient", quto, "the answer of the remainder", remain)
}
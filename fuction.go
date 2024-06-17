package main

import "fmt"

// func max(c1, c2 int) {
//     if (c1 > c2) {
//         fmt.Println("c1 is greater")
//     } else if (c1 == c2) {
//         fmt.Println("both are equal")
//     } else {
//         fmt.Println("c2 is greater")
//     }
// }

// func main() {
//     max(30, 33)
// }



func main(){
	const  a, b int = 40, 22
   
 max_min := func(){
		if(a >= b){
			fmt.Println("a is greater")
		} else if(a == b){
			fmt.Println("both are equal")
		} else {
			fmt.Println("b is greater")
		}
	}
	max_min()
}
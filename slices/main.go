package main

import (
"fmt"
"sort"

)

func main(){
	fmt.Println("this is the slicers")
	var fruits = []string {"apple", "banana", "pineapple"}
	fmt.Printf("type of fruitlist is %T\n", fruits)
	fmt.Println("fruit lists are", fruits)
	// add items
	fruits = append(fruits, "mango", "jack fruits")
	fmt.Println("new fruit lists are", fruits)
	// select the items between

	

	fmt.Println("the new lists are", fruits[2:])
	fruits= fruits[1:3]
	fmt.Println("the new lists are", fruits)

	score := make([]int, 4)
	score[0] = 222
	score[1] = 333
	score[2] = 444
	score[3] = 555


	fmt.Println("scores are", score)

	// memory allocation by using append. this is happend because of make

	score = append(score, 666, 757, 111, 234)
	fmt.Println("new scores are", score)
	// sorting integers

	sort.Ints(score)
	fmt.Println("this is the sorted value of score", score)

	// removing items from lists

	var index int = 2
	score = append(score[:index], score[index +1:]...)
	fmt.Println("final score after removal", score)
}
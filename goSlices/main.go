package main

import (
	"fmt"
)

func main() {
	fruitList := []string{"Apple", "Orange", "Banana"}

	fmt.Printf("Type of fruit is %T\n %v\n", fruitList, fruitList)

	fruitList = append(fruitList, "Pineapple", "Mango", "Tomato")

	fmt.Println(fruitList)

	for i, v := range fruitList {
		fmt.Println("The index and value of fruit list is\t",i,v)
	}

	fruitList = append(fruitList[1:]) // Include position of Idx No.1
	fmt.Println(fruitList)

	//fruitList = append(fruitList[1:4]) // total showing 4-1=3 index and not-include idx No.4
	//fmt.Println(fruitList)
	//
	//
	//highScore := make([]int, 4)
	//
	//highScore[0] = 343
	//highScore[1] = 231
	//highScore[2] = 421
	//highScore[3] = 123
	//highScore = append(highScore,453,532,723)
	//sort.Ints(highScore)
	//fmt.Println(highScore)
	//fmt.Println(sort.IntsAreSorted(highScore))

}

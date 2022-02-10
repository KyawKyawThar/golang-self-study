package main

import "fmt"

func main() {
	fruitList := [3]string{"Apple", "orange", "banana"}
	fmt.Println(fruitList, len(fruitList))

	var vgList [4]string

	vgList[0] = "Beans"
	vgList[1] = "potato"
	vgList[2] = "tomatoes"

	fmt.Println(vgList, len(vgList))
}


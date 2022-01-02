package main

import "fmt"

func main () {

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("value of actual pointer memory address is ",ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("adding pointer value is ",ptr)
	fmt.Println("value of actual pointer memory address is ",*ptr)
	fmt.Println("value of actual pointer memory address is ",&myNumber)
}

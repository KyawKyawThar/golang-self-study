package main

import "fmt"

func main() {
	n := []int{2, 4, 5, 4, 6, 2, 5, 1, 8, 7, 9, 12, 65}
	s := sum(n...)
	fmt.Println("All number", s)

	e1 := even(sum, n...)
	fmt.Println("Even number", e1)

	o1 := odd(sum, n...)
	fmt.Println("Odd number", o1)
}

//function (r receiver) identifier(parameters) (return (s)) {code...}
func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)

	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//Callback func even
func even(f func(ii ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

//Callback fun odd

func odd(f func(yi ...int) int, xi ...int) int {
	var ii []int

	for _, v := range xi {
		if v%2 != 0 {
			ii = append(ii, v)
		}
	}

	return f(ii...)
}

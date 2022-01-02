package main

import "fmt"

//Create and use an anonymous struct

func main() {
	s1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Meta",
		friends: map[string]int{
			"HL":  23,
			"AKK": 27,
		},
		favDrinks: []string{"Martini", "Vanilla", "Cappuccino"},
	}

	//fmt.Println(s1)
	//fmt.Println(s1.friends)
	//fmt.Println(s1.favDrinks)

	// map for range K,V

	for k, v := range s1.friends {
		fmt.Println("for Map", k, v)

	}

	// slice for range idx,v
	for i, val := range s1.favDrinks {
		fmt.Println("for Slice", i, val)

	}
}

package main

import "fmt"

type person struct {
	first    string
	last     string
	favFlour []string
}

func main() {
	p1 := person{
		first: "Meta",
		last:  "Block",
		favFlour: []string{
			"Chocolate", "Martini", "rum and coke",
		},
	}
	p2 := person{
		first:    "HL",
		last:     "Chain",
		favFlour: []string{"Strawberry", "Vanilla", "Cappuccino"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	//fmt.Println("This is a Map", m)
	for k, v := range m {
		fmt.Println("This is a key", k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlour {
			fmt.Println(i, val)
		}
		fmt.Println("--------------")
	}

}

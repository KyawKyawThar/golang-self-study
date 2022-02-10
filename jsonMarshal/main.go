package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p := person{
		First: "Meta",
		Last:  "Chain",
		Age:   25,
	}

	p1 := person{
		First: "Highest",
		Last:  "LeveL",
		Age:   28,
	}

	people := []person{p, p1}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

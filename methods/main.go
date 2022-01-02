package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

//function (r receiver) identifier(parameters) (return (s)) {code...}

func (s secretAgent) speak() { //method
	fmt.Println("I am ", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Highest",
			last:  "Level",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Meta",
			last:  "Chain",
		},
		ltk: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()
}

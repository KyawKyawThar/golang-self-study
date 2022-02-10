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
	fmt.Println("I am ", s.first, s.last, "the secret Agent speak")
}

func (p person) speak() { //method
	fmt.Println("I am ", p.first, p.last, "the person speak")
}

type human interface {
	speak()
}

func bar(h human) {

	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrr", h.(secretAgent).first)

	}
	fmt.Println("I was passed into bar", h)
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

	p1 := person{
		first: "Dr.",
		last:  "HL",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	//Polymorphism (poly=many and mor= changes)
	//and func bar take many types and called polymorphism
	bar(sa1)
	bar(sa2)
	bar(p1)
}

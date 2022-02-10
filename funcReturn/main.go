package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Printf("%T\n", s1)
	fmt.Println(s1)

	b := bar()()
	fmt.Printf("%T\n", b)
	fmt.Println(b)
}

func foo() string {
	return "Hello Gopher"
}

//Anonymous func
func bar() func() int64 {
	return func() int64 {
		return 251
	}
}

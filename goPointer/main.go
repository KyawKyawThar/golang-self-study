package main

import "fmt"

func main() {

	//myNumber := 23
	//var ptr = &myNumber
	//
	//fmt.Println("value of actual pointer memory address is ",ptr)
	//fmt.Println("value of actual pointer is ", *ptr)
	//
	//*ptr = *ptr + 2
	//fmt.Println("adding pointer value is ",ptr)
	//fmt.Println("value of actual pointer memory address is ",*ptr)
	//fmt.Println("value of actual pointer memory address is ",&myNumber)

	//a := 42
	//fmt.Println(a)
	//fmt.Println(&a) //&give you the address
	//
	//fmt.Printf("%T\n", a)
	//fmt.Printf("%T\n", &a)
	//
	//b := &a
	//fmt.Println("This is b", b)
	//fmt.Println(*b) // * give you the stored value at an address when you have the address
	//fmt.Println(*&a)
	//
	//*b = 45
	//fmt.Println(a)

	x := 0

	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)

	fmt.Println("x after", &x)
	fmt.Println("x after", x)

	p1 := person{name: "Kyaw Kyaw"}

	fmt.Println(p1)

	changeName(&p1)
	fmt.Println(p1)
}

func foo(y *int) {

	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)

	*y = 45
	fmt.Println("y after", *y)
	fmt.Println("y after", y)

}

type person struct {
	name string
}

func changeName(p *person) {
	p.name = "HighestLeveL"

	(*p).name = "AOA"
}

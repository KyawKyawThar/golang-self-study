package main

import (
	"fmt"
	"math"
)

//
////Create a type square
//type square struct {
//	length float64
//}
//
////Create a type circle
//type circle struct {
//	radius float64
//}
//
////create a method to each that calculates area and returns it
//func (s square) area() float64 {
//
//	return s.length * s.length
//}
//
////create a method to each that calculates area and returns it
//func (c circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}
//
////Create a type shapes which defines an interface as anything which has the area method
//type shape interface {
//	area() float64
//}
//
////Create a func info which takes type shape and then print the area
//func info(s shape) {
//	fmt.Println(s.area())
//}
//
////create the value of square
////create a value of circle
////use func info to print the area of square
////use func info to print the area of circle
//
//func main() {
//	circ := circle{radius: 422.123}
//	suq := square{length: 15}
//
//	info(circ)
//	info(suq)
//}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())

}

func main() {
	s1 := square{length: 45.11}
	c1 := circle{radius: 25}
	info(s1)
	info(c1)
}

package main

import (
	"fmt"
	"sort"
)

//appending Slice of Slice
func main() {
	x := []int{45, 33, 22, 11, 54}
	fmt.Println(x)

	x = append(x, 22, 33, 44)
	fmt.Println(x)

	y := []int{14, 5, 1996}
	x = append(x, y...) //variadic parameter
	fmt.Println(x)
	fmt.Println(sort.IntsAreSorted(x))
	sort.Ints(x)
	fmt.Println(x)
	fmt.Println(sort.IntsAreSorted(x))

	//Delete idx from Slices

	x = append(x[:3],x[8:]...) //8-3 = 5 //subtract 5 idx from slice
	fmt.Println(x)


	jb := []string{"James","Bond","Chocolate","martini"}
	fmt.Println(jb)
	mp := []string{"Miss","Moneypenny","Strawberry","Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb,mp}
	fmt.Println(xp)

}

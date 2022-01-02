package main

import "fmt"

func main() {
	m := map[string]int{"KyawKyawThar": 25, "Su": 22}

	fmt.Println(m)

	fmt.Println(m["KyawKyawThar"])
	fmt.Println(m["Su"])

	//v, ok := m["sdf"]
	//fmt.Println(v)
	//fmt.Println(ok)

	if v, ok := m["KyawKyawThar"]; ok { //if value this doesn't run
		//For obvious reasons this is called the “comma ok” idiom
		//check key-value include inside a map or not using comma ok syntax.
		fmt.Println("This is the if statement", v)
	}
}

//3263967......4075443
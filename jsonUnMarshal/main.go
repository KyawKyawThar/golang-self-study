package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	bs := []byte(`[{"First":"Meta","Last":"Chain","Age":25},{"First":"Highest","Last":"LeveL","Age":28}]`)

	//fmt.Printf("%T\n", bs)

	var people []Person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", people)

	for i, v := range people {
		fmt.Println("\n Person number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

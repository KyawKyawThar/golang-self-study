package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome from our crypto website")
	fmt.Println("Please rate our website between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	//strings.trimspace is important

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add plus one in rating", numRating + 1)
	}
}

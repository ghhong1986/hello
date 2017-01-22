package hello

import (
	"fmt"
	"strings"
)

func workwithstring() {
	str := "HI, I'm upper case"
	lower := strings.ToLower(str)
	fmt.Println(lower)

	if strings.Contains(lower, "case") {
		fmt.Println("Yes,exists")
	}
	fmt.Println("Characters 3-9:" + str[3:9])

	fmt.Println("First Five: " + str[:5])

	// split a string on a specific character or word
	sentence := "I'm a sentence made up of words"
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)

	// If you were splitting on whitespace, using Fields is better because
	// it will split on more than just the space, but all whitespace chars
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)
}

func workslice() {
	// initializes an empty array
	var empty []int

	// initialize array with values
	alphas := []string{"abc", "def", "ghi", "jkl"}

	// slices can not be modified, a new copy needs to be made
	// here is a example that appends elements to a slice
	empty = append(empty, 123)
	empty = append(empty, 456)
	fmt.Printf("%v \n", empty)

	// append multiple values
	alphas = append(alphas, "pqr", "stu")
	fmt.Printf("%v \n", alphas)

	// get length of a slice
	fmt.Println("Length: ", len(alphas))

	// retrieve a single element from slice
	fmt.Println(alphas[1])

	// retrieve a slice of a slice
	alpha2 := alphas[1:3]
	fmt.Println(alpha2)

	// check if element exists in array
	// there is no function to determine this
	// the only method is to iterate over the array
	// see elemExists func defined below
	if elemExists("def", alphas) {
		fmt.Println("Exists!")
	}
}

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

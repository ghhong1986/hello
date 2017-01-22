package hello

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	Echo("Bonjour tout le monde")

	// call the Say function which returns a string
	fmt.Println(Say("Gopher"))

	// test function with multiple return value
	q, r := Divide(11, 3)

	// this example uses Printf to format output, %v can be used for any type
	fmt.Printf("Quotient: %v, Remainder %v \n", q, r)
}

func TestStruct(t *testing.T) {
	num := 1

	// Go is not picky, conditionals don't require parentheses
	if num > 3 {
		fmt.Println("Many")
	}

	// Go is picky, "else" must be on the same line as closing if bracket
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Many")
	}

	// Switch statement takes conditionals and auto breaks
	switch {
	case num == 1:
		fmt.Println("One")
	case num == 2:
		fmt.Println("Two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}
}

func TestLoop(t *testing.T) {
	// creating an array of strings
	alphas := []string{"abc", "def", "ghi"}

	// standard for loop
	for i := 0; i < len(alphas); i++ {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	// if iterating over a array, this would be how you would actually do it
	// the standard loop would be used if uncommon step function
	// range returns two values, i -> index, val -> value
	for i, val := range alphas {
		fmt.Printf("%d: %s \n", i, val)
	}

	// if you did not care about the value and only wanted the index
	for i := range alphas {
		fmt.Println(i)
	}

	// if you did not care about the index and only the value, you use the _
	// which means don't save this value
	for _, val := range alphas {
		fmt.Println(val)
	}

	// how to use for like a while
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
}

func TestSortExample(t *testing.T) {
	sortExample()
}

func TestJsonExample(t *testing.T) {
	jsonExample()
}

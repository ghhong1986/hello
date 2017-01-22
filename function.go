package hello

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

func Echo(s string) {
	fmt.Println(s)
}

// Function with single return value
// the type of the return value is specifed after function declaration
func Say(s string) string {
	phrase := "Hello " + s
	return phrase
}

// Function with single named return value
// You can specify the return variable name, which initializes it
// the := notation is for new variables, and = for initial ones
// Also you do not need to include the variable in return statement
// it will return the current value of the variable at return
func Say2(s string) (phrase string) {
	phrase = "Hello " + s
	return
}

// Function with multiple parameters and return values
func Divide(x, y float64) (float64, float64) {
	q := math.Trunc(x / y)
	r := math.Mod(x, y)
	return q, r
}

// Function with multiple parameters and named return values
// If the types are the same you can specify the type once at the end
func Divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return
}

func sortExample() {
	abc := []string{"jkl", "ghi", "abc", "def"}
	nums := []int{4, 2, 12, 5, 1, 3}

	// show default slices out of order
	fmt.Println("Raw ABC : ", abc)
	fmt.Println("Raw Nums: ", nums)

	// sort using Strings() method
	sort.Strings(abc)
	fmt.Println("Sorted ABC:", abc)

	// sort using Ints() method
	sort.Ints(nums)
	fmt.Println("Sorted Nums:", nums)

	// reverse sort, need to cast abc as a StringSlice reverse and sort
	sort.Sort(sort.Reverse(sort.StringSlice(abc)))
	fmt.Println("Reverse ABC:", abc)

	//-------------------------------------------------
	// Sort Maps
	//-------------------------------------------------
	// sort map by key
	hash := map[string]int{
		"c": 3,
		"a": 1,
		"b": 2,
		"e": 5,
		"d": 4,
	}

	// view map to see its out of order
	for k, v := range hash {
		fmt.Printf("%s => %v\n", k, v)
	}

	// to sort, create array of keys and sort array
	keys := make([]string, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// display hash by looping ordered keys
	for i := range keys {
		fmt.Printf("%s => %v\n", keys[i], hash[keys[i]])
	}
}

type Animal struct {
	Name  string
	Order string
}

func jsonExample() {
	var jsonBlob = []byte(`[  
    {"Name": "Platypus", "Order": "Monotremata"},  
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}  
]`)
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v \n", animals)
}

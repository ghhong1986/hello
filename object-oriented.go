package hello

import "fmt"

// define Dog object type
type Dog struct {
	Name  string
	Color string
}

// Method for object specify the object the refer to and then
// the method name and rest of normal function definition
func (d Dog) Call() {
	fmt.Printf("Here comes a %s dog, %s.\n", d.Color, d.Name)
}

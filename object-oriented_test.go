package hello

import (
	"testing"
)

func TestObjectOriented(t *testing.T) {
	Spot := Dog{Name: "Spot", Color: "brown"}

	// or, try it like this...
	var Rover Dog
	Rover.Name = "Rover"
	Rover.Color = "light blondish with some dark patches and very scruffy"

	// call object method
	Spot.Call()
	Rover.Call()
}

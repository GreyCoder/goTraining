package main

import (
	"fmt"
)
func main() {
	fruits := [4]string{"Banana", "Orange", "Pineapple", "Strawberry"}

	// Use a 'classic' for loop  to print out each fruit in the array, and the
	// corresponding index.
	for ndx := 0; ndx < len(fruits); ndx++ {
		fmt.Printf("%v: is %s", ndx, fruits[ndx])
	}
	// Use the range keyword to loop over the same array of fruits, again printing
	// out the fruit and the corresponding index.
	for i, n := range fruits {
		fmt.Printf("%v: is %s", i, n)
	}
	// Using the keyword `continue`, skip every other fruit (even ones)

	for i, n := range fruits {
		if i % 2 != 0 {
			continue
		}
		fmt.Printf("%v: is %s", i, n)

	}
}
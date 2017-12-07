package example2

import (
	"fmt"
)
// declare a struct called Movie
// add the following fields:
// - Title (string)
// - Released (bool)
// - Length (int)

type Movie struct {
	Title string
	Released bool
	Length int
}


func Example2() {
	// declare a variable called "movie" of type "Movie"

	movie := Movie{}

	// Set the Title to "Wizard of Oz"
	// Set the Released variable to "true"
	// Set Length to 125

	movie.Title = "Wizard of Oz"
	movie.Released = true
	movie.Length = 125

	// Print the value of "movie" out
	// hint: you can use fmt.Println(movie)
	fmt.Println(movie)
}
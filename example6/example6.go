package main

import "fmt"

func main() {
	x := "George"
	printValue(&x)
	fmt.Println(x)
}

func printValue(s *string) {
	// print the pointer value
	// print the string value
	// change the string value
	fmt.Println(s)
	fmt.Println(*s)
	*s = "Ringo"
}

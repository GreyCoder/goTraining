package main

import "fmt"

func main() {
	parents := []string{"Carol", "Mike"}
	kids := []string{"Marcia", "Jan", "Cindy", "Greg", "Peter", "Bobby"}

	// Create a new slice called family by joining the parents and kids slice together
	family := make([]string, len(parents))
	copy(family, parents)
	family = append (family, kids...)

	family2 := []string{}
	family2 = append(family2, parents...)
	family2 = append(family2, kids...)

	family3 := append(parents, kids...)

	fmt.Println(family)
	fmt.Println(family2)
	fmt.Println(family3)

	// Fix the following bugs
	extras := make([]string, 1, 1)
	extras[0] = "Alice"
	fmt.Println(extras)
}
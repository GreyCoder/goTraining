package main
import "fmt"
func main() {

	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	// Loop through the map and print out the key and the value
	for beatle := range beatles{
		fmt.Printf("Key is %+v and value is %s\n", beatle, beatles[beatle])
	}

	// Look up the key `Bob` and detect that it wasn't found by printing `not found`
	bob, ok := beatles["Bob"]
	if ok {
		fmt.Println(bob)
	}else{
		fmt.Println("not found")
	}
}
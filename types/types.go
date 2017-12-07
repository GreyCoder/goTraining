package main

import (
	"fmt"
)

func main() {
	var byteType byte
	fmt.Println(&byteType)
	for i := 0; i < 10; i++ {
		fmt.Printf("%x\n", byteType)
		byteType++
	}

}
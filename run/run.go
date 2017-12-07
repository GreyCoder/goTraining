package main

import (
	"fmt"
	"goTraining/example1"
	"goTraining/example2"
	"github.com/gopherguides/foo"
)

func main() {

	user := foo.User{
		First: "John",
		Last: "Doe",

	}

//	fmt.Println(user.First)
//	fmt.Println(user.Last)
	fmt.Println(user)
	example1.Example1()
	example2.Example2()
}
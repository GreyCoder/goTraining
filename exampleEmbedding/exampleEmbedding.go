package main

import (
	"errors"
	"fmt"
	"strings"
)

// Given the following two structs, embed an email into the Invite struct
// to make the `main` program below work.
type Invite struct {
	ID int
	Email
}

type Email struct {
	DisplayName string
	Username    string
	Domain      string
}

// Create a method on Email called `Address` that takes no arguments
// It should return a string in the format:
// `DisplayName <Username@Domain>`
// You can use `fmt.Sprintf` to create the return string

type Addressable interface{
	Address() string
}

func main() {
	invite := Invite{
		ID: 1,
		Email: Email{
			DisplayName: "John Smith",
			Username:    "jsmith",
			Domain:      "yahoo.com",
		},
	}
	if err := SendEmail(invite); err != nil {
		fmt.Printf("error sending email: %s")
		return
	}
	fmt.Printf("sent email to : %s", invite.Address())
}

func SendEmail(a Addressable) error {
	// Do some work here
	address := a.Address()
	if address == "" {
		return errors.New("no email address provided")
	}
	return nil
}


func (i Invite) Address() string {
	EmailAddress := []string{i.Email.Username, "@", i.Email.Domain}
	return strings.Join(EmailAddress, "")
}
package main

import "fmt"

//Declare a struct
type customer struct {
	fName    string
	lName    string
	userName string
	password string
	email    string
	phone    int
	address  string
}

func (c customer) userCredentials() (string, string) {
	return c.userName, c.password
}

func (c customer) userAddress() string {
	return c.address
}
func (c customer) allUserInformation() {
	fmt.Printf("First Name: %v\n"+
		"Last Name: %v\n"+
		"Username: %v\n"+
		"Password: %v\n"+
		"Email: %v\n"+
		"Phone: %v\n"+
		"Address: %v\n\n", c.fName, c.lName, c.userName, c.password, c.email, c.phone, c.address)
}

//Create Methods

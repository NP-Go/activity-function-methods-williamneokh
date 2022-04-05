package main

import "fmt"

func main() {
	//Declare new item customer
	customer1 := customer{}
	customer1 = customer{
		fName:    "Micheal",
		lName:    "Jordan",
		userName: "MJ2020",
		password: "1234567",
		email:    "MJ2020@gmail.com",
		phone:    12345678,
		address:  "18227 Capstan Greens Road Cornelius, NC 28031",
	}

	customer1.allUserInformation()

	cUsername, cPassword := customer1.userCredentials()
	fmt.Printf("Username: %v\nPassword: %v\n", cUsername, cPassword)

	cUserAddress := customer1.userAddress()
	fmt.Printf("Customer Address: %v\n", cUserAddress)

	//Insert Code here
}

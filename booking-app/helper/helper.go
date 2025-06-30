package helper

import "fmt"

// we can export a variable or function by make its first letter capital

func GreetUser(conferenceName string, conferenceTickets int, bookedTickets uint) {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Print("Get your ticket here to attend\n")
	fmt.Printf("Total tickets: %d, Available tickets: %d\n", conferenceTickets, conferenceTickets-int(bookedTickets))
	fmt.Printf("Type of conferenceName: %T, Type of conferenceTickets: %T, Type of bookedTickets: %T\n",
		conferenceName, conferenceTickets, bookedTickets)
}

func GetUserInput() (string, string, string, uint) {
	var firstName string // defining type of the variable
	var lastName string
	var email string
	var userTickets uint
	// taking user input
	fmt.Println("Enter your first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email Name: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanf("%d", &userTickets)

	return firstName, lastName, email, userTickets
}

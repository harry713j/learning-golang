package main // everything in go is package

import "fmt" // fmt package Print() function belong here

// main() is the entrypoint
func main() {
	const conferenceName = "Dev Conference"
	const conferenceTickets = 50
	bookedTickets := 0 // go syntatical sugar (only for var)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Print("Get your ticket here to attend\n")
	fmt.Printf("Total tickets: %d, Available tickets: %d\n", conferenceTickets, conferenceTickets-bookedTickets)
	fmt.Printf("Type of conferenceName: %T, Type of conferenceTickets: %T, Type of bookedTickets: %T\n",
		conferenceName, conferenceTickets, bookedTickets)

	var firstName string // defining type of the variable
	var lastName string
	var email string
	var userTickets int
	// taking user input
	fmt.Println("Enter your first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email Name: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanf("%d", &userTickets)

	// fmt.Println(&bookedTickets) // address of the variable (pointer)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get confirmation email at %v.\n",
		firstName, lastName, userTickets, email)

}

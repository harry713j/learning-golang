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

	var userName string // defining type of the variable
	var userTickets int
	// taking user input

	userName = "Heaven"
	userTickets = 2

	fmt.Printf("%v has booked %v tickets\n", userName, userTickets)

}

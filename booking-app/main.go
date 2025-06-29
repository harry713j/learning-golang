package main // everything in go is package

import "fmt" // fmt package Print() function belong here

// main() is the entrypoint
func main() {
	conferenceName := "Dev Conference"
	const conferenceTickets = 50
	var bookedTickets uint = 0 // go syntatical sugar (only for var)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Print("Get your ticket here to attend\n")
	fmt.Printf("Total tickets: %d, Available tickets: %d\n", conferenceTickets, conferenceTickets-bookedTickets)
	fmt.Printf("Type of conferenceName: %T, Type of conferenceTickets: %T, Type of bookedTickets: %T\n",
		conferenceName, conferenceTickets, bookedTickets)

	// fmt.Println(&bookedTickets) // address of the variable (pointer)

	// array for storing the booking details -> array declaration
	// var bookingDetails = [50]string{"Aria", "Ben", "Catty"} // initilize with values
	// var bookingDetails = [50]string{}
	// var bookingDetails [50]string // declare a empty array

	// slices is the abstraction of the array, but it is dynamic in size
	// var bookingDetails = []string{}
	// bookingDetails := []string{}
	var bookingDetails []string

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

	bookedTickets = bookedTickets + userTickets
	// bookingDetails[0] = firstName + " " + lastName // string concatination
	bookingDetails = append(bookingDetails, firstName+" "+lastName) // adding value to the slice

	fmt.Println("Whole array ", bookingDetails)
	fmt.Println("First element: ", bookingDetails[0])
	fmt.Printf("Type of array: %T\n", bookingDetails)
	fmt.Printf("Length of array: %v\n", len(bookingDetails))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get confirmation email at %v.\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("Total tickets: %d, Available tickets: %d\n", conferenceTickets, conferenceTickets-bookedTickets)

}

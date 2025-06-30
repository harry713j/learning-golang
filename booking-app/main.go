package main // everything in go is package

// importing single package
// import "fmt" // fmt package Print() function belong here

// importing multiple package
import (
	"fmt"
	"strings"
)

// main() is the entrypoint
func main() {
	conferenceName := "Dev Conference"
	const conferenceTickets = 50
	var bookedTickets uint = 0 // go syntatical sugar (only for var)

	greetUser(conferenceName, conferenceTickets, bookedTickets)

	// fmt.Println(&bookedTickets) // address of the variable (pointer)

	// array for storing the booking details -> array declaration
	// var bookingDetails = [50]string{"Aria", "Ben", "Catty"} // initilize with values
	// var bookingDetails = [50]string{}
	// var bookingDetails [50]string // declare a empty array

	// slices is the abstraction of the array, but it is dynamic in size
	// var bookingDetails = []string{}
	// bookingDetails := []string{}
	var bookingDetails []string

	// only for loop is exist in GO

	// for i := 0; i < 5; i++ {
	// 	fmt.Print(i, " ")
	// }

	// // for i := range 5{
	// // 	fmt.Print(i, " ")
	// // }

	// fmt.Println("")

	// infinite for-loop syntax for{} or for true {}

	for bookedTickets < conferenceTickets && len(bookingDetails) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUser(firstName, lastName, email, userTickets, conferenceTickets, bookedTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookedTickets = bookedTickets + userTickets
			// fmt.Println("Whole array ", bookingDetails)
			// fmt.Println("First element: ", bookingDetails[0])
			// fmt.Printf("Type of array: %T\n", bookingDetails)
			// fmt.Printf("Length of array: %v\n", len(bookingDetails))

			bookTicket(bookedTickets, userTickets, bookingDetails, firstName, lastName, email, conferenceTickets)

			firstNames := filterFirstName(bookingDetails)
			fmt.Println("All the name of people who booked the tickets : ", firstNames)

		} else {
			fmt.Printf("Please provide valid details to book the tickets")
			continue
		}

		// // if all the tickets are finished booked then we will break out of the loop
		// if bookedTickets == conferenceTickets {
		// 	fmt.Println("All the ticket for the", conferenceName, "is booked!")
		// 	break
		// }

	}

}

// func greetUser() {
// 	fmt.Println("Hello there")
// }

func greetUser(conferenceName string, conferenceTickets int, bookedTickets uint) {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Print("Get your ticket here to attend\n")
	fmt.Printf("Total tickets: %d, Available tickets: %d\n", conferenceTickets, conferenceTickets-int(bookedTickets))
	fmt.Printf("Type of conferenceName: %T, Type of conferenceTickets: %T, Type of bookedTickets: %T\n",
		conferenceName, conferenceTickets, bookedTickets)
}

func filterFirstName(bookings []string) []string {

	firstNames := []string{}
	// for-each loop show only the firstName
	for _, element := range bookings { // _ used for placeholder , variable we want to ignore using
		var names = strings.Fields(element) // Fields() split the string
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

// functions returning list of values
func validateUser(firstName string, lastName string, email string, userTickets uint, conferenceTickets int, bookedTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= uint((conferenceTickets-int(bookedTickets)))

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

func bookTicket(bookedTickets uint, userTickets uint, bookingDetails []string, firstName string, lastName string, email string, conferenceTickets int) {

	// bookingDetails[0] = firstName + " " + lastName // string concatination
	bookingDetails = append(bookingDetails, firstName+" "+lastName) // adding value to the slice

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get confirmation email at %v.\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("Total tickets: %d, Available tickets: %d\n", conferenceTickets, conferenceTickets-int(bookedTickets))
}

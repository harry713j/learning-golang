package main // everything in go is package

// importing single package
// import "fmt" // fmt package Print() function belong here

// importing multiple package
import (
	"booking-app/helper" // importing from our package
	"fmt"
)

// structure definition
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// package level variables, we can't use := for type inference for the package level variables
var conferenceName = "Dev Conference"

const conferenceTickets = 50

var bookedTickets uint = 0
var bookingDetails = make([]UserData, 0) // list of map with size 0, which will grow dynamically

// main() is the entrypoint
func main() {

	helper.GreetUser(conferenceName, conferenceTickets, bookedTickets)

	for bookedTickets < conferenceTickets && len(bookingDetails) < 50 {

		firstName, lastName, email, userTickets := helper.GetUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUser(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookedTickets = bookedTickets + userTickets

			bookingDetails = bookTicket(userTickets, firstName, lastName, email)

			firstNames := filterFirstName()
			fmt.Println("All the name of people who booked the tickets : ", firstNames)

			fmt.Println("Booking Details: ", bookingDetails)

		} else {
			fmt.Printf("Please provide valid details to book the tickets\n")

			if !isValidName {
				fmt.Println("First Name and Last Name should contain atleast 2 characters")
			}

			if !isValidEmail {
				fmt.Println("Email should be valid and contain @")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets should not be greater than remaining tickets")
			}

			continue
		}

		fmt.Println()

	}

}

func filterFirstName() []string {

	firstNames := []string{}
	// for-each loop show only the firstName
	for _, element := range bookingDetails { // _ used for placeholder , variable we want to ignore using
		firstNames = append(firstNames, element.firstName)
	}

	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) []UserData {

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookingDetails = append(bookingDetails, userData) // adding value to the slice

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get confirmation email at %v.\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("Total tickets: %d, Available tickets: %d\n", conferenceTickets, conferenceTickets-int(bookedTickets))

	return bookingDetails
}

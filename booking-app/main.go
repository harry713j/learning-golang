package main // everything in go is package

import (
	"booking-app/helper" // importing from our package
	"fmt"
	"sync"
	"time"
)

// structure definition
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

const conferenceTickets = 50

var conferenceName = "Dev Conference"
var bookedTickets uint = 0
var bookingDetails = make([]UserData, 0) // list of map with size 0, which will grow dynamically
var waitGroup = sync.WaitGroup{}         // will force main thread to wait for the other thread to complete

func main() {

	helper.GreetUser(conferenceName, conferenceTickets, bookedTickets)

	for bookedTickets < conferenceTickets && len(bookingDetails) < 50 {

		firstName, lastName, email, userTickets := helper.GetUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUser(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookedTickets = bookedTickets + userTickets

			bookingDetails = bookTicket(userTickets, firstName, lastName, email)

			waitGroup.Add(1) // add number of goroutine to wait group

			// this will block the normal flow of the program, so we can handle it with
			// goroutine or light-weight thread which will run it in background
			// sendTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email) // go keyword for run it with another go routine

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
	waitGroup.Wait() // wait till wait group counter is 0
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

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// simulating a delay (delay the execution of main thread or goroutine)
	time.Sleep(15 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("######################")
	waitGroup.Done() // remove the goroutine from the wait group
}

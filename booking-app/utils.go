package main // belong to same package

import "strings"

/*
****** SCOPE: Package Level ********
Varibles and functions defined outside any functions can be accessed in all other files that defined within same package
*/

// functions returning list of values
func validateUser(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= uint((conferenceTickets-int(bookedTickets)))

	return isValidName, isValidEmail, isValidTicketNumber
}

// to compile and run whole package `go run package-location`

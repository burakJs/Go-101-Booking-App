package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicketCount uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isGreaterThenRemainingTicket := remainingTicketCount > userTicket
	isValidTicketNumber := isGreaterThenRemainingTicket && userTicket > 0

	return isValidName, isValidEmail, isValidTicketNumber
}

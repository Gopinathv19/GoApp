package helper

import(
	"strings"
)

func IsValidateUserInputs(firstName string, secondName string, ticketsNeedtoBeBooked uint,  emailAddress string, phoneNumber string,remainingTicketsAvailable uint) (bool, bool, bool, bool) {
	isValidUserName := len(firstName) > 2 && len(secondName) > 2
	isValidTicketNumber := ticketsNeedtoBeBooked > 0 && ticketsNeedtoBeBooked <= remainingTicketsAvailable
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidPhnNumber := len(phoneNumber) == 10
	return isValidUserName, isValidTicketNumber, isValidEmail, isValidPhnNumber
}
package main

import (
	"BookingApplication/helper"
	"fmt"
	"time"
)

// Package level variables
const ticketsAvailable uint = 50

var remainingTicketsAvailable uint = 50
var ticketsSaled uint = 0

const travelsName string = "G travels"

type UserData struct {
	firstName             string
	secondName            string
	email                 string
	phoneNumber           string
	ticketsNeedtoBeBooked uint
}

func main() {

	GreetUsers()
	var firstName string
	var secondName string
	var emailAddress string
	var phoneNumber string
	var ticketsNeedtoBeBooked uint
	var fullNames []string

	for {
		// getting the user details by the getUserDetailsFunction

		firstName, secondName, emailAddress, phoneNumber, ticketsNeedtoBeBooked, fullNames = getUserDetails(firstName, secondName, emailAddress, phoneNumber, ticketsNeedtoBeBooked, fullNames)

		// validating the userInputs using the validateUserInputs function

		isValidUserName, isValidTicketNumber, isValidEmail, isValidPhnNumber := helper.IsValidateUserInputs(firstName, secondName, ticketsNeedtoBeBooked, emailAddress, phoneNumber, remainingTicketsAvailable)

		if isValidUserName && isValidEmail && isValidPhnNumber && isValidTicketNumber {
			ticketsSaled += ticketsNeedtoBeBooked
			remainingTicketsAvailable -= ticketsSaled
			userDetailsInStruct := userBookingConfirmation(firstName, secondName, emailAddress, phoneNumber, ticketsNeedtoBeBooked)
			userDetails := putUserDetailsInSlices(userDetailsInStruct)
			// to improve the ticket booking concurrently to the many users we use go key word  to increase the concurrency
			 
			getTickets(firstName, secondName, ticketsNeedtoBeBooked, userDetails)
			go sendEmail(emailAddress, ticketsNeedtoBeBooked)
			 

		} else {
			if !isValidEmail {
				fmt.Printf("Your email format is not good check the email:%v\n", emailAddress)
			}
			if !isValidUserName {
				fmt.Printf("Your username is not in the format check the userName:%v %v \n", firstName, secondName)
			}
			if !isValidPhnNumber {
				fmt.Printf("Your phone number not in the correct format check the phone number:%v\n", phoneNumber)
			}
			if !isValidTicketNumber {
				fmt.Printf("We have only %v this much amount of the tickets so you cant book more than this so book below this not this much %v \n", remainingTicketsAvailable, ticketsNeedtoBeBooked)
			}

			continue
		}
	}

}

func GreetUsers() {
	fmt.Printf("Welcome to the %v make your destiny here\n", travelsName)
	fmt.Printf("Nothing else you need to book and enjoy the journey\n")
	fmt.Printf("We have %v of tickets to chennai out of this we have  %v of tickets\n", ticketsAvailable, remainingTicketsAvailable)
}

func getFirstNames(userDetails []UserData) []string {

	callNames := []string{}
	for _, user := range userDetails {
		firstName := user.firstName
		callNames = append(callNames, firstName)

	}
	return callNames
}

func getUserDetails(firstName string, secondName string, emailAddress string, phoneNumber string, ticketsNeedtoBeBooked uint, fullNames []string) (string, string, string, string, uint, []string) {
	fmt.Printf("Enter the First Name:\n")
	fmt.Scan(&firstName)
	fmt.Printf("Enter the Second Name:\n")
	fmt.Scan(&secondName)
	fullNames = append(fullNames, firstName+" "+secondName)
	fmt.Printf("Enter the email Address:\n")
	fmt.Scan(&emailAddress)
	fmt.Printf("Enter the phone number:\n")
	fmt.Scan(&phoneNumber)
	fmt.Printf("Enter the tickets No tickets to be Booked:\n")
	fmt.Scan(&ticketsNeedtoBeBooked)
	return firstName, secondName, emailAddress, phoneNumber, ticketsNeedtoBeBooked, fullNames
}

func userBookingConfirmation(firstName string, secondName string, email string, phoneNumber string, ticketsNeedtoBeBooked uint) UserData {
	var userDetails = UserData{
		firstName:             firstName,
		secondName:            secondName,
		email:                 email,
		phoneNumber:           phoneNumber,
		ticketsNeedtoBeBooked: ticketsNeedtoBeBooked,
	}

	return userDetails
}

func putUserDetailsInSlices(userDetailsInMap UserData) []UserData {
	var UserDetailsInSlices = make([]UserData, 0)
	UserDetailsInSlices = append(UserDetailsInSlices, userDetailsInMap)
	return UserDetailsInSlices
}

func getTickets(firstName string, secondName string, ticketsNeedtoBeBooked uint, userDetails []UserData) {
	fmt.Printf("Thank you %v %v you have booked tickets and the count is %v \n", firstName, secondName, ticketsNeedtoBeBooked)
	fmt.Printf("The remaining tickets available is %v\n", remainingTicketsAvailable)
	firstNames := getFirstNames(userDetails)
	fmt.Printf("So for booked persons first names are %v\n", firstNames)
	fmt.Printf("The user details of the customer is %v\n", userDetails)
}

func sendEmail(emailAddress string, ticketsNeedtoBeBooked uint) {
	time.Sleep(10 * time.Second)
	fmt.Println("************************************************")
	fmt.Printf("The tickets have sended to your email : %v \n and the total no of tickets you have booked is %v\n", emailAddress, ticketsNeedtoBeBooked)
	fmt.Println("************************************************")
}

/*
*********	Glossory Of the Application   *********

	ticketsAvailable:::Total tickets before booking
	remainingTicketsAvailable:::To dyanamic value change after each booking
	ticketsSaled::: Have the total of the tickets been sold this value is been minused to remainingTicketsAvailabe to get the current ticket count

	var firstName string   ::: firstname of the user
    var secondName string  ::: secondname of the user
	var emailAddress string  ::: emailAddress of the user
	var phoneNumber int    ::: phoneNumber of the user
	var ticketsNeedtoBeBooked uint  ::: count of the tickets the user need to book the tickets
	var fullNames[] string   ::: it is used to put the fullname of the user insteam of concating the first and the last name of the user

	callNames :::  the call name are the slice used to take the first name of the user for some naming convention
	fieldedName ::: the fieldName is used to get the first of the name that has been done by the string fields in built function

*/

/*

 ********  Syntax Explantion that has been used in this program    *********

	Array Operations
	fmt.Print(fullNames)
	fmt.Print(fullNames[0])
	fmt.Print(len(fullNames))

	In order to create the slice we need to declare them as the  var nameofTheSlice[] string if you clearly see there we dont
	mention the size of the array that is the slice

	to append the elements in the slice the function declaration is nameOfTheVariable=append(nameOfTheVariable,value);

	Slice Opperation
	fmt.Print(fullNames)
	fmt.Print(fullNames[0])
	fmt.Print(len(fullNames))

	In go we simply use only one loop
	the one and only loop is for
	the declaration is for{ }

	to declare the for each loop we will do the following steps

	the first step is to declare the for loop

	for index,NameofTheVar := nameOfTheArray{
	    }



   we can also limit the infinite for loop by the conditions

   the procedure is

   for remainingTickets>0{
   //statements
   }

   for len(bookings) <50{
   // statements
   }

   or we can limit that by the combination of the both

   for remainingTickets >0 && len(booikings)<50{
   // statements
   }



   when we need to return the value in the function we follow this type of procedure like this

   func nameoftheFunction()return_type{


   }


   the another datatype that is provided by the goland is called map the map ds is used to store the data in the key value pairs the
   the key value pairs are used in many way so it is important to know about that ds

   the declaration for that map is

   var map_name=make(map[string]string)
   [string]-> type for the key
   string -> type for the value

   time.Sleep -> is the function that allows to stop the current thread execution and makes to sleep for some seconds

*/

/*
	see here we using the index for the range and the
	index variable is not being used so we likely do the _ to implicitly not using them
	 for _,nameOfTheVar:=nameofTheArray{

	 }

*/

/*

************ trash code ***************


			// so for booked person details
			firstNames := getFirstNames(userDetails)
			fmt.Printf("So for booked persons  first names are %v \n", firstNames)



*/

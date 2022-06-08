package main

import (
	"fmt"
	"sync"
	"time"
)

const maxTicketCount int = 50

var conferenceName = "Go Conference"
var remainingTicketCount uint = 50
var bookings = make([]UserData, 0)
var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	// for {

	firstName, lastName, email, userTicketCount := getUserInput()

	noTicketsRemaining := remainingTicketCount == 0
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTicketCount, remainingTicketCount)

	if isValidEmail && isValidName && isValidTicketNumber {
		bookTicket(userTicketCount, firstName, lastName, email)
		user := UserData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTicketCount,
		}
		wg.Add(1)
		go sendTicket(user)

		firstNames := getFirstNames()

		fmt.Printf("The first names of bookings are : %v\n\n", firstNames)

		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {

		if !isValidName {
			fmt.Println("First name or last name is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesnt't contains '@' sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets is invalid")
		}
		fmt.Println("Your input datas are invalid, try again")
		// continue
	}
	wg.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still availiable\n", maxTicketCount, remainingTicketCount)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicketCount uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicketCount)

	return firstName, lastName, email, userTicketCount
}

func bookTicket(userTicketCount uint, firstName string, lastName string, email string) {
	remainingTicketCount -= userTicketCount

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicketCount,
	}

	bookings = append(bookings, userData)
	fmt.Printf("\nList of booking : %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets\n", firstName, lastName, userTicketCount)
	fmt.Printf("You will receive a confirmation email at %v\n", email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTicketCount, conferenceName)
}

func sendTicket(user UserData) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v ", user.numberOfTickets, user.firstName, user.lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n %v\n to email address %v\n", ticket, user.email)
	fmt.Println("#####################")
	wg.Done()
}

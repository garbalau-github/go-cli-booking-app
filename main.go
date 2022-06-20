package main

import (
	"fmt"
	"sync"
	"time"
	"booking-app/helper"
)

const conferenceTickets uint = 50
var remainingTickets uint = 50;
var conferenceName string = "Go Conference"
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var waitGroupFunctions = sync.WaitGroup{}

func main() {

	greetUser()

	var firstName, lastName, email, userTickets = getUserInput()
	var isValidName, isValidEmail, isValidTicketNumber = helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {	

		bookTicket(userTickets, firstName, lastName, email)

		waitGroupFunctions.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		var firstNames = printFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		
		if remainingTickets == 0  {
			fmt.Println("Our conference is booked out. Come back next year!")
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if (!isValidTicketNumber) {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}

	waitGroupFunctions.Wait()

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v this are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func printFirstNames() []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you need? ")
	fmt.Scan(&userTickets)
	
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v, for boooking %v tickets! You will receive a confirmation email soon, check your inbox\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName);
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Simulate generation
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending ticket: \n %v to email address %v\n", ticket, email)
	fmt.Println("########")

	waitGroupFunctions.Done()
}
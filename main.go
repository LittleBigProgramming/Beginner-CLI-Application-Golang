package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			var userData = UserData {
				firstName: firstName,
				lastName: lastName,
				email: email,
				numberOfTickets: userTickets,
			}
			bookTicket(userData)
		
			wg.Add(1)
			go sendTicket(userData)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				return
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	wg.Wait()
}

func printFirstNames() []string {
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
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userData UserData) {
	remainingTickets = remainingTickets - userData.numberOfTickets
	
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userData.firstName, userData.lastName, userData.numberOfTickets, userData.email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// simulate sending a ticket wich may be a long running task
func sendTicket(userData UserData) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userData.numberOfTickets, userData.firstName, userData.lastName)

	fmt.Println("########")
	fmt.Printf("Sending ticket:\n %v \n to email address %v", ticket, userData.email)
	fmt.Println("########")

	wg.Done()
}

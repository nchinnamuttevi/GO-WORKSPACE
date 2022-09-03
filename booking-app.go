package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v remainingTicketsare still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("enter user first name:")
		fmt.Scan(&firstName)

		fmt.Println("enter user last name:")
		fmt.Scan(&lastName)

		fmt.Println("enter user email:")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets:")
		fmt.Scan(&userTickets)

		// user valiudation
		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		//logic for registration
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you User: %v %v for booking %v tickets. You will recive a conformation email sent to %v .\n", firstName, lastName, userTickets, email)
			fmt.Printf(" %v has %v remaining tickets left\n", conferenceName, remainingTickets)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The confirmed attendees first names are:%v\n", firstNames)
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is fully booked please come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name is not long enough")
			}
			if !isValidEmail {
				fmt.Println("Email is not valid, needs @ sign")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining you can't book %v tickets\n", remainingTickets, userTickets)
			}
		}

	}
}

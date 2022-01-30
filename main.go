package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTicket uint = 50
	var remainingTicket uint = conferenceTicket
	var bookings []string

	//print the type of variable
	fmt.Printf("conferenceTicket is type of %T, remainingTicket is type of %T\n", conferenceTicket, remainingTicket)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTicket, "tickets and", remainingTicket, "are still available")
	fmt.Println("Book your ticket to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var numberOfTickets uint

		fmt.Println("Enter your first name ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Println("How many tickets you want")
		fmt.Scan(&numberOfTickets)

		if numberOfTickets >= remainingTicket {
			fmt.Printf("we have only %v tickets, so you can't book %v tickets\n please try again \n", remainingTicket, numberOfTickets)
			continue
		}
		remainingTicket = remainingTicket - numberOfTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, numberOfTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("first names of bookings are : %v\n", firstNames)

		if remainingTicket == 0 {
			//end program
			fmt.Println("Our conferebce is booked out, Thank you for overehelming response")
			break
		}
	}
}

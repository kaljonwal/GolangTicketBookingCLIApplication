package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket = conferenceTicket

	//print the type of variable
	fmt.Printf("conferenceTicket is type of %T, remainingTicket is type of %T\n", conferenceTicket, remainingTicket)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTicket, "tickets and", remainingTicket, "are still available")
	fmt.Println("Book your ticket to attend")

	var userName string
	var numberOfTickets int
	userName = "kalyan"
	numberOfTickets = 2
	fmt.Printf("%v booked the %v tickets\n", userName, numberOfTickets)
}

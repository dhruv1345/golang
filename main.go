package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferencetickets = 50
	var remainingTickets uint = 50

	fmt.Printf("\nCoferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n", conferencetickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Println("We have total of", conferencetickets, "tickets and ", remainingTickets, " are still available")
	fmt.Print("\nGet your tickets here !!!\n")

	var userName string
	var usertickets int

	userName = "Tom"
	usertickets = 2

	fmt.Printf("User %v booked %v tickets.\n\n", userName, usertickets)

}

package main

import "fmt"

func main (){
	var confrenceName = "Go Conference"
	const conferenceTickets = 50 // This values do not change
	var remainingTickets = 50

	fmt.Printf(" Welcome to %v booking app \n", confrenceName)
	fmt.Printf(" We have a total of %v and %v are available \n" conferenceTickets, remainingTickets)
	fmt.Printf(" Get your tickets here ")

	
}

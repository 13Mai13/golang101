package main

import "fmt"

func main (){
	var confrenceName = "Go Conference"
	const conferenceTickets = 50 // This values do not change
	var remainingTickets = 50

	fmt.Println(" Welcome to", confrenceName, "booking app ")
	fmt.Println(" We have a total of", conferenceTickets, "and ", remainingTickets, "are available")
	fmt.Println(" Get your tickets here ")

	
}

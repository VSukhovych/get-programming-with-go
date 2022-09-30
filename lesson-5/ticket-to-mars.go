package main

import (
	"fmt"
	"math/rand"
)

type ticket struct {
	company         string
	tripDays        int
	twoWays         bool
	priceInMillions int
}

func main() {
	tickets := getRandomTickets(10)
	printTickets(tickets)
}

func getRandomTickets(numberOftickets uint) []ticket {
	tickets := make([]ticket, numberOftickets)
	distanceToMarsKm := 62100000
	basicSpeedKmS := 16
	basicPriceMillions := 36

	for i := range tickets {
		company := rand.Intn(2)
		switch company {
		case 0:
			tickets[i].company = "SpaceX"
		case 1:
			tickets[i].company = "Space Adventures"
		case 2:
			tickets[i].company = "Virgin Galactic"
		}

		extraSpeedKmS := rand.Intn(14) + 1
		speedKmS := basicSpeedKmS + extraSpeedKmS

		tripDays := distanceToMarsKm / speedKmS / 60 / 60 / 24
		tickets[i].tripDays = tripDays

		twoWays := rand.Intn(2) == 0
		tickets[i].twoWays = twoWays

		priceMillions := basicPriceMillions + (extraSpeedKmS * 1) // 1 - is a cofficient for pricing an additional speed
		if twoWays {
			priceMillions *= 2
		}
		tickets[i].priceInMillions = priceMillions
	}

	return tickets
}

func printTickets(ticketsToPrint []ticket) {
	fmt.Printf("| %-22v| %-7v| %-10v| %-5v |\n", "Sapce travel company", "Days", "Type", "Price")
	fmt.Printf("======================================================\n")
	for _, ticket := range ticketsToPrint {
		tripType := "One way"
		if ticket.twoWays {
			tripType = "Two ways"
		}
		fmt.Printf("| %-22v| %-7v| %-10v| $%4v |\n", ticket.company, ticket.tripDays, tripType, ticket.priceInMillions)
	}
}

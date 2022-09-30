package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(getRandomDate())
	}
}

func getRandomDate() string {
	year := rand.Intn(2022) + 1
	month := rand.Intn(12) + 1
	leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
	daysInMonth := 31
	switch month {
	case 2:
		daysInMonth = 28
		if leap {
			daysInMonth = 29
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	return fmt.Sprintf("Year: %v; Month: %v; Day: %v", year, month, day)
}
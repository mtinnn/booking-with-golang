package main

import (
	"fmt"
	"strings"
)

func main() {
	const allticket =50
	ticketnumber := 50
	var firtName string
	var lastName string
	var email string
	var countOfTickets int
	var bookings = []string{}
	for {
		// inpuuut
		fmt.Println("welcome")
		fmt.Println("enter your first name")
		fmt.Scan(&firtName)
		fmt.Println("enter your lastname")
		fmt.Scan(&lastName)
		fmt.Println("enter your email")
		fmt.Scan(&email)
		fmt.Println("how mant tickets you want?")
		fmt.Scan(&countOfTickets)


		///// are we have enouf tickets?
		ticketnumber=ticketnumber-countOfTickets
		if ticketnumber<0{
			println("we dont have enough tickets we")
			break
		}

		fmt.Printf("%v %v  you booking %v of %v we will send varify email to %v\n", firtName, lastName, countOfTickets,allticket, email)
		bookings = append(bookings, firtName+" "+lastName)

		namees := []string{}
		for _, name := range bookings {
			var n = strings.Fields(name)
			namees = append(namees, n[0])
			break

		}
		fmt.Println("================================================================")
		fmt.Printf("list of booking names: %v\n", namees)

	}

}

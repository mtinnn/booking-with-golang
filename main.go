package main

import (
	"fmt"
	"strings"
)

func main() {
	
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

		//check the inputs
		validname:=len(firtName)>2&&len(lastName)>2
		validemail:=strings.Contains(email,"@")
		validCountOfTickets:=countOfTickets<ticketnumber

		///// are we have enouf tickets?
	if validCountOfTickets&&validemail&&validname{
		bookings = append(bookings, firtName+" "+lastName)
		
		fmt.Printf("%v %v  you booking %v tickets we will send varify email to %v \n we still have %v tickets \n", firtName, lastName, countOfTickets,email, ticketnumber,)

		namees := []string{}
		for _, name := range bookings {
			var n = strings.Fields(name)
			namees = append(namees, n[0])
			break

		}
		fmt.Println("================================================================")
		fmt.Printf("list of booking names: %v\n", namees)

	}else{

		if !validemail{

			fmt.Println("you email is not correct")
		}
		if !validname{
			println("enter valid name")

		}
		if !validCountOfTickets{
			
			fmt.Printf("we dont have enough tickets we just have %v tickets",ticketnumber)
		}




	}



}}

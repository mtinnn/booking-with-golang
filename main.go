package main

import (
	"fmt"
	"strings"
)
var reminingTicket uint= 50
var bookings = []string{}
var confransName="MAGICK"

func main() {
	
	


	for {
		//get grating
		getgratin()
		// inpuuut
		firtName, lastName,email,countOfTickets:= inputs()
		//check the inputs
		validname,validemail,validCountOfTickets:= checkinput(firtName, lastName,email,countOfTickets)
		//add to list
		booking(validCountOfTickets , validemail ,validname ,firtName ,lastName ,email,countOfTickets)
		
	//print first name
		nameList()

		}
		

	}



	






func getgratin(){
	fmt.Printf("welcome %v  confrance you can booking your tickets from this platform\n",confransName)
	fmt.Printf("we have %v tickets\n",reminingTicket)
	fmt.Println("=====================")
}

func inputs()(string, string,string,uint){
	var firtName string
	var lastName string
	var email string
	var countOfTickets uint

	fmt.Println("enter your first name")
	fmt.Scan(&firtName)
	fmt.Println("enter your lastname")
	fmt.Scan(&lastName)
	fmt.Println("enter your email")
	fmt.Scan(&email)
	fmt.Println("how many tickets you want?")
	fmt.Scan(&countOfTickets)

	return firtName,lastName,email,countOfTickets

}

func checkinput(firtName string, lastName string, email string, countOfTickets uint)(bool,bool,bool){

	validname:=len(firtName)>2&&len(lastName)>2
	validemail:=strings.Contains(email,"@")
	validCountOfTickets:=countOfTickets<reminingTicket


	if !validemail{

		fmt.Println("you email is not correct")
		fmt.Println("=====================")
	}
	if !validname{
		println("enter valid name")
		fmt.Println("=====================")

	}
	if !validCountOfTickets{
		
		fmt.Printf("we dont have enough tickets we just have %v tickets\n",reminingTicket)
		fmt.Println("=====================")
	}
	return validname,validemail,validCountOfTickets

}

func booking(validCountOfTickets bool, validemail bool,validname bool,firtName string,lastName string,email string,countOfTickets uint){
	if validCountOfTickets&&validemail&&validname{
		bookings = append(bookings, firtName+" "+lastName)
		reminingTicket=reminingTicket-countOfTickets
		fmt.Printf("%v %v  you booking %v tickets we will send varify email to %v \n we still have %v tickets \n", firtName, lastName, countOfTickets,email, reminingTicket,)
		fmt.Println("=====================")
}}

func nameList(){
	namees := []string{}
	for _, name := range bookings {
		n := strings.Fields(name)
		namees = append(namees, n[0])
		}
		
		fmt.Printf("who pepele booking:%v\n",namees)
	}






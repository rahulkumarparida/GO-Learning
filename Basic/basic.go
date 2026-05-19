package main

import (
	"basic/calculate"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var confrenceName string = "GO conference"	
const ticketPrice int = 250 // in rupees 
const ticketsCount int= 50
var remainingTickets uint= 50
type UserData struct{
	 firstName string
	 lastName string
	 ticketsBooked uint
	 totalPrice int
}
var bookings = make([]UserData,0)


func main()  {


	greetUser()
	var bookedMembers []string
	
	// for remainingTickets > 0{
			var firstName, lastName,ticketsBooked = userInput()
		

			if remainingTickets <= ticketsBooked && ticketsBooked <= 50 {
				ticketsBooked = remainingTickets
			}
			if ticketsBooked > 50{
				fmt.Println("Sorry the max tickets that can be bought are: ", remainingTickets)
				fmt.Println("Please try again!!")
				// continue
			}
			remainingTickets = remainingTickets - ticketsBooked
			
			var userName = firstName+" "+lastName
			var userData = UserData{
				firstName: firstName,
				lastName: lastName,
				ticketsBooked: ticketsBooked,
				totalPrice: calculate.CalculateTicketPrice(int(ticketsBooked),ticketPrice),
			}
			// Create a data block using map
			//! Instea o using maps we create a struct for more robust handling nad flexibility on types 
			// var userData = make(map[string]string)
			// userData["firstName"]=firstName
			// userData["lastName"]=lastName
			// userData["bookedTickets"]=strconv.FormatUint(uint64(ticketsBooked),10)

			bookings = append(bookings,userData)
			bookedMembers = bookedUsernames()
	
	

			bookingConfirmation(userName,ticketsBooked,confrenceName,remainingTickets)
			
			if remainingTickets == 0 {
				fmt.Println("We are out of tickets.")
				// break
			}
		// }	

		fmt.Println("Booked members:",bookedMembers)	
		wg.Wait()
	}

func greetUser()  {
	fmt.Println("Welcome to the",confrenceName)
	fmt.Printf("Total tickets avaliable %v Remaining tickets: %v\n",ticketsCount ,remainingTickets)
	fmt.Println("Book your tickets to enjoy the shows.")
	fmt.Println(calculate.TestingCalculate())
}

func userInput() (string,string,uint) {
	var firstName string
	var lastName string
	var ticketsBooked uint
	fmt.Print("Enter your firstName:\n")
	fmt.Scan(&firstName)
	fmt.Print("Enter your lastName:\n")
	fmt.Scan(&lastName)
	fmt.Print("\nEnter the number of tickets you want to book:\n")
	fmt.Scan(&ticketsBooked)
	return firstName, lastName, ticketsBooked
}

func bookedUsernames() ([]string)  {
		startName := []string{}
		startName = nil
		for _,name := range bookings {

			// startName = append(startName, name["firstName"])

			startName = append(startName, name.firstName)
			// fmt.Println("FOr:",name)
		}
	return startName
}
func bookingConfirmation(userName string,ticketsBooked uint,confrenceName string,remainingTickets uint)  {
	var price = calculate.CalculateTicketPrice(int(ticketsBooked),ticketPrice)
	wg.Add(1)
	go sendTicket(ticketsBooked,userName, price)
	fmt.Printf("Thank you for purchaing the tickets will see you on the confrence.\n")
	fmt.Println("List of Bookings:",bookings,"\nTotal of ",len(bookings),"bookings")
	fmt.Printf("Remaining tickets: %v\n",remainingTickets)
	fmt.Println(calculate.Message)
}

func sendTicket(userTickets uint, userName string , price int) {
	time.Sleep(5*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v  costed %v", userTickets, userName,price)
	fmt.Println("###########################################################")
	fmt.Printf("\nSent ticket: %v \nto email address %v\n", ticket, "rrorx123@gmail.com\n")
	fmt.Println("###########################################################")
	wg.Done()

}
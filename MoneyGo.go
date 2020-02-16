package main

import (
	"fmt"
)

func main() {
	//Declare a variable that holds the options entered by the user
	var key string

	//Declare a variable to control whether to exit the for loop
	loop := true

	//Show main menu
	for {

		fmt.Println("\n----------My App To Track Spending----------")
		fmt.Println("          1 Income and Expense Detail")
		fmt.Println("          2 Registered Income")
		fmt.Println("          3 Registered Expenditure")
		fmt.Println("          4 Exit")
		fmt.Println("Please enter your order number: ")

		fmt.Scanln(&key)

		switch key {

		case "1" : 
			fmt.Println("------My Income and Expense Detail-------")
		case "2" :
			fmt.Println("------My Income and Expense Detail-------")
		case "3" :
			fmt.Println("------My Income and Expense Detail-------")
		case "4" :
			//break We can't break here because it only stop "switch"， not for loop
			loop = false
		default :
			fmt.Println("Oups! Please enter the correct option!")
		}

		if !loop {
			break
		}

	}

	fmt.Println("Bye, see you next time ^_^ !")
}
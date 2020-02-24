package main

import (
	"fmt"
)

func main() {
	//Declare a variable that holds the options entered by the user
	key := ""

	//Declare a variable to control whether to exit the for loop
	loop := true
	//Declare the scale
	balance := 1000.0
	//Amount of each income and expense
	money := 0.0
	//note of each income and expense
	note := ""
	//Detail of each income and expense, when we use money, only change "details"
	details := "Income/Expenditure\tBalance\t\tAmount\tNote"
	//Show main menu
	for {

		fmt.Println("\n----------My App To Track Spending----------")
		fmt.Println("          1 Income and Expense Detail")
		fmt.Println("          2 Registered Income")
		fmt.Println("          3 Registered Expenditure")
		fmt.Println("          4 Exit")
		fmt.Println("Please enter your order number (1-4): ")

		fmt.Scanln(&key)

		switch key {

		case "1" : 
			fmt.Println("------My Income and Expense Detail-------")
			fmt.Println(details)
		case "2" :
			fmt.Println("Amount of income: ")
			fmt.Scanln(&money)
			balance += money //Modify account balance

			fmt.Println("Note of income: ")
			fmt.Scanln(&note)
			//Splicing revenue into “details”
			details += fmt.Sprintf("\nIncome\t\t\t%v\t\t%v\t%v", balance, money, note)
			fmt.Println("Successfully added new details!")

		case "3" :
			fmt.Println("Amount of Expenditure： ")
			AmountEx:fmt.Scanln(&money)
			if money > balance {
				fmt.Println("Sorry! You don't have enough money！！！\nPlease enter another amount: ")
				goto AmountEx
			} else {
				balance -= money
				fmt.Println("Note of Expenditure： ")
				fmt.Scanln(&note)
				//Splicing revenue into “details”
				details += fmt.Sprintf("\nExpenditure\t\t%v\t\t%v\t%v", balance, money, note)
				fmt.Println("Successfully added new details!")
			}

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
package utils
//Encapsulate the functions of the app into a structure, and then call the methods of the structure to implement accounting and details
//把记账软件的功能,封装到一个结构体中,然后调用该结构体的方法,来实现记账和明细

import (
	"fmt"
)

type FamilyAccount struct {
	//Define required fields

	//Declare a field that holds the options entered by the user
	key string
	//Declare a field to control whether to exit the for loop
	loop bool
	//Declare the scale
	balance float64
	//Amount of each income and expense
	money float64
	//note of each income and expense
	note string
	//Define a field to record whether there is income and expenditure behavior
	flag bool
	//Detail of each income and expense, when we use money, only change "details"
	details string
}

//Write a constructor to give the factory pattern and return a FamilyAccount instance
//编写要给工厂模式的构造方法,返回一个FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {

	return &FamilyAccount{
		key: "",
		loop: true,
		balance: 1000.0,
		money: 0.0,
		note: "",
		flag: false,
		details: "Income/Expenditure\tBalance\t\tAmount\tNote",
	}

}

//Encapsulate showDetails method into structure
func (this *FamilyAccount) showDetails() {
	fmt.Println("------My Income and Expense Detail-------")
	if this.flag {
		//因为我们用的是FamilyAccount结构体里传过来的字段,所以不能直接yongflag, 要用this. , 表示调用这个方法的结构体变量里面的字段
		fmt.Println(this.details)
	} else {
		fmt.Println("No current income and expenditure details!")
	}
}

//Encapsulate method registeredIncome into structure
func (this *FamilyAccount) income() {
	fmt.Println("Amount of income: ")
	fmt.Scanln(&this.money)
	this.balance += this.money //Modify account balance

	fmt.Println("Note of income: ")
	fmt.Scanln(&this.note)
	//Splicing revenue into “details”
	this.details += fmt.Sprintf("\nIncome\t\t\t%v\t\t%v\t%v", this.balance, this.money, this.note)
	fmt.Println("Successfully added new details!")
	this.flag = true
}

//Encapsulate method Registered Expenditure into structure
func (this *FamilyAccount) pay() {
	fmt.Println("Amount of Expenditure： ")
	AmountEx:fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("Sorry! You don't have enough money！！！\nPlease enter another amount: ")
		goto AmountEx
	} else {
		this.balance -= this.money
		fmt.Println("Note of Expenditure： ")
		fmt.Scanln(&this.note)
		//Splicing revenue into “details”
		this.details += fmt.Sprintf("\nExpenditure\t\t%v\t\t%v\t%v", this.balance, this.money, this.note)
		fmt.Println("Successfully added new details!")
		this.flag = true
	}
}

func (this *FamilyAccount) exit() {
	//break We can't break here because it only stop "switch"， not for loop
	fmt.Println("Are you sure you want to quit? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("Please check your input and try again: y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

// Bind the corresponding method to the structure
////Show main menu
func (this *FamilyAccount) MainMenu() {
	for {

		fmt.Println("\n----------My App To Track Spending----------")
		fmt.Println("          1 Income and Expense Detail")
		fmt.Println("          2 Registered Income")
		fmt.Println("          3 Registered Expenditure")
		fmt.Println("          4 Exit")
		fmt.Println("Please enter your order number (1-4): ")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1" : 
			this.showDetails()
			
		case "2" :
			this.income()

		case "3" :
			this.pay()

		case "4" :
			this.exit()

		default :
			fmt.Println("Oups! Please enter the correct option!")
		}

		if !this.loop {
			break
		}

	}

	fmt.Println("Bye, see you next time ^_^ !")
}

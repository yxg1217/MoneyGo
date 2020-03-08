package main

import (
	"/Users/xiaogeyan/Document/github/MoneyGo/utils"
	"fmt"
)
func main() {

	fmt.Println("This is a money management app implemented in an object-oriented way :)")
	//First create a pointer instance of the familyAccount structure, then call its MainMenu, and then call the method in the structure in the menu
	//先创建了一个familyAccount这个结构体的的指针实例,然后在调用它的MainMenu,然后menu里面又调用结构体里面的方法
	utils.NewFamilyAccount().MainMenu() 
}
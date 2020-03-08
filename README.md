# MoneyGo
money management apps in Golang

## 面向过程test
----------My App To Track Spending----------
1 Income and Expense Detail
2 Registered Income
3 Registered Expenditure
4 Exit

---------------------Features 1----------------------
Show Main menu and Exit

Accoding to our features , print Main menu. Exit the app when user choise 4
---------------------Features 2----------------------
Expense Detail and Registered Income

1）因为需要显示明细，我们定义一个变量details string来记录
2）还需要定义变量来记录余额（balance），每次收支的金额（money），每次收支的说明（note）
---------------------Features 3----------------------
Registered Expenditure (just like feature 2)

---------------------Features 4----------------------

## 面向对象编程 Object-Oriented Programming

把记账软件的功能,把所有的字段封装到一个结构体中,然后把相应的方法绑定到结构体中,最后提供一个NewFamilyAccount工厂模式的构造方法,返回一个实例,来实现记账和明细
结构体的名字:FamilyAccount
通过在main方法中创建一个结构体familyAccount实例,再调用菜单,实现记账功能

Encapsulate the functions of the app into a structure, and then call the methods of the structure to implement accounting and details
Structure name: FamilyAccount
Realize the accounting function by creating a structure familyAccount instance in the main method
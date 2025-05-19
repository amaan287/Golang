package main

import (
	"fmt"

	"github.com/amaan287/Golang/Practice/structs/BankAccount"
	book "github.com/amaan287/Golang/Practice/structs/Book"
	employee "github.com/amaan287/Golang/Practice/structs/Employee"
	linkedlist "github.com/amaan287/Golang/Practice/structs/LinkedList"
	rect "github.com/amaan287/Golang/Practice/structs/Rectangle"
	students "github.com/amaan287/Golang/Practice/structs/Students"
)

func main() {
	// var x int = 10
	// var y int = 20

	b := book.CreateBook("Amaan mirza", "Let's go", 212)
	//newb := book.GetBook(&b)
	b.UpdateTitle("Let's go further")
	r := rect.CreateRect(10.22, 5.23)
	//area := r.CalcArea()
	r.Double()
	s := students.CreateStudent("Amaan", 100)
	s.AddTen()
	a := BankAccount.CreateAccout("Amaan", 0)
	a.Deposit(20)
	a.Withdraw(10)
	e := employee.CreateEmployee("Amaan", 19000)
	e.Raise()
	l := linkedlist.LinkedList{}
	l.InsertAtBack(2344)
	l.InsertAtBack(22)
	fmt.Printf("%d", l)

}

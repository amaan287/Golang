package main

import "fmt"
func main(){
	day:=3
switch day{
case 1:
	fmt.Println("Monday")
case 2:
	fmt.Println("Tuesday")
case 3:
	fmt.Println("Wednesday")
case 4:
	fmt.Println("Thrusday")
case 5:
	fmt.Println("Friday")
case 6:
	fmt.Println("Saturday")
default: 
fmt.Println("Sunday")
} 

month := "January"
switch month {
case "January", "December", "February":
	fmt.Println("winter")
case "March","April","June":
	fmt.Println("Summer")
}

temprature:=7

switch{
case temprature<=0:
	fmt.Println("Freezing")
case temprature>0 && temprature<10:
	fmt.Println("Cold")
case temprature>10 && temprature <20:
	fmt.Println("not warm not cold")
case temprature>20 && temprature < 30:
	fmt.Println("warm")
default:
	fmt.Println("Hot")

}


}
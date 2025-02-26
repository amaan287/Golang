package main

import "fmt"

 func main(){
	numbers:=make([]int,3,5)
	numbers = append(numbers, 23)
	numbers = append(numbers, 23)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
 }
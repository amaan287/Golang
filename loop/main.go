package main

import "fmt"
 func main(){

	for i :=0 ; i<10 ; i++{
		fmt.Println(i)
	}

counter :=0
for{
	fmt.Println("Infinite loop")
	counter ++
	if counter ==5 {
		break
	}
}

//^ important
numbers:=[]int{1,2,3,4,5,6}
for index, values:=range numbers{
	fmt.Printf("Index %d, Value: %d \n", index,values)
} 

name:="amaan mirza"
for index,char:=range name{
	fmt.Printf("Index %d, Char: %c \n", index,char)
}

 }
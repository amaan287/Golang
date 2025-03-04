package main

import "fmt"
func main(){
grades:= make(map[string]int)


number := map[int]int{
	1: 1,
	2: 3,
	3: 5,
}
fmt.Println(number)
grades["amaan"]=94
grades["new"]=100

grade,exists:=grades["new"] //? map return two values one is the value and other is the boolean value to check if it exists or not 

if exists{
fmt.Println(grade)
}

// delete(grades,"new") //? delete a enttry in the map
for index,value:=range grades{
	fmt.Printf("Value of the map is %d and index or key will be %s \n",value ,index)
}
fmt.Println(grades)
}
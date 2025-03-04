package main

import "fmt"


func modifyValueByRefrence(num *int){
	*num = *num+20
	fmt.Println(*num)
}


func main(){
	nums :=1
	var pointer *int  =&nums
	fmt.Println(*pointer)
	modifyValueByRefrence(&nums)
} 
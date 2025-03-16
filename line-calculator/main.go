package main

import (
	"fmt"
	"log"
)

func calculations(operator string, number1, number2 float64) (float64, string) {
	var result float64
	switch operator {
	case "*":
		result = number1 * number2
		return result, ""
	case "+":
		result = number1 + number2
		return result, ""
	case "-":
		result = number1 - number2
		return result, ""
	case "/":
		result = number1 / number2
		return result, ""
	}
	return 0, "Cant calculate"
}

func main() {
	var number1, number2 float64
	var operator string
	fmt.Println("Enter an operator to do calculations: *, +, -, /")
	fmt.Scan(&operator)
	fmt.Println("Enter a number: ")
	fmt.Scan(&number1)
	fmt.Println("Enter second number: ")
	fmt.Scan(&number2)
	res, err := calculations(operator, number1, number2)
	if err != "" {
		log.Fatal(err)
	}
	fmt.Println("result: ", res)
	main()
}

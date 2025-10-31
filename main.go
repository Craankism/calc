package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	var result float64

	//Input
	//first number
	fmt.Println("Enter first number: ")
	fmt.Scan(&num1)
	//operator
	fmt.Println("Enter operator: ")
	fmt.Scan(&operator)
	//second number
	fmt.Println("Enter second number: ")
	fmt.Scan(&num2)
	

	//Operator check and calculation
	switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
			if num2 == 0 {
				fmt.Println("Error: Division by zero")
				return
			}
		default:
			fmt.Println("Error: Unknown operator")
			return
	}
	fmt.Printf("Result: %g\n", result)
}
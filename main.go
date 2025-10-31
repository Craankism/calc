package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	var result float64
	firstRun := true

	for {
		//first number or previous result
		if firstRun { 
		fmt.Println("Enter number: ")
		fmt.Scan(&num1)
		firstRun = false
		} else {
			num1 = result
		}
		//operator
		fmt.Println("Enter operator: ")
		fmt.Scan(&operator)
		//second number
		fmt.Println("Enter number: ")
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
					continue
				}
		}
		fmt.Printf("Result: %g\n", result)
	}
}
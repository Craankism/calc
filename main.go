package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	var result float64

	fmt.Scanf("%f %s %f", &num1, &operator, &num2)

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
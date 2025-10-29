package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	var result float64
	var firstInput bool = true

	for {
		if firstInput {
			fmt.Println("Enter calculation (e.g., 3 + 4):")
			_, err := fmt.Scanf("%f %s %f", &num1, &operator, &num2)
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			firstInput = false
		} else {
			fmt.Print("Enter next operator and number (or 'q' to quit): ")
			_, err := fmt.Scanf("%s", &operator)
			if err != nil {
				fmt.Println("Invalid input.")
				return
			}
			num1 = result
		}

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
		default:
			fmt.Println("Error: Unknown operator")
			continue
	}
	fmt.Printf("Result: %g\n", result)
	}
}
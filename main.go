package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num1, num2, result float64
	var operator string
	firstRun := true
	scanner := bufio.NewScanner(os.Stdin)

	for {
		//first number or previous result
		if firstRun {
		fmt.Println("Enter number: ")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &num1)
		//clear calculator
		if scanner.Text() == "c" {
			fmt.Println("Calculator cleared.")
			result = 0
			firstRun = true
			continue
		}
		firstRun = false
		} else {
			num1 = result
		}
		//operator
		fmt.Println("Enter operator: ")
		scanner.Scan()
		//clear calculator
		if scanner.Text() == "c" {
			fmt.Println("Calculator cleared.")
			result = 0
			firstRun = true
			continue
		} else {
			operator = scanner.Text()
		}
		//second number
		fmt.Println("Enter number: ")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &num2)
		//clear calculator
		if scanner.Text() == "c" {
			fmt.Println("Calculator cleared.")
			result = 0
			firstRun = true
			continue
		}

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
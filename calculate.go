package main

import (
	"fmt"
	"strconv"
)

func calculate() {
	//call input with numbers and operators from parser
	numbers, operators := parser(lastTerm)
	fmt.Println("parser_output: ", numbers, operators)

	//start with first number
	//result get later updated inside of the loop
	result, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Println(err)
	}

	//calculate
	for i := 1; i < len(numbers);  i++ {

		b, err := strconv.Atoi(numbers[i])
		if err != nil {
			fmt.Println(err)
		}

		switch operators[i-1] {
		case "+":
			result = result + b
		case "-":
			result = result - b
		case "*":
			result = result * b
		case "/":
			result = result / b
			if b == 0 {
				fmt.Println("Err: No division with zero!")
			}
		default:
			fmt.Println("Error! Wrong operator")
		}
	}
	fmt.Println("result: ", result)
}

package main

import (
	"fmt"
	"strconv"
)

func calculate() {
	//call input with numbers and operators from parser
	numbers, operators := parser(lastTerm)
	fmt.Println("parser_response: ", numbers, operators)

	//start with first number
	//result get later updated inside of the loop
	result, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Println(err)
	}

	//calculate
	for i := 0; i < len(operators); i++ {

		b, err := strconv.Atoi(numbers[i+1])
		if err != nil {
			fmt.Println(err)
		}

		switch operators[i] {
		case "+":
			result = result + b
		case "-":
			result = result - b
		}
	}
	fmt.Println("result: ", result)
}

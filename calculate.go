package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate() string {
	output := parser(lastTerm)
	fmt.Println(output)

	var stack []string
	var result float64

	for i := 0; i < len(output); i++ {
		//numbers get added to stack
		if output[i] != "+" && output[i] != "-" && output[i] != "*" && output[i] != "/" {
			stack = append(stack, output[i])

			//set operator and calculate
		} else {
			//set a and b to last 2 positions stack
			a, err := strconv.ParseFloat(stack[len(stack)-2], 64)
			if err != nil {
				fmt.Println(err)
			}
			b, err := strconv.ParseFloat(stack[len(stack)-1], 64)
			if err != nil {
				fmt.Println(err)
			}

			currentOp := string(output[i])

			switch currentOp {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
				if b == 0 {
					fmt.Println("Err: Can't divide by 0")
				}
			}
			//remove last two numbers from stack and add result to stack
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.FormatFloat(result, 'f', -1, 64))
		}
	}
	if len(output) == 0 && len(stack) == 1 {
		stackStr := strings.Join(stack, "")
		_, err := strconv.ParseFloat(stackStr, 64)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Result: ", result)
	return strconv.FormatFloat(result, 'f', -1, 64)
}

/*pop remaining operators from stack (LIFO = Last-In-First-Out)
for len(stack) > 0 {
	output = append(output, stack[len(stack)-1])
	stack = stack[:len(stack)-1]
}
*/

/*
	//call input with numbers and operators from parser
	rpn := parser(lastTerm)
	fmt.Println("parser_output: ", rpn)

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
*/

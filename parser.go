package main

import (
	"fmt"
)
var normalOperators []string
var specialOperators []string

func parser(input string) (numbers []string, operators []string) {
	for i := 0; i < len(input); i++ {
		//numbers between "0-9"
		if input[i] >= '0' && input[i] <= '9' {
			//if number after number
			num := string(input[i])
			for i+1 < len(input) && input[i+1] >= '0' && input[i+1] <= '9' {
				i++
				num += string(input[i])
			}
			//combine numbers
			numbers = append(numbers, num)			

			//operators "+ - * /"
		} else if input[i] == '+' || input[i] == '-' {
			normalOperators = append(normalOperators, string(input[i]))
		} else if input[i] == '*' || input[i] == '/' {
			specialOperators = append(specialOperators, string(input[i]))
		} else {
			fmt.Println("wrong input")
		}
		operators = append(normalOperators, specialOperators...)
	}
	return numbers, operators
}
/*
Step 4: Evaluate RPN

We evaluate from left to right using a stack:

Push 22 → stack [22]

Push 3 → stack [22, 3]

Operator * → pop 22, 3, compute 22*3=66, push 66 → stack [66]

Push 4 → stack [66, 4]

Push 6 → stack [66, 4, 6]

Operator * → pop 4,6, compute 4*6=24, push 24 → stack [66, 24]

Operator + → pop 66,24, compute 66+24=90, push 90 → stack [90]

✅ Result: 90
*/
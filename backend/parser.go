package main

import (
	"regexp"
	"strings"
)

var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

// validateInput checks if input contains only valid characters: digits, operators, decimal point, and spaces
func validateInput(input string) bool {
	// Whitelist: digits (0-9), operators (+, -, *, /), decimal point (,), and spaces
	validPattern := regexp.MustCompile(`^\d+(,\d+)?([+\-*]\d+(,\d+)?)*$`)
	if !validPattern.MatchString(input) {
		return false
	}
	// Check for max length to prevent DoS
	if len(input) > 1024 {
		return false
	}
	return true
}

func parser(input string) (output []string) {
	if !validateInput(input) {
		return nil
	}
	var stack []string
	input = strings.ReplaceAll(input, ",", ".")
	for i := 0; i < len(input); i++ {
		//numbers between "0-9"
		if input[i] >= '0' && input[i] <= '9' {
			//check if number after number or for . after number
			num := string(input[i])
			for i+1 < len(input) && input[i+1] >= '0' && input[i+1] <= '9' || i+1 < len(input) && input[i+1] == '.' {
				i++
				num += string(input[i])
			}
			//combine numbers and push to output
			output = append(output, num)

			//operators "+ - * /"
		} else if input[i] == '+' || input[i] == '-' || input[i] == '*' || input[i] == '/' {
			currentOp := string(input[i])
			//pop operators with higher or equal priority to output
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[currentOp] {
				//pop operator to output from last index of the stack
				output = append(output, stack[len(stack)-1])
				//remove operator from last index of stack
				stack = stack[:len(stack)-1]

			}
			stack = append(stack, currentOp)
		}
	}
	//pop remaining operators from stack (LIFO = Last-In-First-Out)
	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return output
}

/*
Example input: 22*3+4*6
| Token | Aktion                                                                    | Output         | Stack  |
| ----- | ------------------------------------------------------------------------- | -------------- | ------ |
| 22    | Zahl → in Ausgabe                                                         | 22             |        |
| *     | Stack leer → `*` auf Stack                                                | 22             | *      |
| 3     | Zahl → in Ausgabe                                                         | 22 3           | *      |
| +     | Hat niedrigere Priorität als `*` → pop `*` in Ausgabe, dann `+` auf Stack | 22 3 *         | +      |
| 4     | Zahl → in Ausgabe                                                         | 22 3 * 4       | +      |
| *     | Höhere Priorität als `+` → direkt auf Stack                               | 22 3 * 4       | + *    |
| 6     | Zahl → in Ausgabe                                                         | 22 3 * 4 6     | + *    |
| Ende  | Stack entleeren                                                           | 22 3 * 4 6 * + | (leer) |
*/

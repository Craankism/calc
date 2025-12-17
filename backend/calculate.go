package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func calculate(term string) string {
	output := parser(term)
	fmt.Println(output)

	var stack []string
	var result *big.Float

	for i := 0; i < len(output); i++ {
		if len(output) == 1 {
			return output[0]
			//numbers get added to stack
		} else if output[i] != "+" && output[i] != "-" && output[i] != "*" && output[i] != "/" {
			stack = append(stack, output[i])

			//set operator and calculate
		} else {
			//string to float64
			s2, _ := strconv.ParseFloat(stack[len(stack)-2], 64)
			//float64 to big.Float
			a := big.NewFloat(s2)

			//String into float64
			s1, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
			//float64 to big.Float
			b := big.NewFloat(s1)

			//set zero to 0 for divide check
			zero := big.NewFloat(0)

			currentOp := string(output[i])

			switch currentOp {
			case "+":
				result = new(big.Float).Add(a, b)
			case "-":
				result = new(big.Float).Sub(a, b)
			case "*":
				result = new(big.Float).Mul(a, b)
			case "/":
				result = new(big.Float).Quo(a, b)
				if b.Cmp(zero) == 0 {
					fmt.Println("Err: Can't divide by 0")
				}
			}
			//remove last two numbers from stack and add result to stack
			stack = stack[:len(stack)-2]
			stack = append(stack, result.String())
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
	return result.String()
}

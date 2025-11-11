package main

import "fmt"

func parser(input string) (numbers []string, operators []string) {
	//example: 42 + 17 + 1
	//step 1: check input character by character
	for i := 0; i < len(input); i++ {
		if input[i] >= 48 && input[i] <= 57 {
			numbers = append(numbers, string(input[i]))
		} else if input[i] >= 42 && input[i] <= 47 {
			operators = append(operators, string(input[i]))
		} else {
			fmt.Println("idiot")
		}
	}
	//step 2: scan for number or operator
	//step 3: if number then add to numbers array
	//step 4: if operator then add to operators array
	//step 5: loop till done
	//step 6: return numbers array and operators array
	
	//step 7: combine numbers like "42" instead of "4 2"
	return numbers, operators
}

func main() {
	numbers, operators := parser("42+17+1")
	fmt.Println(numbers, operators)
}

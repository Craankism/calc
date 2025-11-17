package main

var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

var stack []string

func parser(input string) (output []string) {
	for i := 0; i < len(input); i++ {
		//numbers between "0-9"
		if input[i] >= '0' && input[i] <= '9' {
			//check if number after number
			num := string(input[i])
			for i+1 < len(input) && input[i+1] >= '0' && input[i+1] <= '9' {
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

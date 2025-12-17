# calc
simple calculator application for a webserver in go

<p align="center">
  <img src="docs/img/calc.png" alt="Calculator">
</p>

- simple calculations (no exponents and stuff)
- no negative numbers
- basic calculator design with buttons (windows calculator as reference)
- CI/CD with GitHub Actions (release-please)
- containerization with docker and docker-compose

work order:
![Image](docs/img/operating_principle.svg)

[Term](#term)<br>
[Parser](#parser)<br>
[Calculation](#calculation)<br>
[Result](#result)<br>

# Term
<i>Example: 2 + 3 * 4</i><br>

The user can input a term with the buttons of the UI or directly pass it to the text field via keyboard.<br>
The input then gets marshalled to JSON and sends an <b>HTTP POST request</b> to the server. The form data is included in the <b>request body</b>.

Test case:
```bash
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"term": "2+3*4"}'
```
# Parser
<i>Example: 2 3 4 * +</i><br>

<b>Convert Term to Reverse Polish Notation (RPN)</b><br>
First the Term gets umarshalled, so our program can use it.<br>
The program checks whether the input is valid to prevent inappropriate use. Valid inputs are the operators +, -, *, / and numbers 0-9.

Next, it checks the first character of the input. If it is a number, the program checks the following character. If the next character is also a number, the digits are combined into a single number.

For example, if the input is 23 + ..., the program reads the first character as 2, the second character as 3, and the third character as the operator +. Since the first two characters are numbers, they are combined into the number 23.

Numbers are always passed to the output stack.


Step by step operation:<br>
<b><u>First</u>, in our example the number 3 gets appended to the output stack.</b>

<b><u>Second</u>, it checks the next character and finds the operator +. Because the stack is empty at the moment, the operator + gets appended to the stack. This is a different stack, then the output stack.</b>

<b><u>Third</u>, the next character is the number 3, so it gets appended to the output stack.</b>

<b><u>Fourth</u>, next up is the character *. Now the program compares the precedence (priority) of * to the precedence of +. The precedence of * is higher then +, so it will be appended to the stack</b>

<b><u>Fifth</u>, the last character is the number 4, because it's a number, it will be appended to the output stack</b>

<b><u>Sixth</u>, lastly the remaining operators in the stack will be appended to the output stack, with the Last-In-First-Out(LIFO) principle</b>

# Calculation
<i>Example: 14</i><br>

The RPN gets passed to the calculation and its works similar as the parser.

Numbers get directly appended to the stack. <b>Stack:[2 3 4]</b><br>
If the next character is a operator, then the last two number in the stack get "popped" and added back to the stack. In our example the operator is * and the last two numbers are 4 and 3 (3*4=12). <b>Stack:[2 12]</b><br>
The next character is the operator +, so it's the same process as before. (2+12=14). <b>Stack:[14]</b><br>
If there are no characters left, then the result is the stack (14).

If there is only 1 character and it's a number, the number is the result.


# Result
The result now gets marshalled to JSON and responds to the POST request.
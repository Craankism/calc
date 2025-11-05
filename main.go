package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type quickmaffs struct {
	Term string `json:"term"`
}

func userInput(input quickmaffs) string {
	fmt.Println(input.Term)
	return input.Term
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	calcHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			input := r.FormValue("input")
			fmt.Printf("Input: %s", input)
		}
		http.ServeFile(w, r, "index.html")
		var qm quickmaffs
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &qm)
		io.WriteString(w, userInput(qm))
	}
	http.HandleFunc("/calc", calcHandler)

	http.ListenAndServe(":8080", nil)
}

/*
func main() {
	var num1, num2, result big.Float
	var operator string
	firstRun := true
	scanner := bufio.NewScanner(os.Stdin)

	//local host
	http.Handle("/calc/static/", http.StripPrefix("/calc/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/calc/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "index.html")
	})
	http.ListenAndServe(":8080", nil)

	for {
		//first number or previous result
		if firstRun {
		fmt.Println("Enter number: ")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &num1)
		//clear calculator
		if scanner.Text() == "c" {
			fmt.Println("Calculator cleared.")
			result.SetFloat64(0)
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
			result.SetFloat64(0)
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
			result.SetFloat64(0)
			firstRun = true
			continue
		}

		//Operator check and calculation
		switch operator {
			case "+":
				result.Add(&num1, &num2)
			case "-":
				result.Sub(&num1, &num2)
			case "*":
				result.Mul(&num1, &num2)
			case "/":
				result.Quo(&num1, &num2)
				if num2.Cmp(big.NewFloat(0)) == 0 {
					fmt.Println("Error: Division by zero")
					continue
				}
		}
		fmt.Printf("Result: %s\n", result.String())
	}
}*/

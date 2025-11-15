package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// /*
// 	{
// 		"term": "2+2"
// 	}
// */
// type quickmaffs struct {
// 	Term string `json:"term"`
// }

// /*
// 	{
// 		"result": "4"
// 	}
// */
// type calcResponse struct {
// 	Result float64 `json:"result"`
// }

var lastTerm string

func userInput(input quickmaffs) calcResponse {
	calcResponseObject := calcResponse{
		Result: 69,
	}

	fmt.Println("Input: ", input.Term)
	fmt.Println("Response: ", calcResponseObject)
	return calcResponseObject
}

// 	fmt.Println(input.Term)
// 	return calcResponseObject
// }

	calcHandler := func(w http.ResponseWriter, r *http.Request) {
		var qm quickmaffs
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(body, &qm)
		if err != nil {
			fmt.Println(err)
		}
		//parser to calculate
		lastTerm = qm.Term
		calculate()
		
		data, err := json.Marshal(userInput(qm))
		if err != nil {
			fmt.Println(err)
		}
		_, err = w.Write(data)
		if err != nil {
			fmt.Println(err)
		}
	}
	http.HandleFunc("/calc", calcHandler)

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "index.html")
// 	})
// 	http.HandleFunc("/static/style.css", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "static/style.css")
// 	})

// 	calcHandler := func(w http.ResponseWriter, r *http.Request) {
// 		var qm quickmaffs
// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		err = json.Unmarshal(body, &qm)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		data, err := json.Marshal(userInput(qm))
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		_, err = w.Write(data)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// 	http.HandleFunc("/calc", calcHandler)

//
// func main() {
// 	var num1, num2, result big.Float
// 	var operator string
// 	firstRun := true
// 	scanner := bufio.NewScanner(os.Stdin)

// 	//local host
// 	http.Handle("/calc/static/", http.StripPrefix("/calc/static/", http.FileServer(http.Dir("static"))))
// 	http.HandleFunc("/calc/", func(w http.ResponseWriter, r *http.Request){
// 		http.ServeFile(w, r, "index.html")
// 	})
// 	http.ListenAndServe(":8080", nil)

// 	for {
// 		//first number or previous result
// 		if firstRun {
// 		fmt.Println("Enter number: ")
// 		scanner.Scan()
// 		fmt.Sscan(scanner.Text(), &num1)
// 		//clear calculator
// 		if scanner.Text() == "c" {
// 			fmt.Println("Calculator cleared.")
// 			result.SetFloat64(0)
// 			firstRun = true
// 			continue
// 		}
// 		firstRun = false
// 		} else {
// 			num1 = result
// 		}
// 		//operator
// 		fmt.Println("Enter operator: ")
// 		scanner.Scan()
// 		//clear calculator
// 		if scanner.Text() == "c" {
// 			fmt.Println("Calculator cleared.")
// 			result.SetFloat64(0)
// 			firstRun = true
// 			continue
// 		} else {
// 			operator = scanner.Text()
// 		}
// 		//second number
// 		fmt.Println("Enter number: ")
// 		scanner.Scan()
// 		fmt.Sscan(scanner.Text(), &num2)
// 		//clear calculator
// 		if scanner.Text() == "c" {
// 			fmt.Println("Calculator cleared.")
// 			result.SetFloat64(0)
// 			firstRun = true
// 			continue
// 		}

// 		//Operator check and calculation
// 		switch operator {
// 			case "+":
// 				result.Add(&num1, &num2)
// 			case "-":
// 				result.Sub(&num1, &num2)
// 			case "*":
// 				result.Mul(&num1, &num2)
// 			case "/":
// 				result.Quo(&num1, &num2)
// 				if num2.Cmp(big.NewFloat(0)) == 0 {
// 					fmt.Println("Error: Division by zero")
// 					continue
// 				}
// 		}
// 		fmt.Printf("Result: %s\n", result.String())
// 	}
// }

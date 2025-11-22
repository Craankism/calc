package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
	{
		"term": "2+2"
	}
*/
type quickmaffs struct {
	Term string `json:"term"`
}

/*
	{
		"result": "4"
	}
*/
type calcResponse struct {
	Result string `json:"result"`
}

var lastTerm string

func userInput(input quickmaffs, result string) calcResponse {
	calcResponseObject := calcResponse{
		Result: result,
	}
	fmt.Println("Input: ", input.Term)
	fmt.Println("Response: ", calcResponseObject)
	return calcResponseObject
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/static/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/style.css")
	})

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
		result := calculate()
		result = strings.ReplaceAll(result, ".", ",")

		data, err := json.Marshal(userInput(qm, result))
		if err != nil {
			fmt.Println(err)
		}
		_, err = w.Write(data)
		if err != nil {
			fmt.Println(err)
		}
	}
	http.HandleFunc("/calc", calcHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

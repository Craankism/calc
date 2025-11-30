package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func userInput(input quickmaffs, result string) calcResponse {
	calcResponseObject := calcResponse{
		Result: result,
	}
	fmt.Println("Input: ", input.Term)
	fmt.Println("Response: ", calcResponseObject)
	return calcResponseObject
}

func main() {
	env := os.Getenv("ENVIROMENT")
	if env == "dev" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "frontend/dev.index.html")
		})
	} else if env == "prod" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "frontend/index.html")
		})
	}
	http.HandleFunc("/frontend/static/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/static/style.css")
	})

	calcHandler := func(w http.ResponseWriter, r *http.Request) {
		// Limit request body size to 100KB to prevent memory exhaustion
		r.Body = http.MaxBytesReader(w, r.Body, 100*1024)

		var qm quickmaffs
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(body, &qm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "invalid JSON"}`)
			fmt.Println("JSON unmarshal error:", err)
			return
		}

		// Validate input before processing
		if !validateInput(qm.Term) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "invalid characters in expression"}`)
			fmt.Println("Invalid input:", qm.Term)
			return
		}

		//parser to calculate
		result := calculate(qm.Term)
		result = strings.ReplaceAll(result, ".", ",")

		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(userInput(qm, result))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "failed to marshal response"}`)
			fmt.Println("JSON marshal error:", err)
			return
		}
		_, err = w.Write(data)
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	}
	http.HandleFunc("/calc", calcHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

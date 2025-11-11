package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input             string
		expectedNumbers   []string
		expectedOperators []string
	}{
		{"21+15-2", []string{"21", "15", "2"}, []string{"+", "-"}},
		{"3*4/2", []string{"3", "4", "2"}, []string{"*", "/"}},
		{"100-50+25", []string{"100", "50", "25"}, []string{"-", "+"}},
		{"7", []string{"7"}, []string{}},
		{"8+9*10-11/12", []string{"8", "9", "10", "11", "12"}, []string{"+", "*", "-", "/"}},
		{"0+0-0*0/0", []string{"0", "0", "0", "0", "0"}, []string{"+", "-", "*", "/"}},
		{"12345", []string{"12345"}, []string{}},
		{"1+2+3+4+5", []string{"1", "2", "3", "4", "5"}, []string{"+", "+", "+", "+"}},
		{"42-13*7/2+8", []string{"42", "13", "7", "2", "8"}, []string{"-", "*", "/", "+"}},
	}

	for _, test := range tests {
		numbers, operators := parser(test.input)
		if !reflect.DeepEqual(numbers, test.expectedNumbers) {
			t.Errorf("Input: %q, expected numbers %v, got %v", test.input, test.expectedNumbers, numbers)
		}
		if !reflect.DeepEqual(operators, test.expectedOperators) {
			t.Errorf("Input: %q, expected operators %v, got %v", test.input, test.expectedOperators, operators)
		}
	}
}

package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	var testCases = []struct{
		input []string
		result int
		err bool
	}{
		{
			input: []string{"10", "+", "15", "-", "15"},
			result: 10,
		},
		{
			input: []string{"10", "+", "15", "+", "15"},
			result: 40,
		},
		{
			input: []string{"10", "+", "15", "-", "15", "*", "1000"},
			result: -14975,
		},
		{
			input: []string{"1000", "*", "(", "2", "-", "15", ")"},
			result: -13000,
		},
		{
			input: []string{"1000", "*", "(", "2", "+", "8", ")"},
			result: 10000,
		},
	}

	for num, test := range testCases {
		out, _ := calc(test.input)
		if out != test.result {
			t.Errorf("%d != %d\n Test number: %d", out, test.result, num)
		}
	}
}

func TestParser(t *testing.T) {
	var testCases = []struct {
		input string
		result []string
	}{
		{
			input: "5+10",
			result: []string{"5", "+", "10"},
		},
		{
			input: "(1000+5)",
			result: []string{"(","1000", "+", "5", ")"},
		},
		{
			input: "1000+2*(25+100/4*(2*(2+2)))",
			result: []string{"1000", "+", "2", "*", "(", "25","+", "100", "/", "4", "*","(", "2", "*", "(", "2", "+", "2", ")", ")", ")"},
		},
		{
			input: "(1000+5)",
			result: []string{"(","1000", "+", "5", ")"},
		},
	}

	for _, test := range testCases {
		out := tokenize(test.input)
		for num := range out {
			if out[num] != test.result[num] {
				t.Errorf("%s != %s\n Test number: %d", out[num], test.result[num], num)
			}
		}
	}
}

func TestBoth(t *testing.T) {
	var testCases = []struct {
		input string
		result int
		err bool
	}{
		{
			input: "5+10",
			result: 15,
			err: false,
		},
		{
			input: "2+2*2",
			result: 6,
			err: false,
		},
		{
			input: "1000 + 2*(25 + 100/4*(2*(2 + 2)))",
			result: 1450,
			err: false,
		},
		{
			input: "(1000 + 5)",
			result: 1005,
			err: false,
		},
		{
			input: "(1 + 2) - 3",
			result: 0,
			err: false,
		},
		{
			input: "(1 + 2) * 3",
			result: 9,
			err: false,
		},
		{
			input: "1/0",
			result: 0,
			err: true,
		},
		{
			input: "5/-5",
			result: -1,
			err: false,
		},
		{
			input: "5*-5",
			result: -25,
			err: false,
		},
		{
			input: "-5-5",
			result: -10,
			err: false,
		},
		{
			input: "-5/-5",
			result: 1,
			err: false,
		},
	}

	for num, test := range testCases {
		out, err := calc(tokenize(test.input))
		if err == nil && test.err {
			t.Errorf("Expect error\n Test number: %d", num)	
		}

		if err != nil && !test.err {
			t.Errorf("Unexpected error\n Test number: %d", num)	
		}
		
		if err == nil && out != test.result {
			t.Errorf("%d != %d\n Test number: %d", out, test.result, num)
		}
	}
}
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
			err: false,
		},
		{
			input: []string{"10", "+", "15", "+", "15"},
			result: 40,
			err: false,
		},
		{
			input: []string{"10", "+", "15", "-", "15", "*", "1000"},
			result: -14975,
			err: false,
		},
		{
			input: []string{"1000", "*", "(", "2", "-", "15", ")"},
			result: -13000,
			err: false,
		},
		{
			input: []string{"1000", "*", "(", "2", "+", "8", ")"},
			result: 10000,
			err: false,
		},
	}

	for _, test := range testCases {
		out, err := calc(test.input)
		if err == nil && test.err {
			t.Errorf("Something happened")
		}
		if err != nil && !test.err {
			t.Errorf("Something happened")
		}
		if err == nil && out != test.result {
			t.Errorf("Something happened")
		}
	}
}
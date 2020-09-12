package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var lexem map[string]int = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func doOperation(operation string, first, second int) (result int) {
	switch operation {
	case "+":
		result = second + first
	case "-":
		result = second - first
	case "*":
		result= second * first
	case "/":
		result = second / first
	}

	return result
}

func popTwoItems(s *Stack) (first, second interface{}, err error) {
	first, err = s.pop()
	if err != nil {
		return
	}

	second, err = s.pop()
	return
}

func calc(lexems []string) (result int, err error) {
	var digits Stack
	var symbols Stack

	for _, symb := range lexems {
		if priority, is := lexem[symb]; is {
			topSymb, err := symbols.top();
			if err != nil || priority > lexem[topSymb.(string)] {
				symbols.push(symb)
			} else if priority <= lexem[topSymb.(string)] {
				operation, err := symbols.pop()
				if err != nil {
					return result, err
				}

				first, second, err := popTwoItems(&digits)
				if err != nil {
					return result, err
				}

				res := doOperation(operation.(string), first.(int), second.(int))
				digits.push(res)
				symbols.push(symb)
			}
		} else if symb == "(" {
			symbols.push(symb)
		} else if symb == ")" {
			for {
				operation, err := symbols.pop()
				if err != nil {
					return result, err
				}
				if operation == "(" {
					break;
				}

				first, second, err := popTwoItems(&digits)
				if err != nil {
					return result, err
				}

				res := doOperation(operation.(string), first.(int), second.(int))
				digits.push(res)
			}
		} else {  // If digit 
			digit, err := strconv.Atoi(symb)
			if err != nil {
				return result, err
			}
			digits.push(digit)
		}
	}

	for {
		operation, err := symbols.pop()
		if err != nil {
			break;
		}

		first, second, _ := popTwoItems(&digits)
		res := doOperation(operation.(string), first.(int), second.(int))
		digits.push(res)
	}

	res, err := digits.pop()
	return res.(int), err
}

func parser(line string) (result []string) {
	var digit string
	var symbol string

	for _, run := range line {
		if unicode.IsDigit(run) {
			if symbol != "" {
				result = append(result, symbol)
				symbol = ""
			}
			digit += string(run)
			continue
		}

		if digit != "" {
			result = append(result, digit)
			digit = ""
		}

		switch run {
		case '-', '+', '/', '*':
			symbol = string(run)
		case '(', ')':
			if symbol != "" {
				result = append(result, symbol)
				symbol = ""
			}
			result = append(result, string(run))
		}
	}

	if digit != "" {
		result = append(result, digit)
	}

	return result
}

func main() {
	var input string
    fmt.Scanln(&input)
	res, err := calc(parser(input))
	fmt.Println(res, err)
}
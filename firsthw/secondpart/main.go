package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var weight map[string]int = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func invokeOperation(operation string, digits *Stack) (result int, err error) {
	first, second, err := popTwoItems(digits)
	if err != nil {
		return result, err
	}

	switch operation {
	case "+":
		result = second.(int) + first.(int)
	case "-":
		result = second.(int) - first.(int)
	case "*":
		result= second.(int) * first.(int)
	case "/":
		result = second.(int) / first.(int)
	}

	return result, nil
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
	var operations Stack

	for _, oper := range lexems {
		if priority, is := weight[oper]; is {
			topOper, err := operations.top();
			if err != nil || priority > weight[topOper.(string)] {
				operations.push(oper)
			} else if priority <= weight[topOper.(string)] {
				operation, _ := operations.pop()
				res, err := invokeOperation(operation.(string), &digits)
				if err != nil {
					return result, err
				}
				digits.push(res)
				operations.push(oper)
			}
		} else if oper == "(" {
			operations.push(oper)
		} else if oper == ")" {
			for {
				operation, err := operations.pop()
				if err != nil {
					return result, err
				}
				if operation == "(" {
					break;
				}
				res, err := invokeOperation(operation.(string), &digits)
				if err != nil {
					return result, err
				}

				digits.push(res)
			}
		} else {  // If digit 
			digit, err := strconv.Atoi(oper)
			if err != nil {
				return result, err
			}
			digits.push(digit)
		}
	}

	for operations.size() > 0 {
		operation, _ := operations.pop()
		res, err := invokeOperation(operation.(string), &digits)
		if err != nil {
			return result, err
		}
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
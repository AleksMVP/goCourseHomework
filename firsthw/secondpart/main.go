package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var weights map[string]int = map[string]int{
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

func calc(tokens []string) int {
	var digits, operations Stack

	for _, argument := range tokens {
		if digit, err := strconv.Atoi(argument); err == nil {
			digits.push(digit);
		} else if argument == "(" {
			operations.push(argument);
		} else if argument == ")" {
			for {
				operation, err := operations.pop();
				if err != nil {
					panic("Something happened")
				} else if operation == "(" {
					break;
				}
				result, err := invokeOperation(operation.(string), &digits)
				if err != nil {
					panic("Something happened")
				}
				digits.push(result)
			}	
		} else if weight, ok := weights[argument]; ok {
			topOper, err := operations.top()

			if err != nil || weights[topOper.(string)] < weight {
				operations.push(argument)
			} else {
				operation, _ := operations.pop();
				result, err := invokeOperation(operation.(string), &digits)
				if err != nil {
					panic("Something happened")
				}
				digits.push(result)
				operations.push(argument)
			}
		} else {
			panic("Something happened")
		}
	}

	for operations.size() > 0 {
		operation, _ := operations.pop();
		result, err := invokeOperation(operation.(string), &digits)
		if err != nil {
			panic("Something happened")
		}
		digits.push(result)
	}

	result, err := digits.pop()
	if err != nil {
		panic("Something happened")
	}

	return result.(int)
}

func parser(line string) (result []string) {
	var digit string

	for _, run := range line {
		if unicode.IsDigit(run) {
			digit += string(run)
			continue
		} else if digit != "" {
			result = append(result, digit)
			digit = ""
		}

		switch run {
		case '-', '+', '/', '*', '(', ')':
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

	defer func() {
		if r := recover(); r != nil {
            fmt.Println(r)
        }
	}()

	result := calc(parser(input))
	fmt.Println(result)
}
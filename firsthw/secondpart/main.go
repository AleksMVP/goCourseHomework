package main

import (
	"fmt"
	"regexp"
	"strconv"
	"errors"
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
		result = second.(int) * first.(int)
	case "/":
		if (first.(int) == 0) {
			return result, errors.New("Division by zero")
		}
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

func calc(tokens []string) (int, error) {
	var digits, operations Stack

	for _, argument := range tokens {
		if digit, err := strconv.Atoi(argument); err == nil {
			digits.push(digit)
		} else if argument == "(" {
			operations.push(argument)
		} else if argument == ")" {
			for {
				operation, err := operations.pop()
				if err != nil {
					panic("Something happened")
				} else if operation == "(" {
					break
				}
				result, err := invokeOperation(operation.(string), &digits)
				if err != nil {
					return 0, err
				}
				digits.push(result)
			}
		} else if weight, ok := weights[argument]; ok {
			topOper, err := operations.top()

			if err != nil || weights[topOper.(string)] < weight {
				operations.push(argument)
			} else {
				operation, _ := operations.pop()
				result, err := invokeOperation(operation.(string), &digits)
				if err != nil {
					return 0, err
				}
				digits.push(result)
				operations.push(argument)
			}
		} else {
			panic("Something happened")
		}
	}

	for operations.size() > 0 {
		operation, _ := operations.pop()
		result, err := invokeOperation(operation.(string), &digits)
		if err != nil {
			return 0, err
		}
		digits.push(result)
	}

	result, err := digits.pop()
	if err != nil {
		panic("Something happened")
	}

	return result.(int), nil
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

var tokenRegex = regexp.MustCompile(`(?P<legal>\d+|\-d+|\-|\+|\*|/|\(|\))|(?P<ignore>\s)|(?P<error>.)`)

func tokenize(code string) (result []string) {
	groups := tokenRegex.SubexpNames()

	isNegative := false

	legal := indexOf("legal", groups)
	ignore := indexOf("ignore", groups)
	err := indexOf("error", groups)

	var row, col int
	allSub := tokenRegex.FindAllStringSubmatch(code, -1)
	previous := ""
	for _, elem := range allSub {
		if elem[legal] != "" {
			if _, err := strconv.Atoi(elem[legal]); err == nil && isNegative {
				result = append(result, "-" + elem[legal])
				isNegative = false
			} else {
				if (previous == "/" || previous == "*" || previous == "-" || previous == "+") && elem[legal] == "-" {
					isNegative = true
				} else {
					result = append(result, elem[legal])
				}
			}
			col += len(elem[legal])
			previous = elem[legal]
		} else if elem[ignore] != "" {
			if elem[ignore] == "\n" {
				col = 0
				row++
			} else {
				col++
			}
		} else if elem[err] != "" {
			panic(fmt.Sprintf("unexpected '%s' at (%d, %d)", elem[err], row, col))
		}

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

	result, err := calc(tokenize(input))
	if err != nil {
		fmt.Println(err);
	}
	fmt.Println(result)
}

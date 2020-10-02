package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parseArgs(args []string) (Params, error) {
	var param Params

	isInput := true
	for num, arg := range args {
		switch arg {
		case "-c":
			param.c = true
		case "-d":
			param.d = true
		case "-u":
			param.u = true
		case "-i":
			param.i = true
		case "-f":
			fCount, err := strconv.Atoi(args[num+1])
			if err != nil {
				return param, fmt.Errorf("Wrong argument after -f %s", err)
			}
			param.f = fCount
		case "-s":
			sCount, err := strconv.Atoi(args[num+1])
			if err != nil {
				return param, fmt.Errorf("Wrong argument after -s %s", err)
			}
			param.s = sCount
		default:
			if num == 0 || args[num-1] == "-f" || args[num-1] == "-s" {
				break
			}
			if isInput {
				param.input = arg
				isInput = false
			} else {
				param.output = arg
			}

			if !(num == len(args)-1 || num == len(args)-2) {
				return param, errors.New("Wrong arguments")
			}
		}
	}

	return param, nil
}

package main

import (
	"strconv"
	"strings"
)

func uniq(param Params, input []string) (out []string) {
	var inputCopy []string = make([]string, len(input), len(input))
	copy(inputCopy, input)

	if param.f > 0 {
		for num, line := range inputCopy {
			lineInRune := []rune(line)
			spaceCount := 0
			lastSpacePos := 0
			spacePos := 0
			for pos, r := range lineInRune {
				if r == ' ' {
					spaceCount++
					lastSpacePos = pos
				}
				if spaceCount == param.f {
					spacePos = pos
					break
				}
			}

			if spacePos != 0 {
				if spacePos != len(lineInRune) - 1 {
					inputCopy[num] = string(lineInRune[spacePos + 1:])
				} else {
					inputCopy[num] = string(lineInRune[spacePos:])
				}
			} else if spaceCount < param.f && spaceCount != 0 {
				if lastSpacePos != len(lineInRune) - 1 {
					inputCopy[num] = string(lineInRune[lastSpacePos + 1:])
				} else {
					inputCopy[num] = ""
				}
			} else {
				inputCopy[num] = string(lineInRune)
			}
		}
	}

	if param.s > 0 {
		for num, line := range inputCopy {
			lineInRune := []rune(line)
			if param.s < len(lineInRune) {
				inputCopy[num] = string(lineInRune[param.s:])
			} else {
				inputCopy[num] = ""
			}
		}
	}

	var linesCount map[string]int = map[string]int{}
	for _, line := range inputCopy {
		if param.i {
			line = strings.ToLower(line)
		}
		linesCount[line]++
	}

	for num, line := range inputCopy {
		if param.i {
			line = strings.ToLower(line)
		}

		switch {
		case param.c && linesCount[line] != -1:
			out = append(out, strconv.Itoa(linesCount[line])+" "+input[num])
		case param.d && linesCount[line] != -1 && linesCount[line] > 1:
			out = append(out, input[num])
		case param.u && linesCount[line] != -1 && linesCount[line] == 1:
			out = append(out, input[num])
		case !param.c && !param.u && !param.d && linesCount[line] != -1:
			out = append(out, input[num])
		}

		linesCount[line] = -1
	}

	return out
}

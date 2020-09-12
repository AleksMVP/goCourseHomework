package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Params struct {
	c      bool
	d      bool
	u      bool
	i      bool
	f      int
	s      int
	input  string
	output string
}

func read(param Params) (array []string, err error) {
	var reader bufio.Reader

	if param.input == "" {
		reader = *bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(param.input)
		if err != nil {
			return array, err
		}
		defer file.Close()
		reader = *bufio.NewReader(file)
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		array = append(array, strings.Trim(str, "\n"))
	}

	return array, nil
}

func write(param Params, array []string) error {
	var file *os.File

	if param.output == "" {
		file = os.Stdout
	} else {
		var err error
		file, err = os.OpenFile(param.output, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	for _, line := range array {
		file.WriteString(line + "\n")
	}

	return nil
}

func main() {
	param, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	if (param.c && param.d) || (param.c && param.u) || (param.u && param.d) {
		fmt.Println("uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]",
			"Use the parameters -с -d -u individually")
		return
	}

	input, err := read(param)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = write(param, uniq(param, input))

	if err != nil {
		fmt.Println(err)
		return
	}
}

package main

import (
	"testing"
)

func TestUniq(t *testing.T) {
	var testCases = []struct {
		param  Params
		input  []string
		result []string
	}{
		{
			Params{},
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		{
			Params{c: true},
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
			},
		},
		{
			Params{d: true},
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"I love music.",
				"I love music of Kartik.",
			},
		},
		{
			Params{u: true},
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"",
				"Thanks.",
			},
		},
		{
			Params{i: true},
			[]string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
			},
			[]string{
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Kartik.",
				"Thanks.",
			},
		},
		{
			Params{f: 1},
			[]string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"We love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		{
			Params{f: 10},
			[]string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"We love music.",
				"Thanks.",
			},
		},
		{
			Params{s: 1},
			[]string{
				"I love music.",
				"B love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		},
	}

	for _, test := range testCases {
		out := uniq(test.param, test.input)
		for num := range out {
			if test.result[num] != out[num] {
				t.Errorf("%s != %s \n Test number: %d", out[num], test.result[num], num)
			}
		}
	}
}

func TestArgParser(t *testing.T) {
	var testCases = []struct {
		input  []string
		result Params
		err    bool
	}{
		{
			[]string{"-c", "-d", "-u", "-i", "-f", "1", "-s", "1", "hello.txt", "godbye.txt"},
			Params{true, true, true, true, 1, 1, "hello.txt", "godbye.txt"},
			false,
		},
		{
			[]string{"-c", "-d", "-u", "-i", "-f", "1", "-s", "bobo", "hello.txt", "godbye.txt"},
			Params{},
			true,
		},
		{
			[]string{"-c", "-d", "-u", "-i", "-f", "bebe", "-s", "1", "hello.txt", "godbye.txt"},
			Params{},
			true,
		},
		{
			[]string{"-c", "-d", "-u", "-i", "hello.txt", "godbye.txt", "-f", "1", "-s", "1"},
			Params{},
			true,
		},
	}

	for num, test := range testCases {
		param, err := parseArgs(test.input)
		if test.err && err == nil {
			t.Errorf("Expected error Test number: %d", num)
			break
		}

		if param != test.result && err == nil {
			t.Errorf("Test number: %d", num)
		}
	}
}

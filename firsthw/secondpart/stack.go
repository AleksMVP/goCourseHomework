package main

import (
	"errors"
)

type Stack struct {
	array []interface{}
}

func (s *Stack)push(item interface{}) {
	s.array = append(s.array, item)
}

func (s *Stack)pop() (item interface{}, err error) {
	if len(s.array) == 0 {
		return item, errors.New("Stack empty")
	}

	item = s.array[len(s.array) - 1]
	s.array = s.array[:len(s.array) - 1]
	return item, nil
}

func (s *Stack)top() (item interface{}, err error) {
	if len(s.array) == 0 {
		return item, errors.New("Stack empty")
	}

	return s.array[len(s.array) - 1], nil
}

func (s *Stack)size() int {
	return len(s.array)
}
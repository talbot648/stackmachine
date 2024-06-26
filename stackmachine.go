package main

import (
	"errors"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i int) error {
	if !isNumberValid(i) {
		return errors.New("error: pushing invalid number")
	}
	*s = append(*s, i)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("cannot pop from an empty stack")
	}
	LastElementIndex := len(*s) - 1
	LastElement := (*s)[LastElementIndex]
	*s = (*s)[:LastElementIndex]
	return LastElement, nil
}

func (s *Stack) Dup() error {
	if s.isEmpty() {
		return errors.New("error: cannot duplicate empty stack")
	}
	LastElement := (*s)[len(*s)-1]
	*s = append(*s, LastElement)
	return nil

}

func (s *Stack) Add() error {
	if !isStackValid(s) {
		return errors.New("error: cannot add with less than two numbers in stack")
	}
	LastElement, _ := s.Pop()
	secondLastElement, _ := s.Pop()

	totalNumber := LastElement + secondLastElement
	if !isNumberValid(totalNumber) {
		return errors.New("error: total is an invalid value")
	}
	s.Push(totalNumber)
	return nil
}

func (s *Stack) Subtract() error {
	if !isStackValid(s) {
		return errors.New("error: cannot subtract with less than two numbers in stack")
	}
	LastElement, _ := s.Pop()
	secondLastElement, _ := s.Pop()

	subtractedTotal := LastElement - secondLastElement
	if !isNumberValid(subtractedTotal) {
		return errors.New("error: total is an invalid value")
	}
	s.Push(subtractedTotal)

	return nil
}

func (s *Stack) Multiply() error {
	if !isStackValid(s) {
		return errors.New("error: cannot multiply with less than two numbers in stack")

	}
	lastElement, _ := s.Pop()
	secondLastElement, _ := s.Pop()
	multipliedTotal := lastElement * secondLastElement
	if !isNumberValid(multipliedTotal) {
		return errors.New("error: total is an invalid value")
	}
	s.Push(multipliedTotal)
	return nil
}

func (s *Stack) Clear() {
	*s = nil
}

func (s *Stack) Sum() error {
	if s.isEmpty() {
		return errors.New("error: cannot add all numbers of an empty stack")
	}
	total := 0
	for _, i := range *s {
		i, _ = s.Pop()
		total += i
	}
	if !isNumberValid(total) {
		return errors.New("error: total is an invalid value")
	}
	s.Push(total)
	return nil
}

func stackMachine(commands string) (int, error) {
	var stack Stack
	splitString, err := splitString(commands)
	if err != nil {
		return 0, errors.New("")
	}
	for _, i := range splitString {
		switch i {
		case "+":
			err := stack.Add()
			if err != nil {
				return 0, errors.New("")
			}
		case "-":
			err := stack.Subtract()
			if err != nil {
				return 0, errors.New("")
			}
		case "*":
			err := stack.Multiply()
			if err != nil {
				return 0, errors.New("")
			}
		case "POP":
			stack.Pop()
		}
		numToPush, err := strconv.Atoi(i)
		if err == nil {
			err := stack.Push(numToPush)
			if err != nil {
				return 0, errors.New("")
			}
		}
	}
	result, _ := stack.Pop()
	return result, nil
}

func splitString(commands string) ([]string, error) {
	if isStringEmpty(commands) {
		return nil, errors.New("invalid input: string should not be empty")
	}
	splitString := strings.Split(commands, " ")
	return splitString, nil
}

func isStringEmpty(commands string) bool {
	return commands == ""
}

func isNumberValid(i int) bool {
	if i < 0 || i > 50000 {
		return false
	}
	return true
}

func isStackValid(s *Stack) bool {
	return len(*s) > 1
}

func main() {
	// main is unused - run using
	// go test ./...
}

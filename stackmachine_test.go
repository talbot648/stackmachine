package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestStartsWithEmptyStack(t *testing.T) {
	_, err := StackMachine("")

	if err == nil {
		t.Error("expected error due to no results")
	}
}

// Write your own TDD tests here as you develop


/*
  All these tests must pass for completion
*/
func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		commands string
		expected int
		expectedErr error		
	}{
		{commands:"no result", expected:0, expectedErr: errors.New("")},
		{commands:"50000 DUP +", expected: 0, expectedErr: errors.New("") },
		{commands:"5 6 + 2 *", expected: 22, expectedErr: nil },
		{commands:"1 2 3 4 + CLEAR 12 +", expected: 0, expectedErr: errors.New("") },
		{commands:"1 CLEAR 2 3 +", expected: 5, expectedErr: nil },
		{commands:"99", expected: 99, expectedErr: nil },
		{commands:"DOGBANANA", expected: 0, expectedErr: errors.New("") },
		{commands:"5 9 DUP + + 43 - 3 *", expected: 60, expectedErr: nil },
		{commands:"2 5 -", expected: 3, expectedErr: nil },
		{commands:"5 2 -", expected: 0, expectedErr: errors.New("") },
		{commands:"25000 DUP +", expected: 50000, expectedErr: nil },
		{commands:"50000 0 +", expected: 50000, expectedErr: nil },
		{commands:"50000 1 +", expected: 0, expectedErr: errors.New("") },
		{commands:"50001", expected: 0, expectedErr: errors.New("") },
		{commands:"50000 0 *", expected: 0, expectedErr: nil },
		{commands:"1 2 3 4 5 + + + + * 999", expected: 0, expectedErr: errors.New("") },
		{commands:"1 2 - 99 +", expected: 0, expectedErr: errors.New("") },
		{commands:"50000 50000 -", expected: 0, expectedErr: nil },
		{commands:"CLEAR", expected: 0, expectedErr: errors.New("") },
		{commands:"3 4 3 5 5 1 1 1 SUM", expected: 23, expectedErr: nil },
		{commands:"3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, expectedErr: errors.New("") },
		{commands:"3 4 3 5 5 1 1 1 SUM -", expected: 0, expectedErr: errors.New("") },
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s -> %v, %v", test.commands, test.expected, test.expectedErr), func(t *testing.T) {
			
			got, err := StackMachine(test.commands)

			if test.expectedErr != err {
				t.Errorf("got error %v, want %v", test.expectedErr, err)
			} else if got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}

}
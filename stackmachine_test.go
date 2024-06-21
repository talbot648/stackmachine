package main

import (
	"reflect"
	"testing"
)

//Tests to write
/*
given string can be split done
initialise stack done
push to the stack done
pop from the stack done
add the last two numbers together
subtract the last two numbers
duplicate the last number done
empty the stack
multiply the last 2 numbers together with a \*
SUM remove all numbers from stack, add them together and push number onto stack
return the last element of the stack
Report error for invalid inputs
report error for any number below 0 or over 50000 done
report error for empty string done
report error for adding with less than 2 numbers done
report error for subtracting with less than 2 numbers done
report error for multiplying with less than 2 numbers done
report error if clear is last command and stack is empty

*/

// Write your own TDD tests here as you develop
func TestGivenStringCanBeSplit(t *testing.T) {

	givenString := "5 6 8"
	want := []string{"5", "6", "8"}

	stringSplit, _ := splitString(givenString)

	if !reflect.DeepEqual(stringSplit, want) {
		t.Error("expected string to split")
	}
}

func TestSplitStringWithOneIten(t *testing.T) {
	givenString := "5"
	want := []string{"5"}

	stringSplit, _ := splitString(givenString)

	if !reflect.DeepEqual(stringSplit, want) {
		t.Error("expected string to split")
	}
}

func TestReportsErrorWhenSplittingEmptyString(t *testing.T) {

	givenString := ""

	_, err := splitString(givenString)

	if err == nil {
		t.Error("expected error when splitting empty string")
	}
}

func TestPushesAndPopsNumber(t *testing.T) {
	var stack Stack

	numToPush := 5
	stack.Push(numToPush)

	got, _ := stack.Pop()
	want := numToPush

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAcceptsPushingMaximumValidNumber(t *testing.T) {
	var stack Stack

	numToPush := 50000
	err := stack.Push(numToPush)

	if err != nil {
		t.Error("expected no error with pushing the maximum valid number")
	}
}

func TestAcceptsPushingMinimumValidNumber(t *testing.T) {
	var stack Stack

	numToPush := 0
	err := stack.Push(numToPush)

	if err != nil {
		t.Error("expected no error with pushing the minimum valid number")
	}
}

func TestRejectsPushingBelowMinimumValidNumber(t *testing.T) {
	var stack Stack

	numToPush := -1
	err := stack.Push(numToPush)

	if err == nil {
		t.Error("expected error with pushing minimum invalid number")
	}
}

func TestReportsErrorWhenPoppingEmptyStack(t *testing.T) {
	var stack Stack

	_, err := stack.Pop()

	if err == nil {
		t.Error("expected error when popping from and empty stack")
	}
}

func TestLastElementIsPopped(t *testing.T) {
	var stack Stack

	stack.Push(5)
	stack.Push(7)
	want := 7
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestCorectElementIsPoppedWithConsequtivePops(t *testing.T) {
	var stack Stack

	stack.Push(5)
	stack.Push(7)

	want := 5
	stack.Pop()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestStackStartsEmpty(t *testing.T) {
	var stack Stack

	got := stack.isEmpty()
	want := true

	if got != want {
		t.Error("expected stack should be empty")
	}
}

func TestDuplicateNumber(t *testing.T) {
	var stack Stack
	stack.Push(7)

	stack.Dup()
	numOne, _ := stack.Pop()
	numTwo, _ := stack.Pop()

	if numOne != numTwo {
		t.Error("expected last element in stack to be duplicated")
	}
}

func TestReportsErrorWhenDupEmptyStack(t *testing.T) {
	var stack Stack

	err := stack.Dup()

	if err == nil {
		t.Error("expected error when duplicating empty stack")
	}

}

func TestAddNumbers(t *testing.T) {
	var stack Stack
	stack.Push(5)
	stack.Push(7)

	want := 12
	stack.Add()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}

}

func TestAddNumbersToMaximumValidNumber(t *testing.T) {
	var stack Stack
	stack.Push(49999)
	stack.Push(1)

	want := 50000
	stack.Add()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}

}

func TestAddNumbersToMinimumValidNumber(t *testing.T) {
	var stack Stack
	stack.Push(0)
	stack.Push(0)

	want := 0
	stack.Add()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}

}

func TestReportsErrorAddingNumbersWithInvalidStackSize(t *testing.T) {
	var stack Stack
	stack.Push(5)

	err := stack.Add()

	if err == nil {
		t.Error("expected error when adding with one number in stack")
	}
}

func TestReportsErrorWhenTotalIsAboveValidSize(t *testing.T) {
	var stack Stack
	stack.Push(49999)
	stack.Push(2)
	err := stack.Add()

	if err == nil {
		t.Error("expected error when adding numbers above valid value")
	}
}

/*
func TestReportsErrorWhenTotalIsBelowValidSize(t *testing.T) {
	//Arrange
	var stack Stack
	stack.Push(0)
	stack.Push(-2)
	//Act
	fmt.Println(stack)
	err := stack.Add()

	//Assert
	if err == nil {
		t.Error("expected error when adding numbers above valid value")
	}
}
*/

func TestSubtractNumbers(t *testing.T) {
	var stack Stack
	stack.Push(3)
	stack.Push(10)

	want := 7
	stack.Subtract()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}

}

func TestSubtractNumbersToMinimumValidValue(t *testing.T) {

	var stack Stack
	stack.Push(5)
	stack.Push(5)

	want := 0
	stack.Subtract()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}
}

func TestAcceptsTotalSubtractedNumberToMaximumValidValue(t *testing.T) {
	var stack Stack
	stack.Push(0)
	stack.Push(50000)

	want := 50000
	stack.Subtract()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}
}

func TestReportsErrorSubtractingWithInvalidStackSize(t *testing.T) {
	var stack Stack
	stack.Push(5)

	err := stack.Subtract()

	if err == nil {
		t.Error("expected error when subtracting with an invalid stack size below two numbers")
	}
}

func TestReportsErrorWhenSubtractingTotalIsBelowMinimumValidValue(t *testing.T) {
	var stack Stack
	stack.Push(6)
	stack.Push(5)

	err := stack.Subtract()

	if err == nil {
		t.Error("expected error when subtracting to below minimum valid value ")
	}
}

func TestMultiplyLastTwoNumbers(t *testing.T) {
	var stack Stack

	stack.Push(5)
	stack.Push(6)

	want := 30
	stack.Multiply()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}
}

func TestAcceptsMultiplyingToMaximumValidValue(t *testing.T) {
	var stack Stack

	stack.Push(25000)
	stack.Push(2)

	err := stack.Multiply()

	if err != nil {
		t.Error("expected no error when multiplying to maximum valid value")
	}
}

func TestAcceptsMultiplyingToMinimumValidValue(t *testing.T) {
	var stack Stack

	stack.Push(25000)
	stack.Push(0)

	err := stack.Multiply()

	if err != nil {
		t.Error("expected no error when multiplying to minimum valid value")
	}
}

func TestReportsErrorMultiplyingWithInvalidStackSize(t *testing.T) {
	var stack Stack

	stack.Push(5)

	err := stack.Multiply()

	if err == nil {
		t.Error("expected error when multiplying with less than two numbers in stack")
	}
}

func TestReportsErrorMultiplyingAboveMaximumValidValue(t *testing.T) {
	var stack Stack

	stack.Push(25001)
	stack.Push(2)

	err := stack.Multiply()

	if err == nil {
		t.Error("expected error when multiplying with less than two numbers in stack")
	}
}

func TestClearsStack(t *testing.T) {
	var stack Stack

	want := true
	stack.Push(5)
	stack.Push(7)
	stack.Clear()
	got := stack.isEmpty()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}

}

/*
All these tests must pass for completion
*/

/*
func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expected    int
		expectedErr error
	}{
		{name: "empty error", commands: "", expected: 0, expectedErr: errors.New("")},
		{name: "add overflow", commands: "50000 DUP +", expected: 0, expectedErr: errors.New("")},
		{name: "too few add", commands: "99 +", expected: 0, expectedErr: errors.New("")},
		{name: "too few minus", commands: "99 -", expected: 0, expectedErr: errors.New("")},
		{name: "too few multiply", commands: "99 *", expected: 0, expectedErr: errors.New("")},
		{name: "empty stack", commands: "99 CLEAR", expected: 0, expectedErr: errors.New("")},
		{name: "sum single value", commands: "99 SUM", expected: 99, expectedErr: nil},
		{name: "sum empty", commands: "SUM", expected: 0, expectedErr: errors.New("")},
		{name: "normal +*", commands: "5 6 + 2 *", expected: 22, expectedErr: nil},
		{name: "clear too few", commands: "1 2 3 4 + CLEAR 12 +", expected: 0, expectedErr: errors.New("")},
		{name: "normal after clear", commands: "1 CLEAR 2 3 +", expected: 5, expectedErr: nil},
		{name: "single integer", commands: "9876", expected: 9876, expectedErr: nil},
		{name: "invalid command", commands: "DOGBANANA", expected: 0, expectedErr: errors.New("")},
		{name: "normal +-*", commands: "5 9 DUP + + 43 - 3 *", expected: 60, expectedErr: nil},
		{name: "minus", commands: "2 5 -", expected: 3, expectedErr: nil},
		{name: "underflow minus", commands: "5 2 -", expected: 0, expectedErr: errors.New("")},
		{name: "at overflow limit", commands: "25000 DUP +", expected: 50000, expectedErr: nil},
		{name: "at overflow limit single value", commands: "50000 0 +", expected: 50000, expectedErr: nil},
		{name: "overflow plus", commands: "50000 1 +", expected: 0, expectedErr: errors.New("")},
		{name: "overflow single value", commands: "50001", expected: 0, expectedErr: errors.New("")},
		{name: "times zero at overflow limit", commands: "50000 0 *", expected: 0, expectedErr: nil},
		{name: "too few at first", commands: "1 2 3 4 5 + + + + * 999", expected: 0, expectedErr: errors.New("")},
		{name: "normal simple", commands: "1 2 - 99 +", expected: 100, expectedErr: nil},
		{name: "at overflow minus to zero", commands: "50000 50000 -", expected: 0, expectedErr: nil},
		{name: "clear empties stack", commands: "CLEAR", expected: 0, expectedErr: errors.New("")},
		{name: "normal sum", commands: "3 4 3 5 5 1 1 1 SUM", expected: 23, expectedErr: nil},
		{name: "sum after clear stack", commands: "3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, expectedErr: nil},
		{name: "sum then too few", commands: "3 4 3 5 5 1 1 1 SUM -", expected: 0, expectedErr: errors.New("")},
		{name: "fibonacci", commands: "1 2 3 4 5 * * * *", expected: 120, expectedErr: nil},
	}

	for _, test := range tests {

		got, err := StackMachine(test.commands)

		if test.expectedErr != nil {
			if err == nil {
				t.Errorf("%s (%s) Expected error, but received nil", test.name, test.commands)
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("%s (%s) got error %v, want %v", test.name, test.commands, err, test.expectedErr)
			}
		} else if got != test.expected {
			t.Errorf("%s (%s) got %v, want %v", test.name, test.commands, got, test.expected)
		}
	}
}

*/

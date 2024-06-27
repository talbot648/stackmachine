package main

import (
	"errors"
	"fmt"
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
Report error for invalid inputs done
report error for any number below 0 or over 50000 done
report error for empty string done
report error for adding with less than 2 numbers done
report error for subtracting with less than 2 numbers done
report error for multiplying with less than 2 numbers done
report error if clear is last command and stack is empty done

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
	want := errors.New("invalid input: string should not be empty")
	_, got := splitString(givenString)

	if got.Error() != want.Error() {
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
	want := errors.New("error: pushing invalid number")
	got := stack.Push(numToPush)
	if got.Error() != want.Error() {
		t.Error("expected error with pushing minimum invalid number")
	}
}

func TestReportsErrorWhenPoppingEmptyStack(t *testing.T) {
	var stack Stack

	want := errors.New("cannot pop from an empty stack")
	_, got := stack.Pop()

	if got.Error() != want.Error() {
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

	want := errors.New("error: cannot duplicate empty stack")
	got := stack.Dup()

	if got.Error() != want.Error() {
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

	want := errors.New("error: cannot add with less than two numbers in stack")
	got := stack.Add()

	if got.Error() != want.Error() {
		t.Error("expected error when adding with one number in stack")
	}
}

func TestReportsErrorWhenTotalIsAboveValidSize(t *testing.T) {
	var stack Stack

	stack.Push(49999)
	stack.Push(2)

	want := errors.New("error: total is an invalid value")
	got := stack.Add()

	if got.Error() != want.Error() {
		t.Error("expected error when adding numbers above valid value")
	}
}

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

	want := errors.New("error: cannot subtract with less than two numbers in stack")
	got := stack.Subtract()

	if got.Error() != want.Error() {
		t.Error("expected error when subtracting with an invalid stack size below two numbers")
	}
}

func TestReportsErrorWhenSubtractingTotalIsBelowMinimumValidValue(t *testing.T) {
	var stack Stack
	stack.Push(6)
	stack.Push(5)

	want := errors.New("error: total is an invalid value")
	got := stack.Subtract()

	if got.Error() != want.Error() {
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

	want := errors.New("error: cannot multiply with less than two numbers in stack")
	got := stack.Multiply()

	if got.Error() != want.Error() {
		t.Error("expected error when multiplying with less than two numbers in stack")
	}
}

func TestReportsErrorMultiplyingAboveMaximumValidValue(t *testing.T) {
	var stack Stack

	stack.Push(25001)
	stack.Push(2)

	want := errors.New("error: total is an invalid value")
	got := stack.Multiply()

	if got.Error() != want.Error() {
		t.Error("expected error when the total from multiplying two values is above maximum valid value")
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

func TestAcceptsClearingEmptyStack(t *testing.T) {
	var stack Stack

	want := true
	stack.Clear()
	got := stack.isEmpty()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}
}

func TestSumMethodPopsAllElementsInStackBeforeAdding(t *testing.T) {
	var stack Stack

	want := true
	stack.Push(50)
	stack.Push(4)
	stack.Push(5)
	stack.Push(2)
	stack.Sum()

	stack.Pop()
	got := stack.isEmpty()

	if got != want {
		t.Error("expect stack to be empty after popping the total sum value")
	}

}

func TestAddingAllNumbersInStack(t *testing.T) {
	var stack Stack

	want := 28
	stack.Push(10)
	stack.Push(4)
	stack.Push(12)
	stack.Push(2)
	stack.Sum()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}

}

func TestAcceptsAddingAllNumbersToMaximumValidValue(t *testing.T) {
	var stack Stack

	want := 50000
	stack.Push(49990)
	stack.Push(4)
	stack.Push(5)
	stack.Push(1)
	stack.Sum()
	got, _ := stack.Pop()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}
}

func TestReportsErrorWhenAddingAllNumbersOnEmptyStack(t *testing.T) {
	var stack Stack

	want := errors.New("error: cannot add all numbers of an empty stack")
	got := stack.Sum()

	if got.Error() != want.Error() {
		t.Error("expected error when adding all numbers in an empty stack")
	}
}

func TestReportsErrorWhenTotalOfAllNumbersAddedIsAboveMaximumValidValue(t *testing.T) {
	var stack Stack

	stack.Push(49990)
	stack.Push(4)
	stack.Push(5)
	stack.Push(2)

	want := errors.New("error: total is an invalid value")
	got := stack.Sum()

	if got.Error() != want.Error() {
		t.Error("expected error when total is above maximum valid value")
	}
}

func TestReturnsErrorForNoCommand(t *testing.T) {
	_, err := stackMachine("")

	if err == nil {
		t.Error("expected error when calling stack machine with no command")
	}
}

func TestReturnsSingleNumber(t *testing.T) {
	want := 99
	got, _ := stackMachine("99")

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestRejectsTooLargeNumber(t *testing.T) {
	_, err := stackMachine("50001")

	if err == nil {
		t.Error("expected error when inputting too large of a number")
	}
}

func TestReturnsMostRecentNumber(t *testing.T) {
	want := 49
	got, _ := stackMachine("50 49")

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestReturnsSumOfTwoValues(t *testing.T) {
	want := 18
	got, _ := stackMachine("10 8 +")

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestReturnsErrorWhenSumOfTwoValuesIsTooLarge(t *testing.T) {
	_, err := stackMachine("49999 2 +")

	if err == nil {
		t.Error("expected error when the total number from a sum is too high")
	}
}

func TestReturnsDifferenceOfTwoValues(t *testing.T) {
	want := 5
	got, _ := stackMachine("3 8 -")

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestReturnsErrorWhenDiffernceOfTwoValuesIsTooSmall(t *testing.T) {
	_, err := stackMachine("3 2 -")

	if err == nil {
		t.Error("expected error when total difference is below the minimum valid value")
	}
}

func TestReturnsMultipleOfTwoValues(t *testing.T) {
	want := 12
	got, _ := stackMachine("4 3 *")

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestReturnsErrorWhenMultipleOfTwoValuesIsTooLarge(t *testing.T) {
	_, err := stackMachine("25001 2 *")

	if err == nil {
		t.Error("expected error when total value from multipling two numbers is above maximum valid value")
	}
}

func TestReturnsCorrectPoppedNumber(t *testing.T) {

	want := 5
	got, _ := stackMachine("5 7 POP")

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestReturnsErrorWhenPoppingEmptyStack(t *testing.T) {

	_, err := stackMachine("POP")

	if err == nil {
		t.Error("expected error when popping from empty stack")
	}
}

func TestReturnsDuplicatedNumber(t *testing.T) {
	want := 5

	got, gotErr := stackMachine("5 DUP POP")

	if gotErr != nil {
		t.Error("unexpected error")
	}

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestReturnsErrorWhenDuplicatingEmptyStack(t *testing.T) {

	_, err := stackMachine("DUP")

	if err == nil {
		t.Error("Expected error to be returned when duplicating empty stack")
	}
}

func TestReturnSummedStack(t *testing.T) {

	want := 22

	got, _ := stackMachine("5 7 6 4 SUM")

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestReturnErrorSummingEmptyStack(t *testing.T) {

	_, err := stackMachine("SUM")

	if err == nil {
		t.Error("expected error to be returned when summing empty stack")
	}
}

func TestReturnsErrorWhenReturningClearedStack(t *testing.T) {
	want := 0
	got, err := stackMachine("5 6 7 4 CLEAR")
	if err == nil {
		t.Error("expected an error")
	}
	if got != want {
		t.Error("expected 0 from cleared stack")
	}
}

func TestReturnErrorForInvalidInput(t *testing.T) {
	_, err := stackMachine("DOG CHILD 3")
	fmt.Println(err)
	if err == nil {
		t.Error("expected error for invalid input")
	}
}

/*
All these tests must pass for completion
*/

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

		got, err := stackMachine(test.commands)

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

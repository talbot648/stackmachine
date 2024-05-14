# stackmachine

Write an emulator of a simple stack machine.

A stack machine is a system that performs a sequence of simple operations on a stack of integers. Initially the stack is empty. The sequence of operations is given as a string. Operations are separated by single spaces. The following operations may be specified:

- An integer X (from 0 to 50000): the machine pushes X onto the stack;
- "POP": the machine removes the topmost number from the stack;
- "DUP": the machine pushes a duplicate of the topmost number onto the stack;
- "+": the machine pops the two topmost elements from the stack, adds them together and pushes the sum onto the stack;
- "-": the machine pops the two topmost elements from the stack, subtracts the second one from the first (topmost) one and pushes the difference onto the stack.
- "\*" pop the top two elements off the stack, multiply them together, push the result onto the stack
- "CLEAR" empties the stack so it has no elements on it

After processing all the operations, the machine returns the topmost value from the stack.

The machine processes unsigned integers (numbers from 0 to 50000). An overflow in addition or underflow in subtraction causes an error. The machine also reports an error when it tries to perform an operation that expects more numbers on the stack than the stack actually contains. Also, if, after performing all the operations, the stack is empty, the machine reports an error.

For example, given a string "4 5 6 - 7 +", the machine performs the following operations:

```
  operation | comment             | stack
 --------------------------------------------------
            |                     | [empty]
 "4"        | push 4              |
            |                     | 4
 "5"        | push 5              |
            |                     | 4, 5
 "6"        | push 6              |
            |                     | 4, 5, 6
 "-"        | subtract 5 from 6   |
            |                     | 4, 1
 "7"        | push 7              |
            |                     | 4, 1, 7
 "+"        | add 1 and 7         |
            |                     | 4, 8
```

Finally, the machine will return 8.

Given a string "13 DUP 4 POP 5 DUP + DUP + -", the machine performs the following operations:

```
 operation | comment             | stack
 --------------------------------------------------
            |                     | [empty]
 "13"       | push 13             |
            |                     | 13
 "DUP"      | duplicate 13        |
            |                     | 13, 13
 "4"        | push 4              |
            |                     | 13, 13, 4
 "POP"      | pop 4               |
            |                     | 13, 13
 "5"        | push 5              |
            |                     | 13, 13, 5
 "DUP"      | duplicate 5         |
            |                     | 13, 13, 5, 5
 "+"        | add 5 and 5         |
            |                     | 13, 13, 10
 "DUP"      | duplicate 10        |
            |                     | 13, 13, 10, 10
 "+"        | add 10 and 10       |
            |                     | 13, 13, 20
 "-"        | subtract 13 from 20 |
            |                     | 13, 7
```

Finally, the machine will return 7.

Given a string "5 6 + -", the machine reports an error. After the addition, there is only one number on the stack and the subtraction operation expects two.

Given a string "3 DUP 5 - -", the machine reports an error. The second subtraction yields a negative result.

Work:

- git clone this repo https://github.com/bjssacademy/stackmachine
- create the logic inside the function

`func StackMachine( commands string ) (int, error)`

that, given a string 'commands' containing a sequence of operations for the stack machine, returns the result the machine would return after processing the operations. The function should return an error if the machine would report an error while processing the operations.

- Use a test-first TDD approach to drive towards a solution.
- All example tests must pass
- All unit tests in the github repo must pass, prior to completion
- Ask for help on your coding channels

Examples:

1. Given commands = "4 5 6 - 7 +", the function should return 8, as explained above.

2. Given commands = "13 DUP 4 POP 5 DUP + DUP + -" the function should return 7.

3. Given commands = "5 6 + -", the function returns an error

4. Given commands = "3 DUP 5 - -" returns an error

5. Given commands = "50000 DUP +" returns an error

6. Given commands = "5 6 + 2 \*" returns 22

7. Given commands = "1 2 3 4 + CLEAR 12 +" returns an error

8. Given commands = "1 CLEAR 2 3 +" returns 5

Assume that:

- The maximum number of commands is within the range [0..2,000];
- The command input string is a valid sequence of word machine operations.
- Each command is separated by a single space

In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.

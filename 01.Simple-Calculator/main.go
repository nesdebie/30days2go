package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: <operand1> \"<operator>\" <operand2>")
		fmt.Println("Example: " +  os.Args[0] + " 5 \"+\" 3")
		return
	}
	fmt.Println(os.Args[1:])

	rawArgs1 := os.Args[1]
	operator := os.Args[2]
	rawArgs2 := os.Args[3]

	operand1, err1 := strconv.ParseFloat(rawArgs1, 64)
	operand2, err2 := strconv.ParseFloat(rawArgs2, 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid numbers.")
		return
	}
	var result float64
	switch operator {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "*":		
			result = operand1 * operand2
		case "/":		
			if operand2 == 0 {
				fmt.Println("Error: Division by zero")
				return
			}
			result = operand1 / operand2
		default:
			fmt.Println("Invalid operator. Use +, -, *, or /")
			return
	}
	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", operand1, operator, operand2, result)
}

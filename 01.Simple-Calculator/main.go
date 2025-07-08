package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: <num1> \"<operator>\" <num2>")
		fmt.Println("Example: " +  os.Args[0] + " 5 \"+\" 3")
		return
	}
	fmt.Println(os.Args[1:])

	s1 := os.Args[1]
	operator := os.Args[2]
	s2 := os.Args[3]

	num1, err1 := strconv.ParseFloat(s1, 64)
	num2, err2 := strconv.ParseFloat(s2, 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid numbers.")
		return
	}
	var result float64
	switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":		
			result = num1 * num2
		case "/":		
			if num2 == 0 {
				fmt.Println("Error: Division by zero")
				return
			}
			result = num1 / num2
		default:
			fmt.Println("Invalid operator. Use +, -, *, or /")
			return
	}
	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	firstInput := getInput("First Number")
	secondInput := getInput("Second Number")

	var result float64

	//Notice here that we are executing getOperation into operation variable, before evaluating it
	switch operation := getOperation(); operation {
	case "+":
		result = addNumbers(firstInput, secondInput)
	case "-":
		result = subtractNumbers(firstInput, secondInput)
	case "*":
		result = multiplyNumbers(firstInput, secondInput)
	case "/":
		result = divideNumbers(firstInput, secondInput)
	default:
		panic("Invalid operation")
	}
	result = math.Round(result*100)/100
	fmt.Printf("The result is %v\n\n", result)

}

func getInput(prompt string) float64 {
	fmt.Printf("%v: \n", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return value
}

func getOperation() string {
	fmt.Println("Select an operation")
	operation, _ := reader.ReadString('\n')
	return strings.TrimSpace(operation)
}

func addNumbers(first, second float64) float64 {
	return first + second
}

func subtractNumbers(first, second float64) float64 {
	return first - second
}

func multiplyNumbers(first, second float64) float64 {
	return first * second
}

func divideNumbers(first, second float64) float64 {
	return first / second
}

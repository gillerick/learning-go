package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	firstInput := getInput("First Number")
	secondInput := getInput("Second Number")

}

func getInput(prompt string) float64 {
	fmt.Println("%v: ", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return value
}

func getOperation() string {
	fmt.Print("Select an operation")
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
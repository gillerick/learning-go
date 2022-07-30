package main

import "fmt"

func main() {
	multiSum, count := addValues(45, 45, 34, 1)
	fmt.Println("Multi-value sum is", multiSum)
	fmt.Println("Total values is", count)

}

func addValues(values ...int) (int, int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, len(values)
}

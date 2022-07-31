package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)

	//The length variable is the number of characters written into the file
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)

	defer file.Close()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

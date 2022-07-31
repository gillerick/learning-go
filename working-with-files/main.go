package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	defer func(file *os.File) {
		err := file.Close()
		checkError(err)
	}(file)
	defer readFile("./fromString.txt")

}

//Reading a file always produces an array of bytes
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

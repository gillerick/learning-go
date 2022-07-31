package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	response, err := http.Get(url)
	checkError(err)

	fmt.Printf("Response type: %T\n", response) //*http.Response

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	content := string(bytes)
	fmt.Print(content)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

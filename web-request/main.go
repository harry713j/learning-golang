package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("******************* Web Request Demo ************************")

	response, err := http.Get(url)

	handleError(err)

	fmt.Printf("Resoponse tye is %T\n", response)

	// caller's responsibility to close the connection
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	handleError(err)

	content := string(databytes)

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Web Content: ", content)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

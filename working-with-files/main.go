package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("************* File Operations ***************")
	content := "Hello there this is a file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		fmt.Println("Failed to create/access the file")
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		fmt.Println("Failed to write in the file")
		panic(err)
	}

	fmt.Println("Length of content in file is: ", length)
	// close the file
	defer file.Close()

	// read
	readFile("./myfile.txt")

}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	// databyte: return data in byte format
	if err != nil {
		fmt.Println("Failed to read the file")
		panic(err)
	}

	fmt.Println("Content of the file: \n", string(databyte))
}

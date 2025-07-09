package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("HTTP SERVER")

	// mux := http.NewServeMux();
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "Hello")
	// })

	// read from file
	bytedata, err := os.ReadFile("./index.html")

	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, string(bytedata))
	}

	http.HandleFunc("/", handler)
	// any error this will handle
	log.Fatal(http.ListenAndServe(":4000", nil))
}

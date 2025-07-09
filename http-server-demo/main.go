package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

var users = make(map[int]User) // store the user in-memory

func main() {
	fmt.Println("HTTP SERVER")
	const PORT string = ":8080"
	mux := http.NewServeMux() // handle the routes
	// read from file
	bytedata, err := os.ReadFile("./index.html")

	if err != nil {
		panic(err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(bytedata)
	})

	// post
	mux.HandleFunc("POST /users", postUser)

	mux.HandleFunc("GET /users", getUsers)

	// any error this will handle
	fmt.Println("Server is Listening to port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// insert it to the user
	users[len(users)+1] = user

	w.WriteHeader(http.StatusNoContent)
	io.WriteString(w, "User Accepted")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

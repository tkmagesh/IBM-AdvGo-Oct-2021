package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

var users []User = []User{
	{FirstName: "John", LastName: "Doe", Age: 25},
	{FirstName: "Magesh", LastName: "Kuppan", Age: 45},
}

/*
Expose a /users endpoint
	GET -> returns all the users
	POST -> add the user to the users slice with 201 status code and the object added as the response
*/

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
	})
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			FirstName: "John",
			LastName:  "Doe",
			Age:       25,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	})
	http.ListenAndServe(":8080", nil)
}

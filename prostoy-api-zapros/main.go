package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ListOffriends)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type User struct {
	name     string
	lastName string
	age      uint
	country  string
}

func ListOffriends(w http.ResponseWriter, r *http.Request) {
	// u := User{name: "Almaz", lastName: "Halykow", age: 23, country: "Turkemnistan"}
	friends := [3]string{"Soltan", "Archa", "Yhlas"}
	data, _ := json.Marshal(friends)
	dataString := string(data)
	// UserData, _ := json.Marshal(u)
	// UserDataString := string(UserData)
	fmt.Fprintf(w, dataString)
}

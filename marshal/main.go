package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Implement the generic function named marshal that can accept any value
// of type T and marshal that value into JSON.
func marshal[T any](v T) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Declare a struct type named User with two fields, Name and Age.
type User struct {
	Name string
	Age  int
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	u := User{
		Name: "Bill",
		Age:  10,
	}

	// Call the generic marshal function.
	data, err := marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {

	// Construct a value of type User.

	// Print the JSON produced by the marshal function.

	fmt.Println(string(data))
}

package main

import (
	"encoding/json"
	"os"
)

func main() {
	var user User

	user = NewUser("John McClane", 42)

	if user.Valid() {
		json.NewEncoder(os.Stdout).Encode(user)
	} else {
		json.NewEncoder(os.Stdout).Encode(user.Errors)
	}

	user = NewUser("", 42)

	if user.Valid() {
		json.NewEncoder(os.Stdout).Encode(user)
	} else {
		json.NewEncoder(os.Stdout).Encode(user.Errors)
	}
}

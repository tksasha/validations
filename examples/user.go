package main

import (
	"github.com/tksasha/model"
	"github.com/tksasha/validations"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	model.Model
}

func (user *User) Valid() bool {
	user.Errors = validations.NewErrors()

	validations.ValidatePresenceOf(user.Errors, "name", user.Name)

	return user.Errors.Empty()
}

func NewUser(name string, age int) (user User) {
	user.Name = name

	user.Age = age

	user.Errors = validations.NewErrors()

	return
}

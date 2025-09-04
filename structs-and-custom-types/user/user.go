package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
	// User User
}

// creating a method

func (userDetails *User) OutputUserDetails() {
	fmt.Println(userDetails.firstName, userDetails.lastName, userDetails.birthDate)
}

func (userDetails *User) ClearUserName() {
	userDetails.firstName = ""
}

func NewAdmin(email, password string) *Admin{
	return &Admin{
		email: email,
		password: password,
		User: User{
			firstName: "firstName",
			lastName: "lastName",
			birthDate: "___",
			createdAt: time.Now(),
		},
	}
}


func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname, birthdate is required")
	}

	return &User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}
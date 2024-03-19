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

// Adding Method to Struct
func (u *User) OutputUserDetails() {
	fmt.Printf("Created At: %v , %v %v", u.createdAt.Local().Format("2006/01/02 03:04:05 PM"), u.firstName, u.lastName)
}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First Name , Last Name and BirthDate cannot be empty")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

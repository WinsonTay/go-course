package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please Enter your first name.. ")
	userLastName := getUserData("Please Enter your last name .. ")
	userBirthDate := getUserData("Please Enter your birthdate (MM/DD/YYYY)")

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	var value string
	fmt.Print(promptText)
	fmt.Scanln(&value)
	return value
}

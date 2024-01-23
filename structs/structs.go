package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date: ")

	var appUser *user.User

	admin := user.NewAdmin("text@xyz.com", "mayur")
	
	
	admin.ClearUserDetails()

	appUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.OutputUserDetails(appUser)
	appUser.ClearUserDetails()
	user.OutputUserDetails(appUser)
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

package main

import (
	"fmt"

	"examples.com/struct/user"
)




func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")


	appUser, err :=  user.New(userFirstName, userLastName, userBirthDate)

	if err!= nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("sam@mail.com", "testint123")

	fmt.Println(admin)

	fmt.Println(appUser)
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName: userLastName,
	// 	birthDate: userBirthDate,
	// 	createdAt: time.Now()
	// }
	// ... do something awesome with that gathered data!

appUser.OutputUserDetails()
appUser.ClearUserName()
appUser.OutputUserDetails()
}


// func outputUserDetails(userDetails user){
// 	fmt.Println(userDetails.firstName, userDetails.lastName, userDetails.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

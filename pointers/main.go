package main

import "fmt"

func main() {

	var age int
	var testPointer *int
	agePointer := &age


	fmt.Println("Test Pointer:", testPointer)

	fmt.Println("Age Pointer:", agePointer)

	fmt.Println("Age:", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("New Age:",age)

}

func editAgeToAdultYears(age *int)  {
	

	*age -= 18

}
func getAdultYears(age *int) int  {
	return *age -18
}
// func getAdultYears(age int) int {
// 	return age -18
// }
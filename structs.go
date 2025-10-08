package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name     string 
	Age      uint
	isRemote bool
	Street string
	Address
}

type Address struct {
	Street string `json:"My Street"`
	City string
}


func (e *Employee) updateName(newName string) {
	e.Name = newName
}

func (a Address) printAddress(){
	fmt.Println(a)
}

func main() {
	address := Address{
		Street: "123 Abc Street",
		City: "Ota",
	}
	employee1 := Employee{
		Name:     "Samuel",
		Age:      30,
		isRemote: true,
		Street: "San Francisco",
		Address: address ,
	}

	// job := struct{
	// 	title string
	// 	salary int
	// }{
	// 	title : "Software engineer",
	// 	salary: 100000,
	// }

	// fmt.Println(job.title)


	employee1.updateName("Oluwatobi")

	fmt.Println(employee1)


	// Composition
	fmt.Println("Addresss:" , employee1.Street, employee1.City)
	fmt.Println("Addresss:" , employee1.Address.Street, employee1.City)

	// Composition
	employee1.printAddress()

	// employeePtr := &employee1


	// employeePtr.age = 35

	// fmt.Println(*employeePtr)

	employeeCopy := employee1

	employeeCopy.Age = 40

	fmt.Println(employeeCopy)

	fmt.Println(employee1)

	jsonData, _ := json.Marshal(employee1)
	fmt.Println(string(jsonData))
}
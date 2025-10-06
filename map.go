package main

import "fmt"

func main() {

	var m map[string]int


	m =make(map[string]int)
	m2 := map[string]int{
		"name": 4,
	}

	m["route"] = 14
	m["age"] = 25
	m["cgpa"] = 5.0


	i := m["route"]
	myAge := m["age"]
	test := m2["name"]


	delete(m, "cgpa")

	cgpa,ok := m["cgpa"]
	if ok {
		fmt.Println(cgpa)
		
	}else{
		fmt.Println("Cgpa doesnt exist")
	}


	fmt.Println(i, test, myAge, len(m), cgpa)

	fmt.Println(m)

	for key, val := range m{
		fmt.Printf("key: %v, val: %v", key, val)
	}
}
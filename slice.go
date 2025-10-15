package main

import "fmt"

func main() {
	arrStr := [5]string{"1", "2"}
	fruits := []string{"apple", "banana", "carrot"}
	fruits = append(fruits, "date", "egg")
	fmt.Println(arrStr)
	a := make([]int, 5)

	b := make([]int, 0, 5)
	fmt.Printf("%T", a)
	fmt.Printf("%v" ,b)

	original := [5]int{1, 2, 3, 4, 5}
	shared := original[:3]
	shared[0] = 99

	

	fmt.Println(shared, original)

	// numbers := make([]int, 5) // length 5, capacity 5
	numbers := make([]int, 5, 10) // length 5 , capacity 10
	fmt.Println(numbers)

	fmt.Println(fruits)


}
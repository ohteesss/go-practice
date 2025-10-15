package main

import "fmt"

func main() {
	testStr := "I 💗 you"

	// for i := 0; i < len(testStr); i++ {
	// 	fmt.Printf("%v \n", string(testStr[i]))
	// }
	for _,v := range testStr {
		fmt.Printf("%v", string(v))
	}
	for _,v := range testStr {
		fmt.Printf("%v", v)
	}
}
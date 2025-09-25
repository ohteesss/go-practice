package main

import (
	"fmt"
)

func main() {

	// Type casting
	var anything interface{} = 10

	// testing := int(anything) not possible
	testing := anything.(int)

	fmt.Println(testing)

	if n,ok := anything.(int);ok{
		fmt.Println(n, nil)
	}

	tryingInterface("hello")
    tryingInterface(42)
    tryingInterface(false)



	// numbers
	var int = 48
	var f = float64(int)
	u := uint(f)

	// String and 
	var s  = "hello"
	var b []byte  = []byte(s)
	var s2 = string(b)

	fmt.Println(s,b,s2)

	num := 54
	stinga := "tovi"
	fmt.Println("\t what\"s that about")
	fmt.Printf("tring this stuff out: %v and this other p: %v", string(num), stinga)
	fmt.Println(u)
}


func  tryingInterface(x interface{}){


	switch x.(type){
	case int:
		fmt.Println( "this is an integer")
	case string:
		fmt.Println( "this is a string")
	default:
		fmt.Println( "I no sabi bro")

	}
}
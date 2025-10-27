package main

import "fmt"

type gallon float64

func (g gallon) quart() float64{
	return float64( g * 100)
}

func main() {
	i, j := 42, 2701

	p := &i

	fmt.Println(*p)

	*p = 21

	fmt.Println(i)


	p = &j        
	*p = *p / 37  
	fmt.Println(j) 

	var gallon1 gallon

	gallon1 = 50

	fmt.Println(gallon1.quart())
	
}
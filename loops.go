package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	// sum := 1
	// for i := 0; i < 1000; i++ {
	// 	sum += i
	// }

	// fmt.Println(sum)

	// for sum < 1000 {
	// 	sum += sum
	// }

	// for i := 1; i < 10; i++ {
	// 	if i> 5 {
	// 		break
	// 	}
	// 	fmt.Printf("%d ", i)
	// }

	for i := 0; i < 10; i++{
		if i %2 == 0 {
			continue
		}
		fmt.Printf("%v", i)

	}
	for i:=1; i<= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
		
	}

	// fmt.Println(sqrt(2), sqrt(-4))

	// fmt.Println(sum)
}
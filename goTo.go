package main

import "fmt"

func main() {
    fmt.Println(1)  // Prints the number 1
    goto NEXT        // Jumps to the label "NEXT"
    fmt.Println(2)  // This line is skipped
NEXT:
    fmt.Println(3)  // Prints the number 3
	fmt.Println(4)
//     fmt.Println(1)  // Prints the number 1
//     goto outer        // Jumps to the label "NEXT"
//     fmt.Println(2)  // This line is skipped

// outer:
// 	for i:=1; i < 15; i++ {
// 		for j:= 1; j <= i; j++ {
// 			if i + j > 6{
// 				break outer
// 			}
// 			fmt.Println(i, j)
// 		}
// 	}
fmt.Println("done")
}


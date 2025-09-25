package main

import (
	"fmt"
	"sort"
)

func main() {
	// var strArr [3]string

	// strArr[0] = "a"
	// strArr[1] = "b"
	// strArr[2] = "c"

	// primeNumbers := []int{2,3,13,17,5,7,11}

	// sort.Ints(primeNumbers[:])
	// fmt.Print(primeNumbers)

	// Bubble sort
	// for i := 0; i < len(primeNumbers) ; i++{
	// 	for j := i +1; j < len(primeNumbers); j++{
	// 		if primeNumbers[i] > primeNumbers[j] {
	// 			primeNumbers[i], primeNumbers[i+1] = primeNumbers[i+1], primeNumbers[i]
	// 		}
	// 	}

	// }

	// Array copy
	// arrCopy1 := [5]int{1,3,5,7}
	// arrCopy2 := [5]int{2,5,6,1,2}

	// n := copy(arrCopy1[:],arrCopy2[:])

	// fmt.Println(arrCopy1)
	// fmt.Println(arrCopy2)

	// fmt.Println(n)

	// fmt.Println(primeNumbers)
	// fmt.Print(strArr)

	// Array delete in Golang
	// arrDel := [4]int{2,4,5,6}
	// arrDel[2] = 0
	// fmt.Println(arrDel)

	// Array contains or include, return index else returns the length of the array
	// arr := [4]int{2,4,6,8}
	// i := sort.Search(len(arr), func(index int) bool {
	// 	return arr[index] == 7
	// })
	// fmt.Println(i)

	// Array reverse
	arr := [5]int{2, 4, 6, 8, 12}
	// for i, j := 0, len(arr) - 1; i < j; i, j = i+1, j-1 {
	// 	arr[i], arr[j] = arr[j], arr[i]
	// }

	sort.Sort(sort.Reverse(sort.IntSlice(arr[:])))



	fmt.Println(arr)

}
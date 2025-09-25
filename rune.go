package main

import "fmt"

func main() {
	str := "testing Go"
	str2 := "💞 code"


	fmt.Println(str[0])
	fmt.Println(string(str2[0]))
	for _, v := range str {
		fmt.Println(string(v))
	}

	for i,v := range str2{
		if i==0{

			fmt.Println(string(v))
		}
	}

	slicingStr()
	
}


func slicingStr(){
	// strings are read-only slice of bytes
	hello := "hello"
	s := "😂"
	runes := []rune("🍰Have your cake and eat it")
	fmt.Println("Length of s:" ,len(s)) //Length of s: 4(bytems)
	fmt.Println(s[0]) //240(first byte)

	fmt.Printf("%c\n ", runes[0]) //🍰 

	fmt.Printf("%c\n%c\n", s[0], hello[1])

}
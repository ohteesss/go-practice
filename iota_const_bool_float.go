package main

import "fmt"

const name string = "Oluwatobi";

const (
	Sunday = iota
	Monday
	Tuesday
//
	Wednesday
_
	Thursday
)

type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func printer() {
	// fmt.Print(name, GB)
    fmt.Print(GB);

}

func main() {
    var f float32 
    var bVal bool

    // var f2 float64 
    f = 12.34567890123
    compl := complex(11, 10)
    // f2 = 12.34567890123456
    fmt.Printf("bVal: %v \n ",  bVal)
    fmt.Println(f, compl)  // prints "12.345679 12.34567890123456"
    printer()
}
package main

import (
	"fmt"
	"math"
)

// func main() {
// 	var investmentAmount float64 = 1000
// 	var expectedReturnRate = 5.5
// 	var years float64 = 10

// 	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate /100, years)
// 	fmt.Println(futureValue)
// }


func main(){
	 var investmentAmount, expectedReturnRate, years float64
	 

	 fmt.Print("Enter Investment Amount: ")
	 fmt.Scan(&investmentAmount);
	 fmt.Print("Expected Return Rate: ")
	 fmt.Scan(&expectedReturnRate)
	 fmt.Print("Years: ")
	 fmt.Scan(&years)
	 futureValue :=  investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	fmt.Println(futureValue)

}

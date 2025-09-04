package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

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

	futureValue,formattedFV := calculateFutureValues(investmentAmount,expectedReturnRate, years)
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Print(formattedFV)
	fmt.Println(futureValue)

}




// func calculateFutureValues(investmentAmount,expectedReturnRate, years float64) (float64, float64){
// 	fv := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
// 	rfv := fv / math.Pow(1+inflationRate/100, years)

// 	return fv,rfv
// }
func calculateFutureValues(investmentAmount,expectedReturnRate, years float64) (fv float64, rfv float64){
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)


	return fv, rfv
	// return
}


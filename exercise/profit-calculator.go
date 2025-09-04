package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// var revenue, expenses, taxRate float64
	revenue,err1 := getUserInput("Revenue:")
	if err1 != nil {
		fmt.Println("ERROR")
		fmt.Println(err1)
		fmt.Println("-------")
		return
	 }

	expenses, err2 := getUserInput("Expenses:")
	if err2 != nil {
		fmt.Println("ERROR")
		fmt.Println(err2)
		fmt.Println("-------")
		return
	 }
	taxRate,err3 := getUserInput("Tax Rate:")
	if err3 != nil {
		fmt.Println("ERROR")
		fmt.Println(err3)
		fmt.Println("-------")
		return
	 }
	


	EBT,EAT,ratio := calculateFinance(revenue,expenses, taxRate)

	



	fmt.Printf("Earnings Before Tax: %.2f, Earnings After Tax: %.2f, Ration: %.2f ", EBT, EAT, ratio)




}

func storeResults(ebt,eat,ratio float64){
	stringValue := fmt.Sprintf("%.1f %.1f %.1f", ebt, eat, ratio)

	os.WriteFile("calculated-finance.txt", []byte(stringValue), 0644)

}

func getUserInput(text string) (float64, error){
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("user input cannot be negative")
	}
	return userInput, nil
}

func calculateFinance (revenue,expenses, taxRate float64)(float64,float64, float64){
	
	EBT := revenue - expenses
	EAT := EBT * (100 - taxRate)/100
	ratio := EBT/EAT


	return EBT,EAT, ratio
}
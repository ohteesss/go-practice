package main

import "fmt"

func main() {

	// var revenue, expenses, taxRate float64
	revenue := getUserInput("Revenue:")
	expenses := getUserInput("Expenses:")
	taxRate := getUserInput("Tax Rate:")
	


	EBT,EAT,ratio := calculateFinance(revenue,expenses, taxRate)

	



	fmt.Printf("Earnings Before Tax: %.2f, Earnings After Tax: %.2f, Ration: %.2f ", EBT, EAT, ratio)




}

func getUserInput(text string) float64{
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinance (revenue,expenses, taxRate float64)(float64,float64, float64){
	
	EBT := revenue - expenses
	EAT := EBT * (100 - taxRate)/100
	ratio := EBT/EAT

	return EBT,EAT, ratio
}
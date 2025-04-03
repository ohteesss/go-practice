package main

import "fmt"

func main() {

	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	EAT := EBT * (100 - taxRate)/100
	ratio := EBT/EAT



	fmt.Println("Earnings Before Tax, Earnings After Tax, Ration: ", EBT, EAT, ratio)




}
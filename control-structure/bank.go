package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"



func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func writeFloatToFile(value float64, filleName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filleName, []byte(valueText), 0644)
}

func main(){
	 accountBalance, err := getFloatFromFile(balanceFile)

	 if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		// panic("Sorry we can continue!")
	 }

	fmt.Println("Welcome to Go Bank!")
	for  {
		displayMenu()
	
		var choice int
	
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	
		// if choice == 1 {
		// 	fmt.Println("Your account balance is:" , accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
	
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
	
		// 		continue
		// 	}else {
		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount: ", accountBalance)}
		// }else if choice == 3 {
		// 	fmt.Print("How much do you want to withdraw: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)
			
		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
	
		// 		continue
		// 	}
		// 	 if withdrawalAmount > accountBalance {
		// 		fmt.Println("Invalid amount: Withdraw Amount can not be greater than account balance")
	
		// 		continue
		// 	}
		// 	accountBalance -= withdrawalAmount
		// 	fmt.Println("Balance updated! New amount: ", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }

		switch choice {
		case 1: fmt.Println("Your account balance is:" , accountBalance)
		case 2: 
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
	
				continue
			}else {
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)}
			writeFloatToFile(accountBalance, balanceFile)
		case 3:
			fmt.Print("How much do you want to withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
	
				continue
			}
			 if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount: Withdraw Amount can not be greater than account balance")
	
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeFloatToFile(accountBalance, balanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for banking with us!")
			return
		}
		
	

	}

	



}



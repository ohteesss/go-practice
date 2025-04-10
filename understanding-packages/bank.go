package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"packages.com/bank/fileops"
)

const balanceFile = "balance.txt"


func main(){
	 accountBalance, err := fileops.GetFloatFromFile(balanceFile)

	 if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		// panic("Sorry we can continue!")
	 }

	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())
	for  {
		presentOptions()
	
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
			fileops.WriteFloatToFile(accountBalance, balanceFile)
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
			fileops.WriteFloatToFile(accountBalance, balanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for banking with us!")
			return
		}
		
	

	}

	



}



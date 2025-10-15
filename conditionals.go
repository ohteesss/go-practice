package main

import "fmt"

func main() {

	var grade int
	var answer int

	var day int

	var num int

	fmt.Println("What's your input grade?")

	fmt.Scan(&grade)

	if grade > 60 {
		fmt.Println("You passed the exam")
	}else if grade == 60{
		fmt.Println("You scored the cut off score, Let's do a quick test!")
		fmt.Println("What is 50 - 5")
		fmt.Scan(&answer)
		if(answer == 45){
			fmt.Println("You have now passed the exam")
		}else {
			fmt.Println("You were given another chance, but you still failed the exam.")
		}

	}else{
		fmt.Println("Oops! You failed the exam")
	}

	fmt.Println("Input a day of the week ranging from 1 to 7")
	fmt.Scan(&day)

	switch day {
	case 1,2,3,4,5:
		fmt.Println("You should be working today!")
	case 6:
		fmt.Println("Have a lovely Weekend!")
	case 7:
		fmt.Println("Go to Church!")
	default:
		fmt.Println("Use your head and input the right day")
	}

	fmt.Println("Input any number!")

	fmt.Scan(&num)

	switch  {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num > 200:
		fmt.Printf("%d is lesser than 200", num) // Fallthrough happens even when the case evaluates to false
	}



}
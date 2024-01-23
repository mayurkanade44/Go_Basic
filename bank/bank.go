package main

import "fmt"

func main() {
	accountBalance := 1000.0
	fmt.Println("Welcome to go bank")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. exit")

		var choice int
		fmt.Print("Please enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is", accountBalance)
		} else if choice == 2 {
			fmt.Println("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Print("Your updated account balance is ", accountBalance)
		} else if choice == 3 {
			fmt.Println("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Print("Your account balance low")
			} else {
				accountBalance -= withdrawAmount
				fmt.Print("Your updated account balance is ", accountBalance)
			}
		} else {
			fmt.Print("Thank you for banking, goodbye")
			break
		}
	}

}

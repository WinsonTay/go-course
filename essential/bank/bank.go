package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.json"

func main() {
	accountBalance, err := fileops.GetBalanceFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------------")
		panic("Exit")
	}
	fmt.Println("Welcome to Hell Bank")
	for {
		presentOptions()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Balance: %.1f", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("How Much do you want deposit?")
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fileops.WriteBalanceToFile(accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Enter Withdrawal Amount: ")
			fmt.Scan(&withdrawalAmount)
			if accountBalance-withdrawalAmount < 0 {
				fmt.Print("You do not have enough balance to withdraw that amount")
				continue
			} else {
				accountBalance -= withdrawalAmount
				fileops.WriteBalanceToFile(accountBalance)
				fmt.Println("You have withdraw the money, New Balance: ", accountBalance)
			}
		default:
			fmt.Println("Goodbye !")
			return
		}
	}

}

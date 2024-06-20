package main
import (
	"fmt"
)

func main() {
	// Declare a variable to receive user input
	key := ""
	loop := true

	balance := 10000.0
	money := 0.0
	note := ""
	details := "Transaction\tBalance\tAmount\tDetails"

	for {

		fmt.Println("--------------")
		fmt.Println("1. Details")
		fmt.Println("2. Record Incomes")
		fmt.Println("3. Record Spending")
		fmt.Println("4. Exit")
		fmt.Println("Please choose from 1 to 4")

		fmt.Scanln(&key)
		switch key {
			case "1":
				fmt.Println("The current account details:")
				fmt.Println(details)
			case "2":
				fmt.Println("Transaction amount:")
				fmt.Scanln(&money)
				balance += money
				fmt.Println("Transaction notes:")
				fmt.Scanln(&note)
				// concate info to details
				details += fmt.Sprintf("\nIncome\t%v\t%v\t%v", balance, money, note)
			case "3":
				fmt.Println("Spending amount:")
				fmt.Scanln(&money)
				if money > balance {
					fmt.Println("Insufficient fund")
					break
				}
				balance -= money
				fmt.Println("Spending notes:")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\nSpending\t%v\t%v\t%v", balance, money, note)
			case "4":
				fmt.Println("Are you sure to quit? y/n")
				choice := ""
				for {
					fmt.Scanln(&choice)
					if choice == "y" || choice == "n" {
						break
					}
					fmt.Println("Wrong input")
				}
				if choice == "y" {
					loop = false
				}
			default:
				fmt.Println("Please make correct input")
				break
		}
		if !loop {
			break
		}
	}
	fmt.Println("You have exited the system")



}
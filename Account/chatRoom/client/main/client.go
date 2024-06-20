package main
import (
	"fmt"
	"os"
	"chatRoom/client/process"
)

// define 2 variables, one stores user ID, another stores password
var userId int
var userPassword string
var userName string

func main() {
	// receive user's choice
	var key int
	for true {
		fmt.Println("-------------------- Welcome to the chat channel --------------------")
		fmt.Println("\t\t\t 1. Login chat channel")
		fmt.Println("\t\t\t 2. Create an account")
		fmt.Println("\t\t\t 3. Exit the system")
		fmt.Println("\t\t\t Please choose from 1 to 3")
		fmt.Scanf("%d\n", &key)
		switch key {
			case 1:
				fmt.Println("Login chat channel")
				fmt.Println("Please enter your user ID")
				fmt.Scanf("%d\n", &userId)
				fmt.Println("Please enter your user password")
				fmt.Scanf("%s\n", &userPassword)

				userProcess := &process.UserProcess {}
				userProcess.Login(userId, userPassword)
			case 2:
				fmt.Println("Create an account")
				fmt.Println("Please enter an user ID")
				fmt.Scanf("%d\n", &userId)
				fmt.Println("Please enter an user password")
				fmt.Scanf("%s\n", &userPassword)
				fmt.Println("Please enter an username")
				fmt.Scanf("%s\n", &userName)

				userProcess := &process.UserProcess {}
				userProcess.Register(userId, userPassword, userName)
			case 3:
				fmt.Println("Exit the system")
				os.Exit(0)
			default:
				fmt.Println("Wrong input, please enter again!")
		}
	}
}
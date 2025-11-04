package main

import (
	"fmt"
)

func checkName() {
	var name string
	fmt.Println("Welcome to the game, Please enter your name")
	
	if _, err := fmt.Scanln(&name); err != nil {
		fmt.Println("Invalid name:", err)
	}else{
		fmt.Printf("Hello %s welcome to the game!\n", name)
	}


}

func checkBalance(balance int)(int, bool){
	var bet int
	fmt.Printf("Enter your bet, or 0 to quit (balance =$%d)\n", balance)
	fmt.Scanln(&bet)

	if bet == 0{
		fmt.Println("Thank you for playing the game")
		return balance, true
	}

	if bet > balance{
		fmt.Println("Can't place bet bigger than your balance!")
		return balance, false
	}

	balance -= bet
	fmt.Printf("You placed %d bet, have $%d left in the balance",bet,  balance)
	return balance, false

}
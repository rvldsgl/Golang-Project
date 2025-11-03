package main

import(
	"fmt"
)

func welcome (){
	fmt.Println("Welcome to the game, Please enter your name")
	var name string
	var accountBalance int = 200
	gameContinue := true

	for gameContinue {
		if _, err := fmt.Scanln(&name); err != nil {
			fmt.Println("Invalid name:", err)
			gameContinue = false
		} 
		

		fmt.Printf("Hello %s your balance is %d$\n",name, accountBalance)

	// fmt.Printf("Please enter yout bet, or 0 to quit (balance =%d$): ", accountBalance)
	// fmt.Scanln(&accountBalance)
}
}

func main(){
	welcome()
	
}
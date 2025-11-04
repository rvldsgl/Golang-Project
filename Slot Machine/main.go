package main

import(
	"fmt"
)




func main() {
    checkName()
    balance := 200
    gameLoop := true

    for gameLoop {
        newBalance, shouldEnd := checkBalance(balance)
        if shouldEnd {
            gameLoop = true
            break
        }
				balance = newBalance
        
    }
		
    
    fmt.Println("Game Over! Final balance:", balance)
}
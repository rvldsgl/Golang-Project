package main
import (
	// "fmt"
	"math/rand"
)

var symbols = map[string]uint{
    "A": 4,   // lebih jarang muncul → lebih berharga
    "B": 7,
    "C": 12,
    "D": 20,  // paling sering muncul → paling murah
}

var multipliers = map[string]uint{
    "A": 20,  // karena jarang, pengali besar
    "B": 10,
    "C": 5,
    "D": 2,   // sering muncul, pengali kecil
}


func generateSymbolArray(symbols map[string]uint) []string{
	symbolArray := make([]string, 0)
	for key, count:= range symbols{
		var i uint = 1
		for i <= count{
			symbolArray = append(symbolArray,key)
			i += 1
		}
		
	}
	return symbolArray
}

func createBoard() [][]string{
	board := make([][]string, 3)
	for i:= range board{
		board[i] = make([]string, 3)
	}

	symbolArray := generateSymbolArray(symbols)
	for i,j := range board{
		for k := range j{
			board[i][k] = symbolArray[rand.Intn(len(symbolArray)-1)]
		}
	}
	return board
}


// func spinWheel(){
	
// }

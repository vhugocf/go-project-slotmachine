package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
	
	name := ""
	//var name string

	
	fmt.Println("Welcome to Victor Casino")
	fmt.Printf("Type your name: ")
	_, err := fmt.Scanln(&name)

	if err != nil{
		return ""
	}

	fmt.Printf("Welcome %s, let's play the game by the rules!\n", name)
	return name
}

func getBet(balance uint) uint{
	var bet uint
	// while in go
	for true{
		fmt.Printf("Enter your bet, or type '0' to quit this game (your balance =$%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}// end bet
	return bet

}// end getBalance


//complex logic
// A - very rare
// B - rare
// C - least rare
// D - normal

func generateSymbolsArray(symbols map[string]uint) []string{
	
	symbolsArr := []string{}
	for symbols, count := range symbols {
		for i := uint(0); i < (count); i++{
			symbolsArr = append(symbolsArr, symbols)
		}
	}
	return symbolsArr
}

//generate a randon number
func getRandonNumber(min int, max int) int {
	randonNumber := rand.Intn(max - min  + 1) + min
	return randonNumber
}


func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++{
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++{
		selected := map[int]bool{}
		for row := 0; row < rows; row++{
			for true {
				randonIndex := getRandonNumber(0, len(reel)-1)
				_, exist := selected[randonIndex]
				if !exist{
					selected[randonIndex] = true
					result[row] = append(result[row], reel[randonIndex])
					break
				}
			}
		}
	}
	return result
}




func main(){

	symbols := map[string]uint{
		"A" : 2,
		"B" : 10,
		"C" : 16,
		"D" : 31,
	}
	// what money you will win with your bet
	/*
	multipliers := map[string] uint{

		"A" : 30,
		"B" : 15,
		"C" : 7,
		"D" : 2,
	}
	*/
	symbolsArr := generateSymbolsArray(symbols)
	spin := getSpin(symbolsArr, 3,3)

	fmt.Println(spin)
	
	
	balance := uint(200)
	getName()

	for balance > 0{
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance = balance - bet
	}
	fmt.Printf("Your left with, $%d.\n", balance)
}










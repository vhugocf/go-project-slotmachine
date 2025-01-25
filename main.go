package main

import(

	"fmt"
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

func main(){

	symbols := map[string]uint{
		"A" : 3,
		"B" : 6,
		"C" : 11,
		"D" : 25,
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
	fmt.Println(symbolsArr)
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










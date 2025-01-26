package main

import "fmt"



func GetName() string {
	
	name := ""
	//var name string

	
	fmt.Println("Welcome to Victor's Casino")
	fmt.Printf("Type your name: ")
	_, err := fmt.Scanln(&name)

	if err != nil{
		return ""
	}

	fmt.Printf("Welcome %s, let's play the game by the rules!\n", name)
	return name
}

func GetBet(balance uint) uint{
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

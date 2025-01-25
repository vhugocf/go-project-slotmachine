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







func main(){

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










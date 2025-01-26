package main

import (
	"fmt"
	"math/rand"

	//"golang.org/x/text/unicode/rangetable"
)


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


func printSpin(spin [][]string){
	for _, row := range spin{
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row) - 1{
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}


func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin{
		win := true
		checkSymbol := row[0]

		for _, symbol := range row[1:]{
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		}else {
			lines = append(lines, 0)
		}
	
	}
	return lines
}



func main(){

	symbols := map[string]uint{
		"A" : 2,
		"B" : 10,
		"C" : 16,
		"D" : 31,
	}
	// what money you will win with your bet

	multipliers := map[string] uint{

		"A" : 30,
		"B" : 15,
		"C" : 7,
		"D" : 2,
	}

	symbolsArr := generateSymbolsArray(symbols)
	

	
	balance := uint(200)
	GetName()

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}


		balance = balance - bet
		spin := getSpin(symbolsArr, 3,3)
		printSpin(spin)
		
		//here we check if you win and check balance
		winningLines := checkWin(spin, multipliers)

		for i, multi := range winningLines{
			win := multi * bet
			balance = balance + win
			if multi > 0 {
				fmt.Printf("Won $%d, (%dx) on Line #%d\n", win, multi, i + 1)
			}
			}
		}
		fmt.Printf("Your left with, $%d.\n", balance)



	}// end main













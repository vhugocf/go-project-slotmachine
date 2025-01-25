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




func main(){

	getName()
	


}

package main

import (
	"fmt"
	"number-guessing-game/service"
)


func main(){
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.")
	fmt.Printf("\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n\nEnter your choice: ")
	var choice int
	if _,err:=fmt.Scanln((&choice));err!=nil{
		fmt.Println("error: %w",err)
		return
	}

	for{
		if err:=service.Game(choice); err!=nil{
			fmt.Printf("%v\n",err)
		}
		
		var playAgain string
		fmt.Printf("\nDo you wanna play again? [y/n] : ")
		if _,err:=fmt.Scanln(&playAgain); err!=nil{
			fmt.Printf("%v\n",err)
		}
		
		switch playAgain {
		case "y":
			continue
		case "Y":
			continue
		case "n":
			fmt.Println("Thank you for playing!")
			return
		case "N":
			fmt.Println("Thank you for playing!")
			return
		default:
			fmt.Println("invalid option, aborted the game")
			return
		}
	}
}
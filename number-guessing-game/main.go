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
	err:=service.Game(choice)
	if err!=nil{
		fmt.Printf("%v\n",err)
	}
}
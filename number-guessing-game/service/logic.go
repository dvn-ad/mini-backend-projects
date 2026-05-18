package service

import (
	"fmt"
	"math/rand/v2"
	"number-guessing-game/storage"
)

func Game(choice int)error{
	var chances int
	switch choice {
	case 1:
		chances = 10
	case 2:
		chances=5
	case 3:
		chances=3
	default:
		return fmt.Errorf("invalid choice")
	}
	choices:=[]string{"Easy", "Medium", "Hard"}
	fmt.Printf("\nGreat! You have selected the %s difficulty level.\nLet's start the game!\n",choices[choice-1])
	
	answerBot:=rand.IntN(100) + 1
	var answerUser int

	for i:=range(chances){
		fmt.Printf("\nEnter your guess: ")
		if _,err:=fmt.Scanln((&answerUser));err!=nil{
			return err
		}

		if answerUser==answerBot{
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n",i+1)
			if err:=checkHighest(i+1,choice);err!=nil{
				return err
			}

			return nil
		}
		if answerUser<answerBot{
			fmt.Printf("Incorrect! The number is greater than %d.\n",answerUser)
		}else if answerUser>answerBot{
			fmt.Printf("Incorrect! The number is less than %d.\n",answerUser)
		}

		chances--
	}
	return nil
}


func checkHighest(score int, diff int)error{
	scores ,err:= storage.LoadData()
	if err!=nil{
		return err
	}

	switch diff {
	case 1:
		if scores.Easy>score{
			scores.Easy=score
		}
	case 2:
		if scores.Medium>score{
			scores.Medium=score
		}
	case 3:
		if scores.Medium>score{
			scores.Medium=score
		}
	}

	err=storage.SaveData(scores)
	if err!=nil{
		return err
	}
	return nil
}
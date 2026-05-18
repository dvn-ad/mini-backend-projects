package service

import (
	"number-guessing-game/storage"
)



func checkHighest(score int, diff int)error{
	scores ,err:= storage.LoadData()
	if err!=nil{
		return err
	}

	switch diff {
	case 1:
		if scores.Easy<score{
			scores.Easy=score
		}
	case 2:
		if scores.Medium<score{
			scores.Medium=score
		}
	case 3:
		if scores.Medium<score{
			scores.Medium=score
		}
	}

	err=storage.SaveData(scores)
	if err!=nil{
		return err
	}
	return nil
}
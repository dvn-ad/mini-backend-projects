package storage

import (
	"number-guessing-game/model"
	"os"
	"encoding/json"
)

const path = "data/data.json"

func LoadData()(model.Highscore, error){
	if _,err:=os.Stat(path);os.IsNotExist(err){
		return model.Highscore{}, err
	}

	data,err:=os.ReadFile(path)
	if err!=nil{
		return model.Highscore{}, err
	}

	var scores model.Highscore
	if len(data) == 0 {
		return model.Highscore{}, nil
	}
	err = json.Unmarshal(data, &scores)
	if err != nil {
		return model.Highscore{}, err
	}
	return scores, nil
}

func SaveData(scores model.Highscore)error{
	data, err := json.MarshalIndent(scores, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path,data,0644)
	if err!=nil{
		return err
	}
	return nil
}
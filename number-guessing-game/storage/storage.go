package storage

import (
	"encoding/json"
	"number-guessing-game/model"
	"os"
)

const path = "data/data.json"

func LoadData() (model.Highscore, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return model.Highscore{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return model.Highscore{}, err
	}

	var scores []model.Highscore
	if len(data) == 0 {
		return model.Highscore{}, nil
	}
	err = json.Unmarshal(data, &scores)
	if err != nil {
		return model.Highscore{}, err
	}
	if len(scores) > 0 {
		return scores[0], nil
	}
	return model.Highscore{}, nil
}

func SaveData(scores model.Highscore) error {
	data, err := json.MarshalIndent([]model.Highscore{scores}, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
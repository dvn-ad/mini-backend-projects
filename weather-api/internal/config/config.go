package config

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	Port			string
	RedisURL		string
	WeatherAPIKey	string
}

func Load()(*Config, error){
	_=godotenv.Load()

	cfg:=&Config{
		Port: os.Getenv("PORT"),
		RedisURL: os.Getenv("REDIS_URL"),
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
	}

	if cfg.Port==""{
		cfg.Port="8080"
	}
	if cfg.RedisURL==""{
		return nil, fmt.Errorf("REDIS_URL environment variable is required")
	}
	if cfg.WeatherAPIKey==""{
		return nil, fmt.Errorf("WEATHER_API_KEY environment variable is required")
	}
	return cfg,nil
}
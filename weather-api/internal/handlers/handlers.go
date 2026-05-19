package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"weather-api/internal/cache"
	"weather-api/internal/weather"

	"github.com/redis/go-redis/v9"
)

type WeatherHandler struct {
	weatherClient *weather.Client
	cache         *cache.Cache
}

func NewWeatherHandler(wc *weather.Client, ch *cache.Cache) *WeatherHandler {
	return &WeatherHandler{
		weatherClient: wc,
		cache:         ch,
	}
}

func (h *WeatherHandler) HandleWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	city := r.URL.Query().Get("city")
	city = strings.TrimSpace(strings.ToLower(city))

	if city == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "Missing required query parameter: 'city'"}`))
		return
	}

	cacheKey := "weather:" + city
	ctx := r.Context()

	// 1. Cek Cache Redis
	cachedData, err := h.cache.Get(ctx, cacheKey)
	if err == nil {
		w.Header().Set("X-Cache", "HIT")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(cachedData))
		return
	}

	if !errors.Is(err, redis.Nil) {
		log.Printf("Redis error (bypass to database/API): %v", err)
	}

	// 2. Cache Miss: Panggil Resty Client ke Visual Crossing
	weatherData, err := h.weatherClient.FetchWeather(ctx, city)
	if err != nil {
		if strings.Contains(err.Error(), "invalid city") {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "Invalid city or region code provided"}`))
			return
		}
		log.Printf("External API failure: %v", err)
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(`{"error": "Failed to fetch weather data from external provider"}`))
		return
	}

	// 3. Serialize Struct ke JSON String
	jsonData, err := json.Marshal(weatherData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal serialization error"}`))
		return
	}

	// 4. Set ke Redis Cache (TTL: 12 Jam)
	err = h.cache.Set(ctx, cacheKey, string(jsonData), 12*time.Hour)
	if err != nil {
		log.Printf("Failed to write to Redis cache: %v", err)
	}

	w.Header().Set("X-Cache", "MISS")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonData)
}
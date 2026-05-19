package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"weather-api/internal/cache"
	"weather-api/internal/config"
	"weather-api/internal/handlers"
	"weather-api/internal/middleware"
	"weather-api/internal/weather"

	"golang.org/x/time/rate"
)

func main() {
	// 1. Load Config (.env)
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Configuration loading failed: %v", err)
	}

	// 2. Init Redis Connection
	cacheService, err := cache.New(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	log.Println("Connected to Redis successfully.")

	// 3. Init Clients & Handlers
	weatherClient := weather.NewClient(cfg.WeatherAPIKey)
	weatherHandler := handlers.NewWeatherHandler(weatherClient, cacheService)

	// 4. Init Router (Go 1.22+ syntax)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /weather", weatherHandler.HandleWeather)

	// 5. Setup Rate Limiter (5 req/sec, burst 10)
	limiter := middleware.NewIPRateLimiter(rate.Limit(5), 10)
	rateLimitedMux := middleware.RateLimit(limiter)(mux)

	// 6. Define Server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      rateLimitedMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 7. Run Server & Graceful Shutdown
	go func() {
		log.Printf("Weather Server listening on port %s...", cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server listen failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited safely.")
}
package main

import (
	"log"
	"net/http"

	"github.com/ZeroZeroZerooZeroo/WeatherAPIService/internal/config"
	"github.com/ZeroZeroZerooZeroo/WeatherAPIService/internal/handler"
	"github.com/ZeroZeroZerooZeroo/WeatherAPIService/internal/service"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)

	}

	WeatherService := service.NewWeatherService(
		&http.Client{},
		cfg.ApiKey,
		cfg.ApiURL,
	)

	weatherHandler := handler.NewHandler(WeatherService)

	http.HandleFunc("/weather", weatherHandler.GetWeather)

	log.Printf("Сервер запущен на порту %s", cfg.ServerPort)

	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

}

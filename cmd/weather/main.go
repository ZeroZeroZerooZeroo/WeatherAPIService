package main

import (
	"log"

	"github.com/ZeroZeroZerooZeroo/WeatherAPIService/internal/config"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)

	}

}

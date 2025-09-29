package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ZeroZeroZerooZeroo/WeatherAPIService/internal/service"
)

type WeatherHandler struct {
	weatherService *service.WeatherService
}

func NewHandler(ws *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: ws,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Параметр 'city' обязателен", http.StatusBadRequest)
		return
	}

	data, err := h.weatherService.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Ошибка кодирования ответа", http.StatusInternalServerError)
	}
}

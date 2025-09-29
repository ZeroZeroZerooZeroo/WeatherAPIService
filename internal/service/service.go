package service

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ZeroZeroZerooZeroo/WeatherAPIService/internal/models"
)

type WeatherService struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

func NewWeatherService(httpClient *http.Client, apyKey, baseURL string) *WeatherService {
	return &WeatherService{
		httpClient: httpClient,
		apiKey:     apyKey,
		baseURL:    baseURL,
	}
}

func (ws *WeatherService) GetWeather(city string) (*models.WeatherResponse, error) {

}

func (ws *WeatherService) buildURL(city string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ru",
		ws.baseURL, url.QueryEscape(city), ws.apiKey)
}

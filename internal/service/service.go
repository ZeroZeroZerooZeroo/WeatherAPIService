package service

import (
	"encoding/json"
	"fmt"
	"io"
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
	url := ws.buildURL(city)
	resp, err := ws.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса:%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API ошибка: %s, тело: %s", resp.Status, string(body))

	}

	var apiResponse models.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)

	}
	if len(apiResponse.Weather) == 0 {
		return nil, fmt.Errorf("нет данных о погоде")
	}

	weather := &models.WeatherResponse{
		City:        apiResponse.Name,
		Temperature: apiResponse.Main.Temp,
		Description: apiResponse.Weather[0].Description,
	}
	return weather, nil
}

func (ws *WeatherService) buildURL(city string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ru",
		ws.baseURL, url.QueryEscape(city), ws.apiKey)
}

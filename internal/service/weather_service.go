package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"on_esports/internal/model"

	"github.com/go-resty/resty/v2"
)

type WeatherService interface {
	GetWeatherInfo(city string) (*model.WeatherData, error)
}

type weatherService struct {
	apiKey string
}

func NewWeatherService(apiKey string) WeatherService {
	return &weatherService{apiKey: apiKey}
}

func (ws *weatherService) GetWeatherInfo(city string) (*model.WeatherData, error) {
	client := resty.New()

	resp, err := client.R().Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, ws.apiKey))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode())
	}

	if len(resp.Body()) == 0 {
		return nil, errors.New("API response body is empty")
	}

	var weatherData model.WeatherData
	if err := json.Unmarshal(resp.Body(), &weatherData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return &weatherData, nil
}

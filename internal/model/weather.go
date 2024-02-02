package model

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type WeatherInfo struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainInfo struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type WindInfo struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type cloudsInfo struct {
	All int `json:"all"`
}

type sysInfo struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherData struct {
	Coord      Coord         `json:"coord"`
	Weather    []WeatherInfo `json:"weather"`
	Base       string        `json:"base"`
	Main       MainInfo      `json:"main"`
	Visibility int           `json:"visibility"`
	Wind       WindInfo      `json:"wind"`
	Clouds     cloudsInfo    `json:"clouds"`
	Dt         int64         `json:"dt"`
	Sys        sysInfo       `json:"sys"`
	Timezone   int           `json:"timezone"`
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Cod        int           `json:"cod"`
}

func (w *WeatherData) ConvertWeather(city string) *discordgo.MessageEmbed {
	if len(w.Weather) == 0 {
		return nil
	}

	var sunriseTime, sunsetTime string
	if w.Sys.Sunrise > 0 && w.Sys.Sunset > 0 {
		sunriseTime = time.Unix(w.Sys.Sunrise, 0).Format("15:04:05 MST")
		sunsetTime = time.Unix(w.Sys.Sunset, 0).Format("15:04:05 MST")
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Current Weather in " + city,
		Description: fmt.Sprintf("**%s**\n*%s*", w.Weather[0].Description, city),
		Color:       0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Temperature", Value: fmt.Sprintf("%0.2f°C", kelvinToCelsius(w.Main.Temp)), Inline: true},
			{Name: "Feels Like", Value: fmt.Sprintf("%0.2f°C", w.Main.FeelsLike), Inline: true},
			{Name: "Pressure", Value: fmt.Sprintf("%d hPa", w.Main.Pressure), Inline: true},
			{Name: "Humidity", Value: fmt.Sprintf("%d%%", w.Main.Humidity), Inline: true},
		},
	}

	if w.Wind.Speed > 0 && w.Wind.Deg > 0 {
		embed.Fields = append(embed.Fields,
			&discordgo.MessageEmbedField{Name: "Wind Speed", Value: fmt.Sprintf("%0.2f m/s at %d°", w.Wind.Speed, w.Wind.Deg), Inline: true})
	}

	if w.Clouds.All > 0 {
		embed.Fields = append(embed.Fields,
			&discordgo.MessageEmbedField{Name: "Cloudiness", Value: fmt.Sprintf("%d%%", w.Clouds.All), Inline: true})
	}

	if w.Visibility > 0 {
		embed.Fields = append(embed.Fields,
			&discordgo.MessageEmbedField{Name: "Visibility", Value: fmt.Sprintf("%d meters", w.Visibility), Inline: true})
	}

	if len(sunriseTime) > 0 {
		embed.Fields = append(embed.Fields,
			&discordgo.MessageEmbedField{Name: "Sunrise", Value: sunriseTime, Inline: true})
	}

	if len(sunsetTime) > 0 {
		embed.Fields = append(embed.Fields,
			&discordgo.MessageEmbedField{Name: "Sunset", Value: sunsetTime, Inline: true})
	}

	return embed
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

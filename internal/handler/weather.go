package handler

import (
	"on_esports/internal/logger"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) handleWeatherCommand(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	if len(params) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Usage: weather <city>")
		return
	}

	city := strings.Join(params, " ")

	weatherInfo, err := h.WeatherService.GetWeatherInfo(city)
	if err != nil {
		logger.Errorf("h.WeatherService.GetWeatherInfo(): %v", err)
		s.ChannelMessageSend(m.ChannelID, "Internal server")
		return
	}

	embed := weatherInfo.ConvertWeather(city)

	if embed != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	} else {
		s.ChannelMessageSend(m.ChannelID, "No weather information available")
	}
}

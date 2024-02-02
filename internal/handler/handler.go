// handler/handler.go
package handler

import (
	"on_esports/internal/service"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	DiscordSession *discordgo.Session
	PollService    service.PollService
	WeatherService service.WeatherService
}

func New(discordSession *discordgo.Session, pollService service.PollService, weatherService service.WeatherService) *Handler {
	return &Handler{
		DiscordSession: discordSession,
		PollService:    pollService,
		WeatherService: weatherService,
	}
}

func (h *Handler) Start() error {
	h.DiscordSession.AddHandler(h.handleMessageCreate)

	return h.DiscordSession.Open()
}

func (h *Handler) Stop() {
	h.DiscordSession.Close()
}

package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	fields := strings.Fields(m.Content)
	if len(fields) == 0 {
		return
	}

	switch command := fields[0]; command {
	case "!help":
		h.handleHelpCommand(s, m)
	case "!weather":
		h.handleWeatherCommand(s, m, fields[1:])
	case "!poll":
		h.handlePollCommand(s, m, fields[1:])
	case "!vote":
		h.handleVoteCommand(s, m, fields[1:])
	case "!result":
		h.handlePollResults(s, m, fields[1])

	default:
		h.handleHelpCommand(s, m)
	}
}

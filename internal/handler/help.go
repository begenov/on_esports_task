package handler

import "github.com/bwmarrin/discordgo"

func (h *Handler) handleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "Available commands:\n" +
		"1. !poll <question> | <option1>, <option2>, ...\n" +
		"2. !vote <pollID> <optionID>\n" +
		"3. !result <pollID>\n" +
		"4. !weather <city>\n" +
		"5. !help\n"

	s.ChannelMessageSend(m.ChannelID, helpMessage)
}

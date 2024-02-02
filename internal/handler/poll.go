package handler

import (
	"fmt"
	"on_esports/internal/logger"
	"on_esports/internal/service"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) handleVoteCommand(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	if len(params) != 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: vote <pollID> <optionID>")
		return
	}

	pollID, err := strconv.Atoi(params[0])
	if err != nil {
		logger.Errorf("strconv.Atoi(): %v", err)
		s.ChannelMessageSend(m.ChannelID, "Invalid poll ID. Please provide a valid integer")
		return
	}

	optionID, err := strconv.Atoi(params[1])
	if err != nil {
		logger.Errorf("strconv.Atoi(): %v", err)
		s.ChannelMessageSend(m.ChannelID, "Invalid option ID. Please provide a valid integer")
		return
	}

	err = h.PollService.Vote(pollID, optionID)
	if err != nil {
		logger.Errorf("h.PollService.Vote(): %v", err)
		s.ChannelMessageSend(m.ChannelID, "Internal server.")
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Vote registered!")
}

func (h *Handler) handlePollCommand(s *discordgo.Session, m *discordgo.MessageCreate, params []string) {
	if len(params) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: poll <question> | <option1>, <option2> ...")
		return
	}

	index := -1
	for i := range params {
		if params[i] == "|" {
			index = i
			break
		}
	}

	if index == -1 {
		s.ChannelMessageSend(m.ChannelID, "Usage: !poll <question> | <option1>, <option2>, <option3>...")
		return
	}

	question := strings.Join(params[:index], " ")
	optionsString := strings.Join(params[index+1:], " ")

	options := strings.Split(optionsString, " ")
	for i := range options {
		options[i] = strings.TrimSpace(options[i])
	}

	if len(options) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You must provide at least two options")
		return
	}

	creatorID := m.Author.ID
	poll, err := h.PollService.CreatePoll(question, options, creatorID)
	if err != nil {
		logger.Errorf("h.PollService.CreatePoll(): %v", err)
		s.ChannelMessageSend(m.ChannelID, "Internal server")
		return
	}

	response := fmt.Sprintf("Poll created! ID: %v\n%s", poll.ID, h.getPollMessage(poll))
	s.ChannelMessageSend(m.ChannelID, response)
}

func (h *Handler) getPollMessage(poll *service.Poll) string {
	var pollMessage string

	pollMessage += fmt.Sprintf("**%s**\n\n", poll.Question)

	for _, option := range poll.Options {
		pollMessage += fmt.Sprintf("%d. %s\n", option.ID, option.Text)
	}

	return pollMessage
}

func (h *Handler) handlePollResults(s *discordgo.Session, m *discordgo.MessageCreate, pollIDStr string) {

	pollID, err := strconv.Atoi(pollIDStr)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Invalid pollID. Please provide a valid integer")
		return
	}

	poll, err := h.PollService.GetPollResults(pollID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error getting results for pollID %d: %v", pollID, err))
		return
	}

	response := fmt.Sprintf("Results for Poll ID %d:\n%s", poll.ID, h.getPollResultsMessage(poll))
	s.ChannelMessageSend(m.ChannelID, response)
}

func (h *Handler) getPollResultsMessage(poll *service.Poll) string {
	var resultsMessage string

	resultsMessage += fmt.Sprintf("**%s**\n\n", poll.Question)

	for _, option := range poll.Options {
		resultsMessage += fmt.Sprintf("%s: %d votes\n", option.Text, option.Votes)
	}

	return resultsMessage
}

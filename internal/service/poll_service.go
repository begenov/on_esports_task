package service

import (
	"errors"
	"sync"
)

type Poll struct {
	ID        int
	Question  string
	Options   []Option
	CreatorID string
	mu        sync.Mutex
}

type Option struct {
	ID    int
	Text  string
	Votes int
}

type PollService interface {
	CreatePoll(question string, options []string, creatorID string) (*Poll, error)
	Vote(pollID, optionID int) error
	GetPollResults(pollID int) (*Poll, error)
}

type InMemoryPollService struct {
	polls    map[int]*Poll
	nextPoll int
}

func NewInMemoryPollService() *InMemoryPollService {
	return &InMemoryPollService{
		polls:    make(map[int]*Poll),
		nextPoll: 1,
	}
}

func (s *InMemoryPollService) CreatePoll(question string, options []string, creatorID string) (*Poll, error) {
	s.nextPoll++
	poll := &Poll{ID: s.nextPoll, Question: question, CreatorID: creatorID}
	for i, text := range options {
		option := Option{ID: i + 1, Text: text}
		poll.Options = append(poll.Options, option)
	}
	s.polls[poll.ID] = poll
	return poll, nil
}

func (s *InMemoryPollService) Vote(pollID, optionID int) error {
	poll, exists := s.polls[pollID]
	if !exists {
		return errors.New("poll not found")
	}

	poll.mu.Lock()
	defer poll.mu.Unlock()

	if optionID < 1 || optionID > len(poll.Options) {
		return errors.New("invalid option ID")
	}

	poll.Options[optionID-1].Votes++
	return nil
}

func (s *InMemoryPollService) GetPollResults(pollID int) (*Poll, error) {
	poll, exists := s.polls[pollID]
	if !exists {
		return nil, errors.New("poll not found")
	}
	return poll, nil
}

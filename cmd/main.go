package main

import (
	"on_esports/internal/config"
	"on_esports/internal/handler"
	"on_esports/internal/logger"
	"on_esports/internal/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	if err := config.New(".env"); err != nil {
		logger.Fatalf("config.LoadFromFile(): %v", err)
	}

	weatherService := service.NewWeatherService(config.Config.OPEN_WEATHER_MAP_API_KEY)
	pollService := service.NewInMemoryPollService()

	discord, err := discordgo.New("Bot " + config.Config.DISCORD_BOT_TOKEN)
	if err != nil {
		logger.Fatalf("discordgo.New(): %v", err)
		return
	}

	botHandler := handler.New(discord, pollService, weatherService)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		logger.Info("Starting the program...")
		if err := botHandler.Start(); err != nil {
			logger.Fatalf("botHandler.Start(): %v", err)
			return
		}
	}()

	<-stopChan

	botHandler.Stop()

	logger.Info("Program terminated.")
}

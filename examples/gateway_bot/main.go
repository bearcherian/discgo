package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bearcherian/discgo"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	EnvDiscordApiBaseURL    = "DISCORD_API_BASE_URL"
	EnvDiscordApplicationID = "DISCORD_APPLICATION_ID"
	EnvDiscordClientID      = "DISCORD_CLIENT_ID"
	EnvDiscordClientSecret  = "DISCORD_CLIENT_SECRET"
	EnvDiscordBotToken      = "DISCORD_BOT_TOKEN"
	EnvDiscordPublicKey     = "DISCORD_PUBLIC_KEY"
)

func main() {
	godotenv.Load()

	l, _ := zap.NewDevelopment()
	logger := l.Sugar()

	dc := discgo.NewApiClient(
		discgo.WithApiConfig(&discgo.ApiConfig{
			APIBaseURL: os.Getenv(EnvDiscordApiBaseURL),
			BotToken:   os.Getenv(EnvDiscordBotToken),
		}),
		discgo.WithClientLogger(logger),
	)

	gatewayResponse, _, err := dc.GetGatewayBot(context.Background())
	if err != nil {
		logger.Errorw("unable to get gateway info", "error", err)
	}

	gwClient := discgo.NewGateway(
		discgo.WithConfig(discgo.GatewayConfig{
			AuthToken:   "Bot " + os.Getenv(EnvDiscordBotToken),
			GatewayURL:  gatewayResponse.Url,
			ShardID:     0,
			TotalShards: 1,
		}),
		discgo.WithLogger(logger),
	)

	gwClient.AddEventHandler(discgo.EventMessageCreate, func(d interface{}) {
		var messageData discgo.MessageCreateEventData
		discgo.DataToStruct(d, &messageData)

		fmt.Println("message say ", messageData.Content)
	})

	gwClient.Start()

	defer func() {
		dc.Close()
		gwClient.Close()
	}()

	// wait before killing
	awaitSignal()
	logger.Info("shutting down")
}

func awaitSignal() os.Signal {
	// wait for kill
	sigs := make(chan os.Signal, 1)
	done := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		done <- sig
		fmt.Println("recieved ", sig)
	}()

	return <-done
}

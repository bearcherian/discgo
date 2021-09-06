package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bearcherian/discgo"

	"go.uber.org/zap"
)

func main() {

	// *zap.Logger fulfills discgo.Logger
	l, _ := zap.NewDevelopment()
	logger := l.Sugar()

	// create the API client to make HTTP calls
	// we need this to get the proper Gateway URL and configurations
	dc := discgo.NewApiClient(
		discgo.WithApiConfig(&discgo.ApiConfig{
			BotToken: "MY_BOT_TOKEN",
		}),
		discgo.WithClientLogger(logger),
	)

	// GET request to get gateway url and params
	gatewayResponse, _, err := dc.GetGatewayBot(context.Background())
	if err != nil {
		logger.Errorw("unable to get gateway info", "error", err)
	}

	// create a new Gateway websocket client
	gwClient := discgo.NewGateway(
		discgo.WithConfig(discgo.GatewayConfig{
			AuthToken:   "Bot " + "MY_BOT_TOKEN",
			GatewayURL:  gatewayResponse.Url,
			ShardID:     0,
			TotalShards: 1,
		}),
		discgo.WithLogger(logger),
	)

	// adding a new event handler for new Discord messages
	gwClient.AddEventHandler(discgo.EventMessageCreate, func(d interface{}) {
		var messageData discgo.MessageCreateEventData
		discgo.DataToStruct(d, &messageData)

		fmt.Println("message say ", messageData.Content)
	})

	// start the websocket connection to the Gateway.
	// Event handlers can be called before or after the connection is established
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

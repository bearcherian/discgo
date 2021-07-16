package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OpCode int

const (
	GetGatewayBotURL = "/gateway/bot"
)

type SessionStartLimit struct {
	// The total number of session starts the current user is allowed
	Total int `json:"total"`
	// The remaining number of session starts the current user is allowed
	Remaining int `json:"remaining"`
	// The number of milliseconds after which the limit resets
	ResetAfter int `json:"remain_after"`
	// The number of identify requests allowed per 5 seconds
	MaxConcurrency int `json:"max_concurrency"`
}

type GetGatewayBotResponse struct {
	// The WSS URL that can be used for connecting to the gateway
	Url string `json:"url"`
	// The recommended number of shards to use when connecting
	Shards int `json:"shards"`
	// Information on the current session start limit
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

func (d *DiscordClient) GetGatewayBot(ctx context.Context) (*GetGatewayBotResponse, error) {

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s?encoding=json", d.discordConfig.APIBaseURL, GetGatewayBotURL), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to build request: %w", err)
	}

	resp, err := d.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		errBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("reqeust failed: %d - %s: %s", resp.StatusCode, resp.Status, string(errBytes))
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %w", err)
	}

	var response GetGatewayBotResponse
	if err = json.Unmarshal(responseBytes, &response); err != nil {
		return nil, fmt.Errorf("unable to parse JSON: %w", err)
	}

	return &response, nil
}

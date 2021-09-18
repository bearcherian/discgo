package discgo

import (
	"context"
	"fmt"
	"net/http"
)

const (
	endpointGuilds        = "/guilds"
	endpointGuildByID     = "/guilds/%s"
	endpointGuildChannels = "/guilds/%s/channels"
	endpointGuildPreview  = "/guilds/%s/preview"
)

func (a *ApiClient) CreateGuild(ctx context.Context, guild Guild) (Guild, *http.Response, error) {

	var responseGuild Guild
	resp, err := a.Post(ctx, endpointGuilds, guild, &responseGuild)
	return responseGuild, resp, err
}

func (a *ApiClient) GetGuild(ctx context.Context, guildID string) (Guild, *http.Response, error) {
	var responseGuild Guild
	resp, err := a.Get(ctx, fmt.Sprintf(endpointGuildByID, guildID), &responseGuild)
	return responseGuild, resp, err
}

func (a *ApiClient) ModifyGuild(ctx context.Context, guild Guild) (Guild, *http.Response, error) {
	var responseGuild Guild
	if guild.Id == "" {
		return responseGuild, nil, fmt.Errorf("must have a valid Guild Id")
	}
	resp, err := a.Patch(ctx, fmt.Sprintf(endpointGuildByID, guild.Id), guild, &responseGuild)
	return responseGuild, resp, err
}

func (a *ApiClient) DeleteGuild(ctx context.Context, guildID string) (*http.Response, error) {

	resp, err := a.Delete(ctx, fmt.Sprintf("%s%s", endpointGuildByID, guildID), nil, nil)
	return resp, err
}

func (a *ApiClient) GetGuildChannels(ctx context.Context, guildID string) ([]Channel, *http.Response, error) {
	var channels []Channel
	resp, err := a.Get(ctx, fmt.Sprintf("%s%s", endpointGuildChannels, guildID), &channels)

	return channels, resp, err
}

func (a *ApiClient) CreateGuildChannel(ctx context.Context, guildID string, channel Channel) (Channel, *http.Response, error) {
	var respChannel Channel
	resp, err := a.Post(ctx, fmt.Sprintf("%s%s", endpointGuildChannels, guildID), channel, &respChannel)
	return respChannel, resp, err
}

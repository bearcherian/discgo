package discgo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	endpointGuilds             = "/guilds"
	endpointGuildByID          = "/guilds/%s"
	endpointGuildChannels      = "/guilds/%s/channels"
	endpointGuildPreview       = "/guilds/%s/preview"
	endpointGuildThreadsActive = "/guilds/%s/threads/active"
	endpointGuildMembers       = "/guilds/%s/members"
	endpointGuildMemberById    = "/guilds/%s/members/%s"
	endpointGuildMembersSearch = "/guilds/%s/members/search"
)

// CreateGuild Create a new guild. Returns a guild object on success. Fires a Guild Create Gateway event.
// NOTE: This endpoint can be used only by bots in less than 10 guilds.
func (a *ApiClient) CreateGuild(ctx context.Context, guild Guild) (Guild, *http.Response, error) {

	var responseGuild Guild
	resp, err := a.Post(ctx, endpointGuilds, guild, &responseGuild)
	if err != nil {
		return responseGuild, resp, err
	}
	return responseGuild, resp, errorFromResponse(resp)
}

// GetGuild Returns the guild object for the given id. If with_counts is set to true, this endpoint will also
// return approximate_member_count and approximate_presence_count for the guild.
func (a *ApiClient) GetGuild(ctx context.Context, guildID string) (Guild, *http.Response, error) {
	var responseGuild Guild
	resp, err := a.Get(ctx, fmt.Sprintf(endpointGuildByID, guildID), &responseGuild)
	if err != nil {
		return responseGuild, resp, err
	}
	return responseGuild, resp, errorFromResponse(resp)
}

// GetGuildPreview Returns the guild preview object for the given id. If the user is not in the guild, then the guild
// must be lurkable (it must be Discoverable or have a live public stage).
func (a *ApiClient) GetGuildPreview(ctx context.Context, guildID string) (Guild, *http.Response, error) {
	var responseGuild Guild
	resp, err := a.Get(ctx, fmt.Sprintf(endpointGuildPreview, guildID), &responseGuild)
	if err != nil {
		return responseGuild, resp, err
	}
	return responseGuild, resp, errorFromResponse(resp)
}

// ModifyGuild Modify a guild's settings. Requires the MANAGE_GUILD permission. Returns the updated guild object
// on success. Fires a Guild Update Gateway event. All parameters to this endpoint are optional
func (a *ApiClient) ModifyGuild(ctx context.Context, guild Guild) (Guild, *http.Response, error) {
	var responseGuild Guild
	if guild.Id == "" {
		return responseGuild, nil, fmt.Errorf("must have a valid Guild Id")
	}
	resp, err := a.Patch(ctx, fmt.Sprintf(endpointGuildByID, guild.Id), guild, &responseGuild)
	if err != nil {
		return responseGuild, resp, err
	}
	return responseGuild, resp, errorFromResponse(resp)
}

// DeleteGuild Delete a guild permanently. User must be owner. Returns 204 No Content on success.
// Fires a Guild Delete Gateway event.
func (a *ApiClient) DeleteGuild(ctx context.Context, guildID string) (*http.Response, error) {

	resp, err := a.Delete(ctx, fmt.Sprintf(endpointGuildByID, guildID), nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, errorFromResponse(resp)
}

// GetGuildChannels Returns a list of guild channel objects. Does not include threads.
func (a *ApiClient) GetGuildChannels(ctx context.Context, guildID string) ([]Channel, *http.Response, error) {
	var channels []Channel
	resp, err := a.Get(ctx, fmt.Sprintf(endpointGuildChannels, guildID), &channels)
	if err != nil {
		return channels, resp, err
	}
	return channels, resp, errorFromResponse(resp)
}

// CreateGuildChannel Create a new channel object for the guild. Requires the MANAGE_CHANNELS permission. If setting permission overwrites,
// only permissions your bot has in the guild can be allowed/denied. Setting MANAGE_ROLES permission in channels is only
// possible for guild administrators. Returns the new channel object on success. Fires a Channel Create Gateway event.
// All parameters to this endpoint are optional excluding 'name'
func (a *ApiClient) CreateGuildChannel(ctx context.Context, guildID string, channel Channel) (Channel, *http.Response, error) {
	// TODO Reason header
	var respChannel Channel
	resp, err := a.Post(ctx, fmt.Sprintf(endpointGuildChannels, guildID), channel, &respChannel)
	if err != nil {
		return respChannel, resp, err
	}
	return respChannel, resp, errorFromResponse(resp)
}

// ModifyGuildChannelPositions Modify the positions of a set of channel objects for the guild. Requires MANAGE_CHANNELS permission.
// Returns a 204 empty response on success. Fires multiple Channel Update Gateway events.
func (a *ApiClient) ModifyGuildChannelPositions(ctx context.Context, guildID string, channels []Channel) (*http.Response, error) {
	resp, err := a.Patch(ctx, fmt.Sprintf(endpointGuildChannels, guildID), channels, nil)
	if err != nil {
		return resp, err
	}
	return resp, errorFromResponse(resp)
}

// ListActiveThreads Returns all active threads in the guild, including public and private threads.
// Threads are ordered by their id, in descending order.
func (a *ApiClient) ListActiveThreads(ctx context.Context, guildID string) (ActiveThreadsResponse, *http.Response, error) {
	var activeThreads ActiveThreadsResponse
	resp, err := a.Get(ctx, fmt.Sprintf(endpointGuildThreadsActive, guildID), &activeThreads)
	if err != nil {
		return activeThreads, resp, err
	}
	return activeThreads, resp, errorFromResponse(resp)
}

// GetGuildMember Returns a guild member object for the specified user.
func (a *ApiClient) GetGuildMember(ctx context.Context, guildID, userID string) (GuildMember, *http.Response, error) {
	var responseMember GuildMember
	resp, err := a.Get(ctx, fmt.Sprintf(endpointGuildMemberById, guildID, userID), &responseMember)
	if err != nil {
		return responseMember, resp, err
	}
	return responseMember, resp, errorFromResponse(resp)
}

// ListGuildMembers Returns a list of guild member objects that are members of the guild.
// This endpoint is restricted according to whether the GUILD_MEMBERS Privileged Intent is enabled for your application.
// All parameters to this endpoint are optional
func (a *ApiClient) ListGuildMembers(ctx context.Context, guildID string, limit int, after string) ([]GuildMember, *http.Response, error) {
	var members []GuildMember

	url, err := url.Parse(fmt.Sprintf(endpointGuildMembers, guildID))
	if err != nil {
		return members, nil, fmt.Errorf("unable to parse URL: %w", err)
	}

	q := url.Query()
	if limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", limit))
	}
	if after != "" {
		q.Set("after", after)
	}

	url.RawQuery = q.Encode()

	resp, err := a.Get(ctx, url.String(), &members)
	if err != nil {
		return members, resp, err
	}
	return members, resp, errorFromResponse(resp)
}

// SearchGuildMembers Returns a list of guild member objects whose username or nickname starts with a provided string.
// All parameters to this endpoint except for query are optional
func (a *ApiClient) SearchGuildMembers(ctx context.Context, guildID string, query string, limit int) ([]GuildMember, *http.Response, error) {
	var responseMembers []GuildMember

	url, err := url.Parse(fmt.Sprintf(endpointGuildMembersSearch, guildID))
	if err != nil {
		return responseMembers, nil, fmt.Errorf("unable to parse URL: %w", err)
	}

	q := url.Query()
	q.Set("query", query)
	if limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", limit))
	}
	url.RawQuery = q.Encode()

	resp, err := a.Get(ctx, url.String(), &responseMembers)
	if err != nil {
		return responseMembers, resp, err
	}
	return responseMembers, resp, errorFromResponse(resp)
}

func (a *ApiClient) AddGuildMember(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) ModifyGuildMember(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) ModifyCurrentUserNick(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) AddGuildMemberRole(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) RemoveGuildMemberRole(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) RemoveGuildMember(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildBans(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildBan(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) CreateGuildBan(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) RemoveGuildBan(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildRoles(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) CreateGuildRole(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) ModifyGuildRolePositions(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) ModifyGuildRole(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) DeleteGuildRole(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildPruneCount(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) BeginGuildPrune(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildVoiceRegions(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildInvites(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildIntegrations(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) DeleteGuildIntegration(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildWidgetSettings(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) ModifyGuildWidget(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildWidget(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildVanityURL(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildWidgetImage(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) GetGuildWelcomeScreen(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) ModifyGuildWelcomeScreen(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) ModifyCurrentUserVoiceState(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *ApiClient) ModifyUserVoiceState(ctx context.Context, guildID string) (*http.Response, error) {
	return nil, fmt.Errorf("not implemented")
}

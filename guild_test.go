package discgo

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

const (
	fakeServerURL = "http://fakeserver.com"
)

func TestApiClient_CreateGuild(t *testing.T) {

	type args struct {
		ctx   context.Context
		guild Guild
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		want    Guild
		wantErr bool
	}{
		{
			name: "CreateGuild succeeds",
			args: args{
				ctx:   context.Background(),
				guild: Guild{},
			},
			setup: func() {
				gock.New(fakeServerURL).
					Post(endpointGuilds).
					Reply(204).
					JSON(Guild{})
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			got, _, err := a.CreateGuild(tt.args.ctx, tt.args.guild)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.CreateGuild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiClient.CreateGuild() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiClient_GetGuild(t *testing.T) {

	type args struct {
		ctx     context.Context
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		want    Guild
		wantErr bool
	}{
		{
			name: "GetGuild succeeds",
			args: args{
				ctx:     context.Background(),
				guildID: "snowflake",
			},
			setup: func() {
				gock.New(fakeServerURL).
					Get(fmt.Sprintf(endpointGuildByID, "snowflake")).
					Reply(200).
					JSON(map[string]interface{}{"id": "snowflake"})
			},
			want: Guild{Id: "snowflake"},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			got, _, err := a.GetGuild(tt.args.ctx, tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.GetGuild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiClient.GetGuild() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiClient_GetGuildPreview(t *testing.T) {

	type args struct {
		ctx     context.Context
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		want    Guild
		wantErr bool
	}{
		{
			name: "GetGuildPreview succeeds",
			args: args{
				ctx:     context.Background(),
				guildID: "snowflake",
			},
			setup: func() {
				gock.New(fakeServerURL).
					Get(fmt.Sprintf(endpointGuildPreview, "snowflake")).
					Reply(200).
					JSON(map[string]interface{}{"id": "snowflake"})
			},
			want: Guild{Id: "snowflake"},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			got, _, err := a.GetGuildPreview(tt.args.ctx, tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.GetGuildPreview() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiClient.GetGuildPreview() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiClient_ModifyGuild(t *testing.T) {

	type args struct {
		ctx   context.Context
		guild Guild
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		want    Guild
		wantErr bool
	}{
		{
			name: "ModifyGuild succeeds",
			args: args{
				ctx: context.Background(),
				guild: Guild{
					Id:   "snowflake",
					Name: "Discgo Fake Server",
				},
			},
			setup: func() {
				gock.New(fakeServerURL).
					Patch(fmt.Sprintf(endpointGuildByID, "snowflake")).
					Reply(200).
					JSON(map[string]interface{}{"id": "snowflake", "name": "Discgo Fake Server"})
			},
			want: Guild{
				Id:   "snowflake",
				Name: "Discgo Fake Server",
			},
		},
		{
			name: "ModifyGuild fails on missing guild ID",
			args: args{
				ctx: context.Background(),
				guild: Guild{
					Name: "Discgo Fake Server",
				},
			},
			setup: func() {

			},
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			got, _, err := a.ModifyGuild(tt.args.ctx, tt.args.guild)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.ModifyGuild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiClient.ModifyGuild() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiClient_DeleteGuild(t *testing.T) {

	type args struct {
		ctx     context.Context
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		wantErr bool
	}{
		{
			name: "DeleteGuild succeeds",
			args: args{
				ctx:     context.Background(),
				guildID: "snowflake",
			},
			setup: func() {
				gock.New(fakeServerURL).
					Delete(fmt.Sprintf(endpointGuildByID, "snowflake")).
					Reply(204)
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			_, err := a.DeleteGuild(tt.args.ctx, tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.DeleteGuild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestApiClient_GetGuildChannels(t *testing.T) {

	type args struct {
		ctx     context.Context
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		want    []Channel
		wantErr bool
	}{
		{
			name: "GetGuildChannels succeeds",
			args: args{
				ctx:     context.Background(),
				guildID: "snowflake",
			},
			setup: func() {
				gock.New(fakeServerURL).
					Get(fmt.Sprintf(endpointGuildChannels, "snowflake")).
					Reply(200).
					JSON([]map[string]interface{}{
						{
							"id":       "channel_snowflake1",
							"type":     1,
							"guild_id": "snowflake",
							"name":     "general",
						},
						{
							"id":       "channel_snowflake2",
							"type":     1,
							"guild_id": "snowflake",
							"name":     "random",
						},
						{
							"id":       "channel_snowflake3",
							"type":     1,
							"guild_id": "snowflake",
							"name":     "rules",
						},
					})
			},
			want: []Channel{
				{
					Id:      "channel_snowflake1",
					Type:    1,
					GuildId: "snowflake",
					Name:    "general",
				},
				{
					Id:      "channel_snowflake2",
					Type:    1,
					GuildId: "snowflake",
					Name:    "random",
				},
				{
					Id:      "channel_snowflake3",
					Type:    1,
					GuildId: "snowflake",
					Name:    "rules",
				},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			got, _, err := a.GetGuildChannels(tt.args.ctx, tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.GetGuildChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ApiClient.GetGuildChannels() incorrect number of channels, got = %v, want %v", len(got), len(tt.want))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiClient.GetGuildChannels() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiClient_CreateGuildChannel(t *testing.T) {

	type args struct {
		ctx     context.Context
		guildID string
		channel Channel
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		wantErr bool
	}{
		{
			name: "CreateGuildChannel succeeds",
			args: args{
				ctx:     context.Background(),
				guildID: "snowflake",
				channel: Channel{
					Type: 1,
					Name: "some guild channel",
				},
			},
			setup: func() {
				gock.New(fakeServerURL).
					Post(fmt.Sprintf(endpointGuildChannels, "snowflake")).
					Reply(200).
					JSON(map[string]interface{}{
						"id":      "channel_snowflake",
						"guildID": "snowflake",
						"name":    "some guild channel",
						"type":    1,
					})
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			_, _, err := a.CreateGuildChannel(tt.args.ctx, tt.args.guildID, tt.args.channel)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.CreateGuildChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestApiClient_ModifyGuildChannelPositions(t *testing.T) {

	type args struct {
		ctx      context.Context
		guildID  string
		channels []Channel
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		wantErr bool
	}{
		{
			name: "ModifyGuildChannelPositions succeeds",
			args: args{
				ctx:     context.Background(),
				guildID: "snowflake",
				channels: []Channel{
					{
						Type: 1,
						Name: "some guild channel",
						Id:   "some_channel_flake",
					},
					{
						Type: 1,
						Name: "some guild channel2",
						Id:   "some_channel_flake2",
					},
				},
			},
			setup: func() {
				gock.New(fakeServerURL).
					Patch(fmt.Sprintf(endpointGuildChannels, "snowflake")).
					Reply(204)
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			_, err := a.ModifyGuildChannelPositions(tt.args.ctx, tt.args.guildID, tt.args.channels)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.ModifyGuildChannelPositions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestApiClient_ListGuildMembers(t *testing.T) {

	type args struct {
		ctx     context.Context
		guildID string
		limit   int
		after   string
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		wantErr bool
		want    []GuildMember
	}{
		{
			name: "ListGuildMembers succeeds",
			args: args{
				ctx:     context.Background(),
				guildID: "snowflake",
				limit:   50,
				after:   "some_channel_flake",
			},
			want: []GuildMember{
				{
					User: User{
						Id:            "userID",
						Username:      "some_username",
						Discriminator: "4325",
					},
					Nick: "Jimmy",
				},
				{
					User: User{
						Id:            "userID2",
						Username:      "some_other_username",
						Discriminator: "43225",
					},
					Nick: "Johnny",
				},
			},
			setup: func() {
				gock.New(fakeServerURL).
					Get(fmt.Sprintf(endpointGuildMembers, "snowflake")).
					MatchParams(map[string]string{
						"after": "some_channel_flake",
						"limit": "50",
					}).
					Reply(200).
					JSON([]map[string]interface{}{
						{
							"user": map[string]interface{}{
								"id":            "userID",
								"username":      "some_username",
								"discriminator": "4325",
							},
							"nick": "Jimmy",
						},
						{
							"user": map[string]interface{}{
								"id":            "userID2",
								"username":      "some_other_username",
								"discriminator": "43225",
							},
							"nick": "Johnny",
						},
					})
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			tt.setup()

			a := NewApiClient(
				WithApiConfig(&ApiConfig{
					APIBaseURL: fakeServerURL,
				}),
			)

			members, _, err := a.ListGuildMembers(tt.args.ctx, tt.args.guildID, tt.args.limit, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiClient.ListGuildMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(members, tt.want) {
				t.Errorf("ApiClient.ListGuildMembers()\n\twant:\n\t\t%v\n\tgot:\n\t\t%v\n", tt.want, members)
			}

		})
	}
}

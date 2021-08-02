package main

import (
	"strings"
	"testing"
)

func Test_parseObjectFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "multi word type array is array with Title case",
			args: args{s: "array of [guild member](#DOCS_RESOURCES_GUILD/guild-member-object) objects"},
			want: "GuildMember",
		},
		{
			name: "an array of user is []User",
			args: args{s: "array of [user](#DOC_LINK/user-obect) object"},
			want: "User",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseObjectFromString(tt.args.s); got != tt.want {
				t.Errorf("parseObjectFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeFromString(t *testing.T) {
	type args struct {
		s string
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "parses standard string",
			args: args{s: "string"},
			want: "string",
		},
		{
			name: "parses string with markers",
			args: args{s: "?integer**"},
			want: "int",
		},
		{
			name: "parses array type",
			args: args{s: "array of [guild feature](#DOCS_RESOURCES_GUILD/guild-object-guild-features) strings"},
			want: "[]string",
		},
		{
			name: "parses object type string",
			args: args{s: "array of [guild member](#DOCS_RESOURCES_GUILD/guild-member-object) objects"},
			want: "[]GuildMember",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typeFromString(tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("typeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

var roleContents = "### Role Object\n" +
	"\n" +
	"Roles represent a set of permissions attached to a group of users. Roles have unique names, colors, and can be \"pinned\" to the side bar, causing their members to be listed separately. Roles are unique per guild, and can have separate permission profiles for the global context (guild) and channel context. The `@everyone` role has the same ID as the guild it belongs to.\n" +
	"\n" +
	"###### Role Structure\n" +
	"\n" +
	"| Field       | Type                                                                         | Description                                      |\n" +
	"| ----------- | ---------------------------------------------------------------------------- | ------------------------------------------------ |\n" +
	"| id          | snowflake                                                                    | role id                                          |\n" +
	"| name        | string                                                                       | role name                                        |\n" +
	"| color       | integer                                                                      | integer representation of hexadecimal color code |\n" +
	"| hoist       | boolean                                                                      | if this role is pinned in the user listing       |\n" +
	"| guilds      | array of snowflakes                                                          | made up field for testing []string types         |\n" +
	"| role_users  | array of [role user](/#DOC_RESOURCE/role-user-object) objects                | made up field for testing []object types         |\n" +
	"\n"

var roleStructure = "type Role struct {\n" +
	"	// Id role id\n" +
	"	Id string `json:\"id\"`\n" +
	"	// Name role name\n" +
	"	Name string `json:\"name\"`\n" +
	"	// Color integer representation of hexadecimal color code\n" +
	"	Color int `json:\"color\"`\n" +
	"	// Hoist if this role is pinned in the user listing\n" +
	"	Hoist bool `json:\"hoist\"`\n" +
	"	// Guilds made up field for testing []string types\n" +
	"	Guilds []string `json:\"guilds\"`\n" +
	"	// RoleUsers made up field for testing []object types\n" +
	"	RoleUsers []RoleUser `json:\"role_users\"`\n" +
	"}"

func Test_parseFileContents(t *testing.T) {
	type args struct {
		fileContents string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "parse simple struct",
			args: args{
				fileContents: roleContents,
			},
			want: []string{roleStructure},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := parseFileContents(tt.args.fileContents)
			if len(tt.want) != len(got) {
				t.Errorf("parseFileContents(): got %d structs, want %d structs", len(got), len(tt.want))
			}

			for i, gotStruct := range got {
				gotLines := strings.Split(gotStruct, "\n")
				wantLines := strings.Split(tt.want[i], "\n")

				for j, gotLine := range gotLines {
					if wantLines[j] != gotLine {
						t.Errorf("parseFileContents() - struct %d, line %d\n\tGot: '%v'\n\tWant: '%v'", i, j, gotLine, wantLines[j])
					}
				}

			}
		})
	}
}

package main

import (
	"github.com/koudaiii/qucli/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"add-team": func() (cli.Command, error) {
			return &command.AddTeamCommand{
				Meta: *meta,
			}, nil
		},
		"delete-team": func() (cli.Command, error) {
			return &command.DeleteTeamCommand{
				Meta: *meta,
			}, nil
		},
		"add-user": func() (cli.Command, error) {
			return &command.AddUserCommand{
				Meta: *meta,
			}, nil
		},
		"delete-user": func() (cli.Command, error) {
			return &command.DeleteUserCommand{
				Meta: *meta,
			}, nil
		},
		"add-notification": func() (cli.Command, error) {
			return &command.AddNotificationCommand{
				Meta: *meta,
			}, nil
		},
		"delete-notification": func() (cli.Command, error) {
			return &command.DeleteNotificationCommand{
				Meta: *meta,
			}, nil
		},
		"test-notification": func() (cli.Command, error) {
			return &command.TestNotificationCommand{
				Meta: *meta,
			}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},
		"get": func() (cli.Command, error) {
			return &command.GetCommand{
				Meta: *meta,
			}, nil
		},
		"create": func() (cli.Command, error) {
			return &command.CreateCommand{
				Meta: *meta,
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.DeleteCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}

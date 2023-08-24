package commands

import (
	"errors"
	go_groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/groshi-project/grosh/internal/prompts"
	"github.com/urfave/cli/v2"
)

func RegisterCommand(ctx *cli.Context) error {
	args := ctx.Args()
	argsCount := args.Len()
	if argsCount < 1 || argsCount > 2 {
		return ctx.Command.OnUsageError(ctx, errors.New("invalid number of args"), true)
	}

	url := args.Get(0) // required argument url

	username := args.Get(1) // optional argument username

	if username == "" { // if username was not provided as argument
		var err error
		username, err = prompts.ReadString("Username: ")
		if err != nil {
			return err
		}
	}

	password1, err := prompts.ReadPassword("Password: ")
	if err != nil {
		return err
	}

	password2, err := prompts.ReadPassword("Repeat password: ")
	if err != nil {
		return err
	}

	if password1 != password2 {
		return errors.New("passwords does not match")
	}

	groshiClient := go_groshi.NewGroshiAPIClient(url, "")
	if _, err := groshiClient.UserCreate(username, password1); err != nil {
		return err
	}

	output.PlusLogger.Printf("Created new groshi user @%v.", username)
	output.TipLogger.Printf("Now you can authorize using `groshi login %v %v` command.", url, username)

	return nil
}

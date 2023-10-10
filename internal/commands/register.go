package commands

import (
	"errors"
	groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/input"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
)

// RegisterCommand ...
// grosh register <URL> [username]
func RegisterCommand(ctx *cli.Context) error {
	args := ctx.Args()
	argsCount := args.Len()
	if argsCount < 1 || argsCount > 2 {
		return ctx.Command.OnUsageError(ctx, errors.New("invalid number of args"), true)
	}

	// required argument URL:
	url := args.Get(0)

	// optional argument USERNAME:
	username := args.Get(1)

	if username == "" { // if username was not provided
		var err error
		if username, err = input.ReadString("Username: "); err != nil {
			return err
		}
	}

	password1, err := input.ReadPassword("Password: ")
	if err != nil {
		return err
	}

	password2, err := input.ReadPassword("Repeat password: ")
	if err != nil {
		return err
	}

	if password1 != password2 {
		return errors.New("passwords do not match")
	}

	groshiClient := groshi.NewAPIClient(url, "")
	if _, err := groshiClient.UserCreate(username, password1); err != nil {
		return err
	}

	output.Plus.Printf("Created new groshi user @%v.", username)
	output.Tip.Printf("Now you can authorize using `grosh login %v %v` command.", url, username)

	return nil
}

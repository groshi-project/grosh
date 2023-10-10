package commands

import (
	"errors"
	groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/input"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
)

func AuthCommand(ctx *cli.Context) error {
	args := ctx.Args()
	argsCount := args.Len()
	if argsCount < 1 || argsCount > 3 {
		return ctx.Command.OnUsageError(ctx, errors.New("invalid number of arguments"), true)
	}

	// url is a required argument
	url := ctx.Args().Get(0)

	// username and password are optional arguments
	username := ctx.Args().Get(1)
	password := ctx.Args().Get(2)

	// read username from stdout if it was not provided
	if username == "" {
		var err error
		username, err = prompts.ReadString("Username: ")
		if err != nil {
			return err
		}
	}

	// read password from stdout if it was not provided
	if password == "" {
		var err error
		password, err = prompts.ReadPassword("Password: ")
		if err != nil {
			return err
		}
	}

	client := groshi.NewAPIClient(url, "")
	authData, err := client.AuthLogin(username, password)
	if err != nil {
		return err
	}

	output.PlusLogger.Printf("Authorized at groshi server at %v as @%v.", url, username)

	authCredentials := credentials.NewCredentials(url, authData.Token)
	storageFilePath := credentials.GetCredentialsStorageFilePath()
	if err := authCredentials.WriteToFile(storageFilePath); err != nil {
		return err
	}

	output.PlusLogger.Printf("Stored credentials locally at %v.", storageFilePath)

	return nil
}

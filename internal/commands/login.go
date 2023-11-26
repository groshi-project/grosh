package commands

import (
	groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/input"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
)

// LoginCommand logs in user with provided credentials.
// grosh login <URL> [USERNAME]
func LoginCommand(ctx *cli.Context) error {
	args := ctx.Args()

	// required argument URL:
	url := args.Get(0)

	// optional arguments USERNAME and PASSWORD:
	username := args.Get(1)
	password := args.Get(2)

	// read USERNAME from stdout if it was not provided:
	if username == "" {
		var err error
		if username, err = input.ReadString("Username: "); err != nil {
			return err
		}
	}

	// read PASSWORD from stdin if it was not provided:
	if password == "" {
		var err error
		if password, err = input.ReadPassword("Password: "); err != nil {
			return err
		}
	}

	client := groshi.NewAPIClient(url, "")
	authData, err := client.AuthLogin(username, password)
	if err != nil {
		return err
	}

	output.Plus.Printf("Authorized at groshi server at %v as %v.", url, username)

	authCredentials := credentials.New(url, authData.Token)
	storageFilePath := credentials.GetCredentialsStorageFilePath()
	if err := authCredentials.WriteToFile(storageFilePath); err != nil {
		return err
	}

	output.Plus.Printf("Stored credentials locally at %v.", storageFilePath)

	return nil
}

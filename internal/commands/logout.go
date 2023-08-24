package commands

import (
	"errors"
	"github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"os"
)

func CommandLogout(ctx *cli.Context) error {
	credentialsFilePath := credentials.GetCredentialsStorageFilePath()

	authData, err := credentials.NewCredentialsFromFile(credentialsFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("not authorized")
		}
		return err
	}

	groshiClient := go_groshi.NewGroshiAPIClient(authData.URL, authData.JWT)
	if _, err := groshiClient.AuthLogout(); err != nil {
		return err
	}

	output.MinusLogger.Printf("Successfully logged out from groshi at %v", authData.URL)

	if err := os.Remove(credentialsFilePath); err != nil {
		return err
	}

	output.MinusLogger.Printf("Successfully removed credentials from %v", credentialsFilePath)

	return nil
}

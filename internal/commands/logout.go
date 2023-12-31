package commands

import (
	"errors"
	"fmt"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"os"
)

// CommandLogout remove locally stored credentials.
// Usage: groshi logout.
func CommandLogout(_ *cli.Context) error {
	credentialsStorageFilePath := credentials.GetCredentialsStorageFilePath()

	_, err := os.Stat(credentialsStorageFilePath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("credentials storage file %v does not exist", credentialsStorageFilePath)
	}

	if err := os.Remove(credentialsStorageFilePath); err != nil {
		return err
	}
	output.Minus.Printf("Successfully removed credentials storage file %v", credentialsStorageFilePath)

	return nil
}

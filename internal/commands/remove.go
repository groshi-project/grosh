package commands

import (
	groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
)

// RemoveCommand removes transaction.
// Usage: grosh remove <UUID>.
func RemoveCommand(ctx *cli.Context) error {
	// required argument UUID:
	uuid := ctx.Args().Get(0)

	authData, err := credentials.ReadFromDefaultPath()
	if err != nil {
		return err
	}

	groshiClient := groshi.NewAPIClient(authData.URL, authData.Token)
	transaction, err := groshiClient.TransactionsDelete(uuid)
	if err != nil {
		return err
	}
	output.Minus.Printf("Successfully removed transaction %v", transaction.UUID)

	return nil
}

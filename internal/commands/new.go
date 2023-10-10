package commands

import (
	"errors"
	groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/input"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"strings"
)

// NewCommand is
// grosh new [--description=<TEXT>] [--timestamp=<TIME>] <AMOUNT> <CURRENCY>
func NewCommand(ctx *cli.Context) error {
	args := ctx.Args()

	// required argument AMOUNT:
	amountString := args.Get(0)
	amount, err := input.ParseAmount(amountString)
	if err != nil {
		return err
	}

	// required argument CURRENCY:
	currency := args.Get(1)
	currency = strings.ToUpper(currency)
	if len(currency) != 3 {
		return errors.New("invalid currency format")
	}

	// option --description:
	description := ctx.String("description")

	// option --timestamp:
	timestampString := ctx.String("timestamp")
	timestamp, err := input.ParseOptionalTime(timestampString)
	if err != nil {
		return err
	}

	authData, err := credentials.ReadFromDefaultPath()
	if err != nil {
		return err
	}

	groshiClient := groshi.NewAPIClient(authData.URL, authData.Token)
	transaction, err := groshiClient.TransactionsCreate(
		int(amount*100), currency, &description, timestamp,
	)
	if err != nil {
		return err
	}

	output.Plus.Printf(
		"Successfully created a new transaction (uuid: %v).", transaction.UUID,
	)

	return nil
}

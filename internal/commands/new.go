package commands

import (
	go_groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/groshi-project/grosh/internal/timeutil"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
	"time"
)

// NewCommand is
// grosh new <AMOUNT> <CURRENCY> --description=<TEXT> --timestamp=<TIME>
func NewCommand(ctx *cli.Context) error {
	// required arguments
	args := ctx.Args()
	amountRaw := args.Get(0)
	currency := args.Get(1)

	// options
	description := ctx.String("description")
	timestampRaw := ctx.String("timestamp")

	// parse `amountString` as `amount`
	amount, err := strconv.ParseFloat(amountRaw, 32)
	if err != nil {
		return err
	}

	// currency to upper case
	currency = strings.ToUpper(currency) // todo: also validate currency code

	var timestamp *time.Time
	if timestampRaw != "" {
		timestampValue, err := timeutil.ParseDate(timestampRaw)
		if err != nil {
			return err
		}
		timestamp = &timestampValue
	} else {
		timestamp = nil
	}

	authData, err := credentials.NewCredentialsFromCredentialsFile()
	if err != nil {
		return err
	}

	groshiClient := go_groshi.NewGroshiAPIClient(authData.URL, authData.JWT)
	transaction, err := groshiClient.TransactionsCreate(
		int(amount*100), currency, &description, timestamp,
	)
	if err != nil {
		return err
	}

	output.PlusLogger.Printf(
		"Successfully created a new transaction (uuid: %v).", transaction.UUID,
	)

	return nil
}

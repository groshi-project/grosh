package commands

import (
	groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/input"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
)

// ListCommand lists transactions in given period and optionally in given currency.
// Usage: grosh list [--uuid] [--currency=<CURRENCY>] [--end-time=<TIME>] <START-TIME>.
func ListCommand(ctx *cli.Context) error {
	args := ctx.Args()

	// required argument START-TIME:
	startTimeString := args.Get(0)
	startTime, err := input.ParseTime(startTimeString)
	if err != nil {
		return err
	}

	// option --uuid:
	uuid := ctx.Bool("uuid")

	// option --currency:
	currencyString := ctx.String("currency")
	currency, err := input.ParseOptionalCurrency(currencyString)
	if err != nil {
		return err
	}

	// option --end-time:
	endTimeString := ctx.String("end-time")
	endTime, err := input.ParseOptionalTime(endTimeString)
	if err != nil {
		return err
	}

	authData, err := credentials.ReadFromDefaultPath()
	if err != nil {
		return err
	}

	groshiClient := groshi.NewAPIClient(authData.URL, authData.Token)
	transactions, err := groshiClient.TransactionsReadMany(startTime, endTime, currency)
	if err != nil {
		return err
	}

	if err := output.Transactions(transactions, uuid); err != nil {
		return err
	}

	return nil
}

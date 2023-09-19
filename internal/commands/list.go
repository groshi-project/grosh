package commands

import (
	go_groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/groshi-project/grosh/internal/timeutil"
	"github.com/urfave/cli/v2"
	"time"
)

// ListCommand lists all transactions for given period:
// grosh list --end-time=<END-TIME> <START-TIME>
func ListCommand(ctx *cli.Context) error {
	args := ctx.Args()

	// required argument START_TIME
	startTimeString := args.Get(0)
	startTime, err := timeutil.ParseDate(startTimeString)
	if err != nil {
		return err
	}

	// option --end-time:
	endTimeString := ctx.String("end-time")

	var endTime *time.Time
	if endTimeString != "" {
		endTimeValue, err := timeutil.ParseDate(endTimeString)
		if err != nil {
			return err
		}
		endTime = &endTimeValue
	} else {
		endTime = nil
	}

	// option --currency:
	currencyString := ctx.String("currency")

	var currency *string
	if currencyString != "" {
		currency = &currencyString
	} else {
		currency = nil
	}

	authData, err := credentials.NewCredentialsFromCredentialsFile()
	if err != nil {
		return err
	}

	groshiClient := go_groshi.NewGroshiAPIClient(authData.URL, authData.JWT)
	transactions, err := groshiClient.TransactionsReadMany(startTime, endTime, currency)
	if err != nil {
		return err
	}

	if err := output.PrintTransactions(transactions); err != nil {
		return err
	}

	return nil
}

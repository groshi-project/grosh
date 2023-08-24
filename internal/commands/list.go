package commands

import (
	"errors"
	"github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/dates"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"time"
)

func ListCommand(ctx *cli.Context) error {
	args := ctx.Args()
	argsLen := args.Len()
	if argsLen > 1 {
		return ctx.Command.OnUsageError(
			ctx, errors.New("invalid number of arguments"), true,
		)
	}

	var period string
	if argsLen == 1 {
		period = args.Get(0)
	}
	startTimeString := ctx.String("start-time") // required option if period is not set
	endTimeString := ctx.String("end-time")     // optional option

	var startTime time.Time
	var endTime time.Time

	if period != "" { // if period is indicated
		if startTimeString != "" || endTimeString != "" {
			// return error if --start-time or --end-time are indicated
			return ctx.Command.OnUsageError(
				ctx,
				errors.New("PERIOD cannot be indicated if you use --start-time or --end-time options"),
				true,
			)
		}

		var err error
		startTime, endTime, err = dates.ParsePeriod(period)
		if err != nil {
			return err
		}

	} else {
		if startTimeString == "" {
			// return error if --start-time is not indicated
			return ctx.Command.OnUsageError(
				ctx,
				errors.New("if PERIOD is not indicated, at least --start-time option is required"),
				true,
			)
		}

		var err error
		startTime, err = dates.ParseDate(startTimeString)
		if err != nil {
			return err
		}

		if endTimeString == "" { // if --end-time is not indicated, use current time as the default value
			endTime = time.Now()
		}
	}

	authData, err := credentials.NewCredentialsFromCredentialsFile()
	if err != nil {
		return err
	}

	groshiClient := go_groshi.NewGroshiAPIClient(authData.URL, authData.JWT)
	response, err := groshiClient.TransactionsReadMany(startTime, &endTime)
	if err != nil {
		return err
	}

	transactions, err := response.SliceOfMaps()
	if err != nil {
		return err
	}

	if err := output.PrintTransactions(transactions); err != nil {
		return err
	}

	return nil
}

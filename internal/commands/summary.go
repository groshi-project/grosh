package commands

import (
	groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/input"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"time"
)

// SummaryCommand prints transactions summary for given period:
// groshi summary [--end-time=<END-TIME>] <START-TIME> <CURRENCY>
func SummaryCommand(ctx *cli.Context) error {
	args := ctx.Args()

	// START-TIME argument:
	startTimeString := args.Get(0)
	startTime, err := input.ParseTime(startTimeString)
	if err != nil {
		return err
	}

	// CURRENCY argument:
	currencyString := args.Get(1)
	currency, err := input.ParseCurrency(currencyString)
	if err != nil {
		return err
	}

	// --end-time param:
	endTimeString := ctx.String("end-time")
	endTime, err := input.ParseOptionalTime(endTimeString)
	if err != nil {
		return err
	}
	if endTime == nil {
		currentTime := time.Now()
		endTime = &currentTime
	}

	authData, err := credentials.ReadFromDefaultPath()
	if err != nil {
		return err
	}

	groshiClient := groshi.NewAPIClient(authData.URL, authData.Token)
	summary, err := groshiClient.TransactionsReadSummary(currency, startTime, endTime)
	if err != nil {
		return err
	}

	output.Stdout.Printf("There are %v transactions since %v to %v.", summary.TransactionsCount, startTime.Format(time.DateTime), endTime.Format(time.DateTime))
	output.Stdout.Printf("+%v %v, -%v %v.", summary.Income/100, summary.Currency, summary.Outcome/100, summary.Currency)
	output.Stdout.Printf("Total: %v %v.", summary.Total/100, summary.Currency)
	return nil
}

package commands

import (
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
)

// SummaryCommand prints transactions summary for given period:
// grosh summary --end-time=<END-TIME> <START-TIME> <CURRENCY>
func SummaryCommand(ctx *cli.Context) error {
	// required arguments:
	args := ctx.Args()
	startTimeString := args.Get(0)
	currency := args.Get(1)

	// opion --end-time:
	endTimeString := ctx.String("end-time")
	// todo....

	output.Stdout.Println(startTimeString, currency, endTimeString)
	return nil
}

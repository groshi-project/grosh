package commands

import (
	"errors"
	go_groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/dates"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
	"time"
)

func NewCommand(ctx *cli.Context) error {
	// grosh new <amount> <currency> --description= --date  # new transaction
	args := ctx.Args()
	if args.Len() != 2 {
		return ctx.Command.OnUsageError(ctx, errors.New("invalid count of arguments"), true)
	}

	// required arguments
	amountString := args.Get(0)
	currency := args.Get(1)

	// options
	description := ctx.String("description")
	dateString := ctx.String("date")

	// parse `amountString` as `amount`
	amount, err := strconv.ParseFloat(amountString, 32)
	if err != nil {
		return err
	}

	// currency to upper case
	currency = strings.ToUpper(currency) // todo: also validate currency code

	// parse `dateString` as `date`
	var date time.Time
	if dateString != "" {
		date, err = dates.ParseDate(dateString)
		if err != nil {
			return err
		}
	}

	authData, err := credentials.NewCredentialsFromCredentialsFile()
	if err != nil {
		return err
	}
	groshiClient := go_groshi.NewGroshiAPIClient(authData.URL, authData.JWT)
	rawResponse, err := groshiClient.TransactionsCreate(
		int(amount*100), currency, &description, &date,
	)
	if err != nil {
		return err
	}

	transaction, err := rawResponse.Map()
	if err != nil {
		return err
	}

	output.PlusLogger.Printf(
		"Successfully created new transaction (uuid: %v).", transaction["uuid"].(string),
	)

	return nil

}

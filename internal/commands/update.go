package commands

import (
	groshi "github.com/groshi-project/go-groshi"
	"github.com/groshi-project/grosh/internal/credentials"
	"github.com/groshi-project/grosh/internal/input"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"math"
)

// UpdateCommand updates transaction.
// grosh update [--amount=AMOUNT] [--currency=CURRENCY] [--description=DESCRIPTION] [--timestamp=TIME] <UUID>
func UpdateCommand(ctx *cli.Context) error {
	// required argument UUID:
	uuid := ctx.Args().Get(0)

	// --amount option:
	amountString := ctx.String("amount")
	amountFloat, err := input.ParseOptionalAmount(amountString)
	if err != nil {
		return err
	}

	var amountInt *int

	if amountFloat != nil {
		x := int(math.Round(*amountFloat * 100))
		amountInt = &x
	}

	// --currency option:
	currencyString := ctx.String("currency")
	currency, err := input.ParseOptionalCurrency(currencyString)
	if err != nil {
		return err
	}

	// --description option:
	descriptionString := ctx.String("description")
	description := input.ParseOptionalString(descriptionString)

	// --timestamp option:
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
	transaction, err := groshiClient.TransactionsUpdate(uuid, amountInt, currency, description, timestamp)
	if err != nil {
		return err
	}
	output.Plus.Printf("Successfully updated transaction %v", transaction.UUID)

	return nil
}

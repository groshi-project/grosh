package main

import (
	"github.com/groshi-project/grosh/internal/commands"
	"github.com/groshi-project/grosh/internal/middleware"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const ErrorExitCode = 1

func handleUsageError(ctx *cli.Context, err error, _ bool) error {
	output.Error.Println(err)
	output.Stdout.Printf("usage: %v", ctx.Command.UsageText)
	os.Exit(ErrorExitCode)
	return nil
}

func handleCommandNotFoundError(ctx *cli.Context, command string) {
	output.Error.Printf("'%v' is not a %v command. See `%v --help`", command, ctx.App.Name, ctx.App.Name)
	os.Exit(ErrorExitCode)
}

const categoryUser = "USER"
const categoryTransactions = "TRANSACTIONS"

func main() {
	app := &cli.App{
		Name:        "grosh",
		Usage:       "CLI groshi client",
		UsageText:   "grosh command [command options] [arguments...]",
		Version:     "0.1.0",
		Description: "grosh is a simple CLI client for groshi",

		Commands: []*cli.Command{
			// USER category:
			{
				Name:        "register",
				Category:    categoryUser,
				Usage:       "create new groshi user",
				UsageText:   "grosh register <URL> [USERNAME]",
				Description: "description",

				Action:       commands.RegisterCommand,
				OnUsageError: handleUsageError,
			},
			{
				Name:        "login",
				Category:    categoryUser,
				Usage:       "login to groshi server and store credentials",
				UsageText:   "grosh login <URL> [USERNAME]",
				Description: "description",

				Action:       commands.AuthCommand,
				OnUsageError: handleUsageError,
			},
			{
				Name:        "logout",
				Category:    categoryUser,
				Usage:       "remove locally stored credentials",
				UsageText:   "groshi logout",
				Description: "description",

				Action:       commands.CommandLogout,
				OnUsageError: handleUsageError,
			},

			// TRANSACTIONS category:
			{
				Name:        "new",
				Category:    categoryTransactions,
				Usage:       "create new transaction",
				UsageText:   "groshi new [--timestamp=<TIME>] [--description=<TEXT>] <AMOUNT> <CURRENCY>",
				Description: "create new transaction",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "description",
						Usage:   "description of transaction",
						Aliases: []string{"d"},
					},
					&cli.StringFlag{
						Name:  "timestamp",
						Usage: "date of transaction",
					},
				},

				Action:       middleware.ArgumentsCount(2, commands.NewCommand),
				OnUsageError: handleUsageError,
			},
			{
				Name:        "list",
				Category:    categoryTransactions,
				Usage:       "list transactions for given period and optionally in given currency",
				UsageText:   "groshi list [--uuid] [--currency=<CURRENCY>] [--end-time=<TIME>] <START-TIME>",
				Description: "list transactions for given period",

				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "uuid",
						Usage:   "display transaction UUIDs",
						Aliases: []string{"u"},
					},
					&cli.StringFlag{
						Name:    "end-time",
						Usage:   "end time",
						Aliases: []string{"e"},
					},
					&cli.StringFlag{
						Name:    "currency",
						Usage:   "currency",
						Aliases: []string{"c"},
					},
				},

				Action:       middleware.ArgumentsCount(1, commands.ListCommand),
				OnUsageError: handleUsageError,
			},
			{
				Name:        "summary",
				Category:    categoryTransactions,
				Aliases:     []string{"sum"},
				Usage:       "show summary of transactions for given period and optionally in given currency",
				UsageText:   "groshi summary [--end-time=<END-TIME>] <START-TIME> <CURRENCY>",
				Description: "description",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "end-time",
						Usage:   "end time",
						Aliases: []string{"e"},
					},
				},

				Action:       middleware.ArgumentsCount(2, commands.SummaryCommand),
				OnUsageError: handleUsageError,
			},
		},

		OnUsageError:    handleUsageError,
		CommandNotFound: handleCommandNotFoundError,

		HideHelpCommand: true,

		Authors: []*cli.Author{
			{"jieggii", "jieggii@protonmail.com"},
		},
		Copyright: "(c) groshi-project 2023",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

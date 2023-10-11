package main

import (
	"github.com/groshi-project/grosh/internal/commands"
	"github.com/groshi-project/grosh/internal/middleware"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"os"
)

const version = "0.1.0"
const errorExitCode = 1

func handleUsageError(ctx *cli.Context, err error, _ bool) error {
	output.Error.Println(err)
	output.Stdout.Printf("usage: %v", ctx.Command.UsageText)
	os.Exit(errorExitCode)
	return nil
}

func handleCommandNotFoundError(ctx *cli.Context, command string) {
	output.Error.Printf("'%v' is not a %v command. See `%v --help`", command, ctx.App.Name, ctx.App.Name)
	os.Exit(errorExitCode)
}

const categoryUser = "USER"
const categoryTransactions = "TRANSACTIONS"

func main() {
	app := &cli.App{
		Name:      "grosh",
		Usage:     "a command-line client for groshi",
		UsageText: "grosh <command> [COMMAND OPTIONS] [ARGUMENTS...]",
		Version:   version,

		Description: "grosh is a simple yet powerful command-line client for groshi",

		Commands: []*cli.Command{
			// USER category:
			{
				Name:        "register",
				Category:    categoryUser,
				Usage:       "create a new groshi user",
				UsageText:   "grosh register <URL> [USERNAME]",
				Description: "creates a new user at a groshi server",

				Action:       commands.RegisterCommand,
				OnUsageError: handleUsageError,
			},
			{
				Name:        "login",
				Category:    categoryUser,
				Usage:       "login to a groshi server and store credentials locally",
				UsageText:   "grosh login <URL> [USERNAME]",
				Description: "obtains authorization token and stores it aside with server URL in a file",

				Action:       commands.AuthCommand,
				OnUsageError: handleUsageError,
			},
			{
				Name:        "logout",
				Category:    categoryUser,
				Usage:       "remove locally stored credentials",
				UsageText:   "groshi logout",
				Description: "removes file containing user credentials to access groshi server",

				Action:       commands.CommandLogout,
				OnUsageError: handleUsageError,
			},

			// TRANSACTIONS category:
			{
				Name:        "new",
				Category:    categoryTransactions,
				Usage:       "create a new transaction",
				UsageText:   "groshi new [--timestamp=<TIME>] [--description=<TEXT>] <AMOUNT> <CURRENCY>",
				Description: "creates a new transaction",

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
				Usage:       "list transactions in given period and optionally in given currency",
				UsageText:   "groshi list [--uuid] [--currency=<CURRENCY>] [--end-time=<TIME>] <START-TIME>",
				Description: "retrieves list of all transactions in given period and optionally in given currency",

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
				Usage:       "show summary of transactions for given period in given currency",
				UsageText:   "groshi summary [--end-time=<END-TIME>] <START-TIME> <CURRENCY>",
				Description: "retrieves summary of transactions in given period and currency",

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

			{
				Name:        "remove",
				Category:    categoryTransactions,
				Aliases:     []string{"rm"},
				Usage:       "remove transaction",
				UsageText:   "groshi remove <UUID>",
				Description: "removes transaction by its UUID",

				Action:       middleware.ArgumentsCount(1, commands.RemoveCommand),
				OnUsageError: handleUsageError,
			},
		},

		OnUsageError:    handleUsageError,
		CommandNotFound: handleCommandNotFoundError,

		HideHelpCommand: true,

		Authors: []*cli.Author{
			{"jieggii", "jieggii@protonmail.com"},
		},
		Copyright: "(c) https://github.com/groshi-project 2023",
	}

	err := app.Run(os.Args)
	if err != nil {
		output.Error.Println(err)
	}
}

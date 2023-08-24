package main

import (
	"github.com/groshi-project/grosh/internal/commands"
	"github.com/groshi-project/grosh/internal/output"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const ErrorExitCode = 1

func handleUsageError(ctx *cli.Context, err error, _ bool) error {
	output.ErrorLogger.Println(err)
	output.StdoutLogger.Printf("usage: %v", ctx.Command.UsageText)
	os.Exit(ErrorExitCode)
	return nil
}

func handleCommandNotFoundError(ctx *cli.Context, command string) {
	output.ErrorLogger.Printf("'%v' is not a %v command. See `%v --help`", command, ctx.App.Name, ctx.App.Name)
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
				Usage:       "log out from groshi and remove locally stored credentials",
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
				UsageText:   "groshi new [--description] [--date] <AMOUNT> <CURRENCY> ",
				Description: "create new transaction",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "description",
						Usage:   "description of transaction",
						Aliases: []string{"d"},
					},
					&cli.StringFlag{
						Name:  "date",
						Usage: "date of transaction",
					},
				},

				Action:       commands.NewCommand,
				OnUsageError: handleUsageError,
			},
			{
				Name:        "list",
				Category:    categoryTransactions,
				Usage:       "list transactions for given period",
				UsageText:   "groshi list --start-time=<START_TIME> --end-time=<END_TIME> [PERIOD]",
				Description: "list transactions for given period",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "period",
						Usage:   "time period",
						Aliases: []string{"p"},
					},
					&cli.StringFlag{
						Name:    "start-time",
						Usage:   "start time",
						Aliases: []string{"s"},
					},
					&cli.StringFlag{
						Name:    "end-time",
						Usage:   "end time",
						Aliases: []string{"e"},
					},
				},

				Action:       commands.ListCommand,
				OnUsageError: handleUsageError,
			},
		},

		OnUsageError:    handleUsageError,
		CommandNotFound: handleCommandNotFoundError,

		HideHelpCommand: true,

		Authors: []*cli.Author{
			{"jieggii", "jieggii@protonmail.com"},
		},

		// Copyright of the binary if any
		Copyright: "(c) groshi-project 2023",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

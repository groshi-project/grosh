package middleware

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func ArgumentsCount(expectedArgsCount int, action cli.ActionFunc) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		argsLen := ctx.Args().Len()
		if argsLen != expectedArgsCount {
			return fmt.Errorf(
				"invalid number of arguments (expected %v, got %v)\nUsage: %v", expectedArgsCount, argsLen, ctx.Command.UsageText,
			)
		}
		return action(ctx)
	}
}

package middleware

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func ArgumentsCountStrict(count int, action cli.ActionFunc) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		argsLen := ctx.Args().Len()
		if argsLen != count {
			return fmt.Errorf(
				"invalid number of arguments (expected %v, got %v)\nUsage: %v",
				count, argsLen, ctx.Command.UsageText,
			)
		}
		return action(ctx)
	}
}

func ArgumentsCountVariable(minCount int, maxCount int, action cli.ActionFunc) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		argsLen := ctx.Args().Len()
		if argsLen < minCount || argsLen > maxCount {
			return fmt.Errorf(
				"invalid number of arguments (expected from %v to %v, got %v)\nUsage: %v",
				minCount, maxCount, argsLen, ctx.Command.UsageText,
			)
		}
		return action(ctx)
	}
}

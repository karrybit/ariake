package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	rootCommand := newRootCommand(ctx)
	rootCommand.AddCommand(newPrintCommand(ctx))
	rootCommand.AddCommand(newCreateCommand(ctx))
	rootCommand.AddCommand(newDeleteCommand(ctx))
	rootCommand.AddCommand(newPublishCommand(ctx))
	rootCommand.AddCommand(newSubscribeCommand(ctx))
	rootCommand.AddCommand(newResetCommand(ctx))

	if err := rootCommand.ExecuteContext(ctx); err != nil {
		log.Fatalln(err)
	}
}

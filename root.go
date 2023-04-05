package main

import (
	"context"

	"github.com/spf13/cobra"
)

func newRootCommand(ctx context.Context) *cobra.Command {
	rootCommand := cobra.Command{
		Use:   "ariake",
		Short: "ariake is utility tool to support development for google cloud pubsub in local environment",
		Args:  cobra.NoArgs,
	}
	rootCommand.PersistentFlags().String("host", "localhost:8085", "PUBSUB_EMULATOR_HOST")
	rootCommand.PersistentFlags().String("project_id", "dummy", "PUBSUB_PROJECT_ID")
	return &rootCommand
}

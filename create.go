package main

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

func newCreateCommand(ctx context.Context) *cobra.Command {
	command := cobra.Command{
		Use:   "create",
		Short: "Create a topic to register the emulator",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, args []string) {
			host := command.Flag("host").Value.String()
			projectID := command.Flag("project_id").Value.String()
			client, err := setupClient(ctx, host, projectID)
			if err != nil {
				log.Fatalln(err)
			}
			topicID := command.Flag("topic_id").Value.String()
			if _, err := client.CreateTopic(ctx, topicID); err != nil {
				log.Fatalln(err)
			}
		},
	}
	command.Flags().StringP("topic_id", "t", "", "topic id (required)")
	if err := command.MarkFlagRequired("topic_id"); err != nil {
		log.Fatalln(err)
	}
	return &command
}

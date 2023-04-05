package main

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

func newDeleteCommand(ctx context.Context) *cobra.Command {
	command := cobra.Command{
		Use:   "delete",
		Short: "Delete the topic in the emulator",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, args []string) {
			host := command.Flag("host").Value.String()
			projectID := command.Flag("project_id").Value.String()
			client, err := setupClient(ctx, host, projectID)
			if err != nil {
				log.Fatalln(err)
			}
			topicID := command.Flag("topic_id").Value.String()
			topic, err := getTopic(ctx, client, topicID)
			if err != nil {
				log.Fatalln(err)
			}
			if err := topic.Delete(ctx); err != nil {
				log.Fatalln(err)
			}
			subscriptions := topic.Subscriptions(ctx)
			for {
				subscription, err := subscriptions.Next()
				if err == iterator.Done {
					break
				}
				if err := subscription.Delete(ctx); err != nil {
					log.Fatalln(err)
				}
			}
		},
	}
	command.Flags().StringP("topic_id", "t", "", "topic id (required)")
	if err := command.MarkFlagRequired("topic_id"); err != nil {
		log.Fatalln(err)
	}
	return &command
}

package main

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

func newPublishCommand(ctx context.Context) *cobra.Command {
	command := cobra.Command{
		Use:   "publish",
		Short: "Publish a topic to the emulator",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, args []string) {
			host := command.Flag("host").Value.String()
			projectID := command.Flag("project_id").Value.String()
			client, err := setupClient(ctx, host, projectID)
			if err != nil {
				log.Fatalln(err)
			}
			topicID := command.Flag("topic_id").Value.String()
			message := command.Flag("message").Value.String()

			if err := publishTopic(ctx, client, topicID, message); err != nil {
				log.Fatalln(err)
			}
		},
	}
	command.Flags().StringP("topic_id", "t", "", "topic id (required)")
	command.Flags().StringP("message", "m", "", "message (required)")
	if err := command.MarkFlagRequired("topic_id"); err != nil {
		log.Fatalln(err)
	}
	if err := command.MarkFlagRequired("message"); err != nil {
		log.Fatalln(err)
	}
	return &command
}

// Publish a topic to the emulator
func publishTopic(ctx context.Context, client *pubsub.Client, topicID string, message string) error {
	topic, err := getTopic(ctx, client, topicID)
	if err != nil {
		return err
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	result := topic.Publish(ctx, &pubsub.Message{Data: data})
	_, err = result.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}

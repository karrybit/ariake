package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

func newResetCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset all topic and subscription",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, _ []string) {
			host := command.Flag("host").Value.String()
			projectID := command.Flag("project_id").Value.String()
			client, err := setupClient(ctx, host, projectID)
			if err != nil {
				log.Fatalln(err)
			}
			if err := resetResources(ctx, client); err != nil {
				log.Fatalln(err)
			}
		},
	}
}

// Reset all topic and subscription
func resetResources(ctx context.Context, client *pubsub.Client) error {
	if err := resetTopics(ctx, client); err != nil {
		return err
	}
	if err := resetSubscriptions(ctx, client); err != nil {
		return err
	}
	return nil
}

func resetTopics(ctx context.Context, client *pubsub.Client) error {
	topics := client.Topics(ctx)
	for {
		topic, err := topics.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if err := topic.Delete(ctx); err != nil {
			return err
		}
	}
	return nil
}

func resetSubscriptions(ctx context.Context, client *pubsub.Client) error {
	subscriptions := client.Subscriptions(ctx)
	for {
		subscription, err := subscriptions.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if err := subscription.Delete(ctx); err != nil {
			return err
		}
	}
	return nil
}

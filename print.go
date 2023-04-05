package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

func newPrintCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "print",
		Short: "Print the list of the topics registered in the emulator",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, _ []string) {
			host := command.Flag("host").Value.String()
			projectID := command.Flag("project_id").Value.String()
			client, err := setupClient(ctx, host, projectID)
			if err != nil {
				log.Fatalln(err)
			}
			if err := printResources(ctx, client); err != nil {
				log.Fatalln(err)
			}
		},
	}
}

// Print the list of the topics registered in the emulator
func printResources(ctx context.Context, client *pubsub.Client) error {
	topicIDs, err := collectTopicIDs(ctx, client)
	if err != nil {
		return err
	}
	subscriptionIDs, err := collectSubscriptionIDs(ctx, client)
	if err != nil {
		return err
	}
	type resources struct {
		TopicIDs        []string `json:"topics"`
		SubscriptionIDs []string `json:"subscriptions"`
	}
	r := resources{TopicIDs: topicIDs, SubscriptionIDs: subscriptionIDs}
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func collectTopicIDs(ctx context.Context, client *pubsub.Client) ([]string, error) {
	topics := client.Topics(ctx)
	topicIDs := make([]string, 0)
	for {
		topic, err := topics.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		topicIDs = append(topicIDs, topic.ID())
	}
	return topicIDs, nil
}

func collectSubscriptionIDs(ctx context.Context, client *pubsub.Client) ([]string, error) {
	subscriptions := client.Subscriptions(ctx)
	subscriptionIDs := make([]string, 0)
	for {
		subscription, err := subscriptions.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		subscriptionIDs = append(subscriptionIDs, subscription.ID())
	}
	return subscriptionIDs, nil
}

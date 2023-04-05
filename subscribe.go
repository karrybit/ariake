package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newSubscribeCommand(ctx context.Context) *cobra.Command {
	command := cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to topics in the emulator",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, args []string) {
			host := command.Flag("host").Value.String()
			projectID := command.Flag("project_id").Value.String()
			client, err := setupClient(ctx, host, projectID)
			if err != nil {
				log.Fatalln(err)
			}
			topicID := command.Flag("topic_id").Value.String()
			for {
				message, err := subscriptionTopic(ctx, client, topicID)
				if err != nil {
					log.Fatalln(err)
				}
				var out bytes.Buffer
				json.Indent(&out, []byte(*message), "", "  ")
				fmt.Println(out.String())
			}
		},
	}
	command.Flags().StringP("topic_id", "t", "", "topic id (required)")
	if err := command.MarkFlagRequired("topic_id"); err != nil {
		log.Fatalln(err)
	}
	return &command
}

// Subscribe to topics in the emulator
func subscriptionTopic(ctx context.Context, client *pubsub.Client, topicID string) (*string, error) {
	topic, err := getTopic(ctx, client, topicID)
	if err != nil {
		return nil, err
	}
	subscription, err := client.CreateSubscription(ctx, "sub-"+topicID+"-"+uuid.New().String(), pubsub.SubscriptionConfig{Topic: topic})
	if err != nil {
		return nil, err
	}

	cctx, cancel := context.WithCancel(ctx)
	type result struct {
		message string
		err     error
	}
	ch := make(chan result)

	go func() {
		err = subscription.Receive(cctx, func(_ context.Context, message *pubsub.Message) {
			defer cancel()
			message.Ack()
			unquoted, err := strconv.Unquote(string(message.Data))
			ch <- result{message: unquoted, err: err}
		})
	}()

	select {
	case result := <-ch:
		return &result.message, result.err
	}
}

package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// Get the topic in the emulator
func getTopic(ctx context.Context, client *pubsub.Client, topicID string) (*pubsub.Topic, error) {
	topic := client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("topic(%s) is not existed\n", topic.ID())
	}
	return topic, nil
}

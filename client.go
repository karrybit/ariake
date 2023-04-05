package main

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
)

func setupClient(ctx context.Context, host string, projectID string) (*pubsub.Client, error) {
	_, ok := os.LookupEnv("PUBSUB_EMULATOR_HOST")
	if !ok {
		os.Setenv("PUBSUB_EMULATOR_HOST", host)
	}
	return pubsub.NewClient(ctx, projectID)
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func main() {
	secretID := flag.String("secret-id", "", "Your Yandex.Cloud Lockbox ID of the secret")
	versionID := flag.String("version-id", "", "Optional version ID of the secret")
	flag.Parse()

	if *secretID == "" {
		log.Fatal("secret-id is required")
	}

	ctx := context.Background()
	log.Println("Using instance metadata service for authentication")

	// Create SDK with InstanceServiceAccount
	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: ycsdk.InstanceServiceAccount(),
	})
	if err != nil {
		log.Fatal("Failed to build SDK: ", err)
	}

	log.Printf("Requesting payload for secret: %s", *secretID)

	// Prepare the request
	request := &lockbox.GetPayloadRequest{
		SecretId: *secretID,
	}

	// Add version ID if provided
	if *versionID != "" {
		request.VersionId = *versionID
		log.Printf("Using specific version: %s", *versionID)
	}

	// Get the payload of the secret using the correct method
	payload, err := sdk.LockboxPayload().Payload().Get(ctx, request)
	if err != nil {
		log.Fatal("Failed to get secret payload: ", err)
	}

	// Print the payload entries in a more readable format
	fmt.Println("Secret payload:")
	fmt.Printf("Version ID: %s\n", payload.VersionId)
	fmt.Println("Entries:")
	for _, entry := range payload.Entries {
		fmt.Printf("  Key: %s, Value: %s\n", entry.Key, entry.Value)
	}
}

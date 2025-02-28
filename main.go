package main

import (
	"context"
	"flag"
	"log"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func main() {
	folderID := flag.String("folder-id", "", "Your Yandex.Cloud Lockbox ID of the folder")
	flag.Parse()

	if *folderID == "" {
		log.Fatal("folder-id is required")
	}

	ctx := context.Background()
	log.Println("Using instance metadata service for authentication")

	// Create SDK with InstanceServiceAccount which will use metadata token
	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: ycsdk.InstanceServiceAccount(),
	})
	if err != nil {
		log.Fatal("Failed to build SDK: ", err)
	}

	log.Println("Listing secrets in folder:", *folderID)
	p, err := sdk.LockboxSecret().Secret().List(ctx, &lockbox.ListSecretsRequest{
		FolderId: *folderID,
	})
	if err != nil {
		log.Fatal("Failed to list secrets: ", err)
	}

	log.Println("Found secrets:")
	for i, secret := range p.Secrets {
		log.Printf("%d. ID: %s, Name: %s", i+1, secret.Id, secret.Name)
	}
}

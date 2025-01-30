package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tab/smartid"
	"github.com/tab/smartid/config"
)

func main() {
	// Full example of how to use the Smart-ID client
	client, err := smartid.NewClient(
		config.WithRelyingPartyName("DEMO"),
		config.WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000"),
		config.WithCertificateLevel("QUALIFIED"),
		config.WithHashType("SHA512"),
		config.WithInteractionType("displayTextAndPIN"),
		config.WithText("Enter PIN1"),
		config.WithURL("https://sid.demo.sk.ee/smart-id-rp/v2"),
		config.WithTimeout(60),
	)

	if err != nil {
		log.Fatalf("Failed to create Smart-ID client: %v", err)
	}

	ctx := context.Background()
	identity := smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914")
	result, err := client.Authenticate(ctx, identity)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	fmt.Println("result")
	fmt.Println(result)
}

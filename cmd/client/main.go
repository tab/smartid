package main

import (
	"fmt"
	"log"

	"smartid/internal/client"
	"smartid/internal/config"
)

func main() {
	// Full example of how to use the Smart-ID client
	smartIdClient, err := client.NewClient(
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

	identity := client.NewIdentity(client.TypePNO, "EE", "30303039914")
	response, err := smartIdClient.Authenticate(identity)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	fmt.Println("response")
	fmt.Println(response)

	session, err := smartIdClient.Status(response.SessionID)
	if err != nil {
		log.Fatalf("Failed to get session status: %v", err)
	}

	fmt.Println("session")
	fmt.Println(session)
}

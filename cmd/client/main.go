package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tab/smartid"
)

func main() {
	client := smartid.NewClient().
		WithRelyingPartyName("DEMO").
		WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000").
		WithCertificateLevel("QUALIFIED").
		WithHashType("SHA512").
		WithInteractionType("displayTextAndPIN").
		WithText("Enter PIN1").
		WithURL("https://sid.demo.sk.ee/smart-id-rp/v2").
		WithTimeout(60 * time.Second)

	ctx := context.Background()
	identity := smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914")
	result, err := client.Authenticate(ctx, identity)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	fmt.Println("result")
	fmt.Println(result)
}

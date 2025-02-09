package main

import (
	"context"
	"errors"
	"fmt"
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

	identities := []string{
		smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039917"),
	}

	for _, identity := range identities {
		session, resultCh := client.Authenticate(ctx, identity)
		fmt.Println(session)

		result := <-resultCh

		if result.Err != nil {
			var err *smartid.Error

			if ok := errors.As(result.Err, &err); ok {
				fmt.Printf("Authentication failed with code: %s\n", err.Code)
			} else {
				fmt.Printf("Authentication failed: %v\n", result.Err)
			}
		} else {
			fmt.Println(result.Person)
		}
	}
}

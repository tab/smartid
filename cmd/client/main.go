package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/tab/smartid"
	"github.com/tab/smartid/internal/certificates"
)

func main() {
	pinner, err := certificates.NewCertificatePinner("./certs")
	if err != nil {
		fmt.Println("Failed to create certificate pinner:", err)
	}
	tlsConfig := pinner.TLSConfig()

	client := smartid.NewClient().
		WithRelyingPartyName("DEMO").
		WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000").
		WithCertificateLevel("QUALIFIED").
		WithHashType("SHA512").
		WithInteractionType("displayTextAndPIN").
		WithText("Enter PIN1").
		WithURL("https://sid.demo.sk.ee/smart-id-rp/v2").
		WithTimeout(60 * time.Second).
		WithTLSConfig(tlsConfig)

	identities := []string{
		smartid.NewIdentity(smartid.TypePNO, "EE", "40504040001"),
		smartid.NewIdentity(smartid.TypePNO, "LT", "40504040001"),
		smartid.NewIdentity(smartid.TypePNO, "LV", "050405-10009"),
		smartid.NewIdentity(smartid.TypePNO, "BE", "05040400032"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914"),

		// NOTE: USER_REFUSED
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039917"),

		// NOTE: USER_REFUSED_DISPLAYTEXTANDPIN
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039928"),

		// NOTE: USER_REFUSED_VC_CHOICE
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039939"),

		// NOTE: USER_REFUSED_CONFIRMATIONMESSAGE
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039946"),

		// NOTE: USER_REFUSED_CONFIRMATIONMESSAGE_WITH_VC_CHOICE
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039950"),

		// NOTE: USER_REFUSED_CERT_CHOICE
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039961"),

		// NOTE: WRONG_VC
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039972"),

		// NOTE: TIMEOUT
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039983"),
	}

	worker := smartid.NewWorker(client).
		WithConcurrency(50).
		WithQueueSize(100)

	ctx := context.Background()

	worker.Start(ctx)
	defer worker.Stop()

	var wg sync.WaitGroup

	for _, identity := range identities {
		wg.Add(1)

		session, err := client.CreateSession(ctx, identity)
		if err != nil {
			fmt.Println("Error creating session:", err)
			wg.Done()
			continue
		}
		fmt.Println("Session created:", session)

		resultCh := worker.Process(ctx, session.Id)
		go func() {
			defer wg.Done()
			result := <-resultCh
			if result.Err != nil {
				fmt.Println("Error fetching session:", result.Err)
			} else {
				fmt.Println("Fetched person:", result.Person)
			}
		}()
	}

	wg.Wait()
}

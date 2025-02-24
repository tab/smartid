//go:build example
// +build example

package examples

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/tab/smartid"
)

func Example_CreateSession() {
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

	session, err := client.CreateSession(ctx, identity)
	if err != nil {
		fmt.Println("Error creating session:", err)
	}
	fmt.Println("Session created:", session)
}

func Example_FetchSession() {
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

	sessionId := "d3b8f7c3-7e0c-4a4e-9e6b-4b0e6b8e4f4c"

	person, err := client.FetchSession(ctx, sessionId)
	if err != nil {
		fmt.Println("Error fetching session:", err)
		return
	}

	fmt.Println("Person:", person)
}

func Example_ProcessMultipleIdentitiesInBackground() {
	client := smartid.NewClient().
		WithRelyingPartyName("DEMO").
		WithRelyingPartyUUID("00000000-0000-0000-0000-000000000000").
		WithCertificateLevel("QUALIFIED").
		WithHashType("SHA512").
		WithInteractionType("displayTextAndPIN").
		WithText("Enter PIN1").
		WithURL("https://sid.demo.sk.ee/smart-id-rp/v2").
		WithTimeout(60 * time.Second)

	identities := []string{
		smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039917"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039928"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039939"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039946"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039950"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039961"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039972"),
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

func Example_ProcessMultipleIdentitiesInBackground_WithTLS() {
	manager, err := smartid.NewCertificateManager("./certs")
	if err != nil {
		fmt.Println("Failed to create certificate manager:", err)
	}
	tlsConfig := manager.TLSConfig()

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
		smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039917"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039928"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039939"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039946"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039950"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039961"),
		smartid.NewIdentity(smartid.TypePNO, "EE", "30403039972"),
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

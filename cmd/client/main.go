package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/tab/smartid"
	"github.com/tab/smartid/internal/config"
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

	ctx := context.Background()

	worker := smartid.NewWorker(ctx, client)
	worker.WithConfig(config.WorkerConfig{
		Concurrency: 50,
		QueueSize:   100,
	})

	worker.Start()
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

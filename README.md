# Go Smart-ID client

Golang client for the Smart-ID API (https://www.smart-id.com).
It is a simple wrapper around the API, which helps easily integrate Smart-ID authentication into Golang applications.

## Installation

Use `go get` to install the package

```sh
go get github.com/tab/smartid
```

## Usage

### Creating a Client

Create a new client using `NewClient()` and customize its configuration using chainable methods.

```go
package main

import (
  "context"
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

  if err := client.Validate(); err != nil {
    log.Fatal("Invalid configuration:", err)
  }

  // Further processing...
}
```

### Start Authentication

Initiate a new authentication session with the `Smart-ID` provider by calling `CreateSession`.
This function generates a random hash, constructs the session request, and returns a session that includes an identifier and a verification code.

```go
func main() {
  // Create a client...

  identity := smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914")

  session, err := client.CreateSession(context.Background(), identity)
  if err != nil {
    log.Fatal("Error creating session:", err)
  }

  fmt.Println("Session created:", session)
}
```

### Fetch Session

```go
func main() {
  // Create a client...

  person, err := client.FetchSession(context.Background(), sessionId)
  if err != nil {
    log.Fatal("Error fetching session:", err)
  }

  fmt.Println("Session status:", session.State)
}
```

### Async Example

For applications requiring the processing of multiple authentication sessions simultaneously, `Smart-ID` provides a worker model.
Create a worker using `NewWorker`, configure its concurrency and queue size, and then start processing.

```go
package main

import (
  "context"
  "fmt"
  "log"
  "sync"
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

  identities := []string{
    smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914"),
    smartid.NewIdentity(smartid.TypePNO, "EE", "30403039917"),
    smartid.NewIdentity(smartid.TypePNO, "EE", "30403039928"),
    // Add more identities as needed
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
      log.Println("Error creating session:", err)
      wg.Done()
      continue
    }
    fmt.Println("Session created:", session)

    resultCh := worker.Process(ctx, session.Id)
    go func() {
      defer wg.Done()
      result := <-resultCh
      if result.Err != nil {
        log.Println("Error fetching session:", result.Err)
      } else {
        fmt.Println("Fetched person:", result.Person)
      }
    }()
  }

  wg.Wait()
}
```

### Certificate pinning (optional)

```go
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

  // Further processing...
```

### Prepare Identity

Smart-ID requires a properly formatted identity string. Use the `NewIdentity` function to create this string.
It combines the identity type, country code, and the identifier value.

```go
package main

import (
  "fmt"

  "github.com/tab/smartid"
)

func main() {
  identity := smartid.NewIdentity(smartid.TypePNO, "EE", "30303039914")
  fmt.Println("Formatted identity:", identity)
}
```

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Acknowledgements

- [SK ID Solutions](https://www.skidsolutions.eu)
- [Smart-ID](https://www.smart-id.com)

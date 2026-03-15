# omada-go-sdk

Golang SDK for the TP-Link Omada OpenAPI.

## What this repo contains

- Generated API client in `client/`
- Lightweight Omada auth helpers in `client/auth/`
- Runnable usage examples in `examples/`
- Code generation tooling in `tools/`

## Install

Install from GitHub:

```bash
go get github.com/Tohaker/omada-go-sdk
```

## Quick usage

The common flow is:

1. Create a client.
2. Get an access token using Omada auth helpers.
3. Create an API client that injects `Authorization: AccessToken=<token>`.
4. Call generated endpoints.

### Simple inline example

```go
package main

import (
	"context"
	"fmt"
	"log"

	client "github.com/Tohaker/omada-go-sdk/client"
)

func main() {
	ctx := context.Background()

	baseURL := "https://your-omada-controller"
	omadacID := "your-omadac-id"
	clientID := "your-client-id"
	clientSecret := "your-client-secret"

	authClient, err := client.NewClientWithResponses(baseURL)
	if err != nil {
		log.Fatalf("create auth client: %v", err)
	}

	tokens, err := authClient.OmadaTokenByClientCredentials(ctx, omadacID, clientID, clientSecret)
	if err != nil {
		log.Fatalf("get token: %v", err)
	}

	api, err := client.NewClientWithResponses(
		baseURL,
		client.WithOmadaAccessToken(tokens.AccessToken),
	)
	if err != nil {
		log.Fatalf("create api client: %v", err)
	}

	sites, err := api.GetSiteListWithResponse(ctx, omadacID, &client.GetSiteListParams{Page: 1, PageSize: 10})
	if err != nil {
		log.Fatalf("get sites: %v", err)
	}

	fmt.Printf("status=%d\n", sites.StatusCode())
}
```

## Examples

- See the `examples/` directory for runnable examples.
- Start with `examples/create_client.go`.

## Code generation and tools

Generation tooling is documented in [the tools directory](./tools/README.md).

From repo root, regenerate client code with:

```bash
cd tools
go generate ./...
```

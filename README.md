# omada-go-sdk

Go SDK for the TP-Link Omada Open API, auto-generated from the [Omada OpenAPI spec](https://use1-omada-northbound.tplinkcloud.com/doc.html#/home).

> **Note:** The upstream OpenAPI spec is not fully valid. This SDK uses a spec-fixer tool to patch known issues before generation. TP-Link is working on improvements, see [this forum post](https://community.tp-link.com/en/business/forum/topic/859880)

## Installation

```bash
go get github.com/Tohaker/omada-go-sdk
```

## Quick Start

### Client Credentials Flow

```go
package main

import (
	"context"
	"fmt"
	"log"

	omada "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	cfg := omada.NewConfiguration()
	cfg.Servers = omada.ServerConfigurations{
		{URL: "https://use1-omada-northbound.tplinkcloud.com"},
	}
	client := omada.NewAPIClient(cfg)

	// 1. Get access token via client credentials
	tokenResp, _, err := client.AuthorizeAPI.AuthorizeToken(context.Background()).
		GrantType("client_credentials").
		TokenRequest(omada.TokenRequest{
			ClientId:     "your-client-id",
			ClientSecret: "your-client-secret",
			OmadacId:     omada.PtrString("your-omadac-id"),
		}).Execute()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Set the access token for subsequent requests
	cfg.DefaultHeader["Authorization"] = "AccessToken=" + *tokenResp.Result.AccessToken

	// 3. Use API endpoints
	// sites, _, err := client.SiteAPI.GetSites(context.Background(), "omadac-id").Execute()
	fmt.Println("Access token obtained successfully")
}
```

### Authorization Code Flow

```go
// 1. Login
loginResp, _, _ := client.AuthorizeAPI.AuthorizeLogin(context.Background()).
	ClientId("your-client-id").
	OmadacId("your-omadac-id").
	LoginRequest(omada.LoginRequest{
		Username: "admin",
		Password: "password",
	}).Execute()

// 2. Get authorization code
codeResp, _, _ := client.AuthorizeAPI.AuthorizeCode(context.Background()).
	ClientId("your-client-id").
	OmadacId("your-omadac-id").
	ResponseType("code").
	CsrfToken(*loginResp.Result.CsrfToken).
	Cookie("TPOMADA_SESSIONID=" + *loginResp.Result.SessionId).
	Execute()

// 3. Exchange code for token
tokenResp, _, _ := client.AuthorizeAPI.AuthorizeToken(context.Background()).
	GrantType("authorization_code").
	Code(*codeResp.Result).
	TokenRequest(omada.TokenRequest{
		ClientId:     "your-client-id",
		ClientSecret: "your-client-secret",
	}).Execute()
```

## Regions

| Region      | Server URL                                      |
| ----------- | ----------------------------------------------- |
| US          | `https://use1-omada-northbound.tplinkcloud.com` |
| EU          | `https://euw1-omada-northbound.tplinkcloud.com` |
| Singapore   | `https://aps1-omada-northbound.tplinkcloud.com` |
| Self-hosted | Your Omada Controller URL                       |

## Regenerating the SDK

Prerequisites: Go 1.21+, Docker.

```bash
# Run the full generation pipeline
./generator/generate.sh
```

Or use the GitHub Actions workflow: **Actions → Generate SDK → Run workflow** with a version tag (e.g. `v0.1.1`).

## Releasing a new version

This repository uses a two-phase release flow with
[`go-changesets`](https://github.com/jakoblorz/go-changesets):

1. [`changesets-version-pr`](./.github/workflows/changesets-version-pr.yml) reads pending `.changeset/*.md` files and opens/updates a release PR.
2. [`changesets-publish`](./.github/workflows/changesets-publish.yml) runs after merge to `main`, regenerates/verifies the SDK, and publishes tags/releases when `version.txt` is newer than the latest tag.

### Add a changeset

Install locally (optional):

```bash
go install github.com/jakoblorz/go-changesets/cmd/changeset@v0.0.9
```

Create a changeset file interactively:

```bash
changeset add
```

Changeset files are committed under `.changeset/` and consumed automatically by the version PR workflow.

## How It Works

1. **`tools/fix_spec.go`** fetches the upstream OpenAPI spec and fixes known issues (schema names with spaces/special chars, missing security definitions, HTTP→HTTPS)
2. **`openapi-generator-cli`** (pinned to v7.12.0 via Docker) generates the Go client SDK from the fixed spec
3. A separate generation pass adds auth endpoints from `generator/auth-spec.yaml`
4. The result is a single `omada` Go package in the `omada/` subdirectory

## Publishing docs

Docs are published using [`mkdocs-material`](https://squidfunk.github.io/mkdocs-material/).

If you've opened this repository in the `devcontainer`, `mkdocs` and all required plugins will already be installed and configured for you.

Otherwise, if you haven't done so already, follow their [getting started](https://squidfunk.github.io/mkdocs-material/getting-started) guide to set up your local environment.

You must also install the following extra plugins;
- `mkdocs-exclude`

To preview the docs locally, just run

```sh
mkdocs serve
```

Or build to the `/site` directory with 

```sh
mkdocs build
```

Whenever docs are generated with `go generate` and commited to the main branch, the docs will deploy and be available to view at `tohaker.github.io/omada-go-sdk`.


## License

See [LICENSE](LICENSE).

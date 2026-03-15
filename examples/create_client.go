package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	client "github.com/Tohaker/omada-go-sdk/client"
)

// CreateClient demonstrates how to initialize and use the generated clients.

func CreateClient() error {
	baseURL := os.Getenv("BASE_URL")
	omadacID := os.Getenv("OMADA_CUSTOMER_ID")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	macAddress := os.Getenv("GATEWAY_MAC_ADDRESS")

	insecureTLS := true

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureTLS}, //nolint:gosec // Local/self-signed controller usage.
		},
	}

	authClient, err := client.NewClientWithResponses(
		baseURL,
		client.WithHTTPClient(httpClient),
	)
	if err != nil {
		return fmt.Errorf("create auth client: %w", err)
	}

	ctx := context.Background()
	tokens, err := authClient.OmadaTokenByClientCredentials(ctx, omadacID, clientID, clientSecret)
	if err != nil {
		return fmt.Errorf("get access token: %w", err)
	}

	api, err := client.NewClientWithResponses(
		baseURL,
		client.WithHTTPClient(httpClient),
		client.WithRequestEditorFn(client.OmadaAccessTokenRequestEditor(tokens.AccessToken)),
	)
	if err != nil {
		return fmt.Errorf("create api client: %w", err)
	}

	params := &client.GetSiteListParams{
		Page:     1,
		PageSize: 10,
	}

	sites, err := api.GetSiteListWithResponse(ctx, omadacID, params)
	if err != nil {
		return fmt.Errorf("GetSiteList call failed: %w", err)
	}

	site := (*sites.JSON200.Result.Data)[0].SiteId

	fmt.Printf("site %s\n", *site)

	// Example core request.
	gatewayInfo, err := api.GetGatewayInfoWithResponse(ctx, omadacID, *site, macAddress)
	if err != nil {
		return fmt.Errorf("GetGatewayInfo call failed: %w", err)
	}
	fmt.Printf("GetGatewayInfo status=%s body=%s\n", gatewayInfo.Status(), *gatewayInfo.JSON200.Msg)

	return nil
}

func main() {
	if err := CreateClient(); err != nil {
		panic(err)
	}
}

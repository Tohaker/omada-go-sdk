# QuickActionAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWizardDefaultNetwork**](QuickActionAPI.md#getnetworkwizarddefaultnetwork) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/quick-action/network-wizard/default-network | Get default network of the network wizard



## GetNetworkWizardDefaultNetwork

> OperationResponseLanNetworkOpenApiVO GetNetworkWizardDefaultNetwork(ctx, omadacId, siteId).Execute()

Get default network of the network wizard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuickActionAPI.GetNetworkWizardDefaultNetwork(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickActionAPI.GetNetworkWizardDefaultNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWizardDefaultNetwork`: OperationResponseLanNetworkOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `QuickActionAPI.GetNetworkWizardDefaultNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWizardDefaultNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseLanNetworkOpenApiVO**](OperationResponseLanNetworkOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


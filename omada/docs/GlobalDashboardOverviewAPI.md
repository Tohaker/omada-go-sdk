# \GlobalDashboardOverviewAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGernalSettings1**](GlobalDashboardOverviewAPI.md#GetGernalSettings1) | **Get** /openapi/v1/{omadacId}/dashboard/overview-without-client | Get global dashboard overview



## GetGernalSettings1

> OperationResponseGlobalOverViewOpenApiVO GetGernalSettings1(ctx, omadacId).Execute()

Get global dashboard overview



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalDashboardOverviewAPI.GetGernalSettings1(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalDashboardOverviewAPI.GetGernalSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGernalSettings1`: OperationResponseGlobalOverViewOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GlobalDashboardOverviewAPI.GetGernalSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGernalSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseGlobalOverViewOpenApiVO**](OperationResponseGlobalOverViewOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


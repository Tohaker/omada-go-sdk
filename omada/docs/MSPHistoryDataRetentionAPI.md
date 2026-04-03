# MSPHistoryDataRetentionAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMspDataRetention**](MSPHistoryDataRetentionAPI.md#getmspdataretention) | **Get** /openapi/v1/msp/{mspId}/retention | Get MSP history data retention configuration
[**ModifyMspRetention**](MSPHistoryDataRetentionAPI.md#modifymspretention) | **Patch** /openapi/v1/msp/{mspId}/retention | Modify MSP history data retention configuration



## GetMspDataRetention

> OperationResponseHistoryRetention GetMspDataRetention(ctx, mspId).Execute()

Get MSP history data retention configuration



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
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPHistoryDataRetentionAPI.GetMspDataRetention(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPHistoryDataRetentionAPI.GetMspDataRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspDataRetention`: OperationResponseHistoryRetention
	fmt.Fprintf(os.Stdout, "Response from `MSPHistoryDataRetentionAPI.GetMspDataRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspDataRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseHistoryRetention**](OperationResponseHistoryRetention.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMspRetention

> OperationResponseWithoutResult ModifyMspRetention(ctx, mspId).ModifyHistoryRetentionOpenApiVO(modifyHistoryRetentionOpenApiVO).Execute()

Modify MSP history data retention configuration



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
	mspId := "mspId_example" // string | MSP ID
	modifyHistoryRetentionOpenApiVO := *openapiclient.NewModifyHistoryRetentionOpenApiVO(false) // ModifyHistoryRetentionOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPHistoryDataRetentionAPI.ModifyMspRetention(context.Background(), mspId).ModifyHistoryRetentionOpenApiVO(modifyHistoryRetentionOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPHistoryDataRetentionAPI.ModifyMspRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspRetention`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MSPHistoryDataRetentionAPI.ModifyMspRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyHistoryRetentionOpenApiVO** | [**ModifyHistoryRetentionOpenApiVO**](ModifyHistoryRetentionOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


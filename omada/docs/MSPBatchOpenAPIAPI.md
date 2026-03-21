# \MSPBatchOpenAPIAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MspBatchController**](MSPBatchOpenAPIAPI.md#MspBatchController) | **Post** /openapi/v1/msp/{mspId}/batch | Batch Processing MSP OpenAPIs



## MspBatchController

> OperationResponseBatchResponseEntity MspBatchController(ctx, mspId).BatchRequestEntity(batchRequestEntity).Execute()

Batch Processing MSP OpenAPIs



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
	batchRequestEntity := *openapiclient.NewBatchRequestEntity() // BatchRequestEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPBatchOpenAPIAPI.MspBatchController(context.Background(), mspId).BatchRequestEntity(batchRequestEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPBatchOpenAPIAPI.MspBatchController``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MspBatchController`: OperationResponseBatchResponseEntity
	fmt.Fprintf(os.Stdout, "Response from `MSPBatchOpenAPIAPI.MspBatchController`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMspBatchControllerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchRequestEntity** | [**BatchRequestEntity**](BatchRequestEntity.md) |  | 

### Return type

[**OperationResponseBatchResponseEntity**](OperationResponseBatchResponseEntity.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


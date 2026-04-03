# BatchOpenAPIAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchController**](BatchOpenAPIAPI.md#batchcontroller) | **Post** /openapi/v1/{omadacId}/batch | Batch Processing OpenAPIs



## BatchController

> OperationResponseBatchResponseEntity BatchController(ctx, omadacId).BatchRequestEntity(batchRequestEntity).Execute()

Batch Processing OpenAPIs



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
	batchRequestEntity := *openapiclient.NewBatchRequestEntity() // BatchRequestEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchOpenAPIAPI.BatchController(context.Background(), omadacId).BatchRequestEntity(batchRequestEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchOpenAPIAPI.BatchController``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchController`: OperationResponseBatchResponseEntity
	fmt.Fprintf(os.Stdout, "Response from `BatchOpenAPIAPI.BatchController`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchControllerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchRequestEntity** | [**BatchRequestEntity**](BatchRequestEntity.md) |  | 

### Return type

[**OperationResponseBatchResponseEntity**](OperationResponseBatchResponseEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


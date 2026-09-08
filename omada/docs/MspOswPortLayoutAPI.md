# MspOswPortLayoutAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMspOswPortLayoutVO**](MspOswPortLayoutAPI.md#getmsposwportlayoutvo) | **Get** /openapi/v1/msp/{mspId}/port/layout | Get all msp osw port layout.



## GetMspOswPortLayoutVO

> OperationResponse GetMspOswPortLayoutVO(ctx, mspId).Execute()

Get all msp osw port layout.



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
	resp, r, err := apiClient.MspOswPortLayoutAPI.GetMspOswPortLayoutVO(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspOswPortLayoutAPI.GetMspOswPortLayoutVO``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOswPortLayoutVO`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspOswPortLayoutAPI.GetMspOswPortLayoutVO`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOswPortLayoutVORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# LanMulticastAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLanMulticast**](LanMulticastAPI.md#createlanmulticast) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/lan-multicasts | Create new lan multicast
[**DeleteLanMulticast**](LanMulticastAPI.md#deletelanmulticast) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/lan-multicasts/{multicastId} | Delete an existing lan multicast
[**GetGridLanMulticasts**](LanMulticastAPI.md#getgridlanmulticasts) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lan-multicasts | Get lan multicast list
[**ModifyLanMulticast**](LanMulticastAPI.md#modifylanmulticast) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/lan-multicasts/{multicastId} | Modify a lan multicast



## CreateLanMulticast

> ResponseIdVO CreateLanMulticast(ctx, omadacId, siteId).LanMulticastVO(lanMulticastVO).Execute()

Create new lan multicast



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
	lanMulticastVO := *openapiclient.NewLanMulticastVO() // LanMulticastVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanMulticastAPI.CreateLanMulticast(context.Background(), omadacId, siteId).LanMulticastVO(lanMulticastVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanMulticastAPI.CreateLanMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanMulticast`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `LanMulticastAPI.CreateLanMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanMulticastVO** | [**LanMulticastVO**](LanMulticastVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanMulticast

> OperationResponseWithoutResult DeleteLanMulticast(ctx, omadacId, siteId, multicastId).Execute()

Delete an existing lan multicast



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
	multicastId := "multicastId_example" // string | lanMulticastId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanMulticastAPI.DeleteLanMulticast(context.Background(), omadacId, siteId, multicastId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanMulticastAPI.DeleteLanMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLanMulticast`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LanMulticastAPI.DeleteLanMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**multicastId** | **string** | lanMulticastId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridLanMulticasts

> OperationResponse GetGridLanMulticasts(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get lan multicast list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanMulticastAPI.GetGridLanMulticasts(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanMulticastAPI.GetGridLanMulticasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridLanMulticasts`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `LanMulticastAPI.GetGridLanMulticasts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridLanMulticastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



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


## ModifyLanMulticast

> OperationResponseWithoutResult ModifyLanMulticast(ctx, omadacId, siteId, multicastId).LanMulticastVO(lanMulticastVO).Execute()

Modify a lan multicast



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
	multicastId := "multicastId_example" // string | lanMulticastId
	lanMulticastVO := *openapiclient.NewLanMulticastVO() // LanMulticastVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanMulticastAPI.ModifyLanMulticast(context.Background(), omadacId, siteId, multicastId).LanMulticastVO(lanMulticastVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanMulticastAPI.ModifyLanMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanMulticast`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LanMulticastAPI.ModifyLanMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**multicastId** | **string** | lanMulticastId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanMulticastVO** | [**LanMulticastVO**](LanMulticastVO.md) |  | 

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


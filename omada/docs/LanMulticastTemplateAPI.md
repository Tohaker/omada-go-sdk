# \LanMulticastTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLanMulticastTemplate**](LanMulticastTemplateAPI.md#CreateLanMulticastTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-multicasts | Create new lan multicast template
[**DeleteLanMulticastTemplate**](LanMulticastTemplateAPI.md#DeleteLanMulticastTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-multicasts/{multicastId} | Delete an existing lan multicast template
[**GetGridLanMulticastTemplates**](LanMulticastTemplateAPI.md#GetGridLanMulticastTemplates) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-multicasts | Get lan multicast template list
[**ModifyLanMulticastTemplate**](LanMulticastTemplateAPI.md#ModifyLanMulticastTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-multicasts/{multicastId} | Modify a lan multicast template



## CreateLanMulticastTemplate

> ResponseIdVO CreateLanMulticastTemplate(ctx, omadacId, siteTemplateId).LanMulticastVO(lanMulticastVO).Execute()

Create new lan multicast template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	lanMulticastVO := *openapiclient.NewLanMulticastVO() // LanMulticastVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanMulticastTemplateAPI.CreateLanMulticastTemplate(context.Background(), omadacId, siteTemplateId).LanMulticastVO(lanMulticastVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanMulticastTemplateAPI.CreateLanMulticastTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanMulticastTemplate`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `LanMulticastTemplateAPI.CreateLanMulticastTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanMulticastTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanMulticastVO** | [**LanMulticastVO**](LanMulticastVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanMulticastTemplate

> OperationResponseWithoutResult DeleteLanMulticastTemplate(ctx, omadacId, siteTemplateId, multicastId).Execute()

Delete an existing lan multicast template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	multicastId := "multicastId_example" // string | lanMulticastId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanMulticastTemplateAPI.DeleteLanMulticastTemplate(context.Background(), omadacId, siteTemplateId, multicastId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanMulticastTemplateAPI.DeleteLanMulticastTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLanMulticastTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LanMulticastTemplateAPI.DeleteLanMulticastTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**multicastId** | **string** | lanMulticastId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanMulticastTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridLanMulticastTemplates

> OperationResponse GetGridLanMulticastTemplates(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get lan multicast template list



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanMulticastTemplateAPI.GetGridLanMulticastTemplates(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanMulticastTemplateAPI.GetGridLanMulticastTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridLanMulticastTemplates`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `LanMulticastTemplateAPI.GetGridLanMulticastTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridLanMulticastTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyLanMulticastTemplate

> OperationResponseWithoutResult ModifyLanMulticastTemplate(ctx, omadacId, siteTemplateId, multicastId).LanMulticastVO(lanMulticastVO).Execute()

Modify a lan multicast template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	multicastId := "multicastId_example" // string | lanMulticastId
	lanMulticastVO := *openapiclient.NewLanMulticastVO() // LanMulticastVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanMulticastTemplateAPI.ModifyLanMulticastTemplate(context.Background(), omadacId, siteTemplateId, multicastId).LanMulticastVO(lanMulticastVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanMulticastTemplateAPI.ModifyLanMulticastTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanMulticastTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LanMulticastTemplateAPI.ModifyLanMulticastTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**multicastId** | **string** | lanMulticastId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanMulticastTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanMulticastVO** | [**LanMulticastVO**](LanMulticastVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


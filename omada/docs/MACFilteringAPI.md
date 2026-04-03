# MACFilteringAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMacFiltering**](MACFilteringAPI.md#createmacfiltering) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/mac-filters | Create MAC filtering
[**DeleteMacFiltering**](MACFilteringAPI.md#deletemacfiltering) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/mac-filters/{filterId} | Delete MAC filtering
[**GetGridAllowMacFiltering**](MACFilteringAPI.md#getgridallowmacfiltering) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mac-filters/allow | Get allow MAC filtering list
[**GetGridDenyMacFiltering**](MACFilteringAPI.md#getgriddenymacfiltering) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mac-filters/deny | Get deny MAC filtering list
[**GetMacFilteringGeneralSetting**](MACFilteringAPI.md#getmacfilteringgeneralsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mac-filter | Get MAC filtering general setting
[**ModifyMacFiltering**](MACFilteringAPI.md#modifymacfiltering) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/mac-filters/{filterId} | Modify MAC filtering
[**ModifyMacFilteringGeneralSetting**](MACFilteringAPI.md#modifymacfilteringgeneralsetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/mac-filter | Modify MAC filtering general setting



## CreateMacFiltering

> OperationResponseWithoutResult CreateMacFiltering(ctx, omadacId, siteId).MacFiltering(macFiltering).Execute()

Create MAC filtering



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
	macFiltering := *openapiclient.NewMacFiltering(int32(123), "Name_example", int32(123)) // MacFiltering | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringAPI.CreateMacFiltering(context.Background(), omadacId, siteId).MacFiltering(macFiltering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringAPI.CreateMacFiltering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMacFiltering`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringAPI.CreateMacFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMacFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macFiltering** | [**MacFiltering**](MacFiltering.md) |  | 

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


## DeleteMacFiltering

> OperationResponseWithoutResult DeleteMacFiltering(ctx, omadacId, siteId, filterId).Execute()

Delete MAC filtering



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
	filterId := "filterId_example" // string | MAC filtering ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringAPI.DeleteMacFiltering(context.Background(), omadacId, siteId, filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringAPI.DeleteMacFiltering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMacFiltering`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringAPI.DeleteMacFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**filterId** | **string** | MAC filtering ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMacFilteringRequest struct via the builder pattern


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


## GetGridAllowMacFiltering

> OperationResponseGridVOMacFiltering GetGridAllowMacFiltering(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get allow MAC filtering list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringAPI.GetGridAllowMacFiltering(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringAPI.GetGridAllowMacFiltering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAllowMacFiltering`: OperationResponseGridVOMacFiltering
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringAPI.GetGridAllowMacFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAllowMacFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOMacFiltering**](OperationResponseGridVOMacFiltering.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDenyMacFiltering

> OperationResponseGridVOMacFiltering GetGridDenyMacFiltering(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get deny MAC filtering list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringAPI.GetGridDenyMacFiltering(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringAPI.GetGridDenyMacFiltering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDenyMacFiltering`: OperationResponseGridVOMacFiltering
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringAPI.GetGridDenyMacFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDenyMacFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOMacFiltering**](OperationResponseGridVOMacFiltering.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacFilteringGeneralSetting

> OperationResponseMacFilteringGeneralSetting GetMacFilteringGeneralSetting(ctx, omadacId, siteId).Execute()

Get MAC filtering general setting



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
	resp, r, err := apiClient.MACFilteringAPI.GetMacFilteringGeneralSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringAPI.GetMacFilteringGeneralSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacFilteringGeneralSetting`: OperationResponseMacFilteringGeneralSetting
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringAPI.GetMacFilteringGeneralSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacFilteringGeneralSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMacFilteringGeneralSetting**](OperationResponseMacFilteringGeneralSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMacFiltering

> OperationResponseWithoutResult ModifyMacFiltering(ctx, omadacId, siteId, filterId).MacFiltering(macFiltering).Execute()

Modify MAC filtering



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
	filterId := "filterId_example" // string | MAC filtering ID
	macFiltering := *openapiclient.NewMacFiltering(int32(123), "Name_example", int32(123)) // MacFiltering | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringAPI.ModifyMacFiltering(context.Background(), omadacId, siteId, filterId).MacFiltering(macFiltering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringAPI.ModifyMacFiltering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMacFiltering`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringAPI.ModifyMacFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**filterId** | **string** | MAC filtering ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMacFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **macFiltering** | [**MacFiltering**](MacFiltering.md) |  | 

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


## ModifyMacFilteringGeneralSetting

> OperationResponseWithoutResult ModifyMacFilteringGeneralSetting(ctx, omadacId, siteId).MacFilteringGeneralSetting(macFilteringGeneralSetting).Execute()

Modify MAC filtering general setting



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
	macFilteringGeneralSetting := *openapiclient.NewMacFilteringGeneralSetting(false) // MacFilteringGeneralSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringAPI.ModifyMacFilteringGeneralSetting(context.Background(), omadacId, siteId).MacFilteringGeneralSetting(macFilteringGeneralSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringAPI.ModifyMacFilteringGeneralSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMacFilteringGeneralSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringAPI.ModifyMacFilteringGeneralSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMacFilteringGeneralSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macFilteringGeneralSetting** | [**MacFilteringGeneralSetting**](MacFilteringGeneralSetting.md) |  | 

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


# DisableNatAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDisableNat1**](DisableNatAPI.md#adddisablenat1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wired-networks/disable-nat | Create Disable Nat
[**DeleteDisableNat1**](DisableNatAPI.md#deletedisablenat1) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/wired-networks/disable-nat/{disableNatId} | Delete Disable Nat
[**GetDisableNatGrid1**](DisableNatAPI.md#getdisablenatgrid1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wired-networks/disable-nats | Get Disable Nat Grid
[**ModifyDisableNat**](DisableNatAPI.md#modifydisablenat) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/wired-networks/disable-nat/{disableNatId} | Modify Disable Nat



## AddDisableNat1

> OperationResponseWithoutResult AddDisableNat1(ctx, omadacId, siteId).DisableNat(disableNat).Execute()

Create Disable Nat



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
	disableNat := *openapiclient.NewDisableNat("Interface_example", []string{"LanList_example"}, "Name_example", false) // DisableNat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisableNatAPI.AddDisableNat1(context.Background(), omadacId, siteId).DisableNat(disableNat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisableNatAPI.AddDisableNat1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDisableNat1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DisableNatAPI.AddDisableNat1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDisableNat1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disableNat** | [**DisableNat**](DisableNat.md) |  | 

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


## DeleteDisableNat1

> OperationResponseWithoutResult DeleteDisableNat1(ctx, omadacId, siteId, disableNatId).Execute()

Delete Disable Nat



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
	disableNatId := "disableNatId_example" // string | Disable Nat ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisableNatAPI.DeleteDisableNat1(context.Background(), omadacId, siteId, disableNatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisableNatAPI.DeleteDisableNat1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDisableNat1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DisableNatAPI.DeleteDisableNat1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**disableNatId** | **string** | Disable Nat ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDisableNat1Request struct via the builder pattern


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


## GetDisableNatGrid1

> OperationResponseGridVODisableNatDetailOpenApiVO GetDisableNatGrid1(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Disable Nat Grid



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
	resp, r, err := apiClient.DisableNatAPI.GetDisableNatGrid1(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisableNatAPI.GetDisableNatGrid1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisableNatGrid1`: OperationResponseGridVODisableNatDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DisableNatAPI.GetDisableNatGrid1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisableNatGrid1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVODisableNatDetailOpenApiVO**](OperationResponseGridVODisableNatDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDisableNat

> OperationResponseWithoutResult ModifyDisableNat(ctx, omadacId, siteId, disableNatId).DisableNat(disableNat).Execute()

Modify Disable Nat



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
	disableNatId := "disableNatId_example" // string | Disable Nat ID
	disableNat := *openapiclient.NewDisableNat("Interface_example", []string{"LanList_example"}, "Name_example", false) // DisableNat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisableNatAPI.ModifyDisableNat(context.Background(), omadacId, siteId, disableNatId).DisableNat(disableNat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisableNatAPI.ModifyDisableNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDisableNat`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DisableNatAPI.ModifyDisableNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**disableNatId** | **string** | Disable Nat ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDisableNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **disableNat** | [**DisableNat**](DisableNat.md) |  | 

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


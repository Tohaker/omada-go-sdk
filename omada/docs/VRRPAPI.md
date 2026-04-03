# VRRPAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOswVrrp**](VRRPAPI.md#createoswvrrp) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/osw-vrrp | Create Switch Vrrp
[**DeleteOswVrrp**](VRRPAPI.md#deleteoswvrrp) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/osw-vrrp/{oswVrrpId} | Delete Switch Vrrp
[**GetGridOswVrrp**](VRRPAPI.md#getgridoswvrrp) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/osw-vrrp | Get Switch Vrrp list
[**ModifyOswVrrp**](VRRPAPI.md#modifyoswvrrp) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/osw-vrrp/{oswVrrpId} | Modify Switch Vrrp



## CreateOswVrrp

> OperationResponseResIdOpenApiVO CreateOswVrrp(ctx, omadacId, siteId).CreateOswVrrpRequest(createOswVrrpRequest).Execute()

Create Switch Vrrp



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
	createOswVrrpRequest := *openapiclient.NewCreateOswVrrpRequest() // CreateOswVrrpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VRRPAPI.CreateOswVrrp(context.Background(), omadacId, siteId).CreateOswVrrpRequest(createOswVrrpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VRRPAPI.CreateOswVrrp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOswVrrp`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VRRPAPI.CreateOswVrrp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOswVrrpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOswVrrpRequest** | [**CreateOswVrrpRequest**](CreateOswVrrpRequest.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOswVrrp

> OperationResponseWithoutResult DeleteOswVrrp(ctx, omadacId, siteId, oswVrrpId).UserInfoBriefDTO(userInfoBriefDTO).Execute()

Delete Switch Vrrp



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
	oswVrrpId := "oswVrrpId_example" // string | Osw Vrrp ID
	userInfoBriefDTO := *openapiclient.NewUserInfoBriefDTO() // UserInfoBriefDTO |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VRRPAPI.DeleteOswVrrp(context.Background(), omadacId, siteId, oswVrrpId).UserInfoBriefDTO(userInfoBriefDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VRRPAPI.DeleteOswVrrp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOswVrrp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VRRPAPI.DeleteOswVrrp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**oswVrrpId** | **string** | Osw Vrrp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOswVrrpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userInfoBriefDTO** | [**UserInfoBriefDTO**](UserInfoBriefDTO.md) |  | 

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


## GetGridOswVrrp

> OperationResponseGridVOOswVrrpOpenApiVO GetGridOswVrrp(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Switch Vrrp list



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
	resp, r, err := apiClient.VRRPAPI.GetGridOswVrrp(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VRRPAPI.GetGridOswVrrp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridOswVrrp`: OperationResponseGridVOOswVrrpOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VRRPAPI.GetGridOswVrrp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridOswVrrpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOswVrrpOpenApiVO**](OperationResponseGridVOOswVrrpOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOswVrrp

> OperationResponseWithoutResult ModifyOswVrrp(ctx, omadacId, siteId, oswVrrpId).CreateOswVrrpRequest(createOswVrrpRequest).Execute()

Modify Switch Vrrp



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
	oswVrrpId := "oswVrrpId_example" // string | Osw Vrrp ID
	createOswVrrpRequest := *openapiclient.NewCreateOswVrrpRequest() // CreateOswVrrpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VRRPAPI.ModifyOswVrrp(context.Background(), omadacId, siteId, oswVrrpId).CreateOswVrrpRequest(createOswVrrpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VRRPAPI.ModifyOswVrrp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswVrrp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VRRPAPI.ModifyOswVrrp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**oswVrrpId** | **string** | Osw Vrrp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswVrrpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createOswVrrpRequest** | [**CreateOswVrrpRequest**](CreateOswVrrpRequest.md) |  | 

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


# AccessControlAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessControl**](AccessControlAPI.md#getaccesscontrol) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/access-control | Get access control setting
[**ModifyAccessControl**](AccessControlAPI.md#modifyaccesscontrol) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/access-control | Modify access control setting



## GetAccessControl

> OperationResponsePortalAccessControlOpenApiVO GetAccessControl(ctx, omadacId, siteId).Execute()

Get access control setting



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
	resp, r, err := apiClient.AccessControlAPI.GetAccessControl(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessControlAPI.GetAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessControl`: OperationResponsePortalAccessControlOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AccessControlAPI.GetAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponsePortalAccessControlOpenApiVO**](OperationResponsePortalAccessControlOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAccessControl

> OperationResponsePortalAccessControlOpenApiVO ModifyAccessControl(ctx, omadacId, siteId).PortalAccessControlOpenApiVO(portalAccessControlOpenApiVO).Execute()

Modify access control setting



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
	portalAccessControlOpenApiVO := *openapiclient.NewPortalAccessControlOpenApiVO(false, false) // PortalAccessControlOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessControlAPI.ModifyAccessControl(context.Background(), omadacId, siteId).PortalAccessControlOpenApiVO(portalAccessControlOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessControlAPI.ModifyAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAccessControl`: OperationResponsePortalAccessControlOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AccessControlAPI.ModifyAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAccessControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalAccessControlOpenApiVO** | [**PortalAccessControlOpenApiVO**](PortalAccessControlOpenApiVO.md) |  | 

### Return type

[**OperationResponsePortalAccessControlOpenApiVO**](OperationResponsePortalAccessControlOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


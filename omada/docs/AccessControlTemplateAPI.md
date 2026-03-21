# \AccessControlTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessControlTemplate**](AccessControlTemplateAPI.md#GetAccessControlTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/access-control | Get access control setting template
[**ModifyAccessControlTemplate**](AccessControlTemplateAPI.md#ModifyAccessControlTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/access-control | Modify access control setting template



## GetAccessControlTemplate

> OperationResponsePortalAccessControlOpenApiVO GetAccessControlTemplate(ctx, omadacId, siteTemplateId).Execute()

Get access control setting template



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessControlTemplateAPI.GetAccessControlTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessControlTemplateAPI.GetAccessControlTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessControlTemplate`: OperationResponsePortalAccessControlOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AccessControlTemplateAPI.GetAccessControlTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessControlTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponsePortalAccessControlOpenApiVO**](OperationResponsePortalAccessControlOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAccessControlTemplate

> OperationResponsePortalAccessControlOpenApiVO ModifyAccessControlTemplate(ctx, omadacId, siteTemplateId).PortalAccessControlOpenApiVO(portalAccessControlOpenApiVO).Execute()

Modify access control setting template



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
	portalAccessControlOpenApiVO := *openapiclient.NewPortalAccessControlOpenApiVO(false, false) // PortalAccessControlOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessControlTemplateAPI.ModifyAccessControlTemplate(context.Background(), omadacId, siteTemplateId).PortalAccessControlOpenApiVO(portalAccessControlOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessControlTemplateAPI.ModifyAccessControlTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAccessControlTemplate`: OperationResponsePortalAccessControlOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AccessControlTemplateAPI.ModifyAccessControlTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAccessControlTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalAccessControlOpenApiVO** | [**PortalAccessControlOpenApiVO**](PortalAccessControlOpenApiVO.md) |  | 

### Return type

[**OperationResponsePortalAccessControlOpenApiVO**](OperationResponsePortalAccessControlOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# UPnPTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUpnpSettingTemplate**](UPnPTemplateAPI.md#getupnpsettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/upnp | Get UPnP template setting
[**UpdateUpnpSettingTemplate**](UPnPTemplateAPI.md#updateupnpsettingtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/upnp | Modify UPnP template setting



## GetUpnpSettingTemplate

> OperationResponseUpnpSettingOpenApiVO GetUpnpSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get UPnP template setting



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UPnPTemplateAPI.GetUpnpSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UPnPTemplateAPI.GetUpnpSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpnpSettingTemplate`: OperationResponseUpnpSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `UPnPTemplateAPI.GetUpnpSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpnpSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUpnpSettingOpenApiVO**](OperationResponseUpnpSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUpnpSettingTemplate

> OperationResponseWithoutResult UpdateUpnpSettingTemplate(ctx, omadacId, siteTemplateId).UpnpSettingOpenApiVO(upnpSettingOpenApiVO).Execute()

Modify UPnP template setting



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	upnpSettingOpenApiVO := *openapiclient.NewUpnpSettingOpenApiVO(false) // UpnpSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UPnPTemplateAPI.UpdateUpnpSettingTemplate(context.Background(), omadacId, siteTemplateId).UpnpSettingOpenApiVO(upnpSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UPnPTemplateAPI.UpdateUpnpSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUpnpSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `UPnPTemplateAPI.UpdateUpnpSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUpnpSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upnpSettingOpenApiVO** | [**UpnpSettingOpenApiVO**](UpnpSettingOpenApiVO.md) |  | 

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


# UPnPAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUpnpSetting**](UPnPAPI.md#getupnpsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/upnp | Get UPnP setting
[**UpdateUpnpSetting**](UPnPAPI.md#updateupnpsetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/upnp | Modify UPnP setting



## GetUpnpSetting

> OperationResponseUpnpSettingOpenApiVO GetUpnpSetting(ctx, omadacId, siteId).Execute()

Get UPnP setting



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
	resp, r, err := apiClient.UPnPAPI.GetUpnpSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UPnPAPI.GetUpnpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpnpSetting`: OperationResponseUpnpSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `UPnPAPI.GetUpnpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpnpSettingRequest struct via the builder pattern


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


## UpdateUpnpSetting

> OperationResponseWithoutResult UpdateUpnpSetting(ctx, omadacId, siteId).UpnpSettingOpenApiVO(upnpSettingOpenApiVO).Execute()

Modify UPnP setting



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
	upnpSettingOpenApiVO := *openapiclient.NewUpnpSettingOpenApiVO(false) // UpnpSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UPnPAPI.UpdateUpnpSetting(context.Background(), omadacId, siteId).UpnpSettingOpenApiVO(upnpSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UPnPAPI.UpdateUpnpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUpnpSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `UPnPAPI.UpdateUpnpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUpnpSettingRequest struct via the builder pattern


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


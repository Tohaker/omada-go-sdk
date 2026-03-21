# \EoGRETunnelAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGreTunnelConfig**](EoGRETunnelAPI.md#GetGreTunnelConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/gre-tunnel | Get EoGRE tunnel config
[**ModifyGreTunnelConfig**](EoGRETunnelAPI.md#ModifyGreTunnelConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/gre-tunnel | Modify EoGRE tunnel config



## GetGreTunnelConfig

> OperationResponseEoGreTunnelSettingOpenApiVO GetGreTunnelConfig(ctx, omadacId, siteId).Execute()

Get EoGRE tunnel config



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
	resp, r, err := apiClient.EoGRETunnelAPI.GetGreTunnelConfig(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EoGRETunnelAPI.GetGreTunnelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGreTunnelConfig`: OperationResponseEoGreTunnelSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `EoGRETunnelAPI.GetGreTunnelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGreTunnelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseEoGreTunnelSettingOpenApiVO**](OperationResponseEoGreTunnelSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGreTunnelConfig

> OperationResponseWithoutResult ModifyGreTunnelConfig(ctx, omadacId, siteId).UpdateEoGreTunnelSettingOpenApiVO(updateEoGreTunnelSettingOpenApiVO).Execute()

Modify EoGRE tunnel config



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
	updateEoGreTunnelSettingOpenApiVO := *openapiclient.NewUpdateEoGreTunnelSettingOpenApiVO(false, int32(123)) // UpdateEoGreTunnelSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EoGRETunnelAPI.ModifyGreTunnelConfig(context.Background(), omadacId, siteId).UpdateEoGreTunnelSettingOpenApiVO(updateEoGreTunnelSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EoGRETunnelAPI.ModifyGreTunnelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGreTunnelConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `EoGRETunnelAPI.ModifyGreTunnelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyGreTunnelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateEoGreTunnelSettingOpenApiVO** | [**UpdateEoGreTunnelSettingOpenApiVO**](UpdateEoGreTunnelSettingOpenApiVO.md) |  | 

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


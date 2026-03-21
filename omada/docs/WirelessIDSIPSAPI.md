# \WirelessIDSIPSAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWidsConfig**](WirelessIDSIPSAPI.md#GetWidsConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/network-security/wireless-ids | Get wireless IDS config
[**GetWipsConfig**](WirelessIDSIPSAPI.md#GetWipsConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/network-security/wireless-ips | Get wireless IPS config
[**ModifyWidsConfig**](WirelessIDSIPSAPI.md#ModifyWidsConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/network-security/wireless-ids | Modify wireless IDS config
[**ModifyWipsConfig**](WirelessIDSIPSAPI.md#ModifyWipsConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/network-security/wireless-ips | Modify wireless IPS config



## GetWidsConfig

> OperationResponseWidsConfigOpenApiVO GetWidsConfig(ctx, omadacId, siteId).Execute()

Get wireless IDS config



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
	resp, r, err := apiClient.WirelessIDSIPSAPI.GetWidsConfig(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessIDSIPSAPI.GetWidsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWidsConfig`: OperationResponseWidsConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessIDSIPSAPI.GetWidsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWidsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWidsConfigOpenApiVO**](OperationResponseWidsConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWipsConfig

> OperationResponseWipsConfigOpenApiVO GetWipsConfig(ctx, omadacId, siteId).Execute()

Get wireless IPS config



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
	resp, r, err := apiClient.WirelessIDSIPSAPI.GetWipsConfig(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessIDSIPSAPI.GetWipsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWipsConfig`: OperationResponseWipsConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessIDSIPSAPI.GetWipsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWipsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWipsConfigOpenApiVO**](OperationResponseWipsConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWidsConfig

> OperationResponseWithoutResult ModifyWidsConfig(ctx, omadacId, siteId).UpdateWidsConfigOpenApiVO(updateWidsConfigOpenApiVO).Execute()

Modify wireless IDS config



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
	updateWidsConfigOpenApiVO := *openapiclient.NewUpdateWidsConfigOpenApiVO(int32(123), false) // UpdateWidsConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessIDSIPSAPI.ModifyWidsConfig(context.Background(), omadacId, siteId).UpdateWidsConfigOpenApiVO(updateWidsConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessIDSIPSAPI.ModifyWidsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWidsConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessIDSIPSAPI.ModifyWidsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWidsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWidsConfigOpenApiVO** | [**UpdateWidsConfigOpenApiVO**](UpdateWidsConfigOpenApiVO.md) |  | 

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


## ModifyWipsConfig

> OperationResponseWithoutResult ModifyWipsConfig(ctx, omadacId, siteId).UpdateWipsConfigOpenApiVO(updateWipsConfigOpenApiVO).Execute()

Modify wireless IPS config



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
	updateWipsConfigOpenApiVO := *openapiclient.NewUpdateWipsConfigOpenApiVO(false, false, false) // UpdateWipsConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessIDSIPSAPI.ModifyWipsConfig(context.Background(), omadacId, siteId).UpdateWipsConfigOpenApiVO(updateWipsConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessIDSIPSAPI.ModifyWipsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWipsConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessIDSIPSAPI.ModifyWipsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWipsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWipsConfigOpenApiVO** | [**UpdateWipsConfigOpenApiVO**](UpdateWipsConfigOpenApiVO.md) |  | 

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


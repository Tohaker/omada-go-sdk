# \WirelessIDSIPSTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWidsConfigTemplate**](WirelessIDSIPSTemplateAPI.md#GetWidsConfigTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/wireless-ids | Get wireless IDS config template
[**GetWipsConfigTemplate**](WirelessIDSIPSTemplateAPI.md#GetWipsConfigTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/wireless-ips | Get wireless IPS config template
[**ModifyWidsConfigTemplate**](WirelessIDSIPSTemplateAPI.md#ModifyWidsConfigTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/wireless-ids | Modify wireless IDS config template
[**ModifyWipsConfigTemplate**](WirelessIDSIPSTemplateAPI.md#ModifyWipsConfigTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/wireless-ips | Modify wireless IPS config template



## GetWidsConfigTemplate

> OperationResponseWidsConfigOpenApiVO GetWidsConfigTemplate(ctx, omadacId, siteTemplateId).Execute()

Get wireless IDS config template



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
	resp, r, err := apiClient.WirelessIDSIPSTemplateAPI.GetWidsConfigTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessIDSIPSTemplateAPI.GetWidsConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWidsConfigTemplate`: OperationResponseWidsConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessIDSIPSTemplateAPI.GetWidsConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWidsConfigTemplateRequest struct via the builder pattern


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


## GetWipsConfigTemplate

> OperationResponseWipsConfigOpenApiVO GetWipsConfigTemplate(ctx, omadacId, siteTemplateId).Execute()

Get wireless IPS config template



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
	resp, r, err := apiClient.WirelessIDSIPSTemplateAPI.GetWipsConfigTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessIDSIPSTemplateAPI.GetWipsConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWipsConfigTemplate`: OperationResponseWipsConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessIDSIPSTemplateAPI.GetWipsConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWipsConfigTemplateRequest struct via the builder pattern


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


## ModifyWidsConfigTemplate

> OperationResponseWithoutResult ModifyWidsConfigTemplate(ctx, omadacId, siteTemplateId).UpdateWidsConfigOpenApiVO(updateWidsConfigOpenApiVO).Execute()

Modify wireless IDS config template



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
	updateWidsConfigOpenApiVO := *openapiclient.NewUpdateWidsConfigOpenApiVO(int32(123), false) // UpdateWidsConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessIDSIPSTemplateAPI.ModifyWidsConfigTemplate(context.Background(), omadacId, siteTemplateId).UpdateWidsConfigOpenApiVO(updateWidsConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessIDSIPSTemplateAPI.ModifyWidsConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWidsConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessIDSIPSTemplateAPI.ModifyWidsConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWidsConfigTemplateRequest struct via the builder pattern


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


## ModifyWipsConfigTemplate

> OperationResponseWithoutResult ModifyWipsConfigTemplate(ctx, omadacId, siteTemplateId).UpdateWipsConfigOpenApiVO(updateWipsConfigOpenApiVO).Execute()

Modify wireless IPS config template



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
	updateWipsConfigOpenApiVO := *openapiclient.NewUpdateWipsConfigOpenApiVO(false, false, false) // UpdateWipsConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessIDSIPSTemplateAPI.ModifyWipsConfigTemplate(context.Background(), omadacId, siteTemplateId).UpdateWipsConfigOpenApiVO(updateWipsConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessIDSIPSTemplateAPI.ModifyWipsConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWipsConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessIDSIPSTemplateAPI.ModifyWipsConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWipsConfigTemplateRequest struct via the builder pattern


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


# \FirewallTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFirewallSettingTemplate**](FirewallTemplateAPI.md#GetFirewallSettingTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/firewall | Get firewall template setting
[**ModifyFirewallSettingTemplate**](FirewallTemplateAPI.md#ModifyFirewallSettingTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/firewall | Modify firewall template setting
[**ResetFirewallSettingTemplate**](FirewallTemplateAPI.md#ResetFirewallSettingTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/firewall/reset | Reset firewall template setting



## GetFirewallSettingTemplate

> OperationResponseFirewallSetting GetFirewallSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get firewall template setting



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
	resp, r, err := apiClient.FirewallTemplateAPI.GetFirewallSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallTemplateAPI.GetFirewallSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallSettingTemplate`: OperationResponseFirewallSetting
	fmt.Fprintf(os.Stdout, "Response from `FirewallTemplateAPI.GetFirewallSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseFirewallSetting**](OperationResponseFirewallSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyFirewallSettingTemplate

> OperationResponseWithoutResult ModifyFirewallSettingTemplate(ctx, omadacId, siteTemplateId).FirewallSetting(firewallSetting).Execute()

Modify firewall template setting



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
	firewallSetting := *openapiclient.NewFirewallSetting(false, int32(123), int32(123), false, false, false, int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)) // FirewallSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallTemplateAPI.ModifyFirewallSettingTemplate(context.Background(), omadacId, siteTemplateId).FirewallSetting(firewallSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallTemplateAPI.ModifyFirewallSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyFirewallSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirewallTemplateAPI.ModifyFirewallSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyFirewallSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firewallSetting** | [**FirewallSetting**](FirewallSetting.md) |  | 

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


## ResetFirewallSettingTemplate

> OperationResponseWithoutResult ResetFirewallSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Reset firewall template setting



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
	resp, r, err := apiClient.FirewallTemplateAPI.ResetFirewallSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallTemplateAPI.ResetFirewallSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetFirewallSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirewallTemplateAPI.ResetFirewallSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetFirewallSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


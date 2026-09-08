# SNMPAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSnmpIncompatibleDevices**](SNMPAPI.md#getsnmpincompatibledevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/snmp/incompatible-devices | Get the list of site devices that do not support SNMP v3 enhanced configuration
[**GetSnmpSetting**](SNMPAPI.md#getsnmpsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/snmp | Get SNMP setting
[**ModifySnmpSetting**](SNMPAPI.md#modifysnmpsetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/snmp | Modify SNMP setting



## GetSnmpIncompatibleDevices

> OperationResponseGridVODeviceVO GetSnmpIncompatibleDevices(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the list of site devices that do not support SNMP v3 enhanced configuration



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
	resp, r, err := apiClient.SNMPAPI.GetSnmpIncompatibleDevices(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SNMPAPI.GetSnmpIncompatibleDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnmpIncompatibleDevices`: OperationResponseGridVODeviceVO
	fmt.Fprintf(os.Stdout, "Response from `SNMPAPI.GetSnmpIncompatibleDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnmpIncompatibleDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVODeviceVO**](OperationResponseGridVODeviceVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnmpSetting

> OperationResponseSnmpSettingOpenApiVO GetSnmpSetting(ctx, omadacId, siteId).Execute()

Get SNMP setting



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
	resp, r, err := apiClient.SNMPAPI.GetSnmpSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SNMPAPI.GetSnmpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnmpSetting`: OperationResponseSnmpSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SNMPAPI.GetSnmpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnmpSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSnmpSettingOpenApiVO**](OperationResponseSnmpSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySnmpSetting

> OperationResponseWithoutResult ModifySnmpSetting(ctx, omadacId, siteId).SnmpSettingOpenApiVO(snmpSettingOpenApiVO).Execute()

Modify SNMP setting



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
	snmpSettingOpenApiVO := *openapiclient.NewSnmpSettingOpenApiVO(false, false) // SnmpSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SNMPAPI.ModifySnmpSetting(context.Background(), omadacId, siteId).SnmpSettingOpenApiVO(snmpSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SNMPAPI.ModifySnmpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySnmpSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SNMPAPI.ModifySnmpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySnmpSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **snmpSettingOpenApiVO** | [**SnmpSettingOpenApiVO**](SnmpSettingOpenApiVO.md) |  | 

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


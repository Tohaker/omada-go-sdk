# DeviceManagementSettingAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceAccessMgmtSetting**](DeviceManagementSettingAPI.md#getdeviceaccessmgmtsetting) | **Get** /openapi/v1/{omadacId}/device-mgmt | Get device access management setting
[**ModifyDeviceAccessMgmtSetting**](DeviceManagementSettingAPI.md#modifydeviceaccessmgmtsetting) | **Patch** /openapi/v1/{omadacId}/device-mgmt | Modify device access management setting



## GetDeviceAccessMgmtSetting

> OperationResponseDeviceAccessManagementSetting GetDeviceAccessMgmtSetting(ctx, omadacId).Execute()

Get device access management setting



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceManagementSettingAPI.GetDeviceAccessMgmtSetting(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementSettingAPI.GetDeviceAccessMgmtSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAccessMgmtSetting`: OperationResponseDeviceAccessManagementSetting
	fmt.Fprintf(os.Stdout, "Response from `DeviceManagementSettingAPI.GetDeviceAccessMgmtSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAccessMgmtSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseDeviceAccessManagementSetting**](OperationResponseDeviceAccessManagementSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDeviceAccessMgmtSetting

> OperationResponseWithoutResult ModifyDeviceAccessMgmtSetting(ctx, omadacId).DeviceAccessManagementSetting(deviceAccessManagementSetting).Execute()

Modify device access management setting



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
	deviceAccessManagementSetting := *openapiclient.NewDeviceAccessManagementSetting(false, false, false) // DeviceAccessManagementSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceManagementSettingAPI.ModifyDeviceAccessMgmtSetting(context.Background(), omadacId).DeviceAccessManagementSetting(deviceAccessManagementSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementSettingAPI.ModifyDeviceAccessMgmtSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDeviceAccessMgmtSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceManagementSettingAPI.ModifyDeviceAccessMgmtSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDeviceAccessMgmtSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceAccessManagementSetting** | [**DeviceAccessManagementSetting**](DeviceAccessManagementSetting.md) |  | 

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


# IntelliRecoverDeviceAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMonitorDevices**](IntelliRecoverDeviceAPI.md#addmonitordevices) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/monitor-device/add | Add device into intelli recover device list
[**DeleteMonitorDevices**](IntelliRecoverDeviceAPI.md#deletemonitordevices) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/monitor-device/delete | Delete the intelli recover device
[**GetEligibleDeviceList**](IntelliRecoverDeviceAPI.md#geteligibledevicelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/monitor-device/eligible-device-grids | Obtain the list of devices that can be monitored
[**GetGridMonitorDevice**](IntelliRecoverDeviceAPI.md#getgridmonitordevice) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/monitor-device/grids | Get the intelli recover device list
[**GetIntelliRecoverSetting**](IntelliRecoverDeviceAPI.md#getintellirecoversetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/intelli-recover/auto-setting/config | Get intelli recover setting
[**ModifyIntelliRecoverSetting**](IntelliRecoverDeviceAPI.md#modifyintellirecoversetting) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/intelli-recover/auto-setting/modify | Modify intelli recover setting
[**RebootDeviceUplinkPoe**](IntelliRecoverDeviceAPI.md#rebootdeviceuplinkpoe) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/monitor-device/reboot-uplink-poe | Reboot the device uplinkDevice poe
[**VerifyMonitorDevice**](IntelliRecoverDeviceAPI.md#verifymonitordevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/monitor-device/verify | Verify if the device is allowed to be monitored



## AddMonitorDevices

> OperationResponseWithoutResult AddMonitorDevices(ctx, omadacId, siteId).AddMonitorDeviceList(addMonitorDeviceList).Execute()

Add device into intelli recover device list



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
	addMonitorDeviceList := *openapiclient.NewAddMonitorDeviceList() // AddMonitorDeviceList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverDeviceAPI.AddMonitorDevices(context.Background(), omadacId, siteId).AddMonitorDeviceList(addMonitorDeviceList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverDeviceAPI.AddMonitorDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMonitorDevices`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverDeviceAPI.AddMonitorDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMonitorDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addMonitorDeviceList** | [**AddMonitorDeviceList**](AddMonitorDeviceList.md) |  | 

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


## DeleteMonitorDevices

> OperationResponseWithoutResult DeleteMonitorDevices(ctx, omadacId, siteId).DeleteIntelliRecoverDevice(deleteIntelliRecoverDevice).Execute()

Delete the intelli recover device



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
	deleteIntelliRecoverDevice := *openapiclient.NewDeleteIntelliRecoverDevice() // DeleteIntelliRecoverDevice | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverDeviceAPI.DeleteMonitorDevices(context.Background(), omadacId, siteId).DeleteIntelliRecoverDevice(deleteIntelliRecoverDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverDeviceAPI.DeleteMonitorDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMonitorDevices`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverDeviceAPI.DeleteMonitorDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteIntelliRecoverDevice** | [**DeleteIntelliRecoverDevice**](DeleteIntelliRecoverDevice.md) |  | 

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


## GetEligibleDeviceList

> OperationResponseGridVOMonitorDevice GetEligibleDeviceList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Obtain the list of devices that can be monitored



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverDeviceAPI.GetEligibleDeviceList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverDeviceAPI.GetEligibleDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEligibleDeviceList`: OperationResponseGridVOMonitorDevice
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverDeviceAPI.GetEligibleDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEligibleDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOMonitorDevice**](OperationResponseGridVOMonitorDevice.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridMonitorDevice

> OperationResponseGridVOMonitorDevice GetGridMonitorDevice(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the intelli recover device list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverDeviceAPI.GetGridMonitorDevice(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverDeviceAPI.GetGridMonitorDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridMonitorDevice`: OperationResponseGridVOMonitorDevice
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverDeviceAPI.GetGridMonitorDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridMonitorDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOMonitorDevice**](OperationResponseGridVOMonitorDevice.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntelliRecoverSetting

> OperationResponseIntelliRecoverSetting GetIntelliRecoverSetting(ctx, omadacId, siteId).Execute()

Get intelli recover setting



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
	resp, r, err := apiClient.IntelliRecoverDeviceAPI.GetIntelliRecoverSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverDeviceAPI.GetIntelliRecoverSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntelliRecoverSetting`: OperationResponseIntelliRecoverSetting
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverDeviceAPI.GetIntelliRecoverSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelliRecoverSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIntelliRecoverSetting**](OperationResponseIntelliRecoverSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIntelliRecoverSetting

> OperationResponseWithoutResult ModifyIntelliRecoverSetting(ctx, omadacId, siteId).IntelliRecoverSetting(intelliRecoverSetting).Execute()

Modify intelli recover setting



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
	intelliRecoverSetting := *openapiclient.NewIntelliRecoverSetting() // IntelliRecoverSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverDeviceAPI.ModifyIntelliRecoverSetting(context.Background(), omadacId, siteId).IntelliRecoverSetting(intelliRecoverSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverDeviceAPI.ModifyIntelliRecoverSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIntelliRecoverSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverDeviceAPI.ModifyIntelliRecoverSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIntelliRecoverSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **intelliRecoverSetting** | [**IntelliRecoverSetting**](IntelliRecoverSetting.md) |  | 

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


## RebootDeviceUplinkPoe

> OperationResponseWithoutResult RebootDeviceUplinkPoe(ctx, omadacId, siteId).MonitorDevice(monitorDevice).Execute()

Reboot the device uplinkDevice poe



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
	monitorDevice := *openapiclient.NewMonitorDevice("Mac_example") // MonitorDevice | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverDeviceAPI.RebootDeviceUplinkPoe(context.Background(), omadacId, siteId).MonitorDevice(monitorDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverDeviceAPI.RebootDeviceUplinkPoe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootDeviceUplinkPoe`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverDeviceAPI.RebootDeviceUplinkPoe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootDeviceUplinkPoeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **monitorDevice** | [**MonitorDevice**](MonitorDevice.md) |  | 

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


## VerifyMonitorDevice

> OperationResponseWithoutResult VerifyMonitorDevice(ctx, omadacId, siteId).VerifyDevice(verifyDevice).Execute()

Verify if the device is allowed to be monitored



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
	verifyDevice := *openapiclient.NewVerifyDevice() // VerifyDevice | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverDeviceAPI.VerifyMonitorDevice(context.Background(), omadacId, siteId).VerifyDevice(verifyDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverDeviceAPI.VerifyMonitorDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyMonitorDevice`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverDeviceAPI.VerifyMonitorDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMonitorDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **verifyDevice** | [**VerifyDevice**](VerifyDevice.md) |  | 

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


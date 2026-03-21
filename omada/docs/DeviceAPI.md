# \DeviceAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveDevice**](DeviceAPI.md#ActiveDevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/active | Batch active devices
[**ActiveDeviceBySn**](DeviceAPI.md#ActiveDeviceBySn) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/sn-active | Batch active devices by SN
[**ActiveDeviceWithMultiSite**](DeviceAPI.md#ActiveDeviceWithMultiSite) | **Post** /openapi/v1/{omadacId}/multi-devices/active | Batch active devices with multi site
[**AddDevices**](DeviceAPI.md#AddDevices) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/multi-add | Batch add devices
[**AddDevicesByDeviceKey**](DeviceAPI.md#AddDevicesByDeviceKey) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/devicekey-add | Batch add devices by DEVICE KEY.
[**AddDevicesBySn**](DeviceAPI.md#AddDevicesBySn) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/sn-add | Batch add devices by SN
[**AddDevicesBySnWithMultiSite**](DeviceAPI.md#AddDevicesBySnWithMultiSite) | **Post** /openapi/v1/{omadacId}/multi-devices/sn-add | Batch add devices by SN with multi site
[**AddDevicesForMsp**](DeviceAPI.md#AddDevicesForMsp) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/devices/multi-add | Batch add devices in MSP view
[**AddTag**](DeviceAPI.md#AddTag) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/tag | Create new tag
[**AdoptDevice**](DeviceAPI.md#AdoptDevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/start-adopt | Start adopt device
[**BatchAdopt**](DeviceAPI.md#BatchAdopt) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/devices/batch-adopt | batch adopt device
[**CancelRollingUpgrade**](DeviceAPI.md#CancelRollingUpgrade) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/rolling-upgrade-tasks/{taskId} | End the rolling upgrade task
[**CopyDeviceConfiguration**](DeviceAPI.md#CopyDeviceConfiguration) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/copy | Copy configuration
[**DeleteTag**](DeviceAPI.md#DeleteTag) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/devices/tag | Delete an existing tag
[**DownloadDeviceInfo**](DeviceAPI.md#DownloadDeviceInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cmd/devices/{deviceMac}/downloadDeviceInfo/{type} | Download device information.
[**ExistUnSupportRadSecDevice**](DeviceAPI.md#ExistUnSupportRadSecDevice) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/un-support-radsec | Check for unsupported RadSec devices.
[**ExportAddFailDevices**](DeviceAPI.md#ExportAddFailDevices) | **Get** /openapi/v1/{omadacId}/devices/add/{operateId}/export | Export the device file that failed to add in GLOBAL view.
[**ExportGlobalDeviceList**](DeviceAPI.md#ExportGlobalDeviceList) | **Post** /openapi/v1/{omadacId}/devices/export | Export device list in GLOBAL view
[**ExportMspAddFailDevices**](DeviceAPI.md#ExportMspAddFailDevices) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/devices/add/{operateId}/export | Export the device file that failed to add in MSP view.
[**ExportPreConfigDevices**](DeviceAPI.md#ExportPreConfigDevices) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/migrate/exportPreDevice | Export preconfigured devices of site.
[**ExportPreConfigDevicesGlobal**](DeviceAPI.md#ExportPreConfigDevicesGlobal) | **Post** /openapi/v1/{omadacId}/migrate/exportPreDevice | Export preconfigured devices in GLOBAL view.
[**ExportSiteDeviceList**](DeviceAPI.md#ExportSiteDeviceList) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/export | Export device list of site
[**ForceProvisionDevice**](DeviceAPI.md#ForceProvisionDevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/force-provision | Force provision device
[**ForgetDevice**](DeviceAPI.md#ForgetDevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/forget | Forget device
[**ForgetDeviceForMsp**](DeviceAPI.md#ForgetDeviceForMsp) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/devices/{deviceMac}/forget | Forget one device in MSP view.
[**GetAddDevicesStatus**](DeviceAPI.md#GetAddDevicesStatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/add/{operateId}/status | Batch get added device status
[**GetAddDevicesStatusMultiSite**](DeviceAPI.md#GetAddDevicesStatusMultiSite) | **Get** /openapi/v1/{omadacId}/multi-devices/add/{operateId}/status | Batch get added device status with multi site
[**GetAdoptAbleDevicesForGlobalBatchAdopt**](DeviceAPI.md#GetAdoptAbleDevicesForGlobalBatchAdopt) | **Get** /openapi/v1/{omadacId}/devices/batch/info | Batch get info of adopted device in global view.
[**GetAdoptAbleDevicesForMspBatchAdopt**](DeviceAPI.md#GetAdoptAbleDevicesForMspBatchAdopt) | **Get** /openapi/v1/msp/{mspId}/devices/batch/info | Batch get info of adopted device in MSP view.
[**GetAdoptTip**](DeviceAPI.md#GetAdoptTip) | **Get** /openapi/v1/{omadacId}/adopt-tip | Get adopt tip
[**GetAllDeviceBySite**](DeviceAPI.md#GetAllDeviceBySite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/all | Get device list info.
[**GetAutoSelectDevices**](DeviceAPI.md#GetAutoSelectDevices) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/quick-action/network-wizard/auto-select-devices | Get auto select devices in quick-config page.
[**GetCheckFirmwareTaskResult**](DeviceAPI.md#GetCheckFirmwareTaskResult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/all-devices/check-latest-firmware-tasks/{taskId} | Get the result of the given check latest firmware task
[**GetDeviceAdoptResult**](DeviceAPI.md#GetDeviceAdoptResult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/adopt-result | Get device adopt result
[**GetDeviceList**](DeviceAPI.md#GetDeviceList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices | Get site device list
[**GetDeviceRememberMe**](DeviceAPI.md#GetDeviceRememberMe) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/remember | Get device remember Config
[**GetDeviceWhiteList**](DeviceAPI.md#GetDeviceWhiteList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/device-white-list | Get the whitelist list of devices.
[**GetFirmwareInfo**](DeviceAPI.md#GetFirmwareInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/latest-firmware-info | Get the latest firmware info of the device
[**GetForgetProcess**](DeviceAPI.md#GetForgetProcess) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/forget/{forgetId}/status | Get batch forget process
[**GetGlobalKnownDeviceList**](DeviceAPI.md#GetGlobalKnownDeviceList) | **Get** /openapi/v1/{omadacId}/devices/known-devices | Get global known device list
[**GetGlobalUnknownDeviceList**](DeviceAPI.md#GetGlobalUnknownDeviceList) | **Get** /openapi/v1/{omadacId}/devices/unknown-devices | Get global unknown device list
[**GetGridAdoptedBridgeDevicesBySite**](DeviceAPI.md#GetGridAdoptedBridgeDevicesBySite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/grid/bridge-devices/adopted | Get Bridge group grouped devices
[**GetGridAdoptedDevicesStatByGlobal**](DeviceAPI.md#GetGridAdoptedDevicesStatByGlobal) | **Get** /openapi/v1/{omadacId}/devices/stat | Query the statistics for the list of global adopted devices.
[**GetGridPendingBridgeDevicesBySite**](DeviceAPI.md#GetGridPendingBridgeDevicesBySite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/grid/bridge-devices/pending | Get Bridge group ungrouped devices
[**GetGridPendingDevicesBySite**](DeviceAPI.md#GetGridPendingDevicesBySite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/grid/devices/pending | Get adoptable device list of target site
[**GetManuallyUpgradeRes**](DeviceAPI.md#GetManuallyUpgradeRes) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/manually-upgrade-tasks/{taskId} | Get the result of the given manually upgrade task
[**GetMoveSiteProcess**](DeviceAPI.md#GetMoveSiteProcess) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/move/{moveSiteId}/status | Get batch move site process
[**GetOnlineTimeline**](DeviceAPI.md#GetOnlineTimeline) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/timeline | Get device online timeline
[**GetOnlineUpgradeRes**](DeviceAPI.md#GetOnlineUpgradeRes) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/online-upgrade-result | Get online upgrade result
[**GetRollingUpgradeRes**](DeviceAPI.md#GetRollingUpgradeRes) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/rolling-upgrade-tasks/{taskId} | Get the result of the given rolling upgrade task
[**GetSingleForgetProcess**](DeviceAPI.md#GetSingleForgetProcess) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/forget/{forgetId}/status | Get forget process of device
[**GetSingleMoveSiteProcess**](DeviceAPI.md#GetSingleMoveSiteProcess) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/move/{moveSiteId}/status | Get move site process of device
[**GetTags**](DeviceAPI.md#GetTags) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/devices/tag | Get tag list
[**GetUplinkDeviceInfo**](DeviceAPI.md#GetUplinkDeviceInfo) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/uplink-info | Query uplink information for specified device MAC addresses under the site.
[**LocateDevice**](DeviceAPI.md#LocateDevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/locate | Locate device
[**LocatePorts**](DeviceAPI.md#LocatePorts) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/locate/switch-ports | Locate multiple ports of multiple switches
[**ManuallyUpgrade**](DeviceAPI.md#ManuallyUpgrade) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/start-manually-upgrade | Start manually upgrade
[**ModifyDeviceRememberMe**](DeviceAPI.md#ModifyDeviceRememberMe) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/remember | Modify device remember Config
[**ModifyTag**](DeviceAPI.md#ModifyTag) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/devices/tag | Modify an existing tag
[**MoveSite**](DeviceAPI.md#MoveSite) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/site-move | Move site
[**MoveSiteForAps**](DeviceAPI.md#MoveSiteForAps) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/site-move | Move aps to another site
[**OnlineCheckUpgrade**](DeviceAPI.md#OnlineCheckUpgrade) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/all-devices/start-check-latest-firmware | Start check latest firmware for all devices
[**OnlineRollingUpgrade**](DeviceAPI.md#OnlineRollingUpgrade) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/start-rolling-upgrade | Start batch rolling upgrade
[**OnlineRollingUpgradeByQuery**](DeviceAPI.md#OnlineRollingUpgradeByQuery) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/devices/onlineRollingUpgrade | Start batch rolling upgrade By Query
[**OnlineUpgrade**](DeviceAPI.md#OnlineUpgrade) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/start-online-upgrade | Start online upgrade
[**RebootDevice**](DeviceAPI.md#RebootDevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/reboot | Reboot device
[**RetryAddDevice**](DeviceAPI.md#RetryAddDevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/{deviceMac}/add/retry | Retry add device
[**RetryAddDeviceByMsp**](DeviceAPI.md#RetryAddDeviceByMsp) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/devices/{deviceMac}/add/retry | retry add device in msp view
[**SearchGlobalDevice**](DeviceAPI.md#SearchGlobalDevice) | **Get** /openapi/v1/{omadacId}/devices | Global search for devices returns the devices under the sites that have permissions.
[**UnbindActiveDeviceBySn**](DeviceAPI.md#UnbindActiveDeviceBySn) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/sn-unbind-active | Change the license from the old device to the new device
[**UnbindDevice**](DeviceAPI.md#UnbindDevice) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/unbind | Batch unbind devices
[**UnbindDeviceBySn**](DeviceAPI.md#UnbindDeviceBySn) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/sn-unbind | Batch unbind devices by SN
[**UpdateDeviceTag**](DeviceAPI.md#UpdateDeviceTag) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/multi-devices/tag | Set device tag for given devices
[**UploadUpgradeFile**](DeviceAPI.md#UploadUpgradeFile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/devices/upgrade-file | Upload device firmware



## ActiveDevice

> []ActiveDeviceRespVO ActiveDevice(ctx, omadacId, siteId).ActiveDeviceOpenApiVO(activeDeviceOpenApiVO).Execute()

Batch active devices



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
	activeDeviceOpenApiVO := *openapiclient.NewActiveDeviceOpenApiVO([]openapiclient.ActivePairOpenApiDTO{*openapiclient.NewActivePairOpenApiDTO("DeviceMac_example", "LicenseType_example")}, "Category_example") // ActiveDeviceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ActiveDevice(context.Background(), omadacId, siteId).ActiveDeviceOpenApiVO(activeDeviceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ActiveDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActiveDevice`: []ActiveDeviceRespVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ActiveDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **activeDeviceOpenApiVO** | [**ActiveDeviceOpenApiVO**](ActiveDeviceOpenApiVO.md) |  | 

### Return type

[**[]ActiveDeviceRespVO**](ActiveDeviceRespVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActiveDeviceBySn

> []ActiveDeviceSnRespVO ActiveDeviceBySn(ctx, omadacId, siteId).ActiveDeviceSnOpenApiVO(activeDeviceSnOpenApiVO).Execute()

Batch active devices by SN



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
	activeDeviceSnOpenApiVO := *openapiclient.NewActiveDeviceSnOpenApiVO([]openapiclient.ActivePairSnOpenApiDTO{*openapiclient.NewActivePairSnOpenApiDTO("LicenseType_example", "Sn_example")}, "Category_example") // ActiveDeviceSnOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ActiveDeviceBySn(context.Background(), omadacId, siteId).ActiveDeviceSnOpenApiVO(activeDeviceSnOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ActiveDeviceBySn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActiveDeviceBySn`: []ActiveDeviceSnRespVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ActiveDeviceBySn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveDeviceBySnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **activeDeviceSnOpenApiVO** | [**ActiveDeviceSnOpenApiVO**](ActiveDeviceSnOpenApiVO.md) |  | 

### Return type

[**[]ActiveDeviceSnRespVO**](ActiveDeviceSnRespVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActiveDeviceWithMultiSite

> []ActiveDeviceRespVO ActiveDeviceWithMultiSite(ctx, omadacId).ActiveDeviceMultiSiteOpenApiVO(activeDeviceMultiSiteOpenApiVO).Execute()

Batch active devices with multi site



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
	activeDeviceMultiSiteOpenApiVO := *openapiclient.NewActiveDeviceMultiSiteOpenApiVO([]openapiclient.ActivePairMultiSiteOpenApiDTO{*openapiclient.NewActivePairMultiSiteOpenApiDTO("DeviceMac_example", "LicenseType_example")}, "Category_example") // ActiveDeviceMultiSiteOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ActiveDeviceWithMultiSite(context.Background(), omadacId).ActiveDeviceMultiSiteOpenApiVO(activeDeviceMultiSiteOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ActiveDeviceWithMultiSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActiveDeviceWithMultiSite`: []ActiveDeviceRespVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ActiveDeviceWithMultiSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveDeviceWithMultiSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeDeviceMultiSiteOpenApiVO** | [**ActiveDeviceMultiSiteOpenApiVO**](ActiveDeviceMultiSiteOpenApiVO.md) |  | 

### Return type

[**[]ActiveDeviceRespVO**](ActiveDeviceRespVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDevices

> OperationResponseDeviceAddRespOpenApiVO AddDevices(ctx, omadacId, siteId).DeviceListAddOpenApiVO(deviceListAddOpenApiVO).Execute()

Batch add devices



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
	deviceListAddOpenApiVO := *openapiclient.NewDeviceListAddOpenApiVO([]openapiclient.AddDeviceVO{*openapiclient.NewAddDeviceVO()}) // DeviceListAddOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.AddDevices(context.Background(), omadacId, siteId).DeviceListAddOpenApiVO(deviceListAddOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.AddDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDevices`: OperationResponseDeviceAddRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.AddDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceListAddOpenApiVO** | [**DeviceListAddOpenApiVO**](DeviceListAddOpenApiVO.md) |  | 

### Return type

[**OperationResponseDeviceAddRespOpenApiVO**](OperationResponseDeviceAddRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDevicesByDeviceKey

> OperationResponseDeviceAddRespOpenApiVO AddDevicesByDeviceKey(ctx, omadacId, siteId).DeviceListAddByDevicekeyOpenApiVO(deviceListAddByDevicekeyOpenApiVO).Execute()

Batch add devices by DEVICE KEY.



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
	deviceListAddByDevicekeyOpenApiVO := *openapiclient.NewDeviceListAddByDevicekeyOpenApiVO([]openapiclient.AddDeviceByDevicekeyOpenApiVO{*openapiclient.NewAddDeviceByDevicekeyOpenApiVO()}) // DeviceListAddByDevicekeyOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.AddDevicesByDeviceKey(context.Background(), omadacId, siteId).DeviceListAddByDevicekeyOpenApiVO(deviceListAddByDevicekeyOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.AddDevicesByDeviceKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDevicesByDeviceKey`: OperationResponseDeviceAddRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.AddDevicesByDeviceKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDevicesByDeviceKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceListAddByDevicekeyOpenApiVO** | [**DeviceListAddByDevicekeyOpenApiVO**](DeviceListAddByDevicekeyOpenApiVO.md) |  | 

### Return type

[**OperationResponseDeviceAddRespOpenApiVO**](OperationResponseDeviceAddRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDevicesBySn

> OperationResponseDeviceAddRespOpenApiVO AddDevicesBySn(ctx, omadacId, siteId).DeviceListAddBySnOpenApiVO(deviceListAddBySnOpenApiVO).Execute()

Batch add devices by SN



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
	deviceListAddBySnOpenApiVO := *openapiclient.NewDeviceListAddBySnOpenApiVO([]openapiclient.AddDeviceBySnOpenApiVO{*openapiclient.NewAddDeviceBySnOpenApiVO()}) // DeviceListAddBySnOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.AddDevicesBySn(context.Background(), omadacId, siteId).DeviceListAddBySnOpenApiVO(deviceListAddBySnOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.AddDevicesBySn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDevicesBySn`: OperationResponseDeviceAddRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.AddDevicesBySn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDevicesBySnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceListAddBySnOpenApiVO** | [**DeviceListAddBySnOpenApiVO**](DeviceListAddBySnOpenApiVO.md) |  | 

### Return type

[**OperationResponseDeviceAddRespOpenApiVO**](OperationResponseDeviceAddRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDevicesBySnWithMultiSite

> OperationResponseDeviceAddRespOpenApiVO AddDevicesBySnWithMultiSite(ctx, omadacId).DeviceListAddMultiSiteBySnOpenApiVO(deviceListAddMultiSiteBySnOpenApiVO).Execute()

Batch add devices by SN with multi site



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
	deviceListAddMultiSiteBySnOpenApiVO := *openapiclient.NewDeviceListAddMultiSiteBySnOpenApiVO([]openapiclient.AddDeviceWithSiteBySnOpenApiVO{*openapiclient.NewAddDeviceWithSiteBySnOpenApiVO()}) // DeviceListAddMultiSiteBySnOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.AddDevicesBySnWithMultiSite(context.Background(), omadacId).DeviceListAddMultiSiteBySnOpenApiVO(deviceListAddMultiSiteBySnOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.AddDevicesBySnWithMultiSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDevicesBySnWithMultiSite`: OperationResponseDeviceAddRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.AddDevicesBySnWithMultiSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDevicesBySnWithMultiSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceListAddMultiSiteBySnOpenApiVO** | [**DeviceListAddMultiSiteBySnOpenApiVO**](DeviceListAddMultiSiteBySnOpenApiVO.md) |  | 

### Return type

[**OperationResponseDeviceAddRespOpenApiVO**](OperationResponseDeviceAddRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDevicesForMsp

> OperationResponseDeviceAddRespOpenApiVO AddDevicesForMsp(ctx, mspId, customerId, siteId).MspAddDeviceListOpenApiVO(mspAddDeviceListOpenApiVO).Execute()

Batch add devices in MSP view



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	mspAddDeviceListOpenApiVO := *openapiclient.NewMspAddDeviceListOpenApiVO([]openapiclient.AddDeviceVO{*openapiclient.NewAddDeviceVO()}) // MspAddDeviceListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.AddDevicesForMsp(context.Background(), mspId, customerId, siteId).MspAddDeviceListOpenApiVO(mspAddDeviceListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.AddDevicesForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDevicesForMsp`: OperationResponseDeviceAddRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.AddDevicesForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDevicesForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mspAddDeviceListOpenApiVO** | [**MspAddDeviceListOpenApiVO**](MspAddDeviceListOpenApiVO.md) |  | 

### Return type

[**OperationResponseDeviceAddRespOpenApiVO**](OperationResponseDeviceAddRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTag

> TagRespOpenApiVO AddTag(ctx, omadacId, siteId).TagOpenApiVO(tagOpenApiVO).Execute()

Create new tag



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
	tagOpenApiVO := *openapiclient.NewTagOpenApiVO("Name_example") // TagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.AddTag(context.Background(), omadacId, siteId).TagOpenApiVO(tagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.AddTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTag`: TagRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.AddTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagOpenApiVO** | [**TagOpenApiVO**](TagOpenApiVO.md) |  | 

### Return type

[**TagRespOpenApiVO**](TagRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdoptDevice

> OperationResponseWithoutResult AdoptDevice(ctx, omadacId, siteId, deviceMac).AdoptDeviceRequest(adoptDeviceRequest).Execute()

Start adopt device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	adoptDeviceRequest := *openapiclient.NewAdoptDeviceRequest() // AdoptDeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.AdoptDevice(context.Background(), omadacId, siteId, deviceMac).AdoptDeviceRequest(adoptDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.AdoptDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdoptDevice`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.AdoptDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdoptDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **adoptDeviceRequest** | [**AdoptDeviceRequest**](AdoptDeviceRequest.md) |  | 

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


## BatchAdopt

> OperationResponseWithoutResult BatchAdopt(ctx, omadacId, siteId).BatchAdoptDeviceRequest(batchAdoptDeviceRequest).Execute()

batch adopt device



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
	batchAdoptDeviceRequest := *openapiclient.NewBatchAdoptDeviceRequest() // BatchAdoptDeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.BatchAdopt(context.Background(), omadacId, siteId).BatchAdoptDeviceRequest(batchAdoptDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.BatchAdopt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchAdopt`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.BatchAdopt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchAdoptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchAdoptDeviceRequest** | [**BatchAdoptDeviceRequest**](BatchAdoptDeviceRequest.md) |  | 

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


## CancelRollingUpgrade

> OperationResponseWithoutResult CancelRollingUpgrade(ctx, omadacId, siteId, taskId).Execute()

End the rolling upgrade task



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
	taskId := "taskId_example" // string | Task ID. The ID is the return value of 'Start batch rolling upgrade' interface.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.CancelRollingUpgrade(context.Background(), omadacId, siteId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.CancelRollingUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRollingUpgrade`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.CancelRollingUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**taskId** | **string** | Task ID. The ID is the return value of &#39;Start batch rolling upgrade&#39; interface. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRollingUpgradeRequest struct via the builder pattern


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


## CopyDeviceConfiguration

> OperationResponseWithoutResult CopyDeviceConfiguration(ctx, omadacId, siteId).DeviceCopyConfigurationOpenApiVO(deviceCopyConfigurationOpenApiVO).Execute()

Copy configuration



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
	deviceCopyConfigurationOpenApiVO := *openapiclient.NewDeviceCopyConfigurationOpenApiVO("SourceMac_example", "TargetMac_example") // DeviceCopyConfigurationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.CopyDeviceConfiguration(context.Background(), omadacId, siteId).DeviceCopyConfigurationOpenApiVO(deviceCopyConfigurationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.CopyDeviceConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyDeviceConfiguration`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.CopyDeviceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyDeviceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceCopyConfigurationOpenApiVO** | [**DeviceCopyConfigurationOpenApiVO**](DeviceCopyConfigurationOpenApiVO.md) |  | 

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


## DeleteTag

> OperationResponseWithoutResult DeleteTag(ctx, omadacId, siteId).DeleteTagOpenApiVO(deleteTagOpenApiVO).Execute()

Delete an existing tag



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
	deleteTagOpenApiVO := *openapiclient.NewDeleteTagOpenApiVO("TagId_example") // DeleteTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeleteTag(context.Background(), omadacId, siteId).DeleteTagOpenApiVO(deleteTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTag`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeleteTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteTagOpenApiVO** | [**DeleteTagOpenApiVO**](DeleteTagOpenApiVO.md) |  | 

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


## DownloadDeviceInfo

> OperationResponse DownloadDeviceInfo(ctx, omadacId, siteId, deviceMac, type_).Execute()

Download device information.



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	type_ := "type__example" // string | Device Type. Supported type: ap, gateway, switch, olt.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DownloadDeviceInfo(context.Background(), omadacId, siteId, deviceMac, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DownloadDeviceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDeviceInfo`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DownloadDeviceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**type_** | **string** | Device Type. Supported type: ap, gateway, switch, olt. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDeviceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExistUnSupportRadSecDevice

> OperationResponseUnSupportRadSecOpenApiVO ExistUnSupportRadSecDevice(ctx, omadacId, siteId).Execute()

Check for unsupported RadSec devices.



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
	resp, r, err := apiClient.DeviceAPI.ExistUnSupportRadSecDevice(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ExistUnSupportRadSecDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExistUnSupportRadSecDevice`: OperationResponseUnSupportRadSecOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ExistUnSupportRadSecDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExistUnSupportRadSecDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUnSupportRadSecOpenApiVO**](OperationResponseUnSupportRadSecOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAddFailDevices

> OperationResponse ExportAddFailDevices(ctx, omadacId, operateId).Format(format).Execute()

Export the device file that failed to add in GLOBAL view.



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
	operateId := "operateId_example" // string | Operation ID
	format := int32(56) // int32 | File format. 0: CSV, 1: XLSX

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ExportAddFailDevices(context.Background(), omadacId, operateId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ExportAddFailDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAddFailDevices`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ExportAddFailDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**operateId** | **string** | Operation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAddFailDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **int32** | File format. 0: CSV, 1: XLSX | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportGlobalDeviceList

> OperationResponse ExportGlobalDeviceList(ctx, omadacId).ExportGlobalDeviceListOpenApiVO(exportGlobalDeviceListOpenApiVO).Execute()

Export device list in GLOBAL view



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
	exportGlobalDeviceListOpenApiVO := *openapiclient.NewExportGlobalDeviceListOpenApiVO(int32(123), int32(123)) // ExportGlobalDeviceListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ExportGlobalDeviceList(context.Background(), omadacId).ExportGlobalDeviceListOpenApiVO(exportGlobalDeviceListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ExportGlobalDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportGlobalDeviceList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ExportGlobalDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportGlobalDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportGlobalDeviceListOpenApiVO** | [**ExportGlobalDeviceListOpenApiVO**](ExportGlobalDeviceListOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportMspAddFailDevices

> OperationResponse ExportMspAddFailDevices(ctx, mspId, customerId, operateId).Format(format).Execute()

Export the device file that failed to add in MSP view.



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	operateId := "operateId_example" // string | Operation ID
	format := int32(56) // int32 | File format. 0: CSV, 1: XLSX

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ExportMspAddFailDevices(context.Background(), mspId, customerId, operateId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ExportMspAddFailDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportMspAddFailDevices`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ExportMspAddFailDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**operateId** | **string** | Operation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportMspAddFailDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **int32** | File format. 0: CSV, 1: XLSX | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPreConfigDevices

> OperationResponse ExportPreConfigDevices(ctx, omadacId, siteId).Execute()

Export preconfigured devices of site.



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
	resp, r, err := apiClient.DeviceAPI.ExportPreConfigDevices(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ExportPreConfigDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportPreConfigDevices`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ExportPreConfigDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPreConfigDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPreConfigDevicesGlobal

> OperationResponse ExportPreConfigDevicesGlobal(ctx, omadacId).Execute()

Export preconfigured devices in GLOBAL view.



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
	resp, r, err := apiClient.DeviceAPI.ExportPreConfigDevicesGlobal(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ExportPreConfigDevicesGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportPreConfigDevicesGlobal`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ExportPreConfigDevicesGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPreConfigDevicesGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSiteDeviceList

> OperationResponse ExportSiteDeviceList(ctx, omadacId, siteId).ExportGlobalDeviceListOpenApiVO(exportGlobalDeviceListOpenApiVO).Execute()

Export device list of site



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
	exportGlobalDeviceListOpenApiVO := *openapiclient.NewExportGlobalDeviceListOpenApiVO(int32(123), int32(123)) // ExportGlobalDeviceListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ExportSiteDeviceList(context.Background(), omadacId, siteId).ExportGlobalDeviceListOpenApiVO(exportGlobalDeviceListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ExportSiteDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSiteDeviceList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ExportSiteDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSiteDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exportGlobalDeviceListOpenApiVO** | [**ExportGlobalDeviceListOpenApiVO**](ExportGlobalDeviceListOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForceProvisionDevice

> OperationResponseWithoutResult ForceProvisionDevice(ctx, omadacId, siteId, deviceMac).Execute()

Force provision device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ForceProvisionDevice(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ForceProvisionDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForceProvisionDevice`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ForceProvisionDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiForceProvisionDeviceRequest struct via the builder pattern


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


## ForgetDevice

> OperationResponseWithoutResult ForgetDevice(ctx, omadacId, siteId, deviceMac).Execute()

Forget device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ForgetDevice(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ForgetDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgetDevice`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ForgetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiForgetDeviceRequest struct via the builder pattern


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


## ForgetDeviceForMsp

> OperationResponseWithoutResult ForgetDeviceForMsp(ctx, mspId, customerId, siteId, deviceMac).Execute()

Forget one device in MSP view.



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ForgetDeviceForMsp(context.Background(), mspId, customerId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ForgetDeviceForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgetDeviceForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ForgetDeviceForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiForgetDeviceForMspRequest struct via the builder pattern


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


## GetAddDevicesStatus

> OperationResponseDeviceAddProcessRespVO GetAddDevicesStatus(ctx, omadacId, siteId, operateId).Execute()

Batch get added device status



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
	operateId := "operateId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetAddDevicesStatus(context.Background(), omadacId, siteId, operateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetAddDevicesStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddDevicesStatus`: OperationResponseDeviceAddProcessRespVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetAddDevicesStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**operateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddDevicesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseDeviceAddProcessRespVO**](OperationResponseDeviceAddProcessRespVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddDevicesStatusMultiSite

> OperationResponseDeviceAddProcessRespVO GetAddDevicesStatusMultiSite(ctx, omadacId, operateId).Execute()

Batch get added device status with multi site



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
	operateId := "operateId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetAddDevicesStatusMultiSite(context.Background(), omadacId, operateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetAddDevicesStatusMultiSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddDevicesStatusMultiSite`: OperationResponseDeviceAddProcessRespVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetAddDevicesStatusMultiSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**operateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddDevicesStatusMultiSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDeviceAddProcessRespVO**](OperationResponseDeviceAddProcessRespVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdoptAbleDevicesForGlobalBatchAdopt

> OperationResponseListBriefDeviceInfoDetailForBatchAdoptOpenApiVO GetAdoptAbleDevicesForGlobalBatchAdopt(ctx, omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsStatus(sortsStatus).SortsIp(sortsIp).SearchKey(searchKey).FiltersTag(filtersTag).Execute()

Batch get info of adopted device in global view.



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name,mac,ip (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetAdoptAbleDevicesForGlobalBatchAdopt(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsStatus(sortsStatus).SortsIp(sortsIp).SearchKey(searchKey).FiltersTag(filtersTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetAdoptAbleDevicesForGlobalBatchAdopt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdoptAbleDevicesForGlobalBatchAdopt`: OperationResponseListBriefDeviceInfoDetailForBatchAdoptOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetAdoptAbleDevicesForGlobalBatchAdopt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdoptAbleDevicesForGlobalBatchAdoptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStatus** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name,mac,ip | 
 **filtersTag** | **string** | Filter query parameters, support field tag name | 

### Return type

[**OperationResponseListBriefDeviceInfoDetailForBatchAdoptOpenApiVO**](OperationResponseListBriefDeviceInfoDetailForBatchAdoptOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdoptAbleDevicesForMspBatchAdopt

> OperationResponseListBriefDeviceInfoDetailForBatchAdoptOpenApiVO GetAdoptAbleDevicesForMspBatchAdopt(ctx, mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsStatus(sortsStatus).SortsIp(sortsIp).SearchKey(searchKey).FiltersTag(filtersTag).Execute()

Batch get info of adopted device in MSP view.



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
	mspId := "mspId_example" // string | MSP ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name,mac,ip (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetAdoptAbleDevicesForMspBatchAdopt(context.Background(), mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsStatus(sortsStatus).SortsIp(sortsIp).SearchKey(searchKey).FiltersTag(filtersTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetAdoptAbleDevicesForMspBatchAdopt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdoptAbleDevicesForMspBatchAdopt`: OperationResponseListBriefDeviceInfoDetailForBatchAdoptOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetAdoptAbleDevicesForMspBatchAdopt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdoptAbleDevicesForMspBatchAdoptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStatus** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name,mac,ip | 
 **filtersTag** | **string** | Filter query parameters, support field tag name | 

### Return type

[**OperationResponseListBriefDeviceInfoDetailForBatchAdoptOpenApiVO**](OperationResponseListBriefDeviceInfoDetailForBatchAdoptOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdoptTip

> OperationResponseAdoptTipOpenApiVO GetAdoptTip(ctx, omadacId).Execute()

Get adopt tip



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
	resp, r, err := apiClient.DeviceAPI.GetAdoptTip(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetAdoptTip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdoptTip`: OperationResponseAdoptTipOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetAdoptTip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdoptTipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseAdoptTipOpenApiVO**](OperationResponseAdoptTipOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDeviceBySite

> OperationResponseListDeviceInfo GetAllDeviceBySite(ctx, omadacId, siteId).Execute()

Get device list info.



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
	resp, r, err := apiClient.DeviceAPI.GetAllDeviceBySite(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetAllDeviceBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDeviceBySite`: OperationResponseListDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetAllDeviceBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDeviceBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListDeviceInfo**](OperationResponseListDeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoSelectDevices

> OperationResponseQueryDeviceListOpenApiVO GetAutoSelectDevices(ctx, omadacId, siteId).QueryDeviceListOpenApiVO(queryDeviceListOpenApiVO).Execute()

Get auto select devices in quick-config page.



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
	queryDeviceListOpenApiVO := *openapiclient.NewQueryDeviceListOpenApiVO([]string{"Macs_example"}) // QueryDeviceListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetAutoSelectDevices(context.Background(), omadacId, siteId).QueryDeviceListOpenApiVO(queryDeviceListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetAutoSelectDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoSelectDevices`: OperationResponseQueryDeviceListOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetAutoSelectDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoSelectDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryDeviceListOpenApiVO** | [**QueryDeviceListOpenApiVO**](QueryDeviceListOpenApiVO.md) |  | 

### Return type

[**OperationResponseQueryDeviceListOpenApiVO**](OperationResponseQueryDeviceListOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckFirmwareTaskResult

> OperationResponseCheckFirmwareRes GetCheckFirmwareTaskResult(ctx, omadacId, siteId, taskId).Execute()

Get the result of the given check latest firmware task



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
	taskId := "taskId_example" // string | Task ID. The ID is the return value of 'Start check latest firmware for all devices' interface

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetCheckFirmwareTaskResult(context.Background(), omadacId, siteId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetCheckFirmwareTaskResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckFirmwareTaskResult`: OperationResponseCheckFirmwareRes
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetCheckFirmwareTaskResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**taskId** | **string** | Task ID. The ID is the return value of &#39;Start check latest firmware for all devices&#39; interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckFirmwareTaskResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseCheckFirmwareRes**](OperationResponseCheckFirmwareRes.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAdoptResult

> OperationResponseAdoptResult GetDeviceAdoptResult(ctx, omadacId, siteId, deviceMac).Execute()

Get device adopt result



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetDeviceAdoptResult(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDeviceAdoptResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAdoptResult`: OperationResponseAdoptResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDeviceAdoptResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAdoptResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseAdoptResult**](OperationResponseAdoptResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceList

> OperationResponseGridVODeviceInfo GetDeviceList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsStatus(sortsStatus).SortsIp(sortsIp).SearchKey(searchKey).FiltersTag(filtersTag).Execute()

Get site device list



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
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name,mac,ip (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetDeviceList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsStatus(sortsStatus).SortsIp(sortsIp).SearchKey(searchKey).FiltersTag(filtersTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceList`: OperationResponseGridVODeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStatus** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name,mac,ip | 
 **filtersTag** | **string** | Filter query parameters, support field tag name | 

### Return type

[**OperationResponseGridVODeviceInfo**](OperationResponseGridVODeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceRememberMe

> OperationResponseDeviceRememberConfig GetDeviceRememberMe(ctx, omadacId, siteId, deviceMac).Execute()

Get device remember Config



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetDeviceRememberMe(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDeviceRememberMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceRememberMe`: OperationResponseDeviceRememberConfig
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDeviceRememberMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRememberMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseDeviceRememberConfig**](OperationResponseDeviceRememberConfig.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWhiteList

> OperationResponseListWhiteDeviceInfoOpenApiVO GetDeviceWhiteList(ctx, omadacId, siteId).Execute()

Get the whitelist list of devices.



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
	resp, r, err := apiClient.DeviceAPI.GetDeviceWhiteList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDeviceWhiteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceWhiteList`: OperationResponseListWhiteDeviceInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDeviceWhiteList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWhiteListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListWhiteDeviceInfoOpenApiVO**](OperationResponseListWhiteDeviceInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirmwareInfo

> OperationResponseDeviceFirmwareInfo GetFirmwareInfo(ctx, omadacId, siteId, deviceMac).Execute()

Get the latest firmware info of the device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetFirmwareInfo(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetFirmwareInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirmwareInfo`: OperationResponseDeviceFirmwareInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetFirmwareInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirmwareInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseDeviceFirmwareInfo**](OperationResponseDeviceFirmwareInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForgetProcess

> OperationResponseWithoutResult GetForgetProcess(ctx, omadacId, siteId, forgetId).Execute()

Get batch forget process



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
	forgetId := "forgetId_example" // string | Forget ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetForgetProcess(context.Background(), omadacId, siteId, forgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetForgetProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForgetProcess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetForgetProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**forgetId** | **string** | Forget ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForgetProcessRequest struct via the builder pattern


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


## GetGlobalKnownDeviceList

> OperationResponseGridVOGlobalKnownDeviceOpenApiVO GetGlobalKnownDeviceList(ctx, omadacId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).SearchSns(searchSns).FiltersTag(filtersTag).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()

Get global known device list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsMac := "sortsMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchMacs := "searchMacs_example" // string | Fuzzy query parameters, support field mac (optional)
	searchNames := "searchNames_example" // string | Fuzzy query parameters, support field name (optional)
	searchModels := "searchModels_example" // string | Fuzzy query parameters, support field model (optional)
	searchSns := "searchSns_example" // string | Fuzzy query parameters, support field sn (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag ID (optional)
	filtersDeviceSeriesType := "filtersDeviceSeriesType_example" // string | Filter query parameters, support field Device series type. 0: basic; 1: pro. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetGlobalKnownDeviceList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).SearchSns(searchSns).FiltersTag(filtersTag).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetGlobalKnownDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalKnownDeviceList`: OperationResponseGridVOGlobalKnownDeviceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetGlobalKnownDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalKnownDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchMacs** | **string** | Fuzzy query parameters, support field mac | 
 **searchNames** | **string** | Fuzzy query parameters, support field name | 
 **searchModels** | **string** | Fuzzy query parameters, support field model | 
 **searchSns** | **string** | Fuzzy query parameters, support field sn | 
 **filtersTag** | **string** | Filter query parameters, support field tag ID | 
 **filtersDeviceSeriesType** | **string** | Filter query parameters, support field Device series type. 0: basic; 1: pro. | 

### Return type

[**OperationResponseGridVOGlobalKnownDeviceOpenApiVO**](OperationResponseGridVOGlobalKnownDeviceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalUnknownDeviceList

> OperationResponseGridVOGlobalUnknownDeviceOpenApiVO GetGlobalUnknownDeviceList(ctx, omadacId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).Execute()

Get global unknown device list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsMac := "sortsMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchMacs := "searchMacs_example" // string | Fuzzy query parameters, support field mac (optional)
	searchNames := "searchNames_example" // string | Fuzzy query parameters, support field name (optional)
	searchModels := "searchModels_example" // string | Fuzzy query parameters, support field model (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetGlobalUnknownDeviceList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetGlobalUnknownDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalUnknownDeviceList`: OperationResponseGridVOGlobalUnknownDeviceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetGlobalUnknownDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalUnknownDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchMacs** | **string** | Fuzzy query parameters, support field mac | 
 **searchNames** | **string** | Fuzzy query parameters, support field name | 
 **searchModels** | **string** | Fuzzy query parameters, support field model | 

### Return type

[**OperationResponseGridVOGlobalUnknownDeviceOpenApiVO**](OperationResponseGridVOGlobalUnknownDeviceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridAdoptedBridgeDevicesBySite

> OperationResponseAdoptedDeviceGridVODeviceInfo GetGridAdoptedBridgeDevicesBySite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).SearchSns(searchSns).FiltersTag(filtersTag).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()

Get Bridge group grouped devices



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
	sortsMac := "sortsMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchMacs := "searchMacs_example" // string | Fuzzy query parameters, support field mac (optional)
	searchNames := "searchNames_example" // string | Fuzzy query parameters, support field name (optional)
	searchModels := "searchModels_example" // string | Fuzzy query parameters, support field model (optional)
	searchSns := "searchSns_example" // string | Fuzzy query parameters, support field sn (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag ID (optional)
	filtersDeviceSeriesType := "filtersDeviceSeriesType_example" // string | Filter query parameters, support field Device series type. 0: basic; 1: pro. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetGridAdoptedBridgeDevicesBySite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).SearchSns(searchSns).FiltersTag(filtersTag).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetGridAdoptedBridgeDevicesBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAdoptedBridgeDevicesBySite`: OperationResponseAdoptedDeviceGridVODeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetGridAdoptedBridgeDevicesBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAdoptedBridgeDevicesBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchMacs** | **string** | Fuzzy query parameters, support field mac | 
 **searchNames** | **string** | Fuzzy query parameters, support field name | 
 **searchModels** | **string** | Fuzzy query parameters, support field model | 
 **searchSns** | **string** | Fuzzy query parameters, support field sn | 
 **filtersTag** | **string** | Filter query parameters, support field tag ID | 
 **filtersDeviceSeriesType** | **string** | Filter query parameters, support field Device series type. 0: basic; 1: pro. | 

### Return type

[**OperationResponseAdoptedDeviceGridVODeviceInfo**](OperationResponseAdoptedDeviceGridVODeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridAdoptedDevicesStatByGlobal

> OperationResponseGlobalDeviceStatOpenApiVO GetGridAdoptedDevicesStatByGlobal(ctx, omadacId).Execute()

Query the statistics for the list of global adopted devices.



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
	resp, r, err := apiClient.DeviceAPI.GetGridAdoptedDevicesStatByGlobal(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetGridAdoptedDevicesStatByGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAdoptedDevicesStatByGlobal`: OperationResponseGlobalDeviceStatOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetGridAdoptedDevicesStatByGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAdoptedDevicesStatByGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseGlobalDeviceStatOpenApiVO**](OperationResponseGlobalDeviceStatOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridPendingBridgeDevicesBySite

> OperationResponseGridVODeviceInfo GetGridPendingBridgeDevicesBySite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).SearchSns(searchSns).FiltersTag(filtersTag).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()

Get Bridge group ungrouped devices



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
	sortsMac := "sortsMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchMacs := "searchMacs_example" // string | Fuzzy query parameters, support field mac (optional)
	searchNames := "searchNames_example" // string | Fuzzy query parameters, support field name (optional)
	searchModels := "searchModels_example" // string | Fuzzy query parameters, support field model (optional)
	searchSns := "searchSns_example" // string | Fuzzy query parameters, support field sn (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag ID (optional)
	filtersDeviceSeriesType := "filtersDeviceSeriesType_example" // string | Filter query parameters, support field Device series type. 0: basic; 1: pro. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetGridPendingBridgeDevicesBySite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).SearchSns(searchSns).FiltersTag(filtersTag).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetGridPendingBridgeDevicesBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPendingBridgeDevicesBySite`: OperationResponseGridVODeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetGridPendingBridgeDevicesBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridPendingBridgeDevicesBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchMacs** | **string** | Fuzzy query parameters, support field mac | 
 **searchNames** | **string** | Fuzzy query parameters, support field name | 
 **searchModels** | **string** | Fuzzy query parameters, support field model | 
 **searchSns** | **string** | Fuzzy query parameters, support field sn | 
 **filtersTag** | **string** | Filter query parameters, support field tag ID | 
 **filtersDeviceSeriesType** | **string** | Filter query parameters, support field Device series type. 0: basic; 1: pro. | 

### Return type

[**OperationResponseGridVODeviceInfo**](OperationResponseGridVODeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridPendingDevicesBySite

> OperationResponseGridVODeviceInfo GetGridPendingDevicesBySite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get adoptable device list of target site



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
	searchKey := "searchKey_example" // string | searchKey (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetGridPendingDevicesBySite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetGridPendingDevicesBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPendingDevicesBySite`: OperationResponseGridVODeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetGridPendingDevicesBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridPendingDevicesBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | searchKey | 

### Return type

[**OperationResponseGridVODeviceInfo**](OperationResponseGridVODeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManuallyUpgradeRes

> OperationResponseUpgradeRes GetManuallyUpgradeRes(ctx, omadacId, siteId, taskId).Execute()

Get the result of the given manually upgrade task



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
	taskId := "taskId_example" // string | Task ID. The ID is the return value of 'Start manually upgrade' interface.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetManuallyUpgradeRes(context.Background(), omadacId, siteId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetManuallyUpgradeRes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManuallyUpgradeRes`: OperationResponseUpgradeRes
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetManuallyUpgradeRes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**taskId** | **string** | Task ID. The ID is the return value of &#39;Start manually upgrade&#39; interface. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManuallyUpgradeResRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseUpgradeRes**](OperationResponseUpgradeRes.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoveSiteProcess

> OperationResponseWithoutResult GetMoveSiteProcess(ctx, omadacId, siteId, moveSiteId).Execute()

Get batch move site process



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
	moveSiteId := "moveSiteId_example" // string | Move Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetMoveSiteProcess(context.Background(), omadacId, siteId, moveSiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetMoveSiteProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMoveSiteProcess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetMoveSiteProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**moveSiteId** | **string** | Move Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoveSiteProcessRequest struct via the builder pattern


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


## GetOnlineTimeline

> OperationResponseTimelineOpenApiVO GetOnlineTimeline(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get device online timeline



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	dto := *openapiclient.NewTimeIntervalQueryOpenApiVO(int64(123), int64(123)) // TimeIntervalQueryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetOnlineTimeline(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetOnlineTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnlineTimeline`: OperationResponseTimelineOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetOnlineTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnlineTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**TimeIntervalQueryOpenApiVO**](TimeIntervalQueryOpenApiVO.md) |  | 

### Return type

[**OperationResponseTimelineOpenApiVO**](OperationResponseTimelineOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnlineUpgradeRes

> OperationResponseOnlineUpgradeRes GetOnlineUpgradeRes(ctx, omadacId, siteId, deviceMac).Execute()

Get online upgrade result



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetOnlineUpgradeRes(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetOnlineUpgradeRes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnlineUpgradeRes`: OperationResponseOnlineUpgradeRes
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetOnlineUpgradeRes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnlineUpgradeResRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOnlineUpgradeRes**](OperationResponseOnlineUpgradeRes.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollingUpgradeRes

> OperationResponseRollingUpgradeRes GetRollingUpgradeRes(ctx, omadacId, siteId, taskId).Execute()

Get the result of the given rolling upgrade task



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
	taskId := "taskId_example" // string | Task ID. The ID is the return value of 'Start batch rolling upgrade' interface.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetRollingUpgradeRes(context.Background(), omadacId, siteId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetRollingUpgradeRes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRollingUpgradeRes`: OperationResponseRollingUpgradeRes
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetRollingUpgradeRes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**taskId** | **string** | Task ID. The ID is the return value of &#39;Start batch rolling upgrade&#39; interface. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRollingUpgradeResRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseRollingUpgradeRes**](OperationResponseRollingUpgradeRes.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleForgetProcess

> OperationResponseWithoutResult GetSingleForgetProcess(ctx, omadacId, siteId, deviceMac, forgetId).Execute()

Get forget process of device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	forgetId := "forgetId_example" // string | Forget ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetSingleForgetProcess(context.Background(), omadacId, siteId, deviceMac, forgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetSingleForgetProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleForgetProcess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetSingleForgetProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**forgetId** | **string** | Forget ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleForgetProcessRequest struct via the builder pattern


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


## GetSingleMoveSiteProcess

> OperationResponseWithoutResult GetSingleMoveSiteProcess(ctx, omadacId, siteId, deviceMac, moveSiteId).Execute()

Get move site process of device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	moveSiteId := "moveSiteId_example" // string | Move Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetSingleMoveSiteProcess(context.Background(), omadacId, siteId, deviceMac, moveSiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetSingleMoveSiteProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleMoveSiteProcess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetSingleMoveSiteProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**moveSiteId** | **string** | Move Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleMoveSiteProcessRequest struct via the builder pattern


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


## GetTags

> []TagRespOpenApiVO GetTags(ctx, omadacId, siteId).Execute()

Get tag list



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
	resp, r, err := apiClient.DeviceAPI.GetTags(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: []TagRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TagRespOpenApiVO**](TagRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUplinkDeviceInfo

> OperationResponseListDeviceUplinkOpenApiVO GetUplinkDeviceInfo(ctx, omadacId, siteId).QueryDeviceUplinkOpenApiVO(queryDeviceUplinkOpenApiVO).Execute()

Query uplink information for specified device MAC addresses under the site.



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
	queryDeviceUplinkOpenApiVO := *openapiclient.NewQueryDeviceUplinkOpenApiVO() // QueryDeviceUplinkOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetUplinkDeviceInfo(context.Background(), omadacId, siteId).QueryDeviceUplinkOpenApiVO(queryDeviceUplinkOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetUplinkDeviceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUplinkDeviceInfo`: OperationResponseListDeviceUplinkOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetUplinkDeviceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUplinkDeviceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryDeviceUplinkOpenApiVO** | [**QueryDeviceUplinkOpenApiVO**](QueryDeviceUplinkOpenApiVO.md) |  | 

### Return type

[**OperationResponseListDeviceUplinkOpenApiVO**](OperationResponseListDeviceUplinkOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocateDevice

> OperationResponseWithoutResult LocateDevice(ctx, omadacId, siteId, deviceMac).LocateDeviceRequest(locateDeviceRequest).Execute()

Locate device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	locateDeviceRequest := *openapiclient.NewLocateDeviceRequest() // LocateDeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.LocateDevice(context.Background(), omadacId, siteId, deviceMac).LocateDeviceRequest(locateDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.LocateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocateDevice`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.LocateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **locateDeviceRequest** | [**LocateDeviceRequest**](LocateDeviceRequest.md) |  | 

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


## LocatePorts

> OperationResponseLocateOswPortsResultOpenApiVO LocatePorts(ctx, omadacId, siteId).LocateOswPortsOpenApiVO(locateOswPortsOpenApiVO).Execute()

Locate multiple ports of multiple switches



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
	locateOswPortsOpenApiVO := *openapiclient.NewLocateOswPortsOpenApiVO(false) // LocateOswPortsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.LocatePorts(context.Background(), omadacId, siteId).LocateOswPortsOpenApiVO(locateOswPortsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.LocatePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocatePorts`: OperationResponseLocateOswPortsResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.LocatePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocatePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locateOswPortsOpenApiVO** | [**LocateOswPortsOpenApiVO**](LocateOswPortsOpenApiVO.md) |  | 

### Return type

[**OperationResponseLocateOswPortsResultOpenApiVO**](OperationResponseLocateOswPortsResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManuallyUpgrade

> OperationResponseUpgradeTask ManuallyUpgrade(ctx, omadacId, siteId).UpgradeRequest(upgradeRequest).Execute()

Start manually upgrade



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
	upgradeRequest := *openapiclient.NewUpgradeRequest() // UpgradeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ManuallyUpgrade(context.Background(), omadacId, siteId).UpgradeRequest(upgradeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ManuallyUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManuallyUpgrade`: OperationResponseUpgradeTask
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ManuallyUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiManuallyUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upgradeRequest** | [**UpgradeRequest**](UpgradeRequest.md) |  | 

### Return type

[**OperationResponseUpgradeTask**](OperationResponseUpgradeTask.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDeviceRememberMe

> OperationResponseWithoutResult ModifyDeviceRememberMe(ctx, omadacId, siteId, deviceMac).DeviceRememberConfig(deviceRememberConfig).Execute()

Modify device remember Config



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	deviceRememberConfig := *openapiclient.NewDeviceRememberConfig() // DeviceRememberConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ModifyDeviceRememberMe(context.Background(), omadacId, siteId, deviceMac).DeviceRememberConfig(deviceRememberConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ModifyDeviceRememberMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDeviceRememberMe`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ModifyDeviceRememberMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDeviceRememberMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deviceRememberConfig** | [**DeviceRememberConfig**](DeviceRememberConfig.md) |  | 

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


## ModifyTag

> OperationResponseWithoutResult ModifyTag(ctx, omadacId, siteId).ModifyTagOpenApiVO(modifyTagOpenApiVO).Execute()

Modify an existing tag



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
	modifyTagOpenApiVO := *openapiclient.NewModifyTagOpenApiVO("Name_example", "TagId_example") // ModifyTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ModifyTag(context.Background(), omadacId, siteId).ModifyTagOpenApiVO(modifyTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ModifyTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTag`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ModifyTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyTagOpenApiVO** | [**ModifyTagOpenApiVO**](ModifyTagOpenApiVO.md) |  | 

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


## MoveSite

> OperationResponseWithoutResult MoveSite(ctx, omadacId, siteId, deviceMac).DeviceMoveSiteOpenApiVO(deviceMoveSiteOpenApiVO).Execute()

Move site



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	deviceMoveSiteOpenApiVO := *openapiclient.NewDeviceMoveSiteOpenApiVO("TargetSite_example") // DeviceMoveSiteOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.MoveSite(context.Background(), omadacId, siteId, deviceMac).DeviceMoveSiteOpenApiVO(deviceMoveSiteOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.MoveSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.MoveSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deviceMoveSiteOpenApiVO** | [**DeviceMoveSiteOpenApiVO**](DeviceMoveSiteOpenApiVO.md) |  | 

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


## MoveSiteForAps

> OperationResponseMoveSiteIdOpenApiVO MoveSiteForAps(ctx, omadacId, siteId).ApMoveSiteOpenApiVO(apMoveSiteOpenApiVO).Execute()

Move aps to another site



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
	apMoveSiteOpenApiVO := *openapiclient.NewApMoveSiteOpenApiVO("TargetSite_example") // ApMoveSiteOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.MoveSiteForAps(context.Background(), omadacId, siteId).ApMoveSiteOpenApiVO(apMoveSiteOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.MoveSiteForAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveSiteForAps`: OperationResponseMoveSiteIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.MoveSiteForAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveSiteForApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apMoveSiteOpenApiVO** | [**ApMoveSiteOpenApiVO**](ApMoveSiteOpenApiVO.md) |  | 

### Return type

[**OperationResponseMoveSiteIdOpenApiVO**](OperationResponseMoveSiteIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnlineCheckUpgrade

> OperationResponseUpgradeTask OnlineCheckUpgrade(ctx, omadacId, siteId).Execute()

Start check latest firmware for all devices



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
	resp, r, err := apiClient.DeviceAPI.OnlineCheckUpgrade(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.OnlineCheckUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnlineCheckUpgrade`: OperationResponseUpgradeTask
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.OnlineCheckUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnlineCheckUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUpgradeTask**](OperationResponseUpgradeTask.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnlineRollingUpgrade

> OperationResponseUpgradeTask OnlineRollingUpgrade(ctx, omadacId, siteId).RollingUpgradeRequest(rollingUpgradeRequest).Execute()

Start batch rolling upgrade



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
	rollingUpgradeRequest := *openapiclient.NewRollingUpgradeRequest() // RollingUpgradeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.OnlineRollingUpgrade(context.Background(), omadacId, siteId).RollingUpgradeRequest(rollingUpgradeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.OnlineRollingUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnlineRollingUpgrade`: OperationResponseUpgradeTask
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.OnlineRollingUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnlineRollingUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rollingUpgradeRequest** | [**RollingUpgradeRequest**](RollingUpgradeRequest.md) |  | 

### Return type

[**OperationResponseUpgradeTask**](OperationResponseUpgradeTask.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnlineRollingUpgradeByQuery

> OperationResponseUpgradeTask OnlineRollingUpgradeByQuery(ctx, omadacId, siteId).RollingUpgradeRequest(rollingUpgradeRequest).FiltersEcspFirstVersion(filtersEcspFirstVersion).Execute()

Start batch rolling upgrade By Query



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
	rollingUpgradeRequest := *openapiclient.NewRollingUpgradeRequest() // RollingUpgradeRequest | 
	filtersEcspFirstVersion := "filtersEcspFirstVersion_example" // string | Filter query parameters, support field ecsp first version. 1: ecsp v1; 2: ecsp v2 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.OnlineRollingUpgradeByQuery(context.Background(), omadacId, siteId).RollingUpgradeRequest(rollingUpgradeRequest).FiltersEcspFirstVersion(filtersEcspFirstVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.OnlineRollingUpgradeByQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnlineRollingUpgradeByQuery`: OperationResponseUpgradeTask
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.OnlineRollingUpgradeByQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnlineRollingUpgradeByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rollingUpgradeRequest** | [**RollingUpgradeRequest**](RollingUpgradeRequest.md) |  | 
 **filtersEcspFirstVersion** | **string** | Filter query parameters, support field ecsp first version. 1: ecsp v1; 2: ecsp v2 | 

### Return type

[**OperationResponseUpgradeTask**](OperationResponseUpgradeTask.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnlineUpgrade

> OperationResponseWithoutResult OnlineUpgrade(ctx, omadacId, siteId, deviceMac).Execute()

Start online upgrade



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.OnlineUpgrade(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.OnlineUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OnlineUpgrade`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.OnlineUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnlineUpgradeRequest struct via the builder pattern


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


## RebootDevice

> OperationResponseWithoutResult RebootDevice(ctx, omadacId, siteId, deviceMac).Execute()

Reboot device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.RebootDevice(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.RebootDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootDevice`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.RebootDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootDeviceRequest struct via the builder pattern


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


## RetryAddDevice

> OperationResponseRetryAddDeviceRespOpenApiVO RetryAddDevice(ctx, omadacId, siteId, deviceMac).RetryAddDeviceOpenApiVO(retryAddDeviceOpenApiVO).Execute()

Retry add device



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	retryAddDeviceOpenApiVO := *openapiclient.NewRetryAddDeviceOpenApiVO() // RetryAddDeviceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.RetryAddDevice(context.Background(), omadacId, siteId, deviceMac).RetryAddDeviceOpenApiVO(retryAddDeviceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.RetryAddDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryAddDevice`: OperationResponseRetryAddDeviceRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.RetryAddDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryAddDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **retryAddDeviceOpenApiVO** | [**RetryAddDeviceOpenApiVO**](RetryAddDeviceOpenApiVO.md) |  | 

### Return type

[**OperationResponseRetryAddDeviceRespOpenApiVO**](OperationResponseRetryAddDeviceRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryAddDeviceByMsp

> OperationResponseRetryAddDeviceRespOpenApiVO RetryAddDeviceByMsp(ctx, customerId, deviceMac, mspId, siteId).RetryAddDeviceOpenApiVO(retryAddDeviceOpenApiVO).Execute()

retry add device in msp view



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
	customerId := "customerId_example" // string | Omada ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	mspId := "mspId_example" // string | mspId
	siteId := "siteId_example" // string | siteId
	retryAddDeviceOpenApiVO := *openapiclient.NewRetryAddDeviceOpenApiVO() // RetryAddDeviceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.RetryAddDeviceByMsp(context.Background(), customerId, deviceMac, mspId, siteId).RetryAddDeviceOpenApiVO(retryAddDeviceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.RetryAddDeviceByMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryAddDeviceByMsp`: OperationResponseRetryAddDeviceRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.RetryAddDeviceByMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Omada ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**mspId** | **string** | mspId | 
**siteId** | **string** | siteId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryAddDeviceByMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **retryAddDeviceOpenApiVO** | [**RetryAddDeviceOpenApiVO**](RetryAddDeviceOpenApiVO.md) |  | 

### Return type

[**OperationResponseRetryAddDeviceRespOpenApiVO**](OperationResponseRetryAddDeviceRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGlobalDevice

> OperationResponseGlobalSearchResultOpenApiVO SearchGlobalDevice(ctx, omadacId).SearchKey(searchKey).Execute()

Global search for devices returns the devices under the sites that have permissions.



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
	searchKey := "searchKey_example" // string | searchKey

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.SearchGlobalDevice(context.Background(), omadacId).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.SearchGlobalDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGlobalDevice`: OperationResponseGlobalSearchResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.SearchGlobalDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGlobalDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchKey** | **string** | searchKey | 

### Return type

[**OperationResponseGlobalSearchResultOpenApiVO**](OperationResponseGlobalSearchResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindActiveDeviceBySn

> OperationResponseWithoutResult UnbindActiveDeviceBySn(ctx, omadacId, siteId).UnbindActiveDeviceSnOpenApiVO(unbindActiveDeviceSnOpenApiVO).Execute()

Change the license from the old device to the new device



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
	unbindActiveDeviceSnOpenApiVO := *openapiclient.NewUnbindActiveDeviceSnOpenApiVO("NewSn_example", "OldSn_example") // UnbindActiveDeviceSnOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.UnbindActiveDeviceBySn(context.Background(), omadacId, siteId).UnbindActiveDeviceSnOpenApiVO(unbindActiveDeviceSnOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.UnbindActiveDeviceBySn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindActiveDeviceBySn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.UnbindActiveDeviceBySn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindActiveDeviceBySnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unbindActiveDeviceSnOpenApiVO** | [**UnbindActiveDeviceSnOpenApiVO**](UnbindActiveDeviceSnOpenApiVO.md) |  | 

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


## UnbindDevice

> []UnbindDeviceRespVO UnbindDevice(ctx, omadacId, siteId).UnbindDeviceOpenApiVO(unbindDeviceOpenApiVO).Execute()

Batch unbind devices



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
	unbindDeviceOpenApiVO := *openapiclient.NewUnbindDeviceOpenApiVO([]string{"DeviceMacs_example"}) // UnbindDeviceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.UnbindDevice(context.Background(), omadacId, siteId).UnbindDeviceOpenApiVO(unbindDeviceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.UnbindDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindDevice`: []UnbindDeviceRespVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.UnbindDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unbindDeviceOpenApiVO** | [**UnbindDeviceOpenApiVO**](UnbindDeviceOpenApiVO.md) |  | 

### Return type

[**[]UnbindDeviceRespVO**](UnbindDeviceRespVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindDeviceBySn

> []UnbindDeviceSnRespVO UnbindDeviceBySn(ctx, omadacId, siteId).UnbindDeviceSnOpenApiVO(unbindDeviceSnOpenApiVO).Execute()

Batch unbind devices by SN



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
	unbindDeviceSnOpenApiVO := *openapiclient.NewUnbindDeviceSnOpenApiVO([]string{"Sns_example"}) // UnbindDeviceSnOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.UnbindDeviceBySn(context.Background(), omadacId, siteId).UnbindDeviceSnOpenApiVO(unbindDeviceSnOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.UnbindDeviceBySn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindDeviceBySn`: []UnbindDeviceSnRespVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.UnbindDeviceBySn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindDeviceBySnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unbindDeviceSnOpenApiVO** | [**UnbindDeviceSnOpenApiVO**](UnbindDeviceSnOpenApiVO.md) |  | 

### Return type

[**[]UnbindDeviceSnRespVO**](UnbindDeviceSnRespVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceTag

> BatchTagRespOpenApiVO UpdateDeviceTag(ctx, omadacId, siteId).SetTagOpenApiVO(setTagOpenApiVO).Execute()

Set device tag for given devices



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
	setTagOpenApiVO := *openapiclient.NewSetTagOpenApiVO([]string{"TagIds_example"}) // SetTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.UpdateDeviceTag(context.Background(), omadacId, siteId).SetTagOpenApiVO(setTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.UpdateDeviceTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceTag`: BatchTagRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.UpdateDeviceTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setTagOpenApiVO** | [**SetTagOpenApiVO**](SetTagOpenApiVO.md) |  | 

### Return type

[**BatchTagRespOpenApiVO**](BatchTagRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadUpgradeFile

> OperationResponseFile UploadUpgradeFile(ctx, omadacId, siteId).File(file).Execute()

Upload device firmware



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
	file := os.NewFile(1234, "some_file") // *os.File | Upload a file in request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.UploadUpgradeFile(context.Background(), omadacId, siteId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.UploadUpgradeFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadUpgradeFile`: OperationResponseFile
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.UploadUpgradeFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadUpgradeFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | Upload a file in request body. | 

### Return type

[**OperationResponseFile**](OperationResponseFile.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


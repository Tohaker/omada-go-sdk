# NetworkAnalyzeAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadCaptureFile**](NetworkAnalyzeAPI.md#downloadcapturefile) | **Post** /openapi/v1/{omadacId}/files/sites/{siteId}/capture/device-type/{deviceType}/{deviceMac}/download | Download package capture file
[**GetAllCaptureDevices**](NetworkAnalyzeAPI.md#getallcapturedevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/capture/all-devices/device-type/{deviceType} | Get devices that can capture packages
[**GetDeviceInterfaces**](NetworkAnalyzeAPI.md#getdeviceinterfaces) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/capture/device-type/{deviceType}/{deviceMac}/interfaces | Get device interfaces
[**GetLastCaptureDevice**](NetworkAnalyzeAPI.md#getlastcapturedevice) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/capture/lastDevice | Get last capture device
[**QueryPackageCaptureStatus**](NetworkAnalyzeAPI.md#querypackagecapturestatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/capture/device-type/{deviceType}/{deviceMac}/{requestId}/status | Query package capture status
[**StartPackageCapture**](NetworkAnalyzeAPI.md#startpackagecapture) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/capture/device-type/{deviceType}/{deviceMac}/start | Start package capturing
[**StopPackageCapture**](NetworkAnalyzeAPI.md#stoppackagecapture) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/capture/device-type/{deviceType}/{deviceMac}/stop | Stop package capturing



## DownloadCaptureFile

> OperationResponseWithoutResult DownloadCaptureFile(ctx, omadacId, siteId, deviceType, deviceMac).DownloadCaptureFileConfig(downloadCaptureFileConfig).Execute()

Download package capture file



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
	deviceType := "deviceType_example" // string | Device type, and should be a value in [ap, switch, gateway].
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	downloadCaptureFileConfig := *openapiclient.NewDownloadCaptureFileConfig("RequestId_example") // DownloadCaptureFileConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAnalyzeAPI.DownloadCaptureFile(context.Background(), omadacId, siteId, deviceType, deviceMac).DownloadCaptureFileConfig(downloadCaptureFileConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAnalyzeAPI.DownloadCaptureFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadCaptureFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NetworkAnalyzeAPI.DownloadCaptureFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceType** | **string** | Device type, and should be a value in [ap, switch, gateway]. | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCaptureFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **downloadCaptureFileConfig** | [**DownloadCaptureFileConfig**](DownloadCaptureFileConfig.md) |  | 

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


## GetAllCaptureDevices

> OperationResponseGridVODeviceCaptureInfo GetAllCaptureDevices(ctx, omadacId, siteId, deviceType).Page(page).PageSize(pageSize).Execute()

Get devices that can capture packages



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
	deviceType := "deviceType_example" // string | Device type, and should be a value in [ap, switch, gateway].
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAnalyzeAPI.GetAllCaptureDevices(context.Background(), omadacId, siteId, deviceType).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAnalyzeAPI.GetAllCaptureDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCaptureDevices`: OperationResponseGridVODeviceCaptureInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkAnalyzeAPI.GetAllCaptureDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceType** | **string** | Device type, and should be a value in [ap, switch, gateway]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCaptureDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVODeviceCaptureInfo**](OperationResponseGridVODeviceCaptureInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceInterfaces

> OperationResponseDeviceInterfaces GetDeviceInterfaces(ctx, omadacId, siteId, deviceType, deviceMac).Execute()

Get device interfaces



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
	deviceType := "deviceType_example" // string | Device type, and should be a value in [ap, switch, gateway].
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAnalyzeAPI.GetDeviceInterfaces(context.Background(), omadacId, siteId, deviceType, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAnalyzeAPI.GetDeviceInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceInterfaces`: OperationResponseDeviceInterfaces
	fmt.Fprintf(os.Stdout, "Response from `NetworkAnalyzeAPI.GetDeviceInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceType** | **string** | Device type, and should be a value in [ap, switch, gateway]. | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseDeviceInterfaces**](OperationResponseDeviceInterfaces.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastCaptureDevice

> OperationResponseDeviceInterfaces GetLastCaptureDevice(ctx, omadacId, siteId).Execute()

Get last capture device



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
	resp, r, err := apiClient.NetworkAnalyzeAPI.GetLastCaptureDevice(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAnalyzeAPI.GetLastCaptureDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLastCaptureDevice`: OperationResponseDeviceInterfaces
	fmt.Fprintf(os.Stdout, "Response from `NetworkAnalyzeAPI.GetLastCaptureDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastCaptureDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDeviceInterfaces**](OperationResponseDeviceInterfaces.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryPackageCaptureStatus

> OperationResponseCaptureResult QueryPackageCaptureStatus(ctx, omadacId, siteId, deviceType, deviceMac, requestId).Execute()

Query package capture status



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
	deviceType := "deviceType_example" // string | Device type, and should be a value in [ap, switch, gateway].
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	requestId := "requestId_example" // string | A GUID based on the timestamp.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAnalyzeAPI.QueryPackageCaptureStatus(context.Background(), omadacId, siteId, deviceType, deviceMac, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAnalyzeAPI.QueryPackageCaptureStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPackageCaptureStatus`: OperationResponseCaptureResult
	fmt.Fprintf(os.Stdout, "Response from `NetworkAnalyzeAPI.QueryPackageCaptureStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceType** | **string** | Device type, and should be a value in [ap, switch, gateway]. | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**requestId** | **string** | A GUID based on the timestamp. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryPackageCaptureStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**OperationResponseCaptureResult**](OperationResponseCaptureResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPackageCapture

> OperationResponseCaptureInfo StartPackageCapture(ctx, omadacId, siteId, deviceType, deviceMac).PackageCaptureConfig(packageCaptureConfig).Execute()

Start package capturing



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
	deviceType := "deviceType_example" // string | Device type, and should be a value in [ap, switch, gateway].
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	packageCaptureConfig := *openapiclient.NewPackageCaptureConfig() // PackageCaptureConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAnalyzeAPI.StartPackageCapture(context.Background(), omadacId, siteId, deviceType, deviceMac).PackageCaptureConfig(packageCaptureConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAnalyzeAPI.StartPackageCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartPackageCapture`: OperationResponseCaptureInfo
	fmt.Fprintf(os.Stdout, "Response from `NetworkAnalyzeAPI.StartPackageCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceType** | **string** | Device type, and should be a value in [ap, switch, gateway]. | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartPackageCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **packageCaptureConfig** | [**PackageCaptureConfig**](PackageCaptureConfig.md) |  | 

### Return type

[**OperationResponseCaptureInfo**](OperationResponseCaptureInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopPackageCapture

> OperationResponseWithoutResult StopPackageCapture(ctx, omadacId, siteId, deviceType, deviceMac).Execute()

Stop package capturing



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
	deviceType := "deviceType_example" // string | Device type, and should be a value in [ap, switch, gateway].
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAnalyzeAPI.StopPackageCapture(context.Background(), omadacId, siteId, deviceType, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAnalyzeAPI.StopPackageCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopPackageCapture`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NetworkAnalyzeAPI.StopPackageCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceType** | **string** | Device type, and should be a value in [ap, switch, gateway]. | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopPackageCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


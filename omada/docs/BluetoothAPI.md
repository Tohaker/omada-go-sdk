# \BluetoothAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteBtIbeaconConfig**](BluetoothAPI.md#BatchDeleteBtIbeaconConfig) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/devices/config/batch | Batch delete Bluetooth Advertising
[**BatchDeleteBtIbeaconConfigV2**](BluetoothAPI.md#BatchDeleteBtIbeaconConfigV2) | **Delete** /openapi/v2/{omadacId}/sites/{siteId}/setting/iot/devices/config/batch | Batch delete Bluetooth Advertising v2
[**BatchDeleteIotServer**](BluetoothAPI.md#BatchDeleteIotServer) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/servers/batch-delete | Batch Delete IoT Transport Stream
[**BatchModifyBtIbeaconConfig**](BluetoothAPI.md#BatchModifyBtIbeaconConfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/devices/config/batch-config | Batch modify Bluetooth Advertising
[**BatchModifyBtIbeaconConfigV2**](BluetoothAPI.md#BatchModifyBtIbeaconConfigV2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/setting/iot/devices/config/batch-config | Batch modify Bluetooth Advertising
[**BatchModifyIotDeviceTransmitPower**](BluetoothAPI.md#BatchModifyIotDeviceTransmitPower) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/radio/transmit-power | Batch Modify transmit power of IoT devices.
[**CreateBtIbeaconConfig**](BluetoothAPI.md#CreateBtIbeaconConfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/devices/config | Create Bluetooth Advertising
[**CreateBtIbeaconConfigV2**](BluetoothAPI.md#CreateBtIbeaconConfigV2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/setting/iot/devices/config | Create Bluetooth Advertising v2
[**CreateIotServer**](BluetoothAPI.md#CreateIotServer) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/servers | Create IoT Transport Stream
[**DeleteBtIbeaconConfig**](BluetoothAPI.md#DeleteBtIbeaconConfig) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/devices/config/{id} | Delete Bluetooth Advertising
[**DeleteBtIbeaconConfigV2**](BluetoothAPI.md#DeleteBtIbeaconConfigV2) | **Delete** /openapi/v2/{omadacId}/sites/{siteId}/setting/iot/devices/config/{id} | Delete Bluetooth Advertising v2
[**DeleteIotServer**](BluetoothAPI.md#DeleteIotServer) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/servers/{id} | Delete IoT Transport Stream
[**GetGridBtIbeaconDevices**](BluetoothAPI.md#GetGridBtIbeaconDevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/devices | Get Bluetooth Advertising Devices
[**GetGridBtIbeaconList**](BluetoothAPI.md#GetGridBtIbeaconList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/devices/config | Get Bluetooth Advertising
[**GetGridBtIbeaconListV2**](BluetoothAPI.md#GetGridBtIbeaconListV2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/setting/iot/devices/config | Get Bluetooth Advertising v2
[**GetGridIotServer**](BluetoothAPI.md#GetGridIotServer) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/servers | Get IoT Transport Steam
[**GetIotBlueToothAgingTime**](BluetoothAPI.md#GetIotBlueToothAgingTime) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/aging-time | Get Aging Time
[**GetIotDeviceTransmitPowerInfoList**](BluetoothAPI.md#GetIotDeviceTransmitPowerInfoList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/radio/transmit-power | Get IoT devices with transmit power
[**GetIotRadioSetting**](BluetoothAPI.md#GetIotRadioSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/radio | Get Iot Radio Setting
[**ModifyBtIbeaconConfig**](BluetoothAPI.md#ModifyBtIbeaconConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/devices/config/{id} | Modify Bluetooth Advertising
[**ModifyBtIbeaconConfigV2**](BluetoothAPI.md#ModifyBtIbeaconConfigV2) | **Put** /openapi/v2/{omadacId}/sites/{siteId}/setting/iot/devices/config/{id} | Modify Bluetooth Advertising v2
[**ModifyIotBlueToothAgingTime**](BluetoothAPI.md#ModifyIotBlueToothAgingTime) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/aging-time | Modify Aging Time
[**ModifyIotRadioSetting**](BluetoothAPI.md#ModifyIotRadioSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/radio | Modify Iot Radio Setting
[**ModifyIotServer**](BluetoothAPI.md#ModifyIotServer) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/iot/servers/{id} | Modify IoT Transport Steam



## BatchDeleteBtIbeaconConfig

> OperationResponseWithoutResult BatchDeleteBtIbeaconConfig(ctx, omadacId, siteId).BatchDeleteCommonOpenApiVO(batchDeleteCommonOpenApiVO).Execute()

Batch delete Bluetooth Advertising



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
	batchDeleteCommonOpenApiVO := *openapiclient.NewBatchDeleteCommonOpenApiVO() // BatchDeleteCommonOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.BatchDeleteBtIbeaconConfig(context.Background(), omadacId, siteId).BatchDeleteCommonOpenApiVO(batchDeleteCommonOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.BatchDeleteBtIbeaconConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteBtIbeaconConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.BatchDeleteBtIbeaconConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteBtIbeaconConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchDeleteCommonOpenApiVO** | [**BatchDeleteCommonOpenApiVO**](BatchDeleteCommonOpenApiVO.md) |  | 

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


## BatchDeleteBtIbeaconConfigV2

> OperationResponseWithoutResult BatchDeleteBtIbeaconConfigV2(ctx, omadacId, siteId).BatchDeleteCommonOpenApiVO(batchDeleteCommonOpenApiVO).Execute()

Batch delete Bluetooth Advertising v2



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
	batchDeleteCommonOpenApiVO := *openapiclient.NewBatchDeleteCommonOpenApiVO() // BatchDeleteCommonOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.BatchDeleteBtIbeaconConfigV2(context.Background(), omadacId, siteId).BatchDeleteCommonOpenApiVO(batchDeleteCommonOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.BatchDeleteBtIbeaconConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteBtIbeaconConfigV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.BatchDeleteBtIbeaconConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteBtIbeaconConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchDeleteCommonOpenApiVO** | [**BatchDeleteCommonOpenApiVO**](BatchDeleteCommonOpenApiVO.md) |  | 

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


## BatchDeleteIotServer

> OperationResponseWithoutResult BatchDeleteIotServer(ctx, omadacId, siteId).BatchDeleteCommonOpenApiVO(batchDeleteCommonOpenApiVO).Execute()

Batch Delete IoT Transport Stream



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
	batchDeleteCommonOpenApiVO := *openapiclient.NewBatchDeleteCommonOpenApiVO() // BatchDeleteCommonOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.BatchDeleteIotServer(context.Background(), omadacId, siteId).BatchDeleteCommonOpenApiVO(batchDeleteCommonOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.BatchDeleteIotServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteIotServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.BatchDeleteIotServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteIotServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchDeleteCommonOpenApiVO** | [**BatchDeleteCommonOpenApiVO**](BatchDeleteCommonOpenApiVO.md) |  | 

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


## BatchModifyBtIbeaconConfig

> OperationResponseWithoutResult BatchModifyBtIbeaconConfig(ctx, omadacId, siteId).BatchModifyIbeaconOpenApiVO(batchModifyIbeaconOpenApiVO).Execute()

Batch modify Bluetooth Advertising



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
	batchModifyIbeaconOpenApiVO := *openapiclient.NewBatchModifyIbeaconOpenApiVO() // BatchModifyIbeaconOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.BatchModifyBtIbeaconConfig(context.Background(), omadacId, siteId).BatchModifyIbeaconOpenApiVO(batchModifyIbeaconOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.BatchModifyBtIbeaconConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyBtIbeaconConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.BatchModifyBtIbeaconConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyBtIbeaconConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchModifyIbeaconOpenApiVO** | [**BatchModifyIbeaconOpenApiVO**](BatchModifyIbeaconOpenApiVO.md) |  | 

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


## BatchModifyBtIbeaconConfigV2

> OperationResponseWithoutResult BatchModifyBtIbeaconConfigV2(ctx, omadacId, siteId).BatchModifyIbeaconV2OpenApiVO(batchModifyIbeaconV2OpenApiVO).Execute()

Batch modify Bluetooth Advertising



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
	batchModifyIbeaconV2OpenApiVO := *openapiclient.NewBatchModifyIbeaconV2OpenApiVO() // BatchModifyIbeaconV2OpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.BatchModifyBtIbeaconConfigV2(context.Background(), omadacId, siteId).BatchModifyIbeaconV2OpenApiVO(batchModifyIbeaconV2OpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.BatchModifyBtIbeaconConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyBtIbeaconConfigV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.BatchModifyBtIbeaconConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyBtIbeaconConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchModifyIbeaconV2OpenApiVO** | [**BatchModifyIbeaconV2OpenApiVO**](BatchModifyIbeaconV2OpenApiVO.md) |  | 

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


## BatchModifyIotDeviceTransmitPower

> OperationResponseWithoutResult BatchModifyIotDeviceTransmitPower(ctx, omadacId, siteId).BatchConfigIotTransmitPowerOpenApiVO(batchConfigIotTransmitPowerOpenApiVO).Execute()

Batch Modify transmit power of IoT devices.



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
	batchConfigIotTransmitPowerOpenApiVO := *openapiclient.NewBatchConfigIotTransmitPowerOpenApiVO([]string{"Macs_example"}, false, false) // BatchConfigIotTransmitPowerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.BatchModifyIotDeviceTransmitPower(context.Background(), omadacId, siteId).BatchConfigIotTransmitPowerOpenApiVO(batchConfigIotTransmitPowerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.BatchModifyIotDeviceTransmitPower``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyIotDeviceTransmitPower`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.BatchModifyIotDeviceTransmitPower`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyIotDeviceTransmitPowerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchConfigIotTransmitPowerOpenApiVO** | [**BatchConfigIotTransmitPowerOpenApiVO**](BatchConfigIotTransmitPowerOpenApiVO.md) |  | 

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


## CreateBtIbeaconConfig

> OperationResponseWithoutResult CreateBtIbeaconConfig(ctx, omadacId, siteId).ConfigIotBtIbeaconOpenApiVO(configIotBtIbeaconOpenApiVO).Execute()

Create Bluetooth Advertising



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
	configIotBtIbeaconOpenApiVO := *openapiclient.NewConfigIotBtIbeaconOpenApiVO([]string{"MacList_example"}, "Major_example", "Minor_example", "Name_example", "Uuid_example") // ConfigIotBtIbeaconOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.CreateBtIbeaconConfig(context.Background(), omadacId, siteId).ConfigIotBtIbeaconOpenApiVO(configIotBtIbeaconOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.CreateBtIbeaconConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBtIbeaconConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.CreateBtIbeaconConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBtIbeaconConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configIotBtIbeaconOpenApiVO** | [**ConfigIotBtIbeaconOpenApiVO**](ConfigIotBtIbeaconOpenApiVO.md) |  | 

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


## CreateBtIbeaconConfigV2

> OperationResponseWithoutResult CreateBtIbeaconConfigV2(ctx, omadacId, siteId).ConfigIotBtIbeaconV2OpenApiVO(configIotBtIbeaconV2OpenApiVO).Execute()

Create Bluetooth Advertising v2



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
	configIotBtIbeaconV2OpenApiVO := *openapiclient.NewConfigIotBtIbeaconV2OpenApiVO([]string{"MacList_example"}, "Major_example", "Minor_example", "Name_example", "Uuid_example") // ConfigIotBtIbeaconV2OpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.CreateBtIbeaconConfigV2(context.Background(), omadacId, siteId).ConfigIotBtIbeaconV2OpenApiVO(configIotBtIbeaconV2OpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.CreateBtIbeaconConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBtIbeaconConfigV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.CreateBtIbeaconConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBtIbeaconConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configIotBtIbeaconV2OpenApiVO** | [**ConfigIotBtIbeaconV2OpenApiVO**](ConfigIotBtIbeaconV2OpenApiVO.md) |  | 

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


## CreateIotServer

> OperationResponseWithoutResult CreateIotServer(ctx, omadacId, siteId).ConfigIotServerOpenApiVO(configIotServerOpenApiVO).Execute()

Create IoT Transport Stream



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
	configIotServerOpenApiVO := *openapiclient.NewConfigIotServerOpenApiVO("AccessToken_example", int32(123), false, "ClientId_example", []int32{int32(123)}, false, "Name_example", false, int32(123), int32(123), "ServerUrl_example") // ConfigIotServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.CreateIotServer(context.Background(), omadacId, siteId).ConfigIotServerOpenApiVO(configIotServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.CreateIotServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIotServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.CreateIotServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIotServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configIotServerOpenApiVO** | [**ConfigIotServerOpenApiVO**](ConfigIotServerOpenApiVO.md) |  | 

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


## DeleteBtIbeaconConfig

> OperationResponseWithoutResult DeleteBtIbeaconConfig(ctx, omadacId, siteId, id).Execute()

Delete Bluetooth Advertising



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
	id := "id_example" // string | Bluetooth Advertising entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.DeleteBtIbeaconConfig(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.DeleteBtIbeaconConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBtIbeaconConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.DeleteBtIbeaconConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Bluetooth Advertising entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBtIbeaconConfigRequest struct via the builder pattern


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


## DeleteBtIbeaconConfigV2

> OperationResponseWithoutResult DeleteBtIbeaconConfigV2(ctx, omadacId, siteId, id).Execute()

Delete Bluetooth Advertising v2



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
	id := "id_example" // string | Bluetooth Advertising entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.DeleteBtIbeaconConfigV2(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.DeleteBtIbeaconConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBtIbeaconConfigV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.DeleteBtIbeaconConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Bluetooth Advertising entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBtIbeaconConfigV2Request struct via the builder pattern


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


## DeleteIotServer

> OperationResponseWithoutResult DeleteIotServer(ctx, omadacId, siteId, id).Execute()

Delete IoT Transport Stream



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
	id := "id_example" // string | IoT Transport Stream entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.DeleteIotServer(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.DeleteIotServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIotServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.DeleteIotServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | IoT Transport Stream entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIotServerRequest struct via the builder pattern


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


## GetGridBtIbeaconDevices

> GridVOApBtDetailOpenApiVO GetGridBtIbeaconDevices(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Bluetooth Advertising Devices



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
	resp, r, err := apiClient.BluetoothAPI.GetGridBtIbeaconDevices(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetGridBtIbeaconDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridBtIbeaconDevices`: GridVOApBtDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetGridBtIbeaconDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridBtIbeaconDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOApBtDetailOpenApiVO**](GridVOApBtDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridBtIbeaconList

> GridVOIotBtIbeaconOpenApiVO GetGridBtIbeaconList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Bluetooth Advertising



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
	resp, r, err := apiClient.BluetoothAPI.GetGridBtIbeaconList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetGridBtIbeaconList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridBtIbeaconList`: GridVOIotBtIbeaconOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetGridBtIbeaconList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridBtIbeaconListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOIotBtIbeaconOpenApiVO**](GridVOIotBtIbeaconOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridBtIbeaconListV2

> GridVOIotBtIbeaconOpenApiVO GetGridBtIbeaconListV2(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Bluetooth Advertising v2



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
	resp, r, err := apiClient.BluetoothAPI.GetGridBtIbeaconListV2(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetGridBtIbeaconListV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridBtIbeaconListV2`: GridVOIotBtIbeaconOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetGridBtIbeaconListV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridBtIbeaconListV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOIotBtIbeaconOpenApiVO**](GridVOIotBtIbeaconOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridIotServer

> GridVOIotServerOpenApiVO GetGridIotServer(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get IoT Transport Steam



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
	resp, r, err := apiClient.BluetoothAPI.GetGridIotServer(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetGridIotServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridIotServer`: GridVOIotServerOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetGridIotServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridIotServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOIotServerOpenApiVO**](GridVOIotServerOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIotBlueToothAgingTime

> IotAgingTimeOpenApiVO GetIotBlueToothAgingTime(ctx, omadacId, siteId).Execute()

Get Aging Time



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
	resp, r, err := apiClient.BluetoothAPI.GetIotBlueToothAgingTime(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetIotBlueToothAgingTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIotBlueToothAgingTime`: IotAgingTimeOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetIotBlueToothAgingTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIotBlueToothAgingTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IotAgingTimeOpenApiVO**](IotAgingTimeOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIotDeviceTransmitPowerInfoList

> GridVOIotDeviceRadioInfoOpenApiVO GetIotDeviceTransmitPowerInfoList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get IoT devices with transmit power



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field Device Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.GetIotDeviceTransmitPowerInfoList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetIotDeviceTransmitPowerInfoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIotDeviceTransmitPowerInfoList`: GridVOIotDeviceRadioInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetIotDeviceTransmitPowerInfoList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIotDeviceTransmitPowerInfoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field Device Name | 

### Return type

[**GridVOIotDeviceRadioInfoOpenApiVO**](GridVOIotDeviceRadioInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIotRadioSetting

> IotRadioSettingOpenApiVO GetIotRadioSetting(ctx, omadacId, siteId).Execute()

Get Iot Radio Setting



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
	resp, r, err := apiClient.BluetoothAPI.GetIotRadioSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetIotRadioSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIotRadioSetting`: IotRadioSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetIotRadioSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIotRadioSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IotRadioSettingOpenApiVO**](IotRadioSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyBtIbeaconConfig

> OperationResponseWithoutResult ModifyBtIbeaconConfig(ctx, omadacId, siteId, id).ConfigIotBtIbeaconOpenApiVO(configIotBtIbeaconOpenApiVO).Execute()

Modify Bluetooth Advertising



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
	id := "id_example" // string | Bluetooth Advertising entry ID
	configIotBtIbeaconOpenApiVO := *openapiclient.NewConfigIotBtIbeaconOpenApiVO([]string{"MacList_example"}, "Major_example", "Minor_example", "Name_example", "Uuid_example") // ConfigIotBtIbeaconOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.ModifyBtIbeaconConfig(context.Background(), omadacId, siteId, id).ConfigIotBtIbeaconOpenApiVO(configIotBtIbeaconOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.ModifyBtIbeaconConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBtIbeaconConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.ModifyBtIbeaconConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Bluetooth Advertising entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBtIbeaconConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **configIotBtIbeaconOpenApiVO** | [**ConfigIotBtIbeaconOpenApiVO**](ConfigIotBtIbeaconOpenApiVO.md) |  | 

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


## ModifyBtIbeaconConfigV2

> OperationResponseWithoutResult ModifyBtIbeaconConfigV2(ctx, omadacId, siteId, id).ConfigIotBtIbeaconV2OpenApiVO(configIotBtIbeaconV2OpenApiVO).Execute()

Modify Bluetooth Advertising v2



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
	id := "id_example" // string | Bluetooth Advertising entry ID
	configIotBtIbeaconV2OpenApiVO := *openapiclient.NewConfigIotBtIbeaconV2OpenApiVO([]string{"MacList_example"}, "Major_example", "Minor_example", "Name_example", "Uuid_example") // ConfigIotBtIbeaconV2OpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.ModifyBtIbeaconConfigV2(context.Background(), omadacId, siteId, id).ConfigIotBtIbeaconV2OpenApiVO(configIotBtIbeaconV2OpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.ModifyBtIbeaconConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBtIbeaconConfigV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.ModifyBtIbeaconConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Bluetooth Advertising entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBtIbeaconConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **configIotBtIbeaconV2OpenApiVO** | [**ConfigIotBtIbeaconV2OpenApiVO**](ConfigIotBtIbeaconV2OpenApiVO.md) |  | 

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


## ModifyIotBlueToothAgingTime

> OperationResponseWithoutResult ModifyIotBlueToothAgingTime(ctx, omadacId, siteId).IotAgingTimeOpenApiVO(iotAgingTimeOpenApiVO).Execute()

Modify Aging Time



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
	iotAgingTimeOpenApiVO := *openapiclient.NewIotAgingTimeOpenApiVO(int32(123), int32(123)) // IotAgingTimeOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.ModifyIotBlueToothAgingTime(context.Background(), omadacId, siteId).IotAgingTimeOpenApiVO(iotAgingTimeOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.ModifyIotBlueToothAgingTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIotBlueToothAgingTime`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.ModifyIotBlueToothAgingTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIotBlueToothAgingTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iotAgingTimeOpenApiVO** | [**IotAgingTimeOpenApiVO**](IotAgingTimeOpenApiVO.md) |  | 

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


## ModifyIotRadioSetting

> OperationResponseWithoutResult ModifyIotRadioSetting(ctx, omadacId, siteId).IotRadioSettingOpenApiVO(iotRadioSettingOpenApiVO).Execute()

Modify Iot Radio Setting



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
	iotRadioSettingOpenApiVO := *openapiclient.NewIotRadioSettingOpenApiVO(false) // IotRadioSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.ModifyIotRadioSetting(context.Background(), omadacId, siteId).IotRadioSettingOpenApiVO(iotRadioSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.ModifyIotRadioSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIotRadioSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.ModifyIotRadioSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIotRadioSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iotRadioSettingOpenApiVO** | [**IotRadioSettingOpenApiVO**](IotRadioSettingOpenApiVO.md) |  | 

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


## ModifyIotServer

> OperationResponseWithoutResult ModifyIotServer(ctx, omadacId, siteId, id).ConfigIotServerOpenApiVO(configIotServerOpenApiVO).Execute()

Modify IoT Transport Steam



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
	id := "id_example" // string | IoT Transport Stream entry ID
	configIotServerOpenApiVO := *openapiclient.NewConfigIotServerOpenApiVO("AccessToken_example", int32(123), false, "ClientId_example", []int32{int32(123)}, false, "Name_example", false, int32(123), int32(123), "ServerUrl_example") // ConfigIotServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.ModifyIotServer(context.Background(), omadacId, siteId, id).ConfigIotServerOpenApiVO(configIotServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.ModifyIotServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIotServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.ModifyIotServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | IoT Transport Stream entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIotServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **configIotServerOpenApiVO** | [**ConfigIotServerOpenApiVO**](ConfigIotServerOpenApiVO.md) |  | 

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


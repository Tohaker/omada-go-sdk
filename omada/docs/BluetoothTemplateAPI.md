# BluetoothTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteIotServerTemplate**](BluetoothTemplateAPI.md#batchdeleteiotservertemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/servers/batch-delete | Batch delete IoT Transport Stream(Template)
[**CreateIotServerTemplate**](BluetoothTemplateAPI.md#createiotservertemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/servers | Create IoT Transport Stream(Template)
[**DeleteIotServerTemplate**](BluetoothTemplateAPI.md#deleteiotservertemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/servers/{id} | Delete IoT Transport Stream(Template)
[**GetGridIotServerTemplate**](BluetoothTemplateAPI.md#getgridiotservertemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/servers | Get IoT Transport Stream(Template)
[**GetIotBlueToothAgingTimeTemplate**](BluetoothTemplateAPI.md#getiotbluetoothagingtimetemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/aging-time | Get Aging Time(Template)
[**GetIotRadioSettingTemplate**](BluetoothTemplateAPI.md#getiotradiosettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/radio | Get Iot Radio Setting(Template)
[**ModifyIotBlueToothAgingTimeTemplate**](BluetoothTemplateAPI.md#modifyiotbluetoothagingtimetemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/aging-time | Modify Aging Time(Template)
[**ModifyIotRadioSettingTemplate**](BluetoothTemplateAPI.md#modifyiotradiosettingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/radio | Modify Iot Radio Setting(Template)
[**ModifyIotServerTemplate**](BluetoothTemplateAPI.md#modifyiotservertemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/iot/servers/{id} | Modify IoT Transport Stream(Template)



## BatchDeleteIotServerTemplate

> OperationResponseWithoutResult BatchDeleteIotServerTemplate(ctx, omadacId, siteTemplateId).BatchDeleteCommonOpenApiVO(batchDeleteCommonOpenApiVO).Execute()

Batch delete IoT Transport Stream(Template)



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
	batchDeleteCommonOpenApiVO := *openapiclient.NewBatchDeleteCommonOpenApiVO() // BatchDeleteCommonOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothTemplateAPI.BatchDeleteIotServerTemplate(context.Background(), omadacId, siteTemplateId).BatchDeleteCommonOpenApiVO(batchDeleteCommonOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.BatchDeleteIotServerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteIotServerTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.BatchDeleteIotServerTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteIotServerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchDeleteCommonOpenApiVO** | [**BatchDeleteCommonOpenApiVO**](BatchDeleteCommonOpenApiVO.md) |  | 

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


## CreateIotServerTemplate

> OperationResponseWithoutResult CreateIotServerTemplate(ctx, omadacId, siteTemplateId).ConfigIotServerOpenApiVO(configIotServerOpenApiVO).Execute()

Create IoT Transport Stream(Template)



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
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId
	configIotServerOpenApiVO := *openapiclient.NewConfigIotServerOpenApiVO("AccessToken_example", int32(123), false, "ClientId_example", []int32{int32(123)}, false, "Name_example", false, int32(123), int32(123), "ServerUrl_example") // ConfigIotServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothTemplateAPI.CreateIotServerTemplate(context.Background(), omadacId, siteTemplateId).ConfigIotServerOpenApiVO(configIotServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.CreateIotServerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIotServerTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.CreateIotServerTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | siteTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIotServerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configIotServerOpenApiVO** | [**ConfigIotServerOpenApiVO**](ConfigIotServerOpenApiVO.md) |  | 

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


## DeleteIotServerTemplate

> OperationResponseWithoutResult DeleteIotServerTemplate(ctx, omadacId, siteTemplateId, id).Execute()

Delete IoT Transport Stream(Template)



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
	id := "id_example" // string | IoT Transport Stream template entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothTemplateAPI.DeleteIotServerTemplate(context.Background(), omadacId, siteTemplateId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.DeleteIotServerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIotServerTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.DeleteIotServerTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | IoT Transport Stream template entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIotServerTemplateRequest struct via the builder pattern


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


## GetGridIotServerTemplate

> GridVOIotServerOpenApiVO GetGridIotServerTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get IoT Transport Stream(Template)



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothTemplateAPI.GetGridIotServerTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.GetGridIotServerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridIotServerTemplate`: GridVOIotServerOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.GetGridIotServerTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridIotServerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOIotServerOpenApiVO**](GridVOIotServerOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIotBlueToothAgingTimeTemplate

> IotAgingTimeOpenApiVO GetIotBlueToothAgingTimeTemplate(ctx, omadacId, siteTemplateId).Execute()

Get Aging Time(Template)



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
	resp, r, err := apiClient.BluetoothTemplateAPI.GetIotBlueToothAgingTimeTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.GetIotBlueToothAgingTimeTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIotBlueToothAgingTimeTemplate`: IotAgingTimeOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.GetIotBlueToothAgingTimeTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIotBlueToothAgingTimeTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IotAgingTimeOpenApiVO**](IotAgingTimeOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIotRadioSettingTemplate

> IotRadioSettingOpenApiVO GetIotRadioSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get Iot Radio Setting(Template)



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
	resp, r, err := apiClient.BluetoothTemplateAPI.GetIotRadioSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.GetIotRadioSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIotRadioSettingTemplate`: IotRadioSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.GetIotRadioSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIotRadioSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IotRadioSettingOpenApiVO**](IotRadioSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIotBlueToothAgingTimeTemplate

> OperationResponseWithoutResult ModifyIotBlueToothAgingTimeTemplate(ctx, omadacId, siteTemplateId).IotAgingTimeOpenApiVO(iotAgingTimeOpenApiVO).Execute()

Modify Aging Time(Template)



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
	iotAgingTimeOpenApiVO := *openapiclient.NewIotAgingTimeOpenApiVO(int32(123), int32(123)) // IotAgingTimeOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothTemplateAPI.ModifyIotBlueToothAgingTimeTemplate(context.Background(), omadacId, siteTemplateId).IotAgingTimeOpenApiVO(iotAgingTimeOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.ModifyIotBlueToothAgingTimeTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIotBlueToothAgingTimeTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.ModifyIotBlueToothAgingTimeTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIotBlueToothAgingTimeTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iotAgingTimeOpenApiVO** | [**IotAgingTimeOpenApiVO**](IotAgingTimeOpenApiVO.md) |  | 

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


## ModifyIotRadioSettingTemplate

> OperationResponseWithoutResult ModifyIotRadioSettingTemplate(ctx, omadacId, siteTemplateId).IotRadioSettingOpenApiVO(iotRadioSettingOpenApiVO).Execute()

Modify Iot Radio Setting(Template)



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
	iotRadioSettingOpenApiVO := *openapiclient.NewIotRadioSettingOpenApiVO(false) // IotRadioSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothTemplateAPI.ModifyIotRadioSettingTemplate(context.Background(), omadacId, siteTemplateId).IotRadioSettingOpenApiVO(iotRadioSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.ModifyIotRadioSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIotRadioSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.ModifyIotRadioSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIotRadioSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iotRadioSettingOpenApiVO** | [**IotRadioSettingOpenApiVO**](IotRadioSettingOpenApiVO.md) |  | 

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


## ModifyIotServerTemplate

> OperationResponseWithoutResult ModifyIotServerTemplate(ctx, omadacId, siteTemplateId, id).ConfigIotServerOpenApiVO(configIotServerOpenApiVO).Execute()

Modify IoT Transport Stream(Template)



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
	id := "id_example" // string | IoT Transport Stream template entry ID
	configIotServerOpenApiVO := *openapiclient.NewConfigIotServerOpenApiVO("AccessToken_example", int32(123), false, "ClientId_example", []int32{int32(123)}, false, "Name_example", false, int32(123), int32(123), "ServerUrl_example") // ConfigIotServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothTemplateAPI.ModifyIotServerTemplate(context.Background(), omadacId, siteTemplateId, id).ConfigIotServerOpenApiVO(configIotServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothTemplateAPI.ModifyIotServerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIotServerTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BluetoothTemplateAPI.ModifyIotServerTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | IoT Transport Stream template entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIotServerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **configIotServerOpenApiVO** | [**ConfigIotServerOpenApiVO**](ConfigIotServerOpenApiVO.md) |  | 

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


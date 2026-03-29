# \OLTTContAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTcont**](OLTTContAPI.md#AddTcont) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/t-conts/add | Create new t-cont
[**DeleteTcont**](OLTTContAPI.md#DeleteTcont) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/t-conts/delete | Delete an existing t-cont
[**EditTcont**](OLTTContAPI.md#EditTcont) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/t-conts/edit | Modify an existing t-cont
[**GetTcontList**](OLTTContAPI.md#GetTcontList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/t-conts/list | Get t-cont list
[**GetTcontPage**](OLTTContAPI.md#GetTcontPage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/t-conts/page | Get t-cont page



## AddTcont

> OperationResponseDeviceResponseBodyVoid AddTcont(ctx, omadacId, siteId, deviceMac, lineProfileId).TcontDTO(tcontDTO).Execute()

Create new t-cont



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
	lineProfileId := "lineProfileId_example" // string | Line Profile ID should be a number between 0-512
	tcontDTO := *openapiclient.NewTcontDTO(int32(123), int32(123)) // TcontDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTContAPI.AddTcont(context.Background(), omadacId, siteId, deviceMac, lineProfileId).TcontDTO(tcontDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTContAPI.AddTcont``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTcont`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTTContAPI.AddTcont`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**lineProfileId** | **string** | Line Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTcontRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tcontDTO** | [**TcontDTO**](TcontDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTcont

> OperationResponseDeviceResponseBodyTcontDeleteResultDTO DeleteTcont(ctx, omadacId, siteId, deviceMac, lineProfileId).TcontDeleteDTO(tcontDeleteDTO).Execute()

Delete an existing t-cont



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
	lineProfileId := "lineProfileId_example" // string | Line Profile ID should be a number between 0-512
	tcontDeleteDTO := *openapiclient.NewTcontDeleteDTO([]int32{int32(123)}) // TcontDeleteDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTContAPI.DeleteTcont(context.Background(), omadacId, siteId, deviceMac, lineProfileId).TcontDeleteDTO(tcontDeleteDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTContAPI.DeleteTcont``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTcont`: OperationResponseDeviceResponseBodyTcontDeleteResultDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTTContAPI.DeleteTcont`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**lineProfileId** | **string** | Line Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTcontRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tcontDeleteDTO** | [**TcontDeleteDTO**](TcontDeleteDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyTcontDeleteResultDTO**](OperationResponseDeviceResponseBodyTcontDeleteResultDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTcont

> OperationResponseDeviceResponseBodyVoid EditTcont(ctx, omadacId, siteId, deviceMac, lineProfileId).TcontModifyDTO(tcontModifyDTO).Execute()

Modify an existing t-cont



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
	lineProfileId := "lineProfileId_example" // string | Line Profile ID should be a number between 0-512
	tcontModifyDTO := *openapiclient.NewTcontModifyDTO(int32(123)) // TcontModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTContAPI.EditTcont(context.Background(), omadacId, siteId, deviceMac, lineProfileId).TcontModifyDTO(tcontModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTContAPI.EditTcont``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTcont`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTTContAPI.EditTcont`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**lineProfileId** | **string** | Line Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTcontRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tcontModifyDTO** | [**TcontModifyDTO**](TcontModifyDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTcontList

> OperationResponseListTcontVO GetTcontList(ctx, omadacId, siteId, deviceMac, lineProfileId).Dto(dto).Execute()

Get t-cont list



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
	lineProfileId := "lineProfileId_example" // string | Line Profile ID should be a number between 0-512
	dto := *openapiclient.NewTcontListQueryDTO() // TcontListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTContAPI.GetTcontList(context.Background(), omadacId, siteId, deviceMac, lineProfileId).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTContAPI.GetTcontList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTcontList`: OperationResponseListTcontVO
	fmt.Fprintf(os.Stdout, "Response from `OLTTContAPI.GetTcontList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**lineProfileId** | **string** | Line Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcontListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dto** | [**TcontListQueryDTO**](TcontListQueryDTO.md) |  | 

### Return type

[**OperationResponseListTcontVO**](OperationResponseListTcontVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTcontPage

> OperationResponsePageResponseTcontVO GetTcontPage(ctx, omadacId, siteId, deviceMac, lineProfileId).Dto(dto).Execute()

Get t-cont page



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
	lineProfileId := "lineProfileId_example" // string | Line Profile ID should be a number between 0-512
	dto := *openapiclient.NewTcontPageQueryDTO() // TcontPageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTContAPI.GetTcontPage(context.Background(), omadacId, siteId, deviceMac, lineProfileId).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTContAPI.GetTcontPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTcontPage`: OperationResponsePageResponseTcontVO
	fmt.Fprintf(os.Stdout, "Response from `OLTTContAPI.GetTcontPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**lineProfileId** | **string** | Line Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcontPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dto** | [**TcontPageQueryDTO**](TcontPageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseTcontVO**](OperationResponsePageResponseTcontVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


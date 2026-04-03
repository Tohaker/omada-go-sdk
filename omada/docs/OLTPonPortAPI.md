# OLTPonPortAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditPonPort**](OLTPonPortAPI.md#editponport) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/pon-port/edit | Batch modify pon port
[**GetPonPortInformationList**](OLTPonPortAPI.md#getponportinformationlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/pon-port/informations/list | Get pon port information list
[**GetPonPortInformationPage**](OLTPonPortAPI.md#getponportinformationpage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/pon-port/informations/page | Get pon port information page
[**GetPonPortList**](OLTPonPortAPI.md#getponportlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/pon-port/configs/list | Get pon port list
[**GetPonPortListMsp**](OLTPonPortAPI.md#getponportlistmsp) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/olts/{deviceMac}/pon/pon-port/configs/list | Get pon port list(msp mode)
[**GetPonPortPage**](OLTPonPortAPI.md#getponportpage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/pon-port/configs/page | Get pon port page



## EditPonPort

> OperationResponseDeviceResponseBodyReactivePonPortDTO EditPonPort(ctx, omadacId, siteId, deviceMac).PonPortBatchModifyDTO(ponPortBatchModifyDTO).Execute()

Batch modify pon port



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
	ponPortBatchModifyDTO := *openapiclient.NewPonPortBatchModifyDTO([]openapiclient.PonPortModifyDTO{*openapiclient.NewPonPortModifyDTO(int32(123), int32(123), int32(123), int32(123), "PortId_example")}) // PonPortBatchModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTPonPortAPI.EditPonPort(context.Background(), omadacId, siteId, deviceMac).PonPortBatchModifyDTO(ponPortBatchModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTPonPortAPI.EditPonPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPonPort`: OperationResponseDeviceResponseBodyReactivePonPortDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTPonPortAPI.EditPonPort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditPonPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ponPortBatchModifyDTO** | [**PonPortBatchModifyDTO**](PonPortBatchModifyDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyReactivePonPortDTO**](OperationResponseDeviceResponseBodyReactivePonPortDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPonPortInformationList

> OperationResponseListPonPortInformationDTO GetPonPortInformationList(ctx, omadacId, siteId, deviceMac).Execute()

Get pon port information list



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
	resp, r, err := apiClient.OLTPonPortAPI.GetPonPortInformationList(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTPonPortAPI.GetPonPortInformationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPonPortInformationList`: OperationResponseListPonPortInformationDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTPonPortAPI.GetPonPortInformationList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPonPortInformationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListPonPortInformationDTO**](OperationResponseListPonPortInformationDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPonPortInformationPage

> OperationResponsePageResponsePonPortInformationDTO GetPonPortInformationPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get pon port information page



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
	dto := *openapiclient.NewBaseDevicePageQueryRequest() // BaseDevicePageQueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTPonPortAPI.GetPonPortInformationPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTPonPortAPI.GetPonPortInformationPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPonPortInformationPage`: OperationResponsePageResponsePonPortInformationDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTPonPortAPI.GetPonPortInformationPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPonPortInformationPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**BaseDevicePageQueryRequest**](BaseDevicePageQueryRequest.md) |  | 

### Return type

[**OperationResponsePageResponsePonPortInformationDTO**](OperationResponsePageResponsePonPortInformationDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPonPortList

> OperationResponseListPonPortDTO GetPonPortList(ctx, omadacId, siteId, deviceMac).Execute()

Get pon port list



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
	resp, r, err := apiClient.OLTPonPortAPI.GetPonPortList(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTPonPortAPI.GetPonPortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPonPortList`: OperationResponseListPonPortDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTPonPortAPI.GetPonPortList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPonPortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListPonPortDTO**](OperationResponseListPonPortDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPonPortListMsp

> OperationResponseListPonPortDTO GetPonPortListMsp(ctx, siteId, deviceMac, mspId, customerId).Execute()

Get pon port list(msp mode)



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
	siteId := "siteId_example" // string | Site ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	mspId := "mspId_example" // string | mspId
	customerId := "customerId_example" // string | customerId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTPonPortAPI.GetPonPortListMsp(context.Background(), siteId, deviceMac, mspId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTPonPortAPI.GetPonPortListMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPonPortListMsp`: OperationResponseListPonPortDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTPonPortAPI.GetPonPortListMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**mspId** | **string** | mspId | 
**customerId** | **string** | customerId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPonPortListMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseListPonPortDTO**](OperationResponseListPonPortDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPonPortPage

> OperationResponsePageResponsePonPortDTO GetPonPortPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get pon port page



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
	dto := *openapiclient.NewBaseDevicePageQueryRequest() // BaseDevicePageQueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTPonPortAPI.GetPonPortPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTPonPortAPI.GetPonPortPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPonPortPage`: OperationResponsePageResponsePonPortDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTPonPortAPI.GetPonPortPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPonPortPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**BaseDevicePageQueryRequest**](BaseDevicePageQueryRequest.md) |  | 

### Return type

[**OperationResponsePageResponsePonPortDTO**](OperationResponsePageResponsePonPortDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \OLTLineProfileAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLineProfile**](OLTLineProfileAPI.md#AddLineProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/profiles/add | Create new line profile
[**DeleteLineProfile**](OLTLineProfileAPI.md#DeleteLineProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/profiles/delete | Delete an existing line profile
[**EditLineProfile**](OLTLineProfileAPI.md#EditLineProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/profiles/edit | Modify an existing line profile
[**GetLineProfileList**](OLTLineProfileAPI.md#GetLineProfileList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/profiles/list | Get line profile list
[**GetLineProfilePage**](OLTLineProfileAPI.md#GetLineProfilePage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/profiles/page | Get line profile page



## AddLineProfile

> OperationResponseDeviceResponseBodyLineProfileDTO AddLineProfile(ctx, omadacId, siteId, deviceMac).LineProfileDTO(lineProfileDTO).Execute()

Create new line profile



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
	lineProfileDTO := *openapiclient.NewLineProfileDTO(int32(123), "MappingMode_example", "UpstreamFEC_example") // LineProfileDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTLineProfileAPI.AddLineProfile(context.Background(), omadacId, siteId, deviceMac).LineProfileDTO(lineProfileDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTLineProfileAPI.AddLineProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLineProfile`: OperationResponseDeviceResponseBodyLineProfileDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTLineProfileAPI.AddLineProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddLineProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lineProfileDTO** | [**LineProfileDTO**](LineProfileDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyLineProfileDTO**](OperationResponseDeviceResponseBodyLineProfileDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLineProfile

> OperationResponseDeviceResponseBodyLineProfileDeleteResultDTO DeleteLineProfile(ctx, omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()

Delete an existing line profile



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
	intIdListRequest := *openapiclient.NewIntIdListRequest([]int32{int32(123)}) // IntIdListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTLineProfileAPI.DeleteLineProfile(context.Background(), omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTLineProfileAPI.DeleteLineProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLineProfile`: OperationResponseDeviceResponseBodyLineProfileDeleteResultDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTLineProfileAPI.DeleteLineProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteLineProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **intIdListRequest** | [**IntIdListRequest**](IntIdListRequest.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyLineProfileDeleteResultDTO**](OperationResponseDeviceResponseBodyLineProfileDeleteResultDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLineProfile

> OperationResponseDeviceResponseBodyVoid EditLineProfile(ctx, omadacId, siteId, deviceMac).LineProfileModifyDTO(lineProfileModifyDTO).Execute()

Modify an existing line profile



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
	lineProfileModifyDTO := *openapiclient.NewLineProfileModifyDTO(int32(123)) // LineProfileModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTLineProfileAPI.EditLineProfile(context.Background(), omadacId, siteId, deviceMac).LineProfileModifyDTO(lineProfileModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTLineProfileAPI.EditLineProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditLineProfile`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTLineProfileAPI.EditLineProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditLineProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lineProfileModifyDTO** | [**LineProfileModifyDTO**](LineProfileModifyDTO.md) |  | 

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


## GetLineProfileList

> OperationResponseListLineProfileVO GetLineProfileList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get line profile list



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
	dto := *openapiclient.NewLineProfileListQueryDTO() // LineProfileListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTLineProfileAPI.GetLineProfileList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTLineProfileAPI.GetLineProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLineProfileList`: OperationResponseListLineProfileVO
	fmt.Fprintf(os.Stdout, "Response from `OLTLineProfileAPI.GetLineProfileList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetLineProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**LineProfileListQueryDTO**](LineProfileListQueryDTO.md) |  | 

### Return type

[**OperationResponseListLineProfileVO**](OperationResponseListLineProfileVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLineProfilePage

> OperationResponsePageResponseLineProfileVO GetLineProfilePage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get line profile page



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
	dto := *openapiclient.NewLineProfilePageQueryDTO() // LineProfilePageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTLineProfileAPI.GetLineProfilePage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTLineProfileAPI.GetLineProfilePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLineProfilePage`: OperationResponsePageResponseLineProfileVO
	fmt.Fprintf(os.Stdout, "Response from `OLTLineProfileAPI.GetLineProfilePage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetLineProfilePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**LineProfilePageQueryDTO**](LineProfilePageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseLineProfileVO**](OperationResponsePageResponseLineProfileVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


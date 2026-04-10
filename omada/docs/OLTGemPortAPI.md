# OLTGemPortAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGemPort**](OLTGemPortAPI.md#addgemport) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-ports/add | Create new gem port
[**DeleteGemPort**](OLTGemPortAPI.md#deletegemport) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-ports/delete | Delete an existing gem port
[**EditGemPort**](OLTGemPortAPI.md#editgemport) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-ports/edit | Modify an existing gem port
[**GetGemPortList**](OLTGemPortAPI.md#getgemportlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-ports/list | Get gem port list
[**GetGemPortPage**](OLTGemPortAPI.md#getgemportpage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-ports/page | Get gem port page



## AddGemPort

> OperationResponseDeviceResponseBodyVoid AddGemPort(ctx, omadacId, siteId, deviceMac, lineProfileId).GemPortDTO(gemPortDTO).Execute()

Create new gem port



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
	gemPortDTO := *openapiclient.NewGemPortDTO("Encrypt_example", int32(123), int32(123)) // GemPortDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemPortAPI.AddGemPort(context.Background(), omadacId, siteId, deviceMac, lineProfileId).GemPortDTO(gemPortDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemPortAPI.AddGemPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGemPort`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTGemPortAPI.AddGemPort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddGemPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gemPortDTO** | [**GemPortDTO**](GemPortDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGemPort

> OperationResponseDeviceResponseBodyGemPortDeleteResultDTO DeleteGemPort(ctx, omadacId, siteId, deviceMac, lineProfileId).GemPortDeleteDTO(gemPortDeleteDTO).Execute()

Delete an existing gem port



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
	gemPortDeleteDTO := *openapiclient.NewGemPortDeleteDTO([]int32{int32(123)}) // GemPortDeleteDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemPortAPI.DeleteGemPort(context.Background(), omadacId, siteId, deviceMac, lineProfileId).GemPortDeleteDTO(gemPortDeleteDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemPortAPI.DeleteGemPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGemPort`: OperationResponseDeviceResponseBodyGemPortDeleteResultDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTGemPortAPI.DeleteGemPort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteGemPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gemPortDeleteDTO** | [**GemPortDeleteDTO**](GemPortDeleteDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyGemPortDeleteResultDTO**](OperationResponseDeviceResponseBodyGemPortDeleteResultDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditGemPort

> OperationResponseDeviceResponseBodyVoid EditGemPort(ctx, omadacId, siteId, deviceMac, lineProfileId).GemPortModifyDTO(gemPortModifyDTO).Execute()

Modify an existing gem port



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
	gemPortModifyDTO := *openapiclient.NewGemPortModifyDTO(int32(123)) // GemPortModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemPortAPI.EditGemPort(context.Background(), omadacId, siteId, deviceMac, lineProfileId).GemPortModifyDTO(gemPortModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemPortAPI.EditGemPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditGemPort`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTGemPortAPI.EditGemPort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditGemPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gemPortModifyDTO** | [**GemPortModifyDTO**](GemPortModifyDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGemPortList

> OperationResponseListGemPortDTO GetGemPortList(ctx, omadacId, siteId, deviceMac, lineProfileId).QueryParam(queryParam).Execute()

Get gem port list



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
	queryParam := *openapiclient.NewGemPortListQueryDTO() // GemPortListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemPortAPI.GetGemPortList(context.Background(), omadacId, siteId, deviceMac, lineProfileId).QueryParam(queryParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemPortAPI.GetGemPortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGemPortList`: OperationResponseListGemPortDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTGemPortAPI.GetGemPortList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGemPortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queryParam** | [**GemPortListQueryDTO**](GemPortListQueryDTO.md) |  | 

### Return type

[**OperationResponseListGemPortDTO**](OperationResponseListGemPortDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGemPortPage

> OperationResponsePageResponseGemPortDTO GetGemPortPage(ctx, omadacId, siteId, deviceMac, lineProfileId).QueryParam(queryParam).Execute()

Get gem port page



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
	queryParam := *openapiclient.NewGemPortPageQueryDTO() // GemPortPageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemPortAPI.GetGemPortPage(context.Background(), omadacId, siteId, deviceMac, lineProfileId).QueryParam(queryParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemPortAPI.GetGemPortPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGemPortPage`: OperationResponsePageResponseGemPortDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTGemPortAPI.GetGemPortPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGemPortPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queryParam** | [**GemPortPageQueryDTO**](GemPortPageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseGemPortDTO**](OperationResponsePageResponseGemPortDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


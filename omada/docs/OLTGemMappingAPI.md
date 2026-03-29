# \OLTGemMappingAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGemMapping**](OLTGemMappingAPI.md#AddGemMapping) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-mappings/add | Create new gem mapping
[**DeleteGemMapping**](OLTGemMappingAPI.md#DeleteGemMapping) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-mappings/delete | Delete an existing gem mapping
[**EditGemMapping**](OLTGemMappingAPI.md#EditGemMapping) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-mappings/edit | Modify an existing gem mapping
[**GetGemMappingList**](OLTGemMappingAPI.md#GetGemMappingList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-mappings/list | Get gem mapping list
[**GetGemMappingPage**](OLTGemMappingAPI.md#GetGemMappingPage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/line/{lineProfileId}/gem-mappings/page | Get gem mapping page



## AddGemMapping

> OperationResponseDeviceResponseBodyVoid AddGemMapping(ctx, omadacId, siteId, deviceMac, lineProfileId).GemMappingDTO(gemMappingDTO).Execute()

Create new gem mapping



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
	gemMappingDTO := *openapiclient.NewGemMappingDTO(int32(123), int32(123)) // GemMappingDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemMappingAPI.AddGemMapping(context.Background(), omadacId, siteId, deviceMac, lineProfileId).GemMappingDTO(gemMappingDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemMappingAPI.AddGemMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGemMapping`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTGemMappingAPI.AddGemMapping`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddGemMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gemMappingDTO** | [**GemMappingDTO**](GemMappingDTO.md) |  | 

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


## DeleteGemMapping

> OperationResponseDeviceResponseBodyVoid DeleteGemMapping(ctx, omadacId, siteId, deviceMac, lineProfileId).GemMappingDeleteDTO(gemMappingDeleteDTO).Execute()

Delete an existing gem mapping



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
	gemMappingDeleteDTO := *openapiclient.NewGemMappingDeleteDTO([]openapiclient.GemMappingDeleteItem{*openapiclient.NewGemMappingDeleteItem(int32(123), int32(123))}) // GemMappingDeleteDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemMappingAPI.DeleteGemMapping(context.Background(), omadacId, siteId, deviceMac, lineProfileId).GemMappingDeleteDTO(gemMappingDeleteDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemMappingAPI.DeleteGemMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGemMapping`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTGemMappingAPI.DeleteGemMapping`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteGemMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gemMappingDeleteDTO** | [**GemMappingDeleteDTO**](GemMappingDeleteDTO.md) |  | 

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


## EditGemMapping

> OperationResponseDeviceResponseBodyVoid EditGemMapping(ctx, omadacId, siteId, deviceMac, lineProfileId).GemMappingModifyDTO(gemMappingModifyDTO).Execute()

Modify an existing gem mapping



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
	gemMappingModifyDTO := *openapiclient.NewGemMappingModifyDTO(int32(123), int32(123)) // GemMappingModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemMappingAPI.EditGemMapping(context.Background(), omadacId, siteId, deviceMac, lineProfileId).GemMappingModifyDTO(gemMappingModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemMappingAPI.EditGemMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditGemMapping`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTGemMappingAPI.EditGemMapping`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditGemMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gemMappingModifyDTO** | [**GemMappingModifyDTO**](GemMappingModifyDTO.md) |  | 

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


## GetGemMappingList

> OperationResponseListGemMappingDTO GetGemMappingList(ctx, omadacId, siteId, deviceMac, lineProfileId).Dto(dto).Execute()

Get gem mapping list



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
	dto := *openapiclient.NewGemMappingListQueryDTO() // GemMappingListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemMappingAPI.GetGemMappingList(context.Background(), omadacId, siteId, deviceMac, lineProfileId).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemMappingAPI.GetGemMappingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGemMappingList`: OperationResponseListGemMappingDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTGemMappingAPI.GetGemMappingList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGemMappingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dto** | [**GemMappingListQueryDTO**](GemMappingListQueryDTO.md) |  | 

### Return type

[**OperationResponseListGemMappingDTO**](OperationResponseListGemMappingDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGemMappingPage

> OperationResponsePageResponseGemMappingDTO GetGemMappingPage(ctx, omadacId, siteId, deviceMac, lineProfileId).Dto(dto).Execute()

Get gem mapping page



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
	dto := *openapiclient.NewGemMappingPageQueryDTO() // GemMappingPageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTGemMappingAPI.GetGemMappingPage(context.Background(), omadacId, siteId, deviceMac, lineProfileId).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTGemMappingAPI.GetGemMappingPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGemMappingPage`: OperationResponsePageResponseGemMappingDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTGemMappingAPI.GetGemMappingPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGemMappingPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dto** | [**GemMappingPageQueryDTO**](GemMappingPageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseGemMappingDTO**](OperationResponsePageResponseGemMappingDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \OLTTrafficProfileAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrafficProfile**](OLTTrafficProfileAPI.md#AddTrafficProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/traffic/profiles/add | Create new traffic profile
[**DeleteTrafficProfile**](OLTTrafficProfileAPI.md#DeleteTrafficProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/traffic/profiles/delete | Delete an existing traffic profile
[**EditTrafficProfile**](OLTTrafficProfileAPI.md#EditTrafficProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/traffic/profiles/edit | Modify an existing traffic profile
[**GetTrafficProfileList**](OLTTrafficProfileAPI.md#GetTrafficProfileList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/traffic/profiles/list | Get traffic profile list
[**GetTrafficProfilePage**](OLTTrafficProfileAPI.md#GetTrafficProfilePage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/traffic/profiles/page | Get traffic profile page



## AddTrafficProfile

> OperationResponseDeviceResponseBodyTrafficProfileAddResultDTO AddTrafficProfile(ctx, omadacId, siteId, deviceMac).TrafficProfileDTO(trafficProfileDTO).Execute()

Create new traffic profile



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
	trafficProfileDTO := *openapiclient.NewTrafficProfileDTO("InnerPriority_example", "Priority_example", "PriorityPolicy_example", int32(123), "RateLimitStatus_example", int32(123)) // TrafficProfileDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTrafficProfileAPI.AddTrafficProfile(context.Background(), omadacId, siteId, deviceMac).TrafficProfileDTO(trafficProfileDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTrafficProfileAPI.AddTrafficProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTrafficProfile`: OperationResponseDeviceResponseBodyTrafficProfileAddResultDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTTrafficProfileAPI.AddTrafficProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddTrafficProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **trafficProfileDTO** | [**TrafficProfileDTO**](TrafficProfileDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyTrafficProfileAddResultDTO**](OperationResponseDeviceResponseBodyTrafficProfileAddResultDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrafficProfile

> OperationResponseDeviceResponseBodyTrafficProfileDeleteResultDTO DeleteTrafficProfile(ctx, omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()

Delete an existing traffic profile



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
	resp, r, err := apiClient.OLTTrafficProfileAPI.DeleteTrafficProfile(context.Background(), omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTrafficProfileAPI.DeleteTrafficProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrafficProfile`: OperationResponseDeviceResponseBodyTrafficProfileDeleteResultDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTTrafficProfileAPI.DeleteTrafficProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteTrafficProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **intIdListRequest** | [**IntIdListRequest**](IntIdListRequest.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyTrafficProfileDeleteResultDTO**](OperationResponseDeviceResponseBodyTrafficProfileDeleteResultDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTrafficProfile

> OperationResponseDeviceResponseBodyVoid EditTrafficProfile(ctx, omadacId, siteId, deviceMac).TrafficProfileModifyDTO(trafficProfileModifyDTO).Execute()

Modify an existing traffic profile



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
	trafficProfileModifyDTO := *openapiclient.NewTrafficProfileModifyDTO("InnerPriority_example", "Priority_example", "PriorityPolicy_example", int32(123), "RateLimitStatus_example", int32(123)) // TrafficProfileModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTrafficProfileAPI.EditTrafficProfile(context.Background(), omadacId, siteId, deviceMac).TrafficProfileModifyDTO(trafficProfileModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTrafficProfileAPI.EditTrafficProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTrafficProfile`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTTrafficProfileAPI.EditTrafficProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditTrafficProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **trafficProfileModifyDTO** | [**TrafficProfileModifyDTO**](TrafficProfileModifyDTO.md) |  | 

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


## GetTrafficProfileList

> OperationResponseListTrafficProfileVO GetTrafficProfileList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get traffic profile list



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
	dto := *openapiclient.NewTrafficProfileListQueryDTO() // TrafficProfileListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTrafficProfileAPI.GetTrafficProfileList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTrafficProfileAPI.GetTrafficProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrafficProfileList`: OperationResponseListTrafficProfileVO
	fmt.Fprintf(os.Stdout, "Response from `OLTTrafficProfileAPI.GetTrafficProfileList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetTrafficProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**TrafficProfileListQueryDTO**](TrafficProfileListQueryDTO.md) |  | 

### Return type

[**OperationResponseListTrafficProfileVO**](OperationResponseListTrafficProfileVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrafficProfilePage

> OperationResponsePageResponseTrafficProfileVO GetTrafficProfilePage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get traffic profile page



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
	dto := *openapiclient.NewTrafficProfilePageQueryDTO() // TrafficProfilePageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTTrafficProfileAPI.GetTrafficProfilePage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTTrafficProfileAPI.GetTrafficProfilePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrafficProfilePage`: OperationResponsePageResponseTrafficProfileVO
	fmt.Fprintf(os.Stdout, "Response from `OLTTrafficProfileAPI.GetTrafficProfilePage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetTrafficProfilePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**TrafficProfilePageQueryDTO**](TrafficProfilePageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseTrafficProfileVO**](OperationResponsePageResponseTrafficProfileVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


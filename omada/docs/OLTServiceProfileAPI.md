# OLTServiceProfileAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServiceProfile**](OLTServiceProfileAPI.md#addserviceprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/profiles/add | Create new service profile
[**DeleteServiceProfile**](OLTServiceProfileAPI.md#deleteserviceprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/profiles/delete | Delete an existing service profile
[**EditServiceProfile**](OLTServiceProfileAPI.md#editserviceprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/profiles/edit | Modify an existing service profile
[**GetServiceProfileList**](OLTServiceProfileAPI.md#getserviceprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/profiles/list | Get service profile list
[**GetServiceProfilePage**](OLTServiceProfileAPI.md#getserviceprofilepage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/profiles/page | Get service profile page



## AddServiceProfile

> OperationResponseDeviceResponseBodyServiceProfileDTO AddServiceProfile(ctx, omadacId, siteId, deviceMac).ServiceProfileDTO(serviceProfileDTO).Execute()

Create new service profile



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
	serviceProfileDTO := *openapiclient.NewServiceProfileDTO(int32(123), "MacLearning_example", "MulticastForward_example", "MulticastMode_example", "NativeVlan_example", int32(123), int32(123)) // ServiceProfileDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServiceProfileAPI.AddServiceProfile(context.Background(), omadacId, siteId, deviceMac).ServiceProfileDTO(serviceProfileDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServiceProfileAPI.AddServiceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddServiceProfile`: OperationResponseDeviceResponseBodyServiceProfileDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServiceProfileAPI.AddServiceProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddServiceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **serviceProfileDTO** | [**ServiceProfileDTO**](ServiceProfileDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyServiceProfileDTO**](OperationResponseDeviceResponseBodyServiceProfileDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceProfile

> OperationResponseDeviceResponseBodyServiceProfileDeleteResultDTO DeleteServiceProfile(ctx, omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()

Delete an existing service profile



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
	resp, r, err := apiClient.OLTServiceProfileAPI.DeleteServiceProfile(context.Background(), omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServiceProfileAPI.DeleteServiceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceProfile`: OperationResponseDeviceResponseBodyServiceProfileDeleteResultDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServiceProfileAPI.DeleteServiceProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteServiceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **intIdListRequest** | [**IntIdListRequest**](IntIdListRequest.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyServiceProfileDeleteResultDTO**](OperationResponseDeviceResponseBodyServiceProfileDeleteResultDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditServiceProfile

> OperationResponseDeviceResponseBodyVoid EditServiceProfile(ctx, omadacId, siteId, deviceMac).ServiceProfileModifyDTO(serviceProfileModifyDTO).Execute()

Modify an existing service profile



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
	serviceProfileModifyDTO := *openapiclient.NewServiceProfileModifyDTO("MulticastForward_example", int32(123)) // ServiceProfileModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServiceProfileAPI.EditServiceProfile(context.Background(), omadacId, siteId, deviceMac).ServiceProfileModifyDTO(serviceProfileModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServiceProfileAPI.EditServiceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditServiceProfile`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServiceProfileAPI.EditServiceProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditServiceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **serviceProfileModifyDTO** | [**ServiceProfileModifyDTO**](ServiceProfileModifyDTO.md) |  | 

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


## GetServiceProfileList

> OperationResponseListServiceProfileVO GetServiceProfileList(ctx, omadacId, siteId, deviceMac).QueryParam(queryParam).Execute()

Get service profile list



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
	queryParam := *openapiclient.NewServiceProfileListQueryDTO() // ServiceProfileListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServiceProfileAPI.GetServiceProfileList(context.Background(), omadacId, siteId, deviceMac).QueryParam(queryParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServiceProfileAPI.GetServiceProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceProfileList`: OperationResponseListServiceProfileVO
	fmt.Fprintf(os.Stdout, "Response from `OLTServiceProfileAPI.GetServiceProfileList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetServiceProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **queryParam** | [**ServiceProfileListQueryDTO**](ServiceProfileListQueryDTO.md) |  | 

### Return type

[**OperationResponseListServiceProfileVO**](OperationResponseListServiceProfileVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceProfilePage

> OperationResponsePageResponseServiceProfileVO GetServiceProfilePage(ctx, omadacId, siteId, deviceMac).QueryParam(queryParam).Execute()

Get service profile page



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
	queryParam := *openapiclient.NewServiceProfilePageQueryDTO() // ServiceProfilePageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServiceProfileAPI.GetServiceProfilePage(context.Background(), omadacId, siteId, deviceMac).QueryParam(queryParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServiceProfileAPI.GetServiceProfilePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceProfilePage`: OperationResponsePageResponseServiceProfileVO
	fmt.Fprintf(os.Stdout, "Response from `OLTServiceProfileAPI.GetServiceProfilePage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetServiceProfilePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **queryParam** | [**ServiceProfilePageQueryDTO**](ServiceProfilePageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseServiceProfileVO**](OperationResponsePageResponseServiceProfileVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \OLTServicePortProfileAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServicePortProfile**](OLTServicePortProfileAPI.md#AddServicePortProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/add | Create new service port profile
[**AddServicePortProfileDetail**](OLTServicePortProfileAPI.md#AddServicePortProfileDetail) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/detail/add | Create new service port profile detail
[**DeleteServicePortProfile**](OLTServicePortProfileAPI.md#DeleteServicePortProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/delete | Delete an existing service port profile
[**DeleteServicePortProfileDetail**](OLTServicePortProfileAPI.md#DeleteServicePortProfileDetail) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/detail/delete | Delete an existing service port profile detail
[**EditServicePortProfile**](OLTServicePortProfileAPI.md#EditServicePortProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/edit | Modify an existing service port profile
[**EditServicePortProfileDetail**](OLTServicePortProfileAPI.md#EditServicePortProfileDetail) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/detail/edit | Modify an existing service port profile detail
[**GetServicePortProfileList**](OLTServicePortProfileAPI.md#GetServicePortProfileList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/list | Get service port profile list
[**GetServicePortProfileListDetail**](OLTServicePortProfileAPI.md#GetServicePortProfileListDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/detail/list | Get service port profile list detail
[**GetServicePortProfilePage**](OLTServicePortProfileAPI.md#GetServicePortProfilePage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/page | Get service port profile page
[**GetServicePortProfilePageDetail**](OLTServicePortProfileAPI.md#GetServicePortProfilePageDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service-port/profiles/detail/page | Get service port profile page detail



## AddServicePortProfile

> OperationResponseDeviceResponseBodyServicePortProfileModifyDTO AddServicePortProfile(ctx, omadacId, siteId, deviceMac).ServicePortProfileModifyDTO(servicePortProfileModifyDTO).Execute()

Create new service port profile



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
	servicePortProfileModifyDTO := *openapiclient.NewServicePortProfileModifyDTO("ServicePortProfileId_example") // ServicePortProfileModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.AddServicePortProfile(context.Background(), omadacId, siteId, deviceMac).ServicePortProfileModifyDTO(servicePortProfileModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.AddServicePortProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddServicePortProfile`: OperationResponseDeviceResponseBodyServicePortProfileModifyDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.AddServicePortProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddServicePortProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **servicePortProfileModifyDTO** | [**ServicePortProfileModifyDTO**](ServicePortProfileModifyDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyServicePortProfileModifyDTO**](OperationResponseDeviceResponseBodyServicePortProfileModifyDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddServicePortProfileDetail

> OperationResponseDeviceResponseBodyVoid AddServicePortProfileDetail(ctx, omadacId, siteId, deviceMac).ServicePortProfileDetailDTO(servicePortProfileDetailDTO).Execute()

Create new service port profile detail



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
	servicePortProfileDetailDTO := *openapiclient.NewServicePortProfileDetailDTO("AdminStatus_example", "CreationMode_example", int32(123), "EtherType_example", int32(123), "InUse_example", int32(123), int32(123), int32(123), int32(123), "ServicePortId_example", "ServicePortProfileId_example", "StatisticPerformance_example", int32(123), "TagAction_example", int32(123), int32(123)) // ServicePortProfileDetailDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.AddServicePortProfileDetail(context.Background(), omadacId, siteId, deviceMac).ServicePortProfileDetailDTO(servicePortProfileDetailDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.AddServicePortProfileDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddServicePortProfileDetail`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.AddServicePortProfileDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddServicePortProfileDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **servicePortProfileDetailDTO** | [**ServicePortProfileDetailDTO**](ServicePortProfileDetailDTO.md) |  | 

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


## DeleteServicePortProfile

> OperationResponseDeviceResponseBodyVoid DeleteServicePortProfile(ctx, omadacId, siteId, deviceMac).ServicePortProfileDeleteDTO(servicePortProfileDeleteDTO).Execute()

Delete an existing service port profile



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
	servicePortProfileDeleteDTO := *openapiclient.NewServicePortProfileDeleteDTO() // ServicePortProfileDeleteDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.DeleteServicePortProfile(context.Background(), omadacId, siteId, deviceMac).ServicePortProfileDeleteDTO(servicePortProfileDeleteDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.DeleteServicePortProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServicePortProfile`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.DeleteServicePortProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteServicePortProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **servicePortProfileDeleteDTO** | [**ServicePortProfileDeleteDTO**](ServicePortProfileDeleteDTO.md) |  | 

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


## DeleteServicePortProfileDetail

> OperationResponseDeviceResponseBodyVoid DeleteServicePortProfileDetail(ctx, omadacId, siteId, deviceMac).ServicePortProfileDetailDeleteDTO(servicePortProfileDetailDeleteDTO).Execute()

Delete an existing service port profile detail



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
	servicePortProfileDetailDeleteDTO := *openapiclient.NewServicePortProfileDetailDeleteDTO() // ServicePortProfileDetailDeleteDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.DeleteServicePortProfileDetail(context.Background(), omadacId, siteId, deviceMac).ServicePortProfileDetailDeleteDTO(servicePortProfileDetailDeleteDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.DeleteServicePortProfileDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServicePortProfileDetail`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.DeleteServicePortProfileDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteServicePortProfileDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **servicePortProfileDetailDeleteDTO** | [**ServicePortProfileDetailDeleteDTO**](ServicePortProfileDetailDeleteDTO.md) |  | 

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


## EditServicePortProfile

> OperationResponseDeviceResponseBodyVoid EditServicePortProfile(ctx, omadacId, siteId, deviceMac).ServicePortProfileModifyDTO(servicePortProfileModifyDTO).Execute()

Modify an existing service port profile



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
	servicePortProfileModifyDTO := *openapiclient.NewServicePortProfileModifyDTO("ServicePortProfileId_example") // ServicePortProfileModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.EditServicePortProfile(context.Background(), omadacId, siteId, deviceMac).ServicePortProfileModifyDTO(servicePortProfileModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.EditServicePortProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditServicePortProfile`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.EditServicePortProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditServicePortProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **servicePortProfileModifyDTO** | [**ServicePortProfileModifyDTO**](ServicePortProfileModifyDTO.md) |  | 

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


## EditServicePortProfileDetail

> OperationResponseDeviceResponseBodyVoid EditServicePortProfileDetail(ctx, omadacId, siteId, deviceMac).ServicePortProfileDetailDTO(servicePortProfileDetailDTO).Execute()

Modify an existing service port profile detail



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
	servicePortProfileDetailDTO := *openapiclient.NewServicePortProfileDetailDTO("AdminStatus_example", "CreationMode_example", int32(123), "EtherType_example", int32(123), "InUse_example", int32(123), int32(123), int32(123), int32(123), "ServicePortId_example", "ServicePortProfileId_example", "StatisticPerformance_example", int32(123), "TagAction_example", int32(123), int32(123)) // ServicePortProfileDetailDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.EditServicePortProfileDetail(context.Background(), omadacId, siteId, deviceMac).ServicePortProfileDetailDTO(servicePortProfileDetailDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.EditServicePortProfileDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditServicePortProfileDetail`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.EditServicePortProfileDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditServicePortProfileDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **servicePortProfileDetailDTO** | [**ServicePortProfileDetailDTO**](ServicePortProfileDetailDTO.md) |  | 

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


## GetServicePortProfileList

> OperationResponseListServicePortProfileDTO GetServicePortProfileList(ctx, omadacId, siteId, deviceMac).Execute()

Get service port profile list



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
	resp, r, err := apiClient.OLTServicePortProfileAPI.GetServicePortProfileList(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.GetServicePortProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePortProfileList`: OperationResponseListServicePortProfileDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.GetServicePortProfileList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetServicePortProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListServicePortProfileDTO**](OperationResponseListServicePortProfileDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePortProfileListDetail

> OperationResponseListServicePortProfileListDetailDTO GetServicePortProfileListDetail(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get service port profile list detail



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
	dto := *openapiclient.NewServicePortProfileDetailListQueryDTO() // ServicePortProfileDetailListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.GetServicePortProfileListDetail(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.GetServicePortProfileListDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePortProfileListDetail`: OperationResponseListServicePortProfileListDetailDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.GetServicePortProfileListDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetServicePortProfileListDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**ServicePortProfileDetailListQueryDTO**](ServicePortProfileDetailListQueryDTO.md) |  | 

### Return type

[**OperationResponseListServicePortProfileListDetailDTO**](OperationResponseListServicePortProfileListDetailDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePortProfilePage

> OperationResponsePageResponseServicePortProfileDTO GetServicePortProfilePage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get service port profile page



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
	dto := *openapiclient.NewServicePortProfilePageQueryDTO() // ServicePortProfilePageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.GetServicePortProfilePage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.GetServicePortProfilePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePortProfilePage`: OperationResponsePageResponseServicePortProfileDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.GetServicePortProfilePage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetServicePortProfilePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**ServicePortProfilePageQueryDTO**](ServicePortProfilePageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseServicePortProfileDTO**](OperationResponsePageResponseServicePortProfileDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePortProfilePageDetail

> OperationResponsePageResponseServicePortProfileListDetailDTO GetServicePortProfilePageDetail(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get service port profile page detail



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
	dto := *openapiclient.NewServicePortProfileDetailPageQueryDTO("ServicePortProfileId_example") // ServicePortProfileDetailPageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortProfileAPI.GetServicePortProfilePageDetail(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortProfileAPI.GetServicePortProfilePageDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePortProfilePageDetail`: OperationResponsePageResponseServicePortProfileListDetailDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortProfileAPI.GetServicePortProfilePageDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetServicePortProfilePageDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**ServicePortProfileDetailPageQueryDTO**](ServicePortProfileDetailPageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseServicePortProfileListDetailDTO**](OperationResponsePageResponseServicePortProfileListDetailDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \OLTServicePortAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServicePort**](OLTServicePortAPI.md#AddServicePort) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-ports/add | Create new service port
[**ClearStatisticInfo**](OLTServicePortAPI.md#ClearStatisticInfo) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-port/statistics/clear | Batch delete existing statistic info
[**DeleteServicePort**](OLTServicePortAPI.md#DeleteServicePort) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-ports/delete | Batch delete existing service port
[**EditAutoRefreshConfig**](OLTServicePortAPI.md#EditAutoRefreshConfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-port/statistic/configs/edit | Modify an existing auto refresh config
[**EditAutoServicePort**](OLTServicePortAPI.md#EditAutoServicePort) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/auto-service-ports/edit | Batch modify auto service port
[**EditServicePort**](OLTServicePortAPI.md#EditServicePort) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-ports/edit | Modify an existing service port
[**GetAutoRefreshConfig**](OLTServicePortAPI.md#GetAutoRefreshConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-port/statistic/configs | Get auto refresh config
[**GetAutoServicePortConfigPage**](OLTServicePortAPI.md#GetAutoServicePortConfigPage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/auto-service-ports/page | Get auto service port config page
[**GetAutoServicePortList**](OLTServicePortAPI.md#GetAutoServicePortList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/auto-service-ports/list | Get auto service port list
[**GetServicePortList**](OLTServicePortAPI.md#GetServicePortList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-ports/list | Get service port list
[**GetServicePortPage**](OLTServicePortAPI.md#GetServicePortPage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-ports/page | Get service port page
[**GetStatisticInfoDetail**](OLTServicePortAPI.md#GetStatisticInfoDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-port/statistics/detail | Get statistic info detail
[**GetStatisticInfoList**](OLTServicePortAPI.md#GetStatisticInfoList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-port/statistics/list | Get statistic info list
[**GetStatisticInfoPage**](OLTServicePortAPI.md#GetStatisticInfoPage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/service-port/statistics/page | Get statistic info page



## AddServicePort

> OperationResponseDeviceResponseBodyVoid AddServicePort(ctx, omadacId, siteId, deviceMac).ServicePortAddDTO(servicePortAddDTO).Execute()

Create new service port



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
	servicePortAddDTO := *openapiclient.NewServicePortAddDTO("AdminStatus_example", false, "OnuId_example", "StatisticPerformance_example") // ServicePortAddDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.AddServicePort(context.Background(), omadacId, siteId, deviceMac).ServicePortAddDTO(servicePortAddDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.AddServicePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddServicePort`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.AddServicePort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddServicePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **servicePortAddDTO** | [**ServicePortAddDTO**](ServicePortAddDTO.md) |  | 

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


## ClearStatisticInfo

> OperationResponseDeviceResponseBodyVoid ClearStatisticInfo(ctx, omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()

Batch delete existing statistic info



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
	resp, r, err := apiClient.OLTServicePortAPI.ClearStatisticInfo(context.Background(), omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.ClearStatisticInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearStatisticInfo`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.ClearStatisticInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiClearStatisticInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **intIdListRequest** | [**IntIdListRequest**](IntIdListRequest.md) |  | 

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


## DeleteServicePort

> OperationResponseDeviceResponseBodyVoid DeleteServicePort(ctx, omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()

Batch delete existing service port



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
	resp, r, err := apiClient.OLTServicePortAPI.DeleteServicePort(context.Background(), omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.DeleteServicePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServicePort`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.DeleteServicePort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteServicePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **intIdListRequest** | [**IntIdListRequest**](IntIdListRequest.md) |  | 

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


## EditAutoRefreshConfig

> OperationResponseDeviceResponseBodyVoid EditAutoRefreshConfig(ctx, omadacId, siteId, deviceMac).StatisticConfigDTO(statisticConfigDTO).Execute()

Modify an existing auto refresh config



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
	statisticConfigDTO := *openapiclient.NewStatisticConfigDTO() // StatisticConfigDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.EditAutoRefreshConfig(context.Background(), omadacId, siteId, deviceMac).StatisticConfigDTO(statisticConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.EditAutoRefreshConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAutoRefreshConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.EditAutoRefreshConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditAutoRefreshConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **statisticConfigDTO** | [**StatisticConfigDTO**](StatisticConfigDTO.md) |  | 

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


## EditAutoServicePort

> OperationResponseDeviceResponseBodyVoid EditAutoServicePort(ctx, omadacId, siteId, deviceMac).AutoServicePortDTO(autoServicePortDTO).Execute()

Batch modify auto service port



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
	autoServicePortDTO := []openapiclient.AutoServicePortDTO{*openapiclient.NewAutoServicePortDTO()} // []AutoServicePortDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.EditAutoServicePort(context.Background(), omadacId, siteId, deviceMac).AutoServicePortDTO(autoServicePortDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.EditAutoServicePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAutoServicePort`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.EditAutoServicePort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditAutoServicePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **autoServicePortDTO** | [**[]AutoServicePortDTO**](AutoServicePortDTO.md) |  | 

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


## EditServicePort

> OperationResponseDeviceResponseBodyVoid EditServicePort(ctx, omadacId, siteId, deviceMac).ServicePortModifyDTO(servicePortModifyDTO).Execute()

Modify an existing service port



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
	servicePortModifyDTO := *openapiclient.NewServicePortModifyDTO("AdminStatus_example", int32(123), int32(123), "StatisticPerformance_example") // ServicePortModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.EditServicePort(context.Background(), omadacId, siteId, deviceMac).ServicePortModifyDTO(servicePortModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.EditServicePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditServicePort`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.EditServicePort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditServicePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **servicePortModifyDTO** | [**ServicePortModifyDTO**](ServicePortModifyDTO.md) |  | 

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


## GetAutoRefreshConfig

> OperationResponseStatisticConfigDTO GetAutoRefreshConfig(ctx, omadacId, siteId, deviceMac).Execute()

Get auto refresh config



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
	resp, r, err := apiClient.OLTServicePortAPI.GetAutoRefreshConfig(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.GetAutoRefreshConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoRefreshConfig`: OperationResponseStatisticConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.GetAutoRefreshConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAutoRefreshConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseStatisticConfigDTO**](OperationResponseStatisticConfigDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoServicePortConfigPage

> OperationResponsePageResponseAutoServicePortVO GetAutoServicePortConfigPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get auto service port config page



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
	dto := *openapiclient.NewAutoServicePortQueryDTO() // AutoServicePortQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.GetAutoServicePortConfigPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.GetAutoServicePortConfigPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoServicePortConfigPage`: OperationResponsePageResponseAutoServicePortVO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.GetAutoServicePortConfigPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAutoServicePortConfigPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**AutoServicePortQueryDTO**](AutoServicePortQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseAutoServicePortVO**](OperationResponsePageResponseAutoServicePortVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoServicePortList

> OperationResponseListAutoServicePortVO GetAutoServicePortList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get auto service port list



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
	dto := *openapiclient.NewAutoServicePortQueryDTO() // AutoServicePortQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.GetAutoServicePortList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.GetAutoServicePortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoServicePortList`: OperationResponseListAutoServicePortVO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.GetAutoServicePortList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAutoServicePortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**AutoServicePortQueryDTO**](AutoServicePortQueryDTO.md) |  | 

### Return type

[**OperationResponseListAutoServicePortVO**](OperationResponseListAutoServicePortVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePortList

> OperationResponseListServicePortVO GetServicePortList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get service port list



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
	dto := *openapiclient.NewServicePortQueryDTO() // ServicePortQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.GetServicePortList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.GetServicePortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePortList`: OperationResponseListServicePortVO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.GetServicePortList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetServicePortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**ServicePortQueryDTO**](ServicePortQueryDTO.md) |  | 

### Return type

[**OperationResponseListServicePortVO**](OperationResponseListServicePortVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePortPage

> OperationResponsePageResponseServicePortVO GetServicePortPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get service port page



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
	dto := *openapiclient.NewServicePortQueryDTO() // ServicePortQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.GetServicePortPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.GetServicePortPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePortPage`: OperationResponsePageResponseServicePortVO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.GetServicePortPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetServicePortPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**ServicePortQueryDTO**](ServicePortQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseServicePortVO**](OperationResponsePageResponseServicePortVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatisticInfoDetail

> OperationResponseStatisticInfoDetailDTO GetStatisticInfoDetail(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get statistic info detail



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
	dto := *openapiclient.NewStatisticInfoDetailQueryDTO(int32(123)) // StatisticInfoDetailQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.GetStatisticInfoDetail(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.GetStatisticInfoDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatisticInfoDetail`: OperationResponseStatisticInfoDetailDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.GetStatisticInfoDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStatisticInfoDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**StatisticInfoDetailQueryDTO**](StatisticInfoDetailQueryDTO.md) |  | 

### Return type

[**OperationResponseStatisticInfoDetailDTO**](OperationResponseStatisticInfoDetailDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatisticInfoList

> OperationResponseListStatisticInfoDTO GetStatisticInfoList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get statistic info list



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
	dto := *openapiclient.NewStatisticInfoQueryDTO() // StatisticInfoQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.GetStatisticInfoList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.GetStatisticInfoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatisticInfoList`: OperationResponseListStatisticInfoDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.GetStatisticInfoList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStatisticInfoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**StatisticInfoQueryDTO**](StatisticInfoQueryDTO.md) |  | 

### Return type

[**OperationResponseListStatisticInfoDTO**](OperationResponseListStatisticInfoDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatisticInfoPage

> OperationResponsePageResponseStatisticInfoDTO GetStatisticInfoPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get statistic info page



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
	dto := *openapiclient.NewStatisticInfoQueryDTO() // StatisticInfoQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTServicePortAPI.GetStatisticInfoPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTServicePortAPI.GetStatisticInfoPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatisticInfoPage`: OperationResponsePageResponseStatisticInfoDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTServicePortAPI.GetStatisticInfoPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStatisticInfoPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**StatisticInfoQueryDTO**](StatisticInfoQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseStatisticInfoDTO**](OperationResponsePageResponseStatisticInfoDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


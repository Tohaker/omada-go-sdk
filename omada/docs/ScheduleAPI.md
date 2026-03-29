# \ScheduleAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePoeSchedule**](ScheduleAPI.md#CreatePoeSchedule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/poe-schedules | Create new PoE Schedule
[**CreatePortSchedule**](ScheduleAPI.md#CreatePortSchedule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/port-schedules | Create new Port Schedule
[**CreateRebootSchedule1**](ScheduleAPI.md#CreateRebootSchedule1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/reboot-schedules | Create new reboot schedule
[**CreateUpgradeSchedule**](ScheduleAPI.md#CreateUpgradeSchedule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/upgrade-schedules | Create new upgrade schedule
[**DeletePoeSchedule**](ScheduleAPI.md#DeletePoeSchedule) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/poe-schedules/{poeScheduleId} | Delete PoE Schedule
[**DeletePortSchedule**](ScheduleAPI.md#DeletePortSchedule) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/port-schedules/{portScheduleId} | Delete Port Schedule
[**DeleteRebootSchedule1**](ScheduleAPI.md#DeleteRebootSchedule1) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/reboot-schedules/{id} | Delete reboot schedule
[**DeleteUpgradeSchedule**](ScheduleAPI.md#DeleteUpgradeSchedule) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/upgrade-schedules/{id} | Delete upgrade schedule
[**GetPoePortsList**](ScheduleAPI.md#GetPoePortsList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/poe-ports | Get PoE ports list
[**GetPoeScheduleList**](ScheduleAPI.md#GetPoeScheduleList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/poe-schedules | Get PoE schedule list
[**GetPortScheduleList**](ScheduleAPI.md#GetPortScheduleList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/port-schedules | Get port schedule list
[**GetPortSchedulePorts**](ScheduleAPI.md#GetPortSchedulePorts) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/port-status-ports | Get device ports list supporting port schedule
[**GetRebootScheduleList2**](ScheduleAPI.md#GetRebootScheduleList2) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/reboot-schedules | Get reboot schedule list
[**GetUpgradeScheduleList**](ScheduleAPI.md#GetUpgradeScheduleList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/upgrade-schedules | Get upgrade schedule list
[**ModifyPoeSchedule**](ScheduleAPI.md#ModifyPoeSchedule) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/poe-schedules/{poeScheduleId} | Modify PoE Schedule
[**ModifyPortSchedule**](ScheduleAPI.md#ModifyPortSchedule) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/port-schedules/{portScheduleId} | Modify Port Schedule
[**ModifyRebootSchedule1**](ScheduleAPI.md#ModifyRebootSchedule1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/reboot-schedules/{id} | Modify reboot schedule
[**ModifyUpgradeSchedule**](ScheduleAPI.md#ModifyUpgradeSchedule) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/upgrade-schedules/{id} | Modify upgrade schedule



## CreatePoeSchedule

> OperationResponseWithoutResult CreatePoeSchedule(ctx, omadacId, siteId).PoeScheduleOpenApiVO(poeScheduleOpenApiVO).Execute()

Create new PoE Schedule



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
	poeScheduleOpenApiVO := *openapiclient.NewPoeScheduleOpenApiVO("Name_example", map[string][]int32{"key": []int32{int32(123)}}, false, "TurnOnTime_example") // PoeScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.CreatePoeSchedule(context.Background(), omadacId, siteId).PoeScheduleOpenApiVO(poeScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.CreatePoeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePoeSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.CreatePoeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePoeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **poeScheduleOpenApiVO** | [**PoeScheduleOpenApiVO**](PoeScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortSchedule

> OperationResponseWithoutResult CreatePortSchedule(ctx, omadacId, siteId).PortScheduleOpenApiVO(portScheduleOpenApiVO).Execute()

Create new Port Schedule



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
	portScheduleOpenApiVO := *openapiclient.NewPortScheduleOpenApiVO("Name_example", map[string][]int32{"key": []int32{int32(123)}}, false, "TurnOnTime_example") // PortScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.CreatePortSchedule(context.Background(), omadacId, siteId).PortScheduleOpenApiVO(portScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.CreatePortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.CreatePortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portScheduleOpenApiVO** | [**PortScheduleOpenApiVO**](PortScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRebootSchedule1

> OperationResponseWithoutResult CreateRebootSchedule1(ctx, omadacId, siteId).RebootScheduleOpenApiVO(rebootScheduleOpenApiVO).Execute()

Create new reboot schedule



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
	rebootScheduleOpenApiVO := *openapiclient.NewRebootScheduleOpenApiVO([]string{"DeviceMacs_example"}, "Name_example", false, *openapiclient.NewRebootScheduleTimeOpenApiVO(int32(123), int32(123), int32(123))) // RebootScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.CreateRebootSchedule1(context.Background(), omadacId, siteId).RebootScheduleOpenApiVO(rebootScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.CreateRebootSchedule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRebootSchedule1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.CreateRebootSchedule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRebootSchedule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rebootScheduleOpenApiVO** | [**RebootScheduleOpenApiVO**](RebootScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUpgradeSchedule

> OperationResponseWithoutResult CreateUpgradeSchedule(ctx, omadacId, siteId).UpgradeScheduleOpenApiVO(upgradeScheduleOpenApiVO).Execute()

Create new upgrade schedule



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
	upgradeScheduleOpenApiVO := *openapiclient.NewUpgradeScheduleOpenApiVO([]string{"DeviceMacs_example"}, "Name_example", false, int32(123)) // UpgradeScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.CreateUpgradeSchedule(context.Background(), omadacId, siteId).UpgradeScheduleOpenApiVO(upgradeScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.CreateUpgradeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUpgradeSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.CreateUpgradeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upgradeScheduleOpenApiVO** | [**UpgradeScheduleOpenApiVO**](UpgradeScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePoeSchedule

> OperationResponseWithoutResult DeletePoeSchedule(ctx, omadacId, siteId, poeScheduleId).Execute()

Delete PoE Schedule



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
	poeScheduleId := "poeScheduleId_example" // string | PoE Schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.DeletePoeSchedule(context.Background(), omadacId, siteId, poeScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.DeletePoeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePoeSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.DeletePoeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**poeScheduleId** | **string** | PoE Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePoeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePortSchedule

> OperationResponseWithoutResult DeletePortSchedule(ctx, omadacId, siteId, portScheduleId).Execute()

Delete Port Schedule



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
	portScheduleId := "portScheduleId_example" // string | Port Schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.DeletePortSchedule(context.Background(), omadacId, siteId, portScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.DeletePortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.DeletePortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portScheduleId** | **string** | Port Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRebootSchedule1

> OperationResponseWithoutResult DeleteRebootSchedule1(ctx, omadacId, siteId, id).Execute()

Delete reboot schedule



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
	id := "id_example" // string | Reboot Schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.DeleteRebootSchedule1(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.DeleteRebootSchedule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRebootSchedule1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.DeleteRebootSchedule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Reboot Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRebootSchedule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUpgradeSchedule

> OperationResponseWithoutResult DeleteUpgradeSchedule(ctx, omadacId, siteId, id).Execute()

Delete upgrade schedule



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
	id := "id_example" // string | Upgrade Schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.DeleteUpgradeSchedule(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.DeleteUpgradeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUpgradeSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.DeleteUpgradeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Upgrade Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoePortsList

> OperationResponseListPoeDeviceDetail GetPoePortsList(ctx, omadacId, siteId).Execute()

Get PoE ports list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.GetPoePortsList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.GetPoePortsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoePortsList`: OperationResponseListPoeDeviceDetail
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.GetPoePortsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoePortsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListPoeDeviceDetail**](OperationResponseListPoeDeviceDetail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoeScheduleList

> OperationResponseGridVOPoeScheduleQueryOpenApiVO GetPoeScheduleList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get PoE schedule list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.GetPoeScheduleList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.GetPoeScheduleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoeScheduleList`: OperationResponseGridVOPoeScheduleQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.GetPoeScheduleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoeScheduleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOPoeScheduleQueryOpenApiVO**](OperationResponseGridVOPoeScheduleQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortScheduleList

> OperationResponseGridVOPortScheduleQueryOpenApiVO GetPortScheduleList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get port schedule list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.GetPortScheduleList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.GetPortScheduleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortScheduleList`: OperationResponseGridVOPortScheduleQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.GetPortScheduleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortScheduleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOPortScheduleQueryOpenApiVO**](OperationResponseGridVOPortScheduleQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortSchedulePorts

> OperationResponseListPortDeviceDetail GetPortSchedulePorts(ctx, omadacId, siteId).Execute()

Get device ports list supporting port schedule



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.GetPortSchedulePorts(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.GetPortSchedulePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortSchedulePorts`: OperationResponseListPortDeviceDetail
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.GetPortSchedulePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortSchedulePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListPortDeviceDetail**](OperationResponseListPortDeviceDetail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRebootScheduleList2

> OperationResponseListRebootScheduleQueryOpenApiVO GetRebootScheduleList2(ctx, omadacId, siteId).Execute()

Get reboot schedule list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.GetRebootScheduleList2(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.GetRebootScheduleList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRebootScheduleList2`: OperationResponseListRebootScheduleQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.GetRebootScheduleList2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRebootScheduleList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRebootScheduleQueryOpenApiVO**](OperationResponseListRebootScheduleQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpgradeScheduleList

> OperationResponseListUpgradeScheduleQueryOpenApiVO GetUpgradeScheduleList(ctx, omadacId, siteId).Execute()

Get upgrade schedule list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.GetUpgradeScheduleList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.GetUpgradeScheduleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpgradeScheduleList`: OperationResponseListUpgradeScheduleQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.GetUpgradeScheduleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpgradeScheduleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListUpgradeScheduleQueryOpenApiVO**](OperationResponseListUpgradeScheduleQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPoeSchedule

> OperationResponseWithoutResult ModifyPoeSchedule(ctx, omadacId, siteId, poeScheduleId).PoeScheduleOpenApiVO(poeScheduleOpenApiVO).Execute()

Modify PoE Schedule



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
	poeScheduleId := "poeScheduleId_example" // string | PoE Schedule ID
	poeScheduleOpenApiVO := *openapiclient.NewPoeScheduleOpenApiVO("Name_example", map[string][]int32{"key": []int32{int32(123)}}, false, "TurnOnTime_example") // PoeScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.ModifyPoeSchedule(context.Background(), omadacId, siteId, poeScheduleId).PoeScheduleOpenApiVO(poeScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.ModifyPoeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPoeSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.ModifyPoeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**poeScheduleId** | **string** | PoE Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPoeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **poeScheduleOpenApiVO** | [**PoeScheduleOpenApiVO**](PoeScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPortSchedule

> OperationResponseWithoutResult ModifyPortSchedule(ctx, omadacId, siteId, portScheduleId).PortScheduleOpenApiVO(portScheduleOpenApiVO).Execute()

Modify Port Schedule



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
	portScheduleId := "portScheduleId_example" // string | Port Schedule ID
	portScheduleOpenApiVO := *openapiclient.NewPortScheduleOpenApiVO("Name_example", map[string][]int32{"key": []int32{int32(123)}}, false, "TurnOnTime_example") // PortScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.ModifyPortSchedule(context.Background(), omadacId, siteId, portScheduleId).PortScheduleOpenApiVO(portScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.ModifyPortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.ModifyPortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portScheduleId** | **string** | Port Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **portScheduleOpenApiVO** | [**PortScheduleOpenApiVO**](PortScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRebootSchedule1

> OperationResponseWithoutResult ModifyRebootSchedule1(ctx, omadacId, siteId, id).RebootScheduleOpenApiVO(rebootScheduleOpenApiVO).Execute()

Modify reboot schedule



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
	id := "id_example" // string | Reboot Schedule ID
	rebootScheduleOpenApiVO := *openapiclient.NewRebootScheduleOpenApiVO([]string{"DeviceMacs_example"}, "Name_example", false, *openapiclient.NewRebootScheduleTimeOpenApiVO(int32(123), int32(123), int32(123))) // RebootScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.ModifyRebootSchedule1(context.Background(), omadacId, siteId, id).RebootScheduleOpenApiVO(rebootScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.ModifyRebootSchedule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRebootSchedule1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.ModifyRebootSchedule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Reboot Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRebootSchedule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rebootScheduleOpenApiVO** | [**RebootScheduleOpenApiVO**](RebootScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyUpgradeSchedule

> OperationResponseWithoutResult ModifyUpgradeSchedule(ctx, omadacId, siteId, id).UpgradeScheduleOpenApiVO(upgradeScheduleOpenApiVO).Execute()

Modify upgrade schedule



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
	id := "id_example" // string | Upgrade Schedule ID
	upgradeScheduleOpenApiVO := *openapiclient.NewUpgradeScheduleOpenApiVO([]string{"DeviceMacs_example"}, "Name_example", false, int32(123)) // UpgradeScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.ModifyUpgradeSchedule(context.Background(), omadacId, siteId, id).UpgradeScheduleOpenApiVO(upgradeScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.ModifyUpgradeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUpgradeSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.ModifyUpgradeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Upgrade Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUpgradeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **upgradeScheduleOpenApiVO** | [**UpgradeScheduleOpenApiVO**](UpgradeScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


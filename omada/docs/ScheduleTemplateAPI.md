# ScheduleTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPortScheduleTemplate**](ScheduleTemplateAPI.md#addportscheduletemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/port-schedules | Create a new Port Schedule Template
[**CreateRebootSchedule**](ScheduleTemplateAPI.md#createrebootschedule) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/reboot-schedules | Create new reboot schedule template
[**DeleteRebootSchedule**](ScheduleTemplateAPI.md#deleterebootschedule) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/reboot-schedules/{id} | Delete reboot schedule template
[**GetPortScheduleListTemplate**](ScheduleTemplateAPI.md#getportschedulelisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/port-schedules | Get port Schedule list
[**GetRebootScheduleList**](ScheduleTemplateAPI.md#getrebootschedulelist) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/reboot-schedules | Get reboot schedule template list
[**ModifyPortScheduleTemplate**](ScheduleTemplateAPI.md#modifyportscheduletemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/port-schedules/{type}/{portScheduleId} | Modify a Port Schedule Template
[**ModifyRebootSchedule**](ScheduleTemplateAPI.md#modifyrebootschedule) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/reboot-schedules/{id} | Modify reboot schedule template
[**RemovePortScheduleTemplate**](ScheduleTemplateAPI.md#removeportscheduletemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/port-schedules/{type}/{portScheduleId} | Delete Port Schedule Template



## AddPortScheduleTemplate

> OperationResponse AddPortScheduleTemplate(ctx, omadacId, siteTemplateId).PortOrPoeScheduleOpenApiVO(portOrPoeScheduleOpenApiVO).Execute()

Create a new Port Schedule Template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	portOrPoeScheduleOpenApiVO := *openapiclient.NewPortOrPoeScheduleOpenApiVO("Name_example", map[string][]int32{"key": []int32{int32(123)}}, int32(123), false, "TurnOnTime_example") // PortOrPoeScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleTemplateAPI.AddPortScheduleTemplate(context.Background(), omadacId, siteTemplateId).PortOrPoeScheduleOpenApiVO(portOrPoeScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.AddPortScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPortScheduleTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.AddPortScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPortScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portOrPoeScheduleOpenApiVO** | [**PortOrPoeScheduleOpenApiVO**](PortOrPoeScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRebootSchedule

> OperationResponseWithoutResult CreateRebootSchedule(ctx, omadacId, siteTemplateId).RebootScheduleTemplateOpenApiVO(rebootScheduleTemplateOpenApiVO).Execute()

Create new reboot schedule template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	rebootScheduleTemplateOpenApiVO := *openapiclient.NewRebootScheduleTemplateOpenApiVO("Name_example", false, *openapiclient.NewRebootScheduleTimeOpenApiVO(int32(123), int32(123), int32(123))) // RebootScheduleTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleTemplateAPI.CreateRebootSchedule(context.Background(), omadacId, siteTemplateId).RebootScheduleTemplateOpenApiVO(rebootScheduleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.CreateRebootSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRebootSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.CreateRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rebootScheduleTemplateOpenApiVO** | [**RebootScheduleTemplateOpenApiVO**](RebootScheduleTemplateOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRebootSchedule

> OperationResponseWithoutResult DeleteRebootSchedule(ctx, omadacId, siteTemplateId, id).Execute()

Delete reboot schedule template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	id := "id_example" // string | Reboot Schedule Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleTemplateAPI.DeleteRebootSchedule(context.Background(), omadacId, siteTemplateId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.DeleteRebootSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRebootSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.DeleteRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Reboot Schedule Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortScheduleListTemplate

> OperationResponseListPortOrPoeScheduleOpenApiVO GetPortScheduleListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get port Schedule list



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleTemplateAPI.GetPortScheduleListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.GetPortScheduleListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortScheduleListTemplate`: OperationResponseListPortOrPoeScheduleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.GetPortScheduleListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortScheduleListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListPortOrPoeScheduleOpenApiVO**](OperationResponseListPortOrPoeScheduleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRebootScheduleList

> OperationResponseListRebootScheduleTemplateQueryOpenApiVO GetRebootScheduleList(ctx, omadacId, siteTemplateId).Execute()

Get reboot schedule template list



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleTemplateAPI.GetRebootScheduleList(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.GetRebootScheduleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRebootScheduleList`: OperationResponseListRebootScheduleTemplateQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.GetRebootScheduleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRebootScheduleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRebootScheduleTemplateQueryOpenApiVO**](OperationResponseListRebootScheduleTemplateQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPortScheduleTemplate

> OperationResponseWithoutResult ModifyPortScheduleTemplate(ctx, omadacId, siteTemplateId, type_, portScheduleId).PortOrPoeScheduleOpenApiVO(portOrPoeScheduleOpenApiVO).Execute()

Modify a Port Schedule Template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	type_ := "type__example" // string | Port Schedule Type. 0: POE schedule; 1: Enable Schedule.
	portScheduleId := "portScheduleId_example" // string | Port Schedule Template ID
	portOrPoeScheduleOpenApiVO := *openapiclient.NewPortOrPoeScheduleOpenApiVO("Name_example", map[string][]int32{"key": []int32{int32(123)}}, int32(123), false, "TurnOnTime_example") // PortOrPoeScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleTemplateAPI.ModifyPortScheduleTemplate(context.Background(), omadacId, siteTemplateId, type_, portScheduleId).PortOrPoeScheduleOpenApiVO(portOrPoeScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.ModifyPortScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortScheduleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.ModifyPortScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**type_** | **string** | Port Schedule Type. 0: POE schedule; 1: Enable Schedule. | 
**portScheduleId** | **string** | Port Schedule Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **portOrPoeScheduleOpenApiVO** | [**PortOrPoeScheduleOpenApiVO**](PortOrPoeScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRebootSchedule

> OperationResponseWithoutResult ModifyRebootSchedule(ctx, omadacId, siteTemplateId, id).RebootScheduleTemplateOpenApiVO(rebootScheduleTemplateOpenApiVO).Execute()

Modify reboot schedule template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	id := "id_example" // string | Reboot Schedule Template ID
	rebootScheduleTemplateOpenApiVO := *openapiclient.NewRebootScheduleTemplateOpenApiVO("Name_example", false, *openapiclient.NewRebootScheduleTimeOpenApiVO(int32(123), int32(123), int32(123))) // RebootScheduleTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleTemplateAPI.ModifyRebootSchedule(context.Background(), omadacId, siteTemplateId, id).RebootScheduleTemplateOpenApiVO(rebootScheduleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.ModifyRebootSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRebootSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.ModifyRebootSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Reboot Schedule Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRebootScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rebootScheduleTemplateOpenApiVO** | [**RebootScheduleTemplateOpenApiVO**](RebootScheduleTemplateOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePortScheduleTemplate

> OperationResponseWithoutResult RemovePortScheduleTemplate(ctx, omadacId, siteTemplateId, type_, portScheduleId).Execute()

Delete Port Schedule Template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	type_ := "type__example" // string | Port Schedule Type. 0: POE schedule; 1: Enable Schedule.
	portScheduleId := "portScheduleId_example" // string | Port Schedule Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleTemplateAPI.RemovePortScheduleTemplate(context.Background(), omadacId, siteTemplateId, type_, portScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.RemovePortScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePortScheduleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.RemovePortScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**type_** | **string** | Port Schedule Type. 0: POE schedule; 1: Enable Schedule. | 
**portScheduleId** | **string** | Port Schedule Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePortScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


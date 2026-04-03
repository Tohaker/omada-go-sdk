# ScheduleTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRebootSchedule**](ScheduleTemplateAPI.md#createrebootschedule) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/reboot-schedules | Create new reboot schedule template
[**DeleteRebootSchedule**](ScheduleTemplateAPI.md#deleterebootschedule) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/reboot-schedules/{id} | Delete reboot schedule template
[**GetRebootScheduleList1**](ScheduleTemplateAPI.md#getrebootschedulelist1) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/reboot-schedules | Get reboot schedule template list
[**ModifyRebootSchedule**](ScheduleTemplateAPI.md#modifyrebootschedule) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/reboot-schedules/{id} | Modify reboot schedule template



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


## GetRebootScheduleList1

> OperationResponseListRebootScheduleTemplateQueryOpenApiVO GetRebootScheduleList1(ctx, omadacId, siteTemplateId).Execute()

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
	resp, r, err := apiClient.ScheduleTemplateAPI.GetRebootScheduleList1(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleTemplateAPI.GetRebootScheduleList1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRebootScheduleList1`: OperationResponseListRebootScheduleTemplateQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ScheduleTemplateAPI.GetRebootScheduleList1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRebootScheduleList1Request struct via the builder pattern


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


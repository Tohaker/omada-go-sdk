# MspLogAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlertLogsForMsp**](MspLogAPI.md#deletealertlogsformsp) | **Delete** /openapi/v1/msp/{mspId}/logs/alerts/delete | Delete MSP alert log
[**DeleteEventLogsForMsp**](MspLogAPI.md#deleteeventlogsformsp) | **Delete** /openapi/v1/msp/{mspId}/logs/events/delete | Delete MSP event log
[**ExportAuditLogListForGlobal1**](MspLogAPI.md#exportauditloglistforglobal1) | **Post** /openapi/v1/msp/{mspId}/logs/audit/export | Export audit log list in MSP view
[**ExportLogListForMsp**](MspLogAPI.md#exportloglistformsp) | **Post** /openapi/v1/msp/{mspId}/logs/export | Export log list in MSP view
[**GetAlertLogsForMsp**](MspLogAPI.md#getalertlogsformsp) | **Get** /openapi/v1/msp/{mspId}/logs/alerts | Get MSP alert log list
[**GetAuditLogSettingForMsp2**](MspLogAPI.md#getauditlogsettingformsp2) | **Get** /openapi/v1/msp/{mspId}/audit-notification | Get MSP audit log notification
[**GetAuditLogsForMsp**](MspLogAPI.md#getauditlogsformsp) | **Get** /openapi/v1/msp/{mspId}/audit-logs | Get MSP audit log list
[**GetEventLogsForMsp**](MspLogAPI.md#geteventlogsformsp) | **Get** /openapi/v1/msp/{mspId}/logs/events | Get MSP event log list
[**GetLogSettingForMsp**](MspLogAPI.md#getlogsettingformsp) | **Get** /openapi/v1/msp/{mspId}/log-notification | Get MSP log notification
[**GetLogSettingForMspV2**](MspLogAPI.md#getlogsettingformspv2) | **Get** /openapi/v1/msp/{mspId}/log-notification-setting | Get MSP log notification v2
[**ModifyAuditLogSettingMsp**](MspLogAPI.md#modifyauditlogsettingmsp) | **Patch** /openapi/v1/msp/{mspId}/audit-notification | Modify MSP audit log notification
[**ModifyLogSettingMsp**](MspLogAPI.md#modifylogsettingmsp) | **Patch** /openapi/v1/msp/{mspId}/log-notification | Modify MSP log notification
[**ModifyLogSettingMspV2**](MspLogAPI.md#modifylogsettingmspv2) | **Patch** /openapi/v1/msp/{mspId}/log-notification-setting | Modify MSP log notification v2
[**ResetLogSettingMsp**](MspLogAPI.md#resetlogsettingmsp) | **Post** /openapi/v1/msp/{mspId}/reset/log-notification | Reset MSP log notification
[**ResolveAlertForMsp**](MspLogAPI.md#resolvealertformsp) | **Post** /openapi/v1/msp/{mspId}/logs/alerts/resolve | Resolve MSP alert log



## DeleteAlertLogsForMsp

> OperationResponseWithoutResult DeleteAlertLogsForMsp(ctx, mspId).DeleteMspAlertLogListOpenApiVO(deleteMspAlertLogListOpenApiVO).Execute()

Delete MSP alert log



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
	mspId := "mspId_example" // string | MSP ID
	deleteMspAlertLogListOpenApiVO := *openapiclient.NewDeleteMspAlertLogListOpenApiVO(int64(123), "SelectType_example", int64(123)) // DeleteMspAlertLogListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.DeleteAlertLogsForMsp(context.Background(), mspId).DeleteMspAlertLogListOpenApiVO(deleteMspAlertLogListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.DeleteAlertLogsForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertLogsForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.DeleteAlertLogsForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertLogsForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteMspAlertLogListOpenApiVO** | [**DeleteMspAlertLogListOpenApiVO**](DeleteMspAlertLogListOpenApiVO.md) |  | 

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


## DeleteEventLogsForMsp

> OperationResponseWithoutResult DeleteEventLogsForMsp(ctx, mspId).DeleteMspEventLogListOpenApiVO(deleteMspEventLogListOpenApiVO).Execute()

Delete MSP event log



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
	mspId := "mspId_example" // string | MSP ID
	deleteMspEventLogListOpenApiVO := *openapiclient.NewDeleteMspEventLogListOpenApiVO(int64(123), "SelectType_example", int64(123)) // DeleteMspEventLogListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.DeleteEventLogsForMsp(context.Background(), mspId).DeleteMspEventLogListOpenApiVO(deleteMspEventLogListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.DeleteEventLogsForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEventLogsForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.DeleteEventLogsForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventLogsForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteMspEventLogListOpenApiVO** | [**DeleteMspEventLogListOpenApiVO**](DeleteMspEventLogListOpenApiVO.md) |  | 

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


## ExportAuditLogListForGlobal1

> OperationResponse ExportAuditLogListForGlobal1(ctx, mspId).ExportMspLogOpenApiVO(exportMspLogOpenApiVO).Execute()

Export audit log list in MSP view



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
	mspId := "mspId_example" // string | MSP ID
	exportMspLogOpenApiVO := *openapiclient.NewExportMspLogOpenApiVO(int32(123)) // ExportMspLogOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.ExportAuditLogListForGlobal1(context.Background(), mspId).ExportMspLogOpenApiVO(exportMspLogOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.ExportAuditLogListForGlobal1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAuditLogListForGlobal1`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.ExportAuditLogListForGlobal1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAuditLogListForGlobal1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportMspLogOpenApiVO** | [**ExportMspLogOpenApiVO**](ExportMspLogOpenApiVO.md) |  | 

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


## ExportLogListForMsp

> OperationResponse ExportLogListForMsp(ctx, mspId).ExportMspLogOpenApiVO(exportMspLogOpenApiVO).Execute()

Export log list in MSP view



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
	mspId := "mspId_example" // string | MSP ID
	exportMspLogOpenApiVO := *openapiclient.NewExportMspLogOpenApiVO(int32(123)) // ExportMspLogOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.ExportLogListForMsp(context.Background(), mspId).ExportMspLogOpenApiVO(exportMspLogOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.ExportLogListForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLogListForMsp`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.ExportLogListForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLogListForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportMspLogOpenApiVO** | [**ExportMspLogOpenApiVO**](ExportMspLogOpenApiVO.md) |  | 

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


## GetAlertLogsForMsp

> OperationResponseAlertLogGridVOAlertLogOpenApiVO GetAlertLogsForMsp(ctx, mspId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).FiltersResolved(filtersResolved).Execute()

Get MSP alert log list



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
	mspId := "mspId_example" // string | MSP ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field 1679297710438
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field 1681889710438
	filtersModule := "filtersModule_example" // string | Filter query parameters, support field module, it should be a value as follows: System, Device (optional)
	filtersResolved := true // bool | Filter query parameters, support field resolved, it should be a value as follows: true, false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.GetAlertLogsForMsp(context.Background(), mspId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).FiltersResolved(filtersResolved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.GetAlertLogsForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertLogsForMsp`: OperationResponseAlertLogGridVOAlertLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.GetAlertLogsForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertLogsForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersTimeStart** | **int64** | Filter query parameters, support field 1679297710438 | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field 1681889710438 | 
 **filtersModule** | **string** | Filter query parameters, support field module, it should be a value as follows: System, Device | 
 **filtersResolved** | **bool** | Filter query parameters, support field resolved, it should be a value as follows: true, false | 

### Return type

[**OperationResponseAlertLogGridVOAlertLogOpenApiVO**](OperationResponseAlertLogGridVOAlertLogOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogSettingForMsp2

> OperationResponseAuditLogNotificationSettingOpenApiVO GetAuditLogSettingForMsp2(ctx, mspId).Execute()

Get MSP audit log notification



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
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.GetAuditLogSettingForMsp2(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.GetAuditLogSettingForMsp2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogSettingForMsp2`: OperationResponseAuditLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.GetAuditLogSettingForMsp2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogSettingForMsp2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseAuditLogNotificationSettingOpenApiVO**](OperationResponseAuditLogNotificationSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogsForMsp

> OperationResponseGridVOAuditLogOpenApiVO GetAuditLogsForMsp(ctx, mspId).Page(page).PageSize(pageSize).SortsTime(sortsTime).FiltersResult(filtersResult).FiltersLevel(filtersLevel).FiltersAuditTypes(filtersAuditTypes).FiltersTimes(filtersTimes).SearchKey(searchKey).Execute()

Get MSP audit log list



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
	mspId := "mspId_example" // string | MSP ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	sortsTime := "sortsTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersResult := int32(56) // int32 | Filter query parameters, support field result,it should be a value as follows: 0: successful; 1: failed,example:0 (optional)
	filtersLevel := "filtersLevel_example" // string | Filter query parameters, support field level,it should be a value as follows: Error; Warning; Information, example:Error (optional)
	filtersAuditTypes := "filtersAuditTypes_example" // string | Filter query parameters, support field auditTypes, for the values of auditLog type, refer to section 5.2.2 of the Open API Access Guide, example:Log,Cloud Access,User Interface. (optional)
	filtersTimes := "filtersTimes_example" // string | Filter query parameters, support field times, example:[{\"timeStart\":1678060800000,\"timeEnd\":1678665599999}](UrlEncode:%5B%7B%22timeStart%22%3A1678060800000%2C%22timeEnd%22%3A1678665599999%7D%5D).If this parameter is not specified (not included or empty array), the interface will query data within the default time period: [{\"timeStart\":  Current timestamp minus milliseconds of 7 days,\"timeEnd\": Current timestamp}]. (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field content (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.GetAuditLogsForMsp(context.Background(), mspId).Page(page).PageSize(pageSize).SortsTime(sortsTime).FiltersResult(filtersResult).FiltersLevel(filtersLevel).FiltersAuditTypes(filtersAuditTypes).FiltersTimes(filtersTimes).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.GetAuditLogsForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogsForMsp`: OperationResponseGridVOAuditLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.GetAuditLogsForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **sortsTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersResult** | **int32** | Filter query parameters, support field result,it should be a value as follows: 0: successful; 1: failed,example:0 | 
 **filtersLevel** | **string** | Filter query parameters, support field level,it should be a value as follows: Error; Warning; Information, example:Error | 
 **filtersAuditTypes** | **string** | Filter query parameters, support field auditTypes, for the values of auditLog type, refer to section 5.2.2 of the Open API Access Guide, example:Log,Cloud Access,User Interface. | 
 **filtersTimes** | **string** | Filter query parameters, support field times, example:[{\&quot;timeStart\&quot;:1678060800000,\&quot;timeEnd\&quot;:1678665599999}](UrlEncode:%5B%7B%22timeStart%22%3A1678060800000%2C%22timeEnd%22%3A1678665599999%7D%5D).If this parameter is not specified (not included or empty array), the interface will query data within the default time period: [{\&quot;timeStart\&quot;:  Current timestamp minus milliseconds of 7 days,\&quot;timeEnd\&quot;: Current timestamp}]. | 
 **searchKey** | **string** | Fuzzy query parameters, support field content | 

### Return type

[**OperationResponseGridVOAuditLogOpenApiVO**](OperationResponseGridVOAuditLogOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventLogsForMsp

> OperationResponseEventLogGridVOEventLogOpenApiVO GetEventLogsForMsp(ctx, mspId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).Execute()

Get MSP event log list



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
	mspId := "mspId_example" // string | MSP ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field 1679297710438
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field 1681889710438
	filtersModule := "filtersModule_example" // string | Filter query parameters, support field module, it should be a value as follows: System, Device (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.GetEventLogsForMsp(context.Background(), mspId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.GetEventLogsForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventLogsForMsp`: OperationResponseEventLogGridVOEventLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.GetEventLogsForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogsForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersTimeStart** | **int64** | Filter query parameters, support field 1679297710438 | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field 1681889710438 | 
 **filtersModule** | **string** | Filter query parameters, support field module, it should be a value as follows: System, Device | 

### Return type

[**OperationResponseEventLogGridVOEventLogOpenApiVO**](OperationResponseEventLogGridVOEventLogOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogSettingForMsp

> OperationResponseLogNotificationSettingOpenApiVO GetLogSettingForMsp(ctx, mspId).Execute()

Get MSP log notification



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
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.GetLogSettingForMsp(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.GetLogSettingForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogSettingForMsp`: OperationResponseLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.GetLogSettingForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSettingForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseLogNotificationSettingOpenApiVO**](OperationResponseLogNotificationSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogSettingForMspV2

> OperationResponseLogNotificationSettingOpenApiVO GetLogSettingForMspV2(ctx, mspId).Execute()

Get MSP log notification v2



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
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.GetLogSettingForMspV2(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.GetLogSettingForMspV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogSettingForMspV2`: OperationResponseLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.GetLogSettingForMspV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSettingForMspV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseLogNotificationSettingOpenApiVO**](OperationResponseLogNotificationSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAuditLogSettingMsp

> OperationResponseWithoutResult ModifyAuditLogSettingMsp(ctx, mspId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()

Modify MSP audit log notification



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
	mspId := "mspId_example" // string | MSP ID
	auditLogNotificationSettingEditOpenApiVO := *openapiclient.NewAuditLogNotificationSettingEditOpenApiVO([]openapiclient.AuditLogNotificationEditOpenApiVO{*openapiclient.NewAuditLogNotificationEditOpenApiVO("DASHBOARD", false)}, *openapiclient.NewWebhookConfigEditOpenApiVO(false)) // AuditLogNotificationSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.ModifyAuditLogSettingMsp(context.Background(), mspId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.ModifyAuditLogSettingMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAuditLogSettingMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.ModifyAuditLogSettingMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAuditLogSettingMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditLogNotificationSettingEditOpenApiVO** | [**AuditLogNotificationSettingEditOpenApiVO**](AuditLogNotificationSettingEditOpenApiVO.md) |  | 

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


## ModifyLogSettingMsp

> OperationResponseWithoutResult ModifyLogSettingMsp(ctx, mspId).LogNotificationSettingEditOpenApiVO(logNotificationSettingEditOpenApiVO).Execute()

Modify MSP log notification



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
	mspId := "mspId_example" // string | MSP ID
	logNotificationSettingEditOpenApiVO := *openapiclient.NewLogNotificationSettingEditOpenApiVO([]openapiclient.LogNotificationEditOpenApiVO{*openapiclient.NewLogNotificationEditOpenApiVO(false, true, true, "LOGIN_OK")}) // LogNotificationSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.ModifyLogSettingMsp(context.Background(), mspId).LogNotificationSettingEditOpenApiVO(logNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.ModifyLogSettingMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLogSettingMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.ModifyLogSettingMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLogSettingMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logNotificationSettingEditOpenApiVO** | [**LogNotificationSettingEditOpenApiVO**](LogNotificationSettingEditOpenApiVO.md) |  | 

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


## ModifyLogSettingMspV2

> OperationResponseWithoutResult ModifyLogSettingMspV2(ctx, mspId).LogNotificationSettingEditOpenApiV2VO(logNotificationSettingEditOpenApiV2VO).Execute()

Modify MSP log notification v2



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
	mspId := "mspId_example" // string | MSP ID
	logNotificationSettingEditOpenApiV2VO := *openapiclient.NewLogNotificationSettingEditOpenApiV2VO() // LogNotificationSettingEditOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.ModifyLogSettingMspV2(context.Background(), mspId).LogNotificationSettingEditOpenApiV2VO(logNotificationSettingEditOpenApiV2VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.ModifyLogSettingMspV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLogSettingMspV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.ModifyLogSettingMspV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLogSettingMspV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logNotificationSettingEditOpenApiV2VO** | [**LogNotificationSettingEditOpenApiV2VO**](LogNotificationSettingEditOpenApiV2VO.md) |  | 

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


## ResetLogSettingMsp

> OperationResponseWithoutResult ResetLogSettingMsp(ctx, mspId).Execute()

Reset MSP log notification



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
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.ResetLogSettingMsp(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.ResetLogSettingMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetLogSettingMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.ResetLogSettingMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetLogSettingMspRequest struct via the builder pattern


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


## ResolveAlertForMsp

> OperationResponseWithoutResult ResolveAlertForMsp(ctx, mspId).ResolveMspLogListOpenApiVO(resolveMspLogListOpenApiVO).Execute()

Resolve MSP alert log



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
	mspId := "mspId_example" // string | MSP ID
	resolveMspLogListOpenApiVO := *openapiclient.NewResolveMspLogListOpenApiVO(int64(123), "SelectType_example", int64(123)) // ResolveMspLogListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLogAPI.ResolveAlertForMsp(context.Background(), mspId).ResolveMspLogListOpenApiVO(resolveMspLogListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLogAPI.ResolveAlertForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAlertForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLogAPI.ResolveAlertForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveAlertForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolveMspLogListOpenApiVO** | [**ResolveMspLogListOpenApiVO**](ResolveMspLogListOpenApiVO.md) |  | 

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


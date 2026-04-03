# LogAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlertLogsForGlobal**](LogAPI.md#deletealertlogsforglobal) | **Delete** /openapi/v1/{omadacId}/logs/alerts/delete | Delete global alert log
[**DeleteAlertLogsForSite**](LogAPI.md#deletealertlogsforsite) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/logs/alerts/delete | Delete site alert log
[**DeleteEventLogsForGlobal**](LogAPI.md#deleteeventlogsforglobal) | **Delete** /openapi/v1/{omadacId}/logs/events/delete | Delete global event log
[**DeleteEventLogsForSite**](LogAPI.md#deleteeventlogsforsite) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/logs/events/delete | Delete site event log
[**ExportAuditLogListForGlobal**](LogAPI.md#exportauditloglistforglobal) | **Post** /openapi/v1/{omadacId}/logs/audit/export | Export audit log list in global view
[**ExportLogListForGlobal**](LogAPI.md#exportloglistforglobal) | **Post** /openapi/v1/{omadacId}/logs/export | Export log list in global view
[**GetAlertLogsForGlobal**](LogAPI.md#getalertlogsforglobal) | **Get** /openapi/v1/{omadacId}/logs/alerts | Get global alert log list
[**GetAlertLogsForSite**](LogAPI.md#getalertlogsforsite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/logs/alerts | Get site alert log list
[**GetAuditLogSettingForGlobal**](LogAPI.md#getauditlogsettingforglobal) | **Get** /openapi/v1/{omadacId}/audit-notification | Get global audit log notification
[**GetAuditLogSettingForMsp1**](LogAPI.md#getauditlogsettingformsp1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/site/audit-notification | Get site audit log notification
[**GetAuditLogsForGlobal**](LogAPI.md#getauditlogsforglobal) | **Get** /openapi/v1/{omadacId}/audit-logs | Get global audit log list
[**GetAuditLogsForSite**](LogAPI.md#getauditlogsforsite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/audit-logs | Get site audit log list
[**GetEventLogsForGlobal**](LogAPI.md#geteventlogsforglobal) | **Get** /openapi/v1/{omadacId}/logs/events | Get global event log list
[**GetEventLogsForSite**](LogAPI.md#geteventlogsforsite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/logs/events | Get site event log list
[**GetLogSettingForGlobal**](LogAPI.md#getlogsettingforglobal) | **Get** /openapi/v1/{omadacId}/log-notification | Get global log notification
[**GetLogSettingForGlobalV2**](LogAPI.md#getlogsettingforglobalv2) | **Get** /openapi/v2/{omadacId}/log-notification | Get global log notification v2
[**GetLogSettingForSite**](LogAPI.md#getlogsettingforsite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/site/log-notification | Get site log notification
[**GetLogSettingForSiteV2**](LogAPI.md#getlogsettingforsitev2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/site/log-notification | Get site log notification v2
[**GetRemoteLoggingSettingTip**](LogAPI.md#getremoteloggingsettingtip) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/remote-logging/tip | Get site remote logging setting tip
[**GetRemoteLoggingTip**](LogAPI.md#getremoteloggingtip) | **Get** /openapi/v1/{omadacId}/global/controller/setting/syslog/tip | Get customer remote logging tip
[**ModifyAuditLogSettingGlobal**](LogAPI.md#modifyauditlogsettingglobal) | **Patch** /openapi/v1/{omadacId}/audit-notification | Modify global audit log notification
[**ModifyAuditLogSettingSite1**](LogAPI.md#modifyauditlogsettingsite1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/site/audit-notification | Modify site audit log notification
[**ModifyLogSettingGlobal**](LogAPI.md#modifylogsettingglobal) | **Patch** /openapi/v1/{omadacId}/log-notification | Modify global log notification
[**ModifyLogSettingGlobalV2**](LogAPI.md#modifylogsettingglobalv2) | **Patch** /openapi/v2/{omadacId}/log-notification | Modify global log notification v2
[**ModifyLogSettingSite**](LogAPI.md#modifylogsettingsite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/site/log-notification | Modify site log notification
[**ModifyLogSettingSiteV2**](LogAPI.md#modifylogsettingsitev2) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/site/log-notification | Modify site log notification v2
[**ResetLogSettingGlobal**](LogAPI.md#resetlogsettingglobal) | **Post** /openapi/v1/{omadacId}/reset/log-notification | Reset global log notification
[**ResetLogSettingSite**](LogAPI.md#resetlogsettingsite) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/site/reset/log-notification | Reset site log notification
[**ResolveAlertForSite**](LogAPI.md#resolvealertforsite) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/logs/alerts/resolve | Resolve site alert log



## DeleteAlertLogsForGlobal

> OperationResponseWithoutResult DeleteAlertLogsForGlobal(ctx, omadacId).DeleteGlobalAlertLogListOpenApiVO(deleteGlobalAlertLogListOpenApiVO).Execute()

Delete global alert log



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
	deleteGlobalAlertLogListOpenApiVO := *openapiclient.NewDeleteGlobalAlertLogListOpenApiVO(int64(123), "SelectType_example", int64(123)) // DeleteGlobalAlertLogListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.DeleteAlertLogsForGlobal(context.Background(), omadacId).DeleteGlobalAlertLogListOpenApiVO(deleteGlobalAlertLogListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.DeleteAlertLogsForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertLogsForGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.DeleteAlertLogsForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertLogsForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteGlobalAlertLogListOpenApiVO** | [**DeleteGlobalAlertLogListOpenApiVO**](DeleteGlobalAlertLogListOpenApiVO.md) |  | 

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


## DeleteAlertLogsForSite

> OperationResponseWithoutResult DeleteAlertLogsForSite(ctx, omadacId, siteId).DeleteSiteAlertLogListOpenApiVO(deleteSiteAlertLogListOpenApiVO).Execute()

Delete site alert log



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
	deleteSiteAlertLogListOpenApiVO := *openapiclient.NewDeleteSiteAlertLogListOpenApiVO(int64(123), "SelectType_example", int64(123)) // DeleteSiteAlertLogListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.DeleteAlertLogsForSite(context.Background(), omadacId, siteId).DeleteSiteAlertLogListOpenApiVO(deleteSiteAlertLogListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.DeleteAlertLogsForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertLogsForSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.DeleteAlertLogsForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertLogsForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteSiteAlertLogListOpenApiVO** | [**DeleteSiteAlertLogListOpenApiVO**](DeleteSiteAlertLogListOpenApiVO.md) |  | 

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


## DeleteEventLogsForGlobal

> OperationResponseWithoutResult DeleteEventLogsForGlobal(ctx, omadacId).DeleteGlobalEventLogListOpenApiVO(deleteGlobalEventLogListOpenApiVO).Execute()

Delete global event log



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
	deleteGlobalEventLogListOpenApiVO := *openapiclient.NewDeleteGlobalEventLogListOpenApiVO(int64(123), "SelectType_example", int64(123)) // DeleteGlobalEventLogListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.DeleteEventLogsForGlobal(context.Background(), omadacId).DeleteGlobalEventLogListOpenApiVO(deleteGlobalEventLogListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.DeleteEventLogsForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEventLogsForGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.DeleteEventLogsForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventLogsForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteGlobalEventLogListOpenApiVO** | [**DeleteGlobalEventLogListOpenApiVO**](DeleteGlobalEventLogListOpenApiVO.md) |  | 

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


## DeleteEventLogsForSite

> OperationResponseWithoutResult DeleteEventLogsForSite(ctx, omadacId, siteId).DeleteSiteEventLogListOpenApiVO(deleteSiteEventLogListOpenApiVO).Execute()

Delete site event log



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
	deleteSiteEventLogListOpenApiVO := *openapiclient.NewDeleteSiteEventLogListOpenApiVO(int64(123), "SelectType_example", int64(123)) // DeleteSiteEventLogListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.DeleteEventLogsForSite(context.Background(), omadacId, siteId).DeleteSiteEventLogListOpenApiVO(deleteSiteEventLogListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.DeleteEventLogsForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEventLogsForSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.DeleteEventLogsForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventLogsForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteSiteEventLogListOpenApiVO** | [**DeleteSiteEventLogListOpenApiVO**](DeleteSiteEventLogListOpenApiVO.md) |  | 

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


## ExportAuditLogListForGlobal

> OperationResponse ExportAuditLogListForGlobal(ctx, omadacId).ExportLogOpenApiVO(exportLogOpenApiVO).Execute()

Export audit log list in global view



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
	exportLogOpenApiVO := *openapiclient.NewExportLogOpenApiVO(int32(123), []string{"SiteIds_example"}) // ExportLogOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ExportAuditLogListForGlobal(context.Background(), omadacId).ExportLogOpenApiVO(exportLogOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ExportAuditLogListForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAuditLogListForGlobal`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ExportAuditLogListForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAuditLogListForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportLogOpenApiVO** | [**ExportLogOpenApiVO**](ExportLogOpenApiVO.md) |  | 

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


## ExportLogListForGlobal

> OperationResponse ExportLogListForGlobal(ctx, omadacId).ExportLogOpenApiVO(exportLogOpenApiVO).Execute()

Export log list in global view



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
	exportLogOpenApiVO := *openapiclient.NewExportLogOpenApiVO(int32(123), []string{"SiteIds_example"}) // ExportLogOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ExportLogListForGlobal(context.Background(), omadacId).ExportLogOpenApiVO(exportLogOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ExportLogListForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLogListForGlobal`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ExportLogListForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLogListForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportLogOpenApiVO** | [**ExportLogOpenApiVO**](ExportLogOpenApiVO.md) |  | 

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


## GetAlertLogsForGlobal

> OperationResponseAlertLogGridVOAlertLogOpenApiVO GetAlertLogsForGlobal(ctx, omadacId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).FiltersResolved(filtersResolved).Execute()

Get global alert log list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field 1679297710438
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field 1681889710438
	filtersModule := "filtersModule_example" // string | Filter query parameters, support field module, it should be a value as follows: System, Device (optional)
	filtersResolved := true // bool | Filter query parameters, support field resolved, it should be a value as follows: true, false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetAlertLogsForGlobal(context.Background(), omadacId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).FiltersResolved(filtersResolved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetAlertLogsForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertLogsForGlobal`: OperationResponseAlertLogGridVOAlertLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetAlertLogsForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertLogsForGlobalRequest struct via the builder pattern


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


## GetAlertLogsForSite

> OperationResponseAlertLogGridVOAlertLogOpenApiVO GetAlertLogsForSite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).FiltersResolved(filtersResolved).Execute()

Get site alert log list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field 1679297710438
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field 1681889710438
	filtersModule := "filtersModule_example" // string | Filter query parameters, support field module, it should be a value as follows: System, Device, Client (optional)
	filtersResolved := true // bool | Filter query parameters, support field resolved, it should be a value as follows: true, false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetAlertLogsForSite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).FiltersResolved(filtersResolved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetAlertLogsForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertLogsForSite`: OperationResponseAlertLogGridVOAlertLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetAlertLogsForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertLogsForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersTimeStart** | **int64** | Filter query parameters, support field 1679297710438 | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field 1681889710438 | 
 **filtersModule** | **string** | Filter query parameters, support field module, it should be a value as follows: System, Device, Client | 
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


## GetAuditLogSettingForGlobal

> OperationResponseAuditLogNotificationSettingOpenApiVO GetAuditLogSettingForGlobal(ctx, omadacId).Execute()

Get global audit log notification



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetAuditLogSettingForGlobal(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetAuditLogSettingForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogSettingForGlobal`: OperationResponseAuditLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetAuditLogSettingForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogSettingForGlobalRequest struct via the builder pattern


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


## GetAuditLogSettingForMsp1

> OperationResponseAuditLogNotificationSettingOpenApiVO GetAuditLogSettingForMsp1(ctx, omadacId, siteId).Execute()

Get site audit log notification



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
	resp, r, err := apiClient.LogAPI.GetAuditLogSettingForMsp1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetAuditLogSettingForMsp1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogSettingForMsp1`: OperationResponseAuditLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetAuditLogSettingForMsp1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogSettingForMsp1Request struct via the builder pattern


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


## GetAuditLogsForGlobal

> OperationResponseGridVOAuditLogOpenApiVO GetAuditLogsForGlobal(ctx, omadacId).Page(page).PageSize(pageSize).SortsTime(sortsTime).FiltersResult(filtersResult).FiltersLevel(filtersLevel).FiltersAuditTypes(filtersAuditTypes).FiltersTimes(filtersTimes).SearchKey(searchKey).Execute()

Get global audit log list



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
	resp, r, err := apiClient.LogAPI.GetAuditLogsForGlobal(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsTime(sortsTime).FiltersResult(filtersResult).FiltersLevel(filtersLevel).FiltersAuditTypes(filtersAuditTypes).FiltersTimes(filtersTimes).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetAuditLogsForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogsForGlobal`: OperationResponseGridVOAuditLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetAuditLogsForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsForGlobalRequest struct via the builder pattern


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


## GetAuditLogsForSite

> OperationResponseGridVOAuditLogOpenApiVO GetAuditLogsForSite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsTime(sortsTime).FiltersResult(filtersResult).FiltersLevel(filtersLevel).FiltersAuditTypes(filtersAuditTypes).FiltersTimes(filtersTimes).SearchKey(searchKey).Execute()

Get site audit log list



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
	resp, r, err := apiClient.LogAPI.GetAuditLogsForSite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsTime(sortsTime).FiltersResult(filtersResult).FiltersLevel(filtersLevel).FiltersAuditTypes(filtersAuditTypes).FiltersTimes(filtersTimes).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetAuditLogsForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogsForSite`: OperationResponseGridVOAuditLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetAuditLogsForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsForSiteRequest struct via the builder pattern


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


## GetEventLogsForGlobal

> OperationResponseEventLogGridVOEventLogOpenApiVO GetEventLogsForGlobal(ctx, omadacId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).Execute()

Get global event log list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field 1679297710438
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field 1681889710438
	filtersModule := "filtersModule_example" // string | Filter query parameters, support field module, it should be a value as follows: System, Device (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetEventLogsForGlobal(context.Background(), omadacId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetEventLogsForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventLogsForGlobal`: OperationResponseEventLogGridVOEventLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetEventLogsForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogsForGlobalRequest struct via the builder pattern


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


## GetEventLogsForSite

> OperationResponseEventLogGridVOEventLogOpenApiVO GetEventLogsForSite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).Execute()

Get site event log list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field 1679297710438
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field 1681889710438
	filtersModule := "filtersModule_example" // string | Filter query parameters, support field module, it should be a value as follows: System, Device, Client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetEventLogsForSite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersModule(filtersModule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetEventLogsForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventLogsForSite`: OperationResponseEventLogGridVOEventLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetEventLogsForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogsForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersTimeStart** | **int64** | Filter query parameters, support field 1679297710438 | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field 1681889710438 | 
 **filtersModule** | **string** | Filter query parameters, support field module, it should be a value as follows: System, Device, Client | 

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


## GetLogSettingForGlobal

> OperationResponseLogNotificationSettingOpenApiVO GetLogSettingForGlobal(ctx, omadacId).Execute()

Get global log notification



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetLogSettingForGlobal(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetLogSettingForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogSettingForGlobal`: OperationResponseLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetLogSettingForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSettingForGlobalRequest struct via the builder pattern


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


## GetLogSettingForGlobalV2

> OperationResponseLogNotificationSettingOpenApiVO GetLogSettingForGlobalV2(ctx, omadacId).Execute()

Get global log notification v2



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetLogSettingForGlobalV2(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetLogSettingForGlobalV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogSettingForGlobalV2`: OperationResponseLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetLogSettingForGlobalV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSettingForGlobalV2Request struct via the builder pattern


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


## GetLogSettingForSite

> OperationResponseLogNotificationSettingOpenApiVO GetLogSettingForSite(ctx, omadacId, siteId).Execute()

Get site log notification



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
	resp, r, err := apiClient.LogAPI.GetLogSettingForSite(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetLogSettingForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogSettingForSite`: OperationResponseLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetLogSettingForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSettingForSiteRequest struct via the builder pattern


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


## GetLogSettingForSiteV2

> OperationResponseLogNotificationSettingOpenApiVO GetLogSettingForSiteV2(ctx, omadacId, siteId).Execute()

Get site log notification v2



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
	resp, r, err := apiClient.LogAPI.GetLogSettingForSiteV2(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetLogSettingForSiteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogSettingForSiteV2`: OperationResponseLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetLogSettingForSiteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSettingForSiteV2Request struct via the builder pattern


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


## GetRemoteLoggingSettingTip

> OperationResponseSiteRemoteLoggingSetting GetRemoteLoggingSettingTip(ctx, omadacId, siteId).Execute()

Get site remote logging setting tip



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetRemoteLoggingSettingTip(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetRemoteLoggingSettingTip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteLoggingSettingTip`: OperationResponseSiteRemoteLoggingSetting
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetRemoteLoggingSettingTip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteLoggingSettingTipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteRemoteLoggingSetting**](OperationResponseSiteRemoteLoggingSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteLoggingTip

> OperationResponseCustomerRemoteLogTipOpenApiVO GetRemoteLoggingTip(ctx, omadacId).Execute()

Get customer remote logging tip



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetRemoteLoggingTip(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetRemoteLoggingTip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteLoggingTip`: OperationResponseCustomerRemoteLogTipOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetRemoteLoggingTip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteLoggingTipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseCustomerRemoteLogTipOpenApiVO**](OperationResponseCustomerRemoteLogTipOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAuditLogSettingGlobal

> OperationResponseWithoutResult ModifyAuditLogSettingGlobal(ctx, omadacId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()

Modify global audit log notification



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
	auditLogNotificationSettingEditOpenApiVO := *openapiclient.NewAuditLogNotificationSettingEditOpenApiVO([]openapiclient.AuditLogNotificationEditOpenApiVO{*openapiclient.NewAuditLogNotificationEditOpenApiVO("DASHBOARD", false)}, *openapiclient.NewWebhookConfigEditOpenApiVO(false)) // AuditLogNotificationSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ModifyAuditLogSettingGlobal(context.Background(), omadacId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ModifyAuditLogSettingGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAuditLogSettingGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ModifyAuditLogSettingGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAuditLogSettingGlobalRequest struct via the builder pattern


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


## ModifyAuditLogSettingSite1

> OperationResponseWithoutResult ModifyAuditLogSettingSite1(ctx, omadacId, siteId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()

Modify site audit log notification



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
	auditLogNotificationSettingEditOpenApiVO := *openapiclient.NewAuditLogNotificationSettingEditOpenApiVO([]openapiclient.AuditLogNotificationEditOpenApiVO{*openapiclient.NewAuditLogNotificationEditOpenApiVO("DASHBOARD", false)}, *openapiclient.NewWebhookConfigEditOpenApiVO(false)) // AuditLogNotificationSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ModifyAuditLogSettingSite1(context.Background(), omadacId, siteId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ModifyAuditLogSettingSite1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAuditLogSettingSite1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ModifyAuditLogSettingSite1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAuditLogSettingSite1Request struct via the builder pattern


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


## ModifyLogSettingGlobal

> OperationResponseWithoutResult ModifyLogSettingGlobal(ctx, omadacId).LogNotificationSettingEditOpenApiVO(logNotificationSettingEditOpenApiVO).Execute()

Modify global log notification



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
	logNotificationSettingEditOpenApiVO := *openapiclient.NewLogNotificationSettingEditOpenApiVO([]openapiclient.LogNotificationEditOpenApiVO{*openapiclient.NewLogNotificationEditOpenApiVO(false, true, true, "LOGIN_OK")}) // LogNotificationSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ModifyLogSettingGlobal(context.Background(), omadacId).LogNotificationSettingEditOpenApiVO(logNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ModifyLogSettingGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLogSettingGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ModifyLogSettingGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLogSettingGlobalRequest struct via the builder pattern


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


## ModifyLogSettingGlobalV2

> OperationResponseWithoutResult ModifyLogSettingGlobalV2(ctx, omadacId).LogNotificationSettingEditOpenApiV2VO(logNotificationSettingEditOpenApiV2VO).Execute()

Modify global log notification v2



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
	logNotificationSettingEditOpenApiV2VO := *openapiclient.NewLogNotificationSettingEditOpenApiV2VO() // LogNotificationSettingEditOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ModifyLogSettingGlobalV2(context.Background(), omadacId).LogNotificationSettingEditOpenApiV2VO(logNotificationSettingEditOpenApiV2VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ModifyLogSettingGlobalV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLogSettingGlobalV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ModifyLogSettingGlobalV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLogSettingGlobalV2Request struct via the builder pattern


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


## ModifyLogSettingSite

> OperationResponseWithoutResult ModifyLogSettingSite(ctx, omadacId, siteId).LogNotificationSettingEditOpenApiVO(logNotificationSettingEditOpenApiVO).Execute()

Modify site log notification



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
	logNotificationSettingEditOpenApiVO := *openapiclient.NewLogNotificationSettingEditOpenApiVO([]openapiclient.LogNotificationEditOpenApiVO{*openapiclient.NewLogNotificationEditOpenApiVO(false, true, true, "LOGIN_OK")}) // LogNotificationSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ModifyLogSettingSite(context.Background(), omadacId, siteId).LogNotificationSettingEditOpenApiVO(logNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ModifyLogSettingSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLogSettingSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ModifyLogSettingSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLogSettingSiteRequest struct via the builder pattern


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


## ModifyLogSettingSiteV2

> OperationResponseWithoutResult ModifyLogSettingSiteV2(ctx, omadacId, siteId).LogNotificationSettingEditOpenApiV2VO(logNotificationSettingEditOpenApiV2VO).Execute()

Modify site log notification v2



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
	logNotificationSettingEditOpenApiV2VO := *openapiclient.NewLogNotificationSettingEditOpenApiV2VO() // LogNotificationSettingEditOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ModifyLogSettingSiteV2(context.Background(), omadacId, siteId).LogNotificationSettingEditOpenApiV2VO(logNotificationSettingEditOpenApiV2VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ModifyLogSettingSiteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLogSettingSiteV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ModifyLogSettingSiteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLogSettingSiteV2Request struct via the builder pattern


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


## ResetLogSettingGlobal

> OperationResponseWithoutResult ResetLogSettingGlobal(ctx, omadacId).Execute()

Reset global log notification



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ResetLogSettingGlobal(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ResetLogSettingGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetLogSettingGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ResetLogSettingGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetLogSettingGlobalRequest struct via the builder pattern


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


## ResetLogSettingSite

> OperationResponseWithoutResult ResetLogSettingSite(ctx, omadacId, siteId).Execute()

Reset site log notification



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
	resp, r, err := apiClient.LogAPI.ResetLogSettingSite(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ResetLogSettingSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetLogSettingSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ResetLogSettingSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetLogSettingSiteRequest struct via the builder pattern


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


## ResolveAlertForSite

> OperationResponseWithoutResult ResolveAlertForSite(ctx, omadacId, siteId).ResolveSiteLogListOpenApiVO(resolveSiteLogListOpenApiVO).Execute()

Resolve site alert log



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
	resolveSiteLogListOpenApiVO := *openapiclient.NewResolveSiteLogListOpenApiVO(int64(123), "SelectType_example", int64(123)) // ResolveSiteLogListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.ResolveAlertForSite(context.Background(), omadacId, siteId).ResolveSiteLogListOpenApiVO(resolveSiteLogListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.ResolveAlertForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAlertForSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.ResolveAlertForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveAlertForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolveSiteLogListOpenApiVO** | [**ResolveSiteLogListOpenApiVO**](ResolveSiteLogListOpenApiVO.md) |  | 

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


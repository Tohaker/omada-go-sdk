# AuditLogAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAuditLogListForGlobal**](AuditLogAPI.md#exportauditloglistforglobal) | **Post** /openapi/v1/{omadacId}/logs/audit/export | Export audit log list in global view
[**GetAuditLogSettingForGlobal**](AuditLogAPI.md#getauditlogsettingforglobal) | **Get** /openapi/v1/{omadacId}/audit-notification | Get global audit log notification
[**GetAuditLogSettingForMsp1**](AuditLogAPI.md#getauditlogsettingformsp1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/site/audit-notification | Get site audit log notification
[**GetAuditLogsForGlobal**](AuditLogAPI.md#getauditlogsforglobal) | **Get** /openapi/v1/{omadacId}/audit-logs | Get global audit log list
[**GetAuditLogsForSite**](AuditLogAPI.md#getauditlogsforsite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/audit-logs | Get site audit log list
[**ModifyAuditLogSettingGlobal**](AuditLogAPI.md#modifyauditlogsettingglobal) | **Patch** /openapi/v1/{omadacId}/audit-notification | Modify global audit log notification
[**ModifyAuditLogSettingSite1**](AuditLogAPI.md#modifyauditlogsettingsite1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/site/audit-notification | Modify site audit log notification



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
	resp, r, err := apiClient.AuditLogAPI.ExportAuditLogListForGlobal(context.Background(), omadacId).ExportLogOpenApiVO(exportLogOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.ExportAuditLogListForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAuditLogListForGlobal`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.ExportAuditLogListForGlobal`: %v\n", resp)
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
	resp, r, err := apiClient.AuditLogAPI.GetAuditLogSettingForGlobal(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.GetAuditLogSettingForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogSettingForGlobal`: OperationResponseAuditLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.GetAuditLogSettingForGlobal`: %v\n", resp)
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
	resp, r, err := apiClient.AuditLogAPI.GetAuditLogSettingForMsp1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.GetAuditLogSettingForMsp1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogSettingForMsp1`: OperationResponseAuditLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.GetAuditLogSettingForMsp1`: %v\n", resp)
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
	resp, r, err := apiClient.AuditLogAPI.GetAuditLogsForGlobal(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsTime(sortsTime).FiltersResult(filtersResult).FiltersLevel(filtersLevel).FiltersAuditTypes(filtersAuditTypes).FiltersTimes(filtersTimes).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.GetAuditLogsForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogsForGlobal`: OperationResponseGridVOAuditLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.GetAuditLogsForGlobal`: %v\n", resp)
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
	resp, r, err := apiClient.AuditLogAPI.GetAuditLogsForSite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsTime(sortsTime).FiltersResult(filtersResult).FiltersLevel(filtersLevel).FiltersAuditTypes(filtersAuditTypes).FiltersTimes(filtersTimes).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.GetAuditLogsForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogsForSite`: OperationResponseGridVOAuditLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.GetAuditLogsForSite`: %v\n", resp)
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
	resp, r, err := apiClient.AuditLogAPI.ModifyAuditLogSettingGlobal(context.Background(), omadacId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.ModifyAuditLogSettingGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAuditLogSettingGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.ModifyAuditLogSettingGlobal`: %v\n", resp)
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
	resp, r, err := apiClient.AuditLogAPI.ModifyAuditLogSettingSite1(context.Background(), omadacId, siteId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.ModifyAuditLogSettingSite1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAuditLogSettingSite1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.ModifyAuditLogSettingSite1`: %v\n", resp)
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


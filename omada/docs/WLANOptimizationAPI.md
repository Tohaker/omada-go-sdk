# WLANOptimizationAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyPlanningHistory**](WLANOptimizationAPI.md#applyplanninghistory) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory/{id} | apply Planning History
[**BatchDeletePlanningHistory**](WLANOptimizationAPI.md#batchdeleteplanninghistory) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory/batch/delete | Batch Delete PlanningHistory
[**CancelRadioFrequencyPlanning**](WLANOptimizationAPI.md#cancelradiofrequencyplanning) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/rfPlanning/cancel | cancel Radio Frequency Planning
[**DeleteExcludeAps**](WLANOptimizationAPI.md#deleteexcludeaps) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/excludeAps/delete | Delete ExcludeAps
[**DeletePlanningHistory**](WLANOptimizationAPI.md#deleteplanninghistory) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory/{id} | delete Planning History
[**ExportRfPlanningHistory**](WLANOptimizationAPI.md#exportrfplanninghistory) | **Post** /openapi/v1/{omadacId}/files/sites/{siteId}/rfPlanning/history/export | Export PlanningHistory
[**GetExcludeAps**](WLANOptimizationAPI.md#getexcludeaps) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/excludeAps | get ExcludeAps
[**GetExperienceIndex**](WLANOptimizationAPI.md#getexperienceindex) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/experienceIndex | get Experience Index
[**GetGirdPlanningHistorys**](WLANOptimizationAPI.md#getgirdplanninghistorys) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory | get Gird Planning Historys
[**GetLast20PlanningHistorys**](WLANOptimizationAPI.md#getlast20planninghistorys) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/history/last20times | get Last 20 PlanningHistorys
[**GetPlanningHistory**](WLANOptimizationAPI.md#getplanninghistory) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory/{id} | Get Planning History
[**GetPlanningHistorysByTime**](WLANOptimizationAPI.md#getplanninghistorysbytime) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/planningHistoryByTime | Get the planningHistory
[**GetRFPlanningDeployHistory**](WLANOptimizationAPI.md#getrfplanningdeployhistory) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfplanning/history | get RF Planning Deploy History
[**GetRadioFrequencyPlanningConfig**](WLANOptimizationAPI.md#getradiofrequencyplanningconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning | get Radio Frequency Planning Config
[**GetRadioFrequencyPlanningConfigTemplate**](WLANOptimizationAPI.md#getradiofrequencyplanningconfigtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rfPlanning/config | get Radio Frequency Planning Config Template
[**GetRadioFrequencyPlanningResult**](WLANOptimizationAPI.md#getradiofrequencyplanningresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/result | get Radio Frequency Planning Result
[**GetTimeLinePlanningHistorys**](WLANOptimizationAPI.md#gettimelineplanninghistorys) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/history/timeline | Get the planningHistory Time Line
[**ModifyExcludeAps**](WLANOptimizationAPI.md#modifyexcludeaps) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/excludeAps | modify Exclude Aps
[**ModifyRFPlanningDeployConfig**](WLANOptimizationAPI.md#modifyrfplanningdeployconfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/config | modify RF Planning Deploy Config
[**ModifyRFPlanningDeployConfigTemplate**](WLANOptimizationAPI.md#modifyrfplanningdeployconfigtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rfPlanning/config | modify RF Planning Deploy Config Template
[**ModifyRFPlanningScheduleConfig**](WLANOptimizationAPI.md#modifyrfplanningscheduleconfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/schedule/config | modify RF Planning Schedule Config
[**QueryRFPlanningBoard**](WLANOptimizationAPI.md#queryrfplanningboard) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/dashboard | Get RRM AI time line
[**QueryRFPlanningBoardStatus**](WLANOptimizationAPI.md#queryrfplanningboardstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/dashboard/status | Get RRM AI dashboard status
[**QueryRFPlanningDashboardHistory**](WLANOptimizationAPI.md#queryrfplanningdashboardhistory) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/dashboard/history | Get RRM AI dashboard history time line
[**StartOptimization**](WLANOptimizationAPI.md#startoptimization) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/rfPlanning/optimization | start Optimization
[**StartRrmOptimization**](WLANOptimizationAPI.md#startrrmoptimization) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/rfPlanning/rrmOptimization | start Rrm Optimization



## ApplyPlanningHistory

> OperationResponse ApplyPlanningHistory(ctx, omadacId, siteId, id).AppliedConfig(appliedConfig).Execute()

apply Planning History



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
	id := "id_example" // string | planning history ID
	appliedConfig := *openapiclient.NewAppliedConfig() // AppliedConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ApplyPlanningHistory(context.Background(), omadacId, siteId, id).AppliedConfig(appliedConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ApplyPlanningHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyPlanningHistory`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ApplyPlanningHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | planning history ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyPlanningHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **appliedConfig** | [**AppliedConfig**](AppliedConfig.md) |  | 

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


## BatchDeletePlanningHistory

> OperationResponse BatchDeletePlanningHistory(ctx, omadacId, siteId).BatchDeletePlanningHistory(batchDeletePlanningHistory).Execute()

Batch Delete PlanningHistory



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
	batchDeletePlanningHistory := *openapiclient.NewBatchDeletePlanningHistory() // BatchDeletePlanningHistory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.BatchDeletePlanningHistory(context.Background(), omadacId, siteId).BatchDeletePlanningHistory(batchDeletePlanningHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.BatchDeletePlanningHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeletePlanningHistory`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.BatchDeletePlanningHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeletePlanningHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchDeletePlanningHistory** | [**BatchDeletePlanningHistory**](BatchDeletePlanningHistory.md) |  | 

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


## CancelRadioFrequencyPlanning

> OperationResponse CancelRadioFrequencyPlanning(ctx, omadacId, siteId).Execute()

cancel Radio Frequency Planning



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
	resp, r, err := apiClient.WLANOptimizationAPI.CancelRadioFrequencyPlanning(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.CancelRadioFrequencyPlanning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRadioFrequencyPlanning`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.CancelRadioFrequencyPlanning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRadioFrequencyPlanningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExcludeAps

> OperationResponse DeleteExcludeAps(ctx, omadacId, siteId).ExcludeApDeleteOpenApiVO(excludeApDeleteOpenApiVO).Execute()

Delete ExcludeAps



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
	excludeApDeleteOpenApiVO := *openapiclient.NewExcludeApDeleteOpenApiVO() // ExcludeApDeleteOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.DeleteExcludeAps(context.Background(), omadacId, siteId).ExcludeApDeleteOpenApiVO(excludeApDeleteOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.DeleteExcludeAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExcludeAps`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.DeleteExcludeAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExcludeApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludeApDeleteOpenApiVO** | [**ExcludeApDeleteOpenApiVO**](ExcludeApDeleteOpenApiVO.md) |  | 

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


## DeletePlanningHistory

> OperationResponse DeletePlanningHistory(ctx, omadacId, siteId, id).Execute()

delete Planning History



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
	id := "id_example" // string | planning history ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.DeletePlanningHistory(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.DeletePlanningHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePlanningHistory`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.DeletePlanningHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | planning history ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanningHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRfPlanningHistory

> OperationResponse ExportRfPlanningHistory(ctx, omadacId, siteId).ExportRFPlanningHistoryOpenapiVO(exportRFPlanningHistoryOpenapiVO).Execute()

Export PlanningHistory



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
	exportRFPlanningHistoryOpenapiVO := *openapiclient.NewExportRFPlanningHistoryOpenapiVO(int32(123)) // ExportRFPlanningHistoryOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ExportRfPlanningHistory(context.Background(), omadacId, siteId).ExportRFPlanningHistoryOpenapiVO(exportRFPlanningHistoryOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ExportRfPlanningHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportRfPlanningHistory`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ExportRfPlanningHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRfPlanningHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exportRFPlanningHistoryOpenapiVO** | [**ExportRFPlanningHistoryOpenapiVO**](ExportRFPlanningHistoryOpenapiVO.md) |  | 

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


## GetExcludeAps

> OperationResponseListExcludeApVO GetExcludeAps(ctx, omadacId, siteId).CurrentPage(currentPage).CurrentPageSize(currentPageSize).Execute()

get ExcludeAps



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
	currentPage := int32(56) // int32 | Start page number. Start from 1.
	currentPageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.GetExcludeAps(context.Background(), omadacId, siteId).CurrentPage(currentPage).CurrentPageSize(currentPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetExcludeAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExcludeAps`: OperationResponseListExcludeApVO
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetExcludeAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExcludeApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currentPage** | **int32** | Start page number. Start from 1. | 
 **currentPageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseListExcludeApVO**](OperationResponseListExcludeApVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperienceIndex

> ExperienceIndex GetExperienceIndex(ctx, omadacId, siteId).Execute()

get Experience Index



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetExperienceIndex(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetExperienceIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperienceIndex`: ExperienceIndex
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetExperienceIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperienceIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExperienceIndex**](ExperienceIndex.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGirdPlanningHistorys

> GridVORFPlanningHistory GetGirdPlanningHistorys(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

get Gird Planning Historys



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetGirdPlanningHistorys(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetGirdPlanningHistorys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGirdPlanningHistorys`: GridVORFPlanningHistory
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetGirdPlanningHistorys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGirdPlanningHistorysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**GridVORFPlanningHistory**](GridVORFPlanningHistory.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLast20PlanningHistorys

> OperationResponse GetLast20PlanningHistorys(ctx, omadacId, siteId).Execute()

get Last 20 PlanningHistorys



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetLast20PlanningHistorys(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetLast20PlanningHistorys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLast20PlanningHistorys`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetLast20PlanningHistorys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLast20PlanningHistorysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanningHistory

> OperationResponsePlanningHistoryDetail GetPlanningHistory(ctx, omadacId, siteId, id).Execute()

Get Planning History



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
	id := "id_example" // string | planning history ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.GetPlanningHistory(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetPlanningHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanningHistory`: OperationResponsePlanningHistoryDetail
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetPlanningHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | planning history ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanningHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponsePlanningHistoryDetail**](OperationResponsePlanningHistoryDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanningHistorysByTime

> OperationResponseGridVORFPlanningHistory GetPlanningHistorysByTime(ctx, omadacId, siteId).CurrentPage(currentPage).CurrentPageSize(currentPageSize).Start(start).End(end).Execute()

Get the planningHistory



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
	currentPage := int32(56) // int32 | Start page number. Start from 1.
	currentPageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	start := int64(789) // int64 | start time
	end := int64(789) // int64 | end time

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.GetPlanningHistorysByTime(context.Background(), omadacId, siteId).CurrentPage(currentPage).CurrentPageSize(currentPageSize).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetPlanningHistorysByTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanningHistorysByTime`: OperationResponseGridVORFPlanningHistory
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetPlanningHistorysByTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanningHistorysByTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **currentPage** | **int32** | Start page number. Start from 1. | 
 **currentPageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **start** | **int64** | start time | 
 **end** | **int64** | end time | 

### Return type

[**OperationResponseGridVORFPlanningHistory**](OperationResponseGridVORFPlanningHistory.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRFPlanningDeployHistory

> RFPlanningDeployResult GetRFPlanningDeployHistory(ctx, omadacId, siteId).Execute()

get RF Planning Deploy History



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetRFPlanningDeployHistory(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetRFPlanningDeployHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRFPlanningDeployHistory`: RFPlanningDeployResult
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetRFPlanningDeployHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRFPlanningDeployHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RFPlanningDeployResult**](RFPlanningDeployResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadioFrequencyPlanningConfig

> OperationResponsePlanningHistoryDetail GetRadioFrequencyPlanningConfig(ctx, omadacId, siteId).Execute()

get Radio Frequency Planning Config



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetRadioFrequencyPlanningConfig(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetRadioFrequencyPlanningConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadioFrequencyPlanningConfig`: OperationResponsePlanningHistoryDetail
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetRadioFrequencyPlanningConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadioFrequencyPlanningConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponsePlanningHistoryDetail**](OperationResponsePlanningHistoryDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadioFrequencyPlanningConfigTemplate

> OperationResponsePlanningHistoryDetail GetRadioFrequencyPlanningConfigTemplate(ctx, omadacId, siteTemplateId).Execute()

get Radio Frequency Planning Config Template



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetRadioFrequencyPlanningConfigTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetRadioFrequencyPlanningConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadioFrequencyPlanningConfigTemplate`: OperationResponsePlanningHistoryDetail
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetRadioFrequencyPlanningConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadioFrequencyPlanningConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponsePlanningHistoryDetail**](OperationResponsePlanningHistoryDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadioFrequencyPlanningResult

> OperationResponseRFPlanningResult GetRadioFrequencyPlanningResult(ctx, omadacId, siteId).Execute()

get Radio Frequency Planning Result



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetRadioFrequencyPlanningResult(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetRadioFrequencyPlanningResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadioFrequencyPlanningResult`: OperationResponseRFPlanningResult
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetRadioFrequencyPlanningResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadioFrequencyPlanningResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseRFPlanningResult**](OperationResponseRFPlanningResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeLinePlanningHistorys

> OperationResponsePlanningHistoryListVO GetTimeLinePlanningHistorys(ctx, omadacId, siteId).Start(start).End(end).Type_(type_).Execute()

Get the planningHistory Time Line



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
	start := int64(789) // int64 | start time
	end := int64(789) // int64 | end time
	type_ := int32(56) // int32 | type：0-day,1-week

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.GetTimeLinePlanningHistorys(context.Background(), omadacId, siteId).Start(start).End(end).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetTimeLinePlanningHistorys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeLinePlanningHistorys`: OperationResponsePlanningHistoryListVO
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetTimeLinePlanningHistorys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeLinePlanningHistorysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | start time | 
 **end** | **int64** | end time | 
 **type_** | **int32** | type：0-day,1-week | 

### Return type

[**OperationResponsePlanningHistoryListVO**](OperationResponsePlanningHistoryListVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyExcludeAps

> OperationResponse ModifyExcludeAps(ctx, omadacId, siteId).ExcludedAPsConfig(excludedAPsConfig).Execute()

modify Exclude Aps



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
	excludedAPsConfig := *openapiclient.NewExcludedAPsConfig(false) // ExcludedAPsConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ModifyExcludeAps(context.Background(), omadacId, siteId).ExcludedAPsConfig(excludedAPsConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ModifyExcludeAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyExcludeAps`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ModifyExcludeAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyExcludeApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludedAPsConfig** | [**ExcludedAPsConfig**](ExcludedAPsConfig.md) |  | 

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


## ModifyRFPlanningDeployConfig

> OperationResponse ModifyRFPlanningDeployConfig(ctx, omadacId, siteId).RFPlanningDeployConfig(rFPlanningDeployConfig).Execute()

modify RF Planning Deploy Config



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
	rFPlanningDeployConfig := *openapiclient.NewRFPlanningDeployConfig(int32(123)) // RFPlanningDeployConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ModifyRFPlanningDeployConfig(context.Background(), omadacId, siteId).RFPlanningDeployConfig(rFPlanningDeployConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ModifyRFPlanningDeployConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRFPlanningDeployConfig`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ModifyRFPlanningDeployConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRFPlanningDeployConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rFPlanningDeployConfig** | [**RFPlanningDeployConfig**](RFPlanningDeployConfig.md) |  | 

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


## ModifyRFPlanningDeployConfigTemplate

> OperationResponse ModifyRFPlanningDeployConfigTemplate(ctx, omadacId, siteTemplateId).RFPlanningDeployConfig(rFPlanningDeployConfig).Execute()

modify RF Planning Deploy Config Template



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
	rFPlanningDeployConfig := *openapiclient.NewRFPlanningDeployConfig(int32(123)) // RFPlanningDeployConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ModifyRFPlanningDeployConfigTemplate(context.Background(), omadacId, siteTemplateId).RFPlanningDeployConfig(rFPlanningDeployConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ModifyRFPlanningDeployConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRFPlanningDeployConfigTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ModifyRFPlanningDeployConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRFPlanningDeployConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rFPlanningDeployConfig** | [**RFPlanningDeployConfig**](RFPlanningDeployConfig.md) |  | 

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


## ModifyRFPlanningScheduleConfig

> OperationResponse ModifyRFPlanningScheduleConfig(ctx, omadacId, siteId).RFPlanningScheduleConfigOpenApiVO(rFPlanningScheduleConfigOpenApiVO).Execute()

modify RF Planning Schedule Config



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
	rFPlanningScheduleConfigOpenApiVO := *openapiclient.NewRFPlanningScheduleConfigOpenApiVO(false) // RFPlanningScheduleConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ModifyRFPlanningScheduleConfig(context.Background(), omadacId, siteId).RFPlanningScheduleConfigOpenApiVO(rFPlanningScheduleConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ModifyRFPlanningScheduleConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRFPlanningScheduleConfig`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ModifyRFPlanningScheduleConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRFPlanningScheduleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rFPlanningScheduleConfigOpenApiVO** | [**RFPlanningScheduleConfigOpenApiVO**](RFPlanningScheduleConfigOpenApiVO.md) |  | 

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


## QueryRFPlanningBoard

> OperationResponseWlanOptDashBoardOpenApiVO QueryRFPlanningBoard(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get RRM AI time line



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
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.QueryRFPlanningBoard(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.QueryRFPlanningBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryRFPlanningBoard`: OperationResponseWlanOptDashBoardOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.QueryRFPlanningBoard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryRFPlanningBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseWlanOptDashBoardOpenApiVO**](OperationResponseWlanOptDashBoardOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryRFPlanningBoardStatus

> OperationResponseWlanOptDashBoardOpenApiVO QueryRFPlanningBoardStatus(ctx, omadacId, siteId).Execute()

Get RRM AI dashboard status



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
	resp, r, err := apiClient.WLANOptimizationAPI.QueryRFPlanningBoardStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.QueryRFPlanningBoardStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryRFPlanningBoardStatus`: OperationResponseWlanOptDashBoardOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.QueryRFPlanningBoardStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryRFPlanningBoardStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWlanOptDashBoardOpenApiVO**](OperationResponseWlanOptDashBoardOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryRFPlanningDashboardHistory

> OperationResponseWlanOptDashBoardOpenApiVO QueryRFPlanningDashboardHistory(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get RRM AI dashboard history time line



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
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.QueryRFPlanningDashboardHistory(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.QueryRFPlanningDashboardHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryRFPlanningDashboardHistory`: OperationResponseWlanOptDashBoardOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.QueryRFPlanningDashboardHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryRFPlanningDashboardHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseWlanOptDashBoardOpenApiVO**](OperationResponseWlanOptDashBoardOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOptimization

> OperationResponse StartOptimization(ctx, omadacId, siteId).OptimizationStrategy(optimizationStrategy).Execute()

start Optimization



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
	optimizationStrategy := *openapiclient.NewOptimizationStrategy(int32(123)) // OptimizationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.StartOptimization(context.Background(), omadacId, siteId).OptimizationStrategy(optimizationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.StartOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOptimization`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.StartOptimization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **optimizationStrategy** | [**OptimizationStrategy**](OptimizationStrategy.md) |  | 

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


## StartRrmOptimization

> OperationResponse StartRrmOptimization(ctx, omadacId, siteId).OptimizationStrategy(optimizationStrategy).Execute()

start Rrm Optimization



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
	optimizationStrategy := *openapiclient.NewOptimizationStrategy(int32(123)) // OptimizationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.StartRrmOptimization(context.Background(), omadacId, siteId).OptimizationStrategy(optimizationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.StartRrmOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartRrmOptimization`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.StartRrmOptimization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartRrmOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **optimizationStrategy** | [**OptimizationStrategy**](OptimizationStrategy.md) |  | 

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


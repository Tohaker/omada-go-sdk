# \ReportV2API

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTabNetworkReport**](ReportV2API.md#AddTabNetworkReport) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/report/tab | add new tab 
[**DeleteTabNetworkReport**](ReportV2API.md#DeleteTabNetworkReport) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/report/tab/{tabIds} | batch delete tabs
[**EmailReportScheduleV2**](ReportV2API.md#EmailReportScheduleV2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/report/export-schedule-email | Export report for scheduler
[**ExportLaterEmailV2**](ReportV2API.md#ExportLaterEmailV2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/report/export-later-email | Export report once later
[**ExportNowEmailV2**](ReportV2API.md#ExportNowEmailV2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/report/export-now-email | Export report now for email
[**ExportNowLocal**](ReportV2API.md#ExportNowLocal) | **Post** /openapi/v2/{omadacId}/files/sites/{siteId}/report/export-now-local | Export report now
[**GetAllTabs**](ReportV2API.md#GetAllTabs) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/report/allTabs | Get report all tabs
[**GetCardsInfo**](ReportV2API.md#GetCardsInfo) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/report/cards | Get cards info 
[**GetScheduleTimeInfoV2**](ReportV2API.md#GetScheduleTimeInfoV2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/report/export-schedule-email | Get schedule time info
[**GetTabById**](ReportV2API.md#GetTabById) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/report/{tabId} | get all cards in a tab
[**RecordTabNetworkReport**](ReportV2API.md#RecordTabNetworkReport) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/report/reorder | reorder the tab 
[**UpdateTabNetworkReport**](ReportV2API.md#UpdateTabNetworkReport) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/report/tab | update  tab 



## AddTabNetworkReport

> OperationResponse AddTabNetworkReport(ctx, omadacId, siteId).ReportTab(reportTab).Execute()

add new tab 



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
	reportTab := *openapiclient.NewReportTab() // ReportTab | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.AddTabNetworkReport(context.Background(), omadacId, siteId).ReportTab(reportTab).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.AddTabNetworkReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTabNetworkReport`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.AddTabNetworkReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTabNetworkReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reportTab** | [**ReportTab**](ReportTab.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTabNetworkReport

> OperationResponse DeleteTabNetworkReport(ctx, omadacId, siteId, tabIds).Execute()

batch delete tabs



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
	tabIds := "tabIds_example" // string | tab ID list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.DeleteTabNetworkReport(context.Background(), omadacId, siteId, tabIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.DeleteTabNetworkReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTabNetworkReport`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.DeleteTabNetworkReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tabIds** | **string** | tab ID list | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTabNetworkReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailReportScheduleV2

> OperationResponse EmailReportScheduleV2(ctx, omadacId, siteId).NetworkReportScheduleVO(networkReportScheduleVO).Execute()

Export report for scheduler



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
	networkReportScheduleVO := *openapiclient.NewNetworkReportScheduleVO(false) // NetworkReportScheduleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.EmailReportScheduleV2(context.Background(), omadacId, siteId).NetworkReportScheduleVO(networkReportScheduleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.EmailReportScheduleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailReportScheduleV2`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.EmailReportScheduleV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailReportScheduleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkReportScheduleVO** | [**NetworkReportScheduleVO**](NetworkReportScheduleVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLaterEmailV2

> OperationResponse ExportLaterEmailV2(ctx, omadacId, siteId).NetworkReportScheduleLaterVO(networkReportScheduleLaterVO).Execute()

Export report once later



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
	networkReportScheduleLaterVO := *openapiclient.NewNetworkReportScheduleLaterVO(false) // NetworkReportScheduleLaterVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.ExportLaterEmailV2(context.Background(), omadacId, siteId).NetworkReportScheduleLaterVO(networkReportScheduleLaterVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.ExportLaterEmailV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLaterEmailV2`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.ExportLaterEmailV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLaterEmailV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkReportScheduleLaterVO** | [**NetworkReportScheduleLaterVO**](NetworkReportScheduleLaterVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportNowEmailV2

> OperationResponse ExportNowEmailV2(ctx, omadacId, siteId).ReportExportV2(reportExportV2).Execute()

Export report now for email



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
	reportExportV2 := *openapiclient.NewReportExportV2(int64(123), "ReportName_example", int32(123), int64(123), []string{"TabIdList_example"}) // ReportExportV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.ExportNowEmailV2(context.Background(), omadacId, siteId).ReportExportV2(reportExportV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.ExportNowEmailV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportNowEmailV2`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.ExportNowEmailV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportNowEmailV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reportExportV2** | [**ReportExportV2**](ReportExportV2.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportNowLocal

> OperationResponse ExportNowLocal(ctx, omadacId, siteId).ReportExportV2(reportExportV2).Execute()

Export report now



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
	reportExportV2 := *openapiclient.NewReportExportV2(int64(123), "ReportName_example", int32(123), int64(123), []string{"TabIdList_example"}) // ReportExportV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.ExportNowLocal(context.Background(), omadacId, siteId).ReportExportV2(reportExportV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.ExportNowLocal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportNowLocal`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.ExportNowLocal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportNowLocalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reportExportV2** | [**ReportExportV2**](ReportExportV2.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTabs

> OperationResponseListReportTab GetAllTabs(ctx, omadacId, siteId).Execute()

Get report all tabs



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
	resp, r, err := apiClient.ReportV2API.GetAllTabs(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.GetAllTabs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTabs`: OperationResponseListReportTab
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.GetAllTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListReportTab**](OperationResponseListReportTab.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardsInfo

> OperationResponseCardInfoVO GetCardsInfo(ctx, omadacId, siteId).ReportCardQueryVO(reportCardQueryVO).Execute()

Get cards info 



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
	reportCardQueryVO := *openapiclient.NewReportCardQueryVO() // ReportCardQueryVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.GetCardsInfo(context.Background(), omadacId, siteId).ReportCardQueryVO(reportCardQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.GetCardsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardsInfo`: OperationResponseCardInfoVO
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.GetCardsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reportCardQueryVO** | [**ReportCardQueryVO**](ReportCardQueryVO.md) |  | 

### Return type

[**OperationResponseCardInfoVO**](OperationResponseCardInfoVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduleTimeInfoV2

> OperationResponse GetScheduleTimeInfoV2(ctx, omadacId, siteId).Execute()

Get schedule time info



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
	resp, r, err := apiClient.ReportV2API.GetScheduleTimeInfoV2(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.GetScheduleTimeInfoV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduleTimeInfoV2`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.GetScheduleTimeInfoV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleTimeInfoV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTabById

> OperationResponse GetTabById(ctx, omadacId, siteId, tabId).Execute()

get all cards in a tab



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
	tabId := "tabId_example" // string | tab ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.GetTabById(context.Background(), omadacId, siteId, tabId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.GetTabById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTabById`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.GetTabById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tabId** | **string** | tab ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTabByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordTabNetworkReport

> OperationResponse RecordTabNetworkReport(ctx, omadacId, siteId).ReportTab(reportTab).Execute()

reorder the tab 



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
	reportTab := []openapiclient.ReportTab{*openapiclient.NewReportTab()} // []ReportTab | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.RecordTabNetworkReport(context.Background(), omadacId, siteId).ReportTab(reportTab).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.RecordTabNetworkReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordTabNetworkReport`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.RecordTabNetworkReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordTabNetworkReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reportTab** | [**[]ReportTab**](ReportTab.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTabNetworkReport

> OperationResponse UpdateTabNetworkReport(ctx, omadacId, siteId).ReportTab(reportTab).Execute()

update  tab 



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
	reportTab := *openapiclient.NewReportTab() // ReportTab | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportV2API.UpdateTabNetworkReport(context.Background(), omadacId, siteId).ReportTab(reportTab).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportV2API.UpdateTabNetworkReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTabNetworkReport`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportV2API.UpdateTabNetworkReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTabNetworkReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reportTab** | [**ReportTab**](ReportTab.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


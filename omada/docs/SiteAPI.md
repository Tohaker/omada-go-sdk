# \SiteAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScenario**](SiteAPI.md#AddScenario) | **Post** /openapi/v1/{omadacId}/scenarios | Create new scenario
[**AddTag1**](SiteAPI.md#AddTag1) | **Post** /openapi/v1/{omadacId}/sites/tags | Create new site tag
[**BatchSiteCopy**](SiteAPI.md#BatchSiteCopy) | **Post** /openapi/v1/{omadacId}/sites/copy | Batch create sites by copying from existing site
[**BatchSiteImport**](SiteAPI.md#BatchSiteImport) | **Post** /openapi/v1/{omadacId}/sites/multi-import | Batch create sites by importing site backup files from file server
[**CreateNewSite**](SiteAPI.md#CreateNewSite) | **Post** /openapi/v1/{omadacId}/sites | Create new site
[**CreateNewSiteByTemplate**](SiteAPI.md#CreateNewSiteByTemplate) | **Post** /openapi/v1/{omadacId}/sites/template | Create new site from site template
[**DeleteScenario**](SiteAPI.md#DeleteScenario) | **Delete** /openapi/v1/{omadacId}/scenarios | Delete site scenario
[**DeleteSite**](SiteAPI.md#DeleteSite) | **Delete** /openapi/v1/{omadacId}/sites/{siteId} | Delete an existing site
[**DeleteTag1**](SiteAPI.md#DeleteTag1) | **Delete** /openapi/v1/{omadacId}/sites/tags | Delete an existing site tag
[**GetAvailableSiteToBind**](SiteAPI.md#GetAvailableSiteToBind) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/available-bind | Get available site to bind template
[**GetNtpServerStatus**](SiteAPI.md#GetNtpServerStatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/ntp | Get ntp server status
[**GetNtpServerStatusTemplate**](SiteAPI.md#GetNtpServerStatusTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/ntp | Get template ntp server status
[**GetPlacedSitePositions**](SiteAPI.md#GetPlacedSitePositions) | **Get** /openapi/v1/{omadacId}/placed-site-position | Obtain the geographic location information of Placed Sites
[**GetScenarioList**](SiteAPI.md#GetScenarioList) | **Get** /openapi/v1/{omadacId}/scenarios | Get scenario list
[**GetScenarioListDifference**](SiteAPI.md#GetScenarioListDifference) | **Get** /openapi/v1/{omadacId}/scenarios/difference | Get scenario list difference
[**GetSiteDeviceAccountSetting**](SiteAPI.md#GetSiteDeviceAccountSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/device-account | Get site device account setting
[**GetSiteEntity**](SiteAPI.md#GetSiteEntity) | **Get** /openapi/v1/{omadacId}/sites/{siteId} | Get site info
[**GetSiteInfoForAbnormal**](SiteAPI.md#GetSiteInfoForAbnormal) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/global-dashboard/site-map/abnormal | Get the details of a site
[**GetSiteList**](SiteAPI.md#GetSiteList) | **Get** /openapi/v1/{omadacId}/sites | Get site list
[**GetSiteSummaryStatisticByOpenApi**](SiteAPI.md#GetSiteSummaryStatisticByOpenApi) | **Post** /openapi/v1/{omadacId}/sites/statistic | Get sites statistic
[**GetSiteUrlByOpenApi**](SiteAPI.md#GetSiteUrlByOpenApi) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/url | Get site url
[**GetTags1**](SiteAPI.md#GetTags1) | **Get** /openapi/v1/{omadacId}/sites/tags | Get site tag list
[**GetUnplacedSitePositions**](SiteAPI.md#GetUnplacedSitePositions) | **Get** /openapi/v1/{omadacId}/unplaced-site-position | Obtain the geographic location information of unplaced Sites
[**ModifyTag1**](SiteAPI.md#ModifyTag1) | **Patch** /openapi/v1/{omadacId}/sites/tags | Modify an existing site tag
[**UpdateSiteDeviceAccountSetting**](SiteAPI.md#UpdateSiteDeviceAccountSetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/device-account | Update site device account setting
[**UpdateSiteEntity**](SiteAPI.md#UpdateSiteEntity) | **Put** /openapi/v1/{omadacId}/sites/{siteId} | Modify an existing site



## AddScenario

> OperationResponseListString AddScenario(ctx, omadacId).Scenario(scenario).Execute()

Create new scenario



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
	scenario := *openapiclient.NewScenario("Name_example") // Scenario | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.AddScenario(context.Background(), omadacId).Scenario(scenario).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.AddScenario``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddScenario`: OperationResponseListString
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.AddScenario`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scenario** | [**Scenario**](Scenario.md) |  | 

### Return type

[**OperationResponseListString**](OperationResponseListString.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTag1

> SiteTag AddTag1(ctx, omadacId).CreateSiteTagOpenApiVO(createSiteTagOpenApiVO).Execute()

Create new site tag



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
	createSiteTagOpenApiVO := *openapiclient.NewCreateSiteTagOpenApiVO() // CreateSiteTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.AddTag1(context.Background(), omadacId).CreateSiteTagOpenApiVO(createSiteTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.AddTag1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTag1`: SiteTag
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.AddTag1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTag1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSiteTagOpenApiVO** | [**CreateSiteTagOpenApiVO**](CreateSiteTagOpenApiVO.md) |  | 

### Return type

[**SiteTag**](SiteTag.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchSiteCopy

> OperationResponseSiteResultVO BatchSiteCopy(ctx, omadacId).BatchSiteCopyVO(batchSiteCopyVO).Execute()

Batch create sites by copying from existing site



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
	batchSiteCopyVO := *openapiclient.NewBatchSiteCopyVO("SiteNamePrefix_example", "SourceSiteId_example", int32(123)) // BatchSiteCopyVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.BatchSiteCopy(context.Background(), omadacId).BatchSiteCopyVO(batchSiteCopyVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.BatchSiteCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSiteCopy`: OperationResponseSiteResultVO
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.BatchSiteCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchSiteCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchSiteCopyVO** | [**BatchSiteCopyVO**](BatchSiteCopyVO.md) |  | 

### Return type

[**OperationResponseSiteResultVO**](OperationResponseSiteResultVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchSiteImport

> OperationResponseSiteResultVO BatchSiteImport(ctx, omadacId).BatchSiteImportVO(batchSiteImportVO).Execute()

Batch create sites by importing site backup files from file server



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
	batchSiteImportVO := *openapiclient.NewBatchSiteImportVO(*openapiclient.NewFileServerOpenApiVO("Hostname_example", int32(123), "Protocol_example"), []openapiclient.SiteImportOpenApiVO{*openapiclient.NewSiteImportOpenApiVO("FilePath_example", "SiteName_example")}) // BatchSiteImportVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.BatchSiteImport(context.Background(), omadacId).BatchSiteImportVO(batchSiteImportVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.BatchSiteImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSiteImport`: OperationResponseSiteResultVO
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.BatchSiteImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchSiteImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchSiteImportVO** | [**BatchSiteImportVO**](BatchSiteImportVO.md) |  | 

### Return type

[**OperationResponseSiteResultVO**](OperationResponseSiteResultVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewSite

> OperationResponse CreateNewSite(ctx, omadacId).CreateSiteEntity(createSiteEntity).Execute()

Create new site



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
	createSiteEntity := *openapiclient.NewCreateSiteEntity(*openapiclient.NewDeviceAccountSettingOpenApiVO("Password_example", "Username_example"), "Name_example", "Region_example", "Scenario_example", "TimeZone_example") // CreateSiteEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.CreateNewSite(context.Background(), omadacId).CreateSiteEntity(createSiteEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.CreateNewSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewSite`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.CreateNewSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSiteEntity** | [**CreateSiteEntity**](CreateSiteEntity.md) |  | 

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


## CreateNewSiteByTemplate

> OperationResponse CreateNewSiteByTemplate(ctx, omadacId).CreateSiteByTemplate(createSiteByTemplate).Execute()

Create new site from site template



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
	createSiteByTemplate := *openapiclient.NewCreateSiteByTemplate(*openapiclient.NewDeviceAccountSettingOpenApiVO("Password_example", "Username_example"), "Name_example", "Region_example", "Scenario_example", "SiteTemplateId_example") // CreateSiteByTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.CreateNewSiteByTemplate(context.Background(), omadacId).CreateSiteByTemplate(createSiteByTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.CreateNewSiteByTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewSiteByTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.CreateNewSiteByTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSiteByTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSiteByTemplate** | [**CreateSiteByTemplate**](CreateSiteByTemplate.md) |  | 

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


## DeleteScenario

> OperationResponseWithoutResult DeleteScenario(ctx, omadacId).Scenario(scenario).Execute()

Delete site scenario



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
	scenario := *openapiclient.NewScenario("Name_example") // Scenario | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.DeleteScenario(context.Background(), omadacId).Scenario(scenario).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.DeleteScenario``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteScenario`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.DeleteScenario`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scenario** | [**Scenario**](Scenario.md) |  | 

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


## DeleteSite

> OperationResponseWithoutResult DeleteSite(ctx, omadacId, siteId).Execute()

Delete an existing site



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
	resp, r, err := apiClient.SiteAPI.DeleteSite(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.DeleteSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.DeleteSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteRequest struct via the builder pattern


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


## DeleteTag1

> OperationResponseWithoutResult DeleteTag1(ctx, omadacId).DeleteSiteTagOpenApiVO(deleteSiteTagOpenApiVO).Execute()

Delete an existing site tag



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
	deleteSiteTagOpenApiVO := *openapiclient.NewDeleteSiteTagOpenApiVO("TagId_example") // DeleteSiteTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.DeleteTag1(context.Background(), omadacId).DeleteSiteTagOpenApiVO(deleteSiteTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.DeleteTag1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTag1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.DeleteTag1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTag1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteSiteTagOpenApiVO** | [**DeleteSiteTagOpenApiVO**](DeleteSiteTagOpenApiVO.md) |  | 

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


## GetAvailableSiteToBind

> OperationResponseSitesSite GetAvailableSiteToBind(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get available site to bind template



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetAvailableSiteToBind(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetAvailableSiteToBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableSiteToBind`: OperationResponseSitesSite
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetAvailableSiteToBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableSiteToBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseSitesSite**](OperationResponseSitesSite.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNtpServerStatus

> OperationResponse GetNtpServerStatus(ctx, omadacId, siteId).Execute()

Get ntp server status



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
	resp, r, err := apiClient.SiteAPI.GetNtpServerStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetNtpServerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNtpServerStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetNtpServerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNtpServerStatusRequest struct via the builder pattern


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


## GetNtpServerStatusTemplate

> OperationResponse GetNtpServerStatusTemplate(ctx, omadacId, siteTemplateId).Execute()

Get template ntp server status



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
	resp, r, err := apiClient.SiteAPI.GetNtpServerStatusTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetNtpServerStatusTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNtpServerStatusTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetNtpServerStatusTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNtpServerStatusTemplateRequest struct via the builder pattern


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


## GetPlacedSitePositions

> OperationResponseListPlacedSite GetPlacedSitePositions(ctx, omadacId).Execute()

Obtain the geographic location information of Placed Sites



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
	resp, r, err := apiClient.SiteAPI.GetPlacedSitePositions(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetPlacedSitePositions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlacedSitePositions`: OperationResponseListPlacedSite
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetPlacedSitePositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlacedSitePositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseListPlacedSite**](OperationResponseListPlacedSite.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScenarioList

> OperationResponseListString GetScenarioList(ctx, omadacId).Execute()

Get scenario list



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
	resp, r, err := apiClient.SiteAPI.GetScenarioList(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetScenarioList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScenarioList`: OperationResponseListString
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetScenarioList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScenarioListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseListString**](OperationResponseListString.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScenarioListDifference

> OperationResponseSetScenarioDifferenceVO GetScenarioListDifference(ctx, omadacId).Execute()

Get scenario list difference



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
	resp, r, err := apiClient.SiteAPI.GetScenarioListDifference(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetScenarioListDifference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScenarioListDifference`: OperationResponseSetScenarioDifferenceVO
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetScenarioListDifference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScenarioListDifferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseSetScenarioDifferenceVO**](OperationResponseSetScenarioDifferenceVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceAccountSetting

> OperationResponseDeviceAccountSettingOpenApiVO GetSiteDeviceAccountSetting(ctx, omadacId, siteId).Execute()

Get site device account setting



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
	resp, r, err := apiClient.SiteAPI.GetSiteDeviceAccountSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSiteDeviceAccountSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceAccountSetting`: OperationResponseDeviceAccountSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSiteDeviceAccountSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceAccountSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDeviceAccountSettingOpenApiVO**](OperationResponseDeviceAccountSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteEntity

> OperationResponseSiteEntity GetSiteEntity(ctx, omadacId, siteId).Execute()

Get site info



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
	resp, r, err := apiClient.SiteAPI.GetSiteEntity(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSiteEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteEntity`: OperationResponseSiteEntity
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSiteEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteEntity**](OperationResponseSiteEntity.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteInfoForAbnormal

> OperationResponse GetSiteInfoForAbnormal(ctx, omadacId, siteId).Execute()

Get the details of a site



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
	resp, r, err := apiClient.SiteAPI.GetSiteInfoForAbnormal(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSiteInfoForAbnormal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteInfoForAbnormal`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSiteInfoForAbnormal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteInfoForAbnormalRequest struct via the builder pattern


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


## GetSiteList

> OperationResponseGridVOSiteSummaryInfo GetSiteList(ctx, omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).FiltersTag(filtersTag).FiltersType(filtersType).Execute()

Get site list



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
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag ID (optional)
	filtersType := "filtersType_example" // string | Filter query parameters, support field site type. 0: basic site; 1: pro site. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetSiteList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).FiltersTag(filtersTag).FiltersType(filtersType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSiteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteList`: OperationResponseGridVOSiteSummaryInfo
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSiteList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 
 **filtersTag** | **string** | Filter query parameters, support field tag ID | 
 **filtersType** | **string** | Filter query parameters, support field site type. 0: basic site; 1: pro site. | 

### Return type

[**OperationResponseGridVOSiteSummaryInfo**](OperationResponseGridVOSiteSummaryInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSummaryStatisticByOpenApi

> OperationResponseSiteSummaryStatistic GetSiteSummaryStatisticByOpenApi(ctx, omadacId).SiteStatisticList(siteStatisticList).Execute()

Get sites statistic



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
	siteStatisticList := *openapiclient.NewSiteStatisticList() // SiteStatisticList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetSiteSummaryStatisticByOpenApi(context.Background(), omadacId).SiteStatisticList(siteStatisticList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSiteSummaryStatisticByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSummaryStatisticByOpenApi`: OperationResponseSiteSummaryStatistic
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSiteSummaryStatisticByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSummaryStatisticByOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteStatisticList** | [**SiteStatisticList**](SiteStatisticList.md) |  | 

### Return type

[**OperationResponseSiteSummaryStatistic**](OperationResponseSiteSummaryStatistic.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteUrlByOpenApi

> OperationResponseSiteUrlOpenApiVO GetSiteUrlByOpenApi(ctx, omadacId, siteId).Execute()

Get site url



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
	resp, r, err := apiClient.SiteAPI.GetSiteUrlByOpenApi(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetSiteUrlByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteUrlByOpenApi`: OperationResponseSiteUrlOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetSiteUrlByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteUrlByOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteUrlOpenApiVO**](OperationResponseSiteUrlOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags1

> []SiteTag GetTags1(ctx, omadacId).Execute()

Get site tag list



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
	resp, r, err := apiClient.SiteAPI.GetTags1(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetTags1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags1`: []SiteTag
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetTags1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTags1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SiteTag**](SiteTag.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnplacedSitePositions

> OperationResponseUnplacedSitesUnplacedSite GetUnplacedSitePositions(ctx, omadacId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Obtain the geographic location information of unplaced Sites



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
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.GetUnplacedSitePositions(context.Background(), omadacId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.GetUnplacedSitePositions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnplacedSitePositions`: OperationResponseUnplacedSitesUnplacedSite
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.GetUnplacedSitePositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnplacedSitePositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseUnplacedSitesUnplacedSite**](OperationResponseUnplacedSitesUnplacedSite.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTag1

> OperationResponseWithoutResult ModifyTag1(ctx, omadacId).SiteTag(siteTag).Execute()

Modify an existing site tag



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
	siteTag := *openapiclient.NewSiteTag() // SiteTag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.ModifyTag1(context.Background(), omadacId).SiteTag(siteTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.ModifyTag1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTag1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.ModifyTag1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTag1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteTag** | [**SiteTag**](SiteTag.md) |  | 

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


## UpdateSiteDeviceAccountSetting

> OperationResponseWithoutResult UpdateSiteDeviceAccountSetting(ctx, omadacId, siteId).DeviceAccountSettingOpenApiVO(deviceAccountSettingOpenApiVO).Execute()

Update site device account setting



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
	deviceAccountSettingOpenApiVO := *openapiclient.NewDeviceAccountSettingOpenApiVO("Password_example", "Username_example") // DeviceAccountSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.UpdateSiteDeviceAccountSetting(context.Background(), omadacId, siteId).DeviceAccountSettingOpenApiVO(deviceAccountSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.UpdateSiteDeviceAccountSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteDeviceAccountSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.UpdateSiteDeviceAccountSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteDeviceAccountSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceAccountSettingOpenApiVO** | [**DeviceAccountSettingOpenApiVO**](DeviceAccountSettingOpenApiVO.md) |  | 

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


## UpdateSiteEntity

> OperationResponseWithoutResult UpdateSiteEntity(ctx, omadacId, siteId).UpdateSiteEntity(updateSiteEntity).Execute()

Modify an existing site



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
	updateSiteEntity := *openapiclient.NewUpdateSiteEntity("Region_example", "Scenario_example", "TimeZone_example") // UpdateSiteEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteAPI.UpdateSiteEntity(context.Background(), omadacId, siteId).UpdateSiteEntity(updateSiteEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.UpdateSiteEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteEntity`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteAPI.UpdateSiteEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSiteEntity** | [**UpdateSiteEntity**](UpdateSiteEntity.md) |  | 

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


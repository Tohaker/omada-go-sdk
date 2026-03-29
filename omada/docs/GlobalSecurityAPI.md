# \GlobalSecurityAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBlockedCountry**](GlobalSecurityAPI.md#AddBlockedCountry) | **Post** /openapi/v1/{omadacId}/security/blocked-country/add | Add Blocked Country
[**DelBlockedCountry**](GlobalSecurityAPI.md#DelBlockedCountry) | **Delete** /openapi/v1/{omadacId}/security/blocked-country | Delete blocked countries.
[**DeleteGlobalThreatList**](GlobalSecurityAPI.md#DeleteGlobalThreatList) | **Post** /openapi/v1/{omadacId}/ips/threat | Delete Global Threat List
[**GetBlockedCountries**](GlobalSecurityAPI.md#GetBlockedCountries) | **Post** /openapi/v1/{omadacId}/security/blocked-country | Get Blocked Countries
[**GetCountryThreats**](GlobalSecurityAPI.md#GetCountryThreats) | **Post** /openapi/v1/{omadacId}/security/threat-map/threats | Get Country Threats
[**GetGlobalCategory**](GlobalSecurityAPI.md#GetGlobalCategory) | **Post** /openapi/v1/{omadacId}/security/threat-map/threat-count | Get Global Category
[**GetGlobalThreatList**](GlobalSecurityAPI.md#GetGlobalThreatList) | **Get** /openapi/v1/{omadacId}/security/threat-management | Get Global Threat List
[**GetGlobalThreatMap**](GlobalSecurityAPI.md#GetGlobalThreatMap) | **Post** /openapi/v1/{omadacId}/security/threat-map | Get Global Threat Map
[**GetGlobalTopThreatList**](GlobalSecurityAPI.md#GetGlobalTopThreatList) | **Get** /openapi/v1/{omadacId}/security/threat-management/top | Get Global Top Threat List
[**GetThreatCount**](GlobalSecurityAPI.md#GetThreatCount) | **Get** /openapi/v1/{omadacId}/security/threat-management/severity | Get Threat Count
[**OperateGlobalThreats**](GlobalSecurityAPI.md#OperateGlobalThreats) | **Post** /openapi/v1/{omadacId}/ips/threat/ops | Operate Global Threats



## AddBlockedCountry

> OperationResponse AddBlockedCountry(ctx, omadacId).OpsBlockedCountryOpenApiVO(opsBlockedCountryOpenApiVO).Execute()

Add Blocked Country



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
	opsBlockedCountryOpenApiVO := *openapiclient.NewOpsBlockedCountryOpenApiVO() // OpsBlockedCountryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.AddBlockedCountry(context.Background(), omadacId).OpsBlockedCountryOpenApiVO(opsBlockedCountryOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.AddBlockedCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBlockedCountry`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.AddBlockedCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBlockedCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsBlockedCountryOpenApiVO** | [**OpsBlockedCountryOpenApiVO**](OpsBlockedCountryOpenApiVO.md) |  | 

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


## DelBlockedCountry

> OperationResponse DelBlockedCountry(ctx, omadacId).OpsBlockedCountryOpenApiVO(opsBlockedCountryOpenApiVO).Execute()

Delete blocked countries.



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
	opsBlockedCountryOpenApiVO := *openapiclient.NewOpsBlockedCountryOpenApiVO() // OpsBlockedCountryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.DelBlockedCountry(context.Background(), omadacId).OpsBlockedCountryOpenApiVO(opsBlockedCountryOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.DelBlockedCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelBlockedCountry`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.DelBlockedCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelBlockedCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opsBlockedCountryOpenApiVO** | [**OpsBlockedCountryOpenApiVO**](OpsBlockedCountryOpenApiVO.md) |  | 

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


## DeleteGlobalThreatList

> OperationResponse DeleteGlobalThreatList(ctx, omadacId).DeleteGlobalThreatOpenApiVO(deleteGlobalThreatOpenApiVO).Execute()

Delete Global Threat List



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
	deleteGlobalThreatOpenApiVO := *openapiclient.NewDeleteGlobalThreatOpenApiVO() // DeleteGlobalThreatOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.DeleteGlobalThreatList(context.Background(), omadacId).DeleteGlobalThreatOpenApiVO(deleteGlobalThreatOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.DeleteGlobalThreatList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGlobalThreatList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.DeleteGlobalThreatList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalThreatListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteGlobalThreatOpenApiVO** | [**DeleteGlobalThreatOpenApiVO**](DeleteGlobalThreatOpenApiVO.md) |  | 

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


## GetBlockedCountries

> OperationResponseBlockedCountryOpenApiVO GetBlockedCountries(ctx, omadacId).SiteListOpenApiVO(siteListOpenApiVO).Execute()

Get Blocked Countries



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
	siteListOpenApiVO := *openapiclient.NewSiteListOpenApiVO() // SiteListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.GetBlockedCountries(context.Background(), omadacId).SiteListOpenApiVO(siteListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.GetBlockedCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockedCountries`: OperationResponseBlockedCountryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.GetBlockedCountries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockedCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteListOpenApiVO** | [**SiteListOpenApiVO**](SiteListOpenApiVO.md) |  | 

### Return type

[**OperationResponseBlockedCountryOpenApiVO**](OperationResponseBlockedCountryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryThreats

> OperationResponseGridVOIpsThreatOpenApiVO GetCountryThreats(ctx, omadacId).QueryCountryThreatListOpenApiVO(queryCountryThreatListOpenApiVO).Execute()

Get Country Threats



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
	queryCountryThreatListOpenApiVO := *openapiclient.NewQueryCountryThreatListOpenApiVO() // QueryCountryThreatListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.GetCountryThreats(context.Background(), omadacId).QueryCountryThreatListOpenApiVO(queryCountryThreatListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.GetCountryThreats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryThreats`: OperationResponseGridVOIpsThreatOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.GetCountryThreats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryThreatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryCountryThreatListOpenApiVO** | [**QueryCountryThreatListOpenApiVO**](QueryCountryThreatListOpenApiVO.md) |  | 

### Return type

[**OperationResponseGridVOIpsThreatOpenApiVO**](OperationResponseGridVOIpsThreatOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalCategory

> OperationResponseThreatMapCategoryOpenApiVO GetGlobalCategory(ctx, omadacId).QueryThreatMapOpenApiVO(queryThreatMapOpenApiVO).Execute()

Get Global Category



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
	queryThreatMapOpenApiVO := *openapiclient.NewQueryThreatMapOpenApiVO() // QueryThreatMapOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.GetGlobalCategory(context.Background(), omadacId).QueryThreatMapOpenApiVO(queryThreatMapOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.GetGlobalCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalCategory`: OperationResponseThreatMapCategoryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.GetGlobalCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryThreatMapOpenApiVO** | [**QueryThreatMapOpenApiVO**](QueryThreatMapOpenApiVO.md) |  | 

### Return type

[**OperationResponseThreatMapCategoryOpenApiVO**](OperationResponseThreatMapCategoryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalThreatList

> OperationResponseGridVOIpsThreatOpenApiVO GetGlobalThreatList(ctx, omadacId).Archived(archived).Page(page).PageSize(pageSize).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).SiteList(siteList).FiltersSeverity(filtersSeverity).SortsTime(sortsTime).SearchKey(searchKey).Execute()

Get Global Threat List



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
	archived := true // bool | archived
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersStartTime := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	filtersEndTime := int64(789) // int64 | End timestamp, in seconds, such as 1682000000
	siteList := "siteList_example" // string | sites are separated by ','. If no value is passed, all sites are selected by default. (optional)
	filtersSeverity := int32(56) // int32 | Threat Severity, such as 0:Critical, 1: Major, 2:Moderate, 3:Minor, 4:Low (optional)
	sortsTime := "sortsTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field Threat Description/Classification/Classification Description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.GetGlobalThreatList(context.Background(), omadacId).Archived(archived).Page(page).PageSize(pageSize).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).SiteList(siteList).FiltersSeverity(filtersSeverity).SortsTime(sortsTime).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.GetGlobalThreatList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalThreatList`: OperationResponseGridVOIpsThreatOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.GetGlobalThreatList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalThreatListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | archived | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersStartTime** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **filtersEndTime** | **int64** | End timestamp, in seconds, such as 1682000000 | 
 **siteList** | **string** | sites are separated by &#39;,&#39;. If no value is passed, all sites are selected by default. | 
 **filtersSeverity** | **int32** | Threat Severity, such as 0:Critical, 1: Major, 2:Moderate, 3:Minor, 4:Low | 
 **sortsTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field Threat Description/Classification/Classification Description | 

### Return type

[**OperationResponseGridVOIpsThreatOpenApiVO**](OperationResponseGridVOIpsThreatOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalThreatMap

> OperationResponseGetThreatMapOpenApiVO GetGlobalThreatMap(ctx, omadacId).QueryThreatMapOpenApiVO(queryThreatMapOpenApiVO).Execute()

Get Global Threat Map



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
	queryThreatMapOpenApiVO := *openapiclient.NewQueryThreatMapOpenApiVO() // QueryThreatMapOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.GetGlobalThreatMap(context.Background(), omadacId).QueryThreatMapOpenApiVO(queryThreatMapOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.GetGlobalThreatMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalThreatMap`: OperationResponseGetThreatMapOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.GetGlobalThreatMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalThreatMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryThreatMapOpenApiVO** | [**QueryThreatMapOpenApiVO**](QueryThreatMapOpenApiVO.md) |  | 

### Return type

[**OperationResponseGetThreatMapOpenApiVO**](OperationResponseGetThreatMapOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalTopThreatList

> OperationResponseTop5ThreatNumOpenApiVO GetGlobalTopThreatList(ctx, omadacId).StartTime(startTime).EndTime(endTime).Execute()

Get Global Top Threat List



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
	startTime := int64(789) // int64 | Start Time
	endTime := int64(789) // int64 | End Time

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.GetGlobalTopThreatList(context.Background(), omadacId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.GetGlobalTopThreatList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalTopThreatList`: OperationResponseTop5ThreatNumOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.GetGlobalTopThreatList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalTopThreatListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time | 

### Return type

[**OperationResponseTop5ThreatNumOpenApiVO**](OperationResponseTop5ThreatNumOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatCount

> OperationResponseThreatSeverityOpenApiVO GetThreatCount(ctx, omadacId).StartTime(startTime).EndTime(endTime).Execute()

Get Threat Count



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
	startTime := int64(789) // int64 | Start Time
	endTime := int64(789) // int64 | End Time

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.GetThreatCount(context.Background(), omadacId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.GetThreatCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatCount`: OperationResponseThreatSeverityOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.GetThreatCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time | 

### Return type

[**OperationResponseThreatSeverityOpenApiVO**](OperationResponseThreatSeverityOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperateGlobalThreats

> OperationResponse OperateGlobalThreats(ctx, omadacId).OperateGlobalThreatOpenApiVO(operateGlobalThreatOpenApiVO).Execute()

Operate Global Threats



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
	operateGlobalThreatOpenApiVO := *openapiclient.NewOperateGlobalThreatOpenApiVO(int32(123)) // OperateGlobalThreatOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalSecurityAPI.OperateGlobalThreats(context.Background(), omadacId).OperateGlobalThreatOpenApiVO(operateGlobalThreatOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalSecurityAPI.OperateGlobalThreats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperateGlobalThreats`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `GlobalSecurityAPI.OperateGlobalThreats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperateGlobalThreatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operateGlobalThreatOpenApiVO** | [**OperateGlobalThreatOpenApiVO**](OperateGlobalThreatOpenApiVO.md) |  | 

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


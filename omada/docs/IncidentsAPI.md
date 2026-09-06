# IncidentsAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAggregateIncident**](IncidentsAPI.md#deleteaggregateincident) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/incidents | Delete aggregate incident
[**DeleteIncidents**](IncidentsAPI.md#deleteincidents) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/incident | Delete incidents
[**GetAnomalyEventSettingForSite**](IncidentsAPI.md#getanomalyeventsettingforsite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/setting | Get site incident event setting
[**GetCategoryEventCount**](IncidentsAPI.md#getcategoryeventcount) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/category | Get incident category event count
[**GetGridDeviceIncidentList**](IncidentsAPI.md#getgriddeviceincidentlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/incident/{mac} | Get grid device incident List
[**GetGridIncidentList**](IncidentsAPI.md#getgridincidentlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/incidents | Get grid incident List
[**GetGridStackIncidentList**](IncidentsAPI.md#getgridstackincidentlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/incident/stack/{stackId} | Get grid stack incident list
[**GetIncidentDetail**](IncidentsAPI.md#getincidentdetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/incident/{incidentId} | Get incident detail
[**GetIncidentDistribution**](IncidentsAPI.md#getincidentdistribution) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/overview/incident-distribution | Get incident overview distribution
[**GetIncidentsRanking**](IncidentsAPI.md#getincidentsranking) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/incidents/ranking-cards | Get incidents ranking cards
[**GetInfluencingClients**](IncidentsAPI.md#getinfluencingclients) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/{mac}/influencing-clients | Get Influencing Clients by an incident
[**GetInfluencingDevices**](IncidentsAPI.md#getinfluencingdevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/{mac}/influencing-devices | Get Influencing Devices by an incident
[**GetNotificationTemplate1**](IncidentsAPI.md#getnotificationtemplate1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/notification | Get incident notification info
[**GetObjectHealthIncident**](IncidentsAPI.md#getobjecthealthincident) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/health/incident/{mac} | Get device/client health incidents
[**GetOverviewStatistics**](IncidentsAPI.md#getoverviewstatistics) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/overview/statistics | Get incident overview statistics
[**GetRankingCards**](IncidentsAPI.md#getrankingcards) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/overview/ranking-cards | Get incident overview ranking cards
[**GetUnresolvedIncidentCount**](IncidentsAPI.md#getunresolvedincidentcount) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/unresolved-incidents | Get unresolved incident count
[**ModifyAggregateIncidentStatus**](IncidentsAPI.md#modifyaggregateincidentstatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/incidents/status | Modify aggregate incident status
[**ModifyAnomalyEventSettingForSite**](IncidentsAPI.md#modifyanomalyeventsettingforsite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/setting/modify | Modify site incident event setting
[**ModifyIncidentStatus**](IncidentsAPI.md#modifyincidentstatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/incident/status | Modify incident status
[**ModifyNotification**](IncidentsAPI.md#modifynotification) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/notification | Set incident notification info



## DeleteAggregateIncident

> OperationResponseObject DeleteAggregateIncident(ctx, omadacId, siteId).AnomalyTimerSettingVO(anomalyTimerSettingVO).Execute()

Delete aggregate incident



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
	anomalyTimerSettingVO := *openapiclient.NewAnomalyTimerSettingVO(int64(123), int64(123), int32(123)) // AnomalyTimerSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.DeleteAggregateIncident(context.Background(), omadacId, siteId).AnomalyTimerSettingVO(anomalyTimerSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.DeleteAggregateIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAggregateIncident`: OperationResponseObject
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.DeleteAggregateIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAggregateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **anomalyTimerSettingVO** | [**AnomalyTimerSettingVO**](AnomalyTimerSettingVO.md) |  | 

### Return type

[**OperationResponseObject**](OperationResponseObject.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncidents

> OperationResponseObject DeleteIncidents(ctx, omadacId, siteId).DeleteAnomalyListSettingVO(deleteAnomalyListSettingVO).Execute()

Delete incidents



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
	deleteAnomalyListSettingVO := *openapiclient.NewDeleteAnomalyListSettingVO() // DeleteAnomalyListSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.DeleteIncidents(context.Background(), omadacId, siteId).DeleteAnomalyListSettingVO(deleteAnomalyListSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.DeleteIncidents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIncidents`: OperationResponseObject
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.DeleteIncidents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteAnomalyListSettingVO** | [**DeleteAnomalyListSettingVO**](DeleteAnomalyListSettingVO.md) |  | 

### Return type

[**OperationResponseObject**](OperationResponseObject.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnomalyEventSettingForSite

> OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO GetAnomalyEventSettingForSite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).FiltersEnable(filtersEnable).FiltersLevel(filtersLevel).FiltersCategory(filtersCategory).Execute()

Get site incident event setting



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
	filtersEnable := int32(56) // int32 | Filter query parameters, support field enable, it should be a value as follows: 0:enable, 1:disable, example:0. (optional)
	filtersLevel := int32(56) // int32 | Filter query parameters, support field level, it should be a value as follows: 0:Critical, 1:Error, 2:Warning, 3: Info, example:1. (optional)
	filtersCategory := "filtersCategory_example" // string | Filter query parameters, support field category, one or more categories, each category should be a value as follows: 11:Access, 12:Authentication, 13:Roaming, 14:Wireless Network, 15:Wired Network, 16:Link, 17:WAN and Services, 18:Device Status, 19:Security, example: 11 Or 12,13,14 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetAnomalyEventSettingForSite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).FiltersEnable(filtersEnable).FiltersLevel(filtersLevel).FiltersCategory(filtersCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetAnomalyEventSettingForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnomalyEventSettingForSite`: OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetAnomalyEventSettingForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnomalyEventSettingForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersEnable** | **int32** | Filter query parameters, support field enable, it should be a value as follows: 0:enable, 1:disable, example:0. | 
 **filtersLevel** | **int32** | Filter query parameters, support field level, it should be a value as follows: 0:Critical, 1:Error, 2:Warning, 3: Info, example:1. | 
 **filtersCategory** | **string** | Filter query parameters, support field category, one or more categories, each category should be a value as follows: 11:Access, 12:Authentication, 13:Roaming, 14:Wireless Network, 15:Wired Network, 16:Link, 17:WAN and Services, 18:Device Status, 19:Security, example: 11 Or 12,13,14 | 

### Return type

[**OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO**](OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryEventCount

> OperationResponseListAnomalyCategoryEventCountVO GetCategoryEventCount(ctx, omadacId, siteId).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Category(category).Execute()

Get incident category event count



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
	filtersStartTime := int64(789) // int64 | Start time in milliseconds
	filtersEndTime := int64(789) // int64 | End time in milliseconds
	category := int32(56) // int32 | Incident category (11-19)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetCategoryEventCount(context.Background(), omadacId, siteId).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetCategoryEventCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryEventCount`: OperationResponseListAnomalyCategoryEventCountVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetCategoryEventCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryEventCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersStartTime** | **int64** | Start time in milliseconds | 
 **filtersEndTime** | **int64** | End time in milliseconds | 
 **category** | **int32** | Incident category (11-19) | 

### Return type

[**OperationResponseListAnomalyCategoryEventCountVO**](OperationResponseListAnomalyCategoryEventCountVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDeviceIncidentList

> OperationResponseAnomalyGridVOAnomalyAggregateVO GetGridDeviceIncidentList(ctx, omadacId, siteId, mac).Page(page).PageSize(pageSize).FiltersStatus(filtersStatus).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()

Get grid device incident List



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
	mac := "mac_example" // string | mac
	filtersStatus := int32(56) // int32 | Incident status filter (optional)
	filtersStartTime := int64(789) // int64 | Start time in milliseconds (optional)
	filtersEndTime := int64(789) // int64 | End time in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetGridDeviceIncidentList(context.Background(), omadacId, siteId, mac).Page(page).PageSize(pageSize).FiltersStatus(filtersStatus).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetGridDeviceIncidentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDeviceIncidentList`: OperationResponseAnomalyGridVOAnomalyAggregateVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetGridDeviceIncidentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | mac | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDeviceIncidentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 

 **filtersStatus** | **int32** | Incident status filter | 
 **filtersStartTime** | **int64** | Start time in milliseconds | 
 **filtersEndTime** | **int64** | End time in milliseconds | 

### Return type

[**OperationResponseAnomalyGridVOAnomalyAggregateVO**](OperationResponseAnomalyGridVOAnomalyAggregateVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridIncidentList

> OperationResponseAnomalyGridVOAnomalyAggregateVO GetGridIncidentList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).FiltersStatus(filtersStatus).FiltersCategory(filtersCategory).FiltersAnomalyCode(filtersAnomalyCode).Execute()

Get grid incident List



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
	filtersStartTime := int64(789) // int64 | Start time in milliseconds
	filtersEndTime := int64(789) // int64 | End time in milliseconds
	filtersStatus := int32(56) // int32 | Incident status filter (optional)
	filtersCategory := "filtersCategory_example" // string | Incident category filter (optional)
	filtersAnomalyCode := "filtersAnomalyCode_example" // string | Anomaly code filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetGridIncidentList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).FiltersStatus(filtersStatus).FiltersCategory(filtersCategory).FiltersAnomalyCode(filtersAnomalyCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetGridIncidentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridIncidentList`: OperationResponseAnomalyGridVOAnomalyAggregateVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetGridIncidentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridIncidentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersStartTime** | **int64** | Start time in milliseconds | 
 **filtersEndTime** | **int64** | End time in milliseconds | 
 **filtersStatus** | **int32** | Incident status filter | 
 **filtersCategory** | **string** | Incident category filter | 
 **filtersAnomalyCode** | **string** | Anomaly code filter | 

### Return type

[**OperationResponseAnomalyGridVOAnomalyAggregateVO**](OperationResponseAnomalyGridVOAnomalyAggregateVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStackIncidentList

> OperationResponseAnomalyGridVOAnomalyAggregateVO GetGridStackIncidentList(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).FiltersStatus(filtersStatus).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()

Get grid stack incident list



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
	stackId := "stackId_example" // string | Stack ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersStatus := int32(56) // int32 | Incident status filter. 0: unresolved, 1: resolved, 2: ignored, 3: ongoing, null: all (optional)
	filtersStartTime := int64(789) // int64 | Start time in milliseconds (optional)
	filtersEndTime := int64(789) // int64 | End time in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetGridStackIncidentList(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).FiltersStatus(filtersStatus).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetGridStackIncidentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStackIncidentList`: OperationResponseAnomalyGridVOAnomalyAggregateVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetGridStackIncidentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridStackIncidentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersStatus** | **int32** | Incident status filter. 0: unresolved, 1: resolved, 2: ignored, 3: ongoing, null: all | 
 **filtersStartTime** | **int64** | Start time in milliseconds | 
 **filtersEndTime** | **int64** | End time in milliseconds | 

### Return type

[**OperationResponseAnomalyGridVOAnomalyAggregateVO**](OperationResponseAnomalyGridVOAnomalyAggregateVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentDetail

> OperationResponseAnomalyVO GetIncidentDetail(ctx, omadacId, siteId, incidentId).Execute()

Get incident detail



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
	incidentId := "incidentId_example" // string | Incident ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetIncidentDetail(context.Background(), omadacId, siteId, incidentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetIncidentDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncidentDetail`: OperationResponseAnomalyVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetIncidentDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**incidentId** | **string** | Incident ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseAnomalyVO**](OperationResponseAnomalyVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentDistribution

> OperationResponseIncidentDistributionOpenApiVO GetIncidentDistribution(ctx, omadacId, siteId).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()

Get incident overview distribution



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
	filtersStartTime := int64(789) // int64 | Start time in milliseconds
	filtersEndTime := int64(789) // int64 | End time in milliseconds

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetIncidentDistribution(context.Background(), omadacId, siteId).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetIncidentDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncidentDistribution`: OperationResponseIncidentDistributionOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetIncidentDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersStartTime** | **int64** | Start time in milliseconds | 
 **filtersEndTime** | **int64** | End time in milliseconds | 

### Return type

[**OperationResponseIncidentDistributionOpenApiVO**](OperationResponseIncidentDistributionOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentsRanking

> OperationResponseIncidentRankingResultVO GetIncidentsRanking(ctx, omadacId, siteId).IncidentRankingQueryVO(incidentRankingQueryVO).Execute()

Get incidents ranking cards



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
	incidentRankingQueryVO := *openapiclient.NewIncidentRankingQueryVO(int32(1001001), int64(123), int64(123)) // IncidentRankingQueryVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetIncidentsRanking(context.Background(), omadacId, siteId).IncidentRankingQueryVO(incidentRankingQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetIncidentsRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncidentsRanking`: OperationResponseIncidentRankingResultVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetIncidentsRanking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentsRankingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **incidentRankingQueryVO** | [**IncidentRankingQueryVO**](IncidentRankingQueryVO.md) |  | 

### Return type

[**OperationResponseIncidentRankingResultVO**](OperationResponseIncidentRankingResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfluencingClients

> OperationResponseGridVOInfluencingClientVO GetInfluencingClients(ctx, omadacId, siteId, mac).AnomalyCode(anomalyCode).Page(page).PageSize(pageSize).StartTime(startTime).EndTime(endTime).Status(status).SortsName(sortsName).FiltersActive(filtersActive).SearchKey(searchKey).Execute()

Get Influencing Clients by an incident



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
	mac := "mac_example" // string | Client MAC address, like AA-BB-CC-DD-EE-FF
	anomalyCode := "anomalyCode_example" // string | Anomaly Code
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–100.
	startTime := int64(789) // int64 | Start time of the query range, in milliseconds (epoch). If both startTime and endTime are omitted, the server defaults to the last 30 days. (optional)
	endTime := int64(789) // int64 | End time of the query range, in milliseconds (epoch). If both startTime and endTime are omitted, the server defaults to the last 30 days. (optional)
	status := int32(56) // int32 | Status of the incident. 0: unresolved, 1: resolved, 2: ignored, 3: ongoing, null: all (optional)
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersActive := "filtersActive_example" // string | Filter query parameters, support field active: true/false. (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field clientName,mac,ip. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetInfluencingClients(context.Background(), omadacId, siteId, mac).AnomalyCode(anomalyCode).Page(page).PageSize(pageSize).StartTime(startTime).EndTime(endTime).Status(status).SortsName(sortsName).FiltersActive(filtersActive).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetInfluencingClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfluencingClients`: OperationResponseGridVOInfluencingClientVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetInfluencingClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | Client MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfluencingClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anomalyCode** | **string** | Anomaly Code | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–100. | 
 **startTime** | **int64** | Start time of the query range, in milliseconds (epoch). If both startTime and endTime are omitted, the server defaults to the last 30 days. | 
 **endTime** | **int64** | End time of the query range, in milliseconds (epoch). If both startTime and endTime are omitted, the server defaults to the last 30 days. | 
 **status** | **int32** | Status of the incident. 0: unresolved, 1: resolved, 2: ignored, 3: ongoing, null: all | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersActive** | **string** | Filter query parameters, support field active: true/false. | 
 **searchKey** | **string** | Fuzzy query parameters, support field clientName,mac,ip. | 

### Return type

[**OperationResponseGridVOInfluencingClientVO**](OperationResponseGridVOInfluencingClientVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfluencingDevices

> OperationResponseGridVOInfluencingDeviceVO GetInfluencingDevices(ctx, omadacId, siteId, mac).AnomalyCode(anomalyCode).Page(page).PageSize(pageSize).StartTime(startTime).EndTime(endTime).Status(status).SortsName(sortsName).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()

Get Influencing Devices by an incident



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
	mac := "mac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	anomalyCode := "anomalyCode_example" // string | Anomaly Code
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–100.
	startTime := int64(789) // int64 | Start time of the query range, in milliseconds. If both startTime and endTime are omitted, the server defaults to the last 30 days. (optional)
	endTime := int64(789) // int64 | End time of the query range, in milliseconds. If both startTime and endTime are omitted, the server defaults to the last 30 days. (optional)
	status := int32(56) // int32 | Status of the incident. 0: unresolved, 1: resolved, 2: ignored, 3: ongoing, null: all (optional)
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersStatus := "filtersStatus_example" // string | Filter query parameters, support field status (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field deviceName,mac,ip. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetInfluencingDevices(context.Background(), omadacId, siteId, mac).AnomalyCode(anomalyCode).Page(page).PageSize(pageSize).StartTime(startTime).EndTime(endTime).Status(status).SortsName(sortsName).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetInfluencingDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfluencingDevices`: OperationResponseGridVOInfluencingDeviceVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetInfluencingDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfluencingDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anomalyCode** | **string** | Anomaly Code | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–100. | 
 **startTime** | **int64** | Start time of the query range, in milliseconds. If both startTime and endTime are omitted, the server defaults to the last 30 days. | 
 **endTime** | **int64** | End time of the query range, in milliseconds. If both startTime and endTime are omitted, the server defaults to the last 30 days. | 
 **status** | **int32** | Status of the incident. 0: unresolved, 1: resolved, 2: ignored, 3: ongoing, null: all | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersStatus** | **string** | Filter query parameters, support field status | 
 **searchKey** | **string** | Fuzzy query parameters, support field deviceName,mac,ip. | 

### Return type

[**OperationResponseGridVOInfluencingDeviceVO**](OperationResponseGridVOInfluencingDeviceVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationTemplate1

> OperationResponseNotificationOpenApiVO GetNotificationTemplate1(ctx, omadacId, siteId).Execute()

Get incident notification info



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
	resp, r, err := apiClient.IncidentsAPI.GetNotificationTemplate1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetNotificationTemplate1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationTemplate1`: OperationResponseNotificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetNotificationTemplate1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationTemplate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseNotificationOpenApiVO**](OperationResponseNotificationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectHealthIncident

> OperationResponseListAnomalyBriefCountVO GetObjectHealthIncident(ctx, omadacId, siteId, mac).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()

Get device/client health incidents



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
	filtersStartTime := int64(789) // int64 | Start time in milliseconds
	filtersEndTime := int64(789) // int64 | End time in milliseconds
	mac := "mac_example" // string | mac

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetObjectHealthIncident(context.Background(), omadacId, siteId, mac).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetObjectHealthIncident``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectHealthIncident`: OperationResponseListAnomalyBriefCountVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetObjectHealthIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | mac | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectHealthIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersStartTime** | **int64** | Start time in milliseconds | 
 **filtersEndTime** | **int64** | End time in milliseconds | 


### Return type

[**OperationResponseListAnomalyBriefCountVO**](OperationResponseListAnomalyBriefCountVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverviewStatistics

> OperationResponseOverviewStatisticsOpenApiVO GetOverviewStatistics(ctx, omadacId, siteId).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()

Get incident overview statistics



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
	filtersStartTime := int64(1646356605000) // int64 | Query start time, in milliseconds
	filtersEndTime := int64(1646360205000) // int64 | Query end time, in milliseconds

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetOverviewStatistics(context.Background(), omadacId, siteId).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetOverviewStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverviewStatistics`: OperationResponseOverviewStatisticsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetOverviewStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOverviewStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersStartTime** | **int64** | Query start time, in milliseconds | 
 **filtersEndTime** | **int64** | Query end time, in milliseconds | 

### Return type

[**OperationResponseOverviewStatisticsOpenApiVO**](OperationResponseOverviewStatisticsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRankingCards

> OperationResponseIncidentRankingResultVO GetRankingCards(ctx, omadacId, siteId).IncidentRankingQueryVO(incidentRankingQueryVO).Execute()

Get incident overview ranking cards



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
	incidentRankingQueryVO := *openapiclient.NewIncidentRankingQueryVO(int32(1001001), int64(123), int64(123)) // IncidentRankingQueryVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetRankingCards(context.Background(), omadacId, siteId).IncidentRankingQueryVO(incidentRankingQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetRankingCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRankingCards`: OperationResponseIncidentRankingResultVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetRankingCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRankingCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **incidentRankingQueryVO** | [**IncidentRankingQueryVO**](IncidentRankingQueryVO.md) |  | 

### Return type

[**OperationResponseIncidentRankingResultVO**](OperationResponseIncidentRankingResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnresolvedIncidentCount

> ActiveIncidentCountVO GetUnresolvedIncidentCount(ctx, omadacId, siteId).Execute()

Get unresolved incident count



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
	resp, r, err := apiClient.IncidentsAPI.GetUnresolvedIncidentCount(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetUnresolvedIncidentCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnresolvedIncidentCount`: ActiveIncidentCountVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetUnresolvedIncidentCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnresolvedIncidentCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActiveIncidentCountVO**](ActiveIncidentCountVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAggregateIncidentStatus

> OperationResponseObject ModifyAggregateIncidentStatus(ctx, omadacId, siteId).AnomalyTimerSettingVO(anomalyTimerSettingVO).Execute()

Modify aggregate incident status



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
	anomalyTimerSettingVO := *openapiclient.NewAnomalyTimerSettingVO(int64(123), int64(123), int32(123)) // AnomalyTimerSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.ModifyAggregateIncidentStatus(context.Background(), omadacId, siteId).AnomalyTimerSettingVO(anomalyTimerSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.ModifyAggregateIncidentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAggregateIncidentStatus`: OperationResponseObject
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.ModifyAggregateIncidentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAggregateIncidentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **anomalyTimerSettingVO** | [**AnomalyTimerSettingVO**](AnomalyTimerSettingVO.md) |  | 

### Return type

[**OperationResponseObject**](OperationResponseObject.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAnomalyEventSettingForSite

> OperationResponseWithoutResult ModifyAnomalyEventSettingForSite(ctx, omadacId, siteId).AnomalyEventSettingEditOpenApiVO(anomalyEventSettingEditOpenApiVO).Execute()

Modify site incident event setting



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
	anomalyEventSettingEditOpenApiVO := *openapiclient.NewAnomalyEventSettingEditOpenApiVO("01001001") // AnomalyEventSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.ModifyAnomalyEventSettingForSite(context.Background(), omadacId, siteId).AnomalyEventSettingEditOpenApiVO(anomalyEventSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.ModifyAnomalyEventSettingForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAnomalyEventSettingForSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.ModifyAnomalyEventSettingForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAnomalyEventSettingForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **anomalyEventSettingEditOpenApiVO** | [**AnomalyEventSettingEditOpenApiVO**](AnomalyEventSettingEditOpenApiVO.md) |  | 

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


## ModifyIncidentStatus

> OperationResponseObject ModifyIncidentStatus(ctx, omadacId, siteId).AnomalyListSettingVO(anomalyListSettingVO).Execute()

Modify incident status



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
	anomalyListSettingVO := *openapiclient.NewAnomalyListSettingVO() // AnomalyListSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.ModifyIncidentStatus(context.Background(), omadacId, siteId).AnomalyListSettingVO(anomalyListSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.ModifyIncidentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIncidentStatus`: OperationResponseObject
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.ModifyIncidentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIncidentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **anomalyListSettingVO** | [**AnomalyListSettingVO**](AnomalyListSettingVO.md) |  | 

### Return type

[**OperationResponseObject**](OperationResponseObject.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyNotification

> OperationResponseSetNotificationOpenApiVO ModifyNotification(ctx, omadacId, siteId).SetNotificationOpenApiVO(setNotificationOpenApiVO).Execute()

Set incident notification info



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
	setNotificationOpenApiVO := *openapiclient.NewSetNotificationOpenApiVO(*openapiclient.NewAlertEmailSettingVO(false), *openapiclient.NewNotificationConfigurationOpenApiVO(map[string]bool{"key": false}, map[string]bool{"key": false}, map[string]bool{"key": false})) // SetNotificationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.ModifyNotification(context.Background(), omadacId, siteId).SetNotificationOpenApiVO(setNotificationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.ModifyNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyNotification`: OperationResponseSetNotificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.ModifyNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setNotificationOpenApiVO** | [**SetNotificationOpenApiVO**](SetNotificationOpenApiVO.md) |  | 

### Return type

[**OperationResponseSetNotificationOpenApiVO**](OperationResponseSetNotificationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


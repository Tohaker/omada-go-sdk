# \ThreatManagementAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteThreatList**](ThreatManagementAPI.md#DeleteThreatList) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ips/threat | Delete archived ips threat
[**GetThreatDetail**](ThreatManagementAPI.md#GetThreatDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ips/threat/{threatId} | Get threat entry detail
[**GetThreatList**](ThreatManagementAPI.md#GetThreatList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ips/grid/threat-management | Get grid threat list
[**OperateThreats**](ThreatManagementAPI.md#OperateThreats) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ips/threat/multi-ops | Block/Isolate Device/Signature Suppression/Allow ips threat



## DeleteThreatList

> OperationResponse DeleteThreatList(ctx, omadacId, siteId).DeleteIpsThreat(deleteIpsThreat).Execute()

Delete archived ips threat



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
	deleteIpsThreat := *openapiclient.NewDeleteIpsThreat() // DeleteIpsThreat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatManagementAPI.DeleteThreatList(context.Background(), omadacId, siteId).DeleteIpsThreat(deleteIpsThreat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatManagementAPI.DeleteThreatList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteThreatList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatManagementAPI.DeleteThreatList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThreatListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteIpsThreat** | [**DeleteIpsThreat**](DeleteIpsThreat.md) |  | 

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


## GetThreatDetail

> OperationResponseGetIpsThreatDetail GetThreatDetail(ctx, omadacId, siteId, threatId).Time(time).Execute()

Get threat entry detail



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
	threatId := "threatId_example" // string | threatId
	time := int64(789) // int64 | Timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatManagementAPI.GetThreatDetail(context.Background(), omadacId, siteId, threatId).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatManagementAPI.GetThreatDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatDetail`: OperationResponseGetIpsThreatDetail
	fmt.Fprintf(os.Stdout, "Response from `ThreatManagementAPI.GetThreatDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**threatId** | **string** | threatId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **time** | **int64** | Timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseGetIpsThreatDetail**](OperationResponseGetIpsThreatDetail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatList

> OperationResponseGridVOGetGridIpsThreat GetThreatList(ctx, omadacId, siteId).Archived(archived).Page(page).PageSize(pageSize).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).FiltersSeverity(filtersSeverity).SortsTime(sortsTime).SearchKey(searchKey).Execute()

Get grid threat list



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
	archived := true // bool | archived
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersStartTime := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	filtersEndTime := int64(789) // int64 | End timestamp, in seconds, such as 1682000000
	filtersSeverity := int32(56) // int32 | Threat Severity, such as 0:Critical, 1: Major, 2:Moderate, 3:Minor, 4:Low (optional)
	sortsTime := "sortsTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field Threat Description/Classification/Classification Description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatManagementAPI.GetThreatList(context.Background(), omadacId, siteId).Archived(archived).Page(page).PageSize(pageSize).FiltersStartTime(filtersStartTime).FiltersEndTime(filtersEndTime).FiltersSeverity(filtersSeverity).SortsTime(sortsTime).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatManagementAPI.GetThreatList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatList`: OperationResponseGridVOGetGridIpsThreat
	fmt.Fprintf(os.Stdout, "Response from `ThreatManagementAPI.GetThreatList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | archived | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersStartTime** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **filtersEndTime** | **int64** | End timestamp, in seconds, such as 1682000000 | 
 **filtersSeverity** | **int32** | Threat Severity, such as 0:Critical, 1: Major, 2:Moderate, 3:Minor, 4:Low | 
 **sortsTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field Threat Description/Classification/Classification Description | 

### Return type

[**OperationResponseGridVOGetGridIpsThreat**](OperationResponseGridVOGetGridIpsThreat.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperateThreats

> OperationResponse OperateThreats(ctx, omadacId, siteId).IpsOperateThreat(ipsOperateThreat).Execute()

Block/Isolate Device/Signature Suppression/Allow ips threat



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
	ipsOperateThreat := *openapiclient.NewIpsOperateThreat(int32(1)) // IpsOperateThreat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatManagementAPI.OperateThreats(context.Background(), omadacId, siteId).IpsOperateThreat(ipsOperateThreat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatManagementAPI.OperateThreats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperateThreats`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatManagementAPI.OperateThreats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperateThreatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ipsOperateThreat** | [**IpsOperateThreat**](IpsOperateThreat.md) |  | 

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


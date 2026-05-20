# WiFiCallingTrafficAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportBatchFullChannelDetectResultData**](WiFiCallingTrafficAPI.md#exportbatchfullchanneldetectresultdata) | **Get** /openapi/v1/{omadacId}/files/sites/{siteId}/wifi-calling/summary/{format} | Export wi-fi calling traffic summary
[**GetGridWifiCallingTrafficResult**](WiFiCallingTrafficAPI.md#getgridwificallingtrafficresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/wifi-calling/grid/summary | Get Wi-Fi Calling Summary



## ExportBatchFullChannelDetectResultData

> OperationResponseWithoutResult ExportBatchFullChannelDetectResultData(ctx, omadacId, siteId, format).TimeStart(timeStart).TimeEnd(timeEnd).SearchKey(searchKey).Execute()

Export wi-fi calling traffic summary



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
	format := "format_example" // string | export data format
	timeStart := int64(789) // int64 | Filter query parameters, support field time range: start timestamp (second).
	timeEnd := int64(789) // int64 | Filter query parameters, support field time range: end timestamp (second).
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field clientName,clientMac,ip,ssid,domain (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiFiCallingTrafficAPI.ExportBatchFullChannelDetectResultData(context.Background(), omadacId, siteId, format).TimeStart(timeStart).TimeEnd(timeEnd).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiFiCallingTrafficAPI.ExportBatchFullChannelDetectResultData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportBatchFullChannelDetectResultData`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiFiCallingTrafficAPI.ExportBatchFullChannelDetectResultData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**format** | **string** | export data format | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportBatchFullChannelDetectResultDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timeStart** | **int64** | Filter query parameters, support field time range: start timestamp (second). | 
 **timeEnd** | **int64** | Filter query parameters, support field time range: end timestamp (second). | 
 **searchKey** | **string** | Fuzzy query parameters, support field clientName,clientMac,ip,ssid,domain | 

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


## GetGridWifiCallingTrafficResult

> OperationResponseWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO GetGridWifiCallingTrafficResult(ctx, omadacId, siteId).QueryTopSsid(queryTopSsid).QueryTopEPDG(queryTopEPDG).QueryClients(queryClients).QueryClientsByMac(queryClientsByMac).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).TopK(topK).SortsApName(sortsApName).SortsCarrierName(sortsCarrierName).SortsWifiCallingProfileName(sortsWifiCallingProfileName).SortsClientName(sortsClientName).SortsIp(sortsIp).SortsDomain(sortsDomain).SortsTrafficDown(sortsTrafficDown).SortsTrafficUp(sortsTrafficUp).SortsTotalTraffic(sortsTotalTraffic).SortsClientMac(sortsClientMac).SortsPriority(sortsPriority).SortsSsid(sortsSsid).SortsStartTime(sortsStartTime).SortsEndTime(sortsEndTime).SearchKey(searchKey).Execute()

Get Wi-Fi Calling Summary



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
	queryTopSsid := true // bool | Whether to query top k ssids by traffic.
	queryTopEPDG := true // bool | Whether to query top k ePDGs by traffic.
	queryClients := true // bool | Whether to query client history.
	queryClientsByMac := true // bool | Whether to query client history by client MAC.
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field time range: start timestamp (second).
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field time range: end timestamp (second).
	topK := int32(56) // int32 | Get the top-k SSIDs or EPDGs based on voice call traffic statistics. Parameter [topK] should not be null when parameter [queryTopEPDG] or parameter [queryTopEPDG] is true. (optional)
	sortsApName := "sortsApName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsCarrierName := "sortsCarrierName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsWifiCallingProfileName := "sortsWifiCallingProfileName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsClientName := "sortsClientName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsDomain := "sortsDomain_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsTrafficDown := "sortsTrafficDown_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsTrafficUp := "sortsTrafficUp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsTotalTraffic := "sortsTotalTraffic_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsClientMac := "sortsClientMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsPriority := "sortsPriority_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsSsid := "sortsSsid_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStartTime := "sortsStartTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsEndTime := "sortsEndTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field clientName,clientMac,ip,ssid,domain (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiFiCallingTrafficAPI.GetGridWifiCallingTrafficResult(context.Background(), omadacId, siteId).QueryTopSsid(queryTopSsid).QueryTopEPDG(queryTopEPDG).QueryClients(queryClients).QueryClientsByMac(queryClientsByMac).Page(page).PageSize(pageSize).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).TopK(topK).SortsApName(sortsApName).SortsCarrierName(sortsCarrierName).SortsWifiCallingProfileName(sortsWifiCallingProfileName).SortsClientName(sortsClientName).SortsIp(sortsIp).SortsDomain(sortsDomain).SortsTrafficDown(sortsTrafficDown).SortsTrafficUp(sortsTrafficUp).SortsTotalTraffic(sortsTotalTraffic).SortsClientMac(sortsClientMac).SortsPriority(sortsPriority).SortsSsid(sortsSsid).SortsStartTime(sortsStartTime).SortsEndTime(sortsEndTime).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiFiCallingTrafficAPI.GetGridWifiCallingTrafficResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridWifiCallingTrafficResult`: OperationResponseWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiFiCallingTrafficAPI.GetGridWifiCallingTrafficResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridWifiCallingTrafficResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryTopSsid** | **bool** | Whether to query top k ssids by traffic. | 
 **queryTopEPDG** | **bool** | Whether to query top k ePDGs by traffic. | 
 **queryClients** | **bool** | Whether to query client history. | 
 **queryClientsByMac** | **bool** | Whether to query client history by client MAC. | 
 **page** | **int32** |  | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersTimeStart** | **int64** | Filter query parameters, support field time range: start timestamp (second). | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field time range: end timestamp (second). | 
 **topK** | **int32** | Get the top-k SSIDs or EPDGs based on voice call traffic statistics. Parameter [topK] should not be null when parameter [queryTopEPDG] or parameter [queryTopEPDG] is true. | 
 **sortsApName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsCarrierName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsWifiCallingProfileName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsClientName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsDomain** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsTrafficDown** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsTrafficUp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsTotalTraffic** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsClientMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsPriority** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsSsid** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStartTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsEndTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field clientName,clientMac,ip,ssid,domain | 

### Return type

[**OperationResponseWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO**](OperationResponseWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


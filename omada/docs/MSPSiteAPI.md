# MSPSiteAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteList1**](MSPSiteAPI.md#getsitelist1) | **Get** /openapi/v1/msp/{mspId}/sites | Get msp site list
[**GetTags2**](MSPSiteAPI.md#gettags2) | **Get** /openapi/v1/msp/{mspId}/sites/tags | Get msp site tag list



## GetSiteList1

> OperationResponseGridVOSiteSummaryInfo GetSiteList1(ctx, mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).FiltersTag(filtersTag).FiltersType(filtersType).Execute()

Get msp site list



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
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag ID (optional)
	filtersType := "filtersType_example" // string | Filter query parameters, support field site type. 0: basic site; 1: pro site. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSiteAPI.GetSiteList1(context.Background(), mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).FiltersTag(filtersTag).FiltersType(filtersType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSiteAPI.GetSiteList1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteList1`: OperationResponseGridVOSiteSummaryInfo
	fmt.Fprintf(os.Stdout, "Response from `MSPSiteAPI.GetSiteList1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteList1Request struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags2

> []SiteTag GetTags2(ctx, mspId).Execute()

Get msp site tag list



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
	resp, r, err := apiClient.MSPSiteAPI.GetTags2(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSiteAPI.GetTags2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags2`: []SiteTag
	fmt.Fprintf(os.Stdout, "Response from `MSPSiteAPI.GetTags2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTags2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SiteTag**](SiteTag.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \MSPDeviceAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdoptOneForMsp**](MSPDeviceAPI.md#AdoptOneForMsp) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/devices/{deviceMac}/start-adopt | Start adopt device For Msp
[**BatchAdoptForMsp**](MSPDeviceAPI.md#BatchAdoptForMsp) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/cmd/devices/batch-adopt | batch adopt device in msp view
[**GetMspKnownDeviceList**](MSPDeviceAPI.md#GetMspKnownDeviceList) | **Get** /openapi/v1/msp/{mspId}/devices/known-devices | Get MSP known device list
[**GetMspUnknownDeviceList**](MSPDeviceAPI.md#GetMspUnknownDeviceList) | **Get** /openapi/v1/msp/{mspId}/devices/unknown-devices | Get MSP unknown device list
[**GetTags3**](MSPDeviceAPI.md#GetTags3) | **Get** /openapi/v1/msp/{mspId}/devices/tag | Get tag list



## AdoptOneForMsp

> OperationResponseWithoutResult AdoptOneForMsp(ctx, customerId, siteId, deviceMac, mspId).AdoptDeviceRequest(adoptDeviceRequest).Execute()

Start adopt device For Msp



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
	customerId := "customerId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	mspId := "mspId_example" // string | mspId
	adoptDeviceRequest := *openapiclient.NewAdoptDeviceRequest() // AdoptDeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPDeviceAPI.AdoptOneForMsp(context.Background(), customerId, siteId, deviceMac, mspId).AdoptDeviceRequest(adoptDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPDeviceAPI.AdoptOneForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdoptOneForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MSPDeviceAPI.AdoptOneForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**mspId** | **string** | mspId | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdoptOneForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **adoptDeviceRequest** | [**AdoptDeviceRequest**](AdoptDeviceRequest.md) |  | 

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


## BatchAdoptForMsp

> OperationResponseWithoutResult BatchAdoptForMsp(ctx, mspId, customerId, siteId).BatchAdoptDeviceRequest(batchAdoptDeviceRequest).Execute()

batch adopt device in msp view



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
	customerId := "customerId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	batchAdoptDeviceRequest := *openapiclient.NewBatchAdoptDeviceRequest() // BatchAdoptDeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPDeviceAPI.BatchAdoptForMsp(context.Background(), mspId, customerId, siteId).BatchAdoptDeviceRequest(batchAdoptDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPDeviceAPI.BatchAdoptForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchAdoptForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MSPDeviceAPI.BatchAdoptForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchAdoptForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchAdoptDeviceRequest** | [**BatchAdoptDeviceRequest**](BatchAdoptDeviceRequest.md) |  | 

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


## GetMspKnownDeviceList

> OperationResponseGridVOMspKnownDeviceOpenApiVO GetMspKnownDeviceList(ctx, mspId).Page(page).PageSize(pageSize).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).SearchSns(searchSns).FiltersTag(filtersTag).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()

Get MSP known device list



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
	searchMacs := "searchMacs_example" // string | Fuzzy query parameters, support field mac (optional)
	searchNames := "searchNames_example" // string | Fuzzy query parameters, support field name (optional)
	searchModels := "searchModels_example" // string | Fuzzy query parameters, support field model (optional)
	searchSns := "searchSns_example" // string | Fuzzy query parameters, support field sn (optional)
	filtersTag := "filtersTag_example" // string | Filter query parameters, support field tag ID (optional)
	filtersDeviceSeriesType := "filtersDeviceSeriesType_example" // string | Filter query parameters, support field Device series type. 0: basic; 1: pro. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPDeviceAPI.GetMspKnownDeviceList(context.Background(), mspId).Page(page).PageSize(pageSize).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).SearchSns(searchSns).FiltersTag(filtersTag).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPDeviceAPI.GetMspKnownDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspKnownDeviceList`: OperationResponseGridVOMspKnownDeviceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MSPDeviceAPI.GetMspKnownDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspKnownDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchMacs** | **string** | Fuzzy query parameters, support field mac | 
 **searchNames** | **string** | Fuzzy query parameters, support field name | 
 **searchModels** | **string** | Fuzzy query parameters, support field model | 
 **searchSns** | **string** | Fuzzy query parameters, support field sn | 
 **filtersTag** | **string** | Filter query parameters, support field tag ID | 
 **filtersDeviceSeriesType** | **string** | Filter query parameters, support field Device series type. 0: basic; 1: pro. | 

### Return type

[**OperationResponseGridVOMspKnownDeviceOpenApiVO**](OperationResponseGridVOMspKnownDeviceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspUnknownDeviceList

> OperationResponseGridVOMspUnknownDeviceOpenApiVO GetMspUnknownDeviceList(ctx, mspId).Page(page).PageSize(pageSize).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()

Get MSP unknown device list



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
	searchMacs := "searchMacs_example" // string | Fuzzy query parameters, support field mac (optional)
	searchNames := "searchNames_example" // string | Fuzzy query parameters, support field name (optional)
	searchModels := "searchModels_example" // string | Fuzzy query parameters, support field model (optional)
	filtersDeviceSeriesType := "filtersDeviceSeriesType_example" // string | Filter query parameters, support field Device series type. 0: basic; 1: pro. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPDeviceAPI.GetMspUnknownDeviceList(context.Background(), mspId).Page(page).PageSize(pageSize).SearchMacs(searchMacs).SearchNames(searchNames).SearchModels(searchModels).FiltersDeviceSeriesType(filtersDeviceSeriesType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPDeviceAPI.GetMspUnknownDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspUnknownDeviceList`: OperationResponseGridVOMspUnknownDeviceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MSPDeviceAPI.GetMspUnknownDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspUnknownDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchMacs** | **string** | Fuzzy query parameters, support field mac | 
 **searchNames** | **string** | Fuzzy query parameters, support field name | 
 **searchModels** | **string** | Fuzzy query parameters, support field model | 
 **filtersDeviceSeriesType** | **string** | Filter query parameters, support field Device series type. 0: basic; 1: pro. | 

### Return type

[**OperationResponseGridVOMspUnknownDeviceOpenApiVO**](OperationResponseGridVOMspUnknownDeviceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags3

> OperationResponseListTagRespOpenApiVO GetTags3(ctx, mspId).Execute()

Get tag list



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
	resp, r, err := apiClient.MSPDeviceAPI.GetTags3(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPDeviceAPI.GetTags3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags3`: OperationResponseListTagRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MSPDeviceAPI.GetTags3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTags3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseListTagRespOpenApiVO**](OperationResponseListTagRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


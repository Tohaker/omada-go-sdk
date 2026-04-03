# HotspotOperatorsAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHotspotOperator**](HotspotOperatorsAPI.md#createhotspotoperator) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/operators | Create a new hotspot operator
[**DeleteHotspotOperator**](HotspotOperatorsAPI.md#deletehotspotoperator) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/operators/{id} | Delete an existing hotspot operator
[**GetHotspotOperatorDetail**](HotspotOperatorsAPI.md#gethotspotoperatordetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/operators/{id} | Get hotspot operator detail
[**GetHotspotOperatorList**](HotspotOperatorsAPI.md#gethotspotoperatorlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/operators | Get hotspot operators list
[**ModifyHotspotOperator**](HotspotOperatorsAPI.md#modifyhotspotoperator) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/operators/{id} | Modify an existing hotspot operator



## CreateHotspotOperator

> OperationResponse CreateHotspotOperator(ctx, omadacId, siteId).HotspotOperator(hotspotOperator).Execute()

Create a new hotspot operator



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
	hotspotOperator := *openapiclient.NewHotspotOperator("Name_example", "Password_example", []string{"SelectedSites_example"}) // HotspotOperator | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HotspotOperatorsAPI.CreateHotspotOperator(context.Background(), omadacId, siteId).HotspotOperator(hotspotOperator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HotspotOperatorsAPI.CreateHotspotOperator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHotspotOperator`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `HotspotOperatorsAPI.CreateHotspotOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHotspotOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hotspotOperator** | [**HotspotOperator**](HotspotOperator.md) |  | 

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


## DeleteHotspotOperator

> OperationResponse DeleteHotspotOperator(ctx, omadacId, siteId, id).Execute()

Delete an existing hotspot operator



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
	id := "id_example" // string | Hotspot Operator ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HotspotOperatorsAPI.DeleteHotspotOperator(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HotspotOperatorsAPI.DeleteHotspotOperator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteHotspotOperator`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `HotspotOperatorsAPI.DeleteHotspotOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Hotspot Operator ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHotspotOperatorRequest struct via the builder pattern


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


## GetHotspotOperatorDetail

> OperationResponseOperatorResponse GetHotspotOperatorDetail(ctx, omadacId, siteId, id).Execute()

Get hotspot operator detail



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
	id := "id_example" // string | Hotspot Operator ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HotspotOperatorsAPI.GetHotspotOperatorDetail(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HotspotOperatorsAPI.GetHotspotOperatorDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHotspotOperatorDetail`: OperationResponseOperatorResponse
	fmt.Fprintf(os.Stdout, "Response from `HotspotOperatorsAPI.GetHotspotOperatorDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Hotspot Operator ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHotspotOperatorDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOperatorResponse**](OperationResponseOperatorResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHotspotOperatorList

> OperationResponseGridVOOperatorResponse GetHotspotOperatorList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()

Get hotspot operators list



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
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name,note (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HotspotOperatorsAPI.GetHotspotOperatorList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HotspotOperatorsAPI.GetHotspotOperatorList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHotspotOperatorList`: OperationResponseGridVOOperatorResponse
	fmt.Fprintf(os.Stdout, "Response from `HotspotOperatorsAPI.GetHotspotOperatorList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHotspotOperatorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name,note | 

### Return type

[**OperationResponseGridVOOperatorResponse**](OperationResponseGridVOOperatorResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyHotspotOperator

> OperationResponseString ModifyHotspotOperator(ctx, omadacId, siteId, id).HotspotOperator(hotspotOperator).Execute()

Modify an existing hotspot operator



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
	id := "id_example" // string | Hotspot Operator ID
	hotspotOperator := *openapiclient.NewHotspotOperator("Name_example", "Password_example", []string{"SelectedSites_example"}) // HotspotOperator | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HotspotOperatorsAPI.ModifyHotspotOperator(context.Background(), omadacId, siteId, id).HotspotOperator(hotspotOperator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HotspotOperatorsAPI.ModifyHotspotOperator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyHotspotOperator`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `HotspotOperatorsAPI.ModifyHotspotOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Hotspot Operator ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyHotspotOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **hotspotOperator** | [**HotspotOperator**](HotspotOperator.md) |  | 

### Return type

[**OperationResponseString**](OperationResponseString.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


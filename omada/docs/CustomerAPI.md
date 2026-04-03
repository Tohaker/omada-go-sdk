# CustomerAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFavoriteCustomer**](CustomerAPI.md#cancelfavoritecustomer) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/favorites/cancels | Cancel a favorite customer
[**CopyCustomer**](CustomerAPI.md#copycustomer) | **Post** /openapi/v1/msp/{mspId}/customers/copy | Create a customer by copying from existing customer
[**CreateCustomer**](CustomerAPI.md#createcustomer) | **Post** /openapi/v1/msp/{mspId}/customers | Create new customer
[**DeleteCustomer**](CustomerAPI.md#deletecustomer) | **Delete** /openapi/v1/msp/{mspId}/customers/{customerId} | Delete an existing customer
[**FavoriteCustomer**](CustomerAPI.md#favoritecustomer) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/favorites | Favorite a customer
[**GetCustomerDetail**](CustomerAPI.md#getcustomerdetail) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId} | Get customer info
[**GetCustomerList**](CustomerAPI.md#getcustomerlist) | **Get** /openapi/v1/msp/{mspId}/customers | Get customer list
[**GetCustomersInGlobalHaveSiteTemplateModify**](CustomerAPI.md#getcustomersinglobalhavesitetemplatemodify) | **Get** /openapi/v1/{omadacId}/copy/site/template/batch/customers | Obtain the customer ID with permission to modify site templates
[**ModifyCustomer**](CustomerAPI.md#modifycustomer) | **Patch** /openapi/v1/msp/{mspId}/customers/{customerId} | Modify an existing customer



## CancelFavoriteCustomer

> OperationResponseWithoutResult CancelFavoriteCustomer(ctx, mspId, customerId).Execute()

Cancel a favorite customer



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
	customerId := "customerId_example" // string | Customer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CancelFavoriteCustomer(context.Background(), mspId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CancelFavoriteCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelFavoriteCustomer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CancelFavoriteCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelFavoriteCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## CopyCustomer

> OperationResponseCustomerCopyOpenapiResVO CopyCustomer(ctx, mspId).CustomerCopyOpenapiVO(customerCopyOpenapiVO).Execute()

Create a customer by copying from existing customer



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
	customerCopyOpenapiVO := *openapiclient.NewCustomerCopyOpenapiVO("NewCustomerName_example", "SourceCustomerId_example") // CustomerCopyOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CopyCustomer(context.Background(), mspId).CustomerCopyOpenapiVO(customerCopyOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CopyCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyCustomer`: OperationResponseCustomerCopyOpenapiResVO
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CopyCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerCopyOpenapiVO** | [**CustomerCopyOpenapiVO**](CustomerCopyOpenapiVO.md) |  | 

### Return type

[**OperationResponseCustomerCopyOpenapiResVO**](OperationResponseCustomerCopyOpenapiResVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomer

> OperationResponseCreateCustomerOmadacVO CreateCustomer(ctx, mspId).CustomerOpenApiVO(customerOpenApiVO).Execute()

Create new customer



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
	customerOpenApiVO := *openapiclient.NewCustomerOpenApiVO() // CustomerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CreateCustomer(context.Background(), mspId).CustomerOpenApiVO(customerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CreateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomer`: OperationResponseCreateCustomerOmadacVO
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerOpenApiVO** | [**CustomerOpenApiVO**](CustomerOpenApiVO.md) |  | 

### Return type

[**OperationResponseCreateCustomerOmadacVO**](OperationResponseCreateCustomerOmadacVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomer

> OperationResponseWithoutResult DeleteCustomer(ctx, mspId, customerId).Execute()

Delete an existing customer



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
	customerId := "customerId_example" // string | Customer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.DeleteCustomer(context.Background(), mspId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.DeleteCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.DeleteCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## FavoriteCustomer

> OperationResponseWithoutResult FavoriteCustomer(ctx, mspId, customerId).Execute()

Favorite a customer



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
	customerId := "customerId_example" // string | Customer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.FavoriteCustomer(context.Background(), mspId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.FavoriteCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FavoriteCustomer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.FavoriteCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFavoriteCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetCustomerDetail

> OperationResponseCustomerInfoOpenApiVO GetCustomerDetail(ctx, mspId, customerId).Execute()

Get customer info



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
	customerId := "customerId_example" // string | Customer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.GetCustomerDetail(context.Background(), mspId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetCustomerDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerDetail`: OperationResponseCustomerInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetCustomerDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseCustomerInfoOpenApiVO**](OperationResponseCustomerInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerList

> OperationResponseGridVOCustomerInfoOpenApiVO GetCustomerList(ctx, mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()

Get customer list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.GetCustomerList(context.Background(), mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetCustomerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerList`: OperationResponseGridVOCustomerInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetCustomerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseGridVOCustomerInfoOpenApiVO**](OperationResponseGridVOCustomerInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomersInGlobalHaveSiteTemplateModify

> OperationResponseListCustomerInfoOpenApiVO GetCustomersInGlobalHaveSiteTemplateModify(ctx, omadacId).Execute()

Obtain the customer ID with permission to modify site templates



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
	resp, r, err := apiClient.CustomerAPI.GetCustomersInGlobalHaveSiteTemplateModify(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetCustomersInGlobalHaveSiteTemplateModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomersInGlobalHaveSiteTemplateModify`: OperationResponseListCustomerInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetCustomersInGlobalHaveSiteTemplateModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomersInGlobalHaveSiteTemplateModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseListCustomerInfoOpenApiVO**](OperationResponseListCustomerInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCustomer

> OperationResponseWithoutResult ModifyCustomer(ctx, mspId, customerId).CustomerOpenApiVO(customerOpenApiVO).Execute()

Modify an existing customer



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
	customerId := "customerId_example" // string | Customer ID
	customerOpenApiVO := *openapiclient.NewCustomerOpenApiVO() // CustomerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.ModifyCustomer(context.Background(), mspId, customerId).CustomerOpenApiVO(customerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.ModifyCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCustomer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.ModifyCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customerOpenApiVO** | [**CustomerOpenApiVO**](CustomerOpenApiVO.md) |  | 

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


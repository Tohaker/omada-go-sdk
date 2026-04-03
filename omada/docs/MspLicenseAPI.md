# MspLicenseAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAvailableLicenseForMsp**](MspLicenseAPI.md#assignavailablelicenseformsp) | **Get** /openapi/v1/msp/{mspId}/license/assign | Get available assign license info
[**AssignmentLicenseForMsp**](MspLicenseAPI.md#assignmentlicenseformsp) | **Post** /openapi/v1/msp/{mspId}/license/assign | Assign license to a customer
[**GetAutoActive1**](MspLicenseAPI.md#getautoactive1) | **Get** /openapi/v1/msp/{mspId}/license/recycle | Get available recycle license info
[**GetAutoActive2**](MspLicenseAPI.md#getautoactive2) | **Get** /openapi/v1/msp/{mspId}/license/auto-active | Get license auto active
[**GetUsedLicenseForMsp**](MspLicenseAPI.md#getusedlicenseformsp) | **Get** /openapi/v1/msp/{mspId}/license/used | Get used license info list for assign or recycle license
[**ModifyAutoActive1**](MspLicenseAPI.md#modifyautoactive1) | **Post** /openapi/v1/msp/{mspId}/license/auto-active | Modify license auto active
[**RecycleLicenseForMsp**](MspLicenseAPI.md#recyclelicenseformsp) | **Post** /openapi/v1/msp/{mspId}/license/recycle | Recycle license from a customer



## AssignAvailableLicenseForMsp

> LicenseAvailableAssignOpenApiVO AssignAvailableLicenseForMsp(ctx, mspId).Execute()

Get available assign license info



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
	resp, r, err := apiClient.MspLicenseAPI.AssignAvailableLicenseForMsp(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLicenseAPI.AssignAvailableLicenseForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignAvailableLicenseForMsp`: LicenseAvailableAssignOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLicenseAPI.AssignAvailableLicenseForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAvailableLicenseForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseAvailableAssignOpenApiVO**](LicenseAvailableAssignOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignmentLicenseForMsp

> OperationResponseWithoutResult AssignmentLicenseForMsp(ctx, mspId).LicenseAssignmentOpenApiVO(licenseAssignmentOpenApiVO).Execute()

Assign license to a customer



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
	licenseAssignmentOpenApiVO := *openapiclient.NewLicenseAssignmentOpenApiVO("CustomerId_example", []openapiclient.LicenseNumOpenApiVO{*openapiclient.NewLicenseNumOpenApiVO("Type_example")}, *openapiclient.NewLicenseOthersIdsOpenApiVO()) // LicenseAssignmentOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLicenseAPI.AssignmentLicenseForMsp(context.Background(), mspId).LicenseAssignmentOpenApiVO(licenseAssignmentOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLicenseAPI.AssignmentLicenseForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignmentLicenseForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLicenseAPI.AssignmentLicenseForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignmentLicenseForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseAssignmentOpenApiVO** | [**LicenseAssignmentOpenApiVO**](LicenseAssignmentOpenApiVO.md) |  | 

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


## GetAutoActive1

> LicenseAvailableRecycleOpenApiVO GetAutoActive1(ctx, mspId).CustomerId(customerId).Execute()

Get available recycle license info



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
	resp, r, err := apiClient.MspLicenseAPI.GetAutoActive1(context.Background(), mspId).CustomerId(customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLicenseAPI.GetAutoActive1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoActive1`: LicenseAvailableRecycleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLicenseAPI.GetAutoActive1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoActive1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerId** | **string** | Customer ID | 

### Return type

[**LicenseAvailableRecycleOpenApiVO**](LicenseAvailableRecycleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoActive2

> OperationResponseLicenseAutoActiveOpenApiVO GetAutoActive2(ctx, mspId).Execute()

Get license auto active



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
	resp, r, err := apiClient.MspLicenseAPI.GetAutoActive2(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLicenseAPI.GetAutoActive2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoActive2`: OperationResponseLicenseAutoActiveOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLicenseAPI.GetAutoActive2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoActive2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseLicenseAutoActiveOpenApiVO**](OperationResponseLicenseAutoActiveOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsedLicenseForMsp

> GridVOUsedLicenseDetailOpenApiVO GetUsedLicenseForMsp(ctx, mspId).CustomerId(customerId).Category(category).Assign(assign).Page(page).PageSize(pageSize).Execute()

Get used license info list for assign or recycle license



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
	category := "category_example" // string | It should be a value as follows: basic; ap; l2Switch; l3Switch; gateway
	assign := true // bool | Assign license: true; recycle license: false
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLicenseAPI.GetUsedLicenseForMsp(context.Background(), mspId).CustomerId(customerId).Category(category).Assign(assign).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLicenseAPI.GetUsedLicenseForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsedLicenseForMsp`: GridVOUsedLicenseDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspLicenseAPI.GetUsedLicenseForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsedLicenseForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerId** | **string** | Customer ID | 
 **category** | **string** | It should be a value as follows: basic; ap; l2Switch; l3Switch; gateway | 
 **assign** | **bool** | Assign license: true; recycle license: false | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOUsedLicenseDetailOpenApiVO**](GridVOUsedLicenseDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAutoActive1

> OperationResponseWithoutResult ModifyAutoActive1(ctx, mspId).LicenseAutoActiveOpenApiVO(licenseAutoActiveOpenApiVO).Execute()

Modify license auto active



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
	licenseAutoActiveOpenApiVO := *openapiclient.NewLicenseAutoActiveOpenApiVO(false) // LicenseAutoActiveOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLicenseAPI.ModifyAutoActive1(context.Background(), mspId).LicenseAutoActiveOpenApiVO(licenseAutoActiveOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLicenseAPI.ModifyAutoActive1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAutoActive1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLicenseAPI.ModifyAutoActive1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAutoActive1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseAutoActiveOpenApiVO** | [**LicenseAutoActiveOpenApiVO**](LicenseAutoActiveOpenApiVO.md) |  | 

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


## RecycleLicenseForMsp

> OperationResponseWithoutResult RecycleLicenseForMsp(ctx, mspId).LicenseRecycleOpenApiVO(licenseRecycleOpenApiVO).Execute()

Recycle license from a customer



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
	licenseRecycleOpenApiVO := *openapiclient.NewLicenseRecycleOpenApiVO("CustomerId_example", []openapiclient.LicenseNumOpenApiVO{*openapiclient.NewLicenseNumOpenApiVO("Type_example")}, *openapiclient.NewLicenseOthersIdsOpenApiVO()) // LicenseRecycleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspLicenseAPI.RecycleLicenseForMsp(context.Background(), mspId).LicenseRecycleOpenApiVO(licenseRecycleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspLicenseAPI.RecycleLicenseForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecycleLicenseForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspLicenseAPI.RecycleLicenseForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecycleLicenseForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseRecycleOpenApiVO** | [**LicenseRecycleOpenApiVO**](LicenseRecycleOpenApiVO.md) |  | 

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


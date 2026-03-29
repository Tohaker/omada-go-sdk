# \MlagAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMlagGroup**](MlagAPI.md#CreateMlagGroup) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/mlag | Create M-LAG group
[**DeleteMlagGroup**](MlagAPI.md#DeleteMlagGroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/mlag/{mlagId} | Delete M-LAG group
[**GetGridMlagGroup**](MlagAPI.md#GetGridMlagGroup) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mlag | Get M-LAG group list
[**GetMlagCccResult**](MlagAPI.md#GetMlagCccResult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mlag/{mlagId}/mlag-ccc/{type} | Get the configuration consistency check result of the M-LAG Group
[**ListSupportMlagGroupOsws**](MlagAPI.md#ListSupportMlagGroupOsws) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mlag/switches | Get support M-LAG group switch list
[**LocateMlag**](MlagAPI.md#LocateMlag) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/mlag/{mlagId}/locate | Locate M-LAG group
[**ModifyMlagGroup**](MlagAPI.md#ModifyMlagGroup) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/mlag/{mlagId} | Modify M-LAG group
[**RebootMlag**](MlagAPI.md#RebootMlag) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/mlag/{mlagId}/reboot | Reboot M-LAG group



## CreateMlagGroup

> OperationResponseWithoutResult CreateMlagGroup(ctx, omadacId, siteId).MlagConfigOpenApiVO(mlagConfigOpenApiVO).Execute()

Create M-LAG group



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
	mlagConfigOpenApiVO := *openapiclient.NewMlagConfigOpenApiVO("Name_example") // MlagConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlagAPI.CreateMlagGroup(context.Background(), omadacId, siteId).MlagConfigOpenApiVO(mlagConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlagAPI.CreateMlagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMlagGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MlagAPI.CreateMlagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMlagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mlagConfigOpenApiVO** | [**MlagConfigOpenApiVO**](MlagConfigOpenApiVO.md) |  | 

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


## DeleteMlagGroup

> OperationResponseWithoutResult DeleteMlagGroup(ctx, omadacId, siteId, mlagId).Execute()

Delete M-LAG group



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
	mlagId := "mlagId_example" // string | M-LAG group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlagAPI.DeleteMlagGroup(context.Background(), omadacId, siteId, mlagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlagAPI.DeleteMlagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMlagGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MlagAPI.DeleteMlagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mlagId** | **string** | M-LAG group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMlagGroupRequest struct via the builder pattern


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


## GetGridMlagGroup

> OperationResponseGridVOOswMlagVO GetGridMlagGroup(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get M-LAG group list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlagAPI.GetGridMlagGroup(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlagAPI.GetGridMlagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridMlagGroup`: OperationResponseGridVOOswMlagVO
	fmt.Fprintf(os.Stdout, "Response from `MlagAPI.GetGridMlagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridMlagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOswMlagVO**](OperationResponseGridVOOswMlagVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMlagCccResult

> OperationResponseMapStringOswMlagCccVO GetMlagCccResult(ctx, omadacId, siteId, mlagId, type_).Execute()

Get the configuration consistency check result of the M-LAG Group



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
	mlagId := "mlagId_example" // string | M-LAG group ID
	type_ := "type__example" // string | The level of configuration consistency check should be a value as follows: 0:Critical & Significant 1: Critical; 2: Significant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlagAPI.GetMlagCccResult(context.Background(), omadacId, siteId, mlagId, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlagAPI.GetMlagCccResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMlagCccResult`: OperationResponseMapStringOswMlagCccVO
	fmt.Fprintf(os.Stdout, "Response from `MlagAPI.GetMlagCccResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mlagId** | **string** | M-LAG group ID | 
**type_** | **string** | The level of configuration consistency check should be a value as follows: 0:Critical &amp; Significant 1: Critical; 2: Significant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMlagCccResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseMapStringOswMlagCccVO**](OperationResponseMapStringOswMlagCccVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportMlagGroupOsws

> OperationResponseMlagSupportOswVO ListSupportMlagGroupOsws(ctx, omadacId, siteId).Execute()

Get support M-LAG group switch list



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
	resp, r, err := apiClient.MlagAPI.ListSupportMlagGroupOsws(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlagAPI.ListSupportMlagGroupOsws``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportMlagGroupOsws`: OperationResponseMlagSupportOswVO
	fmt.Fprintf(os.Stdout, "Response from `MlagAPI.ListSupportMlagGroupOsws`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportMlagGroupOswsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMlagSupportOswVO**](OperationResponseMlagSupportOswVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocateMlag

> OperationResponseWithoutResult LocateMlag(ctx, omadacId, siteId, mlagId).MlagLocateOpenApiVO(mlagLocateOpenApiVO).Execute()

Locate M-LAG group



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
	mlagId := "mlagId_example" // string | M-LAG group ID
	mlagLocateOpenApiVO := *openapiclient.NewMlagLocateOpenApiVO(false, false) // MlagLocateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlagAPI.LocateMlag(context.Background(), omadacId, siteId, mlagId).MlagLocateOpenApiVO(mlagLocateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlagAPI.LocateMlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocateMlag`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MlagAPI.LocateMlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mlagId** | **string** | M-LAG group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocateMlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mlagLocateOpenApiVO** | [**MlagLocateOpenApiVO**](MlagLocateOpenApiVO.md) |  | 

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


## ModifyMlagGroup

> OperationResponseWithoutResult ModifyMlagGroup(ctx, omadacId, siteId, mlagId).MlagConfigOpenApiVO(mlagConfigOpenApiVO).Execute()

Modify M-LAG group



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
	mlagId := "mlagId_example" // string | M-LAG group ID
	mlagConfigOpenApiVO := *openapiclient.NewMlagConfigOpenApiVO("Name_example") // MlagConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlagAPI.ModifyMlagGroup(context.Background(), omadacId, siteId, mlagId).MlagConfigOpenApiVO(mlagConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlagAPI.ModifyMlagGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMlagGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MlagAPI.ModifyMlagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mlagId** | **string** | M-LAG group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMlagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mlagConfigOpenApiVO** | [**MlagConfigOpenApiVO**](MlagConfigOpenApiVO.md) |  | 

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


## RebootMlag

> OperationResponseWithoutResult RebootMlag(ctx, omadacId, siteId, mlagId).MlagRebootOpenApiVO(mlagRebootOpenApiVO).Execute()

Reboot M-LAG group



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
	mlagId := "mlagId_example" // string | M-LAG group ID
	mlagRebootOpenApiVO := *openapiclient.NewMlagRebootOpenApiVO(false) // MlagRebootOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlagAPI.RebootMlag(context.Background(), omadacId, siteId, mlagId).MlagRebootOpenApiVO(mlagRebootOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlagAPI.RebootMlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootMlag`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MlagAPI.RebootMlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mlagId** | **string** | M-LAG group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootMlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mlagRebootOpenApiVO** | [**MlagRebootOpenApiVO**](MlagRebootOpenApiVO.md) |  | 

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


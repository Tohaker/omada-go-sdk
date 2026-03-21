# \RoutingAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyRouting**](RoutingAPI.md#CreatePolicyRouting) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/routing/policy-routings | Create new policy routing
[**CreateStaticRouting**](RoutingAPI.md#CreateStaticRouting) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/routing/static-routings | Create new static routing
[**DeletePolicyRouting**](RoutingAPI.md#DeletePolicyRouting) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/routing/policy-routings/{policyRoutingId} | Delete policy routing
[**DeleteStaticRouting**](RoutingAPI.md#DeleteStaticRouting) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/routing/static-routings/{staticRoutingId} | Delete static routing
[**GetGridPolicyRouting**](RoutingAPI.md#GetGridPolicyRouting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/routing/policy-routings | Get policy routing list
[**GetGridStaticRouting**](RoutingAPI.md#GetGridStaticRouting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/routing/static-routings | Get static routing list
[**GetStaticRoutingInterfaceList**](RoutingAPI.md#GetStaticRoutingInterfaceList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/routing/static-routings/interfaces | Get static routing interface list
[**ModifyPolicyRouting**](RoutingAPI.md#ModifyPolicyRouting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/routing/policy-routings/{policyRoutingId} | Modify policy routing
[**ModifyPolicyRoutingIndex**](RoutingAPI.md#ModifyPolicyRoutingIndex) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/routing/policy-routings/modifyIndex | Modify policy routing Index
[**ModifyStaticRouting**](RoutingAPI.md#ModifyStaticRouting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/routing/static-routings/{staticRoutingId} | Modify static routing



## CreatePolicyRouting

> OperationResponseWithoutResult CreatePolicyRouting(ctx, omadacId, siteId).PolicyRoutingConfig(policyRoutingConfig).Execute()

Create new policy routing



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
	policyRoutingConfig := *openapiclient.NewPolicyRoutingConfig(false, []string{"DestinationIds_example"}, int32(123), "Name_example", []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // PolicyRoutingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.CreatePolicyRouting(context.Background(), omadacId, siteId).PolicyRoutingConfig(policyRoutingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.CreatePolicyRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicyRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.CreatePolicyRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyRoutingConfig** | [**PolicyRoutingConfig**](PolicyRoutingConfig.md) |  | 

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


## CreateStaticRouting

> OperationResponseWithoutResult CreateStaticRouting(ctx, omadacId, siteId).StaticRoutingConfig(staticRoutingConfig).Execute()

Create new static routing



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
	staticRoutingConfig := *openapiclient.NewStaticRoutingConfig([]string{"Destinations_example"}, int32(123), "Name_example", int32(123), false) // StaticRoutingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.CreateStaticRouting(context.Background(), omadacId, siteId).StaticRoutingConfig(staticRoutingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.CreateStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.CreateStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **staticRoutingConfig** | [**StaticRoutingConfig**](StaticRoutingConfig.md) |  | 

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


## DeletePolicyRouting

> OperationResponseWithoutResult DeletePolicyRouting(ctx, omadacId, siteId, policyRoutingId).Execute()

Delete policy routing



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
	policyRoutingId := "policyRoutingId_example" // string | policyRoutingId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.DeletePolicyRouting(context.Background(), omadacId, siteId, policyRoutingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeletePolicyRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicyRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.DeletePolicyRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**policyRoutingId** | **string** | policyRoutingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRoutingRequest struct via the builder pattern


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


## DeleteStaticRouting

> OperationResponseWithoutResult DeleteStaticRouting(ctx, omadacId, siteId, staticRoutingId).Execute()

Delete static routing



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
	staticRoutingId := "staticRoutingId_example" // string | staticRoutingId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.DeleteStaticRouting(context.Background(), omadacId, siteId, staticRoutingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeleteStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.DeleteStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**staticRoutingId** | **string** | staticRoutingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStaticRoutingRequest struct via the builder pattern


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


## GetGridPolicyRouting

> OperationResponsePolicyRoutingOpenApiGridVOPolicyRoutingInfo GetGridPolicyRouting(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get policy routing list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetGridPolicyRouting(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetGridPolicyRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPolicyRouting`: OperationResponsePolicyRoutingOpenApiGridVOPolicyRoutingInfo
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetGridPolicyRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridPolicyRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponsePolicyRoutingOpenApiGridVOPolicyRoutingInfo**](OperationResponsePolicyRoutingOpenApiGridVOPolicyRoutingInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStaticRouting

> OperationResponseStaticRoutingOpenApiGridVOStaticRoutingInfo GetGridStaticRouting(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get static routing list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetGridStaticRouting(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetGridStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStaticRouting`: OperationResponseStaticRoutingOpenApiGridVOStaticRoutingInfo
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetGridStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseStaticRoutingOpenApiGridVOStaticRoutingInfo**](OperationResponseStaticRoutingOpenApiGridVOStaticRoutingInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaticRoutingInterfaceList

> OperationResponseStaticRoutingInterfaceResult GetStaticRoutingInterfaceList(ctx, omadacId, siteId).Execute()

Get static routing interface list



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
	resp, r, err := apiClient.RoutingAPI.GetStaticRoutingInterfaceList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetStaticRoutingInterfaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStaticRoutingInterfaceList`: OperationResponseStaticRoutingInterfaceResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetStaticRoutingInterfaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticRoutingInterfaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseStaticRoutingInterfaceResult**](OperationResponseStaticRoutingInterfaceResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPolicyRouting

> OperationResponseWithoutResult ModifyPolicyRouting(ctx, omadacId, siteId, policyRoutingId).PolicyRoutingConfig(policyRoutingConfig).Execute()

Modify policy routing



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
	policyRoutingId := "policyRoutingId_example" // string | policyRoutingId
	policyRoutingConfig := *openapiclient.NewPolicyRoutingConfig(false, []string{"DestinationIds_example"}, int32(123), "Name_example", []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // PolicyRoutingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.ModifyPolicyRouting(context.Background(), omadacId, siteId, policyRoutingId).PolicyRoutingConfig(policyRoutingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.ModifyPolicyRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPolicyRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.ModifyPolicyRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**policyRoutingId** | **string** | policyRoutingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPolicyRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **policyRoutingConfig** | [**PolicyRoutingConfig**](PolicyRoutingConfig.md) |  | 

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


## ModifyPolicyRoutingIndex

> OperationResponseWithoutResult ModifyPolicyRoutingIndex(ctx, omadacId, siteId).PolicyRoutingDragSortIndexOpenApiVO(policyRoutingDragSortIndexOpenApiVO).Execute()

Modify policy routing Index



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
	policyRoutingDragSortIndexOpenApiVO := *openapiclient.NewPolicyRoutingDragSortIndexOpenApiVO(map[string]int32{"key": int32(123)}) // PolicyRoutingDragSortIndexOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.ModifyPolicyRoutingIndex(context.Background(), omadacId, siteId).PolicyRoutingDragSortIndexOpenApiVO(policyRoutingDragSortIndexOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.ModifyPolicyRoutingIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPolicyRoutingIndex`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.ModifyPolicyRoutingIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPolicyRoutingIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyRoutingDragSortIndexOpenApiVO** | [**PolicyRoutingDragSortIndexOpenApiVO**](PolicyRoutingDragSortIndexOpenApiVO.md) |  | 

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


## ModifyStaticRouting

> OperationResponseWithoutResult ModifyStaticRouting(ctx, omadacId, siteId, staticRoutingId).StaticRoutingConfig(staticRoutingConfig).Execute()

Modify static routing



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
	staticRoutingId := "staticRoutingId_example" // string | staticRoutingId
	staticRoutingConfig := *openapiclient.NewStaticRoutingConfig([]string{"Destinations_example"}, int32(123), "Name_example", int32(123), false) // StaticRoutingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.ModifyStaticRouting(context.Background(), omadacId, siteId, staticRoutingId).StaticRoutingConfig(staticRoutingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.ModifyStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.ModifyStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**staticRoutingId** | **string** | staticRoutingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **staticRoutingConfig** | [**StaticRoutingConfig**](StaticRoutingConfig.md) |  | 

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


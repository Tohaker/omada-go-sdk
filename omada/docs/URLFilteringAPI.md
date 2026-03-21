# \URLFilteringAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUrlFilteringRule**](URLFilteringAPI.md#CreateUrlFilteringRule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/url-filters | Create URL filtering rule
[**DeleteUrlFilteringRule**](URLFilteringAPI.md#DeleteUrlFilteringRule) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/url-filters/{ruleId} | Delete URL filtering rule
[**GetCategory**](URLFilteringAPI.md#GetCategory) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/url-filters/category | Get Category
[**GetGridEapRule**](URLFilteringAPI.md#GetGridEapRule) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/url-filters/eap | Get URL filtering rule list for eap
[**GetGridGatewayRule**](URLFilteringAPI.md#GetGridGatewayRule) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/url-filters/gateway | Get URL filtering rule list for gateway
[**GetUrlFilterGeneral**](URLFilteringAPI.md#GetUrlFilterGeneral) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/url-filters/globalUrlFilter | Get Content Filter Global
[**ModifyUrlFilterGeneral**](URLFilteringAPI.md#ModifyUrlFilterGeneral) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/url-filters/globalUrlFilter | Edit Content Filter Global
[**ModifyUrlFilteringRule**](URLFilteringAPI.md#ModifyUrlFilteringRule) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/url-filters/{ruleId} | Modify URL filtering rule
[**ModifyUrlFilteringRuleIndex**](URLFilteringAPI.md#ModifyUrlFilteringRuleIndex) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/url-filters | Modify URL filtering rule index



## CreateUrlFilteringRule

> OperationResponseResIdOpenApiVO CreateUrlFilteringRule(ctx, omadacId, siteId).UrlFilteringOpenApiVO(urlFilteringOpenApiVO).Execute()

Create URL filtering rule



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
	urlFilteringOpenApiVO := *openapiclient.NewUrlFilteringOpenApiVO(int32(123), int32(123), "Name_example", int32(123), []string{"SourceIds_example"}, int32(123), false, "Type_example") // UrlFilteringOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringAPI.CreateUrlFilteringRule(context.Background(), omadacId, siteId).UrlFilteringOpenApiVO(urlFilteringOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.CreateUrlFilteringRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUrlFilteringRule`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.CreateUrlFilteringRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUrlFilteringRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **urlFilteringOpenApiVO** | [**UrlFilteringOpenApiVO**](UrlFilteringOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUrlFilteringRule

> OperationResponseWithoutResult DeleteUrlFilteringRule(ctx, omadacId, siteId, ruleId).Execute()

Delete URL filtering rule



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
	ruleId := "ruleId_example" // string | URL filtering rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringAPI.DeleteUrlFilteringRule(context.Background(), omadacId, siteId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.DeleteUrlFilteringRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUrlFilteringRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.DeleteUrlFilteringRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ruleId** | **string** | URL filtering rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUrlFilteringRuleRequest struct via the builder pattern


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


## GetCategory

> OperationResponseUrlCategoryOpenApiVO GetCategory(ctx, omadacId, siteId).Execute()

Get Category



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
	resp, r, err := apiClient.URLFilteringAPI.GetCategory(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.GetCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategory`: OperationResponseUrlCategoryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.GetCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUrlCategoryOpenApiVO**](OperationResponseUrlCategoryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridEapRule

> OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO GetGridEapRule(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get URL filtering rule list for eap



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
	resp, r, err := apiClient.URLFilteringAPI.GetGridEapRule(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.GetGridEapRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridEapRule`: OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.GetGridEapRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridEapRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO**](OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridGatewayRule

> OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO GetGridGatewayRule(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get URL filtering rule list for gateway



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
	resp, r, err := apiClient.URLFilteringAPI.GetGridGatewayRule(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.GetGridGatewayRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridGatewayRule`: OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.GetGridGatewayRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridGatewayRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO**](OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUrlFilterGeneral

> OperationResponseUrlFilterGlobalOpenApiVO GetUrlFilterGeneral(ctx, omadacId, siteId).Execute()

Get Content Filter Global



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
	resp, r, err := apiClient.URLFilteringAPI.GetUrlFilterGeneral(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.GetUrlFilterGeneral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUrlFilterGeneral`: OperationResponseUrlFilterGlobalOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.GetUrlFilterGeneral`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUrlFilterGeneralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUrlFilterGlobalOpenApiVO**](OperationResponseUrlFilterGlobalOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyUrlFilterGeneral

> OperationResponseWithoutResult ModifyUrlFilterGeneral(ctx, omadacId, siteId).UrlFilterGlobalOpenApiVO(urlFilterGlobalOpenApiVO).Execute()

Edit Content Filter Global



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
	urlFilterGlobalOpenApiVO := *openapiclient.NewUrlFilterGlobalOpenApiVO(false, false) // UrlFilterGlobalOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringAPI.ModifyUrlFilterGeneral(context.Background(), omadacId, siteId).UrlFilterGlobalOpenApiVO(urlFilterGlobalOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.ModifyUrlFilterGeneral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUrlFilterGeneral`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.ModifyUrlFilterGeneral`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUrlFilterGeneralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **urlFilterGlobalOpenApiVO** | [**UrlFilterGlobalOpenApiVO**](UrlFilterGlobalOpenApiVO.md) |  | 

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


## ModifyUrlFilteringRule

> OperationResponseWithoutResult ModifyUrlFilteringRule(ctx, omadacId, siteId, ruleId).UrlFilteringOpenApiVO(urlFilteringOpenApiVO).Execute()

Modify URL filtering rule



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
	ruleId := "ruleId_example" // string | URL filtering rule ID
	urlFilteringOpenApiVO := *openapiclient.NewUrlFilteringOpenApiVO(int32(123), int32(123), "Name_example", int32(123), []string{"SourceIds_example"}, int32(123), false, "Type_example") // UrlFilteringOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringAPI.ModifyUrlFilteringRule(context.Background(), omadacId, siteId, ruleId).UrlFilteringOpenApiVO(urlFilteringOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.ModifyUrlFilteringRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUrlFilteringRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.ModifyUrlFilteringRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ruleId** | **string** | URL filtering rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUrlFilteringRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **urlFilteringOpenApiVO** | [**UrlFilteringOpenApiVO**](UrlFilteringOpenApiVO.md) |  | 

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


## ModifyUrlFilteringRuleIndex

> OperationResponseWithoutResult ModifyUrlFilteringRuleIndex(ctx, omadacId, siteId).UrlDragSortIndexOpenapiVO(urlDragSortIndexOpenapiVO).Execute()

Modify URL filtering rule index



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
	urlDragSortIndexOpenapiVO := *openapiclient.NewUrlDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}, "Type_example") // UrlDragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringAPI.ModifyUrlFilteringRuleIndex(context.Background(), omadacId, siteId).UrlDragSortIndexOpenapiVO(urlDragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringAPI.ModifyUrlFilteringRuleIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUrlFilteringRuleIndex`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringAPI.ModifyUrlFilteringRuleIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUrlFilteringRuleIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **urlDragSortIndexOpenapiVO** | [**UrlDragSortIndexOpenapiVO**](UrlDragSortIndexOpenapiVO.md) |  | 

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


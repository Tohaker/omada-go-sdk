# URLFilteringTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUrlFilteringRuleTemplate**](URLFilteringTemplateAPI.md#createurlfilteringruletemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/url-filters | Create URL filtering template rule
[**DeleteUrlFilteringRuleTemplate**](URLFilteringTemplateAPI.md#deleteurlfilteringruletemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/url-filters/{ruleId} | Delete URL filtering template rule
[**GetGridEapRuleTemplate**](URLFilteringTemplateAPI.md#getgrideapruletemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/url-filters/eap | Get URL filtering template rule list for eap
[**GetGridGatewayRuleTemplate**](URLFilteringTemplateAPI.md#getgridgatewayruletemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/url-filters/gateway | Get URL filtering template rule list for gateway
[**ModifyUrlFilteringRuleIndexTemplate**](URLFilteringTemplateAPI.md#modifyurlfilteringruleindextemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/url-filters | Modify URL filtering template rule index
[**ModifyUrlFilteringRuleTemplate**](URLFilteringTemplateAPI.md#modifyurlfilteringruletemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/url-filters/{ruleId} | Modify URL filtering template rule



## CreateUrlFilteringRuleTemplate

> OperationResponseResIdOpenApiVO CreateUrlFilteringRuleTemplate(ctx, omadacId, siteTemplateId).UrlFilteringOpenApiVO(urlFilteringOpenApiVO).Execute()

Create URL filtering template rule



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	urlFilteringOpenApiVO := *openapiclient.NewUrlFilteringOpenApiVO(int32(123), int32(123), "Name_example", int32(123), []string{"SourceIds_example"}, int32(123), false, "Type_example") // UrlFilteringOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringTemplateAPI.CreateUrlFilteringRuleTemplate(context.Background(), omadacId, siteTemplateId).UrlFilteringOpenApiVO(urlFilteringOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringTemplateAPI.CreateUrlFilteringRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUrlFilteringRuleTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringTemplateAPI.CreateUrlFilteringRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUrlFilteringRuleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **urlFilteringOpenApiVO** | [**UrlFilteringOpenApiVO**](UrlFilteringOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUrlFilteringRuleTemplate

> OperationResponseWithoutResult DeleteUrlFilteringRuleTemplate(ctx, omadacId, siteTemplateId, ruleId).Execute()

Delete URL filtering template rule



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ruleId := "ruleId_example" // string | URL filtering rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringTemplateAPI.DeleteUrlFilteringRuleTemplate(context.Background(), omadacId, siteTemplateId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringTemplateAPI.DeleteUrlFilteringRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUrlFilteringRuleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringTemplateAPI.DeleteUrlFilteringRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ruleId** | **string** | URL filtering rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUrlFilteringRuleTemplateRequest struct via the builder pattern


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


## GetGridEapRuleTemplate

> OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO GetGridEapRuleTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get URL filtering template rule list for eap



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringTemplateAPI.GetGridEapRuleTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringTemplateAPI.GetGridEapRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridEapRuleTemplate`: OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringTemplateAPI.GetGridEapRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridEapRuleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO**](OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridGatewayRuleTemplate

> OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO GetGridGatewayRuleTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get URL filtering template rule list for gateway



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringTemplateAPI.GetGridGatewayRuleTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringTemplateAPI.GetGridGatewayRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridGatewayRuleTemplate`: OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringTemplateAPI.GetGridGatewayRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridGatewayRuleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO**](OperationResponseUrlFilteringOpenApiGridVOQueryUrlFilteringOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyUrlFilteringRuleIndexTemplate

> OperationResponseWithoutResult ModifyUrlFilteringRuleIndexTemplate(ctx, omadacId, siteTemplateId).UrlDragSortIndexOpenapiVO(urlDragSortIndexOpenapiVO).Execute()

Modify URL filtering template rule index



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	urlDragSortIndexOpenapiVO := *openapiclient.NewUrlDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}, "Type_example") // UrlDragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringTemplateAPI.ModifyUrlFilteringRuleIndexTemplate(context.Background(), omadacId, siteTemplateId).UrlDragSortIndexOpenapiVO(urlDragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringTemplateAPI.ModifyUrlFilteringRuleIndexTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUrlFilteringRuleIndexTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringTemplateAPI.ModifyUrlFilteringRuleIndexTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUrlFilteringRuleIndexTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **urlDragSortIndexOpenapiVO** | [**UrlDragSortIndexOpenapiVO**](UrlDragSortIndexOpenapiVO.md) |  | 

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


## ModifyUrlFilteringRuleTemplate

> OperationResponseWithoutResult ModifyUrlFilteringRuleTemplate(ctx, omadacId, siteTemplateId, ruleId).UrlFilteringOpenApiVO(urlFilteringOpenApiVO).Execute()

Modify URL filtering template rule



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ruleId := "ruleId_example" // string | URL filtering rule ID
	urlFilteringOpenApiVO := *openapiclient.NewUrlFilteringOpenApiVO(int32(123), int32(123), "Name_example", int32(123), []string{"SourceIds_example"}, int32(123), false, "Type_example") // UrlFilteringOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.URLFilteringTemplateAPI.ModifyUrlFilteringRuleTemplate(context.Background(), omadacId, siteTemplateId, ruleId).UrlFilteringOpenApiVO(urlFilteringOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `URLFilteringTemplateAPI.ModifyUrlFilteringRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUrlFilteringRuleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `URLFilteringTemplateAPI.ModifyUrlFilteringRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ruleId** | **string** | URL filtering rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUrlFilteringRuleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **urlFilteringOpenApiVO** | [**UrlFilteringOpenApiVO**](UrlFilteringOpenApiVO.md) |  | 

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


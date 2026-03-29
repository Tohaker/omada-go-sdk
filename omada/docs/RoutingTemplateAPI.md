# \RoutingTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplatePolicyRouting**](RoutingTemplateAPI.md#CreateTemplatePolicyRouting) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/policy-routings | Create site template&#39;s policy routing
[**CreateTemplateStaticRouting**](RoutingTemplateAPI.md#CreateTemplateStaticRouting) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/static-routings | Create site template&#39;s static routing
[**DeleteTemplatePolicyRouting**](RoutingTemplateAPI.md#DeleteTemplatePolicyRouting) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/policy-routings/{policyRoutingId} | Delete site template&#39;s policy routing
[**DeleteTemplateStaticRouting**](RoutingTemplateAPI.md#DeleteTemplateStaticRouting) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/static-routings/{staticRoutingId} | Delete site template&#39;s static routing
[**GetTemplateGridPolicyRouting**](RoutingTemplateAPI.md#GetTemplateGridPolicyRouting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/policy-routings | Get site template&#39;s policy routing list
[**GetTemplateGridStaticRouting**](RoutingTemplateAPI.md#GetTemplateGridStaticRouting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/static-routings | Get site template&#39;s static routing list
[**ModifyTemplatePolicyRouting**](RoutingTemplateAPI.md#ModifyTemplatePolicyRouting) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/policy-routings/{policyRoutingId} | Modify site template&#39;s policy routing
[**ModifyTemplatePolicyRoutingIndex**](RoutingTemplateAPI.md#ModifyTemplatePolicyRoutingIndex) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/policy-routings/modifyIndex | Modify site template&#39;s policy routing Index
[**ModifyTemplateStaticRouting**](RoutingTemplateAPI.md#ModifyTemplateStaticRouting) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/routing/static-routings/{staticRoutingId} | Modify site template&#39;s static routing



## CreateTemplatePolicyRouting

> OperationResponseWithoutResult CreateTemplatePolicyRouting(ctx, omadacId, siteTemplateId).PolicyRoutingConfig(policyRoutingConfig).Execute()

Create site template's policy routing



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	policyRoutingConfig := *openapiclient.NewPolicyRoutingConfig(false, []string{"DestinationIds_example"}, int32(123), "Name_example", []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // PolicyRoutingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.CreateTemplatePolicyRouting(context.Background(), omadacId, siteTemplateId).PolicyRoutingConfig(policyRoutingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.CreateTemplatePolicyRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplatePolicyRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.CreateTemplatePolicyRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplatePolicyRoutingRequest struct via the builder pattern


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


## CreateTemplateStaticRouting

> OperationResponseWithoutResult CreateTemplateStaticRouting(ctx, omadacId, siteTemplateId).StaticRoutingConfigTemplate(staticRoutingConfigTemplate).Execute()

Create site template's static routing



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	staticRoutingConfigTemplate := *openapiclient.NewStaticRoutingConfigTemplate([]string{"Destinations_example"}, int32(123), "Name_example", int32(123), false) // StaticRoutingConfigTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.CreateTemplateStaticRouting(context.Background(), omadacId, siteTemplateId).StaticRoutingConfigTemplate(staticRoutingConfigTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.CreateTemplateStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.CreateTemplateStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **staticRoutingConfigTemplate** | [**StaticRoutingConfigTemplate**](StaticRoutingConfigTemplate.md) |  | 

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


## DeleteTemplatePolicyRouting

> OperationResponseWithoutResult DeleteTemplatePolicyRouting(ctx, omadacId, siteTemplateId, policyRoutingId).Execute()

Delete site template's policy routing



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	policyRoutingId := "policyRoutingId_example" // string | Policy routing ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.DeleteTemplatePolicyRouting(context.Background(), omadacId, siteTemplateId, policyRoutingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.DeleteTemplatePolicyRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplatePolicyRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.DeleteTemplatePolicyRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**policyRoutingId** | **string** | Policy routing ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplatePolicyRoutingRequest struct via the builder pattern


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


## DeleteTemplateStaticRouting

> OperationResponseWithoutResult DeleteTemplateStaticRouting(ctx, omadacId, siteTemplateId, staticRoutingId).Execute()

Delete site template's static routing



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	staticRoutingId := "staticRoutingId_example" // string | Static routing ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.DeleteTemplateStaticRouting(context.Background(), omadacId, siteTemplateId, staticRoutingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.DeleteTemplateStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplateStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.DeleteTemplateStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**staticRoutingId** | **string** | Static routing ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateStaticRoutingRequest struct via the builder pattern


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


## GetTemplateGridPolicyRouting

> OperationResponsePolicyRoutingOpenApiGridVOPolicyRoutingInfo GetTemplateGridPolicyRouting(ctx, omadacId, siteTemplateId).Execute()

Get site template's policy routing list



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.GetTemplateGridPolicyRouting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.GetTemplateGridPolicyRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateGridPolicyRouting`: OperationResponsePolicyRoutingOpenApiGridVOPolicyRoutingInfo
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.GetTemplateGridPolicyRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateGridPolicyRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetTemplateGridStaticRouting

> OperationResponseStaticRoutingOpenApiGridVOStaticRoutingInfoTemplate GetTemplateGridStaticRouting(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get site template's static routing list



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.GetTemplateGridStaticRouting(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.GetTemplateGridStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateGridStaticRouting`: OperationResponseStaticRoutingOpenApiGridVOStaticRoutingInfoTemplate
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.GetTemplateGridStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateGridStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseStaticRoutingOpenApiGridVOStaticRoutingInfoTemplate**](OperationResponseStaticRoutingOpenApiGridVOStaticRoutingInfoTemplate.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTemplatePolicyRouting

> OperationResponseWithoutResult ModifyTemplatePolicyRouting(ctx, omadacId, siteTemplateId, policyRoutingId).PolicyRoutingConfig(policyRoutingConfig).Execute()

Modify site template's policy routing



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	policyRoutingId := "policyRoutingId_example" // string | Policy routing ID
	policyRoutingConfig := *openapiclient.NewPolicyRoutingConfig(false, []string{"DestinationIds_example"}, int32(123), "Name_example", []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // PolicyRoutingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.ModifyTemplatePolicyRouting(context.Background(), omadacId, siteTemplateId, policyRoutingId).PolicyRoutingConfig(policyRoutingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.ModifyTemplatePolicyRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplatePolicyRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.ModifyTemplatePolicyRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**policyRoutingId** | **string** | Policy routing ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplatePolicyRoutingRequest struct via the builder pattern


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


## ModifyTemplatePolicyRoutingIndex

> OperationResponseWithoutResult ModifyTemplatePolicyRoutingIndex(ctx, omadacId, siteTemplateId).PolicyRoutingDragSortIndexOpenApiVO(policyRoutingDragSortIndexOpenApiVO).Execute()

Modify site template's policy routing Index



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	policyRoutingDragSortIndexOpenApiVO := *openapiclient.NewPolicyRoutingDragSortIndexOpenApiVO(map[string]int32{"key": int32(123)}) // PolicyRoutingDragSortIndexOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.ModifyTemplatePolicyRoutingIndex(context.Background(), omadacId, siteTemplateId).PolicyRoutingDragSortIndexOpenApiVO(policyRoutingDragSortIndexOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.ModifyTemplatePolicyRoutingIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplatePolicyRoutingIndex`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.ModifyTemplatePolicyRoutingIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplatePolicyRoutingIndexRequest struct via the builder pattern


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


## ModifyTemplateStaticRouting

> OperationResponseWithoutResult ModifyTemplateStaticRouting(ctx, omadacId, siteTemplateId, staticRoutingId).StaticRoutingConfigTemplate(staticRoutingConfigTemplate).Execute()

Modify site template's static routing



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	staticRoutingId := "staticRoutingId_example" // string | 
	staticRoutingConfigTemplate := *openapiclient.NewStaticRoutingConfigTemplate([]string{"Destinations_example"}, int32(123), "Name_example", int32(123), false) // StaticRoutingConfigTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingTemplateAPI.ModifyTemplateStaticRouting(context.Background(), omadacId, siteTemplateId, staticRoutingId).StaticRoutingConfigTemplate(staticRoutingConfigTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingTemplateAPI.ModifyTemplateStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RoutingTemplateAPI.ModifyTemplateStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**staticRoutingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **staticRoutingConfigTemplate** | [**StaticRoutingConfigTemplate**](StaticRoutingConfigTemplate.md) |  | 

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


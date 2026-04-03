# SessionLimitTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplateSessionLimitRule**](SessionLimitTemplateAPI.md#createtemplatesessionlimitrule) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/session-limit/rules | Create site template&#39;s session limit rule
[**DeleteTemplateSessionLimitRule**](SessionLimitTemplateAPI.md#deletetemplatesessionlimitrule) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/session-limit/rules/{ruleId} | Delete site template&#39;s session limit rule
[**GetTemplateGridSessionLimitRule**](SessionLimitTemplateAPI.md#gettemplategridsessionlimitrule) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/session-limit/rules | Get site template&#39;s session limit rule list
[**GetTemplateSessionLimit**](SessionLimitTemplateAPI.md#gettemplatesessionlimit) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/session-limit | Get site template&#39;s session limit
[**ModifyTemplateSessionLimit**](SessionLimitTemplateAPI.md#modifytemplatesessionlimit) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/session-limit | Modify site template&#39;s session limit
[**ModifyTemplateSessionLimitRule**](SessionLimitTemplateAPI.md#modifytemplatesessionlimitrule) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/session-limit/rules/{ruleId} | Modify site template&#39;s session limit rule
[**ModifyTemplateSessionLimitRuleIndex**](SessionLimitTemplateAPI.md#modifytemplatesessionlimitruleindex) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/session-limit/rules | Modify site template&#39;s session limit rule index



## CreateTemplateSessionLimitRule

> OperationResponseWithoutResult CreateTemplateSessionLimitRule(ctx, omadacId, siteTemplateId).SessionLimitRuleTemplateOpenApiVO(sessionLimitRuleTemplateOpenApiVO).Execute()

Create site template's session limit rule



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
	sessionLimitRuleTemplateOpenApiVO := *openapiclient.NewSessionLimitRuleTemplateOpenApiVO(int32(123), "Name_example", int32(123), false) // SessionLimitRuleTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitTemplateAPI.CreateTemplateSessionLimitRule(context.Background(), omadacId, siteTemplateId).SessionLimitRuleTemplateOpenApiVO(sessionLimitRuleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitTemplateAPI.CreateTemplateSessionLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateSessionLimitRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitTemplateAPI.CreateTemplateSessionLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateSessionLimitRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sessionLimitRuleTemplateOpenApiVO** | [**SessionLimitRuleTemplateOpenApiVO**](SessionLimitRuleTemplateOpenApiVO.md) |  | 

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


## DeleteTemplateSessionLimitRule

> OperationResponseWithoutResult DeleteTemplateSessionLimitRule(ctx, omadacId, siteTemplateId, ruleId).Execute()

Delete site template's session limit rule



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
	ruleId := "ruleId_example" // string | Session limit rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitTemplateAPI.DeleteTemplateSessionLimitRule(context.Background(), omadacId, siteTemplateId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitTemplateAPI.DeleteTemplateSessionLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplateSessionLimitRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitTemplateAPI.DeleteTemplateSessionLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ruleId** | **string** | Session limit rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateSessionLimitRuleRequest struct via the builder pattern


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


## GetTemplateGridSessionLimitRule

> OperationResponseSessionLimitRuleGridOpenApiVOQuerySessionLimitRuleOpenApiVO GetTemplateGridSessionLimitRule(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get site template's session limit rule list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitTemplateAPI.GetTemplateGridSessionLimitRule(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitTemplateAPI.GetTemplateGridSessionLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateGridSessionLimitRule`: OperationResponseSessionLimitRuleGridOpenApiVOQuerySessionLimitRuleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitTemplateAPI.GetTemplateGridSessionLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateGridSessionLimitRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseSessionLimitRuleGridOpenApiVOQuerySessionLimitRuleOpenApiVO**](OperationResponseSessionLimitRuleGridOpenApiVOQuerySessionLimitRuleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateSessionLimit

> OperationResponseSessionLimitEntity GetTemplateSessionLimit(ctx, omadacId, siteTemplateId).Execute()

Get site template's session limit



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
	resp, r, err := apiClient.SessionLimitTemplateAPI.GetTemplateSessionLimit(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitTemplateAPI.GetTemplateSessionLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateSessionLimit`: OperationResponseSessionLimitEntity
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitTemplateAPI.GetTemplateSessionLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateSessionLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSessionLimitEntity**](OperationResponseSessionLimitEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTemplateSessionLimit

> OperationResponseWithoutResult ModifyTemplateSessionLimit(ctx, omadacId, siteTemplateId).SessionLimitEntity(sessionLimitEntity).Execute()

Modify site template's session limit



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
	sessionLimitEntity := *openapiclient.NewSessionLimitEntity(false) // SessionLimitEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitTemplateAPI.ModifyTemplateSessionLimit(context.Background(), omadacId, siteTemplateId).SessionLimitEntity(sessionLimitEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitTemplateAPI.ModifyTemplateSessionLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateSessionLimit`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitTemplateAPI.ModifyTemplateSessionLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateSessionLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sessionLimitEntity** | [**SessionLimitEntity**](SessionLimitEntity.md) |  | 

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


## ModifyTemplateSessionLimitRule

> OperationResponseWithoutResult ModifyTemplateSessionLimitRule(ctx, omadacId, siteTemplateId, ruleId).SessionLimitRuleTemplateOpenApiVO(sessionLimitRuleTemplateOpenApiVO).Execute()

Modify site template's session limit rule



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
	ruleId := "ruleId_example" // string | Session limit rule ID
	sessionLimitRuleTemplateOpenApiVO := *openapiclient.NewSessionLimitRuleTemplateOpenApiVO(int32(123), "Name_example", int32(123), false) // SessionLimitRuleTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitTemplateAPI.ModifyTemplateSessionLimitRule(context.Background(), omadacId, siteTemplateId, ruleId).SessionLimitRuleTemplateOpenApiVO(sessionLimitRuleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitTemplateAPI.ModifyTemplateSessionLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateSessionLimitRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitTemplateAPI.ModifyTemplateSessionLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ruleId** | **string** | Session limit rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateSessionLimitRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sessionLimitRuleTemplateOpenApiVO** | [**SessionLimitRuleTemplateOpenApiVO**](SessionLimitRuleTemplateOpenApiVO.md) |  | 

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


## ModifyTemplateSessionLimitRuleIndex

> OperationResponseWithoutResult ModifyTemplateSessionLimitRuleIndex(ctx, omadacId, siteTemplateId).TransmissionDragSortIndexOpenapiVO(transmissionDragSortIndexOpenapiVO).Execute()

Modify site template's session limit rule index



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
	transmissionDragSortIndexOpenapiVO := *openapiclient.NewTransmissionDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}) // TransmissionDragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitTemplateAPI.ModifyTemplateSessionLimitRuleIndex(context.Background(), omadacId, siteTemplateId).TransmissionDragSortIndexOpenapiVO(transmissionDragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitTemplateAPI.ModifyTemplateSessionLimitRuleIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateSessionLimitRuleIndex`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitTemplateAPI.ModifyTemplateSessionLimitRuleIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateSessionLimitRuleIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transmissionDragSortIndexOpenapiVO** | [**TransmissionDragSortIndexOpenapiVO**](TransmissionDragSortIndexOpenapiVO.md) |  | 

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


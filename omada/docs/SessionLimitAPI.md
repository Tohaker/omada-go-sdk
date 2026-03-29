# \SessionLimitAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionLimitRule**](SessionLimitAPI.md#CreateSessionLimitRule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/session-limit/rules | Create session limit rule
[**DeleteSessionLimitRule**](SessionLimitAPI.md#DeleteSessionLimitRule) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/session-limit/rules/{ruleId} | Delete session limit rule
[**GetGridSessionLimitRule**](SessionLimitAPI.md#GetGridSessionLimitRule) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/session-limit/rules | Get session limit rule list
[**GetSessionLimit**](SessionLimitAPI.md#GetSessionLimit) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/session-limit | Get session limit
[**ModifySessionLimit**](SessionLimitAPI.md#ModifySessionLimit) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/session-limit | Modify session limit
[**ModifySessionLimitRule**](SessionLimitAPI.md#ModifySessionLimitRule) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/session-limit/rules/{ruleId} | Modify session limit rule
[**ModifySessionLimitRuleIndex**](SessionLimitAPI.md#ModifySessionLimitRuleIndex) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/session-limit/rules | Modify session limit rule index



## CreateSessionLimitRule

> OperationResponseWithoutResult CreateSessionLimitRule(ctx, omadacId, siteId).SessionLimitRuleOpenApiVO(sessionLimitRuleOpenApiVO).Execute()

Create session limit rule



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
	sessionLimitRuleOpenApiVO := *openapiclient.NewSessionLimitRuleOpenApiVO(int32(123), "Name_example", int32(123), false) // SessionLimitRuleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitAPI.CreateSessionLimitRule(context.Background(), omadacId, siteId).SessionLimitRuleOpenApiVO(sessionLimitRuleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitAPI.CreateSessionLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSessionLimitRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitAPI.CreateSessionLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionLimitRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sessionLimitRuleOpenApiVO** | [**SessionLimitRuleOpenApiVO**](SessionLimitRuleOpenApiVO.md) |  | 

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


## DeleteSessionLimitRule

> OperationResponseWithoutResult DeleteSessionLimitRule(ctx, omadacId, siteId, ruleId).Execute()

Delete session limit rule



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
	ruleId := "ruleId_example" // string | Session limit rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitAPI.DeleteSessionLimitRule(context.Background(), omadacId, siteId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitAPI.DeleteSessionLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSessionLimitRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitAPI.DeleteSessionLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ruleId** | **string** | Session limit rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionLimitRuleRequest struct via the builder pattern


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


## GetGridSessionLimitRule

> OperationResponseSessionLimitRuleGridOpenApiVOQuerySessionLimitRuleOpenApiVO GetGridSessionLimitRule(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get session limit rule list



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
	resp, r, err := apiClient.SessionLimitAPI.GetGridSessionLimitRule(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitAPI.GetGridSessionLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSessionLimitRule`: OperationResponseSessionLimitRuleGridOpenApiVOQuerySessionLimitRuleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitAPI.GetGridSessionLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSessionLimitRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseSessionLimitRuleGridOpenApiVOQuerySessionLimitRuleOpenApiVO**](OperationResponseSessionLimitRuleGridOpenApiVOQuerySessionLimitRuleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionLimit

> OperationResponseSessionLimitEntity GetSessionLimit(ctx, omadacId, siteId).Execute()

Get session limit



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
	resp, r, err := apiClient.SessionLimitAPI.GetSessionLimit(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitAPI.GetSessionLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionLimit`: OperationResponseSessionLimitEntity
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitAPI.GetSessionLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSessionLimitEntity**](OperationResponseSessionLimitEntity.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySessionLimit

> OperationResponseWithoutResult ModifySessionLimit(ctx, omadacId, siteId).SessionLimitEntity(sessionLimitEntity).Execute()

Modify session limit



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
	sessionLimitEntity := *openapiclient.NewSessionLimitEntity(false) // SessionLimitEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitAPI.ModifySessionLimit(context.Background(), omadacId, siteId).SessionLimitEntity(sessionLimitEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitAPI.ModifySessionLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySessionLimit`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitAPI.ModifySessionLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySessionLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sessionLimitEntity** | [**SessionLimitEntity**](SessionLimitEntity.md) |  | 

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


## ModifySessionLimitRule

> OperationResponseWithoutResult ModifySessionLimitRule(ctx, omadacId, siteId, ruleId).SessionLimitRuleOpenApiVO(sessionLimitRuleOpenApiVO).Execute()

Modify session limit rule



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
	ruleId := "ruleId_example" // string | Session limit rule ID
	sessionLimitRuleOpenApiVO := *openapiclient.NewSessionLimitRuleOpenApiVO(int32(123), "Name_example", int32(123), false) // SessionLimitRuleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitAPI.ModifySessionLimitRule(context.Background(), omadacId, siteId, ruleId).SessionLimitRuleOpenApiVO(sessionLimitRuleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitAPI.ModifySessionLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySessionLimitRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitAPI.ModifySessionLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ruleId** | **string** | Session limit rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySessionLimitRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sessionLimitRuleOpenApiVO** | [**SessionLimitRuleOpenApiVO**](SessionLimitRuleOpenApiVO.md) |  | 

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


## ModifySessionLimitRuleIndex

> OperationResponseWithoutResult ModifySessionLimitRuleIndex(ctx, omadacId, siteId).TransmissionDragSortIndexOpenapiVO(transmissionDragSortIndexOpenapiVO).Execute()

Modify session limit rule index



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
	transmissionDragSortIndexOpenapiVO := *openapiclient.NewTransmissionDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}) // TransmissionDragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionLimitAPI.ModifySessionLimitRuleIndex(context.Background(), omadacId, siteId).TransmissionDragSortIndexOpenapiVO(transmissionDragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionLimitAPI.ModifySessionLimitRuleIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySessionLimitRuleIndex`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SessionLimitAPI.ModifySessionLimitRuleIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySessionLimitRuleIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transmissionDragSortIndexOpenapiVO** | [**TransmissionDragSortIndexOpenapiVO**](TransmissionDragSortIndexOpenapiVO.md) |  | 

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


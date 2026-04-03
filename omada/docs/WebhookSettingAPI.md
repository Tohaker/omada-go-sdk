# WebhookSettingAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookForGlobal**](WebhookSettingAPI.md#createwebhookforglobal) | **Post** /openapi/v1/{omadacId}/webhook/settings | Create global webhook setting
[**DeleteWebhookForGlobal**](WebhookSettingAPI.md#deletewebhookforglobal) | **Delete** /openapi/v1/{omadacId}/webhook/settings/{webhookId} | Delete global webhook setting
[**GetWebhookForGlobal**](WebhookSettingAPI.md#getwebhookforglobal) | **Get** /openapi/v1/{omadacId}/webhook/settings | Get global webhook setting list.
[**GetWebhookLogsForGlobal**](WebhookSettingAPI.md#getwebhooklogsforglobal) | **Get** /openapi/v1/{omadacId}/webhook/settings/dispatch-logs | Get global webhook dispatch log list
[**ModifyWebhookForGlobal**](WebhookSettingAPI.md#modifywebhookforglobal) | **Patch** /openapi/v1/{omadacId}/webhook/settings/{webhookId} | Modify global webhook setting
[**TestWebhookForGlobal**](WebhookSettingAPI.md#testwebhookforglobal) | **Post** /openapi/v1/{omadacId}/webhook/settings/{webhookId} | Test global webhook setting



## CreateWebhookForGlobal

> OperationResponseWithoutResult CreateWebhookForGlobal(ctx, omadacId).OpenApiWebhookSettingAddVO(openApiWebhookSettingAddVO).Execute()

Create global webhook setting



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
	openApiWebhookSettingAddVO := *openapiclient.NewOpenApiWebhookSettingAddVO("Name_example", int32(123), int32(123), []string{" [http(s)://webhook.site/4a566f9e-0b77-42e2-9a34-a78]"}) // OpenApiWebhookSettingAddVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSettingAPI.CreateWebhookForGlobal(context.Background(), omadacId).OpenApiWebhookSettingAddVO(openApiWebhookSettingAddVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSettingAPI.CreateWebhookForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookForGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WebhookSettingAPI.CreateWebhookForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openApiWebhookSettingAddVO** | [**OpenApiWebhookSettingAddVO**](OpenApiWebhookSettingAddVO.md) |  | 

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


## DeleteWebhookForGlobal

> OperationResponseWithoutResult DeleteWebhookForGlobal(ctx, omadacId, webhookId).Execute()

Delete global webhook setting



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
	webhookId := "webhookId_example" // string | webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSettingAPI.DeleteWebhookForGlobal(context.Background(), omadacId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSettingAPI.DeleteWebhookForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhookForGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WebhookSettingAPI.DeleteWebhookForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**webhookId** | **string** | webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookForGlobalRequest struct via the builder pattern


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


## GetWebhookForGlobal

> OperationResponseListWebhookSetting GetWebhookForGlobal(ctx, omadacId).Execute()

Get global webhook setting list.



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
	resp, r, err := apiClient.WebhookSettingAPI.GetWebhookForGlobal(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSettingAPI.GetWebhookForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookForGlobal`: OperationResponseListWebhookSetting
	fmt.Fprintf(os.Stdout, "Response from `WebhookSettingAPI.GetWebhookForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseListWebhookSetting**](OperationResponseListWebhookSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookLogsForGlobal

> OperationResponseGridVOOpenApiWebhookDispatchLogVO GetWebhookLogsForGlobal(ctx, omadacId).Page(page).PageSize(pageSize).FiltersWebhookId(filtersWebhookId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()

Get global webhook dispatch log list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersWebhookId := "filtersWebhookId_example" // string | Filter query parameters, support field d66aed17ea7238154ba610710d9a2890
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field 1679297710438
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field 1681889710438

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSettingAPI.GetWebhookLogsForGlobal(context.Background(), omadacId).Page(page).PageSize(pageSize).FiltersWebhookId(filtersWebhookId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSettingAPI.GetWebhookLogsForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookLogsForGlobal`: OperationResponseGridVOOpenApiWebhookDispatchLogVO
	fmt.Fprintf(os.Stdout, "Response from `WebhookSettingAPI.GetWebhookLogsForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookLogsForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersWebhookId** | **string** | Filter query parameters, support field d66aed17ea7238154ba610710d9a2890 | 
 **filtersTimeStart** | **int64** | Filter query parameters, support field 1679297710438 | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field 1681889710438 | 

### Return type

[**OperationResponseGridVOOpenApiWebhookDispatchLogVO**](OperationResponseGridVOOpenApiWebhookDispatchLogVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWebhookForGlobal

> OperationResponseWithoutResult ModifyWebhookForGlobal(ctx, omadacId, webhookId).OpenApiWebhookSettingEditVO(openApiWebhookSettingEditVO).Execute()

Modify global webhook setting



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
	webhookId := "webhookId_example" // string | webhook ID
	openApiWebhookSettingEditVO := *openapiclient.NewOpenApiWebhookSettingEditVO(int32(123), int32(123), []string{"[http(s)://webhook.site/4a566f9e-0b77-42e2-9a34-a78]"}) // OpenApiWebhookSettingEditVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSettingAPI.ModifyWebhookForGlobal(context.Background(), omadacId, webhookId).OpenApiWebhookSettingEditVO(openApiWebhookSettingEditVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSettingAPI.ModifyWebhookForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWebhookForGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WebhookSettingAPI.ModifyWebhookForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**webhookId** | **string** | webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWebhookForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **openApiWebhookSettingEditVO** | [**OpenApiWebhookSettingEditVO**](OpenApiWebhookSettingEditVO.md) |  | 

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


## TestWebhookForGlobal

> OperationResponseWithoutResult TestWebhookForGlobal(ctx, omadacId, webhookId).OpenApiWebhookSettingTestVO(openApiWebhookSettingTestVO).Execute()

Test global webhook setting



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
	webhookId := "webhookId_example" // string | webhook ID
	openApiWebhookSettingTestVO := *openapiclient.NewOpenApiWebhookSettingTestVO("Name_example", int32(123), []string{"[http(s)://webhook.site/4a566f9e-0b77-42e2-9a34-a78]"}) // OpenApiWebhookSettingTestVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSettingAPI.TestWebhookForGlobal(context.Background(), omadacId, webhookId).OpenApiWebhookSettingTestVO(openApiWebhookSettingTestVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSettingAPI.TestWebhookForGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestWebhookForGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WebhookSettingAPI.TestWebhookForGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**webhookId** | **string** | webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhookForGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **openApiWebhookSettingTestVO** | [**OpenApiWebhookSettingTestVO**](OpenApiWebhookSettingTestVO.md) |  | 

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


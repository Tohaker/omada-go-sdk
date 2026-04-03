# MspWebhookSettingAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookForMsp**](MspWebhookSettingAPI.md#createwebhookformsp) | **Post** /openapi/v1/msp/{mspId}/webhook/settings | Create MSP webhook setting
[**DeleteWebhookForMsp**](MspWebhookSettingAPI.md#deletewebhookformsp) | **Delete** /openapi/v1/msp/{mspId}/webhook/settings/{webhookId} | Delete MSP webhook setting
[**GetWebhookForMsp**](MspWebhookSettingAPI.md#getwebhookformsp) | **Get** /openapi/v1/msp/{mspId}/webhook/settings | Get MSP webhook setting list
[**GetWebhookLogsForMsp**](MspWebhookSettingAPI.md#getwebhooklogsformsp) | **Get** /openapi/v1/msp/{mspId}/webhook/settings/dispatch-logs | Get MSP webhook dispatch log list
[**ModifyWebhookForMsp**](MspWebhookSettingAPI.md#modifywebhookformsp) | **Patch** /openapi/v1/msp/{mspId}/webhook/settings/{webhookId} | Modify MSP webhook setting
[**TestWebhookFormsp**](MspWebhookSettingAPI.md#testwebhookformsp) | **Post** /openapi/v1/msp/{mspId}/webhook/settings/{webhookId} | Test MSP webhook setting



## CreateWebhookForMsp

> OperationResponseWithoutResult CreateWebhookForMsp(ctx, mspId).OpenApiWebhookSettingAddVO(openApiWebhookSettingAddVO).Execute()

Create MSP webhook setting



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
	openApiWebhookSettingAddVO := *openapiclient.NewOpenApiWebhookSettingAddVO("Name_example", int32(123), int32(123), []string{" [http(s)://webhook.site/4a566f9e-0b77-42e2-9a34-a78]"}) // OpenApiWebhookSettingAddVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspWebhookSettingAPI.CreateWebhookForMsp(context.Background(), mspId).OpenApiWebhookSettingAddVO(openApiWebhookSettingAddVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspWebhookSettingAPI.CreateWebhookForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspWebhookSettingAPI.CreateWebhookForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookForMspRequest struct via the builder pattern


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


## DeleteWebhookForMsp

> OperationResponseWithoutResult DeleteWebhookForMsp(ctx, mspId, webhookId).Execute()

Delete MSP webhook setting



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
	webhookId := "webhookId_example" // string | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspWebhookSettingAPI.DeleteWebhookForMsp(context.Background(), mspId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspWebhookSettingAPI.DeleteWebhookForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhookForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspWebhookSettingAPI.DeleteWebhookForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookForMspRequest struct via the builder pattern


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


## GetWebhookForMsp

> OperationResponseListWebhookSetting GetWebhookForMsp(ctx, mspId).Execute()

Get MSP webhook setting list



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
	resp, r, err := apiClient.MspWebhookSettingAPI.GetWebhookForMsp(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspWebhookSettingAPI.GetWebhookForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookForMsp`: OperationResponseListWebhookSetting
	fmt.Fprintf(os.Stdout, "Response from `MspWebhookSettingAPI.GetWebhookForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookForMspRequest struct via the builder pattern


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


## GetWebhookLogsForMsp

> OperationResponseGridVOOpenApiWebhookDispatchLogVO GetWebhookLogsForMsp(ctx, mspId).Page(page).PageSize(pageSize).FiltersWebhookId(filtersWebhookId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()

Get MSP webhook dispatch log list



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
	filtersWebhookId := "filtersWebhookId_example" // string | Filter query parameters, support field d66aed17ea7238154ba610710d9a2890
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field 1679297710438
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field 1681889710438

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspWebhookSettingAPI.GetWebhookLogsForMsp(context.Background(), mspId).Page(page).PageSize(pageSize).FiltersWebhookId(filtersWebhookId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspWebhookSettingAPI.GetWebhookLogsForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookLogsForMsp`: OperationResponseGridVOOpenApiWebhookDispatchLogVO
	fmt.Fprintf(os.Stdout, "Response from `MspWebhookSettingAPI.GetWebhookLogsForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookLogsForMspRequest struct via the builder pattern


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


## ModifyWebhookForMsp

> OperationResponseWithoutResult ModifyWebhookForMsp(ctx, mspId, webhookId).OpenApiWebhookSettingEditVO(openApiWebhookSettingEditVO).Execute()

Modify MSP webhook setting



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
	webhookId := "webhookId_example" // string | Webhook ID
	openApiWebhookSettingEditVO := *openapiclient.NewOpenApiWebhookSettingEditVO(int32(123), int32(123), []string{"[http(s)://webhook.site/4a566f9e-0b77-42e2-9a34-a78]"}) // OpenApiWebhookSettingEditVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspWebhookSettingAPI.ModifyWebhookForMsp(context.Background(), mspId, webhookId).OpenApiWebhookSettingEditVO(openApiWebhookSettingEditVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspWebhookSettingAPI.ModifyWebhookForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWebhookForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspWebhookSettingAPI.ModifyWebhookForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWebhookForMspRequest struct via the builder pattern


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


## TestWebhookFormsp

> OperationResponseWithoutResult TestWebhookFormsp(ctx, mspId, webhookId).OpenApiWebhookSettingTestVO(openApiWebhookSettingTestVO).Execute()

Test MSP webhook setting



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
	webhookId := "webhookId_example" // string | Webhook ID
	openApiWebhookSettingTestVO := *openapiclient.NewOpenApiWebhookSettingTestVO("Name_example", int32(123), []string{"[http(s)://webhook.site/4a566f9e-0b77-42e2-9a34-a78]"}) // OpenApiWebhookSettingTestVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspWebhookSettingAPI.TestWebhookFormsp(context.Background(), mspId, webhookId).OpenApiWebhookSettingTestVO(openApiWebhookSettingTestVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspWebhookSettingAPI.TestWebhookFormsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestWebhookFormsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspWebhookSettingAPI.TestWebhookFormsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhookFormspRequest struct via the builder pattern


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


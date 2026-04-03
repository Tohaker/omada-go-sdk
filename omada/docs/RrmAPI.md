# RrmAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRrmIncompatibleAps**](RrmAPI.md#getrrmincompatibleaps) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rrm/incompatible | Get incompatible rrm ap list
[**GetRrmSetting**](RrmAPI.md#getrrmsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rrm/config | Get Rrm setting
[**GetRrmSettingTemplate**](RrmAPI.md#getrrmsettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rrm/config | Get Rrm template setting
[**ModifyRrmSetting**](RrmAPI.md#modifyrrmsetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/rrm/config | Modify Rrm setting
[**ModifyRrmSettingTemplate**](RrmAPI.md#modifyrrmsettingtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rrm/config | Modify Rrm template setting



## GetRrmIncompatibleAps

> OperationResponseGridVORrmIncompatibleApOpenApiVO GetRrmIncompatibleAps(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()

Get incompatible rrm ap list



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
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name,mac (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RrmAPI.GetRrmIncompatibleAps(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RrmAPI.GetRrmIncompatibleAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRrmIncompatibleAps`: OperationResponseGridVORrmIncompatibleApOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `RrmAPI.GetRrmIncompatibleAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRrmIncompatibleApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name,mac | 

### Return type

[**OperationResponseGridVORrmIncompatibleApOpenApiVO**](OperationResponseGridVORrmIncompatibleApOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRrmSetting

> OperationResponseRrmSettingOpenApiVO GetRrmSetting(ctx, omadacId, siteId).Execute()

Get Rrm setting



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
	resp, r, err := apiClient.RrmAPI.GetRrmSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RrmAPI.GetRrmSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRrmSetting`: OperationResponseRrmSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `RrmAPI.GetRrmSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRrmSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseRrmSettingOpenApiVO**](OperationResponseRrmSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRrmSettingTemplate

> OperationResponseRrmSettingOpenApiVO GetRrmSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get Rrm template setting



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
	resp, r, err := apiClient.RrmAPI.GetRrmSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RrmAPI.GetRrmSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRrmSettingTemplate`: OperationResponseRrmSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `RrmAPI.GetRrmSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRrmSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseRrmSettingOpenApiVO**](OperationResponseRrmSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRrmSetting

> OperationResponseWithoutResult ModifyRrmSetting(ctx, omadacId, siteId).RrmSettingOpenApiVO(rrmSettingOpenApiVO).Execute()

Modify Rrm setting



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
	rrmSettingOpenApiVO := *openapiclient.NewRrmSettingOpenApiVO(int32(123)) // RrmSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RrmAPI.ModifyRrmSetting(context.Background(), omadacId, siteId).RrmSettingOpenApiVO(rrmSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RrmAPI.ModifyRrmSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRrmSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RrmAPI.ModifyRrmSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRrmSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rrmSettingOpenApiVO** | [**RrmSettingOpenApiVO**](RrmSettingOpenApiVO.md) |  | 

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


## ModifyRrmSettingTemplate

> OperationResponseWithoutResult ModifyRrmSettingTemplate(ctx, omadacId, siteTemplateId).RrmSettingOpenApiVO(rrmSettingOpenApiVO).Execute()

Modify Rrm template setting



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
	rrmSettingOpenApiVO := *openapiclient.NewRrmSettingOpenApiVO(int32(123)) // RrmSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RrmAPI.ModifyRrmSettingTemplate(context.Background(), omadacId, siteTemplateId).RrmSettingOpenApiVO(rrmSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RrmAPI.ModifyRrmSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRrmSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RrmAPI.ModifyRrmSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRrmSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rrmSettingOpenApiVO** | [**RrmSettingOpenApiVO**](RrmSettingOpenApiVO.md) |  | 

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


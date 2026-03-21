# \AbnormalDetectAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnomalyEventSettingForSite**](AbnormalDetectAPI.md#GetAnomalyEventSettingForSite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/setting | Get site anomaly event setting
[**GetAnomalyEventSettingForSiteTemlate**](AbnormalDetectAPI.md#GetAnomalyEventSettingForSiteTemlate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/anomaly/setting | Get site template anomaly event setting
[**GetNotificationTemplate**](AbnormalDetectAPI.md#GetNotificationTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/anomaly/notification | Get abnormal notification template info
[**GetNotificationTemplate1**](AbnormalDetectAPI.md#GetNotificationTemplate1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/notification | Get abnormal notification info
[**ModifyAnomalyEventSettingForSite**](AbnormalDetectAPI.md#ModifyAnomalyEventSettingForSite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/setting/modify | Modify site anomaly event setting
[**ModifyAnomalyEventSettingForSiteTemplate**](AbnormalDetectAPI.md#ModifyAnomalyEventSettingForSiteTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/anomaly/setting/modify | Modify site template anomaly event setting
[**ModifyNotification**](AbnormalDetectAPI.md#ModifyNotification) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/anomaly/notification | Set abnormal notification info
[**ModifyNotificationTemplate**](AbnormalDetectAPI.md#ModifyNotificationTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/anomaly/notification | Set abnormal notification template info



## GetAnomalyEventSettingForSite

> OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO GetAnomalyEventSettingForSite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).FiltersEnable(filtersEnable).FiltersLevel(filtersLevel).FiltersCategory(filtersCategory).Execute()

Get site anomaly event setting



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
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersEnable := int32(56) // int32 | Filter query parameters, support field enable, it should be a value as follows: 0:enable, 1:disable, example:0. (optional)
	filtersLevel := int32(56) // int32 | Filter query parameters, support field level, it should be a value as follows: 0:Critical, 1:Error, 2:Warning, 3: Info, example:1. (optional)
	filtersCategory := "filtersCategory_example" // string | Filter query parameters, support field category, one or more categories, each category should be a value as follows: 0:Networking, 1:Mesh, 2:Access, 3:Roaming, 4:Network Service, 5:Software/Configuration, 6:Hardware, 7:Security, 8:Throughput, 9:Coverage, example: 1 Or 2,3,4 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbnormalDetectAPI.GetAnomalyEventSettingForSite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).FiltersEnable(filtersEnable).FiltersLevel(filtersLevel).FiltersCategory(filtersCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbnormalDetectAPI.GetAnomalyEventSettingForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnomalyEventSettingForSite`: OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AbnormalDetectAPI.GetAnomalyEventSettingForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnomalyEventSettingForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersEnable** | **int32** | Filter query parameters, support field enable, it should be a value as follows: 0:enable, 1:disable, example:0. | 
 **filtersLevel** | **int32** | Filter query parameters, support field level, it should be a value as follows: 0:Critical, 1:Error, 2:Warning, 3: Info, example:1. | 
 **filtersCategory** | **string** | Filter query parameters, support field category, one or more categories, each category should be a value as follows: 0:Networking, 1:Mesh, 2:Access, 3:Roaming, 4:Network Service, 5:Software/Configuration, 6:Hardware, 7:Security, 8:Throughput, 9:Coverage, example: 1 Or 2,3,4 | 

### Return type

[**OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO**](OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnomalyEventSettingForSiteTemlate

> OperationResponseGridVOAnomalyEventSettingOpenApiVO GetAnomalyEventSettingForSiteTemlate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).FiltersEnable(filtersEnable).FiltersLevel(filtersLevel).FiltersCategory(filtersCategory).Execute()

Get site template anomaly event setting



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
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100)
	filtersEnable := int32(56) // int32 | Filter query parameters, support field enable, it should be a value as follows: 0:enable, 1:disable, example:0. (optional)
	filtersLevel := int32(56) // int32 | Filter query parameters, support field level, it should be a value as follows: 0:Critical, 1:Error, 2:Warning, 3: Info, example:1. (optional)
	filtersCategory := "filtersCategory_example" // string | Filter query parameters, support field category, one or more categories, each category should be a value as follows: 0:Networking, 1:Mesh, 2:Access, 3:Roaming, 4:Network Service, 5:Software/Configuration, 6:Hardware, 7:Security, 8:Throughput, 9:Coverage, example: 1 Or 2,3,4 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbnormalDetectAPI.GetAnomalyEventSettingForSiteTemlate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).FiltersEnable(filtersEnable).FiltersLevel(filtersLevel).FiltersCategory(filtersCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbnormalDetectAPI.GetAnomalyEventSettingForSiteTemlate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnomalyEventSettingForSiteTemlate`: OperationResponseGridVOAnomalyEventSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AbnormalDetectAPI.GetAnomalyEventSettingForSiteTemlate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnomalyEventSettingForSiteTemlateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000.(value:10,15,20,30,50,100) | 
 **filtersEnable** | **int32** | Filter query parameters, support field enable, it should be a value as follows: 0:enable, 1:disable, example:0. | 
 **filtersLevel** | **int32** | Filter query parameters, support field level, it should be a value as follows: 0:Critical, 1:Error, 2:Warning, 3: Info, example:1. | 
 **filtersCategory** | **string** | Filter query parameters, support field category, one or more categories, each category should be a value as follows: 0:Networking, 1:Mesh, 2:Access, 3:Roaming, 4:Network Service, 5:Software/Configuration, 6:Hardware, 7:Security, 8:Throughput, 9:Coverage, example: 1 Or 2,3,4 | 

### Return type

[**OperationResponseGridVOAnomalyEventSettingOpenApiVO**](OperationResponseGridVOAnomalyEventSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationTemplate

> OperationResponseNotificationOpenApiVO GetNotificationTemplate(ctx, omadacId, siteTemplateId).Execute()

Get abnormal notification template info



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
	resp, r, err := apiClient.AbnormalDetectAPI.GetNotificationTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbnormalDetectAPI.GetNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationTemplate`: OperationResponseNotificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AbnormalDetectAPI.GetNotificationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseNotificationOpenApiVO**](OperationResponseNotificationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationTemplate1

> OperationResponseNotificationOpenApiVO GetNotificationTemplate1(ctx, omadacId, siteId).Execute()

Get abnormal notification info



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
	resp, r, err := apiClient.AbnormalDetectAPI.GetNotificationTemplate1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbnormalDetectAPI.GetNotificationTemplate1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationTemplate1`: OperationResponseNotificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AbnormalDetectAPI.GetNotificationTemplate1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationTemplate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseNotificationOpenApiVO**](OperationResponseNotificationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAnomalyEventSettingForSite

> OperationResponseWithoutResult ModifyAnomalyEventSettingForSite(ctx, omadacId, siteId).AnomalyEventSettingEditOpenApiVO(anomalyEventSettingEditOpenApiVO).Execute()

Modify site anomaly event setting



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
	anomalyEventSettingEditOpenApiVO := *openapiclient.NewAnomalyEventSettingEditOpenApiVO() // AnomalyEventSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbnormalDetectAPI.ModifyAnomalyEventSettingForSite(context.Background(), omadacId, siteId).AnomalyEventSettingEditOpenApiVO(anomalyEventSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbnormalDetectAPI.ModifyAnomalyEventSettingForSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAnomalyEventSettingForSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AbnormalDetectAPI.ModifyAnomalyEventSettingForSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAnomalyEventSettingForSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **anomalyEventSettingEditOpenApiVO** | [**AnomalyEventSettingEditOpenApiVO**](AnomalyEventSettingEditOpenApiVO.md) |  | 

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


## ModifyAnomalyEventSettingForSiteTemplate

> OperationResponseWithoutResult ModifyAnomalyEventSettingForSiteTemplate(ctx, omadacId, siteTemplateId).AnomalyEventSettingEditOpenApiVO(anomalyEventSettingEditOpenApiVO).Execute()

Modify site template anomaly event setting



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
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId
	anomalyEventSettingEditOpenApiVO := *openapiclient.NewAnomalyEventSettingEditOpenApiVO() // AnomalyEventSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbnormalDetectAPI.ModifyAnomalyEventSettingForSiteTemplate(context.Background(), omadacId, siteTemplateId).AnomalyEventSettingEditOpenApiVO(anomalyEventSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbnormalDetectAPI.ModifyAnomalyEventSettingForSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAnomalyEventSettingForSiteTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AbnormalDetectAPI.ModifyAnomalyEventSettingForSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | siteTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAnomalyEventSettingForSiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **anomalyEventSettingEditOpenApiVO** | [**AnomalyEventSettingEditOpenApiVO**](AnomalyEventSettingEditOpenApiVO.md) |  | 

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


## ModifyNotification

> OperationResponseSetNotificationOpenApiVO ModifyNotification(ctx, omadacId, siteId).SetNotificationOpenApiVO(setNotificationOpenApiVO).Execute()

Set abnormal notification info



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
	setNotificationOpenApiVO := *openapiclient.NewSetNotificationOpenApiVO(*openapiclient.NewNotificationConfigurationOpenApiVO(map[string]bool{"key": false}, map[string]bool{"key": false}, map[string]bool{"key": false}), *openapiclient.NewAlertEmailSettingVO(false)) // SetNotificationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbnormalDetectAPI.ModifyNotification(context.Background(), omadacId, siteId).SetNotificationOpenApiVO(setNotificationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbnormalDetectAPI.ModifyNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyNotification`: OperationResponseSetNotificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AbnormalDetectAPI.ModifyNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setNotificationOpenApiVO** | [**SetNotificationOpenApiVO**](SetNotificationOpenApiVO.md) |  | 

### Return type

[**OperationResponseSetNotificationOpenApiVO**](OperationResponseSetNotificationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyNotificationTemplate

> OperationResponseSetNotificationOpenApiVO ModifyNotificationTemplate(ctx, omadacId, siteTemplateId).SetNotificationOpenApiVO(setNotificationOpenApiVO).Execute()

Set abnormal notification template info



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
	setNotificationOpenApiVO := *openapiclient.NewSetNotificationOpenApiVO(*openapiclient.NewNotificationConfigurationOpenApiVO(map[string]bool{"key": false}, map[string]bool{"key": false}, map[string]bool{"key": false}), *openapiclient.NewAlertEmailSettingVO(false)) // SetNotificationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbnormalDetectAPI.ModifyNotificationTemplate(context.Background(), omadacId, siteTemplateId).SetNotificationOpenApiVO(setNotificationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbnormalDetectAPI.ModifyNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyNotificationTemplate`: OperationResponseSetNotificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AbnormalDetectAPI.ModifyNotificationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyNotificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setNotificationOpenApiVO** | [**SetNotificationOpenApiVO**](SetNotificationOpenApiVO.md) |  | 

### Return type

[**OperationResponseSetNotificationOpenApiVO**](OperationResponseSetNotificationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


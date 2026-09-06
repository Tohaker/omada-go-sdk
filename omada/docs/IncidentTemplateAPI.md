# IncidentTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnomalyEventSettingForSiteTemlate**](IncidentTemplateAPI.md#getanomalyeventsettingforsitetemlate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/anomaly/setting | Get site template anomaly event setting
[**GetNotificationTemplate**](IncidentTemplateAPI.md#getnotificationtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/anomaly/notification | Get abnormal notification template info
[**ModifyAnomalyEventSettingForSiteTemplate**](IncidentTemplateAPI.md#modifyanomalyeventsettingforsitetemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/anomaly/setting/modify | Modify site template anomaly event setting
[**ModifyNotificationTemplate**](IncidentTemplateAPI.md#modifynotificationtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/anomaly/notification | Set abnormal notification template info



## GetAnomalyEventSettingForSiteTemlate

> OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO GetAnomalyEventSettingForSiteTemlate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).FiltersEnable(filtersEnable).FiltersLevel(filtersLevel).FiltersCategory(filtersCategory).Execute()

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
	filtersCategory := "filtersCategory_example" // string | Filter query parameters, support field category, one or more categories, each category should be a value as follows: 11:Access, 12:Authentication, 13:Roaming, 14:Wireless Network, 15:Wired Network, 16:Link, 17:WAN and Services, 18:Device Status, 19:Security, example: 11 Or 12,13,14 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentTemplateAPI.GetAnomalyEventSettingForSiteTemlate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).FiltersEnable(filtersEnable).FiltersLevel(filtersLevel).FiltersCategory(filtersCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTemplateAPI.GetAnomalyEventSettingForSiteTemlate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnomalyEventSettingForSiteTemlate`: OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentTemplateAPI.GetAnomalyEventSettingForSiteTemlate`: %v\n", resp)
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
 **filtersCategory** | **string** | Filter query parameters, support field category, one or more categories, each category should be a value as follows: 11:Access, 12:Authentication, 13:Roaming, 14:Wireless Network, 15:Wired Network, 16:Link, 17:WAN and Services, 18:Device Status, 19:Security, example: 11 Or 12,13,14 | 

### Return type

[**OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO**](OperationResponseAnomalySettingGridVOAnomalyEventSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

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
	resp, r, err := apiClient.IncidentTemplateAPI.GetNotificationTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTemplateAPI.GetNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationTemplate`: OperationResponseNotificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentTemplateAPI.GetNotificationTemplate`: %v\n", resp)
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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	anomalyEventSettingEditOpenApiVO := *openapiclient.NewAnomalyEventSettingEditOpenApiVO("01001001") // AnomalyEventSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentTemplateAPI.ModifyAnomalyEventSettingForSiteTemplate(context.Background(), omadacId, siteTemplateId).AnomalyEventSettingEditOpenApiVO(anomalyEventSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTemplateAPI.ModifyAnomalyEventSettingForSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAnomalyEventSettingForSiteTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IncidentTemplateAPI.ModifyAnomalyEventSettingForSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAnomalyEventSettingForSiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **anomalyEventSettingEditOpenApiVO** | [**AnomalyEventSettingEditOpenApiVO**](AnomalyEventSettingEditOpenApiVO.md) |  | 

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
	setNotificationOpenApiVO := *openapiclient.NewSetNotificationOpenApiVO(*openapiclient.NewAlertEmailSettingVO(false), *openapiclient.NewNotificationConfigurationOpenApiVO(map[string]bool{"key": false}, map[string]bool{"key": false}, map[string]bool{"key": false})) // SetNotificationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentTemplateAPI.ModifyNotificationTemplate(context.Background(), omadacId, siteTemplateId).SetNotificationOpenApiVO(setNotificationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTemplateAPI.ModifyNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyNotificationTemplate`: OperationResponseSetNotificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IncidentTemplateAPI.ModifyNotificationTemplate`: %v\n", resp)
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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


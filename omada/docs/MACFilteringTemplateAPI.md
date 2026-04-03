# MACFilteringTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMacFilteringTemplate**](MACFilteringTemplateAPI.md#createmacfilteringtemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-filters | Create MAC filtering template
[**DeleteMacFilteringTemplate**](MACFilteringTemplateAPI.md#deletemacfilteringtemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-filters/{filterId} | Delete MAC filtering template
[**GetGridAllowMacFilteringTemplate**](MACFilteringTemplateAPI.md#getgridallowmacfilteringtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-filters/allow | Get allow MAC filtering template list
[**GetGridDenyMacFilteringTemplate**](MACFilteringTemplateAPI.md#getgriddenymacfilteringtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-filters/deny | Get deny MAC filtering template list
[**GetMacFilteringGeneralSettingTemplate**](MACFilteringTemplateAPI.md#getmacfilteringgeneralsettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-filter | Get MAC filtering template general setting
[**ModifyMacFilteringGeneralSettingTemplate**](MACFilteringTemplateAPI.md#modifymacfilteringgeneralsettingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-filter | Modify MAC filtering template general setting
[**ModifyMacFilteringTemplate**](MACFilteringTemplateAPI.md#modifymacfilteringtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-filters/{filterId} | Modify MAC filtering template



## CreateMacFilteringTemplate

> OperationResponseWithoutResult CreateMacFilteringTemplate(ctx, omadacId, siteTemplateId).MacFiltering(macFiltering).Execute()

Create MAC filtering template



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
	macFiltering := *openapiclient.NewMacFiltering(int32(123), "Name_example", int32(123)) // MacFiltering | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringTemplateAPI.CreateMacFilteringTemplate(context.Background(), omadacId, siteTemplateId).MacFiltering(macFiltering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringTemplateAPI.CreateMacFilteringTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMacFilteringTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringTemplateAPI.CreateMacFilteringTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMacFilteringTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macFiltering** | [**MacFiltering**](MacFiltering.md) |  | 

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


## DeleteMacFilteringTemplate

> OperationResponseWithoutResult DeleteMacFilteringTemplate(ctx, omadacId, siteTemplateId, filterId).Execute()

Delete MAC filtering template



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
	filterId := "filterId_example" // string | MAC filtering template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringTemplateAPI.DeleteMacFilteringTemplate(context.Background(), omadacId, siteTemplateId, filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringTemplateAPI.DeleteMacFilteringTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMacFilteringTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringTemplateAPI.DeleteMacFilteringTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**filterId** | **string** | MAC filtering template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMacFilteringTemplateRequest struct via the builder pattern


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


## GetGridAllowMacFilteringTemplate

> OperationResponseGridVOMacFiltering GetGridAllowMacFilteringTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get allow MAC filtering template list



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
	resp, r, err := apiClient.MACFilteringTemplateAPI.GetGridAllowMacFilteringTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringTemplateAPI.GetGridAllowMacFilteringTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAllowMacFilteringTemplate`: OperationResponseGridVOMacFiltering
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringTemplateAPI.GetGridAllowMacFilteringTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAllowMacFilteringTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOMacFiltering**](OperationResponseGridVOMacFiltering.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDenyMacFilteringTemplate

> OperationResponseGridVOMacFiltering GetGridDenyMacFilteringTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get deny MAC filtering template list



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
	resp, r, err := apiClient.MACFilteringTemplateAPI.GetGridDenyMacFilteringTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringTemplateAPI.GetGridDenyMacFilteringTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDenyMacFilteringTemplate`: OperationResponseGridVOMacFiltering
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringTemplateAPI.GetGridDenyMacFilteringTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDenyMacFilteringTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOMacFiltering**](OperationResponseGridVOMacFiltering.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacFilteringGeneralSettingTemplate

> OperationResponseMacFilteringGeneralSetting GetMacFilteringGeneralSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get MAC filtering template general setting



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
	resp, r, err := apiClient.MACFilteringTemplateAPI.GetMacFilteringGeneralSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringTemplateAPI.GetMacFilteringGeneralSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacFilteringGeneralSettingTemplate`: OperationResponseMacFilteringGeneralSetting
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringTemplateAPI.GetMacFilteringGeneralSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacFilteringGeneralSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMacFilteringGeneralSetting**](OperationResponseMacFilteringGeneralSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMacFilteringGeneralSettingTemplate

> OperationResponseWithoutResult ModifyMacFilteringGeneralSettingTemplate(ctx, omadacId, siteTemplateId).MacFilteringGeneralSetting(macFilteringGeneralSetting).Execute()

Modify MAC filtering template general setting



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
	macFilteringGeneralSetting := *openapiclient.NewMacFilteringGeneralSetting(false) // MacFilteringGeneralSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringTemplateAPI.ModifyMacFilteringGeneralSettingTemplate(context.Background(), omadacId, siteTemplateId).MacFilteringGeneralSetting(macFilteringGeneralSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringTemplateAPI.ModifyMacFilteringGeneralSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMacFilteringGeneralSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringTemplateAPI.ModifyMacFilteringGeneralSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMacFilteringGeneralSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macFilteringGeneralSetting** | [**MacFilteringGeneralSetting**](MacFilteringGeneralSetting.md) |  | 

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


## ModifyMacFilteringTemplate

> OperationResponseWithoutResult ModifyMacFilteringTemplate(ctx, omadacId, siteTemplateId, filterId).MacFiltering(macFiltering).Execute()

Modify MAC filtering template



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
	filterId := "filterId_example" // string | MAC filtering template ID
	macFiltering := *openapiclient.NewMacFiltering(int32(123), "Name_example", int32(123)) // MacFiltering | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACFilteringTemplateAPI.ModifyMacFilteringTemplate(context.Background(), omadacId, siteTemplateId, filterId).MacFiltering(macFiltering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACFilteringTemplateAPI.ModifyMacFilteringTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMacFilteringTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MACFilteringTemplateAPI.ModifyMacFilteringTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**filterId** | **string** | MAC filtering template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMacFilteringTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **macFiltering** | [**MacFiltering**](MacFiltering.md) |  | 

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


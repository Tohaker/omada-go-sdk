# IPMACBindingTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteIpMacBindingsTemplate**](IPMACBindingTemplateAPI.md#batchdeleteipmacbindingstemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ip-mac-binds/batch-delete | Batch delete IP-MAC bindings template
[**CreateIpMacBindingTemplate**](IPMACBindingTemplateAPI.md#createipmacbindingtemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ip-mac-binds | Create IP-MAC binding template
[**DeleteIpMacBindingTemplate**](IPMACBindingTemplateAPI.md#deleteipmacbindingtemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ip-mac-binds/{bindId} | Delete IP-MAC binding template
[**GetGridIpMacBindingTemplate**](IPMACBindingTemplateAPI.md#getgridipmacbindingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ip-mac-binds | Get IP-MAC binding template list
[**GetIpMacBindingGeneralSettingTemplate**](IPMACBindingTemplateAPI.md#getipmacbindinggeneralsettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ip-mac-bind | Get IP-MAC binding template general setting
[**ModifyIpMacBindingGeneralSettingTemplate**](IPMACBindingTemplateAPI.md#modifyipmacbindinggeneralsettingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ip-mac-bind | Modify IP-MAC binding template general setting
[**ModifyIpMacBindingTemplate**](IPMACBindingTemplateAPI.md#modifyipmacbindingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ip-mac-binds/{bindId} | Modify IP-MAC binding template



## BatchDeleteIpMacBindingsTemplate

> OperationResponseWithoutResult BatchDeleteIpMacBindingsTemplate(ctx, omadacId, siteTemplateId).BatchIds(batchIds).Execute()

Batch delete IP-MAC bindings template



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
	batchIds := *openapiclient.NewBatchIds([]string{"Ids_example"}) // BatchIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingTemplateAPI.BatchDeleteIpMacBindingsTemplate(context.Background(), omadacId, siteTemplateId).BatchIds(batchIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingTemplateAPI.BatchDeleteIpMacBindingsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteIpMacBindingsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingTemplateAPI.BatchDeleteIpMacBindingsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteIpMacBindingsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchIds** | [**BatchIds**](BatchIds.md) |  | 

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


## CreateIpMacBindingTemplate

> OperationResponseWithoutResult CreateIpMacBindingTemplate(ctx, omadacId, siteTemplateId).IPMacBinding(iPMacBinding).Execute()

Create IP-MAC binding template



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
	iPMacBinding := *openapiclient.NewIPMacBinding("InterfaceId_example", int32(123), "Ip_example", "Mac_example", false) // IPMacBinding | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingTemplateAPI.CreateIpMacBindingTemplate(context.Background(), omadacId, siteTemplateId).IPMacBinding(iPMacBinding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingTemplateAPI.CreateIpMacBindingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIpMacBindingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingTemplateAPI.CreateIpMacBindingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIpMacBindingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iPMacBinding** | [**IPMacBinding**](IPMacBinding.md) |  | 

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


## DeleteIpMacBindingTemplate

> OperationResponseWithoutResult DeleteIpMacBindingTemplate(ctx, omadacId, siteTemplateId, bindId).Execute()

Delete IP-MAC binding template



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
	bindId := "bindId_example" // string | IP MAC binding template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingTemplateAPI.DeleteIpMacBindingTemplate(context.Background(), omadacId, siteTemplateId, bindId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingTemplateAPI.DeleteIpMacBindingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIpMacBindingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingTemplateAPI.DeleteIpMacBindingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**bindId** | **string** | IP MAC binding template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpMacBindingTemplateRequest struct via the builder pattern


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


## GetGridIpMacBindingTemplate

> OperationResponseGridVOIPMacBinding GetGridIpMacBindingTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get IP-MAC binding template list



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
	resp, r, err := apiClient.IPMACBindingTemplateAPI.GetGridIpMacBindingTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingTemplateAPI.GetGridIpMacBindingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridIpMacBindingTemplate`: OperationResponseGridVOIPMacBinding
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingTemplateAPI.GetGridIpMacBindingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridIpMacBindingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOIPMacBinding**](OperationResponseGridVOIPMacBinding.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpMacBindingGeneralSettingTemplate

> OperationResponseIPMacBindingGeneralSetting GetIpMacBindingGeneralSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get IP-MAC binding template general setting



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
	resp, r, err := apiClient.IPMACBindingTemplateAPI.GetIpMacBindingGeneralSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingTemplateAPI.GetIpMacBindingGeneralSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpMacBindingGeneralSettingTemplate`: OperationResponseIPMacBindingGeneralSetting
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingTemplateAPI.GetIpMacBindingGeneralSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpMacBindingGeneralSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIPMacBindingGeneralSetting**](OperationResponseIPMacBindingGeneralSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIpMacBindingGeneralSettingTemplate

> OperationResponseWithoutResult ModifyIpMacBindingGeneralSettingTemplate(ctx, omadacId, siteTemplateId).IPMacBindingGeneralSetting(iPMacBindingGeneralSetting).Execute()

Modify IP-MAC binding template general setting



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
	iPMacBindingGeneralSetting := *openapiclient.NewIPMacBindingGeneralSetting(false) // IPMacBindingGeneralSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingTemplateAPI.ModifyIpMacBindingGeneralSettingTemplate(context.Background(), omadacId, siteTemplateId).IPMacBindingGeneralSetting(iPMacBindingGeneralSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingTemplateAPI.ModifyIpMacBindingGeneralSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpMacBindingGeneralSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingTemplateAPI.ModifyIpMacBindingGeneralSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpMacBindingGeneralSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iPMacBindingGeneralSetting** | [**IPMacBindingGeneralSetting**](IPMacBindingGeneralSetting.md) |  | 

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


## ModifyIpMacBindingTemplate

> OperationResponseWithoutResult ModifyIpMacBindingTemplate(ctx, omadacId, siteTemplateId, bindId).IPMacBinding(iPMacBinding).Execute()

Modify IP-MAC binding template



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
	bindId := "bindId_example" // string | IP MAC binding template ID
	iPMacBinding := *openapiclient.NewIPMacBinding("InterfaceId_example", int32(123), "Ip_example", "Mac_example", false) // IPMacBinding | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingTemplateAPI.ModifyIpMacBindingTemplate(context.Background(), omadacId, siteTemplateId, bindId).IPMacBinding(iPMacBinding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingTemplateAPI.ModifyIpMacBindingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpMacBindingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingTemplateAPI.ModifyIpMacBindingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**bindId** | **string** | IP MAC binding template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpMacBindingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iPMacBinding** | [**IPMacBinding**](IPMacBinding.md) |  | 

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


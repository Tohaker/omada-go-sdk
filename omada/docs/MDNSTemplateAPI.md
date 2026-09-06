# MDNSTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMdnsTemplate**](MDNSTemplateAPI.md#createmdnstemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/mdns | Create a new mDNS template rule
[**DeleteMdnsTemplate**](MDNSTemplateAPI.md#deletemdnstemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/mdns/{mdnsId} | Delete an exist mDNS template rule
[**GetMdnsGridTemplate**](MDNSTemplateAPI.md#getmdnsgridtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/mdns | Get mDNS template rule list
[**ModifyMdnsTemplate**](MDNSTemplateAPI.md#modifymdnstemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/mdns/{mdnsId} | Modify an exist mDNS template rule



## CreateMdnsTemplate

> OperationResponseResIdOpenApiVO CreateMdnsTemplate(ctx, omadacId, siteTemplateId).CreateMdnsRuleTemplateOpenApiVO(createMdnsRuleTemplateOpenApiVO).Execute()

Create a new mDNS template rule



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	createMdnsRuleTemplateOpenApiVO := *openapiclient.NewCreateMdnsRuleTemplateOpenApiVO("Name_example", []string{"ProfileIds_example"}, false) // CreateMdnsRuleTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDNSTemplateAPI.CreateMdnsTemplate(context.Background(), omadacId, siteTemplateId).CreateMdnsRuleTemplateOpenApiVO(createMdnsRuleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDNSTemplateAPI.CreateMdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMdnsTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MDNSTemplateAPI.CreateMdnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMdnsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMdnsRuleTemplateOpenApiVO** | [**CreateMdnsRuleTemplateOpenApiVO**](CreateMdnsRuleTemplateOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMdnsTemplate

> OperationResponseWithoutResult DeleteMdnsTemplate(ctx, omadacId, siteTemplateId, mdnsId).Execute()

Delete an exist mDNS template rule



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	mdnsId := "mdnsId_example" // string | mDNS rule template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDNSTemplateAPI.DeleteMdnsTemplate(context.Background(), omadacId, siteTemplateId, mdnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDNSTemplateAPI.DeleteMdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMdnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MDNSTemplateAPI.DeleteMdnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**mdnsId** | **string** | mDNS rule template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMdnsTemplateRequest struct via the builder pattern


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


## GetMdnsGridTemplate

> OperationResponseGridVOMdnsRuleTemplateOpenApiVO GetMdnsGridTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get mDNS template rule list



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDNSTemplateAPI.GetMdnsGridTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDNSTemplateAPI.GetMdnsGridTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMdnsGridTemplate`: OperationResponseGridVOMdnsRuleTemplateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MDNSTemplateAPI.GetMdnsGridTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMdnsGridTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOMdnsRuleTemplateOpenApiVO**](OperationResponseGridVOMdnsRuleTemplateOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMdnsTemplate

> OperationResponseWithoutResult ModifyMdnsTemplate(ctx, omadacId, siteTemplateId, mdnsId).CreateMdnsRuleTemplateOpenApiVO(createMdnsRuleTemplateOpenApiVO).Execute()

Modify an exist mDNS template rule



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	mdnsId := "mdnsId_example" // string | mDNS rule template ID
	createMdnsRuleTemplateOpenApiVO := *openapiclient.NewCreateMdnsRuleTemplateOpenApiVO("Name_example", []string{"ProfileIds_example"}, false) // CreateMdnsRuleTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDNSTemplateAPI.ModifyMdnsTemplate(context.Background(), omadacId, siteTemplateId, mdnsId).CreateMdnsRuleTemplateOpenApiVO(createMdnsRuleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDNSTemplateAPI.ModifyMdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMdnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MDNSTemplateAPI.ModifyMdnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**mdnsId** | **string** | mDNS rule template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMdnsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createMdnsRuleTemplateOpenApiVO** | [**CreateMdnsRuleTemplateOpenApiVO**](CreateMdnsRuleTemplateOpenApiVO.md) |  | 

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


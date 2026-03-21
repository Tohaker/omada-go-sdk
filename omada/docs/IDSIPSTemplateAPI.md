# \IDSIPSTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplateAllowList**](IDSIPSTemplateAPI.md#CreateTemplateAllowList) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/ips/allow-list | Create site template&#39;s allow list
[**DeleteTemplateAllowList**](IDSIPSTemplateAPI.md#DeleteTemplateAllowList) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/ips/allow-list/{entryId} | Delete template allow list
[**EditTemplateAllowList**](IDSIPSTemplateAPI.md#EditTemplateAllowList) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/ips/allow-list | Edit site template&#39;s allow list
[**GetIpsTemplateConfig**](IDSIPSTemplateAPI.md#GetIpsTemplateConfig) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/ips | Get site template IDS/IPS config
[**GetTemplateGridAllowList**](IDSIPSTemplateAPI.md#GetTemplateGridAllowList) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/grid/allow-list | Get site template grid ips allow list
[**ModifyIpsTemplateConfig**](IDSIPSTemplateAPI.md#ModifyIpsTemplateConfig) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/network-security/ips | Modify site template IDS/IPS config



## CreateTemplateAllowList

> OperationResponse CreateTemplateAllowList(ctx, omadacId, siteTemplateId).NewIPSAllowListEntry(newIPSAllowListEntry).Execute()

Create site template's allow list



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
	newIPSAllowListEntry := *openapiclient.NewNewIPSAllowListEntry() // NewIPSAllowListEntry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSTemplateAPI.CreateTemplateAllowList(context.Background(), omadacId, siteTemplateId).NewIPSAllowListEntry(newIPSAllowListEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSTemplateAPI.CreateTemplateAllowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateAllowList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSTemplateAPI.CreateTemplateAllowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateAllowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newIPSAllowListEntry** | [**NewIPSAllowListEntry**](NewIPSAllowListEntry.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplateAllowList

> OperationResponse DeleteTemplateAllowList(ctx, omadacId, siteTemplateId, entryId).Execute()

Delete template allow list



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
	entryId := "entryId_example" // string | Allow entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSTemplateAPI.DeleteTemplateAllowList(context.Background(), omadacId, siteTemplateId, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSTemplateAPI.DeleteTemplateAllowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplateAllowList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSTemplateAPI.DeleteTemplateAllowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**entryId** | **string** | Allow entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateAllowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTemplateAllowList

> OperationResponse EditTemplateAllowList(ctx, omadacId, siteTemplateId).ModifyIPSAllowListEntry(modifyIPSAllowListEntry).Execute()

Edit site template's allow list



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
	modifyIPSAllowListEntry := *openapiclient.NewModifyIPSAllowListEntry("Id_example") // ModifyIPSAllowListEntry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSTemplateAPI.EditTemplateAllowList(context.Background(), omadacId, siteTemplateId).ModifyIPSAllowListEntry(modifyIPSAllowListEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSTemplateAPI.EditTemplateAllowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTemplateAllowList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSTemplateAPI.EditTemplateAllowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTemplateAllowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyIPSAllowListEntry** | [**ModifyIPSAllowListEntry**](ModifyIPSAllowListEntry.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpsTemplateConfig

> OperationResponseIpsInfo GetIpsTemplateConfig(ctx, omadacId, siteTemplateId).Execute()

Get site template IDS/IPS config



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
	resp, r, err := apiClient.IDSIPSTemplateAPI.GetIpsTemplateConfig(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSTemplateAPI.GetIpsTemplateConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpsTemplateConfig`: OperationResponseIpsInfo
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSTemplateAPI.GetIpsTemplateConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpsTemplateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIpsInfo**](OperationResponseIpsInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateGridAllowList

> OperationResponse GetTemplateGridAllowList(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get site template grid ips allow list



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field IP / Subnet (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSTemplateAPI.GetTemplateGridAllowList(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSTemplateAPI.GetTemplateGridAllowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateGridAllowList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSTemplateAPI.GetTemplateGridAllowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateGridAllowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field IP / Subnet | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIpsTemplateConfig

> OperationResponse ModifyIpsTemplateConfig(ctx, omadacId, siteTemplateId).IpsInfo(ipsInfo).Execute()

Modify site template IDS/IPS config



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
	ipsInfo := *openapiclient.NewIpsInfo(false) // IpsInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSTemplateAPI.ModifyIpsTemplateConfig(context.Background(), omadacId, siteTemplateId).IpsInfo(ipsInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSTemplateAPI.ModifyIpsTemplateConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpsTemplateConfig`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSTemplateAPI.ModifyIpsTemplateConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpsTemplateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ipsInfo** | [**IpsInfo**](IpsInfo.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


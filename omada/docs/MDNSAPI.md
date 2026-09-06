# MDNSAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMdns**](MDNSAPI.md#createmdns) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mdns | Create a new mDNS rule
[**DeleteMdns**](MDNSAPI.md#deletemdns) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mdns/{mdnsId} | Delete an exist mDNS rule
[**GetMdnsGrid**](MDNSAPI.md#getmdnsgrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mdns | Get mDNS rule list
[**ModifyMdns**](MDNSAPI.md#modifymdns) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mdns/{mdnsId} | Modify an exist mDNS rule



## CreateMdns

> OperationResponseResIdOpenApiVO CreateMdns(ctx, omadacId, siteId).CreateMdnsRuleOpenApiVO(createMdnsRuleOpenApiVO).Execute()

Create a new mDNS rule



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
	createMdnsRuleOpenApiVO := *openapiclient.NewCreateMdnsRuleOpenApiVO("Name_example", []string{"ProfileIds_example"}, false) // CreateMdnsRuleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDNSAPI.CreateMdns(context.Background(), omadacId, siteId).CreateMdnsRuleOpenApiVO(createMdnsRuleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDNSAPI.CreateMdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMdns`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MDNSAPI.CreateMdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMdnsRuleOpenApiVO** | [**CreateMdnsRuleOpenApiVO**](CreateMdnsRuleOpenApiVO.md) |  | 

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


## DeleteMdns

> OperationResponseString DeleteMdns(ctx, omadacId, siteId, mdnsId).Execute()

Delete an exist mDNS rule



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
	mdnsId := "mdnsId_example" // string | mDNS rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDNSAPI.DeleteMdns(context.Background(), omadacId, siteId, mdnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDNSAPI.DeleteMdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMdns`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `MDNSAPI.DeleteMdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mdnsId** | **string** | mDNS rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseString**](OperationResponseString.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMdnsGrid

> OperationResponseGridVOMdnsRuleOpenApiVO GetMdnsGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get mDNS rule list



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
	resp, r, err := apiClient.MDNSAPI.GetMdnsGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDNSAPI.GetMdnsGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMdnsGrid`: OperationResponseGridVOMdnsRuleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MDNSAPI.GetMdnsGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMdnsGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOMdnsRuleOpenApiVO**](OperationResponseGridVOMdnsRuleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMdns

> OperationResponseWithoutResult ModifyMdns(ctx, omadacId, siteId, mdnsId).CreateMdnsRuleOpenApiVO(createMdnsRuleOpenApiVO).Execute()

Modify an exist mDNS rule



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
	mdnsId := "mdnsId_example" // string | mDNS rule ID
	createMdnsRuleOpenApiVO := *openapiclient.NewCreateMdnsRuleOpenApiVO("Name_example", []string{"ProfileIds_example"}, false) // CreateMdnsRuleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MDNSAPI.ModifyMdns(context.Background(), omadacId, siteId, mdnsId).CreateMdnsRuleOpenApiVO(createMdnsRuleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MDNSAPI.ModifyMdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMdns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MDNSAPI.ModifyMdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mdnsId** | **string** | mDNS rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createMdnsRuleOpenApiVO** | [**CreateMdnsRuleOpenApiVO**](CreateMdnsRuleOpenApiVO.md) |  | 

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


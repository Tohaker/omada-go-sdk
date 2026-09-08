# ServiceTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIgmpTemplate**](ServiceTemplateAPI.md#getigmptemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv/igmp | Get IGMP template setting
[**ModifyIgmpTemplate**](ServiceTemplateAPI.md#modifyigmptemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv/igmp | Modify IGMP template setting



## GetIgmpTemplate

> OperationResponseIgmpTemplateOpenApiVO GetIgmpTemplate(ctx, omadacId, siteTemplateId).Execute()

Get IGMP template setting



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.GetIgmpTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetIgmpTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIgmpTemplate`: OperationResponseIgmpTemplateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetIgmpTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIgmpTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIgmpTemplateOpenApiVO**](OperationResponseIgmpTemplateOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIgmpTemplate

> OperationResponseWithoutResult ModifyIgmpTemplate(ctx, omadacId, siteTemplateId).IgmpTemplateOpenApiVO(igmpTemplateOpenApiVO).Execute()

Modify IGMP template setting



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
	igmpTemplateOpenApiVO := *openapiclient.NewIgmpTemplateOpenApiVO(false, int32(123), "WanPortId_example") // IgmpTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyIgmpTemplate(context.Background(), omadacId, siteTemplateId).IgmpTemplateOpenApiVO(igmpTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyIgmpTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIgmpTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyIgmpTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIgmpTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **igmpTemplateOpenApiVO** | [**IgmpTemplateOpenApiVO**](IgmpTemplateOpenApiVO.md) |  | 

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


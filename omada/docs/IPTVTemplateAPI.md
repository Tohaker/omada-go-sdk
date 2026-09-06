# IPTVTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIptvTemplate**](IPTVTemplateAPI.md#getiptvtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv/iptv | Get IPTV template setting
[**GetIptvTemplateServerSetting**](IPTVTemplateAPI.md#getiptvtemplateserversetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv | Get IPTV setting template
[**ModifyIptvTemplate**](IPTVTemplateAPI.md#modifyiptvtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv/iptv | Modify IPTV template setting
[**ModifyIptvTemplateServerSetting**](IPTVTemplateAPI.md#modifyiptvtemplateserversetting) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv | Modify IPTV setting template



## GetIptvTemplate

> OperationResponseIptvOpenApiVO GetIptvTemplate(ctx, omadacId, siteTemplateId).Execute()

Get IPTV template setting



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
	resp, r, err := apiClient.IPTVTemplateAPI.GetIptvTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVTemplateAPI.GetIptvTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIptvTemplate`: OperationResponseIptvOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IPTVTemplateAPI.GetIptvTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIptvTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIptvOpenApiVO**](OperationResponseIptvOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIptvTemplateServerSetting

> OperationResponseIptvSettingOpenApiVO GetIptvTemplateServerSetting(ctx, omadacId, siteTemplateId).Execute()

Get IPTV setting template



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
	resp, r, err := apiClient.IPTVTemplateAPI.GetIptvTemplateServerSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVTemplateAPI.GetIptvTemplateServerSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIptvTemplateServerSetting`: OperationResponseIptvSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IPTVTemplateAPI.GetIptvTemplateServerSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIptvTemplateServerSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIptvSettingOpenApiVO**](OperationResponseIptvSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIptvTemplate

> OperationResponseWithoutResult ModifyIptvTemplate(ctx, omadacId, siteTemplateId).IptvConfigOpenApiVO(iptvConfigOpenApiVO).Execute()

Modify IPTV template setting



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
	iptvConfigOpenApiVO := *openapiclient.NewIptvConfigOpenApiVO(false, int32(123), []openapiclient.IptvPortConfigOpenApiVO{*openapiclient.NewIptvPortConfigOpenApiVO("PortId_example", int32(123))}, "WanPortId_example") // IptvConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPTVTemplateAPI.ModifyIptvTemplate(context.Background(), omadacId, siteTemplateId).IptvConfigOpenApiVO(iptvConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVTemplateAPI.ModifyIptvTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIptvTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPTVTemplateAPI.ModifyIptvTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIptvTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iptvConfigOpenApiVO** | [**IptvConfigOpenApiVO**](IptvConfigOpenApiVO.md) |  | 

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


## ModifyIptvTemplateServerSetting

> OperationResponseWithoutResult ModifyIptvTemplateServerSetting(ctx, omadacId, siteTemplateId).IptvSettingOpenApiVO(iptvSettingOpenApiVO).Execute()

Modify IPTV setting template



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
	iptvSettingOpenApiVO := *openapiclient.NewIptvSettingOpenApiVO(*openapiclient.NewIgmpOpenApiVO(false, int32(123))) // IptvSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPTVTemplateAPI.ModifyIptvTemplateServerSetting(context.Background(), omadacId, siteTemplateId).IptvSettingOpenApiVO(iptvSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVTemplateAPI.ModifyIptvTemplateServerSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIptvTemplateServerSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPTVTemplateAPI.ModifyIptvTemplateServerSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIptvTemplateServerSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iptvSettingOpenApiVO** | [**IptvSettingOpenApiVO**](IptvSettingOpenApiVO.md) |  | 

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


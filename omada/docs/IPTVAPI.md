# IPTVAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIgmp**](IPTVAPI.md#getigmp) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/igmp | Get the IGMP settings
[**GetIptv**](IPTVAPI.md#getiptv) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/iptv | Get IPTV setting
[**GetMld**](IPTVAPI.md#getmld) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mld | Get the MLD settings
[**ModifyIgmp**](IPTVAPI.md#modifyigmp) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/igmp | Modify the IGMP settings
[**ModifyIptv**](IPTVAPI.md#modifyiptv) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/iptv | Modify IPTV setting
[**ModifyMld**](IPTVAPI.md#modifymld) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mld | Modify the MLD settings



## GetIgmp

> OperationResponseIgmpOpenApiVO GetIgmp(ctx, omadacId, siteId).Execute()

Get the IGMP settings



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
	resp, r, err := apiClient.IPTVAPI.GetIgmp(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVAPI.GetIgmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIgmp`: OperationResponseIgmpOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IPTVAPI.GetIgmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIgmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIgmpOpenApiVO**](OperationResponseIgmpOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIptv

> OperationResponseIptvOpenApiVO GetIptv(ctx, omadacId, siteId).Execute()

Get IPTV setting



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
	resp, r, err := apiClient.IPTVAPI.GetIptv(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVAPI.GetIptv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIptv`: OperationResponseIptvOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IPTVAPI.GetIptv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIptvRequest struct via the builder pattern


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


## GetMld

> OperationResponseMldOpenApiVO GetMld(ctx, omadacId, siteId).Execute()

Get the MLD settings



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
	resp, r, err := apiClient.IPTVAPI.GetMld(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVAPI.GetMld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMld`: OperationResponseMldOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `IPTVAPI.GetMld`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMldOpenApiVO**](OperationResponseMldOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIgmp

> OperationResponseWithoutResult ModifyIgmp(ctx, omadacId, siteId).IgmpOpenApiVO(igmpOpenApiVO).Execute()

Modify the IGMP settings



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
	igmpOpenApiVO := *openapiclient.NewIgmpOpenApiVO(false, int32(123)) // IgmpOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPTVAPI.ModifyIgmp(context.Background(), omadacId, siteId).IgmpOpenApiVO(igmpOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVAPI.ModifyIgmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIgmp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPTVAPI.ModifyIgmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIgmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **igmpOpenApiVO** | [**IgmpOpenApiVO**](IgmpOpenApiVO.md) |  | 

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


## ModifyIptv

> OperationResponseWithoutResult ModifyIptv(ctx, omadacId, siteId).IptvConfigOpenApiVO(iptvConfigOpenApiVO).Execute()

Modify IPTV setting



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
	iptvConfigOpenApiVO := *openapiclient.NewIptvConfigOpenApiVO(false, int32(123), []openapiclient.IptvPortConfigOpenApiVO{*openapiclient.NewIptvPortConfigOpenApiVO("PortId_example", int32(123))}, "WanPortId_example") // IptvConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPTVAPI.ModifyIptv(context.Background(), omadacId, siteId).IptvConfigOpenApiVO(iptvConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVAPI.ModifyIptv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIptv`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPTVAPI.ModifyIptv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIptvRequest struct via the builder pattern


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


## ModifyMld

> OperationResponseWithoutResult ModifyMld(ctx, omadacId, siteId).MldOpenApiVO(mldOpenApiVO).Execute()

Modify the MLD settings



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
	mldOpenApiVO := *openapiclient.NewMldOpenApiVO(false, int32(123)) // MldOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPTVAPI.ModifyMld(context.Background(), omadacId, siteId).MldOpenApiVO(mldOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPTVAPI.ModifyMld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMld`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPTVAPI.ModifyMld`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mldOpenApiVO** | [**MldOpenApiVO**](MldOpenApiVO.md) |  | 

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


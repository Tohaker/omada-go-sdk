# DNSAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearDnsCacheList**](DNSAPI.md#cleardnscachelist) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/dns-cache-data/clear | Clear DNS cache list
[**CreateDdns**](DNSAPI.md#createddns) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/ddns | Create a new Dynamic DNS entry
[**DeleteDdns**](DNSAPI.md#deleteddns) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/ddns/{ddnsId} | Delete an exist Dynamic DNS entry
[**GetDdnsGrid**](DNSAPI.md#getddnsgrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/ddns | Get Dynamic DNS list
[**GetDnsCacheList**](DNSAPI.md#getdnscachelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/dns-cache-data | Get DNS cache list
[**GetDnsCacheSetting**](DNSAPI.md#getdnscachesetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dns-cache | Get DNS cache
[**GetDnsProxy**](DNSAPI.md#getdnsproxy) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dns-proxy | Get DNS proxy setting
[**ModifyDdns**](DNSAPI.md#modifyddns) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/ddns/{ddnsId} | Modify an exist Dynamic DNS entry
[**ModifyDnsCacheSetting**](DNSAPI.md#modifydnscachesetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dns-cache | Modify DNS cache setting
[**ModifyDnsProxy**](DNSAPI.md#modifydnsproxy) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dns-proxy | Modify DNS proxy setting



## ClearDnsCacheList

> OperationResponseWithoutResult ClearDnsCacheList(ctx, omadacId, siteId).Execute()

Clear DNS cache list



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
	resp, r, err := apiClient.DNSAPI.ClearDnsCacheList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.ClearDnsCacheList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearDnsCacheList`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.ClearDnsCacheList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearDnsCacheListRequest struct via the builder pattern


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


## CreateDdns

> OperationResponseResIdOpenApiVO CreateDdns(ctx, omadacId, siteId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()

Create a new Dynamic DNS entry



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
	createDdnsOpenApiVO := *openapiclient.NewCreateDdnsOpenApiVO("InterfacePortId_example", "Password_example", int32(123), false, "Username_example") // CreateDdnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.CreateDdns(context.Background(), omadacId, siteId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.CreateDdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDdns`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.CreateDdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDdnsOpenApiVO** | [**CreateDdnsOpenApiVO**](CreateDdnsOpenApiVO.md) |  | 

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


## DeleteDdns

> OperationResponseWithoutResult DeleteDdns(ctx, omadacId, siteId, ddnsId).Execute()

Delete an exist Dynamic DNS entry



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
	ddnsId := "ddnsId_example" // string | Dynamic DNS entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.DeleteDdns(context.Background(), omadacId, siteId, ddnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.DeleteDdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDdns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.DeleteDdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ddnsId** | **string** | Dynamic DNS entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDdnsRequest struct via the builder pattern


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


## GetDdnsGrid

> OperationResponseDdnsOpenApiGridVODdnsOpenApiVO GetDdnsGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsService(sortsService).SortsUpdateInterval(sortsUpdateInterval).SortsStatus(sortsStatus).Execute()

Get Dynamic DNS list



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
	sortsService := "sortsService_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsUpdateInterval := "sortsUpdateInterval_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.GetDdnsGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsService(sortsService).SortsUpdateInterval(sortsUpdateInterval).SortsStatus(sortsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.GetDdnsGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDdnsGrid`: OperationResponseDdnsOpenApiGridVODdnsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.GetDdnsGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDdnsGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsService** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsUpdateInterval** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStatus** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseDdnsOpenApiGridVODdnsOpenApiVO**](OperationResponseDdnsOpenApiGridVODdnsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsCacheList

> OperationResponseDnsCacheOpenApiVO GetDnsCacheList(ctx, omadacId, siteId).DnsCacheQueryVO(dnsCacheQueryVO).Execute()

Get DNS cache list



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
	dnsCacheQueryVO := *openapiclient.NewDnsCacheQueryOpenApiV2VO() // DnsCacheQueryOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.GetDnsCacheList(context.Background(), omadacId, siteId).DnsCacheQueryVO(dnsCacheQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.GetDnsCacheList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsCacheList`: OperationResponseDnsCacheOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.GetDnsCacheList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsCacheListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsCacheQueryVO** | [**DnsCacheQueryOpenApiV2VO**](DnsCacheQueryOpenApiV2VO.md) |  | 

### Return type

[**OperationResponseDnsCacheOpenApiVO**](OperationResponseDnsCacheOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsCacheSetting

> OperationResponseDnsCacheOpenApiVO GetDnsCacheSetting(ctx, omadacId, siteId).Execute()

Get DNS cache



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
	resp, r, err := apiClient.DNSAPI.GetDnsCacheSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.GetDnsCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsCacheSetting`: OperationResponseDnsCacheOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.GetDnsCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDnsCacheOpenApiVO**](OperationResponseDnsCacheOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsProxy

> OperationResponseDnsProxySettingQueryOpenApiVO GetDnsProxy(ctx, omadacId, siteId).Execute()

Get DNS proxy setting



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
	resp, r, err := apiClient.DNSAPI.GetDnsProxy(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.GetDnsProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsProxy`: OperationResponseDnsProxySettingQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.GetDnsProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDnsProxySettingQueryOpenApiVO**](OperationResponseDnsProxySettingQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDdns

> OperationResponseWithoutResult ModifyDdns(ctx, omadacId, siteId, ddnsId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()

Modify an exist Dynamic DNS entry



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
	ddnsId := "ddnsId_example" // string | Dynamic DNS entry ID
	createDdnsOpenApiVO := *openapiclient.NewCreateDdnsOpenApiVO("InterfacePortId_example", "Password_example", int32(123), false, "Username_example") // CreateDdnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.ModifyDdns(context.Background(), omadacId, siteId, ddnsId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.ModifyDdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDdns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.ModifyDdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ddnsId** | **string** | Dynamic DNS entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createDdnsOpenApiVO** | [**CreateDdnsOpenApiVO**](CreateDdnsOpenApiVO.md) |  | 

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


## ModifyDnsCacheSetting

> OperationResponseWithoutResult ModifyDnsCacheSetting(ctx, omadacId, siteId).DnsCacheOpenApiVO(dnsCacheOpenApiVO).Execute()

Modify DNS cache setting



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
	dnsCacheOpenApiVO := *openapiclient.NewDnsCacheOpenApiVO() // DnsCacheOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.ModifyDnsCacheSetting(context.Background(), omadacId, siteId).DnsCacheOpenApiVO(dnsCacheOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.ModifyDnsCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDnsCacheSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.ModifyDnsCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDnsCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsCacheOpenApiVO** | [**DnsCacheOpenApiVO**](DnsCacheOpenApiVO.md) |  | 

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


## ModifyDnsProxy

> OperationResponseWithoutResult ModifyDnsProxy(ctx, omadacId, siteId).DnsProxySettingOpenApiVO(dnsProxySettingOpenApiVO).Execute()

Modify DNS proxy setting



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
	dnsProxySettingOpenApiVO := *openapiclient.NewDnsProxySettingOpenApiVO(false) // DnsProxySettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.ModifyDnsProxy(context.Background(), omadacId, siteId).DnsProxySettingOpenApiVO(dnsProxySettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.ModifyDnsProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDnsProxy`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.ModifyDnsProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDnsProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsProxySettingOpenApiVO** | [**DnsProxySettingOpenApiVO**](DnsProxySettingOpenApiVO.md) |  | 

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


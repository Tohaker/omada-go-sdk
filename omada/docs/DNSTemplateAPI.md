# DNSTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDdnsTemplate**](DNSTemplateAPI.md#createddnstemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns | Create a new Dynamic DNS template entry
[**DeleteDdnsTemplate**](DNSTemplateAPI.md#deleteddnstemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns/{ddnsId} | Delete an exist Dynamic DNS template entry
[**GetDdnsGridTemplate**](DNSTemplateAPI.md#getddnsgridtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns | Get Dynamic DNS template list
[**GetDnsCacheSettingTemplate**](DNSTemplateAPI.md#getdnscachesettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/dns-cache | Get DNS cache setting template
[**GetDnsProxyTemplate**](DNSTemplateAPI.md#getdnsproxytemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/dns-proxy | Get DNS proxy template setting
[**GetGlobalDdnsUpdateUrlTemplate**](DNSTemplateAPI.md#getglobalddnsupdateurltemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns/update-url | Get DDNS update url
[**ModifyDdnsTemplate**](DNSTemplateAPI.md#modifyddnstemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns/{ddnsId} | Modify an exist Dynamic DNS template entry
[**ModifyDnsCacheSettingTemplate**](DNSTemplateAPI.md#modifydnscachesettingtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/dns-cache | Modify DNS cache setting
[**ModifyDnsProxyTemplate**](DNSTemplateAPI.md#modifydnsproxytemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/dns-proxy | Modify DNS proxy template setting
[**RefreshDnsCacheListTemplate**](DNSTemplateAPI.md#refreshdnscachelisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/dns-cache-data/refresh | Refresh dns cache list



## CreateDdnsTemplate

> OperationResponseResIdOpenApiVO CreateDdnsTemplate(ctx, omadacId, siteTemplateId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()

Create a new Dynamic DNS template entry



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
	createDdnsOpenApiVO := *openapiclient.NewCreateDdnsOpenApiVO("InterfacePortId_example", "Password_example", int32(123), false, "Username_example") // CreateDdnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSTemplateAPI.CreateDdnsTemplate(context.Background(), omadacId, siteTemplateId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.CreateDdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDdnsTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.CreateDdnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDdnsTemplateRequest struct via the builder pattern


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


## DeleteDdnsTemplate

> OperationResponseWithoutResult DeleteDdnsTemplate(ctx, omadacId, siteTemplateId, ddnsId).Execute()

Delete an exist Dynamic DNS template entry



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
	ddnsId := "ddnsId_example" // string | Dynamic DNS entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSTemplateAPI.DeleteDdnsTemplate(context.Background(), omadacId, siteTemplateId, ddnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.DeleteDdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDdnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.DeleteDdnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ddnsId** | **string** | Dynamic DNS entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDdnsTemplateRequest struct via the builder pattern


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


## GetDdnsGridTemplate

> OperationResponseDdnsOpenApiGridVODdnsOpenApiVO GetDdnsGridTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SortsService(sortsService).SortsUpdateInterval(sortsUpdateInterval).SortsStatus(sortsStatus).Execute()

Get Dynamic DNS template list



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
	sortsService := "sortsService_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsUpdateInterval := "sortsUpdateInterval_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSTemplateAPI.GetDdnsGridTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SortsService(sortsService).SortsUpdateInterval(sortsUpdateInterval).SortsStatus(sortsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.GetDdnsGridTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDdnsGridTemplate`: OperationResponseDdnsOpenApiGridVODdnsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.GetDdnsGridTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDdnsGridTemplateRequest struct via the builder pattern


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


## GetDnsCacheSettingTemplate

> OperationResponseDnsCacheOpenApiVO GetDnsCacheSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get DNS cache setting template



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
	resp, r, err := apiClient.DNSTemplateAPI.GetDnsCacheSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.GetDnsCacheSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsCacheSettingTemplate`: OperationResponseDnsCacheOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.GetDnsCacheSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsCacheSettingTemplateRequest struct via the builder pattern


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


## GetDnsProxyTemplate

> OperationResponseDnsProxySettingQueryOpenApiVO GetDnsProxyTemplate(ctx, omadacId, siteTemplateId).Execute()

Get DNS proxy template setting



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
	resp, r, err := apiClient.DNSTemplateAPI.GetDnsProxyTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.GetDnsProxyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsProxyTemplate`: OperationResponseDnsProxySettingQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.GetDnsProxyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsProxyTemplateRequest struct via the builder pattern


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


## GetGlobalDdnsUpdateUrlTemplate

> OperationResponseString GetGlobalDdnsUpdateUrlTemplate(ctx, omadacId, siteTemplateId).Execute()

Get DDNS update url



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
	resp, r, err := apiClient.DNSTemplateAPI.GetGlobalDdnsUpdateUrlTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.GetGlobalDdnsUpdateUrlTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalDdnsUpdateUrlTemplate`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.GetGlobalDdnsUpdateUrlTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalDdnsUpdateUrlTemplateRequest struct via the builder pattern


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


## ModifyDdnsTemplate

> OperationResponseWithoutResult ModifyDdnsTemplate(ctx, omadacId, siteTemplateId, ddnsId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()

Modify an exist Dynamic DNS template entry



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
	ddnsId := "ddnsId_example" // string | Dynamic DNS entry ID
	createDdnsOpenApiVO := *openapiclient.NewCreateDdnsOpenApiVO("InterfacePortId_example", "Password_example", int32(123), false, "Username_example") // CreateDdnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSTemplateAPI.ModifyDdnsTemplate(context.Background(), omadacId, siteTemplateId, ddnsId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.ModifyDdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDdnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.ModifyDdnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ddnsId** | **string** | Dynamic DNS entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDdnsTemplateRequest struct via the builder pattern


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


## ModifyDnsCacheSettingTemplate

> OperationResponseWithoutResult ModifyDnsCacheSettingTemplate(ctx, omadacId, siteTemplateId).DnsCacheOpenApiVO(dnsCacheOpenApiVO).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	dnsCacheOpenApiVO := *openapiclient.NewDnsCacheOpenApiVO() // DnsCacheOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSTemplateAPI.ModifyDnsCacheSettingTemplate(context.Background(), omadacId, siteTemplateId).DnsCacheOpenApiVO(dnsCacheOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.ModifyDnsCacheSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDnsCacheSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.ModifyDnsCacheSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDnsCacheSettingTemplateRequest struct via the builder pattern


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


## ModifyDnsProxyTemplate

> OperationResponseWithoutResult ModifyDnsProxyTemplate(ctx, omadacId, siteTemplateId).DnsProxySettingOpenApiVO(dnsProxySettingOpenApiVO).Execute()

Modify DNS proxy template setting



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
	dnsProxySettingOpenApiVO := *openapiclient.NewDnsProxySettingOpenApiVO(false) // DnsProxySettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSTemplateAPI.ModifyDnsProxyTemplate(context.Background(), omadacId, siteTemplateId).DnsProxySettingOpenApiVO(dnsProxySettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.ModifyDnsProxyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDnsProxyTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.ModifyDnsProxyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDnsProxyTemplateRequest struct via the builder pattern


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


## RefreshDnsCacheListTemplate

> OperationResponseGridVODnsCacheInfoVO RefreshDnsCacheListTemplate(ctx, omadacId, siteTemplateId).Vo(vo).Execute()

Refresh dns cache list



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
	vo := *openapiclient.NewDnsCacheQueryOpenApiVO() // DnsCacheQueryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSTemplateAPI.RefreshDnsCacheListTemplate(context.Background(), omadacId, siteTemplateId).Vo(vo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSTemplateAPI.RefreshDnsCacheListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshDnsCacheListTemplate`: OperationResponseGridVODnsCacheInfoVO
	fmt.Fprintf(os.Stdout, "Response from `DNSTemplateAPI.RefreshDnsCacheListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshDnsCacheListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vo** | [**DnsCacheQueryOpenApiVO**](DnsCacheQueryOpenApiVO.md) |  | 

### Return type

[**OperationResponseGridVODnsCacheInfoVO**](OperationResponseGridVODnsCacheInfoVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ServiceTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPortScheduleTemplate**](ServiceTemplateAPI.md#AddPortScheduleTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/port-schedules | Create a new Port Schedule Template
[**CreateDdnsTemplate**](ServiceTemplateAPI.md#CreateDdnsTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns | Create a new Dynamic DNS template entry
[**CreateMdnsTemplate**](ServiceTemplateAPI.md#CreateMdnsTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/mdns | Create a new mDNS template rule
[**DeleteDdnsTemplate**](ServiceTemplateAPI.md#DeleteDdnsTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns/{ddnsId} | Delete an exist Dynamic DNS template entry
[**DeleteMdnsTemplate**](ServiceTemplateAPI.md#DeleteMdnsTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/mdns/{mdnsId} | Delete an exist mDNS template rule
[**GetDdnsGridTemplate**](ServiceTemplateAPI.md#GetDdnsGridTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns | Get Dynamic DNS template list
[**GetDnsCacheSettingTemplate**](ServiceTemplateAPI.md#GetDnsCacheSettingTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/dns-cache | Get DNS cache setting template
[**GetDnsProxyTemplate**](ServiceTemplateAPI.md#GetDnsProxyTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/dns-proxy | Get DNS proxy template setting
[**GetGlobalDdnsUpdateUrlTemplate**](ServiceTemplateAPI.md#GetGlobalDdnsUpdateUrlTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns/update-url | Get DDNS update url
[**GetIgmpTemplate**](ServiceTemplateAPI.md#GetIgmpTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv/igmp | Get IGMP template setting
[**GetIptvTemplate**](ServiceTemplateAPI.md#GetIptvTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv/iptv | Get IPTV template setting
[**GetIptvTemplateServerSetting**](ServiceTemplateAPI.md#GetIptvTemplateServerSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv | Get IPTV setting template
[**GetMdnsGridTemplate**](ServiceTemplateAPI.md#GetMdnsGridTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/mdns | Get mDNS template rule list
[**GetPortScheduleListTemplate**](ServiceTemplateAPI.md#GetPortScheduleListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/port-schedules | Get port Schedule list
[**GetRebootScheduleList**](ServiceTemplateAPI.md#GetRebootScheduleList) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/rebootSchedules | Get reboot schedule templates
[**GetSnmpSettingTemplate**](ServiceTemplateAPI.md#GetSnmpSettingTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/snmp | Get SNMP template setting
[**GetSshSettingTemplate**](ServiceTemplateAPI.md#GetSshSettingTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ssh | Get SSH template setting
[**GetUpnpSettingTemplate**](ServiceTemplateAPI.md#GetUpnpSettingTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/upnp | Get UPnP template setting
[**ModifyDdnsTemplate**](ServiceTemplateAPI.md#ModifyDdnsTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/ddns/{ddnsId} | Modify an exist Dynamic DNS template entry
[**ModifyDnsCacheSettingTemplate**](ServiceTemplateAPI.md#ModifyDnsCacheSettingTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/dns-cache | Modify DNS cache setting
[**ModifyDnsProxyTemplate**](ServiceTemplateAPI.md#ModifyDnsProxyTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/dns-proxy | Modify DNS proxy template setting
[**ModifyIgmpTemplate**](ServiceTemplateAPI.md#ModifyIgmpTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv/igmp | Modify IGMP template setting
[**ModifyIptvTemplate**](ServiceTemplateAPI.md#ModifyIptvTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv/iptv | Modify IPTV template setting
[**ModifyIptvTemplateServerSetting**](ServiceTemplateAPI.md#ModifyIptvTemplateServerSetting) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/iptv | Modify IPTV setting template
[**ModifyMdnsTemplate**](ServiceTemplateAPI.md#ModifyMdnsTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/mdns/{mdnsId} | Modify an exist mDNS template rule
[**ModifyPortScheduleTemplate**](ServiceTemplateAPI.md#ModifyPortScheduleTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/port-schedules/{type}/{portScheduleId} | Modify a Port Schedule Template
[**ModifyRebootScheduleTemplate**](ServiceTemplateAPI.md#ModifyRebootScheduleTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/rebootSchedules/{rebootScheduleId} | Modify reboot schedule template
[**ModifySnmpSettingTemplate**](ServiceTemplateAPI.md#ModifySnmpSettingTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/snmp | Modify SNMP template setting
[**RefreshDnsCacheListTemplate**](ServiceTemplateAPI.md#RefreshDnsCacheListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/dns-cache-data/refresh | Refresh dns cache list
[**RemovePortScheduleTemplate**](ServiceTemplateAPI.md#RemovePortScheduleTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/port-schedules/{type}/{portScheduleId} | Delete Port Schedule Template
[**RemoveRebootScheduleTemplate**](ServiceTemplateAPI.md#RemoveRebootScheduleTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/rebootSchedules/{rebootScheduleId} | Remove reboot schedule template
[**UpdateSshSettingTemplate**](ServiceTemplateAPI.md#UpdateSshSettingTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ssh | Modify SSH template setting
[**UpdateUpnpSettingTemplate**](ServiceTemplateAPI.md#UpdateUpnpSettingTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/upnp | Modify UPnP template setting



## AddPortScheduleTemplate

> OperationResponse AddPortScheduleTemplate(ctx, omadacId, siteTemplateId).PortOrPoeScheduleOpenApiVO(portOrPoeScheduleOpenApiVO).Execute()

Create a new Port Schedule Template



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
	portOrPoeScheduleOpenApiVO := *openapiclient.NewPortOrPoeScheduleOpenApiVO("Name_example", map[string][]int32{"key": []int32{int32(123)}}, int32(123), false, "TurnOnTime_example") // PortOrPoeScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.AddPortScheduleTemplate(context.Background(), omadacId, siteTemplateId).PortOrPoeScheduleOpenApiVO(portOrPoeScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.AddPortScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPortScheduleTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.AddPortScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPortScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portOrPoeScheduleOpenApiVO** | [**PortOrPoeScheduleOpenApiVO**](PortOrPoeScheduleOpenApiVO.md) |  | 

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
	resp, r, err := apiClient.ServiceTemplateAPI.CreateDdnsTemplate(context.Background(), omadacId, siteTemplateId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.CreateDdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDdnsTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.CreateDdnsTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.ServiceTemplateAPI.CreateMdnsTemplate(context.Background(), omadacId, siteTemplateId).CreateMdnsRuleTemplateOpenApiVO(createMdnsRuleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.CreateMdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMdnsTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.CreateMdnsTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.DeleteDdnsTemplate(context.Background(), omadacId, siteTemplateId, ddnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.DeleteDdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDdnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.DeleteDdnsTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.ServiceTemplateAPI.DeleteMdnsTemplate(context.Background(), omadacId, siteTemplateId, mdnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.DeleteMdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMdnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.DeleteMdnsTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.GetDdnsGridTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SortsService(sortsService).SortsUpdateInterval(sortsUpdateInterval).SortsStatus(sortsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetDdnsGridTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDdnsGridTemplate`: OperationResponseDdnsOpenApiGridVODdnsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetDdnsGridTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.GetDnsCacheSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetDnsCacheSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsCacheSettingTemplate`: OperationResponseDnsCacheOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetDnsCacheSettingTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.GetDnsProxyTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetDnsProxyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsProxyTemplate`: OperationResponseDnsProxySettingQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetDnsProxyTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.GetGlobalDdnsUpdateUrlTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetGlobalDdnsUpdateUrlTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalDdnsUpdateUrlTemplate`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetGlobalDdnsUpdateUrlTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.ServiceTemplateAPI.GetIptvTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetIptvTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIptvTemplate`: OperationResponseIptvOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetIptvTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.GetIptvTemplateServerSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetIptvTemplateServerSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIptvTemplateServerSetting`: OperationResponseIptvSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetIptvTemplateServerSetting`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.GetMdnsGridTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetMdnsGridTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMdnsGridTemplate`: OperationResponseGridVOMdnsRuleTemplateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetMdnsGridTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortScheduleListTemplate

> OperationResponseListPortOrPoeScheduleOpenApiVO GetPortScheduleListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get port Schedule list



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
	resp, r, err := apiClient.ServiceTemplateAPI.GetPortScheduleListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetPortScheduleListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortScheduleListTemplate`: OperationResponseListPortOrPoeScheduleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetPortScheduleListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortScheduleListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListPortOrPoeScheduleOpenApiVO**](OperationResponseListPortOrPoeScheduleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRebootScheduleList

> OperationResponseListRebootScheduleTemplateQueryOpenApiVO GetRebootScheduleList(ctx, omadacId, siteTemplateId).Execute()

Get reboot schedule templates



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
	resp, r, err := apiClient.ServiceTemplateAPI.GetRebootScheduleList(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetRebootScheduleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRebootScheduleList`: OperationResponseListRebootScheduleTemplateQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetRebootScheduleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRebootScheduleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRebootScheduleTemplateQueryOpenApiVO**](OperationResponseListRebootScheduleTemplateQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnmpSettingTemplate

> OperationResponseSnmpSettingOpenApiVO GetSnmpSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get SNMP template setting



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
	resp, r, err := apiClient.ServiceTemplateAPI.GetSnmpSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetSnmpSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnmpSettingTemplate`: OperationResponseSnmpSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetSnmpSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnmpSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSnmpSettingOpenApiVO**](OperationResponseSnmpSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshSettingTemplate

> OperationResponseSSHSetting GetSshSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get SSH template setting



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
	resp, r, err := apiClient.ServiceTemplateAPI.GetSshSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetSshSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshSettingTemplate`: OperationResponseSSHSetting
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetSshSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSSHSetting**](OperationResponseSSHSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpnpSettingTemplate

> OperationResponseUpnpSettingOpenApiVO GetUpnpSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get UPnP template setting



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
	resp, r, err := apiClient.ServiceTemplateAPI.GetUpnpSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.GetUpnpSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpnpSettingTemplate`: OperationResponseUpnpSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.GetUpnpSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpnpSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUpnpSettingOpenApiVO**](OperationResponseUpnpSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyDdnsTemplate(context.Background(), omadacId, siteTemplateId, ddnsId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyDdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDdnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyDdnsTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyDnsCacheSettingTemplate(context.Background(), omadacId, siteTemplateId).DnsCacheOpenApiVO(dnsCacheOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyDnsCacheSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDnsCacheSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyDnsCacheSettingTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyDnsProxyTemplate(context.Background(), omadacId, siteTemplateId).DnsProxySettingOpenApiVO(dnsProxySettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyDnsProxyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDnsProxyTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyDnsProxyTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyIptvTemplate(context.Background(), omadacId, siteTemplateId).IptvConfigOpenApiVO(iptvConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyIptvTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIptvTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyIptvTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyIptvTemplateServerSetting(context.Background(), omadacId, siteTemplateId).IptvSettingOpenApiVO(iptvSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyIptvTemplateServerSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIptvTemplateServerSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyIptvTemplateServerSetting`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyMdnsTemplate(context.Background(), omadacId, siteTemplateId, mdnsId).CreateMdnsRuleTemplateOpenApiVO(createMdnsRuleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyMdnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMdnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyMdnsTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPortScheduleTemplate

> OperationResponseWithoutResult ModifyPortScheduleTemplate(ctx, omadacId, siteTemplateId, type_, portScheduleId).PortOrPoeScheduleOpenApiVO(portOrPoeScheduleOpenApiVO).Execute()

Modify a Port Schedule Template



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
	type_ := "type__example" // string | 
	portScheduleId := "portScheduleId_example" // string | 
	portOrPoeScheduleOpenApiVO := *openapiclient.NewPortOrPoeScheduleOpenApiVO("Name_example", map[string][]int32{"key": []int32{int32(123)}}, int32(123), false, "TurnOnTime_example") // PortOrPoeScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyPortScheduleTemplate(context.Background(), omadacId, siteTemplateId, type_, portScheduleId).PortOrPoeScheduleOpenApiVO(portOrPoeScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyPortScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortScheduleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyPortScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**type_** | **string** |  | 
**portScheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **portOrPoeScheduleOpenApiVO** | [**PortOrPoeScheduleOpenApiVO**](PortOrPoeScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRebootScheduleTemplate

> OperationResponseWithoutResult ModifyRebootScheduleTemplate(ctx, omadacId, siteTemplateId, rebootScheduleId).Page(page).PageSize(pageSize).RebootScheduleOpenApiVO(rebootScheduleOpenApiVO).Execute()

Modify reboot schedule template



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
	rebootScheduleId := "rebootScheduleId_example" // string | Reboot Schedule ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	rebootScheduleOpenApiVO := *openapiclient.NewRebootScheduleOpenApiVO([]string{"DeviceMacs_example"}, "Name_example", false, *openapiclient.NewRebootScheduleTimeOpenApiVO(int32(123), int32(123), int32(123))) // RebootScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.ModifyRebootScheduleTemplate(context.Background(), omadacId, siteTemplateId, rebootScheduleId).Page(page).PageSize(pageSize).RebootScheduleOpenApiVO(rebootScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifyRebootScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRebootScheduleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifyRebootScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**rebootScheduleId** | **string** | Reboot Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRebootScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **rebootScheduleOpenApiVO** | [**RebootScheduleOpenApiVO**](RebootScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySnmpSettingTemplate

> OperationResponseWithoutResult ModifySnmpSettingTemplate(ctx, omadacId, siteTemplateId).SnmpSettingOpenApiVO(snmpSettingOpenApiVO).Execute()

Modify SNMP template setting



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
	snmpSettingOpenApiVO := *openapiclient.NewSnmpSettingOpenApiVO(false, false) // SnmpSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.ModifySnmpSettingTemplate(context.Background(), omadacId, siteTemplateId).SnmpSettingOpenApiVO(snmpSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.ModifySnmpSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySnmpSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.ModifySnmpSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySnmpSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **snmpSettingOpenApiVO** | [**SnmpSettingOpenApiVO**](SnmpSettingOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

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
	resp, r, err := apiClient.ServiceTemplateAPI.RefreshDnsCacheListTemplate(context.Background(), omadacId, siteTemplateId).Vo(vo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.RefreshDnsCacheListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshDnsCacheListTemplate`: OperationResponseGridVODnsCacheInfoVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.RefreshDnsCacheListTemplate`: %v\n", resp)
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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePortScheduleTemplate

> OperationResponseWithoutResult RemovePortScheduleTemplate(ctx, omadacId, siteTemplateId, type_, portScheduleId).Execute()

Delete Port Schedule Template



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
	type_ := "type__example" // string | 
	portScheduleId := "portScheduleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.RemovePortScheduleTemplate(context.Background(), omadacId, siteTemplateId, type_, portScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.RemovePortScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePortScheduleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.RemovePortScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**type_** | **string** |  | 
**portScheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePortScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRebootScheduleTemplate

> OperationResponseWithoutResult RemoveRebootScheduleTemplate(ctx, omadacId, siteTemplateId, rebootScheduleId).Execute()

Remove reboot schedule template



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
	rebootScheduleId := "rebootScheduleId_example" // string | Reboot Schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.RemoveRebootScheduleTemplate(context.Background(), omadacId, siteTemplateId, rebootScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.RemoveRebootScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveRebootScheduleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.RemoveRebootScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**rebootScheduleId** | **string** | Reboot Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRebootScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSshSettingTemplate

> OperationResponseWithoutResult UpdateSshSettingTemplate(ctx, omadacId, siteTemplateId).SSHSetting(sSHSetting).Execute()

Modify SSH template setting



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
	sSHSetting := *openapiclient.NewSSHSetting(false, int32(123)) // SSHSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.UpdateSshSettingTemplate(context.Background(), omadacId, siteTemplateId).SSHSetting(sSHSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.UpdateSshSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSshSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.UpdateSshSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSshSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sSHSetting** | [**SSHSetting**](SSHSetting.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUpnpSettingTemplate

> OperationResponseWithoutResult UpdateUpnpSettingTemplate(ctx, omadacId, siteTemplateId).UpnpSettingOpenApiVO(upnpSettingOpenApiVO).Execute()

Modify UPnP template setting



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
	upnpSettingOpenApiVO := *openapiclient.NewUpnpSettingOpenApiVO(false) // UpnpSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTemplateAPI.UpdateUpnpSettingTemplate(context.Background(), omadacId, siteTemplateId).UpnpSettingOpenApiVO(upnpSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTemplateAPI.UpdateUpnpSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUpnpSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceTemplateAPI.UpdateUpnpSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUpnpSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upnpSettingOpenApiVO** | [**UpnpSettingOpenApiVO**](UpnpSettingOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


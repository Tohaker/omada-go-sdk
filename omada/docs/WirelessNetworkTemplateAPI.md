# \WirelessNetworkTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountWlansTemplate**](WirelessNetworkTemplateAPI.md#CountWlansTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/cmd/countWlans | Check if the number of WLAN groups is out of limit
[**CreateSsidTemplate**](WirelessNetworkTemplateAPI.md#CreateSsidTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids | Create new SSID template
[**CreateWlanGroupTemplate**](WirelessNetworkTemplateAPI.md#CreateWlanGroupTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans | Create new WLAN group template
[**DeleteSsidTemplate**](WirelessNetworkTemplateAPI.md#DeleteSsidTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId} | Delete an existing SSID template
[**DeleteWlanGroupTemplate**](WirelessNetworkTemplateAPI.md#DeleteWlanGroupTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId} | Delete an existing WLAN group template
[**GetGridWlanGroupTemplate**](WirelessNetworkTemplateAPI.md#GetGridWlanGroupTemplate) | **Get** /openapi/v2/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans | Get WLAN group template list paging query
[**GetSsidDetailTemplate**](WirelessNetworkTemplateAPI.md#GetSsidDetailTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId} | Get SSID template detail info
[**GetSsidListTemplate**](WirelessNetworkTemplateAPI.md#GetSsidListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids | Get SSID template list
[**GetWlanGroupListTemplate**](WirelessNetworkTemplateAPI.md#GetWlanGroupListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans | Get WLAN group template list
[**UpdateSsidBasicConfigTemplate**](WirelessNetworkTemplateAPI.md#UpdateSsidBasicConfigTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-basic-config | Update SSID template basic config
[**UpdateSsidDhcpOptionConfigTemplate**](WirelessNetworkTemplateAPI.md#UpdateSsidDhcpOptionConfigTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-dhcp-option | Update SSID template DHCP option 82 config
[**UpdateSsidHotspotV2SettingTemplate**](WirelessNetworkTemplateAPI.md#UpdateSsidHotspotV2SettingTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-hotspotv2 | Update SSID template Hotspot2.0 config
[**UpdateSsidMacFilterConfigTemplate**](WirelessNetworkTemplateAPI.md#UpdateSsidMacFilterConfigTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-mac-filter | Update SSID template mac filter config
[**UpdateSsidMultiCastConfigTemplate**](WirelessNetworkTemplateAPI.md#UpdateSsidMultiCastConfigTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-multicast-config | Update SSID template Multicast/Broadcast management config
[**UpdateSsidRateControlConfigTemplate**](WirelessNetworkTemplateAPI.md#UpdateSsidRateControlConfigTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-rate-control | Update SSID template 802.11 rate control config
[**UpdateSsidRateLimitConfigTemplate**](WirelessNetworkTemplateAPI.md#UpdateSsidRateLimitConfigTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-rate-limit | Update SSID template rate limit config
[**UpdateSsidWlanScheduleTemplate**](WirelessNetworkTemplateAPI.md#UpdateSsidWlanScheduleTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-wlan-schedule | Update SSID template WLAN schedule config
[**UpdateWlanGroupTemplate**](WirelessNetworkTemplateAPI.md#UpdateWlanGroupTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId} | Modify an existing WLAN group template



## CountWlansTemplate

> OperationResponseWlanGroupStatusOpenApiVO CountWlansTemplate(ctx, omadacId, siteTemplateId).Execute()

Check if the number of WLAN groups is out of limit



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
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.CountWlansTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.CountWlansTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWlansTemplate`: OperationResponseWlanGroupStatusOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.CountWlansTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | siteTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountWlansTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWlanGroupStatusOpenApiVO**](OperationResponseWlanGroupStatusOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSsidTemplate

> OperationResponse CreateSsidTemplate(ctx, omadacId, siteTemplateId, wlanId).CreateSsidOpenApiVO(createSsidOpenApiVO).Execute()

Create new SSID template



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
	wlanId := "wlanId_example" // string | WLAN ID
	createSsidOpenApiVO := *openapiclient.NewCreateSsidOpenApiVO(int32(123), false, int32(123), false, false, false, false, "Name_example", int32(123), int32(123), false) // CreateSsidOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.CreateSsidTemplate(context.Background(), omadacId, siteTemplateId, wlanId).CreateSsidOpenApiVO(createSsidOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.CreateSsidTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSsidTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.CreateSsidTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSsidTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createSsidOpenApiVO** | [**CreateSsidOpenApiVO**](CreateSsidOpenApiVO.md) |  | 

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


## CreateWlanGroupTemplate

> OperationResponse CreateWlanGroupTemplate(ctx, omadacId, siteTemplateId).CreateWlanGroupOpenApiVO(createWlanGroupOpenApiVO).Execute()

Create new WLAN group template



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
	createWlanGroupOpenApiVO := *openapiclient.NewCreateWlanGroupOpenApiVO(false, "Name_example") // CreateWlanGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.CreateWlanGroupTemplate(context.Background(), omadacId, siteTemplateId).CreateWlanGroupOpenApiVO(createWlanGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.CreateWlanGroupTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWlanGroupTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.CreateWlanGroupTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWlanGroupTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createWlanGroupOpenApiVO** | [**CreateWlanGroupOpenApiVO**](CreateWlanGroupOpenApiVO.md) |  | 

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


## DeleteSsidTemplate

> OperationResponseWithoutResult DeleteSsidTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).Execute()

Delete an existing SSID template



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.DeleteSsidTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.DeleteSsidTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSsidTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.DeleteSsidTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSsidTemplateRequest struct via the builder pattern


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


## DeleteWlanGroupTemplate

> OperationResponseWithoutResult DeleteWlanGroupTemplate(ctx, omadacId, siteTemplateId, wlanId).Execute()

Delete an existing WLAN group template



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
	wlanId := "wlanId_example" // string | WLAN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.DeleteWlanGroupTemplate(context.Background(), omadacId, siteTemplateId, wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.DeleteWlanGroupTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWlanGroupTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.DeleteWlanGroupTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWlanGroupTemplateRequest struct via the builder pattern


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


## GetGridWlanGroupTemplate

> OperationResponseWlanGroupGridOpenApiVO GetGridWlanGroupTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get WLAN group template list paging query



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.GetGridWlanGroupTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.GetGridWlanGroupTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridWlanGroupTemplate`: OperationResponseWlanGroupGridOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.GetGridWlanGroupTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridWlanGroupTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseWlanGroupGridOpenApiVO**](OperationResponseWlanGroupGridOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidDetailTemplate

> OperationResponseSsidDetailOpenApiVO GetSsidDetailTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).Execute()

Get SSID template detail info



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.GetSsidDetailTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.GetSsidDetailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidDetailTemplate`: OperationResponseSsidDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.GetSsidDetailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidDetailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseSsidDetailOpenApiVO**](OperationResponseSsidDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidListTemplate

> OperationResponseGridVOSsidOpenApiVO GetSsidListTemplate(ctx, omadacId, siteTemplateId, wlanId).Page(page).PageSize(pageSize).Execute()

Get SSID template list



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
	wlanId := "wlanId_example" // string | WLAN ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.GetSsidListTemplate(context.Background(), omadacId, siteTemplateId, wlanId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.GetSsidListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidListTemplate`: OperationResponseGridVOSsidOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.GetSsidListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOSsidOpenApiVO**](OperationResponseGridVOSsidOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWlanGroupListTemplate

> OperationResponseListWlanGroupOpenApiVO GetWlanGroupListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get WLAN group template list



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
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.GetWlanGroupListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.GetWlanGroupListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWlanGroupListTemplate`: OperationResponseListWlanGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.GetWlanGroupListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWlanGroupListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListWlanGroupOpenApiVO**](OperationResponseListWlanGroupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSsidBasicConfigTemplate

> OperationResponseWithoutResult UpdateSsidBasicConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidBasicConfigOpenApiVO(updateSsidBasicConfigOpenApiVO).Execute()

Update SSID template basic config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidBasicConfigOpenApiVO := *openapiclient.NewUpdateSsidBasicConfigOpenApiVO(int32(123), false, false, false, false, "Name_example", int32(123), int32(123), false) // UpdateSsidBasicConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidBasicConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidBasicConfigOpenApiVO(updateSsidBasicConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidBasicConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidBasicConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidBasicConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidBasicConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidBasicConfigOpenApiVO** | [**UpdateSsidBasicConfigOpenApiVO**](UpdateSsidBasicConfigOpenApiVO.md) |  | 

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


## UpdateSsidDhcpOptionConfigTemplate

> OperationResponseWithoutResult UpdateSsidDhcpOptionConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidDhcpOptionOpenApiVO(updateSsidDhcpOptionOpenApiVO).Execute()

Update SSID template DHCP option 82 config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidDhcpOptionOpenApiVO := *openapiclient.NewUpdateSsidDhcpOptionOpenApiVO(false) // UpdateSsidDhcpOptionOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidDhcpOptionConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidDhcpOptionOpenApiVO(updateSsidDhcpOptionOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidDhcpOptionConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidDhcpOptionConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidDhcpOptionConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidDhcpOptionConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidDhcpOptionOpenApiVO** | [**UpdateSsidDhcpOptionOpenApiVO**](UpdateSsidDhcpOptionOpenApiVO.md) |  | 

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


## UpdateSsidHotspotV2SettingTemplate

> OperationResponseWithoutResult UpdateSsidHotspotV2SettingTemplate(ctx, omadacId, wlanId, siteTemplateId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()

Update SSID template Hotspot2.0 config



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
	wlanId := "wlanId_example" // string | WLAN ID
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId
	ssidId := "ssidId_example" // string | ssidId
	updateSsidHotspotV2SettingOpenApiVO := *openapiclient.NewUpdateSsidHotspotV2SettingOpenApiVO(false) // UpdateSsidHotspotV2SettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidHotspotV2SettingTemplate(context.Background(), omadacId, wlanId, siteTemplateId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidHotspotV2SettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidHotspotV2SettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidHotspotV2SettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**wlanId** | **string** | WLAN ID | 
**siteTemplateId** | **string** | siteTemplateId | 
**ssidId** | **string** | ssidId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidHotspotV2SettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidHotspotV2SettingOpenApiVO** | [**UpdateSsidHotspotV2SettingOpenApiVO**](UpdateSsidHotspotV2SettingOpenApiVO.md) |  | 

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


## UpdateSsidMacFilterConfigTemplate

> OperationResponseWithoutResult UpdateSsidMacFilterConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()

Update SSID template mac filter config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMacFilterOpenApiVO := *openapiclient.NewUpdateSsidMacFilterOpenApiVO(false) // UpdateSsidMacFilterOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidMacFilterConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidMacFilterConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidMacFilterConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidMacFilterConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidMacFilterConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidMacFilterOpenApiVO** | [**UpdateSsidMacFilterOpenApiVO**](UpdateSsidMacFilterOpenApiVO.md) |  | 

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


## UpdateSsidMultiCastConfigTemplate

> OperationResponseWithoutResult UpdateSsidMultiCastConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidMultiCastOpenApiVO(updateSsidMultiCastOpenApiVO).Execute()

Update SSID template Multicast/Broadcast management config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMultiCastOpenApiVO := *openapiclient.NewUpdateSsidMultiCastOpenApiVO(false, int32(123), false, false, false) // UpdateSsidMultiCastOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidMultiCastConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidMultiCastOpenApiVO(updateSsidMultiCastOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidMultiCastConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidMultiCastConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidMultiCastConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidMultiCastConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidMultiCastOpenApiVO** | [**UpdateSsidMultiCastOpenApiVO**](UpdateSsidMultiCastOpenApiVO.md) |  | 

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


## UpdateSsidRateControlConfigTemplate

> OperationResponseWithoutResult UpdateSsidRateControlConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidRateControlOpenApiVO(updateSsidRateControlOpenApiVO).Execute()

Update SSID template 802.11 rate control config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidRateControlOpenApiVO := *openapiclient.NewUpdateSsidRateControlOpenApiVO(false, false) // UpdateSsidRateControlOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidRateControlConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidRateControlOpenApiVO(updateSsidRateControlOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidRateControlConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidRateControlConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidRateControlConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidRateControlConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidRateControlOpenApiVO** | [**UpdateSsidRateControlOpenApiVO**](UpdateSsidRateControlOpenApiVO.md) |  | 

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


## UpdateSsidRateLimitConfigTemplate

> OperationResponseWithoutResult UpdateSsidRateLimitConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidRateLimitOpenApiVO(updateSsidRateLimitOpenApiVO).Execute()

Update SSID template rate limit config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidRateLimitOpenApiVO := *openapiclient.NewUpdateSsidRateLimitOpenApiVO() // UpdateSsidRateLimitOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidRateLimitConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidRateLimitOpenApiVO(updateSsidRateLimitOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidRateLimitConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidRateLimitConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidRateLimitConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidRateLimitConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidRateLimitOpenApiVO** | [**UpdateSsidRateLimitOpenApiVO**](UpdateSsidRateLimitOpenApiVO.md) |  | 

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


## UpdateSsidWlanScheduleTemplate

> OperationResponseWithoutResult UpdateSsidWlanScheduleTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidWlanScheduleOpenApiVO(updateSsidWlanScheduleOpenApiVO).Execute()

Update SSID template WLAN schedule config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidWlanScheduleOpenApiVO := *openapiclient.NewUpdateSsidWlanScheduleOpenApiVO(false) // UpdateSsidWlanScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidWlanScheduleTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidWlanScheduleOpenApiVO(updateSsidWlanScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidWlanScheduleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidWlanScheduleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidWlanScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidWlanScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidWlanScheduleOpenApiVO** | [**UpdateSsidWlanScheduleOpenApiVO**](UpdateSsidWlanScheduleOpenApiVO.md) |  | 

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


## UpdateWlanGroupTemplate

> OperationResponseWithoutResult UpdateWlanGroupTemplate(ctx, omadacId, siteTemplateId, wlanId).UpdateWlanGroupOpenApiVO(updateWlanGroupOpenApiVO).Execute()

Modify an existing WLAN group template



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
	wlanId := "wlanId_example" // string | WLAN ID
	updateWlanGroupOpenApiVO := *openapiclient.NewUpdateWlanGroupOpenApiVO("Name_example") // UpdateWlanGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateWlanGroupTemplate(context.Background(), omadacId, siteTemplateId, wlanId).UpdateWlanGroupOpenApiVO(updateWlanGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateWlanGroupTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWlanGroupTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateWlanGroupTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWlanGroupTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateWlanGroupOpenApiVO** | [**UpdateWlanGroupOpenApiVO**](UpdateWlanGroupOpenApiVO.md) |  | 

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


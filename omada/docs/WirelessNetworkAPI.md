# WirelessNetworkAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountWlans**](WirelessNetworkAPI.md#countwlans) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/cmd/countWlans | Check if the number of WLAN groups is out of limit
[**CreateSsid**](WirelessNetworkAPI.md#createssid) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids | Create new SSID
[**CreateWlanGroup**](WirelessNetworkAPI.md#createwlangroup) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans | Create new WLAN group
[**DeleteSsid**](WirelessNetworkAPI.md#deletessid) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId} | Delete an existing SSID
[**DeleteWlanGroup**](WirelessNetworkAPI.md#deletewlangroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId} | Delete an existing WLAN group
[**GetGridWlanGroup**](WirelessNetworkAPI.md#getgridwlangroup) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/wireless-network/wlans | Get WLAN group list paging query
[**GetSiteApExist6G**](WirelessNetworkAPI.md#getsiteapexist6g) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/multi-band-6g/exist | Get the site exists 6G ap
[**GetSsidDetail**](WirelessNetworkAPI.md#getssiddetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId} | Get SSID detail info
[**GetSsidList**](WirelessNetworkAPI.md#getssidlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids | Get SSID list
[**GetSsidsBySite**](WirelessNetworkAPI.md#getssidsbysite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids | Get all ssids of the site
[**GetWlanGroupList**](WirelessNetworkAPI.md#getwlangrouplist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans | Get WLAN group list
[**UpdateSsidBasicConfig**](WirelessNetworkAPI.md#updatessidbasicconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-basic-config | Update SSID basic config
[**UpdateSsidDhcpOptionConfig**](WirelessNetworkAPI.md#updatessiddhcpoptionconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-dhcp-option | Update SSID DHCP option 82 config
[**UpdateSsidHotspotV2Setting**](WirelessNetworkAPI.md#updatessidhotspotv2setting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-hotspotv2 | Update SSID Hotspot2.0 config
[**UpdateSsidMacFilterConfig**](WirelessNetworkAPI.md#updatessidmacfilterconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-mac-filter | Update SSID mac filter config
[**UpdateSsidMultiCastConfig**](WirelessNetworkAPI.md#updatessidmulticastconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-multicast-config | Update SSID Multicast/Broadcast management config
[**UpdateSsidRateControlConfig**](WirelessNetworkAPI.md#updatessidratecontrolconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-rate-control | Update SSID 802.11 rate control config
[**UpdateSsidRateLimitConfig**](WirelessNetworkAPI.md#updatessidratelimitconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-rate-limit | Update SSID rate limit config
[**UpdateSsidWlanSchedule**](WirelessNetworkAPI.md#updatessidwlanschedule) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-wlan-schedule | Update SSID WLAN schedule config
[**UpdateWlanGroup**](WirelessNetworkAPI.md#updatewlangroup) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId} | Modify an existing WLAN group



## CountWlans

> OperationResponseWlanGroupStatusOpenApiVO CountWlans(ctx, omadacId, siteId).Execute()

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
	siteId := "siteId_example" // string | siteId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.CountWlans(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.CountWlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWlans`: OperationResponseWlanGroupStatusOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.CountWlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | siteId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountWlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWlanGroupStatusOpenApiVO**](OperationResponseWlanGroupStatusOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSsid

> OperationResponse CreateSsid(ctx, omadacId, siteId, wlanId).CreateSsidOpenApiVO(createSsidOpenApiVO).Execute()

Create new SSID



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
	wlanId := "wlanId_example" // string | WLAN ID
	createSsidOpenApiVO := *openapiclient.NewCreateSsidOpenApiVO(int32(123), false, int32(123), false, false, false, false, "Name_example", int32(123), int32(123), false) // CreateSsidOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.CreateSsid(context.Background(), omadacId, siteId, wlanId).CreateSsidOpenApiVO(createSsidOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.CreateSsid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSsid`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.CreateSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createSsidOpenApiVO** | [**CreateSsidOpenApiVO**](CreateSsidOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWlanGroup

> OperationResponse CreateWlanGroup(ctx, omadacId, siteId).CreateWlanGroupOpenApiVO(createWlanGroupOpenApiVO).Execute()

Create new WLAN group



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
	createWlanGroupOpenApiVO := *openapiclient.NewCreateWlanGroupOpenApiVO(false, "Name_example") // CreateWlanGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.CreateWlanGroup(context.Background(), omadacId, siteId).CreateWlanGroupOpenApiVO(createWlanGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.CreateWlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWlanGroup`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.CreateWlanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWlanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createWlanGroupOpenApiVO** | [**CreateWlanGroupOpenApiVO**](CreateWlanGroupOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSsid

> OperationResponseWithoutResult DeleteSsid(ctx, omadacId, siteId, wlanId, ssidId).Execute()

Delete an existing SSID



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.DeleteSsid(context.Background(), omadacId, siteId, wlanId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.DeleteSsid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSsid`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.DeleteSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSsidRequest struct via the builder pattern


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


## DeleteWlanGroup

> OperationResponseWithoutResult DeleteWlanGroup(ctx, omadacId, siteId, wlanId).Execute()

Delete an existing WLAN group



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
	wlanId := "wlanId_example" // string | WLAN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.DeleteWlanGroup(context.Background(), omadacId, siteId, wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.DeleteWlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWlanGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.DeleteWlanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWlanGroupRequest struct via the builder pattern


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


## GetGridWlanGroup

> OperationResponseWlanGroupGridOpenApiVO GetGridWlanGroup(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get WLAN group list paging query



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.GetGridWlanGroup(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetGridWlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridWlanGroup`: OperationResponseWlanGroupGridOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetGridWlanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridWlanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseWlanGroupGridOpenApiVO**](OperationResponseWlanGroupGridOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteApExist6G

> OperationResponseApExist6GOpenApiVO GetSiteApExist6G(ctx, omadacId, siteId).Execute()

Get the site exists 6G ap



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
	resp, r, err := apiClient.WirelessNetworkAPI.GetSiteApExist6G(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetSiteApExist6G``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteApExist6G`: OperationResponseApExist6GOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetSiteApExist6G`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteApExist6GRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseApExist6GOpenApiVO**](OperationResponseApExist6GOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidDetail

> OperationResponseSsidDetailOpenApiVO GetSsidDetail(ctx, omadacId, siteId, wlanId, ssidId).Execute()

Get SSID detail info



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.GetSsidDetail(context.Background(), omadacId, siteId, wlanId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetSsidDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidDetail`: OperationResponseSsidDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetSsidDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseSsidDetailOpenApiVO**](OperationResponseSsidDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidList

> OperationResponseGridVOSsidOpenApiVO GetSsidList(ctx, omadacId, siteId, wlanId).Page(page).PageSize(pageSize).Execute()

Get SSID list



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
	wlanId := "wlanId_example" // string | WLAN ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.GetSsidList(context.Background(), omadacId, siteId, wlanId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetSsidList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidList`: OperationResponseGridVOSsidOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetSsidList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOSsidOpenApiVO**](OperationResponseGridVOSsidOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidsBySite

> OperationResponseListWlanSimpleOpenApiVO GetSsidsBySite(ctx, omadacId, siteId).Type_(type_).Execute()

Get all ssids of the site



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
	type_ := int32(56) // int32 | Device Type. Supported type: ap and wireless router. 1: ap, 2: wireless router, 3: ap and wireless router

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.GetSsidsBySite(context.Background(), omadacId, siteId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetSsidsBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidsBySite`: OperationResponseListWlanSimpleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetSsidsBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidsBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **int32** | Device Type. Supported type: ap and wireless router. 1: ap, 2: wireless router, 3: ap and wireless router | 

### Return type

[**OperationResponseListWlanSimpleOpenApiVO**](OperationResponseListWlanSimpleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWlanGroupList

> OperationResponseListWlanGroupOpenApiVO GetWlanGroupList(ctx, omadacId, siteId).Execute()

Get WLAN group list



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
	resp, r, err := apiClient.WirelessNetworkAPI.GetWlanGroupList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetWlanGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWlanGroupList`: OperationResponseListWlanGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetWlanGroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWlanGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListWlanGroupOpenApiVO**](OperationResponseListWlanGroupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSsidBasicConfig

> OperationResponseWithoutResult UpdateSsidBasicConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidBasicConfigOpenApiVO(updateSsidBasicConfigOpenApiVO).Execute()

Update SSID basic config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidBasicConfigOpenApiVO := *openapiclient.NewUpdateSsidBasicConfigOpenApiVO(int32(123), false, false, false, false, "Name_example", int32(123), int32(123), false) // UpdateSsidBasicConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidBasicConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidBasicConfigOpenApiVO(updateSsidBasicConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidBasicConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidBasicConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidBasicConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidBasicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidBasicConfigOpenApiVO** | [**UpdateSsidBasicConfigOpenApiVO**](UpdateSsidBasicConfigOpenApiVO.md) |  | 

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


## UpdateSsidDhcpOptionConfig

> OperationResponseWithoutResult UpdateSsidDhcpOptionConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidDhcpOptionOpenApiVO(updateSsidDhcpOptionOpenApiVO).Execute()

Update SSID DHCP option 82 config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidDhcpOptionOpenApiVO := *openapiclient.NewUpdateSsidDhcpOptionOpenApiVO(false) // UpdateSsidDhcpOptionOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidDhcpOptionConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidDhcpOptionOpenApiVO(updateSsidDhcpOptionOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidDhcpOptionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidDhcpOptionConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidDhcpOptionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidDhcpOptionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidDhcpOptionOpenApiVO** | [**UpdateSsidDhcpOptionOpenApiVO**](UpdateSsidDhcpOptionOpenApiVO.md) |  | 

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


## UpdateSsidHotspotV2Setting

> OperationResponseWithoutResult UpdateSsidHotspotV2Setting(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()

Update SSID Hotspot2.0 config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | ssidId
	updateSsidHotspotV2SettingOpenApiVO := *openapiclient.NewUpdateSsidHotspotV2SettingOpenApiVO(false) // UpdateSsidHotspotV2SettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidHotspotV2Setting(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidHotspotV2Setting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidHotspotV2Setting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidHotspotV2Setting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | ssidId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidHotspotV2SettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidHotspotV2SettingOpenApiVO** | [**UpdateSsidHotspotV2SettingOpenApiVO**](UpdateSsidHotspotV2SettingOpenApiVO.md) |  | 

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


## UpdateSsidMacFilterConfig

> OperationResponseWithoutResult UpdateSsidMacFilterConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()

Update SSID mac filter config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMacFilterOpenApiVO := *openapiclient.NewUpdateSsidMacFilterOpenApiVO(false) // UpdateSsidMacFilterOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidMacFilterConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidMacFilterConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidMacFilterConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidMacFilterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidMacFilterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidMacFilterOpenApiVO** | [**UpdateSsidMacFilterOpenApiVO**](UpdateSsidMacFilterOpenApiVO.md) |  | 

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


## UpdateSsidMultiCastConfig

> OperationResponseWithoutResult UpdateSsidMultiCastConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidMultiCastOpenApiVO(updateSsidMultiCastOpenApiVO).Execute()

Update SSID Multicast/Broadcast management config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMultiCastOpenApiVO := *openapiclient.NewUpdateSsidMultiCastOpenApiVO(false, int32(123), false, false, false) // UpdateSsidMultiCastOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidMultiCastConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidMultiCastOpenApiVO(updateSsidMultiCastOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidMultiCastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidMultiCastConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidMultiCastConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidMultiCastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidMultiCastOpenApiVO** | [**UpdateSsidMultiCastOpenApiVO**](UpdateSsidMultiCastOpenApiVO.md) |  | 

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


## UpdateSsidRateControlConfig

> OperationResponseWithoutResult UpdateSsidRateControlConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidRateControlOpenApiVO(updateSsidRateControlOpenApiVO).Execute()

Update SSID 802.11 rate control config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidRateControlOpenApiVO := *openapiclient.NewUpdateSsidRateControlOpenApiVO(false, false) // UpdateSsidRateControlOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidRateControlConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidRateControlOpenApiVO(updateSsidRateControlOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidRateControlConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidRateControlConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidRateControlConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidRateControlConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidRateControlOpenApiVO** | [**UpdateSsidRateControlOpenApiVO**](UpdateSsidRateControlOpenApiVO.md) |  | 

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


## UpdateSsidRateLimitConfig

> OperationResponseWithoutResult UpdateSsidRateLimitConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidRateLimitOpenApiVO(updateSsidRateLimitOpenApiVO).Execute()

Update SSID rate limit config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidRateLimitOpenApiVO := *openapiclient.NewUpdateSsidRateLimitOpenApiVO() // UpdateSsidRateLimitOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidRateLimitConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidRateLimitOpenApiVO(updateSsidRateLimitOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidRateLimitConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidRateLimitConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidRateLimitConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidRateLimitConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidRateLimitOpenApiVO** | [**UpdateSsidRateLimitOpenApiVO**](UpdateSsidRateLimitOpenApiVO.md) |  | 

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


## UpdateSsidWlanSchedule

> OperationResponseWithoutResult UpdateSsidWlanSchedule(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidWlanScheduleOpenApiVO(updateSsidWlanScheduleOpenApiVO).Execute()

Update SSID WLAN schedule config



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
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidWlanScheduleOpenApiVO := *openapiclient.NewUpdateSsidWlanScheduleOpenApiVO(false) // UpdateSsidWlanScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidWlanSchedule(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidWlanScheduleOpenApiVO(updateSsidWlanScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidWlanSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidWlanSchedule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidWlanSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidWlanScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidWlanScheduleOpenApiVO** | [**UpdateSsidWlanScheduleOpenApiVO**](UpdateSsidWlanScheduleOpenApiVO.md) |  | 

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


## UpdateWlanGroup

> OperationResponseWithoutResult UpdateWlanGroup(ctx, omadacId, siteId, wlanId).UpdateWlanGroupOpenApiVO(updateWlanGroupOpenApiVO).Execute()

Modify an existing WLAN group



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
	wlanId := "wlanId_example" // string | WLAN ID
	updateWlanGroupOpenApiVO := *openapiclient.NewUpdateWlanGroupOpenApiVO("Name_example") // UpdateWlanGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateWlanGroup(context.Background(), omadacId, siteId, wlanId).UpdateWlanGroupOpenApiVO(updateWlanGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateWlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWlanGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateWlanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**wlanId** | **string** | WLAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWlanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateWlanGroupOpenApiVO** | [**UpdateWlanGroupOpenApiVO**](UpdateWlanGroupOpenApiVO.md) |  | 

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


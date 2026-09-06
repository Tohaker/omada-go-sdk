# WirelessNetworkAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountApGroups**](WirelessNetworkAPI.md#countapgroups) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/count-ap-groups | Check if the number of AP Groups is out of limit
[**CountWlans**](WirelessNetworkAPI.md#countwlans) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/cmd/countWlans | Check if the number of WLAN groups is out of limit
[**CreateApGroup**](WirelessNetworkAPI.md#createapgroup) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ap-groups | Create new AP Group
[**CreateSsid**](WirelessNetworkAPI.md#createssid) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids | Create new SSID
[**CreateSsidV2**](WirelessNetworkAPI.md#createssidv2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/wireless-network/ssids | Create new SSID v2
[**CreateWlanGroup**](WirelessNetworkAPI.md#createwlangroup) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans | Create new WLAN Group
[**DeleteApGroup**](WirelessNetworkAPI.md#deleteapgroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/ap-groups/{apGroupId} | Delete an existing AP Group
[**DeleteSsid**](WirelessNetworkAPI.md#deletessid) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId} | Delete an existing SSID
[**DeleteSsidBySite**](WirelessNetworkAPI.md#deletessidbysite) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId} | Delete an existing SSID
[**DeleteWlanGroup**](WirelessNetworkAPI.md#deletewlangroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId} | Delete an existing WLAN Group
[**GetAllApGroupListBySite**](WirelessNetworkAPI.md#getallapgrouplistbysite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ap-groups/aps | Get All AP grouping by ap group
[**GetApGroupInfo**](WirelessNetworkAPI.md#getapgroupinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ap-groups/{apGroupId}/info | Get an existing AP Group info
[**GetApGroupList**](WirelessNetworkAPI.md#getapgrouplist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ap-groups | Get AP Group list
[**GetGridWlanGroup**](WirelessNetworkAPI.md#getgridwlangroup) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/wireless-network/wlans | Get WLAN Group list paging query
[**GetSiteApExist6G**](WirelessNetworkAPI.md#getsiteapexist6g) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/multi-band-6g/exist | Get the site exists 6G ap
[**GetSsidDetail**](WirelessNetworkAPI.md#getssiddetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId} | Get SSID detail info
[**GetSsidDetailV2**](WirelessNetworkAPI.md#getssiddetailv2) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId} | Get SSID detail info
[**GetSsidList**](WirelessNetworkAPI.md#getssidlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids | Get SSID list
[**GetSsidListV2**](WirelessNetworkAPI.md#getssidlistv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/wireless-network/ssids | Get SSID list by site
[**GetSsidsBySite**](WirelessNetworkAPI.md#getssidsbysite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids | Get all ssids of the site
[**GetWlanGroup**](WirelessNetworkAPI.md#getwlangroup) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId} | Get WLAN Group by wlan id
[**GetWlanGroupList**](WirelessNetworkAPI.md#getwlangrouplist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans | Get WLAN Group list
[**QuerySsidBindApGroups**](WirelessNetworkAPI.md#queryssidbindapgroups) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/ap-groups | Query AP Groups bound to the SSID
[**QuerySsidDuplicateNameBySite1**](WirelessNetworkAPI.md#queryssidduplicatenamebysite1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/duplicate-name | Query SSIDs with the same name
[**UpdateApGroup**](WirelessNetworkAPI.md#updateapgroup) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/ap-groups/{apGroupId} | Modify an existing AP Group
[**UpdateSsidBandSteerConfig**](WirelessNetworkAPI.md#updatessidbandsteerconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-band-steer | Update SSID band steer config
[**UpdateSsidBandSteerConfigBySite**](WirelessNetworkAPI.md#updatessidbandsteerconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/band-steer | Update SSID band steer config by site
[**UpdateSsidBasicConfig**](WirelessNetworkAPI.md#updatessidbasicconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-basic-config | Update SSID basic config
[**UpdateSsidBasicConfigBySite**](WirelessNetworkAPI.md#updatessidbasicconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/basic-config | Update SSID basic config by site
[**UpdateSsidBindApGroups**](WirelessNetworkAPI.md#updatessidbindapgroups) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/ap-groups | Update SSID binding ap groups
[**UpdateSsidDhcpOptionConfig**](WirelessNetworkAPI.md#updatessiddhcpoptionconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-dhcp-option | Update SSID DHCP option 82 config
[**UpdateSsidDhcpOptionConfigBySite**](WirelessNetworkAPI.md#updatessiddhcpoptionconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/dhcp-option | Update SSID DHCP option 82 config by site
[**UpdateSsidEnableStatusBySite1**](WirelessNetworkAPI.md#updatessidenablestatusbysite1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/enable | Update SSID Enable Status by site
[**UpdateSsidHotspotV2Setting**](WirelessNetworkAPI.md#updatessidhotspotv2setting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-hotspotv2 | Update SSID Hotspot2.0 config
[**UpdateSsidHotspotV2SettingBySite**](WirelessNetworkAPI.md#updatessidhotspotv2settingbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/hotspotv2 | Update SSID Hotspot2.0 config by site
[**UpdateSsidLoadBalanceConfig**](WirelessNetworkAPI.md#updatessidloadbalanceconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-load-balance | Update SSID load balance config
[**UpdateSsidLoadBalanceConfigBySite**](WirelessNetworkAPI.md#updatessidloadbalanceconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/load-balance | Update SSID load balance config by site
[**UpdateSsidMacFilterConfig**](WirelessNetworkAPI.md#updatessidmacfilterconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-mac-filter | Update SSID mac filter config
[**UpdateSsidMacFilterConfigBySite**](WirelessNetworkAPI.md#updatessidmacfilterconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/mac-filter | Update SSID mac filter config by site
[**UpdateSsidMultiCastConfig**](WirelessNetworkAPI.md#updatessidmulticastconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-multicast-config | Update SSID Multicast/Broadcast management config
[**UpdateSsidMultiCastConfigBySite**](WirelessNetworkAPI.md#updatessidmulticastconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/multicast-config | Update SSID Multicast/Broadcast management config by site
[**UpdateSsidRateControlConfig**](WirelessNetworkAPI.md#updatessidratecontrolconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-rate-control | Update SSID 802.11 rate control config
[**UpdateSsidRateControlConfigBySite**](WirelessNetworkAPI.md#updatessidratecontrolconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/rate-control | Update SSID 802.11 rate control config by site
[**UpdateSsidRateLimitConfig**](WirelessNetworkAPI.md#updatessidratelimitconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-rate-limit | Update SSID rate limit config
[**UpdateSsidRateLimitConfigBySite**](WirelessNetworkAPI.md#updatessidratelimitconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/rate-limit | Update SSID rate limit config by site
[**UpdateSsidWifiCallingConfig**](WirelessNetworkAPI.md#updatessidwificallingconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-wifi-calling | Update SSID wifi calling config
[**UpdateSsidWifiCallingConfigBySite**](WirelessNetworkAPI.md#updatessidwificallingconfigbysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/wifi-calling | Update SSID wifi calling config by site
[**UpdateSsidWlanSchedule**](WirelessNetworkAPI.md#updatessidwlanschedule) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-wlan-schedule | Update SSID WLAN schedule config
[**UpdateSsidWlanScheduleBySite**](WirelessNetworkAPI.md#updatessidwlanschedulebysite) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/ssids/{ssidId}/wlan-schedule | Update SSID WLAN schedule config by site
[**UpdateWlanGroup**](WirelessNetworkAPI.md#updatewlangroup) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wireless-network/wlans/{wlanId} | Modify an existing WLAN Group



## CountApGroups

> OperationResponseApGroupStatusOpenApiVO CountApGroups(ctx, omadacId, siteId).Execute()

Check if the number of AP Groups is out of limit



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
	resp, r, err := apiClient.WirelessNetworkAPI.CountApGroups(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.CountApGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountApGroups`: OperationResponseApGroupStatusOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.CountApGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountApGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseApGroupStatusOpenApiVO**](OperationResponseApGroupStatusOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## CreateApGroup

> OperationResponseAddApGroupResultVO CreateApGroup(ctx, omadacId, siteId).CreateApGroupOpenApiVO(createApGroupOpenApiVO).Execute()

Create new AP Group



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
	createApGroupOpenApiVO := *openapiclient.NewCreateApGroupOpenApiVO("Name_example") // CreateApGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.CreateApGroup(context.Background(), omadacId, siteId).CreateApGroupOpenApiVO(createApGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.CreateApGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApGroup`: OperationResponseAddApGroupResultVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.CreateApGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createApGroupOpenApiVO** | [**CreateApGroupOpenApiVO**](CreateApGroupOpenApiVO.md) |  | 

### Return type

[**OperationResponseAddApGroupResultVO**](OperationResponseAddApGroupResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
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


## CreateSsidV2

> OperationResponseCreateSsidResultVO CreateSsidV2(ctx, omadacId, siteId).CreateSsidOpenApiVO(createSsidOpenApiVO).Execute()

Create new SSID v2



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
	createSsidOpenApiVO := *openapiclient.NewCreateSsidOpenApiVO(int32(123), false, int32(123), false, false, false, false, "Name_example", int32(123), int32(123), false) // CreateSsidOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.CreateSsidV2(context.Background(), omadacId, siteId).CreateSsidOpenApiVO(createSsidOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.CreateSsidV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSsidV2`: OperationResponseCreateSsidResultVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.CreateSsidV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSsidV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createSsidOpenApiVO** | [**CreateSsidOpenApiVO**](CreateSsidOpenApiVO.md) |  | 

### Return type

[**OperationResponseCreateSsidResultVO**](OperationResponseCreateSsidResultVO.md)

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

Create new WLAN Group



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


## DeleteApGroup

> OperationResponseWithoutResult DeleteApGroup(ctx, omadacId, siteId, apGroupId).Execute()

Delete an existing AP Group



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
	apGroupId := "apGroupId_example" // string | AP GROUP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.DeleteApGroup(context.Background(), omadacId, siteId, apGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.DeleteApGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.DeleteApGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apGroupId** | **string** | AP GROUP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApGroupRequest struct via the builder pattern


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


## DeleteSsidBySite

> OperationResponse DeleteSsidBySite(ctx, omadacId, siteId, ssidId).Execute()

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
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.DeleteSsidBySite(context.Background(), omadacId, siteId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.DeleteSsidBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSsidBySite`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.DeleteSsidBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSsidBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponse**](OperationResponse.md)

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

Delete an existing WLAN Group



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


## GetAllApGroupListBySite

> OperationResponseGridVOSsidDeviceOpenApiVO GetAllApGroupListBySite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get All AP grouping by ap group



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
	resp, r, err := apiClient.WirelessNetworkAPI.GetAllApGroupListBySite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetAllApGroupListBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllApGroupListBySite`: OperationResponseGridVOSsidDeviceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetAllApGroupListBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApGroupListBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOSsidDeviceOpenApiVO**](OperationResponseGridVOSsidDeviceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApGroupInfo

> OperationResponseApGroupOpenApiVO GetApGroupInfo(ctx, omadacId, siteId, apGroupId).Execute()

Get an existing AP Group info



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
	apGroupId := "apGroupId_example" // string | AP GROUP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.GetApGroupInfo(context.Background(), omadacId, siteId, apGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetApGroupInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApGroupInfo`: OperationResponseApGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetApGroupInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apGroupId** | **string** | AP GROUP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApGroupInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApGroupOpenApiVO**](OperationResponseApGroupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApGroupList

> OperationResponseApGroupGridVOApGroupOpenApiVO GetApGroupList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get AP Group list



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
	resp, r, err := apiClient.WirelessNetworkAPI.GetApGroupList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetApGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApGroupList`: OperationResponseApGroupGridVOApGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetApGroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseApGroupGridVOApGroupOpenApiVO**](OperationResponseApGroupGridVOApGroupOpenApiVO.md)

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

Get WLAN Group list paging query



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


## GetSsidDetailV2

> OperationResponseSsidDetailOpenApiVO GetSsidDetailV2(ctx, omadacId, siteId, ssidId).Execute()

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
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.GetSsidDetailV2(context.Background(), omadacId, siteId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetSsidDetailV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidDetailV2`: OperationResponseSsidDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetSsidDetailV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidDetailV2Request struct via the builder pattern


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


## GetSsidListV2

> OperationResponseGridVOSsidOpenApiVO GetSsidListV2(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get SSID list by site



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
	resp, r, err := apiClient.WirelessNetworkAPI.GetSsidListV2(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetSsidListV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidListV2`: OperationResponseGridVOSsidOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetSsidListV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidListV2Request struct via the builder pattern


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


## GetWlanGroup

> OperationResponseWlanGroupOpenApiVO GetWlanGroup(ctx, omadacId, siteId, wlanId).Execute()

Get WLAN Group by wlan id



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
	resp, r, err := apiClient.WirelessNetworkAPI.GetWlanGroup(context.Background(), omadacId, siteId, wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.GetWlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWlanGroup`: OperationResponseWlanGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.GetWlanGroup`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWlanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWlanGroupOpenApiVO**](OperationResponseWlanGroupOpenApiVO.md)

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

Get WLAN Group list



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


## QuerySsidBindApGroups

> OperationResponseApGroupDetailOpenApiVO QuerySsidBindApGroups(ctx, omadacId, siteId, ssidId).Execute()

Query AP Groups bound to the SSID



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
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.QuerySsidBindApGroups(context.Background(), omadacId, siteId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.QuerySsidBindApGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySsidBindApGroups`: OperationResponseApGroupDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.QuerySsidBindApGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySsidBindApGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApGroupDetailOpenApiVO**](OperationResponseApGroupDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySsidDuplicateNameBySite1

> OperationResponseDuplicateSsidOpenApiVO QuerySsidDuplicateNameBySite1(ctx, omadacId, siteId).Execute()

Query SSIDs with the same name



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
	resp, r, err := apiClient.WirelessNetworkAPI.QuerySsidDuplicateNameBySite1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.QuerySsidDuplicateNameBySite1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySsidDuplicateNameBySite1`: OperationResponseDuplicateSsidOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.QuerySsidDuplicateNameBySite1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySsidDuplicateNameBySite1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDuplicateSsidOpenApiVO**](OperationResponseDuplicateSsidOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApGroup

> OperationResponseWithoutResult UpdateApGroup(ctx, omadacId, siteId, apGroupId).UpdateApGroupOpenApiVO(updateApGroupOpenApiVO).Execute()

Modify an existing AP Group



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
	apGroupId := "apGroupId_example" // string | AP GROUP ID
	updateApGroupOpenApiVO := *openapiclient.NewUpdateApGroupOpenApiVO("Name_example") // UpdateApGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateApGroup(context.Background(), omadacId, siteId, apGroupId).UpdateApGroupOpenApiVO(updateApGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateApGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateApGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apGroupId** | **string** | AP GROUP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApGroupOpenApiVO** | [**UpdateApGroupOpenApiVO**](UpdateApGroupOpenApiVO.md) |  | 

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


## UpdateSsidBandSteerConfig

> OperationResponseWithoutResult UpdateSsidBandSteerConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidBandSteerOpenApiVO(updateSsidBandSteerOpenApiVO).Execute()

Update SSID band steer config



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
	updateSsidBandSteerOpenApiVO := *openapiclient.NewUpdateSsidBandSteerOpenApiVO(int32(123)) // UpdateSsidBandSteerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidBandSteerConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidBandSteerOpenApiVO(updateSsidBandSteerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidBandSteerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidBandSteerConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidBandSteerConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateSsidBandSteerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidBandSteerOpenApiVO** | [**UpdateSsidBandSteerOpenApiVO**](UpdateSsidBandSteerOpenApiVO.md) |  | 

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


## UpdateSsidBandSteerConfigBySite

> OperationResponseWithoutResult UpdateSsidBandSteerConfigBySite(ctx, omadacId, siteId, ssidId).UpdateSsidBandSteerOpenApiVO(updateSsidBandSteerOpenApiVO).Execute()

Update SSID band steer config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidBandSteerOpenApiVO := *openapiclient.NewUpdateSsidBandSteerOpenApiVO(int32(123)) // UpdateSsidBandSteerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidBandSteerConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidBandSteerOpenApiVO(updateSsidBandSteerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidBandSteerConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidBandSteerConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidBandSteerConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidBandSteerConfigBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateSsidBandSteerOpenApiVO** | [**UpdateSsidBandSteerOpenApiVO**](UpdateSsidBandSteerOpenApiVO.md) |  | 

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


## UpdateSsidBasicConfigBySite

> OperationResponseWithoutResult UpdateSsidBasicConfigBySite(ctx, omadacId, siteId, ssidId).UpdateSsidBasicConfigOpenApiVO(updateSsidBasicConfigOpenApiVO).Execute()

Update SSID basic config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidBasicConfigOpenApiVO := *openapiclient.NewUpdateSsidBasicConfigOpenApiVO(int32(123), false, false, false, false, "Name_example", int32(123), int32(123), false) // UpdateSsidBasicConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidBasicConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidBasicConfigOpenApiVO(updateSsidBasicConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidBasicConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidBasicConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidBasicConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidBasicConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidBindApGroups

> OperationResponseWithoutResult UpdateSsidBindApGroups(ctx, omadacId, siteId, ssidId).UpdateSsidBindApGroupOpenApiVO(updateSsidBindApGroupOpenApiVO).Execute()

Update SSID binding ap groups



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidBindApGroupOpenApiVO := *openapiclient.NewUpdateSsidBindApGroupOpenApiVO([]string{"ApGroupIds_example"}) // UpdateSsidBindApGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidBindApGroups(context.Background(), omadacId, siteId, ssidId).UpdateSsidBindApGroupOpenApiVO(updateSsidBindApGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidBindApGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidBindApGroups`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidBindApGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidBindApGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateSsidBindApGroupOpenApiVO** | [**UpdateSsidBindApGroupOpenApiVO**](UpdateSsidBindApGroupOpenApiVO.md) |  | 

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


## UpdateSsidDhcpOptionConfigBySite

> OperationResponseWithoutResult UpdateSsidDhcpOptionConfigBySite(ctx, omadacId, siteId, ssidId).UpdateSsidDhcpOptionOpenApiVO(updateSsidDhcpOptionOpenApiVO).Execute()

Update SSID DHCP option 82 config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidDhcpOptionOpenApiVO := *openapiclient.NewUpdateSsidDhcpOptionOpenApiVO(false) // UpdateSsidDhcpOptionOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidDhcpOptionConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidDhcpOptionOpenApiVO(updateSsidDhcpOptionOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidDhcpOptionConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidDhcpOptionConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidDhcpOptionConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidDhcpOptionConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidEnableStatusBySite1

> OperationResponseWithoutResult UpdateSsidEnableStatusBySite1(ctx, omadacId, siteId, ssidId).UpdateSsidEnableStatusOpenApiVO(updateSsidEnableStatusOpenApiVO).Execute()

Update SSID Enable Status by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidEnableStatusOpenApiVO := *openapiclient.NewUpdateSsidEnableStatusOpenApiVO(false) // UpdateSsidEnableStatusOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidEnableStatusBySite1(context.Background(), omadacId, siteId, ssidId).UpdateSsidEnableStatusOpenApiVO(updateSsidEnableStatusOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidEnableStatusBySite1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidEnableStatusBySite1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidEnableStatusBySite1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidEnableStatusBySite1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateSsidEnableStatusOpenApiVO** | [**UpdateSsidEnableStatusOpenApiVO**](UpdateSsidEnableStatusOpenApiVO.md) |  | 

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
	ssidId := "ssidId_example" // string | SSID ID
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
**ssidId** | **string** | SSID ID | 

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


## UpdateSsidHotspotV2SettingBySite

> OperationResponseWithoutResult UpdateSsidHotspotV2SettingBySite(ctx, omadacId, siteId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()

Update SSID Hotspot2.0 config by site



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
	ssidId := "ssidId_example" // string | ssidId
	updateSsidHotspotV2SettingOpenApiVO := *openapiclient.NewUpdateSsidHotspotV2SettingOpenApiVO(false) // UpdateSsidHotspotV2SettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidHotspotV2SettingBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidHotspotV2SettingBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidHotspotV2SettingBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidHotspotV2SettingBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | ssidId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidHotspotV2SettingBySiteRequest struct via the builder pattern


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


## UpdateSsidLoadBalanceConfig

> OperationResponseWithoutResult UpdateSsidLoadBalanceConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateSsidLoadBalanceOpenApiVO(updateSsidLoadBalanceOpenApiVO).Execute()

Update SSID load balance config



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
	updateSsidLoadBalanceOpenApiVO := *openapiclient.NewUpdateSsidLoadBalanceOpenApiVO(false) // UpdateSsidLoadBalanceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidLoadBalanceConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateSsidLoadBalanceOpenApiVO(updateSsidLoadBalanceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidLoadBalanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidLoadBalanceConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidLoadBalanceConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateSsidLoadBalanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidLoadBalanceOpenApiVO** | [**UpdateSsidLoadBalanceOpenApiVO**](UpdateSsidLoadBalanceOpenApiVO.md) |  | 

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


## UpdateSsidLoadBalanceConfigBySite

> OperationResponseWithoutResult UpdateSsidLoadBalanceConfigBySite(ctx, omadacId, siteId, ssidId).UpdateSsidLoadBalanceOpenApiVO(updateSsidLoadBalanceOpenApiVO).Execute()

Update SSID load balance config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidLoadBalanceOpenApiVO := *openapiclient.NewUpdateSsidLoadBalanceOpenApiVO(false) // UpdateSsidLoadBalanceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidLoadBalanceConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidLoadBalanceOpenApiVO(updateSsidLoadBalanceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidLoadBalanceConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidLoadBalanceConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidLoadBalanceConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidLoadBalanceConfigBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateSsidLoadBalanceOpenApiVO** | [**UpdateSsidLoadBalanceOpenApiVO**](UpdateSsidLoadBalanceOpenApiVO.md) |  | 

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


## UpdateSsidMacFilterConfigBySite

> OperationResponseWithoutResult UpdateSsidMacFilterConfigBySite(ctx, omadacId, siteId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()

Update SSID mac filter config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMacFilterOpenApiVO := *openapiclient.NewUpdateSsidMacFilterOpenApiVO(false) // UpdateSsidMacFilterOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidMacFilterConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidMacFilterConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidMacFilterConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidMacFilterConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidMacFilterConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidMultiCastConfigBySite

> OperationResponseWithoutResult UpdateSsidMultiCastConfigBySite(ctx, omadacId, siteId, ssidId).UpdateSsidMultiCastOpenApiVO(updateSsidMultiCastOpenApiVO).Execute()

Update SSID Multicast/Broadcast management config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMultiCastOpenApiVO := *openapiclient.NewUpdateSsidMultiCastOpenApiVO(false, int32(123), false, false, false) // UpdateSsidMultiCastOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidMultiCastConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidMultiCastOpenApiVO(updateSsidMultiCastOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidMultiCastConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidMultiCastConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidMultiCastConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidMultiCastConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidRateControlConfigBySite

> OperationResponseWithoutResult UpdateSsidRateControlConfigBySite(ctx, omadacId, siteId, ssidId).UpdateSsidRateControlOpenApiVO(updateSsidRateControlOpenApiVO).Execute()

Update SSID 802.11 rate control config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidRateControlOpenApiVO := *openapiclient.NewUpdateSsidRateControlOpenApiVO(false, false) // UpdateSsidRateControlOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidRateControlConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidRateControlOpenApiVO(updateSsidRateControlOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidRateControlConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidRateControlConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidRateControlConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidRateControlConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidRateLimitConfigBySite

> OperationResponseWithoutResult UpdateSsidRateLimitConfigBySite(ctx, omadacId, siteId, ssidId).UpdateSsidRateLimitOpenApiVO(updateSsidRateLimitOpenApiVO).Execute()

Update SSID rate limit config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidRateLimitOpenApiVO := *openapiclient.NewUpdateSsidRateLimitOpenApiVO() // UpdateSsidRateLimitOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidRateLimitConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidRateLimitOpenApiVO(updateSsidRateLimitOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidRateLimitConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidRateLimitConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidRateLimitConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidRateLimitConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidWifiCallingConfig

> OperationResponseWithoutResult UpdateSsidWifiCallingConfig(ctx, omadacId, siteId, wlanId, ssidId).UpdateWifiCallingOpenApiVO(updateWifiCallingOpenApiVO).Execute()

Update SSID wifi calling config



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
	updateWifiCallingOpenApiVO := *openapiclient.NewUpdateWifiCallingOpenApiVO(false) // UpdateWifiCallingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidWifiCallingConfig(context.Background(), omadacId, siteId, wlanId, ssidId).UpdateWifiCallingOpenApiVO(updateWifiCallingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidWifiCallingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidWifiCallingConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidWifiCallingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateSsidWifiCallingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateWifiCallingOpenApiVO** | [**UpdateWifiCallingOpenApiVO**](UpdateWifiCallingOpenApiVO.md) |  | 

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


## UpdateSsidWifiCallingConfigBySite

> OperationResponseWithoutResult UpdateSsidWifiCallingConfigBySite(ctx, omadacId, siteId, ssidId).UpdateWifiCallingOpenApiVO(updateWifiCallingOpenApiVO).Execute()

Update SSID wifi calling config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateWifiCallingOpenApiVO := *openapiclient.NewUpdateWifiCallingOpenApiVO(false) // UpdateWifiCallingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidWifiCallingConfigBySite(context.Background(), omadacId, siteId, ssidId).UpdateWifiCallingOpenApiVO(updateWifiCallingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidWifiCallingConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidWifiCallingConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidWifiCallingConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidWifiCallingConfigBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateWifiCallingOpenApiVO** | [**UpdateWifiCallingOpenApiVO**](UpdateWifiCallingOpenApiVO.md) |  | 

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


## UpdateSsidWlanScheduleBySite

> OperationResponseWithoutResult UpdateSsidWlanScheduleBySite(ctx, omadacId, siteId, ssidId).UpdateSsidWlanScheduleOpenApiVO(updateSsidWlanScheduleOpenApiVO).Execute()

Update SSID WLAN schedule config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidWlanScheduleOpenApiVO := *openapiclient.NewUpdateSsidWlanScheduleOpenApiVO(false) // UpdateSsidWlanScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkAPI.UpdateSsidWlanScheduleBySite(context.Background(), omadacId, siteId, ssidId).UpdateSsidWlanScheduleOpenApiVO(updateSsidWlanScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkAPI.UpdateSsidWlanScheduleBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidWlanScheduleBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkAPI.UpdateSsidWlanScheduleBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidWlanScheduleBySiteRequest struct via the builder pattern


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

Modify an existing WLAN Group



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


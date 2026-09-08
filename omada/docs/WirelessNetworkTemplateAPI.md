# WirelessNetworkTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountApGroupsTemplate**](WirelessNetworkTemplateAPI.md#countapgroupstemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/count-ap-groups | Check if the number of AP Groups template is out of limit
[**CountWlansTemplate**](WirelessNetworkTemplateAPI.md#countwlanstemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/cmd/countWlans | Check if the number of WLAN Groups is out of limit
[**CreateApGroupTemplate**](WirelessNetworkTemplateAPI.md#createapgrouptemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ap-groups | Create new AP Group template
[**CreateSsidTemplate**](WirelessNetworkTemplateAPI.md#createssidtemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids | Create new SSID template
[**CreateSsidTemplateV2**](WirelessNetworkTemplateAPI.md#createssidtemplatev2) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids | Create new SSID template v2
[**CreateWlanGroupTemplate**](WirelessNetworkTemplateAPI.md#createwlangrouptemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans | Create new WLAN Group template
[**DeleteApGroupTemplate**](WirelessNetworkTemplateAPI.md#deleteapgrouptemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ap-groups/{apGroupId} | Delete an existing AP Group template
[**DeleteSsidTemplate**](WirelessNetworkTemplateAPI.md#deletessidtemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId} | Delete an existing SSID template
[**DeleteSsidTemplateV2**](WirelessNetworkTemplateAPI.md#deletessidtemplatev2) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId} | Delete an existing SSID template
[**DeleteWlanGroupTemplate**](WirelessNetworkTemplateAPI.md#deletewlangrouptemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId} | Delete an existing WLAN Group template
[**GetApGroupTemplateList**](WirelessNetworkTemplateAPI.md#getapgrouptemplatelist) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ap-groups | Get AP Group template list
[**GetGridWlanGroupTemplate**](WirelessNetworkTemplateAPI.md#getgridwlangrouptemplate) | **Get** /openapi/v2/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans | Get WLAN Group template list paging query
[**GetSsidDetailTemplate**](WirelessNetworkTemplateAPI.md#getssiddetailtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId} | Get SSID template detail info
[**GetSsidDetailTemplateV2**](WirelessNetworkTemplateAPI.md#getssiddetailtemplatev2) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId} | Get SSID template detail info
[**GetSsidListTemplate**](WirelessNetworkTemplateAPI.md#getssidlisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids | Get SSID template list
[**GetSsidListTemplateV2**](WirelessNetworkTemplateAPI.md#getssidlisttemplatev2) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids | Get SSID template list v2
[**GetWlanGroupListTemplate**](WirelessNetworkTemplateAPI.md#getwlangrouplisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans | Get WLAN Group template list
[**QuerySsidDuplicateNameBySite**](WirelessNetworkTemplateAPI.md#queryssidduplicatenamebysite) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/duplicate-name | Query Template SSIDs with the same name
[**UpdateApGroupTemplate**](WirelessNetworkTemplateAPI.md#updateapgrouptemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ap-groups/{apGroupId} | Modify an existing AP Group template
[**UpdateSsidBandSteerConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessidbandsteerconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-band-steer | Update SSID template band steer config
[**UpdateSsidBasicConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessidbasicconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-basic-config | Update SSID template basic config
[**UpdateSsidDhcpOptionConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessiddhcpoptionconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-dhcp-option | Update SSID template DHCP option 82 config
[**UpdateSsidEnableStatusBySite**](WirelessNetworkTemplateAPI.md#updatessidenablestatusbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/enable | Update SSID Template Enable Status
[**UpdateSsidHotspotV2SettingTemplate**](WirelessNetworkTemplateAPI.md#updatessidhotspotv2settingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-hotspotv2 | Update SSID template Hotspot2.0 config
[**UpdateSsidLoadBalanceConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessidloadbalanceconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-load-balance | Update SSID template load balance config
[**UpdateSsidMacFilterConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessidmacfilterconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-mac-filter | Update SSID template mac filter config
[**UpdateSsidMultiCastConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessidmulticastconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-multicast-config | Update SSID template Multicast/Broadcast management config
[**UpdateSsidRateControlConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessidratecontrolconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-rate-control | Update SSID template 802.11 rate control config
[**UpdateSsidRateLimitConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessidratelimitconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-rate-limit | Update SSID template rate limit config
[**UpdateSsidTemplateBandSteerConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplatebandsteerconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/band-steer | Update SSID Template band steer config by site
[**UpdateSsidTemplateBasicConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplatebasicconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/basic-config | Update SSID Template basic config by site
[**UpdateSsidTemplateDhcpOptionConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplatedhcpoptionconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/dhcp-option | Update SSID Template DHCP option 82 config by site
[**UpdateSsidTemplateHotspotV2SettingBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplatehotspotv2settingbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/hotspotv2 | Update SSID Template Hotspot2.0 config by site
[**UpdateSsidTemplateLoadBalanceConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplateloadbalanceconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/load-balance | Update SSID Template load balance config by site
[**UpdateSsidTemplateMacFilterConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplatemacfilterconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/mac-filter | Update SSID Template mac filter config by site
[**UpdateSsidTemplateMultiCastConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplatemulticastconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/multicast-config | Update SSID Template Multicast/Broadcast management config by site
[**UpdateSsidTemplateRateControlConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplateratecontrolconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/rate-control | Update SSID Template 802.11 rate control config by site
[**UpdateSsidTemplateRateLimitConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplateratelimitconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/rate-limit | Update SSID Template rate limit config by site
[**UpdateSsidTemplateWifiCallingConfigBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplatewificallingconfigbysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/wifi-calling | Update SSID Template wifi calling config by site
[**UpdateSsidTemplateWlanScheduleBySite**](WirelessNetworkTemplateAPI.md#updatessidtemplatewlanschedulebysite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/ssids/{ssidId}/wlan-schedule | Update SSID Template WLAN schedule config by site
[**UpdateSsidWifiCallingConfigTemplate**](WirelessNetworkTemplateAPI.md#updatessidwificallingconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-wifi-calling | Update SSID template wifi calling config
[**UpdateSsidWlanScheduleTemplate**](WirelessNetworkTemplateAPI.md#updatessidwlanscheduletemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId}/ssids/{ssidId}/update-wlan-schedule | Update SSID template WLAN schedule config
[**UpdateWlanGroupTemplate**](WirelessNetworkTemplateAPI.md#updatewlangrouptemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wireless-network/wlans/{wlanId} | Modify an existing WLAN Group template



## CountApGroupsTemplate

> OperationResponseApGroupStatusOpenApiVO CountApGroupsTemplate(ctx, omadacId, siteTemplateId).Execute()

Check if the number of AP Groups template is out of limit



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
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.CountApGroupsTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.CountApGroupsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountApGroupsTemplate`: OperationResponseApGroupStatusOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.CountApGroupsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountApGroupsTemplateRequest struct via the builder pattern


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


## CountWlansTemplate

> OperationResponseWlanGroupStatusOpenApiVO CountWlansTemplate(ctx, omadacId, siteTemplateId).Execute()

Check if the number of WLAN Groups is out of limit



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
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountWlansTemplateRequest struct via the builder pattern


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


## CreateApGroupTemplate

> OperationResponseAddApGroupResultVO CreateApGroupTemplate(ctx, omadacId, siteTemplateId).CreateApGroupOpenApiVO(createApGroupOpenApiVO).Execute()

Create new AP Group template



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
	createApGroupOpenApiVO := *openapiclient.NewCreateApGroupOpenApiVO("Name_example") // CreateApGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.CreateApGroupTemplate(context.Background(), omadacId, siteTemplateId).CreateApGroupOpenApiVO(createApGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.CreateApGroupTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApGroupTemplate`: OperationResponseAddApGroupResultVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.CreateApGroupTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApGroupTemplateRequest struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSsidTemplateV2

> OperationResponseCreateSsidResultVO CreateSsidTemplateV2(ctx, omadacId, siteTemplateId).CreateSsidOpenApiVO(createSsidOpenApiVO).Execute()

Create new SSID template v2



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
	createSsidOpenApiVO := *openapiclient.NewCreateSsidOpenApiVO(int32(123), false, int32(123), false, false, false, false, "Name_example", int32(123), int32(123), false) // CreateSsidOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.CreateSsidTemplateV2(context.Background(), omadacId, siteTemplateId).CreateSsidOpenApiVO(createSsidOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.CreateSsidTemplateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSsidTemplateV2`: OperationResponseCreateSsidResultVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.CreateSsidTemplateV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSsidTemplateV2Request struct via the builder pattern


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


## CreateWlanGroupTemplate

> OperationResponse CreateWlanGroupTemplate(ctx, omadacId, siteTemplateId).CreateWlanGroupOpenApiVO(createWlanGroupOpenApiVO).Execute()

Create new WLAN Group template



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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApGroupTemplate

> OperationResponseWithoutResult DeleteApGroupTemplate(ctx, omadacId, siteTemplateId, apGroupId).Execute()

Delete an existing AP Group template



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
	apGroupId := "apGroupId_example" // string | AP GROUP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.DeleteApGroupTemplate(context.Background(), omadacId, siteTemplateId, apGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.DeleteApGroupTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApGroupTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.DeleteApGroupTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**apGroupId** | **string** | AP GROUP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApGroupTemplateRequest struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSsidTemplateV2

> OperationResponseWithoutResult DeleteSsidTemplateV2(ctx, omadacId, siteTemplateId, ssidId).Execute()

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
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.DeleteSsidTemplateV2(context.Background(), omadacId, siteTemplateId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.DeleteSsidTemplateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSsidTemplateV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.DeleteSsidTemplateV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSsidTemplateV2Request struct via the builder pattern


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


## DeleteWlanGroupTemplate

> OperationResponseWithoutResult DeleteWlanGroupTemplate(ctx, omadacId, siteTemplateId, wlanId).Execute()

Delete an existing WLAN Group template



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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApGroupTemplateList

> OperationResponseApGroupGridVOApGroupOpenApiVO GetApGroupTemplateList(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get AP Group template list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.GetApGroupTemplateList(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.GetApGroupTemplateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApGroupTemplateList`: OperationResponseApGroupGridVOApGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.GetApGroupTemplateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApGroupTemplateListRequest struct via the builder pattern


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


## GetGridWlanGroupTemplate

> OperationResponseWlanGroupGridOpenApiVO GetGridWlanGroupTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get WLAN Group template list paging query



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
**siteTemplateId** | **string** | Site Template ID | 

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidDetailTemplateV2

> OperationResponseSsidDetailOpenApiVO GetSsidDetailTemplateV2(ctx, omadacId, siteTemplateId, ssidId).Execute()

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
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.GetSsidDetailTemplateV2(context.Background(), omadacId, siteTemplateId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.GetSsidDetailTemplateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidDetailTemplateV2`: OperationResponseSsidDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.GetSsidDetailTemplateV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidDetailTemplateV2Request struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidListTemplateV2

> OperationResponseGridVOSsidOpenApiVO GetSsidListTemplateV2(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get SSID template list v2



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.GetSsidListTemplateV2(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.GetSsidListTemplateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidListTemplateV2`: OperationResponseGridVOSsidOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.GetSsidListTemplateV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidListTemplateV2Request struct via the builder pattern


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


## GetWlanGroupListTemplate

> OperationResponseListWlanGroupOpenApiVO GetWlanGroupListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get WLAN Group template list



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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySsidDuplicateNameBySite

> OperationResponseDuplicateSsidOpenApiVO QuerySsidDuplicateNameBySite(ctx, omadacId, siteTemplateId).Execute()

Query Template SSIDs with the same name



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
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.QuerySsidDuplicateNameBySite(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.QuerySsidDuplicateNameBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySsidDuplicateNameBySite`: OperationResponseDuplicateSsidOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.QuerySsidDuplicateNameBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySsidDuplicateNameBySiteRequest struct via the builder pattern


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


## UpdateApGroupTemplate

> OperationResponseWithoutResult UpdateApGroupTemplate(ctx, omadacId, siteTemplateId, apGroupId).UpdateApGroupOpenApiVO(updateApGroupOpenApiVO).Execute()

Modify an existing AP Group template



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
	apGroupId := "apGroupId_example" // string | AP GROUP ID
	updateApGroupOpenApiVO := *openapiclient.NewUpdateApGroupOpenApiVO("Name_example") // UpdateApGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateApGroupTemplate(context.Background(), omadacId, siteTemplateId, apGroupId).UpdateApGroupOpenApiVO(updateApGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateApGroupTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApGroupTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateApGroupTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**apGroupId** | **string** | AP GROUP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApGroupTemplateRequest struct via the builder pattern


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


## UpdateSsidBandSteerConfigTemplate

> OperationResponseWithoutResult UpdateSsidBandSteerConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidBandSteerOpenApiVO(updateSsidBandSteerOpenApiVO).Execute()

Update SSID template band steer config



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
	updateSsidBandSteerOpenApiVO := *openapiclient.NewUpdateSsidBandSteerOpenApiVO(int32(123)) // UpdateSsidBandSteerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidBandSteerConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidBandSteerOpenApiVO(updateSsidBandSteerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidBandSteerConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidBandSteerConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidBandSteerConfigTemplate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateSsidBandSteerConfigTemplateRequest struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSsidEnableStatusBySite

> OperationResponseWithoutResult UpdateSsidEnableStatusBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidEnableStatusOpenApiVO(updateSsidEnableStatusOpenApiVO).Execute()

Update SSID Template Enable Status



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidEnableStatusOpenApiVO := *openapiclient.NewUpdateSsidEnableStatusOpenApiVO(false) // UpdateSsidEnableStatusOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidEnableStatusBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidEnableStatusOpenApiVO(updateSsidEnableStatusOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidEnableStatusBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidEnableStatusBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidEnableStatusBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidEnableStatusBySiteRequest struct via the builder pattern


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


## UpdateSsidHotspotV2SettingTemplate

> OperationResponseWithoutResult UpdateSsidHotspotV2SettingTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	wlanId := "wlanId_example" // string | WLAN ID
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidHotspotV2SettingOpenApiVO := *openapiclient.NewUpdateSsidHotspotV2SettingOpenApiVO(false) // UpdateSsidHotspotV2SettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidHotspotV2SettingTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()
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
**siteTemplateId** | **string** | Site Template ID | 
**wlanId** | **string** | WLAN ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidHotspotV2SettingTemplateRequest struct via the builder pattern


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


## UpdateSsidLoadBalanceConfigTemplate

> OperationResponseWithoutResult UpdateSsidLoadBalanceConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidLoadBalanceOpenApiVO(updateSsidLoadBalanceOpenApiVO).Execute()

Update SSID template load balance config



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
	updateSsidLoadBalanceOpenApiVO := *openapiclient.NewUpdateSsidLoadBalanceOpenApiVO(false) // UpdateSsidLoadBalanceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidLoadBalanceConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateSsidLoadBalanceOpenApiVO(updateSsidLoadBalanceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidLoadBalanceConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidLoadBalanceConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidLoadBalanceConfigTemplate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateSsidLoadBalanceConfigTemplateRequest struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSsidTemplateBandSteerConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateBandSteerConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidBandSteerOpenApiVO(updateSsidBandSteerOpenApiVO).Execute()

Update SSID Template band steer config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidBandSteerOpenApiVO := *openapiclient.NewUpdateSsidBandSteerOpenApiVO(int32(123)) // UpdateSsidBandSteerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateBandSteerConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidBandSteerOpenApiVO(updateSsidBandSteerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateBandSteerConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateBandSteerConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateBandSteerConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateBandSteerConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateBasicConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateBasicConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidBasicConfigOpenApiVO(updateSsidBasicConfigOpenApiVO).Execute()

Update SSID Template basic config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidBasicConfigOpenApiVO := *openapiclient.NewUpdateSsidBasicConfigOpenApiVO(int32(123), false, false, false, false, "Name_example", int32(123), int32(123), false) // UpdateSsidBasicConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateBasicConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidBasicConfigOpenApiVO(updateSsidBasicConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateBasicConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateBasicConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateBasicConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateBasicConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateDhcpOptionConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateDhcpOptionConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidDhcpOptionOpenApiVO(updateSsidDhcpOptionOpenApiVO).Execute()

Update SSID Template DHCP option 82 config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidDhcpOptionOpenApiVO := *openapiclient.NewUpdateSsidDhcpOptionOpenApiVO(false) // UpdateSsidDhcpOptionOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateDhcpOptionConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidDhcpOptionOpenApiVO(updateSsidDhcpOptionOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateDhcpOptionConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateDhcpOptionConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateDhcpOptionConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateDhcpOptionConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateHotspotV2SettingBySite

> OperationResponseWithoutResult UpdateSsidTemplateHotspotV2SettingBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()

Update SSID Template Hotspot2.0 config by site



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
	ssidId := "ssidId_example" // string | ssidId
	updateSsidHotspotV2SettingOpenApiVO := *openapiclient.NewUpdateSsidHotspotV2SettingOpenApiVO(false) // UpdateSsidHotspotV2SettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateHotspotV2SettingBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidHotspotV2SettingOpenApiVO(updateSsidHotspotV2SettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateHotspotV2SettingBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateHotspotV2SettingBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateHotspotV2SettingBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | ssidId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateHotspotV2SettingBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateLoadBalanceConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateLoadBalanceConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidLoadBalanceOpenApiVO(updateSsidLoadBalanceOpenApiVO).Execute()

Update SSID Template load balance config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidLoadBalanceOpenApiVO := *openapiclient.NewUpdateSsidLoadBalanceOpenApiVO(false) // UpdateSsidLoadBalanceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateLoadBalanceConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidLoadBalanceOpenApiVO(updateSsidLoadBalanceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateLoadBalanceConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateLoadBalanceConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateLoadBalanceConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateLoadBalanceConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateMacFilterConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateMacFilterConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()

Update SSID Template mac filter config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMacFilterOpenApiVO := *openapiclient.NewUpdateSsidMacFilterOpenApiVO(false) // UpdateSsidMacFilterOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateMacFilterConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateMacFilterConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateMacFilterConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateMacFilterConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateMacFilterConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateMultiCastConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateMultiCastConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidMultiCastOpenApiVO(updateSsidMultiCastOpenApiVO).Execute()

Update SSID Template Multicast/Broadcast management config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMultiCastOpenApiVO := *openapiclient.NewUpdateSsidMultiCastOpenApiVO(false, int32(123), false, false, false) // UpdateSsidMultiCastOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateMultiCastConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidMultiCastOpenApiVO(updateSsidMultiCastOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateMultiCastConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateMultiCastConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateMultiCastConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateMultiCastConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateRateControlConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateRateControlConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidRateControlOpenApiVO(updateSsidRateControlOpenApiVO).Execute()

Update SSID Template 802.11 rate control config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidRateControlOpenApiVO := *openapiclient.NewUpdateSsidRateControlOpenApiVO(false, false) // UpdateSsidRateControlOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateRateControlConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidRateControlOpenApiVO(updateSsidRateControlOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateRateControlConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateRateControlConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateRateControlConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateRateControlConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateRateLimitConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateRateLimitConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidRateLimitOpenApiVO(updateSsidRateLimitOpenApiVO).Execute()

Update SSID Template rate limit config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidRateLimitOpenApiVO := *openapiclient.NewUpdateSsidRateLimitOpenApiVO() // UpdateSsidRateLimitOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateRateLimitConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidRateLimitOpenApiVO(updateSsidRateLimitOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateRateLimitConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateRateLimitConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateRateLimitConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateRateLimitConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateWifiCallingConfigBySite

> OperationResponseWithoutResult UpdateSsidTemplateWifiCallingConfigBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateWifiCallingOpenApiVO(updateWifiCallingOpenApiVO).Execute()

Update SSID Template wifi calling config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateWifiCallingOpenApiVO := *openapiclient.NewUpdateWifiCallingOpenApiVO(false) // UpdateWifiCallingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateWifiCallingConfigBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateWifiCallingOpenApiVO(updateWifiCallingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateWifiCallingConfigBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateWifiCallingConfigBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateWifiCallingConfigBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateWifiCallingConfigBySiteRequest struct via the builder pattern


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


## UpdateSsidTemplateWlanScheduleBySite

> OperationResponseWithoutResult UpdateSsidTemplateWlanScheduleBySite(ctx, omadacId, siteTemplateId, ssidId).UpdateSsidWlanScheduleOpenApiVO(updateSsidWlanScheduleOpenApiVO).Execute()

Update SSID Template WLAN schedule config by site



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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidWlanScheduleOpenApiVO := *openapiclient.NewUpdateSsidWlanScheduleOpenApiVO(false) // UpdateSsidWlanScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidTemplateWlanScheduleBySite(context.Background(), omadacId, siteTemplateId, ssidId).UpdateSsidWlanScheduleOpenApiVO(updateSsidWlanScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidTemplateWlanScheduleBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidTemplateWlanScheduleBySite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidTemplateWlanScheduleBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsidTemplateWlanScheduleBySiteRequest struct via the builder pattern


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


## UpdateSsidWifiCallingConfigTemplate

> OperationResponseWithoutResult UpdateSsidWifiCallingConfigTemplate(ctx, omadacId, siteTemplateId, wlanId, ssidId).UpdateWifiCallingOpenApiVO(updateWifiCallingOpenApiVO).Execute()

Update SSID template wifi calling config



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
	updateWifiCallingOpenApiVO := *openapiclient.NewUpdateWifiCallingOpenApiVO(false) // UpdateWifiCallingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessNetworkTemplateAPI.UpdateSsidWifiCallingConfigTemplate(context.Background(), omadacId, siteTemplateId, wlanId, ssidId).UpdateWifiCallingOpenApiVO(updateWifiCallingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessNetworkTemplateAPI.UpdateSsidWifiCallingConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsidWifiCallingConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WirelessNetworkTemplateAPI.UpdateSsidWifiCallingConfigTemplate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateSsidWifiCallingConfigTemplateRequest struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWlanGroupTemplate

> OperationResponseWithoutResult UpdateWlanGroupTemplate(ctx, omadacId, siteTemplateId, wlanId).UpdateWlanGroupOpenApiVO(updateWlanGroupOpenApiVO).Execute()

Modify an existing WLAN Group template



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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


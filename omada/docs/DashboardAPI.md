# DashboardAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchConfigTab**](DashboardAPI.md#batchconfigtab) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/multi-tabs/config | Batch set tab config
[**ConfigBandwidthForWanPorts**](DashboardAPI.md#configbandwidthforwanports) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wan/bandwidth | Set site dashboard wan bandwidth
[**CreateTab**](DashboardAPI.md#createtab) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/tabs | Create new tab
[**GetActiveAps**](DashboardAPI.md#getactiveaps) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/most-active-eaps | Get the most active eap list
[**GetActiveApsV2**](DashboardAPI.md#getactiveapsv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/dashboard/most-active-eaps | Get the most active eap list
[**GetActiveSwitches**](DashboardAPI.md#getactiveswitches) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/most-active-switches | Get most active switch list
[**GetAllNetworkActivity**](DashboardAPI.md#getallnetworkactivity) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/allNetworkActivity | Get grid dashboard open network activity statistic list
[**GetCardTopology**](DashboardAPI.md#getcardtopology) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/card/overview | Get network overview
[**GetChannels**](DashboardAPI.md#getchannels) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/channels | Get channel distribution and usage
[**GetGatewayIspLoad**](DashboardAPI.md#getgatewayispload) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/gateway/isp/load | Get site dashboard isp load inform
[**GetGridDashboardIpsecTunnelStats**](DashboardAPI.md#getgriddashboardipsectunnelstats) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/lpset-tunnel-stats | Get grid dashboard lpsec tunnel statistic list
[**GetGridDashboardOpenVpnTunnelStats**](DashboardAPI.md#getgriddashboardopenvpntunnelstats) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/open-vpn-tunnel-stats | Get grid dashboard open vpn tunnel statistic list
[**GetGridDashboardTunnelStats**](DashboardAPI.md#getgriddashboardtunnelstats) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/vpn-tunnel-stats | Get grid dashboard tunnel statistic list
[**GetInterference**](DashboardAPI.md#getinterference) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/top-interference | Get top interference
[**GetIspLoad**](DashboardAPI.md#getispload) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/isp-load | Get isp load info
[**GetOverview**](DashboardAPI.md#getoverview) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/overview-diagram | Get site overview diagram info
[**GetPoeUsage**](DashboardAPI.md#getpoeusage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/poe-usage | Get poe usage
[**GetRetryAndDroppedRate**](DashboardAPI.md#getretryanddroppedrate) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/retry-dropped-rate | Get retried rate and dropped rate
[**GetSpeedTestV2Result**](DashboardAPI.md#getspeedtestv2result) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/speedTestResult | Get Speed Test Result
[**GetSpeedTestV2ResultDateList**](DashboardAPI.md#getspeedtestv2resultdatelist) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/speedTestResult/dateList | Get date list of speed test results
[**GetSwitchSummary**](DashboardAPI.md#getswitchsummary) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/switch-summary | Get switch summary
[**GetTop5Aps**](DashboardAPI.md#gettop5aps) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/top-aps | Get Top 5 Aps
[**GetTopCpuUsageWithTimeRange**](DashboardAPI.md#gettopcpuusagewithtimerange) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/top-device-cpu-usage | Get top device cpu usage
[**GetTopMemoryUsageWithTimeRange**](DashboardAPI.md#gettopmemoryusagewithtimerange) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/top-device-memory-usage | Get top device memory usage
[**GetTrafficActivities**](DashboardAPI.md#gettrafficactivities) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/traffic-activities | Get traffic activity
[**GetTrafficDistribution**](DashboardAPI.md#gettrafficdistribution) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/traffic-distribution | Get traffic distribution
[**GetWifiSummary**](DashboardAPI.md#getwifisummary) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/wifi-summary | Get wifi summary
[**ListAllTabs**](DashboardAPI.md#listalltabs) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/tabs | Get tab list
[**ListTabs**](DashboardAPI.md#listtabs) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/without-overall-tabs | Get tab without overall tab list
[**RemoveTab**](DashboardAPI.md#removetab) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/tabs/{tabId} | Delete an existing tab
[**StartSpeedTestV2**](DashboardAPI.md#startspeedtestv2) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/speedTest | Start Speed Test
[**UpdateTab**](DashboardAPI.md#updatetab) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/tabs/{tabId} | Modify an existing tab



## BatchConfigTab

> OperationResponseWithoutResult BatchConfigTab(ctx, omadacId, siteId).BatchEditTabs(batchEditTabs).Execute()

Batch set tab config



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
	batchEditTabs := *openapiclient.NewBatchEditTabs() // BatchEditTabs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.BatchConfigTab(context.Background(), omadacId, siteId).BatchEditTabs(batchEditTabs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.BatchConfigTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchConfigTab`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.BatchConfigTab`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchConfigTabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchEditTabs** | [**BatchEditTabs**](BatchEditTabs.md) |  | 

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


## ConfigBandwidthForWanPorts

> OperationResponseWithoutResult ConfigBandwidthForWanPorts(ctx, omadacId, siteId).WanPortBandwidthVO(wanPortBandwidthVO).Execute()

Set site dashboard wan bandwidth



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
	wanPortBandwidthVO := *openapiclient.NewWanPortBandwidthVO() // WanPortBandwidthVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.ConfigBandwidthForWanPorts(context.Background(), omadacId, siteId).WanPortBandwidthVO(wanPortBandwidthVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.ConfigBandwidthForWanPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigBandwidthForWanPorts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.ConfigBandwidthForWanPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigBandwidthForWanPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanPortBandwidthVO** | [**WanPortBandwidthVO**](WanPortBandwidthVO.md) |  | 

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


## CreateTab

> OperationResponseWithoutResult CreateTab(ctx, omadacId, siteId).CreateTabOpenApiVO(createTabOpenApiVO).Execute()

Create new tab



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
	createTabOpenApiVO := *openapiclient.NewCreateTabOpenApiVO() // CreateTabOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.CreateTab(context.Background(), omadacId, siteId).CreateTabOpenApiVO(createTabOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.CreateTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTab`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.CreateTab`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTabOpenApiVO** | [**CreateTabOpenApiVO**](CreateTabOpenApiVO.md) |  | 

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


## GetActiveAps

> OperationResponseListActiveDevice GetActiveAps(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get the most active eap list



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetActiveAps(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetActiveAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveAps`: OperationResponseListActiveDevice
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetActiveAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListActiveDevice**](OperationResponseListActiveDevice.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveApsV2

> OperationResponseGetActiveDeviceV2OpenApiVO GetActiveApsV2(ctx, omadacId, siteId).DeviceNum(deviceNum).Start(start).End(end).Execute()

Get the most active eap list



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
	deviceNum := int32(56) // int32 | The number of most active APs acquired.
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetActiveApsV2(context.Background(), omadacId, siteId).DeviceNum(deviceNum).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetActiveApsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveApsV2`: OperationResponseGetActiveDeviceV2OpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetActiveApsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveApsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceNum** | **int32** | The number of most active APs acquired. | 
 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseGetActiveDeviceV2OpenApiVO**](OperationResponseGetActiveDeviceV2OpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveSwitches

> OperationResponseListActiveDevice GetActiveSwitches(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get most active switch list



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetActiveSwitches(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetActiveSwitches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveSwitches`: OperationResponseListActiveDevice
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetActiveSwitches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveSwitchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListActiveDevice**](OperationResponseListActiveDevice.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNetworkActivity

> OperationResponseNetworkActivityVO GetAllNetworkActivity(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get grid dashboard open network activity statistic list



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetAllNetworkActivity(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetAllNetworkActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllNetworkActivity`: OperationResponseNetworkActivityVO
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetAllNetworkActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNetworkActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseNetworkActivityVO**](OperationResponseNetworkActivityVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardTopology

> OperationResponseListCardOverviewOpenApiVO GetCardTopology(ctx, omadacId, siteId).Execute()

Get network overview



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
	resp, r, err := apiClient.DashboardAPI.GetCardTopology(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetCardTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCardTopology`: OperationResponseListCardOverviewOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetCardTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListCardOverviewOpenApiVO**](OperationResponseListCardOverviewOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannels

> OperationResponseApChannelStats GetChannels(ctx, omadacId, siteId).Execute()

Get channel distribution and usage



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
	resp, r, err := apiClient.DashboardAPI.GetChannels(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannels`: OperationResponseApChannelStats
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseApChannelStats**](OperationResponseApChannelStats.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayIspLoad

> OperationResponseDashboardIspLoadDetailVO GetGatewayIspLoad(ctx, omadacId, siteId).Execute()

Get site dashboard isp load inform



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
	resp, r, err := apiClient.DashboardAPI.GetGatewayIspLoad(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetGatewayIspLoad``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGatewayIspLoad`: OperationResponseDashboardIspLoadDetailVO
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetGatewayIspLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayIspLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDashboardIspLoadDetailVO**](OperationResponseDashboardIspLoadDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDashboardIpsecTunnelStats

> OperationResponseListIpsecVpnStats GetGridDashboardIpsecTunnelStats(ctx, omadacId, siteId).Execute()

Get grid dashboard lpsec tunnel statistic list



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
	resp, r, err := apiClient.DashboardAPI.GetGridDashboardIpsecTunnelStats(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetGridDashboardIpsecTunnelStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDashboardIpsecTunnelStats`: OperationResponseListIpsecVpnStats
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetGridDashboardIpsecTunnelStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDashboardIpsecTunnelStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListIpsecVpnStats**](OperationResponseListIpsecVpnStats.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDashboardOpenVpnTunnelStats

> OperationResponseListDashboardVpnStats GetGridDashboardOpenVpnTunnelStats(ctx, omadacId, siteId).Type_(type_).Execute()

Get grid dashboard open vpn tunnel statistic list



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
	type_ := int32(56) // int32 | type: 0:Server, 1:Client

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetGridDashboardOpenVpnTunnelStats(context.Background(), omadacId, siteId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetGridDashboardOpenVpnTunnelStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDashboardOpenVpnTunnelStats`: OperationResponseListDashboardVpnStats
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetGridDashboardOpenVpnTunnelStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDashboardOpenVpnTunnelStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **int32** | type: 0:Server, 1:Client | 

### Return type

[**OperationResponseListDashboardVpnStats**](OperationResponseListDashboardVpnStats.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDashboardTunnelStats

> OperationResponseListDashboardVpnStats GetGridDashboardTunnelStats(ctx, omadacId, siteId).Type_(type_).Execute()

Get grid dashboard tunnel statistic list



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
	type_ := int32(56) // int32 | Type should be a value as follows: 0:Server,1:Client

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetGridDashboardTunnelStats(context.Background(), omadacId, siteId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetGridDashboardTunnelStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDashboardTunnelStats`: OperationResponseListDashboardVpnStats
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetGridDashboardTunnelStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDashboardTunnelStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **int32** | Type should be a value as follows: 0:Server,1:Client | 

### Return type

[**OperationResponseListDashboardVpnStats**](OperationResponseListDashboardVpnStats.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterference

> OperationResponseApInterferences GetInterference(ctx, omadacId, siteId).Execute()

Get top interference



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
	resp, r, err := apiClient.DashboardAPI.GetInterference(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetInterference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterference`: OperationResponseApInterferences
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetInterference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseApInterferences**](OperationResponseApInterferences.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIspLoad

> OperationResponseListIspLoad GetIspLoad(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get isp load info



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetIspLoad(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetIspLoad``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIspLoad`: OperationResponseListIspLoad
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetIspLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIspLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListIspLoad**](OperationResponseListIspLoad.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverview

> OperationResponseGetDashboardOverview GetOverview(ctx, omadacId, siteId).Execute()

Get site overview diagram info



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
	resp, r, err := apiClient.DashboardAPI.GetOverview(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverview`: OperationResponseGetDashboardOverview
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseGetDashboardOverview**](OperationResponseGetDashboardOverview.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoeUsage

> OperationResponseListPoeUsage GetPoeUsage(ctx, omadacId, siteId).Execute()

Get poe usage



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
	resp, r, err := apiClient.DashboardAPI.GetPoeUsage(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetPoeUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoeUsage`: OperationResponseListPoeUsage
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetPoeUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoeUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListPoeUsage**](OperationResponseListPoeUsage.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetryAndDroppedRate

> OperationResponseRetryDropRate GetRetryAndDroppedRate(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get retried rate and dropped rate



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetRetryAndDroppedRate(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetRetryAndDroppedRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRetryAndDroppedRate`: OperationResponseRetryDropRate
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetRetryAndDroppedRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetryAndDroppedRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseRetryDropRate**](OperationResponseRetryDropRate.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpeedTestV2Result

> OperationResponseSpeedTestV2ResultVO GetSpeedTestV2Result(ctx, omadacId, siteId, gatewayMac).Execute()

Get Speed Test Result



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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetSpeedTestV2Result(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetSpeedTestV2Result``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpeedTestV2Result`: OperationResponseSpeedTestV2ResultVO
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetSpeedTestV2Result`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpeedTestV2ResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSpeedTestV2ResultVO**](OperationResponseSpeedTestV2ResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpeedTestV2ResultDateList

> OperationResponseGridVOSpeedTestV2ResultItemOpenApiVO GetSpeedTestV2ResultDateList(ctx, omadacId, siteId, gatewayMac).OpenApiQuerySpeedTestDateListVO(openApiQuerySpeedTestDateListVO).Execute()

Get date list of speed test results



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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	openApiQuerySpeedTestDateListVO := *openapiclient.NewOpenApiQuerySpeedTestDateListVO(int32(123), int32(123)) // OpenApiQuerySpeedTestDateListVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetSpeedTestV2ResultDateList(context.Background(), omadacId, siteId, gatewayMac).OpenApiQuerySpeedTestDateListVO(openApiQuerySpeedTestDateListVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetSpeedTestV2ResultDateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpeedTestV2ResultDateList`: OperationResponseGridVOSpeedTestV2ResultItemOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetSpeedTestV2ResultDateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpeedTestV2ResultDateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **openApiQuerySpeedTestDateListVO** | [**OpenApiQuerySpeedTestDateListVO**](OpenApiQuerySpeedTestDateListVO.md) |  | 

### Return type

[**OperationResponseGridVOSpeedTestV2ResultItemOpenApiVO**](OperationResponseGridVOSpeedTestV2ResultItemOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchSummary

> OperationResponseSwitchSummary GetSwitchSummary(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get switch summary



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetSwitchSummary(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetSwitchSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchSummary`: OperationResponseSwitchSummary
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetSwitchSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseSwitchSummary**](OperationResponseSwitchSummary.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTop5Aps

> OperationResponseTopApByRtDropVO GetTop5Aps(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get Top 5 Aps



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetTop5Aps(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetTop5Aps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTop5Aps`: OperationResponseTopApByRtDropVO
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetTop5Aps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTop5ApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseTopApByRtDropVO**](OperationResponseTopApByRtDropVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopCpuUsageWithTimeRange

> OperationResponseListCpuUsage GetTopCpuUsageWithTimeRange(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get top device cpu usage



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetTopCpuUsageWithTimeRange(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetTopCpuUsageWithTimeRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopCpuUsageWithTimeRange`: OperationResponseListCpuUsage
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetTopCpuUsageWithTimeRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopCpuUsageWithTimeRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListCpuUsage**](OperationResponseListCpuUsage.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopMemoryUsageWithTimeRange

> OperationResponseListMemUsage GetTopMemoryUsageWithTimeRange(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get top device memory usage



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetTopMemoryUsageWithTimeRange(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetTopMemoryUsageWithTimeRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopMemoryUsageWithTimeRange`: OperationResponseListMemUsage
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetTopMemoryUsageWithTimeRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopMemoryUsageWithTimeRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListMemUsage**](OperationResponseListMemUsage.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrafficActivities

> OperationResponseTrafficActivities GetTrafficActivities(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get traffic activity



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetTrafficActivities(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetTrafficActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrafficActivities`: OperationResponseTrafficActivities
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetTrafficActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrafficActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseTrafficActivities**](OperationResponseTrafficActivities.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrafficDistribution

> OperationResponseTrafficDistribution GetTrafficDistribution(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get traffic distribution



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetTrafficDistribution(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetTrafficDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrafficDistribution`: OperationResponseTrafficDistribution
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetTrafficDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrafficDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseTrafficDistribution**](OperationResponseTrafficDistribution.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWifiSummary

> OperationResponseWifiSummary GetWifiSummary(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get wifi summary



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetWifiSummary(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetWifiSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWifiSummary`: OperationResponseWifiSummary
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetWifiSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWifiSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseWifiSummary**](OperationResponseWifiSummary.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllTabs

> OperationResponseListListTabs ListAllTabs(ctx, omadacId, siteId).Execute()

Get tab list



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
	resp, r, err := apiClient.DashboardAPI.ListAllTabs(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.ListAllTabs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllTabs`: OperationResponseListListTabs
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.ListAllTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListListTabs**](OperationResponseListListTabs.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTabs

> OperationResponseListListTabs ListTabs(ctx, omadacId, siteId).Execute()

Get tab without overall tab list



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
	resp, r, err := apiClient.DashboardAPI.ListTabs(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.ListTabs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTabs`: OperationResponseListListTabs
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.ListTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListListTabs**](OperationResponseListListTabs.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTab

> OperationResponseWithoutResult RemoveTab(ctx, omadacId, siteId, tabId).Execute()

Delete an existing tab



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
	tabId := "tabId_example" // string | Tab ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.RemoveTab(context.Background(), omadacId, siteId, tabId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.RemoveTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveTab`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.RemoveTab`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tabId** | **string** | Tab ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTabRequest struct via the builder pattern


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


## StartSpeedTestV2

> OperationResponseStartSpeedTestV2ResultVO StartSpeedTestV2(ctx, omadacId, siteId, gatewayMac).OpenApiSpeedTestSelectPortsVO(openApiSpeedTestSelectPortsVO).Execute()

Start Speed Test



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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	openApiSpeedTestSelectPortsVO := *openapiclient.NewOpenApiSpeedTestSelectPortsVO() // OpenApiSpeedTestSelectPortsVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.StartSpeedTestV2(context.Background(), omadacId, siteId, gatewayMac).OpenApiSpeedTestSelectPortsVO(openApiSpeedTestSelectPortsVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.StartSpeedTestV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartSpeedTestV2`: OperationResponseStartSpeedTestV2ResultVO
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.StartSpeedTestV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSpeedTestV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **openApiSpeedTestSelectPortsVO** | [**OpenApiSpeedTestSelectPortsVO**](OpenApiSpeedTestSelectPortsVO.md) |  | 

### Return type

[**OperationResponseStartSpeedTestV2ResultVO**](OperationResponseStartSpeedTestV2ResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTab

> OperationResponseWithoutResult UpdateTab(ctx, omadacId, siteId, tabId).UpdateTabOpenApiVO(updateTabOpenApiVO).Execute()

Modify an existing tab



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
	tabId := "tabId_example" // string | Tab ID
	updateTabOpenApiVO := *openapiclient.NewUpdateTabOpenApiVO() // UpdateTabOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.UpdateTab(context.Background(), omadacId, siteId, tabId).UpdateTabOpenApiVO(updateTabOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.UpdateTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTab`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.UpdateTab`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tabId** | **string** | Tab ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateTabOpenApiVO** | [**UpdateTabOpenApiVO**](UpdateTabOpenApiVO.md) |  | 

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


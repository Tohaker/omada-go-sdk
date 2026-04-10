# InsightAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteALLBatchFullChannelDetectHistory**](InsightAPI.md#deleteallbatchfullchanneldetecthistory) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect | Delete all interference detection history
[**DeleteBatchFullChannelDetectHistory**](InsightAPI.md#deletebatchfullchanneldetecthistory) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect/{historyId} | Delete the interference detection history
[**ExportBatchFullChannelDetectResultData**](InsightAPI.md#exportbatchfullchanneldetectresultdata) | **Get** /openapi/v1/{omadacId}/files/sites/{siteId}/full-channel-detect/{historyId} | Export batch interference detection results
[**ExportRogueApsGlobal**](InsightAPI.md#exportrogueapsglobal) | **Post** /openapi/v1/{omadacId}/files/neighbors | Export global Rogue AP scan results
[**GetBatchChannelLoadsResult**](InsightAPI.md#getbatchchannelloadsresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect/{historyId}/channel-load | Get channel utilization results of batch interference detection
[**GetBatchFullChannelDetectApList**](InsightAPI.md#getbatchfullchanneldetectaplist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect/{historyId}/ap-list | Get the AP list of batch interference detection
[**GetBatchFullChannelDetectStatus**](InsightAPI.md#getbatchfullchanneldetectstatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect/status | Get the interference detection status
[**GetBatchWifiInterferencesResult**](InsightAPI.md#getbatchwifiinterferencesresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect/{historyId}/grid/wifi-interference | Get WiFi interference results of batch interference detection
[**GetChannelLoadResult**](InsightAPI.md#getchannelloadresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/full-channel-detect/channel-load | Get channel utilization results of interference detection
[**GetDisableFullChannelDetectApMacList**](InsightAPI.md#getdisablefullchanneldetectapmaclist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect/batch/info | Get the MAC list of APs that cannot perform interference detection
[**GetFullChannelDetectStatus**](InsightAPI.md#getfullchanneldetectstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/full-channel-detect/status | Get the interference detection status of APs
[**GetGridEnableFullChannelDetectApList**](InsightAPI.md#getgridenablefullchanneldetectaplist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/grid/full-channel-detect/ap-list | Get the list of devices that can perform interference detection by page
[**GetGridFullChannelDetectHistory**](InsightAPI.md#getgridfullchanneldetecthistory) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect/grid/scan-history | Get the interference detection history by page
[**GetGridOswRoutingTable**](InsightAPI.md#getgridoswroutingtable) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/routingTable | Get switch routing table
[**GetGridRogueAps**](InsightAPI.md#getgridrogueaps) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/insight/rogueaps | Query the Rogue AP scan results
[**GetGridRouting**](InsightAPI.md#getgridrouting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/insight/routing/{type} | Get the routing table data interface
[**GetGridStackRoutingTable**](InsightAPI.md#getgridstackroutingtable) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stack/{stackId}/routingTable | Get stack routing table
[**GetGridVpnClientStatus**](InsightAPI.md#getgridvpnclientstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/vpn/stats/client | Get VPN Client status list
[**GetGridVpnIpSec**](InsightAPI.md#getgridvpnipsec) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/vpn/stats/ipsec | Query the vpnStats ipsec list
[**GetGridVpnS2SPeersStatus**](InsightAPI.md#getgridvpns2speersstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/vpn/stats/s2s/{vpnId}/peer | Get VPN Site-to-Site&#39;s peers status list
[**GetGridVpnS2SStatus**](InsightAPI.md#getgridvpns2sstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/vpn/stats/s2s | Get VPN Site-to-Site status list
[**GetGridVpnServerClientsStatus**](InsightAPI.md#getgridvpnserverclientsstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/vpn/stats/server/{vpnId}/client | Get VPN Server&#39;s clients status list
[**GetGridVpnServerStatus**](InsightAPI.md#getgridvpnserverstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/vpn/stats/server | Get VPN Server status list
[**GetGridVpnTunnel**](InsightAPI.md#getgridvpntunnel) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/vpn/stats/tunnel | Query the vpnStats tunnel list
[**GetGridWidsData**](InsightAPI.md#getgridwidsdata) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/insight/wids | Query the Wireless IDS entry list
[**GetGridWifiInterfResult**](InsightAPI.md#getgridwifiinterfresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/full-channel-detect/wlan-interference |  Get WiFi interference results of interference detection
[**GetGridWipsBlackList**](InsightAPI.md#getgridwipsblacklist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/insight/wids/blacklist | Get the dynamic blacklist entry data of Wireless IDS
[**GetPortForwardStatus**](InsightAPI.md#getportforwardstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/insight/port-forwarding/{type} | Get Port Forwarding Status
[**GetSpectralScanHistoryResult**](InsightAPI.md#getspectralscanhistoryresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/spectral-scan/aps/{apMac}/spectral-scan-result/history | Get history results of environment scanning
[**GetSpectralScanResult**](InsightAPI.md#getspectralscanresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/spectral-scan/aps/{apMac}/spectral-scan-result | Get current results of environment scanning
[**RemoveWipsBlackList**](InsightAPI.md#removewipsblacklist) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/insight/wids/blacklist/{mac} | Remove the specified Device MAC from the blacklist of the reported device
[**ScanRogueAps**](InsightAPI.md#scanrogueaps) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/rogueaps/scan | Scan Rogue APs
[**StartBatchFullChannelDetect**](InsightAPI.md#startbatchfullchanneldetect) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/full-channel-detect-start | Enable batch interference detection
[**StopSpectralScan**](InsightAPI.md#stopspectralscan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/spectral-scan-stop | Stop environment scanning
[**TerminateVpnTunnel**](InsightAPI.md#terminatevpntunnel) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/vpn/stats/tunnel/terminate | Terminating vpn tunnel
[**TriggerFullChannelDetect**](InsightAPI.md#triggerfullchanneldetect) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/full-channel-detect-start | Enable interference detection
[**TriggerSpectralScan**](InsightAPI.md#triggerspectralscan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/spectral-scan-start | Enable environment scanning



## DeleteALLBatchFullChannelDetectHistory

> OperationResponseWithoutResult DeleteALLBatchFullChannelDetectHistory(ctx, omadacId, siteId).Execute()

Delete all interference detection history



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
	resp, r, err := apiClient.InsightAPI.DeleteALLBatchFullChannelDetectHistory(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.DeleteALLBatchFullChannelDetectHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteALLBatchFullChannelDetectHistory`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.DeleteALLBatchFullChannelDetectHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteALLBatchFullChannelDetectHistoryRequest struct via the builder pattern


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


## DeleteBatchFullChannelDetectHistory

> OperationResponseWithoutResult DeleteBatchFullChannelDetectHistory(ctx, omadacId, siteId, historyId).Execute()

Delete the interference detection history



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
	historyId := "historyId_example" // string | The ID List of full channel detect

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.DeleteBatchFullChannelDetectHistory(context.Background(), omadacId, siteId, historyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.DeleteBatchFullChannelDetectHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBatchFullChannelDetectHistory`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.DeleteBatchFullChannelDetectHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**historyId** | **string** | The ID List of full channel detect | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchFullChannelDetectHistoryRequest struct via the builder pattern


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


## ExportBatchFullChannelDetectResultData

> OperationResponseWithoutResult ExportBatchFullChannelDetectResultData(ctx, omadacId, siteId, historyId).Format(format).Execute()

Export batch interference detection results



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
	historyId := "historyId_example" // string | Customer ID
	format := int32(56) // int32 | export data format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.ExportBatchFullChannelDetectResultData(context.Background(), omadacId, siteId, historyId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.ExportBatchFullChannelDetectResultData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportBatchFullChannelDetectResultData`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.ExportBatchFullChannelDetectResultData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**historyId** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportBatchFullChannelDetectResultDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **int32** | export data format | 

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


## ExportRogueApsGlobal

> OperationResponseWithoutResult ExportRogueApsGlobal(ctx, omadacId).RequiredParametersForExportingRogueAPScanResults(requiredParametersForExportingRogueAPScanResults).Execute()

Export global Rogue AP scan results



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
	requiredParametersForExportingRogueAPScanResults := *openapiclient.NewRequiredParametersForExportingRogueAPScanResults() // RequiredParametersForExportingRogueAPScanResults | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.ExportRogueApsGlobal(context.Background(), omadacId).RequiredParametersForExportingRogueAPScanResults(requiredParametersForExportingRogueAPScanResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.ExportRogueApsGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportRogueApsGlobal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.ExportRogueApsGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRogueApsGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requiredParametersForExportingRogueAPScanResults** | [**RequiredParametersForExportingRogueAPScanResults**](RequiredParametersForExportingRogueAPScanResults.md) |  | 

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


## GetBatchChannelLoadsResult

> OperationResponseListApChannelLoadResult GetBatchChannelLoadsResult(ctx, omadacId, siteId, historyId).Mac(mac).Execute()

Get channel utilization results of batch interference detection



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
	historyId := "historyId_example" // string | Customer ID
	mac := "mac_example" // string | Mac Address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetBatchChannelLoadsResult(context.Background(), omadacId, siteId, historyId).Mac(mac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetBatchChannelLoadsResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchChannelLoadsResult`: OperationResponseListApChannelLoadResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetBatchChannelLoadsResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**historyId** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchChannelLoadsResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mac** | **string** | Mac Address, like AA-BB-CC-DD-EE-FF | 

### Return type

[**OperationResponseListApChannelLoadResult**](OperationResponseListApChannelLoadResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchFullChannelDetectApList

> OperationResponseBatchFullChannelDetectApListOpenApiVO GetBatchFullChannelDetectApList(ctx, omadacId, siteId, historyId).Execute()

Get the AP list of batch interference detection



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
	historyId := "historyId_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetBatchFullChannelDetectApList(context.Background(), omadacId, siteId, historyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetBatchFullChannelDetectApList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchFullChannelDetectApList`: OperationResponseBatchFullChannelDetectApListOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetBatchFullChannelDetectApList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**historyId** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchFullChannelDetectApListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseBatchFullChannelDetectApListOpenApiVO**](OperationResponseBatchFullChannelDetectApListOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchFullChannelDetectStatus

> OperationResponseListBatchFullChannelDetectStatusOpenApiVO GetBatchFullChannelDetectStatus(ctx, omadacId, siteId).QueryBatchFullChannelDetectStatusVO(queryBatchFullChannelDetectStatusVO).Execute()

Get the interference detection status



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
	queryBatchFullChannelDetectStatusVO := *openapiclient.NewQueryBatchFullChannelDetectStatusVO() // QueryBatchFullChannelDetectStatusVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetBatchFullChannelDetectStatus(context.Background(), omadacId, siteId).QueryBatchFullChannelDetectStatusVO(queryBatchFullChannelDetectStatusVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetBatchFullChannelDetectStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchFullChannelDetectStatus`: OperationResponseListBatchFullChannelDetectStatusOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetBatchFullChannelDetectStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchFullChannelDetectStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryBatchFullChannelDetectStatusVO** | [**QueryBatchFullChannelDetectStatusVO**](QueryBatchFullChannelDetectStatusVO.md) |  | 

### Return type

[**OperationResponseListBatchFullChannelDetectStatusOpenApiVO**](OperationResponseListBatchFullChannelDetectStatusOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchWifiInterferencesResult

> OperationResponseGridVOBatchWifiInterferenceOpenApiVO GetBatchWifiInterferencesResult(ctx, omadacId, siteId, historyId).Page(page).PageSize(pageSize).FiltersType(filtersType).FiltersRadioId(filtersRadioId).FiltersMac(filtersMac).Execute()

Get WiFi interference results of batch interference detection



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
	historyId := "historyId_example" // string | Customer ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersType := "filtersType_example" // string | Filter query parameters, support field wifi-interference result type. 0: all device; 1: one device.
	filtersRadioId := "filtersRadioId_example" // string | Filter query parameters, support field radio id. 0: 2.4G; 1: 5G; 3:6G. (optional)
	filtersMac := "filtersMac_example" // string | Filter query parameters, support field AP MAC address, like AA-BB-CC-DD-EE-FF. When parameter [filter.type] is 0, this parameter is invalid. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetBatchWifiInterferencesResult(context.Background(), omadacId, siteId, historyId).Page(page).PageSize(pageSize).FiltersType(filtersType).FiltersRadioId(filtersRadioId).FiltersMac(filtersMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetBatchWifiInterferencesResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchWifiInterferencesResult`: OperationResponseGridVOBatchWifiInterferenceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetBatchWifiInterferencesResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**historyId** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchWifiInterferencesResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersType** | **string** | Filter query parameters, support field wifi-interference result type. 0: all device; 1: one device. | 
 **filtersRadioId** | **string** | Filter query parameters, support field radio id. 0: 2.4G; 1: 5G; 3:6G. | 
 **filtersMac** | **string** | Filter query parameters, support field AP MAC address, like AA-BB-CC-DD-EE-FF. When parameter [filter.type] is 0, this parameter is invalid. | 

### Return type

[**OperationResponseGridVOBatchWifiInterferenceOpenApiVO**](OperationResponseGridVOBatchWifiInterferenceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelLoadResult

> OperationResponseApChannelLoadResult GetChannelLoadResult(ctx, omadacId, siteId, apMac).Execute()

Get channel utilization results of interference detection



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetChannelLoadResult(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetChannelLoadResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelLoadResult`: OperationResponseApChannelLoadResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetChannelLoadResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelLoadResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApChannelLoadResult**](OperationResponseApChannelLoadResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDisableFullChannelDetectApMacList

> OperationResponseDisableFullChanDetectApListOpenApiVO GetDisableFullChannelDetectApMacList(ctx, omadacId, siteId).Execute()

Get the MAC list of APs that cannot perform interference detection



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
	resp, r, err := apiClient.InsightAPI.GetDisableFullChannelDetectApMacList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetDisableFullChannelDetectApMacList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisableFullChannelDetectApMacList`: OperationResponseDisableFullChanDetectApListOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetDisableFullChannelDetectApMacList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisableFullChannelDetectApMacListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDisableFullChanDetectApListOpenApiVO**](OperationResponseDisableFullChanDetectApListOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullChannelDetectStatus

> OperationResponseApFullChannelScanStatus GetFullChannelDetectStatus(ctx, omadacId, siteId, apMac).Execute()

Get the interference detection status of APs



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetFullChannelDetectStatus(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetFullChannelDetectStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFullChannelDetectStatus`: OperationResponseApFullChannelScanStatus
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetFullChannelDetectStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFullChannelDetectStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApFullChannelScanStatus**](OperationResponseApFullChannelScanStatus.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridEnableFullChannelDetectApList

> OperationResponseListEnableFullChannelDetectApInfoOpenApiVO GetGridEnableFullChannelDetectApList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the list of devices that can perform interference detection by page



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
	resp, r, err := apiClient.InsightAPI.GetGridEnableFullChannelDetectApList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridEnableFullChannelDetectApList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridEnableFullChannelDetectApList`: OperationResponseListEnableFullChannelDetectApInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridEnableFullChannelDetectApList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridEnableFullChannelDetectApListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseListEnableFullChannelDetectApInfoOpenApiVO**](OperationResponseListEnableFullChannelDetectApInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridFullChannelDetectHistory

> OperationResponseBatchFullChannelDetectGridBatchFullChannelDetectHistoryOpenApiVO GetGridFullChannelDetectHistory(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the interference detection history by page



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
	resp, r, err := apiClient.InsightAPI.GetGridFullChannelDetectHistory(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridFullChannelDetectHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridFullChannelDetectHistory`: OperationResponseBatchFullChannelDetectGridBatchFullChannelDetectHistoryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridFullChannelDetectHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridFullChannelDetectHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseBatchFullChannelDetectGridBatchFullChannelDetectHistoryOpenApiVO**](OperationResponseBatchFullChannelDetectGridBatchFullChannelDetectHistoryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridOswRoutingTable

> OperationResponseGridVOOswRoutingOpenApiVO GetGridOswRoutingTable(ctx, omadacId, siteId, switchMac).Page(page).PageSize(pageSize).SortsDestinationIp(sortsDestinationIp).FiltersType(filtersType).SearchKey(searchKey).Execute()

Get switch routing table



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
	switchMac := "switchMac_example" // string | Switch MAC address, like AA-BB-CC-DD-EE-FF
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsDestinationIp := "sortsDestinationIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersType := "filtersType_example" // string | Filter query parameters, support field type (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field destinationIp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridOswRoutingTable(context.Background(), omadacId, siteId, switchMac).Page(page).PageSize(pageSize).SortsDestinationIp(sortsDestinationIp).FiltersType(filtersType).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridOswRoutingTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridOswRoutingTable`: OperationResponseGridVOOswRoutingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridOswRoutingTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridOswRoutingTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsDestinationIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersType** | **string** | Filter query parameters, support field type | 
 **searchKey** | **string** | Fuzzy query parameters, support field destinationIp | 

### Return type

[**OperationResponseGridVOOswRoutingOpenApiVO**](OperationResponseGridVOOswRoutingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridRogueAps

> OperationResponseGridVORogueAPScanResultEntry GetGridRogueAps(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Query the Rogue AP scan results



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
	resp, r, err := apiClient.InsightAPI.GetGridRogueAps(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridRogueAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridRogueAps`: OperationResponseGridVORogueAPScanResultEntry
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridRogueAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridRogueApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVORogueAPScanResultEntry**](OperationResponseGridVORogueAPScanResultEntry.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridRouting

> OperationResponseGridVOOswRoutingOpenApiVO GetGridRouting(ctx, omadacId, siteId, type_).Page(page).PageSize(pageSize).Execute()

Get the routing table data interface



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
	type_ := "type__example" // string | type
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridRouting(context.Background(), omadacId, siteId, type_).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridRouting`: OperationResponseGridVOOswRoutingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**type_** | **string** | type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOswRoutingOpenApiVO**](OperationResponseGridVOOswRoutingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStackRoutingTable

> OperationResponseGridVOStackRoutingOpenApiVO GetGridStackRoutingTable(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).SortsDestinationIp(sortsDestinationIp).FiltersType(filtersType).SearchKey(searchKey).Execute()

Get stack routing table



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
	stackId := "stackId_example" // string | Stack ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsDestinationIp := "sortsDestinationIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersType := "filtersType_example" // string | Filter query parameters, support field type (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field destinationIp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridStackRoutingTable(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).SortsDestinationIp(sortsDestinationIp).FiltersType(filtersType).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridStackRoutingTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStackRoutingTable`: OperationResponseGridVOStackRoutingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridStackRoutingTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridStackRoutingTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsDestinationIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersType** | **string** | Filter query parameters, support field type | 
 **searchKey** | **string** | Fuzzy query parameters, support field destinationIp | 

### Return type

[**OperationResponseGridVOStackRoutingOpenApiVO**](OperationResponseGridVOStackRoutingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnClientStatus

> OperationResponseGridVOVpnTunnelStatusVO GetGridVpnClientStatus(ctx, omadacId, siteId).FiltersVpnType(filtersVpnType).Page(page).PageSize(pageSize).Execute()

Get VPN Client status list



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
	filtersVpnType := "filtersVpnType_example" // string | Filter query parameters, support field vpnType. 0: L2TP; 1: PPTP; 3: OpenVPN; 4: WireGuard.
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridVpnClientStatus(context.Background(), omadacId, siteId).FiltersVpnType(filtersVpnType).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridVpnClientStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnClientStatus`: OperationResponseGridVOVpnTunnelStatusVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridVpnClientStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnClientStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersVpnType** | **string** | Filter query parameters, support field vpnType. 0: L2TP; 1: PPTP; 3: OpenVPN; 4: WireGuard. | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOVpnTunnelStatusVO**](OperationResponseGridVOVpnTunnelStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnIpSec

> OperationResponseGridVOOsgVpnIpSecOpenApiVO GetGridVpnIpSec(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Query the vpnStats ipsec list



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
	resp, r, err := apiClient.InsightAPI.GetGridVpnIpSec(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridVpnIpSec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnIpSec`: OperationResponseGridVOOsgVpnIpSecOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridVpnIpSec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnIpSecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOsgVpnIpSecOpenApiVO**](OperationResponseGridVOOsgVpnIpSecOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnS2SPeersStatus

> OperationResponseGetVpnResponseGridVOVpnTunnelRemoteStatusVO GetGridVpnS2SPeersStatus(ctx, omadacId, siteId, vpnId).Page(page).PageSize(pageSize).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()

Get VPN Site-to-Site's peers status list



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
	vpnId := "vpnId_example" // string | VPN ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersStatus := "filtersStatus_example" // string | Filter query parameters, support field status. 0: disconnected; 1: connected. (optional)
	searchKey := "searchKey_example" // string | searchKey (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridVpnS2SPeersStatus(context.Background(), omadacId, siteId, vpnId).Page(page).PageSize(pageSize).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridVpnS2SPeersStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnS2SPeersStatus`: OperationResponseGetVpnResponseGridVOVpnTunnelRemoteStatusVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridVpnS2SPeersStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnS2SPeersStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersStatus** | **string** | Filter query parameters, support field status. 0: disconnected; 1: connected. | 
 **searchKey** | **string** | searchKey | 

### Return type

[**OperationResponseGetVpnResponseGridVOVpnTunnelRemoteStatusVO**](OperationResponseGetVpnResponseGridVOVpnTunnelRemoteStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnS2SStatus

> OperationResponseGetVpnResponseGridVOVpnTunnelStatusVO GetGridVpnS2SStatus(ctx, omadacId, siteId).FiltersVpnType(filtersVpnType).Page(page).PageSize(pageSize).Execute()

Get VPN Site-to-Site status list



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
	filtersVpnType := "filtersVpnType_example" // string | Filter query parameters, support field vpnType. 2: IPSec; 4: WireGuard.
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridVpnS2SStatus(context.Background(), omadacId, siteId).FiltersVpnType(filtersVpnType).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridVpnS2SStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnS2SStatus`: OperationResponseGetVpnResponseGridVOVpnTunnelStatusVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridVpnS2SStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnS2SStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersVpnType** | **string** | Filter query parameters, support field vpnType. 2: IPSec; 4: WireGuard. | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGetVpnResponseGridVOVpnTunnelStatusVO**](OperationResponseGetVpnResponseGridVOVpnTunnelStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnServerClientsStatus

> OperationResponseGridVOVpnTunnelRemoteStatusVO GetGridVpnServerClientsStatus(ctx, omadacId, siteId, vpnId).Page(page).PageSize(pageSize).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()

Get VPN Server's clients status list



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
	vpnId := "vpnId_example" // string | VPN ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersStatus := "filtersStatus_example" // string | Filter query parameters, support field status. 0: disconnected; 1: connected. (optional)
	searchKey := "searchKey_example" // string | searchKey (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridVpnServerClientsStatus(context.Background(), omadacId, siteId, vpnId).Page(page).PageSize(pageSize).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridVpnServerClientsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnServerClientsStatus`: OperationResponseGridVOVpnTunnelRemoteStatusVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridVpnServerClientsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnServerClientsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersStatus** | **string** | Filter query parameters, support field status. 0: disconnected; 1: connected. | 
 **searchKey** | **string** | searchKey | 

### Return type

[**OperationResponseGridVOVpnTunnelRemoteStatusVO**](OperationResponseGridVOVpnTunnelRemoteStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnServerStatus

> OperationResponseVpnTunnelGridVOVpnTunnelStatusVO GetGridVpnServerStatus(ctx, omadacId, siteId).FiltersVpnType(filtersVpnType).Page(page).PageSize(pageSize).FiltersClientMode(filtersClientMode).SortsActiveClients(sortsActiveClients).SortsClients(sortsClients).Execute()

Get VPN Server status list



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
	filtersVpnType := "filtersVpnType_example" // string | Filter query parameters, support field vpnType. 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN.
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersClientMode := "filtersClientMode_example" // string | Filter query parameters, support field Client mode for VPN user. 0: NEM(Network Extension Mode) ; 1: Client. (optional)
	sortsActiveClients := "sortsActiveClients_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsClients := "sortsClients_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridVpnServerStatus(context.Background(), omadacId, siteId).FiltersVpnType(filtersVpnType).Page(page).PageSize(pageSize).FiltersClientMode(filtersClientMode).SortsActiveClients(sortsActiveClients).SortsClients(sortsClients).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridVpnServerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnServerStatus`: OperationResponseVpnTunnelGridVOVpnTunnelStatusVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridVpnServerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnServerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersVpnType** | **string** | Filter query parameters, support field vpnType. 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersClientMode** | **string** | Filter query parameters, support field Client mode for VPN user. 0: NEM(Network Extension Mode) ; 1: Client. | 
 **sortsActiveClients** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsClients** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseVpnTunnelGridVOVpnTunnelStatusVO**](OperationResponseVpnTunnelGridVOVpnTunnelStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnTunnel

> OperationResponseGridVOOsgVpnTunnelOpenApiVO GetGridVpnTunnel(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Query the vpnStats tunnel list



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
	resp, r, err := apiClient.InsightAPI.GetGridVpnTunnel(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnTunnel`: OperationResponseGridVOOsgVpnTunnelOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridVpnTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOsgVpnTunnelOpenApiVO**](OperationResponseGridVOOsgVpnTunnelOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridWidsData

> OperationResponseGridVOWidsDataOpenApiVO GetGridWidsData(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Query the Wireless IDS entry list



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
	resp, r, err := apiClient.InsightAPI.GetGridWidsData(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridWidsData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridWidsData`: OperationResponseGridVOWidsDataOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridWidsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridWidsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOWidsDataOpenApiVO**](OperationResponseGridVOWidsDataOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridWifiInterfResult

> OperationResponseGridVOApWifiInterferenceResult GetGridWifiInterfResult(ctx, omadacId, siteId, apMac).Page(page).PageSize(pageSize).Execute()

 Get WiFi interference results of interference detection



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetGridWifiInterfResult(context.Background(), omadacId, siteId, apMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridWifiInterfResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridWifiInterfResult`: OperationResponseGridVOApWifiInterferenceResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridWifiInterfResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridWifiInterfResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOApWifiInterferenceResult**](OperationResponseGridVOApWifiInterferenceResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridWipsBlackList

> OperationResponseGridVOWipsBlackListOpenApiVO GetGridWipsBlackList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the dynamic blacklist entry data of Wireless IDS



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
	resp, r, err := apiClient.InsightAPI.GetGridWipsBlackList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetGridWipsBlackList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridWipsBlackList`: OperationResponseGridVOWipsBlackListOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetGridWipsBlackList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridWipsBlackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOWipsBlackListOpenApiVO**](OperationResponseGridVOWipsBlackListOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortForwardStatus

> OperationResponseGridVOPortForwardingOpenApiVO GetPortForwardStatus(ctx, omadacId, siteId, type_).Page(page).PageSize(pageSize).Execute()

Get Port Forwarding Status



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
	type_ := "type__example" // string | User/UPnP
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetPortForwardStatus(context.Background(), omadacId, siteId, type_).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetPortForwardStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortForwardStatus`: OperationResponseGridVOPortForwardingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetPortForwardStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**type_** | **string** | User/UPnP | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortForwardStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOPortForwardingOpenApiVO**](OperationResponseGridVOPortForwardingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpectralScanHistoryResult

> OperationResponse GetSpectralScanHistoryResult(ctx, omadacId, siteId, apMac).Execute()

Get history results of environment scanning



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetSpectralScanHistoryResult(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetSpectralScanHistoryResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpectralScanHistoryResult`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetSpectralScanHistoryResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpectralScanHistoryResultRequest struct via the builder pattern


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


## GetSpectralScanResult

> OperationResponse GetSpectralScanResult(ctx, omadacId, siteId, apMac).Times(times).Execute()

Get current results of environment scanning



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF
	times := int32(56) // int32 | The time between the requested spectral scan results.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.GetSpectralScanResult(context.Background(), omadacId, siteId, apMac).Times(times).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.GetSpectralScanResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpectralScanResult`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.GetSpectralScanResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpectralScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **times** | **int32** | The time between the requested spectral scan results. | 

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


## RemoveWipsBlackList

> OperationResponseWithoutResult RemoveWipsBlackList(ctx, omadacId, siteId, mac).RemoveBlackListOpenApiVO(removeBlackListOpenApiVO).Execute()

Remove the specified Device MAC from the blacklist of the reported device



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
	mac := "mac_example" // string | Mac Address
	removeBlackListOpenApiVO := *openapiclient.NewRemoveBlackListOpenApiVO() // RemoveBlackListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.RemoveWipsBlackList(context.Background(), omadacId, siteId, mac).RemoveBlackListOpenApiVO(removeBlackListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.RemoveWipsBlackList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveWipsBlackList`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.RemoveWipsBlackList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | Mac Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWipsBlackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **removeBlackListOpenApiVO** | [**RemoveBlackListOpenApiVO**](RemoveBlackListOpenApiVO.md) |  | 

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


## ScanRogueAps

> OperationResponseWithoutResult ScanRogueAps(ctx, omadacId, siteId).Execute()

Scan Rogue APs



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
	resp, r, err := apiClient.InsightAPI.ScanRogueAps(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.ScanRogueAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanRogueAps`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.ScanRogueAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanRogueApsRequest struct via the builder pattern


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


## StartBatchFullChannelDetect

> OperationResponseWithoutResult StartBatchFullChannelDetect(ctx, omadacId, siteId).StartBatchFullChannelDetectCmdOpenApiVO(startBatchFullChannelDetectCmdOpenApiVO).Execute()

Enable batch interference detection



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
	startBatchFullChannelDetectCmdOpenApiVO := *openapiclient.NewStartBatchFullChannelDetectCmdOpenApiVO("SelectType_example") // StartBatchFullChannelDetectCmdOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.StartBatchFullChannelDetect(context.Background(), omadacId, siteId).StartBatchFullChannelDetectCmdOpenApiVO(startBatchFullChannelDetectCmdOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.StartBatchFullChannelDetect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartBatchFullChannelDetect`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.StartBatchFullChannelDetect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartBatchFullChannelDetectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startBatchFullChannelDetectCmdOpenApiVO** | [**StartBatchFullChannelDetectCmdOpenApiVO**](StartBatchFullChannelDetectCmdOpenApiVO.md) |  | 

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


## StopSpectralScan

> OperationResponseWithoutResult StopSpectralScan(ctx, omadacId, siteId, apMac).Execute()

Stop environment scanning



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.StopSpectralScan(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.StopSpectralScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopSpectralScan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.StopSpectralScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSpectralScanRequest struct via the builder pattern


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


## TerminateVpnTunnel

> OperationResponseWithoutResult TerminateVpnTunnel(ctx, omadacId, siteId).OsgVpnTunnelOpenApiVO(osgVpnTunnelOpenApiVO).Execute()

Terminating vpn tunnel



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
	osgVpnTunnelOpenApiVO := *openapiclient.NewOsgVpnTunnelOpenApiVO() // OsgVpnTunnelOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.TerminateVpnTunnel(context.Background(), omadacId, siteId).OsgVpnTunnelOpenApiVO(osgVpnTunnelOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.TerminateVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TerminateVpnTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.TerminateVpnTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateVpnTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **osgVpnTunnelOpenApiVO** | [**OsgVpnTunnelOpenApiVO**](OsgVpnTunnelOpenApiVO.md) |  | 

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


## TriggerFullChannelDetect

> OperationResponseWithoutResult TriggerFullChannelDetect(ctx, omadacId, siteId, apMac).FullChanScanCommandOpenApiVO(fullChanScanCommandOpenApiVO).Execute()

Enable interference detection



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF
	fullChanScanCommandOpenApiVO := *openapiclient.NewFullChanScanCommandOpenApiVO() // FullChanScanCommandOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.TriggerFullChannelDetect(context.Background(), omadacId, siteId, apMac).FullChanScanCommandOpenApiVO(fullChanScanCommandOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.TriggerFullChannelDetect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerFullChannelDetect`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.TriggerFullChannelDetect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerFullChannelDetectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fullChanScanCommandOpenApiVO** | [**FullChanScanCommandOpenApiVO**](FullChanScanCommandOpenApiVO.md) |  | 

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


## TriggerSpectralScan

> OperationResponseWithoutResult TriggerSpectralScan(ctx, omadacId, siteId, apMac).Execute()

Enable environment scanning



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightAPI.TriggerSpectralScan(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightAPI.TriggerSpectralScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerSpectralScan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `InsightAPI.TriggerSpectralScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerSpectralScanRequest struct via the builder pattern


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


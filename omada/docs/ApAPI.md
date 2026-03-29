# \ApAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchModifyAntSwitchConfig**](ApAPI.md#BatchModifyAntSwitchConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/cmd/aps/config/ant-switch | Batch Modify AP antSwitch Setting
[**BatchModifyApVlanConfig**](ApAPI.md#BatchModifyApVlanConfig) | **Put** /openapi/v2/{omadacId}/sites/{siteId}/aps/vlan | Batch Modify Ap Vlan Config
[**BatchModifyMultiApPorts**](ApAPI.md#BatchModifyMultiApPorts) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/ports/config | Batch modify multiple aps port config
[**BatchModifyPowerSavingConfig**](ApAPI.md#BatchModifyPowerSavingConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/cmd/aps/config/power-saving | Batch Modify ap power saving config
[**ChangeP2pRole**](ApAPI.md#ChangeP2pRole) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/change-role | Change p2p ap role
[**GetAfcConfig**](ApAPI.md#GetAfcConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/afc-config | Get ap AFC config
[**GetAfcInfo**](ApAPI.md#GetAfcInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/afc | Get ap AFC Status Info
[**GetAllApBriefInfosBySite**](ApAPI.md#GetAllApBriefInfosBySite) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/lockToAp/aps | Get Lock to Ap brief APs
[**GetAntSwitchConfig**](ApAPI.md#GetAntSwitchConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ant-switch | Get AP antSwitch Setting
[**GetAntennaGainConfig**](ApAPI.md#GetAntennaGainConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/antenna-gain | Get AP antenna gain config
[**GetApBridgeInfo**](ApAPI.md#GetApBridgeInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/bridge | Get P2P bridge config
[**GetApL3AccessConfig**](ApAPI.md#GetApL3AccessConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/l3access | Get Ap L3Access Config
[**GetApListChannelInfo**](ApAPI.md#GetApListChannelInfo) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/channel-info | Get ap list channel info.
[**GetApLldpConfig**](ApAPI.md#GetApLldpConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/lldp | Get Ap LLDP Config
[**GetApLoadBalanceConfig**](ApAPI.md#GetApLoadBalanceConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/load-balance | Get Ap Load Balance Config
[**GetApOfdmaConfig**](ApAPI.md#GetApOfdmaConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ofdma | Get ap OFDMA config
[**GetApP2pInfo**](ApAPI.md#GetApP2pInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/p2pInfo | Get P2P bridge group info
[**GetApPortList**](ApAPI.md#GetApPortList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ports | Get Ap port list
[**GetApPortVlans**](ApAPI.md#GetApPortVlans) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/port-vlans | Get ap current vlans config
[**GetApQoSConfig**](ApAPI.md#GetApQoSConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/qos | Get Ap QoS Config
[**GetApSnmpConfig**](ApAPI.md#GetApSnmpConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/snmp | Get Ap Snmp Config
[**GetApTrunkSettingConfig**](ApAPI.md#GetApTrunkSettingConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/trunk-setting | Get Ap Trunk Setting Config
[**GetApUplinkConfig**](ApAPI.md#GetApUplinkConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/uplink-config | Get Ap Uplink Config
[**GetApVlanConfig**](ApAPI.md#GetApVlanConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/vlan | Get Ap Vlan Config
[**GetApVlanConfigV2**](ApAPI.md#GetApVlanConfigV2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/aps/{apMac}/vlan | Get Ap Vlan Config v2
[**GetAvailableChannelOfAp**](ApAPI.md#GetAvailableChannelOfAp) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/available-channel | Get available channel list of ap
[**GetChannelLimitConfig**](ApAPI.md#GetChannelLimitConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/channel-limit | Get ap channel limit config
[**GetDownlinkWiredDevices**](ApAPI.md#GetDownlinkWiredDevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/wired-downlink | Get ap downlink(wired) device list
[**GetGeneralConfig2**](ApAPI.md#GetGeneralConfig2) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/general-config | Get ap general config
[**GetGridMeshCandiParentsForAdopt**](ApAPI.md#GetGridMeshCandiParentsForAdopt) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/grid/mesh/candiParents | Get the candidate parent AP of the mesh AP
[**GetIpSettingConfig**](ApAPI.md#GetIpSettingConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ip-setting | Get ap ip setting
[**GetIpv6SettingConfig**](ApAPI.md#GetIpv6SettingConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ipv6-setting | Get ap ipv6 setting
[**GetLanDetail**](ApAPI.md#GetLanDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/lan-traffic-info | Get ap lan traffic info
[**GetMeshStatistics**](ApAPI.md#GetMeshStatistics) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/mesh/statistics | Get ap mesh statistics
[**GetMeshStatisticsForMsp**](ApAPI.md#GetMeshStatisticsForMsp) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/aps/{apMac}/mesh/statistics | Get ap mesh statistics for MSP view
[**GetMultiApPortList**](ApAPI.md#GetMultiApPortList) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/ports/capability | Get multiple aps port list
[**GetOverviewDetail**](ApAPI.md#GetOverviewDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac} | Get ap info
[**GetPowerSavingConfig**](ApAPI.md#GetPowerSavingConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/power-saving | Get ap power saving config
[**GetRFScanResult**](ApAPI.md#GetRFScanResult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/rf-scan-result | Get rf scan result
[**GetRFScanResultV2**](ApAPI.md#GetRFScanResultV2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/aps/{apMac}/rf-scan-result | Get rf scan result v2
[**GetRadiosConfig**](ApAPI.md#GetRadiosConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/radio-config | Get ap radio config
[**GetRadiosDetail**](ApAPI.md#GetRadiosDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/radios | Get ap radio detail
[**GetSpeedTestResults**](ApAPI.md#GetSpeedTestResults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/speed-test-result | Get speed test result
[**GetUplinkWiredDetail**](ApAPI.md#GetUplinkWiredDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/wired-uplink | Get ap uplink(wired) detail
[**GetWlansConfig**](ApAPI.md#GetWlansConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/override | Get ap WLANs override config
[**GetWlansConfigV2**](ApAPI.md#GetWlansConfigV2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/aps/{apMac}/override | Get ap WLANs override config V2
[**ModifyAfcConfig**](ApAPI.md#ModifyAfcConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/afc-config | Modify ap AFC config
[**ModifyAntSwitchConfig**](ApAPI.md#ModifyAntSwitchConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ant-switch | Modify AP antSwitch Setting
[**ModifyAntennaGainConfig**](ApAPI.md#ModifyAntennaGainConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/antenna-gain | Modify AP antenna gain config
[**ModifyApBridgeInfo**](ApAPI.md#ModifyApBridgeInfo) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/bridge | Modify P2P bridge config
[**ModifyApChannelConfig**](ApAPI.md#ModifyApChannelConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/channel-config | Modify ap channel config
[**ModifyApLoadBalanceConfig**](ApAPI.md#ModifyApLoadBalanceConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/load-balance | Modify Ap Load Balance Config
[**ModifyApOfdmaConfig**](ApAPI.md#ModifyApOfdmaConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ofdma | Modify ap OFDMA config
[**ModifyApPort**](ApAPI.md#ModifyApPort) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ports/{port} | Modify Ap Port Config
[**ModifyApQosConfig**](ApAPI.md#ModifyApQosConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/qos | Modify Ap QoS Config
[**ModifyApServicesConfig**](ApAPI.md#ModifyApServicesConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/service-config | Modify AP Service Config
[**ModifyApTrunkSettingConfig**](ApAPI.md#ModifyApTrunkSettingConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/trunk-setting | Modify Ap Trunk Setting Config
[**ModifyApUpLinkConfig**](ApAPI.md#ModifyApUpLinkConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/uplink-config | Modify Ap Uplink Config
[**ModifyApVlanConfig**](ApAPI.md#ModifyApVlanConfig) | **Put** /openapi/v2/{omadacId}/sites/{siteId}/aps/{apMac}/vlan | Modify Ap Vlan Config
[**ModifyApWlanGroup**](ApAPI.md#ModifyApWlanGroup) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/wlan-group | Switch ap&#39;s wlan group
[**ModifyChannelLimitConfig**](ApAPI.md#ModifyChannelLimitConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/channel-limit | Modify ap channel limit config
[**ModifyGeneralConfig2**](ApAPI.md#ModifyGeneralConfig2) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/general-config | Modify ap general config
[**ModifyIpSettingConfig**](ApAPI.md#ModifyIpSettingConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ip-setting | Modify ap ip setting
[**ModifyIpv6SettingConfig**](ApAPI.md#ModifyIpv6SettingConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/ipv6-setting | Modify ap ipv6 setting
[**ModifyPowerSavingConfig**](ApAPI.md#ModifyPowerSavingConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/power-saving | Modify ap power saving config
[**ModifyRadiosConfig**](ApAPI.md#ModifyRadiosConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/radio-config | Modify ap radio config
[**ModifyWlansConfig**](ApAPI.md#ModifyWlansConfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/override | Modify ap WLANs override config
[**ModifyWlansConfigV2**](ApAPI.md#ModifyWlansConfigV2) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/aps/{apMac}/override | Modify ap WLANs override config V2
[**MspMoveToCustomer3**](ApAPI.md#MspMoveToCustomer3) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/cmd/aps/move | Move site to target customer
[**TriggerRadioFrequencyScan**](ApAPI.md#TriggerRadioFrequencyScan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/start-rf-scan | Start rf scan
[**TriggerRadioFrequencyScanV2**](ApAPI.md#TriggerRadioFrequencyScanV2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/aps/{apMac}/start-rf-scan | Start rf scan v2
[**TriggerSpeedTest**](ApAPI.md#TriggerSpeedTest) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/start-speed-test | Start speed test
[**UpdateApLocationUseGps**](ApAPI.md#UpdateApLocationUseGps) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/aps/{apMac}/location-gps | Update ap location info use GPS



## BatchModifyAntSwitchConfig

> OperationResponseWithoutResult BatchModifyAntSwitchConfig(ctx, omadacId, siteId).BatchUpdateApAntSwitchOpenApiVO(batchUpdateApAntSwitchOpenApiVO).Execute()

Batch Modify AP antSwitch Setting



### Example

```go
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
	batchUpdateApAntSwitchOpenApiVO := *openapiclient.NewBatchUpdateApAntSwitchOpenApiVO() // BatchUpdateApAntSwitchOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.BatchModifyAntSwitchConfig(context.Background(), omadacId, siteId).BatchUpdateApAntSwitchOpenApiVO(batchUpdateApAntSwitchOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.BatchModifyAntSwitchConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyAntSwitchConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.BatchModifyAntSwitchConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyAntSwitchConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchUpdateApAntSwitchOpenApiVO** | [**BatchUpdateApAntSwitchOpenApiVO**](BatchUpdateApAntSwitchOpenApiVO.md) |  | 

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


## BatchModifyApVlanConfig

> OperationResponseWithoutResult BatchModifyApVlanConfig(ctx, omadacId, siteId).BatchUpdateApVlanOpenApiVO(batchUpdateApVlanOpenApiVO).Execute()

Batch Modify Ap Vlan Config



### Example

```go
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
	batchUpdateApVlanOpenApiVO := *openapiclient.NewBatchUpdateApVlanOpenApiVO([]string{"ApMacList_example"}) // BatchUpdateApVlanOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.BatchModifyApVlanConfig(context.Background(), omadacId, siteId).BatchUpdateApVlanOpenApiVO(batchUpdateApVlanOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.BatchModifyApVlanConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyApVlanConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.BatchModifyApVlanConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyApVlanConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchUpdateApVlanOpenApiVO** | [**BatchUpdateApVlanOpenApiVO**](BatchUpdateApVlanOpenApiVO.md) |  | 

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


## BatchModifyMultiApPorts

> OperationResponseWithoutResult BatchModifyMultiApPorts(ctx, omadacId, siteId).BatchUpdateMultiApPortsOpenApiVO(batchUpdateMultiApPortsOpenApiVO).Execute()

Batch modify multiple aps port config



### Example

```go
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
	batchUpdateMultiApPortsOpenApiVO := *openapiclient.NewBatchUpdateMultiApPortsOpenApiVO([]string{"ApMacList_example"}, []string{"LanPortList_example"}) // BatchUpdateMultiApPortsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.BatchModifyMultiApPorts(context.Background(), omadacId, siteId).BatchUpdateMultiApPortsOpenApiVO(batchUpdateMultiApPortsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.BatchModifyMultiApPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyMultiApPorts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.BatchModifyMultiApPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyMultiApPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchUpdateMultiApPortsOpenApiVO** | [**BatchUpdateMultiApPortsOpenApiVO**](BatchUpdateMultiApPortsOpenApiVO.md) |  | 

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


## BatchModifyPowerSavingConfig

> OperationResponseWithoutResult BatchModifyPowerSavingConfig(ctx, omadacId, siteId).BatchUpdateApPowerSavingConfigOpenApiVO(batchUpdateApPowerSavingConfigOpenApiVO).Execute()

Batch Modify ap power saving config



### Example

```go
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
	batchUpdateApPowerSavingConfigOpenApiVO := *openapiclient.NewBatchUpdateApPowerSavingConfigOpenApiVO([]string{"ApMacList_example"}) // BatchUpdateApPowerSavingConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.BatchModifyPowerSavingConfig(context.Background(), omadacId, siteId).BatchUpdateApPowerSavingConfigOpenApiVO(batchUpdateApPowerSavingConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.BatchModifyPowerSavingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyPowerSavingConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.BatchModifyPowerSavingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyPowerSavingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchUpdateApPowerSavingConfigOpenApiVO** | [**BatchUpdateApPowerSavingConfigOpenApiVO**](BatchUpdateApPowerSavingConfigOpenApiVO.md) |  | 

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


## ChangeP2pRole

> OperationResponseWithoutResult ChangeP2pRole(ctx, omadacId, siteId).ChangeP2pRole(changeP2pRole).Execute()

Change p2p ap role



### Example

```go
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
	changeP2pRole := *openapiclient.NewChangeP2pRole("ClientMac_example", "MainMac_example") // ChangeP2pRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ChangeP2pRole(context.Background(), omadacId, siteId).ChangeP2pRole(changeP2pRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ChangeP2pRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeP2pRole`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ChangeP2pRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeP2pRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **changeP2pRole** | [**ChangeP2pRole**](ChangeP2pRole.md) |  | 

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


## GetAfcConfig

> OperationResponseAfcConfigOpenApiVO GetAfcConfig(ctx, omadacId, siteId, apMac).Execute()

Get ap AFC config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetAfcConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetAfcConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAfcConfig`: OperationResponseAfcConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetAfcConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAfcConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseAfcConfigOpenApiVO**](OperationResponseAfcConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAfcInfo

> OperationResponseApAfcInfoOpenApiVO GetAfcInfo(ctx, omadacId, siteId, apMac).Execute()

Get ap AFC Status Info



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetAfcInfo(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetAfcInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAfcInfo`: OperationResponseApAfcInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetAfcInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAfcInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApAfcInfoOpenApiVO**](OperationResponseApAfcInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllApBriefInfosBySite

> OperationResponseListApBriefInfo GetAllApBriefInfosBySite(ctx, omadacId, siteId).Execute()

Get Lock to Ap brief APs



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetAllApBriefInfosBySite(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetAllApBriefInfosBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllApBriefInfosBySite`: OperationResponseListApBriefInfo
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetAllApBriefInfosBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApBriefInfosBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListApBriefInfo**](OperationResponseListApBriefInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAntSwitchConfig

> OperationResponseApAntSwitchConfigOpenApiVO GetAntSwitchConfig(ctx, omadacId, siteId, apMac).Execute()

Get AP antSwitch Setting



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetAntSwitchConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetAntSwitchConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAntSwitchConfig`: OperationResponseApAntSwitchConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetAntSwitchConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAntSwitchConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApAntSwitchConfigOpenApiVO**](OperationResponseApAntSwitchConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAntennaGainConfig

> OperationResponseApAnteGainConfig GetAntennaGainConfig(ctx, omadacId, siteId, apMac).Execute()

Get AP antenna gain config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetAntennaGainConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetAntennaGainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAntennaGainConfig`: OperationResponseApAnteGainConfig
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetAntennaGainConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAntennaGainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApAnteGainConfig**](OperationResponseApAnteGainConfig.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApBridgeInfo

> OperationResponseApBridgeConfig GetApBridgeInfo(ctx, omadacId, siteId, apMac).Execute()

Get P2P bridge config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApBridgeInfo(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApBridgeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApBridgeInfo`: OperationResponseApBridgeConfig
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApBridgeInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApBridgeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApBridgeConfig**](OperationResponseApBridgeConfig.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApL3AccessConfig

> OperationResponseApL3AccessConfigOpenApiVO GetApL3AccessConfig(ctx, omadacId, siteId, apMac).Execute()

Get Ap L3Access Config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApL3AccessConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApL3AccessConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApL3AccessConfig`: OperationResponseApL3AccessConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApL3AccessConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApL3AccessConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApL3AccessConfigOpenApiVO**](OperationResponseApL3AccessConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApListChannelInfo

> ApChannelInfoOpenApiVO GetApListChannelInfo(ctx, omadacId, siteId).APMACList(aPMACList).Execute()

Get ap list channel info.



### Example

```go
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
	aPMACList := *openapiclient.NewAPMACList() // APMACList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.GetApListChannelInfo(context.Background(), omadacId, siteId).APMACList(aPMACList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApListChannelInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApListChannelInfo`: ApChannelInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApListChannelInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApListChannelInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aPMACList** | [**APMACList**](APMACList.md) |  | 

### Return type

[**ApChannelInfoOpenApiVO**](ApChannelInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApLldpConfig

> OperationResponseApLldpConfigOpenApiVO GetApLldpConfig(ctx, omadacId, siteId, apMac).Execute()

Get Ap LLDP Config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApLldpConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApLldpConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApLldpConfig`: OperationResponseApLldpConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApLldpConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApLldpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApLldpConfigOpenApiVO**](OperationResponseApLldpConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApLoadBalanceConfig

> OperationResponseApAdvancedLoadBalanceOpenApiVO GetApLoadBalanceConfig(ctx, omadacId, siteId, apMac).Execute()

Get Ap Load Balance Config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApLoadBalanceConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApLoadBalanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApLoadBalanceConfig`: OperationResponseApAdvancedLoadBalanceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApLoadBalanceConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApLoadBalanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApAdvancedLoadBalanceOpenApiVO**](OperationResponseApAdvancedLoadBalanceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApOfdmaConfig

> OperationResponseApOfdmaConfigOpenApiVO GetApOfdmaConfig(ctx, omadacId, siteId, apMac).Execute()

Get ap OFDMA config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApOfdmaConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApOfdmaConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApOfdmaConfig`: OperationResponseApOfdmaConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApOfdmaConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApOfdmaConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApOfdmaConfigOpenApiVO**](OperationResponseApOfdmaConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApP2pInfo

> OperationResponseApP2pInfo GetApP2pInfo(ctx, omadacId, siteId, apMac).Execute()

Get P2P bridge group info



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApP2pInfo(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApP2pInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApP2pInfo`: OperationResponseApP2pInfo
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApP2pInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApP2pInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApP2pInfo**](OperationResponseApP2pInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApPortList

> OperationResponseListAPLANPortList GetApPortList(ctx, omadacId, siteId, apMac).Execute()

Get Ap port list



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApPortList(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApPortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApPortList`: OperationResponseListAPLANPortList
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApPortList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApPortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListAPLANPortList**](OperationResponseListAPLANPortList.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApPortVlans

> OperationResponseGridVOApVlansVO GetApPortVlans(ctx, omadacId, siteId, apMac).Page(page).PageSize(pageSize).Execute()

Get ap current vlans config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApPortVlans(context.Background(), omadacId, siteId, apMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApPortVlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApPortVlans`: OperationResponseGridVOApVlansVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApPortVlans`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApPortVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOApVlansVO**](OperationResponseGridVOApVlansVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApQoSConfig

> OperationResponseApAdvancedQosOpenApiVO GetApQoSConfig(ctx, omadacId, siteId, apMac).Execute()

Get Ap QoS Config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApQoSConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApQoSConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApQoSConfig`: OperationResponseApAdvancedQosOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApQoSConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApQoSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApAdvancedQosOpenApiVO**](OperationResponseApAdvancedQosOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApSnmpConfig

> OperationResponseApSnmpConfigOpenApiVO GetApSnmpConfig(ctx, omadacId, siteId, apMac).Execute()

Get Ap Snmp Config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApSnmpConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApSnmpConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApSnmpConfig`: OperationResponseApSnmpConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApSnmpConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApSnmpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApSnmpConfigOpenApiVO**](OperationResponseApSnmpConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApTrunkSettingConfig

> OperationResponseApTrunkSettingOpenApiVO GetApTrunkSettingConfig(ctx, omadacId, siteId, apMac).Execute()

Get Ap Trunk Setting Config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApTrunkSettingConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApTrunkSettingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApTrunkSettingConfig`: OperationResponseApTrunkSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApTrunkSettingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApTrunkSettingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApTrunkSettingOpenApiVO**](OperationResponseApTrunkSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApUplinkConfig

> OperationResponseListApUplinkConfigOpenApiVO GetApUplinkConfig(ctx, omadacId, siteId, apMac).Execute()

Get Ap Uplink Config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApUplinkConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApUplinkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApUplinkConfig`: OperationResponseListApUplinkConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApUplinkConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApUplinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListApUplinkConfigOpenApiVO**](OperationResponseListApUplinkConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApVlanConfig

> OperationResponseApVlanConfigOpenApiVO GetApVlanConfig(ctx, omadacId, siteId, apMac).Execute()

Get Ap Vlan Config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApVlanConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApVlanConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApVlanConfig`: OperationResponseApVlanConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApVlanConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApVlanConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApVlanConfigOpenApiVO**](OperationResponseApVlanConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApVlanConfigV2

> OperationResponseApVlanConfigV2OpenApiVO GetApVlanConfigV2(ctx, omadacId, siteId, apMac).Execute()

Get Ap Vlan Config v2



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetApVlanConfigV2(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetApVlanConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApVlanConfigV2`: OperationResponseApVlanConfigV2OpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetApVlanConfigV2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetApVlanConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApVlanConfigV2OpenApiVO**](OperationResponseApVlanConfigV2OpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableChannelOfAp

> OperationResponseApAvailableChannelOpenApiVO GetAvailableChannelOfAp(ctx, omadacId, siteId, apMac).Execute()

Get available channel list of ap



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetAvailableChannelOfAp(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetAvailableChannelOfAp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableChannelOfAp`: OperationResponseApAvailableChannelOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetAvailableChannelOfAp`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAvailableChannelOfApRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApAvailableChannelOpenApiVO**](OperationResponseApAvailableChannelOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelLimitConfig

> OperationResponseChannelLimitConfigOpenApiVO GetChannelLimitConfig(ctx, omadacId, siteId, apMac).Execute()

Get ap channel limit config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetChannelLimitConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetChannelLimitConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelLimitConfig`: OperationResponseChannelLimitConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetChannelLimitConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetChannelLimitConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseChannelLimitConfigOpenApiVO**](OperationResponseChannelLimitConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownlinkWiredDevices

> OperationResponseApWiredDownlink GetDownlinkWiredDevices(ctx, omadacId, siteId, apMac).Execute()

Get ap downlink(wired) device list



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetDownlinkWiredDevices(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetDownlinkWiredDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDownlinkWiredDevices`: OperationResponseApWiredDownlink
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetDownlinkWiredDevices`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDownlinkWiredDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApWiredDownlink**](OperationResponseApWiredDownlink.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneralConfig2

> OperationResponseApGeneralConfig GetGeneralConfig2(ctx, omadacId, siteId, apMac).Execute()

Get ap general config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetGeneralConfig2(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetGeneralConfig2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGeneralConfig2`: OperationResponseApGeneralConfig
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetGeneralConfig2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGeneralConfig2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApGeneralConfig**](OperationResponseApGeneralConfig.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridMeshCandiParentsForAdopt

> OperationResponseGridVOCandidateParentForAdoptOpenApiVO GetGridMeshCandiParentsForAdopt(ctx, omadacId, siteId, apMac).Page(page).PageSize(pageSize).Execute()

Get the candidate parent AP of the mesh AP



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetGridMeshCandiParentsForAdopt(context.Background(), omadacId, siteId, apMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetGridMeshCandiParentsForAdopt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridMeshCandiParentsForAdopt`: OperationResponseGridVOCandidateParentForAdoptOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetGridMeshCandiParentsForAdopt`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridMeshCandiParentsForAdoptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOCandidateParentForAdoptOpenApiVO**](OperationResponseGridVOCandidateParentForAdoptOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpSettingConfig

> OperationResponseApIPSetting GetIpSettingConfig(ctx, omadacId, siteId, apMac).Execute()

Get ap ip setting



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetIpSettingConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetIpSettingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpSettingConfig`: OperationResponseApIPSetting
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetIpSettingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetIpSettingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApIPSetting**](OperationResponseApIPSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpv6SettingConfig

> OperationResponseApIPv6Setting GetIpv6SettingConfig(ctx, omadacId, siteId, apMac).Execute()

Get ap ipv6 setting



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetIpv6SettingConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetIpv6SettingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpv6SettingConfig`: OperationResponseApIPv6Setting
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetIpv6SettingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetIpv6SettingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApIPv6Setting**](OperationResponseApIPv6Setting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanDetail

> OperationResponseApLanTraffic GetLanDetail(ctx, omadacId, siteId, apMac).Execute()

Get ap lan traffic info



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetLanDetail(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetLanDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanDetail`: OperationResponseApLanTraffic
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetLanDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetLanDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApLanTraffic**](OperationResponseApLanTraffic.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeshStatistics

> OperationResponseApMeshStatisticsOpenApiVO GetMeshStatistics(ctx, omadacId, siteId, apMac).Execute()

Get ap mesh statistics



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetMeshStatistics(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetMeshStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeshStatistics`: OperationResponseApMeshStatisticsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetMeshStatistics`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMeshStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApMeshStatisticsOpenApiVO**](OperationResponseApMeshStatisticsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeshStatisticsForMsp

> OperationResponseApMeshStatisticsOpenApiVO GetMeshStatisticsForMsp(ctx, mspId, customerId, siteId, apMac).Execute()

Get ap mesh statistics for MSP view



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.GetMeshStatisticsForMsp(context.Background(), mspId, customerId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetMeshStatisticsForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeshStatisticsForMsp`: OperationResponseApMeshStatisticsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetMeshStatisticsForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeshStatisticsForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseApMeshStatisticsOpenApiVO**](OperationResponseApMeshStatisticsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultiApPortList

> OperationResponseListAPLANPortList GetMultiApPortList(ctx, omadacId, siteId).MultiApPortCapOpenApiVO(multiApPortCapOpenApiVO).Execute()

Get multiple aps port list



### Example

```go
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
	multiApPortCapOpenApiVO := *openapiclient.NewMultiApPortCapOpenApiVO([]string{"ApMacList_example"}) // MultiApPortCapOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.GetMultiApPortList(context.Background(), omadacId, siteId).MultiApPortCapOpenApiVO(multiApPortCapOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetMultiApPortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultiApPortList`: OperationResponseListAPLANPortList
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetMultiApPortList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiApPortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **multiApPortCapOpenApiVO** | [**MultiApPortCapOpenApiVO**](MultiApPortCapOpenApiVO.md) |  | 

### Return type

[**OperationResponseListAPLANPortList**](OperationResponseListAPLANPortList.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverviewDetail

> OperationResponseApOverviewInfo GetOverviewDetail(ctx, omadacId, siteId, apMac).Execute()

Get ap info



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetOverviewDetail(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetOverviewDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverviewDetail`: OperationResponseApOverviewInfo
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetOverviewDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOverviewDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApOverviewInfo**](OperationResponseApOverviewInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPowerSavingConfig

> OperationResponseApPowerSavingConfigOpenApiVO GetPowerSavingConfig(ctx, omadacId, siteId, apMac).Execute()

Get ap power saving config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetPowerSavingConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetPowerSavingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPowerSavingConfig`: OperationResponseApPowerSavingConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetPowerSavingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPowerSavingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApPowerSavingConfigOpenApiVO**](OperationResponseApPowerSavingConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRFScanResult

> OperationResponseApRFScanInfo GetRFScanResult(ctx, omadacId, siteId, apMac).Execute()

Get rf scan result



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetRFScanResult(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetRFScanResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRFScanResult`: OperationResponseApRFScanInfo
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetRFScanResult`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetRFScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApRFScanInfo**](OperationResponseApRFScanInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRFScanResultV2

> OperationResponseApRFScanResult GetRFScanResultV2(ctx, omadacId, siteId, apMac).Execute()

Get rf scan result v2



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetRFScanResultV2(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetRFScanResultV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRFScanResultV2`: OperationResponseApRFScanResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetRFScanResultV2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetRFScanResultV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApRFScanResult**](OperationResponseApRFScanResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadiosConfig

> OperationResponseApRadiosConfig GetRadiosConfig(ctx, omadacId, siteId, apMac).Execute()

Get ap radio config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetRadiosConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetRadiosConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadiosConfig`: OperationResponseApRadiosConfig
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetRadiosConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetRadiosConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApRadiosConfig**](OperationResponseApRadiosConfig.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadiosDetail

> OperationResponseApRadiosDetail GetRadiosDetail(ctx, omadacId, siteId, apMac).Execute()

Get ap radio detail



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetRadiosDetail(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetRadiosDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadiosDetail`: OperationResponseApRadiosDetail
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetRadiosDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetRadiosDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApRadiosDetail**](OperationResponseApRadiosDetail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpeedTestResults

> OperationResponseApSpeedTestResultsOpenApiVO GetSpeedTestResults(ctx, omadacId, siteId, apMac).Execute()

Get speed test result



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetSpeedTestResults(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetSpeedTestResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpeedTestResults`: OperationResponseApSpeedTestResultsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetSpeedTestResults`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSpeedTestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApSpeedTestResultsOpenApiVO**](OperationResponseApSpeedTestResultsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUplinkWiredDetail

> OperationResponseApWiredUplink GetUplinkWiredDetail(ctx, omadacId, siteId, apMac).Execute()

Get ap uplink(wired) detail



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetUplinkWiredDetail(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetUplinkWiredDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUplinkWiredDetail`: OperationResponseApWiredUplink
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetUplinkWiredDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetUplinkWiredDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApWiredUplink**](OperationResponseApWiredUplink.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWlansConfig

> OperationResponseApSsidWlansOpenApiVO GetWlansConfig(ctx, omadacId, siteId, apMac).Execute()

Get ap WLANs override config



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetWlansConfig(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetWlansConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWlansConfig`: OperationResponseApSsidWlansOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetWlansConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWlansConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApSsidWlansOpenApiVO**](OperationResponseApSsidWlansOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWlansConfigV2

> OperationResponseApSsidWlansOpenApiV2VO GetWlansConfigV2(ctx, omadacId, siteId, apMac).Execute()

Get ap WLANs override config V2



### Example

```go
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
	resp, r, err := apiClient.ApAPI.GetWlansConfigV2(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.GetWlansConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWlansConfigV2`: OperationResponseApSsidWlansOpenApiV2VO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.GetWlansConfigV2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWlansConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApSsidWlansOpenApiV2VO**](OperationResponseApSsidWlansOpenApiV2VO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAfcConfig

> OperationResponseWithoutResult ModifyAfcConfig(ctx, omadacId, siteId, apMac).UpdateAfcConfigOpenApiVO(updateAfcConfigOpenApiVO).Execute()

Modify ap AFC config



### Example

```go
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
	updateAfcConfigOpenApiVO := *openapiclient.NewUpdateAfcConfigOpenApiVO(false) // UpdateAfcConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyAfcConfig(context.Background(), omadacId, siteId, apMac).UpdateAfcConfigOpenApiVO(updateAfcConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyAfcConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAfcConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyAfcConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyAfcConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateAfcConfigOpenApiVO** | [**UpdateAfcConfigOpenApiVO**](UpdateAfcConfigOpenApiVO.md) |  | 

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


## ModifyAntSwitchConfig

> OperationResponseWithoutResult ModifyAntSwitchConfig(ctx, omadacId, siteId, apMac).UpdateApAntSwitchConfig(updateApAntSwitchConfig).Execute()

Modify AP antSwitch Setting



### Example

```go
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
	updateApAntSwitchConfig := *openapiclient.NewUpdateApAntSwitchConfig() // UpdateApAntSwitchConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyAntSwitchConfig(context.Background(), omadacId, siteId, apMac).UpdateApAntSwitchConfig(updateApAntSwitchConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyAntSwitchConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAntSwitchConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyAntSwitchConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyAntSwitchConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApAntSwitchConfig** | [**UpdateApAntSwitchConfig**](UpdateApAntSwitchConfig.md) |  | 

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


## ModifyAntennaGainConfig

> OperationResponseWithoutResult ModifyAntennaGainConfig(ctx, omadacId, siteId, apMac).UpdateApAnteGainConfig(updateApAnteGainConfig).Execute()

Modify AP antenna gain config



### Example

```go
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
	updateApAnteGainConfig := *openapiclient.NewUpdateApAnteGainConfig() // UpdateApAnteGainConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyAntennaGainConfig(context.Background(), omadacId, siteId, apMac).UpdateApAnteGainConfig(updateApAnteGainConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyAntennaGainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAntennaGainConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyAntennaGainConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyAntennaGainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApAnteGainConfig** | [**UpdateApAnteGainConfig**](UpdateApAnteGainConfig.md) |  | 

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


## ModifyApBridgeInfo

> OperationResponseWithoutResult ModifyApBridgeInfo(ctx, omadacId, siteId, apMac).ApBridgeConfig(apBridgeConfig).Execute()

Modify P2P bridge config



### Example

```go
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
	apBridgeConfig := *openapiclient.NewApBridgeConfig("BridgeSsidName_example", "BridgeSsidPassword_example") // ApBridgeConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApBridgeInfo(context.Background(), omadacId, siteId, apMac).ApBridgeConfig(apBridgeConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApBridgeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApBridgeInfo`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApBridgeInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApBridgeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apBridgeConfig** | [**ApBridgeConfig**](ApBridgeConfig.md) |  | 

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


## ModifyApChannelConfig

> OperationResponseWithoutResult ModifyApChannelConfig(ctx, omadacId, siteId, apMac).UpdateApChannelConfigOpenApiVO(updateApChannelConfigOpenApiVO).Execute()

Modify ap channel config



### Example

```go
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
	updateApChannelConfigOpenApiVO := *openapiclient.NewUpdateApChannelConfigOpenApiVO(false, int32(123)) // UpdateApChannelConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApChannelConfig(context.Background(), omadacId, siteId, apMac).UpdateApChannelConfigOpenApiVO(updateApChannelConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApChannelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApChannelConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApChannelConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApChannelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApChannelConfigOpenApiVO** | [**UpdateApChannelConfigOpenApiVO**](UpdateApChannelConfigOpenApiVO.md) |  | 

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


## ModifyApLoadBalanceConfig

> OperationResponseWithoutResult ModifyApLoadBalanceConfig(ctx, omadacId, siteId, apMac).UpdateApAdvancedLoadBalanceOpenApiVO(updateApAdvancedLoadBalanceOpenApiVO).Execute()

Modify Ap Load Balance Config



### Example

```go
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
	updateApAdvancedLoadBalanceOpenApiVO := *openapiclient.NewUpdateApAdvancedLoadBalanceOpenApiVO() // UpdateApAdvancedLoadBalanceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApLoadBalanceConfig(context.Background(), omadacId, siteId, apMac).UpdateApAdvancedLoadBalanceOpenApiVO(updateApAdvancedLoadBalanceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApLoadBalanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApLoadBalanceConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApLoadBalanceConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApLoadBalanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApAdvancedLoadBalanceOpenApiVO** | [**UpdateApAdvancedLoadBalanceOpenApiVO**](UpdateApAdvancedLoadBalanceOpenApiVO.md) |  | 

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


## ModifyApOfdmaConfig

> OperationResponseWithoutResult ModifyApOfdmaConfig(ctx, omadacId, siteId, apMac).UpdateApOfdmaConfigOpenApiVO(updateApOfdmaConfigOpenApiVO).Execute()

Modify ap OFDMA config



### Example

```go
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
	updateApOfdmaConfigOpenApiVO := *openapiclient.NewUpdateApOfdmaConfigOpenApiVO() // UpdateApOfdmaConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApOfdmaConfig(context.Background(), omadacId, siteId, apMac).UpdateApOfdmaConfigOpenApiVO(updateApOfdmaConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApOfdmaConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApOfdmaConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApOfdmaConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApOfdmaConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApOfdmaConfigOpenApiVO** | [**UpdateApOfdmaConfigOpenApiVO**](UpdateApOfdmaConfigOpenApiVO.md) |  | 

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


## ModifyApPort

> OperationResponseWithoutResult ModifyApPort(ctx, omadacId, siteId, apMac, port).ModifyAPLANPort(modifyAPLANPort).Execute()

Modify Ap Port Config



### Example

```go
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
	port := "port_example" // string | ap port
	modifyAPLANPort := *openapiclient.NewModifyAPLANPort() // ModifyAPLANPort | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApPort(context.Background(), omadacId, siteId, apMac, port).ModifyAPLANPort(modifyAPLANPort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApPort`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | ap port | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyApPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **modifyAPLANPort** | [**ModifyAPLANPort**](ModifyAPLANPort.md) |  | 

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


## ModifyApQosConfig

> OperationResponseWithoutResult ModifyApQosConfig(ctx, omadacId, siteId, apMac).ModifyApAdvancedQosOpenApiVO(modifyApAdvancedQosOpenApiVO).Execute()

Modify Ap QoS Config



### Example

```go
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
	modifyApAdvancedQosOpenApiVO := *openapiclient.NewModifyApAdvancedQosOpenApiVO() // ModifyApAdvancedQosOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApQosConfig(context.Background(), omadacId, siteId, apMac).ModifyApAdvancedQosOpenApiVO(modifyApAdvancedQosOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApQosConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApQosConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApQosConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApQosConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyApAdvancedQosOpenApiVO** | [**ModifyApAdvancedQosOpenApiVO**](ModifyApAdvancedQosOpenApiVO.md) |  | 

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


## ModifyApServicesConfig

> OperationResponseApServicesConfigOpenApiVO ModifyApServicesConfig(ctx, omadacId, siteId, apMac).ApServicesConfigOpenApiVO(apServicesConfigOpenApiVO).Execute()

Modify AP Service Config



### Example

```go
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
	apServicesConfigOpenApiVO := *openapiclient.NewApServicesConfigOpenApiVO() // ApServicesConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApServicesConfig(context.Background(), omadacId, siteId, apMac).ApServicesConfigOpenApiVO(apServicesConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApServicesConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApServicesConfig`: OperationResponseApServicesConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApServicesConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApServicesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apServicesConfigOpenApiVO** | [**ApServicesConfigOpenApiVO**](ApServicesConfigOpenApiVO.md) |  | 

### Return type

[**OperationResponseApServicesConfigOpenApiVO**](OperationResponseApServicesConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyApTrunkSettingConfig

> OperationResponseWithoutResult ModifyApTrunkSettingConfig(ctx, omadacId, siteId, apMac).UpdateApTrunkSettingOpenApiVO(updateApTrunkSettingOpenApiVO).Execute()

Modify Ap Trunk Setting Config



### Example

```go
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
	updateApTrunkSettingOpenApiVO := *openapiclient.NewUpdateApTrunkSettingOpenApiVO() // UpdateApTrunkSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApTrunkSettingConfig(context.Background(), omadacId, siteId, apMac).UpdateApTrunkSettingOpenApiVO(updateApTrunkSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApTrunkSettingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApTrunkSettingConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApTrunkSettingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApTrunkSettingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApTrunkSettingOpenApiVO** | [**UpdateApTrunkSettingOpenApiVO**](UpdateApTrunkSettingOpenApiVO.md) |  | 

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


## ModifyApUpLinkConfig

> ApUplinkConfigOpenApiVO ModifyApUpLinkConfig(ctx, omadacId, siteId, apMac).ApUplinkConfigOpenApiVO(apUplinkConfigOpenApiVO).Execute()

Modify Ap Uplink Config



### Example

```go
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
	apUplinkConfigOpenApiVO := *openapiclient.NewApUplinkConfigOpenApiVO(false) // ApUplinkConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApUpLinkConfig(context.Background(), omadacId, siteId, apMac).ApUplinkConfigOpenApiVO(apUplinkConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApUpLinkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApUpLinkConfig`: ApUplinkConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApUpLinkConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApUpLinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apUplinkConfigOpenApiVO** | [**ApUplinkConfigOpenApiVO**](ApUplinkConfigOpenApiVO.md) |  | 

### Return type

[**ApUplinkConfigOpenApiVO**](ApUplinkConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyApVlanConfig

> OperationResponseApVlanConfigV2OpenApiVO ModifyApVlanConfig(ctx, omadacId, siteId, apMac).UpdateApVlanOpenApiVO(updateApVlanOpenApiVO).Execute()

Modify Ap Vlan Config



### Example

```go
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
	updateApVlanOpenApiVO := *openapiclient.NewUpdateApVlanOpenApiVO() // UpdateApVlanOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApVlanConfig(context.Background(), omadacId, siteId, apMac).UpdateApVlanOpenApiVO(updateApVlanOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApVlanConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApVlanConfig`: OperationResponseApVlanConfigV2OpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApVlanConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApVlanConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApVlanOpenApiVO** | [**UpdateApVlanOpenApiVO**](UpdateApVlanOpenApiVO.md) |  | 

### Return type

[**OperationResponseApVlanConfigV2OpenApiVO**](OperationResponseApVlanConfigV2OpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyApWlanGroup

> OperationResponseWithoutResult ModifyApWlanGroup(ctx, omadacId, siteId, apMac).ApUpdateWlanGroupOpenApiVO(apUpdateWlanGroupOpenApiVO).Execute()

Switch ap's wlan group



### Example

```go
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
	apUpdateWlanGroupOpenApiVO := *openapiclient.NewApUpdateWlanGroupOpenApiVO("WlanGroupId_example") // ApUpdateWlanGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyApWlanGroup(context.Background(), omadacId, siteId, apMac).ApUpdateWlanGroupOpenApiVO(apUpdateWlanGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyApWlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApWlanGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyApWlanGroup`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyApWlanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apUpdateWlanGroupOpenApiVO** | [**ApUpdateWlanGroupOpenApiVO**](ApUpdateWlanGroupOpenApiVO.md) |  | 

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


## ModifyChannelLimitConfig

> OperationResponseWithoutResult ModifyChannelLimitConfig(ctx, omadacId, siteId, apMac).UpdateChannelLimitConfigOpenApiVO(updateChannelLimitConfigOpenApiVO).Execute()

Modify ap channel limit config



### Example

```go
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
	updateChannelLimitConfigOpenApiVO := *openapiclient.NewUpdateChannelLimitConfigOpenApiVO(int32(123)) // UpdateChannelLimitConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyChannelLimitConfig(context.Background(), omadacId, siteId, apMac).UpdateChannelLimitConfigOpenApiVO(updateChannelLimitConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyChannelLimitConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyChannelLimitConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyChannelLimitConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyChannelLimitConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateChannelLimitConfigOpenApiVO** | [**UpdateChannelLimitConfigOpenApiVO**](UpdateChannelLimitConfigOpenApiVO.md) |  | 

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


## ModifyGeneralConfig2

> OperationResponseWithoutResult ModifyGeneralConfig2(ctx, omadacId, siteId, apMac).ApGeneralConfig(apGeneralConfig).Execute()

Modify ap general config



### Example

```go
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
	apGeneralConfig := *openapiclient.NewApGeneralConfig() // ApGeneralConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyGeneralConfig2(context.Background(), omadacId, siteId, apMac).ApGeneralConfig(apGeneralConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyGeneralConfig2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGeneralConfig2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyGeneralConfig2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyGeneralConfig2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apGeneralConfig** | [**ApGeneralConfig**](ApGeneralConfig.md) |  | 

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


## ModifyIpSettingConfig

> OperationResponseWithoutResult ModifyIpSettingConfig(ctx, omadacId, siteId, apMac).ApIPSetting(apIPSetting).Execute()

Modify ap ip setting



### Example

```go
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
	apIPSetting := *openapiclient.NewApIPSetting("Mode_example") // ApIPSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyIpSettingConfig(context.Background(), omadacId, siteId, apMac).ApIPSetting(apIPSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyIpSettingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpSettingConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyIpSettingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyIpSettingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apIPSetting** | [**ApIPSetting**](ApIPSetting.md) |  | 

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


## ModifyIpv6SettingConfig

> OperationResponseWithoutResult ModifyIpv6SettingConfig(ctx, omadacId, siteId, apMac).ApIPv6Setting(apIPv6Setting).Execute()

Modify ap ipv6 setting



### Example

```go
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
	apIPv6Setting := *openapiclient.NewApIPv6Setting(false) // ApIPv6Setting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyIpv6SettingConfig(context.Background(), omadacId, siteId, apMac).ApIPv6Setting(apIPv6Setting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyIpv6SettingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpv6SettingConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyIpv6SettingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyIpv6SettingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apIPv6Setting** | [**ApIPv6Setting**](ApIPv6Setting.md) |  | 

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


## ModifyPowerSavingConfig

> OperationResponseWithoutResult ModifyPowerSavingConfig(ctx, omadacId, siteId, apMac).UpdateApPowerSavingConfigOpenApiVO(updateApPowerSavingConfigOpenApiVO).Execute()

Modify ap power saving config



### Example

```go
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
	updateApPowerSavingConfigOpenApiVO := *openapiclient.NewUpdateApPowerSavingConfigOpenApiVO(false, false) // UpdateApPowerSavingConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyPowerSavingConfig(context.Background(), omadacId, siteId, apMac).UpdateApPowerSavingConfigOpenApiVO(updateApPowerSavingConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyPowerSavingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPowerSavingConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyPowerSavingConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyPowerSavingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateApPowerSavingConfigOpenApiVO** | [**UpdateApPowerSavingConfigOpenApiVO**](UpdateApPowerSavingConfigOpenApiVO.md) |  | 

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


## ModifyRadiosConfig

> OperationResponseWithoutResult ModifyRadiosConfig(ctx, omadacId, siteId, apMac).ApRadiosConfig(apRadiosConfig).Execute()

Modify ap radio config



### Example

```go
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
	apRadiosConfig := *openapiclient.NewApRadiosConfig() // ApRadiosConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyRadiosConfig(context.Background(), omadacId, siteId, apMac).ApRadiosConfig(apRadiosConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyRadiosConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRadiosConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyRadiosConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyRadiosConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apRadiosConfig** | [**ApRadiosConfig**](ApRadiosConfig.md) |  | 

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


## ModifyWlansConfig

> OperationResponseWithoutResult ModifyWlansConfig(ctx, omadacId, siteId, apMac).ApSsidOverrideOpenApiVO(apSsidOverrideOpenApiVO).Execute()

Modify ap WLANs override config



### Example

```go
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
	apSsidOverrideOpenApiVO := *openapiclient.NewApSsidOverrideOpenApiVO([]openapiclient.SsidOverrideOpenApiVO{*openapiclient.NewSsidOverrideOpenApiVO(false, "OverrideSsidName_example", false, false, int32(123))}) // ApSsidOverrideOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyWlansConfig(context.Background(), omadacId, siteId, apMac).ApSsidOverrideOpenApiVO(apSsidOverrideOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyWlansConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWlansConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyWlansConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyWlansConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apSsidOverrideOpenApiVO** | [**ApSsidOverrideOpenApiVO**](ApSsidOverrideOpenApiVO.md) |  | 

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


## ModifyWlansConfigV2

> OperationResponseWithoutResult ModifyWlansConfigV2(ctx, omadacId, siteId, apMac).ApSsidOverrideOpenApiV2VO(apSsidOverrideOpenApiV2VO).Execute()

Modify ap WLANs override config V2



### Example

```go
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
	apSsidOverrideOpenApiV2VO := *openapiclient.NewApSsidOverrideOpenApiV2VO([]openapiclient.SsidOverrideOpenApiV2VO{*openapiclient.NewSsidOverrideOpenApiV2VO(false, false, int32(123))}) // ApSsidOverrideOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.ModifyWlansConfigV2(context.Background(), omadacId, siteId, apMac).ApSsidOverrideOpenApiV2VO(apSsidOverrideOpenApiV2VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.ModifyWlansConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWlansConfigV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.ModifyWlansConfigV2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyWlansConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apSsidOverrideOpenApiV2VO** | [**ApSsidOverrideOpenApiV2VO**](ApSsidOverrideOpenApiV2VO.md) |  | 

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


## MspMoveToCustomer3

> OperationResponseWithoutResult MspMoveToCustomer3(ctx, mspId, customerId, siteId).MoveToCustomerOpenApiVO(moveToCustomerOpenApiVO).Execute()

Move site to target customer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	moveToCustomerOpenApiVO := *openapiclient.NewMoveToCustomerOpenApiVO("Customer_example") // MoveToCustomerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.MspMoveToCustomer3(context.Background(), mspId, customerId, siteId).MoveToCustomerOpenApiVO(moveToCustomerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.MspMoveToCustomer3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MspMoveToCustomer3`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.MspMoveToCustomer3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMspMoveToCustomer3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **moveToCustomerOpenApiVO** | [**MoveToCustomerOpenApiVO**](MoveToCustomerOpenApiVO.md) |  | 

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


## TriggerRadioFrequencyScan

> OperationResponseWithoutResult TriggerRadioFrequencyScan(ctx, omadacId, siteId, apMac).Execute()

Start rf scan



### Example

```go
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
	resp, r, err := apiClient.ApAPI.TriggerRadioFrequencyScan(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.TriggerRadioFrequencyScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerRadioFrequencyScan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.TriggerRadioFrequencyScan`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiTriggerRadioFrequencyScanRequest struct via the builder pattern


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


## TriggerRadioFrequencyScanV2

> OperationResponseWithoutResult TriggerRadioFrequencyScanV2(ctx, omadacId, siteId, apMac).RFScanCommand(rFScanCommand).Execute()

Start rf scan v2



### Example

```go
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
	rFScanCommand := *openapiclient.NewRFScanCommand() // RFScanCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.TriggerRadioFrequencyScanV2(context.Background(), omadacId, siteId, apMac).RFScanCommand(rFScanCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.TriggerRadioFrequencyScanV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerRadioFrequencyScanV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.TriggerRadioFrequencyScanV2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiTriggerRadioFrequencyScanV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rFScanCommand** | [**RFScanCommand**](RFScanCommand.md) |  | 

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


## TriggerSpeedTest

> OperationResponseWithoutResult TriggerSpeedTest(ctx, omadacId, siteId, apMac).SpeedTestCommand(speedTestCommand).Execute()

Start speed test



### Example

```go
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
	speedTestCommand := *openapiclient.NewSpeedTestCommand("ChildMac_example") // SpeedTestCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApAPI.TriggerSpeedTest(context.Background(), omadacId, siteId, apMac).SpeedTestCommand(speedTestCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.TriggerSpeedTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerSpeedTest`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.TriggerSpeedTest`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiTriggerSpeedTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **speedTestCommand** | [**SpeedTestCommand**](SpeedTestCommand.md) |  | 

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


## UpdateApLocationUseGps

> OperationResponseApLocationConfigOpenApiVO UpdateApLocationUseGps(ctx, omadacId, siteId, apMac).Execute()

Update ap location info use GPS



### Example

```go
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
	resp, r, err := apiClient.ApAPI.UpdateApLocationUseGps(context.Background(), omadacId, siteId, apMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApAPI.UpdateApLocationUseGps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApLocationUseGps`: OperationResponseApLocationConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ApAPI.UpdateApLocationUseGps`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateApLocationUseGpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApLocationConfigOpenApiVO**](OperationResponseApLocationConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


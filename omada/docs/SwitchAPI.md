# SwitchAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPortTag**](SwitchAPI.md#addporttag) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/port-tag | Create new switch port label
[**BatchApplySwitchNetwork**](SwitchAPI.md#batchapplyswitchnetwork) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/networks | Batch modify switch vlan interfaces.
[**BatchModifyESLoopbackControl**](SwitchAPI.md#batchmodifyesloopbackcontrol) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/cmd/switches/es/config/loopback | Batch modify switch loopback control (Agile Series)
[**BatchModifyLoopbackControl**](SwitchAPI.md#batchmodifyloopbackcontrol) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/cmd/switches/config/loopback | Batch modify switch loopback control
[**BatchModifySwitchPort**](SwitchAPI.md#batchmodifyswitchport) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/multi-ports/config | Batch modify switch port
[**BatchModifySwitchesPortLag**](SwitchAPI.md#batchmodifyswitchesportlag) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/switches/multi-ports/config | Batch modify switches ports and lags
[**BatchPortPoERecovery**](SwitchAPI.md#batchportpoerecovery) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/multi-ports/poe-recovery | Switch port poe recovery.
[**BatchSetNameForGivenPorts1**](SwitchAPI.md#batchsetnameforgivenports1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/multi-ports/name | Batch set name for given ports
[**BatchSetPoeModeForGivenPorts1**](SwitchAPI.md#batchsetpoemodeforgivenports1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/multi-ports/poe-mode | Batch set poe mode for given ports
[**BatchSetPortStatusForGivenPorts1**](SwitchAPI.md#batchsetportstatusforgivenports1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/multi-ports/status | Batch set status for given ports
[**BatchSetProfileOverrideForGivenPorts1**](SwitchAPI.md#batchsetprofileoverrideforgivenports1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/multi-ports/profile-override | Batch set profile-override for given ports
[**CancelCableTest**](SwitchAPI.md#cancelcabletest) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/switches/{switchMac}/cancel | Cancel cable test
[**ClearOswPortCounters**](SwitchAPI.md#clearoswportcounters) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/multi-ports/clear-counters | Clear the switches ports counters
[**CreateOswVrf**](SwitchAPI.md#createoswvrf) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/vrfs | Create new vrf
[**DeleteOswVrf**](SwitchAPI.md#deleteoswvrf) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/vrfs/{vrfId} | Delete vrf
[**DeletePortTag**](SwitchAPI.md#deleteporttag) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/switches/port-tag | Delete an existing switch port label
[**DeleteSwitchLag**](SwitchAPI.md#deleteswitchlag) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/lags/{lagId} | Delete switch lag
[**GetBatchSelectOswDetailsView**](SwitchAPI.md#getbatchselectoswdetailsview) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/ports/select | Get the switches and ports detail in batches
[**GetBatchSwitchExistNetworks**](SwitchAPI.md#getbatchswitchexistnetworks) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/exist-network | Get the networks intersections existing on multiple switches
[**GetCableTestFullResults**](SwitchAPI.md#getcabletestfullresults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/switches/{switchMac}/full-results | Get the cable test full results
[**GetCableTestIncrementResults**](SwitchAPI.md#getcabletestincrementresults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/switches/{switchMac}/increment-results | Get the cable test increment results
[**GetCableTestLogs**](SwitchAPI.md#getcabletestlogs) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/switches/{switchMac}/logs | Get the cable test logs
[**GetCableTestOswPorts**](SwitchAPI.md#getcabletestoswports) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/switches/{switchMac}/ports | Get the port list used for cable test
[**GetESGeneralConfig**](SwitchAPI.md#getesgeneralconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/es/{switchMac}/general-config | Get switch general config (Agile Series )
[**GetESInfo**](SwitchAPI.md#getesinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/es/{switchMac} | Get switch info (Agile Series)
[**GetESNetworkOverview**](SwitchAPI.md#getesnetworkoverview) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/es/{switchMac}/network-overview | Query Agile Series switch valid network
[**GetGeneralConfig**](SwitchAPI.md#getgeneralconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/general-config | Get switch general config
[**GetGridDhcpUserList**](SwitchAPI.md#getgriddhcpuserlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/dhcp/user-list | Get switch dhcp user list
[**GetGridLldpNeighborTable**](SwitchAPI.md#getgridlldpneighbortable) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/lldp-neighbors | Get switch lldp neighbor table
[**GetGridOspfNeighborTable**](SwitchAPI.md#getgridospfneighbortable) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/ospf-neighbors | Get switch ospf neighbor table
[**GetGridPortAndLagNetworks1**](SwitchAPI.md#getgridportandlagnetworks1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/port-lag-networks | Get the networks used on switch&#39;s ports and LAGs
[**GetGridPortAndLagNetworksDetail**](SwitchAPI.md#getgridportandlagnetworksdetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/networks-detail | Get the networks detail (including the vlan related ports and lags info) used on switch&#39;s ports and LAGs
[**GetGridSwitchesPortsCounters**](SwitchAPI.md#getgridswitchesportscounters) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/ports/counters | Get the switches ports counters information
[**GetGridSwitchesPortsOverview**](SwitchAPI.md#getgridswitchesportsoverview) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/ports/overview | Get the switches ports overview
[**GetGridSwitchesPortsPoe**](SwitchAPI.md#getgridswitchesportspoe) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/ports/poe-info | Get the switches ports poe information
[**GetGridVrf**](SwitchAPI.md#getgridvrf) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/vrfs | Get the vrfs on the switch
[**GetOswDDMInfo**](SwitchAPI.md#getoswddminfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/ddm/info | Get osw ddm info.
[**GetOswDetailInfoListForPortsView**](SwitchAPI.md#getoswdetailinfolistforportsview) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/ports/switch-detail | Get the switches detail for ports view
[**GetOswForGivenLanNetworkIdAndVlan**](SwitchAPI.md#getoswforgivenlannetworkidandvlan) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/vlans/{vlan}/switches | Get the paging query for the osws with given network id and vlan
[**GetOswsDetails**](SwitchAPI.md#getoswsdetails) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/details | Get the details of osws with given omadacid, siteid and macs and stackIds.
[**GetOswsSupportArpDetect**](SwitchAPI.md#getoswssupportarpdetect) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/supportArpDetect | Get the paging query for the osws that support arp detect.
[**GetOswsSupportDhcpSnoop**](SwitchAPI.md#getoswssupportdhcpsnoop) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/supportDhcpSnoop | Get the paging query for the osws that support dhcp snoop.
[**GetPortAndLagNetwork1**](SwitchAPI.md#getportandlagnetwork1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/port-lag-networks/{networkId}/vlan/{vlan} | Get the switch&#39;s ports and LAGs that the network affects
[**GetPortTags**](SwitchAPI.md#getporttags) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/port-tag | Get switch port label list
[**GetSwitchExistNetworks**](SwitchAPI.md#getswitchexistnetworks) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/exist-network | Get the networks existing on the switch
[**GetSwitchInfo**](SwitchAPI.md#getswitchinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac} | Get switch info
[**GetSwitchUsedSdmNum**](SwitchAPI.md#getswitchusedsdmnum) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/sdm-used | Get the used sdm template num on the switch
[**ListSwitchNetworks**](SwitchAPI.md#listswitchnetworks) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/networks | Query switch vlan interface infos.
[**ModifyESGeneralConfig**](SwitchAPI.md#modifyesgeneralconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/switches/es/{switchMac}/general-config | Modify switch general config (Agile Series)
[**ModifyESLoopbackControl**](SwitchAPI.md#modifyesloopbackcontrol) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/es/{switchMac}/config/loopback | Modify switch loopback control (Agile Series)
[**ModifyGeneralConfig**](SwitchAPI.md#modifygeneralconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/general-config | Modify switch general config
[**ModifyLoopbackControl**](SwitchAPI.md#modifyloopbackcontrol) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/config/loopback | Modify switch loopback control
[**ModifyOswVrf**](SwitchAPI.md#modifyoswvrf) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/vrfs/{vrfId} | Modify vrf
[**ModifyPortTag**](SwitchAPI.md#modifyporttag) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/switches/port-tag | Modify an existing switch port label
[**ModifySwitchLag**](SwitchAPI.md#modifyswitchlag) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/lags/{lagId} | Modify switch lag
[**ModifySwitchNetwork**](SwitchAPI.md#modifyswitchnetwork) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/networks/{networkId} | Modify switch network config.
[**ModifySwitchPort**](SwitchAPI.md#modifyswitchport) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/ports/{port} | Modify switch port
[**MspMoveToCustomer**](SwitchAPI.md#mspmovetocustomer) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/cmd/switches/move | Move site to target customer
[**SetNameForGivenPort1**](SwitchAPI.md#setnameforgivenport1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/ports/{port}/name | Set name for given port
[**SetPoeModeForGivenPort1**](SwitchAPI.md#setpoemodeforgivenport1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/ports/{port}/poe-mode | Set poe mode for given port
[**SetPortModeForGivenPort1**](SwitchAPI.md#setportmodeforgivenport1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/ports/{port}/status | Set port status for given port
[**SetProfileForGivenPort1**](SwitchAPI.md#setprofileforgivenport1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/ports/{port}/profile | Set profile for given port
[**SetProfileOverrideForGivenPort1**](SwitchAPI.md#setprofileoverrideforgivenport1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/ports/{port}/profile-override | Set profile-override for given port
[**StartCableTest**](SwitchAPI.md#startcabletest) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/switches/{switchMac}/start | Start cable test



## AddPortTag

> PortTagOpenApiVO AddPortTag(ctx, omadacId, siteId).CreatePortTagOpenApiVO(createPortTagOpenApiVO).Execute()

Create new switch port label



### Example

```go
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
	createPortTagOpenApiVO := *openapiclient.NewCreatePortTagOpenApiVO("Name_example") // CreatePortTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.AddPortTag(context.Background(), omadacId, siteId).CreatePortTagOpenApiVO(createPortTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.AddPortTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPortTag`: PortTagOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.AddPortTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPortTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createPortTagOpenApiVO** | [**CreatePortTagOpenApiVO**](CreatePortTagOpenApiVO.md) |  | 

### Return type

[**PortTagOpenApiVO**](PortTagOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchApplySwitchNetwork

> OperationResponseWithoutResult BatchApplySwitchNetwork(ctx, omadacId, siteId, switchMac).BatchApplyOswNetworkOpenApi(batchApplyOswNetworkOpenApi).Execute()

Batch modify switch vlan interfaces.



### Example

```go
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
	batchApplyOswNetworkOpenApi := *openapiclient.NewBatchApplyOswNetworkOpenApi([]openapiclient.OswNetworkBaseOpenApi{*openapiclient.NewOswNetworkBaseOpenApi("Id_example", int32(123))}) // BatchApplyOswNetworkOpenApi | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchApplySwitchNetwork(context.Background(), omadacId, siteId, switchMac).BatchApplyOswNetworkOpenApi(batchApplyOswNetworkOpenApi).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchApplySwitchNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchApplySwitchNetwork`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchApplySwitchNetwork`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchApplySwitchNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchApplyOswNetworkOpenApi** | [**BatchApplyOswNetworkOpenApi**](BatchApplyOswNetworkOpenApi.md) |  | 

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


## BatchModifyESLoopbackControl

> OperationResponseWithoutResult BatchModifyESLoopbackControl(ctx, omadacId, siteId).EasyManagedSwitchBatchLoopbackControl(easyManagedSwitchBatchLoopbackControl).Execute()

Batch modify switch loopback control (Agile Series)



### Example

```go
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
	easyManagedSwitchBatchLoopbackControl := *openapiclient.NewEasyManagedSwitchBatchLoopbackControl([]string{"SwitchMacList_example"}) // EasyManagedSwitchBatchLoopbackControl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchModifyESLoopbackControl(context.Background(), omadacId, siteId).EasyManagedSwitchBatchLoopbackControl(easyManagedSwitchBatchLoopbackControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchModifyESLoopbackControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyESLoopbackControl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchModifyESLoopbackControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyESLoopbackControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **easyManagedSwitchBatchLoopbackControl** | [**EasyManagedSwitchBatchLoopbackControl**](EasyManagedSwitchBatchLoopbackControl.md) |  | 

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


## BatchModifyLoopbackControl

> OperationResponseWithoutResult BatchModifyLoopbackControl(ctx, omadacId, siteId).SwitchBatchLoopbackControl(switchBatchLoopbackControl).Execute()

Batch modify switch loopback control



### Example

```go
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
	switchBatchLoopbackControl := *openapiclient.NewSwitchBatchLoopbackControl([]string{"SwitchMacList_example"}) // SwitchBatchLoopbackControl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchModifyLoopbackControl(context.Background(), omadacId, siteId).SwitchBatchLoopbackControl(switchBatchLoopbackControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchModifyLoopbackControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyLoopbackControl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchModifyLoopbackControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyLoopbackControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **switchBatchLoopbackControl** | [**SwitchBatchLoopbackControl**](SwitchBatchLoopbackControl.md) |  | 

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


## BatchModifySwitchPort

> OperationResponseString BatchModifySwitchPort(ctx, omadacId, siteId, switchMac).BatchOswPortSettingVO(batchOswPortSettingVO).Execute()

Batch modify switch port



### Example

```go
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
	batchOswPortSettingVO := *openapiclient.NewBatchOswPortSettingVO([]int32{int32(123)}) // BatchOswPortSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchModifySwitchPort(context.Background(), omadacId, siteId, switchMac).BatchOswPortSettingVO(batchOswPortSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchModifySwitchPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifySwitchPort`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchModifySwitchPort`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchModifySwitchPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchOswPortSettingVO** | [**BatchOswPortSettingVO**](BatchOswPortSettingVO.md) |  | 

### Return type

[**OperationResponseString**](OperationResponseString.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchModifySwitchesPortLag

> OperationResponseString BatchModifySwitchesPortLag(ctx, omadacId, siteId).MultiOswPortSettingOpenApiVO(multiOswPortSettingOpenApiVO).Execute()

Batch modify switches ports and lags



### Example

```go
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
	multiOswPortSettingOpenApiVO := *openapiclient.NewMultiOswPortSettingOpenApiVO(false, []openapiclient.OswPortLagListVO{*openapiclient.NewOswPortLagListVO("Mac_example")}) // MultiOswPortSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchModifySwitchesPortLag(context.Background(), omadacId, siteId).MultiOswPortSettingOpenApiVO(multiOswPortSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchModifySwitchesPortLag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifySwitchesPortLag`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchModifySwitchesPortLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifySwitchesPortLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **multiOswPortSettingOpenApiVO** | [**MultiOswPortSettingOpenApiVO**](MultiOswPortSettingOpenApiVO.md) |  | 

### Return type

[**OperationResponseString**](OperationResponseString.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchPortPoERecovery

> OperationResponseOswPoeResultOpenApiVO BatchPortPoERecovery(ctx, omadacId, siteId, switchMac).OswPoeRecoverOpenApiVO(oswPoeRecoverOpenApiVO).Execute()

Switch port poe recovery.



### Example

```go
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
	oswPoeRecoverOpenApiVO := *openapiclient.NewOswPoeRecoverOpenApiVO([]int32{int32(123)}) // OswPoeRecoverOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchPortPoERecovery(context.Background(), omadacId, siteId, switchMac).OswPoeRecoverOpenApiVO(oswPoeRecoverOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchPortPoERecovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchPortPoERecovery`: OperationResponseOswPoeResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchPortPoERecovery`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchPortPoERecoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswPoeRecoverOpenApiVO** | [**OswPoeRecoverOpenApiVO**](OswPoeRecoverOpenApiVO.md) |  | 

### Return type

[**OperationResponseOswPoeResultOpenApiVO**](OperationResponseOswPoeResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchSetNameForGivenPorts1

> OperationResponseWithoutResult BatchSetNameForGivenPorts1(ctx, omadacId, siteId, switchMac).PortNameList(portNameList).Execute()

Batch set name for given ports



### Example

```go
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
	portNameList := *openapiclient.NewPortNameList([]openapiclient.SwitchMultiPortName{*openapiclient.NewSwitchMultiPortName("Name_example", int32(123))}) // PortNameList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchSetNameForGivenPorts1(context.Background(), omadacId, siteId, switchMac).PortNameList(portNameList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchSetNameForGivenPorts1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetNameForGivenPorts1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchSetNameForGivenPorts1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchSetNameForGivenPorts1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **portNameList** | [**PortNameList**](PortNameList.md) |  | 

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


## BatchSetPoeModeForGivenPorts1

> OperationResponseWithoutResult BatchSetPoeModeForGivenPorts1(ctx, omadacId, siteId, switchMac).SwitchPortsPoe(switchPortsPoe).Execute()

Batch set poe mode for given ports



### Example

```go
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
	switchPortsPoe := *openapiclient.NewSwitchPortsPoe(int32(123), []int32{int32(123)}) // SwitchPortsPoe | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchSetPoeModeForGivenPorts1(context.Background(), omadacId, siteId, switchMac).SwitchPortsPoe(switchPortsPoe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchSetPoeModeForGivenPorts1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetPoeModeForGivenPorts1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchSetPoeModeForGivenPorts1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchSetPoeModeForGivenPorts1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **switchPortsPoe** | [**SwitchPortsPoe**](SwitchPortsPoe.md) |  | 

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


## BatchSetPortStatusForGivenPorts1

> OperationResponseWithoutResult BatchSetPortStatusForGivenPorts1(ctx, omadacId, siteId, switchMac).SwitchPortsStatus(switchPortsStatus).Execute()

Batch set status for given ports



### Example

```go
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
	switchPortsStatus := *openapiclient.NewSwitchPortsStatus([]int32{int32(123)}, int32(123)) // SwitchPortsStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchSetPortStatusForGivenPorts1(context.Background(), omadacId, siteId, switchMac).SwitchPortsStatus(switchPortsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchSetPortStatusForGivenPorts1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetPortStatusForGivenPorts1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchSetPortStatusForGivenPorts1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchSetPortStatusForGivenPorts1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **switchPortsStatus** | [**SwitchPortsStatus**](SwitchPortsStatus.md) |  | 

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


## BatchSetProfileOverrideForGivenPorts1

> OperationResponseWithoutResult BatchSetProfileOverrideForGivenPorts1(ctx, omadacId, siteId, switchMac).BatchProfileOverride(batchProfileOverride).Execute()

Batch set profile-override for given ports



### Example

```go
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
	batchProfileOverride := *openapiclient.NewBatchProfileOverride([]int32{int32(123)}, false) // BatchProfileOverride | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.BatchSetProfileOverrideForGivenPorts1(context.Background(), omadacId, siteId, switchMac).BatchProfileOverride(batchProfileOverride).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.BatchSetProfileOverrideForGivenPorts1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetProfileOverrideForGivenPorts1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.BatchSetProfileOverrideForGivenPorts1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchSetProfileOverrideForGivenPorts1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchProfileOverride** | [**BatchProfileOverride**](BatchProfileOverride.md) |  | 

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


## CancelCableTest

> OperationResponseWithoutResult CancelCableTest(ctx, omadacId, siteId, switchMac).Execute()

Cancel cable test



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CancelCableTest(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CancelCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelCableTest`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CancelCableTest`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCancelCableTestRequest struct via the builder pattern


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


## ClearOswPortCounters

> OperationResponseString ClearOswPortCounters(ctx, omadacId, siteId).ClearCountersVO(clearCountersVO).Execute()

Clear the switches ports counters



### Example

```go
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
	clearCountersVO := *openapiclient.NewClearCountersVO([]openapiclient.PortParamVO{*openapiclient.NewPortParamVO(int32(123), "SwitchMac_example")}) // ClearCountersVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ClearOswPortCounters(context.Background(), omadacId, siteId).ClearCountersVO(clearCountersVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ClearOswPortCounters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearOswPortCounters`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ClearOswPortCounters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearOswPortCountersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clearCountersVO** | [**ClearCountersVO**](ClearCountersVO.md) |  | 

### Return type

[**OperationResponseString**](OperationResponseString.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOswVrf

> OperationResponseWithoutResult CreateOswVrf(ctx, omadacId, siteId, switchMac).OswVrfConfigOpenApiVO(oswVrfConfigOpenApiVO).Execute()

Create new vrf



### Example

```go
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
	oswVrfConfigOpenApiVO := *openapiclient.NewOswVrfConfigOpenApiVO(false, false, "Vrf_example") // OswVrfConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateOswVrf(context.Background(), omadacId, siteId, switchMac).OswVrfConfigOpenApiVO(oswVrfConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateOswVrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOswVrf`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateOswVrf`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateOswVrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswVrfConfigOpenApiVO** | [**OswVrfConfigOpenApiVO**](OswVrfConfigOpenApiVO.md) |  | 

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


## DeleteOswVrf

> OperationResponseWithoutResult DeleteOswVrf(ctx, omadacId, siteId, switchMac, vrfId).Execute()

Delete vrf



### Example

```go
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
	vrfId := "vrfId_example" // string | VRF ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.DeleteOswVrf(context.Background(), omadacId, siteId, switchMac, vrfId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteOswVrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOswVrf`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.DeleteOswVrf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**vrfId** | **string** | VRF ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOswVrfRequest struct via the builder pattern


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


## DeletePortTag

> OperationResponseWithoutResult DeletePortTag(ctx, omadacId, siteId).DeletePortTagOpenApiVO(deletePortTagOpenApiVO).Execute()

Delete an existing switch port label



### Example

```go
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
	deletePortTagOpenApiVO := *openapiclient.NewDeletePortTagOpenApiVO("TagId_example") // DeletePortTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.DeletePortTag(context.Background(), omadacId, siteId).DeletePortTagOpenApiVO(deletePortTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeletePortTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortTag`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.DeletePortTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deletePortTagOpenApiVO** | [**DeletePortTagOpenApiVO**](DeletePortTagOpenApiVO.md) |  | 

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


## DeleteSwitchLag

> OperationResponseString DeleteSwitchLag(ctx, omadacId, siteId, switchMac, lagId).Execute()

Delete switch lag



### Example

```go
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
	lagId := "lagId_example" // string | lagId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.DeleteSwitchLag(context.Background(), omadacId, siteId, switchMac, lagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteSwitchLag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSwitchLag`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.DeleteSwitchLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**lagId** | **string** | lagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSwitchLagRequest struct via the builder pattern


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


## GetBatchSelectOswDetailsView

> OperationResponseListOswDetailBatchSelectedVO GetBatchSelectOswDetailsView(ctx, omadacId, siteId).MultiOswPortSelectVO(multiOswPortSelectVO).Execute()

Get the switches and ports detail in batches



### Example

```go
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
	multiOswPortSelectVO := *openapiclient.NewMultiOswPortSelectVO(false, []openapiclient.OswPortLagListVO{*openapiclient.NewOswPortLagListVO("Mac_example")}) // MultiOswPortSelectVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetBatchSelectOswDetailsView(context.Background(), omadacId, siteId).MultiOswPortSelectVO(multiOswPortSelectVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetBatchSelectOswDetailsView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchSelectOswDetailsView`: OperationResponseListOswDetailBatchSelectedVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetBatchSelectOswDetailsView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchSelectOswDetailsViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **multiOswPortSelectVO** | [**MultiOswPortSelectVO**](MultiOswPortSelectVO.md) |  | 

### Return type

[**OperationResponseListOswDetailBatchSelectedVO**](OperationResponseListOswDetailBatchSelectedVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchSwitchExistNetworks

> OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO GetBatchSwitchExistNetworks(ctx, omadacId, siteId).OswMacListVO(oswMacListVO).Execute()

Get the networks intersections existing on multiple switches



### Example

```go
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
	oswMacListVO := *openapiclient.NewOswMacListVO([]string{"SwitchMacList_example"}) // OswMacListVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetBatchSwitchExistNetworks(context.Background(), omadacId, siteId).OswMacListVO(oswMacListVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetBatchSwitchExistNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchSwitchExistNetworks`: OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetBatchSwitchExistNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchSwitchExistNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oswMacListVO** | [**OswMacListVO**](OswMacListVO.md) |  | 

### Return type

[**OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO**](OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCableTestFullResults

> OperationResponseOswCableTestResultWithStatusVO GetCableTestFullResults(ctx, omadacId, siteId, switchMac).Execute()

Get the cable test full results



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetCableTestFullResults(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetCableTestFullResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCableTestFullResults`: OperationResponseOswCableTestResultWithStatusVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetCableTestFullResults`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCableTestFullResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOswCableTestResultWithStatusVO**](OperationResponseOswCableTestResultWithStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCableTestIncrementResults

> OperationResponseOswCableTestResultWithStatusVO GetCableTestIncrementResults(ctx, omadacId, siteId, switchMac).Execute()

Get the cable test increment results



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetCableTestIncrementResults(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetCableTestIncrementResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCableTestIncrementResults`: OperationResponseOswCableTestResultWithStatusVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetCableTestIncrementResults`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCableTestIncrementResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOswCableTestResultWithStatusVO**](OperationResponseOswCableTestResultWithStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCableTestLogs

> OperationResponseListOswCableTestLogOpenApiVO GetCableTestLogs(ctx, omadacId, siteId, switchMac).Execute()

Get the cable test logs



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetCableTestLogs(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetCableTestLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCableTestLogs`: OperationResponseListOswCableTestLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetCableTestLogs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCableTestLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListOswCableTestLogOpenApiVO**](OperationResponseListOswCableTestLogOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCableTestOswPorts

> OperationResponseListOswCableTestPortVO GetCableTestOswPorts(ctx, omadacId, siteId, switchMac).Execute()

Get the port list used for cable test



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetCableTestOswPorts(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetCableTestOswPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCableTestOswPorts`: OperationResponseListOswCableTestPortVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetCableTestOswPorts`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCableTestOswPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListOswCableTestPortVO**](OperationResponseListOswCableTestPortVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetESGeneralConfig

> OperationResponseEasyManagedSwitchGeneralConfigDetail GetESGeneralConfig(ctx, omadacId, siteId, switchMac).Execute()

Get switch general config (Agile Series )



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetESGeneralConfig(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetESGeneralConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetESGeneralConfig`: OperationResponseEasyManagedSwitchGeneralConfigDetail
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetESGeneralConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetESGeneralConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseEasyManagedSwitchGeneralConfigDetail**](OperationResponseEasyManagedSwitchGeneralConfigDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetESInfo

> OperationResponseEasyManageOverviewInfo GetESInfo(ctx, omadacId, siteId, switchMac).Execute()

Get switch info (Agile Series)



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetESInfo(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetESInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetESInfo`: OperationResponseEasyManageOverviewInfo
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetESInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetESInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseEasyManageOverviewInfo**](OperationResponseEasyManageOverviewInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetESNetworkOverview

> OperationResponseListNetworkPortsAssociationVO GetESNetworkOverview(ctx, omadacId, siteId, switchMac).Execute()

Query Agile Series switch valid network



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetESNetworkOverview(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetESNetworkOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetESNetworkOverview`: OperationResponseListNetworkPortsAssociationVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetESNetworkOverview`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetESNetworkOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListNetworkPortsAssociationVO**](OperationResponseListNetworkPortsAssociationVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneralConfig

> OperationResponseSwitchGeneralConfig GetGeneralConfig(ctx, omadacId, siteId, switchMac).Execute()

Get switch general config



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetGeneralConfig(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGeneralConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGeneralConfig`: OperationResponseSwitchGeneralConfig
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGeneralConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGeneralConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSwitchGeneralConfig**](OperationResponseSwitchGeneralConfig.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDhcpUserList

> OperationResponseDhcpUserGridVODhcpUserVO GetGridDhcpUserList(ctx, omadacId, siteId, switchMac).Page(page).PageSize(pageSize).Execute()

Get switch dhcp user list



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetGridDhcpUserList(context.Background(), omadacId, siteId, switchMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridDhcpUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDhcpUserList`: OperationResponseDhcpUserGridVODhcpUserVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridDhcpUserList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridDhcpUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseDhcpUserGridVODhcpUserVO**](OperationResponseDhcpUserGridVODhcpUserVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridLldpNeighborTable

> OperationResponseGridVOOswLldpNeighborVO GetGridLldpNeighborTable(ctx, omadacId, siteId, switchMac).Page(page).PageSize(pageSize).SortsPortId(sortsPortId).SearchKey(searchKey).Execute()

Get switch lldp neighbor table



### Example

```go
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
	sortsPortId := "sortsPortId_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field portId,standardPort,deviceId,systemName (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetGridLldpNeighborTable(context.Background(), omadacId, siteId, switchMac).Page(page).PageSize(pageSize).SortsPortId(sortsPortId).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridLldpNeighborTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridLldpNeighborTable`: OperationResponseGridVOOswLldpNeighborVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridLldpNeighborTable`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridLldpNeighborTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsPortId** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field portId,standardPort,deviceId,systemName | 

### Return type

[**OperationResponseGridVOOswLldpNeighborVO**](OperationResponseGridVOOswLldpNeighborVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridOspfNeighborTable

> OperationResponseOswOspfNeighborGridVOOswOspfNeighborVO GetGridOspfNeighborTable(ctx, omadacId, siteId, switchMac).Page(page).PageSize(pageSize).SortsNeighborInterface(sortsNeighborInterface).FiltersNeighborInterface(filtersNeighborInterface).FiltersProcessId(filtersProcessId).SearchKey(searchKey).Execute()

Get switch ospf neighbor table



### Example

```go
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
	sortsNeighborInterface := "sortsNeighborInterface_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersNeighborInterface := []string{"Inner_example"} // []string | Filter query parameters, support field neighborInterface (optional)
	filtersProcessId := "filtersProcessId_example" // string | Filter query parameters, support field processId (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field neighborIp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetGridOspfNeighborTable(context.Background(), omadacId, siteId, switchMac).Page(page).PageSize(pageSize).SortsNeighborInterface(sortsNeighborInterface).FiltersNeighborInterface(filtersNeighborInterface).FiltersProcessId(filtersProcessId).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridOspfNeighborTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridOspfNeighborTable`: OperationResponseOswOspfNeighborGridVOOswOspfNeighborVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridOspfNeighborTable`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridOspfNeighborTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsNeighborInterface** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersNeighborInterface** | **[]string** | Filter query parameters, support field neighborInterface | 
 **filtersProcessId** | **string** | Filter query parameters, support field processId | 
 **searchKey** | **string** | Fuzzy query parameters, support field neighborIp | 

### Return type

[**OperationResponseOswOspfNeighborGridVOOswOspfNeighborVO**](OperationResponseOswOspfNeighborGridVOOswOspfNeighborVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridPortAndLagNetworks1

> GridVOOswNetworkBriefInfoOpenApiVO GetGridPortAndLagNetworks1(ctx, omadacId, siteId, switchMac).Page(page).PageSize(pageSize).Execute()

Get the networks used on switch's ports and LAGs



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetGridPortAndLagNetworks1(context.Background(), omadacId, siteId, switchMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridPortAndLagNetworks1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPortAndLagNetworks1`: GridVOOswNetworkBriefInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridPortAndLagNetworks1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridPortAndLagNetworks1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOOswNetworkBriefInfoOpenApiVO**](GridVOOswNetworkBriefInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridPortAndLagNetworksDetail

> GridVOOswNetworkDetailInfoOpenApiVO GetGridPortAndLagNetworksDetail(ctx, omadacId, siteId, switchMac).Page(page).PageSize(pageSize).Execute()

Get the networks detail (including the vlan related ports and lags info) used on switch's ports and LAGs



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetGridPortAndLagNetworksDetail(context.Background(), omadacId, siteId, switchMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridPortAndLagNetworksDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPortAndLagNetworksDetail`: GridVOOswNetworkDetailInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridPortAndLagNetworksDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridPortAndLagNetworksDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOOswNetworkDetailInfoOpenApiVO**](GridVOOswNetworkDetailInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSwitchesPortsCounters

> OperationResponseGridVOOswPortsSettingCountersVO GetGridSwitchesPortsCounters(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the switches ports counters information



### Example

```go
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
	resp, r, err := apiClient.SwitchAPI.GetGridSwitchesPortsCounters(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridSwitchesPortsCounters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSwitchesPortsCounters`: OperationResponseGridVOOswPortsSettingCountersVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridSwitchesPortsCounters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSwitchesPortsCountersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOswPortsSettingCountersVO**](OperationResponseGridVOOswPortsSettingCountersVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSwitchesPortsOverview

> OperationResponseGridVOOswPortsSettingOverviewVO GetGridSwitchesPortsOverview(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the switches ports overview



### Example

```go
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
	resp, r, err := apiClient.SwitchAPI.GetGridSwitchesPortsOverview(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridSwitchesPortsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSwitchesPortsOverview`: OperationResponseGridVOOswPortsSettingOverviewVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridSwitchesPortsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSwitchesPortsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOswPortsSettingOverviewVO**](OperationResponseGridVOOswPortsSettingOverviewVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSwitchesPortsPoe

> OperationResponseGridVOOswPortsSettingPoeVO GetGridSwitchesPortsPoe(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the switches ports poe information



### Example

```go
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
	resp, r, err := apiClient.SwitchAPI.GetGridSwitchesPortsPoe(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridSwitchesPortsPoe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSwitchesPortsPoe`: OperationResponseGridVOOswPortsSettingPoeVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridSwitchesPortsPoe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSwitchesPortsPoeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOswPortsSettingPoeVO**](OperationResponseGridVOOswPortsSettingPoeVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVrf

> OperationResponseGridVOOswVrfOpenApiVO GetGridVrf(ctx, omadacId, siteId, switchMac).Execute()

Get the vrfs on the switch



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetGridVrf(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetGridVrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVrf`: OperationResponseGridVOOswVrfOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetGridVrf`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridVrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseGridVOOswVrfOpenApiVO**](OperationResponseGridVOOswVrfOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswDDMInfo

> OperationResponseWithoutResult GetOswDDMInfo(ctx, omadacId, siteId, switchMac).Execute()

Get osw ddm info.



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOswDDMInfo(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOswDDMInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswDDMInfo`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOswDDMInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOswDDMInfoRequest struct via the builder pattern


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


## GetOswDetailInfoListForPortsView

> OperationResponseListOswDetailForPortsViewVO GetOswDetailInfoListForPortsView(ctx, omadacId, siteId).Execute()

Get the switches detail for ports view



### Example

```go
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
	resp, r, err := apiClient.SwitchAPI.GetOswDetailInfoListForPortsView(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOswDetailInfoListForPortsView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswDetailInfoListForPortsView`: OperationResponseListOswDetailForPortsViewVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOswDetailInfoListForPortsView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswDetailInfoListForPortsViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListOswDetailForPortsViewVO**](OperationResponseListOswDetailForPortsViewVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswForGivenLanNetworkIdAndVlan

> OperationResponse GetOswForGivenLanNetworkIdAndVlan(ctx, omadacId, siteId, networkId, vlan).Page(page).PageSize(pageSize).Execute()

Get the paging query for the osws with given network id and vlan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | networkId
	vlan := "vlan_example" // string | vlan

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOswForGivenLanNetworkIdAndVlan(context.Background(), omadacId, siteId, networkId, vlan).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOswForGivenLanNetworkIdAndVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswForGivenLanNetworkIdAndVlan`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOswForGivenLanNetworkIdAndVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | networkId | 
**vlan** | **string** | vlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswForGivenLanNetworkIdAndVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





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


## GetOswsDetails

> OperationResponse GetOswsDetails(ctx, omadacId, siteId).OswDetailOpenApiVO(oswDetailOpenApiVO).Execute()

Get the details of osws with given omadacid, siteid and macs and stackIds.



### Example

```go
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
	oswDetailOpenApiVO := *openapiclient.NewOswDetailOpenApiVO("NetworkId_example", int32(123)) // OswDetailOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOswsDetails(context.Background(), omadacId, siteId).OswDetailOpenApiVO(oswDetailOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOswsDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswsDetails`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOswsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oswDetailOpenApiVO** | [**OswDetailOpenApiVO**](OswDetailOpenApiVO.md) |  | 

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


## GetOswsSupportArpDetect

> OperationResponse GetOswsSupportArpDetect(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the paging query for the osws that support arp detect.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOswsSupportArpDetect(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOswsSupportArpDetect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswsSupportArpDetect`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOswsSupportArpDetect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswsSupportArpDetectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



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


## GetOswsSupportDhcpSnoop

> OperationResponse GetOswsSupportDhcpSnoop(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the paging query for the osws that support dhcp snoop.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOswsSupportDhcpSnoop(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOswsSupportDhcpSnoop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswsSupportDhcpSnoop`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOswsSupportDhcpSnoop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswsSupportDhcpSnoopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



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


## GetPortAndLagNetwork1

> OswPortAndLagNetworkVO GetPortAndLagNetwork1(ctx, omadacId, siteId, switchMac, networkId, vlan).Execute()

Get the switch's ports and LAGs that the network affects



### Example

```go
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
	networkId := "networkId_example" // string | Switch network ID.
	vlan := "vlan_example" // string | VLAN.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetPortAndLagNetwork1(context.Background(), omadacId, siteId, switchMac, networkId, vlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetPortAndLagNetwork1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortAndLagNetwork1`: OswPortAndLagNetworkVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetPortAndLagNetwork1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**networkId** | **string** | Switch network ID. | 
**vlan** | **string** | VLAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortAndLagNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**OswPortAndLagNetworkVO**](OswPortAndLagNetworkVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortTags

> []PortTagOpenApiVO GetPortTags(ctx, omadacId, siteId).Execute()

Get switch port label list



### Example

```go
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
	resp, r, err := apiClient.SwitchAPI.GetPortTags(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetPortTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortTags`: []PortTagOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetPortTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PortTagOpenApiVO**](PortTagOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchExistNetworks

> OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO GetSwitchExistNetworks(ctx, omadacId, siteId, switchMac).Execute()

Get the networks existing on the switch



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetSwitchExistNetworks(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetSwitchExistNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchExistNetworks`: OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetSwitchExistNetworks`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSwitchExistNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO**](OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchInfo

> OperationResponseSwitchOverviewInfo GetSwitchInfo(ctx, omadacId, siteId, switchMac).Execute()

Get switch info



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetSwitchInfo(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetSwitchInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchInfo`: OperationResponseSwitchOverviewInfo
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetSwitchInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSwitchInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSwitchOverviewInfo**](OperationResponseSwitchOverviewInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchUsedSdmNum

> OperationResponseOswSdmApplicationVO GetSwitchUsedSdmNum(ctx, omadacId, siteId, switchMac).Execute()

Get the used sdm template num on the switch



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetSwitchUsedSdmNum(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetSwitchUsedSdmNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchUsedSdmNum`: OperationResponseOswSdmApplicationVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetSwitchUsedSdmNum`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSwitchUsedSdmNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOswSdmApplicationVO**](OperationResponseOswSdmApplicationVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSwitchNetworks

> OperationResponseOswNetworkGridOswNetworkOpenApi ListSwitchNetworks(ctx, omadacId, siteId, switchMac).Page(page).PageSize(pageSize).Execute()

Query switch vlan interface infos.



### Example

```go
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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ListSwitchNetworks(context.Background(), omadacId, siteId, switchMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ListSwitchNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSwitchNetworks`: OperationResponseOswNetworkGridOswNetworkOpenApi
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ListSwitchNetworks`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListSwitchNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseOswNetworkGridOswNetworkOpenApi**](OperationResponseOswNetworkGridOswNetworkOpenApi.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyESGeneralConfig

> OperationResponseWithoutResult ModifyESGeneralConfig(ctx, omadacId, siteId, switchMac).EasyManagedSwitchGeneralConfig(easyManagedSwitchGeneralConfig).Execute()

Modify switch general config (Agile Series)



### Example

```go
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
	easyManagedSwitchGeneralConfig := *openapiclient.NewEasyManagedSwitchGeneralConfig() // EasyManagedSwitchGeneralConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifyESGeneralConfig(context.Background(), omadacId, siteId, switchMac).EasyManagedSwitchGeneralConfig(easyManagedSwitchGeneralConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifyESGeneralConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyESGeneralConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifyESGeneralConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyESGeneralConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **easyManagedSwitchGeneralConfig** | [**EasyManagedSwitchGeneralConfig**](EasyManagedSwitchGeneralConfig.md) |  | 

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


## ModifyESLoopbackControl

> OperationResponseWithoutResult ModifyESLoopbackControl(ctx, omadacId, siteId, switchMac).EasyManagedSwitchLoopbackControl(easyManagedSwitchLoopbackControl).Execute()

Modify switch loopback control (Agile Series)



### Example

```go
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
	easyManagedSwitchLoopbackControl := *openapiclient.NewEasyManagedSwitchLoopbackControl() // EasyManagedSwitchLoopbackControl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifyESLoopbackControl(context.Background(), omadacId, siteId, switchMac).EasyManagedSwitchLoopbackControl(easyManagedSwitchLoopbackControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifyESLoopbackControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyESLoopbackControl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifyESLoopbackControl`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyESLoopbackControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **easyManagedSwitchLoopbackControl** | [**EasyManagedSwitchLoopbackControl**](EasyManagedSwitchLoopbackControl.md) |  | 

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


## ModifyGeneralConfig

> OperationResponseWithoutResult ModifyGeneralConfig(ctx, omadacId, siteId, switchMac).SwitchGeneralConfig(switchGeneralConfig).Execute()

Modify switch general config



### Example

```go
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
	switchGeneralConfig := *openapiclient.NewSwitchGeneralConfig() // SwitchGeneralConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifyGeneralConfig(context.Background(), omadacId, siteId, switchMac).SwitchGeneralConfig(switchGeneralConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifyGeneralConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGeneralConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifyGeneralConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyGeneralConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **switchGeneralConfig** | [**SwitchGeneralConfig**](SwitchGeneralConfig.md) |  | 

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


## ModifyLoopbackControl

> OperationResponseWithoutResult ModifyLoopbackControl(ctx, omadacId, siteId, switchMac).SwitchLoopbackControl(switchLoopbackControl).Execute()

Modify switch loopback control



### Example

```go
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
	switchLoopbackControl := *openapiclient.NewSwitchLoopbackControl() // SwitchLoopbackControl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifyLoopbackControl(context.Background(), omadacId, siteId, switchMac).SwitchLoopbackControl(switchLoopbackControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifyLoopbackControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLoopbackControl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifyLoopbackControl`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyLoopbackControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **switchLoopbackControl** | [**SwitchLoopbackControl**](SwitchLoopbackControl.md) |  | 

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


## ModifyOswVrf

> OperationResponseWithoutResult ModifyOswVrf(ctx, omadacId, siteId, switchMac, vrfId).OswVrfConfigOpenApiVO(oswVrfConfigOpenApiVO).Execute()

Modify vrf



### Example

```go
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
	vrfId := "vrfId_example" // string | VRF ID
	oswVrfConfigOpenApiVO := *openapiclient.NewOswVrfConfigOpenApiVO(false, false, "Vrf_example") // OswVrfConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifyOswVrf(context.Background(), omadacId, siteId, switchMac, vrfId).OswVrfConfigOpenApiVO(oswVrfConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifyOswVrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswVrf`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifyOswVrf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**vrfId** | **string** | VRF ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswVrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswVrfConfigOpenApiVO** | [**OswVrfConfigOpenApiVO**](OswVrfConfigOpenApiVO.md) |  | 

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


## ModifyPortTag

> OperationResponseWithoutResult ModifyPortTag(ctx, omadacId, siteId).PortTagOpenApiVO(portTagOpenApiVO).Execute()

Modify an existing switch port label



### Example

```go
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
	portTagOpenApiVO := *openapiclient.NewPortTagOpenApiVO("Name_example", "TagId_example") // PortTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifyPortTag(context.Background(), omadacId, siteId).PortTagOpenApiVO(portTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifyPortTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortTag`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifyPortTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portTagOpenApiVO** | [**PortTagOpenApiVO**](PortTagOpenApiVO.md) |  | 

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


## ModifySwitchLag

> OperationResponseString ModifySwitchLag(ctx, omadacId, siteId, switchMac, lagId).OswLagSettingVO(oswLagSettingVO).Execute()

Modify switch lag



### Example

```go
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
	lagId := "lagId_example" // string | lagId
	oswLagSettingVO := *openapiclient.NewOswLagSettingVO() // OswLagSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifySwitchLag(context.Background(), omadacId, siteId, switchMac, lagId).OswLagSettingVO(oswLagSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifySwitchLag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchLag`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifySwitchLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**lagId** | **string** | lagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswLagSettingVO** | [**OswLagSettingVO**](OswLagSettingVO.md) |  | 

### Return type

[**OperationResponseString**](OperationResponseString.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySwitchNetwork

> OperationResponseWithoutResult ModifySwitchNetwork(ctx, omadacId, siteId, switchMac, networkId).OswNetworkOpenApi(oswNetworkOpenApi).Execute()

Modify switch network config.



### Example

```go
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
	networkId := "networkId_example" // string | Switch network ID.
	oswNetworkOpenApi := *openapiclient.NewOswNetworkOpenApi("Id_example", int32(123), false, int32(123)) // OswNetworkOpenApi | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifySwitchNetwork(context.Background(), omadacId, siteId, switchMac, networkId).OswNetworkOpenApi(oswNetworkOpenApi).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifySwitchNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchNetwork`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifySwitchNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**networkId** | **string** | Switch network ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswNetworkOpenApi** | [**OswNetworkOpenApi**](OswNetworkOpenApi.md) |  | 

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


## ModifySwitchPort

> OperationResponseString ModifySwitchPort(ctx, omadacId, siteId, switchMac, port).OswPortSettingVO(oswPortSettingVO).Execute()

Modify switch port



### Example

```go
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
	port := "port_example" // string | Port ID
	oswPortSettingVO := *openapiclient.NewOswPortSettingVO() // OswPortSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.ModifySwitchPort(context.Background(), omadacId, siteId, switchMac, port).OswPortSettingVO(oswPortSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.ModifySwitchPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchPort`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.ModifySwitchPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswPortSettingVO** | [**OswPortSettingVO**](OswPortSettingVO.md) |  | 

### Return type

[**OperationResponseString**](OperationResponseString.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MspMoveToCustomer

> OperationResponseWithoutResult MspMoveToCustomer(ctx, mspId, customerId, siteId).MoveToCustomerOpenApiVO(moveToCustomerOpenApiVO).Execute()

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
	resp, r, err := apiClient.SwitchAPI.MspMoveToCustomer(context.Background(), mspId, customerId, siteId).MoveToCustomerOpenApiVO(moveToCustomerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.MspMoveToCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MspMoveToCustomer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.MspMoveToCustomer`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiMspMoveToCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **moveToCustomerOpenApiVO** | [**MoveToCustomerOpenApiVO**](MoveToCustomerOpenApiVO.md) |  | 

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


## SetNameForGivenPort1

> OperationResponseWithoutResult SetNameForGivenPort1(ctx, omadacId, siteId, switchMac, port).SwitchPortName(switchPortName).Execute()

Set name for given port



### Example

```go
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
	port := "port_example" // string | Port ID
	switchPortName := *openapiclient.NewSwitchPortName("Name_example") // SwitchPortName | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.SetNameForGivenPort1(context.Background(), omadacId, siteId, switchMac, port).SwitchPortName(switchPortName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.SetNameForGivenPort1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNameForGivenPort1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.SetNameForGivenPort1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNameForGivenPort1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **switchPortName** | [**SwitchPortName**](SwitchPortName.md) |  | 

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


## SetPoeModeForGivenPort1

> OperationResponseWithoutResult SetPoeModeForGivenPort1(ctx, omadacId, siteId, switchMac, port).SwitchPortPoe(switchPortPoe).Execute()

Set poe mode for given port



### Example

```go
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
	port := "port_example" // string | port
	switchPortPoe := *openapiclient.NewSwitchPortPoe(int32(123)) // SwitchPortPoe | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.SetPoeModeForGivenPort1(context.Background(), omadacId, siteId, switchMac, port).SwitchPortPoe(switchPortPoe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.SetPoeModeForGivenPort1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPoeModeForGivenPort1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.SetPoeModeForGivenPort1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | port | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPoeModeForGivenPort1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **switchPortPoe** | [**SwitchPortPoe**](SwitchPortPoe.md) |  | 

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


## SetPortModeForGivenPort1

> OperationResponseWithoutResult SetPortModeForGivenPort1(ctx, omadacId, siteId, switchMac, port).SwitchPortStatus(switchPortStatus).Execute()

Set port status for given port



### Example

```go
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
	port := "port_example" // string | Port
	switchPortStatus := *openapiclient.NewSwitchPortStatus(int32(123)) // SwitchPortStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.SetPortModeForGivenPort1(context.Background(), omadacId, siteId, switchMac, port).SwitchPortStatus(switchPortStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.SetPortModeForGivenPort1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPortModeForGivenPort1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.SetPortModeForGivenPort1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | Port | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPortModeForGivenPort1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **switchPortStatus** | [**SwitchPortStatus**](SwitchPortStatus.md) |  | 

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


## SetProfileForGivenPort1

> OperationResponseWithoutResult SetProfileForGivenPort1(ctx, omadacId, siteId, switchMac, port).SwitchProfileID(switchProfileID).Execute()

Set profile for given port



### Example

```go
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
	port := "port_example" // string | Port ID
	switchProfileID := *openapiclient.NewSwitchProfileID("ProfileId_example") // SwitchProfileID | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.SetProfileForGivenPort1(context.Background(), omadacId, siteId, switchMac, port).SwitchProfileID(switchProfileID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.SetProfileForGivenPort1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProfileForGivenPort1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.SetProfileForGivenPort1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProfileForGivenPort1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **switchProfileID** | [**SwitchProfileID**](SwitchProfileID.md) |  | 

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


## SetProfileOverrideForGivenPort1

> OperationResponseWithoutResult SetProfileOverrideForGivenPort1(ctx, omadacId, siteId, switchMac, port).ProfileOverride(profileOverride).Execute()

Set profile-override for given port



### Example

```go
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
	port := "port_example" // string | Port ID
	profileOverride := *openapiclient.NewProfileOverride(false) // ProfileOverride | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.SetProfileOverrideForGivenPort1(context.Background(), omadacId, siteId, switchMac, port).ProfileOverride(profileOverride).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.SetProfileOverrideForGivenPort1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProfileOverrideForGivenPort1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.SetProfileOverrideForGivenPort1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProfileOverrideForGivenPort1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **profileOverride** | [**ProfileOverride**](ProfileOverride.md) |  | 

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


## StartCableTest

> OperationResponseWithoutResult StartCableTest(ctx, omadacId, siteId, switchMac).OswCableTestTestingPortOpenApiVO(oswCableTestTestingPortOpenApiVO).Execute()

Start cable test



### Example

```go
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
	oswCableTestTestingPortOpenApiVO := *openapiclient.NewOswCableTestTestingPortOpenApiVO() // OswCableTestTestingPortOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.StartCableTest(context.Background(), omadacId, siteId, switchMac).OswCableTestTestingPortOpenApiVO(oswCableTestTestingPortOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.StartCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCableTest`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.StartCableTest`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiStartCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswCableTestTestingPortOpenApiVO** | [**OswCableTestTestingPortOpenApiVO**](OswCableTestTestingPortOpenApiVO.md) |  | 

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


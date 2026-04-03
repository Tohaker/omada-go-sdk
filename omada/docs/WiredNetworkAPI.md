# WiredNetworkAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchIsolateLanNetwork1**](WiredNetworkAPI.md#batchisolatelannetwork1) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/lan-networks/batch-isolate | Batch isolate network
[**CheckNetworkParamWhenCreate**](WiredNetworkAPI.md#checknetworkparamwhencreate) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/param-check | Check basic parameters when creating network
[**CheckNetworkParamWhenModify**](WiredNetworkAPI.md#checknetworkparamwhenmodify) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/param-check | Check basic parameters when modifying the network
[**CheckParamAndGetPreConfig**](WiredNetworkAPI.md#checkparamandgetpreconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/check | Check all parameters and generate configuration when creating network
[**CheckParamAndGetPreConfigWhenModify**](WiredNetworkAPI.md#checkparamandgetpreconfigwhenmodify) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/check | Check all parameters and generate configuration when modifying the network
[**CheckPortBindingParamWhenCreate**](WiredNetworkAPI.md#checkportbindingparamwhencreate) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/ports-check | Check selected ports when creating network
[**CheckPortBindingParamWhenModify**](WiredNetworkAPI.md#checkportbindingparamwhenmodify) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/ports-check | Check port binding when modifying network
[**CheckVirtualWanUsed1**](WiredNetworkAPI.md#checkvirtualwanused1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/virtual-wans/delete-check | Check Virtual Wan Used
[**CheckWanLanStatus1**](WiredNetworkAPI.md#checkwanlanstatus1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wan-lan-status | Check WAN-LAN status
[**ConfirmCreateVlanNetwork**](WiredNetworkAPI.md#confirmcreatevlannetwork) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/confirm | Confirm create lan network
[**ConfirmModifyVlanNetwork**](WiredNetworkAPI.md#confirmmodifyvlannetwork) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/confirm | Confirm modify lan network
[**CreateLanDns**](WiredNetworkAPI.md#createlandns) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/lan/dns | Create a LAN DNS rule
[**CreateLanNetwork**](WiredNetworkAPI.md#createlannetwork) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks | Create LAN network
[**CreateLanNetworkV2**](WiredNetworkAPI.md#createlannetworkv2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/lan-networks | Create LAN network
[**CreateLanProfile**](WiredNetworkAPI.md#createlanprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/lan-profiles | Create new LAN profile
[**CreateOswLanProfile**](WiredNetworkAPI.md#createoswlanprofile) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/lan-profiles | Create new switch profile
[**CreateVirtualWan**](WiredNetworkAPI.md#createvirtualwan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/virtual-wans | Create virtual WAN
[**CreateVlans1**](WiredNetworkAPI.md#createvlans1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/lan-vlans | Batch create vlans
[**DeleteLanDns**](WiredNetworkAPI.md#deletelandns) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/lan/dns/{dnsId} | Delete an existing LAN DNS rule
[**DeleteLanNetwork**](WiredNetworkAPI.md#deletelannetwork) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks/{networkId} | Delete LAN network
[**DeleteLanProfile**](WiredNetworkAPI.md#deletelanprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/lan-profiles/{profileId} | Delete an existing LAN profile
[**DeleteOswLanProfile**](WiredNetworkAPI.md#deleteoswlanprofile) | **Delete** /openapi/v2/{omadacId}/sites/{siteId}/lan-profiles/{profileId} | Delete an existing switch profile
[**DeleteVirtualWan1**](WiredNetworkAPI.md#deletevirtualwan1) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/virtual-wans/{virtualWanId} | Delete virtual WAN
[**GetAllInterfacesForBatchIsolate1**](WiredNetworkAPI.md#getallinterfacesforbatchisolate1) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/lan-networks/isolate/interfaces | Get interface List
[**GetAllLanNetworks**](WiredNetworkAPI.md#getalllannetworks) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks/all | Get all networks for the omada id and site id
[**GetAllValidVlansInSiteForSwitchOuiBasedVlan1**](WiredNetworkAPI.md#getallvalidvlansinsiteforswitchouibasedvlan1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/valid-vlans | Get valid vlanList for switch oui based vlan
[**GetAutoEffectDevicesWhenCreate**](WiredNetworkAPI.md#getautoeffectdeviceswhencreate) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/auto-effect-devices | Get auto effect devices when creating network
[**GetAutoEffectDevicesWhenModify**](WiredNetworkAPI.md#getautoeffectdeviceswhenmodify) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/auto-effect-devices | Get auto effect devices when modifying network
[**GetAutoSelectDevicesWhenModify**](WiredNetworkAPI.md#getautoselectdeviceswhenmodify) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/auto-select-devices | Get auto select devices when modifying network
[**GetAvailableVirtualWan**](WiredNetworkAPI.md#getavailablevirtualwan) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/available-virtual-wans | Query available virtual WAN list
[**GetAvailableWanPorts1**](WiredNetworkAPI.md#getavailablewanports1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/wan-ports | Get available wan ports
[**GetBandScanResult**](WiredNetworkAPI.md#getbandscanresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/band-scan/{portUuid} | Get band scan result.
[**GetDeliveringProcess**](WiredNetworkAPI.md#getdeliveringprocess) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/process | Get network delivering process and config
[**GetDeliveringProcessWithoutConfig**](WiredNetworkAPI.md#getdeliveringprocesswithoutconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/brief-process | Get network delivering process
[**GetDhcpServerDevCapForNetwork**](WiredNetworkAPI.md#getdhcpserverdevcapfornetwork) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/dhcp-server/capabilities | Get the capabilities of the DHCP server under the network
[**GetDhcpServerInfoForNetwork**](WiredNetworkAPI.md#getdhcpserverinfofornetwork) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/dhcp-server-info | Get the address pool of the DHCP server under the network and the number of available IPs
[**GetGridLanDns**](WiredNetworkAPI.md#getgridlandns) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/lan/dns | Get LAN DNS list
[**GetGridSupportVlanNetworkDevicesBySite**](WiredNetworkAPI.md#getgridsupportvlannetworkdevicesbysite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/devices | Get devices list that can be dhcp server
[**GetGridVirtualWan**](WiredNetworkAPI.md#getgridvirtualwan) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/virtual-wans | Query virtual WAN list
[**GetGridVlanNetworkAffectedDevicePorts**](WiredNetworkAPI.md#getgridvlannetworkaffecteddeviceports) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/vlan/{vlan}/ports | Get the ports of devices that use the network
[**GetGridVlans**](WiredNetworkAPI.md#getgridvlans) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/vlans | Get the networks whose purpose is vlan
[**GetInterfaceLanNetwork**](WiredNetworkAPI.md#getinterfacelannetwork) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks/interface | Get all \&quot;single\&quot;/\&quot;multi\&quot; interface lan network
[**GetInterfaceLanNetworkV2**](WiredNetworkAPI.md#getinterfacelannetworkv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/lan-networks/interface | Get all \&quot;single\&quot;/\&quot;multi\&quot; interface lan network V2
[**GetInternet**](WiredNetworkAPI.md#getinternet) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet | Get internet info
[**GetInternetBasicPortInfo1**](WiredNetworkAPI.md#getinternetbasicportinfo1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/basic-info | Get internet basic info
[**GetInternetLoadBalance1**](WiredNetworkAPI.md#getinternetloadbalance1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/load-balance | Get internet load balance config
[**GetIpptInternet**](WiredNetworkAPI.md#getipptinternet) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/internet | Get internet info by mac
[**GetIspScanResult**](WiredNetworkAPI.md#getispscanresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/isp-scan/{portUuid} | Get Isp scan result
[**GetLanNetwork**](WiredNetworkAPI.md#getlannetwork) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks/{networkId} | Get LAN network
[**GetLanNetworkList**](WiredNetworkAPI.md#getlannetworklist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks | Get LAN network list
[**GetLanNetworkListV2**](WiredNetworkAPI.md#getlannetworklistv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/lan-networks | Get LAN network list
[**GetLanNetworkListV3**](WiredNetworkAPI.md#getlannetworklistv3) | **Get** /openapi/v3/{omadacId}/sites/{siteId}/lan-networks | Get LAN network list V3
[**GetLanProfileList**](WiredNetworkAPI.md#getlanprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lan-profiles | Get LAN profile list
[**GetLoadBalanceWeightStatus**](WiredNetworkAPI.md#getloadbalanceweightstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/load-balance/status | Check Wan Loadbalance status
[**GetLocationAndIspInfo**](WiredNetworkAPI.md#getlocationandispinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/location-isp | Get location and ISP info
[**GetLteWanPortsConfig**](WiredNetworkAPI.md#getltewanportsconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/lte/ports-config | Get internet lte wan ports config
[**GetNetworksWithServersForReservation**](WiredNetworkAPI.md#getnetworkswithserversforreservation) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/server-for-reservation | Get all networks that can reserve IP addresses and their corresponding servers
[**GetOswForGivenLanNetworkIds**](WiredNetworkAPI.md#getoswforgivenlannetworkids) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/switches | Get LAN network list
[**GetOswLanProfileList**](WiredNetworkAPI.md#getoswlanprofilelist) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/lan-profiles | Get switch profile list
[**GetOswPortLanNetInfo**](WiredNetworkAPI.md#getoswportlannetinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lan-switch-setting | Get switch port profile info
[**GetSelectDevicePortsInfo**](WiredNetworkAPI.md#getselectdeviceportsinfo) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/devices/ports | Get the port information of the selected devices when creating network
[**GetSelectDevicePortsInfoWhenModify**](WiredNetworkAPI.md#getselectdeviceportsinfowhenmodify) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/devices/ports | Get the port information of the selected devices when modifying the network
[**GetSelectStackPortsInfo**](WiredNetworkAPI.md#getselectstackportsinfo) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/stack/ports | Get the port information of the selected stacks when creating network
[**GetSelectStackPortsInfoWhenModify**](WiredNetworkAPI.md#getselectstackportsinfowhenmodify) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/stack/ports | Get the port information of the selected stacks when modifying the network
[**GetSimCardBandScanResult**](WiredNetworkAPI.md#getsimcardbandscanresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wan/networks/band-scan/{portUuid}/{simCard} | Obtain the bandScan result of the corresponding SIM card
[**GetSimCardIspScanResult**](WiredNetworkAPI.md#getsimcardispscanresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wan/networks/isp-scan/{portUuid}/{simCard} | Obtain the ispScan result of the corresponding SIM card
[**GetSupportInfo**](WiredNetworkAPI.md#getsupportinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/models | Get supported gateway model list for pre-configuration
[**GetUseLanProfileES**](WiredNetworkAPI.md#getuselanprofilees) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lan-profiles/{profileId}/es | Query Agile Series switches that used the profile
[**GetUseNativeNetworkOsw**](WiredNetworkAPI.md#getusenativenetworkosw) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks/es | Query switches that used the Native Network&#39;s profile
[**GetUseNativeNetworkOswV2**](WiredNetworkAPI.md#getusenativenetworkoswv2) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks/es-native | Query switches that used the Native Network&#39;s profile V2
[**GetVlanNetworkAffectedDevice1**](WiredNetworkAPI.md#getvlannetworkaffecteddevice1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/vlan/{vlan}/devices | Get devices list that use the network
[**GetVlanNetworkAffectedSsid1**](WiredNetworkAPI.md#getvlannetworkaffectedssid1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/networks/{networkId}/vlan/{vlan}/ssids | Get grid ssid list that use the vlan
[**GetWanPortsConfig**](WiredNetworkAPI.md#getwanportsconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/internet/ports-config | Get internet ports config
[**GetWanPortsMaxSpeed1**](WiredNetworkAPI.md#getwanportsmaxspeed1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wan-max-speed | Get wan ports max speed
[**ModifyGatewayModel**](WiredNetworkAPI.md#modifygatewaymodel) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/internet/change-model | Modify gateway mode for pre-configuration
[**ModifyInternetBasicPortInfo1**](WiredNetworkAPI.md#modifyinternetbasicportinfo1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/internet/basic-info | Modify the internet basic configuration
[**ModifyInternetLoadBalance1**](WiredNetworkAPI.md#modifyinternetloadbalance1) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/internet/load-balance | Modify internet load balance config
[**ModifyIpptLteWanPortConfig**](WiredNetworkAPI.md#modifyipptltewanportconfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/internet/lte/ports-config | Modify internet lte port config for ippt
[**ModifyIpptWanMode**](WiredNetworkAPI.md#modifyipptwanmode) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/internet/wan-mode | Modify internet wan mode
[**ModifyLanDns**](WiredNetworkAPI.md#modifylandns) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/lan/dns/{dnsId} | Modify an existing LAN DNS rule
[**ModifyLanNetwork**](WiredNetworkAPI.md#modifylannetwork) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/lan-networks/{networkId} | Modify LAN network
[**ModifyLanNetworkV2**](WiredNetworkAPI.md#modifylannetworkv2) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/lan-networks/{networkId} | Modify LAN network
[**ModifyLanProfile**](WiredNetworkAPI.md#modifylanprofile) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/lan-profiles/{profileId} | Modify a LAN profile
[**ModifyLteWanPortConfig**](WiredNetworkAPI.md#modifyltewanportconfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/internet/lte/ports-config | Modify internet lte port config
[**ModifyOswLanProfile**](WiredNetworkAPI.md#modifyoswlanprofile) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/lan-profiles/{profileId} | Modify a switch profile
[**ModifyVirtualWan**](WiredNetworkAPI.md#modifyvirtualwan) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/virtual-wans/{virtualWanId} | Modify virtual WAN
[**ModifyVirtualWanStatus**](WiredNetworkAPI.md#modifyvirtualwanstatus) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/virtual-wans/{virtualWanId}/status | Modify virtual WAN status
[**ModifyWanPortSettings**](WiredNetworkAPI.md#modifywanportsettings) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/wan/networks/port-setting | Modify Wan Port Settings
[**ModifyWanPorts**](WiredNetworkAPI.md#modifywanports) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/wan-ports | Modify Wan Ports
[**ModifyWanPortsConfig**](WiredNetworkAPI.md#modifywanportsconfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/internet/ports-config | Modify internet ports config
[**NetworkMapping**](WiredNetworkAPI.md#networkmapping) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/network-mapping | Network mapping
[**SpeedTest**](WiredNetworkAPI.md#speedtest) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/speedTest | SpeedTest
[**StartBandScan**](WiredNetworkAPI.md#startbandscan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wan/networks/band-scan | BandScan
[**StartIspScan**](WiredNetworkAPI.md#startispscan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wan/networks/isp-scan | IspScan



## BatchIsolateLanNetwork1

> OperationResponseWithoutResult BatchIsolateLanNetwork1(ctx, omadacId, siteId).BatchIsolateInterfaceOpenApiVO(batchIsolateInterfaceOpenApiVO).Execute()

Batch isolate network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	batchIsolateInterfaceOpenApiVO := *openapiclient.NewBatchIsolateInterfaceOpenApiVO() // BatchIsolateInterfaceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.BatchIsolateLanNetwork1(context.Background(), omadacId, siteId).BatchIsolateInterfaceOpenApiVO(batchIsolateInterfaceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.BatchIsolateLanNetwork1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchIsolateLanNetwork1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.BatchIsolateLanNetwork1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchIsolateLanNetwork1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchIsolateInterfaceOpenApiVO** | [**BatchIsolateInterfaceOpenApiVO**](BatchIsolateInterfaceOpenApiVO.md) |  | 

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


## CheckNetworkParamWhenCreate

> OperationResponseLanNetworkQueryOpenApiV3VO CheckNetworkParamWhenCreate(ctx, omadacId, siteId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

Check basic parameters when creating network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CheckNetworkParamWhenCreate(context.Background(), omadacId, siteId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CheckNetworkParamWhenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckNetworkParamWhenCreate`: OperationResponseLanNetworkQueryOpenApiV3VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CheckNetworkParamWhenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckNetworkParamWhenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseLanNetworkQueryOpenApiV3VO**](OperationResponseLanNetworkQueryOpenApiV3VO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckNetworkParamWhenModify

> OperationResponseLanNetworkQueryOpenApiV3VO CheckNetworkParamWhenModify(ctx, omadacId, siteId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

Check basic parameters when modifying the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CheckNetworkParamWhenModify(context.Background(), omadacId, siteId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CheckNetworkParamWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckNetworkParamWhenModify`: OperationResponseLanNetworkQueryOpenApiV3VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CheckNetworkParamWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckNetworkParamWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseLanNetworkQueryOpenApiV3VO**](OperationResponseLanNetworkQueryOpenApiV3VO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckParamAndGetPreConfig

> OperationResponseVlanPreConfigOpenApiVO CheckParamAndGetPreConfig(ctx, omadacId, siteId).CreateVlanParamOpenApiVO(createVlanParamOpenApiVO).Execute()

Check all parameters and generate configuration when creating network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	createVlanParamOpenApiVO := *openapiclient.NewCreateVlanParamOpenApiVO(*openapiclient.NewSelectPortBindingBriefVO(), *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example")) // CreateVlanParamOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CheckParamAndGetPreConfig(context.Background(), omadacId, siteId).CreateVlanParamOpenApiVO(createVlanParamOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CheckParamAndGetPreConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckParamAndGetPreConfig`: OperationResponseVlanPreConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CheckParamAndGetPreConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckParamAndGetPreConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVlanParamOpenApiVO** | [**CreateVlanParamOpenApiVO**](CreateVlanParamOpenApiVO.md) |  | 

### Return type

[**OperationResponseVlanPreConfigOpenApiVO**](OperationResponseVlanPreConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckParamAndGetPreConfigWhenModify

> OperationResponseVlanPreConfigOpenApiVO CheckParamAndGetPreConfigWhenModify(ctx, omadacId, siteId, networkId).ModifyVlanParamOpenApiVO(modifyVlanParamOpenApiVO).Execute()

Check all parameters and generate configuration when modifying the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	modifyVlanParamOpenApiVO := *openapiclient.NewModifyVlanParamOpenApiVO(*openapiclient.NewSelectPortBindingBriefVO(), *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example")) // ModifyVlanParamOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CheckParamAndGetPreConfigWhenModify(context.Background(), omadacId, siteId, networkId).ModifyVlanParamOpenApiVO(modifyVlanParamOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CheckParamAndGetPreConfigWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckParamAndGetPreConfigWhenModify`: OperationResponseVlanPreConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CheckParamAndGetPreConfigWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckParamAndGetPreConfigWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyVlanParamOpenApiVO** | [**ModifyVlanParamOpenApiVO**](ModifyVlanParamOpenApiVO.md) |  | 

### Return type

[**OperationResponseVlanPreConfigOpenApiVO**](OperationResponseVlanPreConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckPortBindingParamWhenCreate

> OperationResponseListCheckFailedDeviceInfoForVlanVO CheckPortBindingParamWhenCreate(ctx, omadacId, siteId).SelectPortBindingVO(selectPortBindingVO).Execute()

Check selected ports when creating network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	selectPortBindingVO := *openapiclient.NewSelectPortBindingVO(int32(123), int32(123)) // SelectPortBindingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CheckPortBindingParamWhenCreate(context.Background(), omadacId, siteId).SelectPortBindingVO(selectPortBindingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CheckPortBindingParamWhenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPortBindingParamWhenCreate`: OperationResponseListCheckFailedDeviceInfoForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CheckPortBindingParamWhenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckPortBindingParamWhenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectPortBindingVO** | [**SelectPortBindingVO**](SelectPortBindingVO.md) |  | 

### Return type

[**OperationResponseListCheckFailedDeviceInfoForVlanVO**](OperationResponseListCheckFailedDeviceInfoForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckPortBindingParamWhenModify

> OperationResponseListCheckFailedDeviceInfoForVlanVO CheckPortBindingParamWhenModify(ctx, omadacId, siteId, networkId).SelectPortBindingVO(selectPortBindingVO).Execute()

Check port binding when modifying network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	selectPortBindingVO := *openapiclient.NewSelectPortBindingVO(int32(123), int32(123)) // SelectPortBindingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CheckPortBindingParamWhenModify(context.Background(), omadacId, siteId, networkId).SelectPortBindingVO(selectPortBindingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CheckPortBindingParamWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPortBindingParamWhenModify`: OperationResponseListCheckFailedDeviceInfoForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CheckPortBindingParamWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckPortBindingParamWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **selectPortBindingVO** | [**SelectPortBindingVO**](SelectPortBindingVO.md) |  | 

### Return type

[**OperationResponseListCheckFailedDeviceInfoForVlanVO**](OperationResponseListCheckFailedDeviceInfoForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckVirtualWanUsed1

> OperationResponseVirtualWanIdUsedOpenApiVO CheckVirtualWanUsed1(ctx, omadacId, siteId).VirtualWanId(virtualWanId).Execute()

Check Virtual Wan Used



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	virtualWanId := "virtualWanId_example" // string | Virtual WAN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CheckVirtualWanUsed1(context.Background(), omadacId, siteId).VirtualWanId(virtualWanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CheckVirtualWanUsed1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckVirtualWanUsed1`: OperationResponseVirtualWanIdUsedOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CheckVirtualWanUsed1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckVirtualWanUsed1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualWanId** | **string** | Virtual WAN ID | 

### Return type

[**OperationResponseVirtualWanIdUsedOpenApiVO**](OperationResponseVirtualWanIdUsedOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckWanLanStatus1

> CheckWanLanStatusOpenApiVO CheckWanLanStatus1(ctx, omadacId, siteId).Execute()

Check WAN-LAN status



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.CheckWanLanStatus1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CheckWanLanStatus1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckWanLanStatus1`: CheckWanLanStatusOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CheckWanLanStatus1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckWanLanStatus1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CheckWanLanStatusOpenApiVO**](CheckWanLanStatusOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmCreateVlanNetwork

> OperationResponseVlanNetworkIdListVO ConfirmCreateVlanNetwork(ctx, omadacId, siteId).CreateVlanParamOpenApiVO(createVlanParamOpenApiVO).Execute()

Confirm create lan network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	createVlanParamOpenApiVO := *openapiclient.NewCreateVlanParamOpenApiVO(*openapiclient.NewSelectPortBindingBriefVO(), *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example")) // CreateVlanParamOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ConfirmCreateVlanNetwork(context.Background(), omadacId, siteId).CreateVlanParamOpenApiVO(createVlanParamOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ConfirmCreateVlanNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmCreateVlanNetwork`: OperationResponseVlanNetworkIdListVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ConfirmCreateVlanNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmCreateVlanNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVlanParamOpenApiVO** | [**CreateVlanParamOpenApiVO**](CreateVlanParamOpenApiVO.md) |  | 

### Return type

[**OperationResponseVlanNetworkIdListVO**](OperationResponseVlanNetworkIdListVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmModifyVlanNetwork

> OperationResponseVlanNetworkIdListVO ConfirmModifyVlanNetwork(ctx, omadacId, siteId, networkId).ModifyVlanParamOpenApiVO(modifyVlanParamOpenApiVO).Execute()

Confirm modify lan network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	modifyVlanParamOpenApiVO := *openapiclient.NewModifyVlanParamOpenApiVO(*openapiclient.NewSelectPortBindingBriefVO(), *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example")) // ModifyVlanParamOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ConfirmModifyVlanNetwork(context.Background(), omadacId, siteId, networkId).ModifyVlanParamOpenApiVO(modifyVlanParamOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ConfirmModifyVlanNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmModifyVlanNetwork`: OperationResponseVlanNetworkIdListVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ConfirmModifyVlanNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmModifyVlanNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyVlanParamOpenApiVO** | [**ModifyVlanParamOpenApiVO**](ModifyVlanParamOpenApiVO.md) |  | 

### Return type

[**OperationResponseVlanNetworkIdListVO**](OperationResponseVlanNetworkIdListVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanDns

> OperationResponseWithoutResult CreateLanDns(ctx, omadacId, siteId).LanDnsOpenApiVO(lanDnsOpenApiVO).Execute()

Create a LAN DNS rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	lanDnsOpenApiVO := *openapiclient.NewLanDnsOpenApiVO("Domain_example", false, "Name_example", int32(123)) // LanDnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CreateLanDns(context.Background(), omadacId, siteId).LanDnsOpenApiVO(lanDnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CreateLanDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanDns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CreateLanDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanDnsOpenApiVO** | [**LanDnsOpenApiVO**](LanDnsOpenApiVO.md) |  | 

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


## CreateLanNetwork

> ResponseIdVO CreateLanNetwork(ctx, omadacId, siteId).LanNetworkOpenApiVO(lanNetworkOpenApiVO).Execute()

Create LAN network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	lanNetworkOpenApiVO := *openapiclient.NewLanNetworkOpenApiVO(false, "Name_example", int32(123)) // LanNetworkOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CreateLanNetwork(context.Background(), omadacId, siteId).LanNetworkOpenApiVO(lanNetworkOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CreateLanNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanNetwork`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CreateLanNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanNetworkOpenApiVO** | [**LanNetworkOpenApiVO**](LanNetworkOpenApiVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanNetworkV2

> ResponseIdVO CreateLanNetworkV2(ctx, omadacId, siteId).LanNetworkOpenApiV2VO(lanNetworkOpenApiV2VO).Execute()

Create LAN network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	lanNetworkOpenApiV2VO := *openapiclient.NewLanNetworkOpenApiV2VO(false, "Name_example", int32(123)) // LanNetworkOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CreateLanNetworkV2(context.Background(), omadacId, siteId).LanNetworkOpenApiV2VO(lanNetworkOpenApiV2VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CreateLanNetworkV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanNetworkV2`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CreateLanNetworkV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanNetworkV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanNetworkOpenApiV2VO** | [**LanNetworkOpenApiV2VO**](LanNetworkOpenApiV2VO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanProfile

> ResponseIdVO CreateLanProfile(ctx, omadacId, siteId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()

Create new LAN profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	lanProfileConfigOpenApiVO := *openapiclient.NewLanProfileConfigOpenApiVO(int32(123), int32(123), false, false, "Name_example", "NativeNetworkId_example", int32(123), false, false) // LanProfileConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CreateLanProfile(context.Background(), omadacId, siteId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CreateLanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanProfile`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CreateLanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanProfileConfigOpenApiVO** | [**LanProfileConfigOpenApiVO**](LanProfileConfigOpenApiVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOswLanProfile

> ResponseIdVO CreateOswLanProfile(ctx, omadacId, siteId).LanProfileSettingOpenApiVO(lanProfileSettingOpenApiVO).Execute()

Create new switch profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	lanProfileSettingOpenApiVO := *openapiclient.NewLanProfileSettingOpenApiVO(int32(123), int32(123), false, false, "Name_example", int32(123), false, false) // LanProfileSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CreateOswLanProfile(context.Background(), omadacId, siteId).LanProfileSettingOpenApiVO(lanProfileSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CreateOswLanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOswLanProfile`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CreateOswLanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOswLanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanProfileSettingOpenApiVO** | [**LanProfileSettingOpenApiVO**](LanProfileSettingOpenApiVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualWan

> OperationResponseWithoutResult CreateVirtualWan(ctx, omadacId, siteId).VirtualWanConfigOpenApiVO(virtualWanConfigOpenApiVO).Execute()

Create virtual WAN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	virtualWanConfigOpenApiVO := *openapiclient.NewVirtualWanConfigOpenApiVO("Name_example", "PhysicalWanId_example", *openapiclient.NewVirtualWanIpv4SettingConfigOpenApiVO("Proto_example", int32(123))) // VirtualWanConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CreateVirtualWan(context.Background(), omadacId, siteId).VirtualWanConfigOpenApiVO(virtualWanConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CreateVirtualWan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVirtualWan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CreateVirtualWan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualWanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualWanConfigOpenApiVO** | [**VirtualWanConfigOpenApiVO**](VirtualWanConfigOpenApiVO.md) |  | 

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


## CreateVlans1

> OperationResponseWithoutResult CreateVlans1(ctx, omadacId, siteId).CreateVLANs(createVLANs).Execute()

Batch create vlans



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	createVLANs := *openapiclient.NewCreateVLANs(false, "Name_example", "Vlans_example") // CreateVLANs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.CreateVlans1(context.Background(), omadacId, siteId).CreateVLANs(createVLANs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.CreateVlans1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVlans1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.CreateVlans1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVlans1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVLANs** | [**CreateVLANs**](CreateVLANs.md) |  | 

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


## DeleteLanDns

> OperationResponseWithoutResult DeleteLanDns(ctx, omadacId, siteId, dnsId).Execute()

Delete an existing LAN DNS rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	dnsId := "dnsId_example" // string | LAN DNS ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.DeleteLanDns(context.Background(), omadacId, siteId, dnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.DeleteLanDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLanDns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.DeleteLanDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**dnsId** | **string** | LAN DNS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanDnsRequest struct via the builder pattern


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


## DeleteLanNetwork

> OperationResponseWithoutResult DeleteLanNetwork(ctx, omadacId, siteId, networkId).Execute()

Delete LAN network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.DeleteLanNetwork(context.Background(), omadacId, siteId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.DeleteLanNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLanNetwork`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.DeleteLanNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanNetworkRequest struct via the builder pattern


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


## DeleteLanProfile

> OperationResponseWithoutResult DeleteLanProfile(ctx, omadacId, siteId, profileId).Execute()

Delete an existing LAN profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	profileId := "profileId_example" // string | LAN profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.DeleteLanProfile(context.Background(), omadacId, siteId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.DeleteLanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLanProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.DeleteLanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | LAN profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanProfileRequest struct via the builder pattern


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


## DeleteOswLanProfile

> OperationResponseWithoutResult DeleteOswLanProfile(ctx, omadacId, siteId, profileId).Execute()

Delete an existing switch profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	profileId := "profileId_example" // string | Switch profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.DeleteOswLanProfile(context.Background(), omadacId, siteId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.DeleteOswLanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOswLanProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.DeleteOswLanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | Switch profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOswLanProfileRequest struct via the builder pattern


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


## DeleteVirtualWan1

> OperationResponseWithoutResult DeleteVirtualWan1(ctx, omadacId, siteId, virtualWanId).Execute()

Delete virtual WAN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	virtualWanId := "virtualWanId_example" // string | Virtual WAN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.DeleteVirtualWan1(context.Background(), omadacId, siteId, virtualWanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.DeleteVirtualWan1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVirtualWan1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.DeleteVirtualWan1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**virtualWanId** | **string** | Virtual WAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualWan1Request struct via the builder pattern


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


## GetAllInterfacesForBatchIsolate1

> OperationResponseListInterfaceForBatchIsolateOpenApiVO GetAllInterfacesForBatchIsolate1(ctx, omadacId, siteId).SearchKey(searchKey).IsolationFilter(isolationFilter).Execute()

Get interface List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field  (optional)
	isolationFilter := int32(56) // int32 | non-isolation set 0; isolation set 1; don't filter isolation vlan set 2. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetAllInterfacesForBatchIsolate1(context.Background(), omadacId, siteId).SearchKey(searchKey).IsolationFilter(isolationFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetAllInterfacesForBatchIsolate1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllInterfacesForBatchIsolate1`: OperationResponseListInterfaceForBatchIsolateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetAllInterfacesForBatchIsolate1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInterfacesForBatchIsolate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchKey** | **string** | Fuzzy query parameters, support field  | 
 **isolationFilter** | **int32** | non-isolation set 0; isolation set 1; don&#39;t filter isolation vlan set 2. | 

### Return type

[**OperationResponseListInterfaceForBatchIsolateOpenApiVO**](OperationResponseListInterfaceForBatchIsolateOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllLanNetworks

> []LanNetworkVO GetAllLanNetworks(ctx, omadacId, siteId).Execute()

Get all networks for the omada id and site id



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetAllLanNetworks(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetAllLanNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllLanNetworks`: []LanNetworkVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetAllLanNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLanNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]LanNetworkVO**](LanNetworkVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllValidVlansInSiteForSwitchOuiBasedVlan1

> GridVOLanNetworkVlansOpenApiVO GetAllValidVlansInSiteForSwitchOuiBasedVlan1(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get valid vlanList for switch oui based vlan



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetAllValidVlansInSiteForSwitchOuiBasedVlan1(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetAllValidVlansInSiteForSwitchOuiBasedVlan1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllValidVlansInSiteForSwitchOuiBasedVlan1`: GridVOLanNetworkVlansOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetAllValidVlansInSiteForSwitchOuiBasedVlan1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllValidVlansInSiteForSwitchOuiBasedVlan1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**GridVOLanNetworkVlansOpenApiVO**](GridVOLanNetworkVlansOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoEffectDevicesWhenCreate

> OperationResponseAutoEffectDevicesForVlanVO GetAutoEffectDevicesWhenCreate(ctx, omadacId, siteId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

Get auto effect devices when creating network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetAutoEffectDevicesWhenCreate(context.Background(), omadacId, siteId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetAutoEffectDevicesWhenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoEffectDevicesWhenCreate`: OperationResponseAutoEffectDevicesForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetAutoEffectDevicesWhenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoEffectDevicesWhenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseAutoEffectDevicesForVlanVO**](OperationResponseAutoEffectDevicesForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoEffectDevicesWhenModify

> OperationResponseAutoEffectDevicesForVlanVO GetAutoEffectDevicesWhenModify(ctx, omadacId, siteId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

Get auto effect devices when modifying network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetAutoEffectDevicesWhenModify(context.Background(), omadacId, siteId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetAutoEffectDevicesWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoEffectDevicesWhenModify`: OperationResponseAutoEffectDevicesForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetAutoEffectDevicesWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoEffectDevicesWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseAutoEffectDevicesForVlanVO**](OperationResponseAutoEffectDevicesForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoSelectDevicesWhenModify

> OperationResponseAutoSelectDevicesForVlanVO GetAutoSelectDevicesWhenModify(ctx, omadacId, siteId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

Get auto select devices when modifying network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetAutoSelectDevicesWhenModify(context.Background(), omadacId, siteId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetAutoSelectDevicesWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoSelectDevicesWhenModify`: OperationResponseAutoSelectDevicesForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetAutoSelectDevicesWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoSelectDevicesWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseAutoSelectDevicesForVlanVO**](OperationResponseAutoSelectDevicesForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableVirtualWan

> OperationResponseVirtualWanAvailablesOpenApiVO GetAvailableVirtualWan(ctx, omadacId, siteId).Execute()

Query available virtual WAN list



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetAvailableVirtualWan(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetAvailableVirtualWan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableVirtualWan`: OperationResponseVirtualWanAvailablesOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetAvailableVirtualWan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableVirtualWanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVirtualWanAvailablesOpenApiVO**](OperationResponseVirtualWanAvailablesOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableWanPorts1

> OperationResponseAvailableWanResultOpenApiVO GetAvailableWanPorts1(ctx, omadacId, siteId).Function(function).Execute()

Get available wan ports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	function := int32(56) // int32 | function used for wan ports query. 0: ACL, 1: QOS, 2: IP_MAC_BLINDING, 3:IGMP_PROXY, 4: VIRTUAL_WAN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetAvailableWanPorts1(context.Background(), omadacId, siteId).Function(function).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetAvailableWanPorts1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableWanPorts1`: OperationResponseAvailableWanResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetAvailableWanPorts1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableWanPorts1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **function** | **int32** | function used for wan ports query. 0: ACL, 1: QOS, 2: IP_MAC_BLINDING, 3:IGMP_PROXY, 4: VIRTUAL_WAN | 

### Return type

[**OperationResponseAvailableWanResultOpenApiVO**](OperationResponseAvailableWanResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBandScanResult

> OperationResponseBandScanResultOpenApiVO GetBandScanResult(ctx, omadacId, siteId, portUuid).Execute()

Get band scan result.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	portUuid := "portUuid_example" // string | Port uuid to get band scan result

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetBandScanResult(context.Background(), omadacId, siteId, portUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetBandScanResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandScanResult`: OperationResponseBandScanResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetBandScanResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portUuid** | **string** | Port uuid to get band scan result | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseBandScanResultOpenApiVO**](OperationResponseBandScanResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeliveringProcess

> OperationResponseVlanNetworkDeliverDataVO GetDeliveringProcess(ctx, omadacId, siteId, networkId).Execute()

Get network delivering process and config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetDeliveringProcess(context.Background(), omadacId, siteId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetDeliveringProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeliveringProcess`: OperationResponseVlanNetworkDeliverDataVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetDeliveringProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveringProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseVlanNetworkDeliverDataVO**](OperationResponseVlanNetworkDeliverDataVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeliveringProcessWithoutConfig

> OperationResponseVlanNetworkDeliverBriefDataVO GetDeliveringProcessWithoutConfig(ctx, omadacId, siteId, networkId).Execute()

Get network delivering process



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetDeliveringProcessWithoutConfig(context.Background(), omadacId, siteId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetDeliveringProcessWithoutConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeliveringProcessWithoutConfig`: OperationResponseVlanNetworkDeliverBriefDataVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetDeliveringProcessWithoutConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveringProcessWithoutConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseVlanNetworkDeliverBriefDataVO**](OperationResponseVlanNetworkDeliverBriefDataVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDhcpServerDevCapForNetwork

> OperationResponseDhcpServerDevCapForVlanVO GetDhcpServerDevCapForNetwork(ctx, omadacId, siteId, networkId).Execute()

Get the capabilities of the DHCP server under the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetDhcpServerDevCapForNetwork(context.Background(), omadacId, siteId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetDhcpServerDevCapForNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpServerDevCapForNetwork`: OperationResponseDhcpServerDevCapForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetDhcpServerDevCapForNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpServerDevCapForNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseDhcpServerDevCapForVlanVO**](OperationResponseDhcpServerDevCapForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDhcpServerInfoForNetwork

> OperationResponseDhcpServerForVlanVO GetDhcpServerInfoForNetwork(ctx, omadacId, siteId, networkId).Execute()

Get the address pool of the DHCP server under the network and the number of available IPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetDhcpServerInfoForNetwork(context.Background(), omadacId, siteId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetDhcpServerInfoForNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpServerInfoForNetwork`: OperationResponseDhcpServerForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetDhcpServerInfoForNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpServerInfoForNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseDhcpServerForVlanVO**](OperationResponseDhcpServerForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridLanDns

> OperationResponseGridVOLanDnsOpenApiVO GetGridLanDns(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get LAN DNS list



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetGridLanDns(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetGridLanDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridLanDns`: OperationResponseGridVOLanDnsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetGridLanDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridLanDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOLanDnsOpenApiVO**](OperationResponseGridVOLanDnsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSupportVlanNetworkDevicesBySite

> OperationResponseGridVODeviceVO GetGridSupportVlanNetworkDevicesBySite(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get devices list that can be dhcp server



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetGridSupportVlanNetworkDevicesBySite(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetGridSupportVlanNetworkDevicesBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSupportVlanNetworkDevicesBySite`: OperationResponseGridVODeviceVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetGridSupportVlanNetworkDevicesBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSupportVlanNetworkDevicesBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVODeviceVO**](OperationResponseGridVODeviceVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVirtualWan

> OperationResponseVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO GetGridVirtualWan(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Query virtual WAN list



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetGridVirtualWan(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetGridVirtualWan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVirtualWan`: OperationResponseVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetGridVirtualWan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVirtualWanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO**](OperationResponseVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVlanNetworkAffectedDevicePorts

> OperationResponseVlanForPortDisplayGridVOVlanNetworkAffectingDeviceDetailVO GetGridVlanNetworkAffectedDevicePorts(ctx, omadacId, siteId, networkId, vlan).Page(page).PageSize(pageSize).Execute()

Get the ports of devices that use the network



### Example

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
	networkId := "networkId_example" // string | Network ID
	vlan := "vlan_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetGridVlanNetworkAffectedDevicePorts(context.Background(), omadacId, siteId, networkId, vlan).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetGridVlanNetworkAffectedDevicePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVlanNetworkAffectedDevicePorts`: OperationResponseVlanForPortDisplayGridVOVlanNetworkAffectingDeviceDetailVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetGridVlanNetworkAffectedDevicePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 
**vlan** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVlanNetworkAffectedDevicePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





### Return type

[**OperationResponseVlanForPortDisplayGridVOVlanNetworkAffectingDeviceDetailVO**](OperationResponseVlanForPortDisplayGridVOVlanNetworkAffectingDeviceDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVlans

> OperationResponseGridVOVlanLanNetworkForBtachDeleteVO GetGridVlans(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get the networks whose purpose is vlan



### Example

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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetGridVlans(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetGridVlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVlans`: OperationResponseGridVOVlanLanNetworkForBtachDeleteVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetGridVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 


 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseGridVOVlanLanNetworkForBtachDeleteVO**](OperationResponseGridVOVlanLanNetworkForBtachDeleteVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaceLanNetwork

> OperationResponseResponseDataVOLanNetworkQueryOpenApiVO GetInterfaceLanNetwork(ctx, omadacId, siteId).Type_(type_).Execute()

Get all \"single\"/\"multi\" interface lan network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	type_ := int32(56) // int32 | Interface vlan type. When \"type\" is 0, return \"single\" interface lan network, else return \"single\"/\"multi\" interface lan network (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetInterfaceLanNetwork(context.Background(), omadacId, siteId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetInterfaceLanNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterfaceLanNetwork`: OperationResponseResponseDataVOLanNetworkQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetInterfaceLanNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceLanNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **int32** | Interface vlan type. When \&quot;type\&quot; is 0, return \&quot;single\&quot; interface lan network, else return \&quot;single\&quot;/\&quot;multi\&quot; interface lan network | 

### Return type

[**OperationResponseResponseDataVOLanNetworkQueryOpenApiVO**](OperationResponseResponseDataVOLanNetworkQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaceLanNetworkV2

> OperationResponseResponseDataVOLanNetworkQueryOpenApiV2VO GetInterfaceLanNetworkV2(ctx, omadacId, siteId).Type_(type_).Execute()

Get all \"single\"/\"multi\" interface lan network V2



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	type_ := int32(56) // int32 | Interface vlan type. When \"type\" is 0, return \"single\" interface lan network, else return \"single\"/\"multi\" interface lan network (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetInterfaceLanNetworkV2(context.Background(), omadacId, siteId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetInterfaceLanNetworkV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterfaceLanNetworkV2`: OperationResponseResponseDataVOLanNetworkQueryOpenApiV2VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetInterfaceLanNetworkV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceLanNetworkV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **int32** | Interface vlan type. When \&quot;type\&quot; is 0, return \&quot;single\&quot; interface lan network, else return \&quot;single\&quot;/\&quot;multi\&quot; interface lan network | 

### Return type

[**OperationResponseResponseDataVOLanNetworkQueryOpenApiV2VO**](OperationResponseResponseDataVOLanNetworkQueryOpenApiV2VO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternet

> OperationResponseInternetOpenApiVO GetInternet(ctx, omadacId, siteId).Execute()

Get internet info



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetInternet(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetInternet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternet`: OperationResponseInternetOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetInternet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseInternetOpenApiVO**](OperationResponseInternetOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternetBasicPortInfo1

> OperationResponseInternetBaseInfoOpenApiVO GetInternetBasicPortInfo1(ctx, omadacId, siteId).Execute()

Get internet basic info



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetInternetBasicPortInfo1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetInternetBasicPortInfo1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternetBasicPortInfo1`: OperationResponseInternetBaseInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetInternetBasicPortInfo1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternetBasicPortInfo1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseInternetBaseInfoOpenApiVO**](OperationResponseInternetBaseInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternetLoadBalance1

> OperationResponseWanLoadBalanceOpenApiVO GetInternetLoadBalance1(ctx, omadacId, siteId).Execute()

Get internet load balance config



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetInternetLoadBalance1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetInternetLoadBalance1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternetLoadBalance1`: OperationResponseWanLoadBalanceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetInternetLoadBalance1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternetLoadBalance1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWanLoadBalanceOpenApiVO**](OperationResponseWanLoadBalanceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpptInternet

> OperationResponseInternetOpenApiVO GetIpptInternet(ctx, omadacId, siteId, gatewayMac).Execute()

Get internet info by mac



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetIpptInternet(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetIpptInternet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpptInternet`: OperationResponseInternetOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetIpptInternet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpptInternetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseInternetOpenApiVO**](OperationResponseInternetOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIspScanResult

> OperationResponseIspScanResultOpenApiVO GetIspScanResult(ctx, omadacId, siteId, portUuid).Execute()

Get Isp scan result



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	portUuid := "portUuid_example" // string | Port Uuid to be scanned

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetIspScanResult(context.Background(), omadacId, siteId, portUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetIspScanResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIspScanResult`: OperationResponseIspScanResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetIspScanResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portUuid** | **string** | Port Uuid to be scanned | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIspScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseIspScanResultOpenApiVO**](OperationResponseIspScanResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanNetwork

> OperationResponseLanNetworkQueryOpenApiV3VO GetLanNetwork(ctx, omadacId, siteId, networkId).Execute()

Get LAN network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetLanNetwork(context.Background(), omadacId, siteId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetLanNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanNetwork`: OperationResponseLanNetworkQueryOpenApiV3VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetLanNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseLanNetworkQueryOpenApiV3VO**](OperationResponseLanNetworkQueryOpenApiV3VO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanNetworkList

> OperationResponseLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO GetLanNetworkList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get LAN network list



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetLanNetworkList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetLanNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanNetworkList`: OperationResponseLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetLanNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO**](OperationResponseLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanNetworkListV2

> OperationResponseLanNetworkOpenApiV2GridVOLanNetworkQueryOpenApiV2VO GetLanNetworkListV2(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get LAN network list



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetLanNetworkListV2(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetLanNetworkListV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanNetworkListV2`: OperationResponseLanNetworkOpenApiV2GridVOLanNetworkQueryOpenApiV2VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetLanNetworkListV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanNetworkListV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseLanNetworkOpenApiV2GridVOLanNetworkQueryOpenApiV2VO**](OperationResponseLanNetworkOpenApiV2GridVOLanNetworkQueryOpenApiV2VO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanNetworkListV3

> OperationResponseLanNetworkOpenApiV3GridVOLanNetworkQueryOpenApiV3VO GetLanNetworkListV3(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get LAN network list V3



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetLanNetworkListV3(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetLanNetworkListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanNetworkListV3`: OperationResponseLanNetworkOpenApiV3GridVOLanNetworkQueryOpenApiV3VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetLanNetworkListV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanNetworkListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseLanNetworkOpenApiV3GridVOLanNetworkQueryOpenApiV3VO**](OperationResponseLanNetworkOpenApiV3GridVOLanNetworkQueryOpenApiV3VO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanProfileList

> OperationResponseGridVOLanProfileOpenApiVO GetLanProfileList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get LAN profile list



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetLanProfileList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetLanProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanProfileList`: OperationResponseGridVOLanProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetLanProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOLanProfileOpenApiVO**](OperationResponseGridVOLanProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalanceWeightStatus

> OperationResponse GetLoadBalanceWeightStatus(ctx, omadacId, siteId).Execute()

Check Wan Loadbalance status



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetLoadBalanceWeightStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetLoadBalanceWeightStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalanceWeightStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetLoadBalanceWeightStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalanceWeightStatusRequest struct via the builder pattern


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


## GetLocationAndIspInfo

> OperationResponseLocationAndIspInfoOpenApiVO GetLocationAndIspInfo(ctx, omadacId, siteId).Execute()

Get location and ISP info



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetLocationAndIspInfo(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetLocationAndIspInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocationAndIspInfo`: OperationResponseLocationAndIspInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetLocationAndIspInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationAndIspInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseLocationAndIspInfoOpenApiVO**](OperationResponseLocationAndIspInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLteWanPortsConfig

> OperationResponseLteWanSettingOpenApiVO GetLteWanPortsConfig(ctx, omadacId, siteId).Execute()

Get internet lte wan ports config



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetLteWanPortsConfig(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetLteWanPortsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLteWanPortsConfig`: OperationResponseLteWanSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetLteWanPortsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLteWanPortsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseLteWanSettingOpenApiVO**](OperationResponseLteWanSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworksWithServersForReservation

> OperationResponseDhcpServerInfoUnderNetworkVO GetNetworksWithServersForReservation(ctx, omadacId, siteId).Execute()

Get all networks that can reserve IP addresses and their corresponding servers



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetNetworksWithServersForReservation(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetNetworksWithServersForReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworksWithServersForReservation`: OperationResponseDhcpServerInfoUnderNetworkVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetNetworksWithServersForReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksWithServersForReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDhcpServerInfoUnderNetworkVO**](OperationResponseDhcpServerInfoUnderNetworkVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswForGivenLanNetworkIds

> OperationResponseGridVOOswDataVO GetOswForGivenLanNetworkIds(ctx, omadacId, siteId).QueryOswDateByNetworkVO(queryOswDateByNetworkVO).Execute()

Get LAN network list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	queryOswDateByNetworkVO := *openapiclient.NewQueryOswDateByNetworkVO() // QueryOswDateByNetworkVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetOswForGivenLanNetworkIds(context.Background(), omadacId, siteId).QueryOswDateByNetworkVO(queryOswDateByNetworkVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetOswForGivenLanNetworkIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswForGivenLanNetworkIds`: OperationResponseGridVOOswDataVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetOswForGivenLanNetworkIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswForGivenLanNetworkIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryOswDateByNetworkVO** | [**QueryOswDateByNetworkVO**](QueryOswDateByNetworkVO.md) |  | 

### Return type

[**OperationResponseGridVOOswDataVO**](OperationResponseGridVOOswDataVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswLanProfileList

> OperationResponseGridVOLanProfileInfoOpenApiVO GetOswLanProfileList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get switch profile list



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetOswLanProfileList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetOswLanProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswLanProfileList`: OperationResponseGridVOLanProfileInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetOswLanProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswLanProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOLanProfileInfoOpenApiVO**](OperationResponseGridVOLanProfileInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswPortLanNetInfo

> GridVOOswProfileOpenApiVO GetOswPortLanNetInfo(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get switch port profile info



### Example

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
	resp, r, err := apiClient.WiredNetworkAPI.GetOswPortLanNetInfo(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetOswPortLanNetInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswPortLanNetInfo`: GridVOOswProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetOswPortLanNetInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswPortLanNetInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**GridVOOswProfileOpenApiVO**](GridVOOswProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectDevicePortsInfo

> OperationResponseListSelectDeviceForVlanVO GetSelectDevicePortsInfo(ctx, omadacId, siteId).SelectMacsWithVlanVO(selectMacsWithVlanVO).Execute()

Get the port information of the selected devices when creating network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	selectMacsWithVlanVO := *openapiclient.NewSelectMacsWithVlanVO([]string{"Macs_example"}, int32(123)) // SelectMacsWithVlanVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetSelectDevicePortsInfo(context.Background(), omadacId, siteId).SelectMacsWithVlanVO(selectMacsWithVlanVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetSelectDevicePortsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectDevicePortsInfo`: OperationResponseListSelectDeviceForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetSelectDevicePortsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectDevicePortsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectMacsWithVlanVO** | [**SelectMacsWithVlanVO**](SelectMacsWithVlanVO.md) |  | 

### Return type

[**OperationResponseListSelectDeviceForVlanVO**](OperationResponseListSelectDeviceForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectDevicePortsInfoWhenModify

> OperationResponseListSelectDeviceForVlanVO GetSelectDevicePortsInfoWhenModify(ctx, omadacId, siteId, networkId).SelectMacsWithVlanVO(selectMacsWithVlanVO).Execute()

Get the port information of the selected devices when modifying the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	selectMacsWithVlanVO := *openapiclient.NewSelectMacsWithVlanVO([]string{"Macs_example"}, int32(123)) // SelectMacsWithVlanVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetSelectDevicePortsInfoWhenModify(context.Background(), omadacId, siteId, networkId).SelectMacsWithVlanVO(selectMacsWithVlanVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetSelectDevicePortsInfoWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectDevicePortsInfoWhenModify`: OperationResponseListSelectDeviceForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetSelectDevicePortsInfoWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectDevicePortsInfoWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **selectMacsWithVlanVO** | [**SelectMacsWithVlanVO**](SelectMacsWithVlanVO.md) |  | 

### Return type

[**OperationResponseListSelectDeviceForVlanVO**](OperationResponseListSelectDeviceForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectStackPortsInfo

> OperationResponseListSelectStackForVlanVO GetSelectStackPortsInfo(ctx, omadacId, siteId).SelectStacksWithVlanVO(selectStacksWithVlanVO).Execute()

Get the port information of the selected stacks when creating network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	selectStacksWithVlanVO := *openapiclient.NewSelectStacksWithVlanVO([]string{"StackIds_example"}, int32(123)) // SelectStacksWithVlanVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetSelectStackPortsInfo(context.Background(), omadacId, siteId).SelectStacksWithVlanVO(selectStacksWithVlanVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetSelectStackPortsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectStackPortsInfo`: OperationResponseListSelectStackForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetSelectStackPortsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectStackPortsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectStacksWithVlanVO** | [**SelectStacksWithVlanVO**](SelectStacksWithVlanVO.md) |  | 

### Return type

[**OperationResponseListSelectStackForVlanVO**](OperationResponseListSelectStackForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectStackPortsInfoWhenModify

> OperationResponseListSelectStackForVlanVO GetSelectStackPortsInfoWhenModify(ctx, omadacId, siteId, networkId).SelectStacksWithVlanVO(selectStacksWithVlanVO).Execute()

Get the port information of the selected stacks when modifying the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	selectStacksWithVlanVO := *openapiclient.NewSelectStacksWithVlanVO([]string{"StackIds_example"}, int32(123)) // SelectStacksWithVlanVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetSelectStackPortsInfoWhenModify(context.Background(), omadacId, siteId, networkId).SelectStacksWithVlanVO(selectStacksWithVlanVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetSelectStackPortsInfoWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectStackPortsInfoWhenModify`: OperationResponseListSelectStackForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetSelectStackPortsInfoWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectStackPortsInfoWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **selectStacksWithVlanVO** | [**SelectStacksWithVlanVO**](SelectStacksWithVlanVO.md) |  | 

### Return type

[**OperationResponseListSelectStackForVlanVO**](OperationResponseListSelectStackForVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardBandScanResult

> OperationResponseBandScanResultOpenApiVO GetSimCardBandScanResult(ctx, omadacId, siteId, portUuid, simCard).Execute()

Obtain the bandScan result of the corresponding SIM card



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	portUuid := "portUuid_example" // string | The ID of the port.
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetSimCardBandScanResult(context.Background(), omadacId, siteId, portUuid, simCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetSimCardBandScanResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardBandScanResult`: OperationResponseBandScanResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetSimCardBandScanResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portUuid** | **string** | The ID of the port. | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardBandScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseBandScanResultOpenApiVO**](OperationResponseBandScanResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardIspScanResult

> OperationResponseIspScanResultOpenApiVO GetSimCardIspScanResult(ctx, omadacId, siteId, portUuid, simCard).Execute()

Obtain the ispScan result of the corresponding SIM card



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	portUuid := "portUuid_example" // string | The ID of the port.
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetSimCardIspScanResult(context.Background(), omadacId, siteId, portUuid, simCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetSimCardIspScanResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardIspScanResult`: OperationResponseIspScanResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetSimCardIspScanResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portUuid** | **string** | The ID of the port. | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimCardIspScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseIspScanResultOpenApiVO**](OperationResponseIspScanResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportInfo

> OperationResponseSupportOsgModelOpenApiAppVO GetSupportInfo(ctx, omadacId, siteId).Execute()

Get supported gateway model list for pre-configuration



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetSupportInfo(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetSupportInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportInfo`: OperationResponseSupportOsgModelOpenApiAppVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetSupportInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSupportOsgModelOpenApiAppVO**](OperationResponseSupportOsgModelOpenApiAppVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUseLanProfileES

> OperationResponseGridVOUseProfileOswOpenApiVO GetUseLanProfileES(ctx, omadacId, siteId, profileId).Page(page).PageSize(pageSize).Execute()

Query Agile Series switches that used the profile



### Example

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
	profileId := "profileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetUseLanProfileES(context.Background(), omadacId, siteId, profileId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetUseLanProfileES``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUseLanProfileES`: OperationResponseGridVOUseProfileOswOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetUseLanProfileES`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUseLanProfileESRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 




### Return type

[**OperationResponseGridVOUseProfileOswOpenApiVO**](OperationResponseGridVOUseProfileOswOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUseNativeNetworkOsw

> OperationResponseGridVOUseProfileOswOpenApiVO GetUseNativeNetworkOsw(ctx, omadacId, siteId).QueryUseNativeNetworkOswOpenApiVO(queryUseNativeNetworkOswOpenApiVO).Execute()

Query switches that used the Native Network's profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	queryUseNativeNetworkOswOpenApiVO := *openapiclient.NewQueryUseNativeNetworkOswOpenApiVO(int32(123), int32(123)) // QueryUseNativeNetworkOswOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetUseNativeNetworkOsw(context.Background(), omadacId, siteId).QueryUseNativeNetworkOswOpenApiVO(queryUseNativeNetworkOswOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetUseNativeNetworkOsw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUseNativeNetworkOsw`: OperationResponseGridVOUseProfileOswOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetUseNativeNetworkOsw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUseNativeNetworkOswRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryUseNativeNetworkOswOpenApiVO** | [**QueryUseNativeNetworkOswOpenApiVO**](QueryUseNativeNetworkOswOpenApiVO.md) |  | 

### Return type

[**OperationResponseGridVOUseProfileOswOpenApiVO**](OperationResponseGridVOUseProfileOswOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUseNativeNetworkOswV2

> OperationResponseGridVOUseProfileOswOpenApiVO GetUseNativeNetworkOswV2(ctx, omadacId, siteId).QueryUseNativeNetworkOswV2OpenApiVO(queryUseNativeNetworkOswV2OpenApiVO).Execute()

Query switches that used the Native Network's profile V2



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	queryUseNativeNetworkOswV2OpenApiVO := *openapiclient.NewQueryUseNativeNetworkOswV2OpenApiVO(int32(123), int32(123)) // QueryUseNativeNetworkOswV2OpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetUseNativeNetworkOswV2(context.Background(), omadacId, siteId).QueryUseNativeNetworkOswV2OpenApiVO(queryUseNativeNetworkOswV2OpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetUseNativeNetworkOswV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUseNativeNetworkOswV2`: OperationResponseGridVOUseProfileOswOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetUseNativeNetworkOswV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUseNativeNetworkOswV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryUseNativeNetworkOswV2OpenApiVO** | [**QueryUseNativeNetworkOswV2OpenApiVO**](QueryUseNativeNetworkOswV2OpenApiVO.md) |  | 

### Return type

[**OperationResponseGridVOUseProfileOswOpenApiVO**](OperationResponseGridVOUseProfileOswOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlanNetworkAffectedDevice1

> OperationResponseListVlanNetworkAffectingDeviceVO GetVlanNetworkAffectedDevice1(ctx, omadacId, siteId, networkId, vlan).Execute()

Get devices list that use the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	vlan := "vlan_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetVlanNetworkAffectedDevice1(context.Background(), omadacId, siteId, networkId, vlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetVlanNetworkAffectedDevice1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVlanNetworkAffectedDevice1`: OperationResponseListVlanNetworkAffectingDeviceVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetVlanNetworkAffectedDevice1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 
**vlan** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVlanNetworkAffectedDevice1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseListVlanNetworkAffectingDeviceVO**](OperationResponseListVlanNetworkAffectingDeviceVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlanNetworkAffectedSsid1

> OperationResponseGridVOVlanNetworkAffectingSsidVO GetVlanNetworkAffectedSsid1(ctx, omadacId, siteId, networkId, vlan).Page(page).PageSize(pageSize).Execute()

Get grid ssid list that use the vlan



### Example

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
	networkId := "networkId_example" // string | Network ID
	vlan := "vlan_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetVlanNetworkAffectedSsid1(context.Background(), omadacId, siteId, networkId, vlan).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetVlanNetworkAffectedSsid1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVlanNetworkAffectedSsid1`: OperationResponseGridVOVlanNetworkAffectingSsidVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetVlanNetworkAffectedSsid1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 
**vlan** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVlanNetworkAffectedSsid1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





### Return type

[**OperationResponseGridVOVlanNetworkAffectingSsidVO**](OperationResponseGridVOVlanNetworkAffectingSsidVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWanPortsConfig

> OperationResponseWanSettingOpenApiVO GetWanPortsConfig(ctx, omadacId, siteId).Execute()

Get internet ports config



### Example

```go
package main

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
	resp, r, err := apiClient.WiredNetworkAPI.GetWanPortsConfig(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetWanPortsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWanPortsConfig`: OperationResponseWanSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetWanPortsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanPortsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWanSettingOpenApiVO**](OperationResponseWanSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWanPortsMaxSpeed1

> OperationResponseWanMaxSpeedOpenApiVO GetWanPortsMaxSpeed1(ctx, omadacId, siteId).WanId(wanId).Execute()

Get wan ports max speed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	wanId := "wanId_example" // string | WanId, can pass multiple wan ports using comma separator like '{portId1},{portId2}'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.GetWanPortsMaxSpeed1(context.Background(), omadacId, siteId).WanId(wanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.GetWanPortsMaxSpeed1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWanPortsMaxSpeed1`: OperationResponseWanMaxSpeedOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.GetWanPortsMaxSpeed1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanPortsMaxSpeed1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanId** | **string** | WanId, can pass multiple wan ports using comma separator like &#39;{portId1},{portId2}&#39;. | 

### Return type

[**OperationResponseWanMaxSpeedOpenApiVO**](OperationResponseWanMaxSpeedOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGatewayModel

> OperationResponseWithoutResult ModifyGatewayModel(ctx, omadacId, siteId).OsgModelOpenApiVO(osgModelOpenApiVO).Execute()

Modify gateway mode for pre-configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	osgModelOpenApiVO := *openapiclient.NewOsgModelOpenApiVO(int32(123)) // OsgModelOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyGatewayModel(context.Background(), omadacId, siteId).OsgModelOpenApiVO(osgModelOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyGatewayModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGatewayModel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyGatewayModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyGatewayModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **osgModelOpenApiVO** | [**OsgModelOpenApiVO**](OsgModelOpenApiVO.md) |  | 

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


## ModifyInternetBasicPortInfo1

> OperationResponseWithoutResult ModifyInternetBasicPortInfo1(ctx, omadacId, siteId).InternetBaseConfigOpenApiVO(internetBaseConfigOpenApiVO).Execute()

Modify the internet basic configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	internetBaseConfigOpenApiVO := *openapiclient.NewInternetBaseConfigOpenApiVO(false, []string{"WanPortList_example"}) // InternetBaseConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyInternetBasicPortInfo1(context.Background(), omadacId, siteId).InternetBaseConfigOpenApiVO(internetBaseConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyInternetBasicPortInfo1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyInternetBasicPortInfo1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyInternetBasicPortInfo1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyInternetBasicPortInfo1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **internetBaseConfigOpenApiVO** | [**InternetBaseConfigOpenApiVO**](InternetBaseConfigOpenApiVO.md) |  | 

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


## ModifyInternetLoadBalance1

> OperationResponseWithoutResult ModifyInternetLoadBalance1(ctx, omadacId, siteId).WanLoadBalanceOpenApiVO(wanLoadBalanceOpenApiVO).Execute()

Modify internet load balance config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	wanLoadBalanceOpenApiVO := *openapiclient.NewWanLoadBalanceOpenApiVO(false, false, []int32{int32(123)}) // WanLoadBalanceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyInternetLoadBalance1(context.Background(), omadacId, siteId).WanLoadBalanceOpenApiVO(wanLoadBalanceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyInternetLoadBalance1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyInternetLoadBalance1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyInternetLoadBalance1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyInternetLoadBalance1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanLoadBalanceOpenApiVO** | [**WanLoadBalanceOpenApiVO**](WanLoadBalanceOpenApiVO.md) |  | 

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


## ModifyIpptLteWanPortConfig

> OperationResponseWithoutResult ModifyIpptLteWanPortConfig(ctx, omadacId, siteId, gatewayMac).LteWanPortSettingConfigOpenApiV2VO(lteWanPortSettingConfigOpenApiV2VO).Execute()

Modify internet lte port config for ippt



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	lteWanPortSettingConfigOpenApiV2VO := *openapiclient.NewLteWanPortSettingConfigOpenApiV2VO(*openapiclient.NewDialupSettingOpenApiV2VO(int32(123), false, int32(123), int32(123)), false, "PortUuid_example") // LteWanPortSettingConfigOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyIpptLteWanPortConfig(context.Background(), omadacId, siteId, gatewayMac).LteWanPortSettingConfigOpenApiV2VO(lteWanPortSettingConfigOpenApiV2VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyIpptLteWanPortConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpptLteWanPortConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyIpptLteWanPortConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpptLteWanPortConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lteWanPortSettingConfigOpenApiV2VO** | [**LteWanPortSettingConfigOpenApiV2VO**](LteWanPortSettingConfigOpenApiV2VO.md) |  | 

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


## ModifyIpptWanMode

> OperationResponse ModifyIpptWanMode(ctx, omadacId, siteId, gatewayMac).IpptWanModeOpenApiVO(ipptWanModeOpenApiVO).Execute()

Modify internet wan mode



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	ipptWanModeOpenApiVO := *openapiclient.NewIpptWanModeOpenApiVO(false) // IpptWanModeOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyIpptWanMode(context.Background(), omadacId, siteId, gatewayMac).IpptWanModeOpenApiVO(ipptWanModeOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyIpptWanMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpptWanMode`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyIpptWanMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpptWanModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ipptWanModeOpenApiVO** | [**IpptWanModeOpenApiVO**](IpptWanModeOpenApiVO.md) |  | 

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


## ModifyLanDns

> OperationResponseWithoutResult ModifyLanDns(ctx, omadacId, siteId, dnsId).LanDnsOpenApiVO(lanDnsOpenApiVO).Execute()

Modify an existing LAN DNS rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	dnsId := "dnsId_example" // string | LAN DNS ID
	lanDnsOpenApiVO := *openapiclient.NewLanDnsOpenApiVO("Domain_example", false, "Name_example", int32(123)) // LanDnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyLanDns(context.Background(), omadacId, siteId, dnsId).LanDnsOpenApiVO(lanDnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyLanDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanDns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyLanDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**dnsId** | **string** | LAN DNS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanDnsOpenApiVO** | [**LanDnsOpenApiVO**](LanDnsOpenApiVO.md) |  | 

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


## ModifyLanNetwork

> OperationResponseWithoutResult ModifyLanNetwork(ctx, omadacId, siteId, networkId).LanNetworkOpenApiVO(lanNetworkOpenApiVO).Execute()

Modify LAN network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	lanNetworkOpenApiVO := *openapiclient.NewLanNetworkOpenApiVO(false, "Name_example", int32(123)) // LanNetworkOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyLanNetwork(context.Background(), omadacId, siteId, networkId).LanNetworkOpenApiVO(lanNetworkOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyLanNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanNetwork`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyLanNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkOpenApiVO** | [**LanNetworkOpenApiVO**](LanNetworkOpenApiVO.md) |  | 

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


## ModifyLanNetworkV2

> OperationResponseWithoutResult ModifyLanNetworkV2(ctx, omadacId, siteId, networkId).LanNetworkOpenApiV2VO(lanNetworkOpenApiV2VO).Execute()

Modify LAN network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | Network ID
	lanNetworkOpenApiV2VO := *openapiclient.NewLanNetworkOpenApiV2VO(false, "Name_example", int32(123)) // LanNetworkOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyLanNetworkV2(context.Background(), omadacId, siteId, networkId).LanNetworkOpenApiV2VO(lanNetworkOpenApiV2VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyLanNetworkV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanNetworkV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyLanNetworkV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanNetworkV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkOpenApiV2VO** | [**LanNetworkOpenApiV2VO**](LanNetworkOpenApiV2VO.md) |  | 

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


## ModifyLanProfile

> OperationResponseWithoutResult ModifyLanProfile(ctx, omadacId, siteId, profileId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()

Modify a LAN profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	profileId := "profileId_example" // string | lanProfileId
	lanProfileConfigOpenApiVO := *openapiclient.NewLanProfileConfigOpenApiVO(int32(123), int32(123), false, false, "Name_example", "NativeNetworkId_example", int32(123), false, false) // LanProfileConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyLanProfile(context.Background(), omadacId, siteId, profileId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyLanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyLanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | lanProfileId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanProfileConfigOpenApiVO** | [**LanProfileConfigOpenApiVO**](LanProfileConfigOpenApiVO.md) |  | 

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


## ModifyLteWanPortConfig

> OperationResponseWithoutResult ModifyLteWanPortConfig(ctx, omadacId, siteId).LteWanPortSettingConfigOpenApiVO(lteWanPortSettingConfigOpenApiVO).Execute()

Modify internet lte port config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	lteWanPortSettingConfigOpenApiVO := *openapiclient.NewLteWanPortSettingConfigOpenApiVO(false, *openapiclient.NewDialupSettingOpenApiVO(int32(123), int32(123)), false, int32(123), "PortId_example") // LteWanPortSettingConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyLteWanPortConfig(context.Background(), omadacId, siteId).LteWanPortSettingConfigOpenApiVO(lteWanPortSettingConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyLteWanPortConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLteWanPortConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyLteWanPortConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLteWanPortConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lteWanPortSettingConfigOpenApiVO** | [**LteWanPortSettingConfigOpenApiVO**](LteWanPortSettingConfigOpenApiVO.md) |  | 

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


## ModifyOswLanProfile

> OperationResponseWithoutResult ModifyOswLanProfile(ctx, omadacId, siteId, profileId).LanProfileSettingOpenApiVO(lanProfileSettingOpenApiVO).Execute()

Modify a switch profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	profileId := "profileId_example" // string | Switch profile ID
	lanProfileSettingOpenApiVO := *openapiclient.NewLanProfileSettingOpenApiVO(int32(123), int32(123), false, false, "Name_example", int32(123), false, false) // LanProfileSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyOswLanProfile(context.Background(), omadacId, siteId, profileId).LanProfileSettingOpenApiVO(lanProfileSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyOswLanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswLanProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyOswLanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | Switch profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswLanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanProfileSettingOpenApiVO** | [**LanProfileSettingOpenApiVO**](LanProfileSettingOpenApiVO.md) |  | 

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


## ModifyVirtualWan

> OperationResponseWithoutResult ModifyVirtualWan(ctx, omadacId, siteId, virtualWanId).VirtualWanConfigOpenApiVO(virtualWanConfigOpenApiVO).Execute()

Modify virtual WAN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	virtualWanId := "virtualWanId_example" // string | Virtual WAN ID
	virtualWanConfigOpenApiVO := *openapiclient.NewVirtualWanConfigOpenApiVO("Name_example", "PhysicalWanId_example", *openapiclient.NewVirtualWanIpv4SettingConfigOpenApiVO("Proto_example", int32(123))) // VirtualWanConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyVirtualWan(context.Background(), omadacId, siteId, virtualWanId).VirtualWanConfigOpenApiVO(virtualWanConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyVirtualWan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVirtualWan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyVirtualWan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**virtualWanId** | **string** | Virtual WAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVirtualWanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **virtualWanConfigOpenApiVO** | [**VirtualWanConfigOpenApiVO**](VirtualWanConfigOpenApiVO.md) |  | 

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


## ModifyVirtualWanStatus

> OperationResponseWithoutResult ModifyVirtualWanStatus(ctx, omadacId, siteId, virtualWanId).VirtualWanStatusOpenApiVO(virtualWanStatusOpenApiVO).Execute()

Modify virtual WAN status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	virtualWanId := "virtualWanId_example" // string | Virtual WAN ID
	virtualWanStatusOpenApiVO := *openapiclient.NewVirtualWanStatusOpenApiVO(false) // VirtualWanStatusOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyVirtualWanStatus(context.Background(), omadacId, siteId, virtualWanId).VirtualWanStatusOpenApiVO(virtualWanStatusOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyVirtualWanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVirtualWanStatus`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyVirtualWanStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**virtualWanId** | **string** | Virtual WAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVirtualWanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **virtualWanStatusOpenApiVO** | [**VirtualWanStatusOpenApiVO**](VirtualWanStatusOpenApiVO.md) |  | 

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


## ModifyWanPortSettings

> OperationResponseWithoutResult ModifyWanPortSettings(ctx, omadacId, siteId).InternetPortOpenApiVO(internetPortOpenApiVO).Execute()

Modify Wan Port Settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	internetPortOpenApiVO := *openapiclient.NewInternetPortOpenApiVO() // InternetPortOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyWanPortSettings(context.Background(), omadacId, siteId).InternetPortOpenApiVO(internetPortOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyWanPortSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWanPortSettings`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyWanPortSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWanPortSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **internetPortOpenApiVO** | [**InternetPortOpenApiVO**](InternetPortOpenApiVO.md) |  | 

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


## ModifyWanPorts

> OperationResponse ModifyWanPorts(ctx, omadacId, siteId).WanPortsOpenApiVO(wanPortsOpenApiVO).Execute()

Modify Wan Ports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	wanPortsOpenApiVO := *openapiclient.NewWanPortsOpenApiVO(false) // WanPortsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyWanPorts(context.Background(), omadacId, siteId).WanPortsOpenApiVO(wanPortsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyWanPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWanPorts`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyWanPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWanPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanPortsOpenApiVO** | [**WanPortsOpenApiVO**](WanPortsOpenApiVO.md) |  | 

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


## ModifyWanPortsConfig

> OperationResponseWithoutResult ModifyWanPortsConfig(ctx, omadacId, siteId).WanSettingConfigOpenApiVO(wanSettingConfigOpenApiVO).Execute()

Modify internet ports config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	wanSettingConfigOpenApiVO := *openapiclient.NewWanSettingConfigOpenApiVO() // WanSettingConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.ModifyWanPortsConfig(context.Background(), omadacId, siteId).WanSettingConfigOpenApiVO(wanSettingConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.ModifyWanPortsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWanPortsConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.ModifyWanPortsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWanPortsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanSettingConfigOpenApiVO** | [**WanSettingConfigOpenApiVO**](WanSettingConfigOpenApiVO.md) |  | 

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


## NetworkMapping

> OperationResponseInternetOpenApiVO NetworkMapping(ctx, omadacId, siteId).NetworkMappingVO(networkMappingVO).Execute()

Network mapping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkMappingVO := *openapiclient.NewNetworkMappingVO() // NetworkMappingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.NetworkMapping(context.Background(), omadacId, siteId).NetworkMappingVO(networkMappingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.NetworkMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkMapping`: OperationResponseInternetOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.NetworkMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkMappingVO** | [**NetworkMappingVO**](NetworkMappingVO.md) |  | 

### Return type

[**OperationResponseInternetOpenApiVO**](OperationResponseInternetOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeedTest

> OperationResponseWithoutResult SpeedTest(ctx, omadacId, siteId).InternetOpenApiVO(internetOpenApiVO).Execute()

SpeedTest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	internetOpenApiVO := *openapiclient.NewInternetOpenApiVO("SiteId_example") // InternetOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.SpeedTest(context.Background(), omadacId, siteId).InternetOpenApiVO(internetOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.SpeedTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeedTest`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.SpeedTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpeedTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **internetOpenApiVO** | [**InternetOpenApiVO**](InternetOpenApiVO.md) |  | 

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


## StartBandScan

> OperationResponseWithoutResult StartBandScan(ctx, omadacId, siteId).BandScanStart(bandScanStart).Execute()

BandScan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	bandScanStart := *openapiclient.NewBandScanStart("PortUuid_example") // BandScanStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.StartBandScan(context.Background(), omadacId, siteId).BandScanStart(bandScanStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.StartBandScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartBandScan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.StartBandScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartBandScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bandScanStart** | [**BandScanStart**](BandScanStart.md) |  | 

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


## StartIspScan

> OperationResponseWithoutResult StartIspScan(ctx, omadacId, siteId).IspScanStartOpenApiVO(ispScanStartOpenApiVO).Execute()

IspScan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	ispScanStartOpenApiVO := *openapiclient.NewIspScanStartOpenApiVO("PortUuid_example") // IspScanStartOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkAPI.StartIspScan(context.Background(), omadacId, siteId).IspScanStartOpenApiVO(ispScanStartOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkAPI.StartIspScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartIspScan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkAPI.StartIspScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartIspScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ispScanStartOpenApiVO** | [**IspScanStartOpenApiVO**](IspScanStartOpenApiVO.md) |  | 

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


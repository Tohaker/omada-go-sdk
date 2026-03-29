# \WiredNetworkTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchIsolateLanNetwork**](WiredNetworkTemplateAPI.md#BatchIsolateLanNetwork) | **Post** /openapi/v2/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks/batch-isolate | Batch isolate network
[**CheckNetworkTemplateParamWhenCreate**](WiredNetworkTemplateAPI.md#CheckNetworkTemplateParamWhenCreate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/param-check | Check basic parameters when creating network template
[**CheckNetworkTemplateParamWhenModify**](WiredNetworkTemplateAPI.md#CheckNetworkTemplateParamWhenModify) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/param-check | Check basic parameters when modifying the network
[**CheckTemplateParamAndGetPreConfig**](WiredNetworkTemplateAPI.md#CheckTemplateParamAndGetPreConfig) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/check | Check all parameters and generate configuration when creating network template
[**CheckTemplateParamAndGetPreConfigWhenModify**](WiredNetworkTemplateAPI.md#CheckTemplateParamAndGetPreConfigWhenModify) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/check | Check all parameters and generate configuration when modifying the network template
[**CheckTemplatePortBindingParamWhenCreate**](WiredNetworkTemplateAPI.md#CheckTemplatePortBindingParamWhenCreate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/ports-check | Check selected ports when creating network for site template
[**CheckTemplatePortBindingParamWhenModify**](WiredNetworkTemplateAPI.md#CheckTemplatePortBindingParamWhenModify) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/ports-check | Check port binding when modifying network for site template
[**CheckVirtualWanUsed**](WiredNetworkTemplateAPI.md#CheckVirtualWanUsed) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/virtual-wans/delete-check | Check Virtual Wan Used
[**CheckWanLanStatus**](WiredNetworkTemplateAPI.md#CheckWanLanStatus) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wan-lan-status | Check template WAN-LAN status
[**ConfirmCreateVlanNetworkTemplate**](WiredNetworkTemplateAPI.md#ConfirmCreateVlanNetworkTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/confirm | Confirm create lan network template
[**ConfirmModifyVlanNetworkTemplate**](WiredNetworkTemplateAPI.md#ConfirmModifyVlanNetworkTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/confirm | Confirm modify lan network template
[**CreateLanDnsTemplate**](WiredNetworkTemplateAPI.md#CreateLanDnsTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan/dns | Create LAN Dns template
[**CreateLanNetworkTemplate**](WiredNetworkTemplateAPI.md#CreateLanNetworkTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks | Create LAN network template
[**CreateLanProfileTemplate**](WiredNetworkTemplateAPI.md#CreateLanProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles | Create new LAN profile template
[**CreateOswLanProfileTemplate**](WiredNetworkTemplateAPI.md#CreateOswLanProfileTemplate) | **Post** /openapi/v2/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles | Create new switch profile template
[**CreateVirtualWanTemplate**](WiredNetworkTemplateAPI.md#CreateVirtualWanTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/virtual-wans | Create virtual WAN template
[**CreateVlans**](WiredNetworkTemplateAPI.md#CreateVlans) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-vlans | Batch create vlans template
[**DeleteLanDnsTemplate**](WiredNetworkTemplateAPI.md#DeleteLanDnsTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan/dns/{dnsId} | Delete LAN Dns template
[**DeleteLanNetworkTemplate**](WiredNetworkTemplateAPI.md#DeleteLanNetworkTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks/{networkId} | Delete LAN network template
[**DeleteLanProfileTemplate**](WiredNetworkTemplateAPI.md#DeleteLanProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles/{profileId} | Delete an existing LAN profile template
[**DeleteOswLanProfileTemplate**](WiredNetworkTemplateAPI.md#DeleteOswLanProfileTemplate) | **Delete** /openapi/v2/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles/{profileId} | Delete an existing switch profile template
[**DeleteVirtualWan**](WiredNetworkTemplateAPI.md#DeleteVirtualWan) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/virtual-wans/{virtualWanId} | Delete virtual WAN template
[**GetAllInterfacesForBatchIsolate**](WiredNetworkTemplateAPI.md#GetAllInterfacesForBatchIsolate) | **Get** /openapi/v2/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks/isolate/interfaces | Get interface Grid
[**GetAllLanNetworksTemplate**](WiredNetworkTemplateAPI.md#GetAllLanNetworksTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks/all | Get all network templates for the omada id and site template id
[**GetAllValidVlansInSiteForSwitchOuiBasedVlan**](WiredNetworkTemplateAPI.md#GetAllValidVlansInSiteForSwitchOuiBasedVlan) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/valid-vlans | Get valid site template vlanList for switch oui based vlan
[**GetAutoEffectDevicesTemplateWhenCreate**](WiredNetworkTemplateAPI.md#GetAutoEffectDevicesTemplateWhenCreate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/auto-effect-devices | Get auto effect devices when creating network
[**GetAutoEffectDevicesTemplateWhenModify**](WiredNetworkTemplateAPI.md#GetAutoEffectDevicesTemplateWhenModify) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/auto-effect-devices | Get auto effect devices when modifying network
[**GetAutoSelectDevicesTemplateWhenModify**](WiredNetworkTemplateAPI.md#GetAutoSelectDevicesTemplateWhenModify) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/auto-select-devices | Get auto select devices when modifying network
[**GetAvailableVirtualWanTemplate**](WiredNetworkTemplateAPI.md#GetAvailableVirtualWanTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/available-virtual-wans | Query available virtual WAN list for template
[**GetAvailableWanPorts**](WiredNetworkTemplateAPI.md#GetAvailableWanPorts) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/wan-ports | Get available wan ports
[**GetGridSupportVlanNetworkDevicesBySiteTemplate**](WiredNetworkTemplateAPI.md#GetGridSupportVlanNetworkDevicesBySiteTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/devices | Get device template list that can be dhcp server
[**GetGridVirtualWanTemplate**](WiredNetworkTemplateAPI.md#GetGridVirtualWanTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/virtual-wans | Query virtual WAN list template
[**GetInterfaceLanNetworkV21**](WiredNetworkTemplateAPI.md#GetInterfaceLanNetworkV21) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks/interface | Get all \&quot;single\&quot;/\&quot;multi\&quot; interface lan network template
[**GetInternetBasicPortInfo**](WiredNetworkTemplateAPI.md#GetInternetBasicPortInfo) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/internet/basic-info | Get site template internet basic info
[**GetInternetLoadBalance**](WiredNetworkTemplateAPI.md#GetInternetLoadBalance) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/internet/load-balance | Get site template internet load balance config
[**GetInternetTemplate**](WiredNetworkTemplateAPI.md#GetInternetTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/internet | Get internet info
[**GetLanDnsTemplateList**](WiredNetworkTemplateAPI.md#GetLanDnsTemplateList) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan/dns | Get LAN Dns template list
[**GetLanNetworkTemplate**](WiredNetworkTemplateAPI.md#GetLanNetworkTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks/{networkId} | Get LAN network template
[**GetLanNetworkTemplateList**](WiredNetworkTemplateAPI.md#GetLanNetworkTemplateList) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks | Get LAN network template list
[**GetLanNetworkTemplateListV3**](WiredNetworkTemplateAPI.md#GetLanNetworkTemplateListV3) | **Get** /openapi/v3/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks | Get LAN network template list
[**GetLanProfileTemplateList**](WiredNetworkTemplateAPI.md#GetLanProfileTemplateList) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles | Get LAN profile template list
[**GetLoadBalanceWeightStatusTemplate**](WiredNetworkTemplateAPI.md#GetLoadBalanceWeightStatusTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/internet/load-balance/status | Check Wan Loadbalance status
[**GetOswLanProfileTemplateList**](WiredNetworkTemplateAPI.md#GetOswLanProfileTemplateList) | **Get** /openapi/v2/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles | Get switch profile template list
[**GetSelectDeviceTemplatePortsInfo**](WiredNetworkTemplateAPI.md#GetSelectDeviceTemplatePortsInfo) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/devices/ports | Get the port information of the selected device templates when creating network for site template
[**GetSelectDeviceTemplatePortsInfoWhenModify**](WiredNetworkTemplateAPI.md#GetSelectDeviceTemplatePortsInfoWhenModify) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/devices/ports | Get the port information of the selected device templates when modifying the network for site template
[**GetSupportPortsDeviceTemplates**](WiredNetworkTemplateAPI.md#GetSupportPortsDeviceTemplates) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/support-ports | Get grid devices templates that support ports config.
[**GetTemplateGridVlanNetworkAffectedDevicePorts**](WiredNetworkTemplateAPI.md#GetTemplateGridVlanNetworkAffectedDevicePorts) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/vlan/{vlan}/ports | Get the ports of devices template that use the network
[**GetUseLanProfileESTemplateList**](WiredNetworkTemplateAPI.md#GetUseLanProfileESTemplateList) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles/{profileId}/es | Get Use LAN profile ES template list
[**GetVlanNetworkAffectedDevice**](WiredNetworkTemplateAPI.md#GetVlanNetworkAffectedDevice) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/vlan/{vlan}/devices | Get device template list that use the network
[**GetVlanNetworkAffectedSsid**](WiredNetworkTemplateAPI.md#GetVlanNetworkAffectedSsid) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/vlan/{vlan}/ssids | Get grid ssid template list that use the vlan
[**GetWanPortsMaxSpeed**](WiredNetworkTemplateAPI.md#GetWanPortsMaxSpeed) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wan-max-speed | Get wan ports max speed
[**ModifyInternetBasicPortInfo**](WiredNetworkTemplateAPI.md#ModifyInternetBasicPortInfo) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/internet/basic-info | Modify site template internet basic configuration
[**ModifyInternetLoadBalance**](WiredNetworkTemplateAPI.md#ModifyInternetLoadBalance) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/internet/load-balance | Modify site template internet load balance config
[**ModifyLanDnsTemplate**](WiredNetworkTemplateAPI.md#ModifyLanDnsTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan/dns/{dnsId} | Modify LAN Dns template
[**ModifyLanNetworkTemplate**](WiredNetworkTemplateAPI.md#ModifyLanNetworkTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-networks/{networkId} | Modify LAN network template
[**ModifyLanProfileTemplate**](WiredNetworkTemplateAPI.md#ModifyLanProfileTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles/{profileId} | Modify a LAN profile template
[**ModifyOswLanProfileTemplate**](WiredNetworkTemplateAPI.md#ModifyOswLanProfileTemplate) | **Patch** /openapi/v2/{omadacId}/sitetemplates/{siteTemplateId}/lan-profiles/{profileId} | Modify a switch profile template
[**ModifyVirtualWanStatusTemplate**](WiredNetworkTemplateAPI.md#ModifyVirtualWanStatusTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/virtual-wans/{virtualWanId}/status | Modify virtual WAN status template
[**ModifyVirtualWanTemplate**](WiredNetworkTemplateAPI.md#ModifyVirtualWanTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/virtual-wans/{virtualWanId} | Modify virtual WAN template
[**ModifyWanPortsTemplate**](WiredNetworkTemplateAPI.md#ModifyWanPortsTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/wan-ports | Modify Wan Ports



## BatchIsolateLanNetwork

> OperationResponseWithoutResult BatchIsolateLanNetwork(ctx, omadacId, siteTemplateId).BatchIsolateInterfaceOpenApiVO(batchIsolateInterfaceOpenApiVO).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	batchIsolateInterfaceOpenApiVO := *openapiclient.NewBatchIsolateInterfaceOpenApiVO() // BatchIsolateInterfaceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.BatchIsolateLanNetwork(context.Background(), omadacId, siteTemplateId).BatchIsolateInterfaceOpenApiVO(batchIsolateInterfaceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.BatchIsolateLanNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchIsolateLanNetwork`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.BatchIsolateLanNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchIsolateLanNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchIsolateInterfaceOpenApiVO** | [**BatchIsolateInterfaceOpenApiVO**](BatchIsolateInterfaceOpenApiVO.md) |  | 

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


## CheckNetworkTemplateParamWhenCreate

> OperationResponseLanNetworkQueryOpenApiV3VO CheckNetworkTemplateParamWhenCreate(ctx, omadacId, siteTemplateId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

Check basic parameters when creating network template



### Example

```go
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
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CheckNetworkTemplateParamWhenCreate(context.Background(), omadacId, siteTemplateId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CheckNetworkTemplateParamWhenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckNetworkTemplateParamWhenCreate`: OperationResponseLanNetworkQueryOpenApiV3VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CheckNetworkTemplateParamWhenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckNetworkTemplateParamWhenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseLanNetworkQueryOpenApiV3VO**](OperationResponseLanNetworkQueryOpenApiV3VO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckNetworkTemplateParamWhenModify

> OperationResponseLanNetworkQueryOpenApiV3VO CheckNetworkTemplateParamWhenModify(ctx, omadacId, siteTemplateId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	networkId := "networkId_example" // string | Network ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CheckNetworkTemplateParamWhenModify(context.Background(), omadacId, siteTemplateId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CheckNetworkTemplateParamWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckNetworkTemplateParamWhenModify`: OperationResponseLanNetworkQueryOpenApiV3VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CheckNetworkTemplateParamWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckNetworkTemplateParamWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseLanNetworkQueryOpenApiV3VO**](OperationResponseLanNetworkQueryOpenApiV3VO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckTemplateParamAndGetPreConfig

> OperationResponseVlanPreConfigOpenApiVO CheckTemplateParamAndGetPreConfig(ctx, omadacId, siteTemplateId).CreateVlanParamOpenApiVO(createVlanParamOpenApiVO).Execute()

Check all parameters and generate configuration when creating network template



### Example

```go
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
	createVlanParamOpenApiVO := *openapiclient.NewCreateVlanParamOpenApiVO(*openapiclient.NewSelectPortBindingBriefVO(), *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example")) // CreateVlanParamOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CheckTemplateParamAndGetPreConfig(context.Background(), omadacId, siteTemplateId).CreateVlanParamOpenApiVO(createVlanParamOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CheckTemplateParamAndGetPreConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTemplateParamAndGetPreConfig`: OperationResponseVlanPreConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CheckTemplateParamAndGetPreConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTemplateParamAndGetPreConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVlanParamOpenApiVO** | [**CreateVlanParamOpenApiVO**](CreateVlanParamOpenApiVO.md) |  | 

### Return type

[**OperationResponseVlanPreConfigOpenApiVO**](OperationResponseVlanPreConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckTemplateParamAndGetPreConfigWhenModify

> OperationResponseVlanPreConfigOpenApiVO CheckTemplateParamAndGetPreConfigWhenModify(ctx, omadacId, siteTemplateId, networkId).ModifyVlanParamOpenApiVO(modifyVlanParamOpenApiVO).Execute()

Check all parameters and generate configuration when modifying the network template



### Example

```go
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
	networkId := "networkId_example" // string | Network ID
	modifyVlanParamOpenApiVO := *openapiclient.NewModifyVlanParamOpenApiVO(*openapiclient.NewSelectPortBindingBriefVO(), *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example")) // ModifyVlanParamOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CheckTemplateParamAndGetPreConfigWhenModify(context.Background(), omadacId, siteTemplateId, networkId).ModifyVlanParamOpenApiVO(modifyVlanParamOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CheckTemplateParamAndGetPreConfigWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTemplateParamAndGetPreConfigWhenModify`: OperationResponseVlanPreConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CheckTemplateParamAndGetPreConfigWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTemplateParamAndGetPreConfigWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyVlanParamOpenApiVO** | [**ModifyVlanParamOpenApiVO**](ModifyVlanParamOpenApiVO.md) |  | 

### Return type

[**OperationResponseVlanPreConfigOpenApiVO**](OperationResponseVlanPreConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckTemplatePortBindingParamWhenCreate

> OperationResponseListString CheckTemplatePortBindingParamWhenCreate(ctx, omadacId, siteTemplateId).SelectPortBindingVO(selectPortBindingVO).Execute()

Check selected ports when creating network for site template



### Example

```go
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
	selectPortBindingVO := *openapiclient.NewSelectPortBindingVO(int32(123), int32(123)) // SelectPortBindingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CheckTemplatePortBindingParamWhenCreate(context.Background(), omadacId, siteTemplateId).SelectPortBindingVO(selectPortBindingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CheckTemplatePortBindingParamWhenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTemplatePortBindingParamWhenCreate`: OperationResponseListString
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CheckTemplatePortBindingParamWhenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTemplatePortBindingParamWhenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectPortBindingVO** | [**SelectPortBindingVO**](SelectPortBindingVO.md) |  | 

### Return type

[**OperationResponseListString**](OperationResponseListString.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckTemplatePortBindingParamWhenModify

> OperationResponseListString CheckTemplatePortBindingParamWhenModify(ctx, omadacId, siteTemplateId, networkId).SelectPortBindingVO(selectPortBindingVO).Execute()

Check port binding when modifying network for site template



### Example

```go
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
	networkId := "networkId_example" // string | Network ID
	selectPortBindingVO := *openapiclient.NewSelectPortBindingVO(int32(123), int32(123)) // SelectPortBindingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CheckTemplatePortBindingParamWhenModify(context.Background(), omadacId, siteTemplateId, networkId).SelectPortBindingVO(selectPortBindingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CheckTemplatePortBindingParamWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTemplatePortBindingParamWhenModify`: OperationResponseListString
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CheckTemplatePortBindingParamWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTemplatePortBindingParamWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **selectPortBindingVO** | [**SelectPortBindingVO**](SelectPortBindingVO.md) |  | 

### Return type

[**OperationResponseListString**](OperationResponseListString.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckVirtualWanUsed

> OperationResponseVirtualWanIdUsedOpenApiVO CheckVirtualWanUsed(ctx, omadacId, siteTemplateId).VirtualWanId(virtualWanId).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	virtualWanId := "virtualWanId_example" // string | Virtual WAN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CheckVirtualWanUsed(context.Background(), omadacId, siteTemplateId).VirtualWanId(virtualWanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CheckVirtualWanUsed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckVirtualWanUsed`: OperationResponseVirtualWanIdUsedOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CheckVirtualWanUsed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckVirtualWanUsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualWanId** | **string** | Virtual WAN ID | 

### Return type

[**OperationResponseVirtualWanIdUsedOpenApiVO**](OperationResponseVirtualWanIdUsedOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckWanLanStatus

> CheckWanLanStatusOpenApiVO CheckWanLanStatus(ctx, omadacId, siteTemplateId).Execute()

Check template WAN-LAN status



### Example

```go
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
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CheckWanLanStatus(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CheckWanLanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckWanLanStatus`: CheckWanLanStatusOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CheckWanLanStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckWanLanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CheckWanLanStatusOpenApiVO**](CheckWanLanStatusOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmCreateVlanNetworkTemplate

> OperationResponseVlanNetworkIdListVO ConfirmCreateVlanNetworkTemplate(ctx, omadacId, siteTemplateId).CreateVlanParamOpenApiVO(createVlanParamOpenApiVO).Execute()

Confirm create lan network template



### Example

```go
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
	createVlanParamOpenApiVO := *openapiclient.NewCreateVlanParamOpenApiVO(*openapiclient.NewSelectPortBindingBriefVO(), *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example")) // CreateVlanParamOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ConfirmCreateVlanNetworkTemplate(context.Background(), omadacId, siteTemplateId).CreateVlanParamOpenApiVO(createVlanParamOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ConfirmCreateVlanNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmCreateVlanNetworkTemplate`: OperationResponseVlanNetworkIdListVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ConfirmCreateVlanNetworkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmCreateVlanNetworkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVlanParamOpenApiVO** | [**CreateVlanParamOpenApiVO**](CreateVlanParamOpenApiVO.md) |  | 

### Return type

[**OperationResponseVlanNetworkIdListVO**](OperationResponseVlanNetworkIdListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmModifyVlanNetworkTemplate

> OperationResponseVlanNetworkIdListVO ConfirmModifyVlanNetworkTemplate(ctx, omadacId, siteTemplateId, networkId).ModifyVlanParamOpenApiVO(modifyVlanParamOpenApiVO).Execute()

Confirm modify lan network template



### Example

```go
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
	networkId := "networkId_example" // string | Network ID
	modifyVlanParamOpenApiVO := *openapiclient.NewModifyVlanParamOpenApiVO(*openapiclient.NewSelectPortBindingBriefVO(), *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example")) // ModifyVlanParamOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ConfirmModifyVlanNetworkTemplate(context.Background(), omadacId, siteTemplateId, networkId).ModifyVlanParamOpenApiVO(modifyVlanParamOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ConfirmModifyVlanNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmModifyVlanNetworkTemplate`: OperationResponseVlanNetworkIdListVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ConfirmModifyVlanNetworkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmModifyVlanNetworkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyVlanParamOpenApiVO** | [**ModifyVlanParamOpenApiVO**](ModifyVlanParamOpenApiVO.md) |  | 

### Return type

[**OperationResponseVlanNetworkIdListVO**](OperationResponseVlanNetworkIdListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanDnsTemplate

> OperationResponseWithoutResult CreateLanDnsTemplate(ctx, omadacId, siteTemplateId).LanDnsOpenApiVO(lanDnsOpenApiVO).Execute()

Create LAN Dns template



### Example

```go
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
	lanDnsOpenApiVO := *openapiclient.NewLanDnsOpenApiVO("Domain_example", false, "Name_example", int32(123)) // LanDnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CreateLanDnsTemplate(context.Background(), omadacId, siteTemplateId).LanDnsOpenApiVO(lanDnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CreateLanDnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanDnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CreateLanDnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanDnsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanDnsOpenApiVO** | [**LanDnsOpenApiVO**](LanDnsOpenApiVO.md) |  | 

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


## CreateLanNetworkTemplate

> ResponseIdVO CreateLanNetworkTemplate(ctx, omadacId, siteTemplateId).LanNetworkTemplateOpenApiVO(lanNetworkTemplateOpenApiVO).Execute()

Create LAN network template



### Example

```go
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
	lanNetworkTemplateOpenApiVO := *openapiclient.NewLanNetworkTemplateOpenApiVO(false, "Name_example", int32(123)) // LanNetworkTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CreateLanNetworkTemplate(context.Background(), omadacId, siteTemplateId).LanNetworkTemplateOpenApiVO(lanNetworkTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CreateLanNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanNetworkTemplate`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CreateLanNetworkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanNetworkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanNetworkTemplateOpenApiVO** | [**LanNetworkTemplateOpenApiVO**](LanNetworkTemplateOpenApiVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanProfileTemplate

> ResponseIdVO CreateLanProfileTemplate(ctx, omadacId, siteTemplateId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()

Create new LAN profile template



### Example

```go
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
	lanProfileConfigOpenApiVO := *openapiclient.NewLanProfileConfigOpenApiVO(int32(123), int32(123), false, false, "Name_example", "NativeNetworkId_example", int32(123), false, false) // LanProfileConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CreateLanProfileTemplate(context.Background(), omadacId, siteTemplateId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CreateLanProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanProfileTemplate`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CreateLanProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanProfileConfigOpenApiVO** | [**LanProfileConfigOpenApiVO**](LanProfileConfigOpenApiVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOswLanProfileTemplate

> ResponseIdVO CreateOswLanProfileTemplate(ctx, omadacId, siteTemplateId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()

Create new switch profile template



### Example

```go
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
	lanProfileConfigOpenApiVO := *openapiclient.NewLanProfileConfigOpenApiVO(int32(123), int32(123), false, false, "Name_example", "NativeNetworkId_example", int32(123), false, false) // LanProfileConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CreateOswLanProfileTemplate(context.Background(), omadacId, siteTemplateId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CreateOswLanProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOswLanProfileTemplate`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CreateOswLanProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOswLanProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanProfileConfigOpenApiVO** | [**LanProfileConfigOpenApiVO**](LanProfileConfigOpenApiVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualWanTemplate

> OperationResponseWithoutResult CreateVirtualWanTemplate(ctx, omadacId, siteTemplateId).VirtualWanConfigOpenApiVO(virtualWanConfigOpenApiVO).Execute()

Create virtual WAN template



### Example

```go
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
	virtualWanConfigOpenApiVO := *openapiclient.NewVirtualWanConfigOpenApiVO("Name_example", "PhysicalWanId_example", *openapiclient.NewVirtualWanIpv4SettingConfigOpenApiVO("Proto_example", int32(123))) // VirtualWanConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CreateVirtualWanTemplate(context.Background(), omadacId, siteTemplateId).VirtualWanConfigOpenApiVO(virtualWanConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CreateVirtualWanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVirtualWanTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CreateVirtualWanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualWanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualWanConfigOpenApiVO** | [**VirtualWanConfigOpenApiVO**](VirtualWanConfigOpenApiVO.md) |  | 

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


## CreateVlans

> OperationResponseWithoutResult CreateVlans(ctx, omadacId, siteTemplateId).CreateVLANs(createVLANs).Execute()

Batch create vlans template



### Example

```go
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
	createVLANs := *openapiclient.NewCreateVLANs(false, "Name_example", "Vlans_example") // CreateVLANs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.CreateVlans(context.Background(), omadacId, siteTemplateId).CreateVLANs(createVLANs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.CreateVlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVlans`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.CreateVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVLANs** | [**CreateVLANs**](CreateVLANs.md) |  | 

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


## DeleteLanDnsTemplate

> OperationResponseWithoutResult DeleteLanDnsTemplate(ctx, omadacId, siteTemplateId, dnsId).Execute()

Delete LAN Dns template



### Example

```go
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
	dnsId := "dnsId_example" // string | LAN DNS ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.DeleteLanDnsTemplate(context.Background(), omadacId, siteTemplateId, dnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.DeleteLanDnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLanDnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.DeleteLanDnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**dnsId** | **string** | LAN DNS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanDnsTemplateRequest struct via the builder pattern


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


## DeleteLanNetworkTemplate

> OperationResponseWithoutResult DeleteLanNetworkTemplate(ctx, omadacId, siteTemplateId, networkId).Execute()

Delete LAN network template



### Example

```go
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
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.DeleteLanNetworkTemplate(context.Background(), omadacId, siteTemplateId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.DeleteLanNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLanNetworkTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.DeleteLanNetworkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanNetworkTemplateRequest struct via the builder pattern


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


## DeleteLanProfileTemplate

> OperationResponseWithoutResult DeleteLanProfileTemplate(ctx, omadacId, siteTemplateId, profileId).Execute()

Delete an existing LAN profile template



### Example

```go
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
	profileId := "profileId_example" // string | LAN profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.DeleteLanProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.DeleteLanProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLanProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.DeleteLanProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | LAN profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanProfileTemplateRequest struct via the builder pattern


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


## DeleteOswLanProfileTemplate

> OperationResponseWithoutResult DeleteOswLanProfileTemplate(ctx, omadacId, siteTemplateId, profileId).Execute()

Delete an existing switch profile template



### Example

```go
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
	profileId := "profileId_example" // string | LAN profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.DeleteOswLanProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.DeleteOswLanProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOswLanProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.DeleteOswLanProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | LAN profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOswLanProfileTemplateRequest struct via the builder pattern


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


## DeleteVirtualWan

> OperationResponseWithoutResult DeleteVirtualWan(ctx, omadacId, siteTemplateId, virtualWanId).Execute()

Delete virtual WAN template



### Example

```go
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
	virtualWanId := "virtualWanId_example" // string | Virtual WAN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.DeleteVirtualWan(context.Background(), omadacId, siteTemplateId, virtualWanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.DeleteVirtualWan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVirtualWan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.DeleteVirtualWan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**virtualWanId** | **string** | Virtual WAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualWanRequest struct via the builder pattern


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


## GetAllInterfacesForBatchIsolate

> OperationResponseListInterfaceForBatchIsolateOpenApiVO GetAllInterfacesForBatchIsolate(ctx, omadacId, siteTemplateId).SearchKey(searchKey).IsolationFilter(isolationFilter).Execute()

Get interface Grid



### Example

```go
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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field  (optional)
	isolationFilter := int32(56) // int32 | non-isolation set 0; isolation set 1; don't filter isolation vlan set 2. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetAllInterfacesForBatchIsolate(context.Background(), omadacId, siteTemplateId).SearchKey(searchKey).IsolationFilter(isolationFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetAllInterfacesForBatchIsolate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllInterfacesForBatchIsolate`: OperationResponseListInterfaceForBatchIsolateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetAllInterfacesForBatchIsolate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInterfacesForBatchIsolateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchKey** | **string** | Fuzzy query parameters, support field  | 
 **isolationFilter** | **int32** | non-isolation set 0; isolation set 1; don&#39;t filter isolation vlan set 2. | 

### Return type

[**OperationResponseListInterfaceForBatchIsolateOpenApiVO**](OperationResponseListInterfaceForBatchIsolateOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllLanNetworksTemplate

> OperationResponseListLanNetworkVO GetAllLanNetworksTemplate(ctx, omadacId, siteTemplateId).Execute()

Get all network templates for the omada id and site template id



### Example

```go
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
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetAllLanNetworksTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetAllLanNetworksTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllLanNetworksTemplate`: OperationResponseListLanNetworkVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetAllLanNetworksTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLanNetworksTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListLanNetworkVO**](OperationResponseListLanNetworkVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllValidVlansInSiteForSwitchOuiBasedVlan

> GridVOLanNetworkVlansOpenApiVO GetAllValidVlansInSiteForSwitchOuiBasedVlan(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get valid site template vlanList for switch oui based vlan



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetAllValidVlansInSiteForSwitchOuiBasedVlan(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetAllValidVlansInSiteForSwitchOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllValidVlansInSiteForSwitchOuiBasedVlan`: GridVOLanNetworkVlansOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetAllValidVlansInSiteForSwitchOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllValidVlansInSiteForSwitchOuiBasedVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**GridVOLanNetworkVlansOpenApiVO**](GridVOLanNetworkVlansOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoEffectDevicesTemplateWhenCreate

> OperationResponseAutoEffectDevicesForVlanVO GetAutoEffectDevicesTemplateWhenCreate(ctx, omadacId, siteTemplateId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetAutoEffectDevicesTemplateWhenCreate(context.Background(), omadacId, siteTemplateId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetAutoEffectDevicesTemplateWhenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoEffectDevicesTemplateWhenCreate`: OperationResponseAutoEffectDevicesForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetAutoEffectDevicesTemplateWhenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoEffectDevicesTemplateWhenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseAutoEffectDevicesForVlanVO**](OperationResponseAutoEffectDevicesForVlanVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoEffectDevicesTemplateWhenModify

> OperationResponseAutoEffectDevicesForVlanVO GetAutoEffectDevicesTemplateWhenModify(ctx, omadacId, siteTemplateId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	networkId := "networkId_example" // string | Network ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetAutoEffectDevicesTemplateWhenModify(context.Background(), omadacId, siteTemplateId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetAutoEffectDevicesTemplateWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoEffectDevicesTemplateWhenModify`: OperationResponseAutoEffectDevicesForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetAutoEffectDevicesTemplateWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoEffectDevicesTemplateWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseAutoEffectDevicesForVlanVO**](OperationResponseAutoEffectDevicesForVlanVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoSelectDevicesTemplateWhenModify

> OperationResponseAutoSelectDevicesForVlanVO GetAutoSelectDevicesTemplateWhenModify(ctx, omadacId, siteTemplateId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	networkId := "networkId_example" // string | Network ID
	lanNetworkOpenApiV3VO := *openapiclient.NewLanNetworkOpenApiV3VO(int32(123), false, "Name_example") // LanNetworkOpenApiV3VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetAutoSelectDevicesTemplateWhenModify(context.Background(), omadacId, siteTemplateId, networkId).LanNetworkOpenApiV3VO(lanNetworkOpenApiV3VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetAutoSelectDevicesTemplateWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoSelectDevicesTemplateWhenModify`: OperationResponseAutoSelectDevicesForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetAutoSelectDevicesTemplateWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoSelectDevicesTemplateWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkOpenApiV3VO** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

### Return type

[**OperationResponseAutoSelectDevicesForVlanVO**](OperationResponseAutoSelectDevicesForVlanVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableVirtualWanTemplate

> OperationResponseVirtualWanAvailablesOpenApiVO GetAvailableVirtualWanTemplate(ctx, omadacId, siteTemplateId).Execute()

Query available virtual WAN list for template



### Example

```go
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
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetAvailableVirtualWanTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetAvailableVirtualWanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableVirtualWanTemplate`: OperationResponseVirtualWanAvailablesOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetAvailableVirtualWanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableVirtualWanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVirtualWanAvailablesOpenApiVO**](OperationResponseVirtualWanAvailablesOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableWanPorts

> OperationResponseAvailableWanResultOpenApiVO GetAvailableWanPorts(ctx, omadacId, siteTemplateId).Function(function).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	function := int32(56) // int32 | function used for wan ports query. 0: ACL, 1: QOS, 2: IP_MAC_BLINDING, 3:IGMP_PROXY, 4: VIRTUAL_WAN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetAvailableWanPorts(context.Background(), omadacId, siteTemplateId).Function(function).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetAvailableWanPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableWanPorts`: OperationResponseAvailableWanResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetAvailableWanPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableWanPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **function** | **int32** | function used for wan ports query. 0: ACL, 1: QOS, 2: IP_MAC_BLINDING, 3:IGMP_PROXY, 4: VIRTUAL_WAN | 

### Return type

[**OperationResponseAvailableWanResultOpenApiVO**](OperationResponseAvailableWanResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSupportVlanNetworkDevicesBySiteTemplate

> OperationResponseGridVODeviceVO GetGridSupportVlanNetworkDevicesBySiteTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get device template list that can be dhcp server



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetGridSupportVlanNetworkDevicesBySiteTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetGridSupportVlanNetworkDevicesBySiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSupportVlanNetworkDevicesBySiteTemplate`: OperationResponseGridVODeviceVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetGridSupportVlanNetworkDevicesBySiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSupportVlanNetworkDevicesBySiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVODeviceVO**](OperationResponseGridVODeviceVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVirtualWanTemplate

> OperationResponseVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO GetGridVirtualWanTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Query virtual WAN list template



### Example

```go
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
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetGridVirtualWanTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetGridVirtualWanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVirtualWanTemplate`: OperationResponseVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetGridVirtualWanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVirtualWanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO**](OperationResponseVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaceLanNetworkV21

> OperationResponseResponseDataVOLanNetworkQueryOpenApiV2VO GetInterfaceLanNetworkV21(ctx, omadacId, siteTemplateId).Type_(type_).Execute()

Get all \"single\"/\"multi\" interface lan network template



### Example

```go
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
	type_ := int32(56) // int32 | Interface vlan type. When \"type\" is 0, return \"single\" interface lan network, else return \"single\"/\"multi\" interface lan network (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetInterfaceLanNetworkV21(context.Background(), omadacId, siteTemplateId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetInterfaceLanNetworkV21``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterfaceLanNetworkV21`: OperationResponseResponseDataVOLanNetworkQueryOpenApiV2VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetInterfaceLanNetworkV21`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceLanNetworkV21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **int32** | Interface vlan type. When \&quot;type\&quot; is 0, return \&quot;single\&quot; interface lan network, else return \&quot;single\&quot;/\&quot;multi\&quot; interface lan network | 

### Return type

[**OperationResponseResponseDataVOLanNetworkQueryOpenApiV2VO**](OperationResponseResponseDataVOLanNetworkQueryOpenApiV2VO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternetBasicPortInfo

> OperationResponseInternetBaseInfoTemplateOpenApiVO GetInternetBasicPortInfo(ctx, omadacId, siteTemplateId).Execute()

Get site template internet basic info



### Example

```go
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
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetInternetBasicPortInfo(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetInternetBasicPortInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternetBasicPortInfo`: OperationResponseInternetBaseInfoTemplateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetInternetBasicPortInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternetBasicPortInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseInternetBaseInfoTemplateOpenApiVO**](OperationResponseInternetBaseInfoTemplateOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternetLoadBalance

> OperationResponseWanLoadBalanceOpenApiVO GetInternetLoadBalance(ctx, omadacId, siteTemplateId).Execute()

Get site template internet load balance config



### Example

```go
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
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetInternetLoadBalance(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetInternetLoadBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternetLoadBalance`: OperationResponseWanLoadBalanceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetInternetLoadBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternetLoadBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWanLoadBalanceOpenApiVO**](OperationResponseWanLoadBalanceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternetTemplate

> OperationResponseInternetOpenApiVO GetInternetTemplate(ctx, omadacId, siteTemplateId).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetInternetTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetInternetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternetTemplate`: OperationResponseInternetOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetInternetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseInternetOpenApiVO**](OperationResponseInternetOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanDnsTemplateList

> OperationResponseGridVOLanDnsOpenApiVO GetLanDnsTemplateList(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get LAN Dns template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetLanDnsTemplateList(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetLanDnsTemplateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanDnsTemplateList`: OperationResponseGridVOLanDnsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetLanDnsTemplateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanDnsTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOLanDnsOpenApiVO**](OperationResponseGridVOLanDnsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanNetworkTemplate

> OperationResponseLanNetworkTemplateQueryOpenApiV3VO GetLanNetworkTemplate(ctx, omadacId, siteTemplateId, networkId).Execute()

Get LAN network template



### Example

```go
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
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetLanNetworkTemplate(context.Background(), omadacId, siteTemplateId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetLanNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanNetworkTemplate`: OperationResponseLanNetworkTemplateQueryOpenApiV3VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetLanNetworkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanNetworkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseLanNetworkTemplateQueryOpenApiV3VO**](OperationResponseLanNetworkTemplateQueryOpenApiV3VO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanNetworkTemplateList

> OperationResponseLanNetworkOpenApiV2GridVOLanNetworkTemplateQueryOpenApiV2VO GetLanNetworkTemplateList(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get LAN network template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetLanNetworkTemplateList(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetLanNetworkTemplateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanNetworkTemplateList`: OperationResponseLanNetworkOpenApiV2GridVOLanNetworkTemplateQueryOpenApiV2VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetLanNetworkTemplateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanNetworkTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseLanNetworkOpenApiV2GridVOLanNetworkTemplateQueryOpenApiV2VO**](OperationResponseLanNetworkOpenApiV2GridVOLanNetworkTemplateQueryOpenApiV2VO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanNetworkTemplateListV3

> OperationResponseLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO GetLanNetworkTemplateListV3(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get LAN network template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetLanNetworkTemplateListV3(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetLanNetworkTemplateListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanNetworkTemplateListV3`: OperationResponseLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetLanNetworkTemplateListV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanNetworkTemplateListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO**](OperationResponseLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanProfileTemplateList

> OperationResponseGridVOLanProfileOpenApiVO GetLanProfileTemplateList(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get LAN profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetLanProfileTemplateList(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetLanProfileTemplateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanProfileTemplateList`: OperationResponseGridVOLanProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetLanProfileTemplateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanProfileTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOLanProfileOpenApiVO**](OperationResponseGridVOLanProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalanceWeightStatusTemplate

> OperationResponse GetLoadBalanceWeightStatusTemplate(ctx, omadacId, siteTemplateId).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetLoadBalanceWeightStatusTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetLoadBalanceWeightStatusTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalanceWeightStatusTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetLoadBalanceWeightStatusTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalanceWeightStatusTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswLanProfileTemplateList

> OperationResponseGridVOLanProfileOpenApiVO GetOswLanProfileTemplateList(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get switch profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetOswLanProfileTemplateList(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetOswLanProfileTemplateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswLanProfileTemplateList`: OperationResponseGridVOLanProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetOswLanProfileTemplateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswLanProfileTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOLanProfileOpenApiVO**](OperationResponseGridVOLanProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectDeviceTemplatePortsInfo

> OperationResponseListSelectDeviceForVlanTemplateVO GetSelectDeviceTemplatePortsInfo(ctx, omadacId, siteTemplateId).SelectMacsWithVlanVO(selectMacsWithVlanVO).Execute()

Get the port information of the selected device templates when creating network for site template



### Example

```go
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
	selectMacsWithVlanVO := *openapiclient.NewSelectMacsWithVlanVO([]string{"Macs_example"}, int32(123)) // SelectMacsWithVlanVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetSelectDeviceTemplatePortsInfo(context.Background(), omadacId, siteTemplateId).SelectMacsWithVlanVO(selectMacsWithVlanVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetSelectDeviceTemplatePortsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectDeviceTemplatePortsInfo`: OperationResponseListSelectDeviceForVlanTemplateVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetSelectDeviceTemplatePortsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectDeviceTemplatePortsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectMacsWithVlanVO** | [**SelectMacsWithVlanVO**](SelectMacsWithVlanVO.md) |  | 

### Return type

[**OperationResponseListSelectDeviceForVlanTemplateVO**](OperationResponseListSelectDeviceForVlanTemplateVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectDeviceTemplatePortsInfoWhenModify

> OperationResponseListSelectDeviceForVlanTemplateVO GetSelectDeviceTemplatePortsInfoWhenModify(ctx, omadacId, siteTemplateId, networkId).SelectMacsWithVlanVO(selectMacsWithVlanVO).Execute()

Get the port information of the selected device templates when modifying the network for site template



### Example

```go
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
	networkId := "networkId_example" // string | Network ID
	selectMacsWithVlanVO := *openapiclient.NewSelectMacsWithVlanVO([]string{"Macs_example"}, int32(123)) // SelectMacsWithVlanVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetSelectDeviceTemplatePortsInfoWhenModify(context.Background(), omadacId, siteTemplateId, networkId).SelectMacsWithVlanVO(selectMacsWithVlanVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetSelectDeviceTemplatePortsInfoWhenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectDeviceTemplatePortsInfoWhenModify`: OperationResponseListSelectDeviceForVlanTemplateVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetSelectDeviceTemplatePortsInfoWhenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectDeviceTemplatePortsInfoWhenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **selectMacsWithVlanVO** | [**SelectMacsWithVlanVO**](SelectMacsWithVlanVO.md) |  | 

### Return type

[**OperationResponseListSelectDeviceForVlanTemplateVO**](OperationResponseListSelectDeviceForVlanTemplateVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportPortsDeviceTemplates

> OperationResponseListDeviceTemplateForVlanVO GetSupportPortsDeviceTemplates(ctx, omadacId, siteTemplateId).Execute()

Get grid devices templates that support ports config.



### Example

```go
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
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetSupportPortsDeviceTemplates(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetSupportPortsDeviceTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportPortsDeviceTemplates`: OperationResponseListDeviceTemplateForVlanVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetSupportPortsDeviceTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportPortsDeviceTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListDeviceTemplateForVlanVO**](OperationResponseListDeviceTemplateForVlanVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateGridVlanNetworkAffectedDevicePorts

> OperationResponseVlanForPortDisplayGridVOVlanNetworkAffectingDeviceDetailVO GetTemplateGridVlanNetworkAffectedDevicePorts(ctx, omadacId, siteTemplateId, networkId, vlan).Page(page).PageSize(pageSize).Execute()

Get the ports of devices template that use the network



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	networkId := "networkId_example" // string | Network ID
	vlan := "vlan_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetTemplateGridVlanNetworkAffectedDevicePorts(context.Background(), omadacId, siteTemplateId, networkId, vlan).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetTemplateGridVlanNetworkAffectedDevicePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateGridVlanNetworkAffectedDevicePorts`: OperationResponseVlanForPortDisplayGridVOVlanNetworkAffectingDeviceDetailVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetTemplateGridVlanNetworkAffectedDevicePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 
**vlan** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateGridVlanNetworkAffectedDevicePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





### Return type

[**OperationResponseVlanForPortDisplayGridVOVlanNetworkAffectingDeviceDetailVO**](OperationResponseVlanForPortDisplayGridVOVlanNetworkAffectingDeviceDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUseLanProfileESTemplateList

> OperationResponseGridVOUseProfileOswOpenApiVO GetUseLanProfileESTemplateList(ctx, omadacId, siteTemplateId, profileId).Page(page).PageSize(pageSize).Execute()

Get Use LAN profile ES template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | LAN profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetUseLanProfileESTemplateList(context.Background(), omadacId, siteTemplateId, profileId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetUseLanProfileESTemplateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUseLanProfileESTemplateList`: OperationResponseGridVOUseProfileOswOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetUseLanProfileESTemplateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | LAN profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUseLanProfileESTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 




### Return type

[**OperationResponseGridVOUseProfileOswOpenApiVO**](OperationResponseGridVOUseProfileOswOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlanNetworkAffectedDevice

> OperationResponseListVlanNetworkAffectingDeviceVO GetVlanNetworkAffectedDevice(ctx, omadacId, siteTemplateId, networkId, vlan).Execute()

Get device template list that use the network



### Example

```go
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
	networkId := "networkId_example" // string | Network ID
	vlan := "vlan_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetVlanNetworkAffectedDevice(context.Background(), omadacId, siteTemplateId, networkId, vlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetVlanNetworkAffectedDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVlanNetworkAffectedDevice`: OperationResponseListVlanNetworkAffectingDeviceVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetVlanNetworkAffectedDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 
**vlan** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVlanNetworkAffectedDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseListVlanNetworkAffectingDeviceVO**](OperationResponseListVlanNetworkAffectingDeviceVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlanNetworkAffectedSsid

> OperationResponseGridVOVlanNetworkAffectingSsidVO GetVlanNetworkAffectedSsid(ctx, omadacId, siteTemplateId, networkId, vlan).Page(page).PageSize(pageSize).Execute()

Get grid ssid template list that use the vlan



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	networkId := "networkId_example" // string | Network ID
	vlan := "vlan_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetVlanNetworkAffectedSsid(context.Background(), omadacId, siteTemplateId, networkId, vlan).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetVlanNetworkAffectedSsid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVlanNetworkAffectedSsid`: OperationResponseGridVOVlanNetworkAffectingSsidVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetVlanNetworkAffectedSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 
**vlan** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVlanNetworkAffectedSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





### Return type

[**OperationResponseGridVOVlanNetworkAffectingSsidVO**](OperationResponseGridVOVlanNetworkAffectingSsidVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWanPortsMaxSpeed

> OperationResponseWanMaxSpeedOpenApiVO GetWanPortsMaxSpeed(ctx, omadacId, siteTemplateId).WanId(wanId).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	wanId := "wanId_example" // string | WanId, can pass multiple wan ports using comma separator like '{portId1},{portId2}'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.GetWanPortsMaxSpeed(context.Background(), omadacId, siteTemplateId).WanId(wanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.GetWanPortsMaxSpeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWanPortsMaxSpeed`: OperationResponseWanMaxSpeedOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.GetWanPortsMaxSpeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanPortsMaxSpeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanId** | **string** | WanId, can pass multiple wan ports using comma separator like &#39;{portId1},{portId2}&#39;. | 

### Return type

[**OperationResponseWanMaxSpeedOpenApiVO**](OperationResponseWanMaxSpeedOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyInternetBasicPortInfo

> OperationResponseWithoutResult ModifyInternetBasicPortInfo(ctx, omadacId, siteTemplateId).InternetBaseConfigTemplateOpenApiVO(internetBaseConfigTemplateOpenApiVO).Execute()

Modify site template internet basic configuration



### Example

```go
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
	internetBaseConfigTemplateOpenApiVO := *openapiclient.NewInternetBaseConfigTemplateOpenApiVO([]string{"WanPortList_example"}) // InternetBaseConfigTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyInternetBasicPortInfo(context.Background(), omadacId, siteTemplateId).InternetBaseConfigTemplateOpenApiVO(internetBaseConfigTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyInternetBasicPortInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyInternetBasicPortInfo`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyInternetBasicPortInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyInternetBasicPortInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **internetBaseConfigTemplateOpenApiVO** | [**InternetBaseConfigTemplateOpenApiVO**](InternetBaseConfigTemplateOpenApiVO.md) |  | 

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


## ModifyInternetLoadBalance

> OperationResponseWithoutResult ModifyInternetLoadBalance(ctx, omadacId, siteTemplateId).WanLoadBalanceOpenApiVO(wanLoadBalanceOpenApiVO).Execute()

Modify site template internet load balance config



### Example

```go
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
	wanLoadBalanceOpenApiVO := *openapiclient.NewWanLoadBalanceOpenApiVO(false, false, []int32{int32(123)}) // WanLoadBalanceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyInternetLoadBalance(context.Background(), omadacId, siteTemplateId).WanLoadBalanceOpenApiVO(wanLoadBalanceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyInternetLoadBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyInternetLoadBalance`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyInternetLoadBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyInternetLoadBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanLoadBalanceOpenApiVO** | [**WanLoadBalanceOpenApiVO**](WanLoadBalanceOpenApiVO.md) |  | 

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


## ModifyLanDnsTemplate

> OperationResponseWithoutResult ModifyLanDnsTemplate(ctx, omadacId, siteTemplateId, dnsId).LanDnsOpenApiVO(lanDnsOpenApiVO).Execute()

Modify LAN Dns template



### Example

```go
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
	dnsId := "dnsId_example" // string | LAN DNS ID
	lanDnsOpenApiVO := *openapiclient.NewLanDnsOpenApiVO("Domain_example", false, "Name_example", int32(123)) // LanDnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyLanDnsTemplate(context.Background(), omadacId, siteTemplateId, dnsId).LanDnsOpenApiVO(lanDnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyLanDnsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanDnsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyLanDnsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**dnsId** | **string** | LAN DNS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanDnsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanDnsOpenApiVO** | [**LanDnsOpenApiVO**](LanDnsOpenApiVO.md) |  | 

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


## ModifyLanNetworkTemplate

> OperationResponseWithoutResult ModifyLanNetworkTemplate(ctx, omadacId, siteTemplateId, networkId).LanNetworkTemplateOpenApiVO(lanNetworkTemplateOpenApiVO).Execute()

Modify LAN network template



### Example

```go
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
	networkId := "networkId_example" // string | Network ID
	lanNetworkTemplateOpenApiVO := *openapiclient.NewLanNetworkTemplateOpenApiVO(false, "Name_example", int32(123)) // LanNetworkTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyLanNetworkTemplate(context.Background(), omadacId, siteTemplateId, networkId).LanNetworkTemplateOpenApiVO(lanNetworkTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyLanNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanNetworkTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyLanNetworkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanNetworkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanNetworkTemplateOpenApiVO** | [**LanNetworkTemplateOpenApiVO**](LanNetworkTemplateOpenApiVO.md) |  | 

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


## ModifyLanProfileTemplate

> OperationResponseWithoutResult ModifyLanProfileTemplate(ctx, omadacId, siteTemplateId, profileId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()

Modify a LAN profile template



### Example

```go
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
	profileId := "profileId_example" // string | lanProfileId
	lanProfileConfigOpenApiVO := *openapiclient.NewLanProfileConfigOpenApiVO(int32(123), int32(123), false, false, "Name_example", "NativeNetworkId_example", int32(123), false, false) // LanProfileConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyLanProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyLanProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyLanProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | lanProfileId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanProfileConfigOpenApiVO** | [**LanProfileConfigOpenApiVO**](LanProfileConfigOpenApiVO.md) |  | 

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


## ModifyOswLanProfileTemplate

> OperationResponseWithoutResult ModifyOswLanProfileTemplate(ctx, omadacId, siteTemplateId, profileId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()

Modify a switch profile template



### Example

```go
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
	profileId := "profileId_example" // string | lanProfileId
	lanProfileConfigOpenApiVO := *openapiclient.NewLanProfileConfigOpenApiVO(int32(123), int32(123), false, false, "Name_example", "NativeNetworkId_example", int32(123), false, false) // LanProfileConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyOswLanProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).LanProfileConfigOpenApiVO(lanProfileConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyOswLanProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswLanProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyOswLanProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | lanProfileId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswLanProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanProfileConfigOpenApiVO** | [**LanProfileConfigOpenApiVO**](LanProfileConfigOpenApiVO.md) |  | 

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


## ModifyVirtualWanStatusTemplate

> OperationResponseWithoutResult ModifyVirtualWanStatusTemplate(ctx, omadacId, siteTemplateId, virtualWanId).VirtualWanStatusOpenApiVO(virtualWanStatusOpenApiVO).Execute()

Modify virtual WAN status template



### Example

```go
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
	virtualWanId := "virtualWanId_example" // string | Virtual WAN ID
	virtualWanStatusOpenApiVO := *openapiclient.NewVirtualWanStatusOpenApiVO(false) // VirtualWanStatusOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyVirtualWanStatusTemplate(context.Background(), omadacId, siteTemplateId, virtualWanId).VirtualWanStatusOpenApiVO(virtualWanStatusOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyVirtualWanStatusTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVirtualWanStatusTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyVirtualWanStatusTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**virtualWanId** | **string** | Virtual WAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVirtualWanStatusTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **virtualWanStatusOpenApiVO** | [**VirtualWanStatusOpenApiVO**](VirtualWanStatusOpenApiVO.md) |  | 

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


## ModifyVirtualWanTemplate

> OperationResponseWithoutResult ModifyVirtualWanTemplate(ctx, omadacId, siteTemplateId, virtualWanId).VirtualWanConfigOpenApiVO(virtualWanConfigOpenApiVO).Execute()

Modify virtual WAN template



### Example

```go
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
	virtualWanId := "virtualWanId_example" // string | Virtual WAN ID
	virtualWanConfigOpenApiVO := *openapiclient.NewVirtualWanConfigOpenApiVO("Name_example", "PhysicalWanId_example", *openapiclient.NewVirtualWanIpv4SettingConfigOpenApiVO("Proto_example", int32(123))) // VirtualWanConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyVirtualWanTemplate(context.Background(), omadacId, siteTemplateId, virtualWanId).VirtualWanConfigOpenApiVO(virtualWanConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyVirtualWanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVirtualWanTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyVirtualWanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**virtualWanId** | **string** | Virtual WAN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVirtualWanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **virtualWanConfigOpenApiVO** | [**VirtualWanConfigOpenApiVO**](VirtualWanConfigOpenApiVO.md) |  | 

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


## ModifyWanPortsTemplate

> OperationResponseWithoutResult ModifyWanPortsTemplate(ctx, omadacId, siteTemplateId).WanPortsOpenApiVO(wanPortsOpenApiVO).Execute()

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	wanPortsOpenApiVO := *openapiclient.NewWanPortsOpenApiVO(false) // WanPortsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WiredNetworkTemplateAPI.ModifyWanPortsTemplate(context.Background(), omadacId, siteTemplateId).WanPortsOpenApiVO(wanPortsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WiredNetworkTemplateAPI.ModifyWanPortsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWanPortsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WiredNetworkTemplateAPI.ModifyWanPortsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWanPortsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanPortsOpenApiVO** | [**WanPortsOpenApiVO**](WanPortsOpenApiVO.md) |  | 

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


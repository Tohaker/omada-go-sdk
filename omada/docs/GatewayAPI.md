# GatewayAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchModifyPortConfig**](GatewayAPI.md#batchmodifyportconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/multi-ports/config | Batch modify gateway port config
[**ChangeInternetState**](GatewayAPI.md#changeinternetstate) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/internet-state | Change Internet State
[**ChangeIpv6State**](GatewayAPI.md#changeipv6state) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/gateways/{gatewayMac}/ipv6State | Modify IPv6 state
[**ChangeOduMode**](GatewayAPI.md#changeodumode) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/change-mode | Change mode for 5G-Outdoor
[**ChangeOduModeForMsp**](GatewayAPI.md#changeodumodeformsp) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/gateways/{gatewayMac}/change-mode | Change mode for 5G-Outdoor for msp
[**ChangePinSetting**](GatewayAPI.md#changepinsetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/pin | Modify PIN management information of an LTE model
[**DeleteIspFile**](GatewayAPI.md#deleteispfile) | **Delete** /openapi/v1/{omadacId}/files/sites/{siteId}/gateways/{gatewayMac}/isp-upgrade/{fileId} | Delete isp file
[**GetClientDhcpLeaseTimes**](GatewayAPI.md#getclientdhcpleasetimes) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/dhcp-lease-time | Get gateway dhcp client lease time
[**GetEnableWanDetail**](GatewayAPI.md#getenablewandetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/health/gateways/{gatewayMac}/wans/details | Get enable wan port detail
[**GetGatewayInfo1**](GatewayAPI.md#getgatewayinfo1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac} | Get gateway info
[**GetGeneralConfig1**](GatewayAPI.md#getgeneralconfig1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/general-config | Get gateway general config
[**GetGridDhcpUserList1**](GatewayAPI.md#getgriddhcpuserlist1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/dhcp/user-list | Get gateway dhcp user list
[**GetIpptLanIpv6**](GatewayAPI.md#getipptlanipv6) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/network/ipv6 | Get lan ipv6 config for ippt
[**GetLanStatus**](GatewayAPI.md#getlanstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/lan-status | Get gateway lan status
[**GetPinSetting**](GatewayAPI.md#getpinsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/pin | Get PIN setting
[**GetPorts**](GatewayAPI.md#getports) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/ports | Get gateway ports
[**GetSimCardUsed**](GatewayAPI.md#getsimcardused) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/simCardUsed | Get SIM card used
[**GetSpecifiedGatewayInfos**](GatewayAPI.md#getspecifiedgatewayinfos) | **Get** /openapi/v1/{omadacId}/global/gateways/wan-status | Get global gateways wan status
[**GetSsidDetail1**](GatewayAPI.md#getssiddetail1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/wlans/ssid | Get SSID detail info
[**GetWanStatus**](GatewayAPI.md#getwanstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/wan-status | Get gateway wan status
[**IspUpgrade**](GatewayAPI.md#ispupgrade) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/gateways/{gatewayMac}/isp-upgrade | Upgrade isp
[**ModifyConfigAdvanced**](GatewayAPI.md#modifyconfigadvanced) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/advanced | Modify gateway advanced config
[**ModifyConfigCommonAdvanced**](GatewayAPI.md#modifyconfigcommonadvanced) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/advanced/common | Modify gateway advanced common config
[**ModifyConfigGeneral**](GatewayAPI.md#modifyconfiggeneral) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/general | Modify gateway general config
[**ModifyConfigRadios**](GatewayAPI.md#modifyconfigradios) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/radios | Modify gateway radios config
[**ModifyConfigServices**](GatewayAPI.md#modifyconfigservices) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/services | Modify gateway config service
[**ModifyConfigWirelessAdvanced**](GatewayAPI.md#modifyconfigwirelessadvanced) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/advanced/wireless | Modify gateway advanced wireless config
[**ModifyConfigWlans**](GatewayAPI.md#modifyconfigwlans) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/wlans | Modify gateway wlans config
[**ModifyGeneralConfig1**](GatewayAPI.md#modifygeneralconfig1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/general-config | Modify gateway general config
[**ModifyIpptLanIpv6**](GatewayAPI.md#modifyipptlanipv6) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/network/ipv6 | Modify lan ipv6 config for ippt
[**ModifyOsgConfigForMsp**](GatewayAPI.md#modifyosgconfigformsp) | **Patch** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/gateways/{gatewayMac} | Modify gateway config for msp
[**ModifyOsgConfigGeneralForMsp**](GatewayAPI.md#modifyosgconfiggeneralformsp) | **Put** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/gateways/{gatewayMac}/config/general | Modify general gateway config for msp
[**ModifyPortConfig1**](GatewayAPI.md#modifyportconfig1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/ports/{port}/config | Modify gateway port config
[**ModifySshSettingByMac**](GatewayAPI.md#modifysshsettingbymac) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/services/ssh | Modify SSH setting by mac
[**ModifySsidBasicConfig**](GatewayAPI.md#modifyssidbasicconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/wlans/ssid/{ssidId}/basic-config | modify SSID basic config
[**ModifySsidMacFilterConfig**](GatewayAPI.md#modifyssidmacfilterconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/config/wlans/ssid/{ssidId}/mac-filter | Update SSID mac filter config
[**MoveToSite1**](GatewayAPI.md#movetosite1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/gateways/move | Move devices to another site
[**MspMoveToCustomer2**](GatewayAPI.md#mspmovetocustomer2) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/cmd/gateways/move | Move site to target customer
[**RecoveryPoePort**](GatewayAPI.md#recoverypoeport) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/ports/{port}/restart | Recovery gateway poe port
[**UploadIspFile**](GatewayAPI.md#uploadispfile) | **Post** /openapi/v1/{omadacId}/files/sites/{siteId}/gateways/{gatewayMac}/isp-upgrade | Upload isp upgrade file



## BatchModifyPortConfig

> OperationResponseGatewayPortsConfigEntity BatchModifyPortConfig(ctx, omadacId, siteId, gatewayMac).GatewayPortsConfigEntity(gatewayPortsConfigEntity).Execute()

Batch modify gateway port config



### Example

```go
package main

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
	gatewayPortsConfigEntity := *openapiclient.NewGatewayPortsConfigEntity() // GatewayPortsConfigEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.BatchModifyPortConfig(context.Background(), omadacId, siteId, gatewayMac).GatewayPortsConfigEntity(gatewayPortsConfigEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.BatchModifyPortConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyPortConfig`: OperationResponseGatewayPortsConfigEntity
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.BatchModifyPortConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchModifyPortConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gatewayPortsConfigEntity** | [**GatewayPortsConfigEntity**](GatewayPortsConfigEntity.md) |  | 

### Return type

[**OperationResponseGatewayPortsConfigEntity**](OperationResponseGatewayPortsConfigEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeInternetState

> OperationResponseWithoutResult ChangeInternetState(ctx, omadacId, siteId, gatewayMac).ChangeInternetStateOpenApiVO(changeInternetStateOpenApiVO).Execute()

Change Internet State



### Example

```go
package main

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
	changeInternetStateOpenApiVO := *openapiclient.NewChangeInternetStateOpenApiVO() // ChangeInternetStateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ChangeInternetState(context.Background(), omadacId, siteId, gatewayMac).ChangeInternetStateOpenApiVO(changeInternetStateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ChangeInternetState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeInternetState`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ChangeInternetState`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiChangeInternetStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **changeInternetStateOpenApiVO** | [**ChangeInternetStateOpenApiVO**](ChangeInternetStateOpenApiVO.md) |  | 

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


## ChangeIpv6State

> OperationResponseWithoutResult ChangeIpv6State(ctx, omadacId, siteId, gatewayMac).ModifyIpv6State(modifyIpv6State).Execute()

Modify IPv6 state



### Example

```go
package main

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
	modifyIpv6State := *openapiclient.NewModifyIpv6State(int32(123), int32(123)) // ModifyIpv6State | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ChangeIpv6State(context.Background(), omadacId, siteId, gatewayMac).ModifyIpv6State(modifyIpv6State).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ChangeIpv6State``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeIpv6State`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ChangeIpv6State`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiChangeIpv6StateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyIpv6State** | [**ModifyIpv6State**](ModifyIpv6State.md) |  | 

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


## ChangeOduMode

> OperationResponseWithoutResult ChangeOduMode(ctx, omadacId, siteId, gatewayMac).ChangeOduModeVO(changeOduModeVO).Execute()

Change mode for 5G-Outdoor



### Example

```go
package main

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
	changeOduModeVO := *openapiclient.NewChangeOduModeVO() // ChangeOduModeVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ChangeOduMode(context.Background(), omadacId, siteId, gatewayMac).ChangeOduModeVO(changeOduModeVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ChangeOduMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeOduMode`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ChangeOduMode`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiChangeOduModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **changeOduModeVO** | [**ChangeOduModeVO**](ChangeOduModeVO.md) |  | 

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


## ChangeOduModeForMsp

> OperationResponseWithoutResult ChangeOduModeForMsp(ctx, mspId, customerId, siteId, gatewayMac).ChangeOduModeVO(changeOduModeVO).Execute()

Change mode for 5G-Outdoor for msp



### Example

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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	changeOduModeVO := *openapiclient.NewChangeOduModeVO() // ChangeOduModeVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ChangeOduModeForMsp(context.Background(), mspId, customerId, siteId, gatewayMac).ChangeOduModeVO(changeOduModeVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ChangeOduModeForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeOduModeForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ChangeOduModeForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeOduModeForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **changeOduModeVO** | [**ChangeOduModeVO**](ChangeOduModeVO.md) |  | 

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


## ChangePinSetting

> OperationResponseOsgLtePinLockResultOpenApiVO ChangePinSetting(ctx, omadacId, siteId, gatewayMac).OsgLtePinOpenApiVO(osgLtePinOpenApiVO).Execute()

Modify PIN management information of an LTE model



### Example

```go
package main

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
	osgLtePinOpenApiVO := *openapiclient.NewOsgLtePinOpenApiVO() // OsgLtePinOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ChangePinSetting(context.Background(), omadacId, siteId, gatewayMac).OsgLtePinOpenApiVO(osgLtePinOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ChangePinSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangePinSetting`: OperationResponseOsgLtePinLockResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ChangePinSetting`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiChangePinSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgLtePinOpenApiVO** | [**OsgLtePinOpenApiVO**](OsgLtePinOpenApiVO.md) |  | 

### Return type

[**OperationResponseOsgLtePinLockResultOpenApiVO**](OperationResponseOsgLtePinLockResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIspFile

> OperationResponseWithoutResult DeleteIspFile(ctx, omadacId, siteId, gatewayMac, fileId).Execute()

Delete isp file



### Example

```go
package main

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
	fileId := "fileId_example" // string | fileId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.DeleteIspFile(context.Background(), omadacId, siteId, gatewayMac, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.DeleteIspFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIspFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.DeleteIspFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 
**fileId** | **string** | fileId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIspFileRequest struct via the builder pattern


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


## GetClientDhcpLeaseTimes

> OperationResponseListClientDhcpLeaseTimeOpenApiVO GetClientDhcpLeaseTimes(ctx, omadacId, siteId).QueryDhcpLeaseTimeVO(queryDhcpLeaseTimeVO).Execute()

Get gateway dhcp client lease time



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	queryDhcpLeaseTimeVO := *openapiclient.NewQueryDhcpLeaseTimeVO([]openapiclient.QueryDhcpLeaseTimeParamVO{*openapiclient.NewQueryDhcpLeaseTimeParamVO("IpAddress_example", "Mac_example")}) // QueryDhcpLeaseTimeVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GetClientDhcpLeaseTimes(context.Background(), omadacId, siteId).QueryDhcpLeaseTimeVO(queryDhcpLeaseTimeVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetClientDhcpLeaseTimes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientDhcpLeaseTimes`: OperationResponseListClientDhcpLeaseTimeOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetClientDhcpLeaseTimes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientDhcpLeaseTimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryDhcpLeaseTimeVO** | [**QueryDhcpLeaseTimeVO**](QueryDhcpLeaseTimeVO.md) |  | 

### Return type

[**OperationResponseListClientDhcpLeaseTimeOpenApiVO**](OperationResponseListClientDhcpLeaseTimeOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnableWanDetail

> OperationResponseWanDetails GetEnableWanDetail(ctx, omadacId, siteId, gatewayMac).Execute()

Get enable wan port detail



### Example

```go
package main

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
	resp, r, err := apiClient.GatewayAPI.GetEnableWanDetail(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetEnableWanDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnableWanDetail`: OperationResponseWanDetails
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetEnableWanDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetEnableWanDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWanDetails**](OperationResponseWanDetails.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayInfo1

> OperationResponseGatewayInfo GetGatewayInfo1(ctx, omadacId, siteId, gatewayMac).Execute()

Get gateway info



### Example

```go
package main

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
	resp, r, err := apiClient.GatewayAPI.GetGatewayInfo1(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetGatewayInfo1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGatewayInfo1`: OperationResponseGatewayInfo
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetGatewayInfo1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGatewayInfo1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseGatewayInfo**](OperationResponseGatewayInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneralConfig1

> OperationResponseGatewayGeneralConfig GetGeneralConfig1(ctx, omadacId, siteId, gatewayMac).Execute()

Get gateway general config



### Example

```go
package main

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
	resp, r, err := apiClient.GatewayAPI.GetGeneralConfig1(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetGeneralConfig1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGeneralConfig1`: OperationResponseGatewayGeneralConfig
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetGeneralConfig1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGeneralConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseGatewayGeneralConfig**](OperationResponseGatewayGeneralConfig.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDhcpUserList1

> OperationResponseOsgDhcpUserGridVOOsgDhcpUserVO GetGridDhcpUserList1(ctx, omadacId, siteId, gatewayMac).Page(page).PageSize(pageSize).Execute()

Get gateway dhcp user list



### Example

```go
package main

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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GetGridDhcpUserList1(context.Background(), omadacId, siteId, gatewayMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetGridDhcpUserList1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDhcpUserList1`: OperationResponseOsgDhcpUserGridVOOsgDhcpUserVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetGridDhcpUserList1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridDhcpUserList1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseOsgDhcpUserGridVOOsgDhcpUserVO**](OperationResponseOsgDhcpUserGridVOOsgDhcpUserVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpptLanIpv6

> OperationResponseIpv6ForIpptOpenApiVO GetIpptLanIpv6(ctx, omadacId, siteId, gatewayMac).Execute()

Get lan ipv6 config for ippt



### Example

```go
package main

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
	resp, r, err := apiClient.GatewayAPI.GetIpptLanIpv6(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetIpptLanIpv6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpptLanIpv6`: OperationResponseIpv6ForIpptOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetIpptLanIpv6`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetIpptLanIpv6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseIpv6ForIpptOpenApiVO**](OperationResponseIpv6ForIpptOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanStatus

> OperationResponseListLanStatus GetLanStatus(ctx, siteId, gatewayMac, omadacId).Execute()

Get gateway lan status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	siteId := "siteId_example" // string | Site ID
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	omadacId := "omadacId_example" // string | omadacId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GetLanStatus(context.Background(), siteId, gatewayMac, omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetLanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanStatus`: OperationResponseListLanStatus
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetLanStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 
**omadacId** | **string** | omadacId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListLanStatus**](OperationResponseListLanStatus.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPinSetting

> OperationResponseOsgPinDetailOpenApiVO GetPinSetting(ctx, omadacId, siteId, gatewayMac).Execute()

Get PIN setting



### Example

```go
package main

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
	resp, r, err := apiClient.GatewayAPI.GetPinSetting(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetPinSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPinSetting`: OperationResponseOsgPinDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetPinSetting`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPinSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOsgPinDetailOpenApiVO**](OperationResponseOsgPinDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPorts

> OperationResponseGatewayPortInfos GetPorts(ctx, omadacId, siteId, gatewayMac).Execute()

Get gateway ports



### Example

```go
package main

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
	resp, r, err := apiClient.GatewayAPI.GetPorts(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPorts`: OperationResponseGatewayPortInfos
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetPorts`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseGatewayPortInfos**](OperationResponseGatewayPortInfos.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimCardUsed

> OperationResponseOsgSimCardOpenApiVO GetSimCardUsed(ctx, omadacId, siteId, gatewayMac).Execute()

Get SIM card used



### Example

```go
package main

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
	resp, r, err := apiClient.GatewayAPI.GetSimCardUsed(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetSimCardUsed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimCardUsed`: OperationResponseOsgSimCardOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetSimCardUsed`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSimCardUsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOsgSimCardOpenApiVO**](OperationResponseOsgSimCardOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecifiedGatewayInfos

> OperationResponseGridVOGatewayInfos GetSpecifiedGatewayInfos(ctx, omadacId).Page(page).PageSize(pageSize).Execute()

Get global gateways wan status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GetSpecifiedGatewayInfos(context.Background(), omadacId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetSpecifiedGatewayInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpecifiedGatewayInfos`: OperationResponseGridVOGatewayInfos
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetSpecifiedGatewayInfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecifiedGatewayInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOGatewayInfos**](OperationResponseGridVOGatewayInfos.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidDetail1

> OperationResponseSsidDetailOpenApiVO GetSsidDetail1(ctx, omadacId, siteId, gatewayMac).Execute()

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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GetSsidDetail1(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetSsidDetail1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidDetail1`: OperationResponseSsidDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetSsidDetail1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSsidDetail1Request struct via the builder pattern


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


## GetWanStatus

> OperationResponseListGatewayWanStatusEntity GetWanStatus(ctx, omadacId, siteId, gatewayMac).Execute()

Get gateway wan status



### Example

```go
package main

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
	resp, r, err := apiClient.GatewayAPI.GetWanStatus(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetWanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWanStatus`: OperationResponseListGatewayWanStatusEntity
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetWanStatus`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListGatewayWanStatusEntity**](OperationResponseListGatewayWanStatusEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IspUpgrade

> OperationResponseWithoutResult IspUpgrade(ctx, omadacId, siteId, gatewayMac).GatewayIspUpgrade(gatewayIspUpgrade).Execute()

Upgrade isp



### Example

```go
package main

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
	gatewayIspUpgrade := *openapiclient.NewGatewayIspUpgrade() // GatewayIspUpgrade | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.IspUpgrade(context.Background(), omadacId, siteId, gatewayMac).GatewayIspUpgrade(gatewayIspUpgrade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.IspUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IspUpgrade`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.IspUpgrade`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIspUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gatewayIspUpgrade** | [**GatewayIspUpgrade**](GatewayIspUpgrade.md) |  | 

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


## ModifyConfigAdvanced

> OperationResponseWithoutResult ModifyConfigAdvanced(ctx, omadacId, siteId, gatewayMac).OsgConfigAdvancedOpenApiVO(osgConfigAdvancedOpenApiVO).Execute()

Modify gateway advanced config



### Example

```go
package main

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
	osgConfigAdvancedOpenApiVO := *openapiclient.NewOsgConfigAdvancedOpenApiVO() // OsgConfigAdvancedOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyConfigAdvanced(context.Background(), omadacId, siteId, gatewayMac).OsgConfigAdvancedOpenApiVO(osgConfigAdvancedOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyConfigAdvanced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigAdvanced`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyConfigAdvanced`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyConfigAdvancedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigAdvancedOpenApiVO** | [**OsgConfigAdvancedOpenApiVO**](OsgConfigAdvancedOpenApiVO.md) |  | 

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


## ModifyConfigCommonAdvanced

> OperationResponseWithoutResult ModifyConfigCommonAdvanced(ctx, omadacId, siteId, gatewayMac).OsgConfigCommonAdvancedOpenApiVO(osgConfigCommonAdvancedOpenApiVO).Execute()

Modify gateway advanced common config



### Example

```go
package main

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
	osgConfigCommonAdvancedOpenApiVO := *openapiclient.NewOsgConfigCommonAdvancedOpenApiVO() // OsgConfigCommonAdvancedOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyConfigCommonAdvanced(context.Background(), omadacId, siteId, gatewayMac).OsgConfigCommonAdvancedOpenApiVO(osgConfigCommonAdvancedOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyConfigCommonAdvanced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigCommonAdvanced`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyConfigCommonAdvanced`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyConfigCommonAdvancedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigCommonAdvancedOpenApiVO** | [**OsgConfigCommonAdvancedOpenApiVO**](OsgConfigCommonAdvancedOpenApiVO.md) |  | 

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


## ModifyConfigGeneral

> OperationResponseWithoutResult ModifyConfigGeneral(ctx, omadacId, siteId, gatewayMac).GatewayGeneralConfig(gatewayGeneralConfig).Execute()

Modify gateway general config



### Example

```go
package main

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
	gatewayGeneralConfig := *openapiclient.NewGatewayGeneralConfig(int32(123)) // GatewayGeneralConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyConfigGeneral(context.Background(), omadacId, siteId, gatewayMac).GatewayGeneralConfig(gatewayGeneralConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyConfigGeneral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigGeneral`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyConfigGeneral`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyConfigGeneralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gatewayGeneralConfig** | [**GatewayGeneralConfig**](GatewayGeneralConfig.md) |  | 

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


## ModifyConfigRadios

> OperationResponseWithoutResult ModifyConfigRadios(ctx, omadacId, siteId, gatewayMac).OsgConfigRadiosopenApiVO(osgConfigRadiosopenApiVO).Execute()

Modify gateway radios config



### Example

```go
package main

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
	osgConfigRadiosopenApiVO := *openapiclient.NewOsgConfigRadiosopenApiVO() // OsgConfigRadiosopenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyConfigRadios(context.Background(), omadacId, siteId, gatewayMac).OsgConfigRadiosopenApiVO(osgConfigRadiosopenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyConfigRadios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigRadios`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyConfigRadios`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyConfigRadiosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigRadiosopenApiVO** | [**OsgConfigRadiosopenApiVO**](OsgConfigRadiosopenApiVO.md) |  | 

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


## ModifyConfigServices

> OperationResponseWithoutResult ModifyConfigServices(ctx, omadacId, siteId, gatewayMac).OsgConfigServicesOpenApiVO(osgConfigServicesOpenApiVO).Execute()

Modify gateway config service



### Example

```go
package main

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
	osgConfigServicesOpenApiVO := *openapiclient.NewOsgConfigServicesOpenApiVO() // OsgConfigServicesOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyConfigServices(context.Background(), omadacId, siteId, gatewayMac).OsgConfigServicesOpenApiVO(osgConfigServicesOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyConfigServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigServices`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyConfigServices`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyConfigServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigServicesOpenApiVO** | [**OsgConfigServicesOpenApiVO**](OsgConfigServicesOpenApiVO.md) |  | 

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


## ModifyConfigWirelessAdvanced

> OperationResponseWithoutResult ModifyConfigWirelessAdvanced(ctx, omadacId, siteId, gatewayMac).OsgConfigWirelessAdvancedOpenApiVO(osgConfigWirelessAdvancedOpenApiVO).Execute()

Modify gateway advanced wireless config



### Example

```go
package main

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
	osgConfigWirelessAdvancedOpenApiVO := *openapiclient.NewOsgConfigWirelessAdvancedOpenApiVO() // OsgConfigWirelessAdvancedOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyConfigWirelessAdvanced(context.Background(), omadacId, siteId, gatewayMac).OsgConfigWirelessAdvancedOpenApiVO(osgConfigWirelessAdvancedOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyConfigWirelessAdvanced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigWirelessAdvanced`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyConfigWirelessAdvanced`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyConfigWirelessAdvancedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigWirelessAdvancedOpenApiVO** | [**OsgConfigWirelessAdvancedOpenApiVO**](OsgConfigWirelessAdvancedOpenApiVO.md) |  | 

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


## ModifyConfigWlans

> OperationResponseWithoutResult ModifyConfigWlans(ctx, omadacId, siteId, gatewayMac).OsgConfigWlansOpenApiVO(osgConfigWlansOpenApiVO).Execute()

Modify gateway wlans config



### Example

```go
package main

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
	osgConfigWlansOpenApiVO := *openapiclient.NewOsgConfigWlansOpenApiVO([]openapiclient.OsgSsidOverrideOpenApiVO{*openapiclient.NewOsgSsidOverrideOpenApiVO(false, "GlobalSsid_example", false, int32(123), int32(123), "Ssid_example", false, false)}) // OsgConfigWlansOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyConfigWlans(context.Background(), omadacId, siteId, gatewayMac).OsgConfigWlansOpenApiVO(osgConfigWlansOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyConfigWlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigWlans`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyConfigWlans`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyConfigWlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigWlansOpenApiVO** | [**OsgConfigWlansOpenApiVO**](OsgConfigWlansOpenApiVO.md) |  | 

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


## ModifyGeneralConfig1

> OperationResponseWithoutResult ModifyGeneralConfig1(ctx, omadacId, siteId, gatewayMac).GatewayGeneralConfig(gatewayGeneralConfig).Execute()

Modify gateway general config



### Example

```go
package main

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
	gatewayGeneralConfig := *openapiclient.NewGatewayGeneralConfig(int32(123)) // GatewayGeneralConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyGeneralConfig1(context.Background(), omadacId, siteId, gatewayMac).GatewayGeneralConfig(gatewayGeneralConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyGeneralConfig1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGeneralConfig1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyGeneralConfig1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyGeneralConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gatewayGeneralConfig** | [**GatewayGeneralConfig**](GatewayGeneralConfig.md) |  | 

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


## ModifyIpptLanIpv6

> OperationResponseWithoutResult ModifyIpptLanIpv6(ctx, omadacId, siteId, gatewayMac).Ipv6ForIpptOpenApiVO(ipv6ForIpptOpenApiVO).Execute()

Modify lan ipv6 config for ippt



### Example

```go
package main

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
	ipv6ForIpptOpenApiVO := *openapiclient.NewIpv6ForIpptOpenApiVO(false) // Ipv6ForIpptOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyIpptLanIpv6(context.Background(), omadacId, siteId, gatewayMac).Ipv6ForIpptOpenApiVO(ipv6ForIpptOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyIpptLanIpv6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpptLanIpv6`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyIpptLanIpv6`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyIpptLanIpv6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ipv6ForIpptOpenApiVO** | [**Ipv6ForIpptOpenApiVO**](Ipv6ForIpptOpenApiVO.md) |  | 

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


## ModifyOsgConfigForMsp

> OperationResponseWithoutResult ModifyOsgConfigForMsp(ctx, mspId, customerId, siteId, gatewayMac).GatewayOsgMspConfig(gatewayOsgMspConfig).Execute()

Modify gateway config for msp



### Example

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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	gatewayOsgMspConfig := *openapiclient.NewGatewayOsgMspConfig() // GatewayOsgMspConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyOsgConfigForMsp(context.Background(), mspId, customerId, siteId, gatewayMac).GatewayOsgMspConfig(gatewayOsgMspConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyOsgConfigForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOsgConfigForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyOsgConfigForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOsgConfigForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gatewayOsgMspConfig** | [**GatewayOsgMspConfig**](GatewayOsgMspConfig.md) |  | 

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


## ModifyOsgConfigGeneralForMsp

> OperationResponseWithoutResult ModifyOsgConfigGeneralForMsp(ctx, mspId, customerId, siteId, gatewayMac).GatewayMspConfigGeneral(gatewayMspConfigGeneral).Execute()

Modify general gateway config for msp



### Example

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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	gatewayMspConfigGeneral := *openapiclient.NewGatewayMspConfigGeneral() // GatewayMspConfigGeneral | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyOsgConfigGeneralForMsp(context.Background(), mspId, customerId, siteId, gatewayMac).GatewayMspConfigGeneral(gatewayMspConfigGeneral).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyOsgConfigGeneralForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOsgConfigGeneralForMsp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyOsgConfigGeneralForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOsgConfigGeneralForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gatewayMspConfigGeneral** | [**GatewayMspConfigGeneral**](GatewayMspConfigGeneral.md) |  | 

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


## ModifyPortConfig1

> OperationResponseGatewayPortSettingConfig ModifyPortConfig1(ctx, omadacId, siteId, gatewayMac, port).GatewayPortSettingConfig(gatewayPortSettingConfig).Execute()

Modify gateway port config



### Example

```go
package main

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
	port := "port_example" // string | Gateway port number
	gatewayPortSettingConfig := *openapiclient.NewGatewayPortSettingConfig() // GatewayPortSettingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifyPortConfig1(context.Background(), omadacId, siteId, gatewayMac, port).GatewayPortSettingConfig(gatewayPortSettingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifyPortConfig1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortConfig1`: OperationResponseGatewayPortSettingConfig
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifyPortConfig1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | Gateway port number | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gatewayPortSettingConfig** | [**GatewayPortSettingConfig**](GatewayPortSettingConfig.md) |  | 

### Return type

[**OperationResponseGatewayPortSettingConfig**](OperationResponseGatewayPortSettingConfig.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySshSettingByMac

> OperationResponseWithoutResult ModifySshSettingByMac(ctx, omadacId, siteId, gatewayMac).SSHSetting(sSHSetting).Execute()

Modify SSH setting by mac



### Example

```go
package main

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
	sSHSetting := *openapiclient.NewSSHSetting(false, int32(123)) // SSHSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifySshSettingByMac(context.Background(), omadacId, siteId, gatewayMac).SSHSetting(sSHSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifySshSettingByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySshSettingByMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifySshSettingByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifySshSettingByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sSHSetting** | [**SSHSetting**](SSHSetting.md) |  | 

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


## ModifySsidBasicConfig

> OperationResponseWithoutResult ModifySsidBasicConfig(ctx, omadacId, siteId, gatewayMac, ssidId).UpdateSsidBasicConfigForIpptOpenApiVO(updateSsidBasicConfigForIpptOpenApiVO).Execute()

modify SSID basic config



### Example

```go
package main

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
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidBasicConfigForIpptOpenApiVO := *openapiclient.NewUpdateSsidBasicConfigForIpptOpenApiVO(int32(123), false, "Name_example", int32(123)) // UpdateSsidBasicConfigForIpptOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifySsidBasicConfig(context.Background(), omadacId, siteId, gatewayMac, ssidId).UpdateSsidBasicConfigForIpptOpenApiVO(updateSsidBasicConfigForIpptOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifySsidBasicConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySsidBasicConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifySsidBasicConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySsidBasicConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidBasicConfigForIpptOpenApiVO** | [**UpdateSsidBasicConfigForIpptOpenApiVO**](UpdateSsidBasicConfigForIpptOpenApiVO.md) |  | 

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


## ModifySsidMacFilterConfig

> OperationResponseWithoutResult ModifySsidMacFilterConfig(ctx, omadacId, siteId, gatewayMac, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()

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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	ssidId := "ssidId_example" // string | SSID ID
	updateSsidMacFilterOpenApiVO := *openapiclient.NewUpdateSsidMacFilterOpenApiVO(false) // UpdateSsidMacFilterOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.ModifySsidMacFilterConfig(context.Background(), omadacId, siteId, gatewayMac, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.ModifySsidMacFilterConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySsidMacFilterConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.ModifySsidMacFilterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySsidMacFilterConfigRequest struct via the builder pattern


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


## MoveToSite1

> OperationResponseWithoutResult MoveToSite1(ctx, omadacId, siteId).MoveToSite(moveToSite).Execute()

Move devices to another site



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	moveToSite := *openapiclient.NewMoveToSite() // MoveToSite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.MoveToSite1(context.Background(), omadacId, siteId).MoveToSite(moveToSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.MoveToSite1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveToSite1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.MoveToSite1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveToSite1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **moveToSite** | [**MoveToSite**](MoveToSite.md) |  | 

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


## MspMoveToCustomer2

> OperationResponseWithoutResult MspMoveToCustomer2(ctx, mspId, customerId, siteId).MoveToCustomerOpenApiVO(moveToCustomerOpenApiVO).Execute()

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
	resp, r, err := apiClient.GatewayAPI.MspMoveToCustomer2(context.Background(), mspId, customerId, siteId).MoveToCustomerOpenApiVO(moveToCustomerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.MspMoveToCustomer2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MspMoveToCustomer2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.MspMoveToCustomer2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiMspMoveToCustomer2Request struct via the builder pattern


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


## RecoveryPoePort

> OperationResponseOperationResponseWithoutResult RecoveryPoePort(ctx, siteId, gatewayMac, port, omadacId).Execute()

Recovery gateway poe port



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	siteId := "siteId_example" // string | Site ID
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	port := "port_example" // string | Gateway port number
	omadacId := "omadacId_example" // string | omadacId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.RecoveryPoePort(context.Background(), siteId, gatewayMac, port, omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.RecoveryPoePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecoveryPoePort`: OperationResponseOperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.RecoveryPoePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 
**port** | **string** | Gateway port number | 
**omadacId** | **string** | omadacId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryPoePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseOperationResponseWithoutResult**](OperationResponseOperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIspFile

> OperationResponseWithoutResult UploadIspFile(ctx, omadacId, siteId, gatewayMac).UploadSSLKeyRequest(uploadSSLKeyRequest).Execute()

Upload isp upgrade file



### Example

```go
package main

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
	uploadSSLKeyRequest := *openapiclient.NewUploadSSLKeyRequest("TODO") // UploadSSLKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.UploadIspFile(context.Background(), omadacId, siteId, gatewayMac).UploadSSLKeyRequest(uploadSSLKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.UploadIspFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadIspFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.UploadIspFile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUploadIspFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **uploadSSLKeyRequest** | [**UploadSSLKeyRequest**](UploadSSLKeyRequest.md) |  | 

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


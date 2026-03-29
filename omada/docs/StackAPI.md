# \StackAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchApplyStackNetworks**](StackAPI.md#BatchApplyStackNetworks) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/networks | Batch modify stack vlan interface status
[**BatchModifyOswStackPorts**](StackAPI.md#BatchModifyOswStackPorts) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/ports/config | Batch modify stack port
[**BatchModifySwitchStackPortSetting**](StackAPI.md#BatchModifySwitchStackPortSetting) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/stacks/{stackId}/multi-ports/config | Batch modify stack port
[**CancelStackCableTest**](StackAPI.md#CancelStackCableTest) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/stacks/{stackId}/cancel | Cancel stack cable test
[**CreateOswStack**](StackAPI.md#CreateOswStack) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stacks | Create Switch Stack
[**CreateStackStaticRouting**](StackAPI.md#CreateStackStaticRouting) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/staticRoutings | Create stack staticRouting
[**CreateStackVrf**](StackAPI.md#CreateStackVrf) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/vrfs | Create new stack vrf
[**DeleteOswStack**](StackAPI.md#DeleteOswStack) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId} | Delete Switch Stack
[**DeleteOswStackLag**](StackAPI.md#DeleteOswStackLag) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/lags/{lagId} | Delete stack lag
[**DeleteStackStaticRouting**](StackAPI.md#DeleteStackStaticRouting) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/staticRoutings/{staticRoutingId} | Delete stack staticRouting
[**DeleteStackVrf**](StackAPI.md#DeleteStackVrf) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/vrfs/{vrfId} | Delete an existing stack vrf
[**DeleteSwitchStackLagSetting**](StackAPI.md#DeleteSwitchStackLagSetting) | **Delete** /openapi/v2/{omadacId}/sites/{siteId}/stacks/{stackId}/lags/{lagId} | Delete stack lag
[**DetectSwitchStackMembers**](StackAPI.md#DetectSwitchStackMembers) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stacks/detect/{stackId} | Detect switch stack members
[**DetectSwitchStackMembersForAllStacks**](StackAPI.md#DetectSwitchStackMembersForAllStacks) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stacks/detect | Detect switch stacks members for all stacks
[**DownloadStackDeviceInfo**](StackAPI.md#DownloadStackDeviceInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stack/{stackId}/downloadDeviceInfo | Download stack device info.
[**ForceProvisionStack**](StackAPI.md#ForceProvisionStack) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/stack/{stackId}/forceProvision | Force provision stack
[**ForgetStack**](StackAPI.md#ForgetStack) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/stack/{stackId}/forget | Forget stack
[**GetGridDiscoveryStackList**](StackAPI.md#GetGridDiscoveryStackList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/discovery | Get discovery stack list
[**GetGridOswStackList**](StackAPI.md#GetGridOswStackList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks | Get stack list
[**GetGridStackDhcpUserList**](StackAPI.md#GetGridStackDhcpUserList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stack/{stackId}/dhcp/user-list | Get stack dhcp user list
[**GetGridStackLldpNeighborTable**](StackAPI.md#GetGridStackLldpNeighborTable) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stack/{stackId}/lldp-neighbors | Get stack lldp neighbor table
[**GetGridStackOspfNeighborTable**](StackAPI.md#GetGridStackOspfNeighborTable) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stack/{stackId}/ospf-neighbors | Get stack ospf neighbor table
[**GetGridStackPortAndLagNetworks**](StackAPI.md#GetGridStackPortAndLagNetworks) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/port-lag-networks | Get the networks used on stack&#39;s ports and LAGs
[**GetGridStackPortAndLagNetworksDetail**](StackAPI.md#GetGridStackPortAndLagNetworksDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/networks-detail | Get the networks detail (including the vlan related ports and lags info) used on stack&#39;s ports and LAGs
[**GetGridStackStaticRouting**](StackAPI.md#GetGridStackStaticRouting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/staticRoutings | Get grid stack staticRouting
[**GetMspOswStackDetail**](StackAPI.md#GetMspOswStackDetail) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/stack/{stackId} | Get msp stack detail
[**GetMspOswStackList**](StackAPI.md#GetMspOswStackList) | **Get** /openapi/v1/msp/{mspId}/stack | Get msp stack List
[**GetOswStackDDMInfo**](StackAPI.md#GetOswStackDDMInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stack/{stackId}/ddm/info | Get stack ddm info.
[**GetOswStackDetail**](StackAPI.md#GetOswStackDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId} | Get stack detail
[**GetOswStackLag**](StackAPI.md#GetOswStackLag) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/lags/{lagId} | Get stack lag
[**GetOswStackLagList**](StackAPI.md#GetOswStackLagList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/lags | Get stack lag List
[**GetOswStackPortList**](StackAPI.md#GetOswStackPortList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/ports | Get stack port List
[**GetStackCableTestFullResults**](StackAPI.md#GetStackCableTestFullResults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/stacks/{stackId}/full-results | Get stack cable test full results
[**GetStackCableTestIncrementResults**](StackAPI.md#GetStackCableTestIncrementResults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/stacks/{stackId}/increment-results | Get stack cable test increment results
[**GetStackCableTestLogs**](StackAPI.md#GetStackCableTestLogs) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/stacks/{stackId}/logs | Get stack cable test logs
[**GetStackCableTestPorts**](StackAPI.md#GetStackCableTestPorts) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/stacks/{stackId}/ports | Get the port list of stack used for cable test
[**GetStackGridVrf**](StackAPI.md#GetStackGridVrf) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/vrfs | Get stack vrf page
[**GetStackNetworkList**](StackAPI.md#GetStackNetworkList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/networks | Get stack vlan interface List
[**GetStackPortAndLagNetwork**](StackAPI.md#GetStackPortAndLagNetwork) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/port-lag-networks/{networkId}/vlan/{vlan} | Get the stack&#39;s ports and LAGs that the network affects
[**GetStackRememberMe**](StackAPI.md#GetStackRememberMe) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/remember | Get stack remember Config
[**GetStackableSwitches**](StackAPI.md#GetStackableSwitches) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stacks/switches | Get stackable switches
[**LocateOswStack**](StackAPI.md#LocateOswStack) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/stacks/{stackId}/locate | Locate switch stack
[**ModifyOswStack**](StackAPI.md#ModifyOswStack) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId} | Modify Switch Stack
[**ModifyOswStackDetail**](StackAPI.md#ModifyOswStackDetail) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/config | Modify stack detail
[**ModifyOswStackLag**](StackAPI.md#ModifyOswStackLag) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/lags/{lagId} | Modify stack lag
[**ModifyOswStackPort**](StackAPI.md#ModifyOswStackPort) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/ports/{port} | Modify stack port
[**ModifyStackLoopbackControl**](StackAPI.md#ModifyStackLoopbackControl) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/config/loopback | Modify stack loopback control
[**ModifyStackNetwork**](StackAPI.md#ModifyStackNetwork) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/networks/{networkId} | Modify stack vlan interface
[**ModifyStackRememberMe**](StackAPI.md#ModifyStackRememberMe) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/remember | Modify stack remember Config
[**ModifyStackStaticRouting**](StackAPI.md#ModifyStackStaticRouting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/staticRoutings/{staticRoutingId} | Modify stack staticRouting
[**ModifyStackVrf**](StackAPI.md#ModifyStackVrf) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/stacks/{stackId}/vrfs/{vrfId} | Modify an existing stack vrf
[**ModifySwitchStackLagSetting**](StackAPI.md#ModifySwitchStackLagSetting) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/stacks/{stackId}/lags/{lagId} | Modify stack lag
[**ModifySwitchStackPortSetting**](StackAPI.md#ModifySwitchStackPortSetting) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/stacks/{stackId}/ports/{port} | Modify stack port
[**RebootOswStack**](StackAPI.md#RebootOswStack) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/stacks/{stackId}/reboot | Reboot switch stack
[**StartStackCableTest**](StackAPI.md#StartStackCableTest) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cable-test/stacks/{stackId}/start | Start stack cable test



## BatchApplyStackNetworks

> OperationResponseWithoutResult BatchApplyStackNetworks(ctx, omadacId, siteId, stackId).BatchApplyStackNetworkOpenApiVO(batchApplyStackNetworkOpenApiVO).Execute()

Batch modify stack vlan interface status



### Example

```go
package main

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
	batchApplyStackNetworkOpenApiVO := *openapiclient.NewBatchApplyStackNetworkOpenApiVO([]openapiclient.OswNetworkBaseOpenApi{*openapiclient.NewOswNetworkBaseOpenApi("Id_example", int32(123))}) // BatchApplyStackNetworkOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.BatchApplyStackNetworks(context.Background(), omadacId, siteId, stackId).BatchApplyStackNetworkOpenApiVO(batchApplyStackNetworkOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.BatchApplyStackNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchApplyStackNetworks`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.BatchApplyStackNetworks`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchApplyStackNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchApplyStackNetworkOpenApiVO** | [**BatchApplyStackNetworkOpenApiVO**](BatchApplyStackNetworkOpenApiVO.md) |  | 

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


## BatchModifyOswStackPorts

> OperationResponseWithoutResult BatchModifyOswStackPorts(ctx, omadacId, siteId, stackId).BatchStackPortSettingOpenApiVO(batchStackPortSettingOpenApiVO).Execute()

Batch modify stack port



### Example

```go
package main

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
	batchStackPortSettingOpenApiVO := *openapiclient.NewBatchStackPortSettingOpenApiVO([]openapiclient.OswStackUnitVO{*openapiclient.NewOswStackUnitVO("Mac_example", []openapiclient.OswStandPortVO{*openapiclient.NewOswStandPortVO(int32(123), int32(123), int32(123))}, int32(123))}) // BatchStackPortSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.BatchModifyOswStackPorts(context.Background(), omadacId, siteId, stackId).BatchStackPortSettingOpenApiVO(batchStackPortSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.BatchModifyOswStackPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyOswStackPorts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.BatchModifyOswStackPorts`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchModifyOswStackPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchStackPortSettingOpenApiVO** | [**BatchStackPortSettingOpenApiVO**](BatchStackPortSettingOpenApiVO.md) |  | 

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


## BatchModifySwitchStackPortSetting

> OperationResponseWithoutResult BatchModifySwitchStackPortSetting(ctx, omadacId, siteId, stackId).BatchStackPortSettingVO(batchStackPortSettingVO).Execute()

Batch modify stack port



### Example

```go
package main

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
	batchStackPortSettingVO := *openapiclient.NewBatchStackPortSettingVO([]openapiclient.OswStackUnitVO{*openapiclient.NewOswStackUnitVO("Mac_example", []openapiclient.OswStandPortVO{*openapiclient.NewOswStandPortVO(int32(123), int32(123), int32(123))}, int32(123))}) // BatchStackPortSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.BatchModifySwitchStackPortSetting(context.Background(), omadacId, siteId, stackId).BatchStackPortSettingVO(batchStackPortSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.BatchModifySwitchStackPortSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifySwitchStackPortSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.BatchModifySwitchStackPortSetting`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiBatchModifySwitchStackPortSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchStackPortSettingVO** | [**BatchStackPortSettingVO**](BatchStackPortSettingVO.md) |  | 

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


## CancelStackCableTest

> OperationResponseWithoutResult CancelStackCableTest(ctx, omadacId, siteId, stackId).Execute()

Cancel stack cable test



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.CancelStackCableTest(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.CancelStackCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelStackCableTest`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.CancelStackCableTest`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCancelStackCableTestRequest struct via the builder pattern


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


## CreateOswStack

> OperationResponseWithoutResult CreateOswStack(ctx, omadacId, siteId).OswStackConfigOpenApiVO(oswStackConfigOpenApiVO).Execute()

Create Switch Stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	oswStackConfigOpenApiVO := *openapiclient.NewOswStackConfigOpenApiVO([]openapiclient.OswStackMemberVO{*openapiclient.NewOswStackMemberVO("Mac_example", int32(123), []openapiclient.OswStackPortGroupVO{*openapiclient.NewOswStackPortGroupVO(int32(123), []openapiclient.OswStandPortVO{*openapiclient.NewOswStandPortVO(int32(123), int32(123), int32(123))})}, int32(123))}, "Name_example") // OswStackConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.CreateOswStack(context.Background(), omadacId, siteId).OswStackConfigOpenApiVO(oswStackConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.CreateOswStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOswStack`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.CreateOswStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOswStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oswStackConfigOpenApiVO** | [**OswStackConfigOpenApiVO**](OswStackConfigOpenApiVO.md) |  | 

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


## CreateStackStaticRouting

> OperationResponseWithoutResult CreateStackStaticRouting(ctx, omadacId, siteId, stackId).OswStaticRoutingConfigOpenApiVO(oswStaticRoutingConfigOpenApiVO).Execute()

Create stack staticRouting



### Example

```go
package main

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
	oswStaticRoutingConfigOpenApiVO := *openapiclient.NewOswStaticRoutingConfigOpenApiVO([]string{"Destinations_example"}, int32(123), int32(123), "NextHopIp_example", false) // OswStaticRoutingConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.CreateStackStaticRouting(context.Background(), omadacId, siteId, stackId).OswStaticRoutingConfigOpenApiVO(oswStaticRoutingConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.CreateStackStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.CreateStackStaticRouting`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateStackStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswStaticRoutingConfigOpenApiVO** | [**OswStaticRoutingConfigOpenApiVO**](OswStaticRoutingConfigOpenApiVO.md) |  | 

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


## CreateStackVrf

> OperationResponseWithoutResult CreateStackVrf(ctx, omadacId, siteId, stackId).OswVrfVO(oswVrfVO).Execute()

Create new stack vrf



### Example

```go
package main

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
	oswVrfVO := *openapiclient.NewOswVrfVO(false, false, "Vrf_example") // OswVrfVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.CreateStackVrf(context.Background(), omadacId, siteId, stackId).OswVrfVO(oswVrfVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.CreateStackVrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackVrf`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.CreateStackVrf`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateStackVrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswVrfVO** | [**OswVrfVO**](OswVrfVO.md) |  | 

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


## DeleteOswStack

> OperationResponseWithoutResult DeleteOswStack(ctx, omadacId, siteId, stackId).Execute()

Delete Switch Stack



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.DeleteOswStack(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.DeleteOswStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOswStack`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.DeleteOswStack`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteOswStackRequest struct via the builder pattern


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


## DeleteOswStackLag

> OperationResponseWithoutResult DeleteOswStackLag(ctx, omadacId, siteId, stackId, lagId).Execute()

Delete stack lag



### Example

```go
package main

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
	lagId := "lagId_example" // string | LAG ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.DeleteOswStackLag(context.Background(), omadacId, siteId, stackId, lagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.DeleteOswStackLag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOswStackLag`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.DeleteOswStackLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**lagId** | **string** | LAG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOswStackLagRequest struct via the builder pattern


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


## DeleteStackStaticRouting

> OperationResponseWithoutResult DeleteStackStaticRouting(ctx, omadacId, siteId, stackId, staticRoutingId).Execute()

Delete stack staticRouting



### Example

```go
package main

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
	staticRoutingId := "staticRoutingId_example" // string | Static routing ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.DeleteStackStaticRouting(context.Background(), omadacId, siteId, stackId, staticRoutingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.DeleteStackStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.DeleteStackStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**staticRoutingId** | **string** | Static routing ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackStaticRoutingRequest struct via the builder pattern


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


## DeleteStackVrf

> OperationResponseWithoutResult DeleteStackVrf(ctx, omadacId, siteId, stackId, vrfId).Execute()

Delete an existing stack vrf



### Example

```go
package main

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
	vrfId := "vrfId_example" // string | VRF ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.DeleteStackVrf(context.Background(), omadacId, siteId, stackId, vrfId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.DeleteStackVrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackVrf`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.DeleteStackVrf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**vrfId** | **string** | VRF ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackVrfRequest struct via the builder pattern


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


## DeleteSwitchStackLagSetting

> OperationResponseWithoutResult DeleteSwitchStackLagSetting(ctx, omadacId, siteId, stackId, lagId).Execute()

Delete stack lag



### Example

```go
package main

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
	lagId := "lagId_example" // string | Lag ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.DeleteSwitchStackLagSetting(context.Background(), omadacId, siteId, stackId, lagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.DeleteSwitchStackLagSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSwitchStackLagSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.DeleteSwitchStackLagSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**lagId** | **string** | Lag ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSwitchStackLagSettingRequest struct via the builder pattern


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


## DetectSwitchStackMembers

> OperationResponseWithoutResult DetectSwitchStackMembers(ctx, omadacId, siteId, stackId).Execute()

Detect switch stack members



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.DetectSwitchStackMembers(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.DetectSwitchStackMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetectSwitchStackMembers`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.DetectSwitchStackMembers`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDetectSwitchStackMembersRequest struct via the builder pattern


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


## DetectSwitchStackMembersForAllStacks

> OperationResponseWithoutResult DetectSwitchStackMembersForAllStacks(ctx, omadacId, siteId).Execute()

Detect switch stacks members for all stacks



### Example

```go
package main

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
	resp, r, err := apiClient.StackAPI.DetectSwitchStackMembersForAllStacks(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.DetectSwitchStackMembersForAllStacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetectSwitchStackMembersForAllStacks`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.DetectSwitchStackMembersForAllStacks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetectSwitchStackMembersForAllStacksRequest struct via the builder pattern


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


## DownloadStackDeviceInfo

> map[string]interface{} DownloadStackDeviceInfo(ctx, omadacId, siteId, stackId).Execute()

Download stack device info.



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.DownloadStackDeviceInfo(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.DownloadStackDeviceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadStackDeviceInfo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.DownloadStackDeviceInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDownloadStackDeviceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForceProvisionStack

> OperationResponseWithoutResult ForceProvisionStack(ctx, omadacId, siteId, stackId).Execute()

Force provision stack



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ForceProvisionStack(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ForceProvisionStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForceProvisionStack`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ForceProvisionStack`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiForceProvisionStackRequest struct via the builder pattern


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


## ForgetStack

> OperationResponseWithoutResult ForgetStack(ctx, omadacId, siteId, stackId).Execute()

Forget stack



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ForgetStack(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ForgetStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgetStack`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ForgetStack`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiForgetStackRequest struct via the builder pattern


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


## GetGridDiscoveryStackList

> OperationResponseWithoutResult GetGridDiscoveryStackList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get discovery stack list



### Example

```go
package main

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
	resp, r, err := apiClient.StackAPI.GetGridDiscoveryStackList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetGridDiscoveryStackList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDiscoveryStackList`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetGridDiscoveryStackList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDiscoveryStackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

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


## GetGridOswStackList

> OperationResponseGridVOOswStackVO GetGridOswStackList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get stack list



### Example

```go
package main

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
	resp, r, err := apiClient.StackAPI.GetGridOswStackList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetGridOswStackList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridOswStackList`: OperationResponseGridVOOswStackVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetGridOswStackList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridOswStackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOswStackVO**](OperationResponseGridVOOswStackVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStackDhcpUserList

> OperationResponseDhcpUserGridVODhcpUserVO GetGridStackDhcpUserList(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()

Get stack dhcp user list



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetGridStackDhcpUserList(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetGridStackDhcpUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStackDhcpUserList`: OperationResponseDhcpUserGridVODhcpUserVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetGridStackDhcpUserList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridStackDhcpUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseDhcpUserGridVODhcpUserVO**](OperationResponseDhcpUserGridVODhcpUserVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStackLldpNeighborTable

> OperationResponseGridVOOswLldpNeighborVO GetGridStackLldpNeighborTable(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).SortsStandardPort(sortsStandardPort).SearchKey(searchKey).Execute()

Get stack lldp neighbor table



### Example

```go
package main

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
	sortsStandardPort := "sortsStandardPort_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field portId,standardPort,deviceId,systemName (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetGridStackLldpNeighborTable(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).SortsStandardPort(sortsStandardPort).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetGridStackLldpNeighborTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStackLldpNeighborTable`: OperationResponseGridVOOswLldpNeighborVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetGridStackLldpNeighborTable`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridStackLldpNeighborTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsStandardPort** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field portId,standardPort,deviceId,systemName | 

### Return type

[**OperationResponseGridVOOswLldpNeighborVO**](OperationResponseGridVOOswLldpNeighborVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStackOspfNeighborTable

> OperationResponseOswOspfNeighborGridVOOswOspfNeighborVO GetGridStackOspfNeighborTable(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).SortsNeighborInterface(sortsNeighborInterface).FiltersNeighborInterface(filtersNeighborInterface).FiltersProcessId(filtersProcessId).SearchKey(searchKey).Execute()

Get stack ospf neighbor table



### Example

```go
package main

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
	sortsNeighborInterface := "sortsNeighborInterface_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersNeighborInterface := []string{"Inner_example"} // []string | Filter query parameters, support field neighborInterface (optional)
	filtersProcessId := "filtersProcessId_example" // string | Filter query parameters, support field processId (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field neighborIp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetGridStackOspfNeighborTable(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).SortsNeighborInterface(sortsNeighborInterface).FiltersNeighborInterface(filtersNeighborInterface).FiltersProcessId(filtersProcessId).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetGridStackOspfNeighborTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStackOspfNeighborTable`: OperationResponseOswOspfNeighborGridVOOswOspfNeighborVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetGridStackOspfNeighborTable`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridStackOspfNeighborTableRequest struct via the builder pattern


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

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStackPortAndLagNetworks

> GridVOOswNetworkBriefInfoOpenApiVO GetGridStackPortAndLagNetworks(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()

Get the networks used on stack's ports and LAGs



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetGridStackPortAndLagNetworks(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetGridStackPortAndLagNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStackPortAndLagNetworks`: GridVOOswNetworkBriefInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetGridStackPortAndLagNetworks`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridStackPortAndLagNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOOswNetworkBriefInfoOpenApiVO**](GridVOOswNetworkBriefInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStackPortAndLagNetworksDetail

> GridVOOswNetworkDetailInfoOpenApiVO GetGridStackPortAndLagNetworksDetail(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()

Get the networks detail (including the vlan related ports and lags info) used on stack's ports and LAGs



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetGridStackPortAndLagNetworksDetail(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetGridStackPortAndLagNetworksDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStackPortAndLagNetworksDetail`: GridVOOswNetworkDetailInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetGridStackPortAndLagNetworksDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridStackPortAndLagNetworksDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOOswNetworkDetailInfoOpenApiVO**](GridVOOswNetworkDetailInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridStackStaticRouting

> OperationResponseOswStaticRoutingGridVOOswStaticRoutingVO GetGridStackStaticRouting(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()

Get grid stack staticRouting



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetGridStackStaticRouting(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetGridStackStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridStackStaticRouting`: OperationResponseOswStaticRoutingGridVOOswStaticRoutingVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetGridStackStaticRouting`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridStackStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseOswStaticRoutingGridVOOswStaticRoutingVO**](OperationResponseOswStaticRoutingGridVOOswStaticRoutingVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspOswStackDetail

> OperationResponseWithoutResult GetMspOswStackDetail(ctx, mspId, customerId, siteId, stackId).Execute()

Get msp stack detail



### Example

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
	customerId := "customerId_example" // string | customerId
	siteId := "siteId_example" // string | Site ID
	stackId := "stackId_example" // string | Stack ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetMspOswStackDetail(context.Background(), mspId, customerId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetMspOswStackDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOswStackDetail`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetMspOswStackDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | customerId | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOswStackDetailRequest struct via the builder pattern


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


## GetMspOswStackList

> OperationResponseWithoutResult GetMspOswStackList(ctx, mspId).Page(page).PageSize(pageSize).Execute()

Get msp stack List



### Example

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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetMspOswStackList(context.Background(), mspId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetMspOswStackList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOswStackList`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetMspOswStackList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOswStackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

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


## GetOswStackDDMInfo

> OperationResponseListOswDDMInfoOpenApiVO GetOswStackDDMInfo(ctx, omadacId, siteId, stackId).Execute()

Get stack ddm info.



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetOswStackDDMInfo(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetOswStackDDMInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswStackDDMInfo`: OperationResponseListOswDDMInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetOswStackDDMInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOswStackDDMInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListOswDDMInfoOpenApiVO**](OperationResponseListOswDDMInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswStackDetail

> OperationResponseOswStackDetailVO GetOswStackDetail(ctx, omadacId, siteId, stackId).Execute()

Get stack detail



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetOswStackDetail(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetOswStackDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswStackDetail`: OperationResponseOswStackDetailVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetOswStackDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOswStackDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOswStackDetailVO**](OperationResponseOswStackDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswStackLag

> OperationResponseOswStackMemberLagVO GetOswStackLag(ctx, omadacId, siteId, stackId, lagId).Execute()

Get stack lag



### Example

```go
package main

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
	lagId := "lagId_example" // string | LAG ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetOswStackLag(context.Background(), omadacId, siteId, stackId, lagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetOswStackLag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswStackLag`: OperationResponseOswStackMemberLagVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetOswStackLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**lagId** | **string** | LAG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswStackLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseOswStackMemberLagVO**](OperationResponseOswStackMemberLagVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswStackLagList

> OperationResponseListOswStackMemberLagVO GetOswStackLagList(ctx, omadacId, siteId, stackId).Execute()

Get stack lag List



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetOswStackLagList(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetOswStackLagList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswStackLagList`: OperationResponseListOswStackMemberLagVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetOswStackLagList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOswStackLagListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListOswStackMemberLagVO**](OperationResponseListOswStackMemberLagVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswStackPortList

> OperationResponseListOswStackPortVO GetOswStackPortList(ctx, omadacId, siteId, stackId).Execute()

Get stack port List



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetOswStackPortList(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetOswStackPortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswStackPortList`: OperationResponseListOswStackPortVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetOswStackPortList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOswStackPortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListOswStackPortVO**](OperationResponseListOswStackPortVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackCableTestFullResults

> OperationResponseOswCableTestResultWithStatusVO GetStackCableTestFullResults(ctx, omadacId, siteId, stackId).Execute()

Get stack cable test full results



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackCableTestFullResults(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackCableTestFullResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackCableTestFullResults`: OperationResponseOswCableTestResultWithStatusVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackCableTestFullResults`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStackCableTestFullResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOswCableTestResultWithStatusVO**](OperationResponseOswCableTestResultWithStatusVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackCableTestIncrementResults

> OperationResponseOswCableTestResultWithStatusVO GetStackCableTestIncrementResults(ctx, omadacId, siteId, stackId).Execute()

Get stack cable test increment results



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackCableTestIncrementResults(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackCableTestIncrementResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackCableTestIncrementResults`: OperationResponseOswCableTestResultWithStatusVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackCableTestIncrementResults`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStackCableTestIncrementResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOswCableTestResultWithStatusVO**](OperationResponseOswCableTestResultWithStatusVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackCableTestLogs

> OperationResponseListOswCableTestLogOpenApiVO GetStackCableTestLogs(ctx, omadacId, siteId, stackId).Execute()

Get stack cable test logs



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackCableTestLogs(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackCableTestLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackCableTestLogs`: OperationResponseListOswCableTestLogOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackCableTestLogs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStackCableTestLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListOswCableTestLogOpenApiVO**](OperationResponseListOswCableTestLogOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackCableTestPorts

> OperationResponseOswStackCableTestVO GetStackCableTestPorts(ctx, omadacId, siteId, stackId).Execute()

Get the port list of stack used for cable test



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackCableTestPorts(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackCableTestPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackCableTestPorts`: OperationResponseOswStackCableTestVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackCableTestPorts`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStackCableTestPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOswStackCableTestVO**](OperationResponseOswStackCableTestVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackGridVrf

> GridVOOswVrfVO GetStackGridVrf(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()

Get stack vrf page



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackGridVrf(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackGridVrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackGridVrf`: GridVOOswVrfVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackGridVrf`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStackGridVrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**GridVOOswVrfVO**](GridVOOswVrfVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackNetworkList

> OperationResponseOswNetworkGridOswNetworkOpenApi GetStackNetworkList(ctx, omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()

Get stack vlan interface List



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackNetworkList(context.Background(), omadacId, siteId, stackId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackNetworkList`: OperationResponseOswNetworkGridOswNetworkOpenApi
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackNetworkList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStackNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseOswNetworkGridOswNetworkOpenApi**](OperationResponseOswNetworkGridOswNetworkOpenApi.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackPortAndLagNetwork

> OswPortAndLagNetworkVO GetStackPortAndLagNetwork(ctx, omadacId, siteId, stackId, networkId, vlan).Execute()

Get the stack's ports and LAGs that the network affects



### Example

```go
package main

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
	networkId := "networkId_example" // string | Switch network ID.
	vlan := "vlan_example" // string | VLAN.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackPortAndLagNetwork(context.Background(), omadacId, siteId, stackId, networkId, vlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackPortAndLagNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackPortAndLagNetwork`: OswPortAndLagNetworkVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackPortAndLagNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**networkId** | **string** | Switch network ID. | 
**vlan** | **string** | VLAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackPortAndLagNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**OswPortAndLagNetworkVO**](OswPortAndLagNetworkVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackRememberMe

> OperationResponseDeviceRememberConfig GetStackRememberMe(ctx, omadacId, siteId, stackId).Execute()

Get stack remember Config



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackRememberMe(context.Background(), omadacId, siteId, stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackRememberMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackRememberMe`: OperationResponseDeviceRememberConfig
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackRememberMe`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStackRememberMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseDeviceRememberConfig**](OperationResponseDeviceRememberConfig.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackableSwitches

> OperationResponseOswStackSwitchVO GetStackableSwitches(ctx, omadacId, siteId).StackId(stackId).Execute()

Get stackable switches



### Example

```go
package main

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.GetStackableSwitches(context.Background(), omadacId, siteId).StackId(stackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.GetStackableSwitches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackableSwitches`: OperationResponseOswStackSwitchVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.GetStackableSwitches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackableSwitchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stackId** | **string** | Stack ID | 

### Return type

[**OperationResponseOswStackSwitchVO**](OperationResponseOswStackSwitchVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocateOswStack

> OperationResponseLocateResultVO LocateOswStack(ctx, omadacId, siteId, stackId).OswStackLocateOpenApiVO(oswStackLocateOpenApiVO).Execute()

Locate switch stack



### Example

```go
package main

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
	oswStackLocateOpenApiVO := *openapiclient.NewOswStackLocateOpenApiVO(false) // OswStackLocateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.LocateOswStack(context.Background(), omadacId, siteId, stackId).OswStackLocateOpenApiVO(oswStackLocateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.LocateOswStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocateOswStack`: OperationResponseLocateResultVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.LocateOswStack`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiLocateOswStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswStackLocateOpenApiVO** | [**OswStackLocateOpenApiVO**](OswStackLocateOpenApiVO.md) |  | 

### Return type

[**OperationResponseLocateResultVO**](OperationResponseLocateResultVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOswStack

> OperationResponseWithoutResult ModifyOswStack(ctx, omadacId, siteId, stackId).OswStackConfigOpenApiVO(oswStackConfigOpenApiVO).Execute()

Modify Switch Stack



### Example

```go
package main

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
	oswStackConfigOpenApiVO := *openapiclient.NewOswStackConfigOpenApiVO([]openapiclient.OswStackMemberVO{*openapiclient.NewOswStackMemberVO("Mac_example", int32(123), []openapiclient.OswStackPortGroupVO{*openapiclient.NewOswStackPortGroupVO(int32(123), []openapiclient.OswStandPortVO{*openapiclient.NewOswStandPortVO(int32(123), int32(123), int32(123))})}, int32(123))}, "Name_example") // OswStackConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyOswStack(context.Background(), omadacId, siteId, stackId).OswStackConfigOpenApiVO(oswStackConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyOswStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswStack`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyOswStack`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyOswStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswStackConfigOpenApiVO** | [**OswStackConfigOpenApiVO**](OswStackConfigOpenApiVO.md) |  | 

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


## ModifyOswStackDetail

> OperationResponseWithoutResult ModifyOswStackDetail(ctx, omadacId, siteId, stackId).OswStackDetailConfigOpenApiVO(oswStackDetailConfigOpenApiVO).Execute()

Modify stack detail



### Example

```go
package main

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
	oswStackDetailConfigOpenApiVO := *openapiclient.NewOswStackDetailConfigOpenApiVO() // OswStackDetailConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyOswStackDetail(context.Background(), omadacId, siteId, stackId).OswStackDetailConfigOpenApiVO(oswStackDetailConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyOswStackDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswStackDetail`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyOswStackDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyOswStackDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswStackDetailConfigOpenApiVO** | [**OswStackDetailConfigOpenApiVO**](OswStackDetailConfigOpenApiVO.md) |  | 

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


## ModifyOswStackLag

> OperationResponseWithoutResult ModifyOswStackLag(ctx, omadacId, siteId, stackId, lagId).OswStackMemberLagOpenApiVO(oswStackMemberLagOpenApiVO).Execute()

Modify stack lag



### Example

```go
package main

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
	lagId := "lagId_example" // string | LAG ID
	oswStackMemberLagOpenApiVO := *openapiclient.NewOswStackMemberLagOpenApiVO(false) // OswStackMemberLagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyOswStackLag(context.Background(), omadacId, siteId, stackId, lagId).OswStackMemberLagOpenApiVO(oswStackMemberLagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyOswStackLag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswStackLag`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyOswStackLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**lagId** | **string** | LAG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswStackLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswStackMemberLagOpenApiVO** | [**OswStackMemberLagOpenApiVO**](OswStackMemberLagOpenApiVO.md) |  | 

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


## ModifyOswStackPort

> OperationResponseWithoutResult ModifyOswStackPort(ctx, omadacId, siteId, stackId, port).OswPortSettingOpenApiVO(oswPortSettingOpenApiVO).Execute()

Modify stack port



### Example

```go
package main

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
	port := "port_example" // string | Port
	oswPortSettingOpenApiVO := *openapiclient.NewOswPortSettingOpenApiVO(false) // OswPortSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyOswStackPort(context.Background(), omadacId, siteId, stackId, port).OswPortSettingOpenApiVO(oswPortSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyOswStackPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswStackPort`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyOswStackPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**port** | **string** | Port | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswStackPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswPortSettingOpenApiVO** | [**OswPortSettingOpenApiVO**](OswPortSettingOpenApiVO.md) |  | 

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


## ModifyStackLoopbackControl

> OperationResponseWithoutResult ModifyStackLoopbackControl(ctx, omadacId, siteId, stackId).SwitchLoopbackControl(switchLoopbackControl).Execute()

Modify stack loopback control



### Example

```go
package main

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
	switchLoopbackControl := *openapiclient.NewSwitchLoopbackControl() // SwitchLoopbackControl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyStackLoopbackControl(context.Background(), omadacId, siteId, stackId).SwitchLoopbackControl(switchLoopbackControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyStackLoopbackControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyStackLoopbackControl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyStackLoopbackControl`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyStackLoopbackControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **switchLoopbackControl** | [**SwitchLoopbackControl**](SwitchLoopbackControl.md) |  | 

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


## ModifyStackNetwork

> OperationResponseOswStackNetworkModifyRespOpenApiVO ModifyStackNetwork(ctx, omadacId, siteId, stackId, networkId).OswNetworkOpenApi(oswNetworkOpenApi).Execute()

Modify stack vlan interface



### Example

```go
package main

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
	networkId := "networkId_example" // string | Network ID
	oswNetworkOpenApi := *openapiclient.NewOswNetworkOpenApi("Id_example", int32(123), false, int32(123)) // OswNetworkOpenApi | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyStackNetwork(context.Background(), omadacId, siteId, stackId, networkId).OswNetworkOpenApi(oswNetworkOpenApi).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyStackNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyStackNetwork`: OperationResponseOswStackNetworkModifyRespOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyStackNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyStackNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswNetworkOpenApi** | [**OswNetworkOpenApi**](OswNetworkOpenApi.md) |  | 

### Return type

[**OperationResponseOswStackNetworkModifyRespOpenApiVO**](OperationResponseOswStackNetworkModifyRespOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyStackRememberMe

> OperationResponseWithoutResult ModifyStackRememberMe(ctx, omadacId, siteId, stackId).DeviceRememberConfig(deviceRememberConfig).Execute()

Modify stack remember Config



### Example

```go
package main

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
	deviceRememberConfig := *openapiclient.NewDeviceRememberConfig() // DeviceRememberConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyStackRememberMe(context.Background(), omadacId, siteId, stackId).DeviceRememberConfig(deviceRememberConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyStackRememberMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyStackRememberMe`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyStackRememberMe`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyStackRememberMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deviceRememberConfig** | [**DeviceRememberConfig**](DeviceRememberConfig.md) |  | 

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


## ModifyStackStaticRouting

> OperationResponseWithoutResult ModifyStackStaticRouting(ctx, omadacId, siteId, stackId, staticRoutingId).OswStaticRoutingConfigOpenApiVO(oswStaticRoutingConfigOpenApiVO).Execute()

Modify stack staticRouting



### Example

```go
package main

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
	staticRoutingId := "staticRoutingId_example" // string | Static routing ID
	oswStaticRoutingConfigOpenApiVO := *openapiclient.NewOswStaticRoutingConfigOpenApiVO([]string{"Destinations_example"}, int32(123), int32(123), "NextHopIp_example", false) // OswStaticRoutingConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyStackStaticRouting(context.Background(), omadacId, siteId, stackId, staticRoutingId).OswStaticRoutingConfigOpenApiVO(oswStaticRoutingConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyStackStaticRouting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyStackStaticRouting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyStackStaticRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**staticRoutingId** | **string** | Static routing ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyStackStaticRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswStaticRoutingConfigOpenApiVO** | [**OswStaticRoutingConfigOpenApiVO**](OswStaticRoutingConfigOpenApiVO.md) |  | 

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


## ModifyStackVrf

> OperationResponseWithoutResult ModifyStackVrf(ctx, omadacId, siteId, stackId, vrfId).ModifyOswVrfOpenApiVO(modifyOswVrfOpenApiVO).Execute()

Modify an existing stack vrf



### Example

```go
package main

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
	vrfId := "vrfId_example" // string | VRF ID
	modifyOswVrfOpenApiVO := *openapiclient.NewModifyOswVrfOpenApiVO(false) // ModifyOswVrfOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifyStackVrf(context.Background(), omadacId, siteId, stackId, vrfId).ModifyOswVrfOpenApiVO(modifyOswVrfOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifyStackVrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyStackVrf`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifyStackVrf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**vrfId** | **string** | VRF ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyStackVrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **modifyOswVrfOpenApiVO** | [**ModifyOswVrfOpenApiVO**](ModifyOswVrfOpenApiVO.md) |  | 

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


## ModifySwitchStackLagSetting

> OperationResponseWithoutResult ModifySwitchStackLagSetting(ctx, omadacId, siteId, stackId, lagId).StackLagSettingVO(stackLagSettingVO).Execute()

Modify stack lag



### Example

```go
package main

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
	lagId := "lagId_example" // string | Lag ID
	stackLagSettingVO := *openapiclient.NewStackLagSettingVO() // StackLagSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifySwitchStackLagSetting(context.Background(), omadacId, siteId, stackId, lagId).StackLagSettingVO(stackLagSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifySwitchStackLagSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchStackLagSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifySwitchStackLagSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**lagId** | **string** | Lag ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchStackLagSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **stackLagSettingVO** | [**StackLagSettingVO**](StackLagSettingVO.md) |  | 

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


## ModifySwitchStackPortSetting

> OperationResponseWithoutResult ModifySwitchStackPortSetting(ctx, omadacId, siteId, stackId, port).StackPortSettingVO(stackPortSettingVO).Execute()

Modify stack port



### Example

```go
package main

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
	port := "port_example" // string | Port
	stackPortSettingVO := *openapiclient.NewStackPortSettingVO() // StackPortSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.ModifySwitchStackPortSetting(context.Background(), omadacId, siteId, stackId, port).StackPortSettingVO(stackPortSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.ModifySwitchStackPortSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchStackPortSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.ModifySwitchStackPortSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 
**port** | **string** | Port | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchStackPortSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **stackPortSettingVO** | [**StackPortSettingVO**](StackPortSettingVO.md) |  | 

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


## RebootOswStack

> OperationResponseListRebootResultVO RebootOswStack(ctx, omadacId, siteId, stackId).OswStackRebootOpenApiVO(oswStackRebootOpenApiVO).Execute()

Reboot switch stack



### Example

```go
package main

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
	oswStackRebootOpenApiVO := *openapiclient.NewOswStackRebootOpenApiVO() // OswStackRebootOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.RebootOswStack(context.Background(), omadacId, siteId, stackId).OswStackRebootOpenApiVO(oswStackRebootOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.RebootOswStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootOswStack`: OperationResponseListRebootResultVO
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.RebootOswStack`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRebootOswStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswStackRebootOpenApiVO** | [**OswStackRebootOpenApiVO**](OswStackRebootOpenApiVO.md) |  | 

### Return type

[**OperationResponseListRebootResultVO**](OperationResponseListRebootResultVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartStackCableTest

> OperationResponseWithoutResult StartStackCableTest(ctx, omadacId, siteId, stackId).OswCableTestTestingPortOpenApiVO(oswCableTestTestingPortOpenApiVO).Execute()

Start stack cable test



### Example

```go
package main

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
	oswCableTestTestingPortOpenApiVO := *openapiclient.NewOswCableTestTestingPortOpenApiVO() // OswCableTestTestingPortOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackAPI.StartStackCableTest(context.Background(), omadacId, siteId, stackId).OswCableTestTestingPortOpenApiVO(oswCableTestTestingPortOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackAPI.StartStackCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartStackCableTest`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `StackAPI.StartStackCableTest`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiStartStackCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswCableTestTestingPortOpenApiVO** | [**OswCableTestTestingPortOpenApiVO**](OswCableTestTestingPortOpenApiVO.md) |  | 

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


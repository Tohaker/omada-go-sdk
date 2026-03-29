# \SwitchTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPortTagTemplate**](SwitchTemplateAPI.md#AddPortTagTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/port-tag | Create new template port label
[**BatchModifySwitchPortTemplate**](SwitchTemplateAPI.md#BatchModifySwitchPortTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/multi-ports/config | Batch modify switch template port
[**BatchSetNameForGivenPorts**](SwitchTemplateAPI.md#BatchSetNameForGivenPorts) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/multi-ports/name | Batch set name for given switch template ports
[**BatchSetPoeModeForGivenPorts**](SwitchTemplateAPI.md#BatchSetPoeModeForGivenPorts) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/multi-ports/poe-mode | Batch set poe mode for given switch template ports
[**BatchSetPortStatusForGivenPorts**](SwitchTemplateAPI.md#BatchSetPortStatusForGivenPorts) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/multi-ports/status | Batch set status for given switch template ports
[**BatchSetProfileOverrideForGivenPorts**](SwitchTemplateAPI.md#BatchSetProfileOverrideForGivenPorts) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/multi-ports/profile-override | Batch set profile-override for given switch template ports
[**CreateOswVrfTemplate**](SwitchTemplateAPI.md#CreateOswVrfTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/vrfs | Create new vrf template
[**DeleteOswVrfTemplate**](SwitchTemplateAPI.md#DeleteOswVrfTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/vrfs/{vrfId} | Delete vrf template
[**DeletePortTagTemplate**](SwitchTemplateAPI.md#DeletePortTagTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/port-tag | Delete an existing template port label
[**DeleteSwitchLagTemplate**](SwitchTemplateAPI.md#DeleteSwitchLagTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/lags/{lagId} | Delete switch template lag
[**GetBatchSwitchTemplateExistNetworks**](SwitchTemplateAPI.md#GetBatchSwitchTemplateExistNetworks) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/exist-network | Get the networks intersections existing on multiple switch templates
[**GetGridPortAndLagNetworks**](SwitchTemplateAPI.md#GetGridPortAndLagNetworks) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/port-lag-networks | Get the networks used on switch template&#39;s ports and LAGs
[**GetGridVrfTemplate**](SwitchTemplateAPI.md#GetGridVrfTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/vrfs | Get the vrfs on the switch template
[**GetOswForGivenLanNetworkIdAndVlanTemplate**](SwitchTemplateAPI.md#GetOswForGivenLanNetworkIdAndVlanTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/networks/{networkId}/vlans/{vlan}/switches | Get the paging query for the osw templates with given network id and vlan
[**GetOswsDetailsTemplate**](SwitchTemplateAPI.md#GetOswsDetailsTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switches/details | Get the details of osw templates with given omadacid, siteid and macs and stackIds.
[**GetPortAndLagNetwork**](SwitchTemplateAPI.md#GetPortAndLagNetwork) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/port-lag-networks/{networkId}/vlan/{vlan} | Get the switch template&#39;s ports and LAGs that the network affects
[**GetPortTagTemplates**](SwitchTemplateAPI.md#GetPortTagTemplates) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/port-tag | Get template port label list
[**GetSwitchTemplateExistNetworks**](SwitchTemplateAPI.md#GetSwitchTemplateExistNetworks) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/exist-network | Get the networks existing on the switch template
[**GetSwitchTemplateInfo**](SwitchTemplateAPI.md#GetSwitchTemplateInfo) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId} | Get switch template info
[**ModifyOswVrfTemplate**](SwitchTemplateAPI.md#ModifyOswVrfTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/vrfs/{vrfId} | Modify vrf template
[**ModifyPortTagTemplate**](SwitchTemplateAPI.md#ModifyPortTagTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/port-tag | Modify an existing template port label
[**ModifySwitchLagTemplate**](SwitchTemplateAPI.md#ModifySwitchLagTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/lags/{lagId} | Modify switch template lag
[**ModifySwitchPortTemplate**](SwitchTemplateAPI.md#ModifySwitchPortTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/ports/{port} | Modify switch template port
[**SetNameForGivenPort**](SwitchTemplateAPI.md#SetNameForGivenPort) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/ports/{port}/name | Set name for given switch template port
[**SetPoeModeForGivenPort**](SwitchTemplateAPI.md#SetPoeModeForGivenPort) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/ports/{port}/poe-mode | Set poe mode for given switch template port
[**SetPortModeForGivenPort**](SwitchTemplateAPI.md#SetPortModeForGivenPort) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/ports/{port}/status | Set port status for given switch template port
[**SetProfileForGivenPort**](SwitchTemplateAPI.md#SetProfileForGivenPort) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/ports/{port}/profile | Set profile for given switch template port
[**SetProfileOverrideForGivenPort**](SwitchTemplateAPI.md#SetProfileOverrideForGivenPort) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switchtemplates/{deviceTemplateId}/ports/{port}/profile-override | Set profile-override for given switch template port



## AddPortTagTemplate

> PortTagOpenApiVO AddPortTagTemplate(ctx, omadacId, siteTemplateId).CreatePortTagOpenApiVO(createPortTagOpenApiVO).Execute()

Create new template port label



### Example

```go
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
	createPortTagOpenApiVO := *openapiclient.NewCreatePortTagOpenApiVO("Name_example") // CreatePortTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.AddPortTagTemplate(context.Background(), omadacId, siteTemplateId).CreatePortTagOpenApiVO(createPortTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.AddPortTagTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPortTagTemplate`: PortTagOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.AddPortTagTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPortTagTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createPortTagOpenApiVO** | [**CreatePortTagOpenApiVO**](CreatePortTagOpenApiVO.md) |  | 

### Return type

[**PortTagOpenApiVO**](PortTagOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchModifySwitchPortTemplate

> OperationResponseWithoutResult BatchModifySwitchPortTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).BatchOswPortSettingVO(batchOswPortSettingVO).Execute()

Batch modify switch template port



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	batchOswPortSettingVO := *openapiclient.NewBatchOswPortSettingVO([]int32{int32(123)}) // BatchOswPortSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.BatchModifySwitchPortTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).BatchOswPortSettingVO(batchOswPortSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.BatchModifySwitchPortTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifySwitchPortTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.BatchModifySwitchPortTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifySwitchPortTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchOswPortSettingVO** | [**BatchOswPortSettingVO**](BatchOswPortSettingVO.md) |  | 

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


## BatchSetNameForGivenPorts

> OperationResponseWithoutResult BatchSetNameForGivenPorts(ctx, omadacId, siteTemplateId, deviceTemplateId).PortNameList(portNameList).Execute()

Batch set name for given switch template ports



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	portNameList := *openapiclient.NewPortNameList([]openapiclient.SwitchMultiPortName{*openapiclient.NewSwitchMultiPortName("Name_example", int32(123))}) // PortNameList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.BatchSetNameForGivenPorts(context.Background(), omadacId, siteTemplateId, deviceTemplateId).PortNameList(portNameList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.BatchSetNameForGivenPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetNameForGivenPorts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.BatchSetNameForGivenPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchSetNameForGivenPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **portNameList** | [**PortNameList**](PortNameList.md) |  | 

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


## BatchSetPoeModeForGivenPorts

> OperationResponseWithoutResult BatchSetPoeModeForGivenPorts(ctx, omadacId, siteTemplateId, deviceTemplateId).SwitchPortsPoe(switchPortsPoe).Execute()

Batch set poe mode for given switch template ports



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	switchPortsPoe := *openapiclient.NewSwitchPortsPoe(int32(123), []int32{int32(123)}) // SwitchPortsPoe | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.BatchSetPoeModeForGivenPorts(context.Background(), omadacId, siteTemplateId, deviceTemplateId).SwitchPortsPoe(switchPortsPoe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.BatchSetPoeModeForGivenPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetPoeModeForGivenPorts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.BatchSetPoeModeForGivenPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchSetPoeModeForGivenPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **switchPortsPoe** | [**SwitchPortsPoe**](SwitchPortsPoe.md) |  | 

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


## BatchSetPortStatusForGivenPorts

> OperationResponseWithoutResult BatchSetPortStatusForGivenPorts(ctx, omadacId, siteTemplateId, deviceTemplateId).SwitchPortsStatus(switchPortsStatus).Execute()

Batch set status for given switch template ports



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	switchPortsStatus := *openapiclient.NewSwitchPortsStatus([]int32{int32(123)}, int32(123)) // SwitchPortsStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.BatchSetPortStatusForGivenPorts(context.Background(), omadacId, siteTemplateId, deviceTemplateId).SwitchPortsStatus(switchPortsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.BatchSetPortStatusForGivenPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetPortStatusForGivenPorts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.BatchSetPortStatusForGivenPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchSetPortStatusForGivenPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **switchPortsStatus** | [**SwitchPortsStatus**](SwitchPortsStatus.md) |  | 

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


## BatchSetProfileOverrideForGivenPorts

> OperationResponseWithoutResult BatchSetProfileOverrideForGivenPorts(ctx, omadacId, siteTemplateId, deviceTemplateId).BatchProfileOverride(batchProfileOverride).Execute()

Batch set profile-override for given switch template ports



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	batchProfileOverride := *openapiclient.NewBatchProfileOverride([]int32{int32(123)}, false) // BatchProfileOverride | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.BatchSetProfileOverrideForGivenPorts(context.Background(), omadacId, siteTemplateId, deviceTemplateId).BatchProfileOverride(batchProfileOverride).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.BatchSetProfileOverrideForGivenPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetProfileOverrideForGivenPorts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.BatchSetProfileOverrideForGivenPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchSetProfileOverrideForGivenPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchProfileOverride** | [**BatchProfileOverride**](BatchProfileOverride.md) |  | 

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


## CreateOswVrfTemplate

> OperationResponseWithoutResult CreateOswVrfTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).OswVrfConfigOpenApiVO(oswVrfConfigOpenApiVO).Execute()

Create new vrf template



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	oswVrfConfigOpenApiVO := *openapiclient.NewOswVrfConfigOpenApiVO(false, false, "Vrf_example") // OswVrfConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.CreateOswVrfTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).OswVrfConfigOpenApiVO(oswVrfConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.CreateOswVrfTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOswVrfTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.CreateOswVrfTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOswVrfTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswVrfConfigOpenApiVO** | [**OswVrfConfigOpenApiVO**](OswVrfConfigOpenApiVO.md) |  | 

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


## DeleteOswVrfTemplate

> OperationResponseWithoutResult DeleteOswVrfTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId, vrfId).Execute()

Delete vrf template



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	vrfId := "vrfId_example" // string | VRF ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.DeleteOswVrfTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId, vrfId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.DeleteOswVrfTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOswVrfTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.DeleteOswVrfTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**vrfId** | **string** | VRF ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOswVrfTemplateRequest struct via the builder pattern


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


## DeletePortTagTemplate

> OperationResponseWithoutResult DeletePortTagTemplate(ctx, omadacId, siteTemplateId).DeletePortTagOpenApiVO(deletePortTagOpenApiVO).Execute()

Delete an existing template port label



### Example

```go
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
	deletePortTagOpenApiVO := *openapiclient.NewDeletePortTagOpenApiVO("TagId_example") // DeletePortTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.DeletePortTagTemplate(context.Background(), omadacId, siteTemplateId).DeletePortTagOpenApiVO(deletePortTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.DeletePortTagTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortTagTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.DeletePortTagTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortTagTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deletePortTagOpenApiVO** | [**DeletePortTagOpenApiVO**](DeletePortTagOpenApiVO.md) |  | 

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


## DeleteSwitchLagTemplate

> OperationResponseWithoutResult DeleteSwitchLagTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId, lagId).Execute()

Delete switch template lag



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	lagId := "lagId_example" // string | lagId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.DeleteSwitchLagTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId, lagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.DeleteSwitchLagTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSwitchLagTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.DeleteSwitchLagTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**lagId** | **string** | lagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSwitchLagTemplateRequest struct via the builder pattern


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


## GetBatchSwitchTemplateExistNetworks

> OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO GetBatchSwitchTemplateExistNetworks(ctx, omadacId, siteTemplateId).OswMacListVO(oswMacListVO).Execute()

Get the networks intersections existing on multiple switch templates



### Example

```go
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
	oswMacListVO := *openapiclient.NewOswMacListVO([]string{"SwitchMacList_example"}) // OswMacListVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.GetBatchSwitchTemplateExistNetworks(context.Background(), omadacId, siteTemplateId).OswMacListVO(oswMacListVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetBatchSwitchTemplateExistNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchSwitchTemplateExistNetworks`: OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetBatchSwitchTemplateExistNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchSwitchTemplateExistNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oswMacListVO** | [**OswMacListVO**](OswMacListVO.md) |  | 

### Return type

[**OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO**](OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridPortAndLagNetworks

> GridVOOswNetworkBriefInfoOpenApiVO GetGridPortAndLagNetworks(ctx, omadacId, siteTemplateId, deviceTemplateId).Page(page).PageSize(pageSize).Execute()

Get the networks used on switch template's ports and LAGs



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.GetGridPortAndLagNetworks(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetGridPortAndLagNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPortAndLagNetworks`: GridVOOswNetworkBriefInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetGridPortAndLagNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridPortAndLagNetworksRequest struct via the builder pattern


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


## GetGridVrfTemplate

> OperationResponseGridVOOswVrfOpenApiVO GetGridVrfTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).Execute()

Get the vrfs on the switch template



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.GetGridVrfTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetGridVrfTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVrfTemplate`: OperationResponseGridVOOswVrfOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetGridVrfTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVrfTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseGridVOOswVrfOpenApiVO**](OperationResponseGridVOOswVrfOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswForGivenLanNetworkIdAndVlanTemplate

> OperationResponse GetOswForGivenLanNetworkIdAndVlanTemplate(ctx, omadacId, siteTemplateId, networkId, vlan).Page(page).PageSize(pageSize).Execute()

Get the paging query for the osw templates with given network id and vlan



### Example

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
	networkId := "networkId_example" // string | networkId
	vlan := "vlan_example" // string | vlan

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.GetOswForGivenLanNetworkIdAndVlanTemplate(context.Background(), omadacId, siteTemplateId, networkId, vlan).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetOswForGivenLanNetworkIdAndVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswForGivenLanNetworkIdAndVlanTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetOswForGivenLanNetworkIdAndVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**networkId** | **string** | networkId | 
**vlan** | **string** | vlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswForGivenLanNetworkIdAndVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





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


## GetOswsDetailsTemplate

> OperationResponse GetOswsDetailsTemplate(ctx, omadacId, siteTemplateId).OswDetailOpenApiVO(oswDetailOpenApiVO).Execute()

Get the details of osw templates with given omadacid, siteid and macs and stackIds.



### Example

```go
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
	oswDetailOpenApiVO := *openapiclient.NewOswDetailOpenApiVO("NetworkId_example", int32(123)) // OswDetailOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.GetOswsDetailsTemplate(context.Background(), omadacId, siteTemplateId).OswDetailOpenApiVO(oswDetailOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetOswsDetailsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswsDetailsTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetOswsDetailsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswsDetailsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oswDetailOpenApiVO** | [**OswDetailOpenApiVO**](OswDetailOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortAndLagNetwork

> OswPortAndLagNetworkVO GetPortAndLagNetwork(ctx, omadacId, siteTemplateId, deviceTemplateId, networkId, vlan).Execute()

Get the switch template's ports and LAGs that the network affects



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	networkId := "networkId_example" // string | Switch network ID.
	vlan := "vlan_example" // string | VLAN.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.GetPortAndLagNetwork(context.Background(), omadacId, siteTemplateId, deviceTemplateId, networkId, vlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetPortAndLagNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortAndLagNetwork`: OswPortAndLagNetworkVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetPortAndLagNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**networkId** | **string** | Switch network ID. | 
**vlan** | **string** | VLAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortAndLagNetworkRequest struct via the builder pattern


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


## GetPortTagTemplates

> []PortTagOpenApiVO GetPortTagTemplates(ctx, omadacId, siteTemplateId).Execute()

Get template port label list



### Example

```go
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
	resp, r, err := apiClient.SwitchTemplateAPI.GetPortTagTemplates(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetPortTagTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortTagTemplates`: []PortTagOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetPortTagTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortTagTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PortTagOpenApiVO**](PortTagOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchTemplateExistNetworks

> OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO GetSwitchTemplateExistNetworks(ctx, omadacId, siteTemplateId, deviceTemplateId).Execute()

Get the networks existing on the switch template



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.GetSwitchTemplateExistNetworks(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetSwitchTemplateExistNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchTemplateExistNetworks`: OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetSwitchTemplateExistNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchTemplateExistNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO**](OperationResponseLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchTemplateInfo

> OperationResponseSwitchTemplateOverviewInfo GetSwitchTemplateInfo(ctx, omadacId, siteTemplateId, deviceTemplateId).Execute()

Get switch template info



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.GetSwitchTemplateInfo(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.GetSwitchTemplateInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchTemplateInfo`: OperationResponseSwitchTemplateOverviewInfo
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.GetSwitchTemplateInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchTemplateInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSwitchTemplateOverviewInfo**](OperationResponseSwitchTemplateOverviewInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOswVrfTemplate

> OperationResponseWithoutResult ModifyOswVrfTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId, vrfId).OswVrfConfigOpenApiVO(oswVrfConfigOpenApiVO).Execute()

Modify vrf template



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	vrfId := "vrfId_example" // string | VRF ID
	oswVrfConfigOpenApiVO := *openapiclient.NewOswVrfConfigOpenApiVO(false, false, "Vrf_example") // OswVrfConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.ModifyOswVrfTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId, vrfId).OswVrfConfigOpenApiVO(oswVrfConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.ModifyOswVrfTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswVrfTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.ModifyOswVrfTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**vrfId** | **string** | VRF ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswVrfTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswVrfConfigOpenApiVO** | [**OswVrfConfigOpenApiVO**](OswVrfConfigOpenApiVO.md) |  | 

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


## ModifyPortTagTemplate

> OperationResponseWithoutResult ModifyPortTagTemplate(ctx, omadacId, siteTemplateId).PortTagOpenApiVO(portTagOpenApiVO).Execute()

Modify an existing template port label



### Example

```go
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
	portTagOpenApiVO := *openapiclient.NewPortTagOpenApiVO("Name_example", "TagId_example") // PortTagOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.ModifyPortTagTemplate(context.Background(), omadacId, siteTemplateId).PortTagOpenApiVO(portTagOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.ModifyPortTagTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortTagTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.ModifyPortTagTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortTagTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portTagOpenApiVO** | [**PortTagOpenApiVO**](PortTagOpenApiVO.md) |  | 

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


## ModifySwitchLagTemplate

> OperationResponseWithoutResult ModifySwitchLagTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId, lagId).OswLagSettingVO(oswLagSettingVO).Execute()

Modify switch template lag



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	lagId := "lagId_example" // string | lagId
	oswLagSettingVO := *openapiclient.NewOswLagSettingVO() // OswLagSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.ModifySwitchLagTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId, lagId).OswLagSettingVO(oswLagSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.ModifySwitchLagTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchLagTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.ModifySwitchLagTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**lagId** | **string** | lagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchLagTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswLagSettingVO** | [**OswLagSettingVO**](OswLagSettingVO.md) |  | 

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


## ModifySwitchPortTemplate

> OperationResponseWithoutResult ModifySwitchPortTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId, port).OswPortSettingVO(oswPortSettingVO).Execute()

Modify switch template port



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	port := "port_example" // string | Port ID
	oswPortSettingVO := *openapiclient.NewOswPortSettingVO() // OswPortSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.ModifySwitchPortTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId, port).OswPortSettingVO(oswPortSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.ModifySwitchPortTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchPortTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.ModifySwitchPortTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**port** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchPortTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oswPortSettingVO** | [**OswPortSettingVO**](OswPortSettingVO.md) |  | 

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


## SetNameForGivenPort

> OperationResponseWithoutResult SetNameForGivenPort(ctx, omadacId, siteTemplateId, deviceTemplateId, port).SwitchPortName(switchPortName).Execute()

Set name for given switch template port



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	port := "port_example" // string | Port ID
	switchPortName := *openapiclient.NewSwitchPortName("Name_example") // SwitchPortName | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.SetNameForGivenPort(context.Background(), omadacId, siteTemplateId, deviceTemplateId, port).SwitchPortName(switchPortName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.SetNameForGivenPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNameForGivenPort`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.SetNameForGivenPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**port** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNameForGivenPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **switchPortName** | [**SwitchPortName**](SwitchPortName.md) |  | 

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


## SetPoeModeForGivenPort

> OperationResponseWithoutResult SetPoeModeForGivenPort(ctx, omadacId, siteTemplateId, deviceTemplateId, port).SwitchPortPoe(switchPortPoe).Execute()

Set poe mode for given switch template port



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	port := "port_example" // string | port
	switchPortPoe := *openapiclient.NewSwitchPortPoe(int32(123)) // SwitchPortPoe | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.SetPoeModeForGivenPort(context.Background(), omadacId, siteTemplateId, deviceTemplateId, port).SwitchPortPoe(switchPortPoe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.SetPoeModeForGivenPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPoeModeForGivenPort`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.SetPoeModeForGivenPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**port** | **string** | port | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPoeModeForGivenPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **switchPortPoe** | [**SwitchPortPoe**](SwitchPortPoe.md) |  | 

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


## SetPortModeForGivenPort

> OperationResponseWithoutResult SetPortModeForGivenPort(ctx, omadacId, siteTemplateId, deviceTemplateId, port).SwitchPortStatus(switchPortStatus).Execute()

Set port status for given switch template port



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	port := "port_example" // string | Port
	switchPortStatus := *openapiclient.NewSwitchPortStatus(int32(123)) // SwitchPortStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.SetPortModeForGivenPort(context.Background(), omadacId, siteTemplateId, deviceTemplateId, port).SwitchPortStatus(switchPortStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.SetPortModeForGivenPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPortModeForGivenPort`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.SetPortModeForGivenPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**port** | **string** | Port | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPortModeForGivenPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **switchPortStatus** | [**SwitchPortStatus**](SwitchPortStatus.md) |  | 

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


## SetProfileForGivenPort

> OperationResponseWithoutResult SetProfileForGivenPort(ctx, omadacId, siteTemplateId, deviceTemplateId, port).SwitchProfileID(switchProfileID).Execute()

Set profile for given switch template port



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	port := "port_example" // string | Port ID
	switchProfileID := *openapiclient.NewSwitchProfileID("ProfileId_example") // SwitchProfileID | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.SetProfileForGivenPort(context.Background(), omadacId, siteTemplateId, deviceTemplateId, port).SwitchProfileID(switchProfileID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.SetProfileForGivenPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProfileForGivenPort`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.SetProfileForGivenPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**port** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProfileForGivenPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **switchProfileID** | [**SwitchProfileID**](SwitchProfileID.md) |  | 

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


## SetProfileOverrideForGivenPort

> OperationResponseWithoutResult SetProfileOverrideForGivenPort(ctx, omadacId, siteTemplateId, deviceTemplateId, port).ProfileOverride(profileOverride).Execute()

Set profile-override for given switch template port



### Example

```go
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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	port := "port_example" // string | Port ID
	profileOverride := *openapiclient.NewProfileOverride(false) // ProfileOverride | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchTemplateAPI.SetProfileOverrideForGivenPort(context.Background(), omadacId, siteTemplateId, deviceTemplateId, port).ProfileOverride(profileOverride).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchTemplateAPI.SetProfileOverrideForGivenPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProfileOverrideForGivenPort`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SwitchTemplateAPI.SetProfileOverrideForGivenPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**port** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProfileOverrideForGivenPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **profileOverride** | [**ProfileOverride**](ProfileOverride.md) |  | 

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


# \TopologyAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableNetworks**](TopologyAPI.md#GetAvailableNetworks) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/available-network | Get available network
[**GetDeviceLinkTopology**](TopologyAPI.md#GetDeviceLinkTopology) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/devices/{deviceMac}/device-link-topology | Get Device Link Topology
[**GetDevicesOfSsid**](TopologyAPI.md#GetDevicesOfSsid) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/ssids/{ssidId}/devices | Search Devices of Ssid in Topology
[**GetDevicesOfVlan**](TopologyAPI.md#GetDevicesOfVlan) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/{networkId}/vlan/{vlan}/devices | Search Devices of Vlan in Topology
[**GetFilterDevicesOfSsid**](TopologyAPI.md#GetFilterDevicesOfSsid) | **Get** /openapi/v3/{omadacId}/sites/{siteId}/topology/ssids/{ssidId}/devices | Filter Devices of Ssid in Topology
[**GetFilterDevicesOfVlan**](TopologyAPI.md#GetFilterDevicesOfVlan) | **Get** /openapi/v3/{omadacId}/sites/{siteId}/topology/{networkId}/vlan/{vlan}/devices | Filter Devices of Vlan in Topology
[**GetGridDeviceClient**](TopologyAPI.md#GetGridDeviceClient) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/search-device-client | Search Device or Client in Topology
[**GetIsolatedAndPreConfigDevices**](TopologyAPI.md#GetIsolatedAndPreConfigDevices) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/isolated-and-pre-config | Get isolated and preconfigured devices
[**GetMspDeviceLinkTopology**](TopologyAPI.md#GetMspDeviceLinkTopology) | **Get** /openapi/v2/msp/{mspId}/customers/{customerId}/sites/{siteId}/devices/{deviceMac}/device-link-topology | Get Msp Device Link Topology
[**GetTopology**](TopologyAPI.md#GetTopology) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology | Get site topology
[**GetTopologyClients**](TopologyAPI.md#GetTopologyClients) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/devices/{mac}/clients | Get the clients
[**GetTopologyClientsByDevice**](TopologyAPI.md#GetTopologyClientsByDevice) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/devices/{mac}/all-clients | Get the clients by device
[**GetTopologyClientsByDevices**](TopologyAPI.md#GetTopologyClientsByDevices) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/topology/devices/all-clients | Get the clients by devices.
[**GetTopologyDiscoveryStatus**](TopologyAPI.md#GetTopologyDiscoveryStatus) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/discovery/status | Get topology discovery status
[**GetTopologyNodes**](TopologyAPI.md#GetTopologyNodes) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/topology/nodes | Get topology nodes
[**GetV3Topology**](TopologyAPI.md#GetV3Topology) | **Get** /openapi/v3/{omadacId}/sites/{siteId}/topology | Get site topology, version 3
[**RefreshTopology**](TopologyAPI.md#RefreshTopology) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/topology | Refresh site topology
[**SetTopologyNode**](TopologyAPI.md#SetTopologyNode) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/topology/nodes | Set site topology top node



## GetAvailableNetworks

> OperationResponseTopologyAvailableNetworkAndSSID GetAvailableNetworks(ctx, omadacId, siteId).Execute()

Get available network



### Example

```go
package main

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
	resp, r, err := apiClient.TopologyAPI.GetAvailableNetworks(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetAvailableNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableNetworks`: OperationResponseTopologyAvailableNetworkAndSSID
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetAvailableNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseTopologyAvailableNetworkAndSSID**](OperationResponseTopologyAvailableNetworkAndSSID.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLinkTopology

> OperationResponseListTopologyV3OpenApiNodeVO GetDeviceLinkTopology(ctx, omadacId, siteId, deviceMac).Execute()

Get Device Link Topology



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetDeviceLinkTopology(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetDeviceLinkTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLinkTopology`: OperationResponseListTopologyV3OpenApiNodeVO
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetDeviceLinkTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLinkTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListTopologyV3OpenApiNodeVO**](OperationResponseListTopologyV3OpenApiNodeVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesOfSsid

> OperationResponseListString GetDevicesOfSsid(ctx, omadacId, siteId, ssidId).Execute()

Search Devices of Ssid in Topology



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetDevicesOfSsid(context.Background(), omadacId, siteId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetDevicesOfSsid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesOfSsid`: OperationResponseListString
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetDevicesOfSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesOfSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListString**](OperationResponseListString.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesOfVlan

> OperationResponseListString GetDevicesOfVlan(ctx, omadacId, siteId, networkId, vlan).Execute()

Search Devices of Vlan in Topology



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | networkId
	vlan := "vlan_example" // string | vlan

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetDevicesOfVlan(context.Background(), omadacId, siteId, networkId, vlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetDevicesOfVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesOfVlan`: OperationResponseListString
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetDevicesOfVlan`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDevicesOfVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseListString**](OperationResponseListString.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilterDevicesOfSsid

> OperationResponseTopologyFilterDevicesVO GetFilterDevicesOfSsid(ctx, omadacId, siteId, ssidId).Execute()

Filter Devices of Ssid in Topology



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	ssidId := "ssidId_example" // string | SSID ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetFilterDevicesOfSsid(context.Background(), omadacId, siteId, ssidId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetFilterDevicesOfSsid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilterDevicesOfSsid`: OperationResponseTopologyFilterDevicesVO
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetFilterDevicesOfSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ssidId** | **string** | SSID ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterDevicesOfSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseTopologyFilterDevicesVO**](OperationResponseTopologyFilterDevicesVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilterDevicesOfVlan

> OperationResponseTopologyFilterDevicesVO GetFilterDevicesOfVlan(ctx, omadacId, siteId, networkId, vlan).Execute()

Filter Devices of Vlan in Topology



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	networkId := "networkId_example" // string | networkId
	vlan := "vlan_example" // string | vlan

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetFilterDevicesOfVlan(context.Background(), omadacId, siteId, networkId, vlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetFilterDevicesOfVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilterDevicesOfVlan`: OperationResponseTopologyFilterDevicesVO
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetFilterDevicesOfVlan`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFilterDevicesOfVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseTopologyFilterDevicesVO**](OperationResponseTopologyFilterDevicesVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDeviceClient

> OperationResponseGridVOTopologyDeviceClient GetGridDeviceClient(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Search Device or Client in Topology



### Example

```go
package main

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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetGridDeviceClient(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetGridDeviceClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDeviceClient`: OperationResponseGridVOTopologyDeviceClient
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetGridDeviceClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDeviceClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field  | 

### Return type

[**OperationResponseGridVOTopologyDeviceClient**](OperationResponseGridVOTopologyDeviceClient.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsolatedAndPreConfigDevices

> OperationResponseTopologyIsolatedAndPreConfigDevice GetIsolatedAndPreConfigDevices(ctx, omadacId, siteId).Execute()

Get isolated and preconfigured devices



### Example

```go
package main

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
	resp, r, err := apiClient.TopologyAPI.GetIsolatedAndPreConfigDevices(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetIsolatedAndPreConfigDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsolatedAndPreConfigDevices`: OperationResponseTopologyIsolatedAndPreConfigDevice
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetIsolatedAndPreConfigDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIsolatedAndPreConfigDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseTopologyIsolatedAndPreConfigDevice**](OperationResponseTopologyIsolatedAndPreConfigDevice.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspDeviceLinkTopology

> OperationResponseListTopologyV3OpenApiNodeVO GetMspDeviceLinkTopology(ctx, mspId, customerId, siteId, deviceMac).Execute()

Get Msp Device Link Topology



### Example

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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetMspDeviceLinkTopology(context.Background(), mspId, customerId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetMspDeviceLinkTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspDeviceLinkTopology`: OperationResponseListTopologyV3OpenApiNodeVO
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetMspDeviceLinkTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspDeviceLinkTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseListTopologyV3OpenApiNodeVO**](OperationResponseListTopologyV3OpenApiNodeVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopology

> OperationResponseTopologyOpenApiVO GetTopology(ctx, omadacId, siteId).Execute()

Get site topology



### Example

```go
package main

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
	resp, r, err := apiClient.TopologyAPI.GetTopology(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopology`: OperationResponseTopologyOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseTopologyOpenApiVO**](OperationResponseTopologyOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyClients

> OperationResponseListTopologyClient GetTopologyClients(ctx, omadacId, siteId, mac).Execute()

Get the clients



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	mac := "mac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetTopologyClients(context.Background(), omadacId, siteId, mac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetTopologyClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopologyClients`: OperationResponseListTopologyClient
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetTopologyClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListTopologyClient**](OperationResponseListTopologyClient.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyClientsByDevice

> OperationResponseListTopologyClientNode GetTopologyClientsByDevice(ctx, omadacId, siteId, mac).FiltersType(filtersType).Execute()

Get the clients by device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	mac := "mac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	filtersType := "filtersType_example" // string | Filter query parameters, support field filter client group type. 0: except ipc; 1: just ipc; 0,1: all client.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetTopologyClientsByDevice(context.Background(), omadacId, siteId, mac).FiltersType(filtersType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetTopologyClientsByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopologyClientsByDevice`: OperationResponseListTopologyClientNode
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetTopologyClientsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyClientsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filtersType** | **string** | Filter query parameters, support field filter client group type. 0: except ipc; 1: just ipc; 0,1: all client. | 

### Return type

[**OperationResponseListTopologyClientNode**](OperationResponseListTopologyClientNode.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyClientsByDevices

> OperationResponseListTopologyDeviceClients GetTopologyClientsByDevices(ctx, omadacId, siteId).TopologyClientsQuery(topologyClientsQuery).Execute()

Get the clients by devices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	topologyClientsQuery := *openapiclient.NewTopologyClientsQuery() // TopologyClientsQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetTopologyClientsByDevices(context.Background(), omadacId, siteId).TopologyClientsQuery(topologyClientsQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetTopologyClientsByDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopologyClientsByDevices`: OperationResponseListTopologyDeviceClients
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetTopologyClientsByDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyClientsByDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **topologyClientsQuery** | [**TopologyClientsQuery**](TopologyClientsQuery.md) |  | 

### Return type

[**OperationResponseListTopologyDeviceClients**](OperationResponseListTopologyDeviceClients.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyDiscoveryStatus

> OperationResponse GetTopologyDiscoveryStatus(ctx, omadacId, siteId).Execute()

Get topology discovery status



### Example

```go
package main

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
	resp, r, err := apiClient.TopologyAPI.GetTopologyDiscoveryStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetTopologyDiscoveryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopologyDiscoveryStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetTopologyDiscoveryStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyDiscoveryStatusRequest struct via the builder pattern


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


## GetTopologyNodes

> OperationResponseListTopologyRoot GetTopologyNodes(ctx, omadacId, siteId).Execute()

Get topology nodes



### Example

```go
package main

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
	resp, r, err := apiClient.TopologyAPI.GetTopologyNodes(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetTopologyNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopologyNodes`: OperationResponseListTopologyRoot
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetTopologyNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListTopologyRoot**](OperationResponseListTopologyRoot.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV3Topology

> OperationResponseTopologyV3OpenApiVO GetV3Topology(ctx, omadacId, siteId).FiltersDeviceStatus(filtersDeviceStatus).Execute()

Get site topology, version 3



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	filtersDeviceStatus := "filtersDeviceStatus_example" // string | Filter query parameters, support field device status. 0: connected; 1: disconnected; 0,1: both. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.GetV3Topology(context.Background(), omadacId, siteId).FiltersDeviceStatus(filtersDeviceStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetV3Topology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV3Topology`: OperationResponseTopologyV3OpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetV3Topology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV3TopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersDeviceStatus** | **string** | Filter query parameters, support field device status. 0: connected; 1: disconnected; 0,1: both. | 

### Return type

[**OperationResponseTopologyV3OpenApiVO**](OperationResponseTopologyV3OpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshTopology

> OperationResponse RefreshTopology(ctx, omadacId, siteId).Execute()

Refresh site topology



### Example

```go
package main

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
	resp, r, err := apiClient.TopologyAPI.RefreshTopology(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.RefreshTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshTopology`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.RefreshTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTopologyRequest struct via the builder pattern


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


## SetTopologyNode

> OperationResponse SetTopologyNode(ctx, omadacId, siteId).TopologyRootNode(topologyRootNode).Execute()

Set site topology top node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	topologyRootNode := *openapiclient.NewTopologyRootNode() // TopologyRootNode | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopologyAPI.SetTopologyNode(context.Background(), omadacId, siteId).TopologyRootNode(topologyRootNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.SetTopologyNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTopologyNode`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.SetTopologyNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTopologyNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **topologyRootNode** | [**TopologyRootNode**](TopologyRootNode.md) |  | 

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


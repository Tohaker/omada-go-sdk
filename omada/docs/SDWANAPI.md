# \SDWANAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoSelectWanPortRecommendResult**](SDWANAPI.md#AutoSelectWanPortRecommendResult) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/recommend/wans | Auto recommend WAN port.
[**CheckSdWanGroupIpPool**](SDWANAPI.md#CheckSdWanGroupIpPool) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/ipPool/check | Check SD-WAN IP pool conflict.
[**CheckSelectedLanNetworkConflict**](SDWANAPI.md#CheckSelectedLanNetworkConflict) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/lan/check | Check whether selected lanNetwork conflict.
[**CreateSdWanGroup**](SDWANAPI.md#CreateSdWanGroup) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group | Create SD-WAN Group.
[**DeleteSdWanGroup**](SDWANAPI.md#DeleteSdWanGroup) | **Delete** /openapi/v1/{omadacId}/sdwan/sdwan-group/{groupId} | Delete SD-WAN Group.
[**FirstCheckConnection**](SDWANAPI.md#FirstCheckConnection) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/{groupId}/firstCheck | First check SD-WAN group connection.
[**GetCurrentSdWanGroup**](SDWANAPI.md#GetCurrentSdWanGroup) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/{groupId} | Get SD-WAN Group.
[**GetGridSdWanGroup**](SDWANAPI.md#GetGridSdWanGroup) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group | Get SD-WAN Group Grid.
[**GetGridSdWanGroupBrief**](SDWANAPI.md#GetGridSdWanGroupBrief) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/brief | Get brief SD-WAN Group Grid.
[**GetGridSdWanGroupDevices**](SDWANAPI.md#GetGridSdWanGroupDevices) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/candidate/devices | Get SD-WAN candidate devices.
[**GetSdWanGroupDevices**](SDWANAPI.md#GetSdWanGroupDevices) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/saved/devices/{groupId} | Get current SD-WAN devices.
[**ModifyLanIpRange**](SDWANAPI.md#ModifyLanIpRange) | **Put** /openapi/v1/{omadacId}/sdwan/sdwan-group/lan/modify | Modify selected LanNetwork IP.
[**ModifySdWanGroup**](SDWANAPI.md#ModifySdWanGroup) | **Put** /openapi/v1/{omadacId}/sdwan/sdwan-group/{groupId} | Modify SD-WAN Group.



## AutoSelectWanPortRecommendResult

> OperationResponseBatchAutoSelectWanPortResult AutoSelectWanPortRecommendResult(ctx, omadacId).BatchAutoSelectWanPortReq(batchAutoSelectWanPortReq).Execute()

Auto recommend WAN port.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	batchAutoSelectWanPortReq := *openapiclient.NewBatchAutoSelectWanPortReq() // BatchAutoSelectWanPortReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.AutoSelectWanPortRecommendResult(context.Background(), omadacId).BatchAutoSelectWanPortReq(batchAutoSelectWanPortReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.AutoSelectWanPortRecommendResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoSelectWanPortRecommendResult`: OperationResponseBatchAutoSelectWanPortResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.AutoSelectWanPortRecommendResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoSelectWanPortRecommendResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchAutoSelectWanPortReq** | [**BatchAutoSelectWanPortReq**](BatchAutoSelectWanPortReq.md) |  | 

### Return type

[**OperationResponseBatchAutoSelectWanPortResult**](OperationResponseBatchAutoSelectWanPortResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckSdWanGroupIpPool

> OperationResponseWithoutResult CheckSdWanGroupIpPool(ctx, omadacId).SdWanIpPoolRange(sdWanIpPoolRange).Execute()

Check SD-WAN IP pool conflict.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	sdWanIpPoolRange := *openapiclient.NewSdWanIpPoolRange() // SdWanIpPoolRange | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.CheckSdWanGroupIpPool(context.Background(), omadacId).SdWanIpPoolRange(sdWanIpPoolRange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.CheckSdWanGroupIpPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckSdWanGroupIpPool`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.CheckSdWanGroupIpPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckSdWanGroupIpPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdWanIpPoolRange** | [**SdWanIpPoolRange**](SdWanIpPoolRange.md) |  | 

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


## CheckSelectedLanNetworkConflict

> OperationResponseSdWanSelectedLanNetwork CheckSelectedLanNetworkConflict(ctx, omadacId).SdWanSelectedLanNetwork(sdWanSelectedLanNetwork).Execute()

Check whether selected lanNetwork conflict.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	sdWanSelectedLanNetwork := *openapiclient.NewSdWanSelectedLanNetwork() // SdWanSelectedLanNetwork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.CheckSelectedLanNetworkConflict(context.Background(), omadacId).SdWanSelectedLanNetwork(sdWanSelectedLanNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.CheckSelectedLanNetworkConflict``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckSelectedLanNetworkConflict`: OperationResponseSdWanSelectedLanNetwork
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.CheckSelectedLanNetworkConflict`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckSelectedLanNetworkConflictRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdWanSelectedLanNetwork** | [**SdWanSelectedLanNetwork**](SdWanSelectedLanNetwork.md) |  | 

### Return type

[**OperationResponseSdWanSelectedLanNetwork**](OperationResponseSdWanSelectedLanNetwork.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSdWanGroup

> OperationResponseWithoutResult CreateSdWanGroup(ctx, omadacId).SdWanGroup(sdWanGroup).Execute()

Create SD-WAN Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	sdWanGroup := *openapiclient.NewSdWanGroup() // SdWanGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.CreateSdWanGroup(context.Background(), omadacId).SdWanGroup(sdWanGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.CreateSdWanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSdWanGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.CreateSdWanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdWanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdWanGroup** | [**SdWanGroup**](SdWanGroup.md) |  | 

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


## DeleteSdWanGroup

> OperationResponseWithoutResult DeleteSdWanGroup(ctx, omadacId, groupId).Execute()

Delete SD-WAN Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	groupId := "groupId_example" // string | The ID of a SD-WAN Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.DeleteSdWanGroup(context.Background(), omadacId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.DeleteSdWanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSdWanGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.DeleteSdWanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**groupId** | **string** | The ID of a SD-WAN Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdWanGroupRequest struct via the builder pattern


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


## FirstCheckConnection

> OperationResponseSdWanGroupTunnelStatus FirstCheckConnection(ctx, omadacId, groupId).Execute()

First check SD-WAN group connection.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	groupId := "groupId_example" // string | The ID of a SD-WAN Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.FirstCheckConnection(context.Background(), omadacId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.FirstCheckConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FirstCheckConnection`: OperationResponseSdWanGroupTunnelStatus
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.FirstCheckConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**groupId** | **string** | The ID of a SD-WAN Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirstCheckConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSdWanGroupTunnelStatus**](OperationResponseSdWanGroupTunnelStatus.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentSdWanGroup

> OperationResponseSdWanGroup GetCurrentSdWanGroup(ctx, omadacId, groupId).Execute()

Get SD-WAN Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	groupId := "groupId_example" // string | The ID of a SD-WAN Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.GetCurrentSdWanGroup(context.Background(), omadacId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.GetCurrentSdWanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentSdWanGroup`: OperationResponseSdWanGroup
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.GetCurrentSdWanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**groupId** | **string** | The ID of a SD-WAN Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentSdWanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSdWanGroup**](OperationResponseSdWanGroup.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSdWanGroup

> OperationResponseGridVOSdWanGroupDetail GetGridSdWanGroup(ctx, omadacId).Page(page).PageSize(pageSize).Execute()

Get SD-WAN Group Grid.



### Example

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.GetGridSdWanGroup(context.Background(), omadacId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.GetGridSdWanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSdWanGroup`: OperationResponseGridVOSdWanGroupDetail
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.GetGridSdWanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSdWanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 


### Return type

[**OperationResponseGridVOSdWanGroupDetail**](OperationResponseGridVOSdWanGroupDetail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSdWanGroupBrief

> OperationResponseGridVOSdWanGroupBrief GetGridSdWanGroupBrief(ctx, omadacId).Page(page).PageSize(pageSize).Execute()

Get brief SD-WAN Group Grid.



### Example

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.GetGridSdWanGroupBrief(context.Background(), omadacId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.GetGridSdWanGroupBrief``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSdWanGroupBrief`: OperationResponseGridVOSdWanGroupBrief
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.GetGridSdWanGroupBrief`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSdWanGroupBriefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 


### Return type

[**OperationResponseGridVOSdWanGroupBrief**](OperationResponseGridVOSdWanGroupBrief.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSdWanGroupDevices

> OperationResponseGridVOSdWanCandidateDevice GetGridSdWanGroupDevices(ctx, omadacId).QuerySdWanCandidateDevice(querySdWanCandidateDevice).Execute()

Get SD-WAN candidate devices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	querySdWanCandidateDevice := *openapiclient.NewQuerySdWanCandidateDevice(int32(123), int32(123), int32(123)) // QuerySdWanCandidateDevice | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.GetGridSdWanGroupDevices(context.Background(), omadacId).QuerySdWanCandidateDevice(querySdWanCandidateDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.GetGridSdWanGroupDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSdWanGroupDevices`: OperationResponseGridVOSdWanCandidateDevice
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.GetGridSdWanGroupDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSdWanGroupDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **querySdWanCandidateDevice** | [**QuerySdWanCandidateDevice**](QuerySdWanCandidateDevice.md) |  | 

### Return type

[**OperationResponseGridVOSdWanCandidateDevice**](OperationResponseGridVOSdWanCandidateDevice.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdWanGroupDevices

> OperationResponseSdWanCurrentSelectedDeviceInfo GetSdWanGroupDevices(ctx, omadacId, groupId).Execute()

Get current SD-WAN devices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	groupId := "groupId_example" // string | The ID of a SD-WAN Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.GetSdWanGroupDevices(context.Background(), omadacId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.GetSdWanGroupDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSdWanGroupDevices`: OperationResponseSdWanCurrentSelectedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.GetSdWanGroupDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**groupId** | **string** | The ID of a SD-WAN Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdWanGroupDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSdWanCurrentSelectedDeviceInfo**](OperationResponseSdWanCurrentSelectedDeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyLanIpRange

> OperationResponseWithoutResult ModifyLanIpRange(ctx, omadacId).LanNetworkBrief(lanNetworkBrief).Execute()

Modify selected LanNetwork IP.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	lanNetworkBrief := *openapiclient.NewLanNetworkBrief() // LanNetworkBrief | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.ModifyLanIpRange(context.Background(), omadacId).LanNetworkBrief(lanNetworkBrief).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.ModifyLanIpRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLanIpRange`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.ModifyLanIpRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLanIpRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lanNetworkBrief** | [**LanNetworkBrief**](LanNetworkBrief.md) |  | 

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


## ModifySdWanGroup

> OperationResponseWithoutResult ModifySdWanGroup(ctx, omadacId, groupId).SdWanGroup(sdWanGroup).Execute()

Modify SD-WAN Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	groupId := "groupId_example" // string | The ID of a SD-WAN Group
	sdWanGroup := *openapiclient.NewSdWanGroup() // SdWanGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.ModifySdWanGroup(context.Background(), omadacId, groupId).SdWanGroup(sdWanGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.ModifySdWanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySdWanGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.ModifySdWanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**groupId** | **string** | The ID of a SD-WAN Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySdWanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sdWanGroup** | [**SdWanGroup**](SdWanGroup.md) |  | 

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


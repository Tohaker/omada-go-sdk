# SDWANAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoCreateMappingNetwork**](SDWANAPI.md#autocreatemappingnetwork) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/map/network | Auto create mapped network.
[**AutoCreateSdWanGroupName**](SDWANAPI.md#autocreatesdwangroupname) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/auto/groupName | Auto create group name.
[**AutoSelectWanPortRecommendResult**](SDWANAPI.md#autoselectwanportrecommendresult) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/recommend/wans | Auto recommend WAN port.
[**CheckModifiedMappedNetwork**](SDWANAPI.md#checkmodifiedmappednetwork) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/map/check | Check modified mapped network.
[**CheckSdWanGroupIpPool**](SDWANAPI.md#checksdwangroupippool) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/ipPool/check | Check SD-WAN IP pool conflict.
[**CreateSdWanGroup**](SDWANAPI.md#createsdwangroup) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group | Create SD-WAN Group.
[**DeleteSdWanGroup**](SDWANAPI.md#deletesdwangroup) | **Delete** /openapi/v1/{omadacId}/sdwan/sdwan-group/{groupId} | Delete SD-WAN Group.
[**FirstCheckConnection**](SDWANAPI.md#firstcheckconnection) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/{groupId}/firstCheck | First check SD-WAN group connection.
[**GetCurrentSdWanGroup**](SDWANAPI.md#getcurrentsdwangroup) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/{groupId} | Get SD-WAN Group.
[**GetGridSdWanGroup**](SDWANAPI.md#getgridsdwangroup) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group | Get SD-WAN Group Grid.
[**GetGridSdWanGroupBrief**](SDWANAPI.md#getgridsdwangroupbrief) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/brief | Get SD-WAN Group Grid brief info.
[**GetGridSdWanGroupDevices**](SDWANAPI.md#getgridsdwangroupdevices) | **Post** /openapi/v1/{omadacId}/sdwan/sdwan-group/candidate/devices | Get SD-WAN candidate devices.
[**GetSdWanGroupDevices**](SDWANAPI.md#getsdwangroupdevices) | **Get** /openapi/v1/{omadacId}/sdwan/sdwan-group/saved/devices/{groupId} | Get current SD-WAN devices.
[**ModifyLanIpRange**](SDWANAPI.md#modifylaniprange) | **Put** /openapi/v1/{omadacId}/sdwan/sdwan-group/lan/modify | Modify selected LanNetwork IP.
[**ModifySdWanGroup**](SDWANAPI.md#modifysdwangroup) | **Put** /openapi/v1/{omadacId}/sdwan/sdwan-group/{groupId} | Modify SD-WAN Group.
[**ModifySdWanGroupNetWorkMap**](SDWANAPI.md#modifysdwangroupnetworkmap) | **Put** /openapi/v1/{omadacId}/sdwan/sdwan-group/map/{groupId} | Modify SD-WAN Group NAT info.



## AutoCreateMappingNetwork

> OperationResponseSdWanMappedNetworkResult AutoCreateMappingNetwork(ctx, omadacId).SdWanSelectedMapNetwork(sdWanSelectedMapNetwork).Execute()

Auto create mapped network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	sdWanSelectedMapNetwork := *openapiclient.NewSdWanSelectedMapNetwork() // SdWanSelectedMapNetwork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.AutoCreateMappingNetwork(context.Background(), omadacId).SdWanSelectedMapNetwork(sdWanSelectedMapNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.AutoCreateMappingNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreateMappingNetwork`: OperationResponseSdWanMappedNetworkResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.AutoCreateMappingNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreateMappingNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdWanSelectedMapNetwork** | [**SdWanSelectedMapNetwork**](SdWanSelectedMapNetwork.md) |  | 

### Return type

[**OperationResponseSdWanMappedNetworkResult**](OperationResponseSdWanMappedNetworkResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoCreateSdWanGroupName

> OperationResponseString AutoCreateSdWanGroupName(ctx, omadacId).Execute()

Auto create group name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.AutoCreateSdWanGroupName(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.AutoCreateSdWanGroupName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreateSdWanGroupName`: OperationResponseString
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.AutoCreateSdWanGroupName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreateSdWanGroupNameRequest struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckModifiedMappedNetwork

> OperationResponseWithoutResult CheckModifiedMappedNetwork(ctx, omadacId).CheckMappedNetwork(checkMappedNetwork).Execute()

Check modified mapped network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	checkMappedNetwork := *openapiclient.NewCheckMappedNetwork() // CheckMappedNetwork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.CheckModifiedMappedNetwork(context.Background(), omadacId).CheckMappedNetwork(checkMappedNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.CheckModifiedMappedNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckModifiedMappedNetwork`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.CheckModifiedMappedNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckModifiedMappedNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkMappedNetwork** | [**CheckMappedNetwork**](CheckMappedNetwork.md) |  | 

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSdWanGroupBrief

> OperationResponseGridVOSdWanGroupBrief GetGridSdWanGroupBrief(ctx, omadacId).Page(page).PageSize(pageSize).Execute()

Get SD-WAN Group Grid brief info.



### Example

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySdWanGroupNetWorkMap

> OperationResponseWithoutResult ModifySdWanGroupNetWorkMap(ctx, omadacId, groupId).SdWanNatInfo(sdWanNatInfo).Execute()

Modify SD-WAN Group NAT info.



### Example

```go
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
	sdWanNatInfo := *openapiclient.NewSdWanNatInfo() // SdWanNatInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDWANAPI.ModifySdWanGroupNetWorkMap(context.Background(), omadacId, groupId).SdWanNatInfo(sdWanNatInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDWANAPI.ModifySdWanGroupNetWorkMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySdWanGroupNetWorkMap`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SDWANAPI.ModifySdWanGroupNetWorkMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**groupId** | **string** | The ID of a SD-WAN Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySdWanGroupNetWorkMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sdWanNatInfo** | [**SdWanNatInfo**](SdWanNatInfo.md) |  | 

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


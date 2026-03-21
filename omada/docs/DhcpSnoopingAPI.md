# \DhcpSnoopingAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDhcpSnoops**](DhcpSnoopingAPI.md#CreateDhcpSnoops) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/dhcpSnoops | Create new dhcp snoops
[**DeleteDhcpSnoop**](DhcpSnoopingAPI.md#DeleteDhcpSnoop) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/dhcpSnoops/{dhcpSnoopId} | Delete an existing dhcp snoop
[**GetDhcpSnoopImpbs**](DhcpSnoopingAPI.md#GetDhcpSnoopImpbs) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/dhcpSnoops/impbs | Get the impbs with given omadacid, siteid and mac and stackId.
[**GetDhcpSnoopStatus**](DhcpSnoopingAPI.md#GetDhcpSnoopStatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dhcpSnoops/status | Get dhcp snoop status
[**GetGridDhcpSnoops**](DhcpSnoopingAPI.md#GetGridDhcpSnoops) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dhcpSnoops | Get dhcp snoop list
[**ModifyDhcpSnoop**](DhcpSnoopingAPI.md#ModifyDhcpSnoop) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/dhcpSnoops/{dhcpSnoopId} | Modify a dhcp snoop
[**ModifyDhcpSnoopStatus**](DhcpSnoopingAPI.md#ModifyDhcpSnoopStatus) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/dhcpSnoops/status | Modify dhcp snoop status



## CreateDhcpSnoops

> ResponseIdVO CreateDhcpSnoops(ctx, omadacId, siteId).DhcpSnoopVO(dhcpSnoopVO).Execute()

Create new dhcp snoops



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	dhcpSnoopVO := *openapiclient.NewDhcpSnoopVO() // DhcpSnoopVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpSnoopingAPI.CreateDhcpSnoops(context.Background(), omadacId, siteId).DhcpSnoopVO(dhcpSnoopVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpSnoopingAPI.CreateDhcpSnoops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDhcpSnoops`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `DhcpSnoopingAPI.CreateDhcpSnoops`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDhcpSnoopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dhcpSnoopVO** | [**DhcpSnoopVO**](DhcpSnoopVO.md) |  | 

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


## DeleteDhcpSnoop

> OperationResponseWithoutResult DeleteDhcpSnoop(ctx, omadacId, siteId, dhcpSnoopId).Execute()

Delete an existing dhcp snoop



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	dhcpSnoopId := "dhcpSnoopId_example" // string | dhcpSnoopId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpSnoopingAPI.DeleteDhcpSnoop(context.Background(), omadacId, siteId, dhcpSnoopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpSnoopingAPI.DeleteDhcpSnoop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDhcpSnoop`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DhcpSnoopingAPI.DeleteDhcpSnoop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**dhcpSnoopId** | **string** | dhcpSnoopId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDhcpSnoopRequest struct via the builder pattern


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


## GetDhcpSnoopImpbs

> OperationResponse GetDhcpSnoopImpbs(ctx, omadacId, siteId).DhcpSnoopImbpVO(dhcpSnoopImbpVO).Execute()

Get the impbs with given omadacid, siteid and mac and stackId.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	dhcpSnoopImbpVO := *openapiclient.NewDhcpSnoopImbpVO() // DhcpSnoopImbpVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpSnoopingAPI.GetDhcpSnoopImpbs(context.Background(), omadacId, siteId).DhcpSnoopImbpVO(dhcpSnoopImbpVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpSnoopingAPI.GetDhcpSnoopImpbs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpSnoopImpbs`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpSnoopingAPI.GetDhcpSnoopImpbs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpSnoopImpbsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dhcpSnoopImbpVO** | [**DhcpSnoopImbpVO**](DhcpSnoopImbpVO.md) |  | 

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


## GetDhcpSnoopStatus

> OperationResponse GetDhcpSnoopStatus(ctx, omadacId, siteId).Execute()

Get dhcp snoop status



### Example

```go
package main

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
	resp, r, err := apiClient.DhcpSnoopingAPI.GetDhcpSnoopStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpSnoopingAPI.GetDhcpSnoopStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpSnoopStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpSnoopingAPI.GetDhcpSnoopStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpSnoopStatusRequest struct via the builder pattern


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


## GetGridDhcpSnoops

> OperationResponse GetGridDhcpSnoops(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get dhcp snoop list



### Example

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
	resp, r, err := apiClient.DhcpSnoopingAPI.GetGridDhcpSnoops(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpSnoopingAPI.GetGridDhcpSnoops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDhcpSnoops`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpSnoopingAPI.GetGridDhcpSnoops`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDhcpSnoopsRequest struct via the builder pattern


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


## ModifyDhcpSnoop

> OperationResponseWithoutResult ModifyDhcpSnoop(ctx, omadacId, siteId, dhcpSnoopId).DhcpSnoopVO(dhcpSnoopVO).Execute()

Modify a dhcp snoop



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	dhcpSnoopId := "dhcpSnoopId_example" // string | dhcpSnoopId
	dhcpSnoopVO := *openapiclient.NewDhcpSnoopVO() // DhcpSnoopVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpSnoopingAPI.ModifyDhcpSnoop(context.Background(), omadacId, siteId, dhcpSnoopId).DhcpSnoopVO(dhcpSnoopVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpSnoopingAPI.ModifyDhcpSnoop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDhcpSnoop`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DhcpSnoopingAPI.ModifyDhcpSnoop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**dhcpSnoopId** | **string** | dhcpSnoopId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDhcpSnoopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dhcpSnoopVO** | [**DhcpSnoopVO**](DhcpSnoopVO.md) |  | 

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


## ModifyDhcpSnoopStatus

> OperationResponseWithoutResult ModifyDhcpSnoopStatus(ctx, omadacId, siteId).DhcpSnoopStatusVO(dhcpSnoopStatusVO).Execute()

Modify dhcp snoop status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	dhcpSnoopStatusVO := *openapiclient.NewDhcpSnoopStatusVO() // DhcpSnoopStatusVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpSnoopingAPI.ModifyDhcpSnoopStatus(context.Background(), omadacId, siteId).DhcpSnoopStatusVO(dhcpSnoopStatusVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpSnoopingAPI.ModifyDhcpSnoopStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDhcpSnoopStatus`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DhcpSnoopingAPI.ModifyDhcpSnoopStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDhcpSnoopStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dhcpSnoopStatusVO** | [**DhcpSnoopStatusVO**](DhcpSnoopStatusVO.md) |  | 

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


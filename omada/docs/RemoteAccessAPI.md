# RemoteAccessAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTunnel**](RemoteAccessAPI.md#addtunnel) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/remoteAccess/tunnel | Create new remote access tunnel
[**DeleteTunnel**](RemoteAccessAPI.md#deletetunnel) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/remoteAccess/tunnel/{tunnelId} | Delete remote access tunnel information
[**EditTunnel**](RemoteAccessAPI.md#edittunnel) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/remoteAccess/tunnel/{tunnelId} | Modify remote access tunnel information
[**GetTunnel**](RemoteAccessAPI.md#gettunnel) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/remoteAccess/tunnel | Get remote access tunnel
[**GetTunnelStatus**](RemoteAccessAPI.md#gettunnelstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/remoteAccess/tunnel/{tunnelId}/status | Get remote access tunnel&#39;s status
[**GetTunnelsStatus**](RemoteAccessAPI.md#gettunnelsstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/remoteAccess/tunnel/status | Get all remote access tunnel&#39;s status
[**ModifyTunnel**](RemoteAccessAPI.md#modifytunnel) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/remoteAccess/tunnel/{tunnelId}/status | Enable/Disable remote access tunnel



## AddTunnel

> OperationResponseWithoutResult AddTunnel(ctx, omadacId, siteId).NatTraversalTunnelOpenApiVO(natTraversalTunnelOpenApiVO).Execute()

Create new remote access tunnel



### Example

```go
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
	natTraversalTunnelOpenApiVO := *openapiclient.NewNatTraversalTunnelOpenApiVO("AppType_example", "LocalAddress_example", int32(123), "Name_example") // NatTraversalTunnelOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.AddTunnel(context.Background(), omadacId, siteId).NatTraversalTunnelOpenApiVO(natTraversalTunnelOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.AddTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.AddTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **natTraversalTunnelOpenApiVO** | [**NatTraversalTunnelOpenApiVO**](NatTraversalTunnelOpenApiVO.md) |  | 

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


## DeleteTunnel

> OperationResponseWithoutResult DeleteTunnel(ctx, omadacId, siteId, tunnelId).Execute()

Delete remote access tunnel information



### Example

```go
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
	tunnelId := "tunnelId_example" // string | Remote access tunnel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.DeleteTunnel(context.Background(), omadacId, siteId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.DeleteTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.DeleteTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tunnelId** | **string** | Remote access tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTunnelRequest struct via the builder pattern


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


## EditTunnel

> OperationResponseWithoutResult EditTunnel(ctx, omadacId, siteId, tunnelId).NatTraversalTunnelOpenApiVO(natTraversalTunnelOpenApiVO).Execute()

Modify remote access tunnel information



### Example

```go
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
	tunnelId := "tunnelId_example" // string | Remote access tunnel ID
	natTraversalTunnelOpenApiVO := *openapiclient.NewNatTraversalTunnelOpenApiVO("AppType_example", "LocalAddress_example", int32(123), "Name_example") // NatTraversalTunnelOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.EditTunnel(context.Background(), omadacId, siteId, tunnelId).NatTraversalTunnelOpenApiVO(natTraversalTunnelOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.EditTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.EditTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tunnelId** | **string** | Remote access tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **natTraversalTunnelOpenApiVO** | [**NatTraversalTunnelOpenApiVO**](NatTraversalTunnelOpenApiVO.md) |  | 

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


## GetTunnel

> OperationResponseNatTraversalTunnelGridVONatTraversalTunnelVO GetTunnel(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get remote access tunnel



### Example

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
	resp, r, err := apiClient.RemoteAccessAPI.GetTunnel(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.GetTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTunnel`: OperationResponseNatTraversalTunnelGridVONatTraversalTunnelVO
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.GetTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseNatTraversalTunnelGridVONatTraversalTunnelVO**](OperationResponseNatTraversalTunnelGridVONatTraversalTunnelVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTunnelStatus

> OperationResponseNatTraversalSingleTunnelStatusVO GetTunnelStatus(ctx, omadacId, siteId, tunnelId).Execute()

Get remote access tunnel's status



### Example

```go
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
	tunnelId := "tunnelId_example" // string | Remote access tunnel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.GetTunnelStatus(context.Background(), omadacId, siteId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.GetTunnelStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTunnelStatus`: OperationResponseNatTraversalSingleTunnelStatusVO
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.GetTunnelStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tunnelId** | **string** | Remote access tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTunnelStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseNatTraversalSingleTunnelStatusVO**](OperationResponseNatTraversalSingleTunnelStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTunnelsStatus

> OperationResponseNatTraversalTunnelsStatusVO GetTunnelsStatus(ctx, omadacId, siteId).Execute()

Get all remote access tunnel's status



### Example

```go
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
	resp, r, err := apiClient.RemoteAccessAPI.GetTunnelsStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.GetTunnelsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTunnelsStatus`: OperationResponseNatTraversalTunnelsStatusVO
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.GetTunnelsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTunnelsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseNatTraversalTunnelsStatusVO**](OperationResponseNatTraversalTunnelsStatusVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTunnel

> OperationResponseWithoutResult ModifyTunnel(ctx, omadacId, siteId, tunnelId).NatTraversalTunnelOpenVO(natTraversalTunnelOpenVO).Execute()

Enable/Disable remote access tunnel



### Example

```go
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
	tunnelId := "tunnelId_example" // string | Remote access tunnel ID
	natTraversalTunnelOpenVO := *openapiclient.NewNatTraversalTunnelOpenVO() // NatTraversalTunnelOpenVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.ModifyTunnel(context.Background(), omadacId, siteId, tunnelId).NatTraversalTunnelOpenVO(natTraversalTunnelOpenVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.ModifyTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.ModifyTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tunnelId** | **string** | Remote access tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **natTraversalTunnelOpenVO** | [**NatTraversalTunnelOpenVO**](NatTraversalTunnelOpenVO.md) |  | 

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


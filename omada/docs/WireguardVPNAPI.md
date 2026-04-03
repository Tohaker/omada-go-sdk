# WireguardVPNAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePeer**](WireguardVPNAPI.md#createpeer) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguard-peers | Create new wireguard peer
[**CreateWireguard**](WireguardVPNAPI.md#createwireguard) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguards | Create new wireguard
[**DeletePeer**](WireguardVPNAPI.md#deletepeer) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguard-peers/{id} | Delete an existing wireguard peer
[**DeleteWireguard**](WireguardVPNAPI.md#deletewireguard) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguards/{id} | Delete an existing wireguard
[**GetWireguardKey**](WireguardVPNAPI.md#getwireguardkey) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguard-key | Get new private key and public key
[**GetWireguardSummary**](WireguardVPNAPI.md#getwireguardsummary) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguard-summarys | Get all wireguard&#39;s id and name info
[**ListPeer**](WireguardVPNAPI.md#listpeer) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguard-peers | Get wireguard peer list
[**ListWireguard**](WireguardVPNAPI.md#listwireguard) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguards | Get wireguard list
[**ModifyPeer**](WireguardVPNAPI.md#modifypeer) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguard-peers/{id} | Modify an existing wireguard peer
[**ModifyWireguard**](WireguardVPNAPI.md#modifywireguard) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/wireguards/{id} | Modify an existing wireguard



## CreatePeer

> OperationResponseResponseIdVO CreatePeer(ctx, omadacId, siteId).WireguardPeerOpenApiVO(wireguardPeerOpenApiVO).Execute()

Create new wireguard peer



### Example

```go
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
	wireguardPeerOpenApiVO := *openapiclient.NewWireguardPeerOpenApiVO([]string{"AllowAddress_example"}, "InterfaceId_example", int32(123), "Name_example", "PublicKey_example", false) // WireguardPeerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardVPNAPI.CreatePeer(context.Background(), omadacId, siteId).WireguardPeerOpenApiVO(wireguardPeerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.CreatePeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePeer`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.CreatePeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wireguardPeerOpenApiVO** | [**WireguardPeerOpenApiVO**](WireguardPeerOpenApiVO.md) |  | 

### Return type

[**OperationResponseResponseIdVO**](OperationResponseResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWireguard

> OperationResponseResponseIdVO CreateWireguard(ctx, omadacId, siteId).WireguardOpenApiVO(wireguardOpenApiVO).Execute()

Create new wireguard



### Example

```go
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
	wireguardOpenApiVO := *openapiclient.NewWireguardOpenApiVO(int32(123), "LocalIp_example", int32(123), "Name_example", "PrivateKey_example", false) // WireguardOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardVPNAPI.CreateWireguard(context.Background(), omadacId, siteId).WireguardOpenApiVO(wireguardOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.CreateWireguard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWireguard`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.CreateWireguard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWireguardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wireguardOpenApiVO** | [**WireguardOpenApiVO**](WireguardOpenApiVO.md) |  | 

### Return type

[**OperationResponseResponseIdVO**](OperationResponseResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePeer

> OperationResponseWithoutResult DeletePeer(ctx, omadacId, siteId, id).Execute()

Delete an existing wireguard peer



### Example

```go
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
	id := "id_example" // string | Peer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardVPNAPI.DeletePeer(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.DeletePeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePeer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.DeletePeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Peer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePeerRequest struct via the builder pattern


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


## DeleteWireguard

> OperationResponseWithoutResult DeleteWireguard(ctx, omadacId, siteId, id).Execute()

Delete an existing wireguard



### Example

```go
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
	id := "id_example" // string | WireGuard VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardVPNAPI.DeleteWireguard(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.DeleteWireguard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWireguard`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.DeleteWireguard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | WireGuard VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWireguardRequest struct via the builder pattern


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


## GetWireguardKey

> OperationResponseWireguardKeyOpenApiVO GetWireguardKey(ctx, omadacId, siteId).Execute()

Get new private key and public key



### Example

```go
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
	resp, r, err := apiClient.WireguardVPNAPI.GetWireguardKey(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.GetWireguardKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWireguardKey`: OperationResponseWireguardKeyOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.GetWireguardKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWireguardKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWireguardKeyOpenApiVO**](OperationResponseWireguardKeyOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireguardSummary

> OperationResponseResponseDataVOWireguardSummaryOpenApiVO GetWireguardSummary(ctx, omadacId, siteId).Execute()

Get all wireguard's id and name info



### Example

```go
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
	resp, r, err := apiClient.WireguardVPNAPI.GetWireguardSummary(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.GetWireguardSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWireguardSummary`: OperationResponseResponseDataVOWireguardSummaryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.GetWireguardSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWireguardSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseResponseDataVOWireguardSummaryOpenApiVO**](OperationResponseResponseDataVOWireguardSummaryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPeer

> OperationResponseWireguardPeerOpenApiGridVOWireguardPeerDetailOpenApiVO ListPeer(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get wireguard peer list



### Example

```go
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
	searchKey := "searchKey_example" // string | searchKey (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardVPNAPI.ListPeer(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.ListPeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPeer`: OperationResponseWireguardPeerOpenApiGridVOWireguardPeerDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.ListPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | searchKey | 

### Return type

[**OperationResponseWireguardPeerOpenApiGridVOWireguardPeerDetailOpenApiVO**](OperationResponseWireguardPeerOpenApiGridVOWireguardPeerDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWireguard

> OperationResponseGridVOWireguardDetailOpenApiVO ListWireguard(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get wireguard list



### Example

```go
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
	searchKey := "searchKey_example" // string | searchKey (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardVPNAPI.ListWireguard(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.ListWireguard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWireguard`: OperationResponseGridVOWireguardDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.ListWireguard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWireguardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | searchKey | 

### Return type

[**OperationResponseGridVOWireguardDetailOpenApiVO**](OperationResponseGridVOWireguardDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPeer

> OperationResponseWithoutResult ModifyPeer(ctx, omadacId, siteId, id).WireguardPeerOpenApiVO(wireguardPeerOpenApiVO).Execute()

Modify an existing wireguard peer



### Example

```go
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
	id := "id_example" // string | Peer ID
	wireguardPeerOpenApiVO := *openapiclient.NewWireguardPeerOpenApiVO([]string{"AllowAddress_example"}, "InterfaceId_example", int32(123), "Name_example", "PublicKey_example", false) // WireguardPeerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardVPNAPI.ModifyPeer(context.Background(), omadacId, siteId, id).WireguardPeerOpenApiVO(wireguardPeerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.ModifyPeer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPeer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.ModifyPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Peer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **wireguardPeerOpenApiVO** | [**WireguardPeerOpenApiVO**](WireguardPeerOpenApiVO.md) |  | 

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


## ModifyWireguard

> OperationResponseWithoutResult ModifyWireguard(ctx, omadacId, siteId, id).WireguardOpenApiVO(wireguardOpenApiVO).Execute()

Modify an existing wireguard



### Example

```go
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
	id := "id_example" // string | Wireguard ID
	wireguardOpenApiVO := *openapiclient.NewWireguardOpenApiVO(int32(123), "LocalIp_example", int32(123), "Name_example", "PrivateKey_example", false) // WireguardOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardVPNAPI.ModifyWireguard(context.Background(), omadacId, siteId, id).WireguardOpenApiVO(wireguardOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardVPNAPI.ModifyWireguard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWireguard`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `WireguardVPNAPI.ModifyWireguard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Wireguard ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWireguardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **wireguardOpenApiVO** | [**WireguardOpenApiVO**](WireguardOpenApiVO.md) |  | 

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


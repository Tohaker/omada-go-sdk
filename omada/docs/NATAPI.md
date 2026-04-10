# NATAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOtoNat**](NATAPI.md#createotonat) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/nat/one-to-one-nat | Create new One-to-One NAT
[**CreatePortForwarding**](NATAPI.md#createportforwarding) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/nat/port-forwardings | Create new port forwarding
[**DeleteOtoNat**](NATAPI.md#deleteotonat) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/nat/one-to-one-nat/{otonatId} | Delete One-to-One NAT
[**DeletePortForwarding**](NATAPI.md#deleteportforwarding) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/nat/port-forwardings/{portForwardingId} | Delete port forwarding
[**GetAlg**](NATAPI.md#getalg) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/nat/alg | Get ALG Info
[**GetGridOtoNats**](NATAPI.md#getgridotonats) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/nat/one-to-one-nat | Get One-to-One NAT list
[**GetPortForwardingList**](NATAPI.md#getportforwardinglist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/nat/port-forwardings | Get port forwarding list
[**ModifyAlg**](NATAPI.md#modifyalg) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/nat/alg | Modify ALG setting
[**ModifyOtoNat**](NATAPI.md#modifyotonat) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/nat/one-to-one-nat/{otonatId} | Modify One-to-One NAT
[**ModifyPortForwarding**](NATAPI.md#modifyportforwarding) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/nat/port-forwardings/{portForwardingId} | Modify port forwarding



## CreateOtoNat

> OperationResponseWithoutResult CreateOtoNat(ctx, omadacId, siteId).OtoNatOpenApiVO(otoNatOpenApiVO).Execute()

Create new One-to-One NAT



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
	otoNatOpenApiVO := *openapiclient.NewOtoNatOpenApiVO(false, "ExternalIp_example", []string{"InterfaceIds_example"}, "InternalIp_example", "Name_example", false) // OtoNatOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATAPI.CreateOtoNat(context.Background(), omadacId, siteId).OtoNatOpenApiVO(otoNatOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.CreateOtoNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOtoNat`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.CreateOtoNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOtoNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **otoNatOpenApiVO** | [**OtoNatOpenApiVO**](OtoNatOpenApiVO.md) |  | 

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


## CreatePortForwarding

> OperationResponseWithoutResult CreatePortForwarding(ctx, omadacId, siteId).PortForwardingConfig(portForwardingConfig).Execute()

Create new port forwarding



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
	portForwardingConfig := *openapiclient.NewPortForwardingConfig(false, "ForwardIp_example", int32(123), "Name_example", false) // PortForwardingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATAPI.CreatePortForwarding(context.Background(), omadacId, siteId).PortForwardingConfig(portForwardingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.CreatePortForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortForwarding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.CreatePortForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portForwardingConfig** | [**PortForwardingConfig**](PortForwardingConfig.md) |  | 

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


## DeleteOtoNat

> OperationResponseWithoutResult DeleteOtoNat(ctx, omadacId, siteId, otonatId).Execute()

Delete One-to-One NAT



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
	otonatId := "otonatId_example" // string | otonatId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATAPI.DeleteOtoNat(context.Background(), omadacId, siteId, otonatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.DeleteOtoNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOtoNat`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.DeleteOtoNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**otonatId** | **string** | otonatId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOtoNatRequest struct via the builder pattern


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


## DeletePortForwarding

> OperationResponseWithoutResult DeletePortForwarding(ctx, omadacId, siteId, portForwardingId).Execute()

Delete port forwarding



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
	portForwardingId := "portForwardingId_example" // string | portForwardingId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATAPI.DeletePortForwarding(context.Background(), omadacId, siteId, portForwardingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.DeletePortForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortForwarding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.DeletePortForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portForwardingId** | **string** | portForwardingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortForwardingRequest struct via the builder pattern


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


## GetAlg

> OperationResponseGetAlgOpenApiVO GetAlg(ctx, omadacId, siteId).Execute()

Get ALG Info



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
	resp, r, err := apiClient.NATAPI.GetAlg(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.GetAlg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlg`: OperationResponseGetAlgOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.GetAlg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseGetAlgOpenApiVO**](OperationResponseGetAlgOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridOtoNats

> OperationResponseGridVOOtoNatInfoOpenApiVO GetGridOtoNats(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get One-to-One NAT list



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
	resp, r, err := apiClient.NATAPI.GetGridOtoNats(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.GetGridOtoNats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridOtoNats`: OperationResponseGridVOOtoNatInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.GetGridOtoNats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridOtoNatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOOtoNatInfoOpenApiVO**](OperationResponseGridVOOtoNatInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortForwardingList

> OperationResponsePortForwardingGridVOPortForwardingInfo GetPortForwardingList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get port forwarding list



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
	resp, r, err := apiClient.NATAPI.GetPortForwardingList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.GetPortForwardingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortForwardingList`: OperationResponsePortForwardingGridVOPortForwardingInfo
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.GetPortForwardingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortForwardingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponsePortForwardingGridVOPortForwardingInfo**](OperationResponsePortForwardingGridVOPortForwardingInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAlg

> OperationResponseWithoutResult ModifyAlg(ctx, omadacId, siteId).ALGSetting(aLGSetting).Execute()

Modify ALG setting



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
	aLGSetting := *openapiclient.NewALGSetting(false, false, false, false, false) // ALGSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATAPI.ModifyAlg(context.Background(), omadacId, siteId).ALGSetting(aLGSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.ModifyAlg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAlg`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.ModifyAlg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAlgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aLGSetting** | [**ALGSetting**](ALGSetting.md) |  | 

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


## ModifyOtoNat

> OperationResponseWithoutResult ModifyOtoNat(ctx, omadacId, siteId, otonatId).OtoNatOpenApiVO(otoNatOpenApiVO).Execute()

Modify One-to-One NAT



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
	otonatId := "otonatId_example" // string | otonatId
	otoNatOpenApiVO := *openapiclient.NewOtoNatOpenApiVO(false, "ExternalIp_example", []string{"InterfaceIds_example"}, "InternalIp_example", "Name_example", false) // OtoNatOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATAPI.ModifyOtoNat(context.Background(), omadacId, siteId, otonatId).OtoNatOpenApiVO(otoNatOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.ModifyOtoNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOtoNat`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.ModifyOtoNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**otonatId** | **string** | otonatId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOtoNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **otoNatOpenApiVO** | [**OtoNatOpenApiVO**](OtoNatOpenApiVO.md) |  | 

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


## ModifyPortForwarding

> OperationResponseWithoutResult ModifyPortForwarding(ctx, omadacId, siteId, portForwardingId).PortForwardingConfig(portForwardingConfig).Execute()

Modify port forwarding



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
	portForwardingId := "portForwardingId_example" // string | portForwardingId
	portForwardingConfig := *openapiclient.NewPortForwardingConfig(false, "ForwardIp_example", int32(123), "Name_example", false) // PortForwardingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATAPI.ModifyPortForwarding(context.Background(), omadacId, siteId, portForwardingId).PortForwardingConfig(portForwardingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATAPI.ModifyPortForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortForwarding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATAPI.ModifyPortForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portForwardingId** | **string** | portForwardingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **portForwardingConfig** | [**PortForwardingConfig**](PortForwardingConfig.md) |  | 

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


# \NATTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDisableNat**](NATTemplateAPI.md#AddDisableNat) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wired-networks/disable-nat | Create new siteTemplate&#39;s Disable Nat
[**AddOtoNatTemplates**](NATTemplateAPI.md#AddOtoNatTemplates) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/transmission/otonats | Create site template&#39;s otonat
[**CreateTemplatePortForwarding**](NATTemplateAPI.md#CreateTemplatePortForwarding) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/nat/port-forwardings | Create site template&#39;s new port forwarding
[**DeleteDisableNat**](NATTemplateAPI.md#DeleteDisableNat) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wired-networks/disable-nat/{disableNatId} | Delete an existing Disable Nat in siteTemplate
[**DeleteTemplatePortForwarding**](NATTemplateAPI.md#DeleteTemplatePortForwarding) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/nat/port-forwardings/{portForwardingId} | Delete site template&#39;s port forwarding
[**GetDisableNatGrid**](NATTemplateAPI.md#GetDisableNatGrid) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wired-networks/disable-nat | Get siteTemplate&#39;s Disable Nat list
[**GetTemplateAlg**](NATTemplateAPI.md#GetTemplateAlg) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/nat/alg | Get site template&#39;s ALG Info
[**GetTemplatePortForwardingList**](NATTemplateAPI.md#GetTemplatePortForwardingList) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/nat/port-forwardings | Get site template&#39;s port forwarding list
[**ModifyDisableNatTemplate**](NATTemplateAPI.md#ModifyDisableNatTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/wired-networks/disable-nat/{disableNatId} | Modify an existing Disable Nat in siteTemplate
[**ModifyTemplateAlg**](NATTemplateAPI.md#ModifyTemplateAlg) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/nat/alg | Modify site template&#39;s ALG setting
[**ModifyTemplatePortForwarding**](NATTemplateAPI.md#ModifyTemplatePortForwarding) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/nat/port-forwardings/{portForwardingId} | Modify site template&#39;s port forwarding



## AddDisableNat

> OperationResponseWithoutResult AddDisableNat(ctx, omadacId, siteTemplateId).DisableNat(disableNat).Execute()

Create new siteTemplate's Disable Nat



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
	disableNat := *openapiclient.NewDisableNat("Interface_example", []string{"LanList_example"}, "Name_example", false) // DisableNat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.AddDisableNat(context.Background(), omadacId, siteTemplateId).DisableNat(disableNat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.AddDisableNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDisableNat`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.AddDisableNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDisableNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disableNat** | [**DisableNat**](DisableNat.md) |  | 

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


## AddOtoNatTemplates

> OperationResponseWithoutResult AddOtoNatTemplates(ctx, omadacId, siteTemplateId).OtoNatOpenApiVO(otoNatOpenApiVO).Execute()

Create site template's otonat



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
	otoNatOpenApiVO := *openapiclient.NewOtoNatOpenApiVO(false, "ExternalIp_example", []string{"InterfaceIds_example"}, "InternalIp_example", "Name_example", false) // OtoNatOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.AddOtoNatTemplates(context.Background(), omadacId, siteTemplateId).OtoNatOpenApiVO(otoNatOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.AddOtoNatTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOtoNatTemplates`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.AddOtoNatTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOtoNatTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **otoNatOpenApiVO** | [**OtoNatOpenApiVO**](OtoNatOpenApiVO.md) |  | 

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


## CreateTemplatePortForwarding

> OperationResponseWithoutResult CreateTemplatePortForwarding(ctx, omadacId, siteTemplateId).PortForwardingTemplateOpenApiVO(portForwardingTemplateOpenApiVO).Execute()

Create site template's new port forwarding



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
	portForwardingTemplateOpenApiVO := *openapiclient.NewPortForwardingTemplateOpenApiVO(false, "ForwardIp_example", int32(123), []string{"InterfaceWanPortId_example"}, "Name_example", false) // PortForwardingTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.CreateTemplatePortForwarding(context.Background(), omadacId, siteTemplateId).PortForwardingTemplateOpenApiVO(portForwardingTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.CreateTemplatePortForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplatePortForwarding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.CreateTemplatePortForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplatePortForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portForwardingTemplateOpenApiVO** | [**PortForwardingTemplateOpenApiVO**](PortForwardingTemplateOpenApiVO.md) |  | 

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


## DeleteDisableNat

> OperationResponseWithoutResult DeleteDisableNat(ctx, omadacId, siteTemplateId, disableNatId).Execute()

Delete an existing Disable Nat in siteTemplate



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
	disableNatId := "disableNatId_example" // string | Disable Nat ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.DeleteDisableNat(context.Background(), omadacId, siteTemplateId, disableNatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.DeleteDisableNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDisableNat`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.DeleteDisableNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**disableNatId** | **string** | Disable Nat ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDisableNatRequest struct via the builder pattern


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


## DeleteTemplatePortForwarding

> OperationResponseWithoutResult DeleteTemplatePortForwarding(ctx, omadacId, siteTemplateId, portForwardingId).Execute()

Delete site template's port forwarding



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
	portForwardingId := "portForwardingId_example" // string | Port forwarding ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.DeleteTemplatePortForwarding(context.Background(), omadacId, siteTemplateId, portForwardingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.DeleteTemplatePortForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplatePortForwarding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.DeleteTemplatePortForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**portForwardingId** | **string** | Port forwarding ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplatePortForwardingRequest struct via the builder pattern


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


## GetDisableNatGrid

> OperationResponseGridVODisableNatDetailOpenApiVO GetDisableNatGrid(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get siteTemplate's Disable Nat list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.GetDisableNatGrid(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.GetDisableNatGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisableNatGrid`: OperationResponseGridVODisableNatDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.GetDisableNatGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisableNatGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVODisableNatDetailOpenApiVO**](OperationResponseGridVODisableNatDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateAlg

> OperationResponseALGSetting GetTemplateAlg(ctx, omadacId, siteTemplateId).Execute()

Get site template's ALG Info



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
	resp, r, err := apiClient.NATTemplateAPI.GetTemplateAlg(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.GetTemplateAlg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateAlg`: OperationResponseALGSetting
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.GetTemplateAlg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateAlgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseALGSetting**](OperationResponseALGSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatePortForwardingList

> OperationResponsePortForwardingOpenApiGridVOPortForwardingInfoTemplate GetTemplatePortForwardingList(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get site template's port forwarding list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.GetTemplatePortForwardingList(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.GetTemplatePortForwardingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplatePortForwardingList`: OperationResponsePortForwardingOpenApiGridVOPortForwardingInfoTemplate
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.GetTemplatePortForwardingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatePortForwardingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponsePortForwardingOpenApiGridVOPortForwardingInfoTemplate**](OperationResponsePortForwardingOpenApiGridVOPortForwardingInfoTemplate.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDisableNatTemplate

> OperationResponseWithoutResult ModifyDisableNatTemplate(ctx, omadacId, siteTemplateId, disableNatId).DisableNat(disableNat).Execute()

Modify an existing Disable Nat in siteTemplate



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
	disableNatId := "disableNatId_example" // string | Disable Nat ID
	disableNat := *openapiclient.NewDisableNat("Interface_example", []string{"LanList_example"}, "Name_example", false) // DisableNat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.ModifyDisableNatTemplate(context.Background(), omadacId, siteTemplateId, disableNatId).DisableNat(disableNat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.ModifyDisableNatTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDisableNatTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.ModifyDisableNatTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**disableNatId** | **string** | Disable Nat ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDisableNatTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **disableNat** | [**DisableNat**](DisableNat.md) |  | 

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


## ModifyTemplateAlg

> OperationResponseWithoutResult ModifyTemplateAlg(ctx, omadacId, siteTemplateId).ALGSetting(aLGSetting).Execute()

Modify site template's ALG setting



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
	aLGSetting := *openapiclient.NewALGSetting(false, false, false, false, false) // ALGSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.ModifyTemplateAlg(context.Background(), omadacId, siteTemplateId).ALGSetting(aLGSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.ModifyTemplateAlg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateAlg`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.ModifyTemplateAlg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateAlgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aLGSetting** | [**ALGSetting**](ALGSetting.md) |  | 

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


## ModifyTemplatePortForwarding

> OperationResponseWithoutResult ModifyTemplatePortForwarding(ctx, omadacId, siteTemplateId, portForwardingId).PortForwardingTemplateOpenApiVO(portForwardingTemplateOpenApiVO).Execute()

Modify site template's port forwarding



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
	portForwardingId := "portForwardingId_example" // string | Port forwarding ID
	portForwardingTemplateOpenApiVO := *openapiclient.NewPortForwardingTemplateOpenApiVO(false, "ForwardIp_example", int32(123), []string{"InterfaceWanPortId_example"}, "Name_example", false) // PortForwardingTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NATTemplateAPI.ModifyTemplatePortForwarding(context.Background(), omadacId, siteTemplateId, portForwardingId).PortForwardingTemplateOpenApiVO(portForwardingTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NATTemplateAPI.ModifyTemplatePortForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplatePortForwarding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `NATTemplateAPI.ModifyTemplatePortForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**portForwardingId** | **string** | Port forwarding ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplatePortForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **portForwardingTemplateOpenApiVO** | [**PortForwardingTemplateOpenApiVO**](PortForwardingTemplateOpenApiVO.md) |  | 

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


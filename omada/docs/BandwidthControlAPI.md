# \BandwidthControlAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBandwidthCtrlRule**](BandwidthControlAPI.md#CreateBandwidthCtrlRule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/bandwidth-control/rules | Create bandwidth control rule
[**DeleteBandwidthCtrlRule**](BandwidthControlAPI.md#DeleteBandwidthCtrlRule) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/bandwidth-control/rules/{ruleId} | Delete bandwidth control rule
[**GetBandwidthCtrl**](BandwidthControlAPI.md#GetBandwidthCtrl) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/bandwidth-control | Get bandwidth control
[**GetGridBandwidthCtrlRule**](BandwidthControlAPI.md#GetGridBandwidthCtrlRule) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/bandwidth-control/rules | Get bandwidth control rule list
[**ModifyBandwidthCtrl**](BandwidthControlAPI.md#ModifyBandwidthCtrl) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/bandwidth-control | Modify bandwidth control
[**ModifyBandwidthCtrlRule**](BandwidthControlAPI.md#ModifyBandwidthCtrlRule) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/bandwidth-control/rules/{ruleId} | Modify bandwidth control rule
[**ModifyBandwidthCtrlRuleIndex**](BandwidthControlAPI.md#ModifyBandwidthCtrlRuleIndex) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/bandwidth-control/rules | Modify bandwidth control rule index



## CreateBandwidthCtrlRule

> OperationResponseWithoutResult CreateBandwidthCtrlRule(ctx, omadacId, siteId).BandwidthControlRule(bandwidthControlRule).Execute()

Create bandwidth control rule



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
	bandwidthControlRule := *openapiclient.NewBandwidthControlRule(int32(123), int32(123), int32(123), "Name_example", []string{"SourceIds_example"}, false, int32(123), int32(123), []string{"WanPortIds_example"}) // BandwidthControlRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlAPI.CreateBandwidthCtrlRule(context.Background(), omadacId, siteId).BandwidthControlRule(bandwidthControlRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlAPI.CreateBandwidthCtrlRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBandwidthCtrlRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlAPI.CreateBandwidthCtrlRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBandwidthCtrlRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bandwidthControlRule** | [**BandwidthControlRule**](BandwidthControlRule.md) |  | 

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


## DeleteBandwidthCtrlRule

> OperationResponseWithoutResult DeleteBandwidthCtrlRule(ctx, omadacId, siteId, ruleId).Execute()

Delete bandwidth control rule



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
	ruleId := "ruleId_example" // string | Bandwidth control rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlAPI.DeleteBandwidthCtrlRule(context.Background(), omadacId, siteId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlAPI.DeleteBandwidthCtrlRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBandwidthCtrlRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlAPI.DeleteBandwidthCtrlRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ruleId** | **string** | Bandwidth control rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBandwidthCtrlRuleRequest struct via the builder pattern


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


## GetBandwidthCtrl

> OperationResponseBandwidthControl GetBandwidthCtrl(ctx, omadacId, siteId).Execute()

Get bandwidth control



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
	resp, r, err := apiClient.BandwidthControlAPI.GetBandwidthCtrl(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlAPI.GetBandwidthCtrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandwidthCtrl`: OperationResponseBandwidthControl
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlAPI.GetBandwidthCtrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandwidthCtrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseBandwidthControl**](OperationResponseBandwidthControl.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridBandwidthCtrlRule

> OperationResponseGridVOBandwidthControlRule GetGridBandwidthCtrlRule(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get bandwidth control rule list



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
	resp, r, err := apiClient.BandwidthControlAPI.GetGridBandwidthCtrlRule(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlAPI.GetGridBandwidthCtrlRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridBandwidthCtrlRule`: OperationResponseGridVOBandwidthControlRule
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlAPI.GetGridBandwidthCtrlRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridBandwidthCtrlRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOBandwidthControlRule**](OperationResponseGridVOBandwidthControlRule.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyBandwidthCtrl

> OperationResponseWithoutResult ModifyBandwidthCtrl(ctx, omadacId, siteId).BandwidthControl(bandwidthControl).Execute()

Modify bandwidth control



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
	bandwidthControl := *openapiclient.NewBandwidthControl(false) // BandwidthControl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlAPI.ModifyBandwidthCtrl(context.Background(), omadacId, siteId).BandwidthControl(bandwidthControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlAPI.ModifyBandwidthCtrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBandwidthCtrl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlAPI.ModifyBandwidthCtrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBandwidthCtrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bandwidthControl** | [**BandwidthControl**](BandwidthControl.md) |  | 

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


## ModifyBandwidthCtrlRule

> OperationResponseWithoutResult ModifyBandwidthCtrlRule(ctx, omadacId, siteId, ruleId).BandwidthControlRule(bandwidthControlRule).Execute()

Modify bandwidth control rule



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
	ruleId := "ruleId_example" // string | Bandwidth control rule ID
	bandwidthControlRule := *openapiclient.NewBandwidthControlRule(int32(123), int32(123), int32(123), "Name_example", []string{"SourceIds_example"}, false, int32(123), int32(123), []string{"WanPortIds_example"}) // BandwidthControlRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlAPI.ModifyBandwidthCtrlRule(context.Background(), omadacId, siteId, ruleId).BandwidthControlRule(bandwidthControlRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlAPI.ModifyBandwidthCtrlRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBandwidthCtrlRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlAPI.ModifyBandwidthCtrlRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ruleId** | **string** | Bandwidth control rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBandwidthCtrlRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bandwidthControlRule** | [**BandwidthControlRule**](BandwidthControlRule.md) |  | 

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


## ModifyBandwidthCtrlRuleIndex

> OperationResponseWithoutResult ModifyBandwidthCtrlRuleIndex(ctx, omadacId, siteId).TransmissionDragSortIndexOpenapiVO(transmissionDragSortIndexOpenapiVO).Execute()

Modify bandwidth control rule index



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
	transmissionDragSortIndexOpenapiVO := *openapiclient.NewTransmissionDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}) // TransmissionDragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlAPI.ModifyBandwidthCtrlRuleIndex(context.Background(), omadacId, siteId).TransmissionDragSortIndexOpenapiVO(transmissionDragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlAPI.ModifyBandwidthCtrlRuleIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBandwidthCtrlRuleIndex`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlAPI.ModifyBandwidthCtrlRuleIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBandwidthCtrlRuleIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transmissionDragSortIndexOpenapiVO** | [**TransmissionDragSortIndexOpenapiVO**](TransmissionDragSortIndexOpenapiVO.md) |  | 

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


# BandwidthControlTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBandwidthCtrlRuleTemplate**](BandwidthControlTemplateAPI.md#createbandwidthctrlruletemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/bandwidth-control/rules | Create bandwidth control template rule
[**DeleteBandwidthCtrlRuleTemplate**](BandwidthControlTemplateAPI.md#deletebandwidthctrlruletemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/bandwidth-control/rules/{ruleId} | Delete bandwidth control template rule
[**GetBandwidthCtrlTemplate**](BandwidthControlTemplateAPI.md#getbandwidthctrltemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/bandwidth-control | Get bandwidth control template
[**GetGridBandwidthCtrlRuleTemplate**](BandwidthControlTemplateAPI.md#getgridbandwidthctrlruletemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/bandwidth-control/rules | Get bandwidth control template rule list
[**ModifyBandwidthCtrlRuleIndexTemplate**](BandwidthControlTemplateAPI.md#modifybandwidthctrlruleindextemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/bandwidth-control/rules | Modify bandwidth control template rule index
[**ModifyBandwidthCtrlRuleTemplate**](BandwidthControlTemplateAPI.md#modifybandwidthctrlruletemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/bandwidth-control/rules/{ruleId} | Modify bandwidth control template rule
[**ModifyBandwidthCtrlTemplate**](BandwidthControlTemplateAPI.md#modifybandwidthctrltemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/bandwidth-control | Modify bandwidth control template



## CreateBandwidthCtrlRuleTemplate

> OperationResponseWithoutResult CreateBandwidthCtrlRuleTemplate(ctx, omadacId, siteTemplateId).BandwidthControlRule(bandwidthControlRule).Execute()

Create bandwidth control template rule



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
	bandwidthControlRule := *openapiclient.NewBandwidthControlRule(int32(123), int32(123), int32(123), "Name_example", []string{"SourceIds_example"}, false, int32(123), int32(123), []string{"WanPortIds_example"}) // BandwidthControlRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlTemplateAPI.CreateBandwidthCtrlRuleTemplate(context.Background(), omadacId, siteTemplateId).BandwidthControlRule(bandwidthControlRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlTemplateAPI.CreateBandwidthCtrlRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBandwidthCtrlRuleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlTemplateAPI.CreateBandwidthCtrlRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBandwidthCtrlRuleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bandwidthControlRule** | [**BandwidthControlRule**](BandwidthControlRule.md) |  | 

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


## DeleteBandwidthCtrlRuleTemplate

> OperationResponseWithoutResult DeleteBandwidthCtrlRuleTemplate(ctx, omadacId, siteTemplateId, ruleId).Execute()

Delete bandwidth control template rule



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
	ruleId := "ruleId_example" // string | Bandwidth control template rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlTemplateAPI.DeleteBandwidthCtrlRuleTemplate(context.Background(), omadacId, siteTemplateId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlTemplateAPI.DeleteBandwidthCtrlRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBandwidthCtrlRuleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlTemplateAPI.DeleteBandwidthCtrlRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ruleId** | **string** | Bandwidth control template rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBandwidthCtrlRuleTemplateRequest struct via the builder pattern


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


## GetBandwidthCtrlTemplate

> OperationResponseBandwidthControl GetBandwidthCtrlTemplate(ctx, omadacId, siteTemplateId).Execute()

Get bandwidth control template



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
	resp, r, err := apiClient.BandwidthControlTemplateAPI.GetBandwidthCtrlTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlTemplateAPI.GetBandwidthCtrlTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandwidthCtrlTemplate`: OperationResponseBandwidthControl
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlTemplateAPI.GetBandwidthCtrlTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandwidthCtrlTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseBandwidthControl**](OperationResponseBandwidthControl.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridBandwidthCtrlRuleTemplate

> OperationResponseGridVOBandwidthControlRule GetGridBandwidthCtrlRuleTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get bandwidth control template rule list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlTemplateAPI.GetGridBandwidthCtrlRuleTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlTemplateAPI.GetGridBandwidthCtrlRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridBandwidthCtrlRuleTemplate`: OperationResponseGridVOBandwidthControlRule
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlTemplateAPI.GetGridBandwidthCtrlRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridBandwidthCtrlRuleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOBandwidthControlRule**](OperationResponseGridVOBandwidthControlRule.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyBandwidthCtrlRuleIndexTemplate

> OperationResponseWithoutResult ModifyBandwidthCtrlRuleIndexTemplate(ctx, omadacId, siteTemplateId).TransmissionDragSortIndexOpenapiVO(transmissionDragSortIndexOpenapiVO).Execute()

Modify bandwidth control template rule index



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
	transmissionDragSortIndexOpenapiVO := *openapiclient.NewTransmissionDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}) // TransmissionDragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlTemplateAPI.ModifyBandwidthCtrlRuleIndexTemplate(context.Background(), omadacId, siteTemplateId).TransmissionDragSortIndexOpenapiVO(transmissionDragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlTemplateAPI.ModifyBandwidthCtrlRuleIndexTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBandwidthCtrlRuleIndexTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlTemplateAPI.ModifyBandwidthCtrlRuleIndexTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBandwidthCtrlRuleIndexTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transmissionDragSortIndexOpenapiVO** | [**TransmissionDragSortIndexOpenapiVO**](TransmissionDragSortIndexOpenapiVO.md) |  | 

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


## ModifyBandwidthCtrlRuleTemplate

> OperationResponseWithoutResult ModifyBandwidthCtrlRuleTemplate(ctx, omadacId, siteTemplateId, ruleId).BandwidthControlRule(bandwidthControlRule).Execute()

Modify bandwidth control template rule



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
	ruleId := "ruleId_example" // string | Bandwidth control template rule ID
	bandwidthControlRule := *openapiclient.NewBandwidthControlRule(int32(123), int32(123), int32(123), "Name_example", []string{"SourceIds_example"}, false, int32(123), int32(123), []string{"WanPortIds_example"}) // BandwidthControlRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlTemplateAPI.ModifyBandwidthCtrlRuleTemplate(context.Background(), omadacId, siteTemplateId, ruleId).BandwidthControlRule(bandwidthControlRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlTemplateAPI.ModifyBandwidthCtrlRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBandwidthCtrlRuleTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlTemplateAPI.ModifyBandwidthCtrlRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ruleId** | **string** | Bandwidth control template rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBandwidthCtrlRuleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bandwidthControlRule** | [**BandwidthControlRule**](BandwidthControlRule.md) |  | 

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


## ModifyBandwidthCtrlTemplate

> OperationResponseWithoutResult ModifyBandwidthCtrlTemplate(ctx, omadacId, siteTemplateId).BandwidthControl(bandwidthControl).Execute()

Modify bandwidth control template



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
	bandwidthControl := *openapiclient.NewBandwidthControl(false) // BandwidthControl | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BandwidthControlTemplateAPI.ModifyBandwidthCtrlTemplate(context.Background(), omadacId, siteTemplateId).BandwidthControl(bandwidthControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BandwidthControlTemplateAPI.ModifyBandwidthCtrlTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBandwidthCtrlTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BandwidthControlTemplateAPI.ModifyBandwidthCtrlTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBandwidthCtrlTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bandwidthControl** | [**BandwidthControl**](BandwidthControl.md) |  | 

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


# \GatewayQOSTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplateBwCtrl**](GatewayQOSTemplateAPI.md#CreateTemplateBwCtrl) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/bwcs | Create siteTemplate&#39;s new bandwidth control rule
[**CreateTemplateClassRule**](GatewayQOSTemplateAPI.md#CreateTemplateClassRule) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/class-rules | Create siteTemplate&#39;s new class rule
[**DeleteTemplateBwCtrl**](GatewayQOSTemplateAPI.md#DeleteTemplateBwCtrl) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/bwcs/{id} | Delete an existing bandwidth control rule in siteTemplate
[**DeleteTemplateClassRule**](GatewayQOSTemplateAPI.md#DeleteTemplateClassRule) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/class-rules/{id} | Delete an existing class rule in siteTemplate
[**GetTemplateBandwidthCtrlGrid**](GatewayQOSTemplateAPI.md#GetTemplateBandwidthCtrlGrid) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/bwcs | Get siteTemplate&#39;s bandwidth control rule list
[**GetTemplateClassRuleGrid**](GatewayQOSTemplateAPI.md#GetTemplateClassRuleGrid) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/class-rules | Get siteTemplate&#39;s class rule list
[**GetTemplateQosWans**](GatewayQOSTemplateAPI.md#GetTemplateQosWans) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/wans | Get siteTemplate&#39;s WAN ports info for Gateway QoS
[**GetTemplateTagOut**](GatewayQOSTemplateAPI.md#GetTemplateTagOut) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/tag-outbound-traffic | Get siteTemplate&#39;s Tag Outbound Traffic info
[**GetTemplateVoip**](GatewayQOSTemplateAPI.md#GetTemplateVoip) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/voip-prioritization | Get siteTemplate&#39;s VoIP Prioritization info
[**ModifyTemplateBwCtrl**](GatewayQOSTemplateAPI.md#ModifyTemplateBwCtrl) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/bwcs/{id} | Modify an existing bandwidth control rule in siteTemplate
[**ModifyTemplateClassRule**](GatewayQOSTemplateAPI.md#ModifyTemplateClassRule) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/class-rules/{id} | Modify an existing class rule in siteTemplate
[**ModifyTemplateTagOut**](GatewayQOSTemplateAPI.md#ModifyTemplateTagOut) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/tag-outbound-traffic | Modify siteTemplate&#39;s Tag Outbound Traffic info
[**ModifyTemplateVoip**](GatewayQOSTemplateAPI.md#ModifyTemplateVoip) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/qos/gateway/voip-prioritization | Modify siteTemplate&#39;s VoIP Prioritization info



## CreateTemplateBwCtrl

> OperationResponseResponseIdVO CreateTemplateBwCtrl(ctx, omadacId, siteTemplateId).QosBwcOpenApiVO(qosBwcOpenApiVO).Execute()

Create siteTemplate's new bandwidth control rule



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
	qosBwcOpenApiVO := *openapiclient.NewQosBwcOpenApiVO([]int32{int32(123)}, int32(123), int32(123), int32(123), false, false, false, "Wan_example") // QosBwcOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQOSTemplateAPI.CreateTemplateBwCtrl(context.Background(), omadacId, siteTemplateId).QosBwcOpenApiVO(qosBwcOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.CreateTemplateBwCtrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateBwCtrl`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.CreateTemplateBwCtrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateBwCtrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **qosBwcOpenApiVO** | [**QosBwcOpenApiVO**](QosBwcOpenApiVO.md) |  | 

### Return type

[**OperationResponseResponseIdVO**](OperationResponseResponseIdVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplateClassRule

> OperationResponseResponseIdVO CreateTemplateClassRule(ctx, omadacId, siteTemplateId).ClassRuleTemplateOpenApiVO(classRuleTemplateOpenApiVO).Execute()

Create siteTemplate's new class rule



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
	classRuleTemplateOpenApiVO := *openapiclient.NewClassRuleTemplateOpenApiVO(int32(123), "Dscp_example", false, int32(123), "LocalIp_example", "RemoteIp_example", "ServiceType_example") // ClassRuleTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQOSTemplateAPI.CreateTemplateClassRule(context.Background(), omadacId, siteTemplateId).ClassRuleTemplateOpenApiVO(classRuleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.CreateTemplateClassRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateClassRule`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.CreateTemplateClassRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateClassRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **classRuleTemplateOpenApiVO** | [**ClassRuleTemplateOpenApiVO**](ClassRuleTemplateOpenApiVO.md) |  | 

### Return type

[**OperationResponseResponseIdVO**](OperationResponseResponseIdVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplateBwCtrl

> OperationResponseWithoutResult DeleteTemplateBwCtrl(ctx, omadacId, siteTemplateId, id).Execute()

Delete an existing bandwidth control rule in siteTemplate



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
	id := "id_example" // string | Bandwidth control rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQOSTemplateAPI.DeleteTemplateBwCtrl(context.Background(), omadacId, siteTemplateId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.DeleteTemplateBwCtrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplateBwCtrl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.DeleteTemplateBwCtrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Bandwidth control rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateBwCtrlRequest struct via the builder pattern


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


## DeleteTemplateClassRule

> OperationResponseWithoutResult DeleteTemplateClassRule(ctx, omadacId, siteTemplateId, id).Execute()

Delete an existing class rule in siteTemplate



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
	id := "id_example" // string | Class rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQOSTemplateAPI.DeleteTemplateClassRule(context.Background(), omadacId, siteTemplateId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.DeleteTemplateClassRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplateClassRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.DeleteTemplateClassRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Class rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateClassRuleRequest struct via the builder pattern


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


## GetTemplateBandwidthCtrlGrid

> OperationResponseGridVOQosBwcDetailOpenApiVO GetTemplateBandwidthCtrlGrid(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get siteTemplate's bandwidth control rule list



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
	resp, r, err := apiClient.GatewayQOSTemplateAPI.GetTemplateBandwidthCtrlGrid(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.GetTemplateBandwidthCtrlGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateBandwidthCtrlGrid`: OperationResponseGridVOQosBwcDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.GetTemplateBandwidthCtrlGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateBandwidthCtrlGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOQosBwcDetailOpenApiVO**](OperationResponseGridVOQosBwcDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateClassRuleGrid

> OperationResponseGridVOClassRuleTemplateDetailOpenApiVO GetTemplateClassRuleGrid(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get siteTemplate's class rule list



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
	resp, r, err := apiClient.GatewayQOSTemplateAPI.GetTemplateClassRuleGrid(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.GetTemplateClassRuleGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateClassRuleGrid`: OperationResponseGridVOClassRuleTemplateDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.GetTemplateClassRuleGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateClassRuleGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOClassRuleTemplateDetailOpenApiVO**](OperationResponseGridVOClassRuleTemplateDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateQosWans

> OperationResponseResponseDataVOQosBwcWanInfoOpenApiVO GetTemplateQosWans(ctx, omadacId, siteTemplateId).Execute()

Get siteTemplate's WAN ports info for Gateway QoS



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
	resp, r, err := apiClient.GatewayQOSTemplateAPI.GetTemplateQosWans(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.GetTemplateQosWans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateQosWans`: OperationResponseResponseDataVOQosBwcWanInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.GetTemplateQosWans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateQosWansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseResponseDataVOQosBwcWanInfoOpenApiVO**](OperationResponseResponseDataVOQosBwcWanInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateTagOut

> OperationResponseTagOutConfigOpenApiVO GetTemplateTagOut(ctx, omadacId, siteTemplateId).Execute()

Get siteTemplate's Tag Outbound Traffic info



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
	resp, r, err := apiClient.GatewayQOSTemplateAPI.GetTemplateTagOut(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.GetTemplateTagOut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateTagOut`: OperationResponseTagOutConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.GetTemplateTagOut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateTagOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseTagOutConfigOpenApiVO**](OperationResponseTagOutConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateVoip

> OperationResponseVoipOpenApiVO GetTemplateVoip(ctx, omadacId, siteTemplateId).Execute()

Get siteTemplate's VoIP Prioritization info



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
	resp, r, err := apiClient.GatewayQOSTemplateAPI.GetTemplateVoip(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.GetTemplateVoip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateVoip`: OperationResponseVoipOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.GetTemplateVoip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateVoipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVoipOpenApiVO**](OperationResponseVoipOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTemplateBwCtrl

> OperationResponseWithoutResult ModifyTemplateBwCtrl(ctx, omadacId, siteTemplateId, id).QosBwcEditOpenApiVO(qosBwcEditOpenApiVO).Execute()

Modify an existing bandwidth control rule in siteTemplate



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
	id := "id_example" // string | Bandwidth control rule ID
	qosBwcEditOpenApiVO := *openapiclient.NewQosBwcEditOpenApiVO([]int32{int32(123)}, int32(123), int32(123), int32(123), false, false, false) // QosBwcEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQOSTemplateAPI.ModifyTemplateBwCtrl(context.Background(), omadacId, siteTemplateId, id).QosBwcEditOpenApiVO(qosBwcEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.ModifyTemplateBwCtrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateBwCtrl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.ModifyTemplateBwCtrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Bandwidth control rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateBwCtrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **qosBwcEditOpenApiVO** | [**QosBwcEditOpenApiVO**](QosBwcEditOpenApiVO.md) |  | 

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


## ModifyTemplateClassRule

> OperationResponseWithoutResult ModifyTemplateClassRule(ctx, omadacId, siteTemplateId, id).ClassRuleTemplateOpenApiVO(classRuleTemplateOpenApiVO).Execute()

Modify an existing class rule in siteTemplate



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
	id := "id_example" // string | Class rule ID
	classRuleTemplateOpenApiVO := *openapiclient.NewClassRuleTemplateOpenApiVO(int32(123), "Dscp_example", false, int32(123), "LocalIp_example", "RemoteIp_example", "ServiceType_example") // ClassRuleTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQOSTemplateAPI.ModifyTemplateClassRule(context.Background(), omadacId, siteTemplateId, id).ClassRuleTemplateOpenApiVO(classRuleTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.ModifyTemplateClassRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateClassRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.ModifyTemplateClassRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Class rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateClassRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **classRuleTemplateOpenApiVO** | [**ClassRuleTemplateOpenApiVO**](ClassRuleTemplateOpenApiVO.md) |  | 

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


## ModifyTemplateTagOut

> OperationResponseWithoutResult ModifyTemplateTagOut(ctx, omadacId, siteTemplateId).TagOutConfigOpenApiVO(tagOutConfigOpenApiVO).Execute()

Modify siteTemplate's Tag Outbound Traffic info



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
	tagOutConfigOpenApiVO := *openapiclient.NewTagOutConfigOpenApiVO([]openapiclient.TagOutItemOpenApiVO{*openapiclient.NewTagOutItemOpenApiVO(int32(123), "Dscp_example", false)}) // TagOutConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQOSTemplateAPI.ModifyTemplateTagOut(context.Background(), omadacId, siteTemplateId).TagOutConfigOpenApiVO(tagOutConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.ModifyTemplateTagOut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateTagOut`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.ModifyTemplateTagOut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateTagOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagOutConfigOpenApiVO** | [**TagOutConfigOpenApiVO**](TagOutConfigOpenApiVO.md) |  | 

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


## ModifyTemplateVoip

> OperationResponseWithoutResult ModifyTemplateVoip(ctx, omadacId, siteTemplateId).VoipOpenApiVO(voipOpenApiVO).Execute()

Modify siteTemplate's VoIP Prioritization info



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
	voipOpenApiVO := *openapiclient.NewVoipOpenApiVO(false) // VoipOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQOSTemplateAPI.ModifyTemplateVoip(context.Background(), omadacId, siteTemplateId).VoipOpenApiVO(voipOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQOSTemplateAPI.ModifyTemplateVoip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateVoip`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQOSTemplateAPI.ModifyTemplateVoip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateVoipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voipOpenApiVO** | [**VoipOpenApiVO**](VoipOpenApiVO.md) |  | 

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


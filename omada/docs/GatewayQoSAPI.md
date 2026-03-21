# \GatewayQoSAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBwCtrl**](GatewayQoSAPI.md#CreateBwCtrl) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/bwcs | Create new bandwidth control rule
[**CreateClassRule**](GatewayQoSAPI.md#CreateClassRule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/class-rules | Create new class rule
[**DeleteBwCtrl**](GatewayQoSAPI.md#DeleteBwCtrl) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/bwcs/{id} | Delete an existing bandwidth control rule
[**DeleteClassRule**](GatewayQoSAPI.md#DeleteClassRule) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/class-rules/{id} | Delete an existing class rule
[**GetBandwidthCtrlGrid**](GatewayQoSAPI.md#GetBandwidthCtrlGrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/bwcs | Get bandwidth control rule list
[**GetClassRuleGrid**](GatewayQoSAPI.md#GetClassRuleGrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/class-rules | Get class rule list
[**GetQosWans**](GatewayQoSAPI.md#GetQosWans) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/wans | Get WAN ports info for Gateway QoS
[**GetTagOut**](GatewayQoSAPI.md#GetTagOut) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/tag-outbound-traffic | Get Tag Outbound Traffic info
[**GetVoip**](GatewayQoSAPI.md#GetVoip) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/voip-prioritization | Get VoIP Prioritization info
[**ModifyBwCtrl**](GatewayQoSAPI.md#ModifyBwCtrl) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/bwcs/{id} | Modify an existing bandwidth control rule
[**ModifyClassRule**](GatewayQoSAPI.md#ModifyClassRule) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/class-rules/{id} | Modify an existing class rule
[**ModifyTagOut**](GatewayQoSAPI.md#ModifyTagOut) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/tag-outbound-traffic | Modify Tag Outbound Traffic info
[**ModifyVoip**](GatewayQoSAPI.md#ModifyVoip) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/qos/gateway/voip-prioritization | Modify VoIP Prioritization info



## CreateBwCtrl

> OperationResponseResponseIdVO CreateBwCtrl(ctx, omadacId, siteId).QosBwcOpenApiVO(qosBwcOpenApiVO).Execute()

Create new bandwidth control rule



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
	qosBwcOpenApiVO := *openapiclient.NewQosBwcOpenApiVO([]int32{int32(123)}, int32(123), int32(123), int32(123), false, false, false, "Wan_example") // QosBwcOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQoSAPI.CreateBwCtrl(context.Background(), omadacId, siteId).QosBwcOpenApiVO(qosBwcOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.CreateBwCtrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBwCtrl`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.CreateBwCtrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBwCtrlRequest struct via the builder pattern


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


## CreateClassRule

> OperationResponseResponseIdVO CreateClassRule(ctx, omadacId, siteId).ClassRuleOpenApiVO(classRuleOpenApiVO).Execute()

Create new class rule



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
	classRuleOpenApiVO := *openapiclient.NewClassRuleOpenApiVO(int32(123), "Dscp_example", false, int32(123), "LocalIp_example", "RemoteIp_example", "ServiceType_example") // ClassRuleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQoSAPI.CreateClassRule(context.Background(), omadacId, siteId).ClassRuleOpenApiVO(classRuleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.CreateClassRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClassRule`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.CreateClassRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClassRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **classRuleOpenApiVO** | [**ClassRuleOpenApiVO**](ClassRuleOpenApiVO.md) |  | 

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


## DeleteBwCtrl

> OperationResponseWithoutResult DeleteBwCtrl(ctx, omadacId, siteId, id).Execute()

Delete an existing bandwidth control rule



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
	id := "id_example" // string | Bandwidth control rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQoSAPI.DeleteBwCtrl(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.DeleteBwCtrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBwCtrl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.DeleteBwCtrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Bandwidth control rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBwCtrlRequest struct via the builder pattern


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


## DeleteClassRule

> OperationResponseWithoutResult DeleteClassRule(ctx, omadacId, siteId, id).Execute()

Delete an existing class rule



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
	id := "id_example" // string | Class rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQoSAPI.DeleteClassRule(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.DeleteClassRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClassRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.DeleteClassRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Class rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClassRuleRequest struct via the builder pattern


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


## GetBandwidthCtrlGrid

> OperationResponseGridVOQosBwcDetailOpenApiVO GetBandwidthCtrlGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

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
	resp, r, err := apiClient.GatewayQoSAPI.GetBandwidthCtrlGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.GetBandwidthCtrlGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandwidthCtrlGrid`: OperationResponseGridVOQosBwcDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.GetBandwidthCtrlGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandwidthCtrlGridRequest struct via the builder pattern


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


## GetClassRuleGrid

> OperationResponseGridVOClassRuleDetailOpenApiVO GetClassRuleGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get class rule list



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
	resp, r, err := apiClient.GatewayQoSAPI.GetClassRuleGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.GetClassRuleGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClassRuleGrid`: OperationResponseGridVOClassRuleDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.GetClassRuleGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClassRuleGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOClassRuleDetailOpenApiVO**](OperationResponseGridVOClassRuleDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQosWans

> OperationResponseResponseDataVOQosBwcWanInfoOpenApiVO GetQosWans(ctx, omadacId, siteId).Execute()

Get WAN ports info for Gateway QoS



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
	resp, r, err := apiClient.GatewayQoSAPI.GetQosWans(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.GetQosWans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQosWans`: OperationResponseResponseDataVOQosBwcWanInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.GetQosWans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQosWansRequest struct via the builder pattern


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


## GetTagOut

> OperationResponseTagOutConfigOpenApiVO GetTagOut(ctx, omadacId, siteId).Execute()

Get Tag Outbound Traffic info



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
	resp, r, err := apiClient.GatewayQoSAPI.GetTagOut(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.GetTagOut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagOut`: OperationResponseTagOutConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.GetTagOut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagOutRequest struct via the builder pattern


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


## GetVoip

> OperationResponseVoipOpenApiVO GetVoip(ctx, omadacId, siteId).Execute()

Get VoIP Prioritization info



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
	resp, r, err := apiClient.GatewayQoSAPI.GetVoip(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.GetVoip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoip`: OperationResponseVoipOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.GetVoip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipRequest struct via the builder pattern


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


## ModifyBwCtrl

> OperationResponseWithoutResult ModifyBwCtrl(ctx, omadacId, siteId, id).QosBwcEditOpenApiVO(qosBwcEditOpenApiVO).Execute()

Modify an existing bandwidth control rule



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
	id := "id_example" // string | Bandwidth control rule ID
	qosBwcEditOpenApiVO := *openapiclient.NewQosBwcEditOpenApiVO([]int32{int32(123)}, int32(123), int32(123), int32(123), false, false, false) // QosBwcEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQoSAPI.ModifyBwCtrl(context.Background(), omadacId, siteId, id).QosBwcEditOpenApiVO(qosBwcEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.ModifyBwCtrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBwCtrl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.ModifyBwCtrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Bandwidth control rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBwCtrlRequest struct via the builder pattern


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


## ModifyClassRule

> OperationResponseWithoutResult ModifyClassRule(ctx, omadacId, siteId, id).ClassRuleOpenApiVO(classRuleOpenApiVO).Execute()

Modify an existing class rule



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
	id := "id_example" // string | Class rule ID
	classRuleOpenApiVO := *openapiclient.NewClassRuleOpenApiVO(int32(123), "Dscp_example", false, int32(123), "LocalIp_example", "RemoteIp_example", "ServiceType_example") // ClassRuleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQoSAPI.ModifyClassRule(context.Background(), omadacId, siteId, id).ClassRuleOpenApiVO(classRuleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.ModifyClassRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyClassRule`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.ModifyClassRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Class rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyClassRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **classRuleOpenApiVO** | [**ClassRuleOpenApiVO**](ClassRuleOpenApiVO.md) |  | 

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


## ModifyTagOut

> OperationResponseWithoutResult ModifyTagOut(ctx, omadacId, siteId).TagOutConfigOpenApiVO(tagOutConfigOpenApiVO).Execute()

Modify Tag Outbound Traffic info



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
	tagOutConfigOpenApiVO := *openapiclient.NewTagOutConfigOpenApiVO([]openapiclient.TagOutItemOpenApiVO{*openapiclient.NewTagOutItemOpenApiVO(int32(123), "Dscp_example", false)}) // TagOutConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQoSAPI.ModifyTagOut(context.Background(), omadacId, siteId).TagOutConfigOpenApiVO(tagOutConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.ModifyTagOut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTagOut`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.ModifyTagOut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTagOutRequest struct via the builder pattern


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


## ModifyVoip

> OperationResponseWithoutResult ModifyVoip(ctx, omadacId, siteId).VoipOpenApiVO(voipOpenApiVO).Execute()

Modify VoIP Prioritization info



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
	voipOpenApiVO := *openapiclient.NewVoipOpenApiVO(false) // VoipOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayQoSAPI.ModifyVoip(context.Background(), omadacId, siteId).VoipOpenApiVO(voipOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayQoSAPI.ModifyVoip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoip`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayQoSAPI.ModifyVoip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipRequest struct via the builder pattern


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


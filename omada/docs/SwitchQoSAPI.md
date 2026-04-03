# SwitchQoSAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQosRule**](SwitchQoSAPI.md#createqosrule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule | Create switch QoS rule.
[**DeleteQosRule**](SwitchQoSAPI.md#deleteqosrule) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/{qosRuleId} | Delete switch QoS rule.
[**GetAllSelectableDevices**](SwitchQoSAPI.md#getallselectabledevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/selectable-devices/all | Get switch QoS all selectable device&#39;s mac.
[**GetDscpMapping**](SwitchQoSAPI.md#getdscpmapping) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/dscp-mapping | Get switch QoS dscp mapping.
[**GetDscpMappingTemplate**](SwitchQoSAPI.md#getdscpmappingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-qos/qos-rule/dscp-mapping | Get switch QoS dscp mapping in site template.
[**GetOswQosMode**](SwitchQoSAPI.md#getoswqosmode) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-mode | Get switch QoS mode.
[**GetQosRules**](SwitchQoSAPI.md#getqosrules) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule | Get switch QoS rules.
[**GetSchedulerMapping**](SwitchQoSAPI.md#getschedulermapping) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/scheduler-mapping | Get switch QoS queue scheduling.
[**GetSelectableDevices**](SwitchQoSAPI.md#getselectabledevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/selectable-devices | Get switch QoS selectable devices.
[**GetSelectedDeviceBriefInfo**](SwitchQoSAPI.md#getselecteddevicebriefinfo) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/device-info | Get switch QoS selected device brief info.
[**GetUpgradableDevices**](SwitchQoSAPI.md#getupgradabledevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/upgradable-devices | Get switch QoS upgradable devices.
[**ModifyDscpMapping**](SwitchQoSAPI.md#modifydscpmapping) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/dscp-mapping | Modify switch QoS dscp mapping.
[**ModifyOswQosMode**](SwitchQoSAPI.md#modifyoswqosmode) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-mode | Modify switch QoS mode.
[**ModifyQosRule**](SwitchQoSAPI.md#modifyqosrule) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/{qosRuleId} | Modify switch QoS rule.
[**ModifyQosRuleStatus**](SwitchQoSAPI.md#modifyqosrulestatus) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/status/{qosRuleId} | Modify switch QoS rule&#39;s status.
[**ModifySchedulerMapping**](SwitchQoSAPI.md#modifyschedulermapping) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switch-qos/qos-rule/scheduler-mapping | Modify switch QoS queue scheduling.



## CreateQosRule

> OperationResponse CreateQosRule(ctx, omadacId, siteId).OswQosRuleVO(oswQosRuleVO).Execute()

Create switch QoS rule.



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
	oswQosRuleVO := *openapiclient.NewOswQosRuleVO([]int32{int32(123)}, "Name_example", int32(123), false, int32(123)) // OswQosRuleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.CreateQosRule(context.Background(), omadacId, siteId).OswQosRuleVO(oswQosRuleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.CreateQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQosRule`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.CreateQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oswQosRuleVO** | [**OswQosRuleVO**](OswQosRuleVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQosRule

> OperationResponse DeleteQosRule(ctx, omadacId, siteId, qosRuleId).Execute()

Delete switch QoS rule.



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
	qosRuleId := "qosRuleId_example" // string | qosRuleId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.DeleteQosRule(context.Background(), omadacId, siteId, qosRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.DeleteQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteQosRule`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.DeleteQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**qosRuleId** | **string** | qosRuleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSelectableDevices

> OperationResponse GetAllSelectableDevices(ctx, omadacId, siteId).RuleType(ruleType).NetworkId(networkId).Execute()

Get switch QoS all selectable device's mac.



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
	ruleType := int32(56) // int32 | 
	networkId := "networkId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.GetAllSelectableDevices(context.Background(), omadacId, siteId).RuleType(ruleType).NetworkId(networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetAllSelectableDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSelectableDevices`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetAllSelectableDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSelectableDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleType** | **int32** |  | 
 **networkId** | **string** |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDscpMapping

> OperationResponse GetDscpMapping(ctx, omadacId, siteId).Execute()

Get switch QoS dscp mapping.



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
	resp, r, err := apiClient.SwitchQoSAPI.GetDscpMapping(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetDscpMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDscpMapping`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetDscpMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDscpMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDscpMappingTemplate

> OperationResponse GetDscpMappingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get switch QoS dscp mapping in site template.



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
	resp, r, err := apiClient.SwitchQoSAPI.GetDscpMappingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetDscpMappingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDscpMappingTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetDscpMappingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDscpMappingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswQosMode

> OperationResponse GetOswQosMode(ctx, omadacId, siteId).Execute()

Get switch QoS mode.



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
	resp, r, err := apiClient.SwitchQoSAPI.GetOswQosMode(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetOswQosMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswQosMode`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetOswQosMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswQosModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQosRules

> OperationResponse GetQosRules(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get switch QoS rules.



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.GetQosRules(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetQosRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQosRules`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetQosRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQosRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedulerMapping

> OperationResponse GetSchedulerMapping(ctx, omadacId, siteId).Execute()

Get switch QoS queue scheduling.



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
	resp, r, err := apiClient.SwitchQoSAPI.GetSchedulerMapping(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetSchedulerMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedulerMapping`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetSchedulerMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulerMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectableDevices

> OperationResponse GetSelectableDevices(ctx, omadacId, siteId).RuleType(ruleType).NetworkId(networkId).Execute()

Get switch QoS selectable devices.



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
	ruleType := int32(56) // int32 | 
	networkId := "networkId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.GetSelectableDevices(context.Background(), omadacId, siteId).RuleType(ruleType).NetworkId(networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetSelectableDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectableDevices`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetSelectableDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectableDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleType** | **int32** |  | 
 **networkId** | **string** |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectedDeviceBriefInfo

> OperationResponse GetSelectedDeviceBriefInfo(ctx, omadacId, siteId).OswMacListVO(oswMacListVO).Execute()

Get switch QoS selected device brief info.



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
	oswMacListVO := *openapiclient.NewOswMacListVO([]string{"SwitchMacList_example"}) // OswMacListVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.GetSelectedDeviceBriefInfo(context.Background(), omadacId, siteId).OswMacListVO(oswMacListVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetSelectedDeviceBriefInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectedDeviceBriefInfo`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetSelectedDeviceBriefInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectedDeviceBriefInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oswMacListVO** | [**OswMacListVO**](OswMacListVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpgradableDevices

> OperationResponse GetUpgradableDevices(ctx, omadacId, siteId).Execute()

Get switch QoS upgradable devices.



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
	resp, r, err := apiClient.SwitchQoSAPI.GetUpgradableDevices(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.GetUpgradableDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpgradableDevices`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.GetUpgradableDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpgradableDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDscpMapping

> OperationResponse ModifyDscpMapping(ctx, omadacId, siteId).DscpMappingVO(dscpMappingVO).Execute()

Modify switch QoS dscp mapping.



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
	dscpMappingVO := *openapiclient.NewDscpMappingVO(map[string][]int32{"key": []int32{int32(123)}}) // DscpMappingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.ModifyDscpMapping(context.Background(), omadacId, siteId).DscpMappingVO(dscpMappingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.ModifyDscpMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDscpMapping`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.ModifyDscpMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDscpMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dscpMappingVO** | [**DscpMappingVO**](DscpMappingVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOswQosMode

> OperationResponse ModifyOswQosMode(ctx, omadacId, siteId).OswQosModeVO(oswQosModeVO).Execute()

Modify switch QoS mode.



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
	oswQosModeVO := *openapiclient.NewOswQosModeVO() // OswQosModeVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.ModifyOswQosMode(context.Background(), omadacId, siteId).OswQosModeVO(oswQosModeVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.ModifyOswQosMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswQosMode`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.ModifyOswQosMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswQosModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oswQosModeVO** | [**OswQosModeVO**](OswQosModeVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyQosRule

> OperationResponse ModifyQosRule(ctx, omadacId, siteId, qosRuleId).OswQosRuleVO(oswQosRuleVO).Execute()

Modify switch QoS rule.



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
	qosRuleId := "qosRuleId_example" // string | qosRuleId
	oswQosRuleVO := *openapiclient.NewOswQosRuleVO([]int32{int32(123)}, "Name_example", int32(123), false, int32(123)) // OswQosRuleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.ModifyQosRule(context.Background(), omadacId, siteId, qosRuleId).OswQosRuleVO(oswQosRuleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.ModifyQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyQosRule`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.ModifyQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**qosRuleId** | **string** | qosRuleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswQosRuleVO** | [**OswQosRuleVO**](OswQosRuleVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyQosRuleStatus

> OperationResponse ModifyQosRuleStatus(ctx, omadacId, siteId, qosRuleId).OswQosRuleStatusVO(oswQosRuleStatusVO).Execute()

Modify switch QoS rule's status.



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
	qosRuleId := "qosRuleId_example" // string | qosRuleId
	oswQosRuleStatusVO := *openapiclient.NewOswQosRuleStatusVO(false) // OswQosRuleStatusVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.ModifyQosRuleStatus(context.Background(), omadacId, siteId, qosRuleId).OswQosRuleStatusVO(oswQosRuleStatusVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.ModifyQosRuleStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyQosRuleStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.ModifyQosRuleStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**qosRuleId** | **string** | qosRuleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyQosRuleStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswQosRuleStatusVO** | [**OswQosRuleStatusVO**](OswQosRuleStatusVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySchedulerMapping

> OperationResponse ModifySchedulerMapping(ctx, omadacId, siteId).QueueSchedulerMappingVO(queueSchedulerMappingVO).Execute()

Modify switch QoS queue scheduling.



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
	queueSchedulerMappingVO := *openapiclient.NewQueueSchedulerMappingVO() // QueueSchedulerMappingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchQoSAPI.ModifySchedulerMapping(context.Background(), omadacId, siteId).QueueSchedulerMappingVO(queueSchedulerMappingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchQoSAPI.ModifySchedulerMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySchedulerMapping`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SwitchQoSAPI.ModifySchedulerMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySchedulerMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queueSchedulerMappingVO** | [**QueueSchedulerMappingVO**](QueueSchedulerMappingVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


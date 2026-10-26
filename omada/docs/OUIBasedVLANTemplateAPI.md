# OUIBasedVLANTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApOuiBasedVlanTemplate**](OUIBasedVLANTemplateAPI.md#createapouibasedvlantemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ap-oui-rules | Create Ap oui based vlan template
[**CreateSwitchOuiBasedVlanTemplate**](OUIBasedVLANTemplateAPI.md#createswitchouibasedvlantemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules | Create Switch oui based vlan template
[**DeleteApOuiBasedVlanTemplate**](OUIBasedVLANTemplateAPI.md#deleteapouibasedvlantemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ap-oui-rules/{ouiBasedVlanId} | Delete Ap oui based vlan template
[**DeleteSwitchOuiBasedVlanTemplate**](OUIBasedVLANTemplateAPI.md#deleteswitchouibasedvlantemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/{ouiBasedVlanId} | Delete Switch oui based vlan template
[**EnableAutoAddVlanTemplate**](OUIBasedVLANTemplateAPI.md#enableautoaddvlantemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/enable-auto-add-vlan | Enable the site template and device auto add vlan.
[**GetAllSupportAutoAddVlanSwitchListTemplate**](OUIBasedVLANTemplateAPI.md#getallsupportautoaddvlanswitchlisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/switches/support-auto-add-vlan | Get switch template list of supporting oui based vlan with ports(lags) info of vlans in oui based rules but not included in port(lag) vlan.
[**GetGridApOuiBasedVlanTemplate**](OUIBasedVLANTemplateAPI.md#getgridapouibasedvlantemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ap-oui-rules | Get Ap oui based vlan template list
[**GetGridSwitchOuiBasedVlanTemplate**](OUIBasedVLANTemplateAPI.md#getgridswitchouibasedvlantemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules | Get Switch oui based vlan template list
[**GetOuiBasedVlanSwitchListTemplate**](OUIBasedVLANTemplateAPI.md#getouibasedvlanswitchlisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/switches | Get all oui based vlan switch template list.
[**GetWhetherAllSwitchSupportAutoAddVlanTemplate**](OUIBasedVLANTemplateAPI.md#getwhetherallswitchsupportautoaddvlantemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/auto-add-vlan-all-support | Get whether all switches template support oui rule auto add vlan.
[**GetWhetherExistAllowAllPortOrLagsForRuleCreationTemplate**](OUIBasedVLANTemplateAPI.md#getwhetherexistallowallportorlagsforrulecreationtemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/switches/check | Get whether exist at least one port or lag with allow all network tags for oui based vlan creation.
[**GetWhetherExistAllowAllPortOrLagsForRuleModificationTemplate**](OUIBasedVLANTemplateAPI.md#getwhetherexistallowallportorlagsforrulemodificationtemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/switches/{ouiBasedVlanId}/check | Get whether exist at least one port or lag with allow all network tags for oui based vlan modification.
[**GetWhetherSiteOpenAutoAddVlanTemplate**](OUIBasedVLANTemplateAPI.md#getwhethersiteopenautoaddvlantemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/whether-auto-add-vlan | Get whether site template open oui rule auto add vlan.
[**ModifyApOuiBasedVlanTemplate**](OUIBasedVLANTemplateAPI.md#modifyapouibasedvlantemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ap-oui-rules/{ouiBasedVlanId} | Modify Ap oui based vlan template
[**ModifySwitchOuiBasedVlanTemplate**](OUIBasedVLANTemplateAPI.md#modifyswitchouibasedvlantemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/switch-oui-rules/{ouiBasedVlanId} | Modify Switch oui based vlan template



## CreateApOuiBasedVlanTemplate

> OperationResponseResIdOpenApiVO CreateApOuiBasedVlanTemplate(ctx, omadacId, siteTemplateId).OuiBasedVlanApOpenApiVO(ouiBasedVlanApOpenApiVO).Execute()

Create Ap oui based vlan template



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ouiBasedVlanApOpenApiVO := *openapiclient.NewOuiBasedVlanApOpenApiVO(false, "Name_example", []openapiclient.VlanOuiModeOpenApiVO{*openapiclient.NewVlanOuiModeOpenApiVO("OuiProfileId_example", int32(123), int32(123))}, []string{"SsidIdList_example"}) // OuiBasedVlanApOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.CreateApOuiBasedVlanTemplate(context.Background(), omadacId, siteTemplateId).OuiBasedVlanApOpenApiVO(ouiBasedVlanApOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.CreateApOuiBasedVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApOuiBasedVlanTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.CreateApOuiBasedVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApOuiBasedVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ouiBasedVlanApOpenApiVO** | [**OuiBasedVlanApOpenApiVO**](OuiBasedVlanApOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSwitchOuiBasedVlanTemplate

> OperationResponseResIdOpenApiVO CreateSwitchOuiBasedVlanTemplate(ctx, omadacId, siteTemplateId).OuiBasedVlanTemplateOpenApiVO(ouiBasedVlanTemplateOpenApiVO).Execute()

Create Switch oui based vlan template



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ouiBasedVlanTemplateOpenApiVO := *openapiclient.NewOuiBasedVlanTemplateOpenApiVO(false, int32(123), "Name_example", []openapiclient.VlanOuiModeOpenApiVO{*openapiclient.NewVlanOuiModeOpenApiVO("OuiProfileId_example", int32(123), int32(123))}) // OuiBasedVlanTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.CreateSwitchOuiBasedVlanTemplate(context.Background(), omadacId, siteTemplateId).OuiBasedVlanTemplateOpenApiVO(ouiBasedVlanTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.CreateSwitchOuiBasedVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSwitchOuiBasedVlanTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.CreateSwitchOuiBasedVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSwitchOuiBasedVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ouiBasedVlanTemplateOpenApiVO** | [**OuiBasedVlanTemplateOpenApiVO**](OuiBasedVlanTemplateOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApOuiBasedVlanTemplate

> OperationResponseWithoutResult DeleteApOuiBasedVlanTemplate(ctx, omadacId, siteTemplateId, ouiBasedVlanId).Execute()

Delete Ap oui based vlan template



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | The id of oui profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.DeleteApOuiBasedVlanTemplate(context.Background(), omadacId, siteTemplateId, ouiBasedVlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.DeleteApOuiBasedVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApOuiBasedVlanTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.DeleteApOuiBasedVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | The id of oui profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApOuiBasedVlanTemplateRequest struct via the builder pattern


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


## DeleteSwitchOuiBasedVlanTemplate

> OperationResponseWithoutResult DeleteSwitchOuiBasedVlanTemplate(ctx, omadacId, siteTemplateId, ouiBasedVlanId).Execute()

Delete Switch oui based vlan template



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | Oui Based Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.DeleteSwitchOuiBasedVlanTemplate(context.Background(), omadacId, siteTemplateId, ouiBasedVlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.DeleteSwitchOuiBasedVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSwitchOuiBasedVlanTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.DeleteSwitchOuiBasedVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | Oui Based Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSwitchOuiBasedVlanTemplateRequest struct via the builder pattern


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


## EnableAutoAddVlanTemplate

> OperationResponse EnableAutoAddVlanTemplate(ctx, omadacId, siteTemplateId).Execute()

Enable the site template and device auto add vlan.



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.EnableAutoAddVlanTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.EnableAutoAddVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableAutoAddVlanTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.EnableAutoAddVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAutoAddVlanTemplateRequest struct via the builder pattern


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


## GetAllSupportAutoAddVlanSwitchListTemplate

> OperationResponseGridVOOuiBasedVlanDeviceInfoVO GetAllSupportAutoAddVlanSwitchListTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get switch template list of supporting oui based vlan with ports(lags) info of vlans in oui based rules but not included in port(lag) vlan.



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.GetAllSupportAutoAddVlanSwitchListTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.GetAllSupportAutoAddVlanSwitchListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSupportAutoAddVlanSwitchListTemplate`: OperationResponseGridVOOuiBasedVlanDeviceInfoVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.GetAllSupportAutoAddVlanSwitchListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSupportAutoAddVlanSwitchListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOOuiBasedVlanDeviceInfoVO**](OperationResponseGridVOOuiBasedVlanDeviceInfoVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridApOuiBasedVlanTemplate

> OperationResponseGridVOOuiBasedVlanApQueryOpenApiVO GetGridApOuiBasedVlanTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get Ap oui based vlan template list



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.GetGridApOuiBasedVlanTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.GetGridApOuiBasedVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridApOuiBasedVlanTemplate`: OperationResponseGridVOOuiBasedVlanApQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.GetGridApOuiBasedVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridApOuiBasedVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOuiBasedVlanApQueryOpenApiVO**](OperationResponseGridVOOuiBasedVlanApQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSwitchOuiBasedVlanTemplate

> OperationResponseGridVOOuiBasedVlanTemplateSwitchQueryOpenApiVO GetGridSwitchOuiBasedVlanTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get Switch oui based vlan template list



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.GetGridSwitchOuiBasedVlanTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.GetGridSwitchOuiBasedVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSwitchOuiBasedVlanTemplate`: OperationResponseGridVOOuiBasedVlanTemplateSwitchQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.GetGridSwitchOuiBasedVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSwitchOuiBasedVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOuiBasedVlanTemplateSwitchQueryOpenApiVO**](OperationResponseGridVOOuiBasedVlanTemplateSwitchQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOuiBasedVlanSwitchListTemplate

> OperationResponseOuiBasedVlanGridVOSwitchDataVO GetOuiBasedVlanSwitchListTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get all oui based vlan switch template list.



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.GetOuiBasedVlanSwitchListTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.GetOuiBasedVlanSwitchListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOuiBasedVlanSwitchListTemplate`: OperationResponseOuiBasedVlanGridVOSwitchDataVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.GetOuiBasedVlanSwitchListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOuiBasedVlanSwitchListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseOuiBasedVlanGridVOSwitchDataVO**](OperationResponseOuiBasedVlanGridVOSwitchDataVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhetherAllSwitchSupportAutoAddVlanTemplate

> OperationResponseOuiBasedVlanDeviceAutoAddVlanVO GetWhetherAllSwitchSupportAutoAddVlanTemplate(ctx, omadacId, siteTemplateId).Execute()

Get whether all switches template support oui rule auto add vlan.



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.GetWhetherAllSwitchSupportAutoAddVlanTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.GetWhetherAllSwitchSupportAutoAddVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhetherAllSwitchSupportAutoAddVlanTemplate`: OperationResponseOuiBasedVlanDeviceAutoAddVlanVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.GetWhetherAllSwitchSupportAutoAddVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhetherAllSwitchSupportAutoAddVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseOuiBasedVlanDeviceAutoAddVlanVO**](OperationResponseOuiBasedVlanDeviceAutoAddVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhetherExistAllowAllPortOrLagsForRuleCreationTemplate

> OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO GetWhetherExistAllowAllPortOrLagsForRuleCreationTemplate(ctx, omadacId, siteTemplateId).OuiBasedVlanRuleVO(ouiBasedVlanRuleVO).Execute()

Get whether exist at least one port or lag with allow all network tags for oui based vlan creation.



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ouiBasedVlanRuleVO := *openapiclient.NewOuiBasedVlanRuleVO(false, "Name_example", []openapiclient.VlanOuiModeVO{*openapiclient.NewVlanOuiModeVO("OuiProfileId_example", int32(123), int32(123))}, int32(123)) // OuiBasedVlanRuleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.GetWhetherExistAllowAllPortOrLagsForRuleCreationTemplate(context.Background(), omadacId, siteTemplateId).OuiBasedVlanRuleVO(ouiBasedVlanRuleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.GetWhetherExistAllowAllPortOrLagsForRuleCreationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhetherExistAllowAllPortOrLagsForRuleCreationTemplate`: OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.GetWhetherExistAllowAllPortOrLagsForRuleCreationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhetherExistAllowAllPortOrLagsForRuleCreationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ouiBasedVlanRuleVO** | [**OuiBasedVlanRuleVO**](OuiBasedVlanRuleVO.md) |  | 

### Return type

[**OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO**](OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhetherExistAllowAllPortOrLagsForRuleModificationTemplate

> OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO GetWhetherExistAllowAllPortOrLagsForRuleModificationTemplate(ctx, omadacId, siteTemplateId, ouiBasedVlanId).OuiBasedVlanRuleVO(ouiBasedVlanRuleVO).Execute()

Get whether exist at least one port or lag with allow all network tags for oui based vlan modification.



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | Oui Based Vlan ID
	ouiBasedVlanRuleVO := *openapiclient.NewOuiBasedVlanRuleVO(false, "Name_example", []openapiclient.VlanOuiModeVO{*openapiclient.NewVlanOuiModeVO("OuiProfileId_example", int32(123), int32(123))}, int32(123)) // OuiBasedVlanRuleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.GetWhetherExistAllowAllPortOrLagsForRuleModificationTemplate(context.Background(), omadacId, siteTemplateId, ouiBasedVlanId).OuiBasedVlanRuleVO(ouiBasedVlanRuleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.GetWhetherExistAllowAllPortOrLagsForRuleModificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhetherExistAllowAllPortOrLagsForRuleModificationTemplate`: OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.GetWhetherExistAllowAllPortOrLagsForRuleModificationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | Oui Based Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhetherExistAllowAllPortOrLagsForRuleModificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ouiBasedVlanRuleVO** | [**OuiBasedVlanRuleVO**](OuiBasedVlanRuleVO.md) |  | 

### Return type

[**OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO**](OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhetherSiteOpenAutoAddVlanTemplate

> OperationResponseOuiBasedVlanSiteAutoAddVlanVO GetWhetherSiteOpenAutoAddVlanTemplate(ctx, omadacId, siteTemplateId).Execute()

Get whether site template open oui rule auto add vlan.



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.GetWhetherSiteOpenAutoAddVlanTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.GetWhetherSiteOpenAutoAddVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhetherSiteOpenAutoAddVlanTemplate`: OperationResponseOuiBasedVlanSiteAutoAddVlanVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.GetWhetherSiteOpenAutoAddVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhetherSiteOpenAutoAddVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseOuiBasedVlanSiteAutoAddVlanVO**](OperationResponseOuiBasedVlanSiteAutoAddVlanVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyApOuiBasedVlanTemplate

> OperationResponseWithoutResult ModifyApOuiBasedVlanTemplate(ctx, omadacId, siteTemplateId, ouiBasedVlanId).OuiBasedVlanApOpenApiVO(ouiBasedVlanApOpenApiVO).Execute()

Modify Ap oui based vlan template



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | The id of oui profile.
	ouiBasedVlanApOpenApiVO := *openapiclient.NewOuiBasedVlanApOpenApiVO(false, "Name_example", []openapiclient.VlanOuiModeOpenApiVO{*openapiclient.NewVlanOuiModeOpenApiVO("OuiProfileId_example", int32(123), int32(123))}, []string{"SsidIdList_example"}) // OuiBasedVlanApOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.ModifyApOuiBasedVlanTemplate(context.Background(), omadacId, siteTemplateId, ouiBasedVlanId).OuiBasedVlanApOpenApiVO(ouiBasedVlanApOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.ModifyApOuiBasedVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApOuiBasedVlanTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.ModifyApOuiBasedVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | The id of oui profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyApOuiBasedVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ouiBasedVlanApOpenApiVO** | [**OuiBasedVlanApOpenApiVO**](OuiBasedVlanApOpenApiVO.md) |  | 

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


## ModifySwitchOuiBasedVlanTemplate

> OperationResponseWithoutResult ModifySwitchOuiBasedVlanTemplate(ctx, omadacId, siteTemplateId, ouiBasedVlanId).OuiBasedVlanTemplateOpenApiVO(ouiBasedVlanTemplateOpenApiVO).Execute()

Modify Switch oui based vlan template



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | Oui Based Vlan ID
	ouiBasedVlanTemplateOpenApiVO := *openapiclient.NewOuiBasedVlanTemplateOpenApiVO(false, int32(123), "Name_example", []openapiclient.VlanOuiModeOpenApiVO{*openapiclient.NewVlanOuiModeOpenApiVO("OuiProfileId_example", int32(123), int32(123))}) // OuiBasedVlanTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANTemplateAPI.ModifySwitchOuiBasedVlanTemplate(context.Background(), omadacId, siteTemplateId, ouiBasedVlanId).OuiBasedVlanTemplateOpenApiVO(ouiBasedVlanTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANTemplateAPI.ModifySwitchOuiBasedVlanTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchOuiBasedVlanTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANTemplateAPI.ModifySwitchOuiBasedVlanTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | Oui Based Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchOuiBasedVlanTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ouiBasedVlanTemplateOpenApiVO** | [**OuiBasedVlanTemplateOpenApiVO**](OuiBasedVlanTemplateOpenApiVO.md) |  | 

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


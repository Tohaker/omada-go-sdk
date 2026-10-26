# OUIBasedVLANAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApOuiBasedVlan**](OUIBasedVLANAPI.md#createapouibasedvlan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ap-oui-rules | Create Ap oui based vlan
[**CreateSwitchOuiBasedVlan**](OUIBasedVLANAPI.md#createswitchouibasedvlan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules | Create Switch oui based vlan
[**DeleteApOuiBasedVlan**](OUIBasedVLANAPI.md#deleteapouibasedvlan) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/ap-oui-rules/{ouiBasedVlanId} | Delete Ap oui based vlan
[**DeleteSwitchOuiBasedVlan**](OUIBasedVLANAPI.md#deleteswitchouibasedvlan) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/{ouiBasedVlanId} | Delete Switch oui based vlan
[**EnableAutoAddVlan**](OUIBasedVLANAPI.md#enableautoaddvlan) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/enable-auto-add-vlan | Enable the site and device auto add vlan.
[**GetAllSupportAutoAddVlanSwitchList**](OUIBasedVLANAPI.md#getallsupportautoaddvlanswitchlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/switches/support-auto-add-vlan | Get switch list of supporting oui based vlan with ports(lags) info of vlans in oui based rules but not included in port(lag) vlan.
[**GetGridApOuiBasedVlan**](OUIBasedVLANAPI.md#getgridapouibasedvlan) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ap-oui-rules | Get Ap oui based vlan list
[**GetGridSwitchOuiBasedVlan**](OUIBasedVLANAPI.md#getgridswitchouibasedvlan) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules | Get Switch oui based vlan list
[**GetOuiBasedVlanSwitchList**](OUIBasedVLANAPI.md#getouibasedvlanswitchlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/switches | Get all oui based vlan switch list.
[**GetSwitchList**](OUIBasedVLANAPI.md#getswitchlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/support-device | Get switch list of supporting oui based vlan
[**GetWhetherAllSwitchSupportAutoAddVlan**](OUIBasedVLANAPI.md#getwhetherallswitchsupportautoaddvlan) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/auto-add-vlan-all-support | Get whether all switches support oui rule auto add vlan.
[**GetWhetherExistAllowAllPortOrLagsForRuleCreation**](OUIBasedVLANAPI.md#getwhetherexistallowallportorlagsforrulecreation) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/switches/check | Get whether exist at least one port or lag with allow all network tags for oui based vlan creation.
[**GetWhetherExistAllowAllPortOrLagsForRuleModification**](OUIBasedVLANAPI.md#getwhetherexistallowallportorlagsforrulemodification) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/switches/{ouiBasedVlanId}/check | Get whether exist at least one port or lag with allow all network tags for oui based vlan modification.
[**GetWhetherSiteOpenAutoAddVlan**](OUIBasedVLANAPI.md#getwhethersiteopenautoaddvlan) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/whether-auto-add-vlan | Get whether site open oui rule auto add vlan.
[**ModifyApOuiBasedVlan**](OUIBasedVLANAPI.md#modifyapouibasedvlan) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/ap-oui-rules/{ouiBasedVlanId} | Modify Ap oui based vlan
[**ModifySwitchOuiBasedVlan**](OUIBasedVLANAPI.md#modifyswitchouibasedvlan) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/switch-oui-rules/{ouiBasedVlanId} | Modify Switch oui based vlan



## CreateApOuiBasedVlan

> OperationResponseResIdOpenApiVO CreateApOuiBasedVlan(ctx, omadacId, siteId).OuiBasedVlanApOpenApiVO(ouiBasedVlanApOpenApiVO).Execute()

Create Ap oui based vlan



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
	ouiBasedVlanApOpenApiVO := *openapiclient.NewOuiBasedVlanApOpenApiVO(false, "Name_example", []openapiclient.VlanOuiModeOpenApiVO{*openapiclient.NewVlanOuiModeOpenApiVO("OuiProfileId_example", int32(123), int32(123))}, []string{"SsidIdList_example"}) // OuiBasedVlanApOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.CreateApOuiBasedVlan(context.Background(), omadacId, siteId).OuiBasedVlanApOpenApiVO(ouiBasedVlanApOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.CreateApOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApOuiBasedVlan`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.CreateApOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApOuiBasedVlanRequest struct via the builder pattern


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


## CreateSwitchOuiBasedVlan

> OperationResponseResIdOpenApiVO CreateSwitchOuiBasedVlan(ctx, omadacId, siteId).OuiBasedVlanSwitchOpenApiVO(ouiBasedVlanSwitchOpenApiVO).Execute()

Create Switch oui based vlan



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
	ouiBasedVlanSwitchOpenApiVO := *openapiclient.NewOuiBasedVlanSwitchOpenApiVO(false, int32(123), "Name_example", []openapiclient.VlanOuiModeOpenApiVO{*openapiclient.NewVlanOuiModeOpenApiVO("OuiProfileId_example", int32(123), int32(123))}) // OuiBasedVlanSwitchOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.CreateSwitchOuiBasedVlan(context.Background(), omadacId, siteId).OuiBasedVlanSwitchOpenApiVO(ouiBasedVlanSwitchOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.CreateSwitchOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSwitchOuiBasedVlan`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.CreateSwitchOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSwitchOuiBasedVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ouiBasedVlanSwitchOpenApiVO** | [**OuiBasedVlanSwitchOpenApiVO**](OuiBasedVlanSwitchOpenApiVO.md) |  | 

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


## DeleteApOuiBasedVlan

> OperationResponseWithoutResult DeleteApOuiBasedVlan(ctx, omadacId, siteId, ouiBasedVlanId).Execute()

Delete Ap oui based vlan



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
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | Oui Based Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.DeleteApOuiBasedVlan(context.Background(), omadacId, siteId, ouiBasedVlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.DeleteApOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApOuiBasedVlan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.DeleteApOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | Oui Based Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApOuiBasedVlanRequest struct via the builder pattern


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


## DeleteSwitchOuiBasedVlan

> OperationResponseWithoutResult DeleteSwitchOuiBasedVlan(ctx, omadacId, siteId, ouiBasedVlanId).Execute()

Delete Switch oui based vlan



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
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | Oui Based Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.DeleteSwitchOuiBasedVlan(context.Background(), omadacId, siteId, ouiBasedVlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.DeleteSwitchOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSwitchOuiBasedVlan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.DeleteSwitchOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | Oui Based Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSwitchOuiBasedVlanRequest struct via the builder pattern


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


## EnableAutoAddVlan

> OperationResponse EnableAutoAddVlan(ctx, omadacId, siteId).Execute()

Enable the site and device auto add vlan.



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
	resp, r, err := apiClient.OUIBasedVLANAPI.EnableAutoAddVlan(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.EnableAutoAddVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableAutoAddVlan`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.EnableAutoAddVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableAutoAddVlanRequest struct via the builder pattern


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


## GetAllSupportAutoAddVlanSwitchList

> OperationResponseGridVOOuiBasedVlanDeviceInfoVO GetAllSupportAutoAddVlanSwitchList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Name(name).Execute()

Get switch list of supporting oui based vlan with ports(lags) info of vlans in oui based rules but not included in port(lag) vlan.



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
	name := "name_example" // string | Switch name. Support fuzzy search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.GetAllSupportAutoAddVlanSwitchList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetAllSupportAutoAddVlanSwitchList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSupportAutoAddVlanSwitchList`: OperationResponseGridVOOuiBasedVlanDeviceInfoVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetAllSupportAutoAddVlanSwitchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSupportAutoAddVlanSwitchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 


 **name** | **string** | Switch name. Support fuzzy search | 

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


## GetGridApOuiBasedVlan

> OperationResponseGridVOOuiBasedVlanApQueryOpenApiVO GetGridApOuiBasedVlan(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Ap oui based vlan list



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
	resp, r, err := apiClient.OUIBasedVLANAPI.GetGridApOuiBasedVlan(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetGridApOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridApOuiBasedVlan`: OperationResponseGridVOOuiBasedVlanApQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetGridApOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridApOuiBasedVlanRequest struct via the builder pattern


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


## GetGridSwitchOuiBasedVlan

> OperationResponseGridVOOuiBasedVlanSwitchQueryOpenApiVO GetGridSwitchOuiBasedVlan(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Switch oui based vlan list



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
	resp, r, err := apiClient.OUIBasedVLANAPI.GetGridSwitchOuiBasedVlan(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetGridSwitchOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSwitchOuiBasedVlan`: OperationResponseGridVOOuiBasedVlanSwitchQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetGridSwitchOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSwitchOuiBasedVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOuiBasedVlanSwitchQueryOpenApiVO**](OperationResponseGridVOOuiBasedVlanSwitchQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOuiBasedVlanSwitchList

> OperationResponseOuiBasedVlanGridVOSwitchDataVO GetOuiBasedVlanSwitchList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get all oui based vlan switch list.



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
	resp, r, err := apiClient.OUIBasedVLANAPI.GetOuiBasedVlanSwitchList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetOuiBasedVlanSwitchList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOuiBasedVlanSwitchList`: OperationResponseOuiBasedVlanGridVOSwitchDataVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetOuiBasedVlanSwitchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOuiBasedVlanSwitchListRequest struct via the builder pattern


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


## GetSwitchList

> OperationResponseGridVOOuiBasedVlanDeviceInfoVO GetSwitchList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Name(name).Execute()

Get switch list of supporting oui based vlan



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
	name := "name_example" // string | Switch name. Support fuzzy search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.GetSwitchList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetSwitchList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchList`: OperationResponseGridVOOuiBasedVlanDeviceInfoVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetSwitchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 


 **name** | **string** | Switch name. Support fuzzy search | 

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


## GetWhetherAllSwitchSupportAutoAddVlan

> OperationResponseOuiBasedVlanDeviceAutoAddVlanVO GetWhetherAllSwitchSupportAutoAddVlan(ctx, omadacId, siteId).Execute()

Get whether all switches support oui rule auto add vlan.



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
	resp, r, err := apiClient.OUIBasedVLANAPI.GetWhetherAllSwitchSupportAutoAddVlan(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetWhetherAllSwitchSupportAutoAddVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhetherAllSwitchSupportAutoAddVlan`: OperationResponseOuiBasedVlanDeviceAutoAddVlanVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetWhetherAllSwitchSupportAutoAddVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhetherAllSwitchSupportAutoAddVlanRequest struct via the builder pattern


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


## GetWhetherExistAllowAllPortOrLagsForRuleCreation

> OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO GetWhetherExistAllowAllPortOrLagsForRuleCreation(ctx, omadacId, siteId).OuiBasedVlanRuleVO(ouiBasedVlanRuleVO).Execute()

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
	siteId := "siteId_example" // string | Site ID
	ouiBasedVlanRuleVO := *openapiclient.NewOuiBasedVlanRuleVO(false, "Name_example", []openapiclient.VlanOuiModeVO{*openapiclient.NewVlanOuiModeVO("OuiProfileId_example", int32(123), int32(123))}, int32(123)) // OuiBasedVlanRuleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.GetWhetherExistAllowAllPortOrLagsForRuleCreation(context.Background(), omadacId, siteId).OuiBasedVlanRuleVO(ouiBasedVlanRuleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetWhetherExistAllowAllPortOrLagsForRuleCreation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhetherExistAllowAllPortOrLagsForRuleCreation`: OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetWhetherExistAllowAllPortOrLagsForRuleCreation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhetherExistAllowAllPortOrLagsForRuleCreationRequest struct via the builder pattern


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


## GetWhetherExistAllowAllPortOrLagsForRuleModification

> OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO GetWhetherExistAllowAllPortOrLagsForRuleModification(ctx, omadacId, siteId, ouiBasedVlanId).OuiBasedVlanRuleVO(ouiBasedVlanRuleVO).Execute()

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
	siteId := "siteId_example" // string | Site ID
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | Oui Based Vlan ID
	ouiBasedVlanRuleVO := *openapiclient.NewOuiBasedVlanRuleVO(false, "Name_example", []openapiclient.VlanOuiModeVO{*openapiclient.NewVlanOuiModeVO("OuiProfileId_example", int32(123), int32(123))}, int32(123)) // OuiBasedVlanRuleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.GetWhetherExistAllowAllPortOrLagsForRuleModification(context.Background(), omadacId, siteId, ouiBasedVlanId).OuiBasedVlanRuleVO(ouiBasedVlanRuleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetWhetherExistAllowAllPortOrLagsForRuleModification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhetherExistAllowAllPortOrLagsForRuleModification`: OperationResponseOuiBasedVlanOswPortNetworkTagsSettingVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetWhetherExistAllowAllPortOrLagsForRuleModification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | Oui Based Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhetherExistAllowAllPortOrLagsForRuleModificationRequest struct via the builder pattern


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


## GetWhetherSiteOpenAutoAddVlan

> OperationResponseOuiBasedVlanSiteAutoAddVlanVO GetWhetherSiteOpenAutoAddVlan(ctx, omadacId, siteId).Execute()

Get whether site open oui rule auto add vlan.



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
	resp, r, err := apiClient.OUIBasedVLANAPI.GetWhetherSiteOpenAutoAddVlan(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.GetWhetherSiteOpenAutoAddVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhetherSiteOpenAutoAddVlan`: OperationResponseOuiBasedVlanSiteAutoAddVlanVO
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.GetWhetherSiteOpenAutoAddVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhetherSiteOpenAutoAddVlanRequest struct via the builder pattern


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


## ModifyApOuiBasedVlan

> OperationResponseWithoutResult ModifyApOuiBasedVlan(ctx, omadacId, siteId, ouiBasedVlanId).OuiBasedVlanApOpenApiVO(ouiBasedVlanApOpenApiVO).Execute()

Modify Ap oui based vlan



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
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | Oui Based Vlan ID
	ouiBasedVlanApOpenApiVO := *openapiclient.NewOuiBasedVlanApOpenApiVO(false, "Name_example", []openapiclient.VlanOuiModeOpenApiVO{*openapiclient.NewVlanOuiModeOpenApiVO("OuiProfileId_example", int32(123), int32(123))}, []string{"SsidIdList_example"}) // OuiBasedVlanApOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.ModifyApOuiBasedVlan(context.Background(), omadacId, siteId, ouiBasedVlanId).OuiBasedVlanApOpenApiVO(ouiBasedVlanApOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.ModifyApOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApOuiBasedVlan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.ModifyApOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | Oui Based Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyApOuiBasedVlanRequest struct via the builder pattern


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


## ModifySwitchOuiBasedVlan

> OperationResponseWithoutResult ModifySwitchOuiBasedVlan(ctx, omadacId, siteId, ouiBasedVlanId).OuiBasedVlanSwitchOpenApiVO(ouiBasedVlanSwitchOpenApiVO).Execute()

Modify Switch oui based vlan



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
	ouiBasedVlanId := "ouiBasedVlanId_example" // string | Oui Based Vlan ID
	ouiBasedVlanSwitchOpenApiVO := *openapiclient.NewOuiBasedVlanSwitchOpenApiVO(false, int32(123), "Name_example", []openapiclient.VlanOuiModeOpenApiVO{*openapiclient.NewVlanOuiModeOpenApiVO("OuiProfileId_example", int32(123), int32(123))}) // OuiBasedVlanSwitchOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OUIBasedVLANAPI.ModifySwitchOuiBasedVlan(context.Background(), omadacId, siteId, ouiBasedVlanId).OuiBasedVlanSwitchOpenApiVO(ouiBasedVlanSwitchOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OUIBasedVLANAPI.ModifySwitchOuiBasedVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySwitchOuiBasedVlan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OUIBasedVLANAPI.ModifySwitchOuiBasedVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ouiBasedVlanId** | **string** | Oui Based Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySwitchOuiBasedVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ouiBasedVlanSwitchOpenApiVO** | [**OuiBasedVlanSwitchOpenApiVO**](OuiBasedVlanSwitchOpenApiVO.md) |  | 

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


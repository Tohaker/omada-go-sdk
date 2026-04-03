# SiteTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchSyncSiteWithTemplate**](SiteTemplateAPI.md#batchsyncsitewithtemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/multi-sites/sync | Batch sync site settings with the site template
[**BindSiteTemplate**](SiteTemplateAPI.md#bindsitetemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/bind-site | Bind site with a site template
[**BoundSiteExistUnSupportRadSecDevice**](SiteTemplateAPI.md#boundsiteexistunsupportradsecdevice) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devices/un-support-radsec | Check for unsupported RadSec devices in the bound sites.
[**CopySiteTemplate**](SiteTemplateAPI.md#copysitetemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/copy | Create Site Template by copying exist site template
[**CopySiteTemplateToCustomers**](SiteTemplateAPI.md#copysitetemplatetocustomers) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/copy/batch | Copy site template to customers
[**CreateDeviceTemplate**](SiteTemplateAPI.md#createdevicetemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates | create device template
[**CreateNewSiteTemplate**](SiteTemplateAPI.md#createnewsitetemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates | Create new site template
[**DeleteDeviceTemplate**](SiteTemplateAPI.md#deletedevicetemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/{deviceTemplateId} | delete device template
[**DeleteSiteTemplate**](SiteTemplateAPI.md#deletesitetemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId} | Delete an existing site template
[**DeviceTemplateBindDeviceBatch**](SiteTemplateAPI.md#devicetemplatebinddevicebatch) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/{deviceTemplateId}/batch-binddevice | Device template batch bind devices
[**DeviceTemplateSyncConfigBatch**](SiteTemplateAPI.md#devicetemplatesyncconfigbatch) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/{deviceTemplateId}/batch-sync-config | Device template batch sync config to devices
[**DeviceTemplateUnbindDeviceBatch**](SiteTemplateAPI.md#devicetemplateunbinddevicebatch) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/{deviceTemplateId}/batch-unbinddevice | Device template batch unbind devices
[**GetBoundSites**](SiteTemplateAPI.md#getboundsites) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/binding-sites | Get sites info which bound to the site template
[**GetGridDeviceTemplate**](SiteTemplateAPI.md#getgriddevicetemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates | Get deviceTemplate info list
[**GetGridDeviceTemplateAvailableDevices**](SiteTemplateAPI.md#getgriddevicetemplateavailabledevices) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/{deviceTemplateId}/available-bind-devices | Get deviceTemplate&#39;s available devices
[**GetGridDeviceTemplateBoundDevices**](SiteTemplateAPI.md#getgriddevicetemplatebounddevices) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/{deviceTemplateId}/bound-devices | Get deviceTemplate&#39;s bound devices
[**GetGridDeviceTemplateDetail**](SiteTemplateAPI.md#getgriddevicetemplatedetail) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/detail | get device template
[**GetGridDevicesAvailableTemplate**](SiteTemplateAPI.md#getgriddevicesavailabletemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/available-devices/{siteId} | get the devices on the site and the templates that can be bound
[**GetSiteTemplateAllModules**](SiteTemplateAPI.md#getsitetemplateallmodules) | **Get** /openapi/v1/{omadacId}/sitetemplates/all-modules | Get all function modules which can be selected when creating site template
[**GetSiteTemplateConfiguration**](SiteTemplateAPI.md#getsitetemplateconfiguration) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/configuration | get site template settings
[**GetSiteTemplateDstInfo**](SiteTemplateAPI.md#getsitetemplatedstinfo) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/dst-info | Get daylight saving time information for the Site template configuration
[**GetSiteTemplateEntity**](SiteTemplateAPI.md#getsitetemplateentity) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId} | Get site template info
[**GetSiteTemplateGeneralConfig**](SiteTemplateAPI.md#getsitetemplategeneralconfig) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/general/config | get site template general config
[**GetSiteTemplateList**](SiteTemplateAPI.md#getsitetemplatelist) | **Get** /openapi/v1/{omadacId}/sitetemplates | Get site template list
[**GetSiteTemplateModules**](SiteTemplateAPI.md#getsitetemplatemodules) | **Get** /openapi/v1/{omadacId}/sitetemplates/modules | Get function modules which can be selected when creating site template
[**GetSiteTemplateService**](SiteTemplateAPI.md#getsitetemplateservice) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service | get site template service
[**GetSiteTemplateWirelessFeature**](SiteTemplateAPI.md#getsitetemplatewirelessfeature) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/wireless/feature | get site template WirelessFeature
[**GetSitesOverrides**](SiteTemplateAPI.md#getsitesoverrides) | **Get** /openapi/v1/{omadacId}/site-overrides | get site overrides
[**GetValidDeviceModel**](SiteTemplateAPI.md#getvaliddevicemodel) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/available-devicemodels | get valid device template
[**ModifyDeviceTemplate**](SiteTemplateAPI.md#modifydevicetemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/devicetemplates/{deviceTemplateId} | modify device template
[**UnbindSite**](SiteTemplateAPI.md#unbindsite) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/batch-unbind | Batch unbind sites from a site template
[**UpdateSiteTemplateEntity**](SiteTemplateAPI.md#updatesitetemplateentity) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId} | Modify an existing site template



## BatchSyncSiteWithTemplate

> OperationResponseWithoutResult BatchSyncSiteWithTemplate(ctx, omadacId, siteTemplateId).BatchSyncSitesOpenApiVO(batchSyncSitesOpenApiVO).Execute()

Batch sync site settings with the site template



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
	batchSyncSitesOpenApiVO := *openapiclient.NewBatchSyncSitesOpenApiVO([]string{"SiteIds_example"}) // BatchSyncSitesOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.BatchSyncSiteWithTemplate(context.Background(), omadacId, siteTemplateId).BatchSyncSitesOpenApiVO(batchSyncSitesOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.BatchSyncSiteWithTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSyncSiteWithTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.BatchSyncSiteWithTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchSyncSiteWithTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchSyncSitesOpenApiVO** | [**BatchSyncSitesOpenApiVO**](BatchSyncSitesOpenApiVO.md) |  | 

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


## BindSiteTemplate

> OperationResponseWithoutResult BindSiteTemplate(ctx, omadacId, siteTemplateId).BindSiteOpenApiVO(bindSiteOpenApiVO).Execute()

Bind site with a site template



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
	bindSiteOpenApiVO := *openapiclient.NewBindSiteOpenApiVO("SiteId_example") // BindSiteOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.BindSiteTemplate(context.Background(), omadacId, siteTemplateId).BindSiteOpenApiVO(bindSiteOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.BindSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindSiteTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.BindSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindSiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bindSiteOpenApiVO** | [**BindSiteOpenApiVO**](BindSiteOpenApiVO.md) |  | 

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


## BoundSiteExistUnSupportRadSecDevice

> OperationResponseUnSupportRadSecOpenApiVO BoundSiteExistUnSupportRadSecDevice(ctx, omadacId, siteTemplateId).Execute()

Check for unsupported RadSec devices in the bound sites.



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
	resp, r, err := apiClient.SiteTemplateAPI.BoundSiteExistUnSupportRadSecDevice(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.BoundSiteExistUnSupportRadSecDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoundSiteExistUnSupportRadSecDevice`: OperationResponseUnSupportRadSecOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.BoundSiteExistUnSupportRadSecDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoundSiteExistUnSupportRadSecDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUnSupportRadSecOpenApiVO**](OperationResponseUnSupportRadSecOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopySiteTemplate

> CopySiteResultVO CopySiteTemplate(ctx, omadacId, siteTemplateId).CopySiteTemplateOpenApiVO(copySiteTemplateOpenApiVO).Execute()

Create Site Template by copying exist site template



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
	copySiteTemplateOpenApiVO := *openapiclient.NewCopySiteTemplateOpenApiVO("Name_example") // CopySiteTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.CopySiteTemplate(context.Background(), omadacId, siteTemplateId).CopySiteTemplateOpenApiVO(copySiteTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.CopySiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopySiteTemplate`: CopySiteResultVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.CopySiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopySiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **copySiteTemplateOpenApiVO** | [**CopySiteTemplateOpenApiVO**](CopySiteTemplateOpenApiVO.md) |  | 

### Return type

[**CopySiteResultVO**](CopySiteResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopySiteTemplateToCustomers

> OperationResponseCopySiteTemplateOpenApiResultVO CopySiteTemplateToCustomers(ctx, omadacId, siteTemplateId).BatchCopySiteTemplateOpenApiVO(batchCopySiteTemplateOpenApiVO).Execute()

Copy site template to customers



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
	batchCopySiteTemplateOpenApiVO := *openapiclient.NewBatchCopySiteTemplateOpenApiVO("Name_example", []string{"TargetOmadacs_example"}) // BatchCopySiteTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.CopySiteTemplateToCustomers(context.Background(), omadacId, siteTemplateId).BatchCopySiteTemplateOpenApiVO(batchCopySiteTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.CopySiteTemplateToCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopySiteTemplateToCustomers`: OperationResponseCopySiteTemplateOpenApiResultVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.CopySiteTemplateToCustomers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopySiteTemplateToCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchCopySiteTemplateOpenApiVO** | [**BatchCopySiteTemplateOpenApiVO**](BatchCopySiteTemplateOpenApiVO.md) |  | 

### Return type

[**OperationResponseCopySiteTemplateOpenApiResultVO**](OperationResponseCopySiteTemplateOpenApiResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceTemplate

> OperationResponseWithoutResult CreateDeviceTemplate(ctx, omadacId, siteTemplateId).DeviceTemplateAdd(deviceTemplateAdd).Execute()

create device template



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
	deviceTemplateAdd := *openapiclient.NewDeviceTemplateAdd() // DeviceTemplateAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.CreateDeviceTemplate(context.Background(), omadacId, siteTemplateId).DeviceTemplateAdd(deviceTemplateAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.CreateDeviceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.CreateDeviceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceTemplateAdd** | [**DeviceTemplateAdd**](DeviceTemplateAdd.md) |  | 

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


## CreateNewSiteTemplate

> OperationResponse CreateNewSiteTemplate(ctx, omadacId).CreateSiteTemplateEntity(createSiteTemplateEntity).Execute()

Create new site template



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
	createSiteTemplateEntity := *openapiclient.NewCreateSiteTemplateEntity("Name_example") // CreateSiteTemplateEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.CreateNewSiteTemplate(context.Background(), omadacId).CreateSiteTemplateEntity(createSiteTemplateEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.CreateNewSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewSiteTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.CreateNewSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSiteTemplateEntity** | [**CreateSiteTemplateEntity**](CreateSiteTemplateEntity.md) |  | 

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


## DeleteDeviceTemplate

> OperationResponseWithoutResult DeleteDeviceTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).Execute()

delete device template



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.DeleteDeviceTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.DeleteDeviceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeviceTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.DeleteDeviceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceTemplateRequest struct via the builder pattern


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


## DeleteSiteTemplate

> OperationResponseWithoutResult DeleteSiteTemplate(ctx, omadacId, siteTemplateId).Execute()

Delete an existing site template



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
	resp, r, err := apiClient.SiteTemplateAPI.DeleteSiteTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.DeleteSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSiteTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.DeleteSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteTemplateRequest struct via the builder pattern


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


## DeviceTemplateBindDeviceBatch

> OperationResponseBatchBindDeviceResultOpenApiVO DeviceTemplateBindDeviceBatch(ctx, omadacId, siteTemplateId, deviceTemplateId).BatchBindDeviceOpenApiVO(batchBindDeviceOpenApiVO).Execute()

Device template batch bind devices



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	batchBindDeviceOpenApiVO := *openapiclient.NewBatchBindDeviceOpenApiVO() // BatchBindDeviceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.DeviceTemplateBindDeviceBatch(context.Background(), omadacId, siteTemplateId, deviceTemplateId).BatchBindDeviceOpenApiVO(batchBindDeviceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.DeviceTemplateBindDeviceBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceTemplateBindDeviceBatch`: OperationResponseBatchBindDeviceResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.DeviceTemplateBindDeviceBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceTemplateBindDeviceBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchBindDeviceOpenApiVO** | [**BatchBindDeviceOpenApiVO**](BatchBindDeviceOpenApiVO.md) |  | 

### Return type

[**OperationResponseBatchBindDeviceResultOpenApiVO**](OperationResponseBatchBindDeviceResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceTemplateSyncConfigBatch

> BatchBindDeviceResultOpenApiVO DeviceTemplateSyncConfigBatch(ctx, omadacId, siteTemplateId, deviceTemplateId).BatchBindDeviceOpenApiVO(batchBindDeviceOpenApiVO).Execute()

Device template batch sync config to devices



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	batchBindDeviceOpenApiVO := *openapiclient.NewBatchBindDeviceOpenApiVO() // BatchBindDeviceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.DeviceTemplateSyncConfigBatch(context.Background(), omadacId, siteTemplateId, deviceTemplateId).BatchBindDeviceOpenApiVO(batchBindDeviceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.DeviceTemplateSyncConfigBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceTemplateSyncConfigBatch`: BatchBindDeviceResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.DeviceTemplateSyncConfigBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceTemplateSyncConfigBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchBindDeviceOpenApiVO** | [**BatchBindDeviceOpenApiVO**](BatchBindDeviceOpenApiVO.md) |  | 

### Return type

[**BatchBindDeviceResultOpenApiVO**](BatchBindDeviceResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceTemplateUnbindDeviceBatch

> BatchBindDeviceResultOpenApiVO DeviceTemplateUnbindDeviceBatch(ctx, omadacId, siteTemplateId, deviceTemplateId).BatchBindDeviceOpenApiVO(batchBindDeviceOpenApiVO).Execute()

Device template batch unbind devices



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	batchBindDeviceOpenApiVO := *openapiclient.NewBatchBindDeviceOpenApiVO() // BatchBindDeviceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.DeviceTemplateUnbindDeviceBatch(context.Background(), omadacId, siteTemplateId, deviceTemplateId).BatchBindDeviceOpenApiVO(batchBindDeviceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.DeviceTemplateUnbindDeviceBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceTemplateUnbindDeviceBatch`: BatchBindDeviceResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.DeviceTemplateUnbindDeviceBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceTemplateUnbindDeviceBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchBindDeviceOpenApiVO** | [**BatchBindDeviceOpenApiVO**](BatchBindDeviceOpenApiVO.md) |  | 

### Return type

[**BatchBindDeviceResultOpenApiVO**](BatchBindDeviceResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoundSites

> OperationResponseBindSiteBriefOpenApiVO GetBoundSites(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get sites info which bound to the site template



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
	resp, r, err := apiClient.SiteTemplateAPI.GetBoundSites(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetBoundSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBoundSites`: OperationResponseBindSiteBriefOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetBoundSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoundSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseBindSiteBriefOpenApiVO**](OperationResponseBindSiteBriefOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDeviceTemplate

> OperationResponseGridVODeviceTemplateBriefOpenApiVO GetGridDeviceTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get deviceTemplate info list



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
	resp, r, err := apiClient.SiteTemplateAPI.GetGridDeviceTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetGridDeviceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDeviceTemplate`: OperationResponseGridVODeviceTemplateBriefOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetGridDeviceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDeviceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVODeviceTemplateBriefOpenApiVO**](OperationResponseGridVODeviceTemplateBriefOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDeviceTemplateAvailableDevices

> OperationResponseGridVODeviceBindOpenApiVO GetGridDeviceTemplateAvailableDevices(ctx, omadacId, siteTemplateId, deviceTemplateId).Page(page).PageSize(pageSize).Execute()

Get deviceTemplate's available devices



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.GetGridDeviceTemplateAvailableDevices(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetGridDeviceTemplateAvailableDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDeviceTemplateAvailableDevices`: OperationResponseGridVODeviceBindOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetGridDeviceTemplateAvailableDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDeviceTemplateAvailableDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVODeviceBindOpenApiVO**](OperationResponseGridVODeviceBindOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDeviceTemplateBoundDevices

> OperationResponseGridVODeviceBindOpenApiVO GetGridDeviceTemplateBoundDevices(ctx, omadacId, siteTemplateId, deviceTemplateId).Page(page).PageSize(pageSize).Execute()

Get deviceTemplate's bound devices



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.GetGridDeviceTemplateBoundDevices(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetGridDeviceTemplateBoundDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDeviceTemplateBoundDevices`: OperationResponseGridVODeviceBindOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetGridDeviceTemplateBoundDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDeviceTemplateBoundDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVODeviceBindOpenApiVO**](OperationResponseGridVODeviceBindOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDeviceTemplateDetail

> OperationResponseDeviceTemplateGridOpenApiVODeviceTemplateOpenApiQueryVO GetGridDeviceTemplateDetail(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

get device template



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
	resp, r, err := apiClient.SiteTemplateAPI.GetGridDeviceTemplateDetail(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetGridDeviceTemplateDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDeviceTemplateDetail`: OperationResponseDeviceTemplateGridOpenApiVODeviceTemplateOpenApiQueryVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetGridDeviceTemplateDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDeviceTemplateDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseDeviceTemplateGridOpenApiVODeviceTemplateOpenApiQueryVO**](OperationResponseDeviceTemplateGridOpenApiVODeviceTemplateOpenApiQueryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridDevicesAvailableTemplate

> OperationResponseGridVODeviceAvailableTemplateOpenApiVO GetGridDevicesAvailableTemplate(ctx, omadacId, siteId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

get the devices on the site and the templates that can be bound



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.GetGridDevicesAvailableTemplate(context.Background(), omadacId, siteId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetGridDevicesAvailableTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridDevicesAvailableTemplate`: OperationResponseGridVODeviceAvailableTemplateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetGridDevicesAvailableTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridDevicesAvailableTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVODeviceAvailableTemplateOpenApiVO**](OperationResponseGridVODeviceAvailableTemplateOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteTemplateAllModules

> OperationResponseSiteTemplateAllModulesOpenApiVO GetSiteTemplateAllModules(ctx, omadacId).Execute()

Get all function modules which can be selected when creating site template



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateAllModules(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateAllModules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateAllModules`: OperationResponseSiteTemplateAllModulesOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateAllModules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateAllModulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseSiteTemplateAllModulesOpenApiVO**](OperationResponseSiteTemplateAllModulesOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteTemplateConfiguration

> OperationResponseSiteTemplateSettings GetSiteTemplateConfiguration(ctx, omadacId, siteTemplateId).Execute()

get site template settings



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
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateConfiguration(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateConfiguration`: OperationResponseSiteTemplateSettings
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteTemplateSettings**](OperationResponseSiteTemplateSettings.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteTemplateDstInfo

> OperationResponse GetSiteTemplateDstInfo(ctx, omadacId, siteTemplateId).Execute()

Get daylight saving time information for the Site template configuration



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
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateDstInfo(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateDstInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateDstInfo`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateDstInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateDstInfoRequest struct via the builder pattern


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


## GetSiteTemplateEntity

> OperationResponseSiteTemplateBriefOpenApiVO GetSiteTemplateEntity(ctx, omadacId, siteTemplateId).Execute()

Get site template info



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
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateEntity(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateEntity`: OperationResponseSiteTemplateBriefOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteTemplateBriefOpenApiVO**](OperationResponseSiteTemplateBriefOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteTemplateGeneralConfig

> OperationResponseSiteServiceGeneralConfigOpenApiVO GetSiteTemplateGeneralConfig(ctx, omadacId, siteTemplateId).Execute()

get site template general config



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
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateGeneralConfig(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateGeneralConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateGeneralConfig`: OperationResponseSiteServiceGeneralConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateGeneralConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateGeneralConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteServiceGeneralConfigOpenApiVO**](OperationResponseSiteServiceGeneralConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteTemplateList

> OperationResponseGridVOSiteTemplateSummaryVO GetSiteTemplateList(ctx, omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()

Get site template list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateList`: OperationResponseGridVOSiteTemplateSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseGridVOSiteTemplateSummaryVO**](OperationResponseGridVOSiteTemplateSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteTemplateModules

> OperationResponseSiteTemplateModulesOpenApiVO GetSiteTemplateModules(ctx, omadacId).Execute()

Get function modules which can be selected when creating site template



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateModules(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateModules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateModules`: OperationResponseSiteTemplateModulesOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateModules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateModulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseSiteTemplateModulesOpenApiVO**](OperationResponseSiteTemplateModulesOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteTemplateService

> OperationResponseSiteServiceOpenApiVO GetSiteTemplateService(ctx, omadacId, siteTemplateId).Execute()

get site template service



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
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateService(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateService`: OperationResponseSiteServiceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteServiceOpenApiVO**](OperationResponseSiteServiceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteTemplateWirelessFeature

> OperationResponseSiteTemplateWirelessFeature GetSiteTemplateWirelessFeature(ctx, omadacId, siteTemplateId).Execute()

get site template WirelessFeature



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
	resp, r, err := apiClient.SiteTemplateAPI.GetSiteTemplateWirelessFeature(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSiteTemplateWirelessFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateWirelessFeature`: OperationResponseSiteTemplateWirelessFeature
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSiteTemplateWirelessFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateWirelessFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteTemplateWirelessFeature**](OperationResponseSiteTemplateWirelessFeature.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSitesOverrides

> OperationResponseSiteOverrideResultOpenApiVO GetSitesOverrides(ctx, omadacId).SiteIds(siteIds).Execute()

get site overrides



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
	siteIds := []string{"Inner_example"} // []string | Site IDs

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.GetSitesOverrides(context.Background(), omadacId).SiteIds(siteIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetSitesOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSitesOverrides`: OperationResponseSiteOverrideResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetSitesOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteIds** | **[]string** | Site IDs | 

### Return type

[**OperationResponseSiteOverrideResultOpenApiVO**](OperationResponseSiteOverrideResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidDeviceModel

> OperationResponseValidDeviceModelOpenApiVO GetValidDeviceModel(ctx, omadacId, siteTemplateId).Execute()

get valid device template



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
	resp, r, err := apiClient.SiteTemplateAPI.GetValidDeviceModel(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.GetValidDeviceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidDeviceModel`: OperationResponseValidDeviceModelOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.GetValidDeviceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidDeviceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseValidDeviceModelOpenApiVO**](OperationResponseValidDeviceModelOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDeviceTemplate

> OperationResponseWithoutResult ModifyDeviceTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).DeviceTemplateEdit(deviceTemplateEdit).Execute()

modify device template



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	deviceTemplateEdit := *openapiclient.NewDeviceTemplateEdit() // DeviceTemplateEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.ModifyDeviceTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).DeviceTemplateEdit(deviceTemplateEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.ModifyDeviceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDeviceTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.ModifyDeviceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDeviceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deviceTemplateEdit** | [**DeviceTemplateEdit**](DeviceTemplateEdit.md) |  | 

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


## UnbindSite

> OperationResponseWithoutResult UnbindSite(ctx, omadacId, siteTemplateId).BatchUnbindSites(batchUnbindSites).Execute()

Batch unbind sites from a site template



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
	batchUnbindSites := *openapiclient.NewBatchUnbindSites([]string{"SiteIds_example"}) // BatchUnbindSites | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.UnbindSite(context.Background(), omadacId, siteTemplateId).BatchUnbindSites(batchUnbindSites).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.UnbindSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.UnbindSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchUnbindSites** | [**BatchUnbindSites**](BatchUnbindSites.md) |  | 

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


## UpdateSiteTemplateEntity

> OperationResponseWithoutResult UpdateSiteTemplateEntity(ctx, omadacId, siteTemplateId).UpdateSiteByTemplate(updateSiteByTemplate).Execute()

Modify an existing site template



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
	updateSiteByTemplate := *openapiclient.NewUpdateSiteByTemplate("Name_example") // UpdateSiteByTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateAPI.UpdateSiteTemplateEntity(context.Background(), omadacId, siteTemplateId).UpdateSiteByTemplate(updateSiteByTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateAPI.UpdateSiteTemplateEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteTemplateEntity`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateAPI.UpdateSiteTemplateEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteTemplateEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSiteByTemplate** | [**UpdateSiteByTemplate**](UpdateSiteByTemplate.md) |  | 

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


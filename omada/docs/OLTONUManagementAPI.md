# OLTONUManagementAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditOnuGlobalConfigStatus**](OLTONUManagementAPI.md#editonuglobalconfigstatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/global-config/edit | Modify an existing ONU global config status
[**EditOnuInformationAdminStatusConfig**](OLTONUManagementAPI.md#editonuinformationadminstatusconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/informations/admin/edit | Modify an existing ONU admin status config
[**EditOnuInformationDescriptionConfig**](OLTONUManagementAPI.md#editonuinformationdescriptionconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/informations/description/edit | Modify an existing ONU description
[**EditOnuIsolationStatus**](OLTONUManagementAPI.md#editonuisolationstatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/global-config/isolation/edit | Modify an existing ONU  isolation status
[**GetOnuGlobalConfigStatus**](OLTONUManagementAPI.md#getonuglobalconfigstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/global-config/get | Get ONU global config status
[**GetOnuInformationConfigList**](OLTONUManagementAPI.md#getonuinformationconfiglist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/informations/list | Get ONU information config list
[**GetOnuInformationConfigPage**](OLTONUManagementAPI.md#getonuinformationconfigpage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/informations/page | Get ONU information config page
[**GetOnuInformationDetail**](OLTONUManagementAPI.md#getonuinformationdetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/informations/detail/get | Get ONU information detail
[**GetOnuIsolationStatus**](OLTONUManagementAPI.md#getonuisolationstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/global-config/isolation/get | Get ONU isolation status
[**GetOnuRebootStatusConfigList**](OLTONUManagementAPI.md#getonurebootstatusconfiglist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/informations/reboot-status/list | Get ONU reboot status
[**RebootOnu**](OLTONUManagementAPI.md#rebootonu) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-management/informations/onu/reboot | Reboot ONU



## EditOnuGlobalConfigStatus

> OperationResponseDeviceResponseBodyVoid EditOnuGlobalConfigStatus(ctx, omadacId, siteId, deviceMac).OnuGlobalConfigStatusEditDTO(onuGlobalConfigStatusEditDTO).Execute()

Modify an existing ONU global config status



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	onuGlobalConfigStatusEditDTO := *openapiclient.NewOnuGlobalConfigStatusEditDTO("OnuIsolation_example") // OnuGlobalConfigStatusEditDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.EditOnuGlobalConfigStatus(context.Background(), omadacId, siteId, deviceMac).OnuGlobalConfigStatusEditDTO(onuGlobalConfigStatusEditDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.EditOnuGlobalConfigStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuGlobalConfigStatus`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.EditOnuGlobalConfigStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOnuGlobalConfigStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **onuGlobalConfigStatusEditDTO** | [**OnuGlobalConfigStatusEditDTO**](OnuGlobalConfigStatusEditDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOnuInformationAdminStatusConfig

> OperationResponseDeviceResponseBodyVoid EditOnuInformationAdminStatusConfig(ctx, omadacId, siteId, deviceMac).OnuInformationAdminStatusEditConfigDTO(onuInformationAdminStatusEditConfigDTO).Execute()

Modify an existing ONU admin status config



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	onuInformationAdminStatusEditConfigDTO := *openapiclient.NewOnuInformationAdminStatusEditConfigDTO("AdminStatus_example", []string{"Keys_example"}) // OnuInformationAdminStatusEditConfigDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.EditOnuInformationAdminStatusConfig(context.Background(), omadacId, siteId, deviceMac).OnuInformationAdminStatusEditConfigDTO(onuInformationAdminStatusEditConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.EditOnuInformationAdminStatusConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuInformationAdminStatusConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.EditOnuInformationAdminStatusConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOnuInformationAdminStatusConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **onuInformationAdminStatusEditConfigDTO** | [**OnuInformationAdminStatusEditConfigDTO**](OnuInformationAdminStatusEditConfigDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOnuInformationDescriptionConfig

> OperationResponseDeviceResponseBodyVoid EditOnuInformationDescriptionConfig(ctx, omadacId, siteId, deviceMac).OnuInformationDescriptionEditConfigDTO(onuInformationDescriptionEditConfigDTO).Execute()

Modify an existing ONU description



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	onuInformationDescriptionEditConfigDTO := *openapiclient.NewOnuInformationDescriptionEditConfigDTO("Key_example", "OnuDescription_example") // OnuInformationDescriptionEditConfigDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.EditOnuInformationDescriptionConfig(context.Background(), omadacId, siteId, deviceMac).OnuInformationDescriptionEditConfigDTO(onuInformationDescriptionEditConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.EditOnuInformationDescriptionConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuInformationDescriptionConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.EditOnuInformationDescriptionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOnuInformationDescriptionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **onuInformationDescriptionEditConfigDTO** | [**OnuInformationDescriptionEditConfigDTO**](OnuInformationDescriptionEditConfigDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOnuIsolationStatus

> OperationResponseDeviceResponseBodyVoid EditOnuIsolationStatus(ctx, omadacId, siteId, deviceMac).OnuIsolationStatusDTO(onuIsolationStatusDTO).Execute()

Modify an existing ONU  isolation status



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	onuIsolationStatusDTO := *openapiclient.NewOnuIsolationStatusDTO() // OnuIsolationStatusDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.EditOnuIsolationStatus(context.Background(), omadacId, siteId, deviceMac).OnuIsolationStatusDTO(onuIsolationStatusDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.EditOnuIsolationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuIsolationStatus`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.EditOnuIsolationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOnuIsolationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **onuIsolationStatusDTO** | [**OnuIsolationStatusDTO**](OnuIsolationStatusDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuGlobalConfigStatus

> OperationResponseOnuGlobalConfigStatusDTO GetOnuGlobalConfigStatus(ctx, omadacId, siteId, deviceMac).Execute()

Get ONU global config status



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.GetOnuGlobalConfigStatus(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.GetOnuGlobalConfigStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuGlobalConfigStatus`: OperationResponseOnuGlobalConfigStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.GetOnuGlobalConfigStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnuGlobalConfigStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOnuGlobalConfigStatusDTO**](OperationResponseOnuGlobalConfigStatusDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuInformationConfigList

> OperationResponseListOnuInformationConfigDTO GetOnuInformationConfigList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get ONU information config list



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	dto := *openapiclient.NewOnuManagementListQueryRequestDTO("PonPort_example") // OnuManagementListQueryRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.GetOnuInformationConfigList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.GetOnuInformationConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuInformationConfigList`: OperationResponseListOnuInformationConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.GetOnuInformationConfigList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnuInformationConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuManagementListQueryRequestDTO**](OnuManagementListQueryRequestDTO.md) |  | 

### Return type

[**OperationResponseListOnuInformationConfigDTO**](OperationResponseListOnuInformationConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuInformationConfigPage

> OperationResponsePageResponseOnuInformationConfigDTO GetOnuInformationConfigPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get ONU information config page



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	dto := *openapiclient.NewOnuManagementPageQueryRequestDTO("PonPort_example") // OnuManagementPageQueryRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.GetOnuInformationConfigPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.GetOnuInformationConfigPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuInformationConfigPage`: OperationResponsePageResponseOnuInformationConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.GetOnuInformationConfigPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnuInformationConfigPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuManagementPageQueryRequestDTO**](OnuManagementPageQueryRequestDTO.md) |  | 

### Return type

[**OperationResponsePageResponseOnuInformationConfigDTO**](OperationResponsePageResponseOnuInformationConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuInformationDetail

> OperationResponseOnuDetailConfigDTO GetOnuInformationDetail(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get ONU information detail



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	dto := *openapiclient.NewOnuInfoDetailRequestDTO("Key_example") // OnuInfoDetailRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.GetOnuInformationDetail(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.GetOnuInformationDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuInformationDetail`: OperationResponseOnuDetailConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.GetOnuInformationDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnuInformationDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuInfoDetailRequestDTO**](OnuInfoDetailRequestDTO.md) |  | 

### Return type

[**OperationResponseOnuDetailConfigDTO**](OperationResponseOnuDetailConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuIsolationStatus

> OperationResponseOnuIsolationStatusDTO GetOnuIsolationStatus(ctx, omadacId, siteId, deviceMac).Execute()

Get ONU isolation status



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.GetOnuIsolationStatus(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.GetOnuIsolationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuIsolationStatus`: OperationResponseOnuIsolationStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.GetOnuIsolationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnuIsolationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOnuIsolationStatusDTO**](OperationResponseOnuIsolationStatusDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuRebootStatusConfigList

> OperationResponseListOnuInformationRebootStatusConfigDTO GetOnuRebootStatusConfigList(ctx, omadacId, siteId, deviceMac).Execute()

Get ONU reboot status



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.GetOnuRebootStatusConfigList(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.GetOnuRebootStatusConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuRebootStatusConfigList`: OperationResponseListOnuInformationRebootStatusConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.GetOnuRebootStatusConfigList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnuRebootStatusConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListOnuInformationRebootStatusConfigDTO**](OperationResponseListOnuInformationRebootStatusConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootOnu

> OperationResponseDeviceResponseBodyListSingleOnuRebootResponseDTO RebootOnu(ctx, omadacId, siteId, deviceMac).OnuInformationRebootRequestDTO(onuInformationRebootRequestDTO).Execute()

Reboot ONU



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	onuInformationRebootRequestDTO := *openapiclient.NewOnuInformationRebootRequestDTO([]openapiclient.SingleOnuRebootRequestDTO{*openapiclient.NewSingleOnuRebootRequestDTO()}) // OnuInformationRebootRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONUManagementAPI.RebootOnu(context.Background(), omadacId, siteId, deviceMac).OnuInformationRebootRequestDTO(onuInformationRebootRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONUManagementAPI.RebootOnu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootOnu`: OperationResponseDeviceResponseBodyListSingleOnuRebootResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONUManagementAPI.RebootOnu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootOnuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **onuInformationRebootRequestDTO** | [**OnuInformationRebootRequestDTO**](OnuInformationRebootRequestDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyListSingleOnuRebootResponseDTO**](OperationResponseDeviceResponseBodyListSingleOnuRebootResponseDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


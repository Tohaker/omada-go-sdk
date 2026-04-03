# CLIAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyCliConfig**](CLIAPI.md#applycliconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cli/configs/config-id/{configId}/apply | Apply the CLI configuration
[**CreateCliConfig**](CLIAPI.md#createcliconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cli/config/cli-type/{cliType}/save | Create a new CLI configuration
[**DeleteCliConfig**](CLIAPI.md#deletecliconfig) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/cli/configs/cli-type/{cliType} | Delete the CLI configuration
[**ExportCliVariableToFile**](CLIAPI.md#exportclivariabletofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cli/export/variable | Export CLI variable
[**GetCliConfig**](CLIAPI.md#getcliconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cli/config/cli-type/{cliType}/config-id/{configId} | Get the CLI configuration content
[**GetCliDevice**](CLIAPI.md#getclidevice) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cli/config/devices | Get device list when importing device configuration
[**GetCliDeviceByType**](CLIAPI.md#getclidevicebytype) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cli/config/devices/device-type/{deviceType} | Get device list when creating CLI
[**GetCliList**](CLIAPI.md#getclilist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cli/configs/cli-type/{cliType}/device-type/{deviceType} | Get the CLI configuration list
[**GetModelAndVersionListByType**](CLIAPI.md#getmodelandversionlistbytype) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/cli/models/device-type/{deviceType} | Get model and version for model CLI.
[**ImportCliVariableToFile**](CLIAPI.md#importclivariabletofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cli/import/variable | Import CLI variable
[**ModifyCliConfig**](CLIAPI.md#modifycliconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/cli/config/cli-type/{cliType}/config-id/{configId} | Modify the CLI configuration content



## ApplyCliConfig

> OperationResponseWithoutResult ApplyCliConfig(ctx, omadacId, siteId, configId).Execute()

Apply the CLI configuration



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
	configId := "configId_example" // string | CLI Config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.ApplyCliConfig(context.Background(), omadacId, siteId, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.ApplyCliConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyCliConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.ApplyCliConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**configId** | **string** | CLI Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyCliConfigRequest struct via the builder pattern


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


## CreateCliConfig

> OperationResponseWithoutResult CreateCliConfig(ctx, omadacId, siteId, cliType).ModifyCliOpenApiVO(modifyCliOpenApiVO).Execute()

Create a new CLI configuration



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	modifyCliOpenApiVO := *openapiclient.NewModifyCliOpenApiVO("CliConfig_example", "DeviceType_example", "Name_example") // ModifyCliOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.CreateCliConfig(context.Background(), omadacId, siteId, cliType).ModifyCliOpenApiVO(modifyCliOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.CreateCliConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCliConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.CreateCliConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCliConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyCliOpenApiVO** | [**ModifyCliOpenApiVO**](ModifyCliOpenApiVO.md) |  | 

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


## DeleteCliConfig

> OperationResponseWithoutResult DeleteCliConfig(ctx, omadacId, siteId, cliType).DeleteCliOpenApiVO(deleteCliOpenApiVO).Execute()

Delete the CLI configuration



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	deleteCliOpenApiVO := *openapiclient.NewDeleteCliOpenApiVO() // DeleteCliOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.DeleteCliConfig(context.Background(), omadacId, siteId, cliType).DeleteCliOpenApiVO(deleteCliOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.DeleteCliConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCliConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.DeleteCliConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCliConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteCliOpenApiVO** | [**DeleteCliOpenApiVO**](DeleteCliOpenApiVO.md) |  | 

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


## ExportCliVariableToFile

> OperationResponseWithoutResult ExportCliVariableToFile(ctx, omadacId, siteId).ExportCliVarOpenApiVO(exportCliVarOpenApiVO).Execute()

Export CLI variable



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
	exportCliVarOpenApiVO := *openapiclient.NewExportCliVarOpenApiVO([]openapiclient.DeviceExportCliVO{*openapiclient.NewDeviceExportCliVO("DeviceMac_example")}, int32(123)) // ExportCliVarOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.ExportCliVariableToFile(context.Background(), omadacId, siteId).ExportCliVarOpenApiVO(exportCliVarOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.ExportCliVariableToFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCliVariableToFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.ExportCliVariableToFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCliVariableToFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exportCliVarOpenApiVO** | [**ExportCliVarOpenApiVO**](ExportCliVarOpenApiVO.md) |  | 

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


## GetCliConfig

> OperationResponseCliConfigOpenApiVO GetCliConfig(ctx, omadacId, siteId, cliType, configId).Execute()

Get the CLI configuration content



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	configId := "configId_example" // string | CLI Config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.GetCliConfig(context.Background(), omadacId, siteId, cliType, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.GetCliConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCliConfig`: OperationResponseCliConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.GetCliConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 
**configId** | **string** | CLI Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCliConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseCliConfigOpenApiVO**](OperationResponseCliConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCliDevice

> OperationResponseListSupportCliDeviceVO GetCliDevice(ctx, omadacId, siteId).Execute()

Get device list when importing device configuration



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
	resp, r, err := apiClient.CLIAPI.GetCliDevice(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.GetCliDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCliDevice`: OperationResponseListSupportCliDeviceVO
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.GetCliDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCliDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListSupportCliDeviceVO**](OperationResponseListSupportCliDeviceVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCliDeviceByType

> OperationResponseListCliDeviceListVO GetCliDeviceByType(ctx, omadacId, siteId, deviceType).Execute()

Get device list when creating CLI



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
	deviceType := "deviceType_example" // string | Device Type. Supported type: switch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.GetCliDeviceByType(context.Background(), omadacId, siteId, deviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.GetCliDeviceByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCliDeviceByType`: OperationResponseListCliDeviceListVO
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.GetCliDeviceByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceType** | **string** | Device Type. Supported type: switch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCliDeviceByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListCliDeviceListVO**](OperationResponseListCliDeviceListVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCliList

> OperationResponseGridVOCliVO GetCliList(ctx, omadacId, siteId, cliType, deviceType).Page(page).PageSize(pageSize).Execute()

Get the CLI configuration list



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	deviceType := "deviceType_example" // string | Device Type. Supported type: switch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.GetCliList(context.Background(), omadacId, siteId, cliType, deviceType).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.GetCliList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCliList`: OperationResponseGridVOCliVO
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.GetCliList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 
**deviceType** | **string** | Device Type. Supported type: switch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCliListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





### Return type

[**OperationResponseGridVOCliVO**](OperationResponseGridVOCliVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelAndVersionListByType

> OperationResponseListModelAndModelVersionVO GetModelAndVersionListByType(ctx, omadacId, siteId, deviceType).Execute()

Get model and version for model CLI.



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
	deviceType := "deviceType_example" // string | Device Type. Supported type: switch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.GetModelAndVersionListByType(context.Background(), omadacId, siteId, deviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.GetModelAndVersionListByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelAndVersionListByType`: OperationResponseListModelAndModelVersionVO
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.GetModelAndVersionListByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceType** | **string** | Device Type. Supported type: switch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelAndVersionListByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListModelAndModelVersionVO**](OperationResponseListModelAndModelVersionVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCliVariableToFile

> OperationResponseImportCliVarOpenApiVO ImportCliVariableToFile(ctx, omadacId, siteId).ImportCliVariableToFileRequest(importCliVariableToFileRequest).Execute()

Import CLI variable



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
	importCliVariableToFileRequest := *openapiclient.NewImportCliVariableToFileRequest("TODO") // ImportCliVariableToFileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.ImportCliVariableToFile(context.Background(), omadacId, siteId).ImportCliVariableToFileRequest(importCliVariableToFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.ImportCliVariableToFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCliVariableToFile`: OperationResponseImportCliVarOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.ImportCliVariableToFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCliVariableToFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **importCliVariableToFileRequest** | [**ImportCliVariableToFileRequest**](ImportCliVariableToFileRequest.md) |  | 

### Return type

[**OperationResponseImportCliVarOpenApiVO**](OperationResponseImportCliVarOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCliConfig

> OperationResponseWithoutResult ModifyCliConfig(ctx, omadacId, siteId, cliType, configId).ModifyCliOpenApiVO(modifyCliOpenApiVO).Execute()

Modify the CLI configuration content



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	configId := "configId_example" // string | CLI Config ID
	modifyCliOpenApiVO := *openapiclient.NewModifyCliOpenApiVO("CliConfig_example", "DeviceType_example", "Name_example") // ModifyCliOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLIAPI.ModifyCliConfig(context.Background(), omadacId, siteId, cliType, configId).ModifyCliOpenApiVO(modifyCliOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.ModifyCliConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCliConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.ModifyCliConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 
**configId** | **string** | CLI Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCliConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **modifyCliOpenApiVO** | [**ModifyCliOpenApiVO**](ModifyCliOpenApiVO.md) |  | 

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


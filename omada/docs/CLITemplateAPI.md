# \CLITemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyCliConfigTemplate**](CLITemplateAPI.md#ApplyCliConfigTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/cli/configs/config-id/{configId}/apply | Apply the CLI template configuration
[**CreateCliConfigTemplate**](CLITemplateAPI.md#CreateCliConfigTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/cli/config/cli-type/{cliType}/save | Create a new CLI template configuration
[**DeleteCliConfigTemplate**](CLITemplateAPI.md#DeleteCliConfigTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/cli/configs/cli-type/{cliType} | Delete the CLI template configuration
[**GetCliConfigTemplate**](CLITemplateAPI.md#GetCliConfigTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/cli/config/cli-type/{cliType}/config-id/{configId} | Get the CLI template configuration content
[**GetGridCliTemplate**](CLITemplateAPI.md#GetGridCliTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/cli/configs/cli-type/{cliType}/device-type/{deviceType} | Get the CLI template configuration list
[**ModifyCliConfigTemplate**](CLITemplateAPI.md#ModifyCliConfigTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/cli/config/cli-type/{cliType}/config-id/{configId} | Modify the CLI template configuration content



## ApplyCliConfigTemplate

> OperationResponseWithoutResult ApplyCliConfigTemplate(ctx, omadacId, siteTemplateId, configId).Execute()

Apply the CLI template configuration



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
	configId := "configId_example" // string | CLI Config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLITemplateAPI.ApplyCliConfigTemplate(context.Background(), omadacId, siteTemplateId, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLITemplateAPI.ApplyCliConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyCliConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLITemplateAPI.ApplyCliConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**configId** | **string** | CLI Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyCliConfigTemplateRequest struct via the builder pattern


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


## CreateCliConfigTemplate

> OperationResponseWithoutResult CreateCliConfigTemplate(ctx, omadacId, siteTemplateId, cliType).ModifyCliTemplateOpenApiVO(modifyCliTemplateOpenApiVO).Execute()

Create a new CLI template configuration



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	modifyCliTemplateOpenApiVO := *openapiclient.NewModifyCliTemplateOpenApiVO("CliConfig_example", "DeviceType_example", "Name_example") // ModifyCliTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLITemplateAPI.CreateCliConfigTemplate(context.Background(), omadacId, siteTemplateId, cliType).ModifyCliTemplateOpenApiVO(modifyCliTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLITemplateAPI.CreateCliConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCliConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLITemplateAPI.CreateCliConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCliConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyCliTemplateOpenApiVO** | [**ModifyCliTemplateOpenApiVO**](ModifyCliTemplateOpenApiVO.md) |  | 

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


## DeleteCliConfigTemplate

> OperationResponseWithoutResult DeleteCliConfigTemplate(ctx, omadacId, siteTemplateId, cliType).DeleteCliTemplateOpenApiVO(deleteCliTemplateOpenApiVO).Execute()

Delete the CLI template configuration



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	deleteCliTemplateOpenApiVO := *openapiclient.NewDeleteCliTemplateOpenApiVO() // DeleteCliTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLITemplateAPI.DeleteCliConfigTemplate(context.Background(), omadacId, siteTemplateId, cliType).DeleteCliTemplateOpenApiVO(deleteCliTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLITemplateAPI.DeleteCliConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCliConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLITemplateAPI.DeleteCliConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCliConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteCliTemplateOpenApiVO** | [**DeleteCliTemplateOpenApiVO**](DeleteCliTemplateOpenApiVO.md) |  | 

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


## GetCliConfigTemplate

> OperationResponseCliConfigTemplateOpenApiVO GetCliConfigTemplate(ctx, omadacId, siteTemplateId, cliType, configId).Execute()

Get the CLI template configuration content



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	configId := "configId_example" // string | CLI Config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLITemplateAPI.GetCliConfigTemplate(context.Background(), omadacId, siteTemplateId, cliType, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLITemplateAPI.GetCliConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCliConfigTemplate`: OperationResponseCliConfigTemplateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CLITemplateAPI.GetCliConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 
**configId** | **string** | CLI Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCliConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseCliConfigTemplateOpenApiVO**](OperationResponseCliConfigTemplateOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridCliTemplate

> OperationResponseGridVOCliTemplateOpenApiVO GetGridCliTemplate(ctx, omadacId, siteTemplateId, cliType, deviceType).Page(page).PageSize(pageSize).Execute()

Get the CLI template configuration list



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	deviceType := "deviceType_example" // string | Device Type. Supported type: switch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLITemplateAPI.GetGridCliTemplate(context.Background(), omadacId, siteTemplateId, cliType, deviceType).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLITemplateAPI.GetGridCliTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridCliTemplate`: OperationResponseGridVOCliTemplateOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CLITemplateAPI.GetGridCliTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 
**deviceType** | **string** | Device Type. Supported type: switch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridCliTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





### Return type

[**OperationResponseGridVOCliTemplateOpenApiVO**](OperationResponseGridVOCliTemplateOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCliConfigTemplate

> OperationResponseWithoutResult ModifyCliConfigTemplate(ctx, omadacId, siteTemplateId, cliType, configId).ModifyCliTemplateOpenApiVO(modifyCliTemplateOpenApiVO).Execute()

Modify the CLI template configuration content



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
	cliType := "cliType_example" // string | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI.
	configId := "configId_example" // string | CLI Config ID
	modifyCliTemplateOpenApiVO := *openapiclient.NewModifyCliTemplateOpenApiVO("CliConfig_example", "DeviceType_example", "Name_example") // ModifyCliTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CLITemplateAPI.ModifyCliConfigTemplate(context.Background(), omadacId, siteTemplateId, cliType, configId).ModifyCliTemplateOpenApiVO(modifyCliTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLITemplateAPI.ModifyCliConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCliConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CLITemplateAPI.ModifyCliConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**cliType** | **string** | CLI type should be a value as follows: 0: site CLI; 1: device CLI； 2：model CLI. Site template only supports site CLI. Site template only supports site CLI. | 
**configId** | **string** | CLI Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCliConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **modifyCliTemplateOpenApiVO** | [**ModifyCliTemplateOpenApiVO**](ModifyCliTemplateOpenApiVO.md) |  | 

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


# \IPMACBindingAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteIpMacBindings**](IPMACBindingAPI.md#BatchDeleteIpMacBindings) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-binds/batch-delete | Batch delete IP-MAC bindings
[**CheckIpMacBindingExportToDhcpReservation**](IPMACBindingAPI.md#CheckIpMacBindingExportToDhcpReservation) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-binds/check/dhcp-res | Check IP-MAC binding export to DHCP reservation
[**CreateIpMacBinding**](IPMACBindingAPI.md#CreateIpMacBinding) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-binds | Create IP-MAC binding
[**DeleteIpMacBinding**](IPMACBindingAPI.md#DeleteIpMacBinding) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-binds/{bindId} | Delete IP-MAC binding
[**GetGridIpMacBinding**](IPMACBindingAPI.md#GetGridIpMacBinding) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-binds | Get IP-MAC binding list
[**GetIpMacBindingGeneralSetting**](IPMACBindingAPI.md#GetIpMacBindingGeneralSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-bind | Get IP-MAC binding general setting
[**ImportIpMacBindingListFromFile**](IPMACBindingAPI.md#ImportIpMacBindingListFromFile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-binds/import | Import IP-MAC bindings from file
[**ModifyIpMacBinding**](IPMACBindingAPI.md#ModifyIpMacBinding) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-binds/{bindId} | Modify IP-MAC binding
[**ModifyIpMacBindingGeneralSetting**](IPMACBindingAPI.md#ModifyIpMacBindingGeneralSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/ip-mac-bind | Modify IP-MAC binding general setting



## BatchDeleteIpMacBindings

> OperationResponseWithoutResult BatchDeleteIpMacBindings(ctx, omadacId, siteId).BatchIds(batchIds).Execute()

Batch delete IP-MAC bindings



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
	batchIds := *openapiclient.NewBatchIds([]string{"Ids_example"}) // BatchIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingAPI.BatchDeleteIpMacBindings(context.Background(), omadacId, siteId).BatchIds(batchIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.BatchDeleteIpMacBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteIpMacBindings`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.BatchDeleteIpMacBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteIpMacBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchIds** | [**BatchIds**](BatchIds.md) |  | 

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


## CheckIpMacBindingExportToDhcpReservation

> OperationResponse CheckIpMacBindingExportToDhcpReservation(ctx, omadacId, siteId).CheckIpMacBindingExportToDhcpReservationOpenApiVO(checkIpMacBindingExportToDhcpReservationOpenApiVO).Execute()

Check IP-MAC binding export to DHCP reservation



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
	checkIpMacBindingExportToDhcpReservationOpenApiVO := *openapiclient.NewCheckIpMacBindingExportToDhcpReservationOpenApiVO("InterfaceId_example", "Ip_example", "Mac_example") // CheckIpMacBindingExportToDhcpReservationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingAPI.CheckIpMacBindingExportToDhcpReservation(context.Background(), omadacId, siteId).CheckIpMacBindingExportToDhcpReservationOpenApiVO(checkIpMacBindingExportToDhcpReservationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.CheckIpMacBindingExportToDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckIpMacBindingExportToDhcpReservation`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.CheckIpMacBindingExportToDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckIpMacBindingExportToDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **checkIpMacBindingExportToDhcpReservationOpenApiVO** | [**CheckIpMacBindingExportToDhcpReservationOpenApiVO**](CheckIpMacBindingExportToDhcpReservationOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIpMacBinding

> OperationResponseWithoutResult CreateIpMacBinding(ctx, omadacId, siteId).IPMacBinding(iPMacBinding).Execute()

Create IP-MAC binding



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
	iPMacBinding := *openapiclient.NewIPMacBinding("InterfaceId_example", int32(123), "Ip_example", "Mac_example", false) // IPMacBinding | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingAPI.CreateIpMacBinding(context.Background(), omadacId, siteId).IPMacBinding(iPMacBinding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.CreateIpMacBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIpMacBinding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.CreateIpMacBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIpMacBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iPMacBinding** | [**IPMacBinding**](IPMacBinding.md) |  | 

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


## DeleteIpMacBinding

> OperationResponseWithoutResult DeleteIpMacBinding(ctx, omadacId, siteId, bindId).Execute()

Delete IP-MAC binding



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
	bindId := "bindId_example" // string | IP MAC binding ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingAPI.DeleteIpMacBinding(context.Background(), omadacId, siteId, bindId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.DeleteIpMacBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIpMacBinding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.DeleteIpMacBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**bindId** | **string** | IP MAC binding ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpMacBindingRequest struct via the builder pattern


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


## GetGridIpMacBinding

> OperationResponseGridVOIPMacBinding GetGridIpMacBinding(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get IP-MAC binding list



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
	resp, r, err := apiClient.IPMACBindingAPI.GetGridIpMacBinding(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.GetGridIpMacBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridIpMacBinding`: OperationResponseGridVOIPMacBinding
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.GetGridIpMacBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridIpMacBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOIPMacBinding**](OperationResponseGridVOIPMacBinding.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpMacBindingGeneralSetting

> OperationResponseIPMacBindingGeneralSetting GetIpMacBindingGeneralSetting(ctx, omadacId, siteId).Execute()

Get IP-MAC binding general setting



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
	resp, r, err := apiClient.IPMACBindingAPI.GetIpMacBindingGeneralSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.GetIpMacBindingGeneralSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpMacBindingGeneralSetting`: OperationResponseIPMacBindingGeneralSetting
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.GetIpMacBindingGeneralSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpMacBindingGeneralSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIPMacBindingGeneralSetting**](OperationResponseIPMacBindingGeneralSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportIpMacBindingListFromFile

> OperationResponseWithoutResult ImportIpMacBindingListFromFile(ctx, omadacId, siteId).UploadCertificateRequest(uploadCertificateRequest).Execute()

Import IP-MAC bindings from file



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
	uploadCertificateRequest := *openapiclient.NewUploadCertificateRequest("TODO") // UploadCertificateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingAPI.ImportIpMacBindingListFromFile(context.Background(), omadacId, siteId).UploadCertificateRequest(uploadCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.ImportIpMacBindingListFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportIpMacBindingListFromFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.ImportIpMacBindingListFromFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportIpMacBindingListFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadCertificateRequest** | [**UploadCertificateRequest**](UploadCertificateRequest.md) |  | 

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


## ModifyIpMacBinding

> OperationResponseWithoutResult ModifyIpMacBinding(ctx, omadacId, siteId, bindId).IPMacBinding(iPMacBinding).Execute()

Modify IP-MAC binding



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
	bindId := "bindId_example" // string | IP MAC binding ID
	iPMacBinding := *openapiclient.NewIPMacBinding("InterfaceId_example", int32(123), "Ip_example", "Mac_example", false) // IPMacBinding | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingAPI.ModifyIpMacBinding(context.Background(), omadacId, siteId, bindId).IPMacBinding(iPMacBinding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.ModifyIpMacBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpMacBinding`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.ModifyIpMacBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**bindId** | **string** | IP MAC binding ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpMacBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iPMacBinding** | [**IPMacBinding**](IPMacBinding.md) |  | 

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


## ModifyIpMacBindingGeneralSetting

> OperationResponseWithoutResult ModifyIpMacBindingGeneralSetting(ctx, omadacId, siteId).IPMacBindingGeneralSetting(iPMacBindingGeneralSetting).Execute()

Modify IP-MAC binding general setting



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
	iPMacBindingGeneralSetting := *openapiclient.NewIPMacBindingGeneralSetting(false) // IPMacBindingGeneralSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPMACBindingAPI.ModifyIpMacBindingGeneralSetting(context.Background(), omadacId, siteId).IPMacBindingGeneralSetting(iPMacBindingGeneralSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPMACBindingAPI.ModifyIpMacBindingGeneralSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpMacBindingGeneralSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IPMACBindingAPI.ModifyIpMacBindingGeneralSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpMacBindingGeneralSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iPMacBindingGeneralSetting** | [**IPMacBindingGeneralSetting**](IPMacBindingGeneralSetting.md) |  | 

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


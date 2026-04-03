# ArpDetectionAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArpDetects**](ArpDetectionAPI.md#createarpdetects) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/arpDetects | Create new arp detects
[**DeleteArpDetect**](ArpDetectionAPI.md#deletearpdetect) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/arpDetects/{arpDetectId} | Delete an existing arp detect
[**GetArpDetectStatus**](ArpDetectionAPI.md#getarpdetectstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/arpDetects/status | Get arp detect status
[**GetGridArpDetects**](ArpDetectionAPI.md#getgridarpdetects) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/arpDetects | Get arp detect list
[**GetOswPortLagIdVlans**](ArpDetectionAPI.md#getoswportlagidvlans) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/arpDetects/portLagVlans | Get the port or standard port or lag to vlans map of one osw.
[**ImportArpDetectImpbsFromFile**](ArpDetectionAPI.md#importarpdetectimpbsfromfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/arpDetects/impbs/import | Import the static impbs from file.
[**ModifyArpDetect**](ArpDetectionAPI.md#modifyarpdetect) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/arpDetects/{arpDetectId} | Modify a arp detect
[**ModifyArpDetectStatus**](ArpDetectionAPI.md#modifyarpdetectstatus) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/arpDetects/status | Modify arp detect status



## CreateArpDetects

> ResponseIdVO CreateArpDetects(ctx, omadacId, siteId).ArpDetectVO(arpDetectVO).Execute()

Create new arp detects



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
	arpDetectVO := *openapiclient.NewArpDetectVO() // ArpDetectVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArpDetectionAPI.CreateArpDetects(context.Background(), omadacId, siteId).ArpDetectVO(arpDetectVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpDetectionAPI.CreateArpDetects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateArpDetects`: ResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `ArpDetectionAPI.CreateArpDetects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArpDetectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **arpDetectVO** | [**ArpDetectVO**](ArpDetectVO.md) |  | 

### Return type

[**ResponseIdVO**](ResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArpDetect

> OperationResponseWithoutResult DeleteArpDetect(ctx, omadacId, siteId, arpDetectId).Execute()

Delete an existing arp detect



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
	arpDetectId := "arpDetectId_example" // string | arpDetectId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArpDetectionAPI.DeleteArpDetect(context.Background(), omadacId, siteId, arpDetectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpDetectionAPI.DeleteArpDetect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteArpDetect`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ArpDetectionAPI.DeleteArpDetect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**arpDetectId** | **string** | arpDetectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArpDetectRequest struct via the builder pattern


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


## GetArpDetectStatus

> OperationResponse GetArpDetectStatus(ctx, omadacId, siteId).Execute()

Get arp detect status



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
	resp, r, err := apiClient.ArpDetectionAPI.GetArpDetectStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpDetectionAPI.GetArpDetectStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArpDetectStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ArpDetectionAPI.GetArpDetectStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArpDetectStatusRequest struct via the builder pattern


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


## GetGridArpDetects

> OperationResponse GetGridArpDetects(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get arp detect list



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
	resp, r, err := apiClient.ArpDetectionAPI.GetGridArpDetects(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpDetectionAPI.GetGridArpDetects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridArpDetects`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ArpDetectionAPI.GetGridArpDetects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridArpDetectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



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


## GetOswPortLagIdVlans

> OperationResponse GetOswPortLagIdVlans(ctx, omadacId, siteId).ArpDetectOswVO(arpDetectOswVO).Execute()

Get the port or standard port or lag to vlans map of one osw.



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
	arpDetectOswVO := *openapiclient.NewArpDetectOswVO() // ArpDetectOswVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArpDetectionAPI.GetOswPortLagIdVlans(context.Background(), omadacId, siteId).ArpDetectOswVO(arpDetectOswVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpDetectionAPI.GetOswPortLagIdVlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswPortLagIdVlans`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ArpDetectionAPI.GetOswPortLagIdVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswPortLagIdVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **arpDetectOswVO** | [**ArpDetectOswVO**](ArpDetectOswVO.md) |  | 

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


## ImportArpDetectImpbsFromFile

> OperationResponse ImportArpDetectImpbsFromFile(ctx, omadacId, siteId).UploadCertificateRequest(uploadCertificateRequest).Execute()

Import the static impbs from file.



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
	resp, r, err := apiClient.ArpDetectionAPI.ImportArpDetectImpbsFromFile(context.Background(), omadacId, siteId).UploadCertificateRequest(uploadCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpDetectionAPI.ImportArpDetectImpbsFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportArpDetectImpbsFromFile`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ArpDetectionAPI.ImportArpDetectImpbsFromFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportArpDetectImpbsFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadCertificateRequest** | [**UploadCertificateRequest**](UploadCertificateRequest.md) |  | 

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


## ModifyArpDetect

> OperationResponseWithoutResult ModifyArpDetect(ctx, omadacId, siteId, arpDetectId).ArpDetectVO(arpDetectVO).Execute()

Modify a arp detect



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
	arpDetectId := "arpDetectId_example" // string | arpDetectId
	arpDetectVO := *openapiclient.NewArpDetectVO() // ArpDetectVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArpDetectionAPI.ModifyArpDetect(context.Background(), omadacId, siteId, arpDetectId).ArpDetectVO(arpDetectVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpDetectionAPI.ModifyArpDetect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyArpDetect`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ArpDetectionAPI.ModifyArpDetect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**arpDetectId** | **string** | arpDetectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyArpDetectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **arpDetectVO** | [**ArpDetectVO**](ArpDetectVO.md) |  | 

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


## ModifyArpDetectStatus

> OperationResponseWithoutResult ModifyArpDetectStatus(ctx, omadacId, siteId).ArpDetectStatusVO(arpDetectStatusVO).Execute()

Modify arp detect status



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
	arpDetectStatusVO := *openapiclient.NewArpDetectStatusVO() // ArpDetectStatusVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArpDetectionAPI.ModifyArpDetectStatus(context.Background(), omadacId, siteId).ArpDetectStatusVO(arpDetectStatusVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpDetectionAPI.ModifyArpDetectStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyArpDetectStatus`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ArpDetectionAPI.ModifyArpDetectStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyArpDetectStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **arpDetectStatusVO** | [**ArpDetectStatusVO**](ArpDetectStatusVO.md) |  | 

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


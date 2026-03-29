# \CertProfilesTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertProfile**](CertProfilesTemplateAPI.md#CreateCertProfile) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles | Create a new certificate profile Template
[**DeleteCertProfile**](CertProfilesTemplateAPI.md#DeleteCertProfile) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles/{certId} | Delete an exist certificate profile Template
[**DeleteCertProfileFile**](CertProfilesTemplateAPI.md#DeleteCertProfileFile) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles/delete-file/{fileId} | Delete an exist certificate profile Template file
[**GetCertProfileDetail**](CertProfilesTemplateAPI.md#GetCertProfileDetail) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles/{certId} | Get a certificate profile Template detail
[**GetGridCertProfile**](CertProfilesTemplateAPI.md#GetGridCertProfile) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles | Get certificate profile Template list
[**ModifyCertProfile**](CertProfilesTemplateAPI.md#ModifyCertProfile) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles/{certId} | Modify an exist certificate profile Template
[**UploadCaCertFile**](CertProfilesTemplateAPI.md#UploadCaCertFile) | **Post** /openapi/v1/{omadacId}/files/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles/ca-cert | Upload CA profile Template file
[**UploadClientCertFile**](CertProfilesTemplateAPI.md#UploadClientCertFile) | **Post** /openapi/v1/{omadacId}/files/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles/client-cert | Upload client certificate profile Template file
[**UploadClientPrivateKeyFile**](CertProfilesTemplateAPI.md#UploadClientPrivateKeyFile) | **Post** /openapi/v1/{omadacId}/files/sitetemplates/{siteTemplateId}/setting/profiles/cert-profiles/client-private-key | Upload client private key Template file



## CreateCertProfile

> OperationResponseWithoutResult CreateCertProfile(ctx, omadacId, siteTemplateId).CertProfileRequestOpenApiVO(certProfileRequestOpenApiVO).Execute()

Create a new certificate profile Template



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
	certProfileRequestOpenApiVO := *openapiclient.NewCertProfileRequestOpenApiVO(int32(123), "Name_example", int32(123)) // CertProfileRequestOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.CreateCertProfile(context.Background(), omadacId, siteTemplateId).CertProfileRequestOpenApiVO(certProfileRequestOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.CreateCertProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCertProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.CreateCertProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **certProfileRequestOpenApiVO** | [**CertProfileRequestOpenApiVO**](CertProfileRequestOpenApiVO.md) |  | 

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


## DeleteCertProfile

> OperationResponseDeleteCertOpenApiVO DeleteCertProfile(ctx, omadacId, siteTemplateId, certId).Execute()

Delete an exist certificate profile Template



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
	certId := "certId_example" // string | Cert profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.DeleteCertProfile(context.Background(), omadacId, siteTemplateId, certId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.DeleteCertProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertProfile`: OperationResponseDeleteCertOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.DeleteCertProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**certId** | **string** | Cert profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseDeleteCertOpenApiVO**](OperationResponseDeleteCertOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertProfileFile

> OperationResponseWithoutResult DeleteCertProfileFile(ctx, omadacId, siteTemplateId, fileId).Execute()

Delete an exist certificate profile Template file



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
	fileId := "fileId_example" // string | Cert profile file ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.DeleteCertProfileFile(context.Background(), omadacId, siteTemplateId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.DeleteCertProfileFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertProfileFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.DeleteCertProfileFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**fileId** | **string** | Cert profile file ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertProfileFileRequest struct via the builder pattern


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


## GetCertProfileDetail

> OperationResponseCertProfileDetailOpenApiVO GetCertProfileDetail(ctx, omadacId, siteTemplateId, certId).Execute()

Get a certificate profile Template detail



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
	certId := "certId_example" // string | Cert profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.GetCertProfileDetail(context.Background(), omadacId, siteTemplateId, certId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.GetCertProfileDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertProfileDetail`: OperationResponseCertProfileDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.GetCertProfileDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**certId** | **string** | Cert profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertProfileDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseCertProfileDetailOpenApiVO**](OperationResponseCertProfileDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridCertProfile

> OperationResponseCertProfileOpenApiVO GetGridCertProfile(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).FiltersType(filtersType).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()

Get certificate profile Template list



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
	filtersType := int32(56) // int32 | Filter query parameters, support field type, it should be a value as follows: 0: CA Cert; 1: Client Cert. (optional)
	filtersStatus := int32(56) // int32 | Filter query parameters, support field status, it should be a value as follows: 0: Normal; 1: Expired Soon; 2: Expired (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.GetGridCertProfile(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).FiltersType(filtersType).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.GetGridCertProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridCertProfile`: OperationResponseCertProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.GetGridCertProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridCertProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersType** | **int32** | Filter query parameters, support field type, it should be a value as follows: 0: CA Cert; 1: Client Cert. | 
 **filtersStatus** | **int32** | Filter query parameters, support field status, it should be a value as follows: 0: Normal; 1: Expired Soon; 2: Expired | 
 **searchKey** | **string** | Fuzzy query parameters, support field Name | 

### Return type

[**OperationResponseCertProfileOpenApiVO**](OperationResponseCertProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCertProfile

> OperationResponseWithoutResult ModifyCertProfile(ctx, omadacId, siteTemplateId, certId).CertProfileRequestOpenApiVO(certProfileRequestOpenApiVO).Execute()

Modify an exist certificate profile Template



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
	certId := "certId_example" // string | Certificate profile ID
	certProfileRequestOpenApiVO := *openapiclient.NewCertProfileRequestOpenApiVO(int32(123), "Name_example", int32(123)) // CertProfileRequestOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.ModifyCertProfile(context.Background(), omadacId, siteTemplateId, certId).CertProfileRequestOpenApiVO(certProfileRequestOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.ModifyCertProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCertProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.ModifyCertProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**certId** | **string** | Certificate profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCertProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **certProfileRequestOpenApiVO** | [**CertProfileRequestOpenApiVO**](CertProfileRequestOpenApiVO.md) |  | 

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


## UploadCaCertFile

> OperationResponseUploadCertResponseOpenApiVO UploadCaCertFile(ctx, omadacId, siteTemplateId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()

Upload CA profile Template file



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
	data := "data_example" // string | Certificate profile file name
	uploadCaCertFile1Request := *openapiclient.NewUploadCaCertFile1Request("TODO") // UploadCaCertFile1Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.UploadCaCertFile(context.Background(), omadacId, siteTemplateId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.UploadCaCertFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCaCertFile`: OperationResponseUploadCertResponseOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.UploadCaCertFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCaCertFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **string** | Certificate profile file name | 
 **uploadCaCertFile1Request** | [**UploadCaCertFile1Request**](UploadCaCertFile1Request.md) |  | 

### Return type

[**OperationResponseUploadCertResponseOpenApiVO**](OperationResponseUploadCertResponseOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadClientCertFile

> OperationResponseUploadCertResponseOpenApiVO UploadClientCertFile(ctx, omadacId, siteTemplateId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()

Upload client certificate profile Template file



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
	data := "data_example" // string | Certificate profile file name
	uploadCaCertFile1Request := *openapiclient.NewUploadCaCertFile1Request("TODO") // UploadCaCertFile1Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.UploadClientCertFile(context.Background(), omadacId, siteTemplateId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.UploadClientCertFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadClientCertFile`: OperationResponseUploadCertResponseOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.UploadClientCertFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadClientCertFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **string** | Certificate profile file name | 
 **uploadCaCertFile1Request** | [**UploadCaCertFile1Request**](UploadCaCertFile1Request.md) |  | 

### Return type

[**OperationResponseUploadCertResponseOpenApiVO**](OperationResponseUploadCertResponseOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadClientPrivateKeyFile

> OperationResponseUploadCertResponseOpenApiVO UploadClientPrivateKeyFile(ctx, omadacId, siteTemplateId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()

Upload client private key Template file



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
	data := "data_example" // string | Certificate profile file name
	uploadCaCertFile1Request := *openapiclient.NewUploadCaCertFile1Request("TODO") // UploadCaCertFile1Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesTemplateAPI.UploadClientPrivateKeyFile(context.Background(), omadacId, siteTemplateId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesTemplateAPI.UploadClientPrivateKeyFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadClientPrivateKeyFile`: OperationResponseUploadCertResponseOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesTemplateAPI.UploadClientPrivateKeyFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadClientPrivateKeyFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **string** | Certificate profile file name | 
 **uploadCaCertFile1Request** | [**UploadCaCertFile1Request**](UploadCaCertFile1Request.md) |  | 

### Return type

[**OperationResponseUploadCertResponseOpenApiVO**](OperationResponseUploadCertResponseOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


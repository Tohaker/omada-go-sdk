# \CertProfilesAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertProfile1**](CertProfilesAPI.md#CreateCertProfile1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/profiles/cert-profiles | Create a new certificate profile
[**DeleteCertProfile1**](CertProfilesAPI.md#DeleteCertProfile1) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/profiles/cert-profiles/{certId} | Delete an exist certificate profile
[**DeleteCertProfileFile1**](CertProfilesAPI.md#DeleteCertProfileFile1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/profiles/cert-profiles/delete-file/{fileId} | Delete an exist certificate profile file
[**GetAllCertProfile**](CertProfilesAPI.md#GetAllCertProfile) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/profiles/all-cert-profiles | Get certificate profile list
[**GetCertProfileDetail1**](CertProfilesAPI.md#GetCertProfileDetail1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/profiles/cert-profiles/{certId} | Get a certificate profile detail
[**GetGridCertProfile1**](CertProfilesAPI.md#GetGridCertProfile1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/profiles/cert-profiles | Get grid certificate profile list
[**ModifyCertProfile1**](CertProfilesAPI.md#ModifyCertProfile1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/profiles/cert-profiles/{certId} | Modify an exist certificate profile
[**UploadCaCertFile1**](CertProfilesAPI.md#UploadCaCertFile1) | **Post** /openapi/v1/{omadacId}/files/sites/{siteId}/setting/profiles/cert-profiles/ca-cert | Upload CA profile file
[**UploadClientCertFile1**](CertProfilesAPI.md#UploadClientCertFile1) | **Post** /openapi/v1/{omadacId}/files/sites/{siteId}/setting/profiles/cert-profiles/client-cert | Upload client certificate profile file
[**UploadClientPrivateKeyFile1**](CertProfilesAPI.md#UploadClientPrivateKeyFile1) | **Post** /openapi/v1/{omadacId}/files/sites/{siteId}/setting/profiles/cert-profiles/client-private-key | Upload client private key file



## CreateCertProfile1

> OperationResponseWithoutResult CreateCertProfile1(ctx, omadacId, siteId).CertProfileRequestOpenApiVO(certProfileRequestOpenApiVO).Execute()

Create a new certificate profile



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
	certProfileRequestOpenApiVO := *openapiclient.NewCertProfileRequestOpenApiVO(int32(123), "Name_example", int32(123)) // CertProfileRequestOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.CreateCertProfile1(context.Background(), omadacId, siteId).CertProfileRequestOpenApiVO(certProfileRequestOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.CreateCertProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCertProfile1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.CreateCertProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertProfile1Request struct via the builder pattern


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


## DeleteCertProfile1

> OperationResponseDeleteCertOpenApiVO DeleteCertProfile1(ctx, omadacId, siteId, certId).Execute()

Delete an exist certificate profile



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
	certId := "certId_example" // string | Cert profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.DeleteCertProfile1(context.Background(), omadacId, siteId, certId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.DeleteCertProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertProfile1`: OperationResponseDeleteCertOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.DeleteCertProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**certId** | **string** | Cert profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertProfile1Request struct via the builder pattern


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


## DeleteCertProfileFile1

> OperationResponseWithoutResult DeleteCertProfileFile1(ctx, omadacId, siteId, fileId).Execute()

Delete an exist certificate profile file



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
	fileId := "fileId_example" // string | Cert profile file ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.DeleteCertProfileFile1(context.Background(), omadacId, siteId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.DeleteCertProfileFile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertProfileFile1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.DeleteCertProfileFile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**fileId** | **string** | Cert profile file ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertProfileFile1Request struct via the builder pattern


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


## GetAllCertProfile

> OperationResponseCertProfileOpenApiVO GetAllCertProfile(ctx, omadacId, siteId).Execute()

Get certificate profile list



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
	resp, r, err := apiClient.CertProfilesAPI.GetAllCertProfile(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.GetAllCertProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCertProfile`: OperationResponseCertProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.GetAllCertProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCertProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetCertProfileDetail1

> OperationResponseCertProfileDetailOpenApiVO GetCertProfileDetail1(ctx, omadacId, siteId, certId).Execute()

Get a certificate profile detail



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
	certId := "certId_example" // string | Cert profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.GetCertProfileDetail1(context.Background(), omadacId, siteId, certId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.GetCertProfileDetail1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertProfileDetail1`: OperationResponseCertProfileDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.GetCertProfileDetail1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**certId** | **string** | Cert profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertProfileDetail1Request struct via the builder pattern


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


## GetGridCertProfile1

> OperationResponseCertProfileOpenApiVO GetGridCertProfile1(ctx, omadacId, siteId).Page(page).PageSize(pageSize).FiltersType(filtersType).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()

Get grid certificate profile list



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
	filtersType := int32(56) // int32 | Filter query parameters, support field type, it should be a value as follows: 0: CA Cert; 1: Client Cert. (optional)
	filtersStatus := int32(56) // int32 | Filter query parameters, support field status, it should be a value as follows: 0: Normal; 1: Expired Soon; 2: Expired (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.GetGridCertProfile1(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).FiltersType(filtersType).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.GetGridCertProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridCertProfile1`: OperationResponseCertProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.GetGridCertProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridCertProfile1Request struct via the builder pattern


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


## ModifyCertProfile1

> OperationResponseWithoutResult ModifyCertProfile1(ctx, omadacId, siteId, certId).CertProfileRequestOpenApiVO(certProfileRequestOpenApiVO).Execute()

Modify an exist certificate profile



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
	certId := "certId_example" // string | Certificate profile ID
	certProfileRequestOpenApiVO := *openapiclient.NewCertProfileRequestOpenApiVO(int32(123), "Name_example", int32(123)) // CertProfileRequestOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.ModifyCertProfile1(context.Background(), omadacId, siteId, certId).CertProfileRequestOpenApiVO(certProfileRequestOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.ModifyCertProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCertProfile1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.ModifyCertProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**certId** | **string** | Certificate profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCertProfile1Request struct via the builder pattern


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


## UploadCaCertFile1

> OperationResponseUploadCertResponseOpenApiVO UploadCaCertFile1(ctx, omadacId, siteId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()

Upload CA profile file



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
	data := "data_example" // string | Certificate profile file name
	uploadCaCertFile1Request := *openapiclient.NewUploadCaCertFile1Request("TODO") // UploadCaCertFile1Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.UploadCaCertFile1(context.Background(), omadacId, siteId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.UploadCaCertFile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCaCertFile1`: OperationResponseUploadCertResponseOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.UploadCaCertFile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCaCertFile1Request struct via the builder pattern


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


## UploadClientCertFile1

> OperationResponseUploadCertResponseOpenApiVO UploadClientCertFile1(ctx, omadacId, siteId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()

Upload client certificate profile file



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
	data := "data_example" // string | Certificate profile file name
	uploadCaCertFile1Request := *openapiclient.NewUploadCaCertFile1Request("TODO") // UploadCaCertFile1Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.UploadClientCertFile1(context.Background(), omadacId, siteId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.UploadClientCertFile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadClientCertFile1`: OperationResponseUploadCertResponseOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.UploadClientCertFile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadClientCertFile1Request struct via the builder pattern


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


## UploadClientPrivateKeyFile1

> OperationResponseUploadCertResponseOpenApiVO UploadClientPrivateKeyFile1(ctx, omadacId, siteId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()

Upload client private key file



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
	data := "data_example" // string | Certificate profile file name
	uploadCaCertFile1Request := *openapiclient.NewUploadCaCertFile1Request("TODO") // UploadCaCertFile1Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertProfilesAPI.UploadClientPrivateKeyFile1(context.Background(), omadacId, siteId).Data(data).UploadCaCertFile1Request(uploadCaCertFile1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertProfilesAPI.UploadClientPrivateKeyFile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadClientPrivateKeyFile1`: OperationResponseUploadCertResponseOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CertProfilesAPI.UploadClientPrivateKeyFile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadClientPrivateKeyFile1Request struct via the builder pattern


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


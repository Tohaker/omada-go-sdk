# \LicenseAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoActive**](LicenseAPI.md#GetAutoActive) | **Get** /openapi/v1/{omadacId}/license/auto-active | Get license auto active
[**GetAutoRenew**](LicenseAPI.md#GetAutoRenew) | **Get** /openapi/v1/{omadacId}/license/auto-renew | Get license auto renew
[**GetLicenseNumByCategory**](LicenseAPI.md#GetLicenseNumByCategory) | **Get** /openapi/v1/{omadacId}/licenses/available-license | Get available license num
[**GetSiteAutoRenew**](LicenseAPI.md#GetSiteAutoRenew) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/license/auto-renew | Get site license auto renew
[**ModifyAutoActive**](LicenseAPI.md#ModifyAutoActive) | **Post** /openapi/v1/{omadacId}/license/auto-active | Modify license auto active
[**ModifyAutoRenew**](LicenseAPI.md#ModifyAutoRenew) | **Post** /openapi/v1/{omadacId}/license/auto-renew | Modify license auto renew
[**ModifySiteAutoRenew**](LicenseAPI.md#ModifySiteAutoRenew) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/license/auto-renew | Modify site license auto renew



## GetAutoActive

> OperationResponseLicenseAutoActiveOpenApiVO GetAutoActive(ctx, omadacId).Execute()

Get license auto active



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
	resp, r, err := apiClient.LicenseAPI.GetAutoActive(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.GetAutoActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoActive`: OperationResponseLicenseAutoActiveOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.GetAutoActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseLicenseAutoActiveOpenApiVO**](OperationResponseLicenseAutoActiveOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoRenew

> OperationResponseLicenseAutoRenewOpenApiVO GetAutoRenew(ctx, omadacId).Execute()

Get license auto renew



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
	resp, r, err := apiClient.LicenseAPI.GetAutoRenew(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.GetAutoRenew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoRenew`: OperationResponseLicenseAutoRenewOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.GetAutoRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseLicenseAutoRenewOpenApiVO**](OperationResponseLicenseAutoRenewOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseNumByCategory

> LicenseBindAvailableNumOpenApiVO GetLicenseNumByCategory(ctx, omadacId).Category(category).Execute()

Get available license num



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
	category := "category_example" // string | It should be a value as follows: basic; ap; l2Switch; l3Switch; gateway

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseAPI.GetLicenseNumByCategory(context.Background(), omadacId).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.GetLicenseNumByCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseNumByCategory`: LicenseBindAvailableNumOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.GetLicenseNumByCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseNumByCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | It should be a value as follows: basic; ap; l2Switch; l3Switch; gateway | 

### Return type

[**LicenseBindAvailableNumOpenApiVO**](LicenseBindAvailableNumOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteAutoRenew

> OperationResponseLicenseAutoRenewOpenApiVO GetSiteAutoRenew(ctx, omadacId, siteId).Execute()

Get site license auto renew



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
	resp, r, err := apiClient.LicenseAPI.GetSiteAutoRenew(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.GetSiteAutoRenew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAutoRenew`: OperationResponseLicenseAutoRenewOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.GetSiteAutoRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAutoRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseLicenseAutoRenewOpenApiVO**](OperationResponseLicenseAutoRenewOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAutoActive

> OperationResponseWithoutResult ModifyAutoActive(ctx, omadacId).LicenseAutoActiveOpenApiVO(licenseAutoActiveOpenApiVO).Execute()

Modify license auto active



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
	licenseAutoActiveOpenApiVO := *openapiclient.NewLicenseAutoActiveOpenApiVO(false) // LicenseAutoActiveOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseAPI.ModifyAutoActive(context.Background(), omadacId).LicenseAutoActiveOpenApiVO(licenseAutoActiveOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.ModifyAutoActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAutoActive`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.ModifyAutoActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAutoActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseAutoActiveOpenApiVO** | [**LicenseAutoActiveOpenApiVO**](LicenseAutoActiveOpenApiVO.md) |  | 

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


## ModifyAutoRenew

> OperationResponseWithoutResult ModifyAutoRenew(ctx, omadacId).LicenseAutoRenewOpenApiVO(licenseAutoRenewOpenApiVO).Execute()

Modify license auto renew



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
	licenseAutoRenewOpenApiVO := *openapiclient.NewLicenseAutoRenewOpenApiVO(false) // LicenseAutoRenewOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseAPI.ModifyAutoRenew(context.Background(), omadacId).LicenseAutoRenewOpenApiVO(licenseAutoRenewOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.ModifyAutoRenew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAutoRenew`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.ModifyAutoRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAutoRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseAutoRenewOpenApiVO** | [**LicenseAutoRenewOpenApiVO**](LicenseAutoRenewOpenApiVO.md) |  | 

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


## ModifySiteAutoRenew

> OperationResponseWithoutResult ModifySiteAutoRenew(ctx, omadacId, siteId).LicenseAutoRenewOpenApiVO(licenseAutoRenewOpenApiVO).Execute()

Modify site license auto renew



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
	licenseAutoRenewOpenApiVO := *openapiclient.NewLicenseAutoRenewOpenApiVO(false) // LicenseAutoRenewOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseAPI.ModifySiteAutoRenew(context.Background(), omadacId, siteId).LicenseAutoRenewOpenApiVO(licenseAutoRenewOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.ModifySiteAutoRenew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySiteAutoRenew`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.ModifySiteAutoRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySiteAutoRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **licenseAutoRenewOpenApiVO** | [**LicenseAutoRenewOpenApiVO**](LicenseAutoRenewOpenApiVO.md) |  | 

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


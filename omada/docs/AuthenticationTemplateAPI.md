# AuthenticationTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortalTemplate**](AuthenticationTemplateAPI.md#createportaltemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/portal | Create portal template
[**DeletePortalTemplate**](AuthenticationTemplateAPI.md#deleteportaltemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/portal/{portalId} | Delete portal template
[**GetMacAuthSetting**](AuthenticationTemplateAPI.md#getmacauthsetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-auth | Get site MAC-Based Authentication Template info
[**GetMacAuthSsids**](AuthenticationTemplateAPI.md#getmacauthssids) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-auth/ssids | Get ssids that support MAC auth Template
[**GetPortalListTemplate**](AuthenticationTemplateAPI.md#getportallisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/portals | Get portal list in a site template 
[**GetPortalTemplateCustomization**](AuthenticationTemplateAPI.md#getportaltemplatecustomization) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/portal/{portalId}/customization | Get portal template customization
[**GetPortalTemplateDetail**](AuthenticationTemplateAPI.md#getportaltemplatedetail) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/portal/{portalId} | Get portal template detail
[**ModifyPortalTemplate**](AuthenticationTemplateAPI.md#modifyportaltemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/portal/{portalId} | Modify portal template
[**UpdateMacAuthSetting**](AuthenticationTemplateAPI.md#updatemacauthsetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mac-auth | Update site MAC-Based Authentication Template info
[**UploadPortalPageTemplate**](AuthenticationTemplateAPI.md#uploadportalpagetemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/portal/page | Import portal page in a site template
[**UploadPortalPicTemplate**](AuthenticationTemplateAPI.md#uploadportalpictemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/portal/pics | Upload portal picture in a site template 



## CreatePortalTemplate

> OperationResponseWithoutResult CreatePortalTemplate(ctx, omadacId, siteTemplateId).PortalSetting(portalSetting).Execute()

Create portal template



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
	portalSetting := *openapiclient.NewPortalSetting(*openapiclient.NewAuthTimeoutSetting(), int32(123), false, false, int32(123), "Name_example") // PortalSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationTemplateAPI.CreatePortalTemplate(context.Background(), omadacId, siteTemplateId).PortalSetting(portalSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.CreatePortalTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortalTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.CreatePortalTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortalTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalSetting** | [**PortalSetting**](PortalSetting.md) |  | 

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


## DeletePortalTemplate

> OperationResponseWithoutResult DeletePortalTemplate(ctx, omadacId, siteTemplateId, portalId).Execute()

Delete portal template



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
	portalId := "portalId_example" // string | Portal ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationTemplateAPI.DeletePortalTemplate(context.Background(), omadacId, siteTemplateId, portalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.DeletePortalTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortalTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.DeletePortalTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**portalId** | **string** | Portal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortalTemplateRequest struct via the builder pattern


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


## GetMacAuthSetting

> OperationResponseMacAuthOpenApiVO GetMacAuthSetting(ctx, omadacId, siteTemplateId).Execute()

Get site MAC-Based Authentication Template info



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
	resp, r, err := apiClient.AuthenticationTemplateAPI.GetMacAuthSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.GetMacAuthSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacAuthSetting`: OperationResponseMacAuthOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.GetMacAuthSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacAuthSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMacAuthOpenApiVO**](OperationResponseMacAuthOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacAuthSsids

> OperationResponseListWlanSimpleOpenApiVO GetMacAuthSsids(ctx, omadacId, siteTemplateId).Execute()

Get ssids that support MAC auth Template



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
	resp, r, err := apiClient.AuthenticationTemplateAPI.GetMacAuthSsids(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.GetMacAuthSsids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacAuthSsids`: OperationResponseListWlanSimpleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.GetMacAuthSsids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacAuthSsidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListWlanSimpleOpenApiVO**](OperationResponseListWlanSimpleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortalListTemplate

> OperationResponseListPortalResOpenApiVO GetPortalListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get portal list in a site template 



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
	resp, r, err := apiClient.AuthenticationTemplateAPI.GetPortalListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.GetPortalListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalListTemplate`: OperationResponseListPortalResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.GetPortalListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListPortalResOpenApiVO**](OperationResponseListPortalResOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortalTemplateCustomization

> OperationResponsePortalCustomizationPageResOpenApiVO GetPortalTemplateCustomization(ctx, omadacId, siteTemplateId, portalId).Execute()

Get portal template customization



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
	portalId := "portalId_example" // string | Portal ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationTemplateAPI.GetPortalTemplateCustomization(context.Background(), omadacId, siteTemplateId, portalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.GetPortalTemplateCustomization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalTemplateCustomization`: OperationResponsePortalCustomizationPageResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.GetPortalTemplateCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**portalId** | **string** | Portal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalTemplateCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponsePortalCustomizationPageResOpenApiVO**](OperationResponsePortalCustomizationPageResOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortalTemplateDetail

> OperationResponsePortalDetailResOpenApiVO GetPortalTemplateDetail(ctx, omadacId, siteTemplateId, portalId).Execute()

Get portal template detail



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
	portalId := "portalId_example" // string | Portal ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationTemplateAPI.GetPortalTemplateDetail(context.Background(), omadacId, siteTemplateId, portalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.GetPortalTemplateDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalTemplateDetail`: OperationResponsePortalDetailResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.GetPortalTemplateDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**portalId** | **string** | Portal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalTemplateDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponsePortalDetailResOpenApiVO**](OperationResponsePortalDetailResOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPortalTemplate

> OperationResponseWithoutResult ModifyPortalTemplate(ctx, omadacId, siteTemplateId, portalId).PortalSetting(portalSetting).Execute()

Modify portal template



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
	portalId := "portalId_example" // string | Portal ID
	portalSetting := *openapiclient.NewPortalSetting(*openapiclient.NewAuthTimeoutSetting(), int32(123), false, false, int32(123), "Name_example") // PortalSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationTemplateAPI.ModifyPortalTemplate(context.Background(), omadacId, siteTemplateId, portalId).PortalSetting(portalSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.ModifyPortalTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortalTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.ModifyPortalTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**portalId** | **string** | Portal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortalTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **portalSetting** | [**PortalSetting**](PortalSetting.md) |  | 

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


## UpdateMacAuthSetting

> OperationResponseWithoutResult UpdateMacAuthSetting(ctx, omadacId, siteTemplateId).UpdateMacAuthOpenApiVO(updateMacAuthOpenApiVO).Execute()

Update site MAC-Based Authentication Template info



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
	updateMacAuthOpenApiVO := *openapiclient.NewUpdateMacAuthOpenApiVO(false, []string{"Ssids_example"}) // UpdateMacAuthOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationTemplateAPI.UpdateMacAuthSetting(context.Background(), omadacId, siteTemplateId).UpdateMacAuthOpenApiVO(updateMacAuthOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.UpdateMacAuthSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMacAuthSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.UpdateMacAuthSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMacAuthSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMacAuthOpenApiVO** | [**UpdateMacAuthOpenApiVO**](UpdateMacAuthOpenApiVO.md) |  | 

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


## UploadPortalPageTemplate

> OperationResponseImportedPortalPageResOpenApiVO UploadPortalPageTemplate(ctx, omadacId, siteTemplateId).File(file).Execute()

Import portal page in a site template



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
	file := os.NewFile(1234, "some_file") // *os.File | Portal page file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationTemplateAPI.UploadPortalPageTemplate(context.Background(), omadacId, siteTemplateId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.UploadPortalPageTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPortalPageTemplate`: OperationResponseImportedPortalPageResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.UploadPortalPageTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadPortalPageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | Portal page file. | 

### Return type

[**OperationResponseImportedPortalPageResOpenApiVO**](OperationResponseImportedPortalPageResOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPortalPicTemplate

> OperationResponsePortalPictureInfo UploadPortalPicTemplate(ctx, omadacId, siteTemplateId).Type_(type_).File(file).Execute()

Upload portal picture in a site template 



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
	type_ := "type__example" // string | Portal picture type: background/logo/advertisement.
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	file := os.NewFile(1234, "some_file") // *os.File | Portal picture file. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationTemplateAPI.UploadPortalPicTemplate(context.Background(), omadacId, siteTemplateId).Type_(type_).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTemplateAPI.UploadPortalPicTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPortalPicTemplate`: OperationResponsePortalPictureInfo
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationTemplateAPI.UploadPortalPicTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadPortalPicTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Portal picture type: background/logo/advertisement. | 


 **file** | ***os.File** | Portal picture file. | 

### Return type

[**OperationResponsePortalPictureInfo**](OperationResponsePortalPictureInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


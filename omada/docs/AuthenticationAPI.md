# AuthenticationAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPortal**](AuthenticationAPI.md#addportal) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/portal | Add portal
[**DeletePortal**](AuthenticationAPI.md#deleteportal) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/portal/{portalId} | Delete portal
[**GetEapDot1xCandidates**](AuthenticationAPI.md#geteapdot1xcandidates) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dot1x/eap/candidates | Get site EAP 802.1x setting candidates
[**GetEapDot1xSetting**](AuthenticationAPI.md#geteapdot1xsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dot1x/eap | Get site eap 802.1x setting
[**GetGlobalPortalDomainSetting**](AuthenticationAPI.md#getglobalportaldomainsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/portal/setting/domain | Get Portal Domain
[**GetMacAuthSetting1**](AuthenticationAPI.md#getmacauthsetting1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mac-auth | Get site MAC-Based Authentication info
[**GetMacAuthSsids1**](AuthenticationAPI.md#getmacauthssids1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mac-auth/ssids | Get ssids that support MAC auth
[**GetPortalCandidates**](AuthenticationAPI.md#getportalcandidates) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/portal/candidates | Get portal SSIDs and networks
[**GetPortalCustomization**](AuthenticationAPI.md#getportalcustomization) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/portal/{portalId}/customization | Get portal customization
[**GetPortalDetail**](AuthenticationAPI.md#getportaldetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/portal/{portalId} | Get portal detail
[**GetPortalList**](AuthenticationAPI.md#getportallist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/portals | Get portal list in a site
[**GetPortalListWithLogoutEnabled**](AuthenticationAPI.md#getportallistwithlogoutenabled) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/portals/logout-enabled | Gets a list of logout enabled portals on the site
[**GetSwitchDot1xCandidates**](AuthenticationAPI.md#getswitchdot1xcandidates) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dot1x/candidates | Get site switch 802.1x setting candidates
[**GetSwitchDot1xSetting**](AuthenticationAPI.md#getswitchdot1xsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dot1x | Get site switch 802.1x setting
[**ModifyPortal**](AuthenticationAPI.md#modifyportal) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/portal/{portalId} | Modify portal
[**UpdateEapDot1xSetting**](AuthenticationAPI.md#updateeapdot1xsetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/dot1x/eap | Modify site EAP 802.1x setting
[**UpdateMacAuthSetting1**](AuthenticationAPI.md#updatemacauthsetting1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/mac-auth | Update site MAC-Based Authentication info
[**UpdateSwitchDot1xSetting**](AuthenticationAPI.md#updateswitchdot1xsetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/dot1x | Modify site switch 802.1x setting
[**UploadPortalPage**](AuthenticationAPI.md#uploadportalpage) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/portal/page | Import portal page
[**UploadPortalPic**](AuthenticationAPI.md#uploadportalpic) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/portal/pics | Upload portal picture



## AddPortal

> OperationResponseWithoutResult AddPortal(ctx, omadacId, siteId).PortalSetting(portalSetting).Execute()

Add portal



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
	portalSetting := *openapiclient.NewPortalSetting(*openapiclient.NewAuthTimeoutSetting(), int32(123), false, false, int32(123), "Name_example") // PortalSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AddPortal(context.Background(), omadacId, siteId).PortalSetting(portalSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AddPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPortal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AddPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPortalRequest struct via the builder pattern


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


## DeletePortal

> OperationResponseWithoutResult DeletePortal(ctx, omadacId, siteId, portalId).Execute()

Delete portal



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
	portalId := "portalId_example" // string | Portal ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.DeletePortal(context.Background(), omadacId, siteId, portalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeletePortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.DeletePortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portalId** | **string** | Portal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortalRequest struct via the builder pattern


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


## GetEapDot1xCandidates

> OperationResponseListDot1xEapInfoOpenApiVO GetEapDot1xCandidates(ctx, omadacId, siteId).Execute()

Get site EAP 802.1x setting candidates



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
	resp, r, err := apiClient.AuthenticationAPI.GetEapDot1xCandidates(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetEapDot1xCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEapDot1xCandidates`: OperationResponseListDot1xEapInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetEapDot1xCandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEapDot1xCandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListDot1xEapInfoOpenApiVO**](OperationResponseListDot1xEapInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEapDot1xSetting

> OperationResponseDot1xBasicInfoEapOpenApiVO GetEapDot1xSetting(ctx, omadacId, siteId).Execute()

Get site eap 802.1x setting



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
	resp, r, err := apiClient.AuthenticationAPI.GetEapDot1xSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetEapDot1xSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEapDot1xSetting`: OperationResponseDot1xBasicInfoEapOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetEapDot1xSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEapDot1xSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDot1xBasicInfoEapOpenApiVO**](OperationResponseDot1xBasicInfoEapOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalPortalDomainSetting

> OperationResponseGlobalPortalDomainOpenApiVO GetGlobalPortalDomainSetting(ctx, omadacId, siteId).Execute()

Get Portal Domain



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
	resp, r, err := apiClient.AuthenticationAPI.GetGlobalPortalDomainSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetGlobalPortalDomainSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalPortalDomainSetting`: OperationResponseGlobalPortalDomainOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetGlobalPortalDomainSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalPortalDomainSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseGlobalPortalDomainOpenApiVO**](OperationResponseGlobalPortalDomainOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacAuthSetting1

> OperationResponseMacAuthOpenApiVO GetMacAuthSetting1(ctx, omadacId, siteId).Execute()

Get site MAC-Based Authentication info



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
	resp, r, err := apiClient.AuthenticationAPI.GetMacAuthSetting1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetMacAuthSetting1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacAuthSetting1`: OperationResponseMacAuthOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetMacAuthSetting1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacAuthSetting1Request struct via the builder pattern


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


## GetMacAuthSsids1

> OperationResponseListWlanSimpleOpenApiVO GetMacAuthSsids1(ctx, omadacId, siteId).Execute()

Get ssids that support MAC auth



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
	resp, r, err := apiClient.AuthenticationAPI.GetMacAuthSsids1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetMacAuthSsids1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacAuthSsids1`: OperationResponseListWlanSimpleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetMacAuthSsids1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacAuthSsids1Request struct via the builder pattern


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


## GetPortalCandidates

> OperationResponsePortalCandidatesResOpenApiVO GetPortalCandidates(ctx, omadacId, siteId).PortalCandidatesOpenApiVO(portalCandidatesOpenApiVO).Execute()

Get portal SSIDs and networks



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
	portalCandidatesOpenApiVO := *openapiclient.NewPortalCandidatesOpenApiVO(false) // PortalCandidatesOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetPortalCandidates(context.Background(), omadacId, siteId).PortalCandidatesOpenApiVO(portalCandidatesOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetPortalCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalCandidates`: OperationResponsePortalCandidatesResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetPortalCandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalCandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalCandidatesOpenApiVO** | [**PortalCandidatesOpenApiVO**](PortalCandidatesOpenApiVO.md) |  | 

### Return type

[**OperationResponsePortalCandidatesResOpenApiVO**](OperationResponsePortalCandidatesResOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortalCustomization

> OperationResponsePortalCustomizationPageResOpenApiVO GetPortalCustomization(ctx, omadacId, siteId, portalId).Execute()

Get portal customization



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
	portalId := "portalId_example" // string | Portal ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetPortalCustomization(context.Background(), omadacId, siteId, portalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetPortalCustomization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalCustomization`: OperationResponsePortalCustomizationPageResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetPortalCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portalId** | **string** | Portal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalCustomizationRequest struct via the builder pattern


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


## GetPortalDetail

> OperationResponsePortalDetailResOpenApiVO GetPortalDetail(ctx, omadacId, siteId, portalId).Execute()

Get portal detail



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
	portalId := "portalId_example" // string | Portal ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetPortalDetail(context.Background(), omadacId, siteId, portalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetPortalDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalDetail`: OperationResponsePortalDetailResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetPortalDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portalId** | **string** | Portal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalDetailRequest struct via the builder pattern


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


## GetPortalList

> OperationResponseListPortalResOpenApiVO GetPortalList(ctx, omadacId, siteId).Execute()

Get portal list in a site



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
	resp, r, err := apiClient.AuthenticationAPI.GetPortalList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetPortalList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalList`: OperationResponseListPortalResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetPortalList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalListRequest struct via the builder pattern


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


## GetPortalListWithLogoutEnabled

> OperationResponseListString GetPortalListWithLogoutEnabled(ctx, omadacId, siteId).Execute()

Gets a list of logout enabled portals on the site



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
	resp, r, err := apiClient.AuthenticationAPI.GetPortalListWithLogoutEnabled(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetPortalListWithLogoutEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalListWithLogoutEnabled`: OperationResponseListString
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetPortalListWithLogoutEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalListWithLogoutEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListString**](OperationResponseListString.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchDot1xCandidates

> OperationResponseListDot1xSwitchInfoOpenApiVO GetSwitchDot1xCandidates(ctx, omadacId, siteId).Execute()

Get site switch 802.1x setting candidates



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
	resp, r, err := apiClient.AuthenticationAPI.GetSwitchDot1xCandidates(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetSwitchDot1xCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchDot1xCandidates`: OperationResponseListDot1xSwitchInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetSwitchDot1xCandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchDot1xCandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListDot1xSwitchInfoOpenApiVO**](OperationResponseListDot1xSwitchInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchDot1xSetting

> OperationResponseDot1xSwitchResOpenApiVO GetSwitchDot1xSetting(ctx, omadacId, siteId).Execute()

Get site switch 802.1x setting



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
	resp, r, err := apiClient.AuthenticationAPI.GetSwitchDot1xSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetSwitchDot1xSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchDot1xSetting`: OperationResponseDot1xSwitchResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetSwitchDot1xSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchDot1xSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDot1xSwitchResOpenApiVO**](OperationResponseDot1xSwitchResOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPortal

> OperationResponseWithoutResult ModifyPortal(ctx, omadacId, siteId, portalId).PortalSetting(portalSetting).Execute()

Modify portal



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
	portalId := "portalId_example" // string | Portal ID
	portalSetting := *openapiclient.NewPortalSetting(*openapiclient.NewAuthTimeoutSetting(), int32(123), false, false, int32(123), "Name_example") // PortalSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.ModifyPortal(context.Background(), omadacId, siteId, portalId).PortalSetting(portalSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ModifyPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortal`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ModifyPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**portalId** | **string** | Portal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortalRequest struct via the builder pattern


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


## UpdateEapDot1xSetting

> OperationResponseWithoutResult UpdateEapDot1xSetting(ctx, omadacId, siteId).Dot1xEapOpenApiVO(dot1xEapOpenApiVO).Execute()

Modify site EAP 802.1x setting



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
	dot1xEapOpenApiVO := *openapiclient.NewDot1xEapOpenApiVO(false) // Dot1xEapOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateEapDot1xSetting(context.Background(), omadacId, siteId).Dot1xEapOpenApiVO(dot1xEapOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateEapDot1xSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEapDot1xSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateEapDot1xSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEapDot1xSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dot1xEapOpenApiVO** | [**Dot1xEapOpenApiVO**](Dot1xEapOpenApiVO.md) |  | 

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


## UpdateMacAuthSetting1

> OperationResponseWithoutResult UpdateMacAuthSetting1(ctx, omadacId, siteId).UpdateMacAuthOpenApiVO(updateMacAuthOpenApiVO).Execute()

Update site MAC-Based Authentication info



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
	updateMacAuthOpenApiVO := *openapiclient.NewUpdateMacAuthOpenApiVO(false, []string{"Ssids_example"}) // UpdateMacAuthOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateMacAuthSetting1(context.Background(), omadacId, siteId).UpdateMacAuthOpenApiVO(updateMacAuthOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateMacAuthSetting1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMacAuthSetting1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateMacAuthSetting1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMacAuthSetting1Request struct via the builder pattern


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


## UpdateSwitchDot1xSetting

> OperationResponseWithoutResult UpdateSwitchDot1xSetting(ctx, omadacId, siteId).Dot1xSwitchOpenApiVO(dot1xSwitchOpenApiVO).Execute()

Modify site switch 802.1x setting



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
	dot1xSwitchOpenApiVO := *openapiclient.NewDot1xSwitchOpenApiVO(int32(123), int32(123), false, false, int32(123), "RadiusProfileId_example", false) // Dot1xSwitchOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateSwitchDot1xSetting(context.Background(), omadacId, siteId).Dot1xSwitchOpenApiVO(dot1xSwitchOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateSwitchDot1xSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSwitchDot1xSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateSwitchDot1xSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSwitchDot1xSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dot1xSwitchOpenApiVO** | [**Dot1xSwitchOpenApiVO**](Dot1xSwitchOpenApiVO.md) |  | 

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


## UploadPortalPage

> OperationResponseImportedPortalPageResOpenApiVO UploadPortalPage(ctx, omadacId, siteId).File(file).Execute()

Import portal page



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
	file := os.NewFile(1234, "some_file") // *os.File | Portal page file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UploadPortalPage(context.Background(), omadacId, siteId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UploadPortalPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPortalPage`: OperationResponseImportedPortalPageResOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UploadPortalPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadPortalPageRequest struct via the builder pattern


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


## UploadPortalPic

> OperationResponsePortalPictureInfo UploadPortalPic(ctx, omadacId, siteId).Type_(type_).File(file).Execute()

Upload portal picture



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
	siteId := "siteId_example" // string | Site ID
	file := os.NewFile(1234, "some_file") // *os.File | Portal picture file. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UploadPortalPic(context.Background(), omadacId, siteId).Type_(type_).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UploadPortalPic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPortalPic`: OperationResponsePortalPictureInfo
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UploadPortalPic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadPortalPicRequest struct via the builder pattern


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


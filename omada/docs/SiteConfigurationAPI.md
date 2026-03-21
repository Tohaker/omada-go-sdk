# \SiteConfigurationAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBandSteeringSetting**](SiteConfigurationAPI.md#GetBandSteeringSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/band-steering | Get site band steering setting
[**GetBeaconControlSetting**](SiteConfigurationAPI.md#GetBeaconControlSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/beacon-control | Get site beacon control setting
[**GetChannelLimitSetting**](SiteConfigurationAPI.md#GetChannelLimitSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/channel-limit | Get site channel limit setting
[**GetDisasterModeStatus**](SiteConfigurationAPI.md#GetDisasterModeStatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/disaster | Get the status of disaster modes in Japan
[**GetExistSiteSettingByOpenApi**](SiteConfigurationAPI.md#GetExistSiteSettingByOpenApi) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/exist | Query site setting exist or not
[**GetLldpSetting**](SiteConfigurationAPI.md#GetLldpSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/lldp | Get site lldp setting
[**GetMeshSetting**](SiteConfigurationAPI.md#GetMeshSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/mesh | Get site mesh setting
[**GetMulticastRateLimitByOpenApi**](SiteConfigurationAPI.md#GetMulticastRateLimitByOpenApi) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/multicast-rate-limit | Get site multicast rate limit setting
[**GetOmadacDstInfo**](SiteConfigurationAPI.md#GetOmadacDstInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dst-info | Get site DST
[**GetPortalLogoutDomainSetting**](SiteConfigurationAPI.md#GetPortalLogoutDomainSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/portal-logout | Get portal logout domain setting
[**GetRemoteLoggingSetting**](SiteConfigurationAPI.md#GetRemoteLoggingSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/remote-logging | Get site remote logging setting
[**GetRoamingSetting**](SiteConfigurationAPI.md#GetRoamingSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/roaming | Get site roaming setting
[**GetSiteLedSetting**](SiteConfigurationAPI.md#GetSiteLedSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/led | Get site led setting
[**GetSiteRememberSettingByOpenApi**](SiteConfigurationAPI.md#GetSiteRememberSettingByOpenApi) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/remember-device | Get site remember device setting
[**GetSiteSettingCap**](SiteConfigurationAPI.md#GetSiteSettingCap) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/capacity | Get site capabilities.
[**GetSiteSpecification**](SiteConfigurationAPI.md#GetSiteSpecification) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/specification | Get site functional specifications
[**GetSiteSupportSwitchInfo**](SiteConfigurationAPI.md#GetSiteSupportSwitchInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switch-support | Get supported switch types information for the site
[**ModifyDisasterModeStatus**](SiteConfigurationAPI.md#ModifyDisasterModeStatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/disaster | Modify the status of disaster modes in Japan
[**ModifySiteRememberSettingByOpenApi**](SiteConfigurationAPI.md#ModifySiteRememberSettingByOpenApi) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/remember-device | Modify site remember device setting
[**UpdateBandSteeringSetting**](SiteConfigurationAPI.md#UpdateBandSteeringSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/band-steering | Modify site band steering setting
[**UpdateBeaconControlSetting**](SiteConfigurationAPI.md#UpdateBeaconControlSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/beacon-control | Modify site beacon control setting
[**UpdateChannelLimitSetting**](SiteConfigurationAPI.md#UpdateChannelLimitSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/channel-limit | Modify site channel limit setting
[**UpdateLldpSetting**](SiteConfigurationAPI.md#UpdateLldpSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/lldp | Modify site lldp setting
[**UpdateMcastRateLimitSettingByOpenApi**](SiteConfigurationAPI.md#UpdateMcastRateLimitSettingByOpenApi) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/multicast-rate-limit | Modify site multicast rate limit setting
[**UpdateMeshSetting**](SiteConfigurationAPI.md#UpdateMeshSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/mesh | Modify site mesh setting
[**UpdatePortalLogoutDomainSetting**](SiteConfigurationAPI.md#UpdatePortalLogoutDomainSetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/portal-logout | Modify portal logout domain setting
[**UpdateRemoteLoggingSetting**](SiteConfigurationAPI.md#UpdateRemoteLoggingSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/remote-logging | Modify site remote logging setting
[**UpdateRoamingSetting**](SiteConfigurationAPI.md#UpdateRoamingSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/roaming | Modify site roaming setting
[**UpdateSiteLedSetting**](SiteConfigurationAPI.md#UpdateSiteLedSetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/led | Modify site led setting



## GetBandSteeringSetting

> OperationResponseSiteBandSteeringSetting GetBandSteeringSetting(ctx, omadacId, siteId).Execute()

Get site band steering setting



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetBandSteeringSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetBandSteeringSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandSteeringSetting`: OperationResponseSiteBandSteeringSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetBandSteeringSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandSteeringSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteBandSteeringSetting**](OperationResponseSiteBandSteeringSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBeaconControlSetting

> OperationResponseSiteBeaconControlSetting GetBeaconControlSetting(ctx, omadacId, siteId).Execute()

Get site beacon control setting



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetBeaconControlSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetBeaconControlSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBeaconControlSetting`: OperationResponseSiteBeaconControlSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetBeaconControlSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBeaconControlSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteBeaconControlSetting**](OperationResponseSiteBeaconControlSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelLimitSetting

> OperationResponseSiteChannelLimitSetting GetChannelLimitSetting(ctx, omadacId, siteId).Execute()

Get site channel limit setting



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetChannelLimitSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetChannelLimitSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelLimitSetting`: OperationResponseSiteChannelLimitSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetChannelLimitSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelLimitSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteChannelLimitSetting**](OperationResponseSiteChannelLimitSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDisasterModeStatus

> OperationResponseDisasterModeOpenApiVO GetDisasterModeStatus(ctx, omadacId, siteId).Execute()

Get the status of disaster modes in Japan



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
	resp, r, err := apiClient.SiteConfigurationAPI.GetDisasterModeStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetDisasterModeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisasterModeStatus`: OperationResponseDisasterModeOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetDisasterModeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisasterModeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDisasterModeOpenApiVO**](OperationResponseDisasterModeOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExistSiteSettingByOpenApi

> OperationResponseExistSiteSettingOpenApiVO GetExistSiteSettingByOpenApi(ctx, omadacId, siteId).Execute()

Query site setting exist or not



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
	resp, r, err := apiClient.SiteConfigurationAPI.GetExistSiteSettingByOpenApi(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetExistSiteSettingByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExistSiteSettingByOpenApi`: OperationResponseExistSiteSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetExistSiteSettingByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExistSiteSettingByOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseExistSiteSettingOpenApiVO**](OperationResponseExistSiteSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLldpSetting

> OperationResponseSiteLldpSetting GetLldpSetting(ctx, omadacId, siteId).Execute()

Get site lldp setting



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetLldpSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetLldpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLldpSetting`: OperationResponseSiteLldpSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetLldpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLldpSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteLldpSetting**](OperationResponseSiteLldpSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeshSetting

> OperationResponseSiteMeshSetting GetMeshSetting(ctx, omadacId, siteId).Execute()

Get site mesh setting



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetMeshSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetMeshSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeshSetting`: OperationResponseSiteMeshSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetMeshSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeshSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteMeshSetting**](OperationResponseSiteMeshSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMulticastRateLimitByOpenApi

> OperationResponseSiteMulticastRateLimitSetting GetMulticastRateLimitByOpenApi(ctx, omadacId, siteId).Execute()

Get site multicast rate limit setting



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
	siteId := "siteId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetMulticastRateLimitByOpenApi(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetMulticastRateLimitByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMulticastRateLimitByOpenApi`: OperationResponseSiteMulticastRateLimitSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetMulticastRateLimitByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMulticastRateLimitByOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteMulticastRateLimitSetting**](OperationResponseSiteMulticastRateLimitSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOmadacDstInfo

> OperationResponseDstOpenApiVO GetOmadacDstInfo(ctx, omadacId, siteId).Execute()

Get site DST



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
	resp, r, err := apiClient.SiteConfigurationAPI.GetOmadacDstInfo(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetOmadacDstInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOmadacDstInfo`: OperationResponseDstOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetOmadacDstInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOmadacDstInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDstOpenApiVO**](OperationResponseDstOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortalLogoutDomainSetting

> OperationResponsePortalLogoutConfigOpenApiVO GetPortalLogoutDomainSetting(ctx, omadacId, siteId).Execute()

Get portal logout domain setting



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
	resp, r, err := apiClient.SiteConfigurationAPI.GetPortalLogoutDomainSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetPortalLogoutDomainSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalLogoutDomainSetting`: OperationResponsePortalLogoutConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetPortalLogoutDomainSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalLogoutDomainSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponsePortalLogoutConfigOpenApiVO**](OperationResponsePortalLogoutConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteLoggingSetting

> OperationResponseSiteRemoteLoggingSetting GetRemoteLoggingSetting(ctx, omadacId, siteId).Execute()

Get site remote logging setting



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetRemoteLoggingSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetRemoteLoggingSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteLoggingSetting`: OperationResponseSiteRemoteLoggingSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetRemoteLoggingSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteLoggingSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteRemoteLoggingSetting**](OperationResponseSiteRemoteLoggingSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoamingSetting

> OperationResponseSiteRoamingSetting GetRoamingSetting(ctx, omadacId, siteId).Execute()

Get site roaming setting



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetRoamingSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetRoamingSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoamingSetting`: OperationResponseSiteRoamingSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetRoamingSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoamingSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteRoamingSetting**](OperationResponseSiteRoamingSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteLedSetting

> OperationResponseSiteLedSetting GetSiteLedSetting(ctx, omadacId, siteId).Execute()

Get site led setting



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
	siteId := "siteId_example" // string | site_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.GetSiteLedSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetSiteLedSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteLedSetting`: OperationResponseSiteLedSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetSiteLedSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteLedSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteLedSetting**](OperationResponseSiteLedSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteRememberSettingByOpenApi

> OperationResponseSiteRememberDeviceSetting GetSiteRememberSettingByOpenApi(ctx, omadacId, siteId).Execute()

Get site remember device setting



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
	resp, r, err := apiClient.SiteConfigurationAPI.GetSiteRememberSettingByOpenApi(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetSiteRememberSettingByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteRememberSettingByOpenApi`: OperationResponseSiteRememberDeviceSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetSiteRememberSettingByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRememberSettingByOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteRememberDeviceSetting**](OperationResponseSiteRememberDeviceSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSettingCap

> OperationResponseSiteSettingCapOpenApiVO GetSiteSettingCap(ctx, omadacId, siteId).Execute()

Get site capabilities.



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
	resp, r, err := apiClient.SiteConfigurationAPI.GetSiteSettingCap(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetSiteSettingCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSettingCap`: OperationResponseSiteSettingCapOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetSiteSettingCap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSettingCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteSettingCapOpenApiVO**](OperationResponseSiteSettingCapOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSpecification

> OperationResponseSpecificationOpenApiVO GetSiteSpecification(ctx, omadacId, siteId).Execute()

Get site functional specifications



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
	resp, r, err := apiClient.SiteConfigurationAPI.GetSiteSpecification(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetSiteSpecification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSpecification`: OperationResponseSpecificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetSiteSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSpecificationOpenApiVO**](OperationResponseSpecificationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSupportSwitchInfo

> OperationResponseSiteSupportSwitchOpenApiVO GetSiteSupportSwitchInfo(ctx, omadacId, siteId).Execute()

Get supported switch types information for the site



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
	resp, r, err := apiClient.SiteConfigurationAPI.GetSiteSupportSwitchInfo(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.GetSiteSupportSwitchInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSupportSwitchInfo`: OperationResponseSiteSupportSwitchOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.GetSiteSupportSwitchInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSupportSwitchInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSiteSupportSwitchOpenApiVO**](OperationResponseSiteSupportSwitchOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDisasterModeStatus

> OperationResponseWithoutResult ModifyDisasterModeStatus(ctx, omadacId, siteId).ModifyDisasterModeOpenApiVO(modifyDisasterModeOpenApiVO).Execute()

Modify the status of disaster modes in Japan



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
	modifyDisasterModeOpenApiVO := *openapiclient.NewModifyDisasterModeOpenApiVO() // ModifyDisasterModeOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.ModifyDisasterModeStatus(context.Background(), omadacId, siteId).ModifyDisasterModeOpenApiVO(modifyDisasterModeOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.ModifyDisasterModeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDisasterModeStatus`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.ModifyDisasterModeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDisasterModeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyDisasterModeOpenApiVO** | [**ModifyDisasterModeOpenApiVO**](ModifyDisasterModeOpenApiVO.md) |  | 

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


## ModifySiteRememberSettingByOpenApi

> OperationResponseWithoutResult ModifySiteRememberSettingByOpenApi(ctx, omadacId, siteId).SiteRememberDeviceSetting(siteRememberDeviceSetting).Execute()

Modify site remember device setting



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
	siteRememberDeviceSetting := *openapiclient.NewSiteRememberDeviceSetting() // SiteRememberDeviceSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.ModifySiteRememberSettingByOpenApi(context.Background(), omadacId, siteId).SiteRememberDeviceSetting(siteRememberDeviceSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.ModifySiteRememberSettingByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySiteRememberSettingByOpenApi`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.ModifySiteRememberSettingByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySiteRememberSettingByOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteRememberDeviceSetting** | [**SiteRememberDeviceSetting**](SiteRememberDeviceSetting.md) |  | 

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


## UpdateBandSteeringSetting

> OperationResponseWithoutResult UpdateBandSteeringSetting(ctx, omadacId, siteId).SiteBandSteeringSetting(siteBandSteeringSetting).Execute()

Modify site band steering setting



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
	siteBandSteeringSetting := *openapiclient.NewSiteBandSteeringSetting() // SiteBandSteeringSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateBandSteeringSetting(context.Background(), omadacId, siteId).SiteBandSteeringSetting(siteBandSteeringSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateBandSteeringSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBandSteeringSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateBandSteeringSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBandSteeringSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteBandSteeringSetting** | [**SiteBandSteeringSetting**](SiteBandSteeringSetting.md) |  | 

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


## UpdateBeaconControlSetting

> OperationResponseWithoutResult UpdateBeaconControlSetting(ctx, omadacId, siteId).SiteBeaconControlSetting(siteBeaconControlSetting).Execute()

Modify site beacon control setting



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
	siteBeaconControlSetting := *openapiclient.NewSiteBeaconControlSetting() // SiteBeaconControlSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateBeaconControlSetting(context.Background(), omadacId, siteId).SiteBeaconControlSetting(siteBeaconControlSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateBeaconControlSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBeaconControlSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateBeaconControlSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBeaconControlSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteBeaconControlSetting** | [**SiteBeaconControlSetting**](SiteBeaconControlSetting.md) |  | 

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


## UpdateChannelLimitSetting

> OperationResponseWithoutResult UpdateChannelLimitSetting(ctx, omadacId, siteId).SiteChannelLimitSetting(siteChannelLimitSetting).Execute()

Modify site channel limit setting



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
	siteId := "siteId_example" // string | site_id
	siteChannelLimitSetting := *openapiclient.NewSiteChannelLimitSetting() // SiteChannelLimitSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateChannelLimitSetting(context.Background(), omadacId, siteId).SiteChannelLimitSetting(siteChannelLimitSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateChannelLimitSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChannelLimitSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateChannelLimitSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelLimitSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteChannelLimitSetting** | [**SiteChannelLimitSetting**](SiteChannelLimitSetting.md) |  | 

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


## UpdateLldpSetting

> OperationResponseWithoutResult UpdateLldpSetting(ctx, omadacId, siteId).SiteLldpSetting(siteLldpSetting).Execute()

Modify site lldp setting



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
	siteId := "siteId_example" // string | site_id
	siteLldpSetting := *openapiclient.NewSiteLldpSetting() // SiteLldpSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateLldpSetting(context.Background(), omadacId, siteId).SiteLldpSetting(siteLldpSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateLldpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLldpSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateLldpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLldpSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteLldpSetting** | [**SiteLldpSetting**](SiteLldpSetting.md) |  | 

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


## UpdateMcastRateLimitSettingByOpenApi

> OperationResponseWithoutResult UpdateMcastRateLimitSettingByOpenApi(ctx, omadacId, siteId).SiteMulticastRateLimitSetting(siteMulticastRateLimitSetting).Execute()

Modify site multicast rate limit setting



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
	siteMulticastRateLimitSetting := *openapiclient.NewSiteMulticastRateLimitSetting() // SiteMulticastRateLimitSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateMcastRateLimitSettingByOpenApi(context.Background(), omadacId, siteId).SiteMulticastRateLimitSetting(siteMulticastRateLimitSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateMcastRateLimitSettingByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMcastRateLimitSettingByOpenApi`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateMcastRateLimitSettingByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMcastRateLimitSettingByOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteMulticastRateLimitSetting** | [**SiteMulticastRateLimitSetting**](SiteMulticastRateLimitSetting.md) |  | 

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


## UpdateMeshSetting

> OperationResponseWithoutResult UpdateMeshSetting(ctx, omadacId, siteId).SiteMeshSetting(siteMeshSetting).Execute()

Modify site mesh setting



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
	siteId := "siteId_example" // string | site_id
	siteMeshSetting := *openapiclient.NewSiteMeshSetting() // SiteMeshSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateMeshSetting(context.Background(), omadacId, siteId).SiteMeshSetting(siteMeshSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateMeshSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMeshSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateMeshSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMeshSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteMeshSetting** | [**SiteMeshSetting**](SiteMeshSetting.md) |  | 

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


## UpdatePortalLogoutDomainSetting

> OperationResponseWithoutResult UpdatePortalLogoutDomainSetting(ctx, omadacId, siteId).PortalLogoutConfigOpenApiVO(portalLogoutConfigOpenApiVO).Execute()

Modify portal logout domain setting



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
	portalLogoutConfigOpenApiVO := *openapiclient.NewPortalLogoutConfigOpenApiVO() // PortalLogoutConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdatePortalLogoutDomainSetting(context.Background(), omadacId, siteId).PortalLogoutConfigOpenApiVO(portalLogoutConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdatePortalLogoutDomainSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePortalLogoutDomainSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdatePortalLogoutDomainSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortalLogoutDomainSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalLogoutConfigOpenApiVO** | [**PortalLogoutConfigOpenApiVO**](PortalLogoutConfigOpenApiVO.md) |  | 

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


## UpdateRemoteLoggingSetting

> OperationResponseWithoutResult UpdateRemoteLoggingSetting(ctx, omadacId, siteId).SiteRemoteLoggingSetting(siteRemoteLoggingSetting).Execute()

Modify site remote logging setting



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
	siteRemoteLoggingSetting := *openapiclient.NewSiteRemoteLoggingSetting() // SiteRemoteLoggingSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateRemoteLoggingSetting(context.Background(), omadacId, siteId).SiteRemoteLoggingSetting(siteRemoteLoggingSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateRemoteLoggingSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRemoteLoggingSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateRemoteLoggingSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteLoggingSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteRemoteLoggingSetting** | [**SiteRemoteLoggingSetting**](SiteRemoteLoggingSetting.md) |  | 

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


## UpdateRoamingSetting

> OperationResponseWithoutResult UpdateRoamingSetting(ctx, omadacId, siteId).SiteRoamingSetting(siteRoamingSetting).Execute()

Modify site roaming setting



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
	siteRoamingSetting := *openapiclient.NewSiteRoamingSetting(*openapiclient.NewRoamingOpenApiVO(false, false)) // SiteRoamingSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateRoamingSetting(context.Background(), omadacId, siteId).SiteRoamingSetting(siteRoamingSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateRoamingSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoamingSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateRoamingSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoamingSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteRoamingSetting** | [**SiteRoamingSetting**](SiteRoamingSetting.md) |  | 

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


## UpdateSiteLedSetting

> OperationResponseWithoutResult UpdateSiteLedSetting(ctx, omadacId, siteId).SiteLedSetting(siteLedSetting).Execute()

Modify site led setting



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
	siteId := "siteId_example" // string | site_id
	siteLedSetting := *openapiclient.NewSiteLedSetting(false) // SiteLedSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteConfigurationAPI.UpdateSiteLedSetting(context.Background(), omadacId, siteId).SiteLedSetting(siteLedSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteConfigurationAPI.UpdateSiteLedSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteLedSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteConfigurationAPI.UpdateSiteLedSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | site_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteLedSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteLedSetting** | [**SiteLedSetting**](SiteLedSetting.md) |  | 

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


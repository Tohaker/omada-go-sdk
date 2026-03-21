# \SiteTemplateConfigurationAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBandSteeringTemplateSetting**](SiteTemplateConfigurationAPI.md#GetBandSteeringTemplateSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/band-steering | Get site template band steering setting
[**GetBeaconControlTemplateSetting**](SiteTemplateConfigurationAPI.md#GetBeaconControlTemplateSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/beacon-control | Get site template beacon control setting
[**GetChannelLimitTemplateSetting**](SiteTemplateConfigurationAPI.md#GetChannelLimitTemplateSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/channel-limit | Get site template channel limit setting
[**GetExistSiteTemplateSettingByOpenApi**](SiteTemplateConfigurationAPI.md#GetExistSiteTemplateSettingByOpenApi) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/exist | Query site template setting exist or not
[**GetLldpTemplateSetting**](SiteTemplateConfigurationAPI.md#GetLldpTemplateSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lldp | Get site template lldp setting
[**GetMeshTemplateSetting**](SiteTemplateConfigurationAPI.md#GetMeshTemplateSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mesh | Get site template mesh setting
[**GetMulticastRateLimitTemplateByOpenApi**](SiteTemplateConfigurationAPI.md#GetMulticastRateLimitTemplateByOpenApi) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/multicast-rate-limit | Get site template multicast rate limit setting
[**GetRemoteLoggingTemplateSetting**](SiteTemplateConfigurationAPI.md#GetRemoteLoggingTemplateSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/remote-logging | Get site template remote logging setting
[**GetRoamingTemplateSetting**](SiteTemplateConfigurationAPI.md#GetRoamingTemplateSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/roaming | Get site template roaming setting
[**GetSiteTemplateLedSetting**](SiteTemplateConfigurationAPI.md#GetSiteTemplateLedSetting) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/led | Get site template led setting
[**GetSiteTemplateSettingCap**](SiteTemplateConfigurationAPI.md#GetSiteTemplateSettingCap) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/capacity | Get siteTemplate capabilities.
[**GetSiteTemplateSpecification**](SiteTemplateConfigurationAPI.md#GetSiteTemplateSpecification) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/specification | Get siteTemplate functional specifications
[**ModifySiteConfigurationTemplate**](SiteTemplateConfigurationAPI.md#ModifySiteConfigurationTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/configuration | Modify Site template configuration settings
[**ModifySiteServiceTemplate**](SiteTemplateConfigurationAPI.md#ModifySiteServiceTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/service | Modify Site template service settings
[**ModifySiteTemplateGeneralConfig**](SiteTemplateConfigurationAPI.md#ModifySiteTemplateGeneralConfig) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/general/config | Modify the general configuration settings of the Site template
[**ModifySiteTemplateSetting**](SiteTemplateConfigurationAPI.md#ModifySiteTemplateSetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting | Modify Site template settings
[**UpdateBandSteeringTemplateSetting**](SiteTemplateConfigurationAPI.md#UpdateBandSteeringTemplateSetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/band-steering | Modify site template band steering setting
[**UpdateBeaconControlTemplateSetting**](SiteTemplateConfigurationAPI.md#UpdateBeaconControlTemplateSetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/beacon-control | Modify site template beacon control setting
[**UpdateChannelLimitTemplateSetting**](SiteTemplateConfigurationAPI.md#UpdateChannelLimitTemplateSetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/channel-limit | Modify site template channel limit setting
[**UpdateLldpTemplateSetting**](SiteTemplateConfigurationAPI.md#UpdateLldpTemplateSetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/lldp | Modify site template lldp setting
[**UpdateMcastRateLimitSettingTemplateByOpenApi**](SiteTemplateConfigurationAPI.md#UpdateMcastRateLimitSettingTemplateByOpenApi) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/multicast-rate-limit | Modify site template multicast rate limit setting
[**UpdateMeshTemplateSetting**](SiteTemplateConfigurationAPI.md#UpdateMeshTemplateSetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/mesh | Modify site template mesh setting
[**UpdateRemoteLoggingTemplateSetting**](SiteTemplateConfigurationAPI.md#UpdateRemoteLoggingTemplateSetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/remote-logging | Modify site template remote logging setting
[**UpdateRoamingTemplateSetting**](SiteTemplateConfigurationAPI.md#UpdateRoamingTemplateSetting) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/roaming | Modify site template roaming setting
[**UpdateSiteTemplateLedSetting**](SiteTemplateConfigurationAPI.md#UpdateSiteTemplateLedSetting) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/led | Modify site template led setting



## GetBandSteeringTemplateSetting

> OperationResponseSiteBandSteeringSetting GetBandSteeringTemplateSetting(ctx, omadacId, siteTemplateId).Execute()

Get site template band steering setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetBandSteeringTemplateSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetBandSteeringTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBandSteeringTemplateSetting`: OperationResponseSiteBandSteeringSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetBandSteeringTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandSteeringTemplateSettingRequest struct via the builder pattern


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


## GetBeaconControlTemplateSetting

> OperationResponseSiteBeaconControlSetting GetBeaconControlTemplateSetting(ctx, omadacId, siteTemplateId).Execute()

Get site template beacon control setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetBeaconControlTemplateSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetBeaconControlTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBeaconControlTemplateSetting`: OperationResponseSiteBeaconControlSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetBeaconControlTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBeaconControlTemplateSettingRequest struct via the builder pattern


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


## GetChannelLimitTemplateSetting

> OperationResponseSiteChannelLimitSetting GetChannelLimitTemplateSetting(ctx, omadacId, siteTemplateId).Execute()

Get site template channel limit setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetChannelLimitTemplateSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetChannelLimitTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelLimitTemplateSetting`: OperationResponseSiteChannelLimitSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetChannelLimitTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelLimitTemplateSettingRequest struct via the builder pattern


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


## GetExistSiteTemplateSettingByOpenApi

> OperationResponseExistSiteSettingOpenApiVO GetExistSiteTemplateSettingByOpenApi(ctx, omadacId, siteTemplateId).Execute()

Query site template setting exist or not



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetExistSiteTemplateSettingByOpenApi(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetExistSiteTemplateSettingByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExistSiteTemplateSettingByOpenApi`: OperationResponseExistSiteSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetExistSiteTemplateSettingByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExistSiteTemplateSettingByOpenApiRequest struct via the builder pattern


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


## GetLldpTemplateSetting

> OperationResponseSiteLldpSetting GetLldpTemplateSetting(ctx, omadacId, siteTemplateId).Execute()

Get site template lldp setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetLldpTemplateSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetLldpTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLldpTemplateSetting`: OperationResponseSiteLldpSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetLldpTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLldpTemplateSettingRequest struct via the builder pattern


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


## GetMeshTemplateSetting

> OperationResponseSiteMeshSetting GetMeshTemplateSetting(ctx, omadacId, siteTemplateId).Execute()

Get site template mesh setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetMeshTemplateSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetMeshTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeshTemplateSetting`: OperationResponseSiteMeshSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetMeshTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeshTemplateSettingRequest struct via the builder pattern


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


## GetMulticastRateLimitTemplateByOpenApi

> OperationResponseSiteMulticastRateLimitSetting GetMulticastRateLimitTemplateByOpenApi(ctx, omadacId, siteTemplateId).Execute()

Get site template multicast rate limit setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetMulticastRateLimitTemplateByOpenApi(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetMulticastRateLimitTemplateByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMulticastRateLimitTemplateByOpenApi`: OperationResponseSiteMulticastRateLimitSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetMulticastRateLimitTemplateByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMulticastRateLimitTemplateByOpenApiRequest struct via the builder pattern


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


## GetRemoteLoggingTemplateSetting

> OperationResponseSiteRemoteLoggingSetting GetRemoteLoggingTemplateSetting(ctx, omadacId, siteTemplateId).Execute()

Get site template remote logging setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetRemoteLoggingTemplateSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetRemoteLoggingTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteLoggingTemplateSetting`: OperationResponseSiteRemoteLoggingSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetRemoteLoggingTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteLoggingTemplateSettingRequest struct via the builder pattern


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


## GetRoamingTemplateSetting

> OperationResponseSiteRoamingSetting GetRoamingTemplateSetting(ctx, omadacId, siteTemplateId).Execute()

Get site template roaming setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetRoamingTemplateSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetRoamingTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoamingTemplateSetting`: OperationResponseSiteRoamingSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetRoamingTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoamingTemplateSettingRequest struct via the builder pattern


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


## GetSiteTemplateLedSetting

> OperationResponseSiteLedSetting GetSiteTemplateLedSetting(ctx, omadacId, siteTemplateId).Execute()

Get site template led setting



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetSiteTemplateLedSetting(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetSiteTemplateLedSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateLedSetting`: OperationResponseSiteLedSetting
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetSiteTemplateLedSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateLedSettingRequest struct via the builder pattern


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


## GetSiteTemplateSettingCap

> OperationResponseSiteSettingCapOpenApiVO GetSiteTemplateSettingCap(ctx, omadacId, siteTemplateId).Execute()

Get siteTemplate capabilities.



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetSiteTemplateSettingCap(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetSiteTemplateSettingCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateSettingCap`: OperationResponseSiteSettingCapOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetSiteTemplateSettingCap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateSettingCapRequest struct via the builder pattern


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


## GetSiteTemplateSpecification

> OperationResponseSpecificationOpenApiVO GetSiteTemplateSpecification(ctx, omadacId, siteTemplateId).Execute()

Get siteTemplate functional specifications



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
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.GetSiteTemplateSpecification(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.GetSiteTemplateSpecification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteTemplateSpecification`: OperationResponseSpecificationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.GetSiteTemplateSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTemplateSpecificationRequest struct via the builder pattern


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


## ModifySiteConfigurationTemplate

> OperationResponseWithoutResult ModifySiteConfigurationTemplate(ctx, omadacId, siteTemplateId).SiteTemplateOpenApiVO(siteTemplateOpenApiVO).Execute()

Modify Site template configuration settings



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
	siteTemplateOpenApiVO := *openapiclient.NewSiteTemplateOpenApiVO() // SiteTemplateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.ModifySiteConfigurationTemplate(context.Background(), omadacId, siteTemplateId).SiteTemplateOpenApiVO(siteTemplateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.ModifySiteConfigurationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySiteConfigurationTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.ModifySiteConfigurationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySiteConfigurationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteTemplateOpenApiVO** | [**SiteTemplateOpenApiVO**](SiteTemplateOpenApiVO.md) |  | 

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


## ModifySiteServiceTemplate

> OperationResponseWithoutResult ModifySiteServiceTemplate(ctx, omadacId, siteTemplateId).ModifySiteServiceOpenApiVO(modifySiteServiceOpenApiVO).Execute()

Modify Site template service settings



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
	modifySiteServiceOpenApiVO := *openapiclient.NewModifySiteServiceOpenApiVO() // ModifySiteServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.ModifySiteServiceTemplate(context.Background(), omadacId, siteTemplateId).ModifySiteServiceOpenApiVO(modifySiteServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.ModifySiteServiceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySiteServiceTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.ModifySiteServiceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySiteServiceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifySiteServiceOpenApiVO** | [**ModifySiteServiceOpenApiVO**](ModifySiteServiceOpenApiVO.md) |  | 

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


## ModifySiteTemplateGeneralConfig

> OperationResponseWithoutResult ModifySiteTemplateGeneralConfig(ctx, omadacId, siteTemplateId).SiteServiceGeneralConfigOpenApiVO(siteServiceGeneralConfigOpenApiVO).Execute()

Modify the general configuration settings of the Site template



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
	siteServiceGeneralConfigOpenApiVO := *openapiclient.NewSiteServiceGeneralConfigOpenApiVO() // SiteServiceGeneralConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.ModifySiteTemplateGeneralConfig(context.Background(), omadacId, siteTemplateId).SiteServiceGeneralConfigOpenApiVO(siteServiceGeneralConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.ModifySiteTemplateGeneralConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySiteTemplateGeneralConfig`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.ModifySiteTemplateGeneralConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySiteTemplateGeneralConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteServiceGeneralConfigOpenApiVO** | [**SiteServiceGeneralConfigOpenApiVO**](SiteServiceGeneralConfigOpenApiVO.md) |  | 

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


## ModifySiteTemplateSetting

> OperationResponseWithoutResult ModifySiteTemplateSetting(ctx, omadacId, siteTemplateId).SiteTemplateSettingOpenApiVO(siteTemplateSettingOpenApiVO).Execute()

Modify Site template settings



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
	siteTemplateSettingOpenApiVO := *openapiclient.NewSiteTemplateSettingOpenApiVO() // SiteTemplateSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.ModifySiteTemplateSetting(context.Background(), omadacId, siteTemplateId).SiteTemplateSettingOpenApiVO(siteTemplateSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.ModifySiteTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySiteTemplateSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.ModifySiteTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySiteTemplateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteTemplateSettingOpenApiVO** | [**SiteTemplateSettingOpenApiVO**](SiteTemplateSettingOpenApiVO.md) |  | 

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


## UpdateBandSteeringTemplateSetting

> OperationResponseWithoutResult UpdateBandSteeringTemplateSetting(ctx, omadacId, siteTemplateId).SiteBandSteeringSetting(siteBandSteeringSetting).Execute()

Modify site template band steering setting



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
	siteBandSteeringSetting := *openapiclient.NewSiteBandSteeringSetting() // SiteBandSteeringSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateBandSteeringTemplateSetting(context.Background(), omadacId, siteTemplateId).SiteBandSteeringSetting(siteBandSteeringSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateBandSteeringTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBandSteeringTemplateSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateBandSteeringTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBandSteeringTemplateSettingRequest struct via the builder pattern


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


## UpdateBeaconControlTemplateSetting

> OperationResponseWithoutResult UpdateBeaconControlTemplateSetting(ctx, omadacId, siteTemplateId).SiteBeaconControlSetting(siteBeaconControlSetting).Execute()

Modify site template beacon control setting



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
	siteBeaconControlSetting := *openapiclient.NewSiteBeaconControlSetting() // SiteBeaconControlSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateBeaconControlTemplateSetting(context.Background(), omadacId, siteTemplateId).SiteBeaconControlSetting(siteBeaconControlSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateBeaconControlTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBeaconControlTemplateSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateBeaconControlTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBeaconControlTemplateSettingRequest struct via the builder pattern


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


## UpdateChannelLimitTemplateSetting

> OperationResponseWithoutResult UpdateChannelLimitTemplateSetting(ctx, omadacId, siteTemplateId).SiteChannelLimitSetting(siteChannelLimitSetting).Execute()

Modify site template channel limit setting



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
	siteChannelLimitSetting := *openapiclient.NewSiteChannelLimitSetting() // SiteChannelLimitSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateChannelLimitTemplateSetting(context.Background(), omadacId, siteTemplateId).SiteChannelLimitSetting(siteChannelLimitSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateChannelLimitTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChannelLimitTemplateSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateChannelLimitTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelLimitTemplateSettingRequest struct via the builder pattern


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


## UpdateLldpTemplateSetting

> OperationResponseWithoutResult UpdateLldpTemplateSetting(ctx, omadacId, siteTemplateId).SiteLldpSetting(siteLldpSetting).Execute()

Modify site template lldp setting



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
	siteLldpSetting := *openapiclient.NewSiteLldpSetting() // SiteLldpSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateLldpTemplateSetting(context.Background(), omadacId, siteTemplateId).SiteLldpSetting(siteLldpSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateLldpTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLldpTemplateSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateLldpTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLldpTemplateSettingRequest struct via the builder pattern


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


## UpdateMcastRateLimitSettingTemplateByOpenApi

> OperationResponseWithoutResult UpdateMcastRateLimitSettingTemplateByOpenApi(ctx, omadacId, siteTemplateId).SiteMulticastRateLimitSetting(siteMulticastRateLimitSetting).Execute()

Modify site template multicast rate limit setting



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
	siteMulticastRateLimitSetting := *openapiclient.NewSiteMulticastRateLimitSetting() // SiteMulticastRateLimitSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateMcastRateLimitSettingTemplateByOpenApi(context.Background(), omadacId, siteTemplateId).SiteMulticastRateLimitSetting(siteMulticastRateLimitSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateMcastRateLimitSettingTemplateByOpenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMcastRateLimitSettingTemplateByOpenApi`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateMcastRateLimitSettingTemplateByOpenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMcastRateLimitSettingTemplateByOpenApiRequest struct via the builder pattern


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


## UpdateMeshTemplateSetting

> OperationResponseWithoutResult UpdateMeshTemplateSetting(ctx, omadacId, siteTemplateId).SiteMeshSetting(siteMeshSetting).Execute()

Modify site template mesh setting



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
	siteMeshSetting := *openapiclient.NewSiteMeshSetting() // SiteMeshSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateMeshTemplateSetting(context.Background(), omadacId, siteTemplateId).SiteMeshSetting(siteMeshSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateMeshTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMeshTemplateSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateMeshTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMeshTemplateSettingRequest struct via the builder pattern


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


## UpdateRemoteLoggingTemplateSetting

> OperationResponseWithoutResult UpdateRemoteLoggingTemplateSetting(ctx, omadacId, siteTemplateId).SiteRemoteLoggingSetting(siteRemoteLoggingSetting).Execute()

Modify site template remote logging setting



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
	siteRemoteLoggingSetting := *openapiclient.NewSiteRemoteLoggingSetting() // SiteRemoteLoggingSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateRemoteLoggingTemplateSetting(context.Background(), omadacId, siteTemplateId).SiteRemoteLoggingSetting(siteRemoteLoggingSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateRemoteLoggingTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRemoteLoggingTemplateSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateRemoteLoggingTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteLoggingTemplateSettingRequest struct via the builder pattern


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


## UpdateRoamingTemplateSetting

> OperationResponseWithoutResult UpdateRoamingTemplateSetting(ctx, omadacId, siteTemplateId).SiteRoamingSetting(siteRoamingSetting).Execute()

Modify site template roaming setting



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
	siteRoamingSetting := *openapiclient.NewSiteRoamingSetting(*openapiclient.NewRoamingOpenApiVO(false, false)) // SiteRoamingSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateRoamingTemplateSetting(context.Background(), omadacId, siteTemplateId).SiteRoamingSetting(siteRoamingSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateRoamingTemplateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoamingTemplateSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateRoamingTemplateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoamingTemplateSettingRequest struct via the builder pattern


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


## UpdateSiteTemplateLedSetting

> OperationResponseWithoutResult UpdateSiteTemplateLedSetting(ctx, omadacId, siteTemplateId).SiteLedSetting(siteLedSetting).Execute()

Modify site template led setting



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
	siteLedSetting := *openapiclient.NewSiteLedSetting(false) // SiteLedSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteTemplateConfigurationAPI.UpdateSiteTemplateLedSetting(context.Background(), omadacId, siteTemplateId).SiteLedSetting(siteLedSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteTemplateConfigurationAPI.UpdateSiteTemplateLedSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteTemplateLedSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SiteTemplateConfigurationAPI.UpdateSiteTemplateLedSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteTemplateLedSettingRequest struct via the builder pattern


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


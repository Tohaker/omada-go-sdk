# GatewayTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGatewayInfo**](GatewayTemplateAPI.md#getgatewayinfo) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId} | Get gateway template info
[**GetPortsTemplate**](GatewayTemplateAPI.md#getportstemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/ports/config | Get gateway template port info
[**GetSsidDetailTemplate1**](GatewayTemplateAPI.md#getssiddetailtemplate1) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/wlans/ssid | Get SSID detail info
[**ModifyConfigAdvancedTemplate**](GatewayTemplateAPI.md#modifyconfigadvancedtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/advanced | Modify gateway template advanced config
[**ModifyConfigCommonAdvancedTemplate**](GatewayTemplateAPI.md#modifyconfigcommonadvancedtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/advanced/common | Modify gateway template common advanced config
[**ModifyConfigGeneralTemplate**](GatewayTemplateAPI.md#modifyconfiggeneraltemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/general | Modify gateway template general config
[**ModifyConfigRadiosTemplate**](GatewayTemplateAPI.md#modifyconfigradiostemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/radios | Modify gateway template radios config
[**ModifyConfigServicesTemplate**](GatewayTemplateAPI.md#modifyconfigservicestemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/services | Modify gateway template services config
[**ModifyConfigWirelessAdvancedTemplate**](GatewayTemplateAPI.md#modifyconfigwirelessadvancedtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/advanced/wireless | Modify gateway template wireless advanced config
[**ModifyConfigWlansTemplate**](GatewayTemplateAPI.md#modifyconfigwlanstemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/wlans | Modify gateway template wlans config
[**ModifyPortConfig**](GatewayTemplateAPI.md#modifyportconfig) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/ports/{port}/config | Modify gateway template port config
[**ModifySsidBasicConfigTemplate**](GatewayTemplateAPI.md#modifyssidbasicconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/wlans/ssid/{ssidId}/basic-config | modify SSID basic config
[**ModifySsidMacFilterConfigTemplate**](GatewayTemplateAPI.md#modifyssidmacfilterconfigtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/gatewaytemplates/{deviceTemplateId}/config/wlans/ssid/{ssidId}/mac-filter | Update SSID mac filter config



## GetGatewayInfo

> OperationResponseGatewayTemplateInfo GetGatewayInfo(ctx, omadacId, siteTemplateId, deviceTemplateId).Execute()

Get gateway template info



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.GetGatewayInfo(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.GetGatewayInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGatewayInfo`: OperationResponseGatewayTemplateInfo
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.GetGatewayInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseGatewayTemplateInfo**](OperationResponseGatewayTemplateInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortsTemplate

> OperationResponseListOsgPortInfoOpenApiVO GetPortsTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).Execute()

Get gateway template port info



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.GetPortsTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.GetPortsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortsTemplate`: OperationResponseListOsgPortInfoOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.GetPortsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListOsgPortInfoOpenApiVO**](OperationResponseListOsgPortInfoOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsidDetailTemplate1

> OperationResponseSsidDetailOpenApiVO GetSsidDetailTemplate1(ctx, omadacId, siteTemplateId, deviceTemplateId).Execute()

Get SSID detail info



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.GetSsidDetailTemplate1(context.Background(), omadacId, siteTemplateId, deviceTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.GetSsidDetailTemplate1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsidDetailTemplate1`: OperationResponseSsidDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.GetSsidDetailTemplate1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsidDetailTemplate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSsidDetailOpenApiVO**](OperationResponseSsidDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyConfigAdvancedTemplate

> OperationResponseWithoutResult ModifyConfigAdvancedTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).OsgConfigAdvancedOpenApiVO(osgConfigAdvancedOpenApiVO).Execute()

Modify gateway template advanced config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	osgConfigAdvancedOpenApiVO := *openapiclient.NewOsgConfigAdvancedOpenApiVO() // OsgConfigAdvancedOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifyConfigAdvancedTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).OsgConfigAdvancedOpenApiVO(osgConfigAdvancedOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifyConfigAdvancedTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigAdvancedTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifyConfigAdvancedTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyConfigAdvancedTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigAdvancedOpenApiVO** | [**OsgConfigAdvancedOpenApiVO**](OsgConfigAdvancedOpenApiVO.md) |  | 

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


## ModifyConfigCommonAdvancedTemplate

> OperationResponseWithoutResult ModifyConfigCommonAdvancedTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).OsgConfigCommonAdvancedOpenApiVO(osgConfigCommonAdvancedOpenApiVO).Execute()

Modify gateway template common advanced config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	osgConfigCommonAdvancedOpenApiVO := *openapiclient.NewOsgConfigCommonAdvancedOpenApiVO() // OsgConfigCommonAdvancedOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifyConfigCommonAdvancedTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).OsgConfigCommonAdvancedOpenApiVO(osgConfigCommonAdvancedOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifyConfigCommonAdvancedTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigCommonAdvancedTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifyConfigCommonAdvancedTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyConfigCommonAdvancedTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigCommonAdvancedOpenApiVO** | [**OsgConfigCommonAdvancedOpenApiVO**](OsgConfigCommonAdvancedOpenApiVO.md) |  | 

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


## ModifyConfigGeneralTemplate

> OperationResponseWithoutResult ModifyConfigGeneralTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).OsgGeneralConfigOpenApiV2VO(osgGeneralConfigOpenApiV2VO).Execute()

Modify gateway template general config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	osgGeneralConfigOpenApiV2VO := *openapiclient.NewOsgGeneralConfigOpenApiV2VO(int32(123)) // OsgGeneralConfigOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifyConfigGeneralTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).OsgGeneralConfigOpenApiV2VO(osgGeneralConfigOpenApiV2VO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifyConfigGeneralTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigGeneralTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifyConfigGeneralTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyConfigGeneralTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgGeneralConfigOpenApiV2VO** | [**OsgGeneralConfigOpenApiV2VO**](OsgGeneralConfigOpenApiV2VO.md) |  | 

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


## ModifyConfigRadiosTemplate

> OperationResponseWithoutResult ModifyConfigRadiosTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).OsgConfigRadiosopenApiVO(osgConfigRadiosopenApiVO).Execute()

Modify gateway template radios config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	osgConfigRadiosopenApiVO := *openapiclient.NewOsgConfigRadiosopenApiVO() // OsgConfigRadiosopenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifyConfigRadiosTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).OsgConfigRadiosopenApiVO(osgConfigRadiosopenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifyConfigRadiosTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigRadiosTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifyConfigRadiosTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyConfigRadiosTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigRadiosopenApiVO** | [**OsgConfigRadiosopenApiVO**](OsgConfigRadiosopenApiVO.md) |  | 

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


## ModifyConfigServicesTemplate

> OperationResponseWithoutResult ModifyConfigServicesTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).OsgConfigServicesOpenApiVO(osgConfigServicesOpenApiVO).Execute()

Modify gateway template services config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	osgConfigServicesOpenApiVO := *openapiclient.NewOsgConfigServicesOpenApiVO() // OsgConfigServicesOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifyConfigServicesTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).OsgConfigServicesOpenApiVO(osgConfigServicesOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifyConfigServicesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigServicesTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifyConfigServicesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyConfigServicesTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigServicesOpenApiVO** | [**OsgConfigServicesOpenApiVO**](OsgConfigServicesOpenApiVO.md) |  | 

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


## ModifyConfigWirelessAdvancedTemplate

> OperationResponseWithoutResult ModifyConfigWirelessAdvancedTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).OsgConfigWirelessAdvancedOpenApiVO(osgConfigWirelessAdvancedOpenApiVO).Execute()

Modify gateway template wireless advanced config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	osgConfigWirelessAdvancedOpenApiVO := *openapiclient.NewOsgConfigWirelessAdvancedOpenApiVO() // OsgConfigWirelessAdvancedOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifyConfigWirelessAdvancedTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).OsgConfigWirelessAdvancedOpenApiVO(osgConfigWirelessAdvancedOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifyConfigWirelessAdvancedTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigWirelessAdvancedTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifyConfigWirelessAdvancedTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyConfigWirelessAdvancedTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigWirelessAdvancedOpenApiVO** | [**OsgConfigWirelessAdvancedOpenApiVO**](OsgConfigWirelessAdvancedOpenApiVO.md) |  | 

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


## ModifyConfigWlansTemplate

> OperationResponseWithoutResult ModifyConfigWlansTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId).OsgConfigWlansOpenApiVO(osgConfigWlansOpenApiVO).Execute()

Modify gateway template wlans config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	osgConfigWlansOpenApiVO := *openapiclient.NewOsgConfigWlansOpenApiVO([]openapiclient.OsgSsidOverrideOpenApiVO{*openapiclient.NewOsgSsidOverrideOpenApiVO(false, "GlobalSsid_example", false, int32(123), int32(123), "Ssid_example", false, false)}) // OsgConfigWlansOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifyConfigWlansTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId).OsgConfigWlansOpenApiVO(osgConfigWlansOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifyConfigWlansTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyConfigWlansTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifyConfigWlansTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyConfigWlansTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **osgConfigWlansOpenApiVO** | [**OsgConfigWlansOpenApiVO**](OsgConfigWlansOpenApiVO.md) |  | 

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


## ModifyPortConfig

> OperationResponseGatewayTemplatePortSettingConfig ModifyPortConfig(ctx, omadacId, siteTemplateId, deviceTemplateId, port).GatewayTemplatePortSettingConfig(gatewayTemplatePortSettingConfig).Execute()

Modify gateway template port config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Device Template ID
	port := "port_example" // string | Gateway port number
	gatewayTemplatePortSettingConfig := *openapiclient.NewGatewayTemplatePortSettingConfig() // GatewayTemplatePortSettingConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifyPortConfig(context.Background(), omadacId, siteTemplateId, deviceTemplateId, port).GatewayTemplatePortSettingConfig(gatewayTemplatePortSettingConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifyPortConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPortConfig`: OperationResponseGatewayTemplatePortSettingConfig
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifyPortConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Device Template ID | 
**port** | **string** | Gateway port number | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPortConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gatewayTemplatePortSettingConfig** | [**GatewayTemplatePortSettingConfig**](GatewayTemplatePortSettingConfig.md) |  | 

### Return type

[**OperationResponseGatewayTemplatePortSettingConfig**](OperationResponseGatewayTemplatePortSettingConfig.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySsidBasicConfigTemplate

> OperationResponseWithoutResult ModifySsidBasicConfigTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId, ssidId).UpdateSsidBasicConfigForIpptOpenApiVO(updateSsidBasicConfigForIpptOpenApiVO).Execute()

modify SSID basic config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	ssidId := "ssidId_example" // string | Device Template ID
	updateSsidBasicConfigForIpptOpenApiVO := *openapiclient.NewUpdateSsidBasicConfigForIpptOpenApiVO(int32(123), false, "Name_example", int32(123)) // UpdateSsidBasicConfigForIpptOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifySsidBasicConfigTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId, ssidId).UpdateSsidBasicConfigForIpptOpenApiVO(updateSsidBasicConfigForIpptOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifySsidBasicConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySsidBasicConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifySsidBasicConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 
**ssidId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySsidBasicConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidBasicConfigForIpptOpenApiVO** | [**UpdateSsidBasicConfigForIpptOpenApiVO**](UpdateSsidBasicConfigForIpptOpenApiVO.md) |  | 

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


## ModifySsidMacFilterConfigTemplate

> OperationResponseWithoutResult ModifySsidMacFilterConfigTemplate(ctx, omadacId, siteTemplateId, deviceTemplateId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()

Update SSID mac filter config



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
	deviceTemplateId := "deviceTemplateId_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	ssidId := "ssidId_example" // string | Device Template ID
	updateSsidMacFilterOpenApiVO := *openapiclient.NewUpdateSsidMacFilterOpenApiVO(false) // UpdateSsidMacFilterOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayTemplateAPI.ModifySsidMacFilterConfigTemplate(context.Background(), omadacId, siteTemplateId, deviceTemplateId, ssidId).UpdateSsidMacFilterOpenApiVO(updateSsidMacFilterOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayTemplateAPI.ModifySsidMacFilterConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySsidMacFilterConfigTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `GatewayTemplateAPI.ModifySsidMacFilterConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**deviceTemplateId** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 
**ssidId** | **string** | Device Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySsidMacFilterConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateSsidMacFilterOpenApiVO** | [**UpdateSsidMacFilterOpenApiVO**](UpdateSsidMacFilterOpenApiVO.md) |  | 

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


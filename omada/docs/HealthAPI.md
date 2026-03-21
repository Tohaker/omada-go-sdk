# \HealthAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApHealthDetail**](HealthAPI.md#GetApHealthDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/eaps/{apMac}/health/detail | Get ap health detail
[**GetApHealthTimeList**](HealthAPI.md#GetApHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/eaps/{apMac}/health/timeline | Get ap health time line
[**GetClientHealthDetail**](HealthAPI.md#GetClientHealthDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/health/detail | Get client health detail
[**GetClientHealthTimeList**](HealthAPI.md#GetClientHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/health/timeline | Get client health time line
[**GetMspApHealthDetail**](HealthAPI.md#GetMspApHealthDetail) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/eaps/{apMac}/health/detail | Get msp ap health detail
[**GetMspApHealthTimeList**](HealthAPI.md#GetMspApHealthTimeList) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/eaps/{apMac}/health/timeline | Get msp ap health time line
[**GetMspOsgHealthDetail**](HealthAPI.md#GetMspOsgHealthDetail) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/gateways/{gatewayMac}/health/detail | Get msp gateway health detail
[**GetMspOsgHealthTimeList**](HealthAPI.md#GetMspOsgHealthTimeList) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/gateways/{gatewayMac}/health/timeline | Get msp gateway health time line
[**GetMspOswHealthDetail**](HealthAPI.md#GetMspOswHealthDetail) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/switches/{switchMac}/health/detail | Get msp switch health detail
[**GetMspOswHealthTimeList**](HealthAPI.md#GetMspOswHealthTimeList) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/switches/{switchMac}/health/timeline | Get msp switch health time line
[**GetMultiOswHealthTimeList**](HealthAPI.md#GetMultiOswHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/health/timeline | Batch get switch health time line
[**GetOsgHealthDetail**](HealthAPI.md#GetOsgHealthDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/health/detail | Get gateway health detail
[**GetOsgHealthTimeList**](HealthAPI.md#GetOsgHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/health/timeline | Get gateway health time line
[**GetOswHealthDetail**](HealthAPI.md#GetOswHealthDetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/health/detail | Get switch health detail
[**GetOswHealthTimeList**](HealthAPI.md#GetOswHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/health/timeline | Get switch health time line
[**GetSiteClientHealthTimeList**](HealthAPI.md#GetSiteClientHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/health/client/timeline | Get site client health time line
[**GetSiteHealthTimeList**](HealthAPI.md#GetSiteHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/health/timeline | Get site health time line
[**GetWifiClientInfoList**](HealthAPI.md#GetWifiClientInfoList) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/wifi/health/client/list | Get site wifi top k device/client 
[**GetWifiHealthTimeList**](HealthAPI.md#GetWifiHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wifi/health/timeline | Get site wifi health time line
[**GetWifiSubHealthTimeList**](HealthAPI.md#GetWifiSubHealthTimeList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/wifi/sub/health/timeline | Get site wifi sub health detail 



## GetApHealthDetail

> OperationResponseApHealthDetailVO GetApHealthDetail(ctx, omadacId, siteId, apMac).Start(start).End(end).Execute()

Get ap health detail



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetApHealthDetail(context.Background(), omadacId, siteId, apMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetApHealthDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApHealthDetail`: OperationResponseApHealthDetailVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetApHealthDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApHealthDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseApHealthDetailVO**](OperationResponseApHealthDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApHealthTimeList

> OperationResponseTimeScoreListVO GetApHealthTimeList(ctx, omadacId, siteId, apMac).Start(start).End(end).Execute()

Get ap health time line



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
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetApHealthTimeList(context.Background(), omadacId, siteId, apMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetApHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApHealthTimeList`: OperationResponseTimeScoreListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetApHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseTimeScoreListVO**](OperationResponseTimeScoreListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientHealthDetail

> OperationResponseClientHealthDetailVO GetClientHealthDetail(ctx, omadacId, siteId, clientMac).Start(start).End(end).Execute()

Get client health detail



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
	clientMac := "clientMac_example" // string | Client MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetClientHealthDetail(context.Background(), omadacId, siteId, clientMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetClientHealthDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientHealthDetail`: OperationResponseClientHealthDetailVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetClientHealthDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientHealthDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseClientHealthDetailVO**](OperationResponseClientHealthDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientHealthTimeList

> OperationResponseTimeScoreListVO GetClientHealthTimeList(ctx, omadacId, siteId, clientMac).Start(start).End(end).Execute()

Get client health time line



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
	clientMac := "clientMac_example" // string | Client MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetClientHealthTimeList(context.Background(), omadacId, siteId, clientMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetClientHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientHealthTimeList`: OperationResponseTimeScoreListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetClientHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseTimeScoreListVO**](OperationResponseTimeScoreListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspApHealthDetail

> OperationResponseApHealthDetailVO GetMspApHealthDetail(ctx, mspId, customerId, siteId, apMac).Start(start).End(end).Execute()

Get msp ap health detail



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetMspApHealthDetail(context.Background(), mspId, customerId, siteId, apMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetMspApHealthDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspApHealthDetail`: OperationResponseApHealthDetailVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetMspApHealthDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspApHealthDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseApHealthDetailVO**](OperationResponseApHealthDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspApHealthTimeList

> OperationResponseTimeScoreListVO GetMspApHealthTimeList(ctx, mspId, customerId, siteId, apMac).Start(start).End(end).Execute()

Get msp ap health time line



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	apMac := "apMac_example" // string | AP MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetMspApHealthTimeList(context.Background(), mspId, customerId, siteId, apMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetMspApHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspApHealthTimeList`: OperationResponseTimeScoreListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetMspApHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**apMac** | **string** | AP MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspApHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseTimeScoreListVO**](OperationResponseTimeScoreListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspOsgHealthDetail

> OperationResponseOsgHealthDetailVO GetMspOsgHealthDetail(ctx, mspId, customerId, siteId, gatewayMac).Start(start).End(end).WirelessRouter(wirelessRouter).Execute()

Get msp gateway health detail



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000
	wirelessRouter := true // bool | Whether the device is a wireless router (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetMspOsgHealthDetail(context.Background(), mspId, customerId, siteId, gatewayMac).Start(start).End(end).WirelessRouter(wirelessRouter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetMspOsgHealthDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOsgHealthDetail`: OperationResponseOsgHealthDetailVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetMspOsgHealthDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOsgHealthDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 
 **wirelessRouter** | **bool** | Whether the device is a wireless router | 

### Return type

[**OperationResponseOsgHealthDetailVO**](OperationResponseOsgHealthDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspOsgHealthTimeList

> OperationResponseTimeScoreListVO GetMspOsgHealthTimeList(ctx, mspId, customerId, siteId, gatewayMac).Start(start).End(end).Execute()

Get msp gateway health time line



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetMspOsgHealthTimeList(context.Background(), mspId, customerId, siteId, gatewayMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetMspOsgHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOsgHealthTimeList`: OperationResponseTimeScoreListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetMspOsgHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOsgHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseTimeScoreListVO**](OperationResponseTimeScoreListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspOswHealthDetail

> OperationResponseOswHealthDetailVO GetMspOswHealthDetail(ctx, mspId, customerId, siteId, switchMac).Start(start).End(end).Execute()

Get msp switch health detail



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	switchMac := "switchMac_example" // string | Switch MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetMspOswHealthDetail(context.Background(), mspId, customerId, siteId, switchMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetMspOswHealthDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOswHealthDetail`: OperationResponseOswHealthDetailVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetMspOswHealthDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOswHealthDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseOswHealthDetailVO**](OperationResponseOswHealthDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspOswHealthTimeList

> OperationResponseTimeScoreListVO GetMspOswHealthTimeList(ctx, mspId, customerId, siteId, switchMac).Start(start).End(end).Execute()

Get msp switch health time line



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
	mspId := "mspId_example" // string | MSP ID
	customerId := "customerId_example" // string | Customer ID
	siteId := "siteId_example" // string | Site ID
	switchMac := "switchMac_example" // string | Switch MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetMspOswHealthTimeList(context.Background(), mspId, customerId, siteId, switchMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetMspOswHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOswHealthTimeList`: OperationResponseTimeScoreListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetMspOswHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**customerId** | **string** | Customer ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOswHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseTimeScoreListVO**](OperationResponseTimeScoreListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultiOswHealthTimeList

> OperationResponseHealthTimeLineVO GetMultiOswHealthTimeList(ctx, omadacId, siteId).Macs(macs).Start(start).End(end).Execute()

Batch get switch health time line



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
	macs := "macs_example" // string | Switch MAC address list, using comma separator, like AA-BB-CC-DD-EE-FF,FF-EE-DD-CC-BB-AA
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetMultiOswHealthTimeList(context.Background(), omadacId, siteId).Macs(macs).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetMultiOswHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultiOswHealthTimeList`: OperationResponseHealthTimeLineVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetMultiOswHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiOswHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macs** | **string** | Switch MAC address list, using comma separator, like AA-BB-CC-DD-EE-FF,FF-EE-DD-CC-BB-AA | 
 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseHealthTimeLineVO**](OperationResponseHealthTimeLineVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsgHealthDetail

> OperationResponseOsgHealthDetailVO GetOsgHealthDetail(ctx, omadacId, siteId, gatewayMac).Start(start).End(end).WirelessRouter(wirelessRouter).Execute()

Get gateway health detail



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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000
	wirelessRouter := true // bool | Whether the device is a wireless router (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetOsgHealthDetail(context.Background(), omadacId, siteId, gatewayMac).Start(start).End(end).WirelessRouter(wirelessRouter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetOsgHealthDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsgHealthDetail`: OperationResponseOsgHealthDetailVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetOsgHealthDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsgHealthDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 
 **wirelessRouter** | **bool** | Whether the device is a wireless router | 

### Return type

[**OperationResponseOsgHealthDetailVO**](OperationResponseOsgHealthDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsgHealthTimeList

> OperationResponseTimeScoreListVO GetOsgHealthTimeList(ctx, omadacId, siteId, gatewayMac).Start(start).End(end).Execute()

Get gateway health time line



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
	gatewayMac := "gatewayMac_example" // string | Gateway MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetOsgHealthTimeList(context.Background(), omadacId, siteId, gatewayMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetOsgHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsgHealthTimeList`: OperationResponseTimeScoreListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetOsgHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Gateway MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsgHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseTimeScoreListVO**](OperationResponseTimeScoreListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswHealthDetail

> OperationResponseOswHealthDetailVO GetOswHealthDetail(ctx, omadacId, siteId, switchMac).Start(start).End(end).Execute()

Get switch health detail



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
	switchMac := "switchMac_example" // string | Switch MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetOswHealthDetail(context.Background(), omadacId, siteId, switchMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetOswHealthDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswHealthDetail`: OperationResponseOswHealthDetailVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetOswHealthDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswHealthDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseOswHealthDetailVO**](OperationResponseOswHealthDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswHealthTimeList

> OperationResponseTimeScoreListVO GetOswHealthTimeList(ctx, omadacId, siteId, switchMac).Start(start).End(end).Execute()

Get switch health time line



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
	switchMac := "switchMac_example" // string | Switch MAC address, like AA-BB-CC-DD-EE-FF
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetOswHealthTimeList(context.Background(), omadacId, siteId, switchMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetOswHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswHealthTimeList`: OperationResponseTimeScoreListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetOswHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseTimeScoreListVO**](OperationResponseTimeScoreListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteClientHealthTimeList

> OperationResponseClientScoreTimelineListVO GetSiteClientHealthTimeList(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get site client health time line



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
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetSiteClientHealthTimeList(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetSiteClientHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteClientHealthTimeList`: OperationResponseClientScoreTimelineListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetSiteClientHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteClientHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseClientScoreTimelineListVO**](OperationResponseClientScoreTimelineListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteHealthTimeList

> OperationResponseSiteScoreTimelineListVO GetSiteHealthTimeList(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get site health time line



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
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetSiteHealthTimeList(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetSiteHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteHealthTimeList`: OperationResponseSiteScoreTimelineListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetSiteHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseSiteScoreTimelineListVO**](OperationResponseSiteScoreTimelineListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWifiClientInfoList

> WifiClientInfoVO GetWifiClientInfoList(ctx, omadacId, siteId).WifiDeviceAndClientQueryVO(wifiDeviceAndClientQueryVO).Execute()

Get site wifi top k device/client 



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
	wifiDeviceAndClientQueryVO := *openapiclient.NewWifiDeviceAndClientQueryVO(int64(123)) // WifiDeviceAndClientQueryVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetWifiClientInfoList(context.Background(), omadacId, siteId).WifiDeviceAndClientQueryVO(wifiDeviceAndClientQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetWifiClientInfoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWifiClientInfoList`: WifiClientInfoVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetWifiClientInfoList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWifiClientInfoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wifiDeviceAndClientQueryVO** | [**WifiDeviceAndClientQueryVO**](WifiDeviceAndClientQueryVO.md) |  | 

### Return type

[**WifiClientInfoVO**](WifiClientInfoVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWifiHealthTimeList

> OperationResponseTimeScoreListVO GetWifiHealthTimeList(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get site wifi health time line



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
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetWifiHealthTimeList(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetWifiHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWifiHealthTimeList`: OperationResponseTimeScoreListVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetWifiHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWifiHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**OperationResponseTimeScoreListVO**](OperationResponseTimeScoreListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWifiSubHealthTimeList

> WifiHealthDetailVO GetWifiSubHealthTimeList(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get site wifi sub health detail 



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
	start := int64(789) // int64 | Start timestamp, in milliseconds, such as 174951360000
	end := int64(789) // int64 | End timestamp, in milliseconds, such as 1749600000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthAPI.GetWifiSubHealthTimeList(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetWifiSubHealthTimeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWifiSubHealthTimeList`: WifiHealthDetailVO
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetWifiSubHealthTimeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWifiSubHealthTimeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in milliseconds, such as 174951360000 | 
 **end** | **int64** | End timestamp, in milliseconds, such as 1749600000000 | 

### Return type

[**WifiHealthDetailVO**](WifiHealthDetailVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


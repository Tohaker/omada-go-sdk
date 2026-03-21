# \StatisticAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDevice5MinStatistic**](StatisticAPI.md#GetDevice5MinStatistic) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/stat/{deviceMac}/5min | Get device statistic data 5 min
[**GetDeviceDailyStatistic**](StatisticAPI.md#GetDeviceDailyStatistic) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/stat/{deviceMac}/daily | Get device statistic data daily
[**GetDeviceHourlyStatistic**](StatisticAPI.md#GetDeviceHourlyStatistic) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/stat/{deviceMac}/hourly | Get device statistic data hourly
[**GetOltStatChart**](StatisticAPI.md#GetOltStatChart) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stat/olts/{deviceMac}/chart | Get olt statistics chart
[**GetOswRankingCards**](StatisticAPI.md#GetOswRankingCards) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/health/switches/rankingCards | Get switch ranking cards
[**GetOswStackDetailStat**](StatisticAPI.md#GetOswStackDetailStat) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stat/stack/{stackId} | Get switch stack statistics
[**GetStackStatChart**](StatisticAPI.md#GetStackStatChart) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stat/stack/{stackId}/chart | Get stack statistics chart
[**GetSwitchStat**](StatisticAPI.md#GetSwitchStat) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/stat/switches/{deviceMac} | Get switch statistics
[**GetSwitchStatChart**](StatisticAPI.md#GetSwitchStatChart) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stat/switches/{deviceMac}/chart | Get switch statistics chart



## GetDevice5MinStatistic

> OperationResponseBaseDeviceStatDTO GetDevice5MinStatistic(ctx, omadacId, siteId, deviceMac).Type_(type_).StatQueryVO(statQueryVO).Execute()

Get device statistic data 5 min



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	type_ := "type__example" // string | Device Type. Supported type: ap, gateway, switch, olt.
	statQueryVO := *openapiclient.NewStatQueryVO([]string{"Attrs_example"}, int64(123), int64(123)) // StatQueryVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetDevice5MinStatistic(context.Background(), omadacId, siteId, deviceMac).Type_(type_).StatQueryVO(statQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetDevice5MinStatistic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevice5MinStatistic`: OperationResponseBaseDeviceStatDTO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetDevice5MinStatistic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevice5MinStatisticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** | Device Type. Supported type: ap, gateway, switch, olt. | 
 **statQueryVO** | [**StatQueryVO**](StatQueryVO.md) |  | 

### Return type

[**OperationResponseBaseDeviceStatDTO**](OperationResponseBaseDeviceStatDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceDailyStatistic

> OperationResponseBaseDeviceStatDTO GetDeviceDailyStatistic(ctx, omadacId, siteId, deviceMac).Type_(type_).StatQueryVO(statQueryVO).Execute()

Get device statistic data daily



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	type_ := "type__example" // string | Device Type. Supported type: ap, gateway, switch, olt.
	statQueryVO := *openapiclient.NewStatQueryVO([]string{"Attrs_example"}, int64(123), int64(123)) // StatQueryVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetDeviceDailyStatistic(context.Background(), omadacId, siteId, deviceMac).Type_(type_).StatQueryVO(statQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetDeviceDailyStatistic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceDailyStatistic`: OperationResponseBaseDeviceStatDTO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetDeviceDailyStatistic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceDailyStatisticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** | Device Type. Supported type: ap, gateway, switch, olt. | 
 **statQueryVO** | [**StatQueryVO**](StatQueryVO.md) |  | 

### Return type

[**OperationResponseBaseDeviceStatDTO**](OperationResponseBaseDeviceStatDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceHourlyStatistic

> OperationResponseBaseDeviceStatDTO GetDeviceHourlyStatistic(ctx, omadacId, siteId, deviceMac).Type_(type_).StatQueryVO(statQueryVO).Execute()

Get device statistic data hourly



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	type_ := "type__example" // string | Device Type. Supported type: ap, gateway, switch, olt.
	statQueryVO := *openapiclient.NewStatQueryVO([]string{"Attrs_example"}, int64(123), int64(123)) // StatQueryVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetDeviceHourlyStatistic(context.Background(), omadacId, siteId, deviceMac).Type_(type_).StatQueryVO(statQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetDeviceHourlyStatistic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceHourlyStatistic`: OperationResponseBaseDeviceStatDTO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetDeviceHourlyStatistic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceHourlyStatisticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** | Device Type. Supported type: ap, gateway, switch, olt. | 
 **statQueryVO** | [**StatQueryVO**](StatQueryVO.md) |  | 

### Return type

[**OperationResponseBaseDeviceStatDTO**](OperationResponseBaseDeviceStatDTO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOltStatChart

> OperationResponseOltStatOpenApiVO GetOltStatChart(ctx, omadacId, siteId, deviceMac).OltStatQueryOpenApiDTO(oltStatQueryOpenApiDTO).Execute()

Get olt statistics chart



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	oltStatQueryOpenApiDTO := *openapiclient.NewOltStatQueryOpenApiDTO([]string{"Attrs_example"}, int64(123), int64(123)) // OltStatQueryOpenApiDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetOltStatChart(context.Background(), omadacId, siteId, deviceMac).OltStatQueryOpenApiDTO(oltStatQueryOpenApiDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetOltStatChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOltStatChart`: OperationResponseOltStatOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetOltStatChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOltStatChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oltStatQueryOpenApiDTO** | [**OltStatQueryOpenApiDTO**](OltStatQueryOpenApiDTO.md) |  | 

### Return type

[**OperationResponseOltStatOpenApiVO**](OperationResponseOltStatOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswRankingCards

> OperationResponseOswRankingCardsOpenApiVO GetOswRankingCards(ctx, omadacId, siteId).TimeQueryOpenApiVO(timeQueryOpenApiVO).Execute()

Get switch ranking cards



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
	timeQueryOpenApiVO := *openapiclient.NewTimeQueryOpenApiVO(int64(123), int64(123)) // TimeQueryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetOswRankingCards(context.Background(), omadacId, siteId).TimeQueryOpenApiVO(timeQueryOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetOswRankingCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswRankingCards`: OperationResponseOswRankingCardsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetOswRankingCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswRankingCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeQueryOpenApiVO** | [**TimeQueryOpenApiVO**](TimeQueryOpenApiVO.md) |  | 

### Return type

[**OperationResponseOswRankingCardsOpenApiVO**](OperationResponseOswRankingCardsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswStackDetailStat

> OperationResponseListOswStackDetailStatVO GetOswStackDetailStat(ctx, omadacId, siteId, stackId).OswStackStatQueryVO(oswStackStatQueryVO).Execute()

Get switch stack statistics



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
	stackId := "stackId_example" // string | Stack ID
	oswStackStatQueryVO := *openapiclient.NewOswStackStatQueryVO([]string{"Attrs_example"}, int64(123), int32(123), int64(123)) // OswStackStatQueryVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetOswStackDetailStat(context.Background(), omadacId, siteId, stackId).OswStackStatQueryVO(oswStackStatQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetOswStackDetailStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswStackDetailStat`: OperationResponseListOswStackDetailStatVO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetOswStackDetailStat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswStackDetailStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswStackStatQueryVO** | [**OswStackStatQueryVO**](OswStackStatQueryVO.md) |  | 

### Return type

[**OperationResponseListOswStackDetailStatVO**](OperationResponseListOswStackDetailStatVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackStatChart

> OperationResponseListOswStackMemberStatVO GetStackStatChart(ctx, omadacId, siteId, stackId).OswStackStatQueryOpenApiDTO(oswStackStatQueryOpenApiDTO).Execute()

Get stack statistics chart



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
	stackId := "stackId_example" // string | Stack ID
	oswStackStatQueryOpenApiDTO := *openapiclient.NewOswStackStatQueryOpenApiDTO([]string{"Attrs_example"}, int64(123), int64(123)) // OswStackStatQueryOpenApiDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetStackStatChart(context.Background(), omadacId, siteId, stackId).OswStackStatQueryOpenApiDTO(oswStackStatQueryOpenApiDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetStackStatChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackStatChart`: OperationResponseListOswStackMemberStatVO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetStackStatChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**stackId** | **string** | Stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackStatChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswStackStatQueryOpenApiDTO** | [**OswStackStatQueryOpenApiDTO**](OswStackStatQueryOpenApiDTO.md) |  | 

### Return type

[**OperationResponseListOswStackMemberStatVO**](OperationResponseListOswStackMemberStatVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchStat

> OperationResponseStatisticsOswVO GetSwitchStat(ctx, omadacId, siteId, deviceMac).Execute()

Get switch statistics



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetSwitchStat(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetSwitchStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchStat`: OperationResponseStatisticsOswVO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetSwitchStat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseStatisticsOswVO**](OperationResponseStatisticsOswVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchStatChart

> OperationResponseOswStatOpenApiVO GetSwitchStatChart(ctx, omadacId, siteId, deviceMac).OswStatQueryOpenApiDTO(oswStatQueryOpenApiDTO).Execute()

Get switch statistics chart



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	oswStatQueryOpenApiDTO := *openapiclient.NewOswStatQueryOpenApiDTO([]string{"Attrs_example"}, int64(123), int64(123)) // OswStatQueryOpenApiDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticAPI.GetSwitchStatChart(context.Background(), omadacId, siteId, deviceMac).OswStatQueryOpenApiDTO(oswStatQueryOpenApiDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.GetSwitchStatChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchStatChart`: OperationResponseOswStatOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.GetSwitchStatChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchStatChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oswStatQueryOpenApiDTO** | [**OswStatQueryOpenApiDTO**](OswStatQueryOpenApiDTO.md) |  | 

### Return type

[**OperationResponseOswStatOpenApiVO**](OperationResponseOswStatOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


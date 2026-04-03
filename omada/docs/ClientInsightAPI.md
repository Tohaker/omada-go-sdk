# ClientInsightAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApDensity**](ClientInsightAPI.md#getapdensity) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/ap-density | Get ap density.
[**GetClientActivity**](ClientInsightAPI.md#getclientactivity) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/client-activity | Get client activity
[**GetClientNumberForCustomerList**](ClientInsightAPI.md#getclientnumberforcustomerlist) | **Post** /openapi/v1/msp/{mspId}/customers/client-count | Get the Msp customers&#39; client count.
[**GetClientsAssociationActivities**](ClientInsightAPI.md#getclientsassociationactivities) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/client-association-activities | Get clients association activities.
[**GetClientsAssociationTimeDistribution**](ClientInsightAPI.md#getclientsassociationtimedistribution) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/client-association-time-distribution | Get clients association time distribution.
[**GetClientsBandDistribution**](ClientInsightAPI.md#getclientsbanddistribution) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/client-signal-distribution | Get clients signal distribution.
[**GetClientsBubble**](ClientInsightAPI.md#getclientsbubble) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/clients-bubble | Get clients bubble.
[**GetClientsDistribution**](ClientInsightAPI.md#getclientsdistribution) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/client-distribution | Get client distribution.
[**GetClientsFreqDistribution**](ClientInsightAPI.md#getclientsfreqdistribution) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/client-freq-distribution | Get clients freq distribution.
[**GetClientsRssiDistribution**](ClientInsightAPI.md#getclientsrssidistribution) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/client-rssi-distribution | Get clients rssi distribution.
[**GetClientsSsidDistribution**](ClientInsightAPI.md#getclientsssiddistribution) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/client-ssid-distribution | Get clients ssid distribution.
[**GetCurrentClientNum**](ClientInsightAPI.md#getcurrentclientnum) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/current-client-num | Get current client number.
[**GetDeviceClient5MinStat**](ClientInsightAPI.md#getdeviceclient5minstat) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stat/{deviceMac}/client-stat-5min | Get device client 5 min stat.
[**GetGridKnownClients**](ClientInsightAPI.md#getgridknownclients) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/insight/clients | Get known clients list
[**GetGridPastConnections**](ClientInsightAPI.md#getgridpastconnections) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/insight/past-connection | Get client past connection list
[**GetLongestClientUptime**](ClientInsightAPI.md#getlongestclientuptime) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/longest-uptime | Get longest client uptime.
[**GetMostActiveClients**](ClientInsightAPI.md#getmostactiveclients) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/active-clients | Get most active client.
[**GetMspDashboardOverall**](ClientInsightAPI.md#getmspdashboardoverall) | **Get** /openapi/v1/msp/{mspId}/dashboard/client/overview-diagram | Get the msp overview diagram of client.
[**GetPastClientNum**](ClientInsightAPI.md#getpastclientnum) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/past-client-num | Get past client number.
[**GetStackClientStat**](ClientInsightAPI.md#getstackclientstat) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stat/stacks/{stackId}/client-stat | Get stack client stat.



## GetApDensity

> OperationResponseEapRssiChannelDistributionVO GetApDensity(ctx, omadacId, siteId).Execute()

Get ap density.



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
	resp, r, err := apiClient.ClientInsightAPI.GetApDensity(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetApDensity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApDensity`: OperationResponseEapRssiChannelDistributionVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetApDensity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApDensityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseEapRssiChannelDistributionVO**](OperationResponseEapRssiChannelDistributionVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientActivity

> OperationResponseListClientActivitiesVO GetClientActivity(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get client activity



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000 (optional)
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetClientActivity(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientActivity`: OperationResponseListClientActivitiesVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListClientActivitiesVO**](OperationResponseListClientActivitiesVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientNumberForCustomerList

> OperationResponseListCustomerListClientNumVO GetClientNumberForCustomerList(ctx, mspId).CustomerQueryOpenApiVO(customerQueryOpenApiVO).Execute()

Get the Msp customers' client count.



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
	customerQueryOpenApiVO := *openapiclient.NewCustomerQueryOpenApiVO([]string{"CustomerIds_example"}) // CustomerQueryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetClientNumberForCustomerList(context.Background(), mspId).CustomerQueryOpenApiVO(customerQueryOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientNumberForCustomerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientNumberForCustomerList`: OperationResponseListCustomerListClientNumVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientNumberForCustomerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientNumberForCustomerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerQueryOpenApiVO** | [**CustomerQueryOpenApiVO**](CustomerQueryOpenApiVO.md) |  | 

### Return type

[**OperationResponseListCustomerListClientNumVO**](OperationResponseListCustomerListClientNumVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsAssociationActivities

> OperationResponseClientSummaryVO GetClientsAssociationActivities(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get clients association activities.



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetClientsAssociationActivities(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientsAssociationActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsAssociationActivities`: OperationResponseClientSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientsAssociationActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsAssociationActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseClientSummaryVO**](OperationResponseClientSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsAssociationTimeDistribution

> OperationResponseClientAssociationTimeDistributionVO GetClientsAssociationTimeDistribution(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get clients association time distribution.



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetClientsAssociationTimeDistribution(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientsAssociationTimeDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsAssociationTimeDistribution`: OperationResponseClientAssociationTimeDistributionVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientsAssociationTimeDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsAssociationTimeDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseClientAssociationTimeDistributionVO**](OperationResponseClientAssociationTimeDistributionVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsBandDistribution

> OperationResponseClientSignalDistributionVO GetClientsBandDistribution(ctx, omadacId, siteId).Execute()

Get clients signal distribution.



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
	resp, r, err := apiClient.ClientInsightAPI.GetClientsBandDistribution(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientsBandDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsBandDistribution`: OperationResponseClientSignalDistributionVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientsBandDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsBandDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseClientSignalDistributionVO**](OperationResponseClientSignalDistributionVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsBubble

> OperationResponseListActiveClientBubbleVO GetClientsBubble(ctx, omadacId, siteId).Execute()

Get clients bubble.



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
	resp, r, err := apiClient.ClientInsightAPI.GetClientsBubble(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientsBubble``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsBubble`: OperationResponseListActiveClientBubbleVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientsBubble`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsBubbleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListActiveClientBubbleVO**](OperationResponseListActiveClientBubbleVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsDistribution

> OperationResponseDashboardDistributionVO GetClientsDistribution(ctx, omadacId, siteId).Execute()

Get client distribution.



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
	resp, r, err := apiClient.ClientInsightAPI.GetClientsDistribution(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientsDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsDistribution`: OperationResponseDashboardDistributionVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientsDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDashboardDistributionVO**](OperationResponseDashboardDistributionVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsFreqDistribution

> OperationResponseClientSummaryVO GetClientsFreqDistribution(ctx, omadacId, siteId).Execute()

Get clients freq distribution.



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
	resp, r, err := apiClient.ClientInsightAPI.GetClientsFreqDistribution(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientsFreqDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsFreqDistribution`: OperationResponseClientSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientsFreqDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsFreqDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseClientSummaryVO**](OperationResponseClientSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsRssiDistribution

> OperationResponseClientRssiChannelDistributionVO GetClientsRssiDistribution(ctx, omadacId, siteId).Execute()

Get clients rssi distribution.



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
	resp, r, err := apiClient.ClientInsightAPI.GetClientsRssiDistribution(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientsRssiDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsRssiDistribution`: OperationResponseClientRssiChannelDistributionVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientsRssiDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRssiDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseClientRssiChannelDistributionVO**](OperationResponseClientRssiChannelDistributionVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsSsidDistribution

> OperationResponseClientSummaryVO GetClientsSsidDistribution(ctx, omadacId, siteId).Execute()

Get clients ssid distribution.



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
	resp, r, err := apiClient.ClientInsightAPI.GetClientsSsidDistribution(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetClientsSsidDistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsSsidDistribution`: OperationResponseClientSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetClientsSsidDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsSsidDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseClientSummaryVO**](OperationResponseClientSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentClientNum

> OperationResponseClientSummaryVO GetCurrentClientNum(ctx, omadacId, siteId).Execute()

Get current client number.



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
	resp, r, err := apiClient.ClientInsightAPI.GetCurrentClientNum(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetCurrentClientNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentClientNum`: OperationResponseClientSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetCurrentClientNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentClientNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseClientSummaryVO**](OperationResponseClientSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceClient5MinStat

> OperationResponseListDeviceClientStatVO GetDeviceClient5MinStat(ctx, omadacId, siteId, deviceMac).DeviceType(deviceType).DeviceClientStatQueryOpenApiVO(deviceClientStatQueryOpenApiVO).Execute()

Get device client 5 min stat.



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
	deviceType := "deviceType_example" // string | Device type.
	deviceClientStatQueryOpenApiVO := *openapiclient.NewDeviceClientStatQueryOpenApiVO([]string{"Attrs_example"}, int64(123), int64(123)) // DeviceClientStatQueryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetDeviceClient5MinStat(context.Background(), omadacId, siteId, deviceMac).DeviceType(deviceType).DeviceClientStatQueryOpenApiVO(deviceClientStatQueryOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetDeviceClient5MinStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceClient5MinStat`: OperationResponseListDeviceClientStatVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetDeviceClient5MinStat`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDeviceClient5MinStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deviceType** | **string** | Device type. | 
 **deviceClientStatQueryOpenApiVO** | [**DeviceClientStatQueryOpenApiVO**](DeviceClientStatQueryOpenApiVO.md) |  | 

### Return type

[**OperationResponseListDeviceClientStatVO**](OperationResponseListDeviceClientStatVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridKnownClients

> OperationResponseGridVOKnownClientVO GetGridKnownClients(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsLastSeen(sortsLastSeen).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersGuest(filtersGuest).SearchKey(searchKey).Execute()

Get known clients list



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
	sortsLastSeen := "sortsLastSeen_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersTimeStart := "filtersTimeStart_example" // string | Filter query parameters, support field time range: start timestamp (ms). (optional)
	filtersTimeEnd := "filtersTimeEnd_example" // string | Filter query parameters, support field time range: end timestamp (ms). (optional)
	filtersGuest := "filtersGuest_example" // string | Filter query parameters, support field guest: true/false. (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name,mac,ssid. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetGridKnownClients(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsLastSeen(sortsLastSeen).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersGuest(filtersGuest).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetGridKnownClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridKnownClients`: OperationResponseGridVOKnownClientVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetGridKnownClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridKnownClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsLastSeen** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersTimeStart** | **string** | Filter query parameters, support field time range: start timestamp (ms). | 
 **filtersTimeEnd** | **string** | Filter query parameters, support field time range: end timestamp (ms). | 
 **filtersGuest** | **string** | Filter query parameters, support field guest: true/false. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name,mac,ssid. | 

### Return type

[**OperationResponseGridVOKnownClientVO**](OperationResponseGridVOKnownClientVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridPastConnections

> OperationResponseGridVOClientHistoryInfo GetGridPastConnections(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsLastSeen(sortsLastSeen).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersGuest(filtersGuest).SearchKey(searchKey).Execute()

Get client past connection list



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
	sortsLastSeen := "sortsLastSeen_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersTimeStart := "filtersTimeStart_example" // string | Filter query parameters, support field time range: start timestamp (ms). (optional)
	filtersTimeEnd := "filtersTimeEnd_example" // string | Filter query parameters, support field time range: end timestamp (ms). (optional)
	filtersGuest := "filtersGuest_example" // string | Filter query parameters, support field guest: true/false. (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name,mac,ssid. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetGridPastConnections(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsLastSeen(sortsLastSeen).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).FiltersGuest(filtersGuest).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetGridPastConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPastConnections`: OperationResponseGridVOClientHistoryInfo
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetGridPastConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridPastConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsLastSeen** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersTimeStart** | **string** | Filter query parameters, support field time range: start timestamp (ms). | 
 **filtersTimeEnd** | **string** | Filter query parameters, support field time range: end timestamp (ms). | 
 **filtersGuest** | **string** | Filter query parameters, support field guest: true/false. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name,mac,ssid. | 

### Return type

[**OperationResponseGridVOClientHistoryInfo**](OperationResponseGridVOClientHistoryInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongestClientUptime

> OperationResponseClientSummaryVO GetLongestClientUptime(ctx, omadacId, siteId).Execute()

Get longest client uptime.



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
	resp, r, err := apiClient.ClientInsightAPI.GetLongestClientUptime(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetLongestClientUptime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLongestClientUptime`: OperationResponseClientSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetLongestClientUptime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongestClientUptimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseClientSummaryVO**](OperationResponseClientSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMostActiveClients

> OperationResponseListActiveClientVO GetMostActiveClients(ctx, omadacId, siteId).Execute()

Get most active client.



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
	resp, r, err := apiClient.ClientInsightAPI.GetMostActiveClients(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetMostActiveClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMostActiveClients`: OperationResponseListActiveClientVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetMostActiveClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMostActiveClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListActiveClientVO**](OperationResponseListActiveClientVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspDashboardOverall

> OperationResponseMspClientOverallVO GetMspDashboardOverall(ctx, mspId).Execute()

Get the msp overview diagram of client.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetMspDashboardOverall(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetMspDashboardOverall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspDashboardOverall`: OperationResponseMspClientOverallVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetMspDashboardOverall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspDashboardOverallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseMspClientOverallVO**](OperationResponseMspClientOverallVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPastClientNum

> OperationResponseClientSummaryVO GetPastClientNum(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get past client number.



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetPastClientNum(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetPastClientNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPastClientNum`: OperationResponseClientSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetPastClientNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPastClientNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseClientSummaryVO**](OperationResponseClientSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackClientStat

> OperationResponseListDeviceClientStatVO GetStackClientStat(ctx, omadacId, siteId, stackId).DeviceClientStatQueryOpenApiVO(deviceClientStatQueryOpenApiVO).Execute()

Get stack client stat.



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
	deviceClientStatQueryOpenApiVO := *openapiclient.NewDeviceClientStatQueryOpenApiVO([]string{"Attrs_example"}, int64(123), int64(123)) // DeviceClientStatQueryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInsightAPI.GetStackClientStat(context.Background(), omadacId, siteId, stackId).DeviceClientStatQueryOpenApiVO(deviceClientStatQueryOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInsightAPI.GetStackClientStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackClientStat`: OperationResponseListDeviceClientStatVO
	fmt.Fprintf(os.Stdout, "Response from `ClientInsightAPI.GetStackClientStat`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStackClientStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deviceClientStatQueryOpenApiVO** | [**DeviceClientStatQueryOpenApiVO**](DeviceClientStatQueryOpenApiVO.md) |  | 

### Return type

[**OperationResponseListDeviceClientStatVO**](OperationResponseListDeviceClientStatVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


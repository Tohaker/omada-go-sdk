# NetworkReportAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailNetworkReportLater**](NetworkReportAPI.md#emailnetworkreportlater) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networkReport/oneTime | Email Network Report Later
[**EmailNetworkReportSchedule**](NetworkReportAPI.md#emailnetworkreportschedule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/networkReport/schedule | Email Network Report Schedule



## EmailNetworkReportLater

> OperationResponse EmailNetworkReportLater(ctx, omadacId, siteId).NetworkReportScheduleLaterOpenApiVO(networkReportScheduleLaterOpenApiVO).Execute()

Email Network Report Later



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
	networkReportScheduleLaterOpenApiVO := *openapiclient.NewNetworkReportScheduleLaterOpenApiVO("Cards_example", []string{"EmailList_example"}, false, int64(123), int32(123), int32(123), "ReportName_example", int32(123), int64(123), int32(123), int64(123)) // NetworkReportScheduleLaterOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkReportAPI.EmailNetworkReportLater(context.Background(), omadacId, siteId).NetworkReportScheduleLaterOpenApiVO(networkReportScheduleLaterOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkReportAPI.EmailNetworkReportLater``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailNetworkReportLater`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkReportAPI.EmailNetworkReportLater`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailNetworkReportLaterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkReportScheduleLaterOpenApiVO** | [**NetworkReportScheduleLaterOpenApiVO**](NetworkReportScheduleLaterOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailNetworkReportSchedule

> OperationResponse EmailNetworkReportSchedule(ctx, omadacId, siteId).NetworkReportScheduleOpenApiVO(networkReportScheduleOpenApiVO).Execute()

Email Network Report Schedule



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
	networkReportScheduleOpenApiVO := *openapiclient.NewNetworkReportScheduleOpenApiVO("Cards_example", []string{"EmailList_example"}, false, int32(123), int32(123), "ReportName_example", int32(123), int32(123), int32(123)) // NetworkReportScheduleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkReportAPI.EmailNetworkReportSchedule(context.Background(), omadacId, siteId).NetworkReportScheduleOpenApiVO(networkReportScheduleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkReportAPI.EmailNetworkReportSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailNetworkReportSchedule`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkReportAPI.EmailNetworkReportSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailNetworkReportScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkReportScheduleOpenApiVO** | [**NetworkReportScheduleOpenApiVO**](NetworkReportScheduleOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


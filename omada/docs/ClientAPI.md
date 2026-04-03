# ClientAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchModifyClientSetting**](ClientAPI.md#batchmodifyclientsetting) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/config | Batch config clients
[**BlockClient**](ClientAPI.md#blockclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/block | Block the client
[**DeleteClient**](ClientAPI.md#deleteclient) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac} | Delete client
[**DeleteClients**](ClientAPI.md#deleteclients) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/delete | Batch delete clients
[**DisconnectClient**](ClientAPI.md#disconnectclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/disconnect | Disconnect the client
[**ExportClient**](ClientAPI.md#exportclient) | **Post** /openapi/v1/{omadacId}/files/sites/{siteId}/clients/export | Export client list
[**ExportClientListGlobalByCloudAccess**](ClientAPI.md#exportclientlistglobalbycloudaccess) | **Post** /openapi/v1/{omadacId}/files/client-list | Export global client list.
[**GetClientCorrectionList**](ClientAPI.md#getclientcorrectionlist) | **Get** /openapi/v1/{omadacId}/correction-list | Get client correction options list
[**GetClientDetail**](ClientAPI.md#getclientdetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac} | Get client info
[**GetClientDetailStat5Min**](ClientAPI.md#getclientdetailstat5min) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/client-stat-detail/{clientMac}/5Min | Get client statistical data details at a 5-minute interval.
[**GetClientDetailStatDaily**](ClientAPI.md#getclientdetailstatdaily) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/client-stat-detail/{clientMac}/daily | Get client statistical data details at a daily interval.
[**GetClientDetailStatHourly**](ClientAPI.md#getclientdetailstathourly) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/client-stat-detail/{clientMac}/hourly | Get client statistical data details at a hourly interval.
[**GetClientFilteringOptions**](ClientAPI.md#getclientfilteringoptions) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/clients/search-fields-options | Get client list filtering options
[**GetClientHistoryDataEnable**](ClientAPI.md#getclienthistorydataenable) | **Get** /openapi/v1/{omadacId}/controller/client/history-enable | Get History data retention config.
[**GetClientJourney**](ClientAPI.md#getclientjourney) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/client-connection | Get client connection histories
[**GetClientTimeline**](ClientAPI.md#getclienttimeline) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/client-timeline | Get client timeline events
[**GetClientTopology**](ClientAPI.md#getclienttopology) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/client-link-topology | Get client link topology
[**GetGridActiveClients**](ClientAPI.md#getgridactiveclients) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/clients | Get client list
[**GetGridAllClients**](ClientAPI.md#getgridallclients) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/clients | Get all client list
[**GetGridClientHistory**](ClientAPI.md#getgridclienthistory) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/client-history |  Get Client history.
[**GetVigiDetailStat5Min**](ClientAPI.md#getvigidetailstat5min) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vigi-stat-detail/{vigiMac}/5Min | Get VIGI device statistical data details at a 5-minute interval.
[**GetVigiDetailStatDaily**](ClientAPI.md#getvigidetailstatdaily) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vigi-stat-detail/{vigiMac}/daily | Get VIGI device statistical data details at a daily interval.
[**GetVigiDetailStatHourly**](ClientAPI.md#getvigidetailstathourly) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vigi-stat-detail/{vigiMac}/hourly | Get VIGI device statistical data details at a hourly interval.
[**GetVigiJourney**](ClientAPI.md#getvigijourney) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vigis/{vigiMac}/vigi-connection | Get VIGI device connection histories
[**GetVigiTimeline**](ClientAPI.md#getvigitimeline) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vigis/{vigiMac}/vigi-timeline | Get VIGI device timeline events
[**GetVigiTopology**](ClientAPI.md#getvigitopology) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vigis/{vigiMac}/vigi-link-topology | Get VIGI device link topology
[**GlobalClientStatByDevice**](ClientAPI.md#globalclientstatbydevice) | **Post** /openapi/v1/{omadacId}/clients/stat/devices | Get global client statistics by device.
[**GlobalExportAllClientList**](ClientAPI.md#globalexportallclientlist) | **Post** /openapi/v1/{omadacId}/files/all-client-list | Export all client list in GLOBAL view
[**ModifyClientIpSetting**](ClientAPI.md#modifyclientipsetting) | **Patch** /openapi/v1/{omadacId}/network/sites/{siteId}/cmd/clients/{clientMac}/update-ipSetting | Set ip setting for given client
[**MspClientStatByDevice**](ClientAPI.md#mspclientstatbydevice) | **Post** /openapi/v1/msp/{mspId}/clients/stat/devices | Get msp client statistics by device.
[**RebootClient**](ClientAPI.md#rebootclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/reboot | Reboot the client
[**ReconnectClient**](ClientAPI.md#reconnectclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/reconnect | Reconnect the client
[**UnblockClient**](ClientAPI.md#unblockclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/unblock | Unblock the client
[**UpdateClientLockToApSetting**](ClientAPI.md#updateclientlocktoapsetting) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/lock-to-ap | Lock the given client to aps
[**UpdateClientName**](ClientAPI.md#updateclientname) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/name | Set name for given client
[**UpdateClientRateLimitSetting**](ClientAPI.md#updateclientratelimitsetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/clients/{clientMac}/ratelimit | Set ratelimit setting for given client



## BatchModifyClientSetting

> OperationResponseClientDetail BatchModifyClientSetting(ctx, omadacId, siteId).ClientBatchSetting(clientBatchSetting).Execute()

Batch config clients



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
	clientBatchSetting := *openapiclient.NewClientBatchSetting() // ClientBatchSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.BatchModifyClientSetting(context.Background(), omadacId, siteId).ClientBatchSetting(clientBatchSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.BatchModifyClientSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyClientSetting`: OperationResponseClientDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.BatchModifyClientSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyClientSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientBatchSetting** | [**ClientBatchSetting**](ClientBatchSetting.md) |  | 

### Return type

[**OperationResponseClientDetail**](OperationResponseClientDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockClient

> OperationResponseWithoutResult BlockClient(ctx, omadacId, siteId, clientMac).Execute()

Block the client



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
	clientMac := "clientMac_example" // string | Client MAC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.BlockClient(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.BlockClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.BlockClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockClientRequest struct via the builder pattern


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


## DeleteClient

> OperationResponseWithoutResult DeleteClient(ctx, omadacId, siteId, clientMac).Execute()

Delete client



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
	clientMac := "clientMac_example" // string | Client MAC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.DeleteClient(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.DeleteClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRequest struct via the builder pattern


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


## DeleteClients

> OperationResponseClientDetail DeleteClients(ctx, omadacId, siteId).ClientDeleteFilter(clientDeleteFilter).Execute()

Batch delete clients



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
	clientDeleteFilter := *openapiclient.NewClientDeleteFilter() // ClientDeleteFilter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.DeleteClients(context.Background(), omadacId, siteId).ClientDeleteFilter(clientDeleteFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClients`: OperationResponseClientDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.DeleteClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientDeleteFilter** | [**ClientDeleteFilter**](ClientDeleteFilter.md) |  | 

### Return type

[**OperationResponseClientDetail**](OperationResponseClientDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectClient

> OperationResponseWithoutResult DisconnectClient(ctx, omadacId, siteId, clientMac).Execute()

Disconnect the client



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
	clientMac := "clientMac_example" // string | Client MAC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.DisconnectClient(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DisconnectClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.DisconnectClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectClientRequest struct via the builder pattern


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


## ExportClient

> OperationResponseWithoutResult ExportClient(ctx, omadacId, siteId).ClientExportOpenApiVO(clientExportOpenApiVO).Execute()

Export client list



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
	clientExportOpenApiVO := *openapiclient.NewClientExportOpenApiVO(int32(123), []int32{int32(123)}) // ClientExportOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ExportClient(context.Background(), omadacId, siteId).ClientExportOpenApiVO(clientExportOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ExportClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ExportClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientExportOpenApiVO** | [**ClientExportOpenApiVO**](ClientExportOpenApiVO.md) |  | 

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


## ExportClientListGlobalByCloudAccess

> OperationResponse ExportClientListGlobalByCloudAccess(ctx, omadacId).ExportClientListOpenApiVO(exportClientListOpenApiVO).Execute()

Export global client list.



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
	exportClientListOpenApiVO := *openapiclient.NewExportClientListOpenApiVO(int32(123), int32(123), []string{"SiteIds_example"}) // ExportClientListOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ExportClientListGlobalByCloudAccess(context.Background(), omadacId).ExportClientListOpenApiVO(exportClientListOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ExportClientListGlobalByCloudAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportClientListGlobalByCloudAccess`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ExportClientListGlobalByCloudAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportClientListGlobalByCloudAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportClientListOpenApiVO** | [**ExportClientListOpenApiVO**](ExportClientListOpenApiVO.md) |  | 

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


## GetClientCorrectionList

> OperationResponseClientCorrectionOptionListVO GetClientCorrectionList(ctx, omadacId).Execute()

Get client correction options list



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
	resp, r, err := apiClient.ClientAPI.GetClientCorrectionList(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientCorrectionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientCorrectionList`: OperationResponseClientCorrectionOptionListVO
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientCorrectionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCorrectionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseClientCorrectionOptionListVO**](OperationResponseClientCorrectionOptionListVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientDetail

> OperationResponseClientDetail GetClientDetail(ctx, omadacId, siteId, clientMac).Execute()

Get client info



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
	clientMac := "clientMac_example" // string | Client MAC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetClientDetail(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientDetail`: OperationResponseClientDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseClientDetail**](OperationResponseClientDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientDetailStat5Min

> OperationResponseClientStatisticalDataDetail GetClientDetailStat5Min(ctx, omadacId, siteId, clientMac).ClientStatQuery(clientStatQuery).Execute()

Get client statistical data details at a 5-minute interval.



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
	clientMac := "clientMac_example" // string | Client MAC
	clientStatQuery := *openapiclient.NewClientStatQuery(int64(123), int64(123)) // ClientStatQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetClientDetailStat5Min(context.Background(), omadacId, siteId, clientMac).ClientStatQuery(clientStatQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientDetailStat5Min``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientDetailStat5Min`: OperationResponseClientStatisticalDataDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientDetailStat5Min`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientDetailStat5MinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientStatQuery** | [**ClientStatQuery**](ClientStatQuery.md) |  | 

### Return type

[**OperationResponseClientStatisticalDataDetail**](OperationResponseClientStatisticalDataDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientDetailStatDaily

> OperationResponseClientStatisticalDataDetail GetClientDetailStatDaily(ctx, omadacId, siteId, clientMac).ClientStatQuery(clientStatQuery).Execute()

Get client statistical data details at a daily interval.



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
	clientMac := "clientMac_example" // string | Client MAC
	clientStatQuery := *openapiclient.NewClientStatQuery(int64(123), int64(123)) // ClientStatQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetClientDetailStatDaily(context.Background(), omadacId, siteId, clientMac).ClientStatQuery(clientStatQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientDetailStatDaily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientDetailStatDaily`: OperationResponseClientStatisticalDataDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientDetailStatDaily`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientDetailStatDailyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientStatQuery** | [**ClientStatQuery**](ClientStatQuery.md) |  | 

### Return type

[**OperationResponseClientStatisticalDataDetail**](OperationResponseClientStatisticalDataDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientDetailStatHourly

> OperationResponseClientStatisticalDataDetail GetClientDetailStatHourly(ctx, omadacId, siteId, clientMac).ClientStatQuery(clientStatQuery).Execute()

Get client statistical data details at a hourly interval.



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
	clientMac := "clientMac_example" // string | Client MAC
	clientStatQuery := *openapiclient.NewClientStatQuery(int64(123), int64(123)) // ClientStatQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetClientDetailStatHourly(context.Background(), omadacId, siteId, clientMac).ClientStatQuery(clientStatQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientDetailStatHourly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientDetailStatHourly`: OperationResponseClientStatisticalDataDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientDetailStatHourly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientDetailStatHourlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientStatQuery** | [**ClientStatQuery**](ClientStatQuery.md) |  | 

### Return type

[**OperationResponseClientStatisticalDataDetail**](OperationResponseClientStatisticalDataDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientFilteringOptions

> OperationResponseClientFilteringOptions GetClientFilteringOptions(ctx, omadacId, siteId).Execute()

Get client list filtering options



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
	resp, r, err := apiClient.ClientAPI.GetClientFilteringOptions(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientFilteringOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientFilteringOptions`: OperationResponseClientFilteringOptions
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientFilteringOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientFilteringOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseClientFilteringOptions**](OperationResponseClientFilteringOptions.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientHistoryDataEnable

> OperationResponseOmadacClientSettingOpenApiVO GetClientHistoryDataEnable(ctx, omadacId).Execute()

Get History data retention config.



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
	resp, r, err := apiClient.ClientAPI.GetClientHistoryDataEnable(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientHistoryDataEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientHistoryDataEnable`: OperationResponseOmadacClientSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientHistoryDataEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientHistoryDataEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseOmadacClientSettingOpenApiVO**](OperationResponseOmadacClientSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientJourney

> OperationResponseClientConnectionHistories GetClientJourney(ctx, omadacId, siteId, clientMac).Start(start).End(end).Execute()

Get client connection histories



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
	clientMac := "clientMac_example" // string | Client MAC
	start := int64(789) // int64 | Start timestamp, unit: ms
	end := int64(789) // int64 | End timestamp, unit: ms

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetClientJourney(context.Background(), omadacId, siteId, clientMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientJourney``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientJourney`: OperationResponseClientConnectionHistories
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientJourney`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientJourneyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, unit: ms | 
 **end** | **int64** | End timestamp, unit: ms | 

### Return type

[**OperationResponseClientConnectionHistories**](OperationResponseClientConnectionHistories.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientTimeline

> OperationResponseListActivityRecordsOfAClientSSingleConnections GetClientTimeline(ctx, omadacId, siteId, clientMac).Start(start).End(end).Type_(type_).Execute()

Get client timeline events



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
	clientMac := "clientMac_example" // string | Client MAC
	start := int64(789) // int64 | Start timestamp, unit: ms
	end := int64(789) // int64 | End timestamp, unit: ms
	type_ := int32(56) // int32 | Query type, 0: Connection Timeline; 1: Association; 2: Roaming; 3: Disconnection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetClientTimeline(context.Background(), omadacId, siteId, clientMac).Start(start).End(end).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientTimeline`: OperationResponseListActivityRecordsOfAClientSSingleConnections
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, unit: ms | 
 **end** | **int64** | End timestamp, unit: ms | 
 **type_** | **int32** | Query type, 0: Connection Timeline; 1: Association; 2: Roaming; 3: Disconnection. | 

### Return type

[**OperationResponseListActivityRecordsOfAClientSSingleConnections**](OperationResponseListActivityRecordsOfAClientSSingleConnections.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientTopology

> OperationResponseListClientTopologyNodesInfo GetClientTopology(ctx, omadacId, siteId, clientMac).Execute()

Get client link topology



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
	clientMac := "clientMac_example" // string | Client MAC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetClientTopology(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientTopology`: OperationResponseListClientTopologyNodesInfo
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListClientTopologyNodesInfo**](OperationResponseListClientTopologyNodesInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridActiveClients

> OperationResponseClientGridVOClientInfo GetGridActiveClients(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsMac(sortsMac).SortsIp(sortsIp).FiltersWireless(filtersWireless).FiltersRadioId(filtersRadioId).FiltersApMac(filtersApMac).FiltersSwitchMac(filtersSwitchMac).FiltersGatewayMac(filtersGatewayMac).SearchKey(searchKey).Execute()

Get client list



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
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsMac := "sortsMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersWireless := "filtersWireless_example" // string | Filter query parameters, support field wireless: true/false. (optional)
	filtersRadioId := "filtersRadioId_example" // string | Filter query parameters, support field radioId: 0: 2G, 1: 5G1, 2: 5G2, 3: 6G (optional)
	filtersApMac := "filtersApMac_example" // string | Filter query parameters, support field ap mac (optional)
	filtersSwitchMac := "filtersSwitchMac_example" // string | Filter query parameters, support field switch mac (optional)
	filtersGatewayMac := "filtersGatewayMac_example" // string | Filter query parameters, support field gateway mac (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field clientName,clientMac,ip,channel,ssid,apName,apMac,switchMac,switchName,gatewayMac,gatewayName. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetGridActiveClients(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsMac(sortsMac).SortsIp(sortsIp).FiltersWireless(filtersWireless).FiltersRadioId(filtersRadioId).FiltersApMac(filtersApMac).FiltersSwitchMac(filtersSwitchMac).FiltersGatewayMac(filtersGatewayMac).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetGridActiveClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridActiveClients`: OperationResponseClientGridVOClientInfo
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetGridActiveClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridActiveClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersWireless** | **string** | Filter query parameters, support field wireless: true/false. | 
 **filtersRadioId** | **string** | Filter query parameters, support field radioId: 0: 2G, 1: 5G1, 2: 5G2, 3: 6G | 
 **filtersApMac** | **string** | Filter query parameters, support field ap mac | 
 **filtersSwitchMac** | **string** | Filter query parameters, support field switch mac | 
 **filtersGatewayMac** | **string** | Filter query parameters, support field gateway mac | 
 **searchKey** | **string** | Fuzzy query parameters, support field clientName,clientMac,ip,channel,ssid,apName,apMac,switchMac,switchName,gatewayMac,gatewayName. | 

### Return type

[**OperationResponseClientGridVOClientInfo**](OperationResponseClientGridVOClientInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridAllClients

> OperationResponseClientGridVOClientInfo GetGridAllClients(ctx, omadacId, siteId).ClientQueryDataOpenApiVO(clientQueryDataOpenApiVO).Execute()

Get all client list



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
	clientQueryDataOpenApiVO := *openapiclient.NewClientQueryDataOpenApiVO(int32(123), int32(123)) // ClientQueryDataOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetGridAllClients(context.Background(), omadacId, siteId).ClientQueryDataOpenApiVO(clientQueryDataOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetGridAllClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAllClients`: OperationResponseClientGridVOClientInfo
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetGridAllClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAllClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientQueryDataOpenApiVO** | [**ClientQueryDataOpenApiVO**](ClientQueryDataOpenApiVO.md) |  | 

### Return type

[**OperationResponseClientGridVOClientInfo**](OperationResponseClientGridVOClientInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridClientHistory

> OperationResponseGridVOClientHistoryVO GetGridClientHistory(ctx, omadacId, siteId, clientMac).Page(page).PageSize(pageSize).SortsName(sortsName).SortsMac(sortsMac).SortsIp(sortsIp).SortsAuthType(sortsAuthType).SortsSsidOrNetwork(sortsSsidOrNetwork).SortsDownload(sortsDownload).SortsUpload(sortsUpload).SortsStatus(sortsStatus).SortsStart(sortsStart).SortsEnd(sortsEnd).SortsDuration(sortsDuration).SearchKey(searchKey).Execute()

 Get Client history.



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
	clientMac := "clientMac_example" // string | Client MAC
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsMac := "sortsMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsAuthType := "sortsAuthType_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsSsidOrNetwork := "sortsSsidOrNetwork_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsDownload := "sortsDownload_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsUpload := "sortsUpload_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStart := "sortsStart_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsEnd := "sortsEnd_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsDuration := "sortsDuration_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field client_mac,client_name,voucher.code,local_user.user_name,form_name,auth_admin,ssid_name,network_name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetGridClientHistory(context.Background(), omadacId, siteId, clientMac).Page(page).PageSize(pageSize).SortsName(sortsName).SortsMac(sortsMac).SortsIp(sortsIp).SortsAuthType(sortsAuthType).SortsSsidOrNetwork(sortsSsidOrNetwork).SortsDownload(sortsDownload).SortsUpload(sortsUpload).SortsStatus(sortsStatus).SortsStart(sortsStart).SortsEnd(sortsEnd).SortsDuration(sortsDuration).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetGridClientHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridClientHistory`: OperationResponseGridVOClientHistoryVO
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetGridClientHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridClientHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsAuthType** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsSsidOrNetwork** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsDownload** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsUpload** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStatus** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStart** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsEnd** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsDuration** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field client_mac,client_name,voucher.code,local_user.user_name,form_name,auth_admin,ssid_name,network_name | 

### Return type

[**OperationResponseGridVOClientHistoryVO**](OperationResponseGridVOClientHistoryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVigiDetailStat5Min

> OperationResponseClientStatisticalDataDetail GetVigiDetailStat5Min(ctx, omadacId, siteId, vigiMac).ClientStatQuery(clientStatQuery).Execute()

Get VIGI device statistical data details at a 5-minute interval.



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
	vigiMac := "vigiMac_example" // string | VIGI device MAC
	clientStatQuery := *openapiclient.NewClientStatQuery(int64(123), int64(123)) // ClientStatQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetVigiDetailStat5Min(context.Background(), omadacId, siteId, vigiMac).ClientStatQuery(clientStatQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetVigiDetailStat5Min``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVigiDetailStat5Min`: OperationResponseClientStatisticalDataDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetVigiDetailStat5Min`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vigiMac** | **string** | VIGI device MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVigiDetailStat5MinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientStatQuery** | [**ClientStatQuery**](ClientStatQuery.md) |  | 

### Return type

[**OperationResponseClientStatisticalDataDetail**](OperationResponseClientStatisticalDataDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVigiDetailStatDaily

> OperationResponseClientStatisticalDataDetail GetVigiDetailStatDaily(ctx, omadacId, siteId, vigiMac).ClientStatQuery(clientStatQuery).Execute()

Get VIGI device statistical data details at a daily interval.



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
	vigiMac := "vigiMac_example" // string | VIGI device MAC
	clientStatQuery := *openapiclient.NewClientStatQuery(int64(123), int64(123)) // ClientStatQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetVigiDetailStatDaily(context.Background(), omadacId, siteId, vigiMac).ClientStatQuery(clientStatQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetVigiDetailStatDaily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVigiDetailStatDaily`: OperationResponseClientStatisticalDataDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetVigiDetailStatDaily`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vigiMac** | **string** | VIGI device MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVigiDetailStatDailyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientStatQuery** | [**ClientStatQuery**](ClientStatQuery.md) |  | 

### Return type

[**OperationResponseClientStatisticalDataDetail**](OperationResponseClientStatisticalDataDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVigiDetailStatHourly

> OperationResponseClientStatisticalDataDetail GetVigiDetailStatHourly(ctx, omadacId, siteId, vigiMac).ClientStatQuery(clientStatQuery).Execute()

Get VIGI device statistical data details at a hourly interval.



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
	vigiMac := "vigiMac_example" // string | VIGI device MAC
	clientStatQuery := *openapiclient.NewClientStatQuery(int64(123), int64(123)) // ClientStatQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetVigiDetailStatHourly(context.Background(), omadacId, siteId, vigiMac).ClientStatQuery(clientStatQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetVigiDetailStatHourly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVigiDetailStatHourly`: OperationResponseClientStatisticalDataDetail
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetVigiDetailStatHourly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vigiMac** | **string** | VIGI device MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVigiDetailStatHourlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientStatQuery** | [**ClientStatQuery**](ClientStatQuery.md) |  | 

### Return type

[**OperationResponseClientStatisticalDataDetail**](OperationResponseClientStatisticalDataDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVigiJourney

> OperationResponseClientConnectionHistories GetVigiJourney(ctx, omadacId, siteId, vigiMac).Start(start).End(end).Execute()

Get VIGI device connection histories



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
	vigiMac := "vigiMac_example" // string | VIGI device MAC
	start := int64(789) // int64 | Start timestamp, unit: ms
	end := int64(789) // int64 | End timestamp, unit: ms

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetVigiJourney(context.Background(), omadacId, siteId, vigiMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetVigiJourney``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVigiJourney`: OperationResponseClientConnectionHistories
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetVigiJourney`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vigiMac** | **string** | VIGI device MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVigiJourneyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, unit: ms | 
 **end** | **int64** | End timestamp, unit: ms | 

### Return type

[**OperationResponseClientConnectionHistories**](OperationResponseClientConnectionHistories.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVigiTimeline

> OperationResponseListActivityRecordsOfAClientSSingleConnections GetVigiTimeline(ctx, omadacId, siteId, vigiMac).Start(start).End(end).Type_(type_).Execute()

Get VIGI device timeline events



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
	vigiMac := "vigiMac_example" // string | VIGI device MAC
	start := int64(789) // int64 | Start timestamp, unit: ms
	end := int64(789) // int64 | End timestamp, unit: ms
	type_ := int32(56) // int32 | Query type, 0: Connection Timeline; 1: Association; 2: Roaming; 3: Disconnection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetVigiTimeline(context.Background(), omadacId, siteId, vigiMac).Start(start).End(end).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetVigiTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVigiTimeline`: OperationResponseListActivityRecordsOfAClientSSingleConnections
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetVigiTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vigiMac** | **string** | VIGI device MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVigiTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, unit: ms | 
 **end** | **int64** | End timestamp, unit: ms | 
 **type_** | **int32** | Query type, 0: Connection Timeline; 1: Association; 2: Roaming; 3: Disconnection. | 

### Return type

[**OperationResponseListActivityRecordsOfAClientSSingleConnections**](OperationResponseListActivityRecordsOfAClientSSingleConnections.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVigiTopology

> OperationResponseListClientTopologyNodesInfo GetVigiTopology(ctx, omadacId, siteId, vigiMac).Execute()

Get VIGI device link topology



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
	vigiMac := "vigiMac_example" // string | VIGI device MAC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetVigiTopology(context.Background(), omadacId, siteId, vigiMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetVigiTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVigiTopology`: OperationResponseListClientTopologyNodesInfo
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetVigiTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vigiMac** | **string** | VIGI device MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVigiTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListClientTopologyNodesInfo**](OperationResponseListClientTopologyNodesInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalClientStatByDevice

> OperationResponseListDeviceClientNumOpenApiVO GlobalClientStatByDevice(ctx, omadacId).GlobalDevicesQueryOpenApiVO(globalDevicesQueryOpenApiVO).Execute()

Get global client statistics by device.



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
	globalDevicesQueryOpenApiVO := *openapiclient.NewGlobalDevicesQueryOpenApiVO([]openapiclient.GlobalDeviceItem{*openapiclient.NewGlobalDeviceItem("Mac_example", "SiteId_example")}) // GlobalDevicesQueryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GlobalClientStatByDevice(context.Background(), omadacId).GlobalDevicesQueryOpenApiVO(globalDevicesQueryOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GlobalClientStatByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalClientStatByDevice`: OperationResponseListDeviceClientNumOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GlobalClientStatByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalClientStatByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalDevicesQueryOpenApiVO** | [**GlobalDevicesQueryOpenApiVO**](GlobalDevicesQueryOpenApiVO.md) |  | 

### Return type

[**OperationResponseListDeviceClientNumOpenApiVO**](OperationResponseListDeviceClientNumOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalExportAllClientList

> OperationResponse GlobalExportAllClientList(ctx, omadacId).MultiSiteClientExportOpenApiVO(multiSiteClientExportOpenApiVO).Execute()

Export all client list in GLOBAL view



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
	multiSiteClientExportOpenApiVO := *openapiclient.NewMultiSiteClientExportOpenApiVO(int32(123), int32(123), "SelectType_example", []string{"SiteIds_example"}) // MultiSiteClientExportOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GlobalExportAllClientList(context.Background(), omadacId).MultiSiteClientExportOpenApiVO(multiSiteClientExportOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GlobalExportAllClientList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalExportAllClientList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GlobalExportAllClientList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalExportAllClientListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multiSiteClientExportOpenApiVO** | [**MultiSiteClientExportOpenApiVO**](MultiSiteClientExportOpenApiVO.md) |  | 

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


## ModifyClientIpSetting

> OperationResponseWithoutResult ModifyClientIpSetting(ctx, omadacId, siteId, clientMac).ClientIpSetting(clientIpSetting).Execute()

Set ip setting for given client



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
	clientMac := "clientMac_example" // string | Client MAC
	clientIpSetting := *openapiclient.NewClientIpSetting(false) // ClientIpSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ModifyClientIpSetting(context.Background(), omadacId, siteId, clientMac).ClientIpSetting(clientIpSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ModifyClientIpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyClientIpSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ModifyClientIpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyClientIpSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientIpSetting** | [**ClientIpSetting**](ClientIpSetting.md) |  | 

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


## MspClientStatByDevice

> OperationResponseListMspDeviceClientNumOpenApiVO MspClientStatByDevice(ctx, mspId).MspDevicesQueryOpenApiVO(mspDevicesQueryOpenApiVO).Execute()

Get msp client statistics by device.



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
	mspDevicesQueryOpenApiVO := *openapiclient.NewMspDevicesQueryOpenApiVO([]openapiclient.MspDeviceItem{*openapiclient.NewMspDeviceItem("CustomerId_example", "Mac_example", "SiteId_example")}) // MspDevicesQueryOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.MspClientStatByDevice(context.Background(), mspId).MspDevicesQueryOpenApiVO(mspDevicesQueryOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.MspClientStatByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MspClientStatByDevice`: OperationResponseListMspDeviceClientNumOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.MspClientStatByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMspClientStatByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mspDevicesQueryOpenApiVO** | [**MspDevicesQueryOpenApiVO**](MspDevicesQueryOpenApiVO.md) |  | 

### Return type

[**OperationResponseListMspDeviceClientNumOpenApiVO**](OperationResponseListMspDeviceClientNumOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootClient

> OperationResponse RebootClient(ctx, omadacId, siteId, clientMac).ClientReboot(clientReboot).Execute()

Reboot the client



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
	clientMac := "clientMac_example" // string | Client MAC
	clientReboot := *openapiclient.NewClientReboot() // ClientReboot | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.RebootClient(context.Background(), omadacId, siteId, clientMac).ClientReboot(clientReboot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.RebootClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootClient`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.RebootClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientReboot** | [**ClientReboot**](ClientReboot.md) |  | 

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


## ReconnectClient

> OperationResponseWithoutResult ReconnectClient(ctx, omadacId, siteId, clientMac).Execute()

Reconnect the client



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
	clientMac := "clientMac_example" // string | Client MAC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.ReconnectClient(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.ReconnectClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReconnectClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.ReconnectClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiReconnectClientRequest struct via the builder pattern


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


## UnblockClient

> OperationResponseWithoutResult UnblockClient(ctx, omadacId, siteId, clientMac).Execute()

Unblock the client



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
	clientMac := "clientMac_example" // string | Client MAC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UnblockClient(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UnblockClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnblockClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UnblockClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnblockClientRequest struct via the builder pattern


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


## UpdateClientLockToApSetting

> OperationResponseWithoutResult UpdateClientLockToApSetting(ctx, omadacId, siteId, clientMac).ClientLockToApMacListSetting(clientLockToApMacListSetting).Execute()

Lock the given client to aps



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
	clientMac := "clientMac_example" // string | Client MAC
	clientLockToApMacListSetting := *openapiclient.NewClientLockToApMacListSetting(false) // ClientLockToApMacListSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UpdateClientLockToApSetting(context.Background(), omadacId, siteId, clientMac).ClientLockToApMacListSetting(clientLockToApMacListSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UpdateClientLockToApSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClientLockToApSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UpdateClientLockToApSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientLockToApSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientLockToApMacListSetting** | [**ClientLockToApMacListSetting**](ClientLockToApMacListSetting.md) |  | 

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


## UpdateClientName

> OperationResponseWithoutResult UpdateClientName(ctx, omadacId, siteId, clientMac).ClientNameSetting(clientNameSetting).Execute()

Set name for given client



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
	clientMac := "clientMac_example" // string | Client MAC
	clientNameSetting := *openapiclient.NewClientNameSetting("Name_example") // ClientNameSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UpdateClientName(context.Background(), omadacId, siteId, clientMac).ClientNameSetting(clientNameSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UpdateClientName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClientName`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UpdateClientName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientNameSetting** | [**ClientNameSetting**](ClientNameSetting.md) |  | 

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


## UpdateClientRateLimitSetting

> OperationResponseWithoutResult UpdateClientRateLimitSetting(ctx, omadacId, siteId, clientMac).ClientRateLimitSetting(clientRateLimitSetting).Execute()

Set ratelimit setting for given client



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
	clientMac := "clientMac_example" // string | Client MAC
	clientRateLimitSetting := *openapiclient.NewClientRateLimitSetting(int32(123)) // ClientRateLimitSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.UpdateClientRateLimitSetting(context.Background(), omadacId, siteId, clientMac).ClientRateLimitSetting(clientRateLimitSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.UpdateClientRateLimitSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClientRateLimitSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.UpdateClientRateLimitSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientRateLimitSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientRateLimitSetting** | [**ClientRateLimitSetting**](ClientRateLimitSetting.md) |  | 

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


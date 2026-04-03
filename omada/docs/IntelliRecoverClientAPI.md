# IntelliRecoverClientAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMonitorClients**](IntelliRecoverClientAPI.md#addmonitorclients) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/monitor-client/add | Add client into intelli recover client list
[**DeleteMonitorClients**](IntelliRecoverClientAPI.md#deletemonitorclients) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/monitor-client/delete | Delete the intelli recover client
[**GetEligibleClientList**](IntelliRecoverClientAPI.md#geteligibleclientlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/monitor-client/eligible-client-grids | Obtain the list of clients that can be monitored
[**GetGridMonitorClient**](IntelliRecoverClientAPI.md#getgridmonitorclient) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/monitor-client/grids | Get the intelli recover client list
[**RebootClientUplinkPoe**](IntelliRecoverClientAPI.md#rebootclientuplinkpoe) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/monitor-client/reboot-uplink-poe | Reboot the client uplinkDevice poe
[**VerifyMonitorClient**](IntelliRecoverClientAPI.md#verifymonitorclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/monitor-client/verify | Verify if the client is allowed to be monitored



## AddMonitorClients

> OperationResponseWithoutResult AddMonitorClients(ctx, omadacId, siteId).AddMonitorClientList(addMonitorClientList).Execute()

Add client into intelli recover client list



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
	addMonitorClientList := *openapiclient.NewAddMonitorClientList() // AddMonitorClientList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverClientAPI.AddMonitorClients(context.Background(), omadacId, siteId).AddMonitorClientList(addMonitorClientList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverClientAPI.AddMonitorClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMonitorClients`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverClientAPI.AddMonitorClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMonitorClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addMonitorClientList** | [**AddMonitorClientList**](AddMonitorClientList.md) |  | 

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


## DeleteMonitorClients

> OperationResponseWithoutResult DeleteMonitorClients(ctx, omadacId, siteId).DeleteIntelliRecoverClient(deleteIntelliRecoverClient).Execute()

Delete the intelli recover client



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
	deleteIntelliRecoverClient := *openapiclient.NewDeleteIntelliRecoverClient() // DeleteIntelliRecoverClient | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverClientAPI.DeleteMonitorClients(context.Background(), omadacId, siteId).DeleteIntelliRecoverClient(deleteIntelliRecoverClient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverClientAPI.DeleteMonitorClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMonitorClients`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverClientAPI.DeleteMonitorClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteIntelliRecoverClient** | [**DeleteIntelliRecoverClient**](DeleteIntelliRecoverClient.md) |  | 

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


## GetEligibleClientList

> OperationResponseGridVOMonitorClient GetEligibleClientList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Obtain the list of clients that can be monitored



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverClientAPI.GetEligibleClientList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverClientAPI.GetEligibleClientList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEligibleClientList`: OperationResponseGridVOMonitorClient
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverClientAPI.GetEligibleClientList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEligibleClientListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOMonitorClient**](OperationResponseGridVOMonitorClient.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridMonitorClient

> OperationResponseGridVOMonitorClient GetGridMonitorClient(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the intelli recover client list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverClientAPI.GetGridMonitorClient(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverClientAPI.GetGridMonitorClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridMonitorClient`: OperationResponseGridVOMonitorClient
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverClientAPI.GetGridMonitorClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridMonitorClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOMonitorClient**](OperationResponseGridVOMonitorClient.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootClientUplinkPoe

> OperationResponseWithoutResult RebootClientUplinkPoe(ctx, omadacId, siteId).MonitorClient(monitorClient).Execute()

Reboot the client uplinkDevice poe



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
	monitorClient := *openapiclient.NewMonitorClient("Mac_example") // MonitorClient | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverClientAPI.RebootClientUplinkPoe(context.Background(), omadacId, siteId).MonitorClient(monitorClient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverClientAPI.RebootClientUplinkPoe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootClientUplinkPoe`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverClientAPI.RebootClientUplinkPoe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootClientUplinkPoeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **monitorClient** | [**MonitorClient**](MonitorClient.md) |  | 

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


## VerifyMonitorClient

> OperationResponseWithoutResult VerifyMonitorClient(ctx, omadacId, siteId).VerifyClientVO(verifyClientVO).Execute()

Verify if the client is allowed to be monitored



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
	verifyClientVO := *openapiclient.NewVerifyClientVO() // VerifyClientVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntelliRecoverClientAPI.VerifyMonitorClient(context.Background(), omadacId, siteId).VerifyClientVO(verifyClientVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntelliRecoverClientAPI.VerifyMonitorClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyMonitorClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `IntelliRecoverClientAPI.VerifyMonitorClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMonitorClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **verifyClientVO** | [**VerifyClientVO**](VerifyClientVO.md) |  | 

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


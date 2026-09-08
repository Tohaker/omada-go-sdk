# BackupScheduleAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackupScheduleTask**](BackupScheduleAPI.md#getbackupscheduletask) | **Get** /openapi/v1/{omadacId}/backupschedule/task | Get backup schedule task
[**ModifyBackupScheduleTask**](BackupScheduleAPI.md#modifybackupscheduletask) | **Post** /openapi/v1/{omadacId}/backupschedule/task | Modify backup schedule task



## GetBackupScheduleTask

> OperationResponseAutoBackupOpenApiVO GetBackupScheduleTask(ctx, omadacId).Execute()

Get backup schedule task



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
	resp, r, err := apiClient.BackupScheduleAPI.GetBackupScheduleTask(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupScheduleAPI.GetBackupScheduleTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupScheduleTask`: OperationResponseAutoBackupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BackupScheduleAPI.GetBackupScheduleTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupScheduleTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseAutoBackupOpenApiVO**](OperationResponseAutoBackupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyBackupScheduleTask

> OperationResponseWithoutResult ModifyBackupScheduleTask(ctx, omadacId).AutoBackupOpenApiVO(autoBackupOpenApiVO).Execute()

Modify backup schedule task



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
	autoBackupOpenApiVO := *openapiclient.NewAutoBackupOpenApiVO(false) // AutoBackupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupScheduleAPI.ModifyBackupScheduleTask(context.Background(), omadacId).AutoBackupOpenApiVO(autoBackupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupScheduleAPI.ModifyBackupScheduleTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBackupScheduleTask`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupScheduleAPI.ModifyBackupScheduleTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBackupScheduleTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoBackupOpenApiVO** | [**AutoBackupOpenApiVO**](AutoBackupOpenApiVO.md) |  | 

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


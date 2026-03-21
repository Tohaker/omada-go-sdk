# \BackupAndRestoreAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupFileServer**](BackupAndRestoreAPI.md#BackupFileServer) | **Post** /openapi/v1/{omadacId}/maintenance/backup/file-server | Backup controller config to file server
[**BackupSelfServer**](BackupAndRestoreAPI.md#BackupSelfServer) | **Post** /openapi/v1/{omadacId}/maintenance/backup/self-server | Backup controller config to cloud server
[**BackupSitesFileServer**](BackupAndRestoreAPI.md#BackupSitesFileServer) | **Post** /openapi/v1/{omadacId}/sites/maintenance/multi-backup/file-server | Backup multi sites config to file server
[**BackupSitesSelfServer**](BackupAndRestoreAPI.md#BackupSitesSelfServer) | **Post** /openapi/v1/{omadacId}/sites/maintenance/multi-backup/self-server | Backup multi sites config to self server
[**GetBackupResult**](BackupAndRestoreAPI.md#GetBackupResult) | **Get** /openapi/v1/{omadacId}/maintenance/backup/result | Get controller backup result
[**GetRestoreResult**](BackupAndRestoreAPI.md#GetRestoreResult) | **Get** /openapi/v1/{omadacId}/maintenance/restore/result | Get controller restore result
[**GetSelfServerFileList**](BackupAndRestoreAPI.md#GetSelfServerFileList) | **Get** /openapi/v1/{omadacId}/maintenance/backup/files | Get controller backup file list in self server
[**GetSelfServerSiteFileList**](BackupAndRestoreAPI.md#GetSelfServerSiteFileList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/maintenance/backup/files | Get site backup file list in self server
[**GetSiteBackupResult**](BackupAndRestoreAPI.md#GetSiteBackupResult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/backup/result | Get site backup result
[**RestoreFileServer**](BackupAndRestoreAPI.md#RestoreFileServer) | **Post** /openapi/v1/{omadacId}/maintenance/restore/file-server | Restore controller config from file server
[**RestoreSelfServer**](BackupAndRestoreAPI.md#RestoreSelfServer) | **Post** /openapi/v1/{omadacId}/maintenance/restore/self-server | Restore controller config from cloud server
[**RestoreSitesFileServer**](BackupAndRestoreAPI.md#RestoreSitesFileServer) | **Post** /openapi/v1/{omadacId}/sites/maintenance/multi-restore/file-server | Restore multi sites config from file server
[**RestoreSitesSelfServer**](BackupAndRestoreAPI.md#RestoreSitesSelfServer) | **Post** /openapi/v1/{omadacId}/sites/maintenance/multi-restore/self-server | Restore multi sites config from self server



## BackupFileServer

> OperationResponseWithoutResult BackupFileServer(ctx, omadacId).FileServerGlobalBackupVO(fileServerGlobalBackupVO).Execute()

Backup controller config to file server



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
	fileServerGlobalBackupVO := *openapiclient.NewFileServerGlobalBackupVO("FilePath_example", *openapiclient.NewFileServerOpenApiVO("Hostname_example", int32(123), "Protocol_example")) // FileServerGlobalBackupVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.BackupFileServer(context.Background(), omadacId).FileServerGlobalBackupVO(fileServerGlobalBackupVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.BackupFileServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupFileServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.BackupFileServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupFileServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileServerGlobalBackupVO** | [**FileServerGlobalBackupVO**](FileServerGlobalBackupVO.md) |  | 

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


## BackupSelfServer

> OperationResponseWithoutResult BackupSelfServer(ctx, omadacId).SelfGlobalBackupVO(selfGlobalBackupVO).Execute()

Backup controller config to cloud server



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
	selfGlobalBackupVO := *openapiclient.NewSelfGlobalBackupVO(false) // SelfGlobalBackupVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.BackupSelfServer(context.Background(), omadacId).SelfGlobalBackupVO(selfGlobalBackupVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.BackupSelfServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupSelfServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.BackupSelfServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupSelfServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selfGlobalBackupVO** | [**SelfGlobalBackupVO**](SelfGlobalBackupVO.md) |  | 

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


## BackupSitesFileServer

> OperationResponseWithoutResult BackupSitesFileServer(ctx, omadacId).FileServerSiteBackupVO(fileServerSiteBackupVO).Execute()

Backup multi sites config to file server



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
	fileServerSiteBackupVO := *openapiclient.NewFileServerSiteBackupVO("FilePath_example", *openapiclient.NewFileServerOpenApiVO("Hostname_example", int32(123), "Protocol_example"), []string{"SiteIds_example"}) // FileServerSiteBackupVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.BackupSitesFileServer(context.Background(), omadacId).FileServerSiteBackupVO(fileServerSiteBackupVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.BackupSitesFileServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupSitesFileServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.BackupSitesFileServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupSitesFileServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileServerSiteBackupVO** | [**FileServerSiteBackupVO**](FileServerSiteBackupVO.md) |  | 

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


## BackupSitesSelfServer

> OperationResponseWithoutResult BackupSitesSelfServer(ctx, omadacId).BatchSiteBackupVO(batchSiteBackupVO).Execute()

Backup multi sites config to self server



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
	batchSiteBackupVO := *openapiclient.NewBatchSiteBackupVO([]string{"SiteIds_example"}) // BatchSiteBackupVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.BackupSitesSelfServer(context.Background(), omadacId).BatchSiteBackupVO(batchSiteBackupVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.BackupSitesSelfServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupSitesSelfServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.BackupSitesSelfServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupSitesSelfServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchSiteBackupVO** | [**BatchSiteBackupVO**](BatchSiteBackupVO.md) |  | 

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


## GetBackupResult

> BackupResultOpenApiVO GetBackupResult(ctx, omadacId).Execute()

Get controller backup result



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
	resp, r, err := apiClient.BackupAndRestoreAPI.GetBackupResult(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.GetBackupResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupResult`: BackupResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.GetBackupResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupResultOpenApiVO**](BackupResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreResult

> RestoreResultVO GetRestoreResult(ctx, omadacId).Execute()

Get controller restore result



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
	resp, r, err := apiClient.BackupAndRestoreAPI.GetRestoreResult(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.GetRestoreResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestoreResult`: RestoreResultVO
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.GetRestoreResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestoreResultVO**](RestoreResultVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServerFileList

> BackupFileListVO GetSelfServerFileList(ctx, omadacId).Execute()

Get controller backup file list in self server



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
	resp, r, err := apiClient.BackupAndRestoreAPI.GetSelfServerFileList(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.GetSelfServerFileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelfServerFileList`: BackupFileListVO
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.GetSelfServerFileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServerFileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupFileListVO**](BackupFileListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServerSiteFileList

> BackupFileListVO GetSelfServerSiteFileList(ctx, omadacId, siteId).Execute()

Get site backup file list in self server



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
	resp, r, err := apiClient.BackupAndRestoreAPI.GetSelfServerSiteFileList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.GetSelfServerSiteFileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelfServerSiteFileList`: BackupFileListVO
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.GetSelfServerSiteFileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServerSiteFileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BackupFileListVO**](BackupFileListVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteBackupResult

> BackupResultOpenApiVO GetSiteBackupResult(ctx, omadacId, siteId).Execute()

Get site backup result



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
	resp, r, err := apiClient.BackupAndRestoreAPI.GetSiteBackupResult(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.GetSiteBackupResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteBackupResult`: BackupResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.GetSiteBackupResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteBackupResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BackupResultOpenApiVO**](BackupResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFileServer

> OperationResponseWithoutResult RestoreFileServer(ctx, omadacId).FileServerGlobalRestoreVO(fileServerGlobalRestoreVO).Execute()

Restore controller config from file server



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
	fileServerGlobalRestoreVO := *openapiclient.NewFileServerGlobalRestoreVO("FilePath_example", *openapiclient.NewFileServerOpenApiVO("Hostname_example", int32(123), "Protocol_example"), false) // FileServerGlobalRestoreVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.RestoreFileServer(context.Background(), omadacId).FileServerGlobalRestoreVO(fileServerGlobalRestoreVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.RestoreFileServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreFileServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.RestoreFileServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFileServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileServerGlobalRestoreVO** | [**FileServerGlobalRestoreVO**](FileServerGlobalRestoreVO.md) |  | 

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


## RestoreSelfServer

> OperationResponseWithoutResult RestoreSelfServer(ctx, omadacId).SelfGlobalRestoreVO(selfGlobalRestoreVO).Execute()

Restore controller config from cloud server



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
	selfGlobalRestoreVO := *openapiclient.NewSelfGlobalRestoreVO("FileName_example") // SelfGlobalRestoreVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.RestoreSelfServer(context.Background(), omadacId).SelfGlobalRestoreVO(selfGlobalRestoreVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.RestoreSelfServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreSelfServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.RestoreSelfServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSelfServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selfGlobalRestoreVO** | [**SelfGlobalRestoreVO**](SelfGlobalRestoreVO.md) |  | 

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


## RestoreSitesFileServer

> OperationResponseWithoutResult RestoreSitesFileServer(ctx, omadacId).BatchSiteFileServerRestoreVO(batchSiteFileServerRestoreVO).Execute()

Restore multi sites config from file server



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
	batchSiteFileServerRestoreVO := *openapiclient.NewBatchSiteFileServerRestoreVO(*openapiclient.NewFileServerOpenApiVO("Hostname_example", int32(123), "Protocol_example"), []openapiclient.FileServerSiteRestoreVO{*openapiclient.NewFileServerSiteRestoreVO("FilePath_example", "SiteId_example")}) // BatchSiteFileServerRestoreVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.RestoreSitesFileServer(context.Background(), omadacId).BatchSiteFileServerRestoreVO(batchSiteFileServerRestoreVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.RestoreSitesFileServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreSitesFileServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.RestoreSitesFileServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSitesFileServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchSiteFileServerRestoreVO** | [**BatchSiteFileServerRestoreVO**](BatchSiteFileServerRestoreVO.md) |  | 

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


## RestoreSitesSelfServer

> OperationResponseWithoutResult RestoreSitesSelfServer(ctx, omadacId).BatchSiteSelfRestoreVO(batchSiteSelfRestoreVO).Execute()

Restore multi sites config from self server



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
	batchSiteSelfRestoreVO := *openapiclient.NewBatchSiteSelfRestoreVO([]openapiclient.SelfSiteRestoreVO{*openapiclient.NewSelfSiteRestoreVO("FileName_example", "SiteId_example")}) // BatchSiteSelfRestoreVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.RestoreSitesSelfServer(context.Background(), omadacId).BatchSiteSelfRestoreVO(batchSiteSelfRestoreVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.RestoreSitesSelfServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreSitesSelfServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.RestoreSitesSelfServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSitesSelfServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchSiteSelfRestoreVO** | [**BatchSiteSelfRestoreVO**](BatchSiteSelfRestoreVO.md) |  | 

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


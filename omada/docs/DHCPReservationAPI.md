# DHCPReservationAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchExportDhcpReservationListToFile**](DHCPReservationAPI.md#batchexportdhcpreservationlisttofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/files/dhcp/batch-export | Batch Export DHCP reservation list to file
[**CreateDhcpReservation**](DHCPReservationAPI.md#createdhcpreservation) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp | Create a new DHCP reservation
[**DeleteDhcpReservation**](DHCPReservationAPI.md#deletedhcpreservation) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/{mac} | Delete an exist DHCP reservation
[**DownloadDhcpImportResult**](DHCPReservationAPI.md#downloaddhcpimportresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/files/dhcp/import-result | Download DHCP import result
[**ExportDhcpReservation**](DHCPReservationAPI.md#exportdhcpreservation) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/{mac}/export | Export DHCP reservation to IP_MAC Binding
[**ExportDhcpReservationListToFile**](DHCPReservationAPI.md#exportdhcpreservationlisttofile) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/files/dhcp/export | Export DHCP reservation list to file
[**GetDhcpReservationGrid**](DHCPReservationAPI.md#getdhcpreservationgrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp | Get DHCP reservation list
[**GetGridAllDhcpUserList**](DHCPReservationAPI.md#getgridalldhcpuserlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/user-list | Get the dhcp user list of all servers
[**ImportDhcpReservationListFromFile**](DHCPReservationAPI.md#importdhcpreservationlistfromfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/files/dhcp/import | Import DHCP reservation list from file
[**ModifyDhcpReservation**](DHCPReservationAPI.md#modifydhcpreservation) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/{mac} | Modify an exist DHCP reservation
[**ReservationFromDhcpUserList**](DHCPReservationAPI.md#reservationfromdhcpuserlist) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/user-list/reservation | Reservation From Dhcp User List



## BatchExportDhcpReservationListToFile

> OperationResponseWithoutResult BatchExportDhcpReservationListToFile(ctx, omadacId, siteId).DhcpReservationFilterVO(dhcpReservationFilterVO).Execute()

Batch Export DHCP reservation list to file



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
	dhcpReservationFilterVO := *openapiclient.NewDhcpReservationFilterVO(*openapiclient.NewBatchSelectMacsVO("SelectType_example")) // DhcpReservationFilterVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.BatchExportDhcpReservationListToFile(context.Background(), omadacId, siteId).DhcpReservationFilterVO(dhcpReservationFilterVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.BatchExportDhcpReservationListToFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportDhcpReservationListToFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.BatchExportDhcpReservationListToFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportDhcpReservationListToFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dhcpReservationFilterVO** | [**DhcpReservationFilterVO**](DhcpReservationFilterVO.md) |  | 

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


## CreateDhcpReservation

> OperationResponseResIdOpenApiVO CreateDhcpReservation(ctx, omadacId, siteId).CreateDhcpReservationOpenApiVO(createDhcpReservationOpenApiVO).Execute()

Create a new DHCP reservation



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
	createDhcpReservationOpenApiVO := *openapiclient.NewCreateDhcpReservationOpenApiVO("Mac_example", "NetId_example", false) // CreateDhcpReservationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.CreateDhcpReservation(context.Background(), omadacId, siteId).CreateDhcpReservationOpenApiVO(createDhcpReservationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.CreateDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDhcpReservation`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.CreateDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDhcpReservationOpenApiVO** | [**CreateDhcpReservationOpenApiVO**](CreateDhcpReservationOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDhcpReservation

> OperationResponseWithoutResult DeleteDhcpReservation(ctx, omadacId, siteId, mac).Execute()

Delete an exist DHCP reservation



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
	mac := "mac_example" // string | MAC address of the DHCP reservation, format: AA-BB-CC-11-22-33.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.DeleteDhcpReservation(context.Background(), omadacId, siteId, mac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.DeleteDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDhcpReservation`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.DeleteDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | MAC address of the DHCP reservation, format: AA-BB-CC-11-22-33. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDhcpReservationRequest struct via the builder pattern


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


## DownloadDhcpImportResult

> OperationResponse DownloadDhcpImportResult(ctx, omadacId, siteId).Execute()

Download DHCP import result



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
	resp, r, err := apiClient.DHCPReservationAPI.DownloadDhcpImportResult(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.DownloadDhcpImportResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDhcpImportResult`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.DownloadDhcpImportResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDhcpImportResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDhcpReservation

> OperationResponseWithoutResult ExportDhcpReservation(ctx, omadacId, siteId, mac).ExportDhcpReservationOpenApiVO(exportDhcpReservationOpenApiVO).Execute()

Export DHCP reservation to IP_MAC Binding



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
	mac := "mac_example" // string | DHCP MAC address
	exportDhcpReservationOpenApiVO := *openapiclient.NewExportDhcpReservationOpenApiVO("InterfaceId_example") // ExportDhcpReservationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.ExportDhcpReservation(context.Background(), omadacId, siteId, mac).ExportDhcpReservationOpenApiVO(exportDhcpReservationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.ExportDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDhcpReservation`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.ExportDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | DHCP MAC address | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **exportDhcpReservationOpenApiVO** | [**ExportDhcpReservationOpenApiVO**](ExportDhcpReservationOpenApiVO.md) |  | 

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


## ExportDhcpReservationListToFile

> OperationResponseWithoutResult ExportDhcpReservationListToFile(ctx, omadacId, siteId).Execute()

Export DHCP reservation list to file



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
	resp, r, err := apiClient.DHCPReservationAPI.ExportDhcpReservationListToFile(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.ExportDhcpReservationListToFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDhcpReservationListToFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.ExportDhcpReservationListToFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDhcpReservationListToFileRequest struct via the builder pattern


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


## GetDhcpReservationGrid

> OperationResponseDhcpReservationOpenApiGridVODhcpReservationOpenApiVO GetDhcpReservationGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SortsIp(sortsIp).SortsNetName(sortsNetName).SortsDescription(sortsDescription).SortsStatus(sortsStatus).SearchKey(searchKey).Execute()

Get DHCP reservation list



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
	sortsMac := "sortsMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsNetName := "sortsNetName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsDescription := "sortsDescription_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field : mac, ip, description. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.GetDhcpReservationGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SortsIp(sortsIp).SortsNetName(sortsNetName).SortsDescription(sortsDescription).SortsStatus(sortsStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.GetDhcpReservationGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpReservationGrid`: OperationResponseDhcpReservationOpenApiGridVODhcpReservationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.GetDhcpReservationGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpReservationGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsNetName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsDescription** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStatus** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field : mac, ip, description. | 

### Return type

[**OperationResponseDhcpReservationOpenApiGridVODhcpReservationOpenApiVO**](OperationResponseDhcpReservationOpenApiGridVODhcpReservationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridAllDhcpUserList

> OperationResponseDhcpUserGridVODhcpUserVO GetGridAllDhcpUserList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the dhcp user list of all servers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.GetGridAllDhcpUserList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.GetGridAllDhcpUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAllDhcpUserList`: OperationResponseDhcpUserGridVODhcpUserVO
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.GetGridAllDhcpUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAllDhcpUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseDhcpUserGridVODhcpUserVO**](OperationResponseDhcpUserGridVODhcpUserVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportDhcpReservationListFromFile

> OperationResponseWithoutResult ImportDhcpReservationListFromFile(ctx, omadacId, siteId).ImportDhcpReservationListFromFileRequest(importDhcpReservationListFromFileRequest).Execute()

Import DHCP reservation list from file



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
	importDhcpReservationListFromFileRequest := *openapiclient.NewImportDhcpReservationListFromFileRequest() // ImportDhcpReservationListFromFileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.ImportDhcpReservationListFromFile(context.Background(), omadacId, siteId).ImportDhcpReservationListFromFileRequest(importDhcpReservationListFromFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.ImportDhcpReservationListFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportDhcpReservationListFromFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.ImportDhcpReservationListFromFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportDhcpReservationListFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **importDhcpReservationListFromFileRequest** | [**ImportDhcpReservationListFromFileRequest**](ImportDhcpReservationListFromFileRequest.md) |  | 

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


## ModifyDhcpReservation

> OperationResponseWithoutResult ModifyDhcpReservation(ctx, omadacId, siteId, mac).CreateDhcpReservationOpenApiVO(createDhcpReservationOpenApiVO).Execute()

Modify an exist DHCP reservation



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
	mac := "mac_example" // string | MAC address of the DHCP reservation, format: AA-BB-CC-11-22-33.
	createDhcpReservationOpenApiVO := *openapiclient.NewCreateDhcpReservationOpenApiVO("Mac_example", "NetId_example", false) // CreateDhcpReservationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.ModifyDhcpReservation(context.Background(), omadacId, siteId, mac).CreateDhcpReservationOpenApiVO(createDhcpReservationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.ModifyDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDhcpReservation`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.ModifyDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | MAC address of the DHCP reservation, format: AA-BB-CC-11-22-33. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createDhcpReservationOpenApiVO** | [**CreateDhcpReservationOpenApiVO**](CreateDhcpReservationOpenApiVO.md) |  | 

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


## ReservationFromDhcpUserList

> OperationResponseDhcpReservationErrorVO ReservationFromDhcpUserList(ctx, omadacId, siteId).DhcpUserFilterVO(dhcpUserFilterVO).Execute()

Reservation From Dhcp User List



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
	dhcpUserFilterVO := *openapiclient.NewDhcpUserFilterVO([]string{"SelectIps_example"}, *openapiclient.NewBatchSelectMacsVO("SelectType_example")) // DhcpUserFilterVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPReservationAPI.ReservationFromDhcpUserList(context.Background(), omadacId, siteId).DhcpUserFilterVO(dhcpUserFilterVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPReservationAPI.ReservationFromDhcpUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReservationFromDhcpUserList`: OperationResponseDhcpReservationErrorVO
	fmt.Fprintf(os.Stdout, "Response from `DHCPReservationAPI.ReservationFromDhcpUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationFromDhcpUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dhcpUserFilterVO** | [**DhcpUserFilterVO**](DhcpUserFilterVO.md) |  | 

### Return type

[**OperationResponseDhcpReservationErrorVO**](OperationResponseDhcpReservationErrorVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# DataExportAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportSiteRogueApZipFile**](DataExportAPI.md#exportsiterogueapzipfile) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rogue-ap/export/{format} | Export site Rogue AP scan results



## ExportSiteRogueApZipFile

> OperationResponseWithoutResult ExportSiteRogueApZipFile(ctx, omadacId, siteId, format).Page(page).PageSize(pageSize).Execute()

Export site Rogue AP scan results



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
	format := "format_example" // string | The type of the exported file. 0: CSV, 1: xlsx.
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataExportAPI.ExportSiteRogueApZipFile(context.Background(), omadacId, siteId, format).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataExportAPI.ExportSiteRogueApZipFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSiteRogueApZipFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `DataExportAPI.ExportSiteRogueApZipFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**format** | **string** | The type of the exported file. 0: CSV, 1: xlsx. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSiteRogueApZipFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

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


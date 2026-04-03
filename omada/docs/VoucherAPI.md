# VoucherAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearInvalidVouchersInSelectedGroups**](VoucherAPI.md#clearinvalidvouchersinselectedgroups) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/batch/clear-invalid | Delete expired vouchers in voucher groups
[**ClearInvalidVouchersInaGroup**](VoucherAPI.md#clearinvalidvouchersinagroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/{groupId}/clear-invalid | Delete expired vouchers in a voucher group
[**CreateVoucherGroup**](VoucherAPI.md#createvouchergroup) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups | Create a new Voucher Group
[**DeleteSelectedVoucherGroups**](VoucherAPI.md#deleteselectedvouchergroups) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/batch/delete | Delete selected Voucher Groups
[**DeleteVoucher**](VoucherAPI.md#deletevoucher) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/vouchers/{id} | Delete an existing voucher
[**DeleteVoucherGroup**](VoucherAPI.md#deletevouchergroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/{groupId} | Delete an existing Voucher Group
[**DeleteVouchers**](VoucherAPI.md#deletevouchers) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/vouchers/batch/delete | Delete selected vouchers
[**EditVoucherGroupPattern**](VoucherAPI.md#editvouchergrouppattern) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/{groupId}/pattern | Edit voucher group pattern
[**ExportVouchersWithSiteByCloudAccess**](VoucherAPI.md#exportvoucherswithsitebycloudaccess) | **Post** /openapi/v1/{omadacId}/files/hotspot/sites/{siteId}/vouchers/export | Export voucher list to file
[**GetAllTimeVoucherSummary**](VoucherAPI.md#getalltimevouchersummary) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/vouchers/statistics/summary | Get voucher summary
[**GetGridVoucherGroups**](VoucherAPI.md#getgridvouchergroups) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups | Get Voucher Group List
[**GetGroupConfigLimit**](VoucherAPI.md#getgroupconfiglimit) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/config-limit | Get voucher group config limit
[**GetVoucher**](VoucherAPI.md#getvoucher) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/vouchers/{id} | Get a voucher
[**GetVoucherCurrencyCandidates**](VoucherAPI.md#getvouchercurrencycandidates) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/currency-list | Get voucher currency candidates
[**GetVoucherDistributionByDuration**](VoucherAPI.md#getvoucherdistributionbyduration) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/vouchers/statistics/history/distribution/duration | Get voucher distribution by duration
[**GetVoucherDistributionByPrice**](VoucherAPI.md#getvoucherdistributionbyprice) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/vouchers/statistics/history/distribution/unit-price | Get voucher distribution by price
[**GetVoucherGroupDetail**](VoucherAPI.md#getvouchergroupdetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/{groupId} | Get Voucher Group Detail
[**GetVoucherHistoryStatistics**](VoucherAPI.md#getvoucherhistorystatistics) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/vouchers/statistics/history | Get voucher history statistics
[**GetVoucherLogo**](VoucherAPI.md#getvoucherlogo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/files/voucher/logos/{picId} | Get voucher logo
[**ModifyHotspotSetting**](VoucherAPI.md#modifyhotspotsetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/setting | Modify voucher currency
[**PrintSelectedVoucherGroups**](VoucherAPI.md#printselectedvouchergroups) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/batch/print-unused | Get unused vouchers in selected voucher groups
[**PrintVoucherGroup**](VoucherAPI.md#printvouchergroup) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/voucher-groups/{groupId}/print-unused | Get unused vouchers in voucher group
[**PrintVouchers**](VoucherAPI.md#printvouchers) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/vouchers/batch/print | Get selected vouchers in a voucher group
[**UploadVoucherLogo**](VoucherAPI.md#uploadvoucherlogo) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/files/voucher/logos | Upload voucher logo



## ClearInvalidVouchersInSelectedGroups

> OperationResponseWithoutResult ClearInvalidVouchersInSelectedGroups(ctx, omadacId, siteId).SelectVoucherGroupOpenApiVO(selectVoucherGroupOpenApiVO).Execute()

Delete expired vouchers in voucher groups



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
	selectVoucherGroupOpenApiVO := *openapiclient.NewSelectVoucherGroupOpenApiVO(int32(123)) // SelectVoucherGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.ClearInvalidVouchersInSelectedGroups(context.Background(), omadacId, siteId).SelectVoucherGroupOpenApiVO(selectVoucherGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.ClearInvalidVouchersInSelectedGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearInvalidVouchersInSelectedGroups`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.ClearInvalidVouchersInSelectedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearInvalidVouchersInSelectedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectVoucherGroupOpenApiVO** | [**SelectVoucherGroupOpenApiVO**](SelectVoucherGroupOpenApiVO.md) |  | 

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


## ClearInvalidVouchersInaGroup

> OperationResponseWithoutResult ClearInvalidVouchersInaGroup(ctx, omadacId, siteId, groupId).Execute()

Delete expired vouchers in a voucher group



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
	groupId := "groupId_example" // string | Voucher Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.ClearInvalidVouchersInaGroup(context.Background(), omadacId, siteId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.ClearInvalidVouchersInaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearInvalidVouchersInaGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.ClearInvalidVouchersInaGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**groupId** | **string** | Voucher Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearInvalidVouchersInaGroupRequest struct via the builder pattern


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


## CreateVoucherGroup

> OperationResponseCreatedResIdOpenApiVO CreateVoucherGroup(ctx, omadacId, siteId).CreateVoucherGroupOpenApiVO(createVoucherGroupOpenApiVO).Execute()

Create a new Voucher Group



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
	createVoucherGroupOpenApiVO := *openapiclient.NewCreateVoucherGroupOpenApiVO(int32(123), false, []int32{int32(123)}, int32(123), int64(123), int32(123), int32(123), "Name_example", *openapiclient.NewRateLimitOpenApiVO(int32(123)), int32(123), false) // CreateVoucherGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.CreateVoucherGroup(context.Background(), omadacId, siteId).CreateVoucherGroupOpenApiVO(createVoucherGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.CreateVoucherGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoucherGroup`: OperationResponseCreatedResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.CreateVoucherGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoucherGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVoucherGroupOpenApiVO** | [**CreateVoucherGroupOpenApiVO**](CreateVoucherGroupOpenApiVO.md) |  | 

### Return type

[**OperationResponseCreatedResIdOpenApiVO**](OperationResponseCreatedResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSelectedVoucherGroups

> OperationResponseWithoutResult DeleteSelectedVoucherGroups(ctx, omadacId, siteId).SelectVoucherGroupOpenApiVO(selectVoucherGroupOpenApiVO).Execute()

Delete selected Voucher Groups



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
	selectVoucherGroupOpenApiVO := *openapiclient.NewSelectVoucherGroupOpenApiVO(int32(123)) // SelectVoucherGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.DeleteSelectedVoucherGroups(context.Background(), omadacId, siteId).SelectVoucherGroupOpenApiVO(selectVoucherGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.DeleteSelectedVoucherGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSelectedVoucherGroups`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.DeleteSelectedVoucherGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelectedVoucherGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectVoucherGroupOpenApiVO** | [**SelectVoucherGroupOpenApiVO**](SelectVoucherGroupOpenApiVO.md) |  | 

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


## DeleteVoucher

> OperationResponseWithoutResult DeleteVoucher(ctx, omadacId, siteId, id).Execute()

Delete an existing voucher



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
	id := "id_example" // string | Voucher ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.DeleteVoucher(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.DeleteVoucher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVoucher`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.DeleteVoucher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Voucher ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoucherRequest struct via the builder pattern


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


## DeleteVoucherGroup

> OperationResponseWithoutResult DeleteVoucherGroup(ctx, omadacId, siteId, groupId).Execute()

Delete an existing Voucher Group



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
	groupId := "groupId_example" // string | Voucher Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.DeleteVoucherGroup(context.Background(), omadacId, siteId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.DeleteVoucherGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVoucherGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.DeleteVoucherGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**groupId** | **string** | Voucher Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoucherGroupRequest struct via the builder pattern


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


## DeleteVouchers

> OperationResponseWithoutResult DeleteVouchers(ctx, omadacId, siteId).SelectIdsOpenApiVO(selectIdsOpenApiVO).Execute()

Delete selected vouchers



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
	selectIdsOpenApiVO := *openapiclient.NewSelectIdsOpenApiVO("GroupId_example", int32(123)) // SelectIdsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.DeleteVouchers(context.Background(), omadacId, siteId).SelectIdsOpenApiVO(selectIdsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.DeleteVouchers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVouchers`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.DeleteVouchers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVouchersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectIdsOpenApiVO** | [**SelectIdsOpenApiVO**](SelectIdsOpenApiVO.md) |  | 

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


## EditVoucherGroupPattern

> OperationResponse EditVoucherGroupPattern(ctx, omadacId, siteId, groupId).EditVoucherGroupPatternOpenApiVO(editVoucherGroupPatternOpenApiVO).Execute()

Edit voucher group pattern



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
	groupId := "groupId_example" // string | groupId
	editVoucherGroupPatternOpenApiVO := *openapiclient.NewEditVoucherGroupPatternOpenApiVO() // EditVoucherGroupPatternOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.EditVoucherGroupPattern(context.Background(), omadacId, siteId, groupId).EditVoucherGroupPatternOpenApiVO(editVoucherGroupPatternOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.EditVoucherGroupPattern``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditVoucherGroupPattern`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.EditVoucherGroupPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**groupId** | **string** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditVoucherGroupPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **editVoucherGroupPatternOpenApiVO** | [**EditVoucherGroupPatternOpenApiVO**](EditVoucherGroupPatternOpenApiVO.md) |  | 

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


## ExportVouchersWithSiteByCloudAccess

> OperationResponseWithoutResult ExportVouchersWithSiteByCloudAccess(ctx, omadacId, siteId).ExportVoucherOpenApiVO(exportVoucherOpenApiVO).Execute()

Export voucher list to file



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
	exportVoucherOpenApiVO := *openapiclient.NewExportVoucherOpenApiVO(int32(123), int32(123)) // ExportVoucherOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.ExportVouchersWithSiteByCloudAccess(context.Background(), omadacId, siteId).ExportVoucherOpenApiVO(exportVoucherOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.ExportVouchersWithSiteByCloudAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportVouchersWithSiteByCloudAccess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.ExportVouchersWithSiteByCloudAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportVouchersWithSiteByCloudAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exportVoucherOpenApiVO** | [**ExportVoucherOpenApiVO**](ExportVoucherOpenApiVO.md) |  | 

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


## GetAllTimeVoucherSummary

> OperationResponseAllTimeVoucherSummaryOpenApiVO GetAllTimeVoucherSummary(ctx, omadacId, siteId).Execute()

Get voucher summary



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
	resp, r, err := apiClient.VoucherAPI.GetAllTimeVoucherSummary(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetAllTimeVoucherSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTimeVoucherSummary`: OperationResponseAllTimeVoucherSummaryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetAllTimeVoucherSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTimeVoucherSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseAllTimeVoucherSummaryOpenApiVO**](OperationResponseAllTimeVoucherSummaryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVoucherGroups

> OperationResponseGridVOVoucherGroupOpenApiVO GetGridVoucherGroups(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsCreateTime(sortsCreateTime).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).SearchKey(searchKey).Execute()

Get Voucher Group List



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
	sortsCreateTime := "sortsCreateTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field time range: start timestamp (ms). (optional)
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field time range: end timestamp (ms). (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.GetGridVoucherGroups(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsCreateTime(sortsCreateTime).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetGridVoucherGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVoucherGroups`: OperationResponseGridVOVoucherGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetGridVoucherGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVoucherGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsCreateTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersTimeStart** | **int64** | Filter query parameters, support field time range: start timestamp (ms). | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field time range: end timestamp (ms). | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseGridVOVoucherGroupOpenApiVO**](OperationResponseGridVOVoucherGroupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupConfigLimit

> OperationResponseListVoucherConfigLimitOpenApiVO GetGroupConfigLimit(ctx, omadacId, siteId).Execute()

Get voucher group config limit



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
	resp, r, err := apiClient.VoucherAPI.GetGroupConfigLimit(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetGroupConfigLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupConfigLimit`: OperationResponseListVoucherConfigLimitOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetGroupConfigLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupConfigLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListVoucherConfigLimitOpenApiVO**](OperationResponseListVoucherConfigLimitOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucher

> OperationResponseVoucherOpenApiVO GetVoucher(ctx, omadacId, siteId, id).Execute()

Get a voucher



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
	id := "id_example" // string | Voucher ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.GetVoucher(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetVoucher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoucher`: OperationResponseVoucherOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetVoucher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Voucher ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseVoucherOpenApiVO**](OperationResponseVoucherOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucherCurrencyCandidates

> OperationResponseCurrencyCandidatesOpenApiVO GetVoucherCurrencyCandidates(ctx, omadacId, siteId).Execute()

Get voucher currency candidates



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
	resp, r, err := apiClient.VoucherAPI.GetVoucherCurrencyCandidates(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetVoucherCurrencyCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoucherCurrencyCandidates`: OperationResponseCurrencyCandidatesOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetVoucherCurrencyCandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherCurrencyCandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseCurrencyCandidatesOpenApiVO**](OperationResponseCurrencyCandidatesOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucherDistributionByDuration

> OperationResponseGridVOVoucherDurationDistributionOpenApiVO GetVoucherDistributionByDuration(ctx, omadacId, siteId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()

Get voucher distribution by duration



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
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field time range: start timestamp (second).
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field time range: end timestamp (second).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.GetVoucherDistributionByDuration(context.Background(), omadacId, siteId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetVoucherDistributionByDuration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoucherDistributionByDuration`: OperationResponseGridVOVoucherDurationDistributionOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetVoucherDistributionByDuration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherDistributionByDurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersTimeStart** | **int64** | Filter query parameters, support field time range: start timestamp (second). | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field time range: end timestamp (second). | 

### Return type

[**OperationResponseGridVOVoucherDurationDistributionOpenApiVO**](OperationResponseGridVOVoucherDurationDistributionOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucherDistributionByPrice

> OperationResponseGridVOVoucherUnitPriceDistributionOpenApiVO GetVoucherDistributionByPrice(ctx, omadacId, siteId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()

Get voucher distribution by price



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
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field time range: start timestamp (second).
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field time range: end timestamp (second).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.GetVoucherDistributionByPrice(context.Background(), omadacId, siteId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetVoucherDistributionByPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoucherDistributionByPrice`: OperationResponseGridVOVoucherUnitPriceDistributionOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetVoucherDistributionByPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherDistributionByPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersTimeStart** | **int64** | Filter query parameters, support field time range: start timestamp (second). | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field time range: end timestamp (second). | 

### Return type

[**OperationResponseGridVOVoucherUnitPriceDistributionOpenApiVO**](OperationResponseGridVOVoucherUnitPriceDistributionOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucherGroupDetail

> OperationResponseVoucherGroupGridOpenApiVO GetVoucherGroupDetail(ctx, omadacId, siteId, groupId).Page(page).PageSize(pageSize).SortsCode(sortsCode).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()

Get Voucher Group Detail



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
	groupId := "groupId_example" // string | Voucher Group ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsCode := "sortsCode_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersStatus := int64(789) // int64 | Filter query parameters, support field status for vouchers in the voucher group: 0: unused vouchers, 1: in-use vouchers, 2: expired vouchers (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.GetVoucherGroupDetail(context.Background(), omadacId, siteId, groupId).Page(page).PageSize(pageSize).SortsCode(sortsCode).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetVoucherGroupDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoucherGroupDetail`: OperationResponseVoucherGroupGridOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetVoucherGroupDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**groupId** | **string** | Voucher Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherGroupDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsCode** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersStatus** | **int64** | Filter query parameters, support field status for vouchers in the voucher group: 0: unused vouchers, 1: in-use vouchers, 2: expired vouchers | 
 **searchKey** | **string** | Fuzzy query parameters, support field code | 

### Return type

[**OperationResponseVoucherGroupGridOpenApiVO**](OperationResponseVoucherGroupGridOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucherHistoryStatistics

> OperationResponseVoucherStatisticsHistoryOpenApiVO GetVoucherHistoryStatistics(ctx, omadacId, siteId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()

Get voucher history statistics



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
	filtersTimeStart := int64(789) // int64 | Filter query parameters, support field time range: start timestamp (second).
	filtersTimeEnd := int64(789) // int64 | Filter query parameters, support field time range: end timestamp (second).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.GetVoucherHistoryStatistics(context.Background(), omadacId, siteId).FiltersTimeStart(filtersTimeStart).FiltersTimeEnd(filtersTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetVoucherHistoryStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoucherHistoryStatistics`: OperationResponseVoucherStatisticsHistoryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetVoucherHistoryStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherHistoryStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtersTimeStart** | **int64** | Filter query parameters, support field time range: start timestamp (second). | 
 **filtersTimeEnd** | **int64** | Filter query parameters, support field time range: end timestamp (second). | 

### Return type

[**OperationResponseVoucherStatisticsHistoryOpenApiVO**](OperationResponseVoucherStatisticsHistoryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucherLogo

> OperationResponseObject GetVoucherLogo(ctx, omadacId, siteId, picId).Execute()

Get voucher logo



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
	picId := "picId_example" // string | Voucher logo picture id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.GetVoucherLogo(context.Background(), omadacId, siteId, picId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.GetVoucherLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoucherLogo`: OperationResponseObject
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.GetVoucherLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**picId** | **string** | Voucher logo picture id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseObject**](OperationResponseObject.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyHotspotSetting

> OperationResponseWithoutResult ModifyHotspotSetting(ctx, omadacId, siteId).HotspotSiteSettingOpenApiVO(hotspotSiteSettingOpenApiVO).Execute()

Modify voucher currency



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
	hotspotSiteSettingOpenApiVO := *openapiclient.NewHotspotSiteSettingOpenApiVO() // HotspotSiteSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.ModifyHotspotSetting(context.Background(), omadacId, siteId).HotspotSiteSettingOpenApiVO(hotspotSiteSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.ModifyHotspotSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyHotspotSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.ModifyHotspotSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyHotspotSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hotspotSiteSettingOpenApiVO** | [**HotspotSiteSettingOpenApiVO**](HotspotSiteSettingOpenApiVO.md) |  | 

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


## PrintSelectedVoucherGroups

> OperationResponseListVoucherBriefOpenApiVO PrintSelectedVoucherGroups(ctx, omadacId, siteId).SelectVoucherGroupOpenApiVO(selectVoucherGroupOpenApiVO).Execute()

Get unused vouchers in selected voucher groups



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
	selectVoucherGroupOpenApiVO := *openapiclient.NewSelectVoucherGroupOpenApiVO(int32(123)) // SelectVoucherGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.PrintSelectedVoucherGroups(context.Background(), omadacId, siteId).SelectVoucherGroupOpenApiVO(selectVoucherGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.PrintSelectedVoucherGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrintSelectedVoucherGroups`: OperationResponseListVoucherBriefOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.PrintSelectedVoucherGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSelectedVoucherGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectVoucherGroupOpenApiVO** | [**SelectVoucherGroupOpenApiVO**](SelectVoucherGroupOpenApiVO.md) |  | 

### Return type

[**OperationResponseListVoucherBriefOpenApiVO**](OperationResponseListVoucherBriefOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintVoucherGroup

> OperationResponseListVoucherBriefOpenApiVO PrintVoucherGroup(ctx, omadacId, siteId, groupId).Execute()

Get unused vouchers in voucher group



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
	groupId := "groupId_example" // string | Voucher Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.PrintVoucherGroup(context.Background(), omadacId, siteId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.PrintVoucherGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrintVoucherGroup`: OperationResponseListVoucherBriefOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.PrintVoucherGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**groupId** | **string** | Voucher Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintVoucherGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListVoucherBriefOpenApiVO**](OperationResponseListVoucherBriefOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintVouchers

> OperationResponseListVoucherBriefOpenApiVO PrintVouchers(ctx, omadacId, siteId).SelectIdsOpenApiVO(selectIdsOpenApiVO).Execute()

Get selected vouchers in a voucher group



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
	selectIdsOpenApiVO := *openapiclient.NewSelectIdsOpenApiVO("GroupId_example", int32(123)) // SelectIdsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.PrintVouchers(context.Background(), omadacId, siteId).SelectIdsOpenApiVO(selectIdsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.PrintVouchers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrintVouchers`: OperationResponseListVoucherBriefOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.PrintVouchers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintVouchersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **selectIdsOpenApiVO** | [**SelectIdsOpenApiVO**](SelectIdsOpenApiVO.md) |  | 

### Return type

[**OperationResponseListVoucherBriefOpenApiVO**](OperationResponseListVoucherBriefOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVoucherLogo

> OperationResponsePortalPictureInfo UploadVoucherLogo(ctx, omadacId, siteId).Data(data).File(file).Execute()

Upload voucher logo



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
	data := *openapiclient.NewUploadVoucherGroupLogoOpenApiVO() // UploadVoucherGroupLogoOpenApiVO | 
	file := os.NewFile(1234, "some_file") // *os.File | At least one of the file or md5 parameters needs to be passed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.UploadVoucherLogo(context.Background(), omadacId, siteId).Data(data).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.UploadVoucherLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadVoucherLogo`: OperationResponsePortalPictureInfo
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.UploadVoucherLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadVoucherLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**UploadVoucherGroupLogoOpenApiVO**](UploadVoucherGroupLogoOpenApiVO.md) |  | 
 **file** | ***os.File** | At least one of the file or md5 parameters needs to be passed | 

### Return type

[**OperationResponsePortalPictureInfo**](OperationResponsePortalPictureInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


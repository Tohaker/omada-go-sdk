# AuthorizedClientAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthClient**](AuthorizedClientAPI.md#authclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/clients/{clientMac}/auth | Authentication the given client
[**CancelAuthClient**](AuthorizedClientAPI.md#cancelauthclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/clients/{clientMac}/unauth | Cancel authentication the given client
[**DeleteAllInvalidAuthedClients**](AuthorizedClientAPI.md#deleteallinvalidauthedclients) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/authed-records/delete | Delete all invalid authed record
[**DeleteHotspotAuthedClient**](AuthorizedClientAPI.md#deletehotspotauthedclient) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/authed-records/{id} | Delete given authed record
[**DisconnectHotspotAuthedClient**](AuthorizedClientAPI.md#disconnecthotspotauthedclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/authed-records/{id}/disconnect | Disconnect authed client 
[**ExportAuthedClientListGlobalByCloudAccess**](AuthorizedClientAPI.md#exportauthedclientlistglobalbycloudaccess) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/authed-client-list | Export global authed client list
[**ExtendHotspotAuthedClient**](AuthorizedClientAPI.md#extendhotspotauthedclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/authed-records/{id}/period | Modify period for given authed record
[**GetHotspotAuthedClients**](AuthorizedClientAPI.md#gethotspotauthedclients) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/authed-records | Get authentication record list
[**GetHotspotStatistic**](AuthorizedClientAPI.md#gethotspotstatistic) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/dashboard/statistics | Get hotspot statistic



## AuthClient

> OperationResponseWithoutResult AuthClient(ctx, omadacId, siteId, clientMac).Execute()

Authentication the given client



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
	clientMac := "clientMac_example" // string | Client MAC, format: AA-BB-CC-DD-EE-FF.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizedClientAPI.AuthClient(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.AuthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.AuthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client MAC, format: AA-BB-CC-DD-EE-FF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientRequest struct via the builder pattern


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


## CancelAuthClient

> OperationResponseWithoutResult CancelAuthClient(ctx, omadacId, siteId, clientMac).Execute()

Cancel authentication the given client



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
	resp, r, err := apiClient.AuthorizedClientAPI.CancelAuthClient(context.Background(), omadacId, siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.CancelAuthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelAuthClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.CancelAuthClient`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCancelAuthClientRequest struct via the builder pattern


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


## DeleteAllInvalidAuthedClients

> OperationResponseWithoutResult DeleteAllInvalidAuthedClients(ctx, omadacId, siteId).DeleteFilterInfo(deleteFilterInfo).Execute()

Delete all invalid authed record



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
	deleteFilterInfo := *openapiclient.NewDeleteFilterInfo(int64(123), int64(123)) // DeleteFilterInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizedClientAPI.DeleteAllInvalidAuthedClients(context.Background(), omadacId, siteId).DeleteFilterInfo(deleteFilterInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.DeleteAllInvalidAuthedClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllInvalidAuthedClients`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.DeleteAllInvalidAuthedClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllInvalidAuthedClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteFilterInfo** | [**DeleteFilterInfo**](DeleteFilterInfo.md) |  | 

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


## DeleteHotspotAuthedClient

> OperationResponseWithoutResult DeleteHotspotAuthedClient(ctx, omadacId, siteId, id).Execute()

Delete given authed record



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
	id := "id_example" // string | Authed record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizedClientAPI.DeleteHotspotAuthedClient(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.DeleteHotspotAuthedClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteHotspotAuthedClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.DeleteHotspotAuthedClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Authed record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHotspotAuthedClientRequest struct via the builder pattern


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


## DisconnectHotspotAuthedClient

> OperationResponseWithoutResult DisconnectHotspotAuthedClient(ctx, omadacId, siteId, id).Execute()

Disconnect authed client 



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
	id := "id_example" // string | id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizedClientAPI.DisconnectHotspotAuthedClient(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.DisconnectHotspotAuthedClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectHotspotAuthedClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.DisconnectHotspotAuthedClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectHotspotAuthedClientRequest struct via the builder pattern


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


## ExportAuthedClientListGlobalByCloudAccess

> OperationResponse ExportAuthedClientListGlobalByCloudAccess(ctx, omadacId, siteId).ExportAuthedClientOpenApiVO(exportAuthedClientOpenApiVO).Execute()

Export global authed client list



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
	siteId := "siteId_example" // string | siteId
	exportAuthedClientOpenApiVO := *openapiclient.NewExportAuthedClientOpenApiVO() // ExportAuthedClientOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizedClientAPI.ExportAuthedClientListGlobalByCloudAccess(context.Background(), omadacId, siteId).ExportAuthedClientOpenApiVO(exportAuthedClientOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.ExportAuthedClientListGlobalByCloudAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAuthedClientListGlobalByCloudAccess`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.ExportAuthedClientListGlobalByCloudAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | siteId | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAuthedClientListGlobalByCloudAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exportAuthedClientOpenApiVO** | [**ExportAuthedClientOpenApiVO**](ExportAuthedClientOpenApiVO.md) |  | 

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


## ExtendHotspotAuthedClient

> OperationResponseWithoutResult ExtendHotspotAuthedClient(ctx, omadacId, siteId, id).ExtendOpenApiVO(extendOpenApiVO).Execute()

Modify period for given authed record



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
	id := "id_example" // string | Authed record ID
	extendOpenApiVO := *openapiclient.NewExtendOpenApiVO(int64(123)) // ExtendOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizedClientAPI.ExtendHotspotAuthedClient(context.Background(), omadacId, siteId, id).ExtendOpenApiVO(extendOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.ExtendHotspotAuthedClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtendHotspotAuthedClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.ExtendHotspotAuthedClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Authed record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendHotspotAuthedClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **extendOpenApiVO** | [**ExtendOpenApiVO**](ExtendOpenApiVO.md) |  | 

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


## GetHotspotAuthedClients

> OperationResponseGridVOAuthClientOpenApiVO GetHotspotAuthedClients(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsMac(sortsMac).SortsIp(sortsIp).SortsAuthType(sortsAuthType).SortsSsidOrNetwork(sortsSsidOrNetwork).SortsDownload(sortsDownload).SortsUpload(sortsUpload).SortsStatus(sortsStatus).SortsStart(sortsStart).SortsEnd(sortsEnd).SortsDuration(sortsDuration).SearchKey(searchKey).Execute()

Get authentication record list



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
	resp, r, err := apiClient.AuthorizedClientAPI.GetHotspotAuthedClients(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsMac(sortsMac).SortsIp(sortsIp).SortsAuthType(sortsAuthType).SortsSsidOrNetwork(sortsSsidOrNetwork).SortsDownload(sortsDownload).SortsUpload(sortsUpload).SortsStatus(sortsStatus).SortsStart(sortsStart).SortsEnd(sortsEnd).SortsDuration(sortsDuration).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.GetHotspotAuthedClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHotspotAuthedClients`: OperationResponseGridVOAuthClientOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.GetHotspotAuthedClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHotspotAuthedClientsRequest struct via the builder pattern


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

[**OperationResponseGridVOAuthClientOpenApiVO**](OperationResponseGridVOAuthClientOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHotspotStatistic

> OperationResponseHotspotStatisticVO GetHotspotStatistic(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get hotspot statistic



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
	resp, r, err := apiClient.AuthorizedClientAPI.GetHotspotStatistic(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizedClientAPI.GetHotspotStatistic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHotspotStatistic`: OperationResponseHotspotStatisticVO
	fmt.Fprintf(os.Stdout, "Response from `AuthorizedClientAPI.GetHotspotStatistic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHotspotStatisticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseHotspotStatisticVO**](OperationResponseHotspotStatisticVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


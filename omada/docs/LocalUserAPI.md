# LocalUserAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearLocalUserDynamicMac**](LocalUserAPI.md#clearlocaluserdynamicmac) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/localusers/{id}/clear-dynamic-mac | Clear dynamic mac
[**CreateLocalUser**](LocalUserAPI.md#createlocaluser) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/localusers | Create a new local user
[**DeleteLocalUser**](LocalUserAPI.md#deletelocaluser) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/localusers/{id} | Delete an existing localuser
[**DownloadLocalUsers**](LocalUserAPI.md#downloadlocalusers) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/files/hotspot/local-users | Download local user file (excel or csv) by localhost
[**GetLocalUser**](LocalUserAPI.md#getlocaluser) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/localusers/{id} | Get a local user for given localuserId
[**GetLocalUsers**](LocalUserAPI.md#getlocalusers) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/localusers | Get local user list
[**ModifyLocalUser**](LocalUserAPI.md#modifylocaluser) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/localusers/{id} | Modify an existing localuser
[**UploadLocalUsers**](LocalUserAPI.md#uploadlocalusers) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/files/hotspot/local-users | Upload local user file (excel or csv) by localhost



## ClearLocalUserDynamicMac

> OperationResponseWithoutResult ClearLocalUserDynamicMac(ctx, omadacId, siteId, id).Execute()

Clear dynamic mac



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
	resp, r, err := apiClient.LocalUserAPI.ClearLocalUserDynamicMac(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalUserAPI.ClearLocalUserDynamicMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearLocalUserDynamicMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LocalUserAPI.ClearLocalUserDynamicMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiClearLocalUserDynamicMacRequest struct via the builder pattern


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


## CreateLocalUser

> OperationResponseCreatedResIdOpenApiVO CreateLocalUser(ctx, omadacId, siteId).CreateLocalUserOpenApiVO(createLocalUserOpenApiVO).Execute()

Create a new local user



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
	createLocalUserOpenApiVO := *openapiclient.NewCreateLocalUserOpenApiVO(int32(123), false, int64(123), int32(123), "Password_example", []string{"Portals_example"}, *openapiclient.NewRateLimitOpenApiVO(int32(123)), false, "UserName_example") // CreateLocalUserOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalUserAPI.CreateLocalUser(context.Background(), omadacId, siteId).CreateLocalUserOpenApiVO(createLocalUserOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalUserAPI.CreateLocalUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLocalUser`: OperationResponseCreatedResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LocalUserAPI.CreateLocalUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocalUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createLocalUserOpenApiVO** | [**CreateLocalUserOpenApiVO**](CreateLocalUserOpenApiVO.md) |  | 

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


## DeleteLocalUser

> OperationResponseWithoutResult DeleteLocalUser(ctx, omadacId, siteId, id).Execute()

Delete an existing localuser



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
	id := "id_example" // string | Local user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalUserAPI.DeleteLocalUser(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalUserAPI.DeleteLocalUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLocalUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LocalUserAPI.DeleteLocalUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Local user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalUserRequest struct via the builder pattern


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


## DownloadLocalUsers

> OperationResponse DownloadLocalUsers(ctx, omadacId, siteId).Portals(portals).FileType(fileType).SearchKey(searchKey).Execute()

Download local user file (excel or csv) by localhost



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
	portals := "portals_example" // string | Portal ids of the portals to download local user file from.
	fileType := "fileType_example" // string | Local user file format: csv or xlsx.
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name,username (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalUserAPI.DownloadLocalUsers(context.Background(), omadacId, siteId).Portals(portals).FileType(fileType).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalUserAPI.DownloadLocalUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLocalUsers`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalUserAPI.DownloadLocalUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLocalUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portals** | **string** | Portal ids of the portals to download local user file from. | 
 **fileType** | **string** | Local user file format: csv or xlsx. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name,username | 

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


## GetLocalUser

> OperationResponseLocalUserOpenApiVO GetLocalUser(ctx, omadacId, siteId, id).Execute()

Get a local user for given localuserId



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
	id := "id_example" // string | local user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalUserAPI.GetLocalUser(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalUserAPI.GetLocalUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalUser`: OperationResponseLocalUserOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LocalUserAPI.GetLocalUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | local user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseLocalUserOpenApiVO**](OperationResponseLocalUserOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalUsers

> OperationResponseGridVOLocalUserOpenApiVO GetLocalUsers(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()

Get local user list



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field user_name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalUserAPI.GetLocalUsers(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalUserAPI.GetLocalUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalUsers`: OperationResponseGridVOLocalUserOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LocalUserAPI.GetLocalUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field user_name | 

### Return type

[**OperationResponseGridVOLocalUserOpenApiVO**](OperationResponseGridVOLocalUserOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyLocalUser

> OperationResponseWithoutResult ModifyLocalUser(ctx, omadacId, siteId, id).ModifyLocalUserOpenApiVO(modifyLocalUserOpenApiVO).Execute()

Modify an existing localuser



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
	modifyLocalUserOpenApiVO := *openapiclient.NewModifyLocalUserOpenApiVO(int32(123), false, int64(123), int32(123), "Password_example", []string{"Portals_example"}, *openapiclient.NewRateLimitOpenApiVO(int32(123)), false) // ModifyLocalUserOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalUserAPI.ModifyLocalUser(context.Background(), omadacId, siteId, id).ModifyLocalUserOpenApiVO(modifyLocalUserOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalUserAPI.ModifyLocalUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLocalUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LocalUserAPI.ModifyLocalUser`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyLocalUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyLocalUserOpenApiVO** | [**ModifyLocalUserOpenApiVO**](ModifyLocalUserOpenApiVO.md) |  | 

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


## UploadLocalUsers

> OperationResponse UploadLocalUsers(ctx, omadacId, siteId).UploadLocalUsersRequest(uploadLocalUsersRequest).Execute()

Upload local user file (excel or csv) by localhost



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
	uploadLocalUsersRequest := *openapiclient.NewUploadLocalUsersRequest(*openapiclient.NewHotspotPortalsOpenApiVO([]string{"Portals_example"}), "TODO") // UploadLocalUsersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalUserAPI.UploadLocalUsers(context.Background(), omadacId, siteId).UploadLocalUsersRequest(uploadLocalUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalUserAPI.UploadLocalUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLocalUsers`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalUserAPI.UploadLocalUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadLocalUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadLocalUsersRequest** | [**UploadLocalUsersRequest**](UploadLocalUsersRequest.md) |  | 

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


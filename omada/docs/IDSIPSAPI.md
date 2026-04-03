# IDSIPSAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAllowList**](IDSIPSAPI.md#createallowlist) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/allow-list | Create new allow list
[**DeleteAllowList**](IDSIPSAPI.md#deleteallowlist) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/allow-list/{entryId} | Delete allow list
[**DeleteBlockList**](IDSIPSAPI.md#deleteblocklist) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/block-list/{entryId} | Delete block list
[**DeleteSignature**](IDSIPSAPI.md#deletesignature) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/signature/{signId} | Delete signature
[**EditAllowList**](IDSIPSAPI.md#editallowlist) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/allow-list | Edit allow list
[**GetGridAllowList**](IDSIPSAPI.md#getgridallowlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/grid/allow-list | Get grid ips allow list
[**GetGridBlockList**](IDSIPSAPI.md#getgridblocklist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/grid/block-list | Get grid IPS block list
[**GetGridSignature**](IDSIPSAPI.md#getgridsignature) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/signature | Get signature list
[**GetIpsConfig**](IDSIPSAPI.md#getipsconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips | Get site IDS/IPS config
[**ModifyIpsConfig**](IDSIPSAPI.md#modifyipsconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips | Modify site IDS/IPS config
[**ModifySignature**](IDSIPSAPI.md#modifysignature) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/network-security/ips/signature/{signId} | Modify signature



## CreateAllowList

> OperationResponse CreateAllowList(ctx, omadacId, siteId).NewIPSAllowListEntry(newIPSAllowListEntry).Execute()

Create new allow list



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
	newIPSAllowListEntry := *openapiclient.NewNewIPSAllowListEntry() // NewIPSAllowListEntry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.CreateAllowList(context.Background(), omadacId, siteId).NewIPSAllowListEntry(newIPSAllowListEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.CreateAllowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAllowList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.CreateAllowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newIPSAllowListEntry** | [**NewIPSAllowListEntry**](NewIPSAllowListEntry.md) |  | 

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


## DeleteAllowList

> OperationResponse DeleteAllowList(ctx, omadacId, siteId, entryId).Execute()

Delete allow list



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
	entryId := "entryId_example" // string | Allow entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.DeleteAllowList(context.Background(), omadacId, siteId, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.DeleteAllowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllowList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.DeleteAllowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**entryId** | **string** | Allow entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllowListRequest struct via the builder pattern


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


## DeleteBlockList

> OperationResponse DeleteBlockList(ctx, omadacId, siteId, entryId).Execute()

Delete block list



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
	entryId := "entryId_example" // string | Block entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.DeleteBlockList(context.Background(), omadacId, siteId, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.DeleteBlockList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.DeleteBlockList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**entryId** | **string** | Block entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockListRequest struct via the builder pattern


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


## DeleteSignature

> OperationResponse DeleteSignature(ctx, omadacId, siteId, signId).Execute()

Delete signature



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
	signId := "signId_example" // string | Unique Signature ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.DeleteSignature(context.Background(), omadacId, siteId, signId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.DeleteSignature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSignature`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.DeleteSignature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**signId** | **string** | Unique Signature ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSignatureRequest struct via the builder pattern


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


## EditAllowList

> OperationResponse EditAllowList(ctx, omadacId, siteId).ModifyIPSAllowListEntry(modifyIPSAllowListEntry).Execute()

Edit allow list



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
	modifyIPSAllowListEntry := *openapiclient.NewModifyIPSAllowListEntry("Id_example") // ModifyIPSAllowListEntry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.EditAllowList(context.Background(), omadacId, siteId).ModifyIPSAllowListEntry(modifyIPSAllowListEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.EditAllowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAllowList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.EditAllowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAllowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyIPSAllowListEntry** | [**ModifyIPSAllowListEntry**](ModifyIPSAllowListEntry.md) |  | 

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


## GetGridAllowList

> OperationResponse GetGridAllowList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get grid ips allow list



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field IP / Subnet (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.GetGridAllowList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.GetGridAllowList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAllowList`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.GetGridAllowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAllowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field IP / Subnet | 

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


## GetGridBlockList

> OperationResponseGridVOIpsBlockListEntry GetGridBlockList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get grid IPS block list



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field Name/Source Ip/Destination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.GetGridBlockList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.GetGridBlockList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridBlockList`: OperationResponseGridVOIpsBlockListEntry
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.GetGridBlockList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridBlockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field Name/Source Ip/Destination | 

### Return type

[**OperationResponseGridVOIpsBlockListEntry**](OperationResponseGridVOIpsBlockListEntry.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSignature

> OperationResponseGridVOIpsSignatureInfo GetGridSignature(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get signature list



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
	resp, r, err := apiClient.IDSIPSAPI.GetGridSignature(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.GetGridSignature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSignature`: OperationResponseGridVOIpsSignatureInfo
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.GetGridSignature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOIpsSignatureInfo**](OperationResponseGridVOIpsSignatureInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpsConfig

> OperationResponseIpsInfo GetIpsConfig(ctx, omadacId, siteId).Execute()

Get site IDS/IPS config



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
	resp, r, err := apiClient.IDSIPSAPI.GetIpsConfig(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.GetIpsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpsConfig`: OperationResponseIpsInfo
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.GetIpsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIpsInfo**](OperationResponseIpsInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIpsConfig

> OperationResponse ModifyIpsConfig(ctx, omadacId, siteId).IpsInfo(ipsInfo).Execute()

Modify site IDS/IPS config



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
	ipsInfo := *openapiclient.NewIpsInfo(false) // IpsInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.ModifyIpsConfig(context.Background(), omadacId, siteId).IpsInfo(ipsInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.ModifyIpsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpsConfig`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.ModifyIpsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ipsInfo** | [**IpsInfo**](IpsInfo.md) |  | 

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


## ModifySignature

> OperationResponse ModifySignature(ctx, omadacId, siteId, signId).IpsSignatureConfig(ipsSignatureConfig).Execute()

Modify signature



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
	signId := "signId_example" // string | Unique Signature ID
	ipsSignatureConfig := *openapiclient.NewIpsSignatureConfig(int32(1)) // IpsSignatureConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IDSIPSAPI.ModifySignature(context.Background(), omadacId, siteId, signId).IpsSignatureConfig(ipsSignatureConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IDSIPSAPI.ModifySignature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySignature`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IDSIPSAPI.ModifySignature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**signId** | **string** | Unique Signature ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ipsSignatureConfig** | [**IpsSignatureConfig**](IpsSignatureConfig.md) |  | 

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


# MSPSSOAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalUserGroup1**](MSPSSOAPI.md#createexternalusergroup1) | **Post** /openapi/v1/msp/{mspId}/sso/external-user-groups | Create new msp external user group
[**CreateIdpMetadata1**](MSPSSOAPI.md#createidpmetadata1) | **Post** /openapi/v1/msp/{mspId}/sso/saml-idps | Create new msp idp
[**DeleteExternalUserGroup1**](MSPSSOAPI.md#deleteexternalusergroup1) | **Delete** /openapi/v1/msp/{mspId}/sso/external-user-groups/{externalUserGroupId} | Delete an existing msp external user group
[**DeleteIdpMetadata1**](MSPSSOAPI.md#deleteidpmetadata1) | **Delete** /openapi/v1/msp/{mspId}/sso/saml-idps/{idpId} | Delete an existing msp idp
[**GetExternalUserGroupList1**](MSPSSOAPI.md#getexternalusergrouplist1) | **Get** /openapi/v1/msp/{mspId}/sso/external-user-groups | Get msp external user group list
[**GetExternalUserList1**](MSPSSOAPI.md#getexternaluserlist1) | **Get** /openapi/v1/msp/{mspId}/sso/external-users | Get msp external user list
[**GetIdpMetadataList1**](MSPSSOAPI.md#getidpmetadatalist1) | **Get** /openapi/v1/msp/{mspId}/sso/saml-idps | Get msp idp list
[**ModifyExternalUserGroup1**](MSPSSOAPI.md#modifyexternalusergroup1) | **Put** /openapi/v1/msp/{mspId}/sso/external-user-groups/{externalUserGroupId} | Modify an existing msp external user group
[**ModifyIdpMetadata1**](MSPSSOAPI.md#modifyidpmetadata1) | **Put** /openapi/v1/msp/{mspId}/sso/saml-idps/{idpId} | Modify an existing msp idp



## CreateExternalUserGroup1

> OperationResponseResponseIdVO CreateExternalUserGroup1(ctx, mspId).MspExternalUserGroupOpenApiVO(mspExternalUserGroupOpenApiVO).Execute()

Create new msp external user group



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
	mspExternalUserGroupOpenApiVO := *openapiclient.NewMspExternalUserGroupOpenApiVO(false, "CustomerRoleId_example", "Name_example", "RoleId_example") // MspExternalUserGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.CreateExternalUserGroup1(context.Background(), mspId).MspExternalUserGroupOpenApiVO(mspExternalUserGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.CreateExternalUserGroup1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalUserGroup1`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.CreateExternalUserGroup1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalUserGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mspExternalUserGroupOpenApiVO** | [**MspExternalUserGroupOpenApiVO**](MspExternalUserGroupOpenApiVO.md) |  | 

### Return type

[**OperationResponseResponseIdVO**](OperationResponseResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdpMetadata1

> OperationResponseResponseIdVO CreateIdpMetadata1(ctx, mspId).IdpMetadataOpenApiVO(idpMetadataOpenApiVO).Execute()

Create new msp idp



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
	idpMetadataOpenApiVO := *openapiclient.NewIdpMetadataOpenApiVO("EntityId_example", "LoginUrl_example", "Name_example", "X509Certificate_example") // IdpMetadataOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.CreateIdpMetadata1(context.Background(), mspId).IdpMetadataOpenApiVO(idpMetadataOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.CreateIdpMetadata1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdpMetadata1`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.CreateIdpMetadata1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpMetadata1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpMetadataOpenApiVO** | [**IdpMetadataOpenApiVO**](IdpMetadataOpenApiVO.md) |  | 

### Return type

[**OperationResponseResponseIdVO**](OperationResponseResponseIdVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalUserGroup1

> OperationResponseWithoutResult DeleteExternalUserGroup1(ctx, mspId, externalUserGroupId).Execute()

Delete an existing msp external user group



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
	externalUserGroupId := "externalUserGroupId_example" // string | External user group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.DeleteExternalUserGroup1(context.Background(), mspId, externalUserGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.DeleteExternalUserGroup1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalUserGroup1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.DeleteExternalUserGroup1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**externalUserGroupId** | **string** | External user group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalUserGroup1Request struct via the builder pattern


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


## DeleteIdpMetadata1

> OperationResponseWithoutResult DeleteIdpMetadata1(ctx, mspId, idpId).Execute()

Delete an existing msp idp



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
	idpId := "idpId_example" // string | Idp ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.DeleteIdpMetadata1(context.Background(), mspId, idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.DeleteIdpMetadata1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIdpMetadata1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.DeleteIdpMetadata1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**idpId** | **string** | Idp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpMetadata1Request struct via the builder pattern


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


## GetExternalUserGroupList1

> OperationResponseGridVOMspExternalUserGroupDetailOpenApiVO GetExternalUserGroupList1(ctx, mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()

Get msp external user group list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.GetExternalUserGroupList1(context.Background(), mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.GetExternalUserGroupList1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalUserGroupList1`: OperationResponseGridVOMspExternalUserGroupDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.GetExternalUserGroupList1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalUserGroupList1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseGridVOMspExternalUserGroupDetailOpenApiVO**](OperationResponseGridVOMspExternalUserGroupDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalUserList1

> OperationResponseGridVOMspExternalUserDetailOpenApiVO GetExternalUserList1(ctx, mspId).Page(page).PageSize(pageSize).SortsUserName(sortsUserName).SearchKey(searchKey).Execute()

Get msp external user list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsUserName := "sortsUserName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field userName (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.GetExternalUserList1(context.Background(), mspId).Page(page).PageSize(pageSize).SortsUserName(sortsUserName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.GetExternalUserList1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalUserList1`: OperationResponseGridVOMspExternalUserDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.GetExternalUserList1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalUserList1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsUserName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field userName | 

### Return type

[**OperationResponseGridVOMspExternalUserDetailOpenApiVO**](OperationResponseGridVOMspExternalUserDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpMetadataList1

> OperationResponseGridVOIdpMetadataDetailOpenApiVO GetIdpMetadataList1(ctx, mspId).Page(page).PageSize(pageSize).SortsIdpName(sortsIdpName).Execute()

Get msp idp list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsIdpName := "sortsIdpName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.GetIdpMetadataList1(context.Background(), mspId).Page(page).PageSize(pageSize).SortsIdpName(sortsIdpName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.GetIdpMetadataList1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdpMetadataList1`: OperationResponseGridVOIdpMetadataDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.GetIdpMetadataList1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpMetadataList1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsIdpName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseGridVOIdpMetadataDetailOpenApiVO**](OperationResponseGridVOIdpMetadataDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyExternalUserGroup1

> OperationResponseWithoutResult ModifyExternalUserGroup1(ctx, mspId, externalUserGroupId).MspExternalUserGroupOpenApiVO(mspExternalUserGroupOpenApiVO).Execute()

Modify an existing msp external user group



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
	externalUserGroupId := "externalUserGroupId_example" // string | External user group ID
	mspExternalUserGroupOpenApiVO := *openapiclient.NewMspExternalUserGroupOpenApiVO(false, "CustomerRoleId_example", "Name_example", "RoleId_example") // MspExternalUserGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.ModifyExternalUserGroup1(context.Background(), mspId, externalUserGroupId).MspExternalUserGroupOpenApiVO(mspExternalUserGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.ModifyExternalUserGroup1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyExternalUserGroup1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.ModifyExternalUserGroup1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**externalUserGroupId** | **string** | External user group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyExternalUserGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mspExternalUserGroupOpenApiVO** | [**MspExternalUserGroupOpenApiVO**](MspExternalUserGroupOpenApiVO.md) |  | 

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


## ModifyIdpMetadata1

> OperationResponseWithoutResult ModifyIdpMetadata1(ctx, mspId, idpId).IdpMetadataOpenApiVO(idpMetadataOpenApiVO).Execute()

Modify an existing msp idp



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
	idpId := "idpId_example" // string | Idp ID
	idpMetadataOpenApiVO := *openapiclient.NewIdpMetadataOpenApiVO("EntityId_example", "LoginUrl_example", "Name_example", "X509Certificate_example") // IdpMetadataOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPSSOAPI.ModifyIdpMetadata1(context.Background(), mspId, idpId).IdpMetadataOpenApiVO(idpMetadataOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPSSOAPI.ModifyIdpMetadata1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIdpMetadata1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MSPSSOAPI.ModifyIdpMetadata1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**idpId** | **string** | Idp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIdpMetadata1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idpMetadataOpenApiVO** | [**IdpMetadataOpenApiVO**](IdpMetadataOpenApiVO.md) |  | 

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


# SSOAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalUserGroup**](SSOAPI.md#createexternalusergroup) | **Post** /openapi/v1/{omadacId}/sso/external-user-groups | Create new external user group
[**CreateIdpMetadata**](SSOAPI.md#createidpmetadata) | **Post** /openapi/v1/{omadacId}/sso/saml-idps | Create new idp
[**DeleteExternalUserGroup**](SSOAPI.md#deleteexternalusergroup) | **Delete** /openapi/v1/{omadacId}/sso/external-user-groups/{externalUserGroupId} | Delete an existing external user group
[**DeleteIdpMetadata**](SSOAPI.md#deleteidpmetadata) | **Delete** /openapi/v1/{omadacId}/sso/saml-idps/{idpId} | Delete an existing idp
[**GetExternalUserGroupList**](SSOAPI.md#getexternalusergrouplist) | **Get** /openapi/v1/{omadacId}/sso/external-user-groups | Get external user group list
[**GetExternalUserList**](SSOAPI.md#getexternaluserlist) | **Get** /openapi/v1/{omadacId}/sso/external-users | Get external user list
[**GetIdpMetadataList**](SSOAPI.md#getidpmetadatalist) | **Get** /openapi/v1/{omadacId}/sso/saml-idps | Get idp list
[**ModifyExternalUserGroup**](SSOAPI.md#modifyexternalusergroup) | **Put** /openapi/v1/{omadacId}/sso/external-user-groups/{externalUserGroupId} | Modify an existing external user group
[**ModifyIdpMetadata**](SSOAPI.md#modifyidpmetadata) | **Put** /openapi/v1/{omadacId}/sso/saml-idps/{idpId} | Modify an existing idp



## CreateExternalUserGroup

> OperationResponseResponseIdVO CreateExternalUserGroup(ctx, omadacId).ExternalUserGroupOpenApiVO(externalUserGroupOpenApiVO).Execute()

Create new external user group



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
	externalUserGroupOpenApiVO := *openapiclient.NewExternalUserGroupOpenApiVO(false, "Name_example", "RoleId_example") // ExternalUserGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.CreateExternalUserGroup(context.Background(), omadacId).ExternalUserGroupOpenApiVO(externalUserGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.CreateExternalUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalUserGroup`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.CreateExternalUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalUserGroupOpenApiVO** | [**ExternalUserGroupOpenApiVO**](ExternalUserGroupOpenApiVO.md) |  | 

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


## CreateIdpMetadata

> OperationResponseResponseIdVO CreateIdpMetadata(ctx, omadacId).IdpMetadataOpenApiVO(idpMetadataOpenApiVO).Execute()

Create new idp



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
	idpMetadataOpenApiVO := *openapiclient.NewIdpMetadataOpenApiVO("EntityId_example", "LoginUrl_example", "Name_example", "X509Certificate_example") // IdpMetadataOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.CreateIdpMetadata(context.Background(), omadacId).IdpMetadataOpenApiVO(idpMetadataOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.CreateIdpMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdpMetadata`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.CreateIdpMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpMetadataRequest struct via the builder pattern


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


## DeleteExternalUserGroup

> OperationResponseWithoutResult DeleteExternalUserGroup(ctx, omadacId, externalUserGroupId).Execute()

Delete an existing external user group



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
	externalUserGroupId := "externalUserGroupId_example" // string | External user group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.DeleteExternalUserGroup(context.Background(), omadacId, externalUserGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.DeleteExternalUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExternalUserGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.DeleteExternalUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**externalUserGroupId** | **string** | External user group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalUserGroupRequest struct via the builder pattern


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


## DeleteIdpMetadata

> OperationResponseWithoutResult DeleteIdpMetadata(ctx, omadacId, idpId).Execute()

Delete an existing idp



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
	idpId := "idpId_example" // string | IdP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.DeleteIdpMetadata(context.Background(), omadacId, idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.DeleteIdpMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIdpMetadata`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.DeleteIdpMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**idpId** | **string** | IdP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpMetadataRequest struct via the builder pattern


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


## GetExternalUserGroupList

> OperationResponseGridVOExternalUserGroupDetailOpenApiVO GetExternalUserGroupList(ctx, omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()

Get external user group list



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
	omadacId := "omadacId_example" // string | 
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.GetExternalUserGroupList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.GetExternalUserGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalUserGroupList`: OperationResponseGridVOExternalUserGroupDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.GetExternalUserGroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalUserGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field name | 

### Return type

[**OperationResponseGridVOExternalUserGroupDetailOpenApiVO**](OperationResponseGridVOExternalUserGroupDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalUserList

> OperationResponseGridVOExternalUserDetailOpenApiVO GetExternalUserList(ctx, omadacId).Page(page).PageSize(pageSize).SortsUserName(sortsUserName).SearchKey(searchKey).Execute()

Get external user list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsUserName := "sortsUserName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field userName (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.GetExternalUserList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsUserName(sortsUserName).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.GetExternalUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalUserList`: OperationResponseGridVOExternalUserDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.GetExternalUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsUserName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field userName | 

### Return type

[**OperationResponseGridVOExternalUserDetailOpenApiVO**](OperationResponseGridVOExternalUserDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpMetadataList

> OperationResponseGridVOIdpMetadataDetailOpenApiVO GetIdpMetadataList(ctx, omadacId).Page(page).PageSize(pageSize).SortsIdpName(sortsIdpName).Execute()

Get idp list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsIdpName := "sortsIdpName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.GetIdpMetadataList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsIdpName(sortsIdpName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.GetIdpMetadataList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdpMetadataList`: OperationResponseGridVOIdpMetadataDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.GetIdpMetadataList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpMetadataListRequest struct via the builder pattern


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


## ModifyExternalUserGroup

> OperationResponseWithoutResult ModifyExternalUserGroup(ctx, omadacId, externalUserGroupId).ExternalUserGroupOpenApiVO(externalUserGroupOpenApiVO).Execute()

Modify an existing external user group



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
	externalUserGroupId := "externalUserGroupId_example" // string | External user group ID
	externalUserGroupOpenApiVO := *openapiclient.NewExternalUserGroupOpenApiVO(false, "Name_example", "RoleId_example") // ExternalUserGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.ModifyExternalUserGroup(context.Background(), omadacId, externalUserGroupId).ExternalUserGroupOpenApiVO(externalUserGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.ModifyExternalUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyExternalUserGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.ModifyExternalUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**externalUserGroupId** | **string** | External user group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyExternalUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalUserGroupOpenApiVO** | [**ExternalUserGroupOpenApiVO**](ExternalUserGroupOpenApiVO.md) |  | 

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


## ModifyIdpMetadata

> OperationResponseWithoutResult ModifyIdpMetadata(ctx, omadacId, idpId).IdpMetadataOpenApiVO(idpMetadataOpenApiVO).Execute()

Modify an existing idp



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
	idpId := "idpId_example" // string | IdP id
	idpMetadataOpenApiVO := *openapiclient.NewIdpMetadataOpenApiVO("EntityId_example", "LoginUrl_example", "Name_example", "X509Certificate_example") // IdpMetadataOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.ModifyIdpMetadata(context.Background(), omadacId, idpId).IdpMetadataOpenApiVO(idpMetadataOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.ModifyIdpMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIdpMetadata`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.ModifyIdpMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**idpId** | **string** | IdP id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIdpMetadataRequest struct via the builder pattern


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


# UserAndRoleAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewRole**](UserAndRoleAPI.md#createnewrole) | **Post** /openapi/v1/{omadacId}/roles | Create new role
[**CreateNewUser**](UserAndRoleAPI.md#createnewuser) | **Post** /openapi/v1/{omadacId}/users | Create new user
[**DeleteRole**](UserAndRoleAPI.md#deleterole) | **Delete** /openapi/v1/{omadacId}/roles/{roleId} | Delete an existing role
[**DeleteUser**](UserAndRoleAPI.md#deleteuser) | **Delete** /openapi/v1/{omadacId}/users/{userID} | Delete an existing user
[**GetAllCloudUsersExcludeRoot**](UserAndRoleAPI.md#getallcloudusersexcluderoot) | **Get** /openapi/v1/{omadacId}/users/cloud | Get all cloud users exclude owner
[**GetAllLocalUsersExcludeRoot**](UserAndRoleAPI.md#getalllocalusersexcluderoot) | **Get** /openapi/v1/{omadacId}/users/local | Get all local users exclude owner
[**GetAllRoles**](UserAndRoleAPI.md#getallroles) | **Get** /openapi/v1/{omadacId}/roles | Get role list
[**GetAppGridUsers**](UserAndRoleAPI.md#getappgridusers) | **Get** /openapi/v1/{omadacId}/all-users | Get user list for app
[**GetAvailableRole**](UserAndRoleAPI.md#getavailablerole) | **Get** /openapi/v1/{omadacId}/roles/available | Get available roles for creating new user
[**GetGlobalMFAStatus**](UserAndRoleAPI.md#getglobalmfastatus) | **Get** /openapi/v1/{omadacId}/mfa/status | Get the status of the system&#39;s Two-Factor Authentication.
[**GetGridUsers**](UserAndRoleAPI.md#getgridusers) | **Get** /openapi/v1/{omadacId}/users | Get user list
[**GetRole**](UserAndRoleAPI.md#getrole) | **Get** /openapi/v1/{omadacId}/roles/{roleId} | Get role info
[**GetUser**](UserAndRoleAPI.md#getuser) | **Get** /openapi/v1/{omadacId}/users/{userID} | Get user info
[**ModifyGlobalMFAStatus**](UserAndRoleAPI.md#modifyglobalmfastatus) | **Post** /openapi/v1/{omadacId}/mfa/status | Modify the status of the system&#39;s Two-Factor Authentication.
[**ModifyRole**](UserAndRoleAPI.md#modifyrole) | **Put** /openapi/v1/{omadacId}/roles/{roleId} | Modify an existing role
[**ModifyUser**](UserAndRoleAPI.md#modifyuser) | **Put** /openapi/v1/{omadacId}/users/{userID} | Modify an existing user
[**ModifyUserBatch**](UserAndRoleAPI.md#modifyuserbatch) | **Patch** /openapi/v1/{omadacId}/users/batch-temporary-users | Batch modify users.
[**Reinvite**](UserAndRoleAPI.md#reinvite) | **Post** /openapi/v1/{omadacId}/account/users/{userID}/re-invite | Re-Invite an existing user
[**TransferMspRoot**](UserAndRoleAPI.md#transfermsproot) | **Post** /openapi/v1/msp/{mspId}/owner/transfer | Transfer MSP owner permission to an existing user
[**TransferRoot**](UserAndRoleAPI.md#transferroot) | **Post** /openapi/v1/{omadacId}/owner/transfer | Transfer owner permission to an existing user



## CreateNewRole

> OperationResponseCreateRoleResultVO CreateNewRole(ctx, omadacId).ModifyControllerRoleVO(modifyControllerRoleVO).Execute()

Create new role



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
	modifyControllerRoleVO := *openapiclient.NewModifyControllerRoleVO("Name_example") // ModifyControllerRoleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.CreateNewRole(context.Background(), omadacId).ModifyControllerRoleVO(modifyControllerRoleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.CreateNewRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewRole`: OperationResponseCreateRoleResultVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.CreateNewRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyControllerRoleVO** | [**ModifyControllerRoleVO**](ModifyControllerRoleVO.md) |  | 

### Return type

[**OperationResponseCreateRoleResultVO**](OperationResponseCreateRoleResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewUser

> OperationResponseCreateUserResultVO CreateNewUser(ctx, omadacId).CreateUserVO(createUserVO).Execute()

Create new user



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
	createUserVO := *openapiclient.NewCreateUserVO(false, "Name_example", "RoleId_example", int32(123)) // CreateUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.CreateNewUser(context.Background(), omadacId).CreateUserVO(createUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.CreateNewUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewUser`: OperationResponseCreateUserResultVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.CreateNewUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserVO** | [**CreateUserVO**](CreateUserVO.md) |  | 

### Return type

[**OperationResponseCreateUserResultVO**](OperationResponseCreateUserResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> OperationResponseWithoutResult DeleteRole(ctx, omadacId, roleId).Execute()

Delete an existing role



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
	roleId := "roleId_example" // string | Role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.DeleteRole(context.Background(), omadacId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRole`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.DeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## DeleteUser

> OperationResponseWithoutResult DeleteUser(ctx, omadacId, userID).DeleteUserVO(deleteUserVO).Execute()

Delete an existing user



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
	userID := "userID_example" // string | userID
	deleteUserVO := *openapiclient.NewDeleteUserVO() // DeleteUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.DeleteUser(context.Background(), omadacId, userID).DeleteUserVO(deleteUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**userID** | **string** | userID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteUserVO** | [**DeleteUserVO**](DeleteUserVO.md) |  | 

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


## GetAllCloudUsersExcludeRoot

> OperationResponseGetUserListResponseVO GetAllCloudUsersExcludeRoot(ctx, omadacId).Execute()

Get all cloud users exclude owner



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
	resp, r, err := apiClient.UserAndRoleAPI.GetAllCloudUsersExcludeRoot(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetAllCloudUsersExcludeRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCloudUsersExcludeRoot`: OperationResponseGetUserListResponseVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetAllCloudUsersExcludeRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCloudUsersExcludeRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseGetUserListResponseVO**](OperationResponseGetUserListResponseVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllLocalUsersExcludeRoot

> OperationResponseGetUserListResponseVO GetAllLocalUsersExcludeRoot(ctx, omadacId).Execute()

Get all local users exclude owner



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
	resp, r, err := apiClient.UserAndRoleAPI.GetAllLocalUsersExcludeRoot(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetAllLocalUsersExcludeRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllLocalUsersExcludeRoot`: OperationResponseGetUserListResponseVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetAllLocalUsersExcludeRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLocalUsersExcludeRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseGetUserListResponseVO**](OperationResponseGetUserListResponseVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRoles

> OperationResponseListControllerRoleDetailVO GetAllRoles(ctx, omadacId).Execute()

Get role list



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
	resp, r, err := apiClient.UserAndRoleAPI.GetAllRoles(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetAllRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllRoles`: OperationResponseListControllerRoleDetailVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetAllRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseListControllerRoleDetailVO**](OperationResponseListControllerRoleDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppGridUsers

> OperationResponseGridVOControllerUserAppVO GetAppGridUsers(ctx, omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsRoleId(sortsRoleId).SortsEmail(sortsEmail).SearchKey(searchKey).Execute()

Get user list for app



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
	sortsRoleId := "sortsRoleId_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsEmail := "sortsEmail_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field userName (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.GetAppGridUsers(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsRoleId(sortsRoleId).SortsEmail(sortsEmail).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetAppGridUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppGridUsers`: OperationResponseGridVOControllerUserAppVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetAppGridUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppGridUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsRoleId** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsEmail** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field userName | 

### Return type

[**OperationResponseGridVOControllerUserAppVO**](OperationResponseGridVOControllerUserAppVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableRole

> OperationResponseListRoleBriefVO GetAvailableRole(ctx, omadacId).Execute()

Get available roles for creating new user



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
	resp, r, err := apiClient.UserAndRoleAPI.GetAvailableRole(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetAvailableRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableRole`: OperationResponseListRoleBriefVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetAvailableRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseListRoleBriefVO**](OperationResponseListRoleBriefVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalMFAStatus

> OperationResponse GetGlobalMFAStatus(ctx, omadacId).Execute()

Get the status of the system's Two-Factor Authentication.



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
	resp, r, err := apiClient.UserAndRoleAPI.GetGlobalMFAStatus(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetGlobalMFAStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalMFAStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetGlobalMFAStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalMFAStatusRequest struct via the builder pattern


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


## GetGridUsers

> OperationResponseGridVOUserDetailVO GetGridUsers(ctx, omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsRoleId(sortsRoleId).SortsEmail(sortsEmail).SearchKey(searchKey).Execute()

Get user list



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
	sortsRoleId := "sortsRoleId_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsEmail := "sortsEmail_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field userName (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.GetGridUsers(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsRoleId(sortsRoleId).SortsEmail(sortsEmail).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetGridUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridUsers`: OperationResponseGridVOUserDetailVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetGridUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsRoleId** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsEmail** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field userName | 

### Return type

[**OperationResponseGridVOUserDetailVO**](OperationResponseGridVOUserDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> OperationResponseControllerRoleDetailVO GetRole(ctx, omadacId, roleId).Execute()

Get role info



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
	roleId := "roleId_example" // string | Role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.GetRole(context.Background(), omadacId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: OperationResponseControllerRoleDetailVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseControllerRoleDetailVO**](OperationResponseControllerRoleDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> OperationResponseUserDetailVO GetUser(ctx, omadacId, userID).Execute()

Get user info



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
	userID := "userID_example" // string | userID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.GetUser(context.Background(), omadacId, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: OperationResponseUserDetailVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**userID** | **string** | userID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUserDetailVO**](OperationResponseUserDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGlobalMFAStatus

> OperationResponse ModifyGlobalMFAStatus(ctx, omadacId).MFAConfigOpenApiVO(mFAConfigOpenApiVO).Execute()

Modify the status of the system's Two-Factor Authentication.



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
	mFAConfigOpenApiVO := *openapiclient.NewMFAConfigOpenApiVO() // MFAConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.ModifyGlobalMFAStatus(context.Background(), omadacId).MFAConfigOpenApiVO(mFAConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.ModifyGlobalMFAStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGlobalMFAStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.ModifyGlobalMFAStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyGlobalMFAStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mFAConfigOpenApiVO** | [**MFAConfigOpenApiVO**](MFAConfigOpenApiVO.md) |  | 

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


## ModifyRole

> OperationResponse ModifyRole(ctx, omadacId, roleId).ModifyControllerRoleVO(modifyControllerRoleVO).Execute()

Modify an existing role



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
	roleId := "roleId_example" // string | Role ID
	modifyControllerRoleVO := *openapiclient.NewModifyControllerRoleVO("Name_example") // ModifyControllerRoleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.ModifyRole(context.Background(), omadacId, roleId).ModifyControllerRoleVO(modifyControllerRoleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.ModifyRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRole`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.ModifyRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyControllerRoleVO** | [**ModifyControllerRoleVO**](ModifyControllerRoleVO.md) |  | 

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


## ModifyUser

> OperationResponse ModifyUser(ctx, omadacId, userID).ModifyUserVO(modifyUserVO).Execute()

Modify an existing user



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
	userID := "userID_example" // string | userID
	modifyUserVO := *openapiclient.NewModifyUserVO(false, "Name_example", "RoleId_example") // ModifyUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.ModifyUser(context.Background(), omadacId, userID).ModifyUserVO(modifyUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.ModifyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUser`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.ModifyUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**userID** | **string** | userID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyUserVO** | [**ModifyUserVO**](ModifyUserVO.md) |  | 

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


## ModifyUserBatch

> OperationResponseBatchModifyCountVO ModifyUserBatch(ctx, omadacId).BatchModifyUserOpenApiVO(batchModifyUserOpenApiVO).Execute()

Batch modify users.



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
	batchModifyUserOpenApiVO := *openapiclient.NewBatchModifyUserOpenApiVO() // BatchModifyUserOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.ModifyUserBatch(context.Background(), omadacId).BatchModifyUserOpenApiVO(batchModifyUserOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.ModifyUserBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUserBatch`: OperationResponseBatchModifyCountVO
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.ModifyUserBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUserBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchModifyUserOpenApiVO** | [**BatchModifyUserOpenApiVO**](BatchModifyUserOpenApiVO.md) |  | 

### Return type

[**OperationResponseBatchModifyCountVO**](OperationResponseBatchModifyCountVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reinvite

> OperationResponse Reinvite(ctx, omadacId, userID).Execute()

Re-Invite an existing user



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
	userID := "userID_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.Reinvite(context.Background(), omadacId, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.Reinvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Reinvite`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.Reinvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**userID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReinviteRequest struct via the builder pattern


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


## TransferMspRoot

> OperationResponseWithoutResult TransferMspRoot(ctx, mspId).TransferRootOpenApiVO(transferRootOpenApiVO).Execute()

Transfer MSP owner permission to an existing user



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
	transferRootOpenApiVO := *openapiclient.NewTransferRootOpenApiVO("UserId_example") // TransferRootOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.TransferMspRoot(context.Background(), mspId).TransferRootOpenApiVO(transferRootOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.TransferMspRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferMspRoot`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.TransferMspRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferMspRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferRootOpenApiVO** | [**TransferRootOpenApiVO**](TransferRootOpenApiVO.md) |  | 

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


## TransferRoot

> OperationResponseWithoutResult TransferRoot(ctx, omadacId).TransferRootOpenApiVO(transferRootOpenApiVO).Execute()

Transfer owner permission to an existing user



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
	transferRootOpenApiVO := *openapiclient.NewTransferRootOpenApiVO("UserId_example") // TransferRootOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAndRoleAPI.TransferRoot(context.Background(), omadacId).TransferRootOpenApiVO(transferRootOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAndRoleAPI.TransferRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferRoot`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `UserAndRoleAPI.TransferRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferRootOpenApiVO** | [**TransferRootOpenApiVO**](TransferRootOpenApiVO.md) |  | 

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


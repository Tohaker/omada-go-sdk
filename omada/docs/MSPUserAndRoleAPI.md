# MSPUserAndRoleAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerRole**](MSPUserAndRoleAPI.md#createcustomerrole) | **Post** /openapi/v1/msp/{mspId}/account/customer-roles | Create new customer role
[**CreateNewMspRole**](MSPUserAndRoleAPI.md#createnewmsprole) | **Post** /openapi/v1/msp/{mspId}/roles | Create new msp role
[**CreateNewMspUser**](MSPUserAndRoleAPI.md#createnewmspuser) | **Post** /openapi/v1/msp/{mspId}/users | Create new msp user
[**DeleteCustomerRole**](MSPUserAndRoleAPI.md#deletecustomerrole) | **Delete** /openapi/v1/msp/{mspId}/account/customer-roles/{roleId} | Delete an existing customer role
[**DeleteMspRole**](MSPUserAndRoleAPI.md#deletemsprole) | **Delete** /openapi/v1/msp/{mspId}/roles/{roleId} | Delete an existing msp role
[**DeleteMspUser**](MSPUserAndRoleAPI.md#deletemspuser) | **Delete** /openapi/v1/msp/{mspId}/users/{userID} | Delete an existing msp user
[**GetAllCustomerRoles**](MSPUserAndRoleAPI.md#getallcustomerroles) | **Get** /openapi/v1/msp/{mspId}/customer-roles | Get customer role list in msp
[**GetAllMspCloudUsersExcludeRoot**](MSPUserAndRoleAPI.md#getallmspcloudusersexcluderoot) | **Get** /openapi/v1/msp/{mspId}/users/cloud | Get all MSP cloud users exclude owner
[**GetAllMspLocalUsersExcludeRoot**](MSPUserAndRoleAPI.md#getallmsplocalusersexcluderoot) | **Get** /openapi/v1/msp/{mspId}/users/local | Get all MSP local users exclude owner
[**GetAllMspRoles**](MSPUserAndRoleAPI.md#getallmsproles) | **Get** /openapi/v1/msp/{mspId}/roles | Get msp role list
[**GetAvailableCustomerRoles**](MSPUserAndRoleAPI.md#getavailablecustomerroles) | **Get** /openapi/v1/msp/{mspId}/roles/customer/roles/available | Get available customer roles for creating new user
[**GetAvailableMspRoles**](MSPUserAndRoleAPI.md#getavailablemsproles) | **Get** /openapi/v1/msp/{mspId}/roles/available | Get available MSP roles for creating new user
[**GetGridUsers1**](MSPUserAndRoleAPI.md#getgridusers1) | **Get** /openapi/v1/msp/{mspId}/users | Get msp user list
[**GetMFAStatus**](MSPUserAndRoleAPI.md#getmfastatus) | **Get** /openapi/v1/msp/{mspId}/mfa/status | Get the status of the MSP system&#39;s Two-Factor Authentication.
[**GetMspAppGridUsers**](MSPUserAndRoleAPI.md#getmspappgridusers) | **Get** /openapi/v1/msp/{mspId}/all-users | Get msp user list for app
[**GetMspRole**](MSPUserAndRoleAPI.md#getmsprole) | **Get** /openapi/v1/msp/{mspId}/roles/{roleId} | Get msp role info
[**GetUser1**](MSPUserAndRoleAPI.md#getuser1) | **Get** /openapi/v1/msp/{mspId}/users/{userID} | Get msp user info
[**ModifyCustomerRole**](MSPUserAndRoleAPI.md#modifycustomerrole) | **Put** /openapi/v1/msp/{mspId}/account/customer-roles/{roleId} | Modify an existing customer role
[**ModifyMspMFAStatus**](MSPUserAndRoleAPI.md#modifymspmfastatus) | **Post** /openapi/v1/msp/{mspId}/mfa/status | Modify the status of the MSP system&#39;s Two-Factor Authentication.
[**ModifyMspRole**](MSPUserAndRoleAPI.md#modifymsprole) | **Put** /openapi/v1/msp/{mspId}/roles/{roleId} | Modify an existing msp role
[**ModifyMspUser**](MSPUserAndRoleAPI.md#modifymspuser) | **Put** /openapi/v1/msp/{mspId}/users/{userID} | Modify an existing msp user
[**ModifyMspUserBatch**](MSPUserAndRoleAPI.md#modifymspuserbatch) | **Patch** /openapi/v1/msp/{mspId}/users/batch-temporary-users | Batch modify MSP users.



## CreateCustomerRole

> OperationResponseCreateRoleResultVO CreateCustomerRole(ctx, mspId).RoleDetailOpenApiVO(roleDetailOpenApiVO).Execute()

Create new customer role



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
	roleDetailOpenApiVO := *openapiclient.NewRoleDetailOpenApiVO("Name_example") // RoleDetailOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.CreateCustomerRole(context.Background(), mspId).RoleDetailOpenApiVO(roleDetailOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.CreateCustomerRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerRole`: OperationResponseCreateRoleResultVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.CreateCustomerRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleDetailOpenApiVO** | [**RoleDetailOpenApiVO**](RoleDetailOpenApiVO.md) |  | 

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


## CreateNewMspRole

> OperationResponseCreateRoleResultVO CreateNewMspRole(ctx, mspId).ModifyMspRoleVO(modifyMspRoleVO).Execute()

Create new msp role



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
	modifyMspRoleVO := *openapiclient.NewModifyMspRoleVO("Name_example", *openapiclient.NewMspRoleVO()) // ModifyMspRoleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.CreateNewMspRole(context.Background(), mspId).ModifyMspRoleVO(modifyMspRoleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.CreateNewMspRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewMspRole`: OperationResponseCreateRoleResultVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.CreateNewMspRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewMspRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyMspRoleVO** | [**ModifyMspRoleVO**](ModifyMspRoleVO.md) |  | 

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


## CreateNewMspUser

> OperationResponseCreateUserResultVO CreateNewMspUser(ctx, mspId).CreateMspUserVO(createMspUserVO).Execute()

Create new msp user



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
	createMspUserVO := *openapiclient.NewCreateMspUserVO(false, "CustomerRoleId_example", "Name_example", "RoleId_example", int32(123)) // CreateMspUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.CreateNewMspUser(context.Background(), mspId).CreateMspUserVO(createMspUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.CreateNewMspUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewMspUser`: OperationResponseCreateUserResultVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.CreateNewMspUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewMspUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMspUserVO** | [**CreateMspUserVO**](CreateMspUserVO.md) |  | 

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


## DeleteCustomerRole

> OperationResponse DeleteCustomerRole(ctx, mspId, roleId).Execute()

Delete an existing customer role



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
	roleId := "roleId_example" // string | Role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.DeleteCustomerRole(context.Background(), mspId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.DeleteCustomerRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomerRole`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.DeleteCustomerRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerRoleRequest struct via the builder pattern


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


## DeleteMspRole

> OperationResponse DeleteMspRole(ctx, mspId, roleId).Execute()

Delete an existing msp role



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
	roleId := "roleId_example" // string | Role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.DeleteMspRole(context.Background(), mspId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.DeleteMspRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMspRole`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.DeleteMspRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMspRoleRequest struct via the builder pattern


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


## DeleteMspUser

> OperationResponse DeleteMspUser(ctx, mspId, userID).DeleteUserVO(deleteUserVO).Execute()

Delete an existing msp user



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
	userID := "userID_example" // string | userID
	deleteUserVO := *openapiclient.NewDeleteUserVO() // DeleteUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.DeleteMspUser(context.Background(), mspId, userID).DeleteUserVO(deleteUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.DeleteMspUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMspUser`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.DeleteMspUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**userID** | **string** | userID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMspUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteUserVO** | [**DeleteUserVO**](DeleteUserVO.md) |  | 

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


## GetAllCustomerRoles

> OperationResponseListControllerRoleDetailVO GetAllCustomerRoles(ctx, mspId).Execute()

Get customer role list in msp



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetAllCustomerRoles(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetAllCustomerRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCustomerRoles`: OperationResponseListControllerRoleDetailVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetAllCustomerRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCustomerRolesRequest struct via the builder pattern


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


## GetAllMspCloudUsersExcludeRoot

> OperationResponseGetUserListResponseVO GetAllMspCloudUsersExcludeRoot(ctx, mspId).Execute()

Get all MSP cloud users exclude owner



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetAllMspCloudUsersExcludeRoot(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetAllMspCloudUsersExcludeRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMspCloudUsersExcludeRoot`: OperationResponseGetUserListResponseVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetAllMspCloudUsersExcludeRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMspCloudUsersExcludeRootRequest struct via the builder pattern


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


## GetAllMspLocalUsersExcludeRoot

> OperationResponseGetUserListResponseVO GetAllMspLocalUsersExcludeRoot(ctx, mspId).Execute()

Get all MSP local users exclude owner



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetAllMspLocalUsersExcludeRoot(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetAllMspLocalUsersExcludeRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMspLocalUsersExcludeRoot`: OperationResponseGetUserListResponseVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetAllMspLocalUsersExcludeRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMspLocalUsersExcludeRootRequest struct via the builder pattern


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


## GetAllMspRoles

> OperationResponseListMspRoleDetailVO GetAllMspRoles(ctx, mspId).Execute()

Get msp role list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetAllMspRoles(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetAllMspRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMspRoles`: OperationResponseListMspRoleDetailVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetAllMspRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMspRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseListMspRoleDetailVO**](OperationResponseListMspRoleDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableCustomerRoles

> OperationResponseListRoleBriefVO GetAvailableCustomerRoles(ctx, mspId).Execute()

Get available customer roles for creating new user



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetAvailableCustomerRoles(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetAvailableCustomerRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableCustomerRoles`: OperationResponseListRoleBriefVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetAvailableCustomerRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableCustomerRolesRequest struct via the builder pattern


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


## GetAvailableMspRoles

> OperationResponseListRoleBriefVO GetAvailableMspRoles(ctx, mspId).Execute()

Get available MSP roles for creating new user



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetAvailableMspRoles(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetAvailableMspRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableMspRoles`: OperationResponseListRoleBriefVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetAvailableMspRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableMspRolesRequest struct via the builder pattern


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


## GetGridUsers1

> OperationResponseGridVOMspUserDetailVO GetGridUsers1(ctx, mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsRoleId(sortsRoleId).SortsEmail(sortsEmail).SortsCustomerRoleId(sortsCustomerRoleId).SearchKey(searchKey).Execute()

Get msp user list



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
	sortsRoleId := "sortsRoleId_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsEmail := "sortsEmail_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsCustomerRoleId := "sortsCustomerRoleId_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field userName (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetGridUsers1(context.Background(), mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsRoleId(sortsRoleId).SortsEmail(sortsEmail).SortsCustomerRoleId(sortsCustomerRoleId).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetGridUsers1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridUsers1`: OperationResponseGridVOMspUserDetailVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetGridUsers1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridUsers1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsRoleId** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsEmail** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsCustomerRoleId** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field userName | 

### Return type

[**OperationResponseGridVOMspUserDetailVO**](OperationResponseGridVOMspUserDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMFAStatus

> OperationResponseMFAConfigOpenApiVO GetMFAStatus(ctx, mspId).Execute()

Get the status of the MSP system's Two-Factor Authentication.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetMFAStatus(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetMFAStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMFAStatus`: OperationResponseMFAConfigOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetMFAStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMFAStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseMFAConfigOpenApiVO**](OperationResponseMFAConfigOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspAppGridUsers

> OperationResponseGridVOMspUserVO GetMspAppGridUsers(ctx, mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsRoleId(sortsRoleId).SortsEmail(sortsEmail).SortsCustomerRoleId(sortsCustomerRoleId).SearchKey(searchKey).Execute()

Get msp user list for app



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
	sortsRoleId := "sortsRoleId_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsEmail := "sortsEmail_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsCustomerRoleId := "sortsCustomerRoleId_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field userName (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetMspAppGridUsers(context.Background(), mspId).Page(page).PageSize(pageSize).SortsName(sortsName).SortsRoleId(sortsRoleId).SortsEmail(sortsEmail).SortsCustomerRoleId(sortsCustomerRoleId).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetMspAppGridUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspAppGridUsers`: OperationResponseGridVOMspUserVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetMspAppGridUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspAppGridUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsRoleId** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsEmail** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsCustomerRoleId** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field userName | 

### Return type

[**OperationResponseGridVOMspUserVO**](OperationResponseGridVOMspUserVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspRole

> OperationResponseMspRoleDetailVO GetMspRole(ctx, mspId, roleId).Execute()

Get msp role info



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
	mspId := "mspId_example" // string | 
	roleId := "roleId_example" // string | Role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetMspRole(context.Background(), mspId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetMspRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspRole`: OperationResponseMspRoleDetailVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetMspRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMspRoleDetailVO**](OperationResponseMspRoleDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser1

> OperationResponseMspUserDetailVO GetUser1(ctx, mspId, userID).Execute()

Get msp user info



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
	userID := "userID_example" // string | userID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.GetUser1(context.Background(), mspId, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.GetUser1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser1`: OperationResponseMspUserDetailVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.GetUser1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**userID** | **string** | userID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMspUserDetailVO**](OperationResponseMspUserDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCustomerRole

> OperationResponse ModifyCustomerRole(ctx, mspId, roleId).RoleDetailOpenApiVO(roleDetailOpenApiVO).Execute()

Modify an existing customer role



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
	roleId := "roleId_example" // string | Role ID
	roleDetailOpenApiVO := *openapiclient.NewRoleDetailOpenApiVO("Name_example") // RoleDetailOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.ModifyCustomerRole(context.Background(), mspId, roleId).RoleDetailOpenApiVO(roleDetailOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.ModifyCustomerRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCustomerRole`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.ModifyCustomerRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCustomerRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleDetailOpenApiVO** | [**RoleDetailOpenApiVO**](RoleDetailOpenApiVO.md) |  | 

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


## ModifyMspMFAStatus

> OperationResponse ModifyMspMFAStatus(ctx, mspId).MFAConfigOpenApiVO(mFAConfigOpenApiVO).Execute()

Modify the status of the MSP system's Two-Factor Authentication.



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
	mFAConfigOpenApiVO := *openapiclient.NewMFAConfigOpenApiVO() // MFAConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.ModifyMspMFAStatus(context.Background(), mspId).MFAConfigOpenApiVO(mFAConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.ModifyMspMFAStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspMFAStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.ModifyMspMFAStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspMFAStatusRequest struct via the builder pattern


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


## ModifyMspRole

> OperationResponse ModifyMspRole(ctx, mspId, roleId).ModifyRoleVO(modifyRoleVO).Execute()

Modify an existing msp role



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
	roleId := "roleId_example" // string | Role ID
	modifyRoleVO := *openapiclient.NewModifyRoleVO("Name_example") // ModifyRoleVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.ModifyMspRole(context.Background(), mspId, roleId).ModifyRoleVO(modifyRoleVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.ModifyMspRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspRole`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.ModifyMspRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyRoleVO** | [**ModifyRoleVO**](ModifyRoleVO.md) |  | 

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


## ModifyMspUser

> OperationResponse ModifyMspUser(ctx, mspId, userID).ModifyMspUserVO(modifyMspUserVO).Execute()

Modify an existing msp user



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
	userID := "userID_example" // string | userID
	modifyMspUserVO := *openapiclient.NewModifyMspUserVO(false, "CustomerRoleId_example", "Name_example", "RoleId_example") // ModifyMspUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.ModifyMspUser(context.Background(), mspId, userID).ModifyMspUserVO(modifyMspUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.ModifyMspUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspUser`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.ModifyMspUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**userID** | **string** | userID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyMspUserVO** | [**ModifyMspUserVO**](ModifyMspUserVO.md) |  | 

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


## ModifyMspUserBatch

> OperationResponseBatchModifyCountVO ModifyMspUserBatch(ctx, mspId).BatchModifyUserOpenApiVO(batchModifyUserOpenApiVO).Execute()

Batch modify MSP users.



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
	batchModifyUserOpenApiVO := *openapiclient.NewBatchModifyUserOpenApiVO() // BatchModifyUserOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPUserAndRoleAPI.ModifyMspUserBatch(context.Background(), mspId).BatchModifyUserOpenApiVO(batchModifyUserOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPUserAndRoleAPI.ModifyMspUserBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspUserBatch`: OperationResponseBatchModifyCountVO
	fmt.Fprintf(os.Stdout, "Response from `MSPUserAndRoleAPI.ModifyMspUserBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspUserBatchRequest struct via the builder pattern


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


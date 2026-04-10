# SSLVPNAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteLockedSslVpnTunnel**](SSLVPNAPI.md#batchdeletelockedsslvpntunnel) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-tunnels | Batch delete SSL VPN locked tunnel
[**BatchDeleteSslVpnUserGroup**](SSLVPNAPI.md#batchdeletesslvpnusergroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/user-groups | Batch delete SSL VPN user group
[**CreateLockedSslVpnTunnuel**](SSLVPNAPI.md#createlockedsslvpntunnuel) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-tunnels | Create locked SSL VPN tunnel
[**CreateLockedSslVpnUser**](SSLVPNAPI.md#createlockedsslvpnuser) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-users | Create SSL VPN locked user
[**CreateSslVpnResource**](SSLVPNAPI.md#createsslvpnresource) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/resources | Create SSL VPN resource
[**CreateSslVpnResourceGroup**](SSLVPNAPI.md#createsslvpnresourcegroup) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/resource-groups | Create SSL VPN resource group
[**CreateSslVpnUser**](SSLVPNAPI.md#createsslvpnuser) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/users | Create SSL VPN user
[**CreateSslVpnUserGroup**](SSLVPNAPI.md#createsslvpnusergroup) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/user-groups | Create SSL VPN user group
[**DeleteLockedSslVpnTunnel**](SSLVPNAPI.md#deletelockedsslvpntunnel) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-tunnels/{lockTunnelId} | Delete SSL VPN locked tunnel
[**DeleteLockedSslVpnUser**](SSLVPNAPI.md#deletelockedsslvpnuser) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-users/{id} | Delete SSL VPN locked user
[**DeleteSslVpnResource**](SSLVPNAPI.md#deletesslvpnresource) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/resources/{id} | Delete SSL VPN resource
[**DeleteSslVpnResourceGroup**](SSLVPNAPI.md#deletesslvpnresourcegroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/resource-groups/{id} | Delete SSL VPN resource group
[**DeleteSslVpnUser**](SSLVPNAPI.md#deletesslvpnuser) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/users/{id} | Delete SSL VPN user
[**DeleteSslVpnUserGroup**](SSLVPNAPI.md#deletesslvpnusergroup) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/user-groups/{id} | Delete SSL VPN user group
[**DownloadSslVpnCertificate**](SSLVPNAPI.md#downloadsslvpncertificate) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/certificate | Download SSL VPN certificate.
[**GetGridLockedSslVpnServerUser**](SSLVPNAPI.md#getgridlockedsslvpnserveruser) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-users | Get locked user list for SSL VPN server
[**GetGridSslVpnServerResource**](SSLVPNAPI.md#getgridsslvpnserverresource) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/resources | Get resource list for SSL VPN server
[**GetGridSslVpnServerResourceGroup**](SSLVPNAPI.md#getgridsslvpnserverresourcegroup) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/resource-groups | Get resource group list for SSL VPN server
[**GetGridSslVpnServerUser**](SSLVPNAPI.md#getgridsslvpnserveruser) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/users | Get user list for SSL VPN server
[**GetGridSslVpnServerUserGroup**](SSLVPNAPI.md#getgridsslvpnserverusergroup) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/user-groups | Get user group list for SSL VPN server
[**GetGridSslVpnUserInGroup**](SSLVPNAPI.md#getgridsslvpnuseringroup) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/usergroups/{userGroupId}/users | Get SSL VPN user list in group.
[**GetGridSslVpnUserInGroupV2**](SSLVPNAPI.md#getgridsslvpnuseringroupv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/vpn/usergroups/{userGroupId}/users | Get SSL VPN user list in group V2.
[**GetLockedSslVpnTunnuels**](SSLVPNAPI.md#getlockedsslvpntunnuels) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-tunnels | Get locked SSL VPN tunnel.
[**GetRadiusServerInfo**](SSLVPNAPI.md#getradiusserverinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/radius | Get radius server for SSL VPN server
[**GetSslVpnResourceGroupList**](SSLVPNAPI.md#getsslvpnresourcegrouplist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/briefresourcegroups | Get SSL VPN resource group list.
[**GetSslVpnResourceList**](SSLVPNAPI.md#getsslvpnresourcelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/briefresources | Get SSL VPN resource list
[**GetSslVpnServerSetting**](SSLVPNAPI.md#getsslvpnserversetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/setting | Get SSL VPN server setting
[**GetSslVpnUserGroupList**](SSLVPNAPI.md#getsslvpnusergrouplist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/briefusergroups | Get SSL VPN user Group list.
[**ModifyLockedSslVpnTunnuel**](SSLVPNAPI.md#modifylockedsslvpntunnuel) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-tunnels/{lockTunnelId} | Modify locked SSL VPN tunnel
[**ModifyLockedSslVpnUser**](SSLVPNAPI.md#modifylockedsslvpnuser) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/locked-users/{id} | Modify SSL VPN locked user
[**ModifySslVpnResource**](SSLVPNAPI.md#modifysslvpnresource) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/resources/{id} | Modify SSL VPN resource
[**ModifySslVpnResourceGroup**](SSLVPNAPI.md#modifysslvpnresourcegroup) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/resource-groups/{id} | Modify SSL VPN resource group
[**ModifySslVpnServerSetting**](SSLVPNAPI.md#modifysslvpnserversetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/setting | Modify SSL VPN server setting
[**ModifySslVpnUser**](SSLVPNAPI.md#modifysslvpnuser) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/users/{id} | Modify SSL VPN user
[**ModifySslVpnUserGroup**](SSLVPNAPI.md#modifysslvpnusergroup) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ssl-vpn-server/user-groups/{id} | Modify SSL VPN user group



## BatchDeleteLockedSslVpnTunnel

> OperationResponseWithoutResult BatchDeleteLockedSslVpnTunnel(ctx, omadacId, siteId).BatchSelectSslUserVO(batchSelectSslUserVO).Execute()

Batch delete SSL VPN locked tunnel



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
	batchSelectSslUserVO := *openapiclient.NewBatchSelectSslUserVO() // BatchSelectSslUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.BatchDeleteLockedSslVpnTunnel(context.Background(), omadacId, siteId).BatchSelectSslUserVO(batchSelectSslUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.BatchDeleteLockedSslVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteLockedSslVpnTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.BatchDeleteLockedSslVpnTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteLockedSslVpnTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchSelectSslUserVO** | [**BatchSelectSslUserVO**](BatchSelectSslUserVO.md) |  | 

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


## BatchDeleteSslVpnUserGroup

> OperationResponseWithoutResult BatchDeleteSslVpnUserGroup(ctx, omadacId, siteId).BatchSelectSslUserVO(batchSelectSslUserVO).Execute()

Batch delete SSL VPN user group



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
	batchSelectSslUserVO := *openapiclient.NewBatchSelectSslUserVO() // BatchSelectSslUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.BatchDeleteSslVpnUserGroup(context.Background(), omadacId, siteId).BatchSelectSslUserVO(batchSelectSslUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.BatchDeleteSslVpnUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteSslVpnUserGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.BatchDeleteSslVpnUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteSslVpnUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchSelectSslUserVO** | [**BatchSelectSslUserVO**](BatchSelectSslUserVO.md) |  | 

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


## CreateLockedSslVpnTunnuel

> OperationResponseWithoutResult CreateLockedSslVpnTunnuel(ctx, omadacId, siteId).SslVpnLockCreateAndModifyOpenApiVO(sslVpnLockCreateAndModifyOpenApiVO).Execute()

Create locked SSL VPN tunnel



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
	sslVpnLockCreateAndModifyOpenApiVO := *openapiclient.NewSslVpnLockCreateAndModifyOpenApiVO(int32(123), int32(123)) // SslVpnLockCreateAndModifyOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.CreateLockedSslVpnTunnuel(context.Background(), omadacId, siteId).SslVpnLockCreateAndModifyOpenApiVO(sslVpnLockCreateAndModifyOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.CreateLockedSslVpnTunnuel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLockedSslVpnTunnuel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.CreateLockedSslVpnTunnuel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLockedSslVpnTunnuelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sslVpnLockCreateAndModifyOpenApiVO** | [**SslVpnLockCreateAndModifyOpenApiVO**](SslVpnLockCreateAndModifyOpenApiVO.md) |  | 

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


## CreateLockedSslVpnUser

> OperationResponseWithoutResult CreateLockedSslVpnUser(ctx, omadacId, siteId).SslVpnLockConfigOpenApiVO(sslVpnLockConfigOpenApiVO).Execute()

Create SSL VPN locked user



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
	sslVpnLockConfigOpenApiVO := *openapiclient.NewSslVpnLockConfigOpenApiVO(int32(123), int32(123)) // SslVpnLockConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.CreateLockedSslVpnUser(context.Background(), omadacId, siteId).SslVpnLockConfigOpenApiVO(sslVpnLockConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.CreateLockedSslVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLockedSslVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.CreateLockedSslVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLockedSslVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sslVpnLockConfigOpenApiVO** | [**SslVpnLockConfigOpenApiVO**](SslVpnLockConfigOpenApiVO.md) |  | 

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


## CreateSslVpnResource

> OperationResponseWithoutResult CreateSslVpnResource(ctx, omadacId, siteId).SslVpnResourceConfigOpenApiVO(sslVpnResourceConfigOpenApiVO).Execute()

Create SSL VPN resource



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
	sslVpnResourceConfigOpenApiVO := *openapiclient.NewSslVpnResourceConfigOpenApiVO("Name_example", int32(123)) // SslVpnResourceConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.CreateSslVpnResource(context.Background(), omadacId, siteId).SslVpnResourceConfigOpenApiVO(sslVpnResourceConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.CreateSslVpnResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSslVpnResource`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.CreateSslVpnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSslVpnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sslVpnResourceConfigOpenApiVO** | [**SslVpnResourceConfigOpenApiVO**](SslVpnResourceConfigOpenApiVO.md) |  | 

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


## CreateSslVpnResourceGroup

> OperationResponseWithoutResult CreateSslVpnResourceGroup(ctx, omadacId, siteId).SslVpnResourceGroupConfigOpenApiVO(sslVpnResourceGroupConfigOpenApiVO).Execute()

Create SSL VPN resource group



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
	sslVpnResourceGroupConfigOpenApiVO := *openapiclient.NewSslVpnResourceGroupConfigOpenApiVO("Name_example") // SslVpnResourceGroupConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.CreateSslVpnResourceGroup(context.Background(), omadacId, siteId).SslVpnResourceGroupConfigOpenApiVO(sslVpnResourceGroupConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.CreateSslVpnResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSslVpnResourceGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.CreateSslVpnResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSslVpnResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sslVpnResourceGroupConfigOpenApiVO** | [**SslVpnResourceGroupConfigOpenApiVO**](SslVpnResourceGroupConfigOpenApiVO.md) |  | 

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


## CreateSslVpnUser

> OperationResponseWithoutResult CreateSslVpnUser(ctx, omadacId, siteId).SslVpnUserConfigOpenApiVO(sslVpnUserConfigOpenApiVO).Execute()

Create SSL VPN user



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
	sslVpnUserConfigOpenApiVO := *openapiclient.NewSslVpnUserConfigOpenApiVO(int32(123), "Name_example", "Password_example", false, "Validity_example") // SslVpnUserConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.CreateSslVpnUser(context.Background(), omadacId, siteId).SslVpnUserConfigOpenApiVO(sslVpnUserConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.CreateSslVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSslVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.CreateSslVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSslVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sslVpnUserConfigOpenApiVO** | [**SslVpnUserConfigOpenApiVO**](SslVpnUserConfigOpenApiVO.md) |  | 

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


## CreateSslVpnUserGroup

> OperationResponseWithoutResult CreateSslVpnUserGroup(ctx, omadacId, siteId).SslVpnUserGroupConfigOpenApiVO(sslVpnUserGroupConfigOpenApiVO).Execute()

Create SSL VPN user group



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
	sslVpnUserGroupConfigOpenApiVO := *openapiclient.NewSslVpnUserGroupConfigOpenApiVO("Name_example") // SslVpnUserGroupConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.CreateSslVpnUserGroup(context.Background(), omadacId, siteId).SslVpnUserGroupConfigOpenApiVO(sslVpnUserGroupConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.CreateSslVpnUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSslVpnUserGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.CreateSslVpnUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSslVpnUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sslVpnUserGroupConfigOpenApiVO** | [**SslVpnUserGroupConfigOpenApiVO**](SslVpnUserGroupConfigOpenApiVO.md) |  | 

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


## DeleteLockedSslVpnTunnel

> OperationResponseWithoutResult DeleteLockedSslVpnTunnel(ctx, omadacId, siteId, lockTunnelId).Execute()

Delete SSL VPN locked tunnel



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
	lockTunnelId := "lockTunnelId_example" // string | lockTunnelId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.DeleteLockedSslVpnTunnel(context.Background(), omadacId, siteId, lockTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.DeleteLockedSslVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLockedSslVpnTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.DeleteLockedSslVpnTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**lockTunnelId** | **string** | lockTunnelId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLockedSslVpnTunnelRequest struct via the builder pattern


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


## DeleteLockedSslVpnUser

> OperationResponseWithoutResult DeleteLockedSslVpnUser(ctx, omadacId, siteId, id).Execute()

Delete SSL VPN locked user



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
	id := "id_example" // string | Locked user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.DeleteLockedSslVpnUser(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.DeleteLockedSslVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLockedSslVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.DeleteLockedSslVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Locked user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLockedSslVpnUserRequest struct via the builder pattern


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


## DeleteSslVpnResource

> OperationResponseWithoutResult DeleteSslVpnResource(ctx, omadacId, siteId, id).Execute()

Delete SSL VPN resource



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
	id := "id_example" // string | Resource ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.DeleteSslVpnResource(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.DeleteSslVpnResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSslVpnResource`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.DeleteSslVpnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSslVpnResourceRequest struct via the builder pattern


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


## DeleteSslVpnResourceGroup

> OperationResponseWithoutResult DeleteSslVpnResourceGroup(ctx, omadacId, siteId, id).Execute()

Delete SSL VPN resource group



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
	id := "id_example" // string | Resource group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.DeleteSslVpnResourceGroup(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.DeleteSslVpnResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSslVpnResourceGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.DeleteSslVpnResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Resource group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSslVpnResourceGroupRequest struct via the builder pattern


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


## DeleteSslVpnUser

> OperationResponseWithoutResult DeleteSslVpnUser(ctx, omadacId, siteId, id).Execute()

Delete SSL VPN user



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
	id := "id_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.DeleteSslVpnUser(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.DeleteSslVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSslVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.DeleteSslVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSslVpnUserRequest struct via the builder pattern


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


## DeleteSslVpnUserGroup

> OperationResponseWithoutResult DeleteSslVpnUserGroup(ctx, omadacId, siteId, id).Execute()

Delete SSL VPN user group



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
	id := "id_example" // string | User group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.DeleteSslVpnUserGroup(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.DeleteSslVpnUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSslVpnUserGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.DeleteSslVpnUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | User group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSslVpnUserGroupRequest struct via the builder pattern


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


## DownloadSslVpnCertificate

> OperationResponse DownloadSslVpnCertificate(ctx, omadacId, siteId).WanIp(wanIp).Execute()

Download SSL VPN certificate.



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
	wanIp := "wanIp_example" // string | WLAN IP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.DownloadSslVpnCertificate(context.Background(), omadacId, siteId).WanIp(wanIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.DownloadSslVpnCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadSslVpnCertificate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.DownloadSslVpnCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSslVpnCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wanIp** | **string** | WLAN IP | 

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


## GetGridLockedSslVpnServerUser

> OperationResponseGridVOSSLVPNLockEntity GetGridLockedSslVpnServerUser(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).SortsUserName(sortsUserName).SortsIp(sortsIp).SortsLeftLockTime(sortsLeftLockTime).Execute()

Get locked user list for SSL VPN server



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name. (optional)
	sortsUserName := "sortsUserName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsLeftLockTime := "sortsLeftLockTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.GetGridLockedSslVpnServerUser(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).SortsUserName(sortsUserName).SortsIp(sortsIp).SortsLeftLockTime(sortsLeftLockTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetGridLockedSslVpnServerUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridLockedSslVpnServerUser`: OperationResponseGridVOSSLVPNLockEntity
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetGridLockedSslVpnServerUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridLockedSslVpnServerUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name. | 
 **sortsUserName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsLeftLockTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseGridVOSSLVPNLockEntity**](OperationResponseGridVOSSLVPNLockEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSslVpnServerResource

> OperationResponseGridVOSslVpnResourceEntity GetGridSslVpnServerResource(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get resource list for SSL VPN server



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.GetGridSslVpnServerResource(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetGridSslVpnServerResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSslVpnServerResource`: OperationResponseGridVOSslVpnResourceEntity
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetGridSslVpnServerResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSslVpnServerResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name. | 

### Return type

[**OperationResponseGridVOSslVpnResourceEntity**](OperationResponseGridVOSslVpnResourceEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSslVpnServerResourceGroup

> OperationResponseGridVOSslVpnResourceGroup GetGridSslVpnServerResourceGroup(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get resource group list for SSL VPN server



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
	resp, r, err := apiClient.SSLVPNAPI.GetGridSslVpnServerResourceGroup(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetGridSslVpnServerResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSslVpnServerResourceGroup`: OperationResponseGridVOSslVpnResourceGroup
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetGridSslVpnServerResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSslVpnServerResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOSslVpnResourceGroup**](OperationResponseGridVOSslVpnResourceGroup.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSslVpnServerUser

> OperationResponseSslVpnUserOpenApiGridVOSslVpnUserEntity GetGridSslVpnServerUser(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get user list for SSL VPN server



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
	resp, r, err := apiClient.SSLVPNAPI.GetGridSslVpnServerUser(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetGridSslVpnServerUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSslVpnServerUser`: OperationResponseSslVpnUserOpenApiGridVOSslVpnUserEntity
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetGridSslVpnServerUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSslVpnServerUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseSslVpnUserOpenApiGridVOSslVpnUserEntity**](OperationResponseSslVpnUserOpenApiGridVOSslVpnUserEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSslVpnServerUserGroup

> OperationResponseSslVpnUserGroupGridVOSslVpnUserGroupEntity GetGridSslVpnServerUserGroup(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get user group list for SSL VPN server



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.GetGridSslVpnServerUserGroup(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetGridSslVpnServerUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSslVpnServerUserGroup`: OperationResponseSslVpnUserGroupGridVOSslVpnUserGroupEntity
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetGridSslVpnServerUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSslVpnServerUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name. | 

### Return type

[**OperationResponseSslVpnUserGroupGridVOSslVpnUserGroupEntity**](OperationResponseSslVpnUserGroupGridVOSslVpnUserGroupEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSslVpnUserInGroup

> OperationResponseSslVpnUserOpenApiGridVOSslVpnUserEntity GetGridSslVpnUserInGroup(ctx, omadacId, siteId, userGroupId).Page(page).PageSize(pageSize).Execute()

Get SSL VPN user list in group.



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
	userGroupId := "userGroupId_example" // string | User Group Id
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.GetGridSslVpnUserInGroup(context.Background(), omadacId, siteId, userGroupId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetGridSslVpnUserInGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSslVpnUserInGroup`: OperationResponseSslVpnUserOpenApiGridVOSslVpnUserEntity
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetGridSslVpnUserInGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**userGroupId** | **string** | User Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSslVpnUserInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseSslVpnUserOpenApiGridVOSslVpnUserEntity**](OperationResponseSslVpnUserOpenApiGridVOSslVpnUserEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSslVpnUserInGroupV2

> OperationResponseSslVpnUserGridVOVpnUserInfoVO GetGridSslVpnUserInGroupV2(ctx, omadacId, siteId, userGroupId).Page(page).PageSize(pageSize).SearchKey(searchKey).SortsUsername(sortsUsername).Execute()

Get SSL VPN user list in group V2.



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
	userGroupId := "userGroupId_example" // string | User Group Id
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field user name. (optional)
	sortsUsername := "sortsUsername_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effectuser name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.GetGridSslVpnUserInGroupV2(context.Background(), omadacId, siteId, userGroupId).Page(page).PageSize(pageSize).SearchKey(searchKey).SortsUsername(sortsUsername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetGridSslVpnUserInGroupV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSslVpnUserInGroupV2`: OperationResponseSslVpnUserGridVOVpnUserInfoVO
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetGridSslVpnUserInGroupV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**userGroupId** | **string** | User Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSslVpnUserInGroupV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field user name. | 
 **sortsUsername** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effectuser name. | 

### Return type

[**OperationResponseSslVpnUserGridVOVpnUserInfoVO**](OperationResponseSslVpnUserGridVOVpnUserInfoVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLockedSslVpnTunnuels

> OperationResponseGridVOSSLVPNLockEntity GetLockedSslVpnTunnuels(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get locked SSL VPN tunnel.



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
	resp, r, err := apiClient.SSLVPNAPI.GetLockedSslVpnTunnuels(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetLockedSslVpnTunnuels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLockedSslVpnTunnuels`: OperationResponseGridVOSSLVPNLockEntity
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetLockedSslVpnTunnuels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLockedSslVpnTunnuelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOSSLVPNLockEntity**](OperationResponseGridVOSSLVPNLockEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadiusServerInfo

> OperationResponseListRadiusServerBriefInfo GetRadiusServerInfo(ctx, omadacId, siteId).Execute()

Get radius server for SSL VPN server



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
	resp, r, err := apiClient.SSLVPNAPI.GetRadiusServerInfo(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetRadiusServerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadiusServerInfo`: OperationResponseListRadiusServerBriefInfo
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetRadiusServerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadiusServerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRadiusServerBriefInfo**](OperationResponseListRadiusServerBriefInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslVpnResourceGroupList

> OperationResponseListSslVpnResourceGroupBriefInfo GetSslVpnResourceGroupList(ctx, omadacId, siteId).Execute()

Get SSL VPN resource group list.



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
	resp, r, err := apiClient.SSLVPNAPI.GetSslVpnResourceGroupList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetSslVpnResourceGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSslVpnResourceGroupList`: OperationResponseListSslVpnResourceGroupBriefInfo
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetSslVpnResourceGroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslVpnResourceGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListSslVpnResourceGroupBriefInfo**](OperationResponseListSslVpnResourceGroupBriefInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslVpnResourceList

> OperationResponseListSslVpnResourceBriefInfo GetSslVpnResourceList(ctx, omadacId, siteId).Execute()

Get SSL VPN resource list



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
	resp, r, err := apiClient.SSLVPNAPI.GetSslVpnResourceList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetSslVpnResourceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSslVpnResourceList`: OperationResponseListSslVpnResourceBriefInfo
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetSslVpnResourceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslVpnResourceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListSslVpnResourceBriefInfo**](OperationResponseListSslVpnResourceBriefInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslVpnServerSetting

> OperationResponseSslVpnServerSetting GetSslVpnServerSetting(ctx, omadacId, siteId).Execute()

Get SSL VPN server setting



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
	resp, r, err := apiClient.SSLVPNAPI.GetSslVpnServerSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetSslVpnServerSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSslVpnServerSetting`: OperationResponseSslVpnServerSetting
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetSslVpnServerSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslVpnServerSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSslVpnServerSetting**](OperationResponseSslVpnServerSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslVpnUserGroupList

> OperationResponseListSslVpnUserGroupBriefVO GetSslVpnUserGroupList(ctx, omadacId, siteId).Execute()

Get SSL VPN user Group list.



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
	resp, r, err := apiClient.SSLVPNAPI.GetSslVpnUserGroupList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.GetSslVpnUserGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSslVpnUserGroupList`: OperationResponseListSslVpnUserGroupBriefVO
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.GetSslVpnUserGroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslVpnUserGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListSslVpnUserGroupBriefVO**](OperationResponseListSslVpnUserGroupBriefVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyLockedSslVpnTunnuel

> OperationResponseWithoutResult ModifyLockedSslVpnTunnuel(ctx, omadacId, siteId, lockTunnelId).SslVpnLockCreateAndModifyOpenApiVO(sslVpnLockCreateAndModifyOpenApiVO).Execute()

Modify locked SSL VPN tunnel



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
	lockTunnelId := "lockTunnelId_example" // string | lock Tunnel Id
	sslVpnLockCreateAndModifyOpenApiVO := *openapiclient.NewSslVpnLockCreateAndModifyOpenApiVO(int32(123), int32(123)) // SslVpnLockCreateAndModifyOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.ModifyLockedSslVpnTunnuel(context.Background(), omadacId, siteId, lockTunnelId).SslVpnLockCreateAndModifyOpenApiVO(sslVpnLockCreateAndModifyOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.ModifyLockedSslVpnTunnuel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLockedSslVpnTunnuel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.ModifyLockedSslVpnTunnuel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**lockTunnelId** | **string** | lock Tunnel Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLockedSslVpnTunnuelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sslVpnLockCreateAndModifyOpenApiVO** | [**SslVpnLockCreateAndModifyOpenApiVO**](SslVpnLockCreateAndModifyOpenApiVO.md) |  | 

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


## ModifyLockedSslVpnUser

> OperationResponseWithoutResult ModifyLockedSslVpnUser(ctx, omadacId, siteId, id).SslVpnLockModifyOpenApiVO(sslVpnLockModifyOpenApiVO).Execute()

Modify SSL VPN locked user



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
	id := "id_example" // string | Locked user ID
	sslVpnLockModifyOpenApiVO := *openapiclient.NewSslVpnLockModifyOpenApiVO(int32(123), int32(123)) // SslVpnLockModifyOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.ModifyLockedSslVpnUser(context.Background(), omadacId, siteId, id).SslVpnLockModifyOpenApiVO(sslVpnLockModifyOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.ModifyLockedSslVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLockedSslVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.ModifyLockedSslVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Locked user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLockedSslVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sslVpnLockModifyOpenApiVO** | [**SslVpnLockModifyOpenApiVO**](SslVpnLockModifyOpenApiVO.md) |  | 

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


## ModifySslVpnResource

> OperationResponseWithoutResult ModifySslVpnResource(ctx, omadacId, siteId, id).SslVpnResourceModifyOpenApiVO(sslVpnResourceModifyOpenApiVO).Execute()

Modify SSL VPN resource



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
	id := "id_example" // string | Resource ID
	sslVpnResourceModifyOpenApiVO := *openapiclient.NewSslVpnResourceModifyOpenApiVO(int32(123)) // SslVpnResourceModifyOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.ModifySslVpnResource(context.Background(), omadacId, siteId, id).SslVpnResourceModifyOpenApiVO(sslVpnResourceModifyOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.ModifySslVpnResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySslVpnResource`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.ModifySslVpnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySslVpnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sslVpnResourceModifyOpenApiVO** | [**SslVpnResourceModifyOpenApiVO**](SslVpnResourceModifyOpenApiVO.md) |  | 

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


## ModifySslVpnResourceGroup

> OperationResponseWithoutResult ModifySslVpnResourceGroup(ctx, omadacId, siteId, id).SslVpnResourceGroupModifyOpenApiVO(sslVpnResourceGroupModifyOpenApiVO).Execute()

Modify SSL VPN resource group



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
	id := "id_example" // string | Resource group ID
	sslVpnResourceGroupModifyOpenApiVO := *openapiclient.NewSslVpnResourceGroupModifyOpenApiVO() // SslVpnResourceGroupModifyOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.ModifySslVpnResourceGroup(context.Background(), omadacId, siteId, id).SslVpnResourceGroupModifyOpenApiVO(sslVpnResourceGroupModifyOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.ModifySslVpnResourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySslVpnResourceGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.ModifySslVpnResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Resource group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySslVpnResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sslVpnResourceGroupModifyOpenApiVO** | [**SslVpnResourceGroupModifyOpenApiVO**](SslVpnResourceGroupModifyOpenApiVO.md) |  | 

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


## ModifySslVpnServerSetting

> OperationResponseWithoutResult ModifySslVpnServerSetting(ctx, omadacId, siteId).SslVpnServerConfigOpenApiVO(sslVpnServerConfigOpenApiVO).Execute()

Modify SSL VPN server setting



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
	sslVpnServerConfigOpenApiVO := *openapiclient.NewSslVpnServerConfigOpenApiVO(false) // SslVpnServerConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.ModifySslVpnServerSetting(context.Background(), omadacId, siteId).SslVpnServerConfigOpenApiVO(sslVpnServerConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.ModifySslVpnServerSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySslVpnServerSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.ModifySslVpnServerSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySslVpnServerSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sslVpnServerConfigOpenApiVO** | [**SslVpnServerConfigOpenApiVO**](SslVpnServerConfigOpenApiVO.md) |  | 

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


## ModifySslVpnUser

> OperationResponseWithoutResult ModifySslVpnUser(ctx, omadacId, siteId, id).SslVpnUserModifyOpenApiVO(sslVpnUserModifyOpenApiVO).Execute()

Modify SSL VPN user



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
	id := "id_example" // string | User ID
	sslVpnUserModifyOpenApiVO := *openapiclient.NewSslVpnUserModifyOpenApiVO(int32(123), "Password_example", false, "Validity_example") // SslVpnUserModifyOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.ModifySslVpnUser(context.Background(), omadacId, siteId, id).SslVpnUserModifyOpenApiVO(sslVpnUserModifyOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.ModifySslVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySslVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.ModifySslVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySslVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sslVpnUserModifyOpenApiVO** | [**SslVpnUserModifyOpenApiVO**](SslVpnUserModifyOpenApiVO.md) |  | 

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


## ModifySslVpnUserGroup

> OperationResponseWithoutResult ModifySslVpnUserGroup(ctx, omadacId, siteId, id).SslVpnUserGroupModifyOpenApiVO(sslVpnUserGroupModifyOpenApiVO).Execute()

Modify SSL VPN user group



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
	id := "id_example" // string | User group ID
	sslVpnUserGroupModifyOpenApiVO := *openapiclient.NewSslVpnUserGroupModifyOpenApiVO() // SslVpnUserGroupModifyOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLVPNAPI.ModifySslVpnUserGroup(context.Background(), omadacId, siteId, id).SslVpnUserGroupModifyOpenApiVO(sslVpnUserGroupModifyOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLVPNAPI.ModifySslVpnUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySslVpnUserGroup`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSLVPNAPI.ModifySslVpnUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | User group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySslVpnUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sslVpnUserGroupModifyOpenApiVO** | [**SslVpnUserGroupModifyOpenApiVO**](SslVpnUserGroupModifyOpenApiVO.md) |  | 

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


# \VPNAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientToSiteVpnClient**](VPNAPI.md#CreateClientToSiteVpnClient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients | Create client-to-site VPN client
[**CreateClientToSiteVpnServer**](VPNAPI.md#CreateClientToSiteVpnServer) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers | Create client-to-site VPN server
[**CreateIpsecFailover**](VPNAPI.md#CreateIpsecFailover) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ipsec_failovers | Create IPsec failover
[**CreateSiteToSiteVpn**](VPNAPI.md#CreateSiteToSiteVpn) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns | Create site-to-site VPN
[**CreateVpnUser**](VPNAPI.md#CreateVpnUser) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users | Create VPN user
[**CreateVpnUserV2**](VPNAPI.md#CreateVpnUserV2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/vpn/users | Create VPN user V2
[**DeleteClientToSiteVpnClient**](VPNAPI.md#DeleteClientToSiteVpnClient) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients/{vpnId} | Delete client-to-site VPN client
[**DeleteClientToSiteVpnServer**](VPNAPI.md#DeleteClientToSiteVpnServer) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId} | Delete client-to-site VPN server
[**DeleteIpsecFailover**](VPNAPI.md#DeleteIpsecFailover) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ipsec_failovers/{failoverId} | Delete IPsec failover
[**DeleteSiteToSiteVpn**](VPNAPI.md#DeleteSiteToSiteVpn) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/{vpnId} | Delete site-to-site VPN
[**DeleteVpn**](VPNAPI.md#DeleteVpn) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/{vpnId} | Delete VPN
[**DeleteVpnUser**](VPNAPI.md#DeleteVpnUser) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users/{userId} | Delete VPN user
[**GetAllVpnList**](VPNAPI.md#GetAllVpnList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn | Get All VPN list
[**GetClientToSiteVpnClientList**](VPNAPI.md#GetClientToSiteVpnClientList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients | Get client-to-site VPN client list
[**GetClientToSiteVpnServerInfo**](VPNAPI.md#GetClientToSiteVpnServerInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId} | Get client-to-site VPN server info
[**GetClientToSiteVpnServerList**](VPNAPI.md#GetClientToSiteVpnServerList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers | Get client-to-site VPN server list
[**GetClientToSiteVpnServerUserList**](VPNAPI.md#GetClientToSiteVpnServerUserList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId}/users | Get user list for client-to-site VPN server
[**GetGridIpsecFailover**](VPNAPI.md#GetGridIpsecFailover) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ipsec_failovers | Get IPsec failover list
[**GetGridVpnUser**](VPNAPI.md#GetGridVpnUser) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users | Get VPN user list
[**GetSiteToSiteVpnInfo**](VPNAPI.md#GetSiteToSiteVpnInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/{vpnId} | Get site-to-site VPN info
[**GetSiteToSiteVpnList**](VPNAPI.md#GetSiteToSiteVpnList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns | Get site-to-site VPN list
[**GetVpnClientToSiteClientInfo**](VPNAPI.md#GetVpnClientToSiteClientInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients/{vpnId} | Get client-to-site VPN client info
[**ListRemoteSite**](VPNAPI.md#ListRemoteSite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/remoteSites | List Remote Site
[**ModifyClientToSiteVpnClient**](VPNAPI.md#ModifyClientToSiteVpnClient) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients/{vpnId} | Modify client-to-site VPN client
[**ModifyClientToSiteVpnServer**](VPNAPI.md#ModifyClientToSiteVpnServer) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId} | Modify client-to-site VPN server
[**ModifyIpsecFailover**](VPNAPI.md#ModifyIpsecFailover) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ipsec_failovers/{failoverId} | Modify IPsec failover
[**ModifySiteToSiteVpn**](VPNAPI.md#ModifySiteToSiteVpn) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/{vpnId} | Modify site-to-site VPN
[**ModifyVpnUser**](VPNAPI.md#ModifyVpnUser) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users/{userId} | Modify VPN user
[**ModifyVpnUserV2**](VPNAPI.md#ModifyVpnUserV2) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/vpn/users/{userId} | Modify VPN user V2
[**UploadVpnCertificateFile**](VPNAPI.md#UploadVpnCertificateFile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/certificate | Upload VPN certificate file



## CreateClientToSiteVpnClient

> OperationResponseWithoutResult CreateClientToSiteVpnClient(ctx, omadacId, siteId).ClientToSiteVpnClient(clientToSiteVpnClient).Execute()

Create client-to-site VPN client



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
	clientToSiteVpnClient := *openapiclient.NewClientToSiteVpnClient(int32(123), "Name_example", []string{"Wan_example"}) // ClientToSiteVpnClient | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateClientToSiteVpnClient(context.Background(), omadacId, siteId).ClientToSiteVpnClient(clientToSiteVpnClient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateClientToSiteVpnClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClientToSiteVpnClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateClientToSiteVpnClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientToSiteVpnClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientToSiteVpnClient** | [**ClientToSiteVpnClient**](ClientToSiteVpnClient.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientToSiteVpnServer

> OperationResponseWithoutResult CreateClientToSiteVpnServer(ctx, omadacId, siteId).ClientToSiteVpnServer(clientToSiteVpnServer).Execute()

Create client-to-site VPN server



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
	clientToSiteVpnServer := *openapiclient.NewClientToSiteVpnServer(int32(123), *openapiclient.NewIPSubnetsVO("Ip_example", int32(123)), "Name_example", []string{"Wan_example"}) // ClientToSiteVpnServer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateClientToSiteVpnServer(context.Background(), omadacId, siteId).ClientToSiteVpnServer(clientToSiteVpnServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateClientToSiteVpnServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClientToSiteVpnServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateClientToSiteVpnServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientToSiteVpnServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientToSiteVpnServer** | [**ClientToSiteVpnServer**](ClientToSiteVpnServer.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIpsecFailover

> OperationResponseWithoutResult CreateIpsecFailover(ctx, omadacId, siteId).IPsecFailover(iPsecFailover).Execute()

Create IPsec failover



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
	iPsecFailover := *openapiclient.NewIPsecFailover([]string{"Candidates_example"}, "Name_example", "Primary_example") // IPsecFailover | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateIpsecFailover(context.Background(), omadacId, siteId).IPsecFailover(iPsecFailover).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateIpsecFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIpsecFailover`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateIpsecFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIpsecFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iPsecFailover** | [**IPsecFailover**](IPsecFailover.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteToSiteVpn

> OperationResponseWithoutResult CreateSiteToSiteVpn(ctx, omadacId, siteId).SiteToSiteVpn(siteToSiteVpn).Execute()

Create site-to-site VPN



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
	siteToSiteVpn := *openapiclient.NewSiteToSiteVpn("Name_example", int32(123), false) // SiteToSiteVpn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateSiteToSiteVpn(context.Background(), omadacId, siteId).SiteToSiteVpn(siteToSiteVpn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateSiteToSiteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteToSiteVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteToSiteVpn** | [**SiteToSiteVpn**](SiteToSiteVpn.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpnUser

> OperationResponseWithoutResult CreateVpnUser(ctx, omadacId, siteId).VpnUser(vpnUser).Execute()

Create VPN user



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
	vpnUser := *openapiclient.NewVpnUser() // VpnUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateVpnUser(context.Background(), omadacId, siteId).VpnUser(vpnUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnUser** | [**VpnUser**](VpnUser.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpnUserV2

> OperationResponseWithoutResult CreateVpnUserV2(ctx, omadacId, siteId).VpnUserRequest(vpnUserRequest).Execute()

Create VPN user V2



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
	vpnUserRequest := *openapiclient.NewVpnUserRequest("Password_example", []string{"Servers_example"}, "Username_example") // VpnUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateVpnUserV2(context.Background(), omadacId, siteId).VpnUserRequest(vpnUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateVpnUserV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpnUserV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateVpnUserV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnUserV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnUserRequest** | [**VpnUserRequest**](VpnUserRequest.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientToSiteVpnClient

> OperationResponseWithoutResult DeleteClientToSiteVpnClient(ctx, omadacId, siteId, vpnId).Execute()

Delete client-to-site VPN client



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
	vpnId := "vpnId_example" // string | VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.DeleteClientToSiteVpnClient(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DeleteClientToSiteVpnClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClientToSiteVpnClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DeleteClientToSiteVpnClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientToSiteVpnClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientToSiteVpnServer

> OperationResponseWithoutResult DeleteClientToSiteVpnServer(ctx, omadacId, siteId, vpnId).Execute()

Delete client-to-site VPN server



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
	vpnId := "vpnId_example" // string | VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.DeleteClientToSiteVpnServer(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DeleteClientToSiteVpnServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClientToSiteVpnServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DeleteClientToSiteVpnServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientToSiteVpnServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIpsecFailover

> OperationResponseWithoutResult DeleteIpsecFailover(ctx, omadacId, siteId, failoverId).Execute()

Delete IPsec failover



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
	failoverId := "failoverId_example" // string | IPSec failover ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.DeleteIpsecFailover(context.Background(), omadacId, siteId, failoverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DeleteIpsecFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIpsecFailover`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DeleteIpsecFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**failoverId** | **string** | IPSec failover ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpsecFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteToSiteVpn

> OperationResponseWithoutResult DeleteSiteToSiteVpn(ctx, omadacId, siteId, vpnId).Execute()

Delete site-to-site VPN



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
	vpnId := "vpnId_example" // string | VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.DeleteSiteToSiteVpn(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DeleteSiteToSiteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSiteToSiteVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DeleteSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpn

> OperationResponseWithoutResult DeleteVpn(ctx, omadacId, siteId, vpnId).Execute()

Delete VPN



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
	vpnId := "vpnId_example" // string | VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.DeleteVpn(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DeleteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DeleteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpnUser

> OperationResponseWithoutResult DeleteVpnUser(ctx, omadacId, siteId, userId).Execute()

Delete VPN user



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
	userId := "userId_example" // string | VPN user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.DeleteVpnUser(context.Background(), omadacId, siteId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DeleteVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DeleteVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**userId** | **string** | VPN user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllVpnList

> OperationResponseVpnOpenApiGridVOVPN GetAllVpnList(ctx, omadacId, siteId).Execute()

Get All VPN list



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
	resp, r, err := apiClient.VPNAPI.GetAllVpnList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetAllVpnList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllVpnList`: OperationResponseVpnOpenApiGridVOVPN
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetAllVpnList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllVpnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVpnOpenApiGridVOVPN**](OperationResponseVpnOpenApiGridVOVPN.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientToSiteVpnClientList

> OperationResponseVpnOpenApiGridVOClientToSiteVpnClient GetClientToSiteVpnClientList(ctx, omadacId, siteId).Execute()

Get client-to-site VPN client list



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
	resp, r, err := apiClient.VPNAPI.GetClientToSiteVpnClientList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetClientToSiteVpnClientList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientToSiteVpnClientList`: OperationResponseVpnOpenApiGridVOClientToSiteVpnClient
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetClientToSiteVpnClientList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientToSiteVpnClientListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVpnOpenApiGridVOClientToSiteVpnClient**](OperationResponseVpnOpenApiGridVOClientToSiteVpnClient.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientToSiteVpnServerInfo

> OperationResponseClientToSiteVpnServer GetClientToSiteVpnServerInfo(ctx, omadacId, siteId, vpnId).Execute()

Get client-to-site VPN server info



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
	vpnId := "vpnId_example" // string | VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetClientToSiteVpnServerInfo(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetClientToSiteVpnServerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientToSiteVpnServerInfo`: OperationResponseClientToSiteVpnServer
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetClientToSiteVpnServerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientToSiteVpnServerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseClientToSiteVpnServer**](OperationResponseClientToSiteVpnServer.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientToSiteVpnServerList

> OperationResponseVpnOpenApiGridVOClientToSiteVpnServer GetClientToSiteVpnServerList(ctx, omadacId, siteId).Execute()

Get client-to-site VPN server list



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
	resp, r, err := apiClient.VPNAPI.GetClientToSiteVpnServerList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetClientToSiteVpnServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientToSiteVpnServerList`: OperationResponseVpnOpenApiGridVOClientToSiteVpnServer
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetClientToSiteVpnServerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientToSiteVpnServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVpnOpenApiGridVOClientToSiteVpnServer**](OperationResponseVpnOpenApiGridVOClientToSiteVpnServer.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientToSiteVpnServerUserList

> OperationResponseListVpnUserResponse GetClientToSiteVpnServerUserList(ctx, omadacId, siteId, vpnId).Execute()

Get user list for client-to-site VPN server



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
	vpnId := "vpnId_example" // string | VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetClientToSiteVpnServerUserList(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetClientToSiteVpnServerUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientToSiteVpnServerUserList`: OperationResponseListVpnUserResponse
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetClientToSiteVpnServerUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientToSiteVpnServerUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListVpnUserResponse**](OperationResponseListVpnUserResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridIpsecFailover

> OperationResponseGridVOIPsecFailover GetGridIpsecFailover(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get IPsec failover list



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
	resp, r, err := apiClient.VPNAPI.GetGridIpsecFailover(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetGridIpsecFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridIpsecFailover`: OperationResponseGridVOIPsecFailover
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetGridIpsecFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridIpsecFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOIPsecFailover**](OperationResponseGridVOIPsecFailover.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnUser

> OperationResponseVpnUserOpenApiGridVOVpnUserResponse GetGridVpnUser(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get VPN user list



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
	searchKey := "searchKey_example" // string | searchKey (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetGridVpnUser(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetGridVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnUser`: OperationResponseVpnUserOpenApiGridVOVpnUserResponse
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetGridVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | searchKey | 

### Return type

[**OperationResponseVpnUserOpenApiGridVOVpnUserResponse**](OperationResponseVpnUserOpenApiGridVOVpnUserResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteToSiteVpnInfo

> OperationResponseSiteToSiteVpn GetSiteToSiteVpnInfo(ctx, omadacId, siteId, vpnId).Execute()

Get site-to-site VPN info



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
	vpnId := "vpnId_example" // string | VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetSiteToSiteVpnInfo(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetSiteToSiteVpnInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteToSiteVpnInfo`: OperationResponseSiteToSiteVpn
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetSiteToSiteVpnInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteToSiteVpnInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSiteToSiteVpn**](OperationResponseSiteToSiteVpn.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteToSiteVpnList

> OperationResponseListSiteToSiteVpn GetSiteToSiteVpnList(ctx, omadacId, siteId).Execute()

Get site-to-site VPN list



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
	resp, r, err := apiClient.VPNAPI.GetSiteToSiteVpnList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetSiteToSiteVpnList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteToSiteVpnList`: OperationResponseListSiteToSiteVpn
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetSiteToSiteVpnList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteToSiteVpnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListSiteToSiteVpn**](OperationResponseListSiteToSiteVpn.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnClientToSiteClientInfo

> OperationResponseClientToSiteVpnClient GetVpnClientToSiteClientInfo(ctx, omadacId, siteId, vpnId).Execute()

Get client-to-site VPN client info



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
	vpnId := "vpnId_example" // string | VPN ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetVpnClientToSiteClientInfo(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnClientToSiteClientInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnClientToSiteClientInfo`: OperationResponseClientToSiteVpnClient
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnClientToSiteClientInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnClientToSiteClientInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseClientToSiteVpnClient**](OperationResponseClientToSiteVpnClient.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemoteSite

> OperationResponseMapStringObject ListRemoteSite(ctx, omadacId, siteId).Execute()

List Remote Site



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
	resp, r, err := apiClient.VPNAPI.ListRemoteSite(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ListRemoteSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRemoteSite`: OperationResponseMapStringObject
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ListRemoteSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRemoteSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMapStringObject**](OperationResponseMapStringObject.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyClientToSiteVpnClient

> OperationResponseWithoutResult ModifyClientToSiteVpnClient(ctx, omadacId, siteId, vpnId).ClientToSiteVpnClient(clientToSiteVpnClient).Execute()

Modify client-to-site VPN client



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
	vpnId := "vpnId_example" // string | VPN ID
	clientToSiteVpnClient := *openapiclient.NewClientToSiteVpnClient(int32(123), "Name_example", []string{"Wan_example"}) // ClientToSiteVpnClient | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyClientToSiteVpnClient(context.Background(), omadacId, siteId, vpnId).ClientToSiteVpnClient(clientToSiteVpnClient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyClientToSiteVpnClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyClientToSiteVpnClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyClientToSiteVpnClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyClientToSiteVpnClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientToSiteVpnClient** | [**ClientToSiteVpnClient**](ClientToSiteVpnClient.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyClientToSiteVpnServer

> OperationResponseWithoutResult ModifyClientToSiteVpnServer(ctx, omadacId, siteId, vpnId).ClientToSiteVpnServer(clientToSiteVpnServer).Execute()

Modify client-to-site VPN server



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
	vpnId := "vpnId_example" // string | VPN ID
	clientToSiteVpnServer := *openapiclient.NewClientToSiteVpnServer(int32(123), *openapiclient.NewIPSubnetsVO("Ip_example", int32(123)), "Name_example", []string{"Wan_example"}) // ClientToSiteVpnServer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyClientToSiteVpnServer(context.Background(), omadacId, siteId, vpnId).ClientToSiteVpnServer(clientToSiteVpnServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyClientToSiteVpnServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyClientToSiteVpnServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyClientToSiteVpnServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyClientToSiteVpnServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientToSiteVpnServer** | [**ClientToSiteVpnServer**](ClientToSiteVpnServer.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIpsecFailover

> OperationResponseWithoutResult ModifyIpsecFailover(ctx, omadacId, siteId, failoverId).IPsecFailover(iPsecFailover).Execute()

Modify IPsec failover



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
	failoverId := "failoverId_example" // string | IPSec failover ID
	iPsecFailover := *openapiclient.NewIPsecFailover([]string{"Candidates_example"}, "Name_example", "Primary_example") // IPsecFailover | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyIpsecFailover(context.Background(), omadacId, siteId, failoverId).IPsecFailover(iPsecFailover).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyIpsecFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIpsecFailover`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyIpsecFailover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**failoverId** | **string** | IPSec failover ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIpsecFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iPsecFailover** | [**IPsecFailover**](IPsecFailover.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySiteToSiteVpn

> OperationResponseWithoutResult ModifySiteToSiteVpn(ctx, omadacId, siteId, vpnId).SiteToSiteVpn(siteToSiteVpn).Execute()

Modify site-to-site VPN



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
	vpnId := "vpnId_example" // string | VPN ID
	siteToSiteVpn := *openapiclient.NewSiteToSiteVpn("Name_example", int32(123), false) // SiteToSiteVpn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifySiteToSiteVpn(context.Background(), omadacId, siteId, vpnId).SiteToSiteVpn(siteToSiteVpn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifySiteToSiteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySiteToSiteVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifySiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**vpnId** | **string** | VPN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **siteToSiteVpn** | [**SiteToSiteVpn**](SiteToSiteVpn.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyVpnUser

> OperationResponseWithoutResult ModifyVpnUser(ctx, omadacId, siteId, userId).VpnUser(vpnUser).Execute()

Modify VPN user



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
	userId := "userId_example" // string | VPN user ID
	vpnUser := *openapiclient.NewVpnUser() // VpnUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyVpnUser(context.Background(), omadacId, siteId, userId).VpnUser(vpnUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**userId** | **string** | VPN user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpnUser** | [**VpnUser**](VpnUser.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyVpnUserV2

> OperationResponseWithoutResult ModifyVpnUserV2(ctx, omadacId, siteId, userId).VpnUserRequest(vpnUserRequest).Execute()

Modify VPN user V2



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
	userId := "userId_example" // string | VPN user ID
	vpnUserRequest := *openapiclient.NewVpnUserRequest("Password_example", []string{"Servers_example"}, "Username_example") // VpnUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyVpnUserV2(context.Background(), omadacId, siteId, userId).VpnUserRequest(vpnUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyVpnUserV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVpnUserV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyVpnUserV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**userId** | **string** | VPN user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVpnUserV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpnUserRequest** | [**VpnUserRequest**](VpnUserRequest.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVpnCertificateFile

> OperationResponseWithoutResult UploadVpnCertificateFile(ctx, omadacId, siteId).File(file).Execute()

Upload VPN certificate file



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
	file := os.NewFile(1234, "some_file") // *os.File | 
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.UploadVpnCertificateFile(context.Background(), omadacId, siteId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.UploadVpnCertificateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadVpnCertificateFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.UploadVpnCertificateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadVpnCertificateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 



### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


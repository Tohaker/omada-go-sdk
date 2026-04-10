# VPNAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteVpn**](VPNAPI.md#batchdeletevpn) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn | Batch delete VPN
[**BatchDeleteVpnUser**](VPNAPI.md#batchdeletevpnuser) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users | Batch delete VPN user
[**CheckUsedInVpns**](VPNAPI.md#checkusedinvpns) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/checkUsed | Check used function for multiple VPN items
[**CheckUsedVpn**](VPNAPI.md#checkusedvpn) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/{vpnId}/checkUsed | Check used function for single VPN item
[**CheckValueAvailable**](VPNAPI.md#checkvalueavailable) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/checkValue | Check whether vpn value is available
[**CreateClientToSiteVpnClient**](VPNAPI.md#createclienttositevpnclient) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients | Create client-to-site VPN client
[**CreateClientToSiteVpnServer**](VPNAPI.md#createclienttositevpnserver) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers | Create client-to-site VPN server
[**CreateIpsecFailover**](VPNAPI.md#createipsecfailover) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ipsec_failovers | Create IPsec failover
[**CreateS2SAutoVpn**](VPNAPI.md#creates2sautovpn) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/auto | Create site-to-site VPN by auto
[**CreateS2SManualVpn**](VPNAPI.md#creates2smanualvpn) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns | Create site-to-site VPN by manual
[**CreateSiteToSiteVpn**](VPNAPI.md#createsitetositevpn) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns | Create site-to-site VPN
[**CreateVpnClient**](VPNAPI.md#createvpnclient) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients | Create client-to-site VPN client V2
[**CreateVpnServer**](VPNAPI.md#createvpnserver) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers | Create client-to-site VPN server V2
[**CreateVpnUser**](VPNAPI.md#createvpnuser) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users | Create VPN user
[**CreateVpnUserV2**](VPNAPI.md#createvpnuserv2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/vpn/users | Create VPN user V2
[**CreateVpnUserV3**](VPNAPI.md#createvpnuserv3) | **Post** /openapi/v3/{omadacId}/sites/{siteId}/vpn/users | Create VPN user V3
[**DeleteClientToSiteVpnClient**](VPNAPI.md#deleteclienttositevpnclient) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients/{vpnId} | Delete client-to-site VPN client
[**DeleteClientToSiteVpnServer**](VPNAPI.md#deleteclienttositevpnserver) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId} | Delete client-to-site VPN server
[**DeleteIpsecFailover**](VPNAPI.md#deleteipsecfailover) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ipsec_failovers/{failoverId} | Delete IPsec failover
[**DeleteSiteToSiteVpn**](VPNAPI.md#deletesitetositevpn) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/{vpnId} | Delete site-to-site VPN
[**DeleteVpn**](VPNAPI.md#deletevpn) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/{vpnId} | Delete VPN
[**DeleteVpnUser**](VPNAPI.md#deletevpnuser) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users/{userId} | Delete VPN user
[**DeleteVpnV2**](VPNAPI.md#deletevpnv2) | **Delete** /openapi/v2/{omadacId}/sites/{siteId}/vpn/{vpnId} | Delete VPN V2
[**DisconnectSslVpnTunnel**](VPNAPI.md#disconnectsslvpntunnel) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/stats/sslvpn/tunnel/{tunnelId}/disconnect | Disconnect SSL VPN tunnel
[**DownloadVpnCertificate**](VPNAPI.md#downloadvpncertificate) | **Get** /openapi/v1/{omadacId}/files/sites/{siteId}/vpn/{vpnId}/certificate | Download Open VPN or SSL VPN certificate
[**GetAllVpnList**](VPNAPI.md#getallvpnlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn | Get All VPN list
[**GetBriefVpnUserByServerIdList**](VPNAPI.md#getbriefvpnuserbyserveridlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/userlist/server/{serverId} | Get brief VPN user list by VPN server ID without page
[**GetClientToSiteVpnClientList**](VPNAPI.md#getclienttositevpnclientlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients | Get client-to-site VPN client list
[**GetClientToSiteVpnServerInfo**](VPNAPI.md#getclienttositevpnserverinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId} | Get client-to-site VPN server info
[**GetClientToSiteVpnServerList**](VPNAPI.md#getclienttositevpnserverlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers | Get client-to-site VPN server list
[**GetClientToSiteVpnServerUserList**](VPNAPI.md#getclienttositevpnserveruserlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId}/users | Get user list for client-to-site VPN server
[**GetGridIpsecFailover**](VPNAPI.md#getgridipsecfailover) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ipsec_failovers | Get IPsec failover list
[**GetGridVpnClientV2**](VPNAPI.md#getgridvpnclientv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients | Get VPN Client summary list
[**GetGridVpnS2SV2**](VPNAPI.md#getgridvpns2sv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns | Get VPN Site to Site summary list
[**GetGridVpnServerV2**](VPNAPI.md#getgridvpnserverv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers | Get VPN Server summary list
[**GetGridVpnUser**](VPNAPI.md#getgridvpnuser) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users | Get VPN user list
[**GetGridVpnUserV2**](VPNAPI.md#getgridvpnuserv2) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/vpn/users | Get VPN user list V2
[**GetSiteToSiteVpnInfo**](VPNAPI.md#getsitetositevpninfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/{vpnId} | Get site-to-site VPN info
[**GetSiteToSiteVpnList**](VPNAPI.md#getsitetositevpnlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns | Get site-to-site VPN list
[**GetSslVpnUserGroupBriefList**](VPNAPI.md#getsslvpnusergroupbrieflist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/briefusergroups | Get SSL VPN user Group list V2
[**GetVpnAvailableIpPool**](VPNAPI.md#getvpnavailableippool) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ippool | Get available IP pools for VPN
[**GetVpnClientDetailInfo**](VPNAPI.md#getvpnclientdetailinfo) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients/{vpnId} | Get client-to-site VPN client detail info
[**GetVpnClientToSiteClientInfo**](VPNAPI.md#getvpnclienttositeclientinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients/{vpnId} | Get client-to-site VPN client info
[**GetVpnDefaultValue**](VPNAPI.md#getvpndefaultvalue) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/defaultValue | Get default value for VPN
[**GetVpnPreSharedKey**](VPNAPI.md#getvpnpresharedkey) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/psk | Generate a random pre-shared key for IPSec
[**GetVpnS2SDetailInfo**](VPNAPI.md#getvpns2sdetailinfo) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/{vpnId} | Get site-to-site VPN detail info
[**GetVpnServerDetailInfo**](VPNAPI.md#getvpnserverdetailinfo) | **Get** /openapi/v2/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId} | Get client-to-site VPN server detail info
[**GetVpnServerUserListV3**](VPNAPI.md#getvpnserveruserlistv3) | **Get** /openapi/v3/{omadacId}/sites/{siteId}/vpn/{vpnId}/users | Get user list by VPN server ID
[**GetVpnUserList**](VPNAPI.md#getvpnuserlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/userlist/{protocol} | Get VPN user list without page
[**GetVpnUserServerList**](VPNAPI.md#getvpnuserserverlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vpn/userServers | Get VPN server list for user
[**ListRemoteSite**](VPNAPI.md#listremotesite) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/remoteSites | List Remote Site
[**LockSslVpnTunnel**](VPNAPI.md#locksslvpntunnel) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/stats/sslvpn/tunnel/{tunnelId}/lock | Lock SSL VPN tunnel
[**ModifyClientStatus**](VPNAPI.md#modifyclientstatus) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/{vpnId}/status | Modify VPN status
[**ModifyClientToSiteVpnClient**](VPNAPI.md#modifyclienttositevpnclient) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients/{vpnId} | Modify client-to-site VPN client
[**ModifyClientToSiteVpnServer**](VPNAPI.md#modifyclienttositevpnserver) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId} | Modify client-to-site VPN server
[**ModifyIpsecFailover**](VPNAPI.md#modifyipsecfailover) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/ipsec_failovers/{failoverId} | Modify IPsec failover
[**ModifyS2SAutoVpn**](VPNAPI.md#modifys2sautovpn) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/auto/{vpnId} | Modify site-to-site VPN by auto
[**ModifyS2SManualVpn**](VPNAPI.md#modifys2smanualvpn) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/{vpnId} | Modify site-to-site VPN by manual
[**ModifySiteToSiteVpn**](VPNAPI.md#modifysitetositevpn) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/site-to-site-vpns/{vpnId} | Modify site-to-site VPN
[**ModifyVpnClient**](VPNAPI.md#modifyvpnclient) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-clients/{vpnId} | Modify client-to-site VPN client V2
[**ModifyVpnServer**](VPNAPI.md#modifyvpnserver) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/vpn/client-to-site-vpn-servers/{vpnId} | Modify client-to-site VPN server V2
[**ModifyVpnUser**](VPNAPI.md#modifyvpnuser) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/vpn/users/{userId} | Modify VPN user
[**ModifyVpnUserV2**](VPNAPI.md#modifyvpnuserv2) | **Patch** /openapi/v2/{omadacId}/sites/{siteId}/vpn/users/{userId} | Modify VPN user V2
[**ModifyVpnUserV3**](VPNAPI.md#modifyvpnuserv3) | **Patch** /openapi/v3/{omadacId}/sites/{siteId}/vpn/users/{userId} | Modify VPN user V3
[**UploadVpnCertificateFile**](VPNAPI.md#uploadvpncertificatefile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/vpn/certificate | Upload VPN certificate file
[**UploadVpnCertificateFileV2**](VPNAPI.md#uploadvpncertificatefilev2) | **Post** /openapi/v2/{omadacId}/files/sites/{siteId}/vpn/certificate | Upload VPN certificate file V2



## BatchDeleteVpn

> OperationResponseWithoutResult BatchDeleteVpn(ctx, omadacId, siteId).BatchSelectVpnVO(batchSelectVpnVO).Execute()

Batch delete VPN



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
	batchSelectVpnVO := *openapiclient.NewBatchSelectVpnVO() // BatchSelectVpnVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.BatchDeleteVpn(context.Background(), omadacId, siteId).BatchSelectVpnVO(batchSelectVpnVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.BatchDeleteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.BatchDeleteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchSelectVpnVO** | [**BatchSelectVpnVO**](BatchSelectVpnVO.md) |  | 

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


## BatchDeleteVpnUser

> OperationResponseWithoutResult BatchDeleteVpnUser(ctx, omadacId, siteId).BatchSelectVpnUserVO(batchSelectVpnUserVO).Execute()

Batch delete VPN user



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
	batchSelectVpnUserVO := *openapiclient.NewBatchSelectVpnUserVO() // BatchSelectVpnUserVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.BatchDeleteVpnUser(context.Background(), omadacId, siteId).BatchSelectVpnUserVO(batchSelectVpnUserVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.BatchDeleteVpnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteVpnUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.BatchDeleteVpnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteVpnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchSelectVpnUserVO** | [**BatchSelectVpnUserVO**](BatchSelectVpnUserVO.md) |  | 

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


## CheckUsedInVpns

> OperationResponseVpnListUsedFunctionOpenApiVO CheckUsedInVpns(ctx, omadacId, siteId).BatchSelectVpnVO(batchSelectVpnVO).Execute()

Check used function for multiple VPN items



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
	batchSelectVpnVO := *openapiclient.NewBatchSelectVpnVO() // BatchSelectVpnVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CheckUsedInVpns(context.Background(), omadacId, siteId).BatchSelectVpnVO(batchSelectVpnVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CheckUsedInVpns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsedInVpns`: OperationResponseVpnListUsedFunctionOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CheckUsedInVpns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsedInVpnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchSelectVpnVO** | [**BatchSelectVpnVO**](BatchSelectVpnVO.md) |  | 

### Return type

[**OperationResponseVpnListUsedFunctionOpenApiVO**](OperationResponseVpnListUsedFunctionOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckUsedVpn

> OperationResponseFunctionOpenApiVO CheckUsedVpn(ctx, omadacId, siteId, vpnId).Execute()

Check used function for single VPN item



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
	resp, r, err := apiClient.VPNAPI.CheckUsedVpn(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CheckUsedVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsedVpn`: OperationResponseFunctionOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CheckUsedVpn`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCheckUsedVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseFunctionOpenApiVO**](OperationResponseFunctionOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckValueAvailable

> OperationResponseWithoutResult CheckValueAvailable(ctx, omadacId, siteId).VpnValueAvailableVO(vpnValueAvailableVO).Execute()

Check whether vpn value is available



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
	vpnValueAvailableVO := *openapiclient.NewVpnValueAvailableVO(int32(123)) // VpnValueAvailableVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CheckValueAvailable(context.Background(), omadacId, siteId).VpnValueAvailableVO(vpnValueAvailableVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CheckValueAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckValueAvailable`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CheckValueAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckValueAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnValueAvailableVO** | [**VpnValueAvailableVO**](VpnValueAvailableVO.md) |  | 

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIpsecFailover

> OperationResponseWithoutResult CreateIpsecFailover(ctx, omadacId, siteId).IPsecFailoverConfiguration(iPsecFailoverConfiguration).Execute()

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
	iPsecFailoverConfiguration := *openapiclient.NewIPsecFailoverConfiguration([]string{"Candidates_example"}, "Name_example", "Primary_example") // IPsecFailoverConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateIpsecFailover(context.Background(), omadacId, siteId).IPsecFailoverConfiguration(iPsecFailoverConfiguration).Execute()
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


 **iPsecFailoverConfiguration** | [**IPsecFailoverConfiguration**](IPsecFailoverConfiguration.md) |  | 

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


## CreateS2SAutoVpn

> OperationResponseWithoutResult CreateS2SAutoVpn(ctx, omadacId, siteId).VpnSiteToSiteAutoConfigOpenApiVO(vpnSiteToSiteAutoConfigOpenApiVO).Execute()

Create site-to-site VPN by auto



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
	vpnSiteToSiteAutoConfigOpenApiVO := *openapiclient.NewVpnSiteToSiteAutoConfigOpenApiVO("Name_example", "RemoteSite_example", false, int32(123)) // VpnSiteToSiteAutoConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateS2SAutoVpn(context.Background(), omadacId, siteId).VpnSiteToSiteAutoConfigOpenApiVO(vpnSiteToSiteAutoConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateS2SAutoVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateS2SAutoVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateS2SAutoVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateS2SAutoVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnSiteToSiteAutoConfigOpenApiVO** | [**VpnSiteToSiteAutoConfigOpenApiVO**](VpnSiteToSiteAutoConfigOpenApiVO.md) |  | 

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


## CreateS2SManualVpn

> OperationResponseWithoutResult CreateS2SManualVpn(ctx, omadacId, siteId).VpnSiteToSiteManualConfigOpenApiVO(vpnSiteToSiteManualConfigOpenApiVO).Execute()

Create site-to-site VPN by manual



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
	vpnSiteToSiteManualConfigOpenApiVO := *openapiclient.NewVpnSiteToSiteManualConfigOpenApiVO(*openapiclient.NewVpnAdvancedSettingOpenApiVO(), int32(123), "Name_example", "PreSharedKey_example", "PrivateKey_example", "RemoteIp_example", []string{"RemoteSubnet_example"}, int32(123), false, "TunnelIp_example", int32(123), []string{"Wans_example"}) // VpnSiteToSiteManualConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateS2SManualVpn(context.Background(), omadacId, siteId).VpnSiteToSiteManualConfigOpenApiVO(vpnSiteToSiteManualConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateS2SManualVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateS2SManualVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateS2SManualVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateS2SManualVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnSiteToSiteManualConfigOpenApiVO** | [**VpnSiteToSiteManualConfigOpenApiVO**](VpnSiteToSiteManualConfigOpenApiVO.md) |  | 

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpnClient

> OperationResponseWithoutResult CreateVpnClient(ctx, omadacId, siteId).VpnClientConfigOpenApiVO(vpnClientConfigOpenApiVO).Execute()

Create client-to-site VPN client V2



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
	vpnClientConfigOpenApiVO := *openapiclient.NewVpnClientConfigOpenApiVO("Name_example", "RemoteIp_example", "ServerPublicKey_example", int32(123), int32(123), false, *openapiclient.NewVpnCertificateOpenApiVO(), int32(123), []string{"Wans_example"}) // VpnClientConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateVpnClient(context.Background(), omadacId, siteId).VpnClientConfigOpenApiVO(vpnClientConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateVpnClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpnClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateVpnClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnClientConfigOpenApiVO** | [**VpnClientConfigOpenApiVO**](VpnClientConfigOpenApiVO.md) |  | 

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


## CreateVpnServer

> OperationResponseWithoutResult CreateVpnServer(ctx, omadacId, siteId).VpnServerConfigOpenApiVO(vpnServerConfigOpenApiVO).Execute()

Create client-to-site VPN server V2



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
	vpnServerConfigOpenApiVO := *openapiclient.NewVpnServerConfigOpenApiVO(*openapiclient.NewVpnAdvancedSettingOpenApiVO(), false, *openapiclient.NewLockSettingOpenApiVO(false), int32(123), int32(123), "Name_example", *openapiclient.NewLockSettingOpenApiVO(false), "PreSharedKey_example", "PrivateKey_example", "RemoteIp_example", int32(123), int32(123), false, int32(123), []string{"VpnUserList_example"}, []string{"Wans_example"}) // VpnServerConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateVpnServer(context.Background(), omadacId, siteId).VpnServerConfigOpenApiVO(vpnServerConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateVpnServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpnServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateVpnServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnServerConfigOpenApiVO** | [**VpnServerConfigOpenApiVO**](VpnServerConfigOpenApiVO.md) |  | 

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpnUserV3

> OperationResponseVpnUserBriefVO CreateVpnUserV3(ctx, omadacId, siteId).VpnUserConfigVO(vpnUserConfigVO).Execute()

Create VPN user V3



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
	vpnUserConfigVO := *openapiclient.NewVpnUserConfigVO() // VpnUserConfigVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.CreateVpnUserV3(context.Background(), omadacId, siteId).VpnUserConfigVO(vpnUserConfigVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.CreateVpnUserV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpnUserV3`: OperationResponseVpnUserBriefVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.CreateVpnUserV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnUserV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnUserConfigVO** | [**VpnUserConfigVO**](VpnUserConfigVO.md) |  | 

### Return type

[**OperationResponseVpnUserBriefVO**](OperationResponseVpnUserBriefVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpnV2

> OperationResponseWithoutResult DeleteVpnV2(ctx, omadacId, siteId, vpnId).Execute()

Delete VPN V2



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
	resp, r, err := apiClient.VPNAPI.DeleteVpnV2(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DeleteVpnV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVpnV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DeleteVpnV2`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteVpnV2Request struct via the builder pattern


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


## DisconnectSslVpnTunnel

> OperationResponseWithoutResult DisconnectSslVpnTunnel(ctx, omadacId, siteId, tunnelId).Execute()

Disconnect SSL VPN tunnel



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
	tunnelId := "tunnelId_example" // string | SSL VPN tunnel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.DisconnectSslVpnTunnel(context.Background(), omadacId, siteId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DisconnectSslVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectSslVpnTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DisconnectSslVpnTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tunnelId** | **string** | SSL VPN tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectSslVpnTunnelRequest struct via the builder pattern


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


## DownloadVpnCertificate

> OperationResponse DownloadVpnCertificate(ctx, omadacId, siteId, vpnId).Execute()

Download Open VPN or SSL VPN certificate



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
	resp, r, err := apiClient.VPNAPI.DownloadVpnCertificate(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.DownloadVpnCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadVpnCertificate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.DownloadVpnCertificate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDownloadVpnCertificateRequest struct via the builder pattern


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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBriefVpnUserByServerIdList

> OperationResponseListVpnUserBriefVO GetBriefVpnUserByServerIdList(ctx, omadacId, siteId, serverId).Execute()

Get brief VPN user list by VPN server ID without page



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
	serverId := "serverId_example" // string | VPN server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetBriefVpnUserByServerIdList(context.Background(), omadacId, siteId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetBriefVpnUserByServerIdList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBriefVpnUserByServerIdList`: OperationResponseListVpnUserBriefVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetBriefVpnUserByServerIdList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**serverId** | **string** | VPN server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBriefVpnUserByServerIdListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListVpnUserBriefVO**](OperationResponseListVpnUserBriefVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridIpsecFailover

> OperationResponseGridVOIPsecFailoverInformation GetGridIpsecFailover(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

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
	// response from `GetGridIpsecFailover`: OperationResponseGridVOIPsecFailoverInformation
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

[**OperationResponseGridVOIPsecFailoverInformation**](OperationResponseGridVOIPsecFailoverInformation.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnClientV2

> OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO GetGridVpnClientV2(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersVpnType(filtersVpnType).SortsWans(sortsWans).Execute()

Get VPN Client summary list



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
	filtersVpnType := "filtersVpnType_example" // string | Filter query parameters, support field vpnType. 0: L2TP; 1: PPTP; 3: OpenVPN; 4: WireGuard;. (optional)
	sortsWans := "sortsWans_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetGridVpnClientV2(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersVpnType(filtersVpnType).SortsWans(sortsWans).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetGridVpnClientV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnClientV2`: OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetGridVpnClientV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnClientV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name. | 
 **filtersVpnType** | **string** | Filter query parameters, support field vpnType. 0: L2TP; 1: PPTP; 3: OpenVPN; 4: WireGuard;. | 
 **sortsWans** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO**](OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnS2SV2

> OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO GetGridVpnS2SV2(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersVpnType(filtersVpnType).SortsWans(sortsWans).Execute()

Get VPN Site to Site summary list



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
	filtersVpnType := "filtersVpnType_example" // string | Filter query parameters, support field vpnType. 2: IPSec; 4: WireGuard; 5: SSL VPN. (optional)
	sortsWans := "sortsWans_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetGridVpnS2SV2(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersVpnType(filtersVpnType).SortsWans(sortsWans).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetGridVpnS2SV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnS2SV2`: OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetGridVpnS2SV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnS2SV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name. | 
 **filtersVpnType** | **string** | Filter query parameters, support field vpnType. 2: IPSec; 4: WireGuard; 5: SSL VPN. | 
 **sortsWans** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO**](OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnServerV2

> OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO GetGridVpnServerV2(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersVpnType(filtersVpnType).SortsWans(sortsWans).Execute()

Get VPN Server summary list



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
	filtersVpnType := "filtersVpnType_example" // string | Filter query parameters, support field vpnType. 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. (optional)
	sortsWans := "sortsWans_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetGridVpnServerV2(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersVpnType(filtersVpnType).SortsWans(sortsWans).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetGridVpnServerV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnServerV2`: OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetGridVpnServerV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnServerV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field name. | 
 **filtersVpnType** | **string** | Filter query parameters, support field vpnType. 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. | 
 **sortsWans** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO**](OperationResponseVpnSummaryOpenApiGridVOVpnSummaryVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridVpnUserV2

> OperationResponseVpnUserOpenApiGridVOVpnUserInfoVO GetGridVpnUserV2(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersProtocol(filtersProtocol).FiltersClientMode(filtersClientMode).SortsMaxConnections(sortsMaxConnections).SortsValidity(sortsValidity).Execute()

Get VPN user list V2



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
	filtersProtocol := "filtersProtocol_example" // string | Filter query parameters, support field protocol. 0: L2TP or PPTP; 1: OpenVPN; 2: SSL VPN. (optional)
	filtersClientMode := "filtersClientMode_example" // string | Filter query parameters, support field clientMode. 0: Network Extension Mode; 1: Client. (optional)
	sortsMaxConnections := "sortsMaxConnections_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsValidity := "sortsValidity_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetGridVpnUserV2(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersProtocol(filtersProtocol).FiltersClientMode(filtersClientMode).SortsMaxConnections(sortsMaxConnections).SortsValidity(sortsValidity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetGridVpnUserV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridVpnUserV2`: OperationResponseVpnUserOpenApiGridVOVpnUserInfoVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetGridVpnUserV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridVpnUserV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | searchKey | 
 **filtersProtocol** | **string** | Filter query parameters, support field protocol. 0: L2TP or PPTP; 1: OpenVPN; 2: SSL VPN. | 
 **filtersClientMode** | **string** | Filter query parameters, support field clientMode. 0: Network Extension Mode; 1: Client. | 
 **sortsMaxConnections** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsValidity** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseVpnUserOpenApiGridVOVpnUserInfoVO**](OperationResponseVpnUserOpenApiGridVOVpnUserInfoVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslVpnUserGroupBriefList

> OperationResponseSslVpnUserGroupGridVOSslVpnUserGroupBriefVO GetSslVpnUserGroupBriefList(ctx, omadacId, siteId).Execute()

Get SSL VPN user Group list V2



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
	resp, r, err := apiClient.VPNAPI.GetSslVpnUserGroupBriefList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetSslVpnUserGroupBriefList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSslVpnUserGroupBriefList`: OperationResponseSslVpnUserGroupGridVOSslVpnUserGroupBriefVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetSslVpnUserGroupBriefList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslVpnUserGroupBriefListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSslVpnUserGroupGridVOSslVpnUserGroupBriefVO**](OperationResponseSslVpnUserGroupGridVOSslVpnUserGroupBriefVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnAvailableIpPool

> OperationResponseVpnAvailableIpPoolVO GetVpnAvailableIpPool(ctx, omadacId, siteId).Execute()

Get available IP pools for VPN



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
	resp, r, err := apiClient.VPNAPI.GetVpnAvailableIpPool(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnAvailableIpPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnAvailableIpPool`: OperationResponseVpnAvailableIpPoolVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnAvailableIpPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnAvailableIpPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVpnAvailableIpPoolVO**](OperationResponseVpnAvailableIpPoolVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnClientDetailInfo

> OperationResponseVpnClientDetailVO GetVpnClientDetailInfo(ctx, omadacId, siteId, vpnId).Execute()

Get client-to-site VPN client detail info



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
	resp, r, err := apiClient.VPNAPI.GetVpnClientDetailInfo(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnClientDetailInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnClientDetailInfo`: OperationResponseVpnClientDetailVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnClientDetailInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetVpnClientDetailInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseVpnClientDetailVO**](OperationResponseVpnClientDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnDefaultValue

> OperationResponseVpnDefaultValueRespVO GetVpnDefaultValue(ctx, omadacId, siteId).VpnDefaultValueReqVO(vpnDefaultValueReqVO).Execute()

Get default value for VPN



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
	vpnDefaultValueReqVO := *openapiclient.NewVpnDefaultValueReqVO(int32(123)) // VpnDefaultValueReqVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetVpnDefaultValue(context.Background(), omadacId, siteId).VpnDefaultValueReqVO(vpnDefaultValueReqVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnDefaultValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnDefaultValue`: OperationResponseVpnDefaultValueRespVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnDefaultValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnDefaultValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnDefaultValueReqVO** | [**VpnDefaultValueReqVO**](VpnDefaultValueReqVO.md) |  | 

### Return type

[**OperationResponseVpnDefaultValueRespVO**](OperationResponseVpnDefaultValueRespVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnPreSharedKey

> OperationResponseVpnPreSharedKeyVO GetVpnPreSharedKey(ctx, omadacId, siteId).Execute()

Generate a random pre-shared key for IPSec



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
	resp, r, err := apiClient.VPNAPI.GetVpnPreSharedKey(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnPreSharedKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnPreSharedKey`: OperationResponseVpnPreSharedKeyVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnPreSharedKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnPreSharedKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVpnPreSharedKeyVO**](OperationResponseVpnPreSharedKeyVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnS2SDetailInfo

> OperationResponseVpnSiteToSiteDetailOpenApiVO GetVpnS2SDetailInfo(ctx, omadacId, siteId, vpnId).Execute()

Get site-to-site VPN detail info



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
	resp, r, err := apiClient.VPNAPI.GetVpnS2SDetailInfo(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnS2SDetailInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnS2SDetailInfo`: OperationResponseVpnSiteToSiteDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnS2SDetailInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetVpnS2SDetailInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseVpnSiteToSiteDetailOpenApiVO**](OperationResponseVpnSiteToSiteDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnServerDetailInfo

> OperationResponseVpnServerDetailVO GetVpnServerDetailInfo(ctx, omadacId, siteId, vpnId).Execute()

Get client-to-site VPN server detail info



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
	resp, r, err := apiClient.VPNAPI.GetVpnServerDetailInfo(context.Background(), omadacId, siteId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnServerDetailInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnServerDetailInfo`: OperationResponseVpnServerDetailVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnServerDetailInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetVpnServerDetailInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseVpnServerDetailVO**](OperationResponseVpnServerDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnServerUserListV3

> OperationResponseVpnUserServerGridVOVpnUserInfoVO GetVpnServerUserListV3(ctx, omadacId, siteId, vpnId).Page(page).PageSize(pageSize).FiltersClientMode(filtersClientMode).SortsMaxConnections(sortsMaxConnections).Execute()

Get user list by VPN server ID



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	filtersClientMode := "filtersClientMode_example" // string | Filter query parameters, support field clientMode. 0: Network Extension Mode; 1: Client. (optional)
	sortsMaxConnections := "sortsMaxConnections_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetVpnServerUserListV3(context.Background(), omadacId, siteId, vpnId).Page(page).PageSize(pageSize).FiltersClientMode(filtersClientMode).SortsMaxConnections(sortsMaxConnections).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnServerUserListV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnServerUserListV3`: OperationResponseVpnUserServerGridVOVpnUserInfoVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnServerUserListV3`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetVpnServerUserListV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **filtersClientMode** | **string** | Filter query parameters, support field clientMode. 0: Network Extension Mode; 1: Client. | 
 **sortsMaxConnections** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseVpnUserServerGridVOVpnUserInfoVO**](OperationResponseVpnUserServerGridVOVpnUserInfoVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnUserList

> OperationResponseVpnUserOpenApiGridVOVpnUserResponse GetVpnUserList(ctx, omadacId, siteId, protocol).Execute()

Get VPN user list without page



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
	protocol := "protocol_example" // string | 0: L2TP or PPTP; 1: OpenVPN; 2: SSL VPN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.GetVpnUserList(context.Background(), omadacId, siteId, protocol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnUserList`: OperationResponseVpnUserOpenApiGridVOVpnUserResponse
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**protocol** | **string** | 0: L2TP or PPTP; 1: OpenVPN; 2: SSL VPN | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseVpnUserOpenApiGridVOVpnUserResponse**](OperationResponseVpnUserOpenApiGridVOVpnUserResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnUserServerList

> OperationResponseListVpnUserServerBriefVO GetVpnUserServerList(ctx, omadacId, siteId).Execute()

Get VPN server list for user



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
	resp, r, err := apiClient.VPNAPI.GetVpnUserServerList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.GetVpnUserServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnUserServerList`: OperationResponseListVpnUserServerBriefVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.GetVpnUserServerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnUserServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListVpnUserServerBriefVO**](OperationResponseListVpnUserServerBriefVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockSslVpnTunnel

> OperationResponseWithoutResult LockSslVpnTunnel(ctx, omadacId, siteId, tunnelId).Execute()

Lock SSL VPN tunnel



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
	tunnelId := "tunnelId_example" // string | SSL VPN tunnel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.LockSslVpnTunnel(context.Background(), omadacId, siteId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.LockSslVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockSslVpnTunnel`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.LockSslVpnTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**tunnelId** | **string** | SSL VPN tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockSslVpnTunnelRequest struct via the builder pattern


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


## ModifyClientStatus

> OperationResponseWithoutResult ModifyClientStatus(ctx, omadacId, siteId, vpnId).VpnStatusVO(vpnStatusVO).Execute()

Modify VPN status



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
	vpnStatusVO := *openapiclient.NewVpnStatusVO(false) // VpnStatusVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyClientStatus(context.Background(), omadacId, siteId, vpnId).VpnStatusVO(vpnStatusVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyClientStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyClientStatus`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyClientStatus`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyClientStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpnStatusVO** | [**VpnStatusVO**](VpnStatusVO.md) |  | 

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIpsecFailover

> OperationResponseWithoutResult ModifyIpsecFailover(ctx, omadacId, siteId, failoverId).IPsecFailoverConfiguration(iPsecFailoverConfiguration).Execute()

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
	iPsecFailoverConfiguration := *openapiclient.NewIPsecFailoverConfiguration([]string{"Candidates_example"}, "Name_example", "Primary_example") // IPsecFailoverConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyIpsecFailover(context.Background(), omadacId, siteId, failoverId).IPsecFailoverConfiguration(iPsecFailoverConfiguration).Execute()
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



 **iPsecFailoverConfiguration** | [**IPsecFailoverConfiguration**](IPsecFailoverConfiguration.md) |  | 

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


## ModifyS2SAutoVpn

> OperationResponseWithoutResult ModifyS2SAutoVpn(ctx, omadacId, siteId, vpnId).VpnSiteToSiteAutoConfigOpenApiVO(vpnSiteToSiteAutoConfigOpenApiVO).Execute()

Modify site-to-site VPN by auto



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
	vpnSiteToSiteAutoConfigOpenApiVO := *openapiclient.NewVpnSiteToSiteAutoConfigOpenApiVO("Name_example", "RemoteSite_example", false, int32(123)) // VpnSiteToSiteAutoConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyS2SAutoVpn(context.Background(), omadacId, siteId, vpnId).VpnSiteToSiteAutoConfigOpenApiVO(vpnSiteToSiteAutoConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyS2SAutoVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyS2SAutoVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyS2SAutoVpn`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyS2SAutoVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpnSiteToSiteAutoConfigOpenApiVO** | [**VpnSiteToSiteAutoConfigOpenApiVO**](VpnSiteToSiteAutoConfigOpenApiVO.md) |  | 

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


## ModifyS2SManualVpn

> OperationResponseWithoutResult ModifyS2SManualVpn(ctx, omadacId, siteId, vpnId).VpnSiteToSiteManualConfigOpenApiVO(vpnSiteToSiteManualConfigOpenApiVO).Execute()

Modify site-to-site VPN by manual



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
	vpnSiteToSiteManualConfigOpenApiVO := *openapiclient.NewVpnSiteToSiteManualConfigOpenApiVO(*openapiclient.NewVpnAdvancedSettingOpenApiVO(), int32(123), "Name_example", "PreSharedKey_example", "PrivateKey_example", "RemoteIp_example", []string{"RemoteSubnet_example"}, int32(123), false, "TunnelIp_example", int32(123), []string{"Wans_example"}) // VpnSiteToSiteManualConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyS2SManualVpn(context.Background(), omadacId, siteId, vpnId).VpnSiteToSiteManualConfigOpenApiVO(vpnSiteToSiteManualConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyS2SManualVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyS2SManualVpn`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyS2SManualVpn`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyS2SManualVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpnSiteToSiteManualConfigOpenApiVO** | [**VpnSiteToSiteManualConfigOpenApiVO**](VpnSiteToSiteManualConfigOpenApiVO.md) |  | 

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyVpnClient

> OperationResponseWithoutResult ModifyVpnClient(ctx, omadacId, siteId, vpnId).VpnClientConfigOpenApiVO(vpnClientConfigOpenApiVO).Execute()

Modify client-to-site VPN client V2



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
	vpnClientConfigOpenApiVO := *openapiclient.NewVpnClientConfigOpenApiVO("Name_example", "RemoteIp_example", "ServerPublicKey_example", int32(123), int32(123), false, *openapiclient.NewVpnCertificateOpenApiVO(), int32(123), []string{"Wans_example"}) // VpnClientConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyVpnClient(context.Background(), omadacId, siteId, vpnId).VpnClientConfigOpenApiVO(vpnClientConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyVpnClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVpnClient`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyVpnClient`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyVpnClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpnClientConfigOpenApiVO** | [**VpnClientConfigOpenApiVO**](VpnClientConfigOpenApiVO.md) |  | 

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


## ModifyVpnServer

> OperationResponseWithoutResult ModifyVpnServer(ctx, omadacId, siteId, vpnId).VpnServerConfigOpenApiVO(vpnServerConfigOpenApiVO).Execute()

Modify client-to-site VPN server V2



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
	vpnServerConfigOpenApiVO := *openapiclient.NewVpnServerConfigOpenApiVO(*openapiclient.NewVpnAdvancedSettingOpenApiVO(), false, *openapiclient.NewLockSettingOpenApiVO(false), int32(123), int32(123), "Name_example", *openapiclient.NewLockSettingOpenApiVO(false), "PreSharedKey_example", "PrivateKey_example", "RemoteIp_example", int32(123), int32(123), false, int32(123), []string{"VpnUserList_example"}, []string{"Wans_example"}) // VpnServerConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyVpnServer(context.Background(), omadacId, siteId, vpnId).VpnServerConfigOpenApiVO(vpnServerConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyVpnServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVpnServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyVpnServer`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyVpnServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpnServerConfigOpenApiVO** | [**VpnServerConfigOpenApiVO**](VpnServerConfigOpenApiVO.md) |  | 

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

[AccessToken](../README.md#accesstoken)

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

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyVpnUserV3

> OperationResponseWithoutResult ModifyVpnUserV3(ctx, omadacId, siteId, userId).VpnUserConfigVO(vpnUserConfigVO).Execute()

Modify VPN user V3



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
	vpnUserConfigVO := *openapiclient.NewVpnUserConfigVO() // VpnUserConfigVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.ModifyVpnUserV3(context.Background(), omadacId, siteId, userId).VpnUserConfigVO(vpnUserConfigVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.ModifyVpnUserV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVpnUserV3`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.ModifyVpnUserV3`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyVpnUserV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpnUserConfigVO** | [**VpnUserConfigVO**](VpnUserConfigVO.md) |  | 

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


## UploadVpnCertificateFile

> OperationResponseVpnCertificateVO UploadVpnCertificateFile(ctx, omadacId, siteId).File(file).Execute()

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
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.UploadVpnCertificateFile(context.Background(), omadacId, siteId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.UploadVpnCertificateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadVpnCertificateFile`: OperationResponseVpnCertificateVO
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

[**OperationResponseVpnCertificateVO**](OperationResponseVpnCertificateVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVpnCertificateFileV2

> OperationResponseVpnCertificateVO UploadVpnCertificateFileV2(ctx, omadacId, siteId).File(file).Execute()

Upload VPN certificate file V2



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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNAPI.UploadVpnCertificateFileV2(context.Background(), omadacId, siteId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNAPI.UploadVpnCertificateFileV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadVpnCertificateFileV2`: OperationResponseVpnCertificateVO
	fmt.Fprintf(os.Stdout, "Response from `VPNAPI.UploadVpnCertificateFileV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadVpnCertificateFileV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**OperationResponseVpnCertificateVO**](OperationResponseVpnCertificateVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


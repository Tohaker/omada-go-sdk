# ProfilesAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPSKsToPPSKProfile**](ProfilesAPI.md#addpskstoppskprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profile/{profileId}/add-psk | Add PSKs to PPSK profile
[**AutoCreatePPSK**](ProfilesAPI.md#autocreateppsk) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profile/generate-psk | Auto Create PSKs
[**CreateApnProfile**](ProfilesAPI.md#createapnprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/profiles/apn | Create a new APN profile
[**CreateApnProfileByMac**](ProfilesAPI.md#createapnprofilebymac) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/profiles/apns | Create a new APN profile by mac
[**CreateGoogleLdapProfile**](ProfilesAPI.md#creategoogleldapprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/profiles/ldap/google | Create a new google LDAP profile
[**CreateGroupProfile**](ProfilesAPI.md#creategroupprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/profiles/groups | Create a new group profile
[**CreateLdapProfile**](ProfilesAPI.md#createldapprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/profiles/ldap | Create a new LDAP profile
[**CreateMdnsProfile**](ProfilesAPI.md#createmdnsprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/profiles/bonjour-service | Create new Bonjour Service
[**CreateOuiProfile**](ProfilesAPI.md#createouiprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/oui-profiles | Create OUI profile
[**CreatePPSKProfile**](ProfilesAPI.md#createppskprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profile | Create PPSK profile
[**CreatePPSKProfileV2**](ProfilesAPI.md#createppskprofilev2) | **Post** /openapi/v2/{omadacId}/sites/{siteId}/ppsk-profile | Create PPSK profile V2
[**CreateRadiusProfile**](ProfilesAPI.md#createradiusprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/profiles/radius | Create a new Radius profile
[**CreateRadiusUser**](ProfilesAPI.md#createradiususer) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/profiles/radius-server/users | Create a new Build-in Radius profile user
[**CreateRateLimitProfile**](ProfilesAPI.md#createratelimitprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/rate-limit-profiles | Create rate limit profile
[**CreateServiceType**](ProfilesAPI.md#createservicetype) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/profiles/service-type | Create new Gateway QoS Service
[**CreateTimeRangeProfile**](ProfilesAPI.md#createtimerangeprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/time-range-profiles | Create time range profile
[**DeleteAllPPSKProfiles**](ProfilesAPI.md#deleteallppskprofiles) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profiles | Delete All PPSK profiles
[**DeleteApnProfile**](ProfilesAPI.md#deleteapnprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/profiles/apn/{profileId} | Delete an exist APN profile
[**DeleteGoogleLdapProfile**](ProfilesAPI.md#deletegoogleldapprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/profiles/ldap/google/{profileId} | Delete an exist google LDAP profile
[**DeleteGroupProfile**](ProfilesAPI.md#deletegroupprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/profiles/groups/{groupType}/{groupId} | Delete an exist group profile
[**DeleteLdapProfile**](ProfilesAPI.md#deleteldapprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/profiles/ldap/{ldapProfileId} | Delete an exist LDAP profile
[**DeleteMdnsProfile**](ProfilesAPI.md#deletemdnsprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/profiles/bonjour-service/{id} | Delete an existing Bonjour Service
[**DeleteOuiProfile**](ProfilesAPI.md#deleteouiprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/oui-profiles/{ouiId} | Delete OUI Profile
[**DeletePPSKProfile**](ProfilesAPI.md#deleteppskprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profile/{profileId} | Delete PPSK profile
[**DeletePSKsToPPSKProfile**](ProfilesAPI.md#deletepskstoppskprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profile/{profileId}/delete-psk | Delete PSKs to PPSK profile
[**DeleteRadiusProfile**](ProfilesAPI.md#deleteradiusprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/profiles/radius/{radiusProfileId} | Delete an exist Radius profile
[**DeleteRadiusUser**](ProfilesAPI.md#deleteradiususer) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/profiles/radius-server/users/{userId} | Delete an exist Build-in Radius profile user
[**DeleteRateLimitProfile**](ProfilesAPI.md#deleteratelimitprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/rate-limit-profile/{profileId} | Delete rate limit profile
[**DeleteServiceType**](ProfilesAPI.md#deleteservicetype) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/profiles/service-type/{id} | Delete an existing Gateway QoS Service
[**DeleteTimeRangeProfile**](ProfilesAPI.md#deletetimerangeprofile) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/time-range-profile/{profileId} | Delete time range profile
[**GetApnProfileList**](ProfilesAPI.md#getapnprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/apn | Get APN profile list
[**GetApnProfilesForIppt**](ProfilesAPI.md#getapnprofilesforippt) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/profiles/apns | Get APN profile list by mac for ippt
[**GetGoogleLdapProfileList**](ProfilesAPI.md#getgoogleldapprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/ldap/google | Get google LDAP profile list
[**GetGroupProfiles**](ProfilesAPI.md#getgroupprofiles) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/groups | Get group profile list
[**GetGroupProfilesByType**](ProfilesAPI.md#getgroupprofilesbytype) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/groups/{groupType} | Get group profile list by type
[**GetLdapProfileList**](ProfilesAPI.md#getldapprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/ldap | Get LDAP profile list
[**GetOuiProfileFullList**](ProfilesAPI.md#getouiprofilefulllist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/oui-profiles/all | Get OUI profile summary list
[**GetOuiProfileList**](ProfilesAPI.md#getouiprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/oui-profiles | Get OUI profile list
[**GetPPSKProfileDetail**](ProfilesAPI.md#getppskprofiledetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profile/{profileId} | Get PPSK profile detail
[**GetPPSKProfiles**](ProfilesAPI.md#getppskprofiles) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profiles | Get PPSK profiles list
[**GetRadiusProfileList**](ProfilesAPI.md#getradiusprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/radius | Get Radius profile list
[**GetRadiusUserList**](ProfilesAPI.md#getradiususerlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/radius-server/users | Get Build-in Radius profile user list
[**GetRateLimitProfileList**](ProfilesAPI.md#getratelimitprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rate-limit-profiles | Get rate limit profile list
[**GetServiceTypeSummary**](ProfilesAPI.md#getservicetypesummary) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/service-type-summary | Get all Gateway QoS Service&#39;s ID and name info
[**GetTimeRangeList**](ProfilesAPI.md#gettimerangelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/time-range-profiles | Get time range profile list
[**ListMdnsProfile**](ProfilesAPI.md#listmdnsprofile) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/bonjour-service | Get Bonjour Service list
[**ListRateLimitForHotspot**](ProfilesAPI.md#listratelimitforhotspot) | **Get** /openapi/v1/{omadacId}/hotspot/sites/{siteId}/setting/profiles/rateLimits | get the rate limit list
[**ListServiceType**](ProfilesAPI.md#listservicetype) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/profiles/service-type | Get Gateway QoS Service list
[**ModifyApnProfile**](ProfilesAPI.md#modifyapnprofile) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/profiles/apn/{profileId} | Modify an exist APN profile
[**ModifyGoogleLdapProfile**](ProfilesAPI.md#modifygoogleldapprofile) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/profiles/ldap/google/{profileId} | Modify an exist google LDAP profile
[**ModifyGroupProfile**](ProfilesAPI.md#modifygroupprofile) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/profiles/groups/{groupType}/{groupId} | Modify an exist group profile
[**ModifyLdapProfile**](ProfilesAPI.md#modifyldapprofile) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/profiles/ldap/{ldapProfileId} | Modify an exist LDAP profile
[**ModifyMdnsProfile**](ProfilesAPI.md#modifymdnsprofile) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/profiles/bonjour-service/{id} | Modify an existing Bonjour Service
[**ModifyOuiProfile**](ProfilesAPI.md#modifyouiprofile) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/oui-profiles/{ouiId} | Modify OUI Profile
[**ModifyPPSKProfile**](ProfilesAPI.md#modifyppskprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ppsk-profile/{profileId} | Modify PPSK profile
[**ModifyRadiusProfile**](ProfilesAPI.md#modifyradiusprofile) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/profiles/radius/{radiusProfileId} | Modify an exist Radius profile
[**ModifyRadiusUser**](ProfilesAPI.md#modifyradiususer) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/profiles/radius-server/users/{userId} | Modify an exist Build-in Radius profile user
[**ModifyRateLimitProfile**](ProfilesAPI.md#modifyratelimitprofile) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/rate-limit-profile/{profileId} | Modify rate limit profile
[**ModifyServiceType**](ProfilesAPI.md#modifyservicetype) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/profiles/service-type/{id} | Modify an existing Gateway QoS Service
[**ModifyTimeRangeProfile**](ProfilesAPI.md#modifytimerangeprofile) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/time-range-profile/{profileId} | Modify time range profile



## AddPSKsToPPSKProfile

> OperationResponseWithoutResult AddPSKsToPPSKProfile(ctx, omadacId, siteId, profileId).AddPSKsOpenApiVO(addPSKsOpenApiVO).Execute()

Add PSKs to PPSK profile



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
	profileId := "profileId_example" // string | PPSK profile Id
	addPSKsOpenApiVO := *openapiclient.NewAddPSKsOpenApiVO([]openapiclient.PpskSetting{*openapiclient.NewPpskSetting("Name_example", "Psk_example")}) // AddPSKsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.AddPSKsToPPSKProfile(context.Background(), omadacId, siteId, profileId).AddPSKsOpenApiVO(addPSKsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.AddPSKsToPPSKProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPSKsToPPSKProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.AddPSKsToPPSKProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | PPSK profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPSKsToPPSKProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addPSKsOpenApiVO** | [**AddPSKsOpenApiVO**](AddPSKsOpenApiVO.md) |  | 

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


## AutoCreatePPSK

> []PSK AutoCreatePPSK(ctx, omadacId, siteId).PpskAutoCreateSetting(ppskAutoCreateSetting).Execute()

Auto Create PSKs



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
	ppskAutoCreateSetting := *openapiclient.NewPpskAutoCreateSetting(int32(123), int32(123), "Prefix_example") // PpskAutoCreateSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.AutoCreatePPSK(context.Background(), omadacId, siteId).PpskAutoCreateSetting(ppskAutoCreateSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.AutoCreatePPSK``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreatePPSK`: []PSK
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.AutoCreatePPSK`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreatePPSKRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ppskAutoCreateSetting** | [**PpskAutoCreateSetting**](PpskAutoCreateSetting.md) |  | 

### Return type

[**[]PSK**](PSK.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApnProfile

> OperationResponseResIdOpenApiVO CreateApnProfile(ctx, omadacId, siteId).ApnProfileConfig(apnProfileConfig).Execute()

Create a new APN profile



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
	apnProfileConfig := *openapiclient.NewApnProfileConfig(int32(123), int32(123), "Name_example", int32(123)) // ApnProfileConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateApnProfile(context.Background(), omadacId, siteId).ApnProfileConfig(apnProfileConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateApnProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApnProfile`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateApnProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApnProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apnProfileConfig** | [**ApnProfileConfig**](ApnProfileConfig.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApnProfileByMac

> OperationResponseResIdOpenApiVO CreateApnProfileByMac(ctx, omadacId, siteId, gatewayMac).ApnProfileConfig(apnProfileConfig).Execute()

Create a new APN profile by mac



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
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	apnProfileConfig := *openapiclient.NewApnProfileConfig(int32(123), int32(123), "Name_example", int32(123)) // ApnProfileConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateApnProfileByMac(context.Background(), omadacId, siteId, gatewayMac).ApnProfileConfig(apnProfileConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateApnProfileByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApnProfileByMac`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateApnProfileByMac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApnProfileByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apnProfileConfig** | [**ApnProfileConfig**](ApnProfileConfig.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGoogleLdapProfile

> OperationResponseResIdOpenApiVO CreateGoogleLdapProfile(ctx, omadacId, siteId).Config(config).File(file).Execute()

Create a new google LDAP profile



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
	config := *openapiclient.NewCreateGoogleLdapProfileOpenApiVO("BaseDn_example", int32(123), "Cn_example", "Name_example", int32(123), "Server_example", false) // CreateGoogleLdapProfileOpenApiVO | 
	file := os.NewFile(1234, "some_file") // *os.File | Google LDAP certificate file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateGoogleLdapProfile(context.Background(), omadacId, siteId).Config(config).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateGoogleLdapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGoogleLdapProfile`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateGoogleLdapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGoogleLdapProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **config** | [**CreateGoogleLdapProfileOpenApiVO**](CreateGoogleLdapProfileOpenApiVO.md) |  | 
 **file** | ***os.File** | Google LDAP certificate file. | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupProfile

> OperationResponseResIdOpenApiVO CreateGroupProfile(ctx, omadacId, siteId).CreateGroupOpenApiVO(createGroupOpenApiVO).Execute()

Create a new group profile



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
	createGroupOpenApiVO := *openapiclient.NewCreateGroupOpenApiVO("Name_example", int32(123)) // CreateGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateGroupProfile(context.Background(), omadacId, siteId).CreateGroupOpenApiVO(createGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateGroupProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupProfile`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateGroupProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createGroupOpenApiVO** | [**CreateGroupOpenApiVO**](CreateGroupOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLdapProfile

> OperationResponseResIdOpenApiVO CreateLdapProfile(ctx, omadacId, siteId).CreateLdapProfileOpenApiVO(createLdapProfileOpenApiVO).Execute()

Create a new LDAP profile



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
	createLdapProfileOpenApiVO := *openapiclient.NewCreateLdapProfileOpenApiVO("BaseDn_example", int32(123), "Cn_example", int32(123), "Name_example", "Server_example", false, false) // CreateLdapProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateLdapProfile(context.Background(), omadacId, siteId).CreateLdapProfileOpenApiVO(createLdapProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateLdapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLdapProfile`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateLdapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLdapProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createLdapProfileOpenApiVO** | [**CreateLdapProfileOpenApiVO**](CreateLdapProfileOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMdnsProfile

> OperationResponseResponseIdVO CreateMdnsProfile(ctx, omadacId, siteId).BonjourServiceOpenApiVO(bonjourServiceOpenApiVO).Execute()

Create new Bonjour Service



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
	bonjourServiceOpenApiVO := *openapiclient.NewBonjourServiceOpenApiVO("Name_example", []string{"ServiceIds_example"}) // BonjourServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateMdnsProfile(context.Background(), omadacId, siteId).BonjourServiceOpenApiVO(bonjourServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateMdnsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMdnsProfile`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateMdnsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMdnsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bonjourServiceOpenApiVO** | [**BonjourServiceOpenApiVO**](BonjourServiceOpenApiVO.md) |  | 

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


## CreateOuiProfile

> OperationResponseWithoutResult CreateOuiProfile(ctx, omadacId, siteId).OuiProfileOpenApiVO(ouiProfileOpenApiVO).Execute()

Create OUI profile



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
	ouiProfileOpenApiVO := *openapiclient.NewOuiProfileOpenApiVO("Name_example") // OuiProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateOuiProfile(context.Background(), omadacId, siteId).OuiProfileOpenApiVO(ouiProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateOuiProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOuiProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateOuiProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOuiProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ouiProfileOpenApiVO** | [**OuiProfileOpenApiVO**](OuiProfileOpenApiVO.md) |  | 

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


## CreatePPSKProfile

> OperationResponseWithoutResult CreatePPSKProfile(ctx, omadacId, siteId).PpskProfile(ppskProfile).Execute()

Create PPSK profile



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
	ppskProfile := *openapiclient.NewPpskProfile([]openapiclient.PpskSetting{*openapiclient.NewPpskSetting("Name_example", "Psk_example")}, "ProfileName_example") // PpskProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreatePPSKProfile(context.Background(), omadacId, siteId).PpskProfile(ppskProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreatePPSKProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePPSKProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreatePPSKProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePPSKProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ppskProfile** | [**PpskProfile**](PpskProfile.md) |  | 

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


## CreatePPSKProfileV2

> OperationResponseWithoutResult CreatePPSKProfileV2(ctx, omadacId, siteId).PpskProfileV2(ppskProfileV2).Execute()

Create PPSK profile V2



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
	ppskProfileV2 := *openapiclient.NewPpskProfileV2(false, "ProfileName_example") // PpskProfileV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreatePPSKProfileV2(context.Background(), omadacId, siteId).PpskProfileV2(ppskProfileV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreatePPSKProfileV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePPSKProfileV2`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreatePPSKProfileV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePPSKProfileV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ppskProfileV2** | [**PpskProfileV2**](PpskProfileV2.md) |  | 

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


## CreateRadiusProfile

> OperationResponseResIdOpenApiVO CreateRadiusProfile(ctx, omadacId, siteId).CreateRadiusProfileOpenApiVO(createRadiusProfileOpenApiVO).Execute()

Create a new Radius profile



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
	createRadiusProfileOpenApiVO := *openapiclient.NewCreateRadiusProfileOpenApiVO([]openapiclient.RadiusAuthServerOpenApiVO{*openapiclient.NewRadiusAuthServerOpenApiVO(int32(123), "RadiusPwd_example", "RadiusServerIp_example")}, "Name_example", false, false) // CreateRadiusProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateRadiusProfile(context.Background(), omadacId, siteId).CreateRadiusProfileOpenApiVO(createRadiusProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateRadiusProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRadiusProfile`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateRadiusProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRadiusProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRadiusProfileOpenApiVO** | [**CreateRadiusProfileOpenApiVO**](CreateRadiusProfileOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRadiusUser

> OperationResponseResIdOpenApiVO CreateRadiusUser(ctx, omadacId, siteId).CreateRadiusUserOpenApiVO(createRadiusUserOpenApiVO).Execute()

Create a new Build-in Radius profile user



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
	createRadiusUserOpenApiVO := *openapiclient.NewCreateRadiusUserOpenApiVO(int32(123)) // CreateRadiusUserOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateRadiusUser(context.Background(), omadacId, siteId).CreateRadiusUserOpenApiVO(createRadiusUserOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateRadiusUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRadiusUser`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateRadiusUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRadiusUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRadiusUserOpenApiVO** | [**CreateRadiusUserOpenApiVO**](CreateRadiusUserOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRateLimitProfile

> OperationResponseWithoutResult CreateRateLimitProfile(ctx, omadacId, siteId).CreateRateLimitProfileOpenApiVO(createRateLimitProfileOpenApiVO).Execute()

Create rate limit profile



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
	createRateLimitProfileOpenApiVO := *openapiclient.NewCreateRateLimitProfileOpenApiVO(false, "Name_example", false) // CreateRateLimitProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateRateLimitProfile(context.Background(), omadacId, siteId).CreateRateLimitProfileOpenApiVO(createRateLimitProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateRateLimitProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRateLimitProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateRateLimitProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRateLimitProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRateLimitProfileOpenApiVO** | [**CreateRateLimitProfileOpenApiVO**](CreateRateLimitProfileOpenApiVO.md) |  | 

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


## CreateServiceType

> OperationResponseResponseIdVO CreateServiceType(ctx, omadacId, siteId).GatewayQosServiceOpenApiVO(gatewayQosServiceOpenApiVO).Execute()

Create new Gateway QoS Service



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
	gatewayQosServiceOpenApiVO := *openapiclient.NewGatewayQosServiceOpenApiVO("Name_example", int32(123)) // GatewayQosServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateServiceType(context.Background(), omadacId, siteId).GatewayQosServiceOpenApiVO(gatewayQosServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateServiceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceType`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateServiceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayQosServiceOpenApiVO** | [**GatewayQosServiceOpenApiVO**](GatewayQosServiceOpenApiVO.md) |  | 

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


## CreateTimeRangeProfile

> OperationResponseWithoutResult CreateTimeRangeProfile(ctx, omadacId, siteId).CreateTimeRangeProfileOpenApiVO(createTimeRangeProfileOpenApiVO).Execute()

Create time range profile



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
	createTimeRangeProfileOpenApiVO := *openapiclient.NewCreateTimeRangeProfileOpenApiVO(int32(123), "Name_example", []openapiclient.ScheduleTimeOpenApiVO{*openapiclient.NewScheduleTimeOpenApiVO(int32(123), int32(123), int32(123), int32(123), int32(123))}) // CreateTimeRangeProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateTimeRangeProfile(context.Background(), omadacId, siteId).CreateTimeRangeProfileOpenApiVO(createTimeRangeProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateTimeRangeProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTimeRangeProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateTimeRangeProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTimeRangeProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTimeRangeProfileOpenApiVO** | [**CreateTimeRangeProfileOpenApiVO**](CreateTimeRangeProfileOpenApiVO.md) |  | 

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


## DeleteAllPPSKProfiles

> OperationResponseWithoutResult DeleteAllPPSKProfiles(ctx, omadacId, siteId).Execute()

Delete All PPSK profiles



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
	resp, r, err := apiClient.ProfilesAPI.DeleteAllPPSKProfiles(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteAllPPSKProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllPPSKProfiles`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteAllPPSKProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllPPSKProfilesRequest struct via the builder pattern


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


## DeleteApnProfile

> OperationResponseWithoutResult DeleteApnProfile(ctx, omadacId, siteId, profileId).Execute()

Delete an exist APN profile



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
	profileId := "profileId_example" // string | APN profile ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteApnProfile(context.Background(), omadacId, siteId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteApnProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApnProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteApnProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | APN profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApnProfileRequest struct via the builder pattern


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


## DeleteGoogleLdapProfile

> OperationResponseWithoutResult DeleteGoogleLdapProfile(ctx, omadacId, siteId, profileId).Execute()

Delete an exist google LDAP profile



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
	profileId := "profileId_example" // string | google LDAP profile ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteGoogleLdapProfile(context.Background(), omadacId, siteId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteGoogleLdapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGoogleLdapProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteGoogleLdapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | google LDAP profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGoogleLdapProfileRequest struct via the builder pattern


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


## DeleteGroupProfile

> OperationResponseWithoutResult DeleteGroupProfile(ctx, omadacId, siteId, groupId, groupType).Execute()

Delete an exist group profile



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
	groupId := "groupId_example" // string | Group profile id.
	groupType := "groupType_example" // string | groupType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteGroupProfile(context.Background(), omadacId, siteId, groupId, groupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteGroupProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteGroupProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**groupId** | **string** | Group profile id. | 
**groupType** | **string** | groupType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupProfileRequest struct via the builder pattern


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


## DeleteLdapProfile

> OperationResponseWithoutResult DeleteLdapProfile(ctx, omadacId, siteId, ldapProfileId).Execute()

Delete an exist LDAP profile



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
	ldapProfileId := "ldapProfileId_example" // string | LDAP profile ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteLdapProfile(context.Background(), omadacId, siteId, ldapProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteLdapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLdapProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteLdapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ldapProfileId** | **string** | LDAP profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLdapProfileRequest struct via the builder pattern


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


## DeleteMdnsProfile

> OperationResponseWithoutResult DeleteMdnsProfile(ctx, omadacId, siteId, id).Execute()

Delete an existing Bonjour Service



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
	id := "id_example" // string | Bonjour Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteMdnsProfile(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteMdnsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMdnsProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteMdnsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Bonjour Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMdnsProfileRequest struct via the builder pattern


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


## DeleteOuiProfile

> OperationResponseWithoutResult DeleteOuiProfile(ctx, omadacId, siteId, ouiId).Execute()

Delete OUI Profile



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
	ouiId := "ouiId_example" // string | OUI ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteOuiProfile(context.Background(), omadacId, siteId, ouiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteOuiProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOuiProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteOuiProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ouiId** | **string** | OUI ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOuiProfileRequest struct via the builder pattern


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


## DeletePPSKProfile

> OperationResponseWithoutResult DeletePPSKProfile(ctx, omadacId, siteId, profileId).Execute()

Delete PPSK profile



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
	profileId := "profileId_example" // string | PPSK profile Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeletePPSKProfile(context.Background(), omadacId, siteId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeletePPSKProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePPSKProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeletePPSKProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | PPSK profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePPSKProfileRequest struct via the builder pattern


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


## DeletePSKsToPPSKProfile

> OperationResponseWithoutResult DeletePSKsToPPSKProfile(ctx, omadacId, siteId, profileId).DeletePSKsOpenApiVO(deletePSKsOpenApiVO).Execute()

Delete PSKs to PPSK profile



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
	profileId := "profileId_example" // string | PPSK profile Id
	deletePSKsOpenApiVO := *openapiclient.NewDeletePSKsOpenApiVO([]string{"PpskNameList_example"}) // DeletePSKsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeletePSKsToPPSKProfile(context.Background(), omadacId, siteId, profileId).DeletePSKsOpenApiVO(deletePSKsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeletePSKsToPPSKProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePSKsToPPSKProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeletePSKsToPPSKProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | PPSK profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePSKsToPPSKProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deletePSKsOpenApiVO** | [**DeletePSKsOpenApiVO**](DeletePSKsOpenApiVO.md) |  | 

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


## DeleteRadiusProfile

> OperationResponseWithoutResult DeleteRadiusProfile(ctx, omadacId, siteId, radiusProfileId).Execute()

Delete an exist Radius profile



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
	radiusProfileId := "radiusProfileId_example" // string | Radius profile Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteRadiusProfile(context.Background(), omadacId, siteId, radiusProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteRadiusProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRadiusProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteRadiusProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**radiusProfileId** | **string** | Radius profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRadiusProfileRequest struct via the builder pattern


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


## DeleteRadiusUser

> OperationResponseWithoutResult DeleteRadiusUser(ctx, omadacId, siteId, userId).Execute()

Delete an exist Build-in Radius profile user



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
	userId := "userId_example" // string | Build-in Radius profile user Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteRadiusUser(context.Background(), omadacId, siteId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteRadiusUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRadiusUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteRadiusUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**userId** | **string** | Build-in Radius profile user Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRadiusUserRequest struct via the builder pattern


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


## DeleteRateLimitProfile

> OperationResponseWithoutResult DeleteRateLimitProfile(ctx, omadacId, siteId, profileId).Execute()

Delete rate limit profile



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
	profileId := "profileId_example" // string | Profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteRateLimitProfile(context.Background(), omadacId, siteId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteRateLimitProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRateLimitProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteRateLimitProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRateLimitProfileRequest struct via the builder pattern


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


## DeleteServiceType

> OperationResponseWithoutResult DeleteServiceType(ctx, omadacId, siteId, id).Execute()

Delete an existing Gateway QoS Service



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
	id := "id_example" // string | Gateway QoS Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteServiceType(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteServiceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceType`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteServiceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Gateway QoS Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceTypeRequest struct via the builder pattern


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


## DeleteTimeRangeProfile

> OperationResponseWithoutResult DeleteTimeRangeProfile(ctx, omadacId, siteId, profileId).Execute()

Delete time range profile



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
	profileId := "profileId_example" // string | Profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteTimeRangeProfile(context.Background(), omadacId, siteId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteTimeRangeProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTimeRangeProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteTimeRangeProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeRangeProfileRequest struct via the builder pattern


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


## GetApnProfileList

> OperationResponseListApnProfile GetApnProfileList(ctx, omadacId, siteId).Execute()

Get APN profile list



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
	resp, r, err := apiClient.ProfilesAPI.GetApnProfileList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetApnProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApnProfileList`: OperationResponseListApnProfile
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetApnProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApnProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListApnProfile**](OperationResponseListApnProfile.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApnProfilesForIppt

> OperationResponseApnProfileData GetApnProfilesForIppt(ctx, omadacId, siteId, gatewayMac).Execute()

Get APN profile list by mac for ippt



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
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetApnProfilesForIppt(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetApnProfilesForIppt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApnProfilesForIppt`: OperationResponseApnProfileData
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetApnProfilesForIppt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApnProfilesForIpptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseApnProfileData**](OperationResponseApnProfileData.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoogleLdapProfileList

> OperationResponseListLdapProfileOpenApiVO GetGoogleLdapProfileList(ctx, omadacId, siteId).Execute()

Get google LDAP profile list



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
	resp, r, err := apiClient.ProfilesAPI.GetGoogleLdapProfileList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetGoogleLdapProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleLdapProfileList`: OperationResponseListLdapProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetGoogleLdapProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleLdapProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListLdapProfileOpenApiVO**](OperationResponseListLdapProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupProfiles

> OperationResponseListGroupOpenApiVO GetGroupProfiles(ctx, omadacId, siteId).Execute()

Get group profile list



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
	resp, r, err := apiClient.ProfilesAPI.GetGroupProfiles(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetGroupProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupProfiles`: OperationResponseListGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetGroupProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListGroupOpenApiVO**](OperationResponseListGroupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupProfilesByType

> OperationResponseListGroupOpenApiVO GetGroupProfilesByType(ctx, omadacId, siteId, groupType).Execute()

Get group profile list by type



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
	groupType := "groupType_example" // string | Type of group profile, 0:IP Group; 1:IP Port Group; 2：Mac Group; 3:IPv6 Group; 4:IPv6 Port Group; 5:Country Group; 7:Domain Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetGroupProfilesByType(context.Background(), omadacId, siteId, groupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetGroupProfilesByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupProfilesByType`: OperationResponseListGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetGroupProfilesByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**groupType** | **string** | Type of group profile, 0:IP Group; 1:IP Port Group; 2：Mac Group; 3:IPv6 Group; 4:IPv6 Port Group; 5:Country Group; 7:Domain Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupProfilesByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListGroupOpenApiVO**](OperationResponseListGroupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapProfileList

> OperationResponseListLdapProfileOpenApiVO GetLdapProfileList(ctx, omadacId, siteId).Execute()

Get LDAP profile list



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
	resp, r, err := apiClient.ProfilesAPI.GetLdapProfileList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetLdapProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapProfileList`: OperationResponseListLdapProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetLdapProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListLdapProfileOpenApiVO**](OperationResponseListLdapProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOuiProfileFullList

> OperationResponseListOuiProfileSummaryOpenApiVO GetOuiProfileFullList(ctx, omadacId, siteId).Execute()

Get OUI profile summary list



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
	resp, r, err := apiClient.ProfilesAPI.GetOuiProfileFullList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetOuiProfileFullList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOuiProfileFullList`: OperationResponseListOuiProfileSummaryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetOuiProfileFullList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOuiProfileFullListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListOuiProfileSummaryOpenApiVO**](OperationResponseListOuiProfileSummaryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOuiProfileList

> OperationResponseOuiGridVOOuiProfileQueryOpenApiVO GetOuiProfileList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get OUI profile list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetOuiProfileList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetOuiProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOuiProfileList`: OperationResponseOuiGridVOOuiProfileQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetOuiProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOuiProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseOuiGridVOOuiProfileQueryOpenApiVO**](OperationResponseOuiGridVOOuiProfileQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPPSKProfileDetail

> OperationResponsePPSKProfileVO GetPPSKProfileDetail(ctx, omadacId, siteId, profileId).Execute()

Get PPSK profile detail



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
	profileId := "profileId_example" // string | PPSK profile Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetPPSKProfileDetail(context.Background(), omadacId, siteId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetPPSKProfileDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPPSKProfileDetail`: OperationResponsePPSKProfileVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetPPSKProfileDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | PPSK profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPPSKProfileDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponsePPSKProfileVO**](OperationResponsePPSKProfileVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPPSKProfiles

> OperationResponseListPpskProfileBriefInfo GetPPSKProfiles(ctx, omadacId, siteId).Execute()

Get PPSK profiles list



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
	resp, r, err := apiClient.ProfilesAPI.GetPPSKProfiles(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetPPSKProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPPSKProfiles`: OperationResponseListPpskProfileBriefInfo
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetPPSKProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPPSKProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListPpskProfileBriefInfo**](OperationResponseListPpskProfileBriefInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadiusProfileList

> OperationResponseListRadiusProfileOpenApiVO GetRadiusProfileList(ctx, omadacId, siteId).Execute()

Get Radius profile list



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
	resp, r, err := apiClient.ProfilesAPI.GetRadiusProfileList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetRadiusProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadiusProfileList`: OperationResponseListRadiusProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetRadiusProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadiusProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRadiusProfileOpenApiVO**](OperationResponseListRadiusProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadiusUserList

> OperationResponseGridVORadiusUserOpenApiVO GetRadiusUserList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsUsername(sortsUsername).Execute()

Get Build-in Radius profile user list



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
	sortsUsername := "sortsUsername_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetRadiusUserList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsUsername(sortsUsername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetRadiusUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadiusUserList`: OperationResponseGridVORadiusUserOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetRadiusUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadiusUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsUsername** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseGridVORadiusUserOpenApiVO**](OperationResponseGridVORadiusUserOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateLimitProfileList

> OperationResponseListRateLimitProfileOpenApiVO GetRateLimitProfileList(ctx, omadacId, siteId).Execute()

Get rate limit profile list



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
	resp, r, err := apiClient.ProfilesAPI.GetRateLimitProfileList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetRateLimitProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateLimitProfileList`: OperationResponseListRateLimitProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetRateLimitProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRateLimitProfileOpenApiVO**](OperationResponseListRateLimitProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTypeSummary

> OperationResponseResponseDataVOGatewayQosServiceSummaryOpenApiVO GetServiceTypeSummary(ctx, omadacId, siteId).Execute()

Get all Gateway QoS Service's ID and name info



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
	resp, r, err := apiClient.ProfilesAPI.GetServiceTypeSummary(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetServiceTypeSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTypeSummary`: OperationResponseResponseDataVOGatewayQosServiceSummaryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetServiceTypeSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTypeSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseResponseDataVOGatewayQosServiceSummaryOpenApiVO**](OperationResponseResponseDataVOGatewayQosServiceSummaryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeRangeList

> OperationResponseListTimeRangeProfileOpenApiVO GetTimeRangeList(ctx, omadacId, siteId).Execute()

Get time range profile list



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
	resp, r, err := apiClient.ProfilesAPI.GetTimeRangeList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetTimeRangeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeRangeList`: OperationResponseListTimeRangeProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetTimeRangeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeRangeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListTimeRangeProfileOpenApiVO**](OperationResponseListTimeRangeProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMdnsProfile

> OperationResponseResponseDataVOBonjourServiceDetailOpenApiVO ListMdnsProfile(ctx, omadacId, siteId).Execute()

Get Bonjour Service list



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
	resp, r, err := apiClient.ProfilesAPI.ListMdnsProfile(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ListMdnsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMdnsProfile`: OperationResponseResponseDataVOBonjourServiceDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ListMdnsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMdnsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseResponseDataVOBonjourServiceDetailOpenApiVO**](OperationResponseResponseDataVOBonjourServiceDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRateLimitForHotspot

> OperationResponseListRateLimitProfileVO ListRateLimitForHotspot(ctx, omadacId, siteId).Execute()

get the rate limit list



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
	resp, r, err := apiClient.ProfilesAPI.ListRateLimitForHotspot(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ListRateLimitForHotspot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRateLimitForHotspot`: OperationResponseListRateLimitProfileVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ListRateLimitForHotspot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRateLimitForHotspotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRateLimitProfileVO**](OperationResponseListRateLimitProfileVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceType

> OperationResponseGridVOGatewayQosServiceDetailOpenApiVO ListServiceType(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Gateway QoS Service list



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
	resp, r, err := apiClient.ProfilesAPI.ListServiceType(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ListServiceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceType`: OperationResponseGridVOGatewayQosServiceDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ListServiceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOGatewayQosServiceDetailOpenApiVO**](OperationResponseGridVOGatewayQosServiceDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyApnProfile

> OperationResponseWithoutResult ModifyApnProfile(ctx, omadacId, siteId, profileId).ApnProfileConfig(apnProfileConfig).Execute()

Modify an exist APN profile



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
	profileId := "profileId_example" // string | APN profile ID.
	apnProfileConfig := *openapiclient.NewApnProfileConfig(int32(123), int32(123), "Name_example", int32(123)) // ApnProfileConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyApnProfile(context.Background(), omadacId, siteId, profileId).ApnProfileConfig(apnProfileConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyApnProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApnProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyApnProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | APN profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyApnProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apnProfileConfig** | [**ApnProfileConfig**](ApnProfileConfig.md) |  | 

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


## ModifyGoogleLdapProfile

> OperationResponseWithoutResult ModifyGoogleLdapProfile(ctx, omadacId, siteId, profileId).Config(config).File(file).Execute()

Modify an exist google LDAP profile



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
	profileId := "profileId_example" // string | LDAP profile ID.
	config := *openapiclient.NewCreateGoogleLdapProfileOpenApiVO("BaseDn_example", int32(123), "Cn_example", "Name_example", int32(123), "Server_example", false) // CreateGoogleLdapProfileOpenApiVO | 
	file := os.NewFile(1234, "some_file") // *os.File | Google LDAP certificate file. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyGoogleLdapProfile(context.Background(), omadacId, siteId, profileId).Config(config).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyGoogleLdapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGoogleLdapProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyGoogleLdapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | LDAP profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyGoogleLdapProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **config** | [**CreateGoogleLdapProfileOpenApiVO**](CreateGoogleLdapProfileOpenApiVO.md) |  | 
 **file** | ***os.File** | Google LDAP certificate file. | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGroupProfile

> OperationResponseWithoutResult ModifyGroupProfile(ctx, omadacId, siteId, groupType, groupId).CreateGroupOpenApiVO(createGroupOpenApiVO).Execute()

Modify an exist group profile



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
	groupType := "groupType_example" // string | Type of group profile, 0:IP Group; 1:IP Port Group; 2：Mac Group; 3:IPv6 Group; 4:IPv6 Port Group; 5:Country Group; 7:Domain Group.
	groupId := "groupId_example" // string | Group profile id.
	createGroupOpenApiVO := *openapiclient.NewCreateGroupOpenApiVO("Name_example", int32(123)) // CreateGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyGroupProfile(context.Background(), omadacId, siteId, groupType, groupId).CreateGroupOpenApiVO(createGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyGroupProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGroupProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyGroupProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**groupType** | **string** | Type of group profile, 0:IP Group; 1:IP Port Group; 2：Mac Group; 3:IPv6 Group; 4:IPv6 Port Group; 5:Country Group; 7:Domain Group. | 
**groupId** | **string** | Group profile id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyGroupProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **createGroupOpenApiVO** | [**CreateGroupOpenApiVO**](CreateGroupOpenApiVO.md) |  | 

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


## ModifyLdapProfile

> OperationResponseWithoutResult ModifyLdapProfile(ctx, omadacId, siteId, ldapProfileId).CreateLdapProfileOpenApiVO(createLdapProfileOpenApiVO).Execute()

Modify an exist LDAP profile



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
	ldapProfileId := "ldapProfileId_example" // string | LDAP profile ID.
	createLdapProfileOpenApiVO := *openapiclient.NewCreateLdapProfileOpenApiVO("BaseDn_example", int32(123), "Cn_example", int32(123), "Name_example", "Server_example", false, false) // CreateLdapProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyLdapProfile(context.Background(), omadacId, siteId, ldapProfileId).CreateLdapProfileOpenApiVO(createLdapProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyLdapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLdapProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyLdapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ldapProfileId** | **string** | LDAP profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLdapProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLdapProfileOpenApiVO** | [**CreateLdapProfileOpenApiVO**](CreateLdapProfileOpenApiVO.md) |  | 

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


## ModifyMdnsProfile

> OperationResponseWithoutResult ModifyMdnsProfile(ctx, omadacId, siteId, id).BonjourServiceOpenApiVO(bonjourServiceOpenApiVO).Execute()

Modify an existing Bonjour Service



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
	id := "id_example" // string | Bonjour Service ID
	bonjourServiceOpenApiVO := *openapiclient.NewBonjourServiceOpenApiVO("Name_example", []string{"ServiceIds_example"}) // BonjourServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyMdnsProfile(context.Background(), omadacId, siteId, id).BonjourServiceOpenApiVO(bonjourServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyMdnsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMdnsProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyMdnsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Bonjour Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMdnsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bonjourServiceOpenApiVO** | [**BonjourServiceOpenApiVO**](BonjourServiceOpenApiVO.md) |  | 

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


## ModifyOuiProfile

> OperationResponseWithoutResult ModifyOuiProfile(ctx, omadacId, siteId, ouiId).OuiProfileOpenApiVO(ouiProfileOpenApiVO).Execute()

Modify OUI Profile



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
	ouiId := "ouiId_example" // string | OUI ID
	ouiProfileOpenApiVO := *openapiclient.NewOuiProfileOpenApiVO("Name_example") // OuiProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyOuiProfile(context.Background(), omadacId, siteId, ouiId).OuiProfileOpenApiVO(ouiProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyOuiProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOuiProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyOuiProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ouiId** | **string** | OUI ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOuiProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ouiProfileOpenApiVO** | [**OuiProfileOpenApiVO**](OuiProfileOpenApiVO.md) |  | 

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


## ModifyPPSKProfile

> OperationResponseWithoutResult ModifyPPSKProfile(ctx, omadacId, siteId, profileId).PpskProfile(ppskProfile).Execute()

Modify PPSK profile



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
	profileId := "profileId_example" // string | PPSK profile Id
	ppskProfile := *openapiclient.NewPpskProfile([]openapiclient.PpskSetting{*openapiclient.NewPpskSetting("Name_example", "Psk_example")}, "ProfileName_example") // PpskProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyPPSKProfile(context.Background(), omadacId, siteId, profileId).PpskProfile(ppskProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyPPSKProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPPSKProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyPPSKProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | PPSK profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPPSKProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ppskProfile** | [**PpskProfile**](PpskProfile.md) |  | 

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


## ModifyRadiusProfile

> OperationResponseWithoutResult ModifyRadiusProfile(ctx, omadacId, siteId, radiusProfileId).CreateRadiusProfileOpenApiVO(createRadiusProfileOpenApiVO).Execute()

Modify an exist Radius profile



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
	radiusProfileId := "radiusProfileId_example" // string | Radius profile Id
	createRadiusProfileOpenApiVO := *openapiclient.NewCreateRadiusProfileOpenApiVO([]openapiclient.RadiusAuthServerOpenApiVO{*openapiclient.NewRadiusAuthServerOpenApiVO(int32(123), "RadiusPwd_example", "RadiusServerIp_example")}, "Name_example", false, false) // CreateRadiusProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyRadiusProfile(context.Background(), omadacId, siteId, radiusProfileId).CreateRadiusProfileOpenApiVO(createRadiusProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyRadiusProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRadiusProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyRadiusProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**radiusProfileId** | **string** | Radius profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRadiusProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createRadiusProfileOpenApiVO** | [**CreateRadiusProfileOpenApiVO**](CreateRadiusProfileOpenApiVO.md) |  | 

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


## ModifyRadiusUser

> OperationResponseWithoutResult ModifyRadiusUser(ctx, omadacId, siteId, userId).CreateRadiusUserOpenApiVO(createRadiusUserOpenApiVO).Execute()

Modify an exist Build-in Radius profile user



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
	userId := "userId_example" // string | Build-in Radius profile user Id
	createRadiusUserOpenApiVO := *openapiclient.NewCreateRadiusUserOpenApiVO(int32(123)) // CreateRadiusUserOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyRadiusUser(context.Background(), omadacId, siteId, userId).CreateRadiusUserOpenApiVO(createRadiusUserOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyRadiusUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRadiusUser`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyRadiusUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**userId** | **string** | Build-in Radius profile user Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRadiusUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createRadiusUserOpenApiVO** | [**CreateRadiusUserOpenApiVO**](CreateRadiusUserOpenApiVO.md) |  | 

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


## ModifyRateLimitProfile

> OperationResponseWithoutResult ModifyRateLimitProfile(ctx, omadacId, siteId, profileId).UpdateRateLimitProfileOpenApiVO(updateRateLimitProfileOpenApiVO).Execute()

Modify rate limit profile



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
	profileId := "profileId_example" // string | Profile ID
	updateRateLimitProfileOpenApiVO := *openapiclient.NewUpdateRateLimitProfileOpenApiVO(false, "Name_example", false) // UpdateRateLimitProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyRateLimitProfile(context.Background(), omadacId, siteId, profileId).UpdateRateLimitProfileOpenApiVO(updateRateLimitProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyRateLimitProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRateLimitProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyRateLimitProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRateLimitProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateRateLimitProfileOpenApiVO** | [**UpdateRateLimitProfileOpenApiVO**](UpdateRateLimitProfileOpenApiVO.md) |  | 

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


## ModifyServiceType

> OperationResponseWithoutResult ModifyServiceType(ctx, omadacId, siteId, id).GatewayQosServiceOpenApiVO(gatewayQosServiceOpenApiVO).Execute()

Modify an existing Gateway QoS Service



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
	id := "id_example" // string | Gateway QoS Service ID
	gatewayQosServiceOpenApiVO := *openapiclient.NewGatewayQosServiceOpenApiVO("Name_example", int32(123)) // GatewayQosServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyServiceType(context.Background(), omadacId, siteId, id).GatewayQosServiceOpenApiVO(gatewayQosServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyServiceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyServiceType`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyServiceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | Gateway QoS Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyServiceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gatewayQosServiceOpenApiVO** | [**GatewayQosServiceOpenApiVO**](GatewayQosServiceOpenApiVO.md) |  | 

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


## ModifyTimeRangeProfile

> OperationResponseWithoutResult ModifyTimeRangeProfile(ctx, omadacId, siteId, profileId).UpdateTimeRangeProfileOpenApiVO(updateTimeRangeProfileOpenApiVO).Execute()

Modify time range profile



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
	profileId := "profileId_example" // string | Profile ID
	updateTimeRangeProfileOpenApiVO := *openapiclient.NewUpdateTimeRangeProfileOpenApiVO(int32(123), "Name_example", []openapiclient.ScheduleTimeOpenApiVO{*openapiclient.NewScheduleTimeOpenApiVO(int32(123), int32(123), int32(123), int32(123), int32(123))}) // UpdateTimeRangeProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ModifyTimeRangeProfile(context.Background(), omadacId, siteId, profileId).UpdateTimeRangeProfileOpenApiVO(updateTimeRangeProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ModifyTimeRangeProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTimeRangeProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ModifyTimeRangeProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTimeRangeProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateTimeRangeProfileOpenApiVO** | [**UpdateTimeRangeProfileOpenApiVO**](UpdateTimeRangeProfileOpenApiVO.md) |  | 

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


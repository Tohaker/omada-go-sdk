# \ProfilesTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoCreatePPSKTemplate**](ProfilesTemplateAPI.md#AutoCreatePPSKTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ppsk-profile/generate-psk | Auto Create PSK templates
[**CreateApnProfileTemplate**](ProfilesTemplateAPI.md#CreateApnProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/apn | Create a new APN profile template
[**CreateGroupProfileTemplate**](ProfilesTemplateAPI.md#CreateGroupProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/groups | Create a new group profile template
[**CreateLdapProfileTemplate**](ProfilesTemplateAPI.md#CreateLdapProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/ldap | Create a new LDAP profile template
[**CreateMdnsProfileTemplate**](ProfilesTemplateAPI.md#CreateMdnsProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/bonjour-service | Create new Bonjour Service Template
[**CreateOuiProfileTemplate**](ProfilesTemplateAPI.md#CreateOuiProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/oui-profiles | Create OUI profile template
[**CreatePPSKProfileTemplate**](ProfilesTemplateAPI.md#CreatePPSKProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ppsk-profile | Create PPSK profile template
[**CreateRadiusProfileTemplate**](ProfilesTemplateAPI.md#CreateRadiusProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/radius | Create a new Radius profile template
[**CreateRadiusUserTemplate**](ProfilesTemplateAPI.md#CreateRadiusUserTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/radius-server/users | Create a new Build-in Radius profile user template
[**CreateRateLimitProfileTemplate**](ProfilesTemplateAPI.md#CreateRateLimitProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rate-limit-profiles | Create rate limit profile template
[**CreateTemplateServiceType**](ProfilesTemplateAPI.md#CreateTemplateServiceType) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/service-type | Create siteTemplate&#39;s new Gateway QoS Service
[**CreateTimeRangeProfileTemplate**](ProfilesTemplateAPI.md#CreateTimeRangeProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/time-range-profiles | Create time range profile template
[**DeleteAllPPSKProfilesTemplate**](ProfilesTemplateAPI.md#DeleteAllPPSKProfilesTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ppsk-profiles | Delete All PPSK profile templates
[**DeleteApnProfileTemplate**](ProfilesTemplateAPI.md#DeleteApnProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/apn/{profileId} | Delete an exist APN profile template
[**DeleteGroupProfileTemplate**](ProfilesTemplateAPI.md#DeleteGroupProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/groups/{groupType}/{groupId} | Delete an exist group profile template
[**DeleteLdapProfileTemplate**](ProfilesTemplateAPI.md#DeleteLdapProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/ldap/{ldapProfileId} | Delete an exist LDAP profile template
[**DeleteMdnsProfileTemplate**](ProfilesTemplateAPI.md#DeleteMdnsProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/bonjour-service/{id} | Delete an existing Bonjour Service Template
[**DeleteOuiProfileTemplate**](ProfilesTemplateAPI.md#DeleteOuiProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/oui-profiles/{ouiId} | Delete OUI Profile template
[**DeletePPSKProfileTemplate**](ProfilesTemplateAPI.md#DeletePPSKProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ppsk-profile/{profileId} | Delete PPSK profile template
[**DeleteRadiusProfileTemplate**](ProfilesTemplateAPI.md#DeleteRadiusProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/radius/{radiusProfileId} | Delete an exist Radius profile template
[**DeleteRadiusUserTemplate**](ProfilesTemplateAPI.md#DeleteRadiusUserTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/radius-server/users/{userId} | Delete an exist Build-in Radius profile user template
[**DeleteRateLimitProfileTemplate**](ProfilesTemplateAPI.md#DeleteRateLimitProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rate-limit-profile/{profileId} | Delete rate limit profile template
[**DeleteTemplateServiceType**](ProfilesTemplateAPI.md#DeleteTemplateServiceType) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/service-type/{id} | Delete an existing Gateway QoS Service in siteTemplate
[**DeleteTimeRangeProfileTemplate**](ProfilesTemplateAPI.md#DeleteTimeRangeProfileTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/time-range-profile/{profileId} | Delete time range profile template
[**GetApnProfileListTemplate**](ProfilesTemplateAPI.md#GetApnProfileListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/apn | Get APN profile template list
[**GetGroupProfilesByTypeTemplate**](ProfilesTemplateAPI.md#GetGroupProfilesByTypeTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/groups/{groupType} | Get group profile template list by type
[**GetGroupProfilesTemplate**](ProfilesTemplateAPI.md#GetGroupProfilesTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/groups | Get group profile template list
[**GetLdapProfileListTemplate**](ProfilesTemplateAPI.md#GetLdapProfileListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/ldap | Get LDAP profile template list
[**GetOuiProfileFullListTemplate**](ProfilesTemplateAPI.md#GetOuiProfileFullListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/oui-profiles/all | Get OUI profile template summary list
[**GetOuiProfileListTemplate**](ProfilesTemplateAPI.md#GetOuiProfileListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/oui-profiles | Get OUI profile template list
[**GetPPSKProfileDetailTemplate**](ProfilesTemplateAPI.md#GetPPSKProfileDetailTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ppsk-profile/{profileId} | Get PPSK profile template detail
[**GetPPSKProfilesTemplate**](ProfilesTemplateAPI.md#GetPPSKProfilesTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ppsk-profiles | Get PPSK profile templates list
[**GetRadiusProfileListTemplate**](ProfilesTemplateAPI.md#GetRadiusProfileListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/radius | Get Radius profile template list
[**GetRadiusUserListTemplate**](ProfilesTemplateAPI.md#GetRadiusUserListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/radius-server/users | Get Build-in Radius profile user template list
[**GetRateLimitProfileListTemplate**](ProfilesTemplateAPI.md#GetRateLimitProfileListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rate-limit-profiles | Get rate limit profile template list
[**GetTemplateServiceTypeSummary**](ProfilesTemplateAPI.md#GetTemplateServiceTypeSummary) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/service-type-summary | Get all Gateway QoS Service&#39;s ID and name info in siteTemplate
[**GetTimeRangeListTemplate**](ProfilesTemplateAPI.md#GetTimeRangeListTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/time-range-profiles | Get time range profile template list
[**ListMdnsProfileTemplate**](ProfilesTemplateAPI.md#ListMdnsProfileTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/bonjour-service | Get Bonjour Service Template list
[**ListTemplateServiceType**](ProfilesTemplateAPI.md#ListTemplateServiceType) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/service-type | Get siteTemplate&#39;s Gateway QoS Service list
[**ModifyApnProfileTemplate**](ProfilesTemplateAPI.md#ModifyApnProfileTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/apn/{profileId} | Modify an exist APN profile template
[**ModifyGroupProfileTemplate**](ProfilesTemplateAPI.md#ModifyGroupProfileTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/groups/{groupType}/{groupId} | Modify an exist group profile template
[**ModifyLdapProfileTemplate**](ProfilesTemplateAPI.md#ModifyLdapProfileTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/ldap/{ldapProfileId} | Modify an exist LDAP profile template
[**ModifyMdnsProfileTemplate**](ProfilesTemplateAPI.md#ModifyMdnsProfileTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/bonjour-service/{id} | Modify an existing Bonjour Service Template
[**ModifyOuiProfileTemplate**](ProfilesTemplateAPI.md#ModifyOuiProfileTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/oui-profiles/{ouiId} | Modify OUI Profile template
[**ModifyPPSKProfileTemplate**](ProfilesTemplateAPI.md#ModifyPPSKProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ppsk-profile/{profileId} | Modify PPSK profile template
[**ModifyRadiusProfileTemplate**](ProfilesTemplateAPI.md#ModifyRadiusProfileTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/radius/{radiusProfileId} | Modify an exist Radius profile template
[**ModifyRadiusUserTemplate**](ProfilesTemplateAPI.md#ModifyRadiusUserTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/radius-server/users/{userId} | Modify an exist Build-in Radius profile user template
[**ModifyRateLimitProfileTemplate**](ProfilesTemplateAPI.md#ModifyRateLimitProfileTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rate-limit-profile/{profileId} | Modify rate limit profile template
[**ModifyTemplateServiceType**](ProfilesTemplateAPI.md#ModifyTemplateServiceType) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/profiles/service-type/{id} | Modify an existing Gateway QoS Service in siteTemplate
[**ModifyTimeRangeProfileTemplate**](ProfilesTemplateAPI.md#ModifyTimeRangeProfileTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/time-range-profile/{profileId} | Modify time range profile template



## AutoCreatePPSKTemplate

> []PSK AutoCreatePPSKTemplate(ctx, omadacId, siteTemplateId).PpskAutoCreateSetting(ppskAutoCreateSetting).Execute()

Auto Create PSK templates



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	ppskAutoCreateSetting := *openapiclient.NewPpskAutoCreateSetting(int32(123), int32(123), "Prefix_example") // PpskAutoCreateSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.AutoCreatePPSKTemplate(context.Background(), omadacId, siteTemplateId).PpskAutoCreateSetting(ppskAutoCreateSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.AutoCreatePPSKTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreatePPSKTemplate`: []PSK
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.AutoCreatePPSKTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreatePPSKTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ppskAutoCreateSetting** | [**PpskAutoCreateSetting**](PpskAutoCreateSetting.md) |  | 

### Return type

[**[]PSK**](PSK.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApnProfileTemplate

> OperationResponseResIdOpenApiVO CreateApnProfileTemplate(ctx, omadacId, siteTemplateId).ApnProfileConfig(apnProfileConfig).Execute()

Create a new APN profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	apnProfileConfig := *openapiclient.NewApnProfileConfig(int32(123), int32(123), "Name_example", int32(123)) // ApnProfileConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateApnProfileTemplate(context.Background(), omadacId, siteTemplateId).ApnProfileConfig(apnProfileConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateApnProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApnProfileTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateApnProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApnProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apnProfileConfig** | [**ApnProfileConfig**](ApnProfileConfig.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupProfileTemplate

> OperationResponseResIdOpenApiVO CreateGroupProfileTemplate(ctx, omadacId, siteTemplateId).CreateGroupOpenApiVO(createGroupOpenApiVO).Execute()

Create a new group profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	createGroupOpenApiVO := *openapiclient.NewCreateGroupOpenApiVO("Name_example", int32(123)) // CreateGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateGroupProfileTemplate(context.Background(), omadacId, siteTemplateId).CreateGroupOpenApiVO(createGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateGroupProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupProfileTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateGroupProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createGroupOpenApiVO** | [**CreateGroupOpenApiVO**](CreateGroupOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLdapProfileTemplate

> OperationResponseResIdOpenApiVO CreateLdapProfileTemplate(ctx, omadacId, siteTemplateId).CreateLdapProfileOpenApiVO(createLdapProfileOpenApiVO).Execute()

Create a new LDAP profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	createLdapProfileOpenApiVO := *openapiclient.NewCreateLdapProfileOpenApiVO("BaseDn_example", int32(123), "Cn_example", int32(123), "Name_example", "Server_example", false, false) // CreateLdapProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateLdapProfileTemplate(context.Background(), omadacId, siteTemplateId).CreateLdapProfileOpenApiVO(createLdapProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateLdapProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLdapProfileTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateLdapProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLdapProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createLdapProfileOpenApiVO** | [**CreateLdapProfileOpenApiVO**](CreateLdapProfileOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMdnsProfileTemplate

> OperationResponseResponseIdVO CreateMdnsProfileTemplate(ctx, omadacId, siteTemplateId).BonjourServiceOpenApiVO(bonjourServiceOpenApiVO).Execute()

Create new Bonjour Service Template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	bonjourServiceOpenApiVO := *openapiclient.NewBonjourServiceOpenApiVO("Name_example", []string{"ServiceIds_example"}) // BonjourServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateMdnsProfileTemplate(context.Background(), omadacId, siteTemplateId).BonjourServiceOpenApiVO(bonjourServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateMdnsProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMdnsProfileTemplate`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateMdnsProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMdnsProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bonjourServiceOpenApiVO** | [**BonjourServiceOpenApiVO**](BonjourServiceOpenApiVO.md) |  | 

### Return type

[**OperationResponseResponseIdVO**](OperationResponseResponseIdVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOuiProfileTemplate

> OperationResponse CreateOuiProfileTemplate(ctx, omadacId, siteTemplateId).OuiProfileOpenApiVO(ouiProfileOpenApiVO).Execute()

Create OUI profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	ouiProfileOpenApiVO := *openapiclient.NewOuiProfileOpenApiVO("Name_example") // OuiProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateOuiProfileTemplate(context.Background(), omadacId, siteTemplateId).OuiProfileOpenApiVO(ouiProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateOuiProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOuiProfileTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateOuiProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOuiProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ouiProfileOpenApiVO** | [**OuiProfileOpenApiVO**](OuiProfileOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePPSKProfileTemplate

> OperationResponseWithoutResult CreatePPSKProfileTemplate(ctx, omadacId, siteTemplateId).PpskProfile(ppskProfile).Execute()

Create PPSK profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	ppskProfile := *openapiclient.NewPpskProfile([]openapiclient.PpskSetting{*openapiclient.NewPpskSetting("Name_example", "Psk_example")}, "ProfileName_example") // PpskProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreatePPSKProfileTemplate(context.Background(), omadacId, siteTemplateId).PpskProfile(ppskProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreatePPSKProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePPSKProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreatePPSKProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePPSKProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ppskProfile** | [**PpskProfile**](PpskProfile.md) |  | 

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


## CreateRadiusProfileTemplate

> OperationResponseResIdOpenApiVO CreateRadiusProfileTemplate(ctx, omadacId, siteTemplateId).CreateRadiusProfileOpenApiVO(createRadiusProfileOpenApiVO).Execute()

Create a new Radius profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	createRadiusProfileOpenApiVO := *openapiclient.NewCreateRadiusProfileOpenApiVO([]openapiclient.RadiusAuthServerOpenApiVO{*openapiclient.NewRadiusAuthServerOpenApiVO(int32(123), "RadiusPwd_example", "RadiusServerIp_example")}, "Name_example", false, false) // CreateRadiusProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateRadiusProfileTemplate(context.Background(), omadacId, siteTemplateId).CreateRadiusProfileOpenApiVO(createRadiusProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateRadiusProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRadiusProfileTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateRadiusProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRadiusProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRadiusProfileOpenApiVO** | [**CreateRadiusProfileOpenApiVO**](CreateRadiusProfileOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRadiusUserTemplate

> OperationResponseResIdOpenApiVO CreateRadiusUserTemplate(ctx, omadacId, siteTemplateId).CreateRadiusUserOpenApiVO(createRadiusUserOpenApiVO).Execute()

Create a new Build-in Radius profile user template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	createRadiusUserOpenApiVO := *openapiclient.NewCreateRadiusUserOpenApiVO(int32(123)) // CreateRadiusUserOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateRadiusUserTemplate(context.Background(), omadacId, siteTemplateId).CreateRadiusUserOpenApiVO(createRadiusUserOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateRadiusUserTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRadiusUserTemplate`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateRadiusUserTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRadiusUserTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRadiusUserOpenApiVO** | [**CreateRadiusUserOpenApiVO**](CreateRadiusUserOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRateLimitProfileTemplate

> OperationResponseWithoutResult CreateRateLimitProfileTemplate(ctx, omadacId, siteTemplateId).CreateRateLimitProfileOpenApiVO(createRateLimitProfileOpenApiVO).Execute()

Create rate limit profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	createRateLimitProfileOpenApiVO := *openapiclient.NewCreateRateLimitProfileOpenApiVO(false, "Name_example", false) // CreateRateLimitProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateRateLimitProfileTemplate(context.Background(), omadacId, siteTemplateId).CreateRateLimitProfileOpenApiVO(createRateLimitProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateRateLimitProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRateLimitProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateRateLimitProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRateLimitProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRateLimitProfileOpenApiVO** | [**CreateRateLimitProfileOpenApiVO**](CreateRateLimitProfileOpenApiVO.md) |  | 

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


## CreateTemplateServiceType

> OperationResponseResponseIdVO CreateTemplateServiceType(ctx, omadacId, siteTemplateId).GatewayQosServiceOpenApiVO(gatewayQosServiceOpenApiVO).Execute()

Create siteTemplate's new Gateway QoS Service



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	gatewayQosServiceOpenApiVO := *openapiclient.NewGatewayQosServiceOpenApiVO("Name_example", int32(123)) // GatewayQosServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateTemplateServiceType(context.Background(), omadacId, siteTemplateId).GatewayQosServiceOpenApiVO(gatewayQosServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateTemplateServiceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplateServiceType`: OperationResponseResponseIdVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateTemplateServiceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateServiceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayQosServiceOpenApiVO** | [**GatewayQosServiceOpenApiVO**](GatewayQosServiceOpenApiVO.md) |  | 

### Return type

[**OperationResponseResponseIdVO**](OperationResponseResponseIdVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTimeRangeProfileTemplate

> OperationResponseWithoutResult CreateTimeRangeProfileTemplate(ctx, omadacId, siteTemplateId).CreateTimeRangeProfileOpenApiVO(createTimeRangeProfileOpenApiVO).Execute()

Create time range profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	createTimeRangeProfileOpenApiVO := *openapiclient.NewCreateTimeRangeProfileOpenApiVO(int32(123), "Name_example", []openapiclient.ScheduleTimeOpenApiVO{*openapiclient.NewScheduleTimeOpenApiVO(int32(123), int32(123), int32(123), int32(123), int32(123))}) // CreateTimeRangeProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.CreateTimeRangeProfileTemplate(context.Background(), omadacId, siteTemplateId).CreateTimeRangeProfileOpenApiVO(createTimeRangeProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.CreateTimeRangeProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTimeRangeProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.CreateTimeRangeProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTimeRangeProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTimeRangeProfileOpenApiVO** | [**CreateTimeRangeProfileOpenApiVO**](CreateTimeRangeProfileOpenApiVO.md) |  | 

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


## DeleteAllPPSKProfilesTemplate

> OperationResponseWithoutResult DeleteAllPPSKProfilesTemplate(ctx, omadacId, siteTemplateId).Execute()

Delete All PPSK profile templates



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteAllPPSKProfilesTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteAllPPSKProfilesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllPPSKProfilesTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteAllPPSKProfilesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllPPSKProfilesTemplateRequest struct via the builder pattern


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


## DeleteApnProfileTemplate

> OperationResponseWithoutResult DeleteApnProfileTemplate(ctx, omadacId, siteTemplateId, profileId).Execute()

Delete an exist APN profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | APN profile ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteApnProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteApnProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApnProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteApnProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | APN profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApnProfileTemplateRequest struct via the builder pattern


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


## DeleteGroupProfileTemplate

> OperationResponseWithoutResult DeleteGroupProfileTemplate(ctx, omadacId, siteTemplateId, groupId, groupType).Execute()

Delete an exist group profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	groupId := "groupId_example" // string | Group profile id.
	groupType := "groupType_example" // string | groupType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteGroupProfileTemplate(context.Background(), omadacId, siteTemplateId, groupId, groupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteGroupProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteGroupProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**groupId** | **string** | Group profile id. | 
**groupType** | **string** | groupType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupProfileTemplateRequest struct via the builder pattern


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


## DeleteLdapProfileTemplate

> OperationResponseWithoutResult DeleteLdapProfileTemplate(ctx, omadacId, siteTemplateId, ldapProfileId).Execute()

Delete an exist LDAP profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	ldapProfileId := "ldapProfileId_example" // string | LDAP profile ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteLdapProfileTemplate(context.Background(), omadacId, siteTemplateId, ldapProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteLdapProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLdapProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteLdapProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ldapProfileId** | **string** | LDAP profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLdapProfileTemplateRequest struct via the builder pattern


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


## DeleteMdnsProfileTemplate

> OperationResponseWithoutResult DeleteMdnsProfileTemplate(ctx, omadacId, siteTemplateId, id).Execute()

Delete an existing Bonjour Service Template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	id := "id_example" // string | Bonjour Service Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteMdnsProfileTemplate(context.Background(), omadacId, siteTemplateId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteMdnsProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMdnsProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteMdnsProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Bonjour Service Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMdnsProfileTemplateRequest struct via the builder pattern


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


## DeleteOuiProfileTemplate

> OperationResponseWithoutResult DeleteOuiProfileTemplate(ctx, omadacId, siteTemplateId, ouiId).Execute()

Delete OUI Profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	ouiId := "ouiId_example" // string | OUI ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteOuiProfileTemplate(context.Background(), omadacId, siteTemplateId, ouiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteOuiProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOuiProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteOuiProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ouiId** | **string** | OUI ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOuiProfileTemplateRequest struct via the builder pattern


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


## DeletePPSKProfileTemplate

> OperationResponseWithoutResult DeletePPSKProfileTemplate(ctx, omadacId, siteTemplateId, profileId).Execute()

Delete PPSK profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | PPSK profile Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeletePPSKProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeletePPSKProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePPSKProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeletePPSKProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | PPSK profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePPSKProfileTemplateRequest struct via the builder pattern


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


## DeleteRadiusProfileTemplate

> OperationResponseWithoutResult DeleteRadiusProfileTemplate(ctx, omadacId, siteTemplateId, radiusProfileId).Execute()

Delete an exist Radius profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	radiusProfileId := "radiusProfileId_example" // string | Radius profile Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteRadiusProfileTemplate(context.Background(), omadacId, siteTemplateId, radiusProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteRadiusProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRadiusProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteRadiusProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**radiusProfileId** | **string** | Radius profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRadiusProfileTemplateRequest struct via the builder pattern


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


## DeleteRadiusUserTemplate

> OperationResponseWithoutResult DeleteRadiusUserTemplate(ctx, omadacId, siteTemplateId, userId).Execute()

Delete an exist Build-in Radius profile user template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	userId := "userId_example" // string | Build-in Radius profile user Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteRadiusUserTemplate(context.Background(), omadacId, siteTemplateId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteRadiusUserTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRadiusUserTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteRadiusUserTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**userId** | **string** | Build-in Radius profile user Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRadiusUserTemplateRequest struct via the builder pattern


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


## DeleteRateLimitProfileTemplate

> OperationResponseWithoutResult DeleteRateLimitProfileTemplate(ctx, omadacId, siteTemplateId, profileId).Execute()

Delete rate limit profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | Profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteRateLimitProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteRateLimitProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRateLimitProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteRateLimitProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRateLimitProfileTemplateRequest struct via the builder pattern


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


## DeleteTemplateServiceType

> OperationResponseWithoutResult DeleteTemplateServiceType(ctx, omadacId, siteTemplateId, id).Execute()

Delete an existing Gateway QoS Service in siteTemplate



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	id := "id_example" // string | Gateway QoS Service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteTemplateServiceType(context.Background(), omadacId, siteTemplateId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteTemplateServiceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplateServiceType`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteTemplateServiceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Gateway QoS Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateServiceTypeRequest struct via the builder pattern


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


## DeleteTimeRangeProfileTemplate

> OperationResponseWithoutResult DeleteTimeRangeProfileTemplate(ctx, omadacId, siteTemplateId, profileId).Execute()

Delete time range profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | Profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.DeleteTimeRangeProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.DeleteTimeRangeProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTimeRangeProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.DeleteTimeRangeProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeRangeProfileTemplateRequest struct via the builder pattern


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


## GetApnProfileListTemplate

> OperationResponseListApnProfile GetApnProfileListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get APN profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetApnProfileListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetApnProfileListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApnProfileListTemplate`: OperationResponseListApnProfile
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetApnProfileListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApnProfileListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListApnProfile**](OperationResponseListApnProfile.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupProfilesByTypeTemplate

> OperationResponseListGroupOpenApiVO GetGroupProfilesByTypeTemplate(ctx, omadacId, siteTemplateId, groupType).Execute()

Get group profile template list by type



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	groupType := "groupType_example" // string | Type of group profile, 0:IP Group; 1:IP Port Group; 2：Mac Group; 3:IPv6 Group; 4:IPv6 Port Group; 5:Country Group; 7:Domain Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetGroupProfilesByTypeTemplate(context.Background(), omadacId, siteTemplateId, groupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetGroupProfilesByTypeTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupProfilesByTypeTemplate`: OperationResponseListGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetGroupProfilesByTypeTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**groupType** | **string** | Type of group profile, 0:IP Group; 1:IP Port Group; 2：Mac Group; 3:IPv6 Group; 4:IPv6 Port Group; 5:Country Group; 7:Domain Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupProfilesByTypeTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListGroupOpenApiVO**](OperationResponseListGroupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupProfilesTemplate

> OperationResponseListGroupOpenApiVO GetGroupProfilesTemplate(ctx, omadacId, siteTemplateId).Execute()

Get group profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetGroupProfilesTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetGroupProfilesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupProfilesTemplate`: OperationResponseListGroupOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetGroupProfilesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupProfilesTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListGroupOpenApiVO**](OperationResponseListGroupOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapProfileListTemplate

> OperationResponseListLdapProfileOpenApiVO GetLdapProfileListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get LDAP profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetLdapProfileListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetLdapProfileListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapProfileListTemplate`: OperationResponseListLdapProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetLdapProfileListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapProfileListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListLdapProfileOpenApiVO**](OperationResponseListLdapProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOuiProfileFullListTemplate

> OperationResponseListOuiProfileSummaryOpenApiVO GetOuiProfileFullListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get OUI profile template summary list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetOuiProfileFullListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetOuiProfileFullListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOuiProfileFullListTemplate`: OperationResponseListOuiProfileSummaryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetOuiProfileFullListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOuiProfileFullListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListOuiProfileSummaryOpenApiVO**](OperationResponseListOuiProfileSummaryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOuiProfileListTemplate

> OperationResponseOuiGridVOOuiProfileQueryOpenApiVO GetOuiProfileListTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get OUI profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetOuiProfileListTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetOuiProfileListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOuiProfileListTemplate`: OperationResponseOuiGridVOOuiProfileQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetOuiProfileListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOuiProfileListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseOuiGridVOOuiProfileQueryOpenApiVO**](OperationResponseOuiGridVOOuiProfileQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPPSKProfileDetailTemplate

> OperationResponsePPSKProfileVO GetPPSKProfileDetailTemplate(ctx, omadacId, siteTemplateId, profileId).Execute()

Get PPSK profile template detail



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | PPSK profile Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetPPSKProfileDetailTemplate(context.Background(), omadacId, siteTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetPPSKProfileDetailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPPSKProfileDetailTemplate`: OperationResponsePPSKProfileVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetPPSKProfileDetailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | PPSK profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPPSKProfileDetailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponsePPSKProfileVO**](OperationResponsePPSKProfileVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPPSKProfilesTemplate

> OperationResponseListPpskProfileBriefInfo GetPPSKProfilesTemplate(ctx, omadacId, siteTemplateId).Type_(type_).Execute()

Get PPSK profile templates list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	type_ := int32(56) // int32 | PPSK Profile type, should be a value as follows: 0: PPSK Without RADIUS; 1: PPSK With Built-In RADIUS

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetPPSKProfilesTemplate(context.Background(), omadacId, siteTemplateId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetPPSKProfilesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPPSKProfilesTemplate`: OperationResponseListPpskProfileBriefInfo
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetPPSKProfilesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPPSKProfilesTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **int32** | PPSK Profile type, should be a value as follows: 0: PPSK Without RADIUS; 1: PPSK With Built-In RADIUS | 

### Return type

[**OperationResponseListPpskProfileBriefInfo**](OperationResponseListPpskProfileBriefInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadiusProfileListTemplate

> OperationResponseListRadiusProfileOpenApiVO GetRadiusProfileListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get Radius profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetRadiusProfileListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetRadiusProfileListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadiusProfileListTemplate`: OperationResponseListRadiusProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetRadiusProfileListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadiusProfileListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRadiusProfileOpenApiVO**](OperationResponseListRadiusProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadiusUserListTemplate

> OperationResponseGridVORadiusUserOpenApiVO GetRadiusUserListTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SortsUsername(sortsUsername).Execute()

Get Build-in Radius profile user template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsUsername := "sortsUsername_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetRadiusUserListTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SortsUsername(sortsUsername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetRadiusUserListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadiusUserListTemplate`: OperationResponseGridVORadiusUserOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetRadiusUserListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadiusUserListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsUsername** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseGridVORadiusUserOpenApiVO**](OperationResponseGridVORadiusUserOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateLimitProfileListTemplate

> OperationResponseListRateLimitProfileOpenApiVO GetRateLimitProfileListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get rate limit profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetRateLimitProfileListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetRateLimitProfileListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateLimitProfileListTemplate`: OperationResponseListRateLimitProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetRateLimitProfileListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitProfileListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListRateLimitProfileOpenApiVO**](OperationResponseListRateLimitProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateServiceTypeSummary

> OperationResponseResponseDataVOGatewayQosServiceSummaryOpenApiVO GetTemplateServiceTypeSummary(ctx, omadacId, siteTemplateId).Execute()

Get all Gateway QoS Service's ID and name info in siteTemplate



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetTemplateServiceTypeSummary(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetTemplateServiceTypeSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateServiceTypeSummary`: OperationResponseResponseDataVOGatewayQosServiceSummaryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetTemplateServiceTypeSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateServiceTypeSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseResponseDataVOGatewayQosServiceSummaryOpenApiVO**](OperationResponseResponseDataVOGatewayQosServiceSummaryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeRangeListTemplate

> OperationResponseListTimeRangeProfileOpenApiVO GetTimeRangeListTemplate(ctx, omadacId, siteTemplateId).Execute()

Get time range profile template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.GetTimeRangeListTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.GetTimeRangeListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeRangeListTemplate`: OperationResponseListTimeRangeProfileOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.GetTimeRangeListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeRangeListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListTimeRangeProfileOpenApiVO**](OperationResponseListTimeRangeProfileOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMdnsProfileTemplate

> OperationResponseResponseDataVOBonjourServiceDetailOpenApiVO ListMdnsProfileTemplate(ctx, omadacId, siteTemplateId).Execute()

Get Bonjour Service Template list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ListMdnsProfileTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ListMdnsProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMdnsProfileTemplate`: OperationResponseResponseDataVOBonjourServiceDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ListMdnsProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMdnsProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseResponseDataVOBonjourServiceDetailOpenApiVO**](OperationResponseResponseDataVOBonjourServiceDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateServiceType

> OperationResponseGridVOGatewayQosServiceDetailOpenApiVO ListTemplateServiceType(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get siteTemplate's Gateway QoS Service list



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ListTemplateServiceType(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ListTemplateServiceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplateServiceType`: OperationResponseGridVOGatewayQosServiceDetailOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ListTemplateServiceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateServiceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOGatewayQosServiceDetailOpenApiVO**](OperationResponseGridVOGatewayQosServiceDetailOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyApnProfileTemplate

> OperationResponseWithoutResult ModifyApnProfileTemplate(ctx, omadacId, siteTemplateId, profileId).ApnProfileConfig(apnProfileConfig).Execute()

Modify an exist APN profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | APN profile ID.
	apnProfileConfig := *openapiclient.NewApnProfileConfig(int32(123), int32(123), "Name_example", int32(123)) // ApnProfileConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyApnProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).ApnProfileConfig(apnProfileConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyApnProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApnProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyApnProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | APN profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyApnProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apnProfileConfig** | [**ApnProfileConfig**](ApnProfileConfig.md) |  | 

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


## ModifyGroupProfileTemplate

> OperationResponseWithoutResult ModifyGroupProfileTemplate(ctx, omadacId, siteTemplateId, groupType, groupId).CreateGroupOpenApiVO(createGroupOpenApiVO).Execute()

Modify an exist group profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	groupType := "groupType_example" // string | Type of group profile, 0:IP Group; 1:IP Port Group; 2：Mac Group; 3:IPv6 Group; 4:IPv6 Port Group; 5:Country Group; 7:Domain Group.
	groupId := "groupId_example" // string | Group profile id.
	createGroupOpenApiVO := *openapiclient.NewCreateGroupOpenApiVO("Name_example", int32(123)) // CreateGroupOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyGroupProfileTemplate(context.Background(), omadacId, siteTemplateId, groupType, groupId).CreateGroupOpenApiVO(createGroupOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyGroupProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGroupProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyGroupProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**groupType** | **string** | Type of group profile, 0:IP Group; 1:IP Port Group; 2：Mac Group; 3:IPv6 Group; 4:IPv6 Port Group; 5:Country Group; 7:Domain Group. | 
**groupId** | **string** | Group profile id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyGroupProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **createGroupOpenApiVO** | [**CreateGroupOpenApiVO**](CreateGroupOpenApiVO.md) |  | 

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


## ModifyLdapProfileTemplate

> OperationResponseWithoutResult ModifyLdapProfileTemplate(ctx, omadacId, siteTemplateId, ldapProfileId).CreateLdapProfileOpenApiVO(createLdapProfileOpenApiVO).Execute()

Modify an exist LDAP profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	ldapProfileId := "ldapProfileId_example" // string | LDAP profile ID.
	createLdapProfileOpenApiVO := *openapiclient.NewCreateLdapProfileOpenApiVO("BaseDn_example", int32(123), "Cn_example", int32(123), "Name_example", "Server_example", false, false) // CreateLdapProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyLdapProfileTemplate(context.Background(), omadacId, siteTemplateId, ldapProfileId).CreateLdapProfileOpenApiVO(createLdapProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyLdapProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLdapProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyLdapProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ldapProfileId** | **string** | LDAP profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLdapProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLdapProfileOpenApiVO** | [**CreateLdapProfileOpenApiVO**](CreateLdapProfileOpenApiVO.md) |  | 

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


## ModifyMdnsProfileTemplate

> OperationResponseWithoutResult ModifyMdnsProfileTemplate(ctx, omadacId, siteTemplateId, id).BonjourServiceOpenApiVO(bonjourServiceOpenApiVO).Execute()

Modify an existing Bonjour Service Template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	id := "id_example" // string | Bonjour Service Template ID
	bonjourServiceOpenApiVO := *openapiclient.NewBonjourServiceOpenApiVO("Name_example", []string{"ServiceIds_example"}) // BonjourServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyMdnsProfileTemplate(context.Background(), omadacId, siteTemplateId, id).BonjourServiceOpenApiVO(bonjourServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyMdnsProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMdnsProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyMdnsProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Bonjour Service Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMdnsProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bonjourServiceOpenApiVO** | [**BonjourServiceOpenApiVO**](BonjourServiceOpenApiVO.md) |  | 

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


## ModifyOuiProfileTemplate

> OperationResponseWithoutResult ModifyOuiProfileTemplate(ctx, omadacId, siteTemplateId, ouiId).OuiProfileOpenApiVO(ouiProfileOpenApiVO).Execute()

Modify OUI Profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	ouiId := "ouiId_example" // string | OUI ID
	ouiProfileOpenApiVO := *openapiclient.NewOuiProfileOpenApiVO("Name_example") // OuiProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyOuiProfileTemplate(context.Background(), omadacId, siteTemplateId, ouiId).OuiProfileOpenApiVO(ouiProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyOuiProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOuiProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyOuiProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ouiId** | **string** | OUI ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOuiProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ouiProfileOpenApiVO** | [**OuiProfileOpenApiVO**](OuiProfileOpenApiVO.md) |  | 

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


## ModifyPPSKProfileTemplate

> OperationResponseWithoutResult ModifyPPSKProfileTemplate(ctx, omadacId, siteTemplateId, profileId).PpskProfile(ppskProfile).Execute()

Modify PPSK profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | PPSK profile Id
	ppskProfile := *openapiclient.NewPpskProfile([]openapiclient.PpskSetting{*openapiclient.NewPpskSetting("Name_example", "Psk_example")}, "ProfileName_example") // PpskProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyPPSKProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).PpskProfile(ppskProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyPPSKProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPPSKProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyPPSKProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | PPSK profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPPSKProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ppskProfile** | [**PpskProfile**](PpskProfile.md) |  | 

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


## ModifyRadiusProfileTemplate

> OperationResponseWithoutResult ModifyRadiusProfileTemplate(ctx, omadacId, siteTemplateId, radiusProfileId).CreateRadiusProfileOpenApiVO(createRadiusProfileOpenApiVO).Execute()

Modify an exist Radius profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	radiusProfileId := "radiusProfileId_example" // string | Radius profile Id
	createRadiusProfileOpenApiVO := *openapiclient.NewCreateRadiusProfileOpenApiVO([]openapiclient.RadiusAuthServerOpenApiVO{*openapiclient.NewRadiusAuthServerOpenApiVO(int32(123), "RadiusPwd_example", "RadiusServerIp_example")}, "Name_example", false, false) // CreateRadiusProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyRadiusProfileTemplate(context.Background(), omadacId, siteTemplateId, radiusProfileId).CreateRadiusProfileOpenApiVO(createRadiusProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyRadiusProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRadiusProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyRadiusProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**radiusProfileId** | **string** | Radius profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRadiusProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createRadiusProfileOpenApiVO** | [**CreateRadiusProfileOpenApiVO**](CreateRadiusProfileOpenApiVO.md) |  | 

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


## ModifyRadiusUserTemplate

> OperationResponseWithoutResult ModifyRadiusUserTemplate(ctx, omadacId, siteTemplateId, userId).CreateRadiusUserOpenApiVO(createRadiusUserOpenApiVO).Execute()

Modify an exist Build-in Radius profile user template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	userId := "userId_example" // string | Build-in Radius profile user Id
	createRadiusUserOpenApiVO := *openapiclient.NewCreateRadiusUserOpenApiVO(int32(123)) // CreateRadiusUserOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyRadiusUserTemplate(context.Background(), omadacId, siteTemplateId, userId).CreateRadiusUserOpenApiVO(createRadiusUserOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyRadiusUserTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRadiusUserTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyRadiusUserTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**userId** | **string** | Build-in Radius profile user Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRadiusUserTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createRadiusUserOpenApiVO** | [**CreateRadiusUserOpenApiVO**](CreateRadiusUserOpenApiVO.md) |  | 

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


## ModifyRateLimitProfileTemplate

> OperationResponseWithoutResult ModifyRateLimitProfileTemplate(ctx, omadacId, siteTemplateId, profileId).UpdateRateLimitProfileOpenApiVO(updateRateLimitProfileOpenApiVO).Execute()

Modify rate limit profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | Profile ID
	updateRateLimitProfileOpenApiVO := *openapiclient.NewUpdateRateLimitProfileOpenApiVO(false, "Name_example", false) // UpdateRateLimitProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyRateLimitProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).UpdateRateLimitProfileOpenApiVO(updateRateLimitProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyRateLimitProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRateLimitProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyRateLimitProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRateLimitProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateRateLimitProfileOpenApiVO** | [**UpdateRateLimitProfileOpenApiVO**](UpdateRateLimitProfileOpenApiVO.md) |  | 

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


## ModifyTemplateServiceType

> OperationResponseWithoutResult ModifyTemplateServiceType(ctx, omadacId, siteTemplateId, id).GatewayQosServiceOpenApiVO(gatewayQosServiceOpenApiVO).Execute()

Modify an existing Gateway QoS Service in siteTemplate



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	id := "id_example" // string | Gateway QoS Service ID
	gatewayQosServiceOpenApiVO := *openapiclient.NewGatewayQosServiceOpenApiVO("Name_example", int32(123)) // GatewayQosServiceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyTemplateServiceType(context.Background(), omadacId, siteTemplateId, id).GatewayQosServiceOpenApiVO(gatewayQosServiceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyTemplateServiceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTemplateServiceType`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyTemplateServiceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**id** | **string** | Gateway QoS Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTemplateServiceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gatewayQosServiceOpenApiVO** | [**GatewayQosServiceOpenApiVO**](GatewayQosServiceOpenApiVO.md) |  | 

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


## ModifyTimeRangeProfileTemplate

> OperationResponseWithoutResult ModifyTimeRangeProfileTemplate(ctx, omadacId, siteTemplateId, profileId).UpdateTimeRangeProfileOpenApiVO(updateTimeRangeProfileOpenApiVO).Execute()

Modify time range profile template



### Example

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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	profileId := "profileId_example" // string | Profile ID
	updateTimeRangeProfileOpenApiVO := *openapiclient.NewUpdateTimeRangeProfileOpenApiVO(int32(123), "Name_example", []openapiclient.ScheduleTimeOpenApiVO{*openapiclient.NewScheduleTimeOpenApiVO(int32(123), int32(123), int32(123), int32(123), int32(123))}) // UpdateTimeRangeProfileOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesTemplateAPI.ModifyTimeRangeProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).UpdateTimeRangeProfileOpenApiVO(updateTimeRangeProfileOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesTemplateAPI.ModifyTimeRangeProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTimeRangeProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ProfilesTemplateAPI.ModifyTimeRangeProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTimeRangeProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateTimeRangeProfileOpenApiVO** | [**UpdateTimeRangeProfileOpenApiVO**](UpdateTimeRangeProfileOpenApiVO.md) |  | 

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


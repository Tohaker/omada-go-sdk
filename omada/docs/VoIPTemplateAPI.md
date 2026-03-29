# \VoIPTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCallForwardingRule**](VoIPTemplateAPI.md#AddCallForwardingRule) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/call-forwarding | Create new siteTemplate&#39;s Call Forwarding
[**BatchDeleteVoipTelephoneBookTemplate**](VoIPTemplateAPI.md#BatchDeleteVoipTelephoneBookTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/telephone-book/batch-delete | Batch delete existing Contact Person in siteTemplate
[**CreateCallBlockingProfile**](VoIPTemplateAPI.md#CreateCallBlockingProfile) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/call-blocking | Create new call blocking profile template
[**CreateProviderProfile**](VoIPTemplateAPI.md#CreateProviderProfile) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/provider-profiles | Create new provider profile template
[**CreateVoipTelephoneBookTemplate**](VoIPTemplateAPI.md#CreateVoipTelephoneBookTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/telephone-book | Create new siteTemplate&#39;s Contact Person
[**DeleteCallBlockingProfile1**](VoIPTemplateAPI.md#DeleteCallBlockingProfile1) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/call-blocking/delete | Delete call blocking profile template
[**DeleteCallForwardingRule**](VoIPTemplateAPI.md#DeleteCallForwardingRule) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/call-forwarding | Batch delete existing Call Forwarding in siteTemplate
[**DeleteProviderProfiles1**](VoIPTemplateAPI.md#DeleteProviderProfiles1) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/provider-profiles/delete | Delete provider profile(s) template
[**DeleteVoipTelephoneBookTemplate**](VoIPTemplateAPI.md#DeleteVoipTelephoneBookTemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/telephone-book/{contactId} | Delete an existing Contact Person in siteTemplate
[**GetCallBlockingProfiles**](VoIPTemplateAPI.md#GetCallBlockingProfiles) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/call-blocking | Get call blocking profile list template
[**GetCallForwardingRulesGridTemplate**](VoIPTemplateAPI.md#GetCallForwardingRulesGridTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/call-forwarding/grid | Get siteTemplate&#39;s Call Forwarding list
[**GetDndSettings**](VoIPTemplateAPI.md#GetDndSettings) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/dnd/settings | Get DND settings template
[**GetGridProviderProfileList**](VoIPTemplateAPI.md#GetGridProviderProfileList) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/grid/provider-profiles | Get the grid of provider profile list template
[**GetVoipCallLogEnable**](VoIPTemplateAPI.md#GetVoipCallLogEnable) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/call-log-enable | Get siteTemplate&#39;s Voip CallLog Setting
[**GetVoipEmergency**](VoIPTemplateAPI.md#GetVoipEmergency) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/emergency-number-settings | Get Voip Emergency Number Info Template
[**GetVoipTelephoneBookTemplate**](VoIPTemplateAPI.md#GetVoipTelephoneBookTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/telephone-book | Get siteTemplate&#39;s Voip Contact Person list
[**ModifyCallBlockingProfile**](VoIPTemplateAPI.md#ModifyCallBlockingProfile) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/call-blocking/{profileId} | Modify a call blocking profile template
[**ModifyCallForwardingRule**](VoIPTemplateAPI.md#ModifyCallForwardingRule) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/call-forwarding/{forwardingRuleId} | Modify an existing Call Forwarding in siteTemplate
[**ModifyDndSettings**](VoIPTemplateAPI.md#ModifyDndSettings) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/dnd/settings | Modify DND settings template
[**ModifyProviderProfile**](VoIPTemplateAPI.md#ModifyProviderProfile) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/provider-profiles/{profileId} | Modify a provider profile template
[**ModifyVoipCallLogEnable**](VoIPTemplateAPI.md#ModifyVoipCallLogEnable) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/call-log-enable | Modify siteTemplate&#39;s Voip CallLog Setting
[**ModifyVoipEmergency**](VoIPTemplateAPI.md#ModifyVoipEmergency) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/emergency-number-settings | Modify Voip Emergency Number Setting Template
[**ModifyVoipTelephoneBookTemplate**](VoIPTemplateAPI.md#ModifyVoipTelephoneBookTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/voip/telephone-book/{contactId} | Modify an existing Contact Person in siteTemplate



## AddCallForwardingRule

> OperationResponseAddCallForwardingRuleResp AddCallForwardingRule(ctx, omadacId, siteTemplateId).CallForwardingRule(callForwardingRule).Execute()

Create new siteTemplate's Call Forwarding



### Example

```go
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
	callForwardingRule := *openapiclient.NewCallForwardingRule(int32(123), "DestNumber_example", false, int32(123)) // CallForwardingRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.AddCallForwardingRule(context.Background(), omadacId, siteTemplateId).CallForwardingRule(callForwardingRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.AddCallForwardingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCallForwardingRule`: OperationResponseAddCallForwardingRuleResp
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.AddCallForwardingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCallForwardingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **callForwardingRule** | [**CallForwardingRule**](CallForwardingRule.md) |  | 

### Return type

[**OperationResponseAddCallForwardingRuleResp**](OperationResponseAddCallForwardingRuleResp.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchDeleteVoipTelephoneBookTemplate

> OperationResponseWithoutResult BatchDeleteVoipTelephoneBookTemplate(ctx, omadacId, siteTemplateId).VoipTelephoneBookBatchSetting(voipTelephoneBookBatchSetting).Execute()

Batch delete existing Contact Person in siteTemplate



### Example

```go
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
	voipTelephoneBookBatchSetting := *openapiclient.NewVoipTelephoneBookBatchSetting([]string{"ContactIds_example"}, false, "SelectType_example") // VoipTelephoneBookBatchSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.BatchDeleteVoipTelephoneBookTemplate(context.Background(), omadacId, siteTemplateId).VoipTelephoneBookBatchSetting(voipTelephoneBookBatchSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.BatchDeleteVoipTelephoneBookTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteVoipTelephoneBookTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.BatchDeleteVoipTelephoneBookTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteVoipTelephoneBookTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voipTelephoneBookBatchSetting** | [**VoipTelephoneBookBatchSetting**](VoipTelephoneBookBatchSetting.md) |  | 

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


## CreateCallBlockingProfile

> OperationResponseWithoutResult CreateCallBlockingProfile(ctx, omadacId, siteTemplateId).CreateCallBlockingProfileEntity(createCallBlockingProfileEntity).Execute()

Create new call blocking profile template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	createCallBlockingProfileEntity := *openapiclient.NewCreateCallBlockingProfileEntity("ProfileName_example") // CreateCallBlockingProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.CreateCallBlockingProfile(context.Background(), omadacId, siteTemplateId).CreateCallBlockingProfileEntity(createCallBlockingProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.CreateCallBlockingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCallBlockingProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.CreateCallBlockingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallBlockingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createCallBlockingProfileEntity** | [**CreateCallBlockingProfileEntity**](CreateCallBlockingProfileEntity.md) |  | 

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


## CreateProviderProfile

> OperationResponseWithoutResult CreateProviderProfile(ctx, omadacId, siteTemplateId).CreateProviderProfileEntity(createProviderProfileEntity).Execute()

Create new provider profile template



### Example

```go
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
	createProviderProfileEntity := *openapiclient.NewCreateProviderProfileEntity("ProfileName_example", *openapiclient.NewProviderSettingVO(int32(123))) // CreateProviderProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.CreateProviderProfile(context.Background(), omadacId, siteTemplateId).CreateProviderProfileEntity(createProviderProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.CreateProviderProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProviderProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.CreateProviderProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProviderProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createProviderProfileEntity** | [**CreateProviderProfileEntity**](CreateProviderProfileEntity.md) |  | 

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


## CreateVoipTelephoneBookTemplate

> OperationResponseWithoutResult CreateVoipTelephoneBookTemplate(ctx, omadacId, siteTemplateId).VoipContactPersonSettings(voipContactPersonSettings).Execute()

Create new siteTemplate's Contact Person



### Example

```go
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
	voipContactPersonSettings := *openapiclient.NewVoipContactPersonSettings() // VoipContactPersonSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.CreateVoipTelephoneBookTemplate(context.Background(), omadacId, siteTemplateId).VoipContactPersonSettings(voipContactPersonSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.CreateVoipTelephoneBookTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoipTelephoneBookTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.CreateVoipTelephoneBookTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoipTelephoneBookTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voipContactPersonSettings** | [**VoipContactPersonSettings**](VoipContactPersonSettings.md) |  | 

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


## DeleteCallBlockingProfile1

> OperationResponseListProfilesBindedDeviceInfo DeleteCallBlockingProfile1(ctx, omadacId, siteTemplateId).DeleteCallBlockingProfileEntity(deleteCallBlockingProfileEntity).Execute()

Delete call blocking profile template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	deleteCallBlockingProfileEntity := *openapiclient.NewDeleteCallBlockingProfileEntity("ProfileId_example", false) // DeleteCallBlockingProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.DeleteCallBlockingProfile1(context.Background(), omadacId, siteTemplateId).DeleteCallBlockingProfileEntity(deleteCallBlockingProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.DeleteCallBlockingProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCallBlockingProfile1`: OperationResponseListProfilesBindedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.DeleteCallBlockingProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallBlockingProfile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteCallBlockingProfileEntity** | [**DeleteCallBlockingProfileEntity**](DeleteCallBlockingProfileEntity.md) |  | 

### Return type

[**OperationResponseListProfilesBindedDeviceInfo**](OperationResponseListProfilesBindedDeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallForwardingRule

> OperationResponseVoid DeleteCallForwardingRule(ctx, omadacId, siteTemplateId).DeleteCallForwardingRules(deleteCallForwardingRules).Execute()

Batch delete existing Call Forwarding in siteTemplate



### Example

```go
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
	deleteCallForwardingRules := *openapiclient.NewDeleteCallForwardingRules() // DeleteCallForwardingRules | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.DeleteCallForwardingRule(context.Background(), omadacId, siteTemplateId).DeleteCallForwardingRules(deleteCallForwardingRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.DeleteCallForwardingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCallForwardingRule`: OperationResponseVoid
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.DeleteCallForwardingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallForwardingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteCallForwardingRules** | [**DeleteCallForwardingRules**](DeleteCallForwardingRules.md) |  | 

### Return type

[**OperationResponseVoid**](OperationResponseVoid.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProviderProfiles1

> OperationResponseListProfilesBindedDeviceInfo DeleteProviderProfiles1(ctx, omadacId, siteTemplateId).DeleteProviderProfileTemplateEntity(deleteProviderProfileTemplateEntity).Execute()

Delete provider profile(s) template



### Example

```go
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
	deleteProviderProfileTemplateEntity := *openapiclient.NewDeleteProviderProfileTemplateEntity([]string{"ProfileIds_example"}) // DeleteProviderProfileTemplateEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.DeleteProviderProfiles1(context.Background(), omadacId, siteTemplateId).DeleteProviderProfileTemplateEntity(deleteProviderProfileTemplateEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.DeleteProviderProfiles1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProviderProfiles1`: OperationResponseListProfilesBindedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.DeleteProviderProfiles1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderProfiles1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteProviderProfileTemplateEntity** | [**DeleteProviderProfileTemplateEntity**](DeleteProviderProfileTemplateEntity.md) |  | 

### Return type

[**OperationResponseListProfilesBindedDeviceInfo**](OperationResponseListProfilesBindedDeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVoipTelephoneBookTemplate

> OperationResponseWithoutResult DeleteVoipTelephoneBookTemplate(ctx, omadacId, siteTemplateId, contactId).VoipDeleteTelephoneBook(voipDeleteTelephoneBook).Execute()

Delete an existing Contact Person in siteTemplate



### Example

```go
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
	contactId := "contactId_example" // string | Voip Contact Person ID
	voipDeleteTelephoneBook := *openapiclient.NewVoipDeleteTelephoneBook(false) // VoipDeleteTelephoneBook | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.DeleteVoipTelephoneBookTemplate(context.Background(), omadacId, siteTemplateId, contactId).VoipDeleteTelephoneBook(voipDeleteTelephoneBook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.DeleteVoipTelephoneBookTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVoipTelephoneBookTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.DeleteVoipTelephoneBookTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**contactId** | **string** | Voip Contact Person ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoipTelephoneBookTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **voipDeleteTelephoneBook** | [**VoipDeleteTelephoneBook**](VoipDeleteTelephoneBook.md) |  | 

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


## GetCallBlockingProfiles

> OperationResponseListCallBlockingProfileEntity GetCallBlockingProfiles(ctx, omadacId, siteTemplateId).Execute()

Get call blocking profile list template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.GetCallBlockingProfiles(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.GetCallBlockingProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallBlockingProfiles`: OperationResponseListCallBlockingProfileEntity
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.GetCallBlockingProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallBlockingProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListCallBlockingProfileEntity**](OperationResponseListCallBlockingProfileEntity.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallForwardingRulesGridTemplate

> OperationResponseCallForwardingRulesGrid GetCallForwardingRulesGridTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get siteTemplate's Call Forwarding list



### Example

```go
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
	resp, r, err := apiClient.VoIPTemplateAPI.GetCallForwardingRulesGridTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.GetCallForwardingRulesGridTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallForwardingRulesGridTemplate`: OperationResponseCallForwardingRulesGrid
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.GetCallForwardingRulesGridTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallForwardingRulesGridTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseCallForwardingRulesGrid**](OperationResponseCallForwardingRulesGrid.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDndSettings

> OperationResponseDndSettingEntity GetDndSettings(ctx, omadacId, siteTemplateId).Execute()

Get DND settings template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.GetDndSettings(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.GetDndSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDndSettings`: OperationResponseDndSettingEntity
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.GetDndSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDndSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDndSettingEntity**](OperationResponseDndSettingEntity.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridProviderProfileList

> OperationResponseGridVOProviderProfileEntity GetGridProviderProfileList(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get the grid of provider profile list template



### Example

```go
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
	resp, r, err := apiClient.VoIPTemplateAPI.GetGridProviderProfileList(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.GetGridProviderProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridProviderProfileList`: OperationResponseGridVOProviderProfileEntity
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.GetGridProviderProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridProviderProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOProviderProfileEntity**](OperationResponseGridVOProviderProfileEntity.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoipCallLogEnable

> OperationResponseVoipCallLogEnableSetting GetVoipCallLogEnable(ctx, omadacId, siteTemplateId).Execute()

Get siteTemplate's Voip CallLog Setting



### Example

```go
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
	resp, r, err := apiClient.VoIPTemplateAPI.GetVoipCallLogEnable(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.GetVoipCallLogEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipCallLogEnable`: OperationResponseVoipCallLogEnableSetting
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.GetVoipCallLogEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipCallLogEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVoipCallLogEnableSetting**](OperationResponseVoipCallLogEnableSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoipEmergency

> OperationResponseVoipEmergencyNumberSetting GetVoipEmergency(ctx, omadacId, siteTemplateId).Execute()

Get Voip Emergency Number Info Template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.GetVoipEmergency(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.GetVoipEmergency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipEmergency`: OperationResponseVoipEmergencyNumberSetting
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.GetVoipEmergency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipEmergencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVoipEmergencyNumberSetting**](OperationResponseVoipEmergencyNumberSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoipTelephoneBookTemplate

> OperationResponseGridVOVoipTelephoneBookSetting GetVoipTelephoneBookTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get siteTemplate's Voip Contact Person list



### Example

```go
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
	resp, r, err := apiClient.VoIPTemplateAPI.GetVoipTelephoneBookTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.GetVoipTelephoneBookTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipTelephoneBookTemplate`: OperationResponseGridVOVoipTelephoneBookSetting
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.GetVoipTelephoneBookTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipTelephoneBookTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOVoipTelephoneBookSetting**](OperationResponseGridVOVoipTelephoneBookSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCallBlockingProfile

> OperationResponseWithoutResult ModifyCallBlockingProfile(ctx, omadacId, siteTemplateId, profileId).ModifyCallBlockingProfileEntity(modifyCallBlockingProfileEntity).Execute()

Modify a call blocking profile template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	profileId := "profileId_example" // string | Profile ID
	modifyCallBlockingProfileEntity := *openapiclient.NewModifyCallBlockingProfileEntity() // ModifyCallBlockingProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.ModifyCallBlockingProfile(context.Background(), omadacId, siteTemplateId, profileId).ModifyCallBlockingProfileEntity(modifyCallBlockingProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.ModifyCallBlockingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCallBlockingProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.ModifyCallBlockingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCallBlockingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyCallBlockingProfileEntity** | [**ModifyCallBlockingProfileEntity**](ModifyCallBlockingProfileEntity.md) |  | 

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


## ModifyCallForwardingRule

> OperationResponseVoid ModifyCallForwardingRule(ctx, omadacId, siteTemplateId, forwardingRuleId).CallForwardingRule(callForwardingRule).Execute()

Modify an existing Call Forwarding in siteTemplate



### Example

```go
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
	forwardingRuleId := "forwardingRuleId_example" // string | Entry ID of the call forwarding rule to be modified.
	callForwardingRule := *openapiclient.NewCallForwardingRule(int32(123), "DestNumber_example", false, int32(123)) // CallForwardingRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.ModifyCallForwardingRule(context.Background(), omadacId, siteTemplateId, forwardingRuleId).CallForwardingRule(callForwardingRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.ModifyCallForwardingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCallForwardingRule`: OperationResponseVoid
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.ModifyCallForwardingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**forwardingRuleId** | **string** | Entry ID of the call forwarding rule to be modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCallForwardingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **callForwardingRule** | [**CallForwardingRule**](CallForwardingRule.md) |  | 

### Return type

[**OperationResponseVoid**](OperationResponseVoid.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDndSettings

> OperationResponseVoid ModifyDndSettings(ctx, omadacId, siteTemplateId).DndSettingEntity(dndSettingEntity).Execute()

Modify DND settings template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	dndSettingEntity := *openapiclient.NewDndSettingEntity(false) // DndSettingEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.ModifyDndSettings(context.Background(), omadacId, siteTemplateId).DndSettingEntity(dndSettingEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.ModifyDndSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDndSettings`: OperationResponseVoid
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.ModifyDndSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDndSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dndSettingEntity** | [**DndSettingEntity**](DndSettingEntity.md) |  | 

### Return type

[**OperationResponseVoid**](OperationResponseVoid.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyProviderProfile

> OperationResponseWithoutResult ModifyProviderProfile(ctx, omadacId, siteTemplateId, profileId).ModifyProviderProfileEntity(modifyProviderProfileEntity).Execute()

Modify a provider profile template



### Example

```go
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
	modifyProviderProfileEntity := *openapiclient.NewModifyProviderProfileEntity() // ModifyProviderProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.ModifyProviderProfile(context.Background(), omadacId, siteTemplateId, profileId).ModifyProviderProfileEntity(modifyProviderProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.ModifyProviderProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyProviderProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.ModifyProviderProfile`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyProviderProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyProviderProfileEntity** | [**ModifyProviderProfileEntity**](ModifyProviderProfileEntity.md) |  | 

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


## ModifyVoipCallLogEnable

> OperationResponseWithoutResult ModifyVoipCallLogEnable(ctx, omadacId, siteTemplateId).VoipCallLogEnableSetting(voipCallLogEnableSetting).Execute()

Modify siteTemplate's Voip CallLog Setting



### Example

```go
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
	voipCallLogEnableSetting := *openapiclient.NewVoipCallLogEnableSetting(false) // VoipCallLogEnableSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.ModifyVoipCallLogEnable(context.Background(), omadacId, siteTemplateId).VoipCallLogEnableSetting(voipCallLogEnableSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.ModifyVoipCallLogEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoipCallLogEnable`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.ModifyVoipCallLogEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipCallLogEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voipCallLogEnableSetting** | [**VoipCallLogEnableSetting**](VoipCallLogEnableSetting.md) |  | 

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


## ModifyVoipEmergency

> OperationResponseWithoutResult ModifyVoipEmergency(ctx, omadacId, siteTemplateId).VoipEmergencyNumberSetting(voipEmergencyNumberSetting).Execute()

Modify Voip Emergency Number Setting Template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	voipEmergencyNumberSetting := *openapiclient.NewVoipEmergencyNumberSetting(false) // VoipEmergencyNumberSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.ModifyVoipEmergency(context.Background(), omadacId, siteTemplateId).VoipEmergencyNumberSetting(voipEmergencyNumberSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.ModifyVoipEmergency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoipEmergency`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.ModifyVoipEmergency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipEmergencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voipEmergencyNumberSetting** | [**VoipEmergencyNumberSetting**](VoipEmergencyNumberSetting.md) |  | 

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


## ModifyVoipTelephoneBookTemplate

> OperationResponseWithoutResult ModifyVoipTelephoneBookTemplate(ctx, omadacId, siteTemplateId, contactId).VoipContactPersonSettings(voipContactPersonSettings).Execute()

Modify an existing Contact Person in siteTemplate



### Example

```go
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
	contactId := "contactId_example" // string | Voip Contact Person ID
	voipContactPersonSettings := *openapiclient.NewVoipContactPersonSettings() // VoipContactPersonSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPTemplateAPI.ModifyVoipTelephoneBookTemplate(context.Background(), omadacId, siteTemplateId, contactId).VoipContactPersonSettings(voipContactPersonSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPTemplateAPI.ModifyVoipTelephoneBookTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoipTelephoneBookTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPTemplateAPI.ModifyVoipTelephoneBookTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**contactId** | **string** | Voip Contact Person ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipTelephoneBookTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **voipContactPersonSettings** | [**VoipContactPersonSettings**](VoipContactPersonSettings.md) |  | 

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


# \VoIPAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCallForwardingRule1**](VoIPAPI.md#AddCallForwardingRule1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/call-forwarding | Add call forwarding rules
[**BatchDeleteVoipTelephoneBook**](VoIPAPI.md#BatchDeleteVoipTelephoneBook) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/voip/telephone-book/batch-delete | Batch Delete Contact Person
[**BatchModifyVoipDeviceSetting**](VoIPAPI.md#BatchModifyVoipDeviceSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/batch-modify | Batch Modify Voip Device Setting
[**BindTelephoneNumber**](VoIPAPI.md#BindTelephoneNumber) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/{deviceMac}/bind-telephone-number | Bind Telephone Number
[**ClearVoipCallLog**](VoIPAPI.md#ClearVoipCallLog) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/voip/call-log/delete-all | Delete Voip CallLog Data List
[**CreateCallBlockingProfile1**](VoIPAPI.md#CreateCallBlockingProfile1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/call-blocking | Create new call blocking profile
[**CreateDigitMapProfile**](VoIPAPI.md#CreateDigitMapProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/digit-map | Create new digit map profile
[**CreateDigitMapProfileTemplate**](VoIPAPI.md#CreateDigitMapProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/digit-map | Create new digit map profile template
[**CreateProviderProfile1**](VoIPAPI.md#CreateProviderProfile1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/provider-profiles | Create new provider profile
[**CreateVoipTelephoneBook**](VoIPAPI.md#CreateVoipTelephoneBook) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/voip/telephone-book | Create New Contact Person
[**DeleteCallBlockingProfile**](VoIPAPI.md#DeleteCallBlockingProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/call-blocking/delete | Delete call blocking profile
[**DeleteCallForwardingRule1**](VoIPAPI.md#DeleteCallForwardingRule1) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/call-forwarding | Delete call forwarding rules
[**DeleteDigitMapProfile**](VoIPAPI.md#DeleteDigitMapProfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/digit-map/delete | Delete digit map profile
[**DeleteDigitMapProfileTemplate**](VoIPAPI.md#DeleteDigitMapProfileTemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/digit-map/delete | Delete digit map profile template
[**DeleteProviderProfiles**](VoIPAPI.md#DeleteProviderProfiles) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/provider-profiles/delete | Delete provider profile(s)
[**DeleteTelephoneNumber**](VoIPAPI.md#DeleteTelephoneNumber) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/{deviceMac}/telephone-number/{phoneNumberId} | Delete Telephone Number
[**DeleteTelephoneNumber1**](VoIPAPI.md#DeleteTelephoneNumber1) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voice-mail/all | Delete all voice mail
[**DeleteVoipMail**](VoIPAPI.md#DeleteVoipMail) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voice-mail/batch-delete | Batch delete voice mail
[**DeleteVoipTelephoneBook**](VoIPAPI.md#DeleteVoipTelephoneBook) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/voip/telephone-book/{contactId} | Delete Contact Person
[**DownloadTelephoneNumberImportResult**](VoIPAPI.md#DownloadTelephoneNumberImportResult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/telephone-number/import-result | Download telephone number list import result
[**DownloadVoiceMail**](VoIPAPI.md#DownloadVoiceMail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voice-mail/{voiceId}/download | Get voice mail
[**ExportTelephoneNumberListToFile**](VoIPAPI.md#ExportTelephoneNumberListToFile) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/telephone-number/export | Export telephone number list to file
[**GetAllCallBlockingProfiles**](VoIPAPI.md#GetAllCallBlockingProfiles) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/call-blocking | Get All Call Blocking Profiles
[**GetAllDigitMapProfiles**](VoIPAPI.md#GetAllDigitMapProfiles) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/digit-map | Get All Digit Map Profiles
[**GetCallBlockingProfiles1**](VoIPAPI.md#GetCallBlockingProfiles1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/call-blocking | Get call blocking profile list
[**GetCallForwardingRulesGrid**](VoIPAPI.md#GetCallForwardingRulesGrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/call-forwarding/grid | Get call forwarding rules
[**GetDigitMapProfiles**](VoIPAPI.md#GetDigitMapProfiles) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/digit-map | Get digit map profile list
[**GetDigitMapProfilesTemplate**](VoIPAPI.md#GetDigitMapProfilesTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/digit-map | Get digit map profile list template
[**GetDndSettings1**](VoIPAPI.md#GetDndSettings1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/dnd/settings | Get DND settings
[**GetGridProviderProfileList1**](VoIPAPI.md#GetGridProviderProfileList1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/grid/provider-profiles | Get the grid of provider profile list
[**GetGridSimplifiedVoipDevices**](VoIPAPI.md#GetGridSimplifiedVoipDevices) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/simplied-list | Get Grid of Simplified VoIP Devices
[**GetGridTelephoneNumberRegistrationResults**](VoIPAPI.md#GetGridTelephoneNumberRegistrationResults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/grid/telephone-number | Get Grid of Telephone Number Registration Results
[**GetTelephoneNumberBatchConfigList**](VoIPAPI.md#GetTelephoneNumberBatchConfigList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/telephone-number/batch-config-list | Get Telephone Number Config List
[**GetTelephoneNumberRegistrationResults**](VoIPAPI.md#GetTelephoneNumberRegistrationResults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/{deviceMac}/registration-result | Get Telephone Number Registration Results
[**GetUsbInfo**](VoIPAPI.md#GetUsbInfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voice-mail/settings/usbInfo | Get USB info
[**GetVoipCallLog**](VoIPAPI.md#GetVoipCallLog) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/voip/call-log | Get Voip CallLog Data List
[**GetVoipCallLogEnable1**](VoIPAPI.md#GetVoipCallLogEnable1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/voip/call-log-enable | Get Voip CallLog Setting
[**GetVoipDeviceList**](VoIPAPI.md#GetVoipDeviceList) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices | Get Voip Device List
[**GetVoipEmergency1**](VoIPAPI.md#GetVoipEmergency1) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/voip/emergency-number-settings | Get Voip Emergency Number Info
[**GetVoipMailGrid**](VoIPAPI.md#GetVoipMailGrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voice-mail/grid | Get voice mail list
[**GetVoipMailSettings**](VoIPAPI.md#GetVoipMailSettings) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voice-mail/settings | Get voice mail settings
[**GetVoipTelephoneBook**](VoIPAPI.md#GetVoipTelephoneBook) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/voip/telephone-book | Get Voip TelephoneBook List
[**ImportTelephoneNumberListFromFile**](VoIPAPI.md#ImportTelephoneNumberListFromFile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/telephone-number/import | Import telephone number list from file
[**ModifyCallBlockingProfile1**](VoIPAPI.md#ModifyCallBlockingProfile1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/call-blocking/{profileId} | Modify a call blocking profile
[**ModifyCallForwardingRule1**](VoIPAPI.md#ModifyCallForwardingRule1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/call-forwarding | Modify call forwarding rules
[**ModifyDigitMapProfile**](VoIPAPI.md#ModifyDigitMapProfile) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/digit-map/{profileId} | Modify a digit map profile
[**ModifyDigitMapProfileTemplate**](VoIPAPI.md#ModifyDigitMapProfileTemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/voip/digit-map/{profileId} | Modify a digit map profile template
[**ModifyDndSettings1**](VoIPAPI.md#ModifyDndSettings1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/dnd/settings | Modify DND settings
[**ModifyProviderProfile1**](VoIPAPI.md#ModifyProviderProfile1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/provider-profiles/{profileId} | Modify a provider profile
[**ModifyVoipCallLogEnable1**](VoIPAPI.md#ModifyVoipCallLogEnable1) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/voip/call-log-enable | Modify Voip CallLog Setting
[**ModifyVoipDeviceOsgSetting**](VoIPAPI.md#ModifyVoipDeviceOsgSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/osg/{deviceMac} | Modify Voip Device OSG Setting
[**ModifyVoipDeviceSetting**](VoIPAPI.md#ModifyVoipDeviceSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/{deviceMac} | Modify Voip Device AP Setting
[**ModifyVoipEmergency1**](VoIPAPI.md#ModifyVoipEmergency1) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/voip/emergency-number-settings | Modify Voip Emergency Number Setting
[**ModifyVoipTelephoneBook**](VoIPAPI.md#ModifyVoipTelephoneBook) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/voip/telephone-book/{contactId} | Modify Contact Person
[**SetVoipMailSettings**](VoIPAPI.md#SetVoipMailSettings) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voice-mail/settings | Edit voice mail settings
[**UpdateTelephoneNumberList**](VoIPAPI.md#UpdateTelephoneNumberList) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/voip/voip-devices/telephone-number/batch-config-list | Update Telephone Number List



## AddCallForwardingRule1

> OperationResponseAddCallForwardingRuleResp AddCallForwardingRule1(ctx, omadacId, siteId).CallForwardingRule(callForwardingRule).Execute()

Add call forwarding rules



### Example

```go
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
	callForwardingRule := *openapiclient.NewCallForwardingRule(int32(123), "DestNumber_example", false, int32(123)) // CallForwardingRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.AddCallForwardingRule1(context.Background(), omadacId, siteId).CallForwardingRule(callForwardingRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.AddCallForwardingRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCallForwardingRule1`: OperationResponseAddCallForwardingRuleResp
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.AddCallForwardingRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCallForwardingRule1Request struct via the builder pattern


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


## BatchDeleteVoipTelephoneBook

> OperationResponseWithoutResult BatchDeleteVoipTelephoneBook(ctx, omadacId, siteId).VoipTelephoneBookBatchSetting(voipTelephoneBookBatchSetting).Execute()

Batch Delete Contact Person



### Example

```go
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
	voipTelephoneBookBatchSetting := *openapiclient.NewVoipTelephoneBookBatchSetting([]string{"ContactIds_example"}, false, "SelectType_example") // VoipTelephoneBookBatchSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.BatchDeleteVoipTelephoneBook(context.Background(), omadacId, siteId).VoipTelephoneBookBatchSetting(voipTelephoneBookBatchSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.BatchDeleteVoipTelephoneBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteVoipTelephoneBook`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.BatchDeleteVoipTelephoneBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteVoipTelephoneBookRequest struct via the builder pattern


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


## BatchModifyVoipDeviceSetting

> OperationResponseListProfilesBindedDeviceInfo BatchModifyVoipDeviceSetting(ctx, omadacId, siteId).BatchModifyVoipDeviceSettingEntity(batchModifyVoipDeviceSettingEntity).Execute()

Batch Modify Voip Device Setting



### Example

```go
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
	batchModifyVoipDeviceSettingEntity := *openapiclient.NewBatchModifyVoipDeviceSettingEntity([]string{"DeviceMacList_example"}, false) // BatchModifyVoipDeviceSettingEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.BatchModifyVoipDeviceSetting(context.Background(), omadacId, siteId).BatchModifyVoipDeviceSettingEntity(batchModifyVoipDeviceSettingEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.BatchModifyVoipDeviceSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchModifyVoipDeviceSetting`: OperationResponseListProfilesBindedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.BatchModifyVoipDeviceSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchModifyVoipDeviceSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchModifyVoipDeviceSettingEntity** | [**BatchModifyVoipDeviceSettingEntity**](BatchModifyVoipDeviceSettingEntity.md) |  | 

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


## BindTelephoneNumber

> OperationResponseWithoutResult BindTelephoneNumber(ctx, omadacId, siteId, deviceMac).BindNumberList(bindNumberList).Execute()

Bind Telephone Number



### Example

```go
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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	bindNumberList := *openapiclient.NewBindNumberList([]openapiclient.TelephoneNumberWithoutStatusOpenApiVO{*openapiclient.NewTelephoneNumberWithoutStatusOpenApiVO("DeviceMac_example", "Number_example", "ProfileId_example")}) // BindNumberList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.BindTelephoneNumber(context.Background(), omadacId, siteId, deviceMac).BindNumberList(bindNumberList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.BindTelephoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindTelephoneNumber`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.BindTelephoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindTelephoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bindNumberList** | [**BindNumberList**](BindNumberList.md) |  | 

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


## ClearVoipCallLog

> OperationResponseWithoutResult ClearVoipCallLog(ctx, omadacId, siteId).Execute()

Delete Voip CallLog Data List



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.ClearVoipCallLog(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ClearVoipCallLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearVoipCallLog`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ClearVoipCallLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearVoipCallLogRequest struct via the builder pattern


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


## CreateCallBlockingProfile1

> OperationResponseWithoutResult CreateCallBlockingProfile1(ctx, omadacId, siteId).CreateCallBlockingProfileEntity(createCallBlockingProfileEntity).Execute()

Create new call blocking profile



### Example

```go
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
	createCallBlockingProfileEntity := *openapiclient.NewCreateCallBlockingProfileEntity("ProfileName_example") // CreateCallBlockingProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.CreateCallBlockingProfile1(context.Background(), omadacId, siteId).CreateCallBlockingProfileEntity(createCallBlockingProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.CreateCallBlockingProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCallBlockingProfile1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.CreateCallBlockingProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallBlockingProfile1Request struct via the builder pattern


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


## CreateDigitMapProfile

> OperationResponseWithoutResult CreateDigitMapProfile(ctx, omadacId, siteId).CreateDigitMapProfileEntity(createDigitMapProfileEntity).Execute()

Create new digit map profile



### Example

```go
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
	createDigitMapProfileEntity := *openapiclient.NewCreateDigitMapProfileEntity("DigitMap_example", "ProfileName_example") // CreateDigitMapProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.CreateDigitMapProfile(context.Background(), omadacId, siteId).CreateDigitMapProfileEntity(createDigitMapProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.CreateDigitMapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDigitMapProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.CreateDigitMapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitMapProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDigitMapProfileEntity** | [**CreateDigitMapProfileEntity**](CreateDigitMapProfileEntity.md) |  | 

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


## CreateDigitMapProfileTemplate

> OperationResponseWithoutResult CreateDigitMapProfileTemplate(ctx, omadacId, siteTemplateId).CreateDigitMapProfileEntity(createDigitMapProfileEntity).Execute()

Create new digit map profile template



### Example

```go
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
	createDigitMapProfileEntity := *openapiclient.NewCreateDigitMapProfileEntity("DigitMap_example", "ProfileName_example") // CreateDigitMapProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.CreateDigitMapProfileTemplate(context.Background(), omadacId, siteTemplateId).CreateDigitMapProfileEntity(createDigitMapProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.CreateDigitMapProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDigitMapProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.CreateDigitMapProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitMapProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDigitMapProfileEntity** | [**CreateDigitMapProfileEntity**](CreateDigitMapProfileEntity.md) |  | 

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


## CreateProviderProfile1

> OperationResponseWithoutResult CreateProviderProfile1(ctx, omadacId, siteId).CreateProviderProfileEntity(createProviderProfileEntity).Execute()

Create new provider profile



### Example

```go
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
	createProviderProfileEntity := *openapiclient.NewCreateProviderProfileEntity("ProfileName_example", *openapiclient.NewProviderSettingVO(int32(123))) // CreateProviderProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.CreateProviderProfile1(context.Background(), omadacId, siteId).CreateProviderProfileEntity(createProviderProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.CreateProviderProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProviderProfile1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.CreateProviderProfile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProviderProfile1Request struct via the builder pattern


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


## CreateVoipTelephoneBook

> OperationResponseWithoutResult CreateVoipTelephoneBook(ctx, omadacId, siteId).VoipContactPersonSettings(voipContactPersonSettings).Execute()

Create New Contact Person



### Example

```go
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
	voipContactPersonSettings := *openapiclient.NewVoipContactPersonSettings() // VoipContactPersonSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.CreateVoipTelephoneBook(context.Background(), omadacId, siteId).VoipContactPersonSettings(voipContactPersonSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.CreateVoipTelephoneBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoipTelephoneBook`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.CreateVoipTelephoneBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoipTelephoneBookRequest struct via the builder pattern


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


## DeleteCallBlockingProfile

> OperationResponseListProfilesBindedDeviceInfo DeleteCallBlockingProfile(ctx, omadacId, siteId).DeleteCallBlockingProfileEntity(deleteCallBlockingProfileEntity).Execute()

Delete call blocking profile



### Example

```go
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
	deleteCallBlockingProfileEntity := *openapiclient.NewDeleteCallBlockingProfileEntity("ProfileId_example", false) // DeleteCallBlockingProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DeleteCallBlockingProfile(context.Background(), omadacId, siteId).DeleteCallBlockingProfileEntity(deleteCallBlockingProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteCallBlockingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCallBlockingProfile`: OperationResponseListProfilesBindedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteCallBlockingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallBlockingProfileRequest struct via the builder pattern


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


## DeleteCallForwardingRule1

> OperationResponseVoid DeleteCallForwardingRule1(ctx, omadacId, siteId).DeleteCallForwardingRules(deleteCallForwardingRules).Execute()

Delete call forwarding rules



### Example

```go
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
	deleteCallForwardingRules := *openapiclient.NewDeleteCallForwardingRules() // DeleteCallForwardingRules | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DeleteCallForwardingRule1(context.Background(), omadacId, siteId).DeleteCallForwardingRules(deleteCallForwardingRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteCallForwardingRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCallForwardingRule1`: OperationResponseVoid
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteCallForwardingRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallForwardingRule1Request struct via the builder pattern


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


## DeleteDigitMapProfile

> OperationResponseListProfilesBindedDeviceInfo DeleteDigitMapProfile(ctx, omadacId, siteId).DeleteDigitMapProfileEntity(deleteDigitMapProfileEntity).Execute()

Delete digit map profile



### Example

```go
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
	deleteDigitMapProfileEntity := *openapiclient.NewDeleteDigitMapProfileEntity("ProfileId_example", false) // DeleteDigitMapProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DeleteDigitMapProfile(context.Background(), omadacId, siteId).DeleteDigitMapProfileEntity(deleteDigitMapProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteDigitMapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDigitMapProfile`: OperationResponseListProfilesBindedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteDigitMapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitMapProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteDigitMapProfileEntity** | [**DeleteDigitMapProfileEntity**](DeleteDigitMapProfileEntity.md) |  | 

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


## DeleteDigitMapProfileTemplate

> OperationResponseListProfilesBindedDeviceInfo DeleteDigitMapProfileTemplate(ctx, omadacId, siteTemplateId).DeleteDigitMapProfileEntity(deleteDigitMapProfileEntity).Execute()

Delete digit map profile template



### Example

```go
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
	deleteDigitMapProfileEntity := *openapiclient.NewDeleteDigitMapProfileEntity("ProfileId_example", false) // DeleteDigitMapProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DeleteDigitMapProfileTemplate(context.Background(), omadacId, siteTemplateId).DeleteDigitMapProfileEntity(deleteDigitMapProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteDigitMapProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDigitMapProfileTemplate`: OperationResponseListProfilesBindedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteDigitMapProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitMapProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteDigitMapProfileEntity** | [**DeleteDigitMapProfileEntity**](DeleteDigitMapProfileEntity.md) |  | 

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


## DeleteProviderProfiles

> OperationResponseListProfilesBindedDeviceInfo DeleteProviderProfiles(ctx, omadacId, siteId).ConfirmBindedDevicesEntity(confirmBindedDevicesEntity).Execute()

Delete provider profile(s)



### Example

```go
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
	confirmBindedDevicesEntity := *openapiclient.NewConfirmBindedDevicesEntity([]string{"ProfileIds_example"}, false) // ConfirmBindedDevicesEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DeleteProviderProfiles(context.Background(), omadacId, siteId).ConfirmBindedDevicesEntity(confirmBindedDevicesEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteProviderProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProviderProfiles`: OperationResponseListProfilesBindedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteProviderProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **confirmBindedDevicesEntity** | [**ConfirmBindedDevicesEntity**](ConfirmBindedDevicesEntity.md) |  | 

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


## DeleteTelephoneNumber

> OperationResponseWithoutResult DeleteTelephoneNumber(ctx, omadacId, siteId, deviceMac, phoneNumberId).DeleteTelephoneNumber(deleteTelephoneNumber).Execute()

Delete Telephone Number



### Example

```go
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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	phoneNumberId := "phoneNumberId_example" // string | phone Number ID
	deleteTelephoneNumber := *openapiclient.NewDeleteTelephoneNumber(false) // DeleteTelephoneNumber | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DeleteTelephoneNumber(context.Background(), omadacId, siteId, deviceMac, phoneNumberId).DeleteTelephoneNumber(deleteTelephoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteTelephoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTelephoneNumber`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteTelephoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**phoneNumberId** | **string** | phone Number ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTelephoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **deleteTelephoneNumber** | [**DeleteTelephoneNumber**](DeleteTelephoneNumber.md) |  | 

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


## DeleteTelephoneNumber1

> OperationResponseWithoutResult DeleteTelephoneNumber1(ctx, omadacId, siteId).Execute()

Delete all voice mail



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.DeleteTelephoneNumber1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteTelephoneNumber1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTelephoneNumber1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteTelephoneNumber1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTelephoneNumber1Request struct via the builder pattern


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


## DeleteVoipMail

> OperationResponseWithoutResult DeleteVoipMail(ctx, omadacId, siteId).VoiceMailBatchDelete(voiceMailBatchDelete).Execute()

Batch delete voice mail



### Example

```go
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
	voiceMailBatchDelete := *openapiclient.NewVoiceMailBatchDelete([]string{"VoiceMailList_example"}) // VoiceMailBatchDelete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DeleteVoipMail(context.Background(), omadacId, siteId).VoiceMailBatchDelete(voiceMailBatchDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteVoipMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVoipMail`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteVoipMail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoipMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voiceMailBatchDelete** | [**VoiceMailBatchDelete**](VoiceMailBatchDelete.md) |  | 

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


## DeleteVoipTelephoneBook

> OperationResponseWithoutResult DeleteVoipTelephoneBook(ctx, omadacId, siteId, contactId).Execute()

Delete Contact Person



### Example

```go
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
	contactId := "contactId_example" // string | Voip Contact Person ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DeleteVoipTelephoneBook(context.Background(), omadacId, siteId, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DeleteVoipTelephoneBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVoipTelephoneBook`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DeleteVoipTelephoneBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**contactId** | **string** | Voip Contact Person ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoipTelephoneBookRequest struct via the builder pattern


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


## DownloadTelephoneNumberImportResult

> OperationResponse DownloadTelephoneNumberImportResult(ctx, omadacId, siteId).Execute()

Download telephone number list import result



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.DownloadTelephoneNumberImportResult(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DownloadTelephoneNumberImportResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadTelephoneNumberImportResult`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DownloadTelephoneNumberImportResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTelephoneNumberImportResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadVoiceMail

> OperationResponse DownloadVoiceMail(ctx, omadacId, siteId, voiceId).Execute()

Get voice mail



### Example

```go
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
	voiceId := "voiceId_example" // string | voiceId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.DownloadVoiceMail(context.Background(), omadacId, siteId, voiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.DownloadVoiceMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadVoiceMail`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.DownloadVoiceMail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**voiceId** | **string** | voiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadVoiceMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTelephoneNumberListToFile

> OperationResponseWithoutResult ExportTelephoneNumberListToFile(ctx, omadacId, siteId).Execute()

Export telephone number list to file



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.ExportTelephoneNumberListToFile(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ExportTelephoneNumberListToFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportTelephoneNumberListToFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ExportTelephoneNumberListToFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportTelephoneNumberListToFileRequest struct via the builder pattern


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


## GetAllCallBlockingProfiles

> OperationResponseListSimplifiedCallBlockingProfile GetAllCallBlockingProfiles(ctx, omadacId, siteId).Execute()

Get All Call Blocking Profiles



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetAllCallBlockingProfiles(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetAllCallBlockingProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCallBlockingProfiles`: OperationResponseListSimplifiedCallBlockingProfile
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetAllCallBlockingProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCallBlockingProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListSimplifiedCallBlockingProfile**](OperationResponseListSimplifiedCallBlockingProfile.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDigitMapProfiles

> OperationResponseListSimplifiedDigitMapProfile GetAllDigitMapProfiles(ctx, omadacId, siteId).Execute()

Get All Digit Map Profiles



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetAllDigitMapProfiles(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetAllDigitMapProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDigitMapProfiles`: OperationResponseListSimplifiedDigitMapProfile
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetAllDigitMapProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDigitMapProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListSimplifiedDigitMapProfile**](OperationResponseListSimplifiedDigitMapProfile.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallBlockingProfiles1

> OperationResponseListCallBlockingProfileEntity GetCallBlockingProfiles1(ctx, omadacId, siteId).Execute()

Get call blocking profile list



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetCallBlockingProfiles1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetCallBlockingProfiles1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallBlockingProfiles1`: OperationResponseListCallBlockingProfileEntity
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetCallBlockingProfiles1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallBlockingProfiles1Request struct via the builder pattern


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


## GetCallForwardingRulesGrid

> OperationResponseCallForwardingRulesGrid GetCallForwardingRulesGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get call forwarding rules



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetCallForwardingRulesGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetCallForwardingRulesGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallForwardingRulesGrid`: OperationResponseCallForwardingRulesGrid
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetCallForwardingRulesGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallForwardingRulesGridRequest struct via the builder pattern


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


## GetDigitMapProfiles

> OperationResponseListDigitMapProfileInfo GetDigitMapProfiles(ctx, omadacId, siteId).Execute()

Get digit map profile list



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetDigitMapProfiles(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetDigitMapProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDigitMapProfiles`: OperationResponseListDigitMapProfileInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetDigitMapProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitMapProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListDigitMapProfileInfo**](OperationResponseListDigitMapProfileInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDigitMapProfilesTemplate

> OperationResponseListDigitMapProfileInfo GetDigitMapProfilesTemplate(ctx, omadacId, siteTemplateId).Execute()

Get digit map profile list template



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetDigitMapProfilesTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetDigitMapProfilesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDigitMapProfilesTemplate`: OperationResponseListDigitMapProfileInfo
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetDigitMapProfilesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitMapProfilesTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListDigitMapProfileInfo**](OperationResponseListDigitMapProfileInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDndSettings1

> OperationResponseDndSettingEntity GetDndSettings1(ctx, omadacId, siteId).Execute()

Get DND settings



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetDndSettings1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetDndSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDndSettings1`: OperationResponseDndSettingEntity
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetDndSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDndSettings1Request struct via the builder pattern


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


## GetGridProviderProfileList1

> OperationResponseGridVOProviderProfileEntity GetGridProviderProfileList1(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the grid of provider profile list



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetGridProviderProfileList1(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetGridProviderProfileList1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridProviderProfileList1`: OperationResponseGridVOProviderProfileEntity
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetGridProviderProfileList1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridProviderProfileList1Request struct via the builder pattern


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


## GetGridSimplifiedVoipDevices

> OperationResponseGridVOSimplifiedVoipDeviceOpenApiVO GetGridSimplifiedVoipDevices(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Grid of Simplified VoIP Devices



### Example

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
	resp, r, err := apiClient.VoIPAPI.GetGridSimplifiedVoipDevices(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetGridSimplifiedVoipDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSimplifiedVoipDevices`: OperationResponseGridVOSimplifiedVoipDeviceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetGridSimplifiedVoipDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSimplifiedVoipDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOSimplifiedVoipDeviceOpenApiVO**](OperationResponseGridVOSimplifiedVoipDeviceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridTelephoneNumberRegistrationResults

> OperationResponseGridVODeviceTelephoneNumber GetGridTelephoneNumberRegistrationResults(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Grid of Telephone Number Registration Results



### Example

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
	resp, r, err := apiClient.VoIPAPI.GetGridTelephoneNumberRegistrationResults(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetGridTelephoneNumberRegistrationResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridTelephoneNumberRegistrationResults`: OperationResponseGridVODeviceTelephoneNumber
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetGridTelephoneNumberRegistrationResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridTelephoneNumberRegistrationResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVODeviceTelephoneNumber**](OperationResponseGridVODeviceTelephoneNumber.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelephoneNumberBatchConfigList

> OperationResponseGridVOTelephoneNumberWithStatusOpenApiVO GetTelephoneNumberBatchConfigList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Telephone Number Config List



### Example

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
	resp, r, err := apiClient.VoIPAPI.GetTelephoneNumberBatchConfigList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetTelephoneNumberBatchConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelephoneNumberBatchConfigList`: OperationResponseGridVOTelephoneNumberWithStatusOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetTelephoneNumberBatchConfigList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelephoneNumberBatchConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOTelephoneNumberWithStatusOpenApiVO**](OperationResponseGridVOTelephoneNumberWithStatusOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelephoneNumberRegistrationResults

> OperationResponseNumberRegistrationResult GetTelephoneNumberRegistrationResults(ctx, omadacId, siteId, deviceMac).Execute()

Get Telephone Number Registration Results



### Example

```go
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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.GetTelephoneNumberRegistrationResults(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetTelephoneNumberRegistrationResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelephoneNumberRegistrationResults`: OperationResponseNumberRegistrationResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetTelephoneNumberRegistrationResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelephoneNumberRegistrationResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseNumberRegistrationResult**](OperationResponseNumberRegistrationResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsbInfo

> OperationResponseUsbInfoListResponse GetUsbInfo(ctx, omadacId, siteId).Execute()

Get USB info



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetUsbInfo(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetUsbInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsbInfo`: OperationResponseUsbInfoListResponse
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetUsbInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsbInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUsbInfoListResponse**](OperationResponseUsbInfoListResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoipCallLog

> OperationResponseCallLogGridVOVoipCallLog GetVoipCallLog(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsDateTime(sortsDateTime).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()

Get Voip CallLog Data List



### Example

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
	sortsDateTime := "sortsDateTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	filtersStatus := "filtersStatus_example" // string | Filter query parameters, support field status: 0: incoming, 1: outgoing, 3: transfer, 4: miss，5: rejected (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field deviceNumber/telephonyDevice (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.GetVoipCallLog(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsDateTime(sortsDateTime).FiltersStatus(filtersStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetVoipCallLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipCallLog`: OperationResponseCallLogGridVOVoipCallLog
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetVoipCallLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipCallLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 


 **sortsDateTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **filtersStatus** | **string** | Filter query parameters, support field status: 0: incoming, 1: outgoing, 3: transfer, 4: miss，5: rejected | 
 **searchKey** | **string** | Fuzzy query parameters, support field deviceNumber/telephonyDevice | 

### Return type

[**OperationResponseCallLogGridVOVoipCallLog**](OperationResponseCallLogGridVOVoipCallLog.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoipCallLogEnable1

> OperationResponseVoipCallLogEnableSetting GetVoipCallLogEnable1(ctx, omadacId, siteId).Execute()

Get Voip CallLog Setting



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetVoipCallLogEnable1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetVoipCallLogEnable1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipCallLogEnable1`: OperationResponseVoipCallLogEnableSetting
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetVoipCallLogEnable1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipCallLogEnable1Request struct via the builder pattern


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


## GetVoipDeviceList

> OperationResponseGridVOVoipDevice GetVoipDeviceList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Voip Device List



### Example

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
	resp, r, err := apiClient.VoIPAPI.GetVoipDeviceList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetVoipDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipDeviceList`: OperationResponseGridVOVoipDevice
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetVoipDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOVoipDevice**](OperationResponseGridVOVoipDevice.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoipEmergency1

> OperationResponseVoipEmergencyNumberSetting GetVoipEmergency1(ctx, omadacId, siteId).Execute()

Get Voip Emergency Number Info



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetVoipEmergency1(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetVoipEmergency1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipEmergency1`: OperationResponseVoipEmergencyNumberSetting
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetVoipEmergency1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipEmergency1Request struct via the builder pattern


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


## GetVoipMailGrid

> OperationResponseGridVOVoiceMail GetVoipMailGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get voice mail list



### Example

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
	resp, r, err := apiClient.VoIPAPI.GetVoipMailGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetVoipMailGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipMailGrid`: OperationResponseGridVOVoiceMail
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetVoipMailGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipMailGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOVoiceMail**](OperationResponseGridVOVoiceMail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoipMailSettings

> OperationResponseVoiceMailSettingResponse GetVoipMailSettings(ctx, omadacId, siteId).Execute()

Get voice mail settings



### Example

```go
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
	resp, r, err := apiClient.VoIPAPI.GetVoipMailSettings(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetVoipMailSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipMailSettings`: OperationResponseVoiceMailSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetVoipMailSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipMailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseVoiceMailSettingResponse**](OperationResponseVoiceMailSettingResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoipTelephoneBook

> OperationResponseGridVOVoipTelephoneBookSetting GetVoipTelephoneBook(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Voip TelephoneBook List



### Example

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
	resp, r, err := apiClient.VoIPAPI.GetVoipTelephoneBook(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.GetVoipTelephoneBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoipTelephoneBook`: OperationResponseGridVOVoipTelephoneBookSetting
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.GetVoipTelephoneBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoipTelephoneBookRequest struct via the builder pattern


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


## ImportTelephoneNumberListFromFile

> OperationResponseWithoutResult ImportTelephoneNumberListFromFile(ctx, omadacId, siteId).UploadCertificateRequest(uploadCertificateRequest).Execute()

Import telephone number list from file



### Example

```go
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
	uploadCertificateRequest := *openapiclient.NewUploadCertificateRequest("TODO") // UploadCertificateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ImportTelephoneNumberListFromFile(context.Background(), omadacId, siteId).UploadCertificateRequest(uploadCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ImportTelephoneNumberListFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportTelephoneNumberListFromFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ImportTelephoneNumberListFromFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportTelephoneNumberListFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadCertificateRequest** | [**UploadCertificateRequest**](UploadCertificateRequest.md) |  | 

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


## ModifyCallBlockingProfile1

> OperationResponseWithoutResult ModifyCallBlockingProfile1(ctx, omadacId, siteId, profileId).ModifyCallBlockingProfileEntity(modifyCallBlockingProfileEntity).Execute()

Modify a call blocking profile



### Example

```go
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
	modifyCallBlockingProfileEntity := *openapiclient.NewModifyCallBlockingProfileEntity() // ModifyCallBlockingProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyCallBlockingProfile1(context.Background(), omadacId, siteId, profileId).ModifyCallBlockingProfileEntity(modifyCallBlockingProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyCallBlockingProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCallBlockingProfile1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyCallBlockingProfile1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyCallBlockingProfile1Request struct via the builder pattern


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


## ModifyCallForwardingRule1

> OperationResponseVoid ModifyCallForwardingRule1(ctx, omadacId, siteId).CallForwardingRule(callForwardingRule).Execute()

Modify call forwarding rules



### Example

```go
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
	callForwardingRule := *openapiclient.NewCallForwardingRule(int32(123), "DestNumber_example", false, int32(123)) // CallForwardingRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyCallForwardingRule1(context.Background(), omadacId, siteId).CallForwardingRule(callForwardingRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyCallForwardingRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCallForwardingRule1`: OperationResponseVoid
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyCallForwardingRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCallForwardingRule1Request struct via the builder pattern


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


## ModifyDigitMapProfile

> OperationResponseWithoutResult ModifyDigitMapProfile(ctx, omadacId, siteId, profileId).ModifyDigitMapProfileEntity(modifyDigitMapProfileEntity).Execute()

Modify a digit map profile



### Example

```go
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
	profileId := "profileId_example" // string | Profile ID. Default digit map profile cannot be modified.
	modifyDigitMapProfileEntity := *openapiclient.NewModifyDigitMapProfileEntity() // ModifyDigitMapProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyDigitMapProfile(context.Background(), omadacId, siteId, profileId).ModifyDigitMapProfileEntity(modifyDigitMapProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyDigitMapProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDigitMapProfile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyDigitMapProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**profileId** | **string** | Profile ID. Default digit map profile cannot be modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDigitMapProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyDigitMapProfileEntity** | [**ModifyDigitMapProfileEntity**](ModifyDigitMapProfileEntity.md) |  | 

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


## ModifyDigitMapProfileTemplate

> OperationResponseWithoutResult ModifyDigitMapProfileTemplate(ctx, omadacId, siteTemplateId, profileId).ModifyDigitMapProfileEntity(modifyDigitMapProfileEntity).Execute()

Modify a digit map profile template



### Example

```go
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
	profileId := "profileId_example" // string | Profile ID. Default digit map profile cannot be modified.
	modifyDigitMapProfileEntity := *openapiclient.NewModifyDigitMapProfileEntity() // ModifyDigitMapProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyDigitMapProfileTemplate(context.Background(), omadacId, siteTemplateId, profileId).ModifyDigitMapProfileEntity(modifyDigitMapProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyDigitMapProfileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDigitMapProfileTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyDigitMapProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**profileId** | **string** | Profile ID. Default digit map profile cannot be modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDigitMapProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyDigitMapProfileEntity** | [**ModifyDigitMapProfileEntity**](ModifyDigitMapProfileEntity.md) |  | 

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


## ModifyDndSettings1

> OperationResponseVoid ModifyDndSettings1(ctx, omadacId, siteId).DndSettingEntity(dndSettingEntity).Execute()

Modify DND settings



### Example

```go
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
	dndSettingEntity := *openapiclient.NewDndSettingEntity(false) // DndSettingEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyDndSettings1(context.Background(), omadacId, siteId).DndSettingEntity(dndSettingEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyDndSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDndSettings1`: OperationResponseVoid
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyDndSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDndSettings1Request struct via the builder pattern


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


## ModifyProviderProfile1

> OperationResponseWithoutResult ModifyProviderProfile1(ctx, omadacId, siteId, profileId).ModifyProviderProfileEntity(modifyProviderProfileEntity).Execute()

Modify a provider profile



### Example

```go
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
	modifyProviderProfileEntity := *openapiclient.NewModifyProviderProfileEntity() // ModifyProviderProfileEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyProviderProfile1(context.Background(), omadacId, siteId, profileId).ModifyProviderProfileEntity(modifyProviderProfileEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyProviderProfile1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyProviderProfile1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyProviderProfile1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyProviderProfile1Request struct via the builder pattern


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


## ModifyVoipCallLogEnable1

> OperationResponseWithoutResult ModifyVoipCallLogEnable1(ctx, omadacId, siteId).VoipCallLogEnableSetting(voipCallLogEnableSetting).Execute()

Modify Voip CallLog Setting



### Example

```go
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
	voipCallLogEnableSetting := *openapiclient.NewVoipCallLogEnableSetting(false) // VoipCallLogEnableSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyVoipCallLogEnable1(context.Background(), omadacId, siteId).VoipCallLogEnableSetting(voipCallLogEnableSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyVoipCallLogEnable1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoipCallLogEnable1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyVoipCallLogEnable1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipCallLogEnable1Request struct via the builder pattern


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


## ModifyVoipDeviceOsgSetting

> OperationResponseModifyConfirmResultOpenApiVO ModifyVoipDeviceOsgSetting(ctx, omadacId, siteId, deviceMac).ModifyVoipDeviceOsgSettingEntity(modifyVoipDeviceOsgSettingEntity).Execute()

Modify Voip Device OSG Setting



### Example

```go
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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	modifyVoipDeviceOsgSettingEntity := *openapiclient.NewModifyVoipDeviceOsgSettingEntity(false) // ModifyVoipDeviceOsgSettingEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyVoipDeviceOsgSetting(context.Background(), omadacId, siteId, deviceMac).ModifyVoipDeviceOsgSettingEntity(modifyVoipDeviceOsgSettingEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyVoipDeviceOsgSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoipDeviceOsgSetting`: OperationResponseModifyConfirmResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyVoipDeviceOsgSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipDeviceOsgSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyVoipDeviceOsgSettingEntity** | [**ModifyVoipDeviceOsgSettingEntity**](ModifyVoipDeviceOsgSettingEntity.md) |  | 

### Return type

[**OperationResponseModifyConfirmResultOpenApiVO**](OperationResponseModifyConfirmResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyVoipDeviceSetting

> OperationResponseModifyConfirmResultOpenApiVO ModifyVoipDeviceSetting(ctx, omadacId, siteId, deviceMac).ModifyVoipDeviceApSettingEntity(modifyVoipDeviceApSettingEntity).Execute()

Modify Voip Device AP Setting



### Example

```go
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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	modifyVoipDeviceApSettingEntity := *openapiclient.NewModifyVoipDeviceApSettingEntity(false) // ModifyVoipDeviceApSettingEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyVoipDeviceSetting(context.Background(), omadacId, siteId, deviceMac).ModifyVoipDeviceApSettingEntity(modifyVoipDeviceApSettingEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyVoipDeviceSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoipDeviceSetting`: OperationResponseModifyConfirmResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyVoipDeviceSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipDeviceSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyVoipDeviceApSettingEntity** | [**ModifyVoipDeviceApSettingEntity**](ModifyVoipDeviceApSettingEntity.md) |  | 

### Return type

[**OperationResponseModifyConfirmResultOpenApiVO**](OperationResponseModifyConfirmResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyVoipEmergency1

> OperationResponseWithoutResult ModifyVoipEmergency1(ctx, omadacId, siteId).VoipEmergencyNumberSetting(voipEmergencyNumberSetting).Execute()

Modify Voip Emergency Number Setting



### Example

```go
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
	voipEmergencyNumberSetting := *openapiclient.NewVoipEmergencyNumberSetting(false) // VoipEmergencyNumberSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyVoipEmergency1(context.Background(), omadacId, siteId).VoipEmergencyNumberSetting(voipEmergencyNumberSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyVoipEmergency1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoipEmergency1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyVoipEmergency1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipEmergency1Request struct via the builder pattern


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


## ModifyVoipTelephoneBook

> OperationResponseWithoutResult ModifyVoipTelephoneBook(ctx, omadacId, siteId, contactId).VoipContactPersonSettings(voipContactPersonSettings).Execute()

Modify Contact Person



### Example

```go
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
	contactId := "contactId_example" // string | Voip Contact Person ID
	voipContactPersonSettings := *openapiclient.NewVoipContactPersonSettings() // VoipContactPersonSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.ModifyVoipTelephoneBook(context.Background(), omadacId, siteId, contactId).VoipContactPersonSettings(voipContactPersonSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.ModifyVoipTelephoneBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyVoipTelephoneBook`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.ModifyVoipTelephoneBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**contactId** | **string** | Voip Contact Person ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVoipTelephoneBookRequest struct via the builder pattern


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


## SetVoipMailSettings

> OperationResponseWithoutResult SetVoipMailSettings(ctx, omadacId, siteId).VoiceMailSettingRequest(voiceMailSettingRequest).Execute()

Edit voice mail settings



### Example

```go
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
	voiceMailSettingRequest := *openapiclient.NewVoiceMailSettingRequest(false) // VoiceMailSettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.SetVoipMailSettings(context.Background(), omadacId, siteId).VoiceMailSettingRequest(voiceMailSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.SetVoipMailSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVoipMailSettings`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.SetVoipMailSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVoipMailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voiceMailSettingRequest** | [**VoiceMailSettingRequest**](VoiceMailSettingRequest.md) |  | 

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


## UpdateTelephoneNumberList

> OperationResponseWithoutResult UpdateTelephoneNumberList(ctx, omadacId, siteId).TelephoneNumberListUpdateOpenApiVO(telephoneNumberListUpdateOpenApiVO).Execute()

Update Telephone Number List



### Example

```go
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
	telephoneNumberListUpdateOpenApiVO := *openapiclient.NewTelephoneNumberListUpdateOpenApiVO() // TelephoneNumberListUpdateOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoIPAPI.UpdateTelephoneNumberList(context.Background(), omadacId, siteId).TelephoneNumberListUpdateOpenApiVO(telephoneNumberListUpdateOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoIPAPI.UpdateTelephoneNumberList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTelephoneNumberList`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `VoIPAPI.UpdateTelephoneNumberList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTelephoneNumberListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **telephoneNumberListUpdateOpenApiVO** | [**TelephoneNumberListUpdateOpenApiVO**](TelephoneNumberListUpdateOpenApiVO.md) |  | 

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


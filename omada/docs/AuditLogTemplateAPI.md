# AuditLogTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogSettingForMsp**](AuditLogTemplateAPI.md#getauditlogsettingformsp) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/site/audit-notification | Get site template audit log notification
[**ModifyAuditLogSettingSite**](AuditLogTemplateAPI.md#modifyauditlogsettingsite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/site/audit-notification | Modify site template audit log notification



## GetAuditLogSettingForMsp

> OperationResponseAuditLogNotificationSettingOpenApiVO GetAuditLogSettingForMsp(ctx, omadacId, siteTemplateId).Execute()

Get site template audit log notification



### Example

```go
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
	resp, r, err := apiClient.AuditLogTemplateAPI.GetAuditLogSettingForMsp(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogTemplateAPI.GetAuditLogSettingForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogSettingForMsp`: OperationResponseAuditLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `AuditLogTemplateAPI.GetAuditLogSettingForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogSettingForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseAuditLogNotificationSettingOpenApiVO**](OperationResponseAuditLogNotificationSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAuditLogSettingSite

> OperationResponseWithoutResult ModifyAuditLogSettingSite(ctx, omadacId, siteTemplateId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()

Modify site template audit log notification



### Example

```go
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
	auditLogNotificationSettingEditOpenApiVO := *openapiclient.NewAuditLogNotificationSettingEditOpenApiVO([]openapiclient.AuditLogNotificationEditOpenApiVO{*openapiclient.NewAuditLogNotificationEditOpenApiVO("DASHBOARD", false)}, *openapiclient.NewWebhookConfigEditOpenApiVO(false)) // AuditLogNotificationSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogTemplateAPI.ModifyAuditLogSettingSite(context.Background(), omadacId, siteTemplateId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogTemplateAPI.ModifyAuditLogSettingSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAuditLogSettingSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AuditLogTemplateAPI.ModifyAuditLogSettingSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAuditLogSettingSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **auditLogNotificationSettingEditOpenApiVO** | [**AuditLogNotificationSettingEditOpenApiVO**](AuditLogNotificationSettingEditOpenApiVO.md) |  | 

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


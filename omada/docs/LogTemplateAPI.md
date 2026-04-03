# LogTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogSettingForMsp**](LogTemplateAPI.md#getauditlogsettingformsp) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/site/audit-notification | Get site template audit log notification
[**GetLogSettingForSiteTemplate**](LogTemplateAPI.md#getlogsettingforsitetemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/site/log-notification | Get site template log notification
[**ModifyAuditLogSettingSite**](LogTemplateAPI.md#modifyauditlogsettingsite) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/site/audit-notification | Modify site template audit log notification
[**ModifyLogSettingSiteTemplate**](LogTemplateAPI.md#modifylogsettingsitetemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/site/log-notification | Modify site template log notification
[**ResetLogSettingSiteTemplate**](LogTemplateAPI.md#resetlogsettingsitetemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/site/reset/log-notification | Reset site template log notification



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
	resp, r, err := apiClient.LogTemplateAPI.GetAuditLogSettingForMsp(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogTemplateAPI.GetAuditLogSettingForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogSettingForMsp`: OperationResponseAuditLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogTemplateAPI.GetAuditLogSettingForMsp`: %v\n", resp)
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


## GetLogSettingForSiteTemplate

> OperationResponseLogNotificationSettingOpenApiVO GetLogSettingForSiteTemplate(ctx, omadacId, siteTemplateId).Execute()

Get site template log notification



### Example

```go
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
	resp, r, err := apiClient.LogTemplateAPI.GetLogSettingForSiteTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogTemplateAPI.GetLogSettingForSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogSettingForSiteTemplate`: OperationResponseLogNotificationSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `LogTemplateAPI.GetLogSettingForSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSettingForSiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseLogNotificationSettingOpenApiVO**](OperationResponseLogNotificationSettingOpenApiVO.md)

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
	resp, r, err := apiClient.LogTemplateAPI.ModifyAuditLogSettingSite(context.Background(), omadacId, siteTemplateId).AuditLogNotificationSettingEditOpenApiVO(auditLogNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogTemplateAPI.ModifyAuditLogSettingSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAuditLogSettingSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogTemplateAPI.ModifyAuditLogSettingSite`: %v\n", resp)
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


## ModifyLogSettingSiteTemplate

> OperationResponseWithoutResult ModifyLogSettingSiteTemplate(ctx, omadacId, siteTemplateId).LogNotificationSettingEditOpenApiVO(logNotificationSettingEditOpenApiVO).Execute()

Modify site template log notification



### Example

```go
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
	logNotificationSettingEditOpenApiVO := *openapiclient.NewLogNotificationSettingEditOpenApiVO([]openapiclient.LogNotificationEditOpenApiVO{*openapiclient.NewLogNotificationEditOpenApiVO(false, true, true, "LOGIN_OK")}) // LogNotificationSettingEditOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogTemplateAPI.ModifyLogSettingSiteTemplate(context.Background(), omadacId, siteTemplateId).LogNotificationSettingEditOpenApiVO(logNotificationSettingEditOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogTemplateAPI.ModifyLogSettingSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyLogSettingSiteTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogTemplateAPI.ModifyLogSettingSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLogSettingSiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **logNotificationSettingEditOpenApiVO** | [**LogNotificationSettingEditOpenApiVO**](LogNotificationSettingEditOpenApiVO.md) |  | 

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


## ResetLogSettingSiteTemplate

> OperationResponseWithoutResult ResetLogSettingSiteTemplate(ctx, omadacId, siteTemplateId).Execute()

Reset site template log notification



### Example

```go
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
	resp, r, err := apiClient.LogTemplateAPI.ResetLogSettingSiteTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogTemplateAPI.ResetLogSettingSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetLogSettingSiteTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `LogTemplateAPI.ResetLogSettingSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetLogSettingSiteTemplateRequest struct via the builder pattern


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


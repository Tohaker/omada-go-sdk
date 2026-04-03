# SIMTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMailServerTemplate**](SIMTemplateAPI.md#getmailservertemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/mails | get mail server
[**ModifyMailServerTemplate**](SIMTemplateAPI.md#modifymailservertemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/mails/{mailId} | modify mail server
[**ModifySmsPolicySettingTemplate**](SIMTemplateAPI.md#modifysmspolicysettingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/sim/sms/policy | Modify SMS policy setting template
[**ModifySmsRouterCommandTemplate**](SIMTemplateAPI.md#modifysmsroutercommandtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/sim/sms/router-command | Modify SMS router command template
[**QuerySimCardQuotaSettingTemplate**](SIMTemplateAPI.md#querysimcardquotasettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/sim/quota/{simCard} | Get one of Dual-SIM card&#39;s SIM data setting template
[**QuerySmsPolicySettingTemplate**](SIMTemplateAPI.md#querysmspolicysettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/sim/sms/policy | Query SMS policy setting template
[**QuerySmsRouterCommandTemplate**](SIMTemplateAPI.md#querysmsroutercommandtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/sim/sms/router-command | Query SMS router command template



## GetMailServerTemplate

> OperationResponseMailServerOpenApiVO GetMailServerTemplate(ctx, omadacId, siteTemplateId).Execute()

get mail server



### Example

```go
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
	resp, r, err := apiClient.SIMTemplateAPI.GetMailServerTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMTemplateAPI.GetMailServerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailServerTemplate`: OperationResponseMailServerOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SIMTemplateAPI.GetMailServerTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailServerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMailServerOpenApiVO**](OperationResponseMailServerOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMailServerTemplate

> OperationResponseWithoutResult ModifyMailServerTemplate(ctx, omadacId, siteTemplateId, mailId).MailServerOpenApiModifyVO(mailServerOpenApiModifyVO).Execute()

modify mail server



### Example

```go
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
	mailId := "mailId_example" // string | mailId
	mailServerOpenApiModifyVO := *openapiclient.NewMailServerOpenApiModifyVO(false, "Receiver_example", "Sender_example", false) // MailServerOpenApiModifyVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMTemplateAPI.ModifyMailServerTemplate(context.Background(), omadacId, siteTemplateId, mailId).MailServerOpenApiModifyVO(mailServerOpenApiModifyVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMTemplateAPI.ModifyMailServerTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMailServerTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMTemplateAPI.ModifyMailServerTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**mailId** | **string** | mailId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMailServerTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mailServerOpenApiModifyVO** | [**MailServerOpenApiModifyVO**](MailServerOpenApiModifyVO.md) |  | 

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


## ModifySmsPolicySettingTemplate

> OperationResponseWithoutResult ModifySmsPolicySettingTemplate(ctx, omadacId, siteTemplateId).SmaPolicySetting(smaPolicySetting).Execute()

Modify SMS policy setting template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId
	smaPolicySetting := *openapiclient.NewSmaPolicySetting(int32(123)) // SmaPolicySetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMTemplateAPI.ModifySmsPolicySettingTemplate(context.Background(), omadacId, siteTemplateId).SmaPolicySetting(smaPolicySetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMTemplateAPI.ModifySmsPolicySettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySmsPolicySettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMTemplateAPI.ModifySmsPolicySettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | siteTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySmsPolicySettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smaPolicySetting** | [**SmaPolicySetting**](SmaPolicySetting.md) |  | 

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


## ModifySmsRouterCommandTemplate

> OperationResponseWithoutResult ModifySmsRouterCommandTemplate(ctx, omadacId, siteTemplateId).SmsRouterCommand(smsRouterCommand).Execute()

Modify SMS router command template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId
	smsRouterCommand := *openapiclient.NewSmsRouterCommand(false, false, false) // SmsRouterCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMTemplateAPI.ModifySmsRouterCommandTemplate(context.Background(), omadacId, siteTemplateId).SmsRouterCommand(smsRouterCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMTemplateAPI.ModifySmsRouterCommandTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySmsRouterCommandTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMTemplateAPI.ModifySmsRouterCommandTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | siteTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySmsRouterCommandTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smsRouterCommand** | [**SmsRouterCommand**](SmsRouterCommand.md) |  | 

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


## QuerySimCardQuotaSettingTemplate

> OperationResponseSimQuotaSetting QuerySimCardQuotaSettingTemplate(ctx, omadacId, simCard, siteTemplateId).Execute()

Get one of Dual-SIM card's SIM data setting template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMTemplateAPI.QuerySimCardQuotaSettingTemplate(context.Background(), omadacId, simCard, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMTemplateAPI.QuerySimCardQuotaSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySimCardQuotaSettingTemplate`: OperationResponseSimQuotaSetting
	fmt.Fprintf(os.Stdout, "Response from `SIMTemplateAPI.QuerySimCardQuotaSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 
**siteTemplateId** | **string** | siteTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySimCardQuotaSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSimQuotaSetting**](OperationResponseSimQuotaSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySmsPolicySettingTemplate

> OperationResponseSmaPolicySetting QuerySmsPolicySettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Query SMS policy setting template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMTemplateAPI.QuerySmsPolicySettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMTemplateAPI.QuerySmsPolicySettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySmsPolicySettingTemplate`: OperationResponseSmaPolicySetting
	fmt.Fprintf(os.Stdout, "Response from `SIMTemplateAPI.QuerySmsPolicySettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | siteTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySmsPolicySettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSmaPolicySetting**](OperationResponseSmaPolicySetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySmsRouterCommandTemplate

> OperationResponseSmsRouterCommand QuerySmsRouterCommandTemplate(ctx, omadacId, siteTemplateId).Execute()

Query SMS router command template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | siteTemplateId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMTemplateAPI.QuerySmsRouterCommandTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMTemplateAPI.QuerySmsRouterCommandTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySmsRouterCommandTemplate`: OperationResponseSmsRouterCommand
	fmt.Fprintf(os.Stdout, "Response from `SIMTemplateAPI.QuerySmsRouterCommandTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | siteTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySmsRouterCommandTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSmsRouterCommand**](OperationResponseSmsRouterCommand.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


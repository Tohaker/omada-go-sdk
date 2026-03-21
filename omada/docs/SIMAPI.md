# \SIMAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorrectSimQuota**](SIMAPI.md#CorrectSimQuota) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/sim/correct | Correct SIM data
[**CorrectSimQuotaByMac**](SIMAPI.md#CorrectSimQuotaByMac) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/correct | Correct SIM data
[**CreateMailServer**](SIMAPI.md#CreateMailServer) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/mails | Create mail server
[**CreateMailServerByMac**](SIMAPI.md#CreateMailServerByMac) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/mails | Create mail server by mac
[**ExportSmsMessage**](SIMAPI.md#ExportSmsMessage) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/export | Export SMS message
[**ExportSmsMessageBymac**](SIMAPI.md#ExportSmsMessageBymac) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/export | Export SMS message by mac
[**GetGridSimCardSmsInboxMessage**](SIMAPI.md#GetGridSimCardSmsInboxMessage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/inbox/{simCard} | Get one of Dual-SIM card&#39;s SMS inbox message
[**GetGridSimCardSmsInboxMessageByMac**](SIMAPI.md#GetGridSimCardSmsInboxMessageByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/inbox/{simCard} | Get one of Dual-SIM card&#39;s SMS inbox message by mac
[**GetGridSimCardSmsOutboxMessage**](SIMAPI.md#GetGridSimCardSmsOutboxMessage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/outbox/{simCard} | Get one of Dual-SIM card&#39;s SMS outbox message
[**GetGridSimCardSmsOutboxMessageByMac**](SIMAPI.md#GetGridSimCardSmsOutboxMessageByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/outbox/{simCard} | Get one of Dual-SIM card&#39;s SMS outbox message by mac
[**GetGridSmsInboxMessage**](SIMAPI.md#GetGridSmsInboxMessage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/inbox | Get SMS inbox message
[**GetGridSmsInboxMessageByMac**](SIMAPI.md#GetGridSmsInboxMessageByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/inbox | Get SMS inbox message by mac
[**GetGridSmsOutboxMessage**](SIMAPI.md#GetGridSmsOutboxMessage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/outbox | Get SMS outbox message
[**GetGridSmsOutboxMessageByMac**](SIMAPI.md#GetGridSmsOutboxMessageByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/outbox | Get SMS outbox message by mac
[**GetMailServer**](SIMAPI.md#GetMailServer) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/mails | Get mail server
[**GetMailServerByMac**](SIMAPI.md#GetMailServerByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/mails | Get mail server by mac
[**GetSupportSms**](SIMAPI.md#GetSupportSms) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/support | Get support SMS
[**GetSupportSmsByMac**](SIMAPI.md#GetSupportSmsByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/support | Get support SMS
[**ModifyMailServer**](SIMAPI.md#ModifyMailServer) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/mails/{mailId} | Modify mail server
[**ModifySimQuotaSetting**](SIMAPI.md#ModifySimQuotaSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/sim/quota | Modify SIM data setting
[**ModifySimQuotaSettingByMac**](SIMAPI.md#ModifySimQuotaSettingByMac) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/quota | Modify SIM data setting by mac
[**ModifySmsPolicySetting**](SIMAPI.md#ModifySmsPolicySetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/policy | Modify SMS policy setting
[**ModifySmsPolicySettingByMac**](SIMAPI.md#ModifySmsPolicySettingByMac) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/policy | Modify SMS policy setting
[**ModifySmsRouterCommand**](SIMAPI.md#ModifySmsRouterCommand) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/router-command | Modify SMS router command
[**ModifySmsRouterCommandByMac**](SIMAPI.md#ModifySmsRouterCommandByMac) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/router-command | Modify SMS router command by mac
[**OperateSmsMessage**](SIMAPI.md#OperateSmsMessage) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/operate | Operate SMS message
[**OperateSmsMessageByMac**](SIMAPI.md#OperateSmsMessageByMac) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/operate | Operate SMS message by mac
[**QuerySimCardQuotaSetting**](SIMAPI.md#QuerySimCardQuotaSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/quota/{simCard} | Get one of Dual-SIM card&#39;s SIM data setting
[**QuerySimCardQuotaSettingByMac**](SIMAPI.md#QuerySimCardQuotaSettingByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/quota/{simCard} | Get one of Dual-SIM card&#39;s SIM data setting by mac
[**QuerySimQuotaSetting**](SIMAPI.md#QuerySimQuotaSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/quota | Get SIM data setting
[**QuerySimQuotaSettingByMac**](SIMAPI.md#QuerySimQuotaSettingByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/quota | Get SIM data setting by mac
[**QuerySmsPolicySetting**](SIMAPI.md#QuerySmsPolicySetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/policy | Query SMS policy setting
[**QuerySmsPolicySettingByMac**](SIMAPI.md#QuerySmsPolicySettingByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/policy | Query SMS policy setting by mac
[**QuerySmsRouterCommand**](SIMAPI.md#QuerySmsRouterCommand) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/router-command | Query SMS router command
[**QuerySmsRouterCommandByMac**](SIMAPI.md#QuerySmsRouterCommandByMac) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/router-command | Query SMS router command by mac
[**SendSmsMessage**](SIMAPI.md#SendSmsMessage) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/sim/sms/send | Send SMS message
[**SendSmsMessageByMac**](SIMAPI.md#SendSmsMessageByMac) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/gateways/{gatewayMac}/sim/sms/send | Send SMS message



## CorrectSimQuota

> OperationResponseWithoutResult CorrectSimQuota(ctx, omadacId, siteId).CorrectSimQuota(correctSimQuota).Execute()

Correct SIM data



### Example

```go
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
	correctSimQuota := *openapiclient.NewCorrectSimQuota() // CorrectSimQuota | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.CorrectSimQuota(context.Background(), omadacId, siteId).CorrectSimQuota(correctSimQuota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.CorrectSimQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CorrectSimQuota`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.CorrectSimQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorrectSimQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correctSimQuota** | [**CorrectSimQuota**](CorrectSimQuota.md) |  | 

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


## CorrectSimQuotaByMac

> OperationResponseWithoutResult CorrectSimQuotaByMac(ctx, omadacId, siteId, gatewayMac).CorrectSimQuota(correctSimQuota).Execute()

Correct SIM data



### Example

```go
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
	correctSimQuota := *openapiclient.NewCorrectSimQuota() // CorrectSimQuota | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.CorrectSimQuotaByMac(context.Background(), omadacId, siteId, gatewayMac).CorrectSimQuota(correctSimQuota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.CorrectSimQuotaByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CorrectSimQuotaByMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.CorrectSimQuotaByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCorrectSimQuotaByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correctSimQuota** | [**CorrectSimQuota**](CorrectSimQuota.md) |  | 

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


## CreateMailServer

> OperationResponseWithoutResult CreateMailServer(ctx, omadacId, siteId).MailServerOpenApiVO(mailServerOpenApiVO).Execute()

Create mail server



### Example

```go
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
	mailServerOpenApiVO := *openapiclient.NewMailServerOpenApiVO(false, "Receiver_example", "Sender_example", false) // MailServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.CreateMailServer(context.Background(), omadacId, siteId).MailServerOpenApiVO(mailServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.CreateMailServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMailServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.CreateMailServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mailServerOpenApiVO** | [**MailServerOpenApiVO**](MailServerOpenApiVO.md) |  | 

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


## CreateMailServerByMac

> OperationResponseWithoutResult CreateMailServerByMac(ctx, omadacId, siteId, gatewayMac).MailServerOpenApiVO(mailServerOpenApiVO).Execute()

Create mail server by mac



### Example

```go
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
	mailServerOpenApiVO := *openapiclient.NewMailServerOpenApiVO(false, "Receiver_example", "Sender_example", false) // MailServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.CreateMailServerByMac(context.Background(), omadacId, siteId, gatewayMac).MailServerOpenApiVO(mailServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.CreateMailServerByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMailServerByMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.CreateMailServerByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateMailServerByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mailServerOpenApiVO** | [**MailServerOpenApiVO**](MailServerOpenApiVO.md) |  | 

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


## ExportSmsMessage

> OperationResponseWithoutResult ExportSmsMessage(ctx, omadacId, siteId).ExportMessage(exportMessage).Execute()

Export SMS message



### Example

```go
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
	exportMessage := *openapiclient.NewExportMessage(int64(123), int32(123), int64(123), int32(123)) // ExportMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ExportSmsMessage(context.Background(), omadacId, siteId).ExportMessage(exportMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ExportSmsMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSmsMessage`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ExportSmsMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSmsMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exportMessage** | [**ExportMessage**](ExportMessage.md) |  | 

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


## ExportSmsMessageBymac

> OperationResponseWithoutResult ExportSmsMessageBymac(ctx, omadacId, siteId, gatewayMac).ExportMessage(exportMessage).Execute()

Export SMS message by mac



### Example

```go
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
	exportMessage := *openapiclient.NewExportMessage(int64(123), int32(123), int64(123), int32(123)) // ExportMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ExportSmsMessageBymac(context.Background(), omadacId, siteId, gatewayMac).ExportMessage(exportMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ExportSmsMessageBymac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSmsMessageBymac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ExportSmsMessageBymac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiExportSmsMessageBymacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **exportMessage** | [**ExportMessage**](ExportMessage.md) |  | 

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


## GetGridSimCardSmsInboxMessage

> OperationResponseGridVOMessage GetGridSimCardSmsInboxMessage(ctx, omadacId, siteId, simCard).Page(page).PageSize(pageSize).Execute()

Get one of Dual-SIM card's SMS inbox message



### Example

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
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.GetGridSimCardSmsInboxMessage(context.Background(), omadacId, siteId, simCard).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetGridSimCardSmsInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSimCardSmsInboxMessage`: OperationResponseGridVOMessage
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetGridSimCardSmsInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSimCardSmsInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 




### Return type

[**OperationResponseGridVOMessage**](OperationResponseGridVOMessage.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSimCardSmsInboxMessageByMac

> OperationResponseGridVOInboxMessage GetGridSimCardSmsInboxMessageByMac(ctx, omadacId, siteId, gatewayMac, simCard).Page(page).PageSize(pageSize).Execute()

Get one of Dual-SIM card's SMS inbox message by mac



### Example

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
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.GetGridSimCardSmsInboxMessageByMac(context.Background(), omadacId, siteId, gatewayMac, simCard).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetGridSimCardSmsInboxMessageByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSimCardSmsInboxMessageByMac`: OperationResponseGridVOInboxMessage
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetGridSimCardSmsInboxMessageByMac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSimCardSmsInboxMessageByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





### Return type

[**OperationResponseGridVOInboxMessage**](OperationResponseGridVOInboxMessage.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSimCardSmsOutboxMessage

> OperationResponseGridVOMessage GetGridSimCardSmsOutboxMessage(ctx, omadacId, siteId, simCard).Page(page).PageSize(pageSize).Execute()

Get one of Dual-SIM card's SMS outbox message



### Example

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
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.GetGridSimCardSmsOutboxMessage(context.Background(), omadacId, siteId, simCard).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetGridSimCardSmsOutboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSimCardSmsOutboxMessage`: OperationResponseGridVOMessage
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetGridSimCardSmsOutboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSimCardSmsOutboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 




### Return type

[**OperationResponseGridVOMessage**](OperationResponseGridVOMessage.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSimCardSmsOutboxMessageByMac

> OperationResponseOutboxMessageOpenApiGridVOOutboxMessage GetGridSimCardSmsOutboxMessageByMac(ctx, omadacId, siteId, gatewayMac, simCard).Page(page).PageSize(pageSize).Execute()

Get one of Dual-SIM card's SMS outbox message by mac



### Example

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
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.GetGridSimCardSmsOutboxMessageByMac(context.Background(), omadacId, siteId, gatewayMac, simCard).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetGridSimCardSmsOutboxMessageByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSimCardSmsOutboxMessageByMac`: OperationResponseOutboxMessageOpenApiGridVOOutboxMessage
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetGridSimCardSmsOutboxMessageByMac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSimCardSmsOutboxMessageByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 





### Return type

[**OperationResponseOutboxMessageOpenApiGridVOOutboxMessage**](OperationResponseOutboxMessageOpenApiGridVOOutboxMessage.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSmsInboxMessage

> OperationResponseGridVOMessage GetGridSmsInboxMessage(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get SMS inbox message



### Example

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
	resp, r, err := apiClient.SIMAPI.GetGridSmsInboxMessage(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetGridSmsInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSmsInboxMessage`: OperationResponseGridVOMessage
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetGridSmsInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSmsInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOMessage**](OperationResponseGridVOMessage.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSmsInboxMessageByMac

> OperationResponseGridVOInboxMessage GetGridSmsInboxMessageByMac(ctx, omadacId, siteId, gatewayMac).Page(page).PageSize(pageSize).Execute()

Get SMS inbox message by mac



### Example

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
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.GetGridSmsInboxMessageByMac(context.Background(), omadacId, siteId, gatewayMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetGridSmsInboxMessageByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSmsInboxMessageByMac`: OperationResponseGridVOInboxMessage
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetGridSmsInboxMessageByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridSmsInboxMessageByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 




### Return type

[**OperationResponseGridVOInboxMessage**](OperationResponseGridVOInboxMessage.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSmsOutboxMessage

> OperationResponseGridVOMessage GetGridSmsOutboxMessage(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get SMS outbox message



### Example

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
	resp, r, err := apiClient.SIMAPI.GetGridSmsOutboxMessage(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetGridSmsOutboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSmsOutboxMessage`: OperationResponseGridVOMessage
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetGridSmsOutboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridSmsOutboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOMessage**](OperationResponseGridVOMessage.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridSmsOutboxMessageByMac

> OperationResponseOutboxMessageOpenApiGridVOOutboxMessage GetGridSmsOutboxMessageByMac(ctx, omadacId, siteId, gatewayMac).Page(page).PageSize(pageSize).Execute()

Get SMS outbox message by mac



### Example

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
	gatewayMac := "gatewayMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.GetGridSmsOutboxMessageByMac(context.Background(), omadacId, siteId, gatewayMac).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetGridSmsOutboxMessageByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridSmsOutboxMessageByMac`: OperationResponseOutboxMessageOpenApiGridVOOutboxMessage
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetGridSmsOutboxMessageByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetGridSmsOutboxMessageByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 




### Return type

[**OperationResponseOutboxMessageOpenApiGridVOOutboxMessage**](OperationResponseOutboxMessageOpenApiGridVOOutboxMessage.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailServer

> OperationResponseMailServerOpenApiVO GetMailServer(ctx, omadacId, siteId).Execute()

Get mail server



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.GetMailServer(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetMailServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailServer`: OperationResponseMailServerOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetMailServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseMailServerOpenApiVO**](OperationResponseMailServerOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailServerByMac

> OperationResponseMailServerOpenApiVO GetMailServerByMac(ctx, omadacId, siteId, gatewayMac).Execute()

Get mail server by mac



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.GetMailServerByMac(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetMailServerByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailServerByMac`: OperationResponseMailServerOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetMailServerByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMailServerByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseMailServerOpenApiVO**](OperationResponseMailServerOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportSms

> OperationResponseSupportSmsOpenApiVO GetSupportSms(ctx, omadacId, siteId).Execute()

Get support SMS



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.GetSupportSms(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetSupportSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportSms`: OperationResponseSupportSmsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetSupportSms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSupportSmsOpenApiVO**](OperationResponseSupportSmsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportSmsByMac

> OperationResponseSupportSmsOpenApiVO GetSupportSmsByMac(ctx, omadacId, siteId, gatewayMac).Execute()

Get support SMS



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.GetSupportSmsByMac(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.GetSupportSmsByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportSmsByMac`: OperationResponseSupportSmsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.GetSupportSmsByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSupportSmsByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSupportSmsOpenApiVO**](OperationResponseSupportSmsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMailServer

> OperationResponseWithoutResult ModifyMailServer(ctx, omadacId, siteId, mailId).MailServerOpenApiVO(mailServerOpenApiVO).Execute()

Modify mail server



### Example

```go
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
	mailId := "mailId_example" // string | Mail ID
	mailServerOpenApiVO := *openapiclient.NewMailServerOpenApiVO(false, "Receiver_example", "Sender_example", false) // MailServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ModifyMailServer(context.Background(), omadacId, siteId, mailId).MailServerOpenApiVO(mailServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ModifyMailServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMailServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ModifyMailServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mailId** | **string** | Mail ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mailServerOpenApiVO** | [**MailServerOpenApiVO**](MailServerOpenApiVO.md) |  | 

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


## ModifySimQuotaSetting

> OperationResponseWithoutResult ModifySimQuotaSetting(ctx, omadacId, siteId).SimQuotaSetting(simQuotaSetting).Execute()

Modify SIM data setting



### Example

```go
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
	simQuotaSetting := *openapiclient.NewSimQuotaSetting() // SimQuotaSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ModifySimQuotaSetting(context.Background(), omadacId, siteId).SimQuotaSetting(simQuotaSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ModifySimQuotaSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySimQuotaSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ModifySimQuotaSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySimQuotaSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **simQuotaSetting** | [**SimQuotaSetting**](SimQuotaSetting.md) |  | 

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


## ModifySimQuotaSettingByMac

> OperationResponseWithoutResult ModifySimQuotaSettingByMac(ctx, omadacId, siteId, gatewayMac).SimQuotaSetting(simQuotaSetting).Execute()

Modify SIM data setting by mac



### Example

```go
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
	simQuotaSetting := *openapiclient.NewSimQuotaSetting() // SimQuotaSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ModifySimQuotaSettingByMac(context.Background(), omadacId, siteId, gatewayMac).SimQuotaSetting(simQuotaSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ModifySimQuotaSettingByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySimQuotaSettingByMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ModifySimQuotaSettingByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifySimQuotaSettingByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **simQuotaSetting** | [**SimQuotaSetting**](SimQuotaSetting.md) |  | 

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


## ModifySmsPolicySetting

> OperationResponseWithoutResult ModifySmsPolicySetting(ctx, omadacId, siteId).SmaPolicySetting(smaPolicySetting).Execute()

Modify SMS policy setting



### Example

```go
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
	smaPolicySetting := *openapiclient.NewSmaPolicySetting(int32(123)) // SmaPolicySetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ModifySmsPolicySetting(context.Background(), omadacId, siteId).SmaPolicySetting(smaPolicySetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ModifySmsPolicySetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySmsPolicySetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ModifySmsPolicySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySmsPolicySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smaPolicySetting** | [**SmaPolicySetting**](SmaPolicySetting.md) |  | 

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


## ModifySmsPolicySettingByMac

> OperationResponseWithoutResult ModifySmsPolicySettingByMac(ctx, omadacId, siteId, gatewayMac).SmaPolicySetting(smaPolicySetting).Execute()

Modify SMS policy setting



### Example

```go
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
	smaPolicySetting := *openapiclient.NewSmaPolicySetting(int32(123)) // SmaPolicySetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ModifySmsPolicySettingByMac(context.Background(), omadacId, siteId, gatewayMac).SmaPolicySetting(smaPolicySetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ModifySmsPolicySettingByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySmsPolicySettingByMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ModifySmsPolicySettingByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifySmsPolicySettingByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **smaPolicySetting** | [**SmaPolicySetting**](SmaPolicySetting.md) |  | 

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


## ModifySmsRouterCommand

> OperationResponseWithoutResult ModifySmsRouterCommand(ctx, omadacId, siteId).SmsRouterCommand(smsRouterCommand).Execute()

Modify SMS router command



### Example

```go
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
	smsRouterCommand := *openapiclient.NewSmsRouterCommand(false, false, false) // SmsRouterCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ModifySmsRouterCommand(context.Background(), omadacId, siteId).SmsRouterCommand(smsRouterCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ModifySmsRouterCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySmsRouterCommand`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ModifySmsRouterCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySmsRouterCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smsRouterCommand** | [**SmsRouterCommand**](SmsRouterCommand.md) |  | 

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


## ModifySmsRouterCommandByMac

> OperationResponseWithoutResult ModifySmsRouterCommandByMac(ctx, omadacId, siteId, gatewayMac).SmsRouterCommand(smsRouterCommand).Execute()

Modify SMS router command by mac



### Example

```go
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
	smsRouterCommand := *openapiclient.NewSmsRouterCommand(false, false, false) // SmsRouterCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.ModifySmsRouterCommandByMac(context.Background(), omadacId, siteId, gatewayMac).SmsRouterCommand(smsRouterCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.ModifySmsRouterCommandByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySmsRouterCommandByMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.ModifySmsRouterCommandByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifySmsRouterCommandByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **smsRouterCommand** | [**SmsRouterCommand**](SmsRouterCommand.md) |  | 

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


## OperateSmsMessage

> OperationResponseWithoutResult OperateSmsMessage(ctx, omadacId, siteId).OperateMessage(operateMessage).Execute()

Operate SMS message



### Example

```go
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
	operateMessage := *openapiclient.NewOperateMessage(int32(123), int32(123)) // OperateMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.OperateSmsMessage(context.Background(), omadacId, siteId).OperateMessage(operateMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.OperateSmsMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperateSmsMessage`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.OperateSmsMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperateSmsMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **operateMessage** | [**OperateMessage**](OperateMessage.md) |  | 

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


## OperateSmsMessageByMac

> OperationResponseWithoutResult OperateSmsMessageByMac(ctx, omadacId, siteId, gatewayMac).OperateMessage(operateMessage).Execute()

Operate SMS message by mac



### Example

```go
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
	operateMessage := *openapiclient.NewOperateMessage(int32(123), int32(123)) // OperateMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.OperateSmsMessageByMac(context.Background(), omadacId, siteId, gatewayMac).OperateMessage(operateMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.OperateSmsMessageByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperateSmsMessageByMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.OperateSmsMessageByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOperateSmsMessageByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **operateMessage** | [**OperateMessage**](OperateMessage.md) |  | 

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


## QuerySimCardQuotaSetting

> OperationResponseSimQuotaSetting QuerySimCardQuotaSetting(ctx, omadacId, siteId, simCard).Execute()

Get one of Dual-SIM card's SIM data setting



### Example

```go
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
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.QuerySimCardQuotaSetting(context.Background(), omadacId, siteId, simCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.QuerySimCardQuotaSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySimCardQuotaSetting`: OperationResponseSimQuotaSetting
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.QuerySimCardQuotaSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySimCardQuotaSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSimQuotaSetting**](OperationResponseSimQuotaSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySimCardQuotaSettingByMac

> OperationResponseSimQuotaSetting QuerySimCardQuotaSettingByMac(ctx, omadacId, siteId, gatewayMac, simCard).Execute()

Get one of Dual-SIM card's SIM data setting by mac



### Example

```go
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
	simCard := "simCard_example" // string | SIM card. 1: SIM1; 2: SIM2.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.QuerySimCardQuotaSettingByMac(context.Background(), omadacId, siteId, gatewayMac, simCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.QuerySimCardQuotaSettingByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySimCardQuotaSettingByMac`: OperationResponseSimQuotaSetting
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.QuerySimCardQuotaSettingByMac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**gatewayMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**simCard** | **string** | SIM card. 1: SIM1; 2: SIM2. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySimCardQuotaSettingByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseSimQuotaSetting**](OperationResponseSimQuotaSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySimQuotaSetting

> OperationResponseSimQuotaSetting QuerySimQuotaSetting(ctx, omadacId, siteId).Execute()

Get SIM data setting



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.QuerySimQuotaSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.QuerySimQuotaSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySimQuotaSetting`: OperationResponseSimQuotaSetting
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.QuerySimQuotaSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySimQuotaSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSimQuotaSetting**](OperationResponseSimQuotaSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySimQuotaSettingByMac

> OperationResponseSimQuotaSetting QuerySimQuotaSettingByMac(ctx, omadacId, siteId, gatewayMac).Execute()

Get SIM data setting by mac



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.QuerySimQuotaSettingByMac(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.QuerySimQuotaSettingByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySimQuotaSettingByMac`: OperationResponseSimQuotaSetting
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.QuerySimQuotaSettingByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiQuerySimQuotaSettingByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSimQuotaSetting**](OperationResponseSimQuotaSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySmsPolicySetting

> OperationResponseSmaPolicySetting QuerySmsPolicySetting(ctx, omadacId, siteId).Execute()

Query SMS policy setting



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.QuerySmsPolicySetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.QuerySmsPolicySetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySmsPolicySetting`: OperationResponseSmaPolicySetting
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.QuerySmsPolicySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySmsPolicySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSmaPolicySetting**](OperationResponseSmaPolicySetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySmsPolicySettingByMac

> OperationResponseSmaPolicySetting QuerySmsPolicySettingByMac(ctx, omadacId, siteId, gatewayMac).Execute()

Query SMS policy setting by mac



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.QuerySmsPolicySettingByMac(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.QuerySmsPolicySettingByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySmsPolicySettingByMac`: OperationResponseSmaPolicySetting
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.QuerySmsPolicySettingByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiQuerySmsPolicySettingByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSmaPolicySetting**](OperationResponseSmaPolicySetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySmsRouterCommand

> OperationResponseSmsRouterCommand QuerySmsRouterCommand(ctx, omadacId, siteId).Execute()

Query SMS router command



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.QuerySmsRouterCommand(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.QuerySmsRouterCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySmsRouterCommand`: OperationResponseSmsRouterCommand
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.QuerySmsRouterCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySmsRouterCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSmsRouterCommand**](OperationResponseSmsRouterCommand.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySmsRouterCommandByMac

> OperationResponseSmsRouterCommand QuerySmsRouterCommandByMac(ctx, omadacId, siteId, gatewayMac).Execute()

Query SMS router command by mac



### Example

```go
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
	resp, r, err := apiClient.SIMAPI.QuerySmsRouterCommandByMac(context.Background(), omadacId, siteId, gatewayMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.QuerySmsRouterCommandByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySmsRouterCommandByMac`: OperationResponseSmsRouterCommand
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.QuerySmsRouterCommandByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiQuerySmsRouterCommandByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSmsRouterCommand**](OperationResponseSmsRouterCommand.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSmsMessage

> OperationResponseWithoutResult SendSmsMessage(ctx, omadacId, siteId).SendMessage(sendMessage).Execute()

Send SMS message



### Example

```go
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
	sendMessage := *openapiclient.NewSendMessage("Receiver_example", int32(123)) // SendMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.SendSmsMessage(context.Background(), omadacId, siteId).SendMessage(sendMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.SendSmsMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendSmsMessage`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.SendSmsMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendMessage** | [**SendMessage**](SendMessage.md) |  | 

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


## SendSmsMessageByMac

> OperationResponseWithoutResult SendSmsMessageByMac(ctx, omadacId, siteId, gatewayMac).SendMessage(sendMessage).Execute()

Send SMS message



### Example

```go
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
	sendMessage := *openapiclient.NewSendMessage("Receiver_example", int32(123)) // SendMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SIMAPI.SendSmsMessageByMac(context.Background(), omadacId, siteId, gatewayMac).SendMessage(sendMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SIMAPI.SendSmsMessageByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendSmsMessageByMac`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SIMAPI.SendSmsMessageByMac`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSendSmsMessageByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sendMessage** | [**SendMessage**](SendMessage.md) |  | 

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


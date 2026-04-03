# MspSettingAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCertificate**](MspSettingAPI.md#deletecertificate) | **Delete** /openapi/v1/msp/{mspId}/system/setting/certificate/{cerId} | Delete an existing msp certificate
[**DeleteSSLKey**](MspSettingAPI.md#deletesslkey) | **Delete** /openapi/v1/msp/{mspId}/system/setting/ssl-key/{keyId} | Delete an existing msp SSL key
[**DisableMspControllerUpgradeNotification**](MspSettingAPI.md#disablemspcontrollerupgradenotification) | **Post** /openapi/v1/msp/{mspId}/notification/disable | Turn off the software update push switch in MSP mode
[**EvaluateMspFeedback**](MspSettingAPI.md#evaluatemspfeedback) | **Post** /openapi/v1/msp/{mspId}/users/feedback/evaluates | Evaluate MSP feedback
[**GetGeneralSetting**](MspSettingAPI.md#getgeneralsetting) | **Get** /openapi/v1/msp/{mspId}/general-setting | Get msp general setting
[**GetMspDstInfo**](MspSettingAPI.md#getmspdstinfo) | **Get** /openapi/v1/msp/{mspId}/dst-info | Get msp DST
[**GetMspMailServer**](MspSettingAPI.md#getmspmailserver) | **Get** /openapi/v1/msp/{mspId}/settings/mail-server | Get MSP mail server
[**GetMspMailServerStatusGlobal**](MspSettingAPI.md#getmspmailserverstatusglobal) | **Get** /openapi/v1/msp/{mspId}/account/mail-status | Get MSP mailServer status
[**GetMspRadiusProxy**](MspSettingAPI.md#getmspradiusproxy) | **Get** /openapi/v1/msp/{mspId}/settings/network/radius-proxy | Get Msp RADIUS proxy server setting
[**GetMspRadiusServer**](MspSettingAPI.md#getmspradiusserver) | **Get** /openapi/v1/msp/{mspId}/settings/network/radius-server | Get Msp Built-In RADIUS server setting
[**GetMspUserInterface**](MspSettingAPI.md#getmspuserinterface) | **Get** /openapi/v1/msp/{mspId}/settings/user-interface | Get msp user interface
[**GetPrivacyPolicyMsp**](MspSettingAPI.md#getprivacypolicymsp) | **Get** /openapi/v1/msp/{mspId}/privacy-policy | Get msp privacy policy
[**GetRemoteLoggingSetting1**](MspSettingAPI.md#getremoteloggingsetting1) | **Get** /openapi/v1/msp/{mspId}/remote-logging | Get msp remote logging setting
[**GetUiInterface1**](MspSettingAPI.md#getuiinterface1) | **Get** /openapi/v1/msp/{mspId}/ui-interface | Get msp ui interface
[**ModifyGeneralSetting**](MspSettingAPI.md#modifygeneralsetting) | **Patch** /openapi/v1/msp/{mspId}/general-setting | Modify msp general setting
[**ModifyMspMailServer**](MspSettingAPI.md#modifymspmailserver) | **Patch** /openapi/v1/msp/{mspId}/settings/mail-server | Modify MSP mail server
[**ModifyMspRadiusServer**](MspSettingAPI.md#modifymspradiusserver) | **Patch** /openapi/v1/msp/{mspId}/settings/network/radius-server | Modify Msp Built-In RADIUS server setting
[**ModifyMspUserInterface**](MspSettingAPI.md#modifymspuserinterface) | **Patch** /openapi/v1/msp/{mspId}/settings/user-interface | Modify msp user interface
[**ModifyRadiusProxy1**](MspSettingAPI.md#modifyradiusproxy1) | **Patch** /openapi/v1/msp/{mspId}/settings/network/radius-proxy | Modify Msp RADIUS proxy server setting
[**ModifyRemoteLoggingSetting**](MspSettingAPI.md#modifyremoteloggingsetting) | **Patch** /openapi/v1/msp/{mspId}/remote-logging | Modify msp remote logging setting
[**ModifyUiInterface1**](MspSettingAPI.md#modifyuiinterface1) | **Patch** /openapi/v1/msp/{mspId}/ui-interface | Modify msp ui interface
[**SendTestMailForMspV2**](MspSettingAPI.md#sendtestmailformspv2) | **Post** /openapi/v1/msp/{mspId}/settings/test-mail | Send test mail for MSP
[**UploadCertificate**](MspSettingAPI.md#uploadcertificate) | **Post** /openapi/v1/msp/{mspId}/system/setting/certificate | Upload msp certificate
[**UploadSSLKey**](MspSettingAPI.md#uploadsslkey) | **Post** /openapi/v1/msp/{mspId}/system/setting/ssl-key | Upload msp SSL key



## DeleteCertificate

> OperationResponse DeleteCertificate(ctx, mspId, cerId).Execute()

Delete an existing msp certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	cerId := "cerId_example" // string | Certificate ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.DeleteCertificate(context.Background(), mspId, cerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.DeleteCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertificate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.DeleteCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**cerId** | **string** | Certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateRequest struct via the builder pattern


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


## DeleteSSLKey

> OperationResponse DeleteSSLKey(ctx, mspId, keyId).Execute()

Delete an existing msp SSL key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	keyId := "keyId_example" // string | SSL Key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.DeleteSSLKey(context.Background(), mspId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.DeleteSSLKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSSLKey`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.DeleteSSLKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 
**keyId** | **string** | SSL Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSLKeyRequest struct via the builder pattern


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


## DisableMspControllerUpgradeNotification

> OperationResponseWithoutResult DisableMspControllerUpgradeNotification(ctx, mspId).Execute()

Turn off the software update push switch in MSP mode



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.DisableMspControllerUpgradeNotification(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.DisableMspControllerUpgradeNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableMspControllerUpgradeNotification`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.DisableMspControllerUpgradeNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableMspControllerUpgradeNotificationRequest struct via the builder pattern


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


## EvaluateMspFeedback

> OperationResponseWithoutResult EvaluateMspFeedback(ctx, mspId).FeedbackOpenApiVO(feedbackOpenApiVO).Execute()

Evaluate MSP feedback



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	feedbackOpenApiVO := *openapiclient.NewFeedbackOpenApiVO("Content_example", int32(123)) // FeedbackOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.EvaluateMspFeedback(context.Background(), mspId).FeedbackOpenApiVO(feedbackOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.EvaluateMspFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateMspFeedback`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.EvaluateMspFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateMspFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feedbackOpenApiVO** | [**FeedbackOpenApiVO**](FeedbackOpenApiVO.md) |  | 

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


## GetGeneralSetting

> OperationResponseMspGeneralSettingOpenApiVO GetGeneralSetting(ctx, mspId).Execute()

Get msp general setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetGeneralSetting(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetGeneralSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGeneralSetting`: OperationResponseMspGeneralSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetGeneralSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneralSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseMspGeneralSettingOpenApiVO**](OperationResponseMspGeneralSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspDstInfo

> OperationResponseDstOpenApiVO GetMspDstInfo(ctx, mspId).Execute()

Get msp DST



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetMspDstInfo(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetMspDstInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspDstInfo`: OperationResponseDstOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetMspDstInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspDstInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseDstOpenApiVO**](OperationResponseDstOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspMailServer

> OperationResponseSendMailServerOpenApiVO GetMspMailServer(ctx, mspId).Execute()

Get MSP mail server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetMspMailServer(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetMspMailServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspMailServer`: OperationResponseSendMailServerOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetMspMailServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspMailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseSendMailServerOpenApiVO**](OperationResponseSendMailServerOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspMailServerStatusGlobal

> OperationResponse GetMspMailServerStatusGlobal(ctx, mspId).Execute()

Get MSP mailServer status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetMspMailServerStatusGlobal(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetMspMailServerStatusGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspMailServerStatusGlobal`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetMspMailServerStatusGlobal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspMailServerStatusGlobalRequest struct via the builder pattern


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


## GetMspRadiusProxy

> OperationResponseRadiusProxyServerSetting GetMspRadiusProxy(ctx, mspId).Execute()

Get Msp RADIUS proxy server setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetMspRadiusProxy(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetMspRadiusProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspRadiusProxy`: OperationResponseRadiusProxyServerSetting
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetMspRadiusProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspRadiusProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseRadiusProxyServerSetting**](OperationResponseRadiusProxyServerSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspRadiusServer

> OperationResponseBuiltInRadiusServerSetting GetMspRadiusServer(ctx, mspId).Execute()

Get Msp Built-In RADIUS server setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetMspRadiusServer(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetMspRadiusServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspRadiusServer`: OperationResponseBuiltInRadiusServerSetting
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetMspRadiusServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspRadiusServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseBuiltInRadiusServerSetting**](OperationResponseBuiltInRadiusServerSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspUserInterface

> OperationResponseMspUserInterfaceOpenApiVO GetMspUserInterface(ctx, mspId).Execute()

Get msp user interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetMspUserInterface(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetMspUserInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspUserInterface`: OperationResponseMspUserInterfaceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetMspUserInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspUserInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseMspUserInterfaceOpenApiVO**](OperationResponseMspUserInterfaceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivacyPolicyMsp

> OperationResponsePrivacyPolicy GetPrivacyPolicyMsp(ctx, mspId).Execute()

Get msp privacy policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetPrivacyPolicyMsp(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetPrivacyPolicyMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivacyPolicyMsp`: OperationResponsePrivacyPolicy
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetPrivacyPolicyMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivacyPolicyMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponsePrivacyPolicy**](OperationResponsePrivacyPolicy.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteLoggingSetting1

> OperationResponseMspRemoteLoggingSettingOpenApiVO GetRemoteLoggingSetting1(ctx, mspId).Execute()

Get msp remote logging setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetRemoteLoggingSetting1(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetRemoteLoggingSetting1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteLoggingSetting1`: OperationResponseMspRemoteLoggingSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetRemoteLoggingSetting1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteLoggingSetting1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseMspRemoteLoggingSettingOpenApiVO**](OperationResponseMspRemoteLoggingSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUiInterface1

> OperationResponseMspUiInterfaceOpenApiVO GetUiInterface1(ctx, mspId).Execute()

Get msp ui interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.GetUiInterface1(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.GetUiInterface1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUiInterface1`: OperationResponseMspUiInterfaceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.GetUiInterface1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUiInterface1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseMspUiInterfaceOpenApiVO**](OperationResponseMspUiInterfaceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGeneralSetting

> OperationResponseWithoutResult ModifyGeneralSetting(ctx, mspId).ModifyMspGeneralSettingOpenApiVO(modifyMspGeneralSettingOpenApiVO).Execute()

Modify msp general setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	modifyMspGeneralSettingOpenApiVO := *openapiclient.NewModifyMspGeneralSettingOpenApiVO() // ModifyMspGeneralSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.ModifyGeneralSetting(context.Background(), mspId).ModifyMspGeneralSettingOpenApiVO(modifyMspGeneralSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.ModifyGeneralSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGeneralSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.ModifyGeneralSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyGeneralSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyMspGeneralSettingOpenApiVO** | [**ModifyMspGeneralSettingOpenApiVO**](ModifyMspGeneralSettingOpenApiVO.md) |  | 

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


## ModifyMspMailServer

> OperationResponseWithoutResult ModifyMspMailServer(ctx, mspId).ModifyMailServerOpenApiVO(modifyMailServerOpenApiVO).Execute()

Modify MSP mail server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	modifyMailServerOpenApiVO := *openapiclient.NewModifyMailServerOpenApiVO(false, int32(123), "SmtpServer_example", false) // ModifyMailServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.ModifyMspMailServer(context.Background(), mspId).ModifyMailServerOpenApiVO(modifyMailServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.ModifyMspMailServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspMailServer`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.ModifyMspMailServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspMailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyMailServerOpenApiVO** | [**ModifyMailServerOpenApiVO**](ModifyMailServerOpenApiVO.md) |  | 

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


## ModifyMspRadiusServer

> OperationResponse ModifyMspRadiusServer(ctx, mspId).BuiltInRadiusServerSetting(builtInRadiusServerSetting).Execute()

Modify Msp Built-In RADIUS server setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	builtInRadiusServerSetting := *openapiclient.NewBuiltInRadiusServerSetting(int32(123), false, "Secret_example", int32(123), false) // BuiltInRadiusServerSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.ModifyMspRadiusServer(context.Background(), mspId).BuiltInRadiusServerSetting(builtInRadiusServerSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.ModifyMspRadiusServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspRadiusServer`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.ModifyMspRadiusServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspRadiusServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **builtInRadiusServerSetting** | [**BuiltInRadiusServerSetting**](BuiltInRadiusServerSetting.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMspUserInterface

> OperationResponseWithoutResult ModifyMspUserInterface(ctx, mspId).MspUserInterfaceOpenApiVO(mspUserInterfaceOpenApiVO).Execute()

Modify msp user interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	mspUserInterfaceOpenApiVO := *openapiclient.NewMspUserInterfaceOpenApiVO(int32(123)) // MspUserInterfaceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.ModifyMspUserInterface(context.Background(), mspId).MspUserInterfaceOpenApiVO(mspUserInterfaceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.ModifyMspUserInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspUserInterface`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.ModifyMspUserInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspUserInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mspUserInterfaceOpenApiVO** | [**MspUserInterfaceOpenApiVO**](MspUserInterfaceOpenApiVO.md) |  | 

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


## ModifyRadiusProxy1

> OperationResponseRadiusProxyServerSetting ModifyRadiusProxy1(ctx, mspId).RadiusProxyServerSetting(radiusProxyServerSetting).Execute()

Modify Msp RADIUS proxy server setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	radiusProxyServerSetting := *openapiclient.NewRadiusProxyServerSetting(false, int32(123), "Secret_example") // RadiusProxyServerSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.ModifyRadiusProxy1(context.Background(), mspId).RadiusProxyServerSetting(radiusProxyServerSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.ModifyRadiusProxy1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRadiusProxy1`: OperationResponseRadiusProxyServerSetting
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.ModifyRadiusProxy1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRadiusProxy1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radiusProxyServerSetting** | [**RadiusProxyServerSetting**](RadiusProxyServerSetting.md) |  | 

### Return type

[**OperationResponseRadiusProxyServerSetting**](OperationResponseRadiusProxyServerSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRemoteLoggingSetting

> OperationResponseWithoutResult ModifyRemoteLoggingSetting(ctx, mspId).ModifyMspRemoteLoggingSettingOpenApiVO(modifyMspRemoteLoggingSettingOpenApiVO).Execute()

Modify msp remote logging setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	modifyMspRemoteLoggingSettingOpenApiVO := *openapiclient.NewModifyMspRemoteLoggingSettingOpenApiVO(false) // ModifyMspRemoteLoggingSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.ModifyRemoteLoggingSetting(context.Background(), mspId).ModifyMspRemoteLoggingSettingOpenApiVO(modifyMspRemoteLoggingSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.ModifyRemoteLoggingSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRemoteLoggingSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.ModifyRemoteLoggingSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRemoteLoggingSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyMspRemoteLoggingSettingOpenApiVO** | [**ModifyMspRemoteLoggingSettingOpenApiVO**](ModifyMspRemoteLoggingSettingOpenApiVO.md) |  | 

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


## ModifyUiInterface1

> OperationResponseWithoutResult ModifyUiInterface1(ctx, mspId).MspUiInterfaceOpenApiVO(mspUiInterfaceOpenApiVO).Execute()

Modify msp ui interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	mspUiInterfaceOpenApiVO := *openapiclient.NewMspUiInterfaceOpenApiVO(false, false, int32(123), false, false) // MspUiInterfaceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.ModifyUiInterface1(context.Background(), mspId).MspUiInterfaceOpenApiVO(mspUiInterfaceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.ModifyUiInterface1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUiInterface1`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.ModifyUiInterface1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUiInterface1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mspUiInterfaceOpenApiVO** | [**MspUiInterfaceOpenApiVO**](MspUiInterfaceOpenApiVO.md) |  | 

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


## SendTestMailForMspV2

> OperationResponse SendTestMailForMspV2(ctx, mspId).SendMailServerOpenApiVO(sendMailServerOpenApiVO).Execute()

Send test mail for MSP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	sendMailServerOpenApiVO := *openapiclient.NewSendMailServerOpenApiVO(false, int32(123), "SmtpServer_example", false) // SendMailServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.SendTestMailForMspV2(context.Background(), mspId).SendMailServerOpenApiVO(sendMailServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.SendTestMailForMspV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestMailForMspV2`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.SendTestMailForMspV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestMailForMspV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendMailServerOpenApiVO** | [**SendMailServerOpenApiVO**](SendMailServerOpenApiVO.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCertificate

> OperationResponse UploadCertificate(ctx, mspId).CerName(cerName).UploadCertificateRequest(uploadCertificateRequest).Execute()

Upload msp certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	cerName := "cerName_example" // string | 
	uploadCertificateRequest := *openapiclient.NewUploadCertificateRequest("TODO") // UploadCertificateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.UploadCertificate(context.Background(), mspId).CerName(cerName).UploadCertificateRequest(uploadCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.UploadCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCertificate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.UploadCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cerName** | **string** |  | 
 **uploadCertificateRequest** | [**UploadCertificateRequest**](UploadCertificateRequest.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSSLKey

> OperationResponse UploadSSLKey(ctx, mspId).KeyName(keyName).UploadCertificateRequest(uploadCertificateRequest).Execute()

Upload msp SSL key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	mspId := "mspId_example" // string | MSP ID
	keyName := "keyName_example" // string | 
	uploadCertificateRequest := *openapiclient.NewUploadCertificateRequest("TODO") // UploadCertificateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MspSettingAPI.UploadSSLKey(context.Background(), mspId).KeyName(keyName).UploadCertificateRequest(uploadCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MspSettingAPI.UploadSSLKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadSSLKey`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `MspSettingAPI.UploadSSLKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSSLKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyName** | **string** |  | 
 **uploadCertificateRequest** | [**UploadCertificateRequest**](UploadCertificateRequest.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


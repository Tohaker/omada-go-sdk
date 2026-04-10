# ControllerSettingsAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFavoriteSite**](ControllerSettingsAPI.md#cancelfavoritesite) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/favorites/cancels | Cancel favorite site
[**EvaluateFeedback**](ControllerSettingsAPI.md#evaluatefeedback) | **Post** /openapi/v1/{omadacId}/users/feedback/evaluates | Evaluate a customer feedback
[**FavoriteSite**](ControllerSettingsAPI.md#favoritesite) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/favorites | Favorite site
[**GetClientActiveTimeout**](ControllerSettingsAPI.md#getclientactivetimeout) | **Get** /openapi/v1/{omadacId}/controller/setting/active-timeout | Get client active timeout
[**GetClientDetailInformationSetting**](ControllerSettingsAPI.md#getclientdetailinformationsetting) | **Get** /openapi/v1/{omadacId}/client-detail-information | Get client detail information settings
[**GetClientRecognitionEnable**](ControllerSettingsAPI.md#getclientrecognitionenable) | **Get** /openapi/v1/{omadacId}/controller/setting/client-recognition | Get whether client recognition is enabled
[**GetExpImprove**](ControllerSettingsAPI.md#getexpimprove) | **Get** /openapi/v1/{omadacId}/global/controller/setting/exp-improve | Get experience improvement
[**GetGernalSettings**](ControllerSettingsAPI.md#getgernalsettings) | **Get** /openapi/v1/{omadacId}/global/controller/setting/general | Get general settings
[**GetMailServerStatus**](ControllerSettingsAPI.md#getmailserverstatus) | **Get** /openapi/v1/{omadacId}/mail/status | Get mailServer status
[**GetOmadacDstInfo1**](ControllerSettingsAPI.md#getomadacdstinfo1) | **Get** /openapi/v1/{omadacId}/dst-info | Get controller DST
[**GetPrivacyPolicy**](ControllerSettingsAPI.md#getprivacypolicy) | **Get** /openapi/v1/{omadacId}/privacy-policy | Get privacy policy
[**GetRemoteLogging**](ControllerSettingsAPI.md#getremotelogging) | **Get** /openapi/v1/{omadacId}/global/controller/setting/syslog | Get remote logging
[**GetRetention**](ControllerSettingsAPI.md#getretention) | **Get** /openapi/v1/{omadacId}/controller/setting/retention | Get retention
[**GetUiInterface**](ControllerSettingsAPI.md#getuiinterface) | **Get** /openapi/v1/{omadacId}/controller/setting/ui-interface | Get UI interface
[**ModfiyClientRecognitionEnable**](ControllerSettingsAPI.md#modfiyclientrecognitionenable) | **Patch** /openapi/v1/{omadacId}/controller/setting/client-recognition | Modify client recognition enable
[**ModifyClientActiveTimeout**](ControllerSettingsAPI.md#modifyclientactivetimeout) | **Patch** /openapi/v1/{omadacId}/controller/setting/active-timeout | Modify an existing client active timeout
[**ModifyClientDetailInformationSetting**](ControllerSettingsAPI.md#modifyclientdetailinformationsetting) | **Patch** /openapi/v1/{omadacId}/client-detail-information | Modify client detail information settings
[**ModifyExpImprove**](ControllerSettingsAPI.md#modifyexpimprove) | **Patch** /openapi/v1/{omadacId}/global/controller/setting/exp-improve | Modify an existing experience improvement
[**ModifyGeneralSettings**](ControllerSettingsAPI.md#modifygeneralsettings) | **Patch** /openapi/v1/{omadacId}/global/controller/setting/general | Modify an existing general settings
[**ModifyRemoteLogging**](ControllerSettingsAPI.md#modifyremotelogging) | **Patch** /openapi/v1/{omadacId}/global/controller/setting/syslog | Modify an existing remote logging
[**ModifyRetention1**](ControllerSettingsAPI.md#modifyretention1) | **Patch** /openapi/v1/{omadacId}/controller/setting/retention | Modify an existing retention
[**ModifyUiInterface**](ControllerSettingsAPI.md#modifyuiinterface) | **Patch** /openapi/v1/{omadacId}/controller/setting/ui-interface | Modify an existing UI interface
[**SendTestEmail**](ControllerSettingsAPI.md#sendtestemail) | **Post** /openapi/v1/{omadacId}/settings/test-mail | Send test mail for controller



## CancelFavoriteSite

> OperationResponseWithoutResult CancelFavoriteSite(ctx, omadacId, siteId).Execute()

Cancel favorite site



### Example

```go
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
	resp, r, err := apiClient.ControllerSettingsAPI.CancelFavoriteSite(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.CancelFavoriteSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelFavoriteSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.CancelFavoriteSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelFavoriteSiteRequest struct via the builder pattern


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


## EvaluateFeedback

> OperationResponseWithoutResult EvaluateFeedback(ctx, omadacId).FeedbackOpenApiVO(feedbackOpenApiVO).Execute()

Evaluate a customer feedback



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	feedbackOpenApiVO := *openapiclient.NewFeedbackOpenApiVO("Content_example", int32(123)) // FeedbackOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.EvaluateFeedback(context.Background(), omadacId).FeedbackOpenApiVO(feedbackOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.EvaluateFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateFeedback`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.EvaluateFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateFeedbackRequest struct via the builder pattern


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


## FavoriteSite

> OperationResponseWithoutResult FavoriteSite(ctx, omadacId, siteId).Execute()

Favorite site



### Example

```go
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
	resp, r, err := apiClient.ControllerSettingsAPI.FavoriteSite(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.FavoriteSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FavoriteSite`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.FavoriteSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFavoriteSiteRequest struct via the builder pattern


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


## GetClientActiveTimeout

> OperationResponseClientActiveTimeout GetClientActiveTimeout(ctx, omadacId).Execute()

Get client active timeout



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetClientActiveTimeout(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetClientActiveTimeout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientActiveTimeout`: OperationResponseClientActiveTimeout
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetClientActiveTimeout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientActiveTimeoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseClientActiveTimeout**](OperationResponseClientActiveTimeout.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientDetailInformationSetting

> OperationResponseClientDetailInformationSettingVO GetClientDetailInformationSetting(ctx, omadacId).Execute()

Get client detail information settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetClientDetailInformationSetting(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetClientDetailInformationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientDetailInformationSetting`: OperationResponseClientDetailInformationSettingVO
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetClientDetailInformationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientDetailInformationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseClientDetailInformationSettingVO**](OperationResponseClientDetailInformationSettingVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientRecognitionEnable

> OperationResponseClientRecognitionEnableOpenApiVO GetClientRecognitionEnable(ctx, omadacId).Execute()

Get whether client recognition is enabled



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetClientRecognitionEnable(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetClientRecognitionEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRecognitionEnable`: OperationResponseClientRecognitionEnableOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetClientRecognitionEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRecognitionEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseClientRecognitionEnableOpenApiVO**](OperationResponseClientRecognitionEnableOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpImprove

> OperationResponseExperienceImprovementProgram GetExpImprove(ctx, omadacId).Execute()

Get experience improvement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetExpImprove(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetExpImprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpImprove`: OperationResponseExperienceImprovementProgram
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetExpImprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpImproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseExperienceImprovementProgram**](OperationResponseExperienceImprovementProgram.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGernalSettings

> OperationResponseGeneralSetting GetGernalSettings(ctx, omadacId).Execute()

Get general settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetGernalSettings(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetGernalSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGernalSettings`: OperationResponseGeneralSetting
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetGernalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGernalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseGeneralSetting**](OperationResponseGeneralSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailServerStatus

> OperationResponse GetMailServerStatus(ctx, omadacId).Execute()

Get mailServer status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetMailServerStatus(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetMailServerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailServerStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetMailServerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailServerStatusRequest struct via the builder pattern


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


## GetOmadacDstInfo1

> OperationResponseDstOpenApiVO GetOmadacDstInfo1(ctx, omadacId).Execute()

Get controller DST



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetOmadacDstInfo1(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetOmadacDstInfo1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOmadacDstInfo1`: OperationResponseDstOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetOmadacDstInfo1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOmadacDstInfo1Request struct via the builder pattern


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


## GetPrivacyPolicy

> OperationResponsePrivacyPolicy GetPrivacyPolicy(ctx, omadacId).Execute()

Get privacy policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetPrivacyPolicy(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetPrivacyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivacyPolicy`: OperationResponsePrivacyPolicy
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetPrivacyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivacyPolicyRequest struct via the builder pattern


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


## GetRemoteLogging

> OperationResponseRemoteLog GetRemoteLogging(ctx, omadacId).Execute()

Get remote logging



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetRemoteLogging(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetRemoteLogging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteLogging`: OperationResponseRemoteLog
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetRemoteLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseRemoteLog**](OperationResponseRemoteLog.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetention

> OperationResponseHistoryRetention GetRetention(ctx, omadacId).Execute()

Get retention



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetRetention(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRetention`: OperationResponseHistoryRetention
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseHistoryRetention**](OperationResponseHistoryRetention.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUiInterface

> OperationResponseUIInterface GetUiInterface(ctx, omadacId).Execute()

Get UI interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.GetUiInterface(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.GetUiInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUiInterface`: OperationResponseUIInterface
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.GetUiInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUiInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseUIInterface**](OperationResponseUIInterface.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModfiyClientRecognitionEnable

> OperationResponseWithoutResult ModfiyClientRecognitionEnable(ctx, omadacId).ClientRecognitionEnableOpenApiVO(clientRecognitionEnableOpenApiVO).Execute()

Modify client recognition enable



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	clientRecognitionEnableOpenApiVO := *openapiclient.NewClientRecognitionEnableOpenApiVO(false) // ClientRecognitionEnableOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.ModfiyClientRecognitionEnable(context.Background(), omadacId).ClientRecognitionEnableOpenApiVO(clientRecognitionEnableOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.ModfiyClientRecognitionEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModfiyClientRecognitionEnable`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.ModfiyClientRecognitionEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModfiyClientRecognitionEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRecognitionEnableOpenApiVO** | [**ClientRecognitionEnableOpenApiVO**](ClientRecognitionEnableOpenApiVO.md) |  | 

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


## ModifyClientActiveTimeout

> OperationResponse ModifyClientActiveTimeout(ctx, omadacId).ClientActiveTimeout(clientActiveTimeout).Execute()

Modify an existing client active timeout



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	clientActiveTimeout := *openapiclient.NewClientActiveTimeout(int32(123)) // ClientActiveTimeout | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.ModifyClientActiveTimeout(context.Background(), omadacId).ClientActiveTimeout(clientActiveTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.ModifyClientActiveTimeout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyClientActiveTimeout`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.ModifyClientActiveTimeout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyClientActiveTimeoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientActiveTimeout** | [**ClientActiveTimeout**](ClientActiveTimeout.md) |  | 

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


## ModifyClientDetailInformationSetting

> OperationResponseWithoutResult ModifyClientDetailInformationSetting(ctx, omadacId).ClientDetailInformationSettingVO(clientDetailInformationSettingVO).Execute()

Modify client detail information settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	clientDetailInformationSettingVO := *openapiclient.NewClientDetailInformationSettingVO() // ClientDetailInformationSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.ModifyClientDetailInformationSetting(context.Background(), omadacId).ClientDetailInformationSettingVO(clientDetailInformationSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.ModifyClientDetailInformationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyClientDetailInformationSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.ModifyClientDetailInformationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyClientDetailInformationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientDetailInformationSettingVO** | [**ClientDetailInformationSettingVO**](ClientDetailInformationSettingVO.md) |  | 

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


## ModifyExpImprove

> OperationResponse ModifyExpImprove(ctx, omadacId).ExperienceImprovementProgram(experienceImprovementProgram).Execute()

Modify an existing experience improvement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	experienceImprovementProgram := *openapiclient.NewExperienceImprovementProgram(false) // ExperienceImprovementProgram | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.ModifyExpImprove(context.Background(), omadacId).ExperienceImprovementProgram(experienceImprovementProgram).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.ModifyExpImprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyExpImprove`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.ModifyExpImprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyExpImproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experienceImprovementProgram** | [**ExperienceImprovementProgram**](ExperienceImprovementProgram.md) |  | 

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


## ModifyGeneralSettings

> OperationResponse ModifyGeneralSettings(ctx, omadacId).GeneralSetting(generalSetting).Execute()

Modify an existing general settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	generalSetting := *openapiclient.NewGeneralSetting("Name_example", "Region_example", "TimeZone_example") // GeneralSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.ModifyGeneralSettings(context.Background(), omadacId).GeneralSetting(generalSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.ModifyGeneralSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGeneralSettings`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.ModifyGeneralSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyGeneralSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generalSetting** | [**GeneralSetting**](GeneralSetting.md) |  | 

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


## ModifyRemoteLogging

> OperationResponse ModifyRemoteLogging(ctx, omadacId).RemoteLog(remoteLog).Execute()

Modify an existing remote logging



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	remoteLog := *openapiclient.NewRemoteLog(false) // RemoteLog | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.ModifyRemoteLogging(context.Background(), omadacId).RemoteLog(remoteLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.ModifyRemoteLogging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRemoteLogging`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.ModifyRemoteLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRemoteLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteLog** | [**RemoteLog**](RemoteLog.md) |  | 

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


## ModifyRetention1

> OperationResponse ModifyRetention1(ctx, omadacId).HistoryRetention(historyRetention).Execute()

Modify an existing retention



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	historyRetention := *openapiclient.NewHistoryRetention(false) // HistoryRetention | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.ModifyRetention1(context.Background(), omadacId).HistoryRetention(historyRetention).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.ModifyRetention1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRetention1`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.ModifyRetention1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRetention1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **historyRetention** | [**HistoryRetention**](HistoryRetention.md) |  | 

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


## ModifyUiInterface

> OperationResponse ModifyUiInterface(ctx, omadacId).UIInterface(uIInterface).Execute()

Modify an existing UI interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	uIInterface := *openapiclient.NewUIInterface(int32(123)) // UIInterface | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.ModifyUiInterface(context.Background(), omadacId).UIInterface(uIInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.ModifyUiInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyUiInterface`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.ModifyUiInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUiInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uIInterface** | [**UIInterface**](UIInterface.md) |  | 

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


## SendTestEmail

> OperationResponse SendTestEmail(ctx, omadacId).SendMailServerOpenApiVO(sendMailServerOpenApiVO).Execute()

Send test mail for controller



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	sendMailServerOpenApiVO := *openapiclient.NewSendMailServerOpenApiVO(false, int32(123), "SmtpServer_example", false) // SendMailServerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ControllerSettingsAPI.SendTestEmail(context.Background(), omadacId).SendMailServerOpenApiVO(sendMailServerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ControllerSettingsAPI.SendTestEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestEmail`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ControllerSettingsAPI.SendTestEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestEmailRequest struct via the builder pattern


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


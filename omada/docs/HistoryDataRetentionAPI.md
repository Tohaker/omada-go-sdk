# HistoryDataRetentionAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataRetention**](HistoryDataRetentionAPI.md#getdataretention) | **Get** /openapi/v1/{omadacId}/retention | Get history data retention configuration
[**GetMspClientDetailInfoSetting**](HistoryDataRetentionAPI.md#getmspclientdetailinfosetting) | **Get** /openapi/v1/msp/{mspId}/client-detail-information | Get MSP client detail information setting.
[**ModifyMspClientDetailInfoSetting**](HistoryDataRetentionAPI.md#modifymspclientdetailinfosetting) | **Patch** /openapi/v1/msp/{mspId}/client-detail-information | Modify MSP client detail information setting.
[**ModifyRetention**](HistoryDataRetentionAPI.md#modifyretention) | **Patch** /openapi/v1/{omadacId}/retention | Modify history data retention configuration



## GetDataRetention

> OperationResponseHistoryRetention GetDataRetention(ctx, omadacId).Execute()

Get history data retention configuration



### Example

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
	resp, r, err := apiClient.HistoryDataRetentionAPI.GetDataRetention(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryDataRetentionAPI.GetDataRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataRetention`: OperationResponseHistoryRetention
	fmt.Fprintf(os.Stdout, "Response from `HistoryDataRetentionAPI.GetDataRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRetentionRequest struct via the builder pattern


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


## GetMspClientDetailInfoSetting

> OperationResponse GetMspClientDetailInfoSetting(ctx, mspId).Execute()

Get MSP client detail information setting.



### Example

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
	resp, r, err := apiClient.HistoryDataRetentionAPI.GetMspClientDetailInfoSetting(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryDataRetentionAPI.GetMspClientDetailInfoSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspClientDetailInfoSetting`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `HistoryDataRetentionAPI.GetMspClientDetailInfoSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspClientDetailInfoSettingRequest struct via the builder pattern


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


## ModifyMspClientDetailInfoSetting

> OperationResponse ModifyMspClientDetailInfoSetting(ctx, mspId).MspClientDetailInfoSettingVO(mspClientDetailInfoSettingVO).Execute()

Modify MSP client detail information setting.



### Example

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
	mspClientDetailInfoSettingVO := *openapiclient.NewMspClientDetailInfoSettingVO() // MspClientDetailInfoSettingVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryDataRetentionAPI.ModifyMspClientDetailInfoSetting(context.Background(), mspId).MspClientDetailInfoSettingVO(mspClientDetailInfoSettingVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryDataRetentionAPI.ModifyMspClientDetailInfoSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMspClientDetailInfoSetting`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `HistoryDataRetentionAPI.ModifyMspClientDetailInfoSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMspClientDetailInfoSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mspClientDetailInfoSettingVO** | [**MspClientDetailInfoSettingVO**](MspClientDetailInfoSettingVO.md) |  | 

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


## ModifyRetention

> OperationResponseWithoutResult ModifyRetention(ctx, omadacId).ModifyHistoryRetentionOpenApiVO(modifyHistoryRetentionOpenApiVO).Execute()

Modify history data retention configuration



### Example

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
	modifyHistoryRetentionOpenApiVO := *openapiclient.NewModifyHistoryRetentionOpenApiVO(false) // ModifyHistoryRetentionOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryDataRetentionAPI.ModifyRetention(context.Background(), omadacId).ModifyHistoryRetentionOpenApiVO(modifyHistoryRetentionOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryDataRetentionAPI.ModifyRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRetention`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `HistoryDataRetentionAPI.ModifyRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyHistoryRetentionOpenApiVO** | [**ModifyHistoryRetentionOpenApiVO**](ModifyHistoryRetentionOpenApiVO.md) |  | 

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

